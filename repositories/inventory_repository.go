package repo

import (
	"errors"
	appinit "service2/init"
	"service2/models"
)

// CreateInventory inserts a new inventory record into the database
func CreateInventory(inventory *models.Inventory) error {
	return appinit.DB.Create(inventory).Error
}

// GetInventoryByID fetches an inventory record by ID
func GetInventoryByID(id uint) (models.Inventory, error) {
	var inventory models.Inventory
	if err := appinit.DB.First(&inventory, id).Error; err != nil {
		return inventory, err
	}
	return inventory, nil
}

// GetAllInventories fetches all inventory records
func GetAllInventories() ([]models.Inventory, error) {
	var inventories []models.Inventory
	if err := appinit.DB.Find(&inventories).Error; err != nil {
		return nil, err
	}
	return inventories, nil
}

// UpdateInventory updates an inventory record, creating a new one if it doesn't exist
func UpdateInventory(id uint, quantity int) (models.Inventory, error) {
	var inventory models.Inventory

	// Check if inventory exists
	if err := appinit.DB.First(&inventory, id).Error; err != nil {
		// If not found, create a new inventory record
		newInventory := models.Inventory{
			ID:       id,
			Quantity: quantity,
		}

		if err := appinit.DB.Create(&newInventory).Error; err != nil {
			return newInventory, err
		}
		return newInventory, nil
	}

	// If found, update the quantity
	inventory.Quantity += quantity
	if err := appinit.DB.Save(&inventory).Error; err != nil {
		return inventory, err
	}

	return inventory, nil
}

// DeleteInventory deletes an inventory record by ID
func DeleteInventory(id uint) error {
	if result := appinit.DB.Delete(&models.Inventory{}, id); result.Error != nil {
		return result.Error
	}
	if result := appinit.DB.RowsAffected; result == 0 {
		return errors.New("inventory not found")
	}
	return nil
}
