package repo

import (
	appinit "service2/init"
	"service2/models"
)

func CreateSKU(sku *models.SKU) error {
	return appinit.DB.Create(sku).Error
}

func GetAllSKUs() ([]models.SKU, error) {
	var skus []models.SKU
	if err := appinit.DB.Find(&skus).Error; err != nil {
		return nil, err
	}
	return skus, nil
}

func GetSKUByID(id uint) (models.SKU, error) {
	var sku models.SKU
	if err := appinit.DB.First(&sku, id).Error; err != nil {
		return sku, err
	}
	return sku, nil
}

func UpdateSKU(sku *models.SKU) error {
	return appinit.DB.Save(sku).Error
}

func DeleteSKU(id uint) error {
	return appinit.DB.Delete(&models.SKU{}, id).Error
}
