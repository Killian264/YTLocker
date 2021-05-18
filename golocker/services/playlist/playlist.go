package playlist

import (
	"reflect"

	"github.com/Killian264/YTLocker/golocker/helpers/parsers"
	"github.com/Killian264/YTLocker/golocker/models"
)

// New creates a new playlist
func (s *PlaylistManager) New(playlist models.Playlist, user models.User) (models.Playlist, error) {
	ytPlaylist, err := s.playlist.Create(playlist.Title, playlist.Description)
	if err != nil || ytPlaylist == nil {
		return models.Playlist{}, nil
	}

	playlist.YoutubeID = ytPlaylist.Id
	playlist.Thumbnails = parsers.ParseYTThumbnails(ytPlaylist.Snippet.Thumbnails)

	playlist, err = s.data.NewPlaylist(user.ID, playlist)

	return playlist, err
}

// Get gets a playlist given an id
func (s *PlaylistManager) Get(playlistID uint64) (models.Playlist, error) {
	playlist, err := s.data.GetPlaylist(playlistID)

	if reflect.DeepEqual(playlist, models.Playlist{}) {
		return models.Playlist{}, nil
	}

	return playlist, err
}

// Insert adds a video to a playlist
func (s *PlaylistManager) Insert(playlist models.Playlist, video models.Video) error {
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

// ProcessNewVideo processes subscriptions for a new video
func (s *PlaylistManager) ProcessNewVideo(channel models.Channel, video models.Video) error {
	ids, err := s.data.GetAllPlaylistsSubscribedTo(channel)
	if err != nil {
		return err
	}

	for _, id := range ids {

		playlist, err := s.Get(id)
		if err != nil {
			return err
		}

		exists, err := s.data.PlaylistHasVideo(playlist.ID, video.ID)
		if err != nil {
			return err
		}

		if exists {
			continue
		}

		err = s.Insert(playlist, video)
		if err != nil {
			return err
		}

	}

	return nil
}

// GetAllUserPlaylists returns all playlist for a user
func (s *PlaylistManager) GetAllUserPlaylists(user models.User) ([]models.Playlist, error) {
	playlists, err := s.data.GetAllUserPlaylists(user.ID)

	if playlists == nil {
		return []models.Playlist{}, nil
	}

	return playlists, err
}

// Subscribe subscribes a playlist to a channel, channel uploads will be automatically added to playlist
func (s *PlaylistManager) Subscribe(playlist models.Playlist, channel models.Channel) error {
	return s.data.NewPlaylistChannel(playlist.ID, channel.ID)
}

// Unsubscribe removes a channel subscription from a playlist, new videos on that channel will no longer be added
func (s *PlaylistManager) Unsubscribe(playlist models.Playlist, channel models.Channel) error {
	return s.data.RemovePlaylistChannel(playlist.ID, channel.ID)
}

// GetAllVideos gets an array of all the video id's in a playlist
func (s *PlaylistManager) GetAllVideos(playlist models.Playlist) ([]uint64, error) {
	return s.data.GetAllPlaylistVideos(playlist.ID)
}

// GetAllChannels gets an array of all the channel id's in a playlist
func (s *PlaylistManager) GetAllChannels(playlist models.Playlist) ([]uint64, error) {
	return s.data.GetAllPlaylistChannels(playlist.ID)
}

// GetAllThumbnails gets all thumbnail information
func (s *PlaylistManager) GetAllThumbnails(playlist models.Playlist) ([]models.Thumbnail, error) {
	return s.data.GetThumbnails(playlist.ID, "playlists")
}
