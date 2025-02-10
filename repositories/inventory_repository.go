package repo

import (
	"errors"
	appinit "service2/init"
	"service2/models"
)

func CreateInventory(inventory *models.Inventory) error {
	return appinit.DB.Create(inventory).Error
}

func GetInventoryByID(id uint) (models.Inventory, error) {
	var inventory models.Inventory
	if err := appinit.DB.First(&inventory, id).Error; err != nil {
		return inventory, err
	}
	return inventory, nil
}

func GetAllInventories() ([]models.Inventory, error) {
	var inventories []models.Inventory
	if err := appinit.DB.Find(&inventories).Error; err != nil {
		return nil, err
	}
	return inventories, nil
}

func UpdateInventory(id uint, quantity int) (models.Inventory, error) {
	var inventory models.Inventory

	if err := appinit.DB.First(&inventory, id).Error; err != nil {
		newInventory := models.Inventory{
			ID:       id,
			Quantity: quantity,
		}

		if err := appinit.DB.Create(&newInventory).Error; err != nil {
			return newInventory, err
		}
		return newInventory, nil
	}
	inventory.Quantity += quantity
	if err := appinit.DB.Save(&inventory).Error; err != nil {
		return inventory, err
	}

	return inventory, nil
}

func DeleteInventory(id uint) error {
	if result := appinit.DB.Delete(&models.Inventory{}, id); result.Error != nil {
		return result.Error
	}
	if result := appinit.DB.RowsAffected; result == 0 {
		return errors.New("inventory not found")
	}
	return nil
}
