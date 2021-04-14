package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/Killian264/YTLocker/golocker/data"
	"github.com/Killian264/YTLocker/golocker/interfaces"
	"github.com/Killian264/YTLocker/golocker/models"
	"github.com/Killian264/YTLocker/golocker/parsers"
	"github.com/Killian264/YTLocker/golocker/services/playlist"
	"github.com/Killian264/YTLocker/golocker/services/subscribe"
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

	logger.Print("Creating Playlist Service...")

	s.InitalizePlaylistService(
		"secrets/",
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
	router    *mux.Router
	data      *data.Data
	logger    *log.Logger
	youtube   *ytservice.YTService
	playlist  *playlist.Playlister
	subscribe *subscribe.Subscriber
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

func (s *Services) InitalizePlaylistService(secretsDir string) {

	logger := log.New(os.Stdout, "Subscriber: ", log.Ldate|log.Ltime|log.Lshortfile)
	logger.SetPrefix("PlaylistService: ")

	playlister := playlist.NewPlaylister(interfaces.IPlaylistData(s.data), logger)

	clientData, err := readInClientSecret(fmt.Sprintf("%s%s", secretsDir, "client_secret.json"))
	if err != nil {
		s.logger.Fatalf("Unable to read client secret file: %v", err)
	}

	tokenData, err := readInAccessToken(fmt.Sprintf("%s%s", secretsDir, "access_secret.json"))
	if err != nil {
		s.logger.Fatalf("Unable to read access secret file: %v", err)
	}

	err = playlister.SetDefaultConfig(clientData)
	if err != nil {
		logger.Print(err)
	}

	err = playlister.SetDefaultToken(tokenData)
	if err != nil {
		logger.Print(err)
	}

	s.playlist = playlister

	s.playlist.Initalize()
	s.playlist.CreatePlaylist()

}

func readInClientSecret(path string) (models.YoutubeClientConfig, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return models.YoutubeClientConfig{}, err
	}

	return parsers.ParseClientJson(string(b))
}

func readInAccessToken(path string) (models.YoutubeToken, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return models.YoutubeToken{}, err
	}

	return parsers.ParseAccessTokenJson(string(b))
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

/// FOR TESTING ONLY

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

	parsedChannel := parsers.ParseYTChannel(channel)

	err = s.data.NewChannel(&parsedChannel)

	if err != nil {
		log.Print(err)
	}

	log.Print("Channel ID: ", parsedChannel.ChannelID)

	log.Print("Description: ", parsedChannel.Description, "\n\n")

	for _, thumbnail := range parsedChannel.Thumbnails {
		log.Print(thumbnail.Height)
	}
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

	parsedVideo, channelID := parsers.ParseYTVideo(video)

	s.data.NewVideo(&parsedVideo, channelID)

}
