package services

import (
	"log"

	"github.com/Killian264/YTLocker/golocker/data"
	"github.com/Killian264/YTLocker/golocker/interfaces"
	"github.com/Killian264/YTLocker/golocker/services/ytservice"
	"github.com/gorilla/mux"
)

// Services to be injected into handlers and cron jobs
type Services struct {
	Router    *mux.Router
	Data      *data.Data
	Logger    *log.Logger
	Youtube   *ytservice.YTService
	Subscribe interfaces.ISubscription
	Channel   interfaces.IChannel
	Playlist  interfaces.IPlaylistManager
	User      interfaces.IUser
}
