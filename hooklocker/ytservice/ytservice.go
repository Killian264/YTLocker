package ytservice

import (
	"log"
	"net/http"

	"google.golang.org/api/googleapi/transport"
	"google.golang.org/api/youtube/v3"
)

type YTService struct {
	youtubeService *youtube.Service
	channelService *youtube.ChannelsService
	videoService   *youtube.VideosService
	logger         *log.Logger
}

// InitializeServices creates the yt service
func NewYoutubeService(apiKey string, logger *log.Logger) *YTService {

	service := YTService{}

	youtubeClient := &http.Client{
		Transport: &transport.APIKey{Key: apiKey},
	}

	youtubeService, err := youtube.New(youtubeClient)

	if err != nil {
		panic("error creating youtube service")
	}

	service.youtubeService = youtubeService

	service.videoService = youtube.NewVideosService(youtubeService)

	service.channelService = youtube.NewChannelsService(youtubeService)

	service.logger = logger

	return &service
}

// GetVideoById gets a youtube video by id
func (s *YTService) GetVideo(videoID string) (*youtube.Video, error) {

	parts := []string{"snippet", "contentDetails"}
	call := s.videoService.List(parts)
	call.Id(videoID)

	response, err := call.Do()
	if err != nil {
		s.logger.Print("Youtube data api error: ", err)
		return nil, err
	}

	if response == nil || len(response.Items) == 0 {
		s.logger.Printf("Video with id: %s not found\n", videoID)
		return nil, nil
	}

	video := response.Items[0]

	return video, nil
}

// GetChannelById gets a youtube channel by id
func (s *YTService) GetChannel(channelID string) (*youtube.Channel, error) {

	parts := []string{"id", "snippet"}
	call := s.channelService.List(parts)
	call.Id(channelID)

	response, err := call.Do()
	if err != nil {
		s.logger.Print("youtube data api error: ", err)
		return nil, err
	}

	if response == nil || len(response.Items) == 0 {
		s.logger.Print("Channel with id: ", channelID, " not found\n")
		return nil, nil
	}

	channel := response.Items[0]

	return channel, nil
}
