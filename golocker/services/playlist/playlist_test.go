package playlist

import (
	"testing"

	"github.com/Killian264/YTLocker/golocker/data"
	"github.com/Killian264/YTLocker/golocker/models"
	"github.com/stretchr/testify/assert"
)

var user = &models.User{
	Username: "Killian",
	Email:    "killiandebacker@gmail.com",
	Password: "one-two-three",
}

var user2 = &models.User{
	Username: "Killian",
	Email:    "killiandebacker@gmail.com",
	Password: "one-two-three",
}

var playlist = &models.Playlist{
	Title:       "New Playlist",
	Description: "Cool new playlist!!!",
}

var channel = &models.Channel{
	YoutubeID:   "this is a youtube id",
	Title:       "This is a channel title",
	Description: "This is a channel description",
}

var video = &models.Video{
	YoutubeID:   "this is a youtube id",
	Title:       "This is a video title",
	Description: "This is a video description",
}

func Test_Create_Playlist(t *testing.T) {

	service := createMockServices(t)

	playlist, err := service.New(playlist, user)
	assert.Nil(t, err)
	assert.NotNil(t, playlist)

	created, err := service.Get(user, playlist.ID)
	assert.Nil(t, err)
	assert.NotNil(t, created)

	assert.Equal(t, user.ID, playlist.UserID)

	PlaylistsAreEqual(t, playlist, created)

}

func Test_Only_Get_Users_Playlist(t *testing.T) {

	service := createMockServices(t)

	playlist, err := service.New(playlist, user)
	assert.Nil(t, err)
	assert.NotNil(t, playlist)

	created, err := service.Get(user2, playlist.ID)
	assert.Nil(t, err)
	assert.Nil(t, created)

}

func Test_Playlist_Insert(t *testing.T) {

	service := createMockServices(t)

	playlist, err := service.New(playlist, user)

	err = service.Insert(playlist, video)
	assert.Nil(t, err)

	playlist.Videos = append(playlist.Videos, *video)

	created, err := service.Get(user, playlist.ID)

	PlaylistsAreEqual(t, playlist, created)

}

func Test_Playlist_Subscribe(t *testing.T) {

	service := createMockServices(t)

	playlist, err := service.New(playlist, user)

	err = service.Subscribe(playlist, channel)
	assert.Nil(t, err)

	playlist.Channels = append(playlist.Channels, *channel)

	created, err := service.Get(user, playlist.ID)

	PlaylistsAreEqual(t, playlist, created)

}

func Test_Playlist_UnSubscribe(t *testing.T) {

	service := createMockServices(t)

	playlist, err := service.New(playlist, user)

	err = service.Subscribe(playlist, channel)

	err = service.Unsubscribe(playlist, channel)
	assert.Nil(t, err)

	created, err := service.Get(user, playlist.ID)

	PlaylistsAreEqual(t, playlist, created)

}

func Test_ProcessNewVideo(t *testing.T) {

	service := createMockServices(t)

	expected, _ := service.New(playlist, user)

	err := service.Subscribe(expected, channel)
	assert.Nil(t, err)

	err = service.ProcessNewVideo(channel, video)
	assert.Nil(t, err)

	expected.Channels = append(expected.Channels, *channel)
	expected.Videos = append(expected.Videos, *video)

	created, _ := service.Get(user, expected.ID)

	PlaylistsAreEqual(t, expected, created)
	assert.Equal(t, 1, len(created.Videos))

}

func Test_IgnoreDuplicates_ProcessNewVideo(t *testing.T) {

	service := createMockServices(t)

	playlist, err := service.New(playlist, user)

	err = service.Subscribe(playlist, channel)

	err = service.ProcessNewVideo(channel, video)
	assert.Nil(t, err)

	err = service.ProcessNewVideo(channel, video)
	assert.Nil(t, err)

	created, err := service.Get(user, playlist.ID)

	assert.Equal(t, 1, len(created.Videos))

}

func Test_Get_All_Playlists(t *testing.T) {

	service := createMockServices(t)

	service.New(playlist, user)
	service.New(playlist, user)

	playlists, err := service.GetAllUserPlaylists(user)
	assert.Nil(t, err)

	assert.Equal(t, 2, len(*playlists))

}

func createMockServices(t *testing.T) *PlaylistManager {

	data := data.InMemorySQLiteConnect()

	data.NewUser(user)

	data.NewUser(user2)

	data.NewChannel(channel)

	data.NewVideo(channel, video)

	return NewFakePlaylist(data)

}

func PlaylistsAreEqual(t *testing.T, playlist1 *models.Playlist, playlist2 *models.Playlist) {
	assert.Equal(t, len(playlist1.Thumbnails), len(playlist2.Thumbnails))
	assert.Equal(t, len(playlist1.Videos), len(playlist2.Videos))

	// Encoding decoding to database loses some information for datetimes
	playlist1.CreatedAt = playlist2.CreatedAt
	playlist1.UpdatedAt = playlist2.UpdatedAt
	playlist1.Thumbnails = playlist2.Thumbnails
	playlist1.Videos = playlist2.Videos
	playlist1.Channels = playlist2.Channels

	assert.Equal(t, playlist1, playlist2)
}
