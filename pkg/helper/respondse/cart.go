package respondse

type Cartres struct {
	Product_Id   uint   `json:"product_item_id"`
	ProductName  string `json:"product_name"`
	Prize        uint   `json:"prize"`
	Qty_in_stock uint   `json:"-"`
	Qty          uint   `json:"qty"`
	Sub_Total    uint   `json:"sub_total"`
}
