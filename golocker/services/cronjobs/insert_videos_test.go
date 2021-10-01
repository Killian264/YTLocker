package cronjobs

import (
	"log"
	"os"
	"testing"

	"github.com/Killian264/YTLocker/golocker/data"
	"github.com/Killian264/YTLocker/golocker/models"
	"github.com/Killian264/YTLocker/golocker/services/playlist"
	"github.com/Killian264/YTLocker/golocker/services/user"
	"github.com/Killian264/YTLocker/golocker/services/ytmanager"
	"github.com/stretchr/testify/assert"
)

var testUserModel = models.User{
	Username: "killian",
	Email:    "killiandebacker@gmail.com",
	Picture:  "https://lh3.googleusercontent.com/a/default-user=s96-c",
}

var testPlaylistModel = models.Playlist{
	Title:       "Hsdlakfjaskd",
	Description: "sdlfjaklsdf",
}

func Test_Run(t *testing.T) {
	data := data.InMemorySQLiteConnect()

	managerService := ytmanager.FakeNewYoutubeManager(data)
	playlistService := playlist.NewFakePlaylist(data)
	userService := user.NewUser(data)

	job := NewInsertVideosJob(managerService, playlistService, data, log.New(os.Stdout, "Test: ", log.Lshortfile))

	user, _ := userService.Login(testUserModel)

	testPlaylist, _ := playlistService.New(testPlaylistModel, user)
	channel, _ := managerService.NewChannel("any-id-works")

	managerService.NewVideo(channel, "any-id-works")
	playlistService.Subscribe(testPlaylist, channel)

	err := job.Run()

	assert.Nil(t, err)

	videos, _ := playlistService.GetAllVideos(testPlaylist)

	assert.Equal(t, 1, len(videos))
}
