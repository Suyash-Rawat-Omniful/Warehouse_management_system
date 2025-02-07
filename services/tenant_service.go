package services

import (
	"errors"
	"service2/models"
	repo "service2/repositories"
)

// CreateTenant handles tenant creation
func CreateTenant(tenant *models.Tenant) error {
	return repo.CreateTenant(tenant)
}

// GetAllTenants retrieves all tenants
func GetAllTenants() ([]models.Tenant, error) {
	return repo.GetAllTenants()
}

// GetTenant retrieves a single tenant by ID
func GetTenant(id string) (*models.Tenant, error) {
	return repo.GetTenant(id)
}

// UpdateTenant updates an existing tenant
func UpdateTenant(id string, updatedData *models.Tenant) (*models.Tenant, error) {
	existingTenant, err := repo.GetTenant(id)
	if err != nil {
		return nil, errors.New("tenant not found")
	}

	// Update fields
	existingTenant.Name = updatedData.Name
	existingTenant.Email = updatedData.Email

	err = repo.UpdateTenant(existingTenant)
	if err != nil {
		return nil, err
	}
	return existingTenant, nil
}

// DeleteTenant deletes a tenant by ID
func DeleteTenant(id string) error {
	return repo.DeleteTenant(id)
}
