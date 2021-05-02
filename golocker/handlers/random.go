package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Killian264/YTLocker/golocker/helpers/parsers"
	"github.com/Killian264/YTLocker/golocker/models"
	"github.com/Killian264/YTLocker/golocker/services"
)

// HandleSubscriptionNoError handles a new subscription request wrap in a middleware that handles errors
func CreatePlaylist(w http.ResponseWriter, r *http.Request, s *services.Services) error {

	user := GetUserFromRequest(r)

	playlist := models.Playlist{}

	err := json.NewDecoder(r.Body).Decode(&playlist)
	if err != nil {
		return err
	}

	playlist = parsers.ParsePlaylist(playlist)

	created, err := s.Playlist.New(&playlist, &user)
	if err != nil {
		return err

	}

	marshal, err := json.Marshal(created)
	if err != nil {
		return nil
	}

	w.Write(marshal)

	return err

}

// HandleSubscriptionNoError handles a new subscription request wrap in a middleware that handles errors
func TestHandler(w http.ResponseWriter, r *http.Request, s *services.Services) error {

	user := GetUserFromRequest(r)

	s.Logger.Print(user)

	return nil

}
