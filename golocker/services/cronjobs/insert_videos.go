package cronjobs

import (
	"fmt"
	"log"

	"github.com/Killian264/YTLocker/golocker/models"
	"github.com/Killian264/YTLocker/golocker/services"
)

type IJob interface {
	Run(s *services.Services, l *log.Logger) error
}

func Run(s *services.Services, l *log.Logger) error {

	l.Print("Starting Insert Videos CronJob: -------------")

	err := saveWorkUnits(s)
	if err != nil {
		l.Print("Failed to save work units error: ", err)
		return err
	}

	for true {

		work, err := s.Data.GetFirstSubscriptionWorkUnitByStatus("created")
		if err != nil {
			l.Print("Failed to get work unit: ", err)
			continue
		}

		if work == nil {
			break
		}

		err = doWork(s, work)
		if err != nil {

			l.Print("Failed to process work unit ID: ", work.ID, err)

			err = s.Data.UpdateSubscriptionWorkUnitStatus(work, "error")
			if err != nil {
				l.Print("Failed to update work unit to error ID: ", work.ID, err)
			}

			continue
		}

		err = s.Data.UpdateSubscriptionWorkUnitStatus(work, "complete")
		if err != nil {
			l.Print("Failed to update work unit ID: ", work.ID, err)
			continue
		}

	}

	l.Print("Completed Insert Videos CronJob: -------------")

	return nil

}

func doWork(s *services.Services, workUnit *models.SubscriptionWorkUnit) error {

	fmt.Print(workUnit)

	channel, err := s.Youtube.GetChannel(workUnit.ChannelID)
	if err != nil {
		return err
	}
	if channel == nil {
		return fmt.Errorf("could not find channel")
	}

	video, err := s.Youtube.GetVideo(workUnit.VideoID)
	if err != nil {
		return err
	}
	if video == nil {
		return fmt.Errorf("could not find video")
	}

	return s.Playlist.ProcessNewVideo(channel, video)

}

func saveWorkUnits(s *services.Services) error {

	videos, err := s.Youtube.GetAllVideosFromLast24Hours()
	if err != nil {
		return err
	}

	for _, video := range *videos {

		work, err := s.Data.GetSubscriptionWorkUnit(video.ID, video.ChannelID)
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

		err = s.Data.NewSubscriptionWorkUnit(work)
		if err != nil {
			return err
		}

	}

	return nil

}
