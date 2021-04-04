package data

import (
	"fmt"
	"log"

	"github.com/Killian264/YTLocker/golocker/models"
	uuid "github.com/satori/go.uuid"
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

func (d *Data) GetChannel(channelID string) {

	channel := models.Channel{
		ChannelID: channelID,
	}

	// result := d.gormDB.Joins("JOIN thumbnails ON thumbnails.owner_id = channels.id AND thumbnails.owner_type = 'channels' ").Find(&channel.Thumbnails).Where(&channel).First(&channel)

	result := d.gormDB.Where(&channel).First(&channel)

	result = d.gormDB.Joins("thumbnails").Find(&channel.Thumbnails)

	if result.Error != nil {
		log.Print(result.Error)
	}

	log.Print("Channel ID: ", channel.ChannelID)

	log.Print("Description: ", channel.Description, "\n\n")

	for _, thumbnail := range channel.Thumbnails {
		log.Print(thumbnail.Height)
	}
}

func (d *Data) NewChannel(channel *youtube.Channel) {

	ytThumbnails := []*youtube.Thumbnail{}

	ytThumbnails = append(ytThumbnails, channel.Snippet.Thumbnails.Default)
	ytThumbnails = append(ytThumbnails, channel.Snippet.Thumbnails.Standard)
	ytThumbnails = append(ytThumbnails, channel.Snippet.Thumbnails.Medium)
	ytThumbnails = append(ytThumbnails, channel.Snippet.Thumbnails.High)
	ytThumbnails = append(ytThumbnails, channel.Snippet.Thumbnails.Maxres)

	thumbnails := []models.Thumbnail{}

	for _, thumbnail := range ytThumbnails {
		if thumbnail == nil {
			continue
		}
		thumbnails = append(thumbnails, models.Thumbnail{
			UUID:   uuid.NewV4().String(),
			URL:    thumbnail.Url,
			Width:  uint(thumbnail.Width),
			Height: uint(thumbnail.Width),
		})
	}

	dbChannel := models.Channel{
		UUID:        uuid.NewV4().String(),
		ChannelID:   channel.Id,
		Title:       channel.Snippet.Title,
		Description: channel.Snippet.Description,

		Thumbnails: thumbnails,
	}

	result := d.gormDB.Create(&dbChannel)

	if result.Error != nil {
		log.Print(result.Error)
	}
}

func (d *Data) NewVideo(video *youtube.Video) {

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

	err = gormDB.AutoMigrate(
		&models.User{},
		&models.Playlist{},
		&models.Channel{},
		&models.Video{},
		&models.Thumbnail{},
		&models.SubscriptionRequest{},
	)

	log.Print(err)

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
