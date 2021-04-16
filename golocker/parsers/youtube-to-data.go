package parsers

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"time"

	"github.com/Killian264/YTLocker/golocker/models"
	"golang.org/x/oauth2"
	"google.golang.org/api/youtube/v3"
)

// ParseYTHook parses the xml data sent with YT Subscription webhooks
func ParseYTHook(hookXML string) (models.YTHookPush, error) {
	var hook models.YTHookPush
	err := xml.Unmarshal([]byte(hookXML), &hook)
	if err != nil {
		return models.YTHookPush{}, nil
	}
	return hook, nil
}

// ParseYTChannel parses a youtube.Channel into the db channel model
func ParseYTChannel(channel *youtube.Channel) models.Channel {
	thumbnails := ParseYTThumbnails(channel.Snippet.Thumbnails)

	parsed := models.Channel{
		ChannelID:   channel.Id,
		Title:       channel.Snippet.Title,
		Description: channel.Snippet.Description,

		Thumbnails: thumbnails,
	}

	return parsed
}

// ParseYTVideo parses a youtube.Video into the db video model
// returns parsed video, channelID
func ParseYTVideo(video *youtube.Video) (models.Video, string) {
	thumbnails := ParseYTThumbnails(video.Snippet.Thumbnails)

	parsed := models.Video{
		VideoID:     video.Id,
		Title:       video.Snippet.Title,
		Description: video.Snippet.Description,

		Thumbnails: thumbnails,
	}

	return parsed, video.Snippet.ChannelId
}

// ParseYTThumbnails parses a youtube.ThumbnailDetails into an array of db thumbnail models
func ParseYTThumbnails(details *youtube.ThumbnailDetails) []models.Thumbnail {
	ytThumbnails := []*youtube.Thumbnail{}

	ytThumbnails = append(ytThumbnails, details.Default)
	ytThumbnails = append(ytThumbnails, details.Standard)
	ytThumbnails = append(ytThumbnails, details.Medium)
	ytThumbnails = append(ytThumbnails, details.High)
	ytThumbnails = append(ytThumbnails, details.Maxres)

	thumbnails := []models.Thumbnail{}

	for _, thumbnail := range ytThumbnails {
		if thumbnail == nil {
			continue
		}
		thumbnails = append(thumbnails, models.Thumbnail{
			URL:    thumbnail.Url,
			Width:  uint(thumbnail.Width),
			Height: uint(thumbnail.Width),
		})
	}

	return thumbnails
}

// Heavily edited version of code originally found in google oauth2 ConfigFromJSON
func ParseClientJson(jsonData string) (models.YoutubeClientConfig, error) {
	type config struct {
		ClientID     string   `json:"client_id"`
		ClientSecret string   `json:"client_secret"`
		RedirectURIs []string `json:"redirect_uris"`
		AuthURI      string   `json:"auth_uri"`
		TokenURI     string   `json:"token_uri"`
	}

	var wrapper struct {
		Installed *config `json:"installed"`
	}

	if err := json.Unmarshal([]byte(jsonData), &wrapper); err != nil {
		return models.YoutubeClientConfig{}, err
	}

	if wrapper.Installed == nil {
		return models.YoutubeClientConfig{}, fmt.Errorf("oauth2/google: no credentials found")
	}

	c := wrapper.Installed

	if len(c.RedirectURIs) < 1 {
		return models.YoutubeClientConfig{}, errors.New("oauth2/google: missing redirect URL in the client_credentials.json")
	}

	return models.YoutubeClientConfig{
		ClientID:     c.ClientID,
		ClientSecret: c.ClientSecret,
		RedirectURL:  c.RedirectURIs[0],
		AuthURL:      c.AuthURI,
		TokenURL:     c.TokenURI,
	}, nil
}

func ParseAccessTokenJson(jsonData string) (models.YoutubeToken, error) {
	var token struct {
		AccessToken  string `json:"access_token"`
		TokenType    string `json:"token_type"`
		RefreshToken string `json:"refresh_token"`
		Expiry       string `json:"expiry"`
	}

	err := json.Unmarshal([]byte(jsonData), &token)
	if err != nil {
		return models.YoutubeToken{}, err
	}

	return models.YoutubeToken{
		AccessToken:  token.AccessToken,
		TokenType:    token.TokenType,
		RefreshToken: token.RefreshToken,
		Expiry:       token.Expiry,
	}, nil

}

func ParseYoutubeClient(config models.YoutubeClientConfig) oauth2.Config {
	return oauth2.Config{
		ClientID:     config.ClientID,
		ClientSecret: config.ClientSecret,
		RedirectURL:  config.RedirectURL,
		Scopes:       []string{config.Scope},
		Endpoint: oauth2.Endpoint{
			AuthURL:  config.AuthURL,
			TokenURL: config.TokenURL,
		},
	}
}

func ParseYoutubeToken(token models.YoutubeToken) oauth2.Token {

	expiry, err := time.Parse("2006-01-02T15:04:05.0000000-07:00", token.Expiry)

	if err != nil {
		expiry = time.Now()
	}

	return oauth2.Token{
		AccessToken:  token.AccessToken,
		TokenType:    token.TokenType,
		RefreshToken: token.RefreshToken,
		Expiry:       expiry,
	}
}
