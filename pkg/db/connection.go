package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	config "github.com/Akshayvij07/ecommerce/pkg/config"
	domain "github.com/Akshayvij07/ecommerce/pkg/domain"
)

func ConnectDatabase(cfg config.Config) (*gorm.DB, error) {
	fmt.Println(cfg)
	psqlInfo := fmt.Sprintf("host=%s user=%s dbname=%s port=%s password=%s", cfg.DBHost, cfg.DBUser, cfg.DBName, cfg.DBPort, cfg.DBPassword)
	db, dbErr := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{
		SkipDefaultTransaction: true,
	})

	if dbErr != nil {
		return db, dbErr
	}

	db.AutoMigrate(
		&domain.Users{},
		&domain.UserStatus{},
		&domain.Admin{},
		&domain.Address{},
		&domain.UserAddress{},
		&domain.Category{},
		//&category.SubCategory{},
		&domain.Product{},
		&domain.Cart{},
		&domain.CartItem{},
		&domain.Orders{},
		&domain.OrderStatus{},
		&domain.OrderLine{},
		&domain.PaymentDetails{},
		&domain.PaymentMethod{},
		domain.Coupon{},
		domain.WishList{},
	)
	//update triggers
	err := db.Exec(cartTotalPriceUpdate).Error
	if err != nil {
		return db, err
	}
	err = db.Exec(cartTotalPriceUpateTrigger).Error
	if err != nil {
		return db, err
	}

	return db, nil

}
