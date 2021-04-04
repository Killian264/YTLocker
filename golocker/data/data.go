package data

import (
	"fmt"

	"github.com/Killian264/YTLocker/golocker/models"
	"google.golang.org/api/youtube/v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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
		&models.Video{},
		&models.Thumbnail{},
		&models.SubscriptionRequest{},
	)

	d.gormDB = gormDB
}

func (d *Data) SaveSubscription(request *models.SubscriptionRequest) error {
	return nil
}

func (d *Data) GetSubscription(secret string, channelID string) (*models.SubscriptionRequest, error) {
	return nil, nil
}
func (d *Data) ChannelExists(channelID string) (bool, error) {
	return true, nil
}
func (d *Data) SaveVideo(video *youtube.Video) error {
	return nil
}

func (d *Data) InactivateAllSubscriptions() error {
	return nil
}

func (d *Data) GetInactiveSubscription() (*models.SubscriptionRequest, error) {
	return nil, nil
}

func (d *Data) DeleteSubscription(*models.SubscriptionRequest) error {
	return nil
}
