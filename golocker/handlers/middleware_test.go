package handlers

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/Killian264/YTLocker/golocker/models"
	service "github.com/Killian264/YTLocker/golocker/services"
	"github.com/gorilla/context"
	"github.com/stretchr/testify/assert"
)

var user = models.User{
	Username: "Killian",
	Password: "askdfj23823qqqq",
	Email:    "killiandebacker@gmail.com",
}

var user2 = models.User{
	Username: "Killianqq",
	Password: "askdfj2d3823qqqq",
	Email:    "killiandebackwerer@gmail.com",
}

var playlist = models.Playlist{
	Title:       "wowee a cool playlist",
	Description: "this is a cool playlist",
}

var handleUser = models.User{}
var handlePlaylist = models.Playlist{}

var handler = func(w http.ResponseWriter, r *http.Request, s *service.Services) Response {
	handleUser = GetUserFromRequest(r)

	_, ok := context.GetOk(r, "playlist")
	if !ok {
		return BlankResponse(nil)
	}

	handlePlaylist = GetPlaylistFromRequest(r)
	return BlankResponse(nil)
}

func Test_User_Authenticator(t *testing.T) {

	services := service.NewMockServices()
	Authenticator := CreateUserAuthenticator(services)

	expected, _ := services.User.Register(user)
	bearer, _ := services.User.Login(user.Email, user.Password)

	req, _ := http.NewRequest("GET", "/user/information/", nil)
	req.Header["Authorization"] = []string{bearer}

	fake := FakeRequest{
		Services: services,
		Route:    "/user/information/",
		Request:  req,
		Handler:  Authenticator(handler),
	}

	res := SendFakeRequest(fake)
	assert.Equal(t, 200, res.StatusCode)

	actual := handleUser
	assert.Equal(t, expected.ID, actual.ID)

}

func Test_Playlist_Authenticator(t *testing.T) {

	s := service.NewMockServices()

	// user 1
	savedUser, _ := s.User.Register(user)
	bearer, _ := s.User.Login(user.Email, user.Password)
	playlist, _ := s.Playlist.New(&playlist, &savedUser)

	// user 2
	s.User.Register(user2)
	bearer2, _ := s.User.Login(user2.Email, user2.Password)

	// should get playlist with correct user
	Send_Authenticated_Playlist_Request(t, s, playlist, bearer)
	assert.Equal(t, playlist.ID, handlePlaylist.ID)

	handlePlaylist = models.Playlist{}

	// should not get playlist with invalid user
	Send_Authenticated_Playlist_Request(t, s, playlist, bearer2)
	assert.Empty(t, handlePlaylist)

}

func Send_Authenticated_Playlist_Request(t *testing.T, s *service.Services, playlist *models.Playlist, bearer string) {

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
