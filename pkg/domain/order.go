package domain

import "time"

type PaymentMethod struct {
	ID            uint   `gorm:"primaryKey"`
	PaymentMethod string `json:"payment_method"`
}

type PaymentStatus struct {
	ID            uint   `gorm:"primaryKey"`
	PaymentStatus string `json:"payment_status,omitempty"`
}

type Orders struct {
	ID                uint          `gorm:"primaryKey"`
	UsersID           int           `json:"users_id"`
	Users             Users         `gorm:"foreignKey:UsersID" json:"-"`
	OrderDate         time.Time     `json:"order_date"`
	PaymentMethodID   uint          `json:"payment_method_id"`
	PaymentMethod     PaymentMethod `gorm:"foreignKey:PaymentMethodID" json:"-"`
	ShippingAddressID uint          `json:"shipping_address_id"`
	UserAddress       UserAddress   `gorm:"foreignKey:ShippingAddressID" json:"-"`
	OrderTotal        float64       `json:"order_total"`
	OrderStatusID     uint          `json:"order_status_id"`
	OrderStatus       OrderStatus   `gorm:"foreignKey:OrderStatusID" json:"-"`
	DeliveryUpdatedAt time.Time     `json:"delivery_time"`
	Discount          float64       `json:"discount"`
	CouponCode        string        `json:"coupon_code"`
}

type OrderLine struct {
	ID        uint    `gorm:"primaryKey"`
	ProductID uint    `json:"product_id"`
	Product   Product ` json:"-"`
	OrderID   uint    `json:"order_Id"`
	Order     Orders
	Qty       int     `json:"qty"`
	Price     float64 `json:"price"`
}

type OrderStatus struct {
	ID          uint   `gorm:"primaryKey"`
	OrderStatus string `json:"order_status"`
}

type DeliveryStatus struct {
	ID     uint   `json:"id" gorm:"primaryKey"`
	Status string `json:"status"`
}

type PaymentDetails struct {
	ID              uint          `gorm:"primaryKey" json:"id,omitempty"`
	OrdersID        uint          `json:"order_id,omitempty"`
	Orders          Orders        `gorm:"foreignKey:OrdersID" json:"-"`
	OrderTotal      float64       `json:"order_total"`
	PaymentMethodID uint          `json:"payment_method_id"`
	PaymentMethod   PaymentMethod `gorm:"foreignKey:PaymentMethodID"`
	PaymentStatusID uint          `json:"payment_status_id,omitempty"`
	PaymentStatus   PaymentStatus `gorm:"foreignKey:PaymentStatusID" json:"-"`
	UpdatedAt       time.Time
}

/*type Invoice struct {
	ID          uint
	InvoiceNumber string
	UsersID     uint
	Users       Users
	OrderID     uint
	Orders      Orders
	Date        time.Time
	TotalAmount float64
}*/

type OrderEnhancement struct {
	ID             uint
	OrderCompleted bool
	Users          Users
	OrderID        uint
	Orders         Orders
	OrderStatusID  uint
	OrderStatus    OrderStatus
}
