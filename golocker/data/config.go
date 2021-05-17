package data

import "github.com/Killian264/YTLocker/golocker/models"

func (d *Data) NewYoutubeClientConfig(config *models.YoutubeClientConfig) error {
	config.ID = d.rand.ID()

	result := d.db.Create(&config)

	return result.Error
}

func (d *Data) NewYoutubeToken(token *models.YoutubeToken) error {
	token.ID = d.rand.ID()

	result := d.db.Create(&token)

	return result.Error
}

func (d *Data) GetFirstYoutubeClientConfig() (models.YoutubeClientConfig, error) {
	config := models.YoutubeClientConfig{}

	result := d.db.First(&config)

	if result.Error != nil || notFound(result.Error) {
		return models.YoutubeClientConfig{}, removeNotFound(result.Error)
	}

	return config, nil
}
func (d *Data) GetFirstYoutubeToken() (models.YoutubeToken, error) {
	token := models.YoutubeToken{}

	result := d.db.First(&token)

	if result.Error != nil || notFound(result.Error) {
		return models.YoutubeToken{}, removeNotFound(result.Error)
	}

	return token, nil
}