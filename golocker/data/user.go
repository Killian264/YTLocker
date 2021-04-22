package data

import (
	"github.com/Killian264/YTLocker/golocker/models"
	uuid "github.com/satori/go.uuid"
)

func (d *Data) NewUser(user *models.User) error {

	gormDB := d.gormDB

	user.UUID = uuid.NewV4().String()

	result := gormDB.Create(user)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (d *Data) GetFirstUser() (*models.User, error) {
	user := models.User{}

	result := d.gormDB.First(&user)

	if result.Error != nil || NotFound(result.Error) {
		return nil, RemoveNotFound(result.Error)
	}

	return &user, nil
}

func (d *Data) GetUserByEmail(email string) (*models.User, error) {
	user := models.User{
		Email: email,
	}

	result := d.gormDB.Where(&user).First(&user)

	if result.Error != nil || NotFound(result.Error) {
		return nil, RemoveNotFound(result.Error)
	}

	return &user, nil
}
