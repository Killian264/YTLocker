package data

import (
	"fmt"

	"github.com/Killian264/YTLocker/golocker/models"
)

func (d *Data) NewUser(user models.User) (models.User, error) {
	db := d.db

	user.ID = d.rand.ID()

	result := db.Create(&user)

	return user, result.Error
}

func (d *Data) GetUser(ID uint64) (models.User, error) {
	user := models.User{ID: ID}

	result := d.db.Model(user).Preload("Session").First(&user)

	return user, result.Error
}

func (d *Data) UpdateUser(user models.User) (models.User, error) {
	result := d.db.Model(&user).Select("Username", "Picture").Updates(user)

	if result.Error != nil {
		return models.User{}, result.Error
	}

	return d.GetUser(user.ID)
}

// GetUserByEmail gets user by email
func (d *Data) GetUserByEmail(email string) (models.User, error) {
	user := models.User{}

	result := d.db.Where("email = ?", email).First(&user)

	return user, result.Error
}

// SaveSession saves the session to the user
func (d *Data) NewUserSession(user models.User, session models.Session) (models.Session, error) {
	session.ID = d.rand.ID()

	err := d.db.Model(&user).Association("Session").Replace(&session)

	return session, err
}

// GetSession returns the session associated with the bearer if it is the current user session
func (d *Data) GetSession(bearer string) (models.Session, error) {
	passed := models.Session{}

	result := d.db.Where("bearer = ?", bearer).First(&passed)
	if result.Error != nil {
		return models.Session{}, result.Error
	}

	user := models.User{
		ID: passed.UserID,
	}

	current := models.Session{}

	err := d.db.Model(&user).Association("Session").Find(&current)
	if err != nil {
		return models.Session{}, err
	}

	if passed.ID != current.ID {
		return models.Session{}, fmt.Errorf("session is not current session")
	}

	return current, nil
}
