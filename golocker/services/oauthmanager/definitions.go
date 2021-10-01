package oauthmanager

import (
	"github.com/Killian264/YTLocker/golocker/helpers/parsers"
	"github.com/Killian264/YTLocker/golocker/models"
	"golang.org/x/oauth2"
)

type IOauthManagerData interface {
	NewYoutubeClientConfig(config *models.YoutubeClientConfig) error
	// NewYoutubeToken(token *models.YoutubeToken, isUserToken bool) error

	GetBaseClientConfig() (models.YoutubeClientConfig, error)
	GetBaseToken() (models.YoutubeToken, error)
}

// OauthManager manages oauth information
type OauthManager struct {
	data   IOauthManagerData
	config oauth2.Config
	token  oauth2.Token
}

// NewOauthManager creates an oauth manager
func NewOauthManager(data IOauthManagerData, secretsDir string, redirectUrl string) *OauthManager {
	manager := &OauthManager{
		data: data,
	}

	config, token := manager.readInOauthConfigData(secretsDir)

	manager.config = parsers.ParseYoutubeClient(config)
	manager.token = parsers.ParseYoutubeToken(token)
	manager.config.RedirectURL = redirectUrl

	return manager
}

// NewFakeOauthManager creates a fake oauth manager
// NOTE: These secrets are fake
func NewFakeOauthManager(data IOauthManagerData) *OauthManager {
	configData := models.YoutubeClientConfig{
		ClientID:     "11223534584-asdfhasdjfhwieyrwqejhkflasd.apps.googleusercontent.com",
		ClientSecret: "qwerHSwer_asdhwuerJHFDJqkqw",
		RedirectURL:  "urn:ietf:wg:oauth:2.0:oob",
		AuthURL:      "https://accounts.google.com/o/oauth2/auth",
		TokenURL:     "https://oauth2.googleapis.com/token",
	}
	tokenData := models.YoutubeToken{
		AccessToken:  "sa23.345234524623sdfasdfq-qegehgower9505034jfeworrjwertw_qqwerjfldssgert345sdgdgew-bheiyqeotleqjrljdfluao23423_QwekjfuI023kjasdfwer",
		TokenType:    "Bearer",
		RefreshToken: "asdfjqwekj23//2342329asqq-ajfdki22399jjIjiJIWJFfwerw_qwefdasferw_zwaehwejlkWW",
		Expiry:       "2021-04-13T23:30:06.1139442-05:00",
	}

	manager := &OauthManager{
		data: data,
	}

	manager.config = parsers.ParseYoutubeClient(configData)
	manager.token = parsers.ParseYoutubeToken(tokenData)
	manager.config.RedirectURL = "https://ytlocker.com/"

	return manager
}
