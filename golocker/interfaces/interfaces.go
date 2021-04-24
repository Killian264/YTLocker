package interfaces

import (
	"net/http"

	"github.com/Killian264/YTLocker/golocker/models"
	"google.golang.org/api/youtube/v3"
)

type ISubscription interface {
	SetYTPubSubUrl(url string)
	SetSubscribeUrl(base string, path string)

	GetSubscription(channel *models.Channel) (*models.SubscriptionRequest, error)

	Subscribe(channel *models.Channel) (*models.SubscriptionRequest, error)

	HandleChallenge(request *models.SubscriptionRequest) (bool, error)
	HandleVideoPush(push *models.YTHookPush, secret string) error

	ResubscribeAll() error
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
	GetVideo(videoID string) (*youtube.Video, error)
	GetChannel(channelID string) (*youtube.Channel, error)
}

type IYoutubePlaylistHelper interface {
	Initalize(configData models.YoutubeClientConfig, tokenData models.YoutubeToken)
	Create(playlist models.Playlist) (models.Playlist, error)
	Insert(playlist models.Playlist, video models.Video) error
}

// type IPlaylistHelperData interface {
// 	GetFirstYoutubeClientConfig() (*models.YoutubeClientConfig, error)
// 	GetFirstYoutubeToken() (*models.YoutubeToken, error)
// }

// type IPlaylistManager interface {
// 	Create(playlist *models.Playlist, user *models.User) (*models.Playlist, error)
// 	Insert(playlist *models.Playlist, video *models.Video) error
// 	Subscribe(playlist *models.Playlist, channel *models.Channel)
// 	Unsubscribe(playlist *models.Playlist, channel *models.Channel)
// }
