package handlers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Killian264/YTLocker/golocker/models"
	"github.com/Killian264/YTLocker/golocker/services"
	service "github.com/Killian264/YTLocker/golocker/services"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

var user = models.User{
	Username: "killian",
	Email:    "killiandebacker@gmail.com",
	Picture:  "https://lh3.googleusercontent.com/a/default-user=s96-c",
}

var user2 = models.User{
	Username: "killian",
	Email:    "killiandebacker2@gmail.com",
	Picture:  "https://lh3.googleusercontent.com/a/default-user=s96-c",
}

var playlist = models.Playlist{
	Title:       "wowee a cool playlist",
	Description: "this is a cool playlist",
}

var userGottenFromRequest = models.User{}
var playlistGottenFromRequest = models.Playlist{}

var handler = func(w http.ResponseWriter, r *http.Request, s *service.Services) Response {
	userGottenFromRequest = GetUserFromRequest(r)

	_, ok := context.GetOk(r, "playlist")
	if !ok {
		return BlankResponse(nil)
	}

	playlistGottenFromRequest = GetPlaylistFromRequest(r)
	return BlankResponse(nil)
}

func Test_User_Authenticator(t *testing.T) {
	s := service.NewMockServices()
	Authenticator := CreateUserAuthenticator(s)

	bearer, _ := s.User.GenerateTemporarySessionBearer()

	expected, _ := s.User.Login(user, bearer)

	req, _ := http.NewRequest("GET", "/user/information/", nil)
	req.Header["Authorization"] = []string{expected.Session.Bearer}

	fake := FakeRequest{
		Services: s,
		Route:    "/user/information/",
		Request:  req,
		Handler:  Authenticator(handler),
	}

	res := SendFakeRequest(fake)
	assert.Equal(t, 200, res.StatusCode)

	actual := userGottenFromRequest
	assert.Equal(t, expected.ID, actual.ID)
}

func Test_User_Authenticator_Fails(t *testing.T) {
	s := service.NewMockServices()
	Authenticator := CreateUserAuthenticator(s)

	bearer, _ := s.User.GenerateTemporarySessionBearer()

	s.User.Login(user, bearer)

	req, _ := http.NewRequest("GET", "/user/information/", nil)
	req.Header["Authorization"] = []string{"banans"}

	fake := FakeRequest{
		Services: s,
		Route:    "/user/information/",
		Request:  req,
		Handler:  Authenticator(handler),
	}

	res := SendFakeRequest(fake)
	assert.Equal(t, 401, res.StatusCode)

	assert.Equal(t, userGottenFromRequest, models.User{})
}

func Test_Playlist_Authenticator(t *testing.T) {
	s := service.NewMockServices()

	bearer, _ := s.User.GenerateTemporarySessionBearer()

	savedUser1, _ := s.User.Login(user, bearer)
	savedUser2, _ := s.User.Login(user2, bearer)

	playlist, _ := s.Playlist.New(playlist, savedUser1)

	// should get playlist with correct user
	Send_Authenticated_Playlist_Request(t, s, playlist, savedUser1.Session.Bearer)
	assert.Equal(t, playlist.ID, playlistGottenFromRequest.ID)

	playlistGottenFromRequest = models.Playlist{}

	// should not get playlist with invalid user
	Send_Authenticated_Playlist_Request(t, s, playlist, savedUser2.Session.Bearer)
	assert.Empty(t, playlistGottenFromRequest)
}

func Send_Authenticated_Playlist_Request(t *testing.T, s *service.Services, playlist models.Playlist, bearer string) {
	UserAuthenticator := CreateUserAuthenticator(s)
	PlaylistAuthenticator := CreatePlaylistAuthenticator(s)

	req, _ := http.NewRequest("GET", fmt.Sprintf("/playlist/%d/information/", playlist.ID), nil)
	req.Header["Authorization"] = []string{bearer}

	fake := FakeRequest{
		Services: s,
		Route:    "/playlist/{playlist_id}/information/",
		Request:  req,
		Handler:  UserAuthenticator(PlaylistAuthenticator(handler)),
	}

	SendFakeRequest(fake)
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

	userGottenFromRequest = models.User{}
	playlistGottenFromRequest = models.Playlist{}

	handler := func(w http.ResponseWriter, r *http.Request) {
		res := request.Handler(w, r, request.Services)

		w.WriteHeader(res.Status)
		fmt.Println("FakeRequest got: ", res.Message)
	}

	router.HandleFunc(request.Route, handler)
	router.ServeHTTP(rr, request.Request)
	return rr.Result()
}
