package subscribe

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"

	"github.com/Killian264/YTLocker/golocker/interfaces"
	"github.com/Killian264/YTLocker/golocker/mocks"
	"github.com/Killian264/YTLocker/golocker/models"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/api/youtube/v3"
)

var (
	pushHandlerBase  = "/subscribe/{secret}/"
	pushHandlerPrint = "/subscribe/%s/"
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

	assert.Nil(t, err)

	got.Secret = sub.Secret

	assert.Equal(t, sub, *got)

}

func TestSubscribe(t *testing.T) {

	service, data, _ := createMockServices()

	channelID := "channel-id"

	sub, err := service.CreateSubscription(channelID)
	assert.Nil(t, err)

	handler := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(201)
		assert.Equal(t, sub.Topic, r.FormValue("hub.topic"))
		assert.Contains(t, r.FormValue("hub.callback"), fmt.Sprintf(pushHandlerPrint, sub.Secret))
		assert.Equal(t, fmt.Sprint(sub.LeaseSeconds), r.FormValue("hub.lease_seconds"))
	}

	server := httptest.NewServer(http.HandlerFunc(handler))

	service.SetYTPubSubUrl(server.URL)

	data.On("ChannelExists", channelID).Return(true, nil)
	data.On("SaveSubscription", sub).Return(nil)

	err = service.Subscribe(sub)
	assert.Nil(t, err)

}

func TestValidChallenge(t *testing.T) {

	service, data, _ := createMockServices()
	challenge := "test-challenge"

	sub, err := service.CreateSubscription("channel-id")
	assert.Nil(t, err)

	url := createFakeChallenge(fmt.Sprintf(pushHandlerPrint, sub.Secret), challenge, sub.ChannelID, fmt.Sprint(sub.LeaseSeconds))
	req, err := http.NewRequest("GET", url, nil)
	assert.Nil(t, err)

	data.On("GetSubscription", sub.Secret, sub.ChannelID).Return(sub, nil)

	res := sendFakeRequest(service, *req)
	bytes, err := ioutil.ReadAll(res.Body)
	assert.Nil(t, err)
	body := string(bytes)

	assert.Nil(t, err, "should not error")
	assert.Equal(t, challenge, body, "route should respond with challenge")
	assert.Equal(t, 200, res.StatusCode)

}

func TestInvalidChallenge(t *testing.T) {

	service, data, _ := createMockServices()
	challenge := "test-challenge"

	sub, err := service.CreateSubscription("channel-id")
	assert.Nil(t, err)

	url := createFakeChallenge(fmt.Sprintf(pushHandlerPrint, sub.Secret), challenge, sub.ChannelID, fmt.Sprint(sub.LeaseSeconds))
	req, err := http.NewRequest("GET", url, nil)
	assert.Nil(t, err)

	data.On("GetSubscription", sub.Secret, sub.ChannelID).Return(nil, nil)

	res := sendFakeRequest(service, *req)
	bytes, err := ioutil.ReadAll(res.Body)
	assert.Nil(t, err)

	body := string(bytes)

	assert.Nil(t, err, "should not error")
	assert.Equal(t, "", body, "route should not respond with challenge")
	assert.Equal(t, 500, res.StatusCode)

}

func TestNewVideoWithSave(t *testing.T) {

	service, data, yt := createMockServices()

	secret := "test-secret"

	bodyBytes := []byte(pushXML)

	req, err := http.NewRequest("GET", fmt.Sprintf("/subscribe/%s/", secret), bytes.NewBuffer(bodyBytes))
	assert.Nil(t, err)

	video := youtube.Video{
		Id: "VIDEO_ID",
		Snippet: &youtube.VideoSnippet{
			ChannelId: "CHANNEL_ID",
		},
	}

	yt.On("GetVideo", "VIDEO_ID").Return(&video, nil)
	data.On("SaveVideo", &video).Return(nil)

	response := sendFakeRequest(service, *req)

	assert.Equal(t, 200, response.StatusCode)

}

func TestNewVideoWithoutSave(t *testing.T) {

	service, _, yt := createMockServices()

	secret := "test-secret"

	bodyBytes := []byte(pushXML)

	req, err := http.NewRequest("GET", fmt.Sprintf("/subscribe/%s/", secret), bytes.NewBuffer(bodyBytes))
	assert.Nil(t, err)

	yt.On("GetVideo", "VIDEO_ID").Return(nil, nil)

	response := sendFakeRequest(service, *req)

	assert.Equal(t, 500, response.StatusCode)

}

func TestResubscribeAll(t *testing.T) {
	service, data, _ := createMockServices()

	// Subscribe Function
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(201)
	}))
	data.On("ChannelExists", mock.Anything).Return(true, nil)
	data.On("SaveSubscription", mock.Anything).Return(nil)
	service.SetYTPubSubUrl(server.URL)

	sub := models.SubscriptionRequest{
		ID:           213,
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
func createFakeChallenge(route string, challenge string, channelID string, lease_seconds string) string {

	topic := "https://www.youtube.com/xml/feeds/videos.xml?channel_id=" + channelID

	values := url.Values{}
	values.Add("hub.challenge", challenge)
	values.Add("hub.topic", topic)
	values.Add("hub.lease_seconds", lease_seconds)

	params := values.Encode()

	return fmt.Sprintf("%s?%s", route, params)

}

func createMockServices() (*Subscriber, *mocks.ISubscriptionData, *mocks.IYoutubeService) {

	dataMock := &mocks.ISubscriptionData{}
	ytMock := &mocks.IYoutubeService{}

	data := interfaces.ISubscriptionData(dataMock)
	yt := interfaces.IYoutubeService(ytMock)
	logger := log.New(os.Stdout, "Subscriber: ", log.Ldate|log.Ltime|log.Lshortfile)

	service := &Subscriber{
		dataService: data,
		ytService:   yt,
		logger:      logger,
	}

	service.SetSubscribeUrl("", pushHandlerBase)

	return service, dataMock, ytMock
}

func sendFakeRequest(subscriber *Subscriber, req http.Request) *http.Response {

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc(
		pushHandlerBase,
		func(w http.ResponseWriter, r *http.Request) {
			err := subscriber.HandleSubscription(w, r)

			if err != nil {
				fmt.Print(err.Error(), "\n")
				w.WriteHeader(500)
			}
		},
	)

	router.ServeHTTP(rr, &req)

	return rr.Result()
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
