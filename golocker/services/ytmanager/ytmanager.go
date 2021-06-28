package ytmanager

import (
	"fmt"
	"log"
	"reflect"
	"time"

	"github.com/Killian264/YTLocker/golocker/helpers/parsers"
	"github.com/Killian264/YTLocker/golocker/models"
)

// NewVideo fetches and saves to the db a video from a saved channel with a videoID
func (s *YoutubeManager) NewVideo(channel models.Channel, videoID string) (models.Video, error) {
	saved, _ := s.GetVideoByID(videoID)
	if !reflect.DeepEqual(saved, models.Video{}) {
		return saved, nil
	}

	ytVideo, err := s.yt.GetVideo(channel.YoutubeID, videoID)
	if err != nil {
		return models.Video{}, err
	}
	if ytVideo == nil {
		return models.Video{}, fmt.Errorf("Video does not exist")
	}

	video, channelID := parsers.ParseYTVideo(ytVideo)
	if channelID != channel.YoutubeID {
		return models.Video{}, fmt.Errorf("Wrong channel provided for video")
	}

	video, err = s.data.NewVideo(channel, video)
	if err != nil {
		return models.Video{}, err
	}

	return video, nil
}

func (s *YoutubeManager) GetChannelIdFromUsername(username string) (string, error) {
	channelID, err := s.yt.GetChannelIDByUsername(username)
	if channelID == "" {
		return "", fmt.Errorf("Video does not exist")
	}

	return channelID, err
}

// GetVideo gets a video from the db
func (s *YoutubeManager) GetVideo(ID uint64) (models.Video, error) {
	return s.data.GetVideo(ID)
}

// GetVideoByID gets a video from the db
func (s *YoutubeManager) GetVideoByID(youtubeID string) (models.Video, error) {
	return s.data.GetVideoByID(youtubeID)
}

// NewChannel gets and saves to the db a new channel
func (s *YoutubeManager) NewChannel(channelID string) (models.Channel, error) {
	saved, _ := s.GetChannelByID(channelID)
	if !reflect.DeepEqual(saved, models.Channel{}) {
		return saved, nil
	}

	ytChannel, err := s.yt.GetChannel(channelID)
	if err != nil {
		return models.Channel{}, err
	}
	if ytChannel == nil {
		return models.Channel{}, fmt.Errorf("Channel does not exist")
	}

	channel := parsers.ParseYTChannel(ytChannel)

	channel, err = s.data.NewChannel(channel)
	if err != nil {
		return models.Channel{}, err
	}

	return channel, nil
}

// GetChannel gets a channel from the db
func (s *YoutubeManager) GetChannel(ID uint64) (models.Channel, error) {
	return s.data.GetChannel(ID)
}

// GetChannelByID gets a channel from the db
func (s *YoutubeManager) GetChannelByID(youtubeID string) (models.Channel, error) {
	return s.data.GetChannelByID(youtubeID)
}

// GetAllVideosFromLast24Hours gets the ids of all videos from the last 24 hours
func (s *YoutubeManager) GetAllVideosFromLast24Hours() ([]uint64, error) {
	return s.data.GetVideosFromLast24Hours()
}

func (s *YoutubeManager) CheckForMissedUploads(l *log.Logger) error {
	ids, err := s.data.GetAllChannels()
	if err != nil {
		return err
	}

	after := time.Now().AddDate(0, 0, -1)

	for _, id := range ids {
		channel, err := s.GetChannel(id)
		if err != nil {
			return err
		}

		response, err := s.yt.GetLastVideosFromChannel(channel.YoutubeID, "", after)
		if err != nil {
			return err
		}

		videos := parsers.ParseSearchResponseIntoVideos(response)

		for _, video := range videos {
			_, err = s.data.NewVideo(channel, video)
			if err != nil {
				l.Printf("MissedUploads: Error processing video: %v", err)
			}
		}
	}

	return nil
}

// GetAllVideos gets an array of all the video id's in a playlist
func (s *YoutubeManager) GetAllChannelVideos(channel models.Channel) ([]uint64, error) {
	return s.data.GetAllChannelVideos(channel.ID)
}

// GetAllThumbnails gets all thumbnail information
func (s *YoutubeManager) GetAllVideoThumbnails(video models.Video) ([]models.Thumbnail, error) {
	return s.data.GetThumbnails(video.ID, "videos")
}

// GetAllThumbnails gets all thumbnail information
func (s *YoutubeManager) GetAllChannelThumbnails(channel models.Channel) ([]models.Thumbnail, error) {
	return s.data.GetThumbnails(channel.ID, "channels")
}