package data

import (
	"fmt"

	"github.com/Killian264/YTLocker/golocker/models"
	uuid "github.com/satori/go.uuid"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

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

	err = gormDB.AutoMigrate(
		&models.User{},
		&models.Playlist{},
		&models.Channel{},
		&models.Video{},
		&models.Thumbnail{},
		&models.SubscriptionRequest{},
		&models.YoutubeClientConfig{},
		&models.YoutubeToken{},
	)

	d.gormDB = gormDB
}

func (d *Data) GetChannel(channelID string) (*models.Channel, error) {

	channel := models.Channel{
		ChannelID: channelID,
	}

	result := d.gormDB.Where(&channel).First(&channel)

	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected != 1 {
		return nil, nil
	}

	result = d.gormDB.Where(models.Thumbnail{OwnerID: channel.ID, OwnerType: "channels"}).Find(&channel.Thumbnails)

	if result.Error != nil {
		return nil, result.Error
	}

	return &channel, nil
}

func (d *Data) NewChannel(channel *models.Channel) error {

	channel.UUID = uuid.NewV4().String()

	for _, thumbnail := range channel.Thumbnails {
		thumbnail.UUID = uuid.NewV4().String()
	}

	result := d.gormDB.Create(&channel)

	return result.Error
}

func (d *Data) NewVideo(video *models.Video, channelID string) error {

	channel := models.Channel{}

	result := d.gormDB.Where("channel_id = ?", channelID).First(&channel)

	if result.Error != nil {
		return result.Error
	}

	video.UUID = uuid.NewV4().String()
	video.ChannelID = channel.ID

	for _, thumbnail := range video.Thumbnails {
		thumbnail.UUID = uuid.NewV4().String()
	}

	result = d.gormDB.Create(&video)

	return result.Error
}

func (d *Data) NewSubscription(request *models.SubscriptionRequest) error {
	request.UUID = uuid.NewV4().String()
	result := d.gormDB.Create(request)
	return result.Error
}

func (d *Data) GetSubscription(secret string, channelID string) (*models.SubscriptionRequest, error) {

	request := models.SubscriptionRequest{}

	result := d.gormDB.Where("channel_id = ? AND secret = ? ", channelID, secret).First(&request)

	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected != 1 {
		return nil, nil
	}

	return &request, result.Error

}

func (d *Data) GetChannelFromYoutubeId(channelID string) (*models.Channel, error) {
	channel := models.Channel{
		ChannelID: channelID,
	}

	result := d.gormDB.Where(&channel).First(&channel)

	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected != 1 {
		return nil, nil
	}

	return &channel, nil
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

// IPLAYLISTDATA
func (d *Data) NewYoutubeClientConfig(config *models.YoutubeClientConfig) error {
	config.UUID = uuid.NewV4().String()

	result := d.gormDB.Create(&config)

	return result.Error
}
func (d *Data) GetYoutubeClientConfigByClientID(clientID string) (*models.YoutubeClientConfig, error) {
	return nil, nil
}
func (d *Data) NewYoutubeToken(token *models.YoutubeToken) error {
	token.UUID = uuid.NewV4().String()

	result := d.gormDB.Create(&token)

	return result.Error
}
func (d *Data) GetYoutubeTokenByAccessToken(accessToken string) (*models.YoutubeToken, error) {
	return nil, nil
}

func (d *Data) GetFirstYoutubeClientConfig() (*models.YoutubeClientConfig, error) {
	config := models.YoutubeClientConfig{}

	result := d.gormDB.First(&config)

	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, nil
	}

	return &config, nil
}
func (d *Data) GetFirstYoutubeToken() (*models.YoutubeToken, error) {
	token := models.YoutubeToken{}

	result := d.gormDB.First(&token)

	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, nil
	}

	return &token, nil
}
