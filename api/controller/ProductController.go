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
	"strconv"
)

type ProductController struct {
	service service.ProductService
}

func (c *ProductController) Init(e *echo.Echo) {
	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(auth.JWTCustomClaims)
		},
		SigningKey:  []byte("secret"),
		TokenLookup: "header:user",
	}
	e.GET("/product", c.getAllProducts)
	e.GET("/product/:id", c.getProductById)
	e.POST("/product", c.createProduct)
	e.DELETE("/product", c.removeProduct)
	e.GET("/productByCustomer", c.getProductByCustomer, echojwt.WithConfig(config))
}

func (c *ProductController) getAllProducts(ctx echo.Context) error {
	json.NewEncoder(ctx.Response()).Encode(c.service.GetAllProducts())

	return nil
}

func (c *ProductController) createProduct(ctx echo.Context) error {
	var product entity.Product

	json.NewDecoder(ctx.Request().Body).Decode(&product)
	c.service.CreateProduct(product)

	return nil
}

func (c *ProductController) removeProduct(ctx echo.Context) error {
	var product entity.Product

	json.NewDecoder(ctx.Request().Body).Decode(&product)
	c.service.RemoveProduct(product)

	return nil
}

func (c *ProductController) getProductById(ctx echo.Context) error {
	id := ctx.Param("id")
	newId, err := strconv.Atoi(id)
	if err != nil {
		panic(err)
	}
	ctx.JSON(http.StatusOK, c.service.GetProductById(newId))
	return nil
}

func (c *ProductController) getProductByCustomer(ctx echo.Context) error {
	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(*auth.JWTCustomClaims)

	products := c.service.GetProductByCustomerId(claims.Id)

	ctx.JSON(http.StatusOK, products)

	return nil
}
