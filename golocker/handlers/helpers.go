package handlers

import (
	"net/http"

	"github.com/Killian264/YTLocker/golocker/models"
	"github.com/gorilla/context"
)

type Response struct {
	Redirect    bool
	RedirectUrl string
	Status      int
	Data        interface{}
	Message     string
}

// NewResponse creates a new api response
func NewResponse(status int, data interface{}, message string) Response {
	return Response{
		Redirect:    false,
		RedirectUrl: "",
		Status:      status,
		Data:        data,
		Message:     message,
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

func NewRedirectResponse(url string, message string) Response {
	return Response{
		Redirect:    true,
		RedirectUrl: url,
		Status:      http.StatusTemporaryRedirect,
		Data:        nil,
		Message:     message,
	}
}

// GetUserFromRequest gets the user from the request, user is set by user authenticator
func GetUserFromRequest(r *http.Request) models.User {
	return context.Get(r, "user").(models.User)
}

// GetPlaylistFromRequest gets the playlist set from the request, playlist is set by playlistauthenticator
func GetPlaylistFromRequest(r *http.Request) models.Playlist {
	return context.Get(r, "playlist").(models.Playlist)
}
