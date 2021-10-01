package ytservice

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/Killian264/YTLocker/golocker/models"
	"golang.org/x/oauth2"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

type YTPlaylist struct {
	playlists *youtube.PlaylistsService
	items     *youtube.PlaylistItemsService
	channels  *youtube.ChannelsService
	token     oauth2.Token
}

// Initalize sets the oauth and token information for the next requests.
// Config data must match the playlist that is being inserted into
// Token data should be app level information
func (s *YTPlaylist) Initialize(config oauth2.Config, token oauth2.Token) error {
	client := config.Client(context.Background(), &token)

	service, err := youtube.NewService(oauth2.NoContext, option.WithHTTPClient(client))
	if err != nil {
		return err
	}

	s.playlists = youtube.NewPlaylistsService(service)
	s.items = youtube.NewPlaylistItemsService(service)
	s.channels = youtube.NewChannelsService(service)
	s.token = token

	return nil
}

// GetUser gets the user information
func (s *YTPlaylist) GetUser() (models.OAuthUserInfo, error) {
	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + url.QueryEscape(s.token.AccessToken))
	if err != nil {
		return models.OAuthUserInfo{}, fmt.Errorf("failed to get additional user information: " + err.Error())
	}
	defer resp.Body.Close()

	response, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return models.OAuthUserInfo{}, fmt.Errorf("failed to get read additional user information: " + err.Error())
	}

	details := models.OAuthUserInfo{}

	err = json.Unmarshal(response, &details)
	if err != nil {
		err = fmt.Errorf("failed to parse user information: " + err.Error())
	}

	return details, err
}

// GetChannel gets the channel of the user
func (s *YTPlaylist) GetChannel() (*youtube.Channel, error) {
	parts := []string{"snippet", "contentDetails", "statistics"}

	call := s.channels.List(parts)
	call.Mine(true)
	response, err := call.Do()

	if err != nil {
		return nil, err
	}

	return response.Items[0], err
}

// Create creates a playlist
func (s *YTPlaylist) Create(title string, description string) (*youtube.Playlist, error) {
	parts := []string{"id", "snippet", "status"}

	ytPlaylist := &youtube.Playlist{
		Snippet: &youtube.PlaylistSnippet{
			Title:       title,
			Description: fmt.Sprint(description, "\n\n", "Auto-generated by YTLocker."),
		},
		Status: &youtube.PlaylistStatus{
			PrivacyStatus: "unlisted",
		},
	}

	call := s.playlists.Insert(parts, ytPlaylist)
	response, err := call.Do()

	if err != nil {
		return nil, err
	}

	return response, err
}

// Insert inserts a video into a given playlist
func (s *YTPlaylist) Insert(playlistID string, videoID string) error {
	parts := []string{"id", "snippet"}

	item := &youtube.PlaylistItem{
		Id: videoID,
		Snippet: &youtube.PlaylistItemSnippet{
			PlaylistId: playlistID,
			ResourceId: &youtube.ResourceId{
				Kind:    "youtube#video",
				VideoId: videoID,
			},
			// Position: 0,
			// ForceSendFields: []string{"Position"},
		},
	}

	call := s.items.Insert(parts, item)

	_, err := call.Do()
	if err != nil {
		return err
	}

	return err
}
