package subscribe

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/Killian264/YTLocker/golocker/helpers/parsers"
	"github.com/Killian264/YTLocker/golocker/interfaces"
	"github.com/Killian264/YTLocker/golocker/models"
)

type Subscriber struct {
	pushSubscribeURL string
	pushHandlerURL   string
	ytService        interfaces.IYoutubeService
	dataService      interfaces.ISubscriptionData
	logger           *log.Logger
}

func NewSubscriber(data interfaces.ISubscriptionData, yt interfaces.IYoutubeService, logger *log.Logger) *Subscriber {
	return &Subscriber{
		dataService: data,
		ytService:   yt,
		logger:      logger,
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

// CreateSubscription creates a new channel subscription to a channel feed for the channel with id channelId
func (s *Subscriber) CreateSubscription(channelID string) (*models.SubscriptionRequest, error) {
	leaseSeconds := 691200

	topic := fmt.Sprintf("https://www.youtube.com/xml/feeds/videos.xml?channel_id=%s", channelID)

	secret, err := generateSecret(64)

	if err != nil {
		return nil, err
	}

	request := models.SubscriptionRequest{
		ChannelID:    channelID,
		LeaseSeconds: uint(leaseSeconds),
		Topic:        topic,
		Secret:       secret,
		Active:       true,
	}

	return &request, nil
}

// Subscribe subscribes to a Subscription feed
func (s *Subscriber) Subscribe(request *models.SubscriptionRequest) error {

	channel, err := s.dataService.GetChannel(request.ChannelID)
	if err != nil {
		return err
	}

	if channel == nil {
		return fmt.Errorf("Failed to find channel with id %s", request.ChannelID)
	}

	err = s.postSubscription(request, s.pushSubscribeURL, s.pushHandlerURL)
	if err != nil {
		return err
	}

	return s.dataService.NewSubscription(request)
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

	run := true

	for run {

		old, err := s.dataService.GetInactiveSubscription()
		if err != nil {
			s.logger.Print("ERROR: Failed to get an inactive subscription")
			s.logger.Print(err)
			continue
		}

		if old == nil {
			run = false
			continue
		}

		new, err := s.CreateSubscription(old.ChannelID)
		if err != nil {
			s.logger.Print("ERROR: Failed to create a new subscription for channel id: ", old.ChannelID)
			s.logger.Print(err)
			continue
		}

		err = s.Subscribe(new)
		if err != nil {
			s.logger.Print("ERROR: Failed to subscribe with channel id: ", old.ChannelID)
			s.logger.Print(err)
			continue
		}

		err = s.dataService.DeleteSubscription(old)
		if err != nil {
			s.logger.Print("ERROR: Failed to delete subscription with uuid: ", old.UUID)
			s.logger.Print(err)
			continue
		}

	}

	return nil
}

func (s *Subscriber) HandleChallenge(request *models.SubscriptionRequest) (bool, error) {
	saved, err := s.dataService.GetSubscription(request.Secret, request.ChannelID)

	if err != nil {
		return false, fmt.Errorf("Failed to get subsciption with secret: '%s' and id: '%s'", request.Secret, request.ChannelID)
	}

	if saved == nil {
		return false, fmt.Errorf("Invalid secret or channel id: '%s' and id: '%s'", request.Secret, request.ChannelID)
	}

	if !saved.Active {
		log.Printf("Subscriber: Warning using inactive subscription")
	}

	if request.LeaseSeconds != saved.LeaseSeconds {
		log.Printf("Subscriber: Warning lease seconds do not match ")
	}

	return true, nil
}

func (s *Subscriber) HandleVideoPush(push *models.YTHookPush, secret string) error {

	saved, err := s.dataService.GetSubscription(secret, push.Video.ChannelID)

	if err != nil || saved == nil {
		return fmt.Errorf("Failed to get subsciption with secret: '%s' and id: '%s'", secret, push.Video.ChannelID)
	}

	video, err := s.ytService.GetVideo(push.Video.VideoID)
	if err != nil {
		return err
	}

	if video == nil {
		return fmt.Errorf("Failed to get video with id: %s from channel: %s", push.Video.VideoID, push.Video.ChannelID)
	}

	parsedVideo, channelID := parsers.ParseYTVideo(video)

	err = s.dataService.NewVideo(&parsedVideo, channelID)

	if err != nil {
		return fmt.Errorf("Failed to save new video with video id: '%s' from channel: '%s'", push.Video.VideoID, push.Video.ChannelID)
	}

	return nil
}

func generateSecret(n int) (string, error) {
	b := make([]byte, 64)
	_, err := rand.Read(b)

	h := sha256.New()
	h.Write(b)

	return fmt.Sprintf("%x", h.Sum(nil)), err
}
