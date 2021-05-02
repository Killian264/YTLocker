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
	NewPlaylist(userID uint64, playlist *models.Playlist) error
	GetPlaylist(userID uint64, playlistID uint64) (*models.Playlist, error)

	NewPlaylistVideo(playlistID uint64, videoID uint64) error
	NewPlaylistChannel(playlistID uint64, channelID uint64) error
	RemovePlaylistChannel(playlistID uint64, channelID uint64) error

	PlaylistHasVideo(playlistID uint64, videoID uint64) (bool, error)

	GetFirstYoutubeClientConfig() (*models.YoutubeClientConfig, error)
	GetFirstYoutubeToken() (*models.YoutubeToken, error)

	GetAllPlaylistsSubscribedTo(channel *models.Channel) (*[]models.Playlist, error)
}

// PlaylistManager manages playlists
type PlaylistManager struct {
	playlist IYTPlaylist
	data     IPlaylistManagerData
}

// NewPlaylist creates a new playlist
func NewPlaylist(yt IYTPlaylist, data IPlaylistManagerData) *PlaylistManager {

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

// New creates a new playlist
func (s *PlaylistManager) New(playlist *models.Playlist, user *models.User) (*models.Playlist, error) {

	ytPlaylist, err := s.playlist.Create(playlist.Title, playlist.Description)
	if err != nil || ytPlaylist == nil {
		return nil, nil
	}

	created := *playlist

	created.YoutubeID = ytPlaylist.Id
	created.Thumbnails = parsers.ParseYTThumbnails(ytPlaylist.Snippet.Thumbnails)

	s.data.NewPlaylist(user.ID, &created)

	return &created, nil

}

// Get gets a playlist given an id
func (s *PlaylistManager) Get(user *models.User, playlistID uint64) (*models.Playlist, error) {

	return s.data.GetPlaylist(user.ID, playlistID)

}

// Insert adds a video to a playlist
func (s *PlaylistManager) Insert(playlist *models.Playlist, video *models.Video) error {

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

// Subscribe subscribes a playlist to a channel, channel uploads will be automatically added to playlist
func (s *PlaylistManager) Subscribe(playlist *models.Playlist, channel *models.Channel) error {

	return s.data.NewPlaylistChannel(playlist.ID, channel.ID)

}

// Unsubscribe removes a channel subscription from a playlist, new videos on that channel will no longer be added
func (s *PlaylistManager) Unsubscribe(playlist *models.Playlist, channel *models.Channel) error {

	return s.data.RemovePlaylistChannel(playlist.ID, channel.ID)

}

// ProcessNewVideo processes subscriptions for a new video
func (s *PlaylistManager) ProcessNewVideo(channel *models.Channel, video *models.Video) error {

	playlists, err := s.data.GetAllPlaylistsSubscribedTo(channel)
	if err != nil {
		return err
	}

	for _, playlist := range *playlists {

		exists, err := s.data.PlaylistHasVideo(playlist.ID, video.ID)
		if err != nil {
			return err
		}

		if exists {
			continue
		}

		err = s.Insert(&playlist, video)
		if err != nil {
			return err
		}

	}

	return nil

}
