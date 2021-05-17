package cronjobs

import (
	"fmt"
	"log"

	"github.com/Killian264/YTLocker/golocker/models"
	"github.com/Killian264/YTLocker/golocker/services"
)

type IJob interface {
	Run() error
}

type InsertVideosJob struct {
	s *services.Services
	l *log.Logger
}

func NewInsertVideosJob(s *services.Services, l *log.Logger) *InsertVideosJob {
	return &InsertVideosJob{
		s: s,
		l: l,
	}
}

func (j InsertVideosJob) Run() error {

	err := j.saveWorkUnits()
	if err != nil {
		j.l.Print("Failed to save work units error: ", err)
		return err
	}

	for true {

		work, err := j.s.Data.GetFirstSubscriptionWorkUnitByStatus("created")
		if err != nil {
			j.l.Print("Failed to get work unit: ", err)
			continue
		}
		if work == nil {
			break
		}

		err = j.processVideo(work)
		if err != nil {
			j.l.Print("Failed to process work unit ID: ", work.ID, err)
		}

		status := "complete"
		if err != nil {
			status = "error"
		}

		err = j.s.Data.UpdateSubscriptionWorkUnitStatus(work, status)
		if err != nil {
			j.l.Print("Failed to update work unit ID: ", work.ID, err)
			continue
		}

	}

	return nil

}

func (j InsertVideosJob) processVideo(workUnit *models.SubscriptionWorkUnit) error {

	channel, err := j.s.Youtube.GetChannel(workUnit.ChannelID)
	if err != nil {
		return err
	}
	if channel == nil {
		return fmt.Errorf("could not find channel")
	}

	video, err := j.s.Youtube.GetVideo(workUnit.VideoID)
	if err != nil {
		return err
	}
	if video == nil {
		return fmt.Errorf("could not find video")
	}

	return j.s.Playlist.ProcessNewVideo(*channel, *video)

}

func (j InsertVideosJob) saveWorkUnits() error {

	videos, err := j.s.Youtube.GetAllVideosFromLast24Hours()
	if err != nil {
		return err
	}

	for _, video := range *videos {

		work, err := j.s.Data.GetSubscriptionWorkUnit(video.ID, video.ChannelID)
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

		err = j.s.Data.NewSubscriptionWorkUnit(work)
		if err != nil {
			return err
		}

	}

	return nil

}
