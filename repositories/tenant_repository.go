package repo

import (
	"errors"
	appinit "service2/init"
	"service2/models"
)

// CreateTenant saves a new tenant record
func CreateTenant(tenant *models.Tenant) error {
	return appinit.DB.Create(tenant).Error
}

// GetAllTenants retrieves all tenants
func GetAllTenants() ([]models.Tenant, error) {
	var tenants []models.Tenant
	err := appinit.DB.Find(&tenants).Error
	return tenants, err
}

// GetTenant retrieves a single tenant by ID
func GetTenant(id string) (*models.Tenant, error) {
	var tenant models.Tenant
	if err := appinit.DB.First(&tenant, id).Error; err != nil {
		return nil, errors.New("tenant not found")
	}
	return &tenant, nil
}

// UpdateTenant updates an existing tenant
func UpdateTenant(tenant *models.Tenant) error {
	return appinit.DB.Save(tenant).Error
}

// DeleteTenant deletes a tenant by ID
func DeleteTenant(id string) error {
	if err := appinit.DB.Delete(&models.Tenant{}, id).Error; err != nil {
		return err
	}
	return nil
}
