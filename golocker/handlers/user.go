package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Killian264/YTLocker/golocker/helpers/parsers"
	"github.com/Killian264/YTLocker/golocker/models"
	"github.com/Killian264/YTLocker/golocker/services"
)

type login struct {
	Email    string
	Password string
}

type BearerResponse struct {
	Bearer string
}

// UserRegister registers a new user,
// user requirement specififed in parsers
// Returns BlankResponse
func UserRegister(w http.ResponseWriter, r *http.Request, s *services.Services) Response {
	user := models.User{}

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		return NewResponse(http.StatusBadRequest, nil, "failed to parse user information")
	}

	parsed, errorString := parsers.ParseAndValidateUser(user)
	if errorString != "" {
		return NewResponse(http.StatusBadRequest, nil, errorString)
	}

	valid, err := s.User.ValidEmail(parsed.Email)
	if err != nil {
		return BlankResponse(err)
	}

	if !valid {
		return NewResponse(http.StatusBadRequest, nil, "user with email already exists")
	}

	user, err = s.User.Register(user)

	return BlankResponse(err)
}

// UserLogin logs in a user,
// Returns BearerResponse
func UserLogin(w http.ResponseWriter, r *http.Request, s *services.Services) Response {
	info := login{}

	err := json.NewDecoder(r.Body).Decode(&info)
	if err != nil {
		return NewResponse(http.StatusBadRequest, nil, "failed to parse login information")
	}

	parsed := login{
		Email:    parsers.SanitizeString(info.Email),
		Password: parsers.SanitizeString(info.Password),
	}

	bearer, err := s.User.Login(parsed.Email, parsed.Password)
	if err != nil {
		return BlankResponse(err)
	}

	return NewResponse(http.StatusOK, BearerResponse{Bearer: bearer}, "")

}

// UserInformation not including playlists
// returns user information
func UserInformation(w http.ResponseWriter, r *http.Request, s *services.Services) Response {
	user := GetUserFromRequest(r)

	user.Password = ""

	return NewResponse(http.StatusOK, user, "")

}
