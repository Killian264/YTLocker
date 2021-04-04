package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Killian264/YTLocker/golocker/data"
	"github.com/Killian264/YTLocker/golocker/youtube"
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
	youtube *youtube.YTService
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
	service := new(youtube.YTService)
	service.InitializeServices(apiKey)
	s.youtube = service
}

// InitializeDatabase creates DB Connection for app
func (s *Services) InitializeDatabase(username string, password string, ip string, port string, name string) {

	db := new(data.Data)

	logger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{},
	)

	db.Initialize(username, password, ip, port, name, logger)

	s.data = db
}

// Run starts the application
func (a *Services) Run(host string, port string) {
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", host, port), a.router))
}
