package ytservice

import (
	"net/http"

	"google.golang.org/api/googleapi/transport"
	"google.golang.org/api/youtube/v3"
)

type YTService struct {
	youtubeService *youtube.Service
	channelService *youtube.ChannelsService
	videoService   *youtube.VideosService
	searchService  *youtube.SearchService
}

// NewYoutubeService creates the yt service
func NewYoutubeService(apiKey string) *YTService {

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

	service.searchService = youtube.NewSearchService(youtubeService)

	return &service
}

// GetLastVideosFromChannel gets the last 25 videos from a channel.
// pageToken is blank or a pagetoken given by response
func (s *YTService) GetLastVideosFromChannel(channelID string, pageToken string) (*youtube.SearchListResponse, error) {

	parts := []string{"snippet", "contentDetails"}
	call := s.searchService.List(parts)

	call.ChannelId(channelID)
	call.Order("date")
	call.MaxResults(25)

	if pageToken != "" {
		call.PageToken(pageToken)
	}

	response, err := call.Do()
	if err != nil {
		return nil, err
	}

	return response, nil

}

// GetVideo gets a youtube video by it's youtube channel id and youtube video id
func (s *YTService) GetVideo(channelID string, videoID string) (*youtube.Video, error) {

	parts := []string{"snippet", "contentDetails"}
	call := s.videoService.List(parts)
	call.Id(videoID)

	response, err := call.Do()
	if err != nil {
		return nil, err
	}

	if response == nil || len(response.Items) == 0 {
		return nil, nil
	}

	video := response.Items[0]

	return video, nil
}

// GetChannel gets a youtube channel by it's youtube id
func (s *YTService) GetChannel(channelID string) (*youtube.Channel, error) {

	parts := []string{"id", "snippet"}
	call := s.channelService.List(parts)
	call.Id(channelID)

	response, err := call.Do()
	if err != nil {
		return nil, err
	}

	if response == nil || len(response.Items) == 0 {
		return nil, nil
	}

	channel := response.Items[0]

	return channel, nil
}
