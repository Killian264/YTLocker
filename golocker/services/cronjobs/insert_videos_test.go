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

var playlist = &models.Playlist{
	Title:       "Hsdlakfjaskd",
	Description: "sdlfjaklsdf",
}

func Test_Run(t *testing.T) {

	s, j := createServices()

	user, err := s.User.Register(user)
	assert.Nil(t, err)

	channel, err := s.Youtube.NewChannel("any-id-works")
	assert.Nil(t, err)

	_, err = s.Youtube.NewVideo(channel, "any-id-works")
	assert.Nil(t, err)

	playlist, err = s.Playlist.New(playlist, &user)
	assert.Nil(t, err)

	s.Playlist.Subscribe(playlist, channel)

	err = j.Run()
	assert.Nil(t, err)

	playlist, err = s.Playlist.Get(&user, playlist.ID)
	assert.Nil(t, err)

	assert.NotEqual(t, 2, len(playlist.Videos))

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
