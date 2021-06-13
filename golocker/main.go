package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/Killian264/YTLocker/golocker/data"
	"github.com/Killian264/YTLocker/golocker/handlers"
	"github.com/Killian264/YTLocker/golocker/helpers/parsers"
	"github.com/Killian264/YTLocker/golocker/models"
	"github.com/Killian264/YTLocker/golocker/services"
	"github.com/Killian264/YTLocker/golocker/services/cronjobs"
	playlistserivce "github.com/Killian264/YTLocker/golocker/services/playlist"
	"github.com/Killian264/YTLocker/golocker/services/subscribe"
	"github.com/Killian264/YTLocker/golocker/services/user"
	"github.com/Killian264/YTLocker/golocker/services/ytmanager"
	"github.com/Killian264/YTLocker/golocker/services/ytservice"
	"github.com/robfig/cron"

	muxhandler "github.com/gorilla/handlers"

	"github.com/gorilla/mux"
)

func main() {

	logger := log.New(os.Stdout, "Main: ", log.Ldate|log.Ltime)

	logger.Println("Running... ----------------------------")

	s := NewServices(logger)

	Run(
		s,
		os.Getenv("GO_API_HOST"),
		os.Getenv("GO_API_PORT"),
	)

}

func NewServices(logger *log.Logger) *services.Services {

	s := &services.Services{
		Logger: logger,
	}

	youtubeHelper := InitializeYTService(
		os.Getenv("YOUTUBE_API_KEY"),
	)

	playlistHelper := InitalizePlaylistHelper()

	s.Router = InitializeRouter()

	s.Data = InitializeDatabase(
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_TCP_PORT"),
		os.Getenv("MYSQL_DATABASE"),
	)

	ReadInSecrets(
		s.Data,
		"secrets/",
	)

	s.Youtube = InitalizeYoutubeManager(
		s.Data,
		youtubeHelper,
	)

	s.Subscribe = InitalizeSubscribeService(
		s.Data,
		s.Youtube,
		os.Getenv("GO_API_URL"),
	)

	s.Playlist = InitalizePlaylistManager(
		s.Data,
		playlistHelper,
	)

	s.User = InitalizeUserService(
		s.Data,
	)

	InitializeRoutes(s, s.Router, os.Getenv("ADMIN_BEARER"))

	InitializeCronJobs(s)

	return s
}

func InitalizePlaylistHelper() *ytservice.YTPlaylist {

	return &ytservice.YTPlaylist{}

}

func InitalizeUserService(data user.IUserData) *user.User {
	service := user.NewUser(
		data,
	)

	return service
}

func InitalizePlaylistManager(data playlistserivce.IPlaylistManagerData, yt playlistserivce.IYTPlaylist) *playlistserivce.PlaylistManager {
	service := playlistserivce.NewPlaylist(
		yt,
		data,
	)

	return service
}

func InitalizeYoutubeManager(data ytmanager.IYoutubeManagerData, yt ytmanager.IYTService) *ytmanager.YoutubeManager {
	service := ytmanager.NewYoutubeManager(
		data,
		yt,
	)

	return service
}

func InitalizeSubscribeService(data subscribe.ISubscriptionData, yt subscribe.IYoutubeManager, appURL string) *subscribe.Subscriber {

	service := subscribe.NewSubscriber(data, yt)

	service.SetYTPubSubUrl("https://pubsubhubbub.appspot.com/subscribe")
	service.SetSubscribeUrl(appURL, "/subscribe/{secret}")

	return service

}

// InitializeRouter Creates Router for app
func InitializeRouter() *mux.Router {

	router := mux.NewRouter()

	router.Use(muxhandler.RecoveryHandler())

	return router

}

