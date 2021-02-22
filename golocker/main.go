package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/Killian264/YTLocker/parsers"

	"github.com/Killian264/YTLocker/db"

	"github.com/Killian264/YTLocker/models"

	"github.com/gorilla/mux"

	"google.golang.org/api/googleapi/transport"
	"google.golang.org/api/youtube/v3"
)

/* Main */
func main() {

	log.Print("YTLocker: Starting...")

	a := App{}

	log.Print("YTLocker: Creating Router...")

	a.InitializeRouter()

	log.Print("YTLocker: Creating DB Connection...")

	a.InitializeDatabase(
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_TCP_PORT"),
		os.Getenv("MYSQL_DATABASE"),
	)

	log.Print("YTLocker: Initalizing YTService...")

	a.InitializeYTService(
		os.Getenv("YOUTUBE_API_KEY"),
	)

	log.Print("YTLocker: Running...")

	a.Run(
		os.Getenv("GO_API_HOST"),
		os.Getenv("GO_API_PORT"),
	)

	log.Print("YTLocker: Exiting...")

}

// App contains services for handlers
type App struct {
	Router    *mux.Router
	DB        *db.DB
	Logger    *log.Logger
	YTService *YTService
}

// InitializeRouter Creates Router for app
func (a *App) InitializeRouter() {

	a.Router = mux.NewRouter()

	a.InitializeRoutes()
}

// InitializeYTService Creates YTService for app
func (a *App) InitializeYTService(apiKey string) {
	service := new(YTService)
	service.InitializeServices(apiKey)
	a.YTService = service
}

// InitializeDatabase creates DB Connection for app
func (a *App) InitializeDatabase(username string, password string, ip string, port string, name string) {

	db := new(db.DB)

	db.Initialize(username, password, ip, port, name)

	a.DB = db
}

// Run starts the application
func (a *App) Run(host string, port string) {
	log.Fatal(http.ListenAndServe(":"+port, a.Router))
}

// InitializeRoutes creates the routes
func (a *App) InitializeRoutes() {
	a.Router.HandleFunc("/", a.TestHandler)

	a.Router.HandleFunc("/channel/{channel_id}", a.ChannelHandler)
}

// ChannelHandler handler to mess around with yt api
func (a *App) ChannelHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	channelId := vars["channel_id"]

	ytChannel := a.YTService.GetChannelById(channelId)

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
func (a *App) TestHandler(w http.ResponseWriter, r *http.Request) {

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
		Body:         body,
		HubChallenge: hubChallenge,
		HubTopic:     hubTopic,
	}
	a.DB.Create(&request)

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

// YTService wraps the yt api for common commands to be used by the application
type YTService struct {
	youtubeService *youtube.Service
	channelService *youtube.ChannelsService
}

// InitializeServices creates the yt service
func (s *YTService) InitializeServices(apiKey string) {

	youtubeClient := &http.Client{
		Transport: &transport.APIKey{Key: apiKey},
	}

	youtubeService, err := youtube.New(youtubeClient)

	if err != nil {
		panic("error creating youtube service")
	}

	s.youtubeService = youtubeService

	channelService := youtube.NewChannelsService(youtubeService)

	s.channelService = channelService

}

// Before this is run there should be a check to see if a channel exists already with this id
// To do this the best course of action would be to make a interface for the db object and begin building a db layer
func (s *YTService) GetChannelById(id string) *youtube.Channel {
	log.Print("Searching for channel with ID: ", id, "\n")

	parts := []string{"id", "snippet"}
	call := s.channelService.List(parts)
	call.Id(id)

	response, err := call.Do()
	if err != nil {
		log.Print("youtube data api error: ", err)
	}

	if response == nil || len(response.Items) == 0 {
		log.Print("Channel with id: ", id, " not found\n")
		return nil
	}

	channel := response.Items[0]

	log.Print("Found channel with ID: ", id, "\n")
	log.Print("Channel Title: ", channel.Snippet.Title, "\n")

	return channel
}

// response2 := playlistsList(s.youtubeService, "snippet,contentDetails", id)

// for _, playlist := range response2.Items {
// 	playlistId := playlist.Id
// 	playlistTitle := playlist.Snippet.Title

// 	// Print the playlist ID and title for the playlist resource.
// 	fmt.Println(playlistId, ": ", playlistTitle)
// }
func playlistsList(service *youtube.Service, part string, channelId string) *youtube.PlaylistListResponse {
	call := service.Playlists.List(strings.Split(part, ","))
	if channelId != "" {
		call = call.ChannelId(channelId)
	}
	call = call.MaxResults(2)
	response, err := call.Do()
	if err != nil {
		fmt.Print("\n\nyoutube data api error\n\n", err)
	}
	return response
}
