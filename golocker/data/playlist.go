package data

import (
	"github.com/Killian264/YTLocker/golocker/models"
	uuid "github.com/satori/go.uuid"
)

func (d *Data) NewYoutubeClientConfig(config *models.YoutubeClientConfig) error {
	config.UUID = uuid.NewV4().String()

	result := d.gormDB.Create(&config)

	return result.Error
}

func (d *Data) NewYoutubeToken(token *models.YoutubeToken) error {
	token.UUID = uuid.NewV4().String()

	result := d.gormDB.Create(&token)

	return result.Error
}

func (d *Data) GetFirstYoutubeClientConfig() (*models.YoutubeClientConfig, error) {
	config := models.YoutubeClientConfig{}

	result := d.gormDB.First(&config)

	if result.Error != nil || NotFound(result.Error) {
		return nil, RemoveNotFound(result.Error)
	}

	return &config, nil
}
func (d *Data) GetFirstYoutubeToken() (*models.YoutubeToken, error) {
	token := models.YoutubeToken{}

	result := d.gormDB.First(&token)

	if result.Error != nil || NotFound(result.Error) {
		return nil, RemoveNotFound(result.Error)
	}

	return &token, nil
}
