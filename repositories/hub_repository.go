package repo

import (
	appinit "service2/init"
	"service2/models"
)

func CreateHub(hub *models.Hub) error {
	// ctx := context.Background()
	// HUB_ID := hub.ID
	// TENANT_ID := hub.TenantID
	// fields := map[string]interface{}{
	// 	"hub_id":    HUB_ID,
	// 	"tenant_id": TENANT_ID,
	// }
	// var key string = "hub:" + strconv.Itoa(int(HUB_ID))
	// count, err := redis.Client.HSetAll(ctx, key, fields)
	// if err != nil {
	// 	fmt.Println("can't save the hub in redis : ", fields)
	// } else {
	// 	fmt.Println("save the hub in redis : ", fields, " and count is ", count)
	// }
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
