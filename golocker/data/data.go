package data

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/Killian264/YTLocker/golocker/models"
	uuid "github.com/satori/go.uuid"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Data struct {
	gormDB *gorm.DB
}

func SQLiteConnectAndInitalize() *Data {

	logBase := log.New(os.Stdout, "Data: ", log.Lshortfile)

	logger := logger.New(
		logBase,
		logger.Config{},
	)

	gormDB, err := gorm.Open(sqlite.Open(`file:memdb1?mode=memory`), &gorm.Config{
		Logger: logger,
	})

	if err != nil {
		panic("Error creating db connection")
	}

	data := Data{
		gormDB: gormDB,
	}

	err = data.initalize()

	if err != nil {
		panic("error initializing db")
	}

	return &data

}

func MySQLConnectAndInitialize(username string, password string, ip string, port string, name string, logger logger.Interface) *Data {

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

	data := Data{
		gormDB: gormDB,
	}

	err = data.initalize()

	if err != nil {
		panic("error initializing db")
	}

	return &data
}

func (d *Data) initalize() error {

	err := d.gormDB.AutoMigrate(
		&models.User{},
		&models.Playlist{},
		&models.Channel{},
		&models.Video{},
		&models.Thumbnail{},
		&models.SubscriptionRequest{},
		&models.YoutubeClientConfig{},
		&models.YoutubeToken{},
	)

	return err
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

	if result.RowsAffected == 0 {
		return nil, nil
	}

	return &request, result.Error

}

func (d *Data) GetSubscriptionFromChannelID(channelID string) (*models.SubscriptionRequest, error) {

	request := models.SubscriptionRequest{}

	result := d.gormDB.Where("channel_id = ?", channelID).First(&request)

	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
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

	if result.RowsAffected == 0 {
		return nil, nil
	}

	return &channel, nil
}

func (d *Data) InactivateAllSubscriptions() error {

	result := d.gormDB.Model(&models.SubscriptionRequest{}).Where(&models.SubscriptionRequest{Active: true}).Update("active", false)

	return result.Error

}

func (d *Data) GetInactiveSubscription() (*models.SubscriptionRequest, error) {

	sub := models.SubscriptionRequest{}

	result := d.gormDB.Where("active = false").First(&sub)

	if result.Error != nil && !strings.Contains(result.Error.Error(), "record not found") {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, nil
	}

	return &sub, nil
}
func (d *Data) DeleteSubscription(sub *models.SubscriptionRequest) error {

	result := d.gormDB.Where(&models.SubscriptionRequest{UUID: sub.UUID}).Delete(&models.SubscriptionRequest{UUID: sub.UUID})

	return result.Error

}

func (d *Data) GetUserByEmail(email string) (*models.User, error) {
	user := models.User{
		Email: email,
	}

	result := d.gormDB.Where(&user).First(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, nil
	}

	return &user, nil
}

// IPLAYLISTDATA
func (d *Data) NewYoutubeClientConfig(config *models.YoutubeClientConfig) error {
	config.UUID = uuid.NewV4().String()

	result := d.gormDB.Create(&config)

	return result.Error
}

func (d *Data) NewYoutubeToken(token *models.YoutubeToken) error {
	token.UUID = uuid.NewV4().String()

	result := d.gormDB.Create(&token)

	return result.Error
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

func (d *Data) GetFirstUser() (*models.User, error) {
	user := models.User{}

	result := d.gormDB.First(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, nil
	}

	return &user, nil
}
