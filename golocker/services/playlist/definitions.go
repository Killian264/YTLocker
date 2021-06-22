package playlist

import (
	"github.com/Killian264/YTLocker/golocker/helpers/parsers"
	"github.com/Killian264/YTLocker/golocker/models"
	"github.com/Killian264/YTLocker/golocker/services/ytservice"
	"golang.org/x/oauth2"
	"google.golang.org/api/youtube/v3"
)
type IYTPlaylist interface {
	Initialize(config oauth2.Config, token oauth2.Token) error
	Create(title string, description string) (*youtube.Playlist, error)
	Insert(playlistID string, videoID string) error
}

type IPlaylistManagerData interface {
	NewPlaylist(userID uint64, playlist models.Playlist) (models.Playlist, error)
	GetPlaylist(playlistID uint64) (models.Playlist, error)
	DeletePlaylist(ID uint64) (error)

	GetAllPlaylistVideos(ID uint64) ([]uint64, error) 
	GetAllPlaylistChannels(ID uint64) ([]uint64, error)
	GetThumbnails(ID uint64, ownerType string) ([]models.Thumbnail, error)

	NewPlaylistVideo(playlistID uint64, videoID uint64) error
	NewPlaylistChannel(playlistID uint64, channelID uint64) error
	RemovePlaylistChannel(playlistID uint64, channelID uint64) error

	PlaylistHasVideo(playlistID uint64, videoID uint64) (bool, error)
	GetAllPlaylistsSubscribedTo(channel models.Channel) ([]uint64, error)
	GetAllUserPlaylists(userID uint64) ([]models.Playlist, error)

	// Config
	GetFirstYoutubeClientConfig() (models.YoutubeClientConfig, error)
	GetFirstYoutubeToken() (models.YoutubeToken, error)

	GetLastestPlaylistVideos(userID uint64) ([]uint64, error)
}

// PlaylistManager manages playlists
type PlaylistManager struct {
	playlist IYTPlaylist
	data     IPlaylistManagerData
}

// NewPlaylist creates a new playlist
func NewPlaylist(yt IYTPlaylist, data IPlaylistManagerData) *PlaylistManager {
	configData, err := data.GetFirstYoutubeClientConfig()
	if err != nil || configData == (models.YoutubeClientConfig{}) {
		panic("Failed to get playlist information")
	}
	tokenData, err := data.GetFirstYoutubeToken()
	if err != nil || tokenData == (models.YoutubeToken{}) {
		panic("Failed to get playlist information")
	}

	config := parsers.ParseYoutubeClient(configData)
	token := parsers.ParseYoutubeToken(tokenData)

	yt.Initialize(config, token)

	return &PlaylistManager{
		playlist: yt,
		data:     data,
	}
}

// NewFakePlaylist creates a fake playlist service with youtube operations mocked
func NewFakePlaylist(data IPlaylistManagerData) *PlaylistManager {
	configData := models.YoutubeClientConfig{
		ClientID:     "11223534584-asdfhasdjfhwieyrwqejhkflasd.apps.googleusercontent.com",
		ClientSecret: "qwerHSwer_asdhwuerJHFDJqkqw",
		RedirectURL:  "urn:ietf:wg:oauth:2.0:oob",
		AuthURL:      "https://accounts.google.com/o/oauth2/auth",
		TokenURL:     "https://oauth2.googleapis.com/token",
	}
	tokenData := models.YoutubeToken{
		AccessToken:  "sa23.345234524623sdfasdfq-qegehgower9505034jfeworrjwertw_qqwerjfldssgert345sdgdgew-bheiyqeotleqjrljdfluao23423_QwekjfuI023kjasdfwer",
		TokenType:    "Bearer",
		RefreshToken: "asdfjqwekj23//2342329asqq-ajfdki22399jjIjiJIWJFfwerw_qwefdasferw_zwaehwejlkWW",
		Expiry:       "2021-04-13T23:30:06.1139442-05:00",
	}

	config := parsers.ParseYoutubeClient(configData)
	token := parsers.ParseYoutubeToken(tokenData)

	yt := &ytservice.YTPlaylistFake{}

	yt.Initialize(config, token)

	return &PlaylistManager{
		playlist: yt,
		data:     data,
	}
}