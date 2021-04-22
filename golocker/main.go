package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/Killian264/YTLocker/golocker/data"
	"github.com/Killian264/YTLocker/golocker/handlers"
	"github.com/Killian264/YTLocker/golocker/helpers/parsers"
	"github.com/Killian264/YTLocker/golocker/models"
	"github.com/Killian264/YTLocker/golocker/services"
	"github.com/Killian264/YTLocker/golocker/services/ytservice"
	"gorm.io/gorm/logger"

	muxhandler "github.com/gorilla/handlers"

	"github.com/gorilla/mux"
)

/* Main */
func main() {

	logger := log.New(os.Stdout, "Main: ", log.Lshortfile)

	logger.Print("-----------------------------------------")

	logger.Print("Starting...")

	s := NewServices(logger)

	logger.Print("Running...")

	logger.Print("-----------------------------------------")

	Run(
		&s,
		os.Getenv("GO_API_HOST"),
		os.Getenv("GO_API_PORT"),
	)

	logger.Print("Exiting...")

}

func NewServices(logger *log.Logger) services.Services {

	s := services.Services{
		Logger: logger,
	}

	logger.Print("Creating Router...")
	InitializeRouter(&s)

	logger.Print("Creating Routes...")
	InitializeRoutes(&s)

	logger.Print("Creating Data...")
	InitializeDatabase(
		&s,
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_TCP_PORT"),
		os.Getenv("MYSQL_DATABASE"),
	)

	logger.Print("Creating Playlister...")
	InitalizePlaylistService(
		&s,
		"secrets/",
	)

	logger.Print("Creating Youtube...")
	InitializeYTService(
		&s,
		os.Getenv("YOUTUBE_API_KEY"),
	)

	logger.Print("Creating Subscribe...")
	InitalizeSubscribeService(
		&s,
		os.Getenv("GO_API_URL"),
	)

	return s
}

func InitalizeSubscribeService(s *services.Services, apiURL string) {

	// service := subscribe.NewSubscriber(
	// 	interfaces.ISubscriptionData(s.Data),
	// 	interfaces.IYoutubeService(s.Youtube),
	// 	log.New(os.Stdout, "Sub: ", log.Lshortfile),
	// )

	// service.SetYTPubSubUrl("https://pubsubhubbub.appspot.com/subscribe")
	// service.SetSubscribeUrl(apiURL, "/subscribe/{secret}")

	// s.Subscribe = service

}

// InitializeRouter Creates Router for app
func InitializeRouter(s *services.Services) {

	s.Router = mux.NewRouter()

	s.Router.Use(muxhandler.RecoveryHandler())

	// s.Router.Use(handlers.LoggingMiddlewareTest)

}

// InitializeRoutes creates the routes
func InitializeRoutes(s *services.Services) {

	logger := log.New(os.Stdout, "Han: ", log.Lshortfile)

	ServiceInjector := handlers.CreateServiceInjector(s)

	ErrorHandler := handlers.CreateErrorHandler(logger)

	s.Router.HandleFunc("/subscribe/{secret}", ErrorHandler(ServiceInjector(handlers.HandleYoutubePush)))

}

// InitializeYTService Creates YTService for app
func InitializeYTService(s *services.Services, apiKey string) {

	s.Youtube = ytservice.NewYoutubeService(apiKey)

}

func InitalizePlaylistService(s *services.Services, secretsDir string) {

	clientData, err := readInClientSecret(fmt.Sprintf("%s%s", secretsDir, "client_secret.json"))
	if err != nil {
		s.Logger.Fatalf("Unable to read client secret file: %v", err)
	}

	tokenData, err := readInAccessToken(fmt.Sprintf("%s%s", secretsDir, "access_secret.json"))
	if err != nil {
		s.Logger.Fatalf("Unable to read access secret file: %v", err)
	}

	err = s.Data.NewYoutubeClientConfig(&clientData)
	if err != nil && !strings.Contains(err.Error(), "Duplicate entry") {
		s.Logger.Fatal(err)
	}

	err = s.Data.NewYoutubeToken(&tokenData)
	if err != nil && !strings.Contains(err.Error(), "Duplicate entry") {
		s.Logger.Fatal(err)
	}

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
func InitializeDatabase(s *services.Services, username string, password string, ip string, port string, name string) {

	logBase := log.New(os.Stdout, "Data: ", log.Lshortfile)

	logger := logger.New(
		logBase,
		logger.Config{},
	)

	s.Data = data.MySQLConnectAndInitialize(username, password, ip, port, name, logger)

}

func Run(s *services.Services, host string, port string) {
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", host, port), s.Router))
}
