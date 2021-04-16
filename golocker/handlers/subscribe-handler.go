package handlers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/Killian264/YTLocker/golocker/interfaces"
	"github.com/Killian264/YTLocker/golocker/models"
	"github.com/Killian264/YTLocker/golocker/parsers"
	"github.com/Killian264/YTLocker/golocker/services"
	"github.com/gorilla/mux"
)

// HandleSubscriptionNoError handles a new subscription request wrap in a middleware that handles errors
func HandleYoutubePush(w http.ResponseWriter, r *http.Request, s *services.Services) error {

	challenge := r.URL.Query().Get("hub.challenge")

	if challenge != "" {
		return handleChallenge(w, r, s.Subscribe)
	}

	return handleNewVideoPush(w, r, s.Subscribe)
}

func handleChallenge(w http.ResponseWriter, r *http.Request, s interfaces.ISubscription) error {

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

	isValid, err := s.HandleChallenge(&request)

	if err != nil {
		return err
	}

	if !isValid {
		return fmt.Errorf("Invalid")
	}

	fmt.Fprintf(w, challenge)
	return nil
}

func handleNewVideoPush(w http.ResponseWriter, r *http.Request, s interfaces.ISubscription) error {

	secret := mux.Vars(r)["secret"]
	bytes, err := ioutil.ReadAll(r.Body)

	if err != nil {
		return err
	}

	body := string(bytes)

	push, err := parsers.ParseYTHook(body)
	if err != nil {
		return err
	}

	return s.HandleVideoPush(&push, secret)
}
