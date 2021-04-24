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

// SQLiteConnectAndInitalize is for testing only, rand is diffent to work with SQLite ints
func InMemorySQLiteConnect() *Data {

	logger := logger.New(
		log.New(os.Stdout, "Data: ", log.Lshortfile),
		logger.Config{},
	)

	db, err := gorm.Open(sqlite.Open(`file:memdb1?mode=memory`),
		&gorm.Config{Logger: logger},
	)
	if err != nil {
		panic("Error creating db connection")
	}

	data := Data{
		db:   db,
		rand: DataRand(&TestRand{}),
	}

	err = data.createTables()
	if err != nil {
		panic("error initializing db")
	}

	return &data

}

// TODO: figure out what is wrong with clear

// InMemoryMySQLConnect to the in memory test db in the docker compose
func InMemoryMySQLConnect() *Data {

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", "user", "password", "localhost", "9906", "YTLockerDB")

	logger := logger.New(
		log.New(os.Stdout, "Data: ", log.Lshortfile),
		logger.Config{},
	)

	db, err := gorm.Open(mysql.Open(connectionString),
		&gorm.Config{Logger: logger},
	)
	if err != nil {
		panic("Error creating db connection")
	}

	data := Data{
		db:   db,
		rand: DataRand(&ActualRand{}),
	}

	err = data.dropTables()
	if err != nil {
		panic("error initializing db")
	}

	err = data.createTables()
	if err != nil {
		panic("error initializing db")
	}

	return &data
}

// Baisc MYSQL connect
func MySQLConnect(username string, password string, ip string, port string, name string, logger logger.Interface) *Data {

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, ip, port, name)

	db, err := gorm.Open(mysql.Open(connectionString),
		&gorm.Config{Logger: logger},
	)
	if err != nil {
		panic("Error creating db connection")
	}

	data := Data{
		db:   db,
		rand: DataRand(&ActualRand{}),
	}

	err = data.createTables()
	if err != nil {
		panic("error initializing db")
	}

	return &data
}

func (d *Data) createTables() error {

	return d.db.AutoMigrate(
		&models.User{},
		&models.Playlist{},
		&models.Channel{},
		&models.Video{},
		&models.Thumbnail{},
		&models.SubscriptionRequest{},
		&models.YoutubeClientConfig{},
		&models.YoutubeToken{},
	)

}

func (d *Data) dropTables() error {

	return d.db.Migrator().DropTable(
		&models.User{},
		&models.Playlist{},
		&models.Channel{},
		&models.Video{},
		&models.Thumbnail{},
		&models.SubscriptionRequest{},
		&models.YoutubeClientConfig{},
		&models.YoutubeToken{},
	)

}

// Does not work
func (d *Data) clearTables() {
	d.db.Unscoped().Where("1 = 1").Delete(&models.User{})
	d.db.Unscoped().Where("1 = 1").Delete(&models.Playlist{})
	d.db.Unscoped().Where("1 = 1").Delete(&models.Channel{})
	d.db.Unscoped().Where("1 = 1").Delete(&models.Video{})
	d.db.Unscoped().Where("1 = 1").Delete(&models.Thumbnail{})
	d.db.Unscoped().Where("1 = 1").Delete(&models.SubscriptionRequest{})
	d.db.Unscoped().Where("1 = 1").Delete(&models.YoutubeClientConfig{})
	d.db.Unscoped().Where("1 = 1").Delete(&models.YoutubeToken{})
	return
}
