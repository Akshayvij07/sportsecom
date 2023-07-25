package respondse

import "time"

type SalesReport struct {
	Id             string
	Name           string
	Payment_method string
	OrderDate      time.Time
	Order_Total    int
	Mobile         string
	HouseNumber    string
	Pincode        string
}

type AdminDashboard struct {
	CompletedOrders int     `json:"completed_orders,omitempty"`
	PendingOrders   int     `json:"pending_orders,omitempty"`
	CancelledOrders int     `json:"cancelled_orders,omitempty"`
	TotalOrders     int     `json:"total_orders,omitempty"`
	TotalOrderItems int     `json:"total_order_items,omitempty"`
	OrderValue      float64 `json:"order_value,omitempty"`
	CreditedAmount  float64 `json:"credited_amount,omitempty"`
	PendingAmount   float64 `json:"pending_amount,omitempty"`
	TotalUsers      int     `json:"total_users,omitempty"`
	VerifiedUsers   int     `json:"verified_users,omitempty"`
	OrderedUsers    int     `json:"ordered_users,omitempty"`
}
