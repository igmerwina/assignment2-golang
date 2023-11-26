package services

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"gorestapi/config"
	"gorestapi/model"
)

func Hello(c echo.Context) error {
	return c.String(http.StatusOK, "Assignment 2 - IGM Erwin A\n")
}

func GetAllOrder(c echo.Context) error {
	db := config.GetDB()

	var order []model.Order
	db.Preload("Items").Find(&order)

	fmt.Println("GetAllOrder")

	return c.JSON(http.StatusOK, order)
}

func CreateOrder(c echo.Context) error {
	db := config.GetDB()
	order := model.Order{}

	fmt.Println(order)

	if err := c.Bind(&order); err != nil {
		return err
	}

	db.Debug().Create(&order)

	fmt.Println("CreateOrder")
	return c.JSON(http.StatusOK, order)
}
