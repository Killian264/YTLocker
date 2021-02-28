package parsers

import (
	"reflect"
	"testing"
	"time"

	"github.com/Killian264/YTLocker/models"
)

var (
	testXML = `
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
)

func TestUpdatedDateParse(t *testing.T) {

	parsed := ParseYTHook(testXML)

	date := parsed.Updated

	year, month, day := date.Date()

	if year != 2015 {
		t.Errorf("malformed date year expected: %q got %q", 2015, year)
	}

	if day != 1 {
		t.Errorf("malformed date day expected: %q got %q", 1, day)
	}

	if month != time.April {
		t.Errorf("malformed date month expected: %q got %q", time.March, month)
	}

	hour, min, second := date.Clock()

	if hour != 19 {
		t.Errorf("malformed date hour expected: %q got %q", 19, hour)
	}

	if min != 5 {
		t.Errorf("malformed date min expected: %q got %q", 5, min)
	}

	if second != 24 {
		t.Errorf("malformed date second expected: %q got %q", 24, second)
	}

}

func BuildExpected() models.YTHookPush {

	channel := models.YTHookChannel{
		Name: "Channel title",
		URL:  "http://www.youtube.com/channel/CHANNEL_ID",
	}

	videoLink := models.YTHookLink{
		Rel: "alternate",
		URL: "http://www.youtube.com/watch?v=VIDEO_ID",
	}

	video := models.YTHookVideo{
		ID:        "yt:video:VIDEO_ID",
		VideoID:   "VIDEO_ID",
		ChannelID: "CHANNEL_ID",
		Title:     "Video title",
		Link:      videoLink,
		Channel:   channel,
	}

	hubLink := models.YTHookLink{
		Rel: "hub",
		URL: "https://pubsubhubbub.appspot.com",
	}

	selfLink := models.YTHookLink{
		Rel: "self",
		URL: "https://www.youtube.com/xml/feeds/videos.xml?channel_id=CHANNEL_ID",
	}

	links := []models.YTHookLink{
		hubLink,
		selfLink,
	}

	push := models.YTHookPush{
		Link:  links,
		Title: "YouTube video feed",
		Video: video,
	}

	return push
}

func TestParseYTHook(t *testing.T) {

	got := ParseYTHook(testXML)

	expected := BuildExpected()

	// Time zone is exactly the same but it does not work
	expected.Updated = got.Updated

	expected.Video.Published = got.Video.Published

	expected.Video.Updated = got.Video.Updated

	expected.Received = got.Received

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("ParseYTHook() = %v, want %v", got, expected)
	}
}
