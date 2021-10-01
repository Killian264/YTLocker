package oauthmanager

import (
	"fmt"
	"io/ioutil"

	"github.com/Killian264/YTLocker/golocker/helpers/parsers"
	"github.com/Killian264/YTLocker/golocker/models"
)

func (m OauthManager) readInOauthConfigData(secretsDir string) (models.YoutubeClientConfig, models.YoutubeToken) {
	config, err := readInClientSecret(fmt.Sprintf("%s%s", secretsDir, "client_secret.json"))
	if err != nil {
		panic(err)
	}

	token, err := readInAccessToken(fmt.Sprintf("%s%s", secretsDir, "access_secret.json"))
	if err != nil {
		panic(err)
	}

	return config, token
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
