package subscribe

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"

	"github.com/Killian264/YTLocker/hooklocker/interfaces"
	"github.com/Killian264/YTLocker/hooklocker/models"
	"github.com/Killian264/YTLocker/hooklocker/parsers"
	"github.com/gorilla/mux"
)

// var YoutubeSubscribeUrl = "https://pubsubhubbub.appspot.com/subscribe"

type Subscriber struct {
	pushSubscribeURL string
	pushHandlerURL   string
	ytService        interfaces.IYoutubeService
	dataService      interfaces.ISubscriptionData
	logger           *log.Logger
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
func (s *Subscriber) CreateSubscription(channelID string) (models.SubscriptionRequest, error) {
	leaseSeconds := 691200

	topic := fmt.Sprintf("https://www.youtube.com/xml/feeds/videos.xml?channel_id=%s", channelID)

	secret, err := generateSecret(64)

	if err != nil {
		return models.SubscriptionRequest{}, err
	}

	request := models.SubscriptionRequest{
		ChannelID:    channelID,
		LeaseSeconds: leaseSeconds,
		Topic:        topic,
		Secret:       secret,
		Active:       true,
	}

	return request, nil
}

// Subscribe subscribes to a Subscription feed
func (s *Subscriber) Subscribe(request models.SubscriptionRequest) error {

	exists, err := s.dataService.ChannelExists(request.ChannelID)
	if err != nil {
		return err
	}

	if !exists {
		return fmt.Errorf("Failed to find channel with id %s", request.ChannelID)
	}

	err = s.postSubscription(request, s.pushSubscribeURL, s.pushHandlerURL)
	if err != nil {
		return err
	}

	return s.dataService.SaveSubscription(request)
}

func (s *Subscriber) postSubscription(request models.SubscriptionRequest, pushSubscribeURL string, pushHandlerURL string) error {
	callback := strings.Replace(pushHandlerURL, "{secret}", request.Secret, -1)

	resp, err := http.PostForm(pushSubscribeURL,
		url.Values{
			"hub.callback":      {callback},
			"hub.topic":         {request.Topic},
			"hub.verify":        {"async"},
			"hub.mode":          {"subscribe"},
			"hub.lease_seconds": {strconv.Itoa(request.LeaseSeconds)},
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

		if reflect.DeepEqual(old, models.SubscriptionRequest{}) {
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
			s.logger.Print("ERROR: Failed to delete subscription with id: ", old.ID)
			s.logger.Print(err)
			continue
		}

	}

	return nil
}

// HandleSubscriptionNoError handles a new subscription request wrap in a middleware that handles errors
func (s *Subscriber) HandleSubscription(w http.ResponseWriter, r *http.Request) error {

	challenge := r.URL.Query().Get("hub.challenge")

	if challenge != "" {
		s.logger.Print("Challenge Request Recieved\n")
		return s.handleChallenge(w, r)
	}

	s.logger.Print("New Video Request Recieved\n")
	return s.handleNewVideo(w, r)
}

func (s *Subscriber) handleChallenge(w http.ResponseWriter, r *http.Request) error {

	secret := mux.Vars(r)["secret"]
	challenge := r.URL.Query().Get("hub.challenge")
	topic := r.URL.Query().Get("hub.topic")
	channelID := strings.Replace(topic, "https://www.youtube.com/xml/feeds/videos.xml?channel_id=", "", 1)
	leaseStr := r.URL.Query().Get("hub.lease_seconds")

	lease, err := strconv.Atoi(leaseStr)

	if err != nil {
		return fmt.Errorf("Failed to parse lease_seconds got: %s", leaseStr)
	}

	request := models.SubscriptionRequest{
		ChannelID:    channelID,
		Topic:        topic,
		Secret:       secret,
		LeaseSeconds: lease,
		Active:       true,
	}

	isValid, err := s.subscriptionIsValid(request)
	if err != nil {
		return err
	}

	if isValid {
		fmt.Fprintf(w, challenge)
	}

	return nil
}

func (s *Subscriber) subscriptionIsValid(request models.SubscriptionRequest) (bool, error) {
	saved, err := s.dataService.GetSubscription(request.Secret, request.ChannelID)

	if err != nil {
		return false, fmt.Errorf("Failed to get subsciption with secret: '%s' and id: '%s'", request.Secret, request.ChannelID)
	}

	if saved.Secret != request.Secret || saved.ChannelID != request.ChannelID {
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

func (s *Subscriber) handleNewVideo(w http.ResponseWriter, r *http.Request) error {
	bytes, err := ioutil.ReadAll(r.Body)

	if err != nil {
		return err
	}

	body := string(bytes)

	push, err := parsers.ParseYTHook(body)
	if err != nil {
		return err
	}

	return s.videoPushed(push)
}

func (s *Subscriber) videoPushed(push models.YTHookPush) error {

	video, err := s.ytService.GetVideo(push.Video.VideoID)
	if err != nil {
		return err
	}

	if video != nil {
		return fmt.Errorf("Failed to get video with id: %s from channel: %s", push.Video.VideoID, push.Video.ChannelID)
	}

	err = s.dataService.SaveVideo(*video)

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
