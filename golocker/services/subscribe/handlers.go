package subscribe

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/Killian264/YTLocker/golocker/models"
	"github.com/Killian264/YTLocker/golocker/parsers"
	"github.com/gorilla/mux"
)

// HandleSubscriptionNoError handles a new subscription request wrap in a middleware that handles errors
func (s *Subscriber) HandleSubscription(w http.ResponseWriter, r *http.Request) error {

	challenge := r.URL.Query().Get("hub.challenge")

	if challenge != "" {
		s.logger.Print("Challenge Request Recieved\n")
		return s.handleChallenge(w, r)
	}

	s.logger.Print("New Video Request Recieved\n")
	return s.handlePush(w, r)
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
		LeaseSeconds: uint(lease),
		Active:       true,
	}

	isValid, err := s.subscriptionIsValid(&request)
	if err == nil && isValid {
		fmt.Fprintf(w, challenge)
	}

	return err
}

func (s *Subscriber) handlePush(w http.ResponseWriter, r *http.Request) error {
	bytes, err := ioutil.ReadAll(r.Body)

	if err != nil {
		return err
	}

	body := string(bytes)

	push, err := parsers.ParseYTHook(body)
	if err != nil {
		return err
	}

	return s.videoPushed(&push)
}
