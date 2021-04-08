package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Killian264/YTLocker/golocker/data"
	"github.com/Killian264/YTLocker/golocker/models"
)

//TODO: implement service features
func HandleRegistration(w http.ResponseWriter, r *http.Request, d *data.Data) error {
	//body, err := ioutil.ReadAll(r.Body)
	var usr models.User
	err := json.NewDecoder(r.Body).Decode(&usr)

	if err != nil {
		return err
	}

	//service stuff
	if usr.Username == "" {
		return fmt.Errorf("Username invlaid: %s", usr.Username)
	}
	if usr.Password == "" {
		return fmt.Errorf("Password invlaid: %s", usr.Password)
	}
	if usr.Email == "" {
		return fmt.Errorf("Email invlaid: %s", usr.Email)
	}
	//end

	return d.Create(usr)
}
