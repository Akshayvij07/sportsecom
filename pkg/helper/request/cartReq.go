package request

type Cartreq struct {
	ProductId int
	UserID    int `json:"-"`
}

type Addcount struct {
	UserID    int `json:"-"`
	ProductId int `json:"product_id" binding:"required"`
	Count     int `json:"count" binding:"omitempty,gte=1"`
}

type CartItems struct {
	ProductId    int
	Qty          int
	Prize        int
	Qty_In_Stock int
}
