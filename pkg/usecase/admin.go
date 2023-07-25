package usecase

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/Akshayvij07/ecommerce/pkg/domain"
	"github.com/Akshayvij07/ecommerce/pkg/helper/request"
	"github.com/Akshayvij07/ecommerce/pkg/helper/respondse"
	interfaces "github.com/Akshayvij07/ecommerce/pkg/repository/interface"
	services "github.com/Akshayvij07/ecommerce/pkg/usecase/interface"
	"github.com/golang-jwt/jwt"
	"github.com/jinzhu/copier"
	"golang.org/x/crypto/bcrypt"
)

type AdminUsecase struct {
	AdminRepo interfaces.AdminRepository
}

func NewAdminUseCase(repo interfaces.AdminRepository) services.AdminUsecase {
	return &AdminUsecase{
		AdminRepo: repo,
	}
}

func (c *AdminUsecase) SaveAdmin(ctx context.Context, admin request.Admin) error {
	if admin, err := c.AdminRepo.FindAdmin(ctx, admin.Email); err != nil {
		return err
	} else if admin.ID != 0 {
		return errors.New("Already exist with the same details")
	}
	//generate password

	hashP, err := bcrypt.GenerateFromPassword([]byte(admin.Password), 10)
	if err != nil {
		return errors.New("failed to generate hashed password for admin")
	}
	admin.Password = string(hashP)
	return c.AdminRepo.SaveAdmin(ctx, admin)
}

func (c *AdminUsecase) LoginAdmin(ctx context.Context, admin request.AdminLogin) (string, error) {
	fmt.Println(admin)
	DbAdmin, err := c.AdminRepo.FindAdmin(ctx, admin.Email)

	// if err != nil {
	// 	return "", errors.New("cant find admin")

	// } else if DbAdmin.Email == "" {
	// 	return "", errors.New("This user not found")
	// }
	if err != nil {
		return "test1", err
	} else if DbAdmin.ID == 0 {
		return "test2", fmt.Errorf("user not found")
	}

	if DbAdmin.Email == "" {
		return "", fmt.Errorf("user not found")
	}

	if bcrypt.CompareHashAndPassword([]byte(DbAdmin.Password), []byte(admin.Password)) != nil {
		return "", errors.New("incorrect password")

	}

	claims := &jwt.MapClaims{
		"id":  DbAdmin.ID,
		"exp": time.Now().Add(time.Hour * 48).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte("SECRET"))

	if err != nil {
		return "", err
	}
	return ss, nil
}

func (c *AdminUsecase) FindAllUser(ctx context.Context, pagination request.Pagination) (users []respondse.UserValue, err error) {

	users, err = c.AdminRepo.FindAllUser(ctx, pagination)

	if err != nil {
		return nil, err
	}
	var respond []respondse.UserValue
	copier.Copy(&respond, &users)
	return respond, nil
}

func (c *AdminUsecase) BlockUser(body request.BlockUser, adminId int) error {
	err := c.AdminRepo.BlockUser(body, adminId)
	return err
}

func (c *AdminUsecase) UnBlockUser(id int) error {
	err := c.AdminRepo.UnBlockUser(id)
	return err
}
func (c *AdminUsecase) FindUserbyId(ctx context.Context, userID int) (domain.Users, error) {
	user, err := c.AdminRepo.FindUserbyId(ctx, userID)
	return user, err
}
func (c *AdminUsecase) ViewSalesReport(ctx context.Context) ([]respondse.SalesReport, error) {
	report, err := c.AdminRepo.ViewSalesReport(ctx)
	return report, err
}
