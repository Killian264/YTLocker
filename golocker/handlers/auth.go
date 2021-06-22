package handlers

import (
	"net/http"
	"strconv"

	"github.com/Killian264/YTLocker/golocker/services"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
)

// CreateUserAuthenticator returns a route wrapper that authenticates the user and adds them to the session
func CreateUserAuthenticator(s *services.Services) func(next ServiceHandler) ServiceHandler {
	return func(next ServiceHandler) ServiceHandler {
		return func(w http.ResponseWriter, r *http.Request, s *services.Services) Response {
			header := r.Header["Authorization"]

			if len(header) != 1 {
				return NewResponse(http.StatusUnauthorized, nil, "no authorization header")
			}

			token := header[0]

			user, err := s.User.GetUserFromBearer(token)
			if user == nil || err != nil {
				return NewResponse(http.StatusUnauthorized, nil, "invalid authorization header")
			}

			context.Set(r, "user", *user)

			return next(w, r, s)
		}
	}
}

// CreateUserAuthenticator returns a route wrapper that authenticates the user and adds them to the session
func CreatePlaylistAuthenticator(s *services.Services) func(next ServiceHandler) ServiceHandler {
	return func(next ServiceHandler) ServiceHandler {
		return func(w http.ResponseWriter, r *http.Request, s *services.Services) Response {
			idStr := mux.Vars(r)["playlist_id"]

			id, err := strconv.ParseUint(idStr, 10, 64)
			if err != nil {
				return NewResponse(http.StatusForbidden, nil, "invalid playlist id")
			}

			user := GetUserFromRequest(r)

			playlist, err := s.Playlist.Get(id)
			if err != nil {
				return BlankResponse(err)
			}

			if user.ID != playlist.UserID {
				return NewResponse(http.StatusForbidden, nil, "playlist does not exist")
			}

			context.Set(r, "playlist", playlist)

			return next(w, r, s)
		}
	}
}

// CreateUserAuthenticator returns a route wrapper that authenticate the admin bearer
func CreateAdminAuthenticator(s *services.Services, bearer string) func(next Handler) Handler {
	return func(next Handler) Handler {
		return func(w http.ResponseWriter, r *http.Request) {
			header := r.Header["Authorization"]

			if len(header) != 1 {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("no authorization header"))
				return
			}

			token := header[0]

			if token != bearer {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("invalid bearer"))
				return
			}

			next(w, r)

		}
	}
}
