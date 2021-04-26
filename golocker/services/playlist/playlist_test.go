package playlist

import (
	"testing"

	"github.com/Killian264/YTLocker/golocker/data"
	"github.com/Killian264/YTLocker/golocker/models"
	"github.com/Killian264/YTLocker/golocker/services/ytservice"
	"github.com/stretchr/testify/assert"
)

var user = &models.User{
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

	playlist, err := service.Create(playlist, user)
	assert.Nil(t, err)
	assert.NotNil(t, playlist)

	created, err := service.Get(playlist.ID)
	assert.Nil(t, err)
	assert.NotNil(t, created)

	PlaylistsAreEqual(t, playlist, created)

}

func Test_Playlist_Insert(t *testing.T) {

	service := createMockServices(t)

	playlist, err := service.Create(playlist, user)

	err = service.Insert(playlist, video)
	assert.Nil(t, err)

	playlist.Videos = append(playlist.Videos, *video)

	created, err := service.Get(playlist.ID)

	PlaylistsAreEqual(t, playlist, created)

}

func Test_Playlist_Subscribe(t *testing.T) {

	service := createMockServices(t)

	playlist, err := service.Create(playlist, user)

	err = service.Subscribe(playlist, channel)
	assert.Nil(t, err)

	playlist.Channels = append(playlist.Channels, *channel)

	created, err := service.Get(playlist.ID)

	PlaylistsAreEqual(t, playlist, created)

}

func Test_Playlist_UnSubscribe(t *testing.T) {

	service := createMockServices(t)

	playlist, err := service.Create(playlist, user)

	err = service.Subscribe(playlist, channel)

	err = service.Unsubscribe(playlist, channel)
	assert.Nil(t, err)

	created, err := service.Get(playlist.ID)

	PlaylistsAreEqual(t, playlist, created)

}

func createMockServices(t *testing.T) *Playlist {

	data := data.InMemorySQLiteConnect()
	playlist := &ytservice.YTPlaylistFake{}

	data.NewYoutubeClientConfig(&models.YoutubeClientConfig{
		ClientID:     "11223534584-asdfhasdjfhwieyrwqejhkflasd.apps.googleusercontent.com",
		ClientSecret: "qwerHSwer_asdhwuerJHFDJqkqw",
		RedirectURL:  "urn:ietf:wg:oauth:2.0:oob",
		AuthURL:      "https://accounts.google.com/o/oauth2/auth",
		TokenURL:     "https://oauth2.googleapis.com/token",
	})

	data.NewYoutubeToken(&models.YoutubeToken{
		AccessToken:  "sa23.345234524623sdfasdfq-qegehgower9505034jfeworrjwertw_qqwerjfldssgert345sdgdgew-bheiyqeotleqjrljdfluao23423_QwekjfuI023kjasdfwer",
		TokenType:    "Bearer",
		RefreshToken: "asdfjqwekj23//2342329asqq-ajfdki22399jjIjiJIWJFfwerw_qwefdasferw_zwaehwejlkWW",
		Expiry:       "2021-04-13T23:30:06.1139442-05:00",
	})

	data.NewUser(user)

	data.NewChannel(channel)

	return NewPlaylist(
		playlist,
		data,
	)
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
