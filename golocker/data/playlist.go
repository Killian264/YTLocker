package data

import (
	"github.com/Killian264/YTLocker/golocker/models"
	"gorm.io/gorm"
)

func (d *Data) NewPlaylist(userID uint64, playlist models.Playlist) (models.Playlist, error) {
	playlist.ID = d.rand.ID()

	playlist.UserID = userID

	for _, thumbnail := range playlist.Thumbnails {
		thumbnail.ID = d.rand.ID()
	}

	result := d.db.Create(&playlist)

	return playlist, result.Error
}

func (d *Data) GetPlaylist(playlistID uint64) (models.Playlist, error) {
	playlist := models.Playlist{}

	result := d.db.First(&playlist, playlistID)

	if result.Error != nil || notFound(result.Error) {
		return models.Playlist{}, removeNotFound(result.Error)
	}

	return playlist, nil
}

// PlaylistColorIsValid checks if a color is valid retuns wasFound, idOfFound, error
func (d *Data) PlaylistColorIsValid(userID uint64, color string) (bool, uint64, error) {
	playlist := models.Playlist{}

	result := d.db.Where("user_id = ? AND color = ?", userID, color).First(&playlist)

	if(result.Error == gorm.ErrRecordNotFound){
		return true, 0, nil
	}

	if result.Error != nil {
		return false, 0, result.Error
	}

	return false, playlist.ID, nil
}

func (d *Data) UpdatePlaylist(playlist models.Playlist) (models.Playlist, error) {
	result := d.db.Model(&playlist).Select("Title", "Description", "Color").Updates(playlist)

	if(result.Error != nil){
		return models.Playlist{}, result.Error
	}

	return d.GetPlaylist(playlist.ID)
}

func (d *Data) DeletePlaylist(ID uint64) (error) {
	result := d.db.Delete(&models.Playlist{}, ID)

	return result.Error
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

func (d *Data) GetAllPlaylistsSubscribedTo(channel models.Channel) ([]uint64, error) {
	playlists := []OnlyID{}

	join := "INNER JOIN playlist_channel ON playlist_channel.channel_id = ? AND playlist_channel.playlist_id = playlists.id"

	result := d.db.Model(models.Playlist{}).Joins(join, channel.ID).Find(&playlists)

	if result.Error != nil || notFound(result.Error) {
		return []uint64{}, removeNotFound(result.Error)
	}

	return parseOnlyIDArray(playlists), nil;
}

func (d *Data) PlaylistHasVideo(playlistID uint64, videoID uint64) (bool, error) {
	playlist := &models.Playlist{}

	join := "INNER JOIN playlist_video ON playlist_video.video_id = ? AND playlist_video.playlist_id = playlists.id"

	result := d.db.Model(playlist).Joins(join, videoID).First(playlist, playlistID)

	return !notFound(result.Error), removeNotFound(result.Error)
}

func (d *Data) GetAllUserPlaylists(userID uint64) ([]models.Playlist, error) {
	playlists := []models.Playlist{}

	result := d.db.Where("user_id = ?", userID).Find(&playlists)

	if result.Error != nil || notFound(result.Error) {
		return playlists, removeNotFound(result.Error)
	}

	return playlists, nil
}

func (d *Data) GetAllPlaylistVideos(ID uint64) ([]uint64, error) {
	videos := []OnlyID{}

	result := d.db.Raw(
		`SELECT PV.video_id AS id FROM playlists P 
		JOIN playlist_video PV
			ON P.id = PV.playlist_id
		JOIN videos AS V
			ON PV.video_id = V.id
		WHERE P.id = ?
		ORDER BY V.created_at DESC;`, 
		ID,
	).Scan(&videos);

	if removeNotFound(result.Error) != nil {
		return nil, result.Error
	}

	return parseOnlyIDArray(videos), nil;
}


func (d *Data) GetAllPlaylistChannels(ID uint64) ([]uint64, error) {
	channels := []OnlyID{}

	result := d.db.Raw(
		`SELECT C.channel_id AS id FROM playlists P 
		JOIN playlist_channel C
			ON P.id = C.playlist_id
		WHERE P.id = ?;`, 
		ID,
	).Scan(&channels);

	if removeNotFound(result.Error) != nil {
		return nil, result.Error
	}

	return parseOnlyIDArray(channels), nil;
}

func (d *Data) GetLastestPlaylistVideos(userID uint64) ([]uint64, error) {
	videos := []OnlyID{}

	result := d.db.Raw(
		`SELECT DISTINCT 
			PV.video_id AS id,
			V.created_at
		FROM playlists P 
		JOIN playlist_video PV
			ON P.id = PV.playlist_id
		JOIN videos AS V
			ON PV.video_id = V.id
		WHERE P.user_id = ?
		AND P.deleted_at IS NULL
		ORDER BY V.created_at DESC
		LIMIT 30
		;`, userID,
	).Scan(&videos);

	if removeNotFound(result.Error) != nil {
		return nil, result.Error
	}

	return parseOnlyIDArray(videos), nil;
}