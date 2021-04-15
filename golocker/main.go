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

	logger := log.New(os.Stdout, "Main: ", log.Lshortfile)

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

	logBase := log.New(os.Stdout, "Data: ", log.Lshortfile)

	logger := logger.New(
		logBase,
		logger.Config{},
	)

	db.Initialize(username, password, ip, port, name, logger)

	s.data = db
}

func (a *Services) Run(host string, port string) {
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", host, port), a.router))
}
