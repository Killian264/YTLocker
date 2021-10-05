package services

import (
	"log"
	"os"

	"github.com/Killian264/YTLocker/golocker/data"
	"github.com/Killian264/YTLocker/golocker/services/cronjobs"
	"github.com/Killian264/YTLocker/golocker/services/oauthmanager"
	"github.com/Killian264/YTLocker/golocker/services/playlist"
	"github.com/Killian264/YTLocker/golocker/services/subscribe"
	"github.com/Killian264/YTLocker/golocker/services/user"
	"github.com/Killian264/YTLocker/golocker/services/ytmanager"
	"github.com/gorilla/mux"
)

type Config struct {
	WebBaseUrl     string
	WebLoginUrl    string
	WebRedirectUrl string
	EncryptionKey  string
}

// Services to be injected into handlers and cron jobs
type Services struct {
	Router *mux.Router
	Data   *data.Data
	Logger *log.Logger

	Youtube      *ytmanager.YoutubeManager
	User         *user.User
	Subscribe    *subscribe.Subscriber
	Playlist     *playlist.PlaylistManager
	OauthManager *oauthmanager.OauthManager
	Cronjob      *cronjobs.CronJobManager
	Config       Config
}

func NewMockServices() *Services {
	data := data.InMemorySQLiteConnect()

	managerService := ytmanager.FakeNewYoutubeManager(data)
	playlistService := playlist.NewFakePlaylist(data)
	userService := user.NewUser(data)

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
