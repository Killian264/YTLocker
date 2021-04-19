package ytservice

import (
	"context"
	"fmt"

	"github.com/Killian264/YTLocker/golocker/helpers/parsers"
	"github.com/Killian264/YTLocker/golocker/models"
	"google.golang.org/api/youtube/v3"
)

type YTPlaylist struct {
	playlist *youtube.PlaylistsService
	items    *youtube.PlaylistItemsService
}

func (s *YTPlaylist) Initalize(configData models.YoutubeClientConfig, tokenData models.YoutubeToken) error {

	config := parsers.ParseYoutubeClient(configData)
	token := parsers.ParseYoutubeToken(tokenData)

	client := config.Client(context.Background(), &token)
	service, err := youtube.New(client)
	if err != nil {
		return err
	}

	playlist := youtube.NewPlaylistsService(service)
	items := youtube.NewPlaylistItemsService(service)

	s.playlist = playlist
	s.items = items

	return nil
}

func (s *YTPlaylist) Create(playlist models.Playlist) (models.Playlist, error) {

	parts := []string{"id", "snippet", "status"}

	ytPlaylist := youtube.Playlist{
		Snippet: &youtube.PlaylistSnippet{
			Title:       playlist.Title,
			Description: fmt.Sprint(playlist.Description, "\n\n", "Auto-generated by YTLocker."),
		},
		Status: &youtube.PlaylistStatus{
			PrivacyStatus: "unlisted",
		},
	}

	call := s.playlist.Insert(parts, &ytPlaylist)
	response, err := call.Do()
	if err != nil {
		return playlist, err
	}

	playlist.PlaylistID = response.Id

	return playlist, err
}

func (s *YTPlaylist) Insert(playlist models.Playlist, video models.Video) error {

	parts := []string{"id", "snippet"}

	item := &youtube.PlaylistItem{
		Id: video.VideoID,
		Snippet: &youtube.PlaylistItemSnippet{
			PlaylistId: playlist.PlaylistID,
			ResourceId: &youtube.ResourceId{
				Kind:    "youtube#video",
				VideoId: video.VideoID,
			},
		},
	}

	call := s.items.Insert(parts, item)

	_, err := call.Do()
	if err != nil {
		return err
	}

	return err
}
