package handlers

import (
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/Killian264/YTLocker/golocker/models"
	"github.com/Killian264/YTLocker/golocker/services"
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

type FakeRequest struct {
	Services *services.Services
	Request  *http.Request
	Route    string
	Handler  ServiceHandler
}

// SendFakeRequest sends a fake request the request will return 500 status code on error
func SendFakeRequest(request FakeRequest) *http.Response {
	rr := httptest.NewRecorder()
	router := mux.NewRouter()

	handler := func(w http.ResponseWriter, r *http.Request) {
		res := request.Handler(w, r, request.Services)

		w.WriteHeader(res.Status)
		fmt.Println("FakeRequest got: ", res.Message)
	}

	router.HandleFunc(request.Route, handler)
	router.ServeHTTP(rr, request.Request)
	return rr.Result()
}
