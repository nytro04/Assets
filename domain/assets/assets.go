package assets

type Asset struct {
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	ProductId int64  `json:"product_id"`
	Price     int64  `json:"price"`
}
