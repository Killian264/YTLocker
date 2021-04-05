package interfaces

import (
	"net/http"

	"github.com/Killian264/YTLocker/golocker/models"
	"google.golang.org/api/youtube/v3"
)

// ISubscriptionData Service
type ISubscriptionData interface {
	NewSubscription(request *models.SubscriptionRequest) error
	GetSubscription(secret string, channelID string) (*models.SubscriptionRequest, error)
	GetChannel(channelID string) (*models.Channel, error)
	NewVideo(video models.Video, channelID string) error

	InactivateAllSubscriptions() error
	GetInactiveSubscription() (*models.SubscriptionRequest, error)
	DeleteSubscription(*models.SubscriptionRequest) error
}

// IYoutubeService Service
type IYoutubeService interface {
	GetVideo(videoID string) (*youtube.Video, error)
}

// ISubscription for readability only
type ISubscription interface {
	SetYTPubSubUrl(url string)
	SetSubscribeUrl(base string, path string)
	CreateSubscription(channelID string) (*models.SubscriptionRequest, error)
	Subscribe(request *models.SubscriptionRequest) error
	HandleSubscription(w http.ResponseWriter, r *http.Request) error
	ResubscribeAll() error
}
