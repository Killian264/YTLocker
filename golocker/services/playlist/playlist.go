package playlist

import (
	"fmt"
	"reflect"

	"github.com/Killian264/YTLocker/golocker/helpers/parsers"
	"github.com/Killian264/YTLocker/golocker/models"
)

// New creates a new playlist
// TODO: Add unit tests to s.playlist.Initialize
func (s *PlaylistManager) New(playlist models.Playlist, user models.User) (models.Playlist, error) {
	isValidColor, _, err := s.data.PlaylistColorIsValid(user.ID, playlist.Color)
	if err != nil {
		return models.Playlist{}, err
	}

	if !isValidColor {
		return models.Playlist{}, fmt.Errorf("Duplicate playlist colors are not allowed.")
	}

	playlistService, err := s.oauth.InitializeYTService(s.playlist, playlist.YoutubeAccountID)
	if err != nil {
		return models.Playlist{}, err
	}

	ytPlaylist, err := playlistService.Create(playlist.Title, playlist.Description)
	if err != nil || ytPlaylist == nil {
		return models.Playlist{}, err
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

// Updates playlist information
func (s *PlaylistManager) Update(playlist models.Playlist) (models.Playlist, error) {
	isColorInUse, playlistUsingColor, err := s.data.PlaylistColorIsValid(playlist.UserID, playlist.Color)
	if err != nil {
		return models.Playlist{}, err
	}

	// if the color is in use from and from another playlist
	if !isColorInUse && playlistUsingColor != playlist.ID {
		return models.Playlist{}, fmt.Errorf("Duplicate playlist colors are not allowed.")
	}

	return s.data.UpdatePlaylist(playlist)
}

// Insert adds a video to a playlist
// TODO: Add unit tests to s.playlist.Initialize
func (s *PlaylistManager) Insert(playlist models.Playlist, video models.Video) error {
	playlistService, err := s.oauth.InitializeYTService(s.playlist, playlist.YoutubeAccountID)
	if err != nil {
		return fmt.Errorf("failed to initialize ytservice: " + err.Error())
	}

	err = playlistService.Insert(playlist.YoutubeID, video.YoutubeID)
	if err != nil {
		return fmt.Errorf("failed to insert video on youtube: " + err.Error())
	}

	err = s.data.NewPlaylistVideo(playlist.ID, video.ID)
	if err != nil {
		return fmt.Errorf("failed to create playlist video: " + err.Error())
	}

	return nil
}

// Insert adds a video to a playlist
func (s *PlaylistManager) Delete(playlist models.Playlist) error {
	err := s.data.DeletePlaylist(playlist.ID)
	if err == nil {
		return err
	}

	return nil
}

// ProcessNewVideo processes subscriptions for a new video
func (s *PlaylistManager) ProcessNewVideo(channel models.Channel, video models.Video) error {
	ids, err := s.data.GetAllPlaylistsSubscribedTo(channel)
	if err != nil {
		return fmt.Errorf("failed to get playlists subscribed to channel: " + err.Error())
	}

	for _, id := range ids {
		playlist, err := s.Get(id)
		if err != nil {
			return fmt.Errorf("failed to get playlist: " + err.Error())
		}

		if !playlist.Active {
			continue
		}

		// playlist created after video
		if playlist.CreatedAt.After(video.CreatedAt) {
			continue
		}

		exists, err := s.data.PlaylistHasVideo(playlist.ID, video.ID)
		if err != nil {
			return fmt.Errorf("failed to check if playlist has video: " + err.Error())
		}

		if exists {
			continue
		}

		err = s.Insert(playlist, video)
		if err != nil {
			return fmt.Errorf("failed to insert video: " + err.Error())
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

// GetLastestPlaylistVideos gets the last 30 videos for a user
func (s *PlaylistManager) GetLastestPlaylistVideos(user models.User) ([]uint64, error) {
	return s.data.GetLastestPlaylistVideos(user.ID)
}

// CopyPlaylist copies a playlist, videos are not copied
func (s *PlaylistManager) CopyPlaylist(playlist models.Playlist) (models.Playlist, error) {
	playlist, err := s.data.GetPlaylistForCopy(playlist)
	if err != nil {
		return models.Playlist{}, err
	}

	playlistService, err := s.oauth.InitializeYTService(s.playlist, playlist.YoutubeAccountID)
	if err != nil {
		return models.Playlist{}, err
	}

	ytPlaylist, err := playlistService.Create(playlist.Title, playlist.Description)
	if err != nil || ytPlaylist == nil {
		return models.Playlist{}, err
	}

	playlist.YoutubeID = ytPlaylist.Id
	playlist.Thumbnails = parsers.ParseYTThumbnails(ytPlaylist.Snippet.Thumbnails)

	return s.data.NewPlaylist(playlist.UserID, playlist)
}

// CopyPlaylist copies a playlist, videos are not copied
func (s *PlaylistManager) RefreshPlaylist(playlist models.Playlist) (models.Playlist, error) {
	playlistService, err := s.oauth.InitializeYTService(s.playlist, playlist.YoutubeAccountID)
	if err != nil {
		return models.Playlist{}, err
	}

	allFoundVideos, err := playlistService.GetPlaylistVideos(playlist.YoutubeID)
	if err != nil {
		return models.Playlist{}, err
	}

	allSavedVideos, err := s.data.GetAllPlaylistVideoYoutubeIds(playlist.ID)
	if err != nil {
		return models.Playlist{}, err
	}

	foundVideosMap := map[string]bool{}
	for _, video := range allFoundVideos {
		foundVideosMap[video] = true
	}

	for _, video := range allSavedVideos {
		_, exist := foundVideosMap[video.YoutubeID]
		if !exist {
			err = s.data.RemovePlaylistVideo(playlist.ID, video.ID)
			if err != nil {
				return models.Playlist{}, err
			}
		}
	}

	return s.data.GetPlaylist(playlist.UserID)
}
