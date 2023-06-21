package domain

import (
	"time"
)

type Product struct {
	Id           uint   `gorm:"primaryKey;unique;not null"`
	ProductName  string `gorm:"unique;not null"`
	Description  string
	Brand        string
	Prize        int
	Qty_in_stock int
	CategoryId   uint
	//SubCategoryId uint
	//SubCategory   Category
	Category   Category
	Created_at time.Time
	Updated_at time.Time
}

type Category struct {
	Id           uint   `gorm:"primaryKey;unique;not null"`
	CategoryName string `gorm:"unique;not null"`
	Created_at   time.Time
	Updated_at   time.Time
}
