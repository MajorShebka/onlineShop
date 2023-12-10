package main

import (
	"OnlineShop/api/controller"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())
	(&controller.CustomerController{}).Init(e)
	(&controller.ProductController{}).Init(e)
	e.Logger.Fatal(e.Start(":8080"))
}
