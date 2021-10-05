package user

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"time"

	"github.com/Killian264/YTLocker/golocker/data"
	"github.com/Killian264/YTLocker/golocker/models"
)

type User struct {
	data IUserData
}

type IUserData interface {
	GetUserByEmail(email string) (models.User, error)
	GetUser(ID uint64) (models.User, error)
	NewUser(user models.User) (models.User, error)
	UpdateUser(user models.User) (models.User, error)

	NewUserSession(user models.User, session models.Session) (models.Session, error)
	GetSession(bearer string) (models.Session, error)

	NewTemporarySession(bearer string) (models.TemporarySession, error)
	GetTemporarySession(bearer string) (models.TemporarySession, error)
}

func NewUser(data IUserData) *User {
	return &User{
		data: data,
	}
}

// GetUserFromEmail gets a user via email
func (u *User) GetUserFromEmail(email string) (models.User, error) {
	return u.data.GetUserByEmail(email)
}

// GetUserFromBearer gets the user from the bearer
// Redirects to login for bad bearers should be handled by a middleware
func (u *User) GetUserFromBearer(bearer string) (models.User, error) {
	session, err := u.data.GetSession(bearer)
	if err != nil {
		return models.User{}, err
	}

	if !validSession(session) {
		return models.User{}, fmt.Errorf("Invalid session")
	}

	return u.data.GetUser(session.UserID)
}

// Login returns the user, bearer, error
func (u *User) Login(userInfo models.User, bearer string) (models.User, error) {
	tempSession, err := u.data.GetTemporarySession(bearer)
	if err != nil {
		return models.User{}, err
	}

	// older than 5 minutes
	if tempSession.CreatedAt.Before(time.Now().Add(time.Minute * 5 * -1)) {
		return models.User{}, fmt.Errorf("temp session is too old")
	}

	user, err := u.GetUserFromEmail(userInfo.Email)
	if err != nil && err != data.ErrorNotFound {
		return models.User{}, err
	}

	if err == data.ErrorNotFound {
		user, err = u.data.NewUser(userInfo)
		if err != nil {
			return models.User{}, err
		}
	}

	if user.Username != userInfo.Username || user.Picture != userInfo.Picture {
		userInfo.ID = user.ID
		user, err = u.data.UpdateUser(userInfo)
		if err != nil {
			return models.User{}, err
		}
	}

	session, err := u.data.NewUserSession(user, models.Session{Bearer: bearer})
	user.Session = session

	return user, err
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

	session, err = u.data.NewUserSession(user, session)
	if err != nil {
		return models.Session{}, err
	}

	return session, nil
}

// GenerateTemporarySessionBearer creates a temporary session variable needed for login
func (u *User) GenerateTemporarySessionBearer() (string, error) {
	bearer, err := generateSecret()
	if err != nil {
		return "", err
	}

	_, err = u.data.NewTemporarySession(bearer)
	if err != nil {
		return "", err
	}

	return bearer, nil
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
