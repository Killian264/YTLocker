package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"fmt"
	
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	"gorm.io/driver/mysql"
  	"gorm.io/gorm"
)

/* Helper Functions */
func SetENV(location string){
	err := godotenv.Load(location)

	if err != nil{
		panic("Error setting ENV.")
	}
}


/* Main */
func main() {

	SetENV(".env")

	a := App{}

	a.InitializeRouter()
	a.InitializeDatabase(
		os.Getenv("db_username"),
		os.Getenv("db_password"),
		os.Getenv("db_ip"),
		os.Getenv("db_port"), 
		os.Getenv("db_name"),
	)

	a.Run(
		os.Getenv("go_api_host"),
		os.Getenv("go_api_port"),
	)
}



/* Application Structure */
type App struct {
	Router *mux.Router
	DB *gorm.DB
}

func (a *App) InitializeRouter(){
	a.Router = mux.NewRouter()

	a.InitializeRoutes();
}

func (a *App) InitializeDatabase(username string, password string, ip string, port string, name string){
	connection_string := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, ip, port, name)

	db, err := gorm.Open(mysql.Open(connection_string), &gorm.Config{})

	if err != nil{
		panic("Error creating db connection")
	}

	a.DB = db
}

func (a *App) HomeHandler(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(`{"name": "killian"}`)
}

func (a *App) Run(host string, port string) {
    log.Fatal(http.ListenAndServe(host + ":" + port, a.Router))
}

func (a *App) InitializeRoutes(){
	a.Router.HandleFunc("/", a.HomeHandler)
	
}