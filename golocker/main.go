package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/Killian264/YTLocker/golocker/data"
	"github.com/Killian264/YTLocker/golocker/models"
	"github.com/Killian264/YTLocker/golocker/parsers"
	"github.com/Killian264/YTLocker/golocker/services/subscribe"
	"github.com/Killian264/YTLocker/golocker/services/ytservice"
	"gorm.io/gorm/logger"

	"github.com/gorilla/mux"
)

/* Main */
func main() {

	log.Print("\n\n")

	logger := log.New(os.Stdout, "", log.Lshortfile)

	logger.Print("Starting...")

	services := NewServices(logger)

	logger.Print("Running...")

	services.Run(
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
	subscribe *subscribe.Subscriber
}

func NewServices(logger *log.Logger) Services {

	s := Services{
		logger: logger,
	}

	logger.Print("Creating Router...")

	s.InitializeRouter()

	logger.Print("Creating Routes...")

	s.InitializeRoutes()

	logger.Print("Creating Data...")

	s.InitializeDatabase(
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_TCP_PORT"),
		os.Getenv("MYSQL_DATABASE"),
	)

	logger.Print("Creating Playlister...")

	s.InitalizePlaylistService(
		"secrets/",
	)

	logger.Print("Creating Youtube...")

	s.InitializeYTService(
		os.Getenv("YOUTUBE_API_KEY"),
	)

	return s
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

	clientData, err := readInClientSecret(fmt.Sprintf("%s%s", secretsDir, "client_secret.json"))
	if err != nil {
		s.logger.Fatalf("Unable to read client secret file: %v", err)
	}

	tokenData, err := readInAccessToken(fmt.Sprintf("%s%s", secretsDir, "access_secret.json"))
	if err != nil {
		s.logger.Fatalf("Unable to read access secret file: %v", err)
	}

	err = s.data.NewYoutubeClientConfig(&clientData)
	if err != nil && !strings.Contains(err.Error(), "Duplicate entry") {
		s.logger.Print(err)
	}

	err = s.data.NewYoutubeToken(&tokenData)
	if err != nil && !strings.Contains(err.Error(), "Duplicate entry") {
		s.logger.Print(err)
	}

	// playlist := models.Playlist{
	// 	Title:       "Cool Playlist",
	// 	Description: "Super COOL PLAYLIST THAT was autogenerated by magic!!!!!",
	// }

	// _, err = s.playlist.CreatePlaylist(playlist)
	// if err != nil {
	// 	s.logger.Print(err)
	// }

	// playlist := models.Playlist{
	// 	PlaylistID: "qqq",
	// }

	// err = s.playlist.PlaylistInsert(playlist, models.Video{
	// 	VideoID: "qqq",
	// })
	// if err != nil {
	// 	s.logger.Print(err)
	// }

	// err = s.playlist.PlaylistInsert(playlist, models.Video{
	// 	VideoID: "qqq",
	// })
	// if err != nil {
	// 	s.logger.Print(err)
	// }

}

// configData, err := s.dataService.GetFirstYoutubeClientConfig()
// if err != nil {
// 	return err
// }

// tokenData, err := s.dataService.GetFirstYoutubeToken()
// if err != nil {
// 	return err
// }

// config := parsers.ParseYoutubeClient(*configData)
// token := parsers.ParseYoutubeToken(*tokenData)

/*

for every user
	for each playlist
		get all videos in last day .......

*/

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
