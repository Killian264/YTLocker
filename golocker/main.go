package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

/* Helper Functions */
func SetENV(location string) {
	err := godotenv.Load(location)

	if err != nil {
		panic("Error setting ENV.")
	}
}

// TODO: ADD data of playlist managed by ytlocker

type User struct {
	gorm.Model
	ID int `gorm:"primaryKey"`
	// UUID     uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	UserName string
	Email    string
	Password string
	Salt     string
	Color    string

	Playlists []Playlist
}

type Video struct {
	gorm.Model
	ID          int `gorm:"primaryKey"`
	VideoID     string
	Title       string
	Description string

	Playlists  []Playlist  `gorm:"many2many:playlist_video;"`
	Thumbnails []Thumbnail `gorm:"polymorphic:Owner;"`

	ChannelID int

	PublishedAt time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Playlist struct {
	gorm.Model
	ID int `gorm:"primaryKey"`
	// UUID       uuid.UUID `gorm:"index:type:uuid;default:uuid_generate_v4()"`
	PlaylistID string `gorm:"index"`
	Name       string
	Color      string

	Videos        []Video `gorm:"many2many:playlist_video;"`
	Subscriptions []Subscription

	UserID int

	CreatedAt time.Time
	UpdatedAt time.Time
}

type Subscription struct {
	gorm.Model
	ID int `gorm:"primaryKey"`
	// UUID uuid.UUID `gorm:"index:type:uuid;default:uuid_generate_v4()"`

	ChannelID  int
	UserID     int
	PlaylistID int

	CreatedAt time.Time
	UpdatedAt time.Time
}

type Channel struct {
	gorm.Model
	ID          int    `gorm:"primaryKey"`
	ChannelID   string `gorm:"index"`
	Title       string
	Description string

	Videos     []Video
	Thumbnails []Thumbnail `gorm:"polymorphic:Owner;"`
}

type Thumbnail struct {
	gorm.Model
	ID     int `gorm:"primaryKey"`
	URL    string
	Width  int
	Height int

	OwnerID   int
	OwnerType ThumbnailType
}

type ThumbnailType struct {
	ID   int `gorm:"primaryKey"`
	Type string
}

/* Main */
func main() {

	SetENV("../.env")

	a := App{}

	a.InitializeRouter()
	a.InitializeDatabase(
		os.Getenv("MYSQL_ROOT_USER"),
		os.Getenv("MYSQL_ROOT_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_TCP_PORT"),
		os.Getenv("MYSQL_DATABASE"),
	)

	a.Run(
		os.Getenv("GO_API_HOST"),
		os.Getenv("GO_API_PORT"),
	)
}

/* Application Structure */
type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

func (a *App) InitializeRouter() {
	a.Router = mux.NewRouter()

	a.InitializeRoutes()
}

func (a *App) InitializeDatabase(username string, password string, ip string, port string, name string) {

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, ip, port, name)

	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if err != nil {
		panic("Error creating db connection")
	}

	db.AutoMigrate(
		&User{},
		&Playlist{},
		&Channel{},
		&Subscription{},
		&Video{},
		&Thumbnail{},
		&ThumbnailType{},
	)

	a.DB = db
}

func (a *App) HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(`{"name": "killian"}`)
}

func (a *App) Run(host string, port string) {
	log.Fatal(http.ListenAndServe(host+":"+port, a.Router))
}

func (a *App) InitializeRoutes() {
	a.Router.HandleFunc("/", a.HomeHandler)

}
