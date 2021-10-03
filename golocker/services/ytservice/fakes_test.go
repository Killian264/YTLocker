package ytservice

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"golang.org/x/oauth2"
)

var config = oauth2.Config{}
var token = oauth2.Token{}

func Test_Initialize_With_Base_Token(t *testing.T) {
	service := createMockService()

	token2, err := service.Initialize(config, token)
	assert.Nil(t, err)
	assert.Equal(t, token.AccessToken, token2.AccessToken)
}

func Test_Initialize_Expired_Is_Refreshed(t *testing.T) {
	service := createMockService()

	token.Expiry = time.Now().AddDate(-1, 1, 1)

	token2, err := service.Initialize(config, token)
	assert.Nil(t, err)
	assert.NotEqual(t, token.AccessToken, token2.AccessToken)
	assert.NotEqual(t, token.Expiry, token2.Expiry)
}

func Test_Initialize_Cannot_Refresh_Without_Refresh_Token(t *testing.T) {
	service := createMockService()

	token.Expiry = time.Now().AddDate(-1, 1, 1)
	token.RefreshToken = ""

	_, err := service.Initialize(config, token)
	assert.NotNil(t, err)
}

func Test_Get_User_Returns_Different_Users(t *testing.T) {
	service := createMockService()

	service.Initialize(config, token)
	user, err := service.GetUser()
	assert.Nil(t, err)

	token.AccessToken = "237051094851092380436034"

	service.Initialize(config, token)
	user2, err := service.GetUser()
	assert.Nil(t, err)

	assert.NotEqual(t, user.Email, user2.Email)
}

func Test_Get_User_Always_Returns_Same(t *testing.T) {
	service := createMockService()

	token, err := service.Initialize(config, token)
	user, err := service.GetUser()
	assert.Nil(t, err)
	assert.NotEmpty(t, user.Email)

	service.Initialize(config, token)
	user2, err := service.GetUser()
	assert.Nil(t, err)

	assert.Equal(t, user.Email, user2.Email)

	token.Expiry = time.Now().AddDate(-1, 1, 1)

	service.Initialize(config, token)
	user3, err := service.GetUser()
	assert.Nil(t, err)

	assert.Equal(t, user2.Email, user3.Email)
}

func Test_Get_Channel_Always_Returns_Same(t *testing.T) {
	service := createMockService()

	service.Initialize(config, token)
	channel, err := service.GetChannel()
	assert.Nil(t, err)

	service.Initialize(config, token)
	channel2, err := service.GetChannel()
	assert.Nil(t, err)

	assert.Equal(t, channel.Snippet.Title, channel2.Snippet.Title)
}

func createMockService() *YTPlaylistFake {
	config = oauth2.Config{
		ClientID:     "11223534584-asdfhasdjfhwieyrwqejhkflasd.apps.googleusercontent.com",
		ClientSecret: "qwerHSwer_asdhwuerJHFDJqkqw",
		RedirectURL:  "https://ytlocker.com/",
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://accounts.google.com/o/oauth2/auth",
			TokenURL: "https://oauth2.googleapis.com/token",
		},
	}

	token = oauth2.Token{
		AccessToken:  "sa23.345234524623sdfasdfq-qegehgower9505034jfeworrjwertw_qqwerjfldssgert345sdgdgew-bheiyqeotleqjrljdfluao23423_QwekjfuI023kjasdfwer",
		TokenType:    "Bearer",
		RefreshToken: "asdfjqwekj23//2342329asqq-ajfdki22399jjIjiJIWJFfwerw_qwefdasferw_zwaehwejlkWW",
		Expiry:       time.Now().AddDate(0, 0, 7),
	}

	return NewYTPlaylistFake()
}
