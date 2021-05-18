package subscribe

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Killian264/YTLocker/golocker/data"
	"github.com/Killian264/YTLocker/golocker/models"
	"github.com/Killian264/YTLocker/golocker/services/ytmanager"
	"github.com/stretchr/testify/assert"
)

var hook = &models.YTHookPush{
	Video: models.YTHookVideo{
		VideoID:   "test-video-id",
		ChannelID: "test-channel-id",
	},
}

var channel = &models.Channel{
	YoutubeID:   "test-channel-id",
	Title:       "not used",
	Description: "not used",
}
var channel2 = &models.Channel{
	YoutubeID:   "test-channel-id2",
	Title:       "not used",
	Description: "not used",
}

func Test_IgnoreDuplicates_Challenge(t *testing.T) {

	service := createMockServices(t)

	sub, err := service.Subscribe(channel)
	assert.Nil(t, err)

	sub2, err := service.Subscribe(channel)
	assert.Nil(t, err)

	assert.Equal(t, sub.ID, sub2.ID)
}

func Test_Valid_Challenge(t *testing.T) {

	service := createMockServices(t)

	sub, err := service.Subscribe(channel)
	assert.Nil(t, err)

	valid, err := service.HandleChallenge(sub, channel.YoutubeID)
	assert.Nil(t, err)
	assert.True(t, valid)
}

func Test_InValid_Challenge(t *testing.T) {
	service := createMockServices(t)

	// fake id specified in ytservice fakes
	valid, err := service.HandleChallenge(&models.SubscriptionRequest{
		ChannelID:    uint64(23423),
		LeaseSeconds: 23423,
		Topic:        "random.com/url",
		Secret:       "one-two-three",
	}, "fake-channel-id")

	assert.NotNil(t, err)
	assert.False(t, valid)
}

func Test_Valid_Video_Push(t *testing.T) {

	service := createMockServices(t)

	sub, _ := service.Subscribe(channel)
	err := service.HandleVideoPush(hook, sub.Secret)

	assert.Nil(t, err)
}

func Test_InValid_Video_Push(t *testing.T) {

	service := createMockServices(t)

	err := service.HandleVideoPush(hook, "super fake secret")

	assert.NotNil(t, err)
}

func Test_InValid_Video_Video_Push(t *testing.T) {

	service := createMockServices(t)

	hook.Video.VideoID = "fake-video-id"

	sub, _ := service.Subscribe(channel)
	err := service.HandleVideoPush(hook, sub.Secret)

	assert.NotNil(t, err)
}

func Test_ResubscribeAll(t *testing.T) {

	service := createMockServices(t)

	sub1, _ := service.Subscribe(channel)
	sub2, _ := service.Subscribe(channel2)

	err := service.ResubscribeAll()
	assert.Nil(t, err)

	sub3, err := service.GetSubscription(channel)
	sub4, _ := service.GetSubscription(channel2)

	assert.NotEqual(t, sub1.ID, sub3.ID)
	assert.NotEqual(t, sub2.ID, sub4.ID)
}

func createMockServices(t *testing.T) *Subscriber {

	db := data.InMemorySQLiteConnect()

	manager := ytmanager.FakeNewYoutubeManager(db)

	service := NewSubscriber(
		db,
		manager,
	)

	service.SetSubscribeUrl("", "/subscribe/{secret}/")
	service.SetYTPubSubUrl(youtubePubSub(t))

	new, _ := manager.NewChannel(channel.YoutubeID)
	new2, _ := manager.NewChannel(channel2.YoutubeID)

	channel = &new
	channel2 = &new2

	return service
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
