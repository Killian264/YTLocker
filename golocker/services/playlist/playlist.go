package playlist

import (
	"github.com/Killian264/YTLocker/golocker/helpers/parsers"
	"github.com/Killian264/YTLocker/golocker/models"
	"golang.org/x/oauth2"
	"google.golang.org/api/youtube/v3"
)

type IYTPlaylist interface {
	Initialize(config oauth2.Config, token oauth2.Token) error
	Create(title string, description string) (*youtube.Playlist, error)
	Insert(playlistID string, videoID string) error
}

type IPlaylistManagerData interface {
	NewPlaylist(playlist *models.Playlist) error
	GetPlaylist(ID uint64) (*models.Playlist, error)

	NewPlaylistVideo(playlistID uint64, videoID uint64) error
	NewPlaylistChannel(playlistID uint64, channelID uint64) error
	RemovePlaylistChannel(playlistID uint64, channelID uint64) error

	GetFirstYoutubeClientConfig() (*models.YoutubeClientConfig, error)
	GetFirstYoutubeToken() (*models.YoutubeToken, error)
}

type Playlist struct {
	playlist IYTPlaylist
	data     IPlaylistManagerData
}

func NewPlaylist(ytplaylist IYTPlaylist, data IPlaylistManagerData) *Playlist {

	configData, err := data.GetFirstYoutubeClientConfig()
	if err != nil || configData == nil {
		panic("Failed to get playlist information")
	}
	tokenData, err := data.GetFirstYoutubeToken()
	if err != nil || tokenData == nil {
		panic("Failed to get playlist information")
	}

	config := parsers.ParseYoutubeClient(*configData)
	token := parsers.ParseYoutubeToken(*tokenData)

	ytplaylist.Initialize(config, token)

	return &Playlist{
		playlist: ytplaylist,
		data:     data,
	}
}

func (s *Playlist) Create(playlist *models.Playlist, user *models.User) (*models.Playlist, error) {

	ytPlaylist, err := s.playlist.Create(playlist.Title, playlist.Description)
	if err != nil || ytPlaylist == nil {
		return nil, nil
	}

	created := *playlist

	created.YoutubeID = ytPlaylist.Id
	created.Thumbnails = parsers.ParseYTThumbnails(ytPlaylist.Snippet.Thumbnails)

	s.data.NewPlaylist(&created)

	return &created, nil

}

func (s *Playlist) Get(ID uint64) (*models.Playlist, error) {

	return s.data.GetPlaylist(ID)

}

func (s *Playlist) Insert(playlist *models.Playlist, video *models.Video) error {

	err := s.playlist.Insert(playlist.YoutubeID, video.YoutubeID)
	if err != nil {
		return err
	}

	err = s.data.NewPlaylistVideo(playlist.ID, video.ID)
	if err == nil {
		return err
	}

	return nil

}

func (s *Playlist) Subscribe(playlist *models.Playlist, channel *models.Channel) error {

	return s.data.NewPlaylistChannel(playlist.ID, channel.ID)

}
func (s *Playlist) Unsubscribe(playlist *models.Playlist, channel *models.Channel) error {

	return s.data.RemovePlaylistChannel(playlist.ID, channel.ID)

}
