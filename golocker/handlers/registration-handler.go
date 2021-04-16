package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Killian264/YTLocker/golocker/models"
	"github.com/Killian264/YTLocker/golocker/services/user"
)

//TODO: implement service features
func HandleRegistration(w http.ResponseWriter, r *http.Request, u *user.User) error {
	//body, err := ioutil.ReadAll(r.Body)
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		return err
	}

	return u.RegisterUser(&user)
}
