package ytmanager

import (
	"fmt"
	"log"
	"time"

	"github.com/Killian264/YTLocker/golocker/helpers/parsers"
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
	GetLastVideosFromChannel(channelID string, pageToken string, after time.Time) (*youtube.SearchListResponse, error)
}

type IYoutubeManagerData interface {
	NewChannel(channel *models.Channel) error
	GetChannel(ID uint64) (*models.Channel, error)
	GetChannelByID(channelID string) (*models.Channel, error)

	NewVideo(channel *models.Channel, video *models.Video) error
	GetVideo(ID uint64) (*models.Video, error)
	GetVideoByID(videoID string) (*models.Video, error)

	GetVideosFromLast24Hours() (*[]models.Video, error)

	GetAllChannels() (*[]models.Channel, error)
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

// NewVideo fetches and saves to the db a video from a saved channel with a videoID
func (s *YoutubeManager) NewVideo(channel *models.Channel, videoID string) (*models.Video, error) {

	saved, _ := s.GetVideoByID(videoID)
	if saved != nil {
		return saved, nil
	}

	ytVideo, err := s.yt.GetVideo(channel.YoutubeID, videoID)
	if err != nil {
		return nil, err
	}
	if ytVideo == nil {
		return nil, fmt.Errorf("Video does not exist")
	}

	video, channelID := parsers.ParseYTVideo(ytVideo)
	if channelID != channel.YoutubeID {
		return nil, fmt.Errorf("Wrong channel provided for video")
	}

	err = s.data.NewVideo(channel, &video)
	if err != nil {
		return nil, err
	}

	return &video, nil
}

// GetVideo gets a video from the db
func (s *YoutubeManager) GetVideo(ID uint64) (*models.Video, error) {
	return s.data.GetVideo(ID)
}

// GetVideoByID gets a video from the db
func (s *YoutubeManager) GetVideoByID(youtubeID string) (*models.Video, error) {
	return s.data.GetVideoByID(youtubeID)
}

// NewChannel gets and saves to the db a new channel
func (s *YoutubeManager) NewChannel(channelID string) (*models.Channel, error) {

	saved, _ := s.GetChannelByID(channelID)
	if saved != nil {
		return saved, nil
	}

	ytChannel, err := s.yt.GetChannel(channelID)
	if err != nil {
		return nil, err
	}
	if ytChannel == nil {
		return nil, fmt.Errorf("Channel does not exist")
	}

	channel := parsers.ParseYTChannel(ytChannel)

	err = s.data.NewChannel(&channel)
	if err != nil {
		return nil, err
	}

	return &channel, nil
}

// GetChannel gets a channel from the db
func (s *YoutubeManager) GetChannel(ID uint64) (*models.Channel, error) {
	return s.data.GetChannel(ID)
}

// GetChannelByID gets a channel from the db
func (s *YoutubeManager) GetChannelByID(youtubeID string) (*models.Channel, error) {
	return s.data.GetChannelByID(youtubeID)
}

func (s *YoutubeManager) GetAllVideosFromLast24Hours() (*[]models.Video, error) {
	return s.data.GetVideosFromLast24Hours()
}

func (s *YoutubeManager) CheckForMissedUploads(l *log.Logger) error {

	channels, err := s.data.GetAllChannels()
	if err != nil {
		return err
	}

	after := time.Now().AddDate(0, 0, -1)

	for _, channel := range *channels {

		response, err := s.yt.GetLastVideosFromChannel(channel.YoutubeID, "", after)
		if err != nil {
			return err
		}

		videos := parsers.ParseSearchResponseIntoVideos(response)

		for _, video := range videos {

			saved, err := s.GetVideoByID(video.YoutubeID)
			if err == nil && saved != nil {
				continue
			}

			if err != nil {
				l.Printf("MissedUploads: Error getting video %v", err)
			}

			err = s.data.NewVideo(&channel, &video)
			if err != nil {
				l.Printf("MissedUploads: Error processing video: %v", err)
			}

		}

	}

	return nil

}
