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

//TODO: implement service features
func HandleRegistration(w http.ResponseWriter, r *http.Request, s *services.Services) error {

	user := models.User{}

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.Write([]byte("failed to parse user information"))
		w.WriteHeader(400)
		return nil
	}

	parsed, errorString := parsers.ParseAndValidateUser(user)
	if errorString != "" {
		w.Write([]byte(errorString))
		w.WriteHeader(400)
		return nil
	}

	valid, err := s.User.ValidEmail(parsed.Email)
	if err != nil {
		return err
	}

	if !valid {
		w.Write([]byte("user with email already exists"))
		w.WriteHeader(400)
		return nil
	}

	return s.User.RegisterUser(&user)
}

func HandleLogin(w http.ResponseWriter, r *http.Request, s *services.Services) error {

	info := login{}

	err := json.NewDecoder(r.Body).Decode(&info)
	if err != nil {
		w.Write([]byte("failed to parse login information"))
		w.WriteHeader(400)
		return nil
	}

	parsed := login{
		Email:    parsers.SanitizeString(info.Email),
		Password: parsers.SanitizeString(info.Password),
	}

	bearer, err := s.User.Login(parsed.Email, parsed.Password)

}
