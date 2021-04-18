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
