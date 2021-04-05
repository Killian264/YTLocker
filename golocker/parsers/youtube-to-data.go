package parsers

import (
	"encoding/xml"

	"github.com/Killian264/YTLocker/golocker/models"
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
