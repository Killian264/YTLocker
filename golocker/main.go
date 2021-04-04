package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/Killian264/YTLocker/golocker/data"
	"github.com/Killian264/YTLocker/golocker/parsers"
	"github.com/Killian264/YTLocker/golocker/youtube"
	"gorm.io/gorm/logger"

	"github.com/Killian264/YTLocker/golocker/models"

	"github.com/gorilla/mux"
)

/* Main */
func main() {

	log.SetPrefix("YTLocker: ")

	log.Print("Starting...")

	s := Services{}

	log.Print("Creating Router...")

	s.InitializeRouter()

	log.Print("Creating Routes...")

	s.InitializeRoutes()

	log.Print("Creating Data Service...")

	s.InitializeDatabase(
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_TCP_PORT"),
		os.Getenv("MYSQL_DATABASE"),
	)

	log.Print("Creating Youtube Service...")

	s.InitializeYTService(
		os.Getenv("YOUTUBE_API_KEY"),
	)

	log.Print("Running...")

	s.Run(
		os.Getenv("GO_API_HOST"),
		os.Getenv("GO_API_PORT"),
	)

	log.Print("Exiting...")

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
	s.router.HandleFunc("/", s.SubscriptionHandler)

	s.router.HandleFunc("/channel/{channel_id}", s.ChannelHandler)
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

// ChannelHandler handler to mess around with yt api
func (s *Services) ChannelHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	channelId := vars["channel_id"]

	ytChannel := s.youtube.GetChannelById(channelId)

	fmt.Print(ytChannel.Snippet.Title)
	fmt.Print("\n\n")
	fmt.Print(ytChannel.Snippet.Description)
	fmt.Print("\n\n")

	channelJSON, err := json.Marshal(ytChannel)
	if err != nil {
		panic("error creating object json")
	}
	fmt.Println(string(channelJSON))

	fmt.Print("\n\n\n\n")

	channel := parsers.ParseChannelIntoDBModels(ytChannel)

	fmt.Print(channel.Title)
	fmt.Print("\n\n")
	fmt.Print(channel.Description)
}

// TestHandler recieves a yt hook and parses it
func (a *Services) SubscriptionHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Print("Subcription Request Recieved\n")

	challenge := r.URL.Query().Get("hub.challenge")

	bytes, err := ioutil.ReadAll(r.Body)

	if err != nil {
		panic("error reading body")
	}

	body := string(bytes)

	if challenge == "" {
		// parse data
		// send to service
		// return
	}

	leaseSeconds, err := strconv.Atoi(r.URL.Query().Get("hub.lease_seconds"))

	if err != nil {
		panic("Failed to parse new subscription lease seconds")
	}

	for k, v := range r.URL.Query() {
		log.Printf("key=%v, value=%v \n", k, v)
	}

	request := models.YTHookTopic{
		Body:         body,
		Challenge:    r.URL.Query().Get("hub.challenge"),
		Topic:        r.URL.Query().Get("hub.topic"),
		LeaseSeconds: leaseSeconds,
		Token:        r.URL.Query().Get("hub.verify_token"),
	}
	fmt.Printf("Request: %v \n", request)

	fmt.Print("Request Saved\n")

}

// curl "https://pubsubhubbub.appspot.com/subscribe" ^
//   -H "authority: pubsubhubbub.appspot.com" ^
//   -H "cache-control: max-age=0" ^
//   -H "upgrade-insecure-requests: 1" ^
//   -H "origin: https://pubsubhubbub.appspot.com" ^
//   -H "content-type: application/x-www-form-urlencoded" ^
//   -H "user-agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.190 Safari/537.36" ^
//   -H "accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9" ^
//   -H "sec-fetch-site: same-origin" ^
//   -H "sec-fetch-mode: navigate" ^
//   -H "sec-fetch-user: ?1" ^
//   -H "sec-fetch-dest: document" ^
//   -H "referer: https://pubsubhubbub.appspot.com/subscribe" ^
//   -H "accept-language: en-US,en;q=0.9" ^
//   --data-raw "hub.callback=https^%^3A^%^2F^%^2Fdroplet.ytlocker.com^%^2F&hub.topic=https^%^3A^%^2F^%^2Fwww.youtube.com^%^2Fxml^%^2Ffeeds^%^2Fvideos.xml^%^3Fchannel_id^%^3DUCfJvn8LAFkRRPJNt8tTJumA&hub.verify=async&hub.mode=subscribe&hub.verify_token=&hub.secret=&hub.lease_seconds=691200" ^
//   --compressed

// subscribe
// recieve subscribe
// create playlist
//
