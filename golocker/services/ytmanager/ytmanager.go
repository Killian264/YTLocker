package ytmanager

import (
	"fmt"

	"github.com/Killian264/YTLocker/golocker/helpers/parsers"
	"github.com/Killian264/YTLocker/golocker/models"
	"google.golang.org/api/youtube/v3"
)

type YoutubeManager struct {
	yt   IYTService
	data IYoutubeManagerData
}

type IYTService interface {
	GetVideo(channel *models.Channel, videoID string) (*youtube.Video, error)
	GetChannel(channelID string) (*youtube.Channel, error)
}

type IYoutubeManagerData interface {
	NewChannel(channel *models.Channel) error
	GetChannel(ID uint64) (*models.Channel, error)
	GetChannelByID(channelID string) (*models.Channel, error)

	NewVideo(channel *models.Channel, video *models.Video) error
	GetVideo(ID uint64) (*models.Video, error)
	GetVideoByID(videoID string) (*models.Video, error)
}

func NewYoutubeManager(data IYoutubeManagerData, ytservice IYTService) YoutubeManager {

	return YoutubeManager{
		yt:   ytservice,
		data: data,
	}

}
func (s *YoutubeManager) NewVideo(channel *models.Channel, videoID string) (*models.Video, error) {
	ytVideo, err := s.yt.GetVideo(channel, videoID)
	if err != nil || ytVideo == nil {
		return nil, err
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
func (s *YoutubeManager) GetVideo(ID uint64) (*models.Video, error) {
	return s.data.GetVideo(ID)
}
func (s *YoutubeManager) GetVideoByID(youtubeID string) (*models.Video, error) {
	return s.data.GetVideoByID(youtubeID)
}

func (s *YoutubeManager) NewChannel(channelID string) (*models.Channel, error) {

	ytChannel, err := s.yt.GetChannel(channelID)
	if err != nil || ytChannel == nil {
		return nil, err
	}

	channel := parsers.ParseYTChannel(ytChannel)

	err = s.data.NewChannel(&channel)
	if err != nil {
		return nil, err
	}

	return &channel, nil
}
func (s *YoutubeManager) GetChannel(ID uint64) (*models.Channel, error) {
	return s.data.GetChannel(ID)
}
func (s *YoutubeManager) GetChannelByID(youtubeID string) (*models.Channel, error) {
	return s.data.GetChannelByID(youtubeID)
}
