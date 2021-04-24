package subscribe

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Killian264/YTLocker/golocker/data"
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

var channel = models.Channel{
	YoutubeID:   "test-channel-id",
	Title:       "not used",
	Description: "not used",
}
var channel2 = models.Channel{
	YoutubeID:   "test-channel-id2",
	Title:       "not used",
	Description: "not used",
}

func Test_Valid_Challenge(t *testing.T) {

	service, yt := createMockServices(t)
	yt.On("GetChannelByYoutubeID", channel.YoutubeID).Return(&channel, nil)

	sub, err := service.Subscribe(&channel)
	assert.Nil(t, err)

	valid, err := service.HandleChallenge(sub, channel.YoutubeID)
	assert.Nil(t, err)
	assert.True(t, valid)
}

func Test_InValid_Challenge(t *testing.T) {
	service, yt := createMockServices(t)
	yt.On("GetChannelByYoutubeID", "random fake id").Return(&channel, nil)

	valid, err := service.HandleChallenge(&models.SubscriptionRequest{
		ChannelID:    uint64(23423),
		LeaseSeconds: 23423,
		Topic:        "random.com/url",
		Secret:       "one-two-three",
	}, "random fake id")

	assert.NotNil(t, err)
	assert.False(t, valid)
}

func Test_Valid_Video_Push(t *testing.T) {

	service, yt := createMockServices(t)

	yt.On("GetChannelByYoutubeID", "test-channel-id").Return(&models.Channel{}, nil)
	yt.On("CreateVideo", &models.Channel{}, "test-video-id").Return(&models.Video{}, nil)

	sub, _ := service.Subscribe(&channel)
	err := service.HandleVideoPush(&hook, sub.Secret)

	assert.Nil(t, err)
}

func Test_InValid_Video_Push(t *testing.T) {

	service, yt := createMockServices(t)
	yt.On("GetChannelByYoutubeID", "test-channel-id").Return(&models.Channel{}, nil)

	err := service.HandleVideoPush(&hook, "super fake secret")

	assert.NotNil(t, err)
}

func Test_InValid_Video_Video_Push(t *testing.T) {

	service, yt := createMockServices(t)
	yt.On("GetChannelByYoutubeID", "test-channel-id").Return(&models.Channel{}, nil)
	yt.On("CreateVideo", &models.Channel{}, "test-video-id").Return(nil, fmt.Errorf("123"))

	sub, _ := service.Subscribe(&channel)
	err := service.HandleVideoPush(&hook, sub.Secret)

	assert.NotNil(t, err)
}

func Test_ResubscribeAll(t *testing.T) {

	service, yt := createMockServices(t)

	yt.On("GetChannelByID", mock.Anything).Return(&channel, nil).Once()
	yt.On("GetChannelByID", mock.Anything).Return(&channel2, nil).Once()

	sub1, _ := service.Subscribe(&channel)
	sub2, _ := service.Subscribe(&channel2)

	err := service.ResubscribeAll()
	assert.Nil(t, err)

	sub3, _ := service.GetSubscription(&channel)
	sub4, _ := service.GetSubscription(&channel2)

	assert.NotEqual(t, sub1.ID, sub3.ID)
	assert.NotEqual(t, sub2.ID, sub4.ID)
}

func createMockServices(t *testing.T) (*Subscriber, *mocks.IYoutubeManager) {

	db := data.InMemorySQLiteConnect()
	yt := &mocks.IYoutubeManager{}

	service := NewSubscriber(
		ISubscriptionData(db),
		IYoutubeManager(yt),
	)

	service.SetSubscribeUrl("", "/subscribe/{secret}/")
	service.SetYTPubSubUrl(youtubePubSub(t))

	db.NewChannel(&channel)
	db.NewChannel(&channel2)

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
