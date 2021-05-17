package playlist

import (
	"fmt"
	"testing"

	"github.com/Killian264/YTLocker/golocker/data"
	"github.com/Killian264/YTLocker/golocker/models"
	"github.com/stretchr/testify/assert"
)

var user = models.User{
	Username: "Killian",
	Email:    "killiandebacker@gmail.com",
	Password: "one-two-three",
}

var user2 = models.User{
	Username: "Killian",
	Email:    "killiandebacker@gmail.com",
	Password: "one-two-three",
}

var playlist = models.Playlist{
	Title:       "New Playlist",
	Description: "Cool new playlist!!!",
}

var channel = models.Channel{
	YoutubeID:   "this is a youtube id",
	Title:       "This is a channel title",
	Description: "This is a channel description",
}

var video = models.Video{
	YoutubeID:   "this is a youtube id",
	Title:       "This is a video title",
	Description: "This is a video description",
}

func Test_Create_Playlist(t *testing.T) {
	s := createMockServices(t)

	playlist, err := s.New(playlist, user)
	assert.Nil(t, err)
	assert.NotNil(t, playlist)

	PlaylistExpectedIsActual(t, s, playlist, user)
}

func Test_Only_Get_Users_Playlist(t *testing.T) {
	s := createMockServices(t)

	playlist, err := s.New(playlist, user)
	assert.Nil(t, err)
	assert.NotNil(t, playlist)

	searched, err := s.Get(user2, playlist.ID)
	assert.Nil(t, err)
	assert.Equal(t, models.Playlist{}, searched)
}

func Test_Playlist_Insert(t *testing.T) {
	s := createMockServices(t)

	playlist, err := s.New(playlist, user)

	err = s.Insert(playlist, video)
	assert.Nil(t, err)

	playlist.Videos = append(playlist.Videos, video)

	PlaylistExpectedIsActual(t, s, playlist, user)
}

func Test_Playlist_Subscribe(t *testing.T) {
	s := createMockServices(t)

	playlist, err := s.New(playlist, user)

	err = s.Subscribe(playlist, channel)
	assert.Nil(t, err)

	playlist.Channels = append(playlist.Channels, channel)

	PlaylistExpectedIsActual(t, s, playlist, user)
}

func Test_Playlist_UnSubscribe(t *testing.T) {
	s := createMockServices(t)

	playlist, err := s.New(playlist, user)

	err = s.Subscribe(playlist, channel)
	err = s.Unsubscribe(playlist, channel)
	assert.Nil(t, err)

	PlaylistExpectedIsActual(t, s, playlist, user)
}

func Test_ProcessNewVideo(t *testing.T) {
	s := createMockServices(t)

	expected, _ := s.New(playlist, user)

	fmt.Println(playlist.UserID, user.ID)

	err := s.Subscribe(expected, channel)

	err = s.ProcessNewVideo(channel, video)
	assert.Nil(t, err)

	expected.Channels = append(expected.Channels, channel)
	expected.Videos = append(expected.Videos, video)

	PlaylistExpectedIsActual(t, s, expected, user)
}

func Test_Ensure_ProcessNewVideo_Only_Adds_Only_To_Single_Subscribed(t *testing.T) {
	s := createMockServices(t)

	first, _ := s.New(playlist, user)
	expected, _ := s.New(playlist, user)
	third, _ := s.New(playlist, user)

	s.Subscribe(expected, channel)
	s.ProcessNewVideo(channel, video)

	expected.Channels = append(expected.Channels, channel)
	expected.Videos = append(expected.Videos, video)

	PlaylistExpectedIsActual(t, s, expected, user)
	PlaylistExpectedIsActual(t, s, first, user)
	PlaylistExpectedIsActual(t, s, third, user)
}

