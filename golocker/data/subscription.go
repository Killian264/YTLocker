package data

import (
	"github.com/Killian264/YTLocker/golocker/models"
	uuid "github.com/satori/go.uuid"
)

func (d *Data) NewSubscription(request *models.SubscriptionRequest) error {
	request.UUID = uuid.NewV4().String()

	result := d.gormDB.Create(request)

	return result.Error
}

func (d *Data) GetSubscription(channelID string) (*models.SubscriptionRequest, error) {

	request := models.SubscriptionRequest{}

	result := d.gormDB.Where("channel_id = ?", channelID).First(&request)

	if result.Error != nil || NotFound(result.Error) {
		return nil, RemoveNotFound(result.Error)
	}

	return &request, nil

}

func (d *Data) InactivateAllSubscriptions() error {

	result := d.gormDB.Model(&models.SubscriptionRequest{}).Where(&models.SubscriptionRequest{Active: true}).Update("active", false)

	return result.Error
}

func (d *Data) GetInactiveSubscription() (*models.SubscriptionRequest, error) {

	sub := models.SubscriptionRequest{}

	result := d.gormDB.Where("active = false").First(&sub)

	if result.Error != nil || NotFound(result.Error) {
		return nil, RemoveNotFound(result.Error)
	}

	return &sub, nil
}

func (d *Data) DeleteSubscription(sub *models.SubscriptionRequest) error {

	result := d.gormDB.Where(&models.SubscriptionRequest{UUID: sub.UUID}).Delete(&models.SubscriptionRequest{UUID: sub.UUID})

	return result.Error

}
