package main

import (
	"gorestapi/config"
	"gorestapi/services"

	"github.com/labstack/echo/v4"
)

func main() {
	config.Connect()

	e := echo.New()

	e.GET("/", services.Hello)
	e.GET("/getall", services.GetAllOrder)
	e.POST("/create", services.CreateOrder)
	e.PUT("/update/:orderId", services.UpdateOrder)
	e.DELETE("/delete/:orderId", services.DeleteOrder)

	e.Logger.Fatal(e.Start(":8021"))
}
