package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Killian264/YTLocker/golocker/models"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
)

type Response struct {
	Status  int
	Data    interface{}
	Message string
}

// NewResponse creates a new api response
func NewResponse(status int, data interface{}, message string) Response {
	return Response{
		Status:  status,
		Data:    data,
		Message: message,
	}
}

// BlankResponse creates a successful or error response depending on err, error can be nil
func BlankResponse(err error) Response {
	status := http.StatusOK
	message := ""

	if err != nil {
		status = http.StatusInternalServerError
		message = err.Error()
	}

	return NewResponse(status, nil, message)
}

// GetUserFromRequest gets the user from the request, user is set by user authenticator
func GetUserFromRequest(r *http.Request) models.User {
	return context.Get(r, "user").(models.User)
}

// GetPlaylistFromRequest gets the playlist set from the request, playlist is set by playlistauthenticator
func GetPlaylistFromRequest(r *http.Request) models.Playlist {
	return context.Get(r, "playlist").(models.Playlist)
}

func GetUintFromRequest(r *http.Request, key string) (uint64, error) {
	idStr := mux.Vars(r)[key]
	if idStr == "" {
		return 0, fmt.Errorf("No value provided")
	}

	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("Value was invalid")
	}
	return id, nil
}