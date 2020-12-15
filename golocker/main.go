package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

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

// type Video struct {
// 	gorm.Model
// 	video_id   string
// 	channel_id string
// 	title      string
// }

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
	connection_string := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, ip, port, name)

	db, err := gorm.Open(mysql.Open(connection_string), &gorm.Config{})

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
