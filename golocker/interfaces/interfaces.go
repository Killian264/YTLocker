package interfaces

import (
	"net/http"

	"github.com/Killian264/YTLocker/golocker/models"
)

//// SUBSCRIBE //////////////////////////////////////////////////////////

// ISubscriptionData Data Requirements
type ISubscriptionData interface {
	NewSubscription(request *models.SubscriptionRequest) error
	GetSubscription(channelID uint64, secret string) (*models.SubscriptionRequest, error)

	InactivateAllSubscriptions() error
	GetInactiveSubscription() (*models.SubscriptionRequest, error)
	DeleteSubscription(sub *models.SubscriptionRequest) error
}

type IYoutubeManager interface {
	CreateVideo(channel *models.Channel, videoID string) (*models.Video, error)
	GetChannelByID(ID uint64) (*models.Channel, error)
	GetChannelByYoutubeID(youtubeID string) (*models.Channel, error)
}

// ISubscription Requirements
type ISubscription interface {
	SetYTPubSubUrl(url string)
	SetSubscribeUrl(base string, path string)

	GetSubscription(channel *models.Channel) (*models.SubscriptionRequest, error)

	Subscribe(channel *models.Channel) (*models.SubscriptionRequest, error)

	HandleChallenge(request *models.SubscriptionRequest) (bool, error)
	HandleVideoPush(push *models.YTHookPush, secret string) error

	ResubscribeAll() error
}

//// PLAYLIST //////////////////////////////////////////////////////////

type IPlaylistHelperData interface {
	GetFirstYoutubeClientConfig() (*models.YoutubeClientConfig, error)
	GetFirstYoutubeToken() (*models.YoutubeToken, error)
}

type IYoutubePlaylistService interface {
	Initalize(configData models.YoutubeClientConfig, tokenData models.YoutubeToken)
	Create(playlist models.Playlist) (models.Playlist, error)
	Insert(playlist models.Playlist, video models.Video) error
}

type IPlaylistHelper interface {
	Create(user *models.User, playlist *models.Playlist) (*models.Playlist, error)
	Insert(playlist *models.Playlist, video *models.Video) error
}

//// USER //////////////////////////////////////////////////////////

type IUserData interface {
	GetUserByEmail(email string) (*models.User, error)
	GetUserByID(ID uint64) (*models.User, error)
	NewUser(user *models.User) error
	GetFirstUser() (*models.User, error)
}

type IUser interface {
	GetUserFromRequest(r *http.Request) (*models.User, error)
	RegisterUser(user *models.User) error
	ValidEmail(email string) (bool, error)
	GetUserByID(ID uint64) (*models.User, error)
}

//// NEED WORK //////////////////////////////////////////////////////////

// IPlaylistManager
type IPlaylistManager interface {
	Create(playlist *models.Playlist, user *models.User) (*models.Playlist, error)
	Insert(playlist *models.Playlist, video *models.Video) error
	Subscribe(playlist *models.Playlist, channel *models.Channel)
	Unsubscribe(playlist *models.Playlist, channel *models.Channel)
}

type IChannel interface {
	GetOrCreateChannel(channelID string)
}
