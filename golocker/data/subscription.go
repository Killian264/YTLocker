package data

import (
	"github.com/Killian264/YTLocker/golocker/models"
)

func (d *Data) NewSubscription(request *models.SubscriptionRequest) error {
	request.ID = d.rand.ID()

	result := d.db.Create(request)

	return result.Error
}

func (d *Data) GetSubscription(channelID uint64, secret string) (*models.SubscriptionRequest, error) {

	request := models.SubscriptionRequest{ChannelID: channelID, Secret: secret}

	result := d.db.Where(&request).First(&request)

	if result.Error != nil || notFound(result.Error) {
		return nil, removeNotFound(result.Error)
	}

	return &request, nil

}

func (d *Data) InactivateAllSubscriptions() error {

	result := d.db.Model(&models.SubscriptionRequest{}).Where(&models.SubscriptionRequest{Active: true}).Update("active", false)

	return result.Error
}

func (d *Data) GetInactiveSubscription() (*models.SubscriptionRequest, error) {

	sub := models.SubscriptionRequest{}

	result := d.db.Where("active = false").First(&sub)

	if result.Error != nil || notFound(result.Error) {
		return nil, removeNotFound(result.Error)
	}

	return &sub, nil
}

func (d *Data) DeleteSubscription(sub *models.SubscriptionRequest) error {

	result := d.db.Delete(sub)

	return result.Error

}
