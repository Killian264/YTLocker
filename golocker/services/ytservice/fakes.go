package ytservice

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/Killian264/YTLocker/golocker/models"
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

func (s *YTSerivceFake) GetChannelIDByUsername(username string) (string, error) {
	if username == "fake-channel-username" {
		return "", nil
	}

	if username == "error-channel-username" {
		return "", fmt.Errorf("error")
	}

	return "jasdfkjasdklflkasdfk", nil
}

// GetLastVideosFromChannel fake impl, gets the last 25 videos from a channel AFTER some time
// pageToken is blank or a pagetoken given by response
func (s *YTSerivceFake) GetLastVideosFromChannel(channelID string, pageToken string, after time.Time) (*youtube.SearchListResponse, error) {
	thumbnails := getThumbnails()

	return &youtube.SearchListResponse{
		Items: []*youtube.SearchResult{
			{
				Id: &youtube.ResourceId{
					Kind:    "youtube#video",
					VideoId: "video-id-one",
				},
				Snippet: &youtube.SearchResultSnippet{
					ChannelId:   channelID,
					Title:       "Video Name 1",
					Description: "Video Description",
					Thumbnails:  &thumbnails,
				},
			},
			{
				Id: &youtube.ResourceId{
					Kind:    "youtube#video",
					VideoId: "video-id-two",
				},
				Snippet: &youtube.SearchResultSnippet{
					ChannelId:   channelID,
					Title:       "Video Name 1",
					Description: "Video Description",
					Thumbnails:  &thumbnails,
				},
			},
		},
	}, nil
}

type YTPlaylistFake struct {
	initalized bool
	email      string
	hashmap    map[string]string
}

func NewYTPlaylistFake() *YTPlaylistFake {
	baseAccessToken := "sa23.345234524623sdfasdfq-qegehgower9505034jfeworrjwertw_qqwerjfldssgert345sdgdgew-bheiyqeotleqjrljdfluao23423_QwekjfuI023kjasdfwer"

	hashmap := map[string]string{
		baseAccessToken: "dev-locker@ytlocker.com",
	}

	return &YTPlaylistFake{
		hashmap: hashmap,
	}
}

func (s *YTPlaylistFake) Initialize(config oauth2.Config, token oauth2.Token) (oauth2.Token, error) {
	s.initalized = true
	s.email = ""

	if token.AccessToken == "" || token.TokenType == "" {
		return oauth2.Token{}, fmt.Errorf("Invalid token")
	}

	email, emailFound := s.hashmap[token.AccessToken]

	if !emailFound {
		email = "killian@ytlocker.com" + fmt.Sprint(rand.Int())
		s.hashmap[token.AccessToken] = email
	}

	// is expired
	if token.Expiry.Before(time.Now()) {
		if token.RefreshToken == "" {
			return oauth2.Token{}, fmt.Errorf("Token was expired but did not have refresh token")
		}

		token.AccessToken = "2395034850394609158-6980349076032958-12358-4315=345814598-12" + fmt.Sprint(rand.Int())
		token.Expiry = time.Now().AddDate(0, 0, 1)

		s.hashmap[token.AccessToken] = email
	}

	s.email = email

	return token, nil
}

// TODO: Make actual fake implementation
func (s *YTPlaylistFake) GetUser() (models.OAuthUserInfo, error) {
	if !s.initalized {
		panic("initialize must be ran on playlist")
	}

	return models.OAuthUserInfo{
		Email:         s.email,
		VerifiedEmail: true,
		Picture:       "ytlocker.com" + fmt.Sprint(rand.Int()),
	}, nil
}

// TODO: Make actual fake implementation
func (s *YTPlaylistFake) GetChannel() (*youtube.Channel, error) {
	if !s.initalized {
		panic("initialize must be ran on playlist")
	}

	return &youtube.Channel{
		Snippet: &youtube.ChannelSnippet{
			Title: strings.SplitAfter(s.email, "@")[0],
			Thumbnails: &youtube.ThumbnailDetails{
				Standard: &youtube.Thumbnail{
					Url:    "ytlocker.com",
					Height: 200,
					Width:  200,
				},
			},
		},
	}, nil
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

func (s *YTPlaylistFake) GetPlaylistVideos(playlistId string) ([]string, error) {
	return []string{}, nil
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
