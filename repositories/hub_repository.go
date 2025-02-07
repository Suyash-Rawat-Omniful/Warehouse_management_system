package repo

import (
	appinit "service2/init"
	"service2/models"
)

func CreateHub(hub *models.Hub) error {
	return appinit.DB.Create(hub).Error
}

func GetAllHubs() ([]models.Hub, error) {
	var hubs []models.Hub
	if err := appinit.DB.Find(&hubs).Error; err != nil {
		return nil, err
	}
	return hubs, nil
}

func GetHubByID(id uint) (models.Hub, error) {
	var hub models.Hub
	if err := appinit.DB.First(&hub, id).Error; err != nil {
		return hub, err
	}
	return hub, nil
}

func UpdateHub(hub *models.Hub) error {
	return appinit.DB.Save(hub).Error
}

func DeleteHub(id uint) error {
	return appinit.DB.Delete(&models.Hub{}, id).Error
}
