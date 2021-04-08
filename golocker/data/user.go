package data

import (
	"fmt"

	"github.com/Killian264/YTLocker/golocker/models"
	uuid "github.com/satori/go.uuid"
)

func (d *Data) Create(obj interface{}) error {

	value, ok := obj.(models.User)
	if ok {
		d.CreateUser(&value)
	}

	switch parsed := obj.(type) {
	case nil:
		return fmt.Errorf("object passed to create is nil")
	case models.User:
		return d.CreateUser(&parsed)
	}

	return fmt.Errorf("create does not exist for type")
}

// CreateUser creates a user
// Required Fields:
// Username,
// Password,
// Email,
func (d *Data) CreateUser(user *models.User) error {

	gormDB := d.gormDB

	user.UUID = uuid.NewV4().String()

	result := gormDB.Create(user)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
