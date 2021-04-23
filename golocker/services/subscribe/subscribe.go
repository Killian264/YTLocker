package subscribe

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/Killian264/YTLocker/golocker/interfaces"
	"github.com/Killian264/YTLocker/golocker/models"
)

type Subscriber struct {
	pushSubscribeURL string
	pushHandlerURL   string
	ytmanager        interfaces.IYoutubeManager
	dataService      interfaces.ISubscriptionData
}

func NewSubscriber(data interfaces.ISubscriptionData, yt interfaces.IYoutubeManager) *Subscriber {
	return &Subscriber{
		dataService: data,
		ytmanager:   yt,
	}
}

func (s *Subscriber) SetYTPubSubUrl(url string) {
	s.pushSubscribeURL = url
}

func (s *Subscriber) SetSubscribeUrl(base string, path string) {

	if !strings.Contains(path, "{secret}") {
		panic("path must contain secret parameter")
	}

	base = strings.Trim(base, "/")
	path = strings.Trim(path, "/")

	s.pushHandlerURL = fmt.Sprintf("%s/%s/", base, path)
}

// Subscribe subscribes to a Subscription feed for a given channel
func (s *Subscriber) Subscribe(channelID string) (*models.SubscriptionRequest, error) {

	channel, err := s.ytmanager.GetChannel(channelID)
	if err != nil {
		return nil, err
	}
	if channel == nil {
		return nil, fmt.Errorf("Failed to find channel with id %s", channelID)
	}

	request, err := createSubscription(channelID)
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

func createSubscription(channelID string) (*models.SubscriptionRequest, error) {

	secret, err := generateSecret(64)
	if err != nil {
		return nil, err
	}

	return &models.SubscriptionRequest{
		ChannelID:    channelID,
		LeaseSeconds: uint(691200),
		Topic:        fmt.Sprintf("https://www.youtube.com/xml/feeds/videos.xml?channel_id=%s", channelID),
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

		_, err = s.Subscribe(old.ChannelID)
		if err != nil {
			return err
		}

		err = s.dataService.DeleteSubscription(old)
		if err != nil {
			return err
		}
	}

	return nil
}

// GetSubscription gets a subscription request
func (s *Subscriber) GetSubscription(channelID string) (*models.SubscriptionRequest, error) {
	return s.dataService.GetSubscription(channelID)
}

// HandleChallenge handles a challenge on a new subscription
func (s *Subscriber) HandleChallenge(request *models.SubscriptionRequest) (bool, error) {

	err := s.validSubscription(request.Secret, request.ChannelID)

	return err == nil, err
}

// HandleVideoPush handles a new video push from youtube
func (s *Subscriber) HandleVideoPush(push *models.YTHookPush, secret string) error {

	err := s.validSubscription(secret, push.Video.ChannelID)
	if err != nil {
		return err
	}

	_, err = s.ytmanager.CreateVideo(push.Video.VideoID, push.Video.ChannelID)
	if err != nil {
		return fmt.Errorf("Failed to save new video with video id: '%s' from channel: '%s'", push.Video.VideoID, push.Video.ChannelID)
	}

	return nil
}

func (s *Subscriber) validSubscription(secret string, channelID string) error {

	saved, err := s.dataService.GetSubscription(channelID)
	if err != nil {
		return err
	}

	if saved == nil {
		return fmt.Errorf("Failed to get subsciption with secret: '%s' and id: '%s'", secret, channelID)
	}

	if saved.Secret != secret {
		return fmt.Errorf("Invalid secret: '%s' and id: '%s'", secret, channelID)
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
