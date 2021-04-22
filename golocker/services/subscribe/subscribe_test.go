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

func TestSubscribeValidChannel(t *testing.T) {

	service, yt := createMockServices(t)

	yt.On("GetChannel", "test-channel-id").Return(&models.Channel{}, nil)

	actual, err := service.Subscribe("test-channel-id")
	assert.Nil(t, err)

	saved, err := service.GetSubscription("test-channel-id")
	assert.Nil(t, err)

	assert.Equal(t, actual.ChannelID, saved.ChannelID)
	assert.Equal(t, actual.Secret, saved.Secret)
}

func TestSubscribeInValidChannel(t *testing.T) {

	service, yt := createMockServices(t)

	yt.On("GetChannel", "test-channel-id").Return(nil, nil)

	_, err := service.Subscribe("test-channel-id")
	assert.NotNil(t, err)

	sub, err := service.GetSubscription("test-channel-id")
	assert.Nil(t, sub)
	assert.Nil(t, err)
}

func TestValidChallenge(t *testing.T) {

	service, yt := createMockServices(t)

	yt.On("GetChannel", "test-channel-id").Return(&models.Channel{}, nil)

	sub, _ := service.Subscribe("test-channel-id")

	valid, err := service.HandleChallenge(sub)
	assert.Nil(t, err)
	assert.True(t, valid)
}

func TestInValidChallenge(t *testing.T) {

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

func TestHandleValidVideoPush(t *testing.T) {

	service, yt := createMockServices(t)

	yt.On("GetChannel", "test-channel-id").Return(&models.Channel{}, nil)
	yt.On("CreateVideo", "test-video-id", "test-channel-id").Return(&models.Video{}, nil)

	sub, _ := service.Subscribe("test-channel-id")

	err := service.HandleVideoPush(&models.YTHookPush{
		Video: models.YTHookVideo{
			VideoID:   "test-video-id",
			ChannelID: "test-channel-id",
		},
	}, sub.Secret)

	assert.Nil(t, err)
}

func TestHandleInvalidVideoPush(t *testing.T) {

	service, _ := createMockServices(t)

	err := service.HandleVideoPush(&models.YTHookPush{
		Video: models.YTHookVideo{
			VideoID:   "test-video-id",
			ChannelID: "test-channel-id",
		},
	}, "super fake secret")

	assert.NotNil(t, err)
}

func TestHandleInvalidVideoVideoPush(t *testing.T) {

	service, yt := createMockServices(t)

	yt.On("GetChannel", "test-channel-id").Return(&models.Channel{}, nil)
	yt.On("CreateVideo", "test-video-id", "test-channel-id").Return(nil, fmt.Errorf("123"))

	sub, _ := service.Subscribe("test-channel-id")

	err := service.HandleVideoPush(&models.YTHookPush{
		Video: models.YTHookVideo{
			VideoID:   "test-video-id",
			ChannelID: "test-channel-id",
		},
	}, sub.Secret)

	assert.NotNil(t, err)
}

func TestResubscribeAll(t *testing.T) {

	service, yt := createMockServices(t)

	yt.On("GetChannel", mock.Anything).Return(&models.Channel{}, nil)

	sub1, _ := service.Subscribe("test-channel-id")
	sub2, _ := service.Subscribe("test-channel-id2")

	err := service.ResubscribeAll()
	assert.Nil(t, err)

	sub3, _ := service.GetSubscription("test-channel-id")
	sub4, _ := service.GetSubscription("test-channel-id2")

	assert.NotEqual(t, sub1.Secret, sub3.Secret)
	assert.NotEqual(t, sub2.Secret, sub4.Secret)
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