func Test_Ensure_ProcessNewVideo_Only_Adds_To_All_Subscribed(t *testing.T) {
	s := createMockServices(t)

	first, _ := s.New(playlist, user)
	expected, _ := s.New(playlist, user)
	third, _ := s.New(playlist, user)

	s.Subscribe(first, channel)
	s.Subscribe(expected, channel)
	s.Subscribe(third, channel)

	s.ProcessNewVideo(channel, video)

	first.Channels = append(first.Channels, channel)
	first.Videos = append(first.Videos, video)

	expected.Channels = append(expected.Channels, channel)
	expected.Videos = append(expected.Videos, video)

	third.Channels = append(third.Channels, channel)
	third.Videos = append(third.Videos, video)

	PlaylistExpectedIsActual(t, s, expected, user)
	PlaylistExpectedIsActual(t, s, first, user)
	PlaylistExpectedIsActual(t, s, third, user)
}

func Test_IgnoreDuplicates_ProcessNewVideo(t *testing.T) {
	s := createMockServices(t)

	expected, _ := s.New(playlist, user)

	s.Subscribe(expected, channel)

	s.ProcessNewVideo(channel, video)
	s.ProcessNewVideo(channel, video)

	expected.Channels = append(expected.Channels, channel)
	expected.Videos = append(expected.Videos, video)

	PlaylistExpectedIsActual(t, s, expected, user)
}

func Test_Get_All_Playlists(t *testing.T) {
	service := createMockServices(t)

	service.New(playlist, user)
	service.New(playlist, user)

	playlists, err := service.GetAllUserPlaylists(user)
	assert.Nil(t, err)

	assert.Equal(t, 2, len(playlists))
}

func Test_Get_Playlist_Videos(t *testing.T) {
	service := createMockServices(t)

	playlist, _ := service.New(playlist, user)
	service.Insert(playlist, video)

	videos, err := service.GetAllVideos(playlist)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(videos))
}

func Test_Get_Playlist_Channels(t *testing.T) {
	service := createMockServices(t)

	playlist, _ := service.New(playlist, user)
	service.Subscribe(playlist, channel)

	channels, err := service.GetAllChannels(playlist)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(channels))
}

func Test_Get_Playlist_Thumbnails(t *testing.T) {
	service := createMockServices(t)

	playlist, _ := service.New(playlist, user)

	thumbnails, err := service.GetAllThumbnails(playlist)
	assert.Nil(t, err)
	assert.Equal(t, 5, len(thumbnails))
}

func createMockServices(t *testing.T) *PlaylistManager {

	data := data.InMemorySQLiteConnect()

	data.NewUser(&user)

	data.NewUser(&user2)

	data.NewChannel(&channel)

	data.NewVideo(&channel, &video)

	return NewFakePlaylist(data)

}

func PlaylistsAreEqual(t *testing.T, playlist1 models.Playlist, playlist2 models.Playlist) {
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

func PlaylistExpectedIsActual(t *testing.T, s *PlaylistManager, playlist models.Playlist, user models.User) {
	assert.Equal(t, user.ID, playlist.UserID)

	found, err := s.Get(user, playlist.ID)
	assert.NotNil(t, found)
	assert.Nil(t, err)

	thumbnails, err := s.GetAllThumbnails(found)
	assert.Nil(t, err)

	channels, err := s.GetAllChannels(found)
	assert.Nil(t, err)

	videos, err := s.GetAllVideos(found)
	assert.Nil(t, err)

	found.Thumbnails = thumbnails

	// Encoding decoding to database loses some information for datetimes
	assert.Equal(t, len(playlist.Thumbnails), len(thumbnails))
	assert.Equal(t, len(playlist.Channels), len(channels))
	assert.Equal(t, len(playlist.Videos), len(videos))

	playlist.CreatedAt = found.CreatedAt
	playlist.UpdatedAt = found.UpdatedAt
	playlist.Thumbnails = found.Thumbnails
	playlist.Videos = found.Videos
	playlist.Channels = found.Channels

	assert.Equal(t, playlist, found)
}
