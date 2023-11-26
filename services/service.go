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

	response := model.DataResponse{
		Data: order,
	}

	fmt.Println("GetAllOrder")

	return c.JSON(http.StatusOK, response)
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

func UpdateOrder(c echo.Context) error {
	db := config.GetDB()
	id := c.Param("orderId")

	var order model.Order
	if err := db.First(&order, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Order not found"})
	}

	var updatedOrder model.Order
	if err := c.Bind(&updatedOrder); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	db.Model(&order).Updates(&updatedOrder)
	return c.JSON(http.StatusOK, order)
}

func DeleteOrder(c echo.Context) error {
	db := config.GetDB()

	order := model.Order{}

	delResp := model.DeleteStruct{
		Status:  http.StatusOK,
		Message: "Delete Data Success",
		Success: true,
	}

	if err := c.Bind(&order); err != nil {
		return err
	}

	paramId := c.Param("orderId")

	// Delete related items first
	if err := db.Where("order_id = ?", paramId).Delete(&model.Item{}).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
	}

	// Then delete the order
	if err := db.Delete(&model.Order{}, paramId).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
	}

	db.Model(&order).Where("order_id = ?", paramId).Delete(&order)

	fmt.Println("DeleteOrder")

	return c.JSON(http.StatusOK, delResp)
}
