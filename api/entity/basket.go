package entity

type Basket struct {
	CustomerId int `json:"customerId"`
	ProductId  int `json:"productId"`
}

func (b *Basket) TableName() string {
	return "customers_products"
}
