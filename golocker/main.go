package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Killian264/YTLocker/golocker/data"
	"github.com/Killian264/YTLocker/golocker/services/ytservice"
	"gorm.io/gorm/logger"

	"github.com/gorilla/mux"
)

/* Main */
func main() {

	logger := log.New(os.Stdout, "Subscriber: ", log.Ldate|log.Ltime|log.Lshortfile)

	logger.SetPrefix("YTLocker: ")

	logger.Print("Starting...")

	s := Services{}

	logger.Print("Creating Router...")

	s.InitializeRouter()

	logger.Print("Creating Routes...")

	s.InitializeRoutes()

	logger.Print("Creating Data Service...")

	s.InitializeDatabase(
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_TCP_PORT"),
		os.Getenv("MYSQL_DATABASE"),
	)

	logger.Print("Creating Youtube Service...")

	s.InitializeYTService(
		os.Getenv("YOUTUBE_API_KEY"),
	)

	logger.Print("Running...")

	s.Run(
		os.Getenv("GO_API_HOST"),
		os.Getenv("GO_API_PORT"),
	)

	logger.Print("Exiting...")

}

// App contains services for handlers
type Services struct {
	router  *mux.Router
	data    *data.Data
	logger  *log.Logger
	youtube *ytservice.YTService
}

// InitializeRouter Creates Router for app
func (s *Services) InitializeRouter() {

	s.router = mux.NewRouter()

}

// InitializeRoutes creates the routes
func (s *Services) InitializeRoutes() {

	s.router.HandleFunc("/channel/{channel_id}", func(w http.ResponseWriter, r *http.Request) {
		s.ChannelHandler(w, r)
	})

	s.router.HandleFunc("/video/{video_id}", func(w http.ResponseWriter, r *http.Request) {
		s.VideoHandler(w, r)
	})

	s.router.HandleFunc("/channel2/{channel_id}", func(w http.ResponseWriter, r *http.Request) {
		s.Channel2Handler(w, r)
	})

}

// InitializeYTService Creates YTService for app
func (s *Services) InitializeYTService(apiKey string) {
	logger := log.New(os.Stdout, "Subscriber: ", log.Ldate|log.Ltime|log.Lshortfile)
	logger.SetPrefix("YTService: ")

	ytService := ytservice.NewYoutubeService(apiKey, logger)
	s.youtube = ytService
}

// InitializeDatabase creates DB Connection for app
func (s *Services) InitializeDatabase(username string, password string, ip string, port string, name string) {

	db := new(data.Data)

	logger := logger.New(
		s.logger,
		logger.Config{},
	)

	db.Initialize(username, password, ip, port, name, logger)

	s.data = db
}

func (a *Services) Run(host string, port string) {
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", host, port), a.router))
}

func (s *Services) Channel2Handler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	s.data.GetChannel(vars["channel_id"])
}

// ChannelHandler handler to mess around with yt api
func (s *Services) ChannelHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	channel, err := s.youtube.GetChannel(vars["channel_id"])
	if err != nil || channel == nil {
		return
	}

	fmt.Print("====================================================================\n")
	fmt.Print(channel.Snippet.Title)
	fmt.Print("\n\n")
	fmt.Print(channel.Snippet.Description)
	fmt.Print("\n\n")
	fmt.Print("====================================================================\n")

	s.data.NewChannel(channel)
}

// VideoHandler handler to mess around with yt api
func (s *Services) VideoHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	video, err := s.youtube.GetVideo(vars["video_id"])
	if err != nil || video == nil {
		return
	}

	fmt.Print("====================================================================\n")
	fmt.Print(video.Snippet.Title)
	fmt.Print("\n\n")
	fmt.Print(video.Snippet.Description)
	fmt.Print("\n\n")
	fmt.Print("====================================================================\n")

	s.data.NewVideo(video)

}
