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
	apiV1.POST("/skus", controllers.CreateSKU)
	apiV1.GET("/skus", controllers.GetAllSKUs)
	apiV1.GET("/skus/:id", controllers.GetSKU)
	apiV1.PUT("/skus/:id", controllers.UpdateSKU)
	apiV1.DELETE("/skus/:id", controllers.DeleteSKU)

	// Tenant Routes
	apiV1.POST("/tenants", controllers.CreateTenant)
	apiV1.GET("/tenants", controllers.GetAllTenants)
	apiV1.GET("/tenants/:id", controllers.GetTenant)
	apiV1.PUT("/tenants/:id", controllers.UpdateTenant)
	apiV1.DELETE("/tenants/:id", controllers.DeleteTenant)

	// Hub Routes
	apiV1.GET("/hubs/validate/:id", controllers.ValidateHub)
	apiV1.POST("/hubs/multiple", controllers.CreateMultipleHubs)
	apiV1.POST("/hubs", controllers.CreateHub)
	apiV1.GET("/hubs", controllers.GetAllHubs)
	apiV1.GET("/hubs/:id", controllers.GetHub)
	apiV1.PUT("/hubs/:id", controllers.UpdateHub)
	apiV1.DELETE("/hubs/:id", controllers.DeleteHub)
	// Inventory Routes
	apiV1.POST("/inventories", controllers.CreateInventory)
	apiV1.GET("/inventories", controllers.GetAllInventories)
	apiV1.GET("/inventories/:id", controllers.GetInventory)
	apiV1.PUT("/inventories/:id", controllers.UpdateInventory)
	apiV1.DELETE("/inventories/:id", controllers.DeleteInventory)
	apiV1.GET("/inventories/validate/:id", controllers.CheckAndDecrementInventory)

	return nil
}
