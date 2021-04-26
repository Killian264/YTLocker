package ytservice

import (
	"fmt"

	"golang.org/x/oauth2"
	"google.golang.org/api/youtube/v3"
)

type YTSerivceFake struct{}

// GetVideo gets a video with predetermined data |
// videoID: "fake-video-id" returns nil, nil
// videoID: "error-video-id" returns nil, error
func (s *YTSerivceFake) GetVideo(channelID string, videoID string) (*youtube.Video, error) {

	if videoID == "fake-video-id" {
		return nil, nil
	}

	if videoID == "error-video-id" {
		return nil, fmt.Errorf("error")
	}

	thumbnails := getThumbnails()

	return &youtube.Video{
		Id: videoID,
		Snippet: &youtube.VideoSnippet{
			Title:       "wow cool title",
			Description: "wow that is a super cool description",
			ChannelId:   channelID,
			Thumbnails:  &thumbnails,
		},
	}, nil

}

// GetChannel gets a channel with predetermined data |
// videoID: "fake-channel-id" returns nil, nil
// videoID: "error-channel-id" returns nil, error
func (s *YTSerivceFake) GetChannel(channelID string) (*youtube.Channel, error) {

	if channelID == "fake-channel-id" {
		return nil, nil
	}

	if channelID == "error-channel-id" {
		return nil, fmt.Errorf("error")
	}

	thumbnails := getThumbnails()

	return &youtube.Channel{
		Id: channelID,
		Snippet: &youtube.ChannelSnippet{
			Title:       "wow cool title",
			Description: "wow that is a super cool description",
			Thumbnails:  &thumbnails,
		},
	}, nil

}

type YTPlaylistFake struct {
	initalized bool
}

func (s *YTPlaylistFake) Initialize(config oauth2.Config, token oauth2.Token) error {
	s.initalized = true

	return nil
}
func (s *YTPlaylistFake) Create(title string, description string) (*youtube.Playlist, error) {

	if !s.initalized {
		panic("initialize must be ran on playlist")
	}

	thumbnails := getThumbnails()

	return &youtube.Playlist{
		Id: "simple-playlist-id",
		Snippet: &youtube.PlaylistSnippet{
			Title:       title,
			Description: description,
			Thumbnails:  &thumbnails,
		},
	}, nil

}
func (s *YTPlaylistFake) Insert(playlistID string, videoID string) error {

	if !s.initalized {
		panic("initialize must be ran on playlist")
	}

	return nil
}

func getThumbnails() youtube.ThumbnailDetails {
	return youtube.ThumbnailDetails{
		Default: &youtube.Thumbnail{
			Url:    "ytlocker.com",
			Height: 200,
			Width:  200,
		},
		Standard: &youtube.Thumbnail{
			Url:    "ytlocker.com",
			Height: 200,
			Width:  200,
		},
		Medium: &youtube.Thumbnail{
			Url:    "ytlocker.com",
			Height: 200,
			Width:  200,
		},
		High: &youtube.Thumbnail{
			Url:    "ytlocker.com",
			Height: 200,
			Width:  200,
		},
		Maxres: &youtube.Thumbnail{
			Url:    "ytlocker.com",
			Height: 200,
			Width:  200,
		},
	}
}
