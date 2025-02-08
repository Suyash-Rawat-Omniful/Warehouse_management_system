package controllers

import (
	"fmt"
	"net/http"
	appinit "service2/init"
	"service2/models"
	"service2/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateInventory creates a new inventory record
func CreateInventory(ctx *gin.Context) {
	var inventory models.Inventory
	if err := ctx.ShouldBindJSON(&inventory); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := services.CreateInventory(&inventory); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create inventory"})
		return
	}

	ctx.JSON(http.StatusCreated, inventory)
}

// GetInventory retrieves an inventory record by ID
func GetInventory(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32) // Convert string to uint
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid inventory ID"})
		return
	}

	inventory, err := services.GetInventoryByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Inventory not found"})
		return
	}

	ctx.JSON(http.StatusOK, inventory)
}

// GetAllInventories retrieves all inventory records
func GetAllInventories(ctx *gin.Context) {
	inventories, err := services.GetAllInventories()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve inventories"})
		return
	}

	ctx.JSON(http.StatusOK, inventories)
}

// UpdateInventory updates an inventory record
func UpdateInventory(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32) // Convert string to uint
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid inventory ID"})
		return
	}

	var requestData struct {
		Quantity int `json:"quantity"`
	}

	if err := ctx.ShouldBindJSON(&requestData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedInventory, err := services.UpdateInventory(uint(id), requestData.Quantity)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update inventory"})
		return
	}

	ctx.JSON(http.StatusOK, updatedInventory)
}

// DeleteInventory deletes an inventory record by ID
func DeleteInventory(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32) // Convert string to uint
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid inventory ID"})
		return
	}

	if err := services.DeleteInventory(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete inventory"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Inventory deleted successfully"})
}

// CheckAndDecrementInventory checks if inventory has enough quantity and decrements it if possible
func CheckAndDecrementInventory(ctx *gin.Context) {
	var requestData struct {
		SKUID         string `json:"sku_id"`
		GivenQuantity int    `json:"given_quantity"`
	}
	fmt.Print("hit aayi hai:", ctx.Request.Body, "\n\n")

	// Bind JSON request to struct
	if err := ctx.ShouldBindJSON(&requestData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Print(requestData.SKUID, "(ye sku id hai)--->", requestData.GivenQuantity, "ye given_quantity hai\n\n\n\n")

	// Find inventory item by SKU ID
	var inventory models.Inventory
	if err := appinit.DB.Where("sku_id = ?", requestData.SKUID).First(&inventory).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Inventory item not found"})
		return
	}

	fmt.Print("inventory bhi mil gayi\n\n\n")

	// Check if quantity is sufficient
	if inventory.Quantity < requestData.GivenQuantity {
		ctx.JSON(http.StatusConflict, gin.H{"error": "Not enough inventory available"})
		return
	}

	fmt.Print("decrement bhi karliya\n\n\n")

	// Decrement inventory quantity
	inventory.Quantity -= requestData.GivenQuantity
	if err := appinit.DB.Save(&inventory).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update inventory"})
		return
	}

	// Success response
	ctx.JSON(http.StatusOK, gin.H{"message": "Inventory updated successfully"})
}
