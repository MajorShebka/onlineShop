package service

import (
	"OnlineShop/api/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ProductService struct {
}

func (s *ProductService) GetAllProducts() []entity.Product {
	databaseUrl := "postgres://root:root@localhost/online_shop"
	db, err := gorm.Open(postgres.Open(databaseUrl), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database!")
	}

	products := make([]entity.Product, 0)

	err = db.Find(&products).Error
	if err != nil {
		panic(err)
	}

	return products
}

func (s *ProductService) CreateProduct(p entity.Product) {
	databaseUrl := "postgres://root:root@localhost/online_shop"
	db, err := gorm.Open(postgres.Open(databaseUrl), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database!")
	}

	err = db.Create(&p).Error
	if err != nil {
		panic(err)
	}
}

func (s *ProductService) RemoveProduct(product entity.Product) {
	databaseUrl := "postgres://root:root@localhost/online_shop"
	db, err := gorm.Open(postgres.Open(databaseUrl), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database!")
	}

	err = db.Delete(&product).Error
	if err != nil {
		panic(err)
	}
}

func (s *ProductService) GetProductById(id int) entity.Product {
	databaseUrl := "postgres://root:root@localhost/online_shop"
	db, err := gorm.Open(postgres.Open(databaseUrl), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database!")
	}

	var product entity.Product

	err = db.Where("id = ?", id).Take(&product).Error
	if err != nil {
		panic(err)
	}

	return product
}

func (s *ProductService) GetProductByCustomerId(id int) []entity.Product {
	databaseUrl := "postgres://root:root@localhost/online_shop"
	db, err := gorm.Open(postgres.Open(databaseUrl), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database!")
	}
	products := make([]entity.Product, 0)
	//SELECT
	//  p.id as product_id
	//  ,p.name
	//  ,p.price
	//  ,r.stripe_data
	//FROM products p
	//LEFT JOIN purchases r ON p.id = r.product_id AND r.user_id = <your user>
	err = db.Table("products").Joins("inner join customers_products on customers_products.product_id=products.id and customers_products.customer_id=?", id).Find(&products).Error
	if err != nil {
		panic(err)
	}

	return products
}
