package request

type Update struct {
	OrderId  int `json:"order_id" binding:"required"`
	StatusId int `json:"status_id" binding:"required"`
}

type RazorPayRequest struct {
	RazorPayPaymentId  string
	RazorPayOrderId    string
	Razorpay_signature string
}
