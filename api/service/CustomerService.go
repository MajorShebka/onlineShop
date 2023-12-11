package service

import (
	"OnlineShop/entity"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type CustomerService struct {
}

func GetDb() *gorm.DB {
	databaseUrl := "postgres://root:root@pg/root"
	db, err := gorm.Open(postgres.Open(databaseUrl), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database!")
	}

	return db
}

func (s *CustomerService) CreateCustomer(customer entity.Customer) {
	db := GetDb()

	hashedPass, _ := bcrypt.GenerateFromPassword([]byte(customer.Password), 1)
	customer.Password = string(hashedPass)

	err := db.Create(&customer).Error
	if err != nil {
		panic(err)
	}
}

func (s *CustomerService) GetAllCustomers() []entity.Customer {
	db := GetDb()

	customers := make([]entity.Customer, 0)

	err := db.Preload("Products").Find(&customers).Error
	if err != nil {
		panic(err)
	}

	return customers
}

func (s *CustomerService) SignIn(customer entity.Customer) entity.Customer {
	db := GetDb()
	pass := customer.Password

	err := db.Preload("Products").Where("login = ?", customer.Login).Take(&customer).Error
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
	db := GetDb()

	err := db.Delete(&customer).Error
	if err != nil {
		panic(err)
	}
}

func (s *CustomerService) AddProduct(basket entity.Basket) {
	db := GetDb()

	err := db.Create(&basket).Error
	if err != nil {
		panic(err)
	}
}
