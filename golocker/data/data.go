package data

import (
	"fmt"
	"log"
	"os"

	"github.com/Killian264/YTLocker/golocker/models"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Data struct {
	db   *gorm.DB
	rand DataRand
}

// SQLiteConnectAndInitalize is for testing only, rand is diffent to work with SQLite ints
func SQLiteConnectAndInitalize() *Data {

	logger := logger.New(
		log.New(os.Stdout, "Data: ", log.Lshortfile),
		logger.Config{},
	)

	db, err := gorm.Open(sqlite.Open(`file:memdb1?mode=memory`), &gorm.Config{
		Logger: logger,
	})

	if err != nil {
		panic("Error creating db connection")
	}

	data := Data{
		db:   db,
		rand: DataRand(&TestRand{}),
	}

	err = data.initalize()

	if err != nil {
		panic("error initializing db")
	}

	return &data

}

func MySQLConnectAndInitialize(username string, password string, ip string, port string, name string, logger logger.Interface) *Data {

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, ip, port, name)

	db, err := gorm.Open(
		mysql.Open(connectionString),
		&gorm.Config{
			Logger: logger,
		},
	)

	if err != nil {
		panic("Error creating db connection")
	}

	data := Data{
		db:   db,
		rand: DataRand(&ActualRand{}),
	}

	err = data.initalize()

	if err != nil {
		panic("error initializing db")
	}

	return &data
}

func (d *Data) initalize() error {

	err := d.db.AutoMigrate(
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
		YoutubeID: channelID,
	}

	result := d.db.Where(&channel).First(&channel)

	if result.Error != nil || notFound(result.Error) {
		return nil, removeNotFound(result.Error)
	}

	result = d.db.Where(models.Thumbnail{OwnerID: channel.ID, OwnerType: "channels"}).Find(&channel.Thumbnails)

	if result.Error != nil || notFound(result.Error) {
		return nil, removeNotFound(result.Error)
	}

	return &channel, nil
}

func (d *Data) NewChannel(channel *models.Channel) error {

	channel.ID = d.rand.ID()

	for _, thumbnail := range channel.Thumbnails {
		thumbnail.ID = d.rand.ID()
	}

	result := d.db.Create(&channel)

	return result.Error
}

func (d *Data) NewVideo(video *models.Video, channelID string) error {

	channel := models.Channel{}

	result := d.db.Where("channel_id = ?", channelID).First(&channel)

	if result.Error != nil {
		return result.Error
	}

	video.ID = d.rand.ID()
	video.ChannelID = channel.ID

	for _, thumbnail := range video.Thumbnails {
		thumbnail.ID = d.rand.ID()
	}

	result = d.db.Create(&video)

	return result.Error
}

func (d *Data) GetChannelFromYoutubeID(channelID string) (*models.Channel, error) {
	channel := models.Channel{
		YoutubeID: channelID,
	}

	result := d.db.Where(&channel).First(&channel)

	if result.Error != nil || notFound(result.Error) {
		return nil, removeNotFound(result.Error)
	}

	return &channel, nil
}
