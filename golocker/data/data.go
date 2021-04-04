package data

import (
	"fmt"

	"github.com/Killian264/YTLocker/golocker/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	uuid "github.com/satori/go.uuid"
)

// Request
type Request struct {
	ID           int
	ChannelID    string
	LeaseSeconds int
	Secret       string
	Mode         string
	Active       bool
}

type Data struct {
	gormDB *gorm.DB
}

func (d *Data) Initialize(username string, password string, ip string, port string, name string, logger logger.Interface) {

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, ip, port, name)

	gormDB, err := gorm.Open(
		mysql.Open(connectionString),
		&gorm.Config{
			Logger: logger,
		},
	)

	if err != nil {
		panic("Error creating db connection")
	}

	gormDB.AutoMigrate(
		&models.User{},
		&models.Playlist{},
		&models.Channel{},
		&models.Subscription{},
		&models.Video{},
		&models.Thumbnail{},
		&models.ThumbnailType{},
		&models.Request{},
	)

	d.gormDB = gormDB
}

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

	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(encryptedPassword)
	user.UUID = uuid.NewV4().String()

	result := gormDB.Create(user)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
