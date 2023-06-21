package usecase

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/Akshayvij07/ecommerce/pkg/domain"
	"github.com/Akshayvij07/ecommerce/pkg/helper/request"
	"github.com/Akshayvij07/ecommerce/pkg/helper/respondse"
	interfaces "github.com/Akshayvij07/ecommerce/pkg/repository/interface"
	services "github.com/Akshayvij07/ecommerce/pkg/usecase/interface"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"

	"github.com/jinzhu/copier"
)

type userUseCase struct {
	userRepo interfaces.UserRepository
}

func NewUserUseCase(repo interfaces.UserRepository) services.UserUseCase {
	return &userUseCase{
		userRepo: repo,
	}
}

func (c *userUseCase) SignUp(ctx context.Context, user request.UserSign) (respondse.UserValue, error) {
	//check user already registered or not\
	if user.Password != user.ConfirmPassword {
		return respondse.UserValue{}, fmt.Errorf("The password is not Matching ")
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return respondse.UserValue{}, err
	}
	user.Password = string(hash)
	userData, err := c.userRepo.UserSignup(ctx, user)
	return userData, err

}

func (c *userUseCase) Login(ctx context.Context, user request.Login) (string, error) {
	userData, err := c.userRepo.UserLogin(ctx, user.Email)
	//var userStatus domain.Users
	if err != nil {
		return "test1", err
	} else if userData.ID == 0 {
		return "test2", fmt.Errorf("user not found")
	}

	if userData.Email == "" {
		return "", fmt.Errorf("user not found")
	}

	fmt.Println("db", userData.Password, "user", user.Password)

	err = bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(user.Password))
	if err != nil {
		return "test 4", fmt.Errorf("incorrect Password")
	}

	fmt.Println("user_id on jwt generate ", userData.ID)
	claims := jwt.MapClaims{
		"id":  userData.ID,
		"exp": time.Now().Add(time.Hour * 48).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	cc, err := token.SignedString([]byte("SECRET"))
	if err != nil {
		return "test 6", err
	}

	return cc, nil
}

func (c *userUseCase) OtpLogin(mobno string) (string, error) {
	id, err := c.userRepo.OtpLogin(mobno)
	if err != nil {
		return "", err
	} else if id == 0 {
		return "", errors.New("user not exist with given mobile number")
	}

	fmt.Println("user_id on otp_login", id)
	claims := jwt.MapClaims{
		"id":  id,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte("SECRET"))
	if err != nil {
		return "", err
	}

	return ss, nil
}

func (c *userUseCase) ChangePassword(ctx context.Context, User request.Password) error {
	userData, err := c.userRepo.FindUser(ctx, User.UserID)
	if err != nil {
		return err
	} else if userData.ID == 0 {
		return fmt.Errorf("user not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(User.OldPassword))
	if err != nil {
		return fmt.Errorf("old password is incorrect")
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(User.NewPassword), 10)
	if err != nil {
		return fmt.Errorf("password is not hashed")
	}
	User.NewPassword = string(hash)
	err = c.userRepo.UpdatePassword(ctx, User.UserID, User.NewPassword)
	return err

}

func (c *userUseCase) AddAdress(ctx context.Context, UserID int, AddressReq request.AddressReq, isDefault bool) error {
	address, err := c.userRepo.FindAddressByUserDetails(ctx, AddressReq, UserID)
	if err != nil {
		return err
	} else if address.ID != 0 { // user have already this address exist
		return errors.New("user have already this address exist with same details")
	}
	user, err := c.userRepo.AddAdress(ctx, UserID, AddressReq)
	fmt.Println(user.ID)

	//creating a user address with this given value
	var userAdress = domain.UserAddress{
		UsersID:   uint(UserID),
		AddressID: user.ID,
		IsDefault: isDefault,
	}

	// then update the address with user
	err = c.userRepo.SaveUserAddress(ctx, userAdress)

	if err != nil {
		return err
	}
	log.Printf("successfully user address stored for user with user_id %v", UserID)
	return nil
}

func (c *userUseCase) UpdateAdress(ctx context.Context, UserID int, AddressReq request.AddressReq) error {
	// first validate the addessId is valid or not
	address, err := c.userRepo.FindUserAddressById(ctx, UserID)
	fmt.Println(address)
	if err != nil {
		return err
	} else if address.ID == 0 {
		return errors.New("invalid address id")

	}
	copier.Copy(&address, &AddressReq)
	if err := c.userRepo.UpdateAddress(ctx, address); err != nil {
		return err
	}
	if AddressReq.IsDefault != nil && *AddressReq.IsDefault {
		userAddress := domain.UserAddress{
			UsersID:   uint(UserID),
			AddressID: address.ID,
			IsDefault: *AddressReq.IsDefault,
		}

		err := c.userRepo.UpdateUserAdress(ctx, userAddress)
		if err != nil {
			return err
		}
	}
	log.Printf("successfully address saved for user with user_id %v", UserID)
	return nil

}

func (c *userUseCase) VeiwAdress(ctx context.Context, UserID int) ([]respondse.Address, error) {
	adress, err := c.userRepo.VeiwAdress(ctx, UserID)
	return adress, err
}
func (c *userUseCase) AddToWishList(ctx context.Context, wishList domain.WishList) error {

	product, err := c.userRepo.FindProduct(ctx, wishList.ProductID)
	if err != nil {
		return err
	} else if product.Id == 0 {
		return errors.New("invalid product_id")
	}

	checkWishList, err := c.userRepo.FindWishListItem(ctx, wishList.ProductID, wishList.UserID)
	if err != nil {
		return err
	} else if checkWishList.ID != 0 {
		return errors.New("product is  already exist on wishlist")
	}

	if err := c.userRepo.SaveWishListItem(ctx, wishList); err != nil {
		return err
	}

	return nil
}

func (c *userUseCase) RemoveFromWishList(ctx context.Context, wishList domain.WishList) error {

	product, err := c.userRepo.FindProduct(ctx, wishList.ProductID)
	if err != nil {
		return err
	} else if product.Id == 0 {
		return errors.New("invalid product_id")
	}

	// check the productItem already exist on wishlist for user
	wishList, err = c.userRepo.FindWishListItem(ctx, wishList.ProductID, wishList.UserID)
	if err != nil {
		return err
	} else if wishList.ID == 0 {
		return errors.New("productItem not found in wishlist")
	}

	return c.userRepo.RemoveWishListItem(ctx, wishList)
}

func (c *userUseCase) ListWishlist(ctx context.Context, userID uint) ([]respondse.Wishlist, error) {
	return c.userRepo.FindAllWishListItemsByUserID(ctx, userID)
}
