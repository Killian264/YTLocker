package interfaces

import (
	"net/http"

	"github.com/Killian264/YTLocker/golocker/models"
)

//// SUBSCRIBE //////////////////////////////////////////////////////////

// ISubscriptionData Data Requirements
type ISubscriptionData interface {
	NewSubscription(request *models.SubscriptionRequest) error
	GetSubscription(channelID string) (*models.SubscriptionRequest, error)

	InactivateAllSubscriptions() error
	GetInactiveSubscription() (*models.SubscriptionRequest, error)
	DeleteSubscription(sub *models.SubscriptionRequest) error
}

type IYoutubeManager interface {
	CreateVideo(videoID string, channelID string) (*models.Video, error)
	GetChannel(channelID string) (*models.Channel, error)
}

// ISubscription Requirements
type ISubscription interface {
	SetYTPubSubUrl(url string)
	SetSubscribeUrl(base string, path string)

	HandleChallenge(request *models.SubscriptionRequest) (bool, error)
	HandleVideoPush(push *models.YTHookPush, secret string) error

	Subscribe(channelID string) (*models.SubscriptionRequest, error)
	GetSubscription(channelID string) (*models.SubscriptionRequest, error)

	ResubscribeAll() error
}

//// PLAYLIST //////////////////////////////////////////////////////////

// IPlaylistData Data Requirements
type IPlaylistHelperData interface {
	GetFirstYoutubeClientConfig() (*models.YoutubeClientConfig, error)
	GetFirstYoutubeToken() (*models.YoutubeToken, error)
}

// IPlaylistService Youtube Service Requirements
type IYoutubePlaylistService interface {
	Initalize(configData models.YoutubeClientConfig, tokenData models.YoutubeToken)
	Create(playlist models.Playlist) (models.Playlist, error)
	Insert(playlist models.Playlist, video models.Video) error
}

// IPlaylist Requirements
type IPlaylistHelper interface {
	Create(user *models.User, playlist *models.Playlist) (*models.Playlist, error)
	Insert(playlist *models.Playlist, video *models.Video) error
}

//// NEED WORK //////////////////////////////////////////////////////////

// IPlaylistManager
type IPlaylistManager interface {
	Create(user *models.User, playlist *models.Playlist) (*models.Playlist, error)
	Insert(playlist *models.Playlist, video *models.Video) error
	Subscribe(channel *models.Channel, playlist *models.Playlist)
	Unsubscribe(channel *models.Channel, playlist *models.Playlist)
}

type IChannel interface {
	GetOrCreateChannel(channelID string)
}

type IUserData interface {
	GetUserByEmail(email string) (*models.User, error)
	NewUser(user *models.User) error
	GetFirstUser() (*models.User, error)
}

type IUser interface {
	GetUserFromRequest(r *http.Request) (*models.User, error)
	RegisterUser(user *models.User) error
	ValidEmail(email string) (bool, error)
}
