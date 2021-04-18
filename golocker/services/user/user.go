package user

import (
	"fmt"
	"net/http"
	"regexp"

	"github.com/Killian264/YTLocker/golocker/interfaces"
	"github.com/Killian264/YTLocker/golocker/models"
	"github.com/Killian264/YTLocker/golocker/parsers"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	data interfaces.IUserData
}

func NewUser(data interfaces.IUserData) *User {
	return &User{
		data: data,
	}
}

func (u *User) RegisterUser(user *models.User) error {

	err := parsers.ValidateStringArray([]string{user.Username, user.Email, user.Password})
	if err != nil {
		return err
	}

	err = judgePasswordStrength(user.Password)
	if err != nil {
		return err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)

	return u.data.NewUser(user)
}

func judgePasswordStrength(pass string) error {
	//TODO: agree and expand on password strength requirements, essentially placeholder
	re := regexp.MustCompile(`\d`)

	if len(re.FindAll([]byte(pass), -1)) > 6 {
		return nil
	}

	return fmt.Errorf("Password strength does not meet requirements")
}

func (u *User) ValidEmail(email string) (bool, error) {
	result, err := u.data.GetUserByEmail(email)
	if err != nil {
		return false, err
	}

	return result == nil, nil
}

// TEMP until actual login is implemented
func (u *User) GetUserFromRequest(r *http.Request) (*models.User, error) {

	header := r.Header["Authorization"]

	if len(header) != 1 {
		return nil, fmt.Errorf("No or Invalid Authorization Header")
	}

	token := header[0]

	if token != "TEMP_API_BEARER" {
		return nil, fmt.Errorf("Invalid Bearer")
	}

	return u.data.GetFirstUser()

}
