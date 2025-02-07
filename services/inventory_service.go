package services

import (
	"service2/models"
	repo "service2/repositories"
)

func CreateInventory(inventory *models.Inventory) error {
	return repo.CreateInventory(inventory)
}

func GetInventoryByID(id uint) (models.Inventory, error) {
	return repo.GetInventoryByID(id)
}

func GetAllInventories() ([]models.Inventory, error) {
	return repo.GetAllInventories()
}

func UpdateInventory(id uint, quantity int) (models.Inventory, error) {
	return repo.UpdateInventory(id, quantity)
}

func DeleteInventory(id uint) error {
	return repo.DeleteInventory(id)
}
