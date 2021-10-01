package data

import (
	"github.com/Killian264/YTLocker/golocker/models"
)

var (
	YTLOCKER_TOKEN_TYPE = "ytlocker"
	USER_TOKEN_TYPE     = "user"
)

// NewYoutubeClientConfig creates a new client config
func (d *Data) NewYoutubeClientConfig(config *models.YoutubeClientConfig) error {
	config.ID = d.rand.ID()

	result := d.db.Create(&config)

	return result.Error
}

// NewYoutubeToken creates a new token isUserToken is true if the token is for a user account
// func (d *Data) NewYoutubeToken(token *models.YoutubeToken, isUserToken bool) error {
// 	token.ID = d.rand.ID()

// 	token.Type = YTLOCKER_TOKEN_TYPE
// 	if isUserToken {
// 		token.Type = USER_TOKEN_TYPE
// 	}

// 	result := d.db.Create(&token)

// 	return result.Error
// }

// GetBaseClientConfig gets the base client config for the program
func (d *Data) GetBaseClientConfig() (models.YoutubeClientConfig, error) {
	config := models.YoutubeClientConfig{}

	result := d.db.First(&config)

	return config, result.Error
}

// GetBaseToken gets the base token for the program
func (d *Data) GetBaseToken() (models.YoutubeToken, error) {
	token := models.YoutubeToken{}

	result := d.db.Where("type = ?", YTLOCKER_TOKEN_TYPE).First(&token)

	return token, result.Error
}
