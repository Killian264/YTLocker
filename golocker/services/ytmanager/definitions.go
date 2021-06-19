package ytmanager

import (
	"time"

	"github.com/Killian264/YTLocker/golocker/models"
	"github.com/Killian264/YTLocker/golocker/services/ytservice"
	"google.golang.org/api/youtube/v3"
)

type YoutubeManager struct {
	yt   IYTService
	data IYoutubeManagerData
}

type IYTService interface {
	GetVideo(channelID string, videoID string) (*youtube.Video, error)
	GetChannel(channelID string) (*youtube.Channel, error)
	GetChannelIDByUsername(channelID string) (string, error)
	GetLastVideosFromChannel(channelID string, pageToken string, after time.Time) (*youtube.SearchListResponse, error)
}

type IYoutubeManagerData interface {
	NewChannel(channel models.Channel) (models.Channel, error)
	GetChannel(ID uint64) (models.Channel, error)
	GetChannelByID(channelID string) (models.Channel, error)

	NewVideo(channel models.Channel, video models.Video) (models.Video, error)
	GetVideo(ID uint64) (models.Video, error)
	GetVideoByID(videoID string) (models.Video, error)

	GetVideosFromLast24Hours() ([]uint64, error)

	GetAllChannels() ([]uint64, error)

	GetAllChannelVideos(ID uint64) ([]uint64, error)

	GetThumbnails(ID uint64, ownerType string) ([]models.Thumbnail, error)
}

// NewYoutubeManager creates a new YoutubeManager and does any initilization work
func NewYoutubeManager(data IYoutubeManagerData, ytservice IYTService) *YoutubeManager {
	return &YoutubeManager{
		yt:   ytservice,
		data: data,
	}
}

func FakeNewYoutubeManager(data IYoutubeManagerData) *YoutubeManager {
	return NewYoutubeManager(data, &ytservice.YTSerivceFake{})
}