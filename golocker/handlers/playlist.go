package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Killian264/YTLocker/golocker/helpers/parsers"
	"github.com/Killian264/YTLocker/golocker/models"
	"github.com/Killian264/YTLocker/golocker/services"
	"github.com/gorilla/mux"
)

// HandleSubscriptionNoError handles a new subscription request wrap in a middleware that handles errors
func PlaylistCreate(w http.ResponseWriter, r *http.Request, s *services.Services) Response {
	user := GetUserFromRequest(r)

	playlist := models.Playlist{}

	err := json.NewDecoder(r.Body).Decode(&playlist)
	if err != nil {
		return BlankResponse(err)
	}

	playlist = parsers.ParsePlaylist(playlist)

	created, err := s.Playlist.New(playlist, user)
	if err != nil {
		return BlankResponse(err)
	}

	return NewResponse(http.StatusOK, created, "")
}

type PlaylistListItem struct {
	models.Playlist
	Channels []uint64
	Videos []uint64
}

// HandleSubscriptionNoError handles a new subscription request wrap in a middleware that handles errors
func PlaylistList(w http.ResponseWriter, r *http.Request, s *services.Services) Response {
	user := GetUserFromRequest(r)

	playlists, err := s.Playlist.GetAllUserPlaylists(user)
	if err != nil {
		return BlankResponse(err)
	}

	items := []PlaylistListItem{}

	for _, playlist := range(playlists) {

		var item PlaylistListItem

		item.Playlist = playlist

		thumbnails, err := s.Playlist.GetAllThumbnails(playlist)
		if err != nil {
			return BlankResponse(err)
		}
		channels, err := s.Playlist.GetAllChannels(playlist)
		if err != nil {
			return BlankResponse(err)
		}
		videos, err := s.Playlist.GetAllVideos(playlist)
		if err != nil {
			return BlankResponse(err)
		}

		item.Thumbnails = thumbnails
		item.Channels = channels
		item.Videos = videos

		items = append(items, item)
	}

	return NewResponse(http.StatusOK, items, "")
}

func PlaylistAddSubscription(w http.ResponseWriter, r *http.Request, s *services.Services) Response {
	playlist := GetPlaylistFromRequest(r)

	channelID := mux.Vars(r)["channel_id"]

	channel, err := s.Youtube.NewChannel(channelID)
	if err != nil {
		return BlankResponse(err)
	}

	_, err = s.Subscribe.Subscribe(&channel)
	if err != nil {
		return BlankResponse(err)
	}

	err = s.Playlist.Subscribe(playlist, channel)
	return BlankResponse(err)
}

func PlaylistRemoveSubscription(w http.ResponseWriter, r *http.Request, s *services.Services) Response {
	playlist := GetPlaylistFromRequest(r)

	channelID := mux.Vars(r)["channel_id"]

	channel, err := s.Youtube.GetChannelByID(channelID)
	if err != nil {
		return BlankResponse(err)
	}

	err = s.Playlist.Unsubscribe(playlist, channel)
	return BlankResponse(err)
}

func PlaylistLatestVideos(w http.ResponseWriter, r *http.Request, s *services.Services) Response {
	user := GetUserFromRequest(r)

	videos, err := s.Playlist.GetLastestPlaylistVideos(user)
	if( err != nil ){
		return BlankResponse(err)
	}

	return NewResponse(http.StatusOK, videos, "")
}

func PlaylistDelete(w http.ResponseWriter, r *http.Request, s *services.Services) Response {
	playlist := GetPlaylistFromRequest(r)

	err := s.Playlist.Delete(playlist)

	return BlankResponse(err)
}