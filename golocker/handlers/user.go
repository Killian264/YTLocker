package handlers

import (
	"net/http"

	"github.com/Killian264/YTLocker/golocker/services"
)

type BearerResponse struct {
	Bearer string
}

// UserSessionCreate creates a session required for oauth
func UserSessionCreate(w http.ResponseWriter, r *http.Request, s *services.Services) Response {
	bearer, err := s.User.GenerateTemporarySessionBearer()
	if err != nil {
		return NewResponse(http.StatusBadRequest, nil, "failed to create temp session")
	}

	return NewResponse(http.StatusOK, BearerResponse{Bearer: bearer}, "")
}

// UserSessionRefresh refreshes a user session while they are logged in
func UserSessionRefresh(w http.ResponseWriter, r *http.Request, s *services.Services) Response {
	user := GetUserFromRequest(r)

	session, err := s.User.RefreshSession(user)
	if err != nil {
		return NewResponse(http.StatusBadRequest, nil, "failed to refresh session")
	}

	return NewResponse(http.StatusOK, BearerResponse{Bearer: session.Bearer}, "")
}

// UserInformation includes session information, does not include playlist information
func UserInformation(w http.ResponseWriter, r *http.Request, s *services.Services) Response {
	user := GetUserFromRequest(r)

	return NewResponse(http.StatusOK, user, "")
}
