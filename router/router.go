package router

import (
	"context"
	"service2/controllers"

	"github.com/omniful/go_commons/http"
)

// Initialize sets up all API routes following REST conventions
func Initialize(ctx context.Context, s *http.Server) (err error) {
	apiV1 := s.Engine.Group("/api/V1")

	// SKU Routes
	apiV1.GET("skus/validate/:id", controllers.ValidateSKU)
	apiV1.GET("/skus/byTenant/:id", controllers.FetchSKUsByTenant)
	apiV1.GET("/skus/byHub/:id", controllers.FetchSKUsInHub)
	apiV1.POST("/skus", controllers.CreateSKU)       // Create SKU
	apiV1.GET("/skus", controllers.GetAllSKUs)       // Get all SKUs
	apiV1.GET("/skus/:id", controllers.GetSKU)       // Get SKU by ID
	apiV1.PUT("/skus/:id", controllers.UpdateSKU)    // Update SKU
	apiV1.DELETE("/skus/:id", controllers.DeleteSKU) // Delete SKU

	// Tenant Routes
	apiV1.POST("/tenants", controllers.CreateTenant)       // Create Tenant
	apiV1.GET("/tenants", controllers.GetAllTenants)       // Get all Tenants
	apiV1.GET("/tenants/:id", controllers.GetTenant)       // Get Tenant by ID
	apiV1.PUT("/tenants/:id", controllers.UpdateTenant)    // Update Tenant
	apiV1.DELETE("/tenants/:id", controllers.DeleteTenant) // Delete Tenant

	// Hub Routes
	apiV1.GET("/hubs/validate/:id", controllers.ValidateHub)
	apiV1.POST("/hubs/multiple", controllers.CreateMultipleHubs) // Create Hub
	apiV1.POST("/hubs", controllers.CreateHub)                   // Create Hub
	apiV1.GET("/hubs", controllers.GetAllHubs)                   // Get all Hubs
	apiV1.GET("/hubs/:id", controllers.GetHub)                   // Get Hub by ID
	apiV1.PUT("/hubs/:id", controllers.UpdateHub)                // Update Hub
	apiV1.DELETE("/hubs/:id", controllers.DeleteHub)             // Delete Hub
	// Inventory Routes
	apiV1.POST("/inventories", controllers.CreateInventory)
	apiV1.GET("/inventories", controllers.GetAllInventories)
	apiV1.GET("/inventories/:id", controllers.GetInventory)
	apiV1.PUT("/inventories/:id", controllers.UpdateInventory)
	apiV1.DELETE("/inventories/:id", controllers.DeleteInventory)
	apiV1.GET("/inventories/validate/:id", controllers.CheckAndDecrementInventory)

	return nil
}
