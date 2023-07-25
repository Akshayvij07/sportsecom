package request

type Cartreq struct {
	UserID int    `json:"-"`
	Sku    string `json:"sku" binding:"required"`
}

type Addcount struct {
	UserID int    `json:"-"`
	Sku    string `json:"sku" binding:"required"`
	Count  int    `json:"count" binding:"omitempty,gte=1"`
}

type CartItems struct {
	ProductId    int
	Qty          int
	Prize        int
	Qty_In_Stock int
}
