package playlist

import (
	"github.com/Killian264/YTLocker/golocker/data"
	"github.com/Killian264/YTLocker/golocker/models"
	"github.com/Killian264/YTLocker/golocker/services/oauthmanager"
	"github.com/Killian264/YTLocker/golocker/services/ytservice"
	"golang.org/x/oauth2"
	"google.golang.org/api/youtube/v3"
)

type IYTPlaylist interface {
	Initialize(config oauth2.Config, token oauth2.Token) error
	Create(title string, description string) (*youtube.Playlist, error)
	Insert(playlistID string, videoID string) error
}

type IOauthManager interface {
	GetBaseConfig() oauth2.Config
	GetBaseToken() oauth2.Token
}

type IPlaylistManagerData interface {
	NewPlaylist(userID uint64, playlist models.Playlist) (models.Playlist, error)
	GetPlaylist(playlistID uint64) (models.Playlist, error)
	UpdatePlaylist(playlist models.Playlist) (models.Playlist, error)
	DeletePlaylist(ID uint64) error

	PlaylistColorIsValid(userID uint64, color string) (bool, uint64, error)

	GetAllPlaylistVideos(ID uint64) ([]uint64, error)
	GetAllPlaylistChannels(ID uint64) ([]uint64, error)
	GetThumbnails(ID uint64, ownerType string) ([]models.Thumbnail, error)

	NewPlaylistVideo(playlistID uint64, videoID uint64) error
	NewPlaylistChannel(playlistID uint64, channelID uint64) error
	RemovePlaylistChannel(playlistID uint64, channelID uint64) error

	PlaylistHasVideo(playlistID uint64, videoID uint64) (bool, error)
	GetAllPlaylistsSubscribedTo(channel models.Channel) ([]uint64, error)
	GetAllUserPlaylists(userID uint64) ([]models.Playlist, error)

	GetLastestPlaylistVideos(userID uint64) ([]uint64, error)
}

// PlaylistManager manages playlists
type PlaylistManager struct {
	playlist IYTPlaylist
	oauth    IOauthManager
	data     IPlaylistManagerData
}

// NewPlaylist creates a new playlist
func NewPlaylist(yt IYTPlaylist, oauth IOauthManager, data IPlaylistManagerData) *PlaylistManager {
	config := oauth.GetBaseConfig()
	token := oauth.GetBaseToken()

	yt.Initialize(config, token)

	return &PlaylistManager{
		playlist: yt,
		oauth:    oauth,
		data:     data,
	}
}

// NewFakePlaylist creates a fake playlist service with youtube operations mocked
// NOTE: These secrets are fake
func NewFakePlaylist(data *data.Data) *PlaylistManager {
	return NewPlaylist(
		&ytservice.YTPlaylistFake{},
		oauthmanager.NewFakeOauthManager(data),
		data,
	)
}
