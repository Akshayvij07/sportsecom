package category

import "time"

type SubCategory struct {
	Id           uint `gorm:"primaryKey;unique;not null"`
	CategoryId   uint
	SubCategoryName string `gorm:"unique;not null"`
	Created_at   time.Time
	Updated_at   time.Time
}
