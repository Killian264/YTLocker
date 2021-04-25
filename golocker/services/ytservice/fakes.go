package ytservice

import (
	"fmt"

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
