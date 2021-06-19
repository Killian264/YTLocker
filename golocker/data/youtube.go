package data

import (
	"time"

	"github.com/Killian264/YTLocker/golocker/models"
	"gorm.io/gorm"
)

type Data struct {
	db   *gorm.DB
	rand DataRand
}

func (d *Data) NewChannel(channel models.Channel) (models.Channel, error) {

	channel.ID = d.rand.ID()

	for _, thumbnail := range channel.Thumbnails {
		thumbnail.ID = d.rand.ID()
	}

	result := d.db.Create(&channel)

	return channel, result.Error
}

func (d *Data) GetChannel(ID uint64) (models.Channel, error) {

	channel := models.Channel{}

	result := d.db.First(&channel, ID)

	if result.Error != nil || notFound(result.Error) {
		return models.Channel{}, removeNotFound(result.Error)
	}

	return channel, nil
}

func (d *Data) GetChannelByID(channelID string) (models.Channel, error) {

	channel := models.Channel{}

	result := d.db.Where("channels.youtube_id = ?", channelID).First(&channel)

	if result.Error != nil || notFound(result.Error) {
		return models.Channel{}, removeNotFound(result.Error)
	}

	return channel, nil
}

func (d *Data) NewVideo(channel models.Channel, video models.Video) (models.Video, error) {

	video.ID = d.rand.ID()
	video.ChannelID = channel.ID

	for _, thumbnail := range video.Thumbnails {
		thumbnail.ID = d.rand.ID()
	}

	result := d.db.Create(&video)

	return video, result.Error
}

func (d *Data) GetVideo(ID uint64) (models.Video, error) {

	video := models.Video{}

	result := d.db.First(&video, ID)

	if result.Error != nil || notFound(result.Error) {
		return models.Video{}, removeNotFound(result.Error)
	}

	return video, nil
}

func (d *Data) GetVideoByID(videoID string) (models.Video, error) {

	video := models.Video{}

	result := d.db.Where("videos.youtube_id = ?", videoID).First(&video)

	if result.Error != nil || notFound(result.Error) {
		return models.Video{}, removeNotFound(result.Error)
	}

	return video, nil
}

func (d *Data) GetVideosFromLast24Hours() ([]uint64, error) {
	videos := []OnlyID{}

	dayAgo := time.Now().AddDate(0, 0, -1)

	result := d.db.Model(models.Video{}).Where("updated_at > ?", dayAgo).Order("updated_at asc").Find(&videos)

	if result.Error != nil || notFound(result.Error) {
		return []uint64{}, removeNotFound(result.Error)
	}

	return parseOnlyIDArray(videos), nil;
}

func (d *Data) GetAllChannels() ([]uint64, error) {

	channels := []OnlyID{}

	result := d.db.Model(models.Channel{}).Find(&channels)

	if result.Error != nil || notFound(result.Error) {
		return []uint64{}, removeNotFound(result.Error)
	}

	return parseOnlyIDArray(channels), nil;

}

func (d *Data) GetAllChannelVideos(ID uint64) ([]uint64, error){
	videos := []OnlyID{}

	result := d.db.Raw(
		`SELECT
			V.id AS id 
		FROM channels AS C
		JOIN videos AS V
			ON C.id = V.channel_id
		WHERE C.id = ?
		;`, ID,
	).Scan(&videos);

	if removeNotFound(result.Error) != nil {
		return nil, result.Error
	}

	return parseOnlyIDArray(videos), nil;
}
