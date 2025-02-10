package controllers

import (
	"fmt"
	"net/http"
	appinit "service2/init"
	"service2/models"
	"service2/redis"
	"service2/services"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

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

func GetAllSKUs(ctx *gin.Context) {
	skus, err := services.GetAllSKUs()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve SKUs"})
		return
	}

	ctx.JSON(http.StatusOK, skus)
}

func GetSKU(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid SKU ID"})
		return
	}

	SKUFromRedis, err := redis.Client.HGetAll(ctx, "sku:"+idStr)

	if err == nil && len(SKUFromRedis) > 0 {
		ctx.JSON(http.StatusOK, SKUFromRedis)
		return
	}

	sku, err := services.GetSKUByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "SKU not found"})
		return
	}

	ID := strconv.FormatUint(uint64(sku.ID), 10)
	fields := map[string]interface{}{
		"ID":         sku.ID,
		"Product_ID": sku.ProductID,
		"Price":      sku.Price,
		"Name":       sku.Name,
		"Fragile":    sku.Fragile,
	}
	_, _ = redis.Client.HSetAll(ctx, "sku:"+ID, fields)

	ctx.JSON(http.StatusOK, sku)
}

func UpdateSKU(ctx *gin.Context) {
	var sku models.SKU
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
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
func DeleteSKU(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
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

func FetchSKUsByTenant(c *gin.Context) {
	tenantID := c.Param("id")

	fmt.Println("tenant_id is :", tenantID)
	tenantIDInt, err := strconv.Atoi(tenantID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid tenant_id"})
		return
	}

	var hubs []models.Hub
	err = appinit.DB.Where("tenant_id = ?", tenantIDInt).Find(&hubs).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch hubs"})
		return
	}

	var inventories []models.Inventory
	err = appinit.DB.Where("hub_id IN (?)", getHubIDs(hubs)).Find(&inventories).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch inventories"})
		return
	}

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

	c.JSON(http.StatusOK, skus)
}

func getHubIDs(hubs []models.Hub) []uint {
	var hubIDs []uint
	for _, hub := range hubs {
		hubIDs = append(hubIDs, hub.ID)
	}
	return hubIDs
}

func FetchSKUsInHub(c *gin.Context) {
	hubID := c.Param("id")
	hubIDInt, err := strconv.Atoi(hubID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid hub_id"})
		return
	}

	var inventories []models.Inventory
	err = appinit.DB.Where("hub_id = ?", hubIDInt).Find(&inventories).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch inventories"})
		return
	}

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

	c.JSON(http.StatusOK, skus)
}

func ValidateSKU(c *gin.Context) {
	skuID := c.Param("id")
	skuIDInt, err := strconv.Atoi(skuID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid sku_id"})
		return
	}

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

	c.JSON(http.StatusOK, gin.H{"message": "SKU exists"})
}
