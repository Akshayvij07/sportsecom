package respondse

import "time"

type OrderResponse struct {
	ID                uint      `json:"order_ID"`
	UsersID            uint      `json:"-"`
	OrderDate         time.Time `json:"order_date"`
	PaymentMethodID   uint      `json:"payment_method_id"`
	PaymentMethod     string    `json:"PaymentMethod"`
	ShippingAddressID uint      `json:"shipping_address_id"`
	House_number      string    `json:"house_number"`
	Street            string    `json:"street"`
	City              string    `json:"city"`
	District          string    `json:"district"`
	Pincode           int       `json:"pin_code"`
	Landmark          string    `json:"land_mark"`
	//Discount          float64   `json:"discount"`
	OrderTotal        float64   `json:"order_total"`
	OrderStatusID     uint      `json:"order_status_id"`
	OrderStatus       string    `json:"orderStatus"`
	DeliveryUpdatedAt time.Time `json:"expected_delivery_time"`
}




