package category

import "time"

type Category struct {
	Id           uint   `gorm:"primaryKey;unique;not null"`
	CategoryName string `gorm:"unique;not null"`
	Created_at   time.Time
	Updated_at   time.Time
}
