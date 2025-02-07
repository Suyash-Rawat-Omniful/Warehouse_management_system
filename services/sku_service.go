package services

import (
	"service2/models"
	repo "service2/repositories"
)

func CreateSKU(sku *models.SKU) error {
	return repo.CreateSKU(sku)
}

func GetAllSKUs() ([]models.SKU, error) {
	return repo.GetAllSKUs()
}

func GetSKUByID(id uint) (models.SKU, error) {
	return repo.GetSKUByID(id)
}

func UpdateSKU(sku *models.SKU) error {
	return repo.UpdateSKU(sku)
}

func DeleteSKU(id uint) error {
	return repo.DeleteSKU(id)
}
