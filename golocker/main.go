package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/Killian264/YTLocker/golocker/data"
	"github.com/Killian264/YTLocker/golocker/helpers/parsers"
	"github.com/Killian264/YTLocker/golocker/models"
	"github.com/Killian264/YTLocker/golocker/services"
	"github.com/Killian264/YTLocker/golocker/services/subscribe"
	"github.com/Killian264/YTLocker/golocker/services/ytmanager"
	"github.com/Killian264/YTLocker/golocker/services/ytservice"
	"gorm.io/gorm/logger"

	muxhandler "github.com/gorilla/handlers"

	"github.com/gorilla/mux"
)

func main() {

	logger := log.New(os.Stdout, "Main: ", log.Lshortfile)

	logger.Println("----------------------------")

	services := NewServices(logger)

	logger.Println("----------------------------")

	Run(
		services,
		os.Getenv("GO_API_HOST"),
		os.Getenv("GO_API_PORT"),
	)

}

func NewServices(logger *log.Logger) *services.Services {

	services := &services.Services{
		Logger: logger,
	}

	_ = InitalizePlaylistService()

	ytservice := InitializeYTService(
		os.Getenv("YOUTUBE_API_KEY"),
	)

	services.Router = InitializeRouter()

	services.Data = InitializeDatabase(
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_TCP_PORT"),
		os.Getenv("MYSQL_DATABASE"),
	)

	ReadInSecrets(
		services.Data,
		"secrets/",
	)

	services.Youtube = InitalizeYoutubeManager(
		services.Data,
		ytservice,
	)

	services.Subscribe = InitalizeSubscribeService(
		services.Data,
		services.Youtube,
		os.Getenv("GO_API_URL"),
	)

	InitializeRoutes(services.Router)

	return services
}

func InitalizePlaylistService() *ytservice.YTPlaylist {

	return &ytservice.YTPlaylist{}

}

func InitalizeYoutubeManager(data ytmanager.IYoutubeManagerData, yt ytmanager.IYTService) *ytmanager.YoutubeManager {
	service := ytmanager.NewYoutubeManager(
		data,
		yt,
	)

	return service
}

func InitalizeSubscribeService(data subscribe.ISubscriptionData, yt subscribe.IYoutubeManager, appURL string) *subscribe.Subscriber {

	service := subscribe.NewSubscriber(data, yt)

	service.SetYTPubSubUrl("https://pubsubhubbub.appspot.com/subscribe")
	service.SetSubscribeUrl(appURL, "/subscribe/{secret}")

	return service

}

// InitializeRouter Creates Router for app
func InitializeRouter() *mux.Router {

	router := mux.NewRouter()

	router.Use(muxhandler.RecoveryHandler())

	// s.Router.Use(handlers.LoggingMiddlewareTest)

	return router

}

// InitializeRoutes creates the routes
func InitializeRoutes(router *mux.Router) {

	// logger := log.New(os.Stdout, "Han: ", log.Lshortfile)

	// ServiceInjector := handlers.CreateServiceInjector(s)

	// ErrorHandler := handlers.CreateErrorHandler(logger)

	// s.Router.HandleFunc("/subscribe/{secret}", ErrorHandler(ServiceInjector(handlers.HandleYoutubePush)))

}

// InitializeYTService Creates YTService for app
func InitializeYTService(apiKey string) *ytservice.YTService {

	return ytservice.NewYoutubeService(apiKey)

}

func ReadInSecrets(data *data.Data, secretsPath string) {

	clientData, err := readInClientSecret(fmt.Sprintf("%s%s", secretsPath, "client_secret.json"))
	if err != nil {
		panic(err)
	}

	tokenData, err := readInAccessToken(fmt.Sprintf("%s%s", secretsPath, "access_secret.json"))
	if err != nil {
		panic(err)
	}

	err = data.NewYoutubeClientConfig(&clientData)
	if err != nil && !strings.Contains(err.Error(), "Duplicate entry") {
		panic(err)
	}

	err = data.NewYoutubeToken(&tokenData)
	if err != nil && !strings.Contains(err.Error(), "Duplicate entry") {
		panic(err)
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
func InitializeDatabase(username string, password string, ip string, port string, name string) *data.Data {

	logBase := log.New(os.Stdout, "Data: ", log.Lshortfile)

	logger := logger.New(
		logBase,
		logger.Config{},
	)

	return data.MySQLConnect(username, password, ip, port, name, logger)

}

func Run(s *services.Services, host string, port string) {
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", host, port), s.Router))
}
