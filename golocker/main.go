package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/Killian264/YTLocker/models"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	logFolder  = "../logs"
	apiLogFile = "apilogs.txt"
	dbLogFile  = "dblogs.txt"
)

/* Helper Functions */
func SetENV(location string) {
	err := godotenv.Load(location)

	if err != nil {
		panic("Error setting ENV.")
	}
}

/* Main */
func main() {

	SetENV("../.env")

	a := App{}

	a.InitializeRouter()
	a.InitializeDatabase(
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_TCP_PORT"),
		os.Getenv("MYSQL_DATABASE"),
		logFolder+dbLogFile,
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
	Logger *log.Logger
}

func (a *App) InitializeRouter() {
	a.Router = mux.NewRouter()

	a.InitializeRoutes()
}

func (a *App) InitializeDatabase(username string, password string, ip string, port string, name string, logFileLoc string) {

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, ip, port, name)

	// logFile, err := os.OpenFile(
	// 	logFileLoc,
	// 	os.O_APPEND|os.O_CREATE|os.O_WRONLY,
	// 	0644,
	// )

	// if err != nil {
	// 	panic("Error opening or creating database log file.")
	// }

	// logger := logger.New(
	// 	log.New(logFile, "\r\n", log.LstdFlags),
	// 	logger.Config{
	// 		Colorful: true,
	// 		LogLevel: logger.Warn,
	// 	},
	// )

	db, err := gorm.Open(
		mysql.Open(connectionString),
		&gorm.Config{
			// Logger: logger,
		},
	)

	if err != nil {
		panic("Error creating db connection")
	}

	db.AutoMigrate(
		&models.User{},
		&models.Playlist{},
		&models.Channel{},
		&models.Subscription{},
		&models.Video{},
		&models.Thumbnail{},
		&models.ThumbnailType{},
		&models.Request{},
	)

	a.DB = db
}

func (a *App) Run(host string, port string) {
	log.Fatal(http.ListenAndServe(":"+port, a.Router))
}

func (a *App) InitializeRoutes() {
	a.Router.HandleFunc("/", a.HomeHandler)
}

func (a *App) HomeHandler(w http.ResponseWriter, r *http.Request) {

	hubChallenge := r.URL.Query().Get("hub.challenge")

	hubTopic := r.URL.Query().Get("hub.topic")

	bytes, err := ioutil.ReadAll(r.Body)

	if err != nil {
		panic("error reading body")
	}

	body := string(bytes)

	request := models.Request{
		Body:         body,
		HubChallenge: hubChallenge,
		HubTopic:     hubTopic,
	}

	a.DB.Create(&request)

	fmt.Fprintf(w, hubChallenge)

}
