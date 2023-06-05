package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/Akshayvij07/ecommerce/pkg/domain"
	"github.com/Akshayvij07/ecommerce/pkg/helper/request"
	"github.com/Akshayvij07/ecommerce/pkg/helper/respondse"
	interfaces "github.com/Akshayvij07/ecommerce/pkg/repository/interface"
	"gorm.io/gorm"
)

type AdminDataBase struct {
	DB *gorm.DB
}

func NewAdminRepository(DB *gorm.DB) interfaces.AdminRepository {
	return &AdminDataBase{
		DB: DB,
	}
}

func (c *AdminDataBase) FindAdmin(ctx context.Context, Email string) (domain.Admin, error) {
	var admin domain.Admin
	if c.DB.Raw("SELECT * FROM admins WHERE email=?", Email).Scan(&admin).Error != nil {
		return admin, errors.New("faild to find admin")
	}
	return admin, nil
}

func (c *AdminDataBase) SaveAdmin(ctx context.Context, admin request.Admin) error {
	Query := `INSERT INTO admins(admin_name,email,password)
	VALUES($1, $2, $3) RETURNING *`

	if c.DB.Exec(Query, admin.AdminName, admin.Email, admin.Password).Error != nil {
		return errors.New("failed to create admin")
	}
	return nil
}

func (c *AdminDataBase) FindAllUser(ctx context.Context, pagination request.Pagination) (users []respondse.UserValue, err error) {
	limit := pagination.Page
	offset := (pagination.PerPage - 1) * limit

	query := `SELECT * FROM users ORDER BY name DESC LIMIT $1 OFFSET $2`
	err = c.DB.Raw(query, limit, offset).Scan(&users).Error
	return users, err
}

func (c *AdminDataBase) BlockUser(body request.BlockUser, AdminId int) error {
	//start transaction
	tx := c.DB.Begin()
	//checking if the user exist
	var exist bool
	if err := tx.Raw("SELECT EXISTS(SELECT 1 FROM users WHERE  id= $1)", body.UserID).Scan(&exist).Error; err != nil {
		tx.Rollback()
		return err
	}
	if !exist {
		tx.Rollback()
		return fmt.Errorf("no Such User")
	}

	//Execute the first Sql command (UPDATE)
	if err := tx.Exec("UPDATE users SET is_blocked = true WHERE id =?", body.UserID).Error; err != nil {
		tx.Rollback()
		return err
	}
	//commit the transaction
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}
	//if all commands were executed successfully ,return nil
	return nil
}

func (c *AdminDataBase) UnBlockUser(id int) error {
	tx := c.DB.Begin()
	var exist bool

	if err := tx.Raw("SELECT EXISTS(SELECT 1 FROM users WHERE id=$1 AND is_blocked=true)", id).Scan(&exist).Error; err != nil {
		tx.Rollback()
		return err
	}
	if !exist {
		tx.Rollback()
		return fmt.Errorf("no such user to unblock")

	}
	if err := tx.Exec("UPDATE users SET is_blocked = false WHERE id=$1", id).Error; err != nil {
		tx.Rollback()
		return err
	}
	query := "UPDATE user_statuses SET reason_for_blocking=$1,blocked_at=NULL,blocked_by=$2 WHERE users_id=$3"
	if err := tx.Exec(query, "", 0, id).Error; err != nil {
		tx.Rollback()
		return err
	}
	return nil
}
func (c *AdminDataBase) FindUserbyId(ctx context.Context, userId int) (domain.Users, error) {
	var user domain.Users

	FindbyId := `SELECT *FROM users WHERE id=$1;`

	err := c.DB.Raw(FindbyId, userId).Scan(&user).Error
	if user.ID == 0 {
		return domain.Users{}, fmt.Errorf("user Not Found")
	}
	return user, err
}
