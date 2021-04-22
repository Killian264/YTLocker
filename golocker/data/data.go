package data

import (
	"fmt"
	"log"
	"os"

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

	logger := logger.New(
		log.New(os.Stdout, "Data: ", log.Lshortfile),
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

	if result.Error != nil || NotFound(result.Error) {
		return nil, RemoveNotFound(result.Error)
	}

	result = d.gormDB.Where(models.Thumbnail{OwnerID: channel.ID, OwnerType: "channels"}).Find(&channel.Thumbnails)

	if result.Error != nil || NotFound(result.Error) {
		return nil, RemoveNotFound(result.Error)
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

func (d *Data) GetChannelFromYoutubeID(channelID string) (*models.Channel, error) {
	channel := models.Channel{
		ChannelID: channelID,
	}

	result := d.gormDB.Where(&channel).First(&channel)

	if result.Error != nil || NotFound(result.Error) {
		return nil, RemoveNotFound(result.Error)
	}

	return &channel, nil
}
