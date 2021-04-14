package playlist

import (
	"log"
	"os"
	"testing"

	"github.com/Killian264/YTLocker/golocker/interfaces"
	"github.com/Killian264/YTLocker/golocker/mocks"
	"github.com/Killian264/YTLocker/golocker/models"
	"github.com/stretchr/testify/assert"
)

func TestSetDefaultConfig(t *testing.T) {
	service, data := createMockServices()

	input := models.YoutubeClientConfig{
		ClientID:     "11223534584-asdfhasdjfhwieyrwqejhkflasd.apps.googleusercontent.com",
		ClientSecret: "qwerHSwer_asdhwuerJHFDJqkqw",
		RedirectURL:  "urn:ietf:wg:oauth:2.0:oob",
		Scope:        "youtube.com/wowee",
		AuthURL:      "https://accounts.google.com/o/oauth2/auth",
		TokenURL:     "https://oauth2.googleapis.com/token",
	}

	data.On("GetYoutubeClientConfigByClientID", "11223534584-asdfhasdjfhwieyrwqejhkflasd.apps.googleusercontent.com").Return(nil, nil)
	data.On("NewYoutubeClientConfig", &input).Return(nil)

	err := service.SetDefaultConfig(input)

	assert.Nil(t, err)
}

func TestSetDuplicateDefaultConfig(t *testing.T) {
	service, data := createMockServices()

	input := models.YoutubeClientConfig{
		ClientID:     "11223534584-asdfhasdjfhwieyrwqejhkflasd.apps.googleusercontent.com",
		ClientSecret: "qwerHSwer_asdhwuerJHFDJqkqw",
		RedirectURL:  "urn:ietf:wg:oauth:2.0:oob",
		Scope:        "youtube.com/wowee",
		AuthURL:      "https://accounts.google.com/o/oauth2/auth",
		TokenURL:     "https://oauth2.googleapis.com/token",
	}

	data.On("GetYoutubeClientConfigByClientID", "11223534584-asdfhasdjfhwieyrwqejhkflasd.apps.googleusercontent.com").Return(&input, nil)

	err := service.SetDefaultConfig(input)

	assert.NotNil(t, err)
}

func TestSetDefaultToken(t *testing.T) {
	service, data := createMockServices()

	input := models.YoutubeToken{
		AccessToken:  "sa23.345234524623sdfasdfq-qege",
		TokenType:    "Bearer",
		RefreshToken: "asdfjqwekj23//2342329asqq-ajfdki22399jjIjiJIWJFfwerw_qwefdasferw_zwaehwejlkWW",
		Expiry:       "2021-04-13T23:30:06.1139442-05:00",
	}

	data.On("GetYoutubeTokenByAccessToken", "sa23.345234524623sdfasdfq-qege").Return(nil, nil)
	data.On("NewYoutubeToken", &input).Return(nil)

	err := service.SetDefaultToken(input)
	assert.Nil(t, err)
}

func TestSetDuplicateDefaultToken(t *testing.T) {
	service, data := createMockServices()

	input := models.YoutubeToken{
		AccessToken:  "sa23.345234524623sdfasdfq-qege",
		TokenType:    "Bearer",
		RefreshToken: "asdfjqwekj23//2342329asqq-ajfdki22399jjIjiJIWJFfwerw_qwefdasferw_zwaehwejlkWW",
		Expiry:       "2021-04-13T23:30:06.1139442-05:00",
	}

	data.On("GetYoutubeTokenByAccessToken", "sa23.345234524623sdfasdfq-qege").Return(&input, nil)

	err := service.SetDefaultToken(input)
	assert.NotNil(t, err)
}

func createMockServices() (*Playlister, *mocks.IPlaylistData) {
	dataMock := &mocks.IPlaylistData{}
	service := NewPlaylister(
		interfaces.IPlaylistData(dataMock),
		log.New(os.Stdout, "Subscriber: ", log.Ldate|log.Ltime|log.Lshortfile),
	)

	return service, dataMock
}
