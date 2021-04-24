package ytmanager

import "github.com/Killian264/YTLocker/golocker/models"

type IYoutubeManager interface {
	CreateVideo(videoID string, channel *models.Channel) (*models.Video, error)
	GetVideoByID(ID uint64) (*models.Video, error)
	GetVideoByYoutubeID(youtubeID string) (*models.Video, error)

	CreateChannel(channelID string) (*models.Channel, error)
	GetChannelByID(ID uint64) (*models.Channel, error)
	GetChannelByYoutubeID(youtubeID string) (*models.Channel, error)
}
