package oauthmanager

import (
	"github.com/Killian264/YTLocker/golocker/data"
	"github.com/Killian264/YTLocker/golocker/helpers/parsers"
	"github.com/Killian264/YTLocker/golocker/models"
	"github.com/Killian264/YTLocker/golocker/services/ytservice"
	"golang.org/x/oauth2"
	"google.golang.org/api/youtube/v3"
)

type IOauthManagerData interface {
	NewYoutubeAccount(account models.YoutubeAccount) (models.YoutubeAccount, error)
	NewUserYoutubeAccount(userID uint64, accountID uint64) error

	GetAccount(accountId uint64) (models.YoutubeAccount, error)
	GetAccountByEmail(email string) (models.YoutubeAccount, error)
	GetAccountFromToken(token models.YoutubeToken) (models.YoutubeAccount, error)
	GetUserYoutubeAccounts(user models.User) ([]uint64, error)
	UpdateAccount(account models.YoutubeAccount) (models.YoutubeAccount, error)
}

type IYoutubeService interface {
	Initialize(config oauth2.Config, token oauth2.Token) error
	GetUser() (models.OAuthUserInfo, error)
	GetChannel() (*youtube.Channel, error)
}

// OauthManager manages oauth information
type OauthManager struct {
	data    IOauthManagerData
	youtube IYoutubeService
	config  oauth2.Config
	token   oauth2.Token
	account models.YoutubeAccount
}

// NewOauthManager creates an oauth manager
// secretsDir - the file directory to read the secrets from
// redirectUrl - the config redirect url, should be the callback url
func NewOauthManager(dataService IOauthManagerData, youtube IYoutubeService, secretsDir string, redirectUrl string) *OauthManager {
	manager := &OauthManager{
		data:    dataService,
		youtube: youtube,
	}

	config, token := manager.readInOauthConfigData(secretsDir)

	manager.initializeBaseData(config, token, redirectUrl)

	return manager
}

// NewFakeOauthManager creates a fake oauth manager
// NOTE: These secrets are fake
func NewFakeOauthManager(data IOauthManagerData) *OauthManager {
	config := models.YoutubeClientConfig{
		ClientID:     "11223534584-asdfhasdjfhwieyrwqejhkflasd.apps.googleusercontent.com",
		ClientSecret: "qwerHSwer_asdhwuerJHFDJqkqw",
		RedirectURL:  "urn:ietf:wg:oauth:2.0:oob",
		AuthURL:      "https://accounts.google.com/o/oauth2/auth",
		TokenURL:     "https://oauth2.googleapis.com/token",
	}
	token := models.YoutubeToken{
		AccessToken:  "sa23.345234524623sdfasdfq-qegehgower9505034jfeworrjwertw_qqwerjfldssgert345sdgdgew-bheiyqeotleqjrljdfluao23423_QwekjfuI023kjasdfwer",
		TokenType:    "Bearer",
		RefreshToken: "asdfjqwekj23//2342329asqq-ajfdki22399jjIjiJIWJFfwerw_qwefdasferw_zwaehwejlkWW",
		Expiry:       "2021-04-13T23:30:06.1139442-05:00",
	}

	manager := &OauthManager{
		data:    data,
		youtube: &ytservice.YTPlaylistFake{},
	}

	manager.initializeBaseData(config, token, "https://ytlocker.com/")

	return manager
}

func (m *OauthManager) initializeBaseData(config models.YoutubeClientConfig, token models.YoutubeToken, redirectUrl string) {
	m.config = parsers.ParseYoutubeClient(config)
	m.token = parsers.ParseYoutubeToken(token)

	account, err := m.data.GetAccountFromToken(token)
	if err != nil && err == data.ErrorNotFound {
		account, err = m.CreateAccount(m.token, "manage")
		if err != nil {
			panic("failed to create token account err: " + err.Error())
		}
	}

	m.account = account

	m.config.RedirectURL = redirectUrl
}
