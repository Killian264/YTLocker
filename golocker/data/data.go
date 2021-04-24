package data

import (
	"github.com/Killian264/YTLocker/golocker/models"
	"gorm.io/gorm"
)

type Data struct {
	db   *gorm.DB
	rand DataRand
}

func (d *Data) GetChannel(channelID string) (*models.Channel, error) {

	channel := models.Channel{
		YoutubeID: channelID,
	}

	result := d.db.Where(&channel).First(&channel)

	if result.Error != nil || notFound(result.Error) {
		return nil, removeNotFound(result.Error)
	}

	result = d.db.Where(models.Thumbnail{OwnerID: channel.ID, OwnerType: "channels"}).Find(&channel.Thumbnails)

	if result.Error != nil || notFound(result.Error) {
		return nil, removeNotFound(result.Error)
	}

	return &channel, nil
}

func (d *Data) NewChannel(channel *models.Channel) error {

	channel.ID = d.rand.ID()

	for _, thumbnail := range channel.Thumbnails {
		thumbnail.ID = d.rand.ID()
	}

	result := d.db.Create(&channel)

	return result.Error
}

func (d *Data) NewVideo(video *models.Video, channelID string) error {

	channel := models.Channel{}

	result := d.db.Where("channel_id = ?", channelID).First(&channel)

	if result.Error != nil {
		return result.Error
	}

	video.ID = d.rand.ID()
	video.ChannelID = channel.ID

	for _, thumbnail := range video.Thumbnails {
		thumbnail.ID = d.rand.ID()
	}

	result = d.db.Create(&video)

	return result.Error
}

func (d *Data) GetChannelFromYoutubeID(channelID string) (*models.Channel, error) {
	channel := models.Channel{
		YoutubeID: channelID,
	}

	result := d.db.Where(&channel).First(&channel)

	if result.Error != nil || notFound(result.Error) {
		return nil, removeNotFound(result.Error)
	}

	return &channel, nil
}
