package controllers

import (
	"net/http"
	appinit "service2/init"
	"service2/models"
	redis "service2/redis"
	"service2/services"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func CreateMultipleHubs(ctx *gin.Context) {
	var hub models.Hub

	// Bind JSON once before loop
	if err := ctx.ShouldBindJSON(&hub); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		return
	}

	for i := 100; i <= 100000; i++ {
		newHub := hub
		newHub.ID += uint(i)

		if err := services.CreateHub(&newHub); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create Hub"})
			return
		}

	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Multiple hubs created successfully"})
}

func CreateHub(ctx *gin.Context) {
	var hub models.Hub
	if err := ctx.ShouldBindJSON(&hub); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := services.CreateHub(&hub); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create Hub"})
		return
	}

	ctx.JSON(http.StatusCreated, hub)
}

func GetAllHubs(ctx *gin.Context) {
	hubs, err := services.GetAllHubs()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve Hubs"})
		return
	}

	ctx.JSON(http.StatusOK, hubs)
}

func GetHub(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Hub ID"})
		return
	}

	HubFromRedis, err := redis.Client.HGetAll(ctx, "hub:"+idStr)

	if err == nil && len(HubFromRedis) > 0 {
		ctx.JSON(http.StatusOK, HubFromRedis)
		return
	}

	hub, err := services.GetHubByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Hub not found"})
		return
	}
	ID := strconv.FormatUint(uint64(hub.ID), 10)
	fields := map[string]interface{}{
		"ID":        hub.ID,
		"Tenant_ID": hub.TenantID,
	}
	_, _ = redis.Client.HSetAll(ctx, "hub:"+ID, fields)
	ctx.JSON(http.StatusOK, hub)
}

func UpdateHub(ctx *gin.Context) {
	var hub models.Hub
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Hub ID"})
		return
	}

	if err := ctx.ShouldBindJSON(&hub); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hub.ID = uint(id)
	if err := services.UpdateHub(&hub); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update Hub"})
		return
	}

	ctx.JSON(http.StatusOK, hub)
}

func DeleteHub(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Hub ID"})
		return
	}

	if err := services.DeleteHub(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete Hub"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Hub deleted"})
}

func ValidateHub(c *gin.Context) {
	hubID := c.Param("id")
	hubIDInt, err := strconv.Atoi(hubID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid hub_id"})
		return
	}
	var hub models.Hub
	err = appinit.DB.Where("id = ?", hubIDInt).First(&hub).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Hub not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to check hub existence"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Hub exists"})
}
