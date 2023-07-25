package repository

import (
	"context"
	"errors"
	"fmt"
	"time"

	domain "github.com/Akshayvij07/ecommerce/pkg/domain"
	"github.com/Akshayvij07/ecommerce/pkg/helper/request"
	"github.com/Akshayvij07/ecommerce/pkg/helper/respondse"
	interfaces "github.com/Akshayvij07/ecommerce/pkg/repository/interface"
	"gorm.io/gorm"
)

type userDatabase struct {
	DB *gorm.DB
}

func NewUserRepository(DBase *gorm.DB) interfaces.UserRepository {
	return &userDatabase{
		DB: DBase,
	}
}

func (c *userDatabase) UserSignup(ctx context.Context, user request.UserSign) (userData respondse.UserValue, err error) {
	insertQ := `INSERT INTO users (name,email,mobile,password)VALUES($1,$2,$3,$4)
					RETURNING 	id,name,email,mobile`
	err = c.DB.Raw(insertQ, user.Name, user.Email, user.Phone, user.Password).Scan(&userData).Error
	return userData, err
}

func (c *userDatabase) UserLogin(ctx context.Context, Email string) (domain.Users, error) {
	var usersData domain.Users
	err := c.DB.Raw("SELECT * FROM users WHERE email=?", Email).Scan(&usersData).Error

	return usersData, err
}

func (c *userDatabase) FindUser(ctx context.Context, UsersId int) (domain.Users, error) {
	var usersData domain.Users
	err := c.DB.Raw("SELECT * FROM users WHERE id =?", UsersId).Scan(&usersData).Error

	return usersData, err
}

func (c *userDatabase) OtpLogin(mbnum string) (int, error) {
	var id int
	query := "SELECT id FROM users WHERE mobile=?"
	err := c.DB.Raw(query, mbnum).Scan(&id).Error
	return id, err
}
func (c *userDatabase) FindAddressByUserDetails(ctx context.Context, AddressReq request.AddressReq, userID int) (domain.Address, error) {
	var address domain.Address
	query := `SELECT a.id, a.house_number,a.street,a.city,a.district,a.pincode,a.landmark,ua.is_default
	 FROM user_addresses ua JOIN addresses a ON ua.address_id=a.id AND house_number=$1 
	 AND street=$2 AND city=$3 AND district=$4 AND pincode=$5 AND landmark=$6
	 WHERE ua.users_id=$7`
	// find the address with house,land_mark,pincode,coutry_id
	/*query := `SELECT * FROM addresses a JOIN user_addresses ua
	ON a.id = ua.address_id AND users_id = ? `AND house_number=?
	AND street=? AND city=? AND district=? AND pincode=? AND landmark=?*/
	if c.DB.Raw(query, AddressReq.HouseNumber, AddressReq.Street, AddressReq.City, AddressReq.District, AddressReq.Pincode, AddressReq.Landmark, uint(userID)).Scan(&address).Error != nil {
		return address, errors.New("faild to find the address")
	}
	return address, nil
	/*query := `SELECT a.id, a.house_number,a.street,a.city,a.district,a.pincode,a.landmark,ua.is_default
	 FROM user_addresses ua JOIN addresses a ON ua.address_id=a.id
	 WHERE ua.users_id=?`
	if c.DB.Raw(query, UserID).Scan(&addresses).Error != nil {
		return address, errors.New("faild to get address of user")
	}

	return address, nil*/
}
func (c *userDatabase) FindUserAddressById(ctx context.Context, userId int) (domain.Address, error) {
	var address domain.Address
	query := `SELECT a.id, a.house_number,a.street,a.city,a.district,a.pincode,a.landmark,ua.is_default
	 FROM user_addresses ua JOIN addresses a ON ua.address_id=a.id
	 WHERE ua.users_id=?`
	if c.DB.Raw(query, uint(userId)).Scan(&address).Error != nil {
		return address, errors.New("faild to find the address")
	}
	return address, nil
}

func (c *userDatabase) AddAdress(ctx context.Context, UserID int, AddressReq request.AddressReq) (Address domain.Address, err error) {
	var address domain.Address
	query := `INSERT INTO addresses (house_number,street,city,district,pincode,landmark,created_at) 
	VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`

	createdAt := time.Now()

	if c.DB.Raw(query,
		AddressReq.HouseNumber, AddressReq.Street, AddressReq.City, AddressReq.District,
		AddressReq.Pincode, AddressReq.Landmark, createdAt,
	).Scan(&address).Error != nil {
		return address, errors.New("faild to insert address on database")
	}
	return address, nil
}

func (c *userDatabase) SaveUserAddress(ctx context.Context, userAddress domain.UserAddress) error {

	// first check user's first address is this or not
	var userID uint
	query := `SELECT address_id FROM user_addresses WHERE users_id = $1`
	err := c.DB.Raw(query, userAddress.UsersID).Scan(&userID).Error
	if err != nil {
		return fmt.Errorf("faild to check user have already address exit or not with user_id %v", userAddress.UsersID)
	}

	// if the given address is need to set default  then remove all other from default
	if userID == 0 { // it means user have no other addresses
		userAddress.IsDefault = true
	} else if userAddress.IsDefault {
		query := `UPDATE user_addresses SET is_default = 'f' WHERE users_id = ?`
		if c.DB.Raw(query, userAddress.UsersID).Scan(&userAddress).Error != nil {
			return errors.New("faild to remove default status of address")
		}
	}

	query = `INSERT INTO user_addresses (users_id,address_id,is_default) VALUES ($1, $2, $3)`
	err = c.DB.Exec(query, userAddress.UsersID, userAddress.AddressID, userAddress.IsDefault).Error
	if err != nil {
		return errors.New("faild to inser userAddress on database")
	}
	return nil
}

