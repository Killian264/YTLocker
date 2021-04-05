package subscribe

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/Killian264/YTLocker/golocker/interfaces"
	"github.com/Killian264/YTLocker/golocker/mocks"
	"github.com/Killian264/YTLocker/golocker/models"
	"github.com/Killian264/YTLocker/golocker/parsers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/api/youtube/v3"
)

func TestCreateSubscription(t *testing.T) {

	service, _, _ := createMockServices()

	sub := models.SubscriptionRequest{
		ChannelID:    "superchannelid",
		LeaseSeconds: 8 * 24 * 60 * 60,
		Topic:        "https://www.youtube.com/xml/feeds/videos.xml?channel_id=superchannelid",
		Secret:       "supersecrettokenthatshouldbelikeidk64characterslongorsomething",
		Active:       true,
	}

	got, err := service.CreateSubscription(sub.ChannelID)
	got.Secret = sub.Secret

	assert.Nil(t, err)
	assert.Equal(t, sub, *got)
}

func TestSubscribe(t *testing.T) {

	service, data, _ := createMockServices()

	sub, err := service.CreateSubscription("channel-id")
	assert.Nil(t, err)

	service.SetSubscribeUrl("", "/subscribe/{secret}/")

	handler := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(201)
		assert.Equal(t, sub.Topic, r.FormValue("hub.topic"))
		assert.Contains(t, r.FormValue("hub.callback"), fmt.Sprintf("/subscribe/%s/", sub.Secret))
		assert.Equal(t, fmt.Sprint(sub.LeaseSeconds), r.FormValue("hub.lease_seconds"))
	}

	server := httptest.NewServer(http.HandlerFunc(handler))

	service.SetYTPubSubUrl(server.URL)

	data.On("GetChannel", "channel-id").Return(&models.Channel{}, nil)
	data.On("NewSubscription", sub).Return(nil)

	err = service.Subscribe(sub)
	assert.Nil(t, err)
}

func TestHandleChallenge(t *testing.T) {

	service, data, _ := createMockServices()

	sub, err := service.CreateSubscription("channel-id")
	assert.Nil(t, err)

	data.On("GetSubscription", sub.Secret, sub.ChannelID).Return(sub, nil).Once()
	isValid, err := service.HandleChallenge(sub)
	assert.Nil(t, err)
	assert.True(t, isValid)

	data.On("GetSubscription", sub.Secret, sub.ChannelID).Return(nil, nil).Once()
	isValid, err = service.HandleChallenge(sub)
	assert.NotNil(t, err)
	assert.False(t, isValid)

	data.On("GetSubscription", sub.Secret, sub.ChannelID).Return(nil, fmt.Errorf("hello")).Once()
	isValid, err = service.HandleChallenge(sub)
	assert.NotNil(t, err)
	assert.False(t, isValid)
}

func TestNewVideoWithSave(t *testing.T) {

	service, data, yt := createMockServices()

	push, err := parsers.ParseYTHook(pushXML)
	assert.Nil(t, err)

	video := youtube.Video{
		Id: "VIDEO_ID",
		Snippet: &youtube.VideoSnippet{
			ChannelId:   "CHANNEL_ID",
			Title:       "Test title",
			Description: "Test description",
			Thumbnails:  &youtube.ThumbnailDetails{},
		},
	}

	parsed, channelID := parsers.ParseYTVideo(&video)

	yt.On("GetVideo", "VIDEO_ID").Return(&video, nil).Once()
	data.On("NewVideo", parsed, channelID).Return(nil).Once()
	data.On("GetSubscription", "test-secret", push.Video.ChannelID).Return(&models.SubscriptionRequest{}, nil).Once()

	err = service.HandleVideoPush(&push, "test-secret")
	assert.Nil(t, err)

	data.AssertExpectations(t)
}

func TestNewVideoWithoutSave(t *testing.T) {

	service, data, yt := createMockServices()

	push, err := parsers.ParseYTHook(pushXML)
	assert.Nil(t, err)

	yt.On("GetVideo", "VIDEO_ID").Return(nil, nil)
	data.On("GetSubscription", "test-secret", push.Video.ChannelID).Return(&models.SubscriptionRequest{}, nil).Once()

	err = service.HandleVideoPush(&push, "test-secret")
	assert.NotNil(t, err)

	data.AssertExpectations(t)
}

func TestResubscribeAll(t *testing.T) {
	service, data, _ := createMockServices()

	// Subscribe Function
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(201)
	}))
	data.On("GetChannel", mock.Anything).Return(&models.Channel{}, nil)
	data.On("NewSubscription", mock.Anything, mock.Anything).Return(nil)
	service.SetYTPubSubUrl(server.URL)

	sub := models.SubscriptionRequest{
		UUID:         "213",
		ChannelID:    "superchannelid",
		LeaseSeconds: 8 * 24 * 60 * 60,
		Topic:        "https://www.youtube.com/xml/feeds/videos.xml?channel_id=superchannelid",
		Secret:       "supersecrettokenthatshouldbelikeidk64characterslongorsomething",
		Active:       true,
	}

	// ResubscribeAll
	data.On("InactivateAllSubscriptions").Return(nil).Once()
	data.On("GetInactiveSubscription").Return(&sub, nil).Once()
	data.On("GetInactiveSubscription").Return(nil, nil).Once()
	data.On("DeleteSubscription", &sub).Return(nil).Once()

	err := service.ResubscribeAll()
	assert.Nil(t, err)

	data.AssertExpectations(t)
}

// **************************************** HELPERS **************************************** //
// **************************************** HELPERS **************************************** //
func createMockServices() (*Subscriber, *mocks.ISubscriptionData, *mocks.IYoutubeService) {

	dataMock := &mocks.ISubscriptionData{}
	ytMock := &mocks.IYoutubeService{}

	service := NewSubscriber(
		interfaces.ISubscriptionData(dataMock),
		interfaces.IYoutubeService(ytMock),
		log.New(os.Stdout, "Subscriber: ", log.Ldate|log.Ltime|log.Lshortfile),
	)

	return service, dataMock, ytMock
}

var pushXML = `
	<feed xmlns:yt="http://www.youtube.com/xml/schemas/2015" xmlns="http://www.w3.org/2005/Atom">
		<link rel="hub" href="https://pubsubhubbub.appspot.com"/>
		<link rel="self" href="https://www.youtube.com/xml/feeds/videos.xml?channel_id=CHANNEL_ID"/>
		<title>YouTube video feed</title>
		<updated>2015-04-01T19:05:24+00:00</updated>
		<entry>
			<id>yt:video:VIDEO_ID</id>
			<yt:videoId>VIDEO_ID</yt:videoId>
			<yt:channelId>CHANNEL_ID</yt:channelId>
			<title>Video title</title>
			<link rel="alternate" href="http://www.youtube.com/watch?v=VIDEO_ID"/>
			<author>
				<name>Channel title</name>
				<uri>http://www.youtube.com/channel/CHANNEL_ID</uri>
			</author>
			<published>2015-03-06T21:40:57+00:00</published>
			<updated>2015-03-09T19:05:24+00:00</updated>
		</entry>
	</feed>
	`
