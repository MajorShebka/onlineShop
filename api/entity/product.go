package entity

type Product struct {
	Id        int        `json:"id"`
	Name      string     `json:"name"`
	Desc      string     `json:"desc"`
	Type      string     `json:"type"`
	Price     float64    `json:"price"`
	Image     string     `json:"image"`
	Customers []Customer `json:"customers" gorm:"many2many:customers_products;"`
}
