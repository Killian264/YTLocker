package data

import (
	"github.com/Killian264/YTLocker/golocker/models"
	"gorm.io/gorm"
)

type Data struct {
	db   *gorm.DB
	rand DataRand
}

func (d *Data) NewChannel(channel *models.Channel) error {

	channel.ID = d.rand.ID()

	for _, thumbnail := range channel.Thumbnails {
		thumbnail.ID = d.rand.ID()
	}

	result := d.db.Create(&channel)

	return result.Error
}

func (d *Data) GetChannel(ID uint64) (*models.Channel, error) {

	channel := models.Channel{
		ID: ID,
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

func (d *Data) GetChannelByID(channelID string) (*models.Channel, error) {

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

func (d *Data) NewVideo(channel *models.Channel, video *models.Video) error {

	video.ID = d.rand.ID()
	video.ChannelID = channel.ID

	for _, thumbnail := range video.Thumbnails {
		thumbnail.ID = d.rand.ID()
	}

	result := d.db.Create(&video)

	return result.Error
}

func (d *Data) GetVideo(ID uint64) (*models.Video, error) {

	video := models.Video{
		ID: ID,
	}

	result := d.db.Where(&video).First(&video)

	if result.Error != nil || notFound(result.Error) {
		return nil, removeNotFound(result.Error)
	}

	result = d.db.Where(models.Thumbnail{OwnerID: video.ID, OwnerType: "videos"}).Find(&video.Thumbnails)

	if result.Error != nil || notFound(result.Error) {
		return nil, removeNotFound(result.Error)
	}

	return &video, nil
}

func (d *Data) GetVideoByID(videoID string) (*models.Video, error) {

	video := models.Video{
		YoutubeID: videoID,
	}

	result := d.db.Where(&video).First(&video)

	if result.Error != nil || notFound(result.Error) {
		return nil, removeNotFound(result.Error)
	}

	result = d.db.Where(models.Thumbnail{OwnerID: video.ID, OwnerType: "videos"}).Find(&video.Thumbnails)

	if result.Error != nil || notFound(result.Error) {
		return nil, removeNotFound(result.Error)
	}

	return &video, nil
}
