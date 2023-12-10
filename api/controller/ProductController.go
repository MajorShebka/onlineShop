package controller

import (
	"OnlineShop/api/entity"
	"OnlineShop/api/service"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type ProductController struct {
	service service.ProductService
}

func (c *ProductController) Init(e *echo.Echo) {
	e.GET("/product", c.getAllProducts)
	e.GET("/product/:id", c.getProductById)
	e.POST("/product", c.createProduct)
	e.DELETE("/product", c.removeProduct)
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
