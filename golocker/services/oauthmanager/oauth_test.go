package oauthmanager

import (
	"testing"
	"time"

	"github.com/Killian264/YTLocker/golocker/data"
	"github.com/Killian264/YTLocker/golocker/helpers/parsers"
	"github.com/Killian264/YTLocker/golocker/models"
	"github.com/Killian264/YTLocker/golocker/services/user"
	"github.com/Killian264/YTLocker/golocker/services/ytservice"
	"github.com/stretchr/testify/assert"
	"golang.org/x/oauth2"
)

var validUser = models.User{}
var validToken = oauth2.Token{}

func Test_Get_Base_Methods(t *testing.T) {
	service, service2 := createMockServices2()

	config := service.GetBaseConfig()
	account, err := service.GetBaseYoutubeAccount()
	assert.Nil(t, err)

	assert.NotEmpty(t, config)
	assert.NotEmpty(t, account)
	assert.NotEmpty(t, account.YoutubeToken)
	assert.NotEmpty(t, account.YoutubeToken.ID)

	// There should only be one base account
	config2 := service2.GetBaseConfig()
	account2, err := service2.GetBaseYoutubeAccount()
	assert.Nil(t, err)

	assert.Equal(t, account.ID, account2.ID)
	assert.Equal(t, config.ClientID, config2.ClientID)
	assert.NotEmpty(t, account2.YoutubeToken.ID)
}

func Test_Create_Account(t *testing.T) {
	service := createMockServices()

	account, err := service.GetLoginAccount(validToken, "view")
	assert.Nil(t, err)
	assert.NotEmpty(t, account.Email)
	assert.NotEmpty(t, account.Username)

	account2, err := service.GetAccountByEmail(account.Email)
	assert.Nil(t, err)

	assert.Equal(t, account.Email, account2.Email)

	assert.NotEmpty(t, account.YoutubeToken.AccessToken)
	assert.Equal(t, account.YoutubeToken.AccessToken, account2.YoutubeToken.AccessToken)

	account3, err := service.GetAccountById(account.ID)
	assert.Nil(t, err)

	assert.Equal(t, account.Email, account3.Email)

	assert.NotEmpty(t, account3.YoutubeToken.AccessToken)
	assert.NotEmpty(t, account3.YoutubeToken.ID)
	assert.Equal(t, account.YoutubeToken.AccessToken, account3.YoutubeToken.AccessToken)
}

func Test_Create_Account_Expired_Token(t *testing.T) {
	service := createMockServices()

	validToken.Expiry = time.Now().AddDate(-1, 1, 1)

	account, err := service.GetLoginAccount(validToken, "view")
	assert.Nil(t, err)
	assert.NotEmpty(t, account.Email)
	assert.NotEmpty(t, account.Username)

	account2, err := service.GetAccountByEmail(account.Email)
	assert.Nil(t, err)

	assert.Equal(t, account.Email, account2.Email)

	assert.NotEmpty(t, account.YoutubeToken.AccessToken)
	assert.Equal(t, account.YoutubeToken.AccessToken, account2.YoutubeToken.AccessToken)
}

func Test_Account_Cannot_Be_Created_Twice(t *testing.T) {
	service := createMockServices()

	account, err := service.GetLoginAccount(validToken, "view")
	assert.Nil(t, err)
	assert.NotEmpty(t, account.Email)

	account2, err := service.GetLoginAccount(validToken, "view")
	assert.Equal(t, account.Email, account2.Email)
}

func Test_Link_Account(t *testing.T) {
	service := createMockServices()

	account, _ := service.GetLoginAccount(validToken, "view")

	err := service.LinkAccount(validUser, account)
	assert.Nil(t, err)
}

func Test_Get_User_YoutubeAccounts(t *testing.T) {
	service := createMockServices()

	accounts, err := service.GetUserAccountList(validUser)
	assert.Nil(t, err)
	assert.Equal(t, 0, len(accounts))

	account, err := service.GetLoginAccount(validToken, "view")
	assert.Nil(t, err)

	err = service.LinkAccount(validUser, account)
	assert.Nil(t, err)

	accounts, err = service.GetUserAccountList(validUser)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(accounts))
}

func Test_Get_User_Account(t *testing.T) {
	service := createMockServices()

	account, err := service.GetLoginAccount(validToken, "view")
	assert.Nil(t, err)

	err = service.LinkAccount(validUser, account)
	assert.Nil(t, err)

	foundAccount, err := service.GetUserAccount(models.User{ID: validUser.ID}, account.ID)
	assert.Nil(t, err)
	assert.Equal(t, account.ID, foundAccount.ID)
}

func Test_Refresh_Token(t *testing.T) {
	service := createMockServices()

	account, err := service.GetLoginAccount(validToken, PERMISSION_LEVEL_VIEW)
	assert.Nil(t, err)

	oauth2Token := parsers.ParseYoutubeToken(account.YoutubeToken)
	oauth2Token.AccessToken = "aslkdjfasdf"

	account2, err := service.RefreshToken(account, oauth2Token, PERMISSION_LEVEL_VIEW)
	assert.Nil(t, err)

	assert.Equal(t, account.Email, account2.Email)
	assert.NotEqual(t, account.YoutubeToken.AccessToken, account2.YoutubeToken.AccessToken)
}

func Test_Initialize_Youtube_Service(t *testing.T) {
	service := createMockServices()
	youtube := ytservice.NewYTPlaylistFake()

	account, err := service.GetLoginAccount(validToken, PERMISSION_LEVEL_VIEW)
	assert.Nil(t, err)

	account.YoutubeToken.Expiry = time.Now().AddDate(-1, 0, 0).String()

	returnedService, err := service.InitializeYTService(youtube, account.ID)
	assert.Nil(t, err)

	returnedService.GetUser()

	account2, err := service.GetAccountByEmail(account.Email)

	assert.Equal(t, account.Email, account2.Email)
	assert.NotEqual(t, account.YoutubeToken.AccessToken, account2.YoutubeToken.AccessToken)
}

func createMockServices() *OauthManager {
	db := data.InMemoryMySQLConnect()

	validUser = models.User{
		Username: "killian",
		Email:    "killiandebacker@gmail.com",
		Picture:  "https://lh3.googleusercontent.com/a/default-user=s96-c",
	}

	validToken = oauth2.Token{
		AccessToken:  "23593045903245",
		TokenType:    "Bearer",
		Expiry:       time.Now().AddDate(0, 0, 1),
		RefreshToken: "q30572309458320945",
	}

	userService := user.NewUser(db)

	bearer, err := userService.GenerateTemporarySessionBearer()
	if err != nil {
		panic("err should not happen")
	}

	createdUser, err := userService.Login(validUser, bearer)
	if err != nil {
		panic("err should not happen")
	}
	validUser = createdUser

	service := NewFakeOauthManager(
		db,
	)

	return service
}

// yeah im lazy
func createMockServices2() (*OauthManager, *OauthManager) {
	db := data.InMemorySQLiteConnect()

	userService := user.NewUser(db)

	bearer, err := userService.GenerateTemporarySessionBearer()
	if err != nil {
		panic("err should not happen")
	}

	createdUser, err := userService.Login(validUser, bearer)
	if err != nil {
		panic("err should not happen")
	}

	validUser = createdUser

	service := NewFakeOauthManager(
		db,
	)

	service2 := NewFakeOauthManager(
		db,
	)

	return service, service2
}
