package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Killian264/YTLocker/golocker/data"
	"github.com/Killian264/YTLocker/golocker/handlers"
	"github.com/Killian264/YTLocker/golocker/services"
	"github.com/Killian264/YTLocker/golocker/services/cronjobs"
	"github.com/Killian264/YTLocker/golocker/services/oauthmanager"
	"github.com/Killian264/YTLocker/golocker/services/playlist"
	"github.com/Killian264/YTLocker/golocker/services/subscribe"
	"github.com/Killian264/YTLocker/golocker/services/user"
	"github.com/Killian264/YTLocker/golocker/services/ytmanager"
	"github.com/Killian264/YTLocker/golocker/services/ytservice"

	muxhandler "github.com/gorilla/handlers"

	"github.com/gorilla/mux"
)

func main() {
	logger := log.New(os.Stdout, "Main: ", log.Ldate|log.Ltime)

	logger.Println("Running...")

	s := NewServices(logger)

	logger.Println("Services Created...")

	StartServer(
		s,
		os.Getenv("GO_API_HOST"),
		os.Getenv("GO_API_PORT"),
	)
}

func NewServices(logger *log.Logger) *services.Services {
	youtubeHelperService := ytservice.NewYoutubeService(
		os.Getenv("YOUTUBE_API_KEY"),
	)

	playlistHelperService := &ytservice.YTPlaylist{}

	routerService := InitializeRouter()

	dataService := InitializeDatabase(
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_TCP_PORT"),
		os.Getenv("MYSQL_DATABASE"),
	)

	oauthManagerService := oauthmanager.NewOauthManager(
		dataService,
		playlistHelperService,
		"secrets/",
		"http://localhost:8080/"+"user/oauth/callback",
	)

	youtubeManagerService := ytmanager.NewYoutubeManager(
		dataService,
		youtubeHelperService,
	)

	subscriptionService := InitializeSubscribeService(
		dataService,
		youtubeManagerService,
		os.Getenv("GO_API_URL"),
	)

	playlistService := playlist.NewPlaylist(
		playlistHelperService,
		oauthManagerService,
		dataService,
	)

	userService := user.NewUser(
		dataService,
	)

	cronjobContainer := cronjobs.NewCronJobManager(
		youtubeManagerService,
		playlistService,
		subscriptionService,
		dataService,
	)

	config := services.Config{
		WebBaseUrl:     os.Getenv("WEB_URL"),
		WebLoginUrl:    os.Getenv("WEB_URL") + "/login",
		WebRedirectUrl: os.Getenv("WEB_URL") + "/redirect",
	}

	serviceContainer := &services.Services{
		Logger:       logger,
		Router:       routerService,
		Youtube:      youtubeManagerService,
		User:         userService,
		Subscribe:    subscriptionService,
		Playlist:     playlistService,
		OauthManager: oauthManagerService,
		Cronjob:      cronjobContainer,
		Config:       config,
	}

	InitializeRoutes(serviceContainer, cronjobContainer, os.Getenv("ADMIN_BEARER"))

	return serviceContainer
}

// InitializeDatabase creates DB Connection for app
func InitializeDatabase(username string, password string, ip string, port string, name string) *data.Data {
	logger := log.New(os.Stdout, "Data: ", log.Ldate|log.Ltime)

	return data.MySQLConnect(username, password, ip, port, name, logger)
}

func InitializeSubscribeService(data subscribe.ISubscriptionData, yt subscribe.IYoutubeManager, appURL string) *subscribe.Subscriber {
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
func InitializeRoutes(serviceContainer *services.Services, cronjobContainer *cronjobs.CronJobManager, adminBearer string) {
	logger := log.New(os.Stdout, "Hand: ", log.Ldate|log.Ltime)

	Injector := handlers.CreateServiceInjector(serviceContainer)
	Errors := handlers.CreateResponseWriter(logger)
	UserAuth := handlers.CreateUserAuthenticator(serviceContainer)
	PlaylistAuth := handlers.CreatePlaylistAuthenticator(serviceContainer)
	// AccountAuth := handlers.CreateAccountAuthenticator(serviceContainer)

	router := serviceContainer.Router

	router.HandleFunc("/user/information", Errors(Injector(UserAuth(handlers.UserInformation))))
	router.HandleFunc("/user/oauth/login", Errors(Injector(handlers.OAuthAuthenticate)))
	router.HandleFunc("/user/oauth/callback", Errors(Injector(handlers.OAuthAuthenticateCallback)))
	router.HandleFunc("/user/session/refresh", Errors(Injector(UserAuth(handlers.UserSessionRefresh))))

	router.HandleFunc("/playlist/create", Errors(Injector(UserAuth(handlers.PlaylistCreate))))
	router.HandleFunc("/playlist/list", Errors(Injector(UserAuth(handlers.PlaylistList))))
	router.HandleFunc("/playlist/videos/latest", Errors(Injector(UserAuth(handlers.PlaylistLatestVideos))))
	router.HandleFunc("/playlist/{playlist_id}/update", Errors(Injector(UserAuth(PlaylistAuth(handlers.PlaylistUpdate)))))
	router.HandleFunc("/playlist/{playlist_id}/subscribe/{channel_id}", Errors(Injector(UserAuth(PlaylistAuth(handlers.PlaylistAddSubscription)))))
	router.HandleFunc("/playlist/{playlist_id}/unsubscribe/{channel_id}", Errors(Injector(UserAuth(PlaylistAuth(handlers.PlaylistRemoveSubscription)))))
	router.HandleFunc("/playlist/{playlist_id}/delete", Errors(Injector(UserAuth(PlaylistAuth(handlers.PlaylistDelete)))))

	router.HandleFunc("/video/{video_id}", Errors(Injector(UserAuth(handlers.GetVideo))))
	router.HandleFunc("/channel/search", Errors(Injector(UserAuth(handlers.SearchChannel))))
	router.HandleFunc("/channel/get/{channel_id}", Errors(Injector(UserAuth(handlers.GetChannel))))

	router.HandleFunc("/account/list", Errors(Injector(UserAuth(handlers.AccountList))))
	// router.HandleFunc("/account/{account_id}", Errors(Injector(UserAuth(AccountAuth(handlers.AccountGet)))))

	SubscribeErrors := handlers.CreateSubscribeHandler(logger)
	AdminAuth := handlers.CreateAdminAuthenticator(serviceContainer, adminBearer)

	router.HandleFunc("/subscribe/{secret}/", SubscribeErrors(Injector(handlers.HandleYoutubePush)))

	router.HandleFunc("/admin/check/uploads", AdminAuth(func(rw http.ResponseWriter, r *http.Request) {
		cronjobContainer.CheckForMissedUploads()
	}))

	router.HandleFunc("/admin/update/subscriptions", AdminAuth(func(rw http.ResponseWriter, r *http.Request) {
		cronjobContainer.ResubscribeAllSubscriptions()
	}))
}

func StartServer(serviceContainer *services.Services, host string, port string) {
	headersOk := muxhandler.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	originsOk := muxhandler.AllowedOrigins([]string{"*"})
	methodsOk := muxhandler.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS", "DELETE"})

	router := muxhandler.CORS(originsOk, headersOk, methodsOk)(serviceContainer.Router)

	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", host, port), router))
}
