package cronjobs

import (
	"log"
	"os"

	"github.com/Killian264/YTLocker/golocker/data"
	"github.com/Killian264/YTLocker/golocker/services/playlist"
	"github.com/Killian264/YTLocker/golocker/services/subscribe"
	"github.com/Killian264/YTLocker/golocker/services/ytmanager"
	"github.com/robfig/cron"
)

// PlaylistManager manages playlists
type CronJobManager struct {
	InsertVideosJob             func()
	ResubscribeAllSubscriptions func()
	CheckForMissedUploads       func()

	logger *log.Logger
	runner *cron.Cron
}

func NewCronJobManager(youtube *ytmanager.YoutubeManager, playlist *playlist.PlaylistManager, subscriber *subscribe.Subscriber, data *data.Data) *CronJobManager {
	logger := log.New(os.Stdout, "Cron: ", log.Ldate|log.Ltime)

	job := NewInsertVideosJob(youtube, playlist, data, logger)

	manager := &CronJobManager{
		InsertVideosJob: func() {
			job.Run()
		},
		ResubscribeAllSubscriptions: func() {
			err := subscriber.ResubscribeAll()
			if err != nil {
				logger.Print(err)
			}
		},
		CheckForMissedUploads: func() {
			err := youtube.CheckForMissedUploads(logger)
			if err != nil {
				logger.Print(err)
			}
		},
		runner: cron.New(),
		logger: log.New(os.Stdout, "Cron: ", log.Ldate|log.Ltime),
	}

	manager.run()

	return manager
}

func (m CronJobManager) run() {
	m.runner.AddFunc("@every 2m", func() {
		m.logger.Print("Starting Insert Videos: --------------")
		m.InsertVideosJob()
		m.logger.Print("Finished Insert Videos: --------------")
	})

	m.runner.AddFunc("@weekly", func() {
		m.logger.Print("Starting Resubscribe: ----------------")
		m.ResubscribeAllSubscriptions()
		m.logger.Print("Finished Resubscribe: ----------------")
	})

	m.runner.AddFunc("@every 6h", func() {
		m.logger.Print("Starting Missed Uploads: -------------")
		m.CheckForMissedUploads()
		m.logger.Print("Finished Missed Uploads: -------------")
	})

	go m.runner.Run()
}
