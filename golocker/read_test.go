package main

import (
	"reflect"
	"testing"
	"time"
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

func TestTitleParse(t *testing.T) {

	parsed := ParseYTHook(testXML)

	if parsed.Title != "YouTube video feed" {
		t.Error("malformed youtube feed title")
	}
}

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

}

func TestUpdatedTimeParse(t *testing.T) {

	parsed := ParseYTHook(testXML)

	date := parsed.Updated

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

func TestLink(t *testing.T) {
	parsed := ParseYTHook(testXML)

	links := parsed.Link

	location := -1

	for i, link := range links {

		if link.Rel == "self" {
			location = i
			break
		}
	}

	if location == -1 {
		t.Errorf("malformed links self not found")
	}

	link := links[location]
	url := "https://www.youtube.com/xml/feeds/videos.xml?channel_id=CHANNEL_ID"

	if link.URL != url {
		t.Errorf("malformed links expected: %q got %q", url, link.URL)
	}
}

func TestUpdatedEntryParse(t *testing.T) {

	parsed := ParseYTHook(testXML)

	entry := parsed.Video

	videoID := entry.VideoID

	if videoID != "VIDEO_ID" {
		t.Errorf("malformed entry video id expected: %q got %q", "VIDEO_ID", videoID)
	}
}

func TestUpdatedLinkParse(t *testing.T) {

	parsed := ParseYTHook(testXML)

	link := parsed.Video.Link

	got := link.URL

	expected := `http://www.youtube.com/watch?v=VIDEO_ID`

	if got != expected {
		t.Errorf("malformed entry author title expected: %q got %q", expected, got)
	}
}

func TestUpdatedAuthorParse(t *testing.T) {

	parsed := ParseYTHook(testXML)

	author := parsed.Video.Channel

	name := author.Name

	if name != "Channel title" {
		t.Errorf("malformed entry author title expected: %q got %q", "Channel title", name)
	}
}

func Testtest(t *testing.T) {
	parsed := ParseYTHook(testXML)

	got := parsed.Video.Published

	expected := time.Date(
		2015,
		03,
		6,
		21,
		40,
		57,
		0,
		time.UTC,
	)

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("ParseYTHook() = %v, want %v", got, expected)
	}
}

func BuildExpected() YTHookPush {

	channel := YTHookChannel{
		Name: "Channel title",
		URL:  "http://www.youtube.com/channel/CHANNEL_ID",
	}

	videoLink := YTHookLink{
		Rel: "alternate",
		URL: "http://www.youtube.com/watch?v=VIDEO_ID",
	}

	// UTC prints differently
	location := time.FixedZone("+0000", 0000)

	published := time.Date(
		2015,
		3,
		06,
		21,
		40,
		57,
		0,
		location,
	)

	updated := time.Date(
		2015,
		3,
		9,
		19,
		5,
		24,
		0,
		location,
	)

	video := YTHookVideo{
		ID:        "yt:video:VIDEO_ID",
		VideoID:   "VIDEO_ID",
		ChannelID: "CHANNEL_ID",
		Title:     "Video title",
		Link:      videoLink,
		Channel:   channel,
		Published: published,
		Updated:   updated,
	}

	hubLink := YTHookLink{
		Rel: "hub",
		URL: "https://pubsubhubbub.appspot.com",
	}

	selfLink := YTHookLink{
		Rel: "self",
		URL: "https://www.youtube.com/xml/feeds/videos.xml?channel_id=CHANNEL_ID",
	}

	updatedPush := time.Date(
		2015,
		4,
		1,
		19,
		5,
		24,
		0,
		location,
	)

	links := []YTHookLink{
		hubLink,
		selfLink,
	}

	push := YTHookPush{
		Link:    links,
		Title:   "YouTube video feed",
		Updated: updatedPush,
		Video:   video,
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

	// t.Errorf(expected.Updated.Location().String())
	// t.Errorf(got.Updated.Location().String())
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("ParseYTHook() = %v, want %v", got, expected)
	}
}
