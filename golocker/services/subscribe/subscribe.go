package subscribe

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/Killian264/YTLocker/golocker/models"
)

type Subscriber struct {
	pushSubscribeURL string
	pushHandlerURL   string
	ytmanager        IYoutubeManager
	dataService      ISubscriptionData
}

type ISubscriptionData interface {
	NewSubscription(request *models.SubscriptionRequest) error
	GetSubscription(channelID uint64, secret string) (*models.SubscriptionRequest, error)

	InactivateAllSubscriptions() error
	GetInactiveSubscription() (*models.SubscriptionRequest, error)
	DeleteSubscription(sub *models.SubscriptionRequest) error
}

type IYoutubeManager interface {
	NewVideo(channel *models.Channel, videoID string) (*models.Video, error)
	GetChannel(ID uint64) (*models.Channel, error)
	GetChannelByID(youtubeID string) (*models.Channel, error)
}

// NewSubscriber creates a new subscriber
func NewSubscriber(data ISubscriptionData, yt IYoutubeManager) *Subscriber {
	return &Subscriber{
		dataService: data,
		ytmanager:   yt,
	}
}

// SetYTPubSubUrl set the pub sub url tha this api will call to subscribe/unsubscribe
func (s *Subscriber) SetYTPubSubUrl(url string) {
	s.pushSubscribeURL = url
}

// SetSubscribeUrl set the route this application is using to handle challenges
func (s *Subscriber) SetSubscribeUrl(base string, path string) {

	if !strings.Contains(path, "{secret}") {
		panic("path must contain secret parameter")
	}

	base = strings.Trim(base, "/")
	path = strings.Trim(path, "/")

	s.pushHandlerURL = fmt.Sprintf("%s/%s/", base, path)
}

// Subscribe subscribes to a Subscription feed for a given channel
func (s *Subscriber) Subscribe(channel *models.Channel) (*models.SubscriptionRequest, error) {

	sub, err := s.GetSubscription(channel)
	if err != nil || sub != nil {
		return sub, err
	}

	request, err := createSubscription(channel)
	if err != nil {
		return nil, err
	}

	err = s.postSubscription(request, s.pushSubscribeURL, s.pushHandlerURL)
	if err != nil {
		return nil, err
	}

	err = s.dataService.NewSubscription(request)
	if err != nil {
		return nil, err
	}

	return request, nil
}

func createSubscription(channel *models.Channel) (*models.SubscriptionRequest, error) {

	secret, err := generateSecret(64)
	if err != nil {
		return nil, err
	}

	return &models.SubscriptionRequest{
		ChannelID:    channel.ID,
		LeaseSeconds: 691200,
		Topic:        fmt.Sprintf("https://www.youtube.com/xml/feeds/videos.xml?channel_id=%s", channel.YoutubeID),
		Secret:       secret,
		Active:       true,
	}, nil
}

func (s *Subscriber) postSubscription(request *models.SubscriptionRequest, pushSubscribeURL string, pushHandlerURL string) error {
	callback := strings.Replace(pushHandlerURL, "{secret}", request.Secret, -1)

	resp, err := http.PostForm(pushSubscribeURL,
		url.Values{
			"hub.callback":      {callback},
			"hub.topic":         {request.Topic},
			"hub.verify":        {"async"},
			"hub.mode":          {"subscribe"},
			"hub.lease_seconds": {strconv.Itoa(int(request.LeaseSeconds))},
		},
	)

	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("Failed to subscribe to channel status code: %s", resp.Status)
	}

	return nil
}

// ResubscribeAll resubscribes to all youtube hook subscriptions
func (s *Subscriber) ResubscribeAll() error {

	err := s.dataService.InactivateAllSubscriptions()
	if err != nil {
		return fmt.Errorf("Failed to inactivate all subscriptions got error %s", err.Error())
	}

	for true {

		old, err := s.dataService.GetInactiveSubscription()
		if err != nil || old == nil {
			return err
		}

		channel, err := s.ytmanager.GetChannel(old.ChannelID)
		if err != nil {
			return err
		}
		if channel == nil {
			return fmt.Errorf("Failed to find channel with id %s", channel.YoutubeID)
		}

		err = s.dataService.DeleteSubscription(old)
		if err != nil {
			return err
		}

		_, err = s.Subscribe(channel)
		if err != nil {
			return err
		}
	}

	return nil
}

// GetSubscription gets a subscription request
func (s *Subscriber) GetSubscription(channel *models.Channel) (*models.SubscriptionRequest, error) {
	return s.dataService.GetSubscription(channel.ID, "")
}

// HandleChallenge handles a challenge on a new subscription
func (s *Subscriber) HandleChallenge(request *models.SubscriptionRequest, channelID string) (bool, error) {

	channel, err := s.ytmanager.GetChannelByID(channelID)
	if err != nil {
		return false, err
	}
	if channel == nil {
		return false, fmt.Errorf("Failed to get channe with channelID: %s", channelID)
	}

	err = s.validSubscription(channel, request.Secret)

	return err == nil, err
}

// HandleVideoPush handles a new video push from youtube
func (s *Subscriber) HandleVideoPush(push *models.YTHookPush, secret string) error {

	channel, err := s.ytmanager.GetChannelByID(push.Video.ChannelID)
	if err != nil {
		return err
	}
	if channel == nil {
		return fmt.Errorf("Failed to get channe with channelID: %s", push.Video.ChannelID)
	}

	err = s.validSubscription(channel, secret)
	if err != nil {
		return err
	}

	_, err = s.ytmanager.NewVideo(channel, push.Video.VideoID)
	if err != nil {
		return fmt.Errorf("Failed to save new video with video id: '%s' from channel: '%s'", push.Video.VideoID, push.Video.ChannelID)
	}

	return nil
}

func (s *Subscriber) validSubscription(channel *models.Channel, secret string) error {

	saved, err := s.dataService.GetSubscription(channel.ID, secret)
	if err != nil {
		return err
	}

	if saved == nil {
		return fmt.Errorf("Failed to get subsciption with secret: '%s' and id: '%s'", secret, channel.YoutubeID)
	}

	if saved.Secret != secret {
		return fmt.Errorf("Invalid secret: '%s' and id: '%s'", secret, channel.YoutubeID)
	}

	return nil

}

func generateSecret(n int) (string, error) {
	h := sha256.New()
	b := make([]byte, 64)

	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	_, err = h.Write(b)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", b), nil
}
