package data

import "github.com/Killian264/YTLocker/golocker/models"

func (d *Data) NewSubscriptionWorkUnit(workUnit *models.SubscriptionWorkUnit) error {
	workUnit.ID = d.rand.ID()

	result := d.db.Create(&workUnit)

	return result.Error
}

func (d *Data) GetFirstSubscriptionWorkUnitByStatus(status string) (*models.SubscriptionWorkUnit, error) {
	workUnit := &models.SubscriptionWorkUnit{}

	result := d.db.Where("status = ?", status).First(workUnit)

	if result.Error != nil || notFound(result.Error) {
		return nil, removeNotFound(result.Error)
	}

	return workUnit, nil
}

func (d *Data) UpdateSubscriptionWorkUnitStatus(workUnit *models.SubscriptionWorkUnit, status string) error {
	result := d.db.Model(&models.SubscriptionWorkUnit{}).Where("id = ?", workUnit.ID).Update("status", status)

	return result.Error
}

func (d *Data) GetSubscriptionWorkUnit(videoID uint64, channelID uint64) (*models.SubscriptionWorkUnit, error) {
	workUnit := &models.SubscriptionWorkUnit{
		VideoID:   videoID,
		ChannelID: channelID,
	}

	result := d.db.Where(workUnit).First(workUnit)

	if result.Error != nil || notFound(result.Error) {
		return nil, removeNotFound(result.Error)
	}

	return workUnit, nil
}
