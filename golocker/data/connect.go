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

// SQLiteConnectAndInitalize creates an in memory SQLite db.
// For testing purposes only.
// SQLite supports up to 128 int keys, referential integrity is not checked.
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

// InMemoryMySQLConnect connects the the in memory test db
// TODO: figure out why clear does not work
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

// MySQLConnect connects to a mysql db
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

// TODO: fix
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
