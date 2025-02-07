package controllers

import (
	"net/http"
	"service2/models"
	"service2/services"

	"github.com/gin-gonic/gin"
)

// CreateTenant creates a new tenant
func CreateTenant(c *gin.Context) {
	var tenant models.Tenant
	if err := c.ShouldBindJSON(&tenant); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := services.CreateTenant(&tenant); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create tenant"})
		return
	}

	c.JSON(http.StatusCreated, tenant)
}

// GetAllTenants retrieves all tenants
func GetAllTenants(c *gin.Context) {
	tenants, err := services.GetAllTenants()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch tenants"})
		return
	}
	c.JSON(http.StatusOK, tenants)
}

// GetTenant retrieves a single tenant by ID
func GetTenant(c *gin.Context) {
	id := c.Param("id")

	tenant, err := services.GetTenant(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Tenant not found"})
		return
	}

	c.JSON(http.StatusOK, tenant)
}

// UpdateTenant updates an existing tenant
func UpdateTenant(c *gin.Context) {
	id := c.Param("id")

	var tenant models.Tenant
	if err := c.ShouldBindJSON(&tenant); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedTenant, err := services.UpdateTenant(id, &tenant)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update tenant"})
		return
	}

	c.JSON(http.StatusOK, updatedTenant)
}

// DeleteTenant deletes a tenant by ID
func DeleteTenant(c *gin.Context) {
	id := c.Param("id")

	if err := services.DeleteTenant(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete tenant"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Tenant deleted successfully"})
}
