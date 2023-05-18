package controllers

import (
	"assignment-2/database"
	"assignment-2/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllOrder(c *gin.Context) {
	var db = database.GetDB()

	var orders []models.OrderResponse
	err := db.Model(&models.Order{}).Preload("Items").Find(&orders).Error

	if err != nil {
		fmt.Println("Error getting order data :", err.Error())
	}

	c.JSON(http.StatusOK, gin.H{"Data": orders})
}

func CreateOrder(c *gin.Context) {
	var db = database.GetDB()

	var input models.Order
	var items []models.Item

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error(), "success": false})
		return
	}

	for _, item := range input.Items {
		items = append(items, models.Item{
			ItemCode:    item.ItemCode,
			Description: item.Description,
			Quantity:    item.Quantity,
		})
	}

	orderInput := models.Order{
		CustomerName: input.CustomerName,
		CreatedAt:    input.CreatedAt,
		Items:        items,
	}
	db.Create(&orderInput)
	c.JSON(http.StatusCreated, gin.H{"messages": "Create Data Success", "success": true})
}

func UpdateOrder(c *gin.Context) {
	var db = database.GetDB()
	tx := db.Begin()

	var input models.Order
	var items []models.Item
	if err := c.ShouldBindJSON(&input); err != nil {
		tx.Rollback()
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error(), "success": false})
		return
	}

	// Delete item
	tx.Delete(&models.Item{}, "order_id = ?", c.Param("id"))

	// Find Order
	var order models.Order
	err := tx.Model(&models.Order{}).Preload("Items").First(&order, "order_id = ?", c.Param("id")).Error
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusBadRequest, gin.H{"message": "Error get order data", "success": false})
		return
	}

	// Update Order
	errUpdate := tx.Model(&order).Updates(&input).Error
	if errUpdate != nil {
		tx.Rollback()
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error update order", "success": false})
		return
	}

	for _, _item := range input.Items {
		items = append(items, models.Item{
			ItemCode:    _item.ItemCode,
			Description: _item.Description,
			Quantity:    _item.Quantity,
		})
	}

	// Insert item
	tx.Model(&order).Association("Items").Replace(items)

	var orderResponse models.OrderUpdateResponse
	_ = tx.Model(&models.OrderUpdateResponse{}).Preload("Items").First(&orderResponse, "order_id = ?", c.Param("id")).Error

	tx.Commit()
	c.JSON(http.StatusOK, gin.H{"data": orderResponse, "message": "Update data success", "success": true})
}

func DeleteOrder(c *gin.Context) {
	var db = database.GetDB()

	var order models.Order
	err := db.First(&order, "order_id = ?", c.Param("id")).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error get order data", "success": false})
		return
	}

	db.Unscoped().Delete(&order)

	c.JSON(http.StatusOK, gin.H{"messages": "Delete Data Success", "success": true})
}
