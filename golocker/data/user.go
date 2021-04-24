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

func (d *Data) GetUserByID(ID uint64) (*models.User, error) {
	user := models.User{ID: ID}

	result := d.db.First(&user)

	if result.Error != nil || notFound(result.Error) {
		return nil, removeNotFound(result.Error)
	}

	return &user, nil
}

func (d *Data) GetFirstUser() (*models.User, error) {
	user := models.User{}

	result := d.db.First(&user)

	if result.Error != nil || notFound(result.Error) {
		return nil, removeNotFound(result.Error)
	}

	return &user, nil
}

func (d *Data) GetUserByEmail(email string) (*models.User, error) {
	user := models.User{
		Email: email,
	}

	result := d.db.Where(&user).First(&user)

	if result.Error != nil || notFound(result.Error) {
		return nil, removeNotFound(result.Error)
	}

	return &user, nil
}
