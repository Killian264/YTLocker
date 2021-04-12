package user

import (
	"fmt"
	"regexp"

	"github.com/Killian264/YTLocker/golocker/data"
	"github.com/Killian264/YTLocker/golocker/models"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	dataService data.Data
}

func NewUser(data data.Data) *User {
	return &User{
		dataService: data,
	}
}

func (u *User) RegisterUser(user *models.User) error {

	err := validateString(user.Username)

	if err != nil {
		return err
	}

	err = validateString(user.Password)

	if err != nil {
		return err
	}

	err = validateString(user.Email)

	if err != nil {
		return err
	}

	err = judgePasswordStrength(user.Password)

	if err != nil {
		return err
	}

	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	user.Password = string(encryptedPassword)

	result, err := u.dataService.GetUserByEmail(user.Email)

	if err != nil {
		return err
	}

	if result != nil {
		return fmt.Errorf("A user has already registered under that email")
	}

	return u.dataService.CreateUser(user)
}

func validateString(str string) error {
	//TODO: may need separate sanatize function later anyway?
	if str == "" {
		return fmt.Errorf("Registration information cannot be empty")
	}

	re := regexp.MustCompile(`<(.|\n)*?>`)

	result := re.Find([]byte(str))

	if result != nil {
		return fmt.Errorf("Registration information cannot contain: " + string(result))
	}
	return nil
}

func judgePasswordStrength(pass string) error {
	//TODO: agree and expand on password strength requirements, essentially placeholder
	re := regexp.MustCompile(`\d`)

	if len(re.FindAll([]byte(pass), -1)) > 6 {
		return nil
	}

	return fmt.Errorf("Password strength does not meet requirements")
}
