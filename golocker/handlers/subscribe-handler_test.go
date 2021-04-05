package handlers

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/Killian264/YTLocker/golocker/interfaces"
	"github.com/Killian264/YTLocker/golocker/mocks"
	"github.com/Killian264/YTLocker/golocker/models"
	"github.com/Killian264/YTLocker/golocker/parsers"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestHandleValidChallenge(t *testing.T) {

	service := mocks.ISubscription{}
	challenge := "test-challenge"

	sub := &models.SubscriptionRequest{
		ChannelID:    "channel_id",
		LeaseSeconds: uint(12345),
		Topic:        "https://www.youtube.com/xml/feeds/videos.xml?channel_id=channel_id",
		Secret:       "super_secret",
		Active:       true,
	}

	url := createChallenge(fmt.Sprintf("/subscribe/%s/", sub.Secret), challenge, sub.ChannelID, fmt.Sprint(sub.LeaseSeconds))
	req, err := http.NewRequest("GET", url, nil)

	service.On("HandleChallenge", sub).Return(true, nil)

	res := sendFakeRequest(interfaces.ISubscription(&service), req, "/subscribe/{secret}/")

	body, err := ioutil.ReadAll(res.Body)
	assert.Nil(t, err)

	assert.Equal(t, challenge, string(body), "route should respond with challenge")
	assert.Equal(t, 200, res.StatusCode)
}

func TestHandleInValidChallenge(t *testing.T) {

	service := mocks.ISubscription{}
	challenge := "test-challenge"

	sub := &models.SubscriptionRequest{
		ChannelID:    "channel_id",
		LeaseSeconds: uint(12345),
		Topic:        "https://www.youtube.com/xml/feeds/videos.xml?channel_id=channel_id",
		Secret:       "super_secret",
		Active:       true,
	}

	url := createChallenge(fmt.Sprintf("/subscribe/%s/", sub.Secret), challenge, sub.ChannelID, fmt.Sprint(sub.LeaseSeconds))
	req, err := http.NewRequest("GET", url, nil)

	// return invalid
	service.On("HandleChallenge", sub).Return(false, nil).Once()

	res := sendFakeRequest(interfaces.ISubscription(&service), req, "/subscribe/{secret}/")

	body, err := ioutil.ReadAll(res.Body)
	assert.Nil(t, err)

	assert.Equal(t, "", string(body), "route should respond with challenge")
	assert.Equal(t, 500, res.StatusCode)

	// return error
	service.On("HandleChallenge", sub).Return(true, fmt.Errorf("hello")).Once()

	res = sendFakeRequest(interfaces.ISubscription(&service), req, "/subscribe/{secret}/")

	body, err = ioutil.ReadAll(res.Body)
	assert.Nil(t, err)

	assert.Equal(t, "", string(body), "route should respond with challenge")
	assert.Equal(t, 500, res.StatusCode)
}

func TestNewVideoPush(t *testing.T) {

	service := mocks.ISubscription{}

	body := []byte(pushXML)

	push, err := parsers.ParseYTHook(pushXML)
	assert.Nil(t, err)

	req, err := http.NewRequest("GET", fmt.Sprintf("/subscribe/%s/", "test-secret"), bytes.NewBuffer(body))

	service.On("HandleVideoPush", &push, "test-secret").Return(nil).Once()

	res := sendFakeRequest(interfaces.ISubscription(&service), req, "/subscribe/{secret}/")

	assert.Equal(t, 200, res.StatusCode)
}

func createChallenge(route string, challenge string, channelID string, lease_seconds string) string {

	topic := "https://www.youtube.com/xml/feeds/videos.xml?channel_id=" + channelID

	values := url.Values{}
	values.Add("hub.challenge", challenge)
	values.Add("hub.topic", topic)
	values.Add("hub.lease_seconds", lease_seconds)

	params := values.Encode()

	return fmt.Sprintf("%s?%s", route, params)

}

func sendFakeRequest(service interfaces.ISubscription, req *http.Request, route string) *http.Response {
	rr := httptest.NewRecorder()
	router := mux.NewRouter()

	handler := func(w http.ResponseWriter, r *http.Request) {
		err := HandleYoutubePush(w, r, service)

		if err != nil {
			fmt.Print(err.Error(), "\n")
			w.WriteHeader(500)
		}
	}

	router.HandleFunc(route, handler)
	router.ServeHTTP(rr, req)
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
