package models

import "time"

// Incoming requests
type YTHookTopic struct {
	Challenge    string
	Topic        string
	LeaseSeconds int
	Token        string
	Mode         string
	Body         string
}

// YTHookPush is the struct defining a YT hook push notification
type YTHookPush struct {
	Link    []YTHookLink `xml:"link"`
	Title   string       `xml:"title"`
	Updated time.Time    `xml:"updated"`
	Video   YTHookVideo  `xml:"entry"`
}

// YTHookLink holds link data for a YTHookPush
type YTHookLink struct {
	Rel string `xml:"rel,attr"`
	URL string `xml:"href,attr"`
}

// YTHookVideo holds video data for a YTHookPush
type YTHookVideo struct {
	VideoID   string        `xml:"videoId"`
	ChannelID string        `xml:"channelId"`
	Title     string        `xml:"title"`
	Link      YTHookLink    `xml:"link"`
	Channel   YTHookChannel `xml:"author"`
	Published time.Time     `xml:"published"`
	Updated   time.Time     `xml:"updated"`
}

// YTHookChannel holds channel data for a YTHookPush
type YTHookChannel struct {
	Name string `xml:"name"`
	URL  string `xml:"uri"`
}
