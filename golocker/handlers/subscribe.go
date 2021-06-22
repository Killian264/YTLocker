package handlers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/Killian264/YTLocker/golocker/helpers/parsers"
	"github.com/Killian264/YTLocker/golocker/models"
	"github.com/Killian264/YTLocker/golocker/services"
	"github.com/gorilla/mux"
)

// HandleSubscriptionNoError handles a new subscription request
func HandleYoutubePush(w http.ResponseWriter, r *http.Request, s *services.Services) Response {
	challenge := r.URL.Query().Get("hub.challenge")

	if challenge != "" {
		return handleChallenge(w, r, s)
	}

	return handleNewVideoPush(w, r, s)
}

func handleChallenge(w http.ResponseWriter, r *http.Request, s *services.Services) Response {
	secret := mux.Vars(r)["secret"]
	challenge := r.URL.Query().Get("hub.challenge")
	topic := r.URL.Query().Get("hub.topic")
	channelID := strings.Replace(topic, "https://www.youtube.com/xml/feeds/videos.xml?channel_id=", "", 1)
	leaseStr := r.URL.Query().Get("hub.lease_seconds")

	lease, err := strconv.Atoi(leaseStr)

	if err != nil {
		return BlankResponse(fmt.Errorf("Failed to parse lease_seconds got: %s", leaseStr))
	}

	request := models.SubscriptionRequest{
		Topic:        topic,
		Secret:       secret,
		LeaseSeconds: lease,
		Active:       true,
	}

	isValid, err := s.Subscribe.HandleChallenge(&request, channelID)

	if err != nil {
		return BlankResponse(err)
	}

	if !isValid {
		return BlankResponse(fmt.Errorf("Invalid"))
	}

	fmt.Fprintf(w, challenge)

	return BlankResponse(nil)
}

func handleNewVideoPush(w http.ResponseWriter, r *http.Request, s *services.Services) Response {
	secret := mux.Vars(r)["secret"]
	bytes, err := ioutil.ReadAll(r.Body)

	if err != nil {
		return BlankResponse(err)
	}

	body := string(bytes)

	push, err := parsers.ParseYTHook(body)
	if err != nil {
		return BlankResponse(err)
	}

	return BlankResponse(s.Subscribe.HandleVideoPush(&push, secret))
}
