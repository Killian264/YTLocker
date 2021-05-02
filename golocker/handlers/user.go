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
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("failed to parse user information"))
		return nil
	}

	parsed, errorString := parsers.ParseAndValidateUser(user)
	if errorString != "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(errorString))
		return nil
	}

	valid, err := s.User.ValidEmail(parsed.Email)
	if err != nil {
		return err
	}

	if !valid {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("user with email already exists"))
		return nil
	}

	user, err = s.User.Register(user)

	return err
}

func HandleLogin(w http.ResponseWriter, r *http.Request, s *services.Services) error {

	info := login{}

	err := json.NewDecoder(r.Body).Decode(&info)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("failed to parse login information"))
		return nil
	}

	parsed := login{
		Email:    parsers.SanitizeString(info.Email),
		Password: parsers.SanitizeString(info.Password),
	}

	bearer, err := s.User.Login(parsed.Email, parsed.Password)
	if err != nil {
		return err
	}

	w.Write([]byte(bearer))

	return nil

}

func GetUserInformation(w http.ResponseWriter, r *http.Request, s *services.Services) error {

	user := GetUserFromRequest(r)

	user.Password = ""

	marshal, err := json.Marshal(user)
	if err != nil {
		return nil
	}

	w.Write(marshal)

	return nil

}
