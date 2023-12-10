package service

import (
	"OnlineShop/api/entity"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type CustomerService struct {
}

func (s *CustomerService) CreateCustomer(customer entity.Customer) {
	databaseUrl := "postgres://root:root@localhost/online_shop"
	db, err := gorm.Open(postgres.Open(databaseUrl), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database!")
	}

	hashedPass, _ := bcrypt.GenerateFromPassword([]byte(customer.Password), 1)
	customer.Password = string(hashedPass)

	err = db.Create(&customer).Error
	if err != nil {
		panic(err)
	}
}

func (s *CustomerService) GetAllCustomers() []entity.Customer {
	databaseUrl := "postgres://root:root@localhost/online_shop"
	db, err := gorm.Open(postgres.Open(databaseUrl), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database!")
	}

	customers := make([]entity.Customer, 0)

	err = db.Preload("Products").Find(&customers).Error
	if err != nil {
		panic(err)
	}

	return customers
}

func (s *CustomerService) SignIn(customer entity.Customer) entity.Customer {
	databaseUrl := "postgres://root:root@localhost/online_shop"
	db, err := gorm.Open(postgres.Open(databaseUrl), &gorm.Config{})
	pass := customer.Password
	if err != nil {
		panic("Failed to connect database!")
	}

	err = db.Preload("Products").Where("login = ?", customer.Login).Take(&customer).Error
	if err != nil {
		panic(err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(customer.Password), []byte(pass))
	if err != nil {
		panic(err)
	}

	return customer
}

func (s *CustomerService) RemoveCustomer(customer entity.Customer) {
	databaseUrl := "postgres://root:root@localhost/online_shop"
	db, err := gorm.Open(postgres.Open(databaseUrl), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database!")
	}

	err = db.Delete(&customer).Error
	if err != nil {
		panic(err)
	}
}

func (s *CustomerService) AddProduct(basket entity.Basket) {
	databaseUrl := "postgres://root:root@localhost/online_shop"
	db, err := gorm.Open(postgres.Open(databaseUrl), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database!")
	}

	err = db.Create(&basket).Error
	if err != nil {
		panic(err)
	}
}
