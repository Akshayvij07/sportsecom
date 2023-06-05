package domain

import (
	"time"

	"github.com/Akshayvij07/ecommerce/pkg/domain/category"
)



type Product struct {
	Id           uint   `gorm:"primaryKey;unique;not null"`
	ProductName  string `gorm:"unique;not null"`
	Description  string
	Brand        string
	Prize        int
	Qty_in_stock int
	CategoryId   uint
	SubCategoryId uint
	Category     category.Category
	Created_at   time.Time
	Updated_at   time.Time
}