// InitializeRoutes creates the routes
func InitializeRoutes(services *services.Services, router *mux.Router, adminBearer string) {

	logger := log.New(os.Stdout, "Hand: ", log.Ldate|log.Ltime)

	Injector := handlers.CreateServiceInjector(services)
	Errors := handlers.CreateResponseWriter(logger)
	UserAuth := handlers.CreateUserAuthenticator(services)
	PlaylistAuth := handlers.CreatePlaylistAuthenticator(services)
	SubscribeErrors := handlers.CreateSubscribeHandler(logger)
	AdminAuth := handlers.CreateAdminAuthenticator(services, adminBearer)

	router.HandleFunc("/subscribe/{secret}/", SubscribeErrors(Injector(handlers.HandleYoutubePush)))

	router.HandleFunc("/user/login", Errors(Injector(handlers.UserLogin)))
	router.HandleFunc("/user/register", Errors(Injector(handlers.UserRegister)))
	router.HandleFunc("/user/information", Errors(Injector(UserAuth(handlers.UserInformation))))

	router.HandleFunc("/playlist/create", Errors(Injector(UserAuth(handlers.PlaylistCreate))))
	router.HandleFunc("/playlist/list", Errors(Injector(UserAuth(handlers.PlaylistList))))
	router.HandleFunc("/playlist/videos/latest", Errors(Injector(UserAuth(handlers.PlaylistLatestVideos))))
	router.HandleFunc("/playlist/{playlist_id}/subscribe/{channel_id}", Errors(Injector(UserAuth(PlaylistAuth(handlers.PlaylistAddSubscription)))))
	router.HandleFunc("/playlist/{playlist_id}/unsubscribe/{channel_id}", Errors(Injector(UserAuth(PlaylistAuth(handlers.PlaylistRemoveSubscription)))))

	router.HandleFunc("/video/{video_id}", Errors(Injector(UserAuth(handlers.GetVideo))))
	router.HandleFunc("/channel/{channel_id}", Errors(Injector(UserAuth(handlers.GetChannel))))

	router.HandleFunc("/admin/check/uploads", AdminAuth(func(rw http.ResponseWriter, r *http.Request) {
		checkForMissedUploads(services, logger)
	}))

	router.HandleFunc("/admin/update/subscriptions", AdminAuth(func(rw http.ResponseWriter, r *http.Request) {
		resubscribeAllSubscriptions(services, logger)
	}))

}

// InitializeYTService Creates YTService for app
func InitializeYTService(apiKey string) *ytservice.YTService {

	return ytservice.NewYoutubeService(apiKey)

}

func ReadInSecrets(data *data.Data, secretsPath string) {

	clientData, err := readInClientSecret(fmt.Sprintf("%s%s", secretsPath, "client_secret.json"))
	if err != nil {
		panic(err)
	}

	tokenData, err := readInAccessToken(fmt.Sprintf("%s%s", secretsPath, "access_secret.json"))
	if err != nil {
		panic(err)
	}

	err = data.NewYoutubeClientConfig(&clientData)
	if err != nil && !strings.Contains(err.Error(), "Duplicate entry") {
		panic(err)
	}

	err = data.NewYoutubeToken(&tokenData)
	if err != nil && !strings.Contains(err.Error(), "Duplicate entry") {
		panic(err)
	}

}

func readInClientSecret(path string) (models.YoutubeClientConfig, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return models.YoutubeClientConfig{}, err
	}

	return parsers.ParseClientJson(string(b))
}

func readInAccessToken(path string) (models.YoutubeToken, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return models.YoutubeToken{}, err
	}

	return parsers.ParseAccessTokenJson(string(b))
}

// InitializeDatabase creates DB Connection for app
func InitializeDatabase(username string, password string, ip string, port string, name string) *data.Data {

	logger := log.New(os.Stdout, "Data: ", log.Ldate|log.Ltime)

	return data.MySQLConnect(username, password, ip, port, name, logger)

}

func Run(s *services.Services, host string, port string) {

	headersOk := muxhandler.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	originsOk := muxhandler.AllowedOrigins([]string{"*"})
	methodsOk := muxhandler.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	router := muxhandler.CORS(originsOk, headersOk, methodsOk)(s.Router)

	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", host, port), router))
}

func InitializeCronJobs(service *services.Services) {

	c := cron.New()

	logger := log.New(os.Stdout, "Cron: ", log.Ldate|log.Ltime)

	job := cronjobs.NewInsertVideosJob(service, logger)

	c.AddFunc("@every 2m", func() {
		insertVideos(job, logger)
	})

	c.AddFunc("@weekly", func() {
		resubscribeAllSubscriptions(service, logger)
	})

	c.AddFunc("@every 6h", func() {
		checkForMissedUploads(service, logger)
	})

	c.Start()

}

func insertVideos(job *cronjobs.InsertVideosJob, logger *log.Logger) {
	logger.Print("Starting Insert Videos: --------------")

	job.Run()

	logger.Print("Finished Insert Videos: --------------")
}

func checkForMissedUploads(service *services.Services, logger *log.Logger) {
	logger.Print("Starting Missed Uploads: -------------")

	err := service.Youtube.CheckForMissedUploads(logger)
	if err != nil {
		logger.Print(err)
	}

	logger.Print("Finished Missed Uploads: -------------")
}

func resubscribeAllSubscriptions(service *services.Services, logger *log.Logger) {
	logger.Print("Starting Resubscribe: ----------------")

	err := service.Subscribe.ResubscribeAll()
	if err != nil {
		logger.Print(err)
	}

	logger.Print("Finished Resubscribe: ----------------")
}
