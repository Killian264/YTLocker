package playlist

import (
	"testing"

	"github.com/Killian264/YTLocker/golocker/data"
	"github.com/Killian264/YTLocker/golocker/models"
	"github.com/stretchr/testify/assert"
)

var user = models.User{}

// playlists can be considered unique
var playlist = models.Playlist{}
var playlist2 = models.Playlist{}
var playlist3 = models.Playlist{}

var channel = models.Channel{}
var video = models.Video{}

func Test_Create_Playlist(t *testing.T) {
	s := createMockServices(t)

	playlist, err := s.New(playlist, user)
	assert.Nil(t, err)
	assert.NotNil(t, playlist)

	playlistExpectedIsActual(t, s, playlist, user)
}

func Test_Create_Playlist_Fails_On_Duplicate_Color(t *testing.T) {
	s := createMockServices(t)

	s.New(playlist, user)

	_, err := s.New(playlist, user)
	assert.NotNil(t, err)
}

func Test_Update_Playlist(t *testing.T) {
	s := createMockServices(t)
	
	playlist1, _ := s.New(playlist, user)

	playlist1.Title = "something else"
	playlist1.Description = "something else"
	playlist1.Color = "something else"

	playlist2, err := s.Update(playlist1)
	assert.Nil(t, err)
	assert.NotNil(t, playlist2)

	playlistExpectedIsActual(t, s, playlist1, user)
}

func Test_Update_Playlist_Fails_On_Duplicate_Color(t *testing.T) {
	s := createMockServices(t)
	
	playlist1, err := s.New(playlist, user)
	playlist2, err := s.New(playlist2, user)

	playlist1.Title = "something else"
	playlist1.Description = "something else"
	playlist1.Color = playlist2.Color

	playlist2, err = s.Update(playlist1)
	assert.NotNil(t, err)
}

func Test_Playlist_Insert(t *testing.T) {
	s := createMockServices(t)

	playlist, err := s.New(playlist, user)

	err = s.Insert(playlist, video)
	assert.Nil(t, err)

	playlist.Videos = append(playlist.Videos, video)

	playlistExpectedIsActual(t, s, playlist, user)
}

func Test_Delete_Playlist(t *testing.T) {
	s := createMockServices(t)

	s.New(playlist, user)

	err := s.Delete(playlist)
	assert.Nil(t, err)

	playlist, err := s.Get(playlist.ID)
	assert.Nil(t, err)
	assert.Equal(t, models.Playlist{}, playlist)
}

func Test_Playlist_Subscribe(t *testing.T) {
	s := createMockServices(t)

	playlist, err := s.New(playlist, user)

	err = s.Subscribe(playlist, channel)
	assert.Nil(t, err)

	playlist.Channels = append(playlist.Channels, channel)

	playlistExpectedIsActual(t, s, playlist, user)
}

func Test_Playlist_UnSubscribe(t *testing.T) {
	s := createMockServices(t)

	playlist, err := s.New(playlist, user)

	err = s.Subscribe(playlist, channel)
	err = s.Unsubscribe(playlist, channel)
	assert.Nil(t, err)

	playlistExpectedIsActual(t, s, playlist, user)
}

func Test_ProcessNewVideo(t *testing.T) {
	s := createMockServices(t)

	expected, _ := s.New(playlist, user)

	err := s.Subscribe(expected, channel)

	err = s.ProcessNewVideo(channel, video)
	assert.Nil(t, err)

	expected.Channels = append(expected.Channels, channel)
	expected.Videos = append(expected.Videos, video)

	playlistExpectedIsActual(t, s, expected, user)
}

func Test_Ensure_ProcessNewVideo_Only_Adds_Only_To_Single_Subscribed(t *testing.T) {
	s := createMockServices(t)

	first, _ := s.New(playlist, user)
	expected, _ := s.New(playlist2, user)
	third, _ := s.New(playlist3, user)

	s.Subscribe(expected, channel)
	s.ProcessNewVideo(channel, video)

	expected.Channels = append(expected.Channels, channel)
	expected.Videos = append(expected.Videos, video)

	playlistExpectedIsActual(t, s, expected, user)
	playlistExpectedIsActual(t, s, first, user)
	playlistExpectedIsActual(t, s, third, user)
}

func Test_Ensure_ProcessNewVideo_Only_Adds_To_All_Subscribed(t *testing.T) {
	s := createMockServices(t)

	first, _ := s.New(playlist, user)
	expected, _ := s.New(playlist2, user)
	third, _ := s.New(playlist3, user)

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

	playlistExpectedIsActual(t, s, expected, user)
	playlistExpectedIsActual(t, s, first, user)
	playlistExpectedIsActual(t, s, third, user)
}

func Test_IgnoreDuplicates_ProcessNewVideo(t *testing.T) {
	s := createMockServices(t)

	expected, _ := s.New(playlist, user)

	s.Subscribe(expected, channel)

	s.ProcessNewVideo(channel, video)
	s.ProcessNewVideo(channel, video)

	expected.Channels = append(expected.Channels, channel)
	expected.Videos = append(expected.Videos, video)

	playlistExpectedIsActual(t, s, expected, user)
}

func Test_Get_All_Playlists(t *testing.T) {
	service := createMockServices(t)

	service.New(playlist, user)
	service.New(playlist2, user)

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

	setDefaultModels();

	data.NewUser(&user)
	data.NewUser(&user)

	newChannel, _ := data.NewChannel(channel)
	channel = newChannel

	newVideo, _ := data.NewVideo(channel, video)
	video = newVideo

	return NewFakePlaylist(data)
}

func playlistExpectedIsActual(t *testing.T, s *PlaylistManager, playlist models.Playlist, user models.User) {
	assert.Equal(t, user.ID, playlist.UserID)

	found, err := s.Get(playlist.ID)
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

func setDefaultModels(){
	user = models.User{
		Username: "Killian",
		Email:    "killiandebacker@gmail.com",
		Password: "one-two-three",
	}
	
	playlist = models.Playlist{
		Title:       "New Playlist",
		Description: "Cool new playlist!!!",
		Color: "red-1",
	}

	playlist2 = models.Playlist{
		Title:       "New Playlist",
		Description: "Cool new playlist!!!",
		Color: "red-2",
	}

	playlist3 = models.Playlist{
		Title:       "New Playlist",
		Description: "Cool new playlist!!!",
		Color: "red-3",
	}
	
	channel = models.Channel{
		YoutubeID:   "this is a youtube id",
		Title:       "This is a channel title",
		Description: "This is a channel description",
	}
	
	video = models.Video{
		YoutubeID:   "this is a youtube id",
		Title:       "This is a video title",
		Description: "This is a video description",
	}
}