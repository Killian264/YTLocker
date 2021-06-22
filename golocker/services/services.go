package services

import (
	"log"
	"os"

	"github.com/Killian264/YTLocker/golocker/data"
	"github.com/Killian264/YTLocker/golocker/services/playlist"
	"github.com/Killian264/YTLocker/golocker/services/subscribe"
	userservice "github.com/Killian264/YTLocker/golocker/services/user"
	"github.com/Killian264/YTLocker/golocker/services/ytmanager"
	"github.com/gorilla/mux"
)

// Services to be injected into handlers and cron jobs
type Services struct {
	Router *mux.Router
	Data   *data.Data
	Logger *log.Logger

	Youtube   *ytmanager.YoutubeManager
	User      *userservice.User
	Subscribe *subscribe.Subscriber
	Playlist  *playlist.PlaylistManager
}

func NewMockServices() *Services {
	data := data.InMemorySQLiteConnect()

	managerService := ytmanager.FakeNewYoutubeManager(data)
	playlistService := playlist.NewFakePlaylist(data)
	userService := userservice.NewUser(data)

	logger := log.New(os.Stdout, "Test: ", log.Lshortfile)

	service := &Services{
		Logger:   logger,
		Data:     data,
		Playlist: playlistService,
		Youtube:  managerService,
		User:     userService,
	}

	return service
}
