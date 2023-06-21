package respondse

import "time"

type UserValue struct {
	ID       uint      `json:"id" gorm:"unique;not null"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
	Created  time.Time `json:"created_time"`
}
type Address struct {
	ID          uint   `json:"id" gorm:"primaryKey;unique"`
	HouseNumber string `json:"house_number"`
	Street      string `json:"street"`
	City        string `json:"city"`
	District    string `json:"district"`
	Pincode     string `json:"pincode"`
	Landmark    string `json:"landmark"`
	IsDefault   bool   `json:"is_default"`
}

type Wishlist struct {
	ProductID   uint   `json:"product_item_id"`
	ProductName string `json:"product_name"`
	Prize       uint   `json:"prize"`
	Image       string `json:"image"`
	QtyInStock  uint   `json:"qty_in_stock"`
}
