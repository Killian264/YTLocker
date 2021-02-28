package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/Killian264/YTLocker/parsers"
	"github.com/Killian264/YTLocker/youtube"
	"gorm.io/gorm/logger"

	"github.com/Killian264/YTLocker/db"

	"github.com/Killian264/YTLocker/models"

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
	data    *db.DB
	logger  *log.Logger
	youtube *youtube.YTService
}

// InitializeRouter Creates Router for app
func (s *Services) InitializeRouter() {

	s.router = mux.NewRouter()

}

// InitializeRoutes creates the routes
func (s *Services) InitializeRoutes() {
	s.router.HandleFunc("/", s.TestHandler)

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

	db := new(db.DB)

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
func (a *Services) TestHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Print("Request Recieved\n")

	hubChallenge := r.URL.Query().Get("hub.challenge")

	hubTopic := r.URL.Query().Get("hub.topic")

	bytes, err := ioutil.ReadAll(r.Body)

	if err != nil {
		panic("error reading body")
	}

	body := string(bytes)

	for k, v := range r.URL.Query() {
		log.Printf("key=%v, value=%v \n", k, v)
	}

	request := models.Request{
		Body:      body,
		Challenge: hubChallenge,
		Topic:     hubTopic,
	}
	a.data.Create(&request)

	fmt.Print("Request Saved\n")

	if hubChallenge == "" {
		hubChallenge = "YOUTUBE"
	}

	fmt.Print("FROM: ", hubChallenge, "\n\n")

	if hubChallenge != "" {
		fmt.Fprintf(w, hubChallenge)
		fmt.Print("Replied to Subscription\n")
	}

	// hook := ParseYTHook(body)

	// client := &http.Client{
	// 	Transport: &transport.APIKey{Key: os.Getenv("YOUTUBE_API_KEY")},
	// }

	// service, err := youtube.New(client)
	// if err != nil {
	// 	log.Fatalf("Error creating new YouTube client: %v", err)
	// }

	// parts := []string{"id", "snippet"}
	// channelService := youtube.NewChannelsService(service)

	// call := channelService.List(parts)
	// call.Id(hook.Video.ChannelID)

	// response, err := call.Do()
	// if err != nil {
	// 	fmt.Print("youtube data api error")
	// }

	// channel := response.Items[0]

	// if response.Items[]

	// Make the API call to YouTube.
	// call := service.Search.List(paramThingies).
	// 	Q(*query).
	// 	MaxResults(*maxResults)
	// response, err := call.Do()
	// handleError(err, "")

}
