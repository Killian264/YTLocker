package parsers

import (
	"encoding/json"
	"testing"

	"google.golang.org/api/youtube/v3"
)

var (
	ytChannelJSON = `
	{
		"etag": "EZNMEEQZ_13eVeGFXNTr0wa96aE",
		"id": "UC_x5XG1OV2P6uZZ5FSM9Ttw",
		"kind": "youtube#channel",
		"snippet": {
			"country": "US",
			"customUrl": "googledevelopers",
			"description": "The Google Developers channel features talks from events, educational series, best practices, tips, and the latest updates across our products and platforms.",
			"localized": {
				"description": "The Google Developers channel features talks from events, educational series, best practices, tips, and the latest updates across our products and platforms.",
				"title": "Google Developers"
			},
			"publishedAt": "2007-08-23T00:34:43Z",
			"thumbnails": {
				"default": {
					"height": 88,
					"url": "https://yt3.ggpht.com/ytc/AAUvwngOju7AKiAvKEs1wtsZN366tyNPyMq3nD8eFkMF7bE=s88-c-k-c0x00ffffff-no-rj",
					"width": 88
				},
				"high": {
					"height": 800,
					"url": "https://yt3.ggpht.com/ytc/AAUvwngOju7AKiAvKEs1wtsZN366tyNPyMq3nD8eFkMF7bE=s800-c-k-c0x00ffffff-no-rj",
					"width": 800
				},
				"medium": {
					"height": 240,
					"url": "https://yt3.ggpht.com/ytc/AAUvwngOju7AKiAvKEs1wtsZN366tyNPyMq3nD8eFkMF7bE=s240-c-k-c0x00ffffff-no-rj",
					"width": 240
				}
			},
			"title": "Google Developers"
		}
	}
	  
	`
)

func ytChannelCreate(ytJSON string) *youtube.Channel {

	ytChannel := new(youtube.Channel)

	json.Unmarshal([]byte(ytJSON), ytChannel)

	return ytChannel
}

func TestYtChannelJSONParse(t *testing.T) {

	ytChannel := ytChannelCreate(ytChannelJSON)

	if ytChannel == nil {
		t.Errorf("error parsing yt channel")
	}

	ytSnippet := ytChannel.Snippet

	ytName := ytSnippet.Title

	name := "Google Developers"

	if ytName != name {
		t.Errorf("malformed yt channel JSON parse expected: %q got %q", name, ytName)
	}

	ytLocalized := ytSnippet.Localized.Title

	if ytLocalized != name {
		t.Errorf("malformed yt channel JSON parse expected: %q got %q", name, ytLocalized)
	}

	url := "https://yt3.ggpht.com/ytc/AAUvwngOju7AKiAvKEs1wtsZN366tyNPyMq3nD8eFkMF7bE=s88-c-k-c0x00ffffff-no-rj"

	ytThumbnailUrl := ytSnippet.Thumbnails.Default.Url

	if ytThumbnailUrl != url {
		t.Errorf("malformed yt channel JSON parse expected: %q got %q", ytThumbnailUrl, url)
	}

}

func YtChannelParse(t *testing.T) {
	ytChannel := ytChannelCreate(ytChannelJSON)

	channel := ParseChannelIntoModels(ytChannel)

	ytSnippet := ytChannel.Snippet

	ytName := ytSnippet.Title
	name := channel.Title

	if ytName != name {
		t.Errorf("malformed yt channel parse expected: %q got %q", name, ytName)
	}

	ytThumbnail := ytSnippet.Thumbnails.Default
	ytThumbnailResolution := ytThumbnail.Height * ytThumbnail.Width

	for _, thumbnail := range channel.Thumbnails {
		resolution := thumbnail.Height * thumbnail.Width

		if resolution == int(ytThumbnailResolution) {
			if thumbnail.URL != ytThumbnail.Url {
				t.Errorf("malformed yt channel parse expected: %q got %q", ytThumbnail.Url, thumbnail.URL)
			}
			break
		}
	}

}
