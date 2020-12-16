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

	"github.com/google/uuid"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	UserName  string
	Email     string
	Password  string
	Salt      string
	Playlists []Playlist
	Color     string
}

type Video struct {
	gorm.Model
	ID          string
	VideoID     string
	Channel     Channel
	Title       string
	Description string
	Playlists   []Playlist `gorm:"many2many:playlist_video;"`
	Thumbnails  []Thumbnail
	PublishedAt time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Playlist struct {
	gorm.Model
	ID            uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	PlaylistID    string    `gorm:"index"`
	Name          string
	Color         string
	Videos        []Video `gorm:"many2many:playlist_video;"`
	Subscriptions []Subscription
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type Subscription struct {
	gorm.Model
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Channel   Channel
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Channel struct {
	gorm.Model
	ID          string
	ChannelID   string
	Title       string
	Description string
	Thumbnails  []Thumbnail
}

type Thumbnail struct {
	gorm.Model
	ID     string
	URL    string
	Width  int
	Height int
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

	logger := logger.New()

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, ip, port, name)

	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if err != nil {
		panic("Error creating db connection")
	}

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
