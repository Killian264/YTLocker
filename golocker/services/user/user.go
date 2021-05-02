package user

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"time"

	"github.com/Killian264/YTLocker/golocker/models"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	data IUserData
}

type IUserData interface {
	GetUserByEmail(email string) (*models.User, error)
	GetUser(ID uint64) (*models.User, error)
	NewUser(user *models.User) error

	NewUserSession(user *models.User, session *models.Session) error
	GetSession(bearer string) (*models.Session, error)
}

func NewUser(data IUserData) *User {
	return &User{
		data: data,
	}
}

// RegisterUser adds a new user given the user information
// TODO: move existance checks on a model to seperate layer called by api
func (u *User) Register(user models.User) (models.User, error) {

	hashed, err := hashPassword(user.Password)
	if err != nil {
		return models.User{}, err
	}

	user.Password = hashed

	err = u.data.NewUser(&user)
	if err != nil {
		return models.User{}, err
	}

	session, err := u.RefreshSession(user)

	user.Session = session

	return user, err
}

// GetUserByID gets a user by id
func (u *User) GetUser(ID uint64) (*models.User, error) {

	return u.data.GetUser(ID)

}

// ValidEmail checks to see if an email is a duplicate
func (u *User) ValidEmail(email string) (bool, error) {
	result, err := u.data.GetUserByEmail(email)
	if err != nil {
		return false, err
	}

	return result == nil, nil
}

// GetUserFromBearer gets the user from the bearer
// Redirects to login for bad bearers should be handled by a middleware
func (u *User) GetUserFromBearer(bearer string) (*models.User, error) {

	session, err := u.data.GetSession(bearer)
	if err != nil {
		return nil, err
	}

	if session == nil || !validSession(*session) {
		return nil, fmt.Errorf("Invalid session")
	}

	return u.data.GetUser(session.UserID)

}

// Login returns the bearer if the user information is correct
func (u *User) Login(email string, password string) (string, error) {

	user, err := u.data.GetUserByEmail(email)
	if err != nil {
		return "", err
	}

	if user == nil {
		return "", fmt.Errorf("could not find user")
	}

	if user.Email != email {
		return "", fmt.Errorf("Invalid")
	}

	if comparePassword(user.Password, password) != nil {
		return "", fmt.Errorf("Invalid password")
	}

	session, err := u.RefreshSession(*user)
	if err != nil {
		return "", err
	}

	return session.Bearer, nil

}

// Refreshes the current user session and returns the new session
func (u *User) RefreshSession(user models.User) (models.Session, error) {

	bearer, err := generateSecret()
	if err != nil {
		return models.Session{}, err
	}

	session := models.Session{
		Bearer: bearer,
	}

	err = u.data.NewUserSession(&user, &session)
	if err != nil {
		return models.Session{}, err
	}

	return session, nil

}

func comparePassword(hash string, password string) error {

	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

}

func hashPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	return string(hashed), err
}

func generateSecret() (string, error) {
	h := sha256.New()
	b := make([]byte, 64)

	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	_, err = h.Write(b)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", b), nil
}

func validSession(session models.Session) bool {

	oneDayAgo := time.Now().AddDate(0, 0, -1)

	return oneDayAgo.Before(session.CreatedAt)

}
