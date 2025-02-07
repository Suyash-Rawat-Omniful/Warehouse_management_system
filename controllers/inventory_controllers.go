package controllers

import (
	"net/http"
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
