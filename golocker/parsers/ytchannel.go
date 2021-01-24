package parsers

import (
	"github.com/Killian264/YTLocker/models"
	"google.golang.org/api/youtube/v3"
)

func ParseChannelIntoDBModels(channel *youtube.Channel) models.Channel {

	thumbnailDetails := channel.Snippet.Thumbnails

	high := thumbnailDetails.High
	medium := thumbnailDetails.Medium
	standard := thumbnailDetails.Standard

	thumbnails := []*youtube.Thumbnail{high, medium, standard}

	parsedThumbnails := []models.Thumbnail{}
	for i, thumbnail := range thumbnails {

		if thumbnail == nil {
			continue
		}

		parsedThumbnails[i] = models.Thumbnail{
			URL:    thumbnail.Url,
			Width:  int(thumbnail.Width),
			Height: int(thumbnail.Height),
		}
	}

	parsedChannel := models.Channel{
		ChannelID:   channel.Id,
		Title:       channel.Snippet.Title,
		Description: channel.Snippet.Description,
		Thumbnails:  parsedThumbnails,
	}

	return parsedChannel
}
