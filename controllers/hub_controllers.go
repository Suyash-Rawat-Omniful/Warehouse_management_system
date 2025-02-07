package controllers

import (
	"net/http"
	appinit "service2/init"
	"service2/models"
	"service2/services"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// CreateHub creates a new hub
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

// GetAllHubs retrieves all hubs
func GetAllHubs(ctx *gin.Context) {
	hubs, err := services.GetAllHubs()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve Hubs"})
		return
	}

	ctx.JSON(http.StatusOK, hubs)
}

// GetHub retrieves a single hub by ID
func GetHub(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Hub ID"})
		return
	}

	hub, err := services.GetHubByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Hub not found"})
		return
	}

	ctx.JSON(http.StatusOK, hub)
}

// UpdateHub updates an existing hub
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

// DeleteHub deletes a hub by ID
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

	// Convert to proper data type
	hubIDInt, err := strconv.Atoi(hubID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid hub_id"})
		return
	}

	// Check if the Hub exists in the database
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

	// If hub exists, return success message
	c.JSON(http.StatusOK, gin.H{"message": "Hub exists"})
}
