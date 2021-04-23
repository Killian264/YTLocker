package subscribe

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Killian264/YTLocker/golocker/data"
	"github.com/Killian264/YTLocker/golocker/interfaces"
	"github.com/Killian264/YTLocker/golocker/mocks"
	"github.com/Killian264/YTLocker/golocker/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var hook = models.YTHookPush{
	Video: models.YTHookVideo{
		VideoID:   "test-video-id",
		ChannelID: "test-channel-id",
	},
}

func Test_Valid_Challenge(t *testing.T) {

	service, yt := createMockServices(t)

	yt.On("GetChannel", "test-channel-id").Return(&models.Channel{}, nil)

	sub, _ := service.Subscribe("test-channel-id")

	valid, err := service.HandleChallenge(sub)
	assert.Nil(t, err)
	assert.True(t, valid)
}

func Test_InValid_Challenge(t *testing.T) {

	service, _ := createMockServices(t)

	valid, err := service.HandleChallenge(&models.SubscriptionRequest{
		ChannelID:    "test-channel-id",
		LeaseSeconds: 23423,
		Topic:        "random.com/url",
		Secret:       "one-two-three",
	})
	assert.NotNil(t, err)
	assert.False(t, valid)
}

func Test_InValid_Channel_Challenge(t *testing.T) {

	service, yt := createMockServices(t)

	yt.On("GetChannel", "test-channel-id").Return(nil, nil)

	_, err := service.Subscribe("test-channel-id")
	assert.NotNil(t, err)

	sub, _ := service.GetSubscription("test-channel-id")
	assert.Nil(t, sub)
}

func Test_Valid_Video_Push(t *testing.T) {

	service, yt := createMockServices(t)

	yt.On("GetChannel", "test-channel-id").Return(&models.Channel{}, nil)
	yt.On("CreateVideo", "test-video-id", "test-channel-id").Return(&models.Video{}, nil)

	sub, _ := service.Subscribe("test-channel-id")
	err := service.HandleVideoPush(&hook, sub.Secret)

	assert.Nil(t, err)
}

func Test_InValid_Video_Push(t *testing.T) {

	service, _ := createMockServices(t)

	err := service.HandleVideoPush(&hook, "super fake secret")

	assert.NotNil(t, err)
}

func Test_InValid_Video_Video_Push(t *testing.T) {

	service, yt := createMockServices(t)

	yt.On("GetChannel", "test-channel-id").Return(&models.Channel{}, nil)
	yt.On("CreateVideo", "test-video-id", "test-channel-id").Return(nil, fmt.Errorf("123"))

	sub, _ := service.Subscribe("test-channel-id")
	err := service.HandleVideoPush(&hook, sub.Secret)

	assert.NotNil(t, err)
}

func Test_ResubscribeAll(t *testing.T) {

	service, yt := createMockServices(t)

	yt.On("GetChannel", mock.Anything).Return(&models.Channel{}, nil)

	sub1, _ := service.Subscribe("test-channel-id")
	sub2, _ := service.Subscribe("test-channel-id2")

	err := service.ResubscribeAll()
	assert.Nil(t, err)

	sub3, _ := service.GetSubscription("test-channel-id")
	sub4, _ := service.GetSubscription("test-channel-id2")

	assert.NotEqual(t, sub1.UUID, sub3.UUID)
	assert.NotEqual(t, sub2.UUID, sub4.UUID)
}

func createMockServices(t *testing.T) (*Subscriber, *mocks.IYoutubeManager) {

	db := data.SQLiteConnectAndInitalize()
	yt := &mocks.IYoutubeManager{}

	service := NewSubscriber(
		interfaces.ISubscriptionData(db),
		interfaces.IYoutubeManager(yt),
	)

	service.SetSubscribeUrl("", "/subscribe/{secret}/")
	service.SetYTPubSubUrl(youtubePubSub(t))

	return service, yt
}

func youtubePubSub(t *testing.T) string {

	handler := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(201)
		assert.NotEmpty(t, r.FormValue("hub.topic"))
		assert.NotEmpty(t, r.FormValue("hub.callback"))
		assert.NotEmpty(t, r.FormValue("hub.lease_seconds"))
		assert.NotEmpty(t, r.FormValue("hub.mode"))
	}

	server := httptest.NewServer(http.HandlerFunc(handler))

	return server.URL
}
