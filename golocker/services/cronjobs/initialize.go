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
			logger.Print("Starting Insert Videos: --------------")
			job.Run()
			logger.Print("Finished Insert Videos: --------------")
		},
		ResubscribeAllSubscriptions: func() {
			logger.Print("Starting Resubscribe: ----------------")
			err := subscriber.ResubscribeAll()
			if err != nil {
				logger.Print(err)
			}
			logger.Print("Finished Resubscribe: ----------------")
		},
		CheckForMissedUploads: func() {
			logger.Print("Starting Missed Uploads: -------------")
			err := youtube.CheckForMissedUploads(logger)
			if err != nil {
				logger.Print(err)
			}
			logger.Print("Finished Missed Uploads: -------------")
		},
		runner: cron.New(),
		logger: log.New(os.Stdout, "Cron: ", log.Ldate|log.Ltime),
	}

	manager.run()

	return manager
}

func (m CronJobManager) run() {
	m.runner.AddFunc("@every 2m", func() {
		m.InsertVideosJob()
	})

	m.runner.AddFunc("@weekly", func() {
		m.ResubscribeAllSubscriptions()
	})

	m.runner.AddFunc("@every 6h", func() {
		m.CheckForMissedUploads()
	})

	go m.runner.Run()
}
