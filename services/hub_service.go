package services

import (
	"service2/models"
	repo "service2/repositories"
)

func CreateHub(hub *models.Hub) error {
	return repo.CreateHub(hub)
}

func GetAllHubs() ([]models.Hub, error) {
	return repo.GetAllHubs()
}

func GetHubByID(id uint) (models.Hub, error) {
	return repo.GetHubByID(id)
}

func UpdateHub(hub *models.Hub) error {
	return repo.UpdateHub(hub)
}

func DeleteHub(id uint) error {
	return repo.DeleteHub(id)
}
