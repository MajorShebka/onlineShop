package entity

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type Customer struct {
	Id       int       `json:"id"`
	Login    string    `json:"login"`
	Password string    `json:"password"`
	Email    string    `json:"email"`
	Role     string    `json:"role"`
	Products []Product `json:"products" gorm:"many2many:customers_products;"`
}

func (c *Customer) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Email, validation.Required, is.Email),
		validation.Field(&c.Role, validation.Required),
		validation.Field(&c.Login, validation.Required, validation.Length(4, 20)),
		validation.Field(&c.Password, validation.Required, validation.Length(8, 20)),
	)
}
