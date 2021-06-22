package cronjobs

import (
	"log"
	"os"
	"testing"

	"github.com/Killian264/YTLocker/golocker/data"
	"github.com/Killian264/YTLocker/golocker/models"
	"github.com/Killian264/YTLocker/golocker/services"
	playlistservice "github.com/Killian264/YTLocker/golocker/services/playlist"
	userservice "github.com/Killian264/YTLocker/golocker/services/user"
	"github.com/Killian264/YTLocker/golocker/services/ytmanager"
	"github.com/stretchr/testify/assert"
)

var user = models.User{
	Username: "sdfakd",
	Email:    "sdkfsak",
	Password: "sladfk11111111",
}

var playlist = models.Playlist{
	Title:       "Hsdlakfjaskd",
	Description: "sdlfjaklsdf",
}

func Test_Run(t *testing.T) {
	s, j := createServices()

	user, _ := s.User.Register(user)

	playlist, _ = s.Playlist.New(playlist, user)
	channel, _ := s.Youtube.NewChannel("any-id-works")

	s.Youtube.NewVideo(channel, "any-id-works")
	s.Playlist.Subscribe(playlist, channel)

	err := j.Run()
	assert.Nil(t, err)

	videos, _ := s.Playlist.GetAllVideos(playlist)

	assert.Equal(t, 1, len(videos))
}

func createServices() (*services.Services, IJob) {
	data := data.InMemorySQLiteConnect()

	managerService := ytmanager.FakeNewYoutubeManager(data)
	playlistService := playlistservice.NewFakePlaylist(data)
	userService := userservice.NewUser(data)

	logger := log.New(os.Stdout, "Test: ", log.Lshortfile)

	service := &services.Services{
		Data:     data,
		Playlist: playlistService,
		Youtube:  managerService,
		User:     userService,
	}

	job := NewInsertVideosJob(service, logger)

	return service, job
}
