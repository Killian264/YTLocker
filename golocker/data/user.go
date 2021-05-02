package data

import (
	"github.com/Killian264/YTLocker/golocker/models"
)

func (d *Data) NewUser(user *models.User) error {

	db := d.db

	user.ID = d.rand.ID()

	result := db.Create(user)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (d *Data) GetUser(ID uint64) (*models.User, error) {
	user := models.User{ID: ID}

	result := d.db.Model(user).Preload("Session").First(&user)

	if result.Error != nil || notFound(result.Error) {
		return nil, removeNotFound(result.Error)
	}

	return &user, nil
}

// GetUserByEmail gets user by email
func (d *Data) GetUserByEmail(email string) (*models.User, error) {
	user := models.User{}

	result := d.db.Where("email = ?", email).First(&user)

	if result.Error != nil || notFound(result.Error) {
		return nil, removeNotFound(result.Error)
	}

	return &user, nil
}

// SaveSession saves the session to the user
func (d *Data) NewUserSession(user *models.User, session *models.Session) error {

	session.ID = d.rand.ID()

	return d.db.Model(user).Association("Session").Replace(session)

}

// GetSession returns the session associated with the bearer if it is the current user session
func (d *Data) GetSession(bearer string) (*models.Session, error) {
	passed := models.Session{}

	result := d.db.Where("bearer = ?", bearer).First(&passed)
	if result.Error != nil || notFound(result.Error) {
		return nil, removeNotFound(result.Error)
	}

	user := models.User{
		ID: passed.UserID,
	}

	current := models.Session{}

	err := d.db.Model(&user).Association("Session").Find(&current)
	if err != nil || notFound(err) {
		return nil, removeNotFound(err)
	}

	if passed.ID != current.ID {
		return nil, nil
	}

	return &current, nil
}
