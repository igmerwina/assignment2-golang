package main

import (
	"gorestapi/config"
	"gorestapi/services"

	"github.com/labstack/echo/v4"
)

func main() {
	config.Connect()

	e := echo.New()

	e.GET("/", services.GetAllOrder)
	e.POST("/create", services.CreateOrder)
	// e.PUT("/update", services.UpdateEmployee)
	// e.DELETE("/:id", services.DeleteEmployee)

	e.Logger.Fatal(e.Start(":8021"))
}
