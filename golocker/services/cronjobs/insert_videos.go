package cronjobs

import (
	"fmt"
	"log"
	"reflect"

	"github.com/Killian264/YTLocker/golocker/data"
	"github.com/Killian264/YTLocker/golocker/models"
	"github.com/Killian264/YTLocker/golocker/services/playlist"
	"github.com/Killian264/YTLocker/golocker/services/ytmanager"
)

type IJob interface {
	Run() error
}

type InsertVideosJob struct {
	youtube  *ytmanager.YoutubeManager
	playlist *playlist.PlaylistManager
	data     *data.Data
	logger   *log.Logger
}

func NewInsertVideosJob(youtube *ytmanager.YoutubeManager, playlist *playlist.PlaylistManager, data *data.Data, logger *log.Logger) *InsertVideosJob {
	return &InsertVideosJob{
		youtube:  youtube,
		playlist: playlist,
		data:     data,
		logger:   logger,
	}
}

func (j InsertVideosJob) Run() error {
	err := j.saveWorkUnits()
	if err != nil {
		j.logger.Print("Failed to save work units error: ", err)
		return err
	}

	for true {
		work, err := j.data.GetFirstSubscriptionWorkUnitByStatus("created")
		if err != nil {
			j.logger.Print("Failed to get work unit: ", err)
			continue
		}
		if work == nil {
			break
		}

		err = j.processVideo(work)
		if err != nil {
			j.logger.Print("Failed to process work unit ID: ", work.ID, err)
		}

		status := "complete"
		if err != nil {
			status = "error"
		}

		err = j.data.UpdateSubscriptionWorkUnitStatus(work, status)
		if err != nil {
			j.logger.Print("Failed to update work unit ID: ", work.ID, err)
			continue
		}
	}

	return nil
}

func (j InsertVideosJob) processVideo(workUnit *models.SubscriptionWorkUnit) error {
	channel, err := j.youtube.GetChannel(workUnit.ChannelID)
	if err != nil {
		return err
	}
	if reflect.DeepEqual(channel, models.Channel{}) {
		return fmt.Errorf("could not find channel")
	}

	video, err := j.youtube.GetVideo(workUnit.VideoID)
	if err != nil {
		return err
	}
	if reflect.DeepEqual(video, models.Video{}) {
		return fmt.Errorf("could not find video")
	}

	return j.playlist.ProcessNewVideo(channel, video)
}

func (j InsertVideosJob) saveWorkUnits() error {
	ids, err := j.youtube.GetAllVideosFromLast24Hours()
	if err != nil {
		return err
	}

	for _, id := range ids {
		video, err := j.youtube.GetVideo(id)
		if err != nil {
			return err
		}

		work, err := j.data.GetSubscriptionWorkUnit(video.ID, video.ChannelID)
		if err != nil {
			return err
		}

		if work != nil {
			continue
		}

		work = &models.SubscriptionWorkUnit{
			ChannelID: video.ChannelID,
			VideoID:   video.ID,
			Status:    "created",
		}

		err = j.data.NewSubscriptionWorkUnit(work)
		if err != nil {
			return err
		}
	}

	return nil
}
