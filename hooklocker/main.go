package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Killian264/YTLocker/hooklocker/ytservice"
	"github.com/gorilla/mux"
)

// var YoutubeSubscribeUrl = "https://pubsubhubbub.appspot.com/subscribe"
/* Main */
func main() {

	host := os.Getenv("GO_API_HOST")
	port := os.Getenv("GO_API_PORT")
	key := os.Getenv("YOUTUBE_API_KEY")

	logger := log.New(os.Stdout, "Subscriber: ", log.Ldate|log.Ltime|log.Lshortfile)
	logger.SetPrefix("YTService: ")

	ytService := ytservice.NewYoutubeService(key, logger)

	router := mux.NewRouter()

	router.HandleFunc("/channel/{channel_id}", func(w http.ResponseWriter, r *http.Request) {
		ChannelHandler(w, r, ytService)
	})

	router.HandleFunc("/video/{video_id}", func(w http.ResponseWriter, r *http.Request) {
		VideoHandler(w, r, ytService)
	})

	http.ListenAndServe(fmt.Sprintf("%s:%s", host, port), router)
}

// ChannelHandler handler to mess around with yt api
func ChannelHandler(w http.ResponseWriter, r *http.Request, ytService *ytservice.YTService) {

	vars := mux.Vars(r)
	channel, err := ytService.GetChannel(vars["channel_id"])
	if err != nil || channel == nil {
		return
	}

	fmt.Print("====================================================================\n")
	fmt.Print(channel.Snippet.Title)
	fmt.Print("\n\n")
	fmt.Print(channel.Snippet.Description)
	fmt.Print("\n\n")
	fmt.Print("====================================================================\n")

}

// VideoHandler handler to mess around with yt api
func VideoHandler(w http.ResponseWriter, r *http.Request, ytService *ytservice.YTService) {

	vars := mux.Vars(r)
	video, err := ytService.GetVideo(vars["video_id"])
	if err != nil || video == nil {
		return
	}

	fmt.Print("====================================================================\n")
	fmt.Print(video.Snippet.Title)
	fmt.Print("\n\n")
	fmt.Print(video.Snippet.Description)
	fmt.Print("\n\n")
	fmt.Print("====================================================================\n")

}
