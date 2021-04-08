package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"

	"github.com/Killian264/YTLocker/golocker/data"
	"github.com/Killian264/YTLocker/golocker/models"
	"golang.org/x/crypto/bcrypt"
)

//TODO: implement service features
func HandleRegistration(w http.ResponseWriter, r *http.Request, d *data.Data) error {
	//body, err := ioutil.ReadAll(r.Body)
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		return err
	}

	//service stuff

	if user.Username == "" {
		return fmt.Errorf("User cannot register with empty username")
	}
	if user.Password == "" {
		return fmt.Errorf("User cannot register with empty password")
	}
	if user.Email == "" {
		return fmt.Errorf("User cannot register with empty email")
	}

	user.Username = SanitizeString(user.Username)

	user.Password = SanitizeString(user.Password)

	user.Email = SanitizeString(user.Email)

	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	user.Password = string(encryptedPassword)

	//end

	return d.CreateUser(&user)
}

//TODO: Move later
func SanitizeString(str string) string {
	re := regexp.MustCompile(`<(.|\n)*?>`)
	return re.ReplaceAllString(str, "")
}
