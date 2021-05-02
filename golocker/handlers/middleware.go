package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/Killian264/YTLocker/golocker/models"
	"github.com/Killian264/YTLocker/golocker/services"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
)

type ServiceHandler func(w http.ResponseWriter, r *http.Request, s *services.Services) error
type ErrorHandler func(w http.ResponseWriter, r *http.Request) error
type Handler func(w http.ResponseWriter, r *http.Request)

// CreateServiceInjector returns a route wrapper that injects services
func CreateServiceInjector(s *services.Services) func(next ServiceHandler) ErrorHandler {
	return func(next ServiceHandler) ErrorHandler {
		return func(w http.ResponseWriter, r *http.Request) error {
			return next(w, r, s)
		}
	}
}

// CreateErrorHandler returns a route wrapper that handles errors
func CreateErrorHandler(l *log.Logger) func(next ErrorHandler) Handler {
	return func(next ErrorHandler) Handler {
		return func(w http.ResponseWriter, r *http.Request) {
			err := next(w, r)

			if err == nil {
				return
			}

			l.Printf("\nERROR occurred on ROUTE: '%s' \nERROR: '%s'", r.URL, err.Error())

			w.WriteHeader(http.StatusInternalServerError)

			w.Write([]byte("An Error Occurred"))

		}
	}
}

// CreateUserAuthenticator returns a route wrapper that authenticates the user and adds them to the session
func CreateUserAuthenticator(s *services.Services) func(next ServiceHandler) ServiceHandler {
	return func(next ServiceHandler) ServiceHandler {
		return func(w http.ResponseWriter, r *http.Request, s *services.Services) error {

			header := r.Header["Authorization"]

			if len(header) != 1 {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("no authorization header"))
				return nil
			}

			token := header[0]

			user, err := s.User.GetUserFromBearer(token)
			if err != nil {
				return err
			}

			if user == nil {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("invalid authorization header"))
				return nil
			}

			context.Set(r, "user", *user)

			return next(w, r, s)

		}
	}
}

// CreateUserAuthenticator returns a route wrapper that authenticates the user and adds them to the session
func CreatePlaylistAuthenticator(s *services.Services) func(next ServiceHandler) ServiceHandler {
	return func(next ServiceHandler) ServiceHandler {
		return func(w http.ResponseWriter, r *http.Request, s *services.Services) error {

			idStr := mux.Vars(r)["playlist_id"]
			if idStr == "" {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("no playlist id provided"))
				return nil
			}

			id, err := strconv.ParseUint(idStr, 10, 64)
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("invalid playlist id"))
				return nil
			}

			user := GetUserFromRequest(r)

			playlist, err := s.Playlist.Get(&user, id)
			if err != nil {
				return err
			}

			if playlist == nil {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("playlist does not exist"))
				return nil
			}

			context.Set(r, "playlist", *playlist)

			return next(w, r, s)

		}
	}
}

// GetUserFromRequest gets the user from the request, user is set by user authenticator
func GetUserFromRequest(r *http.Request) models.User {
	return context.Get(r, "user").(models.User)
}

// GetPlaylistFromRequest gets the playlist set from the request, playlist is set by playlistauthenticator
func GetPlaylistFromRequest(r *http.Request) models.Playlist {
	return context.Get(r, "playlist").(models.Playlist)
}
