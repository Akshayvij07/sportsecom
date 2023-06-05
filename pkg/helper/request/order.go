package request

type Update struct {
	OrderId  int `json:"order_id" binding:"required"`
	StatusId int `json:"status_id" binding:"required"`
}
