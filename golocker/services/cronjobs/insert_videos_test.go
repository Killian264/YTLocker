package cronjobs

import (
	"log"
	"os"
	"testing"
	"time"

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
	CreatedAt:   time.Now().AddDate(-1, 1, 1),
	Active:      true,
}

var testAccountModel = models.YoutubeAccount{
	ID:              12342,
	Username:        "asdfsadf",
	Email:           "asdfasdfsdaf@cool.com",
	Picture:         "asdjfasdf",
	PermissionLevel: "manage",
	YoutubeToken: models.YoutubeToken{
		AccessToken:  "asdfasd",
		RefreshToken: "asdfkasdf",
		Expiry:       "12/12/12",
		TokenType:    "Bearer",
	},
}

func Test_Run(t *testing.T) {
	data := data.InMemorySQLiteConnect()

	managerService := ytmanager.FakeNewYoutubeManager(data)
	playlistService := playlist.NewFakePlaylist(data)
	userService := user.NewUser(data)

	job := NewInsertVideosJob(managerService, playlistService, data, log.New(os.Stdout, "Test: ", log.Lshortfile))

	bearer, err := userService.GenerateTemporarySessionBearer()
	assert.Nil(t, err)

	user, err := userService.Login(testUserModel, bearer)
	assert.Nil(t, err)

	account, err := data.NewYoutubeAccount(testAccountModel)
	assert.Nil(t, err)

	testPlaylistModel.YoutubeAccountID = account.ID

	testPlaylist, err := playlistService.New(testPlaylistModel, user)
	assert.Nil(t, err)
	channel, err := managerService.NewChannel("any-id-works")
	assert.Nil(t, err)

	_, err = managerService.NewVideo(channel, "any-id-works")
	assert.Nil(t, err)
	err = playlistService.Subscribe(testPlaylist, channel)
	assert.Nil(t, err)

	err = job.Run()
	assert.Nil(t, err)

	videos, err := playlistService.GetAllVideos(testPlaylist)
	assert.Nil(t, err)

	assert.Equal(t, 1, len(videos))
}
