package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Killian264/YTLocker/golocker/models"
	"github.com/Killian264/YTLocker/golocker/services"
)

//TODO: implement service features
func HandleRegistration(w http.ResponseWriter, r *http.Request, s services.Services) error {
	//body, err := ioutil.ReadAll(r.Body)
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		return err
	}

	valid, err := s.User.ValidEmail(user.Email)
	if err != nil {
		return err
	}

	if !valid {
		return fmt.Errorf("Email is invalid or already exists")
	}

	return s.User.RegisterUser(&user)
}
