package oauthmanager

import (
	"testing"

	"github.com/Killian264/YTLocker/golocker/data"
	"github.com/Killian264/YTLocker/golocker/models"
	"github.com/Killian264/YTLocker/golocker/services/user"
	"github.com/stretchr/testify/assert"
	"golang.org/x/oauth2"
)

var validUser = models.User{
	Username: "killian",
	Email:    "killiandebacker@gmail.com",
	Picture:  "https://lh3.googleusercontent.com/a/default-user=s96-c",
}

func Test_Get_Base_Methods(t *testing.T) {
	service, service2 := createMockServices2()

	config := service.GetBaseConfig()
	token := service.GetBaseToken()
	account := service.GetBaseYoutubeAccount()

	assert.NotEmpty(t, config)
	assert.NotEmpty(t, token)
	assert.NotEmpty(t, account)

	// There should only be one base account
	config2 := service2.GetBaseConfig()
	token2 := service2.GetBaseToken()
	account2 := service2.GetBaseYoutubeAccount()

	assert.Equal(t, account.ID, account2.ID)
	assert.Equal(t, config, config2)
	assert.Equal(t, token, token2)
}

func Test_Create_Account(t *testing.T) {
	service := createMockServices()

	account, err := service.CreateAccount(oauth2.Token{}, "view")

	assert.Nil(t, err)
	assert.NotEmpty(t, account)
}

func Test_Account_Cannot_Be_Created_Twice(t *testing.T) {
	service := createMockServices()

	service.CreateAccount(oauth2.Token{}, "view")

	_, err := service.CreateAccount(oauth2.Token{}, "view")
	assert.NotNil(t, err)
}

func Test_Link_Account(t *testing.T) {
	service := createMockServices()

	account, _ := service.CreateAccount(oauth2.Token{}, "view")

	err := service.LinkAccount(validUser, account)
	assert.Nil(t, err)
}

func Test_Get_User_YoutubeAccounts(t *testing.T) {
	service := createMockServices()

	account, _ := service.CreateAccount(oauth2.Token{}, "view")

	service.LinkAccount(validUser, account)

	accounts, err := service.GetUserAccountList(validUser)
	assert.Nil(t, err)
	assert.NotEmpty(t, accounts)
	assert.Equal(t, len(accounts), 1)
}

func Test_Get_Account_From_Token(t *testing.T) {
	service := createMockServices()

	account, _ := service.CreateAccount(oauth2.Token{}, "view")

	found, err := service.GetAccountByAccessToken(account.YoutubeToken)

	assert.Nil(t, err)
	assert.Equal(t, found.ID, account.ID)
}

func Test_Get_Account_From_Token_Fails(t *testing.T) {
	service := createMockServices()

	_, err := service.GetAccountByAccessToken(models.YoutubeToken{
		AccessToken: "bananas",
	})

	assert.NotNil(t, err)
	assert.Equal(t, err, data.ErrorNotFound)
}

func Test_Link_Base_Accounts(t *testing.T) {
	service := createMockServices()

	err := service.LinkBaseAccounts(validUser, oauth2.Token{})
	assert.Nil(t, err)

	accountList, _ := service.GetUserAccountList(validUser)

	assert.Equal(t, len(accountList), 2)

	err = service.LinkBaseAccounts(validUser, oauth2.Token{})
	assert.Nil(t, err)

	assert.Equal(t, len(accountList), 2)
}

func createMockServices() *OauthManager {
	db := data.InMemorySQLiteConnect()

	userService := user.NewUser(db)

	validUser, _ = userService.Login(validUser)

	service := NewFakeOauthManager(
		db,
	)

	return service
}

// yeah im lazy
func createMockServices2() (*OauthManager, *OauthManager) {
	db := data.InMemorySQLiteConnect()

	userService := user.NewUser(db)

	validUser, _ = userService.Login(validUser)

	service := NewFakeOauthManager(
		db,
	)

	service2 := NewFakeOauthManager(
		db,
	)

	return service, service2
}
