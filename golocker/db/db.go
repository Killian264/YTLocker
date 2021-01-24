package db

import (
	"fmt"

	"github.com/Killian264/YTLocker/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	uuid "github.com/satori/go.uuid"
)

type DB struct {
	gormDB *gorm.DB
}

func (db *DB) Initialize(username string, password string, ip string, port string, name string) {

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, ip, port, name)
	// logFile, err := os.OpenFile(
	// 	logFileLoc,
	// 	os.O_APPEND|os.O_CREATE|os.O_WRONLY,
	// 	0644,
	// )

	// if err != nil {
	// 	panic("Error opening or creating database log file.")
	// }

	// logger := logger.New(
	// 	log.New(logFile, "\r\n", log.LstdFlags),
	// 	logger.Config{
	// 		Colorful: true,
	// 		LogLevel: logger.Warn,
	// 	},
	// )

	gormDB, err := gorm.Open(
		mysql.Open(connectionString),
		&gorm.Config{
			// Logger: logger,
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

	db.gormDB = gormDB
}

func (db *DB) Create(obj interface{}) error {

	value, ok := obj.(models.User)
	if ok {
		db.CreateUser(&value)
	}

	switch parsed := obj.(type) {
	case nil:
		return fmt.Errorf("object passed to create is nil")
	case models.User:
		return db.CreateUser(&parsed)
	}

	return fmt.Errorf("create does not exist for type")
}

// CreateUser creates a user
// Required Fields:
// Username,
// Password,
// Email,
func (db *DB) CreateUser(user *models.User) error {

	gormDB := db.gormDB

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
