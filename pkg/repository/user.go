package repository

import (
	"context"

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

func (c *userDatabase) AddAdress(ctx context.Context, UserID int, address request.AddressReq) (domain.Address, error) {
	var existaddress, newAddress domain.Address

	findaddressbyUser := "SELECT *FROM addresses WHERE user_id=?"

	err := c.DB.Raw(findaddressbyUser, UserID).Scan(&existaddress).Error
	if err != nil {
		return domain.Address{}, err
	}

	if existaddress.ID == 0 || existaddress.UserID == 0 {
		AddAddressQuery := `	INSERT INTO addresses(
			user_id, house_number, street, city, district, pincode, landmark) 
			VALUES($1,$2,$3,$4,$5,$6, $7) RETURNING *`

		err := c.DB.Raw(AddAddressQuery, UserID, address.HouseNumber, address.Street, address.City, address.District, address.Pincode, address.Landmark).Scan(&newAddress).Error
		return newAddress, err
	} else {
		//	address is already there, Edit it
		EditAddressQuery := `	UPDATE addresses SET
								house_number = $1, street = $2, city = $3, district = $4, pincode = $5, landmark = $6
								WHERE user_id = $7
								RETURNING *`
		err := c.DB.Raw(EditAddressQuery, address.HouseNumber, address.Street, address.City, address.District, address.Pincode, address.Landmark, UserID).Scan(&newAddress).Error

		return newAddress, err
	}
}

func (c *userDatabase) UpdateAdress(ctx context.Context, UserID int, address request.AddressReq) (domain.Address, error) {
	var updated domain.Address

	updateQuery := `UPDATE addresses SET
								house_number = $1, street = $2, city = $3, district = $4, pincode = $5, landmark = $6
								WHERE user_id = $7
								RETURNING *`
	err := c.DB.Raw(updateQuery, address.HouseNumber, address.Street, address.City, address.District, address.Pincode, address.Landmark, UserID).Scan(&updated).Error
	return updated, err

}

func (c *userDatabase) VeiwAdress(ctx context.Context, UserID int) (domain.Address, error) {
	var Adress domain.Address

	quary := `SELECT *FROM addresses WHERE user_id=?`

	err := c.DB.Raw(quary, UserID).Scan(&Adress).Error
	return Adress, err
}

func (c *userDatabase) UpdatePassword(ctx context.Context, UserID int, Password string) error {
	var User domain.Users

	updateQuery := `UPDATE users SET
								password = $1
								RETURNING *`
	err := c.DB.Raw(updateQuery, Password).Scan(&User).Error
	return err

}
