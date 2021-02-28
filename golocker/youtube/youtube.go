package youtube

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"google.golang.org/api/googleapi/transport"
	"google.golang.org/api/youtube/v3"
)

// YTService wraps the yt api for common commands to be used by the application
type YTService struct {
	youtubeService *youtube.Service
	channelService *youtube.ChannelsService
}

// InitializeServices creates the yt service
func (s *YTService) InitializeServices(apiKey string) {

	youtubeClient := &http.Client{
		Transport: &transport.APIKey{Key: apiKey},
	}

	youtubeService, err := youtube.New(youtubeClient)

	if err != nil {
		panic("error creating youtube service")
	}

	s.youtubeService = youtubeService

	channelService := youtube.NewChannelsService(youtubeService)

	s.channelService = channelService

}

// Before this is run there should be a check to see if a channel exists already with this id
// To do this the best course of action would be to make a interface for the db object and begin building a db layer
func (s *YTService) GetChannelById(id string) *youtube.Channel {
	log.Print("Searching for channel with ID: ", id, "\n")

	parts := []string{"id", "snippet"}
	call := s.channelService.List(parts)
	call.Id(id)

	response, err := call.Do()
	if err != nil {
		log.Print("youtube data api error: ", err)
	}

	if response == nil || len(response.Items) == 0 {
		log.Print("Channel with id: ", id, " not found\n")
		return nil
	}

	channel := response.Items[0]

	log.Print("Found channel with ID: ", id, "\n")
	log.Print("Channel Title: ", channel.Snippet.Title, "\n")

	return channel
}

// response2 := playlistsList(s.youtubeService, "snippet,contentDetails", id)

// for _, playlist := range response2.Items {
// 	playlistId := playlist.Id
// 	playlistTitle := playlist.Snippet.Title

// 	// Print the playlist ID and title for the playlist resource.
// 	fmt.Println(playlistId, ": ", playlistTitle)
// }
func playlistsList(service *youtube.Service, part string, channelId string) *youtube.PlaylistListResponse {
	call := service.Playlists.List(strings.Split(part, ","))
	if channelId != "" {
		call = call.ChannelId(channelId)
	}
	call = call.MaxResults(2)
	response, err := call.Do()
	if err != nil {
		fmt.Print("\n\nyoutube data api error\n\n", err)
	}
	return response
}
