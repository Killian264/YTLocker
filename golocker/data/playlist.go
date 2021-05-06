package data

import (
	"github.com/Killian264/YTLocker/golocker/models"
	"gorm.io/gorm/clause"
)

func (d *Data) NewYoutubeClientConfig(config *models.YoutubeClientConfig) error {
	config.ID = d.rand.ID()

	result := d.db.Create(&config)

	return result.Error
}

func (d *Data) NewYoutubeToken(token *models.YoutubeToken) error {
	token.ID = d.rand.ID()

	result := d.db.Create(&token)

	return result.Error
}

func (d *Data) GetFirstYoutubeClientConfig() (*models.YoutubeClientConfig, error) {
	config := models.YoutubeClientConfig{}

	result := d.db.First(&config)

	if result.Error != nil || notFound(result.Error) {
		return nil, removeNotFound(result.Error)
	}

	return &config, nil
}
func (d *Data) GetFirstYoutubeToken() (*models.YoutubeToken, error) {
	token := models.YoutubeToken{}

	result := d.db.First(&token)

	if result.Error != nil || notFound(result.Error) {
		return nil, removeNotFound(result.Error)
	}

	return &token, nil
}

func (d *Data) NewPlaylist(userID uint64, playlist *models.Playlist) error {

	playlist.ID = d.rand.ID()

	playlist.UserID = userID

	for _, thumbnail := range playlist.Thumbnails {
		thumbnail.ID = d.rand.ID()
	}

	result := d.db.Create(&playlist)

	return result.Error

}

func (d *Data) GetPlaylist(userID uint64, playlistID uint64) (*models.Playlist, error) {
	playlist := &models.Playlist{ID: playlistID}

	result := d.db.Preload(clause.Associations).Where(playlist).First(playlist)

	if result.Error != nil || notFound(result.Error) {
		return nil, removeNotFound(result.Error)
	}

	if userID != playlist.UserID {
		return nil, nil
	}

	return playlist, nil
}

func (d *Data) NewPlaylistVideo(playlistID uint64, videoID uint64) error {

	playlist := &models.Playlist{ID: playlistID}
	video := &models.Video{ID: videoID}

	return d.db.Model(playlist).Association("Videos").Append(video)

}

func (d *Data) NewPlaylistChannel(playlistID uint64, channelID uint64) error {

	playlist := &models.Playlist{ID: playlistID}
	channel := &models.Channel{ID: channelID}

	return d.db.Model(playlist).Association("Channels").Append(channel)

}

func (d *Data) RemovePlaylistChannel(playlistID uint64, channelID uint64) error {

	playlist := &models.Playlist{ID: playlistID}
	channel := &models.Channel{ID: channelID}

	return d.db.Model(playlist).Association("Channels").Delete(channel)

}

func (d *Data) GetAllPlaylistsSubscribedTo(channel *models.Channel) (*[]models.Playlist, error) {
	playlists := &[]models.Playlist{}

	join := "INNER JOIN playlist_channel ON playlist_channel.channel_id = ? AND playlist_channel.playlist_id = playlists.id"

	result := d.db.Joins(join, channel.ID).Find(playlists)

	if result.Error != nil || notFound(result.Error) {
		return nil, removeNotFound(result.Error)
	}

	return playlists, nil
}

func (d *Data) PlaylistHasVideo(playlistID uint64, videoID uint64) (bool, error) {

	playlist := &models.Playlist{ID: playlistID}

	join := "INNER JOIN playlist_video ON playlist_video.video_id = ? AND playlist_video.playlist_id = playlists.id"

	result := d.db.Model(playlist).Joins(join, videoID).First(playlist)

	return !notFound(result.Error), removeNotFound(result.Error)

}

func (d *Data) GetAllUserPlaylists(userID uint64) (*[]models.Playlist, error) {

	playlists := &[]models.Playlist{}

	result := d.db.Preload(clause.Associations).Where("user_id = ?", userID).Find(playlists)

	if result.Error != nil || notFound(result.Error) {
		return playlists, removeNotFound(result.Error)
	}

	return playlists, nil

}
