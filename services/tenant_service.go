package services

import (
	"errors"
	"service2/models"
	repo "service2/repositories"
)

func CreateTenant(tenant *models.Tenant) error {
	return repo.CreateTenant(tenant)
}

func GetAllTenants() ([]models.Tenant, error) {
	return repo.GetAllTenants()
}

func GetTenant(id string) (*models.Tenant, error) {
	return repo.GetTenant(id)
}

func UpdateTenant(id string, updatedData *models.Tenant) (*models.Tenant, error) {
	existingTenant, err := repo.GetTenant(id)
	if err != nil {
		return nil, errors.New("tenant not found")
	}

	existingTenant.Name = updatedData.Name
	existingTenant.Email = updatedData.Email

	err = repo.UpdateTenant(existingTenant)
	if err != nil {
		return nil, err
	}
	return existingTenant, nil
}

func DeleteTenant(id string) error {
	return repo.DeleteTenant(id)
}
