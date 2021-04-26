package services

import (
	"log"
	"net/http"

	"github.com/Killian264/YTLocker/golocker/data"
	"github.com/Killian264/YTLocker/golocker/models"
	"github.com/gorilla/mux"
	"golang.org/x/oauth2"
	"google.golang.org/api/youtube/v3"
)

// Services to be injected into handlers and cron jobs
type Services struct {
	Router    *mux.Router
	Data      *data.Data
	Logger    *log.Logger
	Youtube   IYoutubeManager
	User      IUser
	Subscribe ISubscription
}

type ISubscription interface {
	SetYTPubSubUrl(url string)
	SetSubscribeUrl(base string, path string)

	Subscribe(channel *models.Channel) (*models.SubscriptionRequest, error)
	GetSubscription(channel *models.Channel) (*models.SubscriptionRequest, error)

	ResubscribeAll() error

	HandleChallenge(request *models.SubscriptionRequest, channelID string) (bool, error)
	HandleVideoPush(push *models.YTHookPush, secret string) error
}

type IYoutubeManager interface {
	NewVideo(channel *models.Channel, videoID string) (*models.Video, error)
	GetVideo(ID uint64) (*models.Video, error)
	GetVideoByID(youtubeID string) (*models.Video, error)

	NewChannel(channelID string) (*models.Channel, error)
	GetChannel(ID uint64) (*models.Channel, error)
	GetChannelByID(youtubeID string) (*models.Channel, error)
}

type IUser interface {
	GetUserFromRequest(r *http.Request) (*models.User, error)
	RegisterUser(user *models.User) error
	ValidEmail(email string) (bool, error)
	GetUserByID(ID uint64) (*models.User, error)
}

// Helper Services
type IYoutubeHelper interface {
	GetLastVideosFromChannel(channelID string, pageToken string) (*youtube.SearchListResponse, error)
	GetVideo(channelID string, videoID string) (*youtube.Video, error)
	GetChannel(channelID string) (*youtube.Channel, error)
}

type IYoutubePlaylistHelper interface {
	Initialize(config oauth2.Config, token oauth2.Token) error
	Create(title string, description string) (*youtube.Playlist, error)
	Insert(playlistID string, videoID string) error
}

// type IPlaylistHelperData interface {
// 	GetFirstYoutubeClientConfig() (*models.YoutubeClientConfig, error)
// 	GetFirstYoutubeToken() (*models.YoutubeToken, error)
// }

type IPlaylistManager interface {
	New(playlist *models.Playlist, user *models.User) (*models.Playlist, error)
	Get(ID uint64) (*models.Playlist, error)

	Insert(playlist *models.Playlist, video *models.Video) error

	Subscribe(playlist *models.Playlist, channel *models.Channel) error
	Unsubscribe(playlist *models.Playlist, channel *models.Channel) error
}