func (c *userDatabase) UpdateUserAdress(ctx context.Context, userAddress domain.UserAddress) error {
	//	var updated domain.UserAddress

	// updateQuery := `UPDATE addresses SET
	// 							house_number = $1, street = $2, city = $3, district = $4, pincode = $5, landmark = $6
	// 							WHERE user_id = $7
	// 							RETURNING *`
	// err := c.DB.Raw(updateQuery, address.HouseNumber, address.Street, address.City, address.District, address.Pincode, address.Landmark, UserID).Scan(&updated).Error
	// return updated, err
	// if it need to set default the change the old default
	if userAddress.IsDefault {

		query := `UPDATE user_addresses SET is_default = 'f' WHERE users_id = ?`
		if c.DB.Raw(query, userAddress.UsersID).Scan(&userAddress).Error != nil {
			return errors.New("faild to remove default status of address")
		}
	}

	// update the user address
	query := `UPDATE user_addresses SET is_default = ? WHERE address_id=? AND users_id=?`
	if c.DB.Raw(query, userAddress.IsDefault, userAddress.AddressID, userAddress.UsersID).Scan(&userAddress).Error != nil {
		return errors.New("faild to update user address")
	}
	return nil

}

func (c *userDatabase) UpdateAddress(ctx context.Context, address domain.Address) error {

	query := `UPDATE addresses SET house_number=$1, street=$2, city=$3, 
	district=$4, pincode=$5,landmark=$6, updated_at = $7 WHERE id=$8`

	updatedAt := time.Now()
	if c.DB.Raw(query, address.HouseNumber, address.Street, address.City,
		address.District, address.Pincode,
		address.Landmark, updatedAt, address.ID).Scan(&address).Error != nil {
		return errors.New("faild to update the address for edit address")
	}
	return nil
}

func (c *userDatabase) VeiwAdress(ctx context.Context, UserID int) ([]respondse.Address, error) {
	//var Address []domain.Address

	var addresses []respondse.Address

	query := `SELECT a.id, a.house_number,a.street,a.city,a.district,a.pincode,a.landmark,ua.is_default
	 FROM user_addresses ua JOIN addresses a ON ua.address_id=a.id 
	 WHERE ua.users_id=?`
	if c.DB.Raw(query, UserID).Scan(&addresses).Error != nil {
		return addresses, errors.New("faild to get address of user")
	}

	return addresses, nil
}

func (c *userDatabase) UpdatePassword(ctx context.Context, UserID int, Password string) error {
	var User domain.Users

	updateQuery := `UPDATE users SET
								password = $1
								RETURNING *`
	err := c.DB.Raw(updateQuery, Password).Scan(&User).Error
	return err

}

func (c *userDatabase) FindWishListItem(ctx context.Context, productID, userID uint) (domain.WishList, error) {

	var wishList domain.WishList
	query := `SELECT * FROM wish_lists WHERE user_id=? AND product_id=?`
	if c.DB.Raw(query, userID, productID).Scan(&wishList).Error != nil {
		return wishList, errors.New("faild to find wishlist item")
	}
	return wishList, nil
}

func (c *userDatabase) FindAllWishListItemsByUserID(ctx context.Context, userID uint) ([]respondse.Wishlist, error) {

	var wishLists []respondse.Wishlist

	favourite := ` SELECT *
	FROM products p
	JOIN wish_lists w ON w.product_id = p.id
	WHERE w.user_id = ?`

	if c.DB.Raw(favourite, userID).Scan(&wishLists).Error != nil {
		return wishLists, errors.New("faild to get wish_list items")
	}
	return wishLists, nil
}

func (c *userDatabase) SaveWishListItem(ctx context.Context, wishList domain.WishList) error {

	query := `INSERT INTO wish_lists (user_id, product_id) VALUES ($1, $2)`

	if c.DB.Raw(query, wishList.UserID, wishList.ProductID).Scan(&wishList).Error != nil {
		return errors.New("faild to insert a product into whishlist")
	}
	return nil
}

func (c *userDatabase) RemoveWishListItem(ctx context.Context, wishList domain.WishList) error {

	query := `DELETE FROM wish_lists WHERE id=?`
	if c.DB.Raw(query, wishList.ID).Scan(&wishList).Error != nil {
		return errors.New("faild to delete product")
	}
	return nil
}

func (c *userDatabase) FindProduct(ctx context.Context, id uint) (respondse.Product, error) {
	var product respondse.Product
	query := `SELECT p.id,p.product_name as name,p.description,p.brand,p.prize,p.category_id,p.qty_in_stock,c.category_name,p.created_at,p.updated_at FROM products p 
		JOIN categories c ON p.category_id=c.id WHERE p.id=$1`
	err := c.DB.Raw(query, id).Scan(&product).Error
	return product, err
}

func (c *userDatabase) GetInvoice(ctx context.Context, UserId int) (respondse.Invoice, error) {
	var Invoice respondse.Invoice
	query := ` SELECT o.users_id,u.name,o.order_date,o.payment_method_id,pm.payment_method,o.shipping_address_id,
	o.discount,o.order_total,o.order_status_id,os.order_status
	FROM orders o
	JOIN users u ON o.users_id = u.id
	JOIN payment_methods pm ON o.payment_method_id = pm.id
	JOIN user_addresses ua ON o.shipping_address_id = ua.id  
	JOIN order_statuses os ON o.order_status_id = os.id
	

	WHERE o.users_id = $1 `
	err := c.DB.Raw(query, UserId).Scan(&Invoice).Error
	return Invoice, err
}
