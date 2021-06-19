package handlers

import (
	"net/http"
	"reflect"
	"strconv"

	"github.com/Killian264/YTLocker/golocker/models"
	"github.com/Killian264/YTLocker/golocker/services"
	"github.com/gorilla/mux"
)

// UserInformation not including playlists
// returns user information
func GetVideo(w http.ResponseWriter, r *http.Request, s *services.Services) Response {
	idString := mux.Vars(r)["video_id"]

	id, err := strconv.ParseUint(idString, 10, 64)

	if err != nil {
		return NewResponse(http.StatusBadRequest, nil, "invalid video id")
	}

	video, err := s.Youtube.GetVideo(id)
	if err != nil {
		return BlankResponse(err)
	}

	if reflect.DeepEqual(video, models.Video{}){
		return NewResponse(http.StatusBadRequest, nil, "video does not exist")
	}

	thumbnails, err := s.Youtube.GetAllVideoThumbnails(video)
	if err != nil {
		return BlankResponse(err)
	}

	video.Thumbnails = thumbnails

	return NewResponse(http.StatusOK, video, "")
}

type ChannelItem struct {
	models.Channel
	Videos []uint64
}

func GetChannel(w http.ResponseWriter, r *http.Request, s *services.Services) Response {
	idString := mux.Vars(r)["channel_id"]

	id, err := strconv.ParseUint(idString, 10, 64)
	if err != nil {
		return NewResponse(http.StatusBadRequest, nil, "invalid channel id")
	}

	channel, err := s.Youtube.GetChannel(id)
	if err != nil {
		return BlankResponse(err)
	}

	if reflect.DeepEqual(channel, models.Channel{}){
		return NewResponse(http.StatusBadRequest, nil, "channel does not exist")
	}

	thumbnails, err := s.Youtube.GetAllChannelThumbnails(channel)
	if err != nil {
		return BlankResponse(err)
	}
	
	videos, err := s.Youtube.GetAllChannelVideos(channel)
	if err != nil {
		return BlankResponse(err)
	}
	
	item := ChannelItem{}
	item.Channel = channel
	item.Thumbnails = thumbnails
	item.Videos = videos

	return NewResponse(http.StatusOK, item, "")
}

func SearchChannel(w http.ResponseWriter, r *http.Request, s *services.Services) Response {
	kind := r.URL.Query().Get("kind") // username or id
	query := r.URL.Query().Get("query") // the actual search param

	channelID := query

	if( kind == "username"){
		searchedID, err := s.Youtube.GetChannelIdFromUsername(query)
		if(err != nil){
			s.Logger.Print(err)
			return NewResponse(http.StatusBadRequest, nil, "channel does not exist")
		}

		channelID = searchedID
	}

	channel, err := s.Youtube.NewChannel(channelID)
	if(err != nil){
		return BlankResponse(err)
	}

	thumbnails, err := s.Youtube.GetAllChannelThumbnails(channel)
	if err != nil {
		return BlankResponse(err)
	}

	item := ChannelItem{}
	item.Channel = channel
	item.Thumbnails = thumbnails

	return NewResponse(http.StatusOK, item, "")
}
