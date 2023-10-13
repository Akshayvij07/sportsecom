package domain

import "time"

type Users struct {
	ID        uint   `gorm:"primaryKey;unique;not null"`
	Name      string `json:"name" binding:"required"`
	Email     string `json:"email" binding:"required,email" gorm:"unique;not null"`
	Mobile    string `json:"mobile" binding:"required,eq=10" gorm:"unique; not null"`
	Password  string `json:"password" gorm:"not null"`
	IsBlocked bool   `gorm:"default:false"`
	CreatedAt time.Time
}

type UserStatus struct {
	ID                uint `gorm:"primaryKey"`
	UsersID           uint
	Users             Users
	BlockedAt         time.Time
	BlockedBy         uint
	ReasonForBlocking string
}

/*type Cart struct {
	Id          uint    `gorm:"primaryKey;unique;not null"`
	UsersID     uint    `json:"users_id"`
	Users       Users   `gorm:"foreignKey:users_id"`
	Total_price float64 `json:"total_price" gorm:"not null"`
}

type CartItem struct {
	Id        uint `json:"id" gorm:"primaryKey;not null"`
	CartId    uint `json:"cart_id"`
	Cart      Cart
	ProductId uint    `json:"product_id" gorm:"not null"`
	Product   Product `json:"-"`
	Qty       uint    `json:"qty" gorm:"not null"`
	prize     int
}*/

type UserAddress struct {
	ID        uint `json:"id" gorm:"primaryKey;unique"`
	UsersID   uint `json:"users_id" gorm:"not null"`
	Users     Users
	AddressID uint `json:"address_id" gorm:"not null"`
	Address   Address
	IsDefault bool `json:"is_default"`
}

type Address struct {
	ID          uint      `json:"id" gorm:"primaryKey;unique"`
	HouseNumber string    `json:"house_number"`
	Street      string    `json:"street"`
	City        string    `json:"city"`
	District    string    `json:"district"`
	Pincode     string    `json:"pincode"`
	Landmark    string    `json:"landmark"`
	CreatedAt   time.Time `json:"created_at" gorm:"not null"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Wallet struct {
	ID          uint  `json:"wallet_id" gorm:"primaryKey;not null"`
	UsersID     uint  `json:"users_id" gorm:"not null"`
	Users       Users `json:"-"`
	TotalAmount uint  `json:"total_amount" gorm:"not null"`
}

type TransactionType string

const (
	Debit  TransactionType = "DEBIT"
	Credit TransactionType = "CREDIT"
)

type Transaction struct {
	TransactionID   uint            `json:"transaction_id" gorm:"primaryKey;not null"`
	WalletID        uint            `json:"wallet_id" gorm:"not null"`
	Wallet          Wallet          `json:"-"`
	TransactionDate time.Time       `json:"transaction_time" gorm:"not null"`
	Amount          uint            `json:"amount" gorm:"not null"`
	TransactionType TransactionType `json:"transaction_type" gorm:"not null"`
}
