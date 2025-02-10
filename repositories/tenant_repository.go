package repo

import (
	"errors"
	appinit "service2/init"
	"service2/models"
)

func CreateTenant(tenant *models.Tenant) error {
	return appinit.DB.Create(tenant).Error
}

func GetAllTenants() ([]models.Tenant, error) {
	var tenants []models.Tenant
	err := appinit.DB.Find(&tenants).Error
	return tenants, err
}

func GetTenant(id string) (*models.Tenant, error) {
	var tenant models.Tenant
	if err := appinit.DB.First(&tenant, id).Error; err != nil {
		return nil, errors.New("tenant not found")
	}
	return &tenant, nil
}

func UpdateTenant(tenant *models.Tenant) error {
	return appinit.DB.Save(tenant).Error
}

func DeleteTenant(id string) error {
	if err := appinit.DB.Delete(&models.Tenant{}, id).Error; err != nil {
		return err
	}
	return nil
}
