package controller

import (
	"OnlineShop/api/auth"
	"OnlineShop/api/entity"
	"OnlineShop/api/service"
	"encoding/json"
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"net/http"
)

type CustomerController struct {
	service service.CustomerService
}

func (c *CustomerController) Init(e *echo.Echo) {
	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(auth.JWTCustomClaims)
		},
		SigningKey:  []byte("secret"),
		TokenLookup: "header:user",
	}
	e.POST("/customer", c.createCustomer)
	e.GET("/customer", c.getCustomers)
	e.DELETE("/customer", c.removeCustomer)
	e.POST("/productToCustomer", c.addProductToCustomer, echojwt.WithConfig(config))
	e.POST("/signIn", c.signIn)
}

func (c *CustomerController) createCustomer(ctx echo.Context) error {
	r := ctx.Request()
	var customer entity.Customer

	json.NewDecoder(r.Body).Decode(&customer)
	c.service.CreateCustomer(customer)

	return nil
}

func (c *CustomerController) getCustomers(ctx echo.Context) error {
	ctx.JSON(200, c.service.GetAllCustomers())
	return nil
}

func (c *CustomerController) removeCustomer(ctx echo.Context) error {
	var customer entity.Customer

	json.NewDecoder(ctx.Request().Body).Decode(&customer)
	c.service.RemoveCustomer(customer)

	return nil
}

func (c *CustomerController) addProductToCustomer(ctx echo.Context) error {
	var basket entity.Basket
	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(*auth.JWTCustomClaims)

	err := json.NewDecoder(ctx.Request().Body).Decode(&basket)
	if err != nil {
		panic(err)
	}
	basket.CustomerId = claims.Id
	c.service.AddProduct(basket)

	return nil
}

func (c *CustomerController) signIn(ctx echo.Context) error {
	var customer entity.Customer

	json.NewDecoder(ctx.Request().Body).Decode(&customer)
	customer = c.service.SignIn(customer)

	token := auth.GenerateToken(customer.Id)
	ctx.JSON(http.StatusOK, echo.Map{
		"token": token,
	})

	return nil
}
