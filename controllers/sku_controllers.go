package controllers

import (
	"fmt"
	"net/http"
	appinit "service2/init"
	"service2/models"
	"service2/services"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// CreateSKU creates a new SKU
func CreateSKU(ctx *gin.Context) {
	var sku models.SKU
	if err := ctx.ShouldBindJSON(&sku); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := services.CreateSKU(&sku); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create SKU"})
		return
	}

	ctx.JSON(http.StatusCreated, sku)
}

// GetAllSKUs retrieves all SKUs
func GetAllSKUs(ctx *gin.Context) {
	skus, err := services.GetAllSKUs()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve SKUs"})
		return
	}

	ctx.JSON(http.StatusOK, skus)
}

// GetSKU retrieves a SKU by ID
func GetSKU(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32) // Convert string to uint
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid SKU ID"})
		return
	}

	sku, err := services.GetSKUByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "SKU not found"})
		return
	}

	ctx.JSON(http.StatusOK, sku)
}

// UpdateSKU updates a SKU
func UpdateSKU(ctx *gin.Context) {
	var sku models.SKU
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32) // Convert string to uint
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid SKU ID"})
		return
	}

	if err := ctx.ShouldBindJSON(&sku); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sku.ID = uint(id)
	if err := services.UpdateSKU(&sku); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update SKU"})
		return
	}

	ctx.JSON(http.StatusOK, sku)
}

// DeleteSKU deletes a SKU by ID
func DeleteSKU(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32) // Convert string to uint
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid SKU ID"})
		return
	}

	if err := services.DeleteSKU(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete SKU"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "SKU deleted"})
}

// FetchSKUsByTenant fetches SKUs for a specific tenant
func FetchSKUsByTenant(c *gin.Context) {
	tenantID := c.Param("id")

	//print the tenant id
	fmt.Println("tenant_id is :", tenantID)

	// Convert to proper data type
	tenantIDInt, err := strconv.Atoi(tenantID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid tenant_id"})
		return
	}

	// Fetch hubs for the specific tenant
	var hubs []models.Hub
	err = appinit.DB.Where("tenant_id = ?", tenantIDInt).Find(&hubs).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch hubs"})
		return
	}

	// Fetch inventories for the hubs
	var inventories []models.Inventory
	err = appinit.DB.Where("hub_id IN (?)", getHubIDs(hubs)).Find(&inventories).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch inventories"})
		return
	}

	// Fetch SKUs for the inventories
	var skuIDs []uint
	for _, inventory := range inventories {
		skuIDs = append(skuIDs, inventory.Sku_id)
	}

	var skus []models.SKU
	err = appinit.DB.Where("id IN (?)", skuIDs).Find(&skus).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch SKUs"})
		return
	}

	// Return SKUs as response
	c.JSON(http.StatusOK, skus)
}

// FetchSKUsBySeller fetches SKUs for a specific seller

// Helper function to extract hub IDs from hubs array
func getHubIDs(hubs []models.Hub) []uint {
	var hubIDs []uint
	for _, hub := range hubs {
		hubIDs = append(hubIDs, hub.ID)
	}
	return hubIDs
}

func FetchSKUsInHub(c *gin.Context) {
	hubID := c.Param("id")

	// Convert to proper data type
	hubIDInt, err := strconv.Atoi(hubID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid hub_id"})
		return
	}

	// Fetch inventories for the given hub
	var inventories []models.Inventory
	err = appinit.DB.Where("hub_id = ?", hubIDInt).Find(&inventories).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch inventories"})
		return
	}

	// Fetch SKUs for the inventories
	var skuIDs []uint
	for _, inventory := range inventories {
		skuIDs = append(skuIDs, inventory.Sku_id)
	}

	var skus []models.SKU
	err = appinit.DB.Where("id IN (?)", skuIDs).Find(&skus).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch SKUs"})
		return
	}

	// Return SKUs as response
	c.JSON(http.StatusOK, skus)
}

func ValidateSKU(c *gin.Context) {
	skuID := c.Param("id")
	fmt.Println("validate sku called on id -> ", skuID)

	// Convert to proper data type
	skuIDInt, err := strconv.Atoi(skuID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid sku_id"})
		return
	}

	// Check if the SKU exists in the database
	var sku models.SKU
	err = appinit.DB.Where("id = ?", skuIDInt).First(&sku).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "SKU not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to check SKU existence"})
		}
		return
	}

	// If SKU exists, return success message
	c.JSON(http.StatusOK, gin.H{"message": "SKU exists"})
}
