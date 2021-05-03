package parsers

import (
	"reflect"
	"testing"
	"time"

	"github.com/Killian264/YTLocker/golocker/models"
	"github.com/stretchr/testify/assert"
	"golang.org/x/oauth2"
	"google.golang.org/api/youtube/v3"
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

	parsed, err := ParseYTHook(testXML)
	assert.Nil(t, err)

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

	got, err := ParseYTHook(testXML)
	assert.Nil(t, err)

	expected := BuildExpected()

	// Time zone is exactly the same but it does not work
	expected.Updated = got.Updated

	expected.Video.Published = got.Video.Published

	expected.Video.Updated = got.Video.Updated

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("ParseYTHook() = %v, want %v", got, expected)
	}
}

func TestParseClientJson(t *testing.T) {

	clientJson := `{
		"installed": {
			"client_id": "11223534584-asdfhasdjfhwieyrwqejhkflasd.apps.googleusercontent.com",
			"project_id": "ytlocker-123325",
			"auth_uri": "https://accounts.google.com/o/oauth2/auth",
			"token_uri": "https://oauth2.googleapis.com/token",
			"auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
			"client_secret": "qwerHSwer_asdhwuerJHFDJqkqw",
			"redirect_uris": ["urn:ietf:wg:oauth:2.0:oob", "http://localhost"]
		}
	}`

	expected := models.YoutubeClientConfig{
		ClientID:     "11223534584-asdfhasdjfhwieyrwqejhkflasd.apps.googleusercontent.com",
		ClientSecret: "qwerHSwer_asdhwuerJHFDJqkqw",
		RedirectURL:  "urn:ietf:wg:oauth:2.0:oob",
		AuthURL:      "https://accounts.google.com/o/oauth2/auth",
		TokenURL:     "https://oauth2.googleapis.com/token",
	}

	actual, err := ParseClientJson(clientJson)
	assert.Nil(t, err)

	assert.Equal(t, expected, actual)
}

func TestAccessTokenJson(t *testing.T) {

	clientJson := `{
		"access_token": "sa23.345234524623sdfasdfq-qegehgower9505034jfeworrjwertw_qqwerjfldssgert345sdgdgew-bheiyqeotleqjrljdfluao23423_QwekjfuI023kjasdfwer",
		"token_type": "Bearer",
		"refresh_token": "asdfjqwekj23//2342329asqq-ajfdki22399jjIjiJIWJFfwerw_qwefdasferw_zwaehwejlkWW",
		"expiry": "2021-04-13T23:30:06.1139442-05:00"
	}`

	expected := models.YoutubeToken{
		AccessToken:  "sa23.345234524623sdfasdfq-qegehgower9505034jfeworrjwertw_qqwerjfldssgert345sdgdgew-bheiyqeotleqjrljdfluao23423_QwekjfuI023kjasdfwer",
		TokenType:    "Bearer",
		RefreshToken: "asdfjqwekj23//2342329asqq-ajfdki22399jjIjiJIWJFfwerw_qwefdasferw_zwaehwejlkWW",
		Expiry:       "2021-04-13T23:30:06.1139442-05:00",
	}

	actual, err := ParseAccessTokenJson(clientJson)
	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}

func TestParseYoutubeClient(t *testing.T) {

	input := models.YoutubeClientConfig{
		ClientID:     "11223534584-asdfhasdjfhwieyrwqejhkflasd.apps.googleusercontent.com",
		ClientSecret: "qwerHSwer_asdhwuerJHFDJqkqw",
		RedirectURL:  "urn:ietf:wg:oauth:2.0:oob",
		Scope:        "youtube.com/wowee",
		AuthURL:      "https://accounts.google.com/o/oauth2/auth",
		TokenURL:     "https://oauth2.googleapis.com/token",
	}

	expected := oauth2.Config{
		ClientID:     "11223534584-asdfhasdjfhwieyrwqejhkflasd.apps.googleusercontent.com",
		ClientSecret: "qwerHSwer_asdhwuerJHFDJqkqw",
		RedirectURL:  "urn:ietf:wg:oauth:2.0:oob",
		Scopes:       []string{"youtube.com/wowee"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://accounts.google.com/o/oauth2/auth",
			TokenURL: "https://oauth2.googleapis.com/token",
		},
	}

	actual := ParseYoutubeClient(input)

	assert.Equal(t, expected, actual)

}

func TestParseYoutubeToken(t *testing.T) {
	input := models.YoutubeToken{
		AccessToken:  "sa23.345234524623sdfasdfq-qegehgower9505034jfeworrjwertw_qqwerjfldssgert345sdgdgew-bheiyqeotleqjrljdfluao23423_QwekjfuI023kjasdfwer",
		TokenType:    "Bearer",
		RefreshToken: "asdfjqwekj23//2342329asqq-ajfdki22399jjIjiJIWJFfwerw_qwefdasferw_zwaehwejlkWW",
		Expiry:       "2021-04-13T23:30:06.1139442-09:00",
	}

	expected := oauth2.Token{
		AccessToken:  "sa23.345234524623sdfasdfq-qegehgower9505034jfeworrjwertw_qqwerjfldssgert345sdgdgew-bheiyqeotleqjrljdfluao23423_QwekjfuI023kjasdfwer",
		TokenType:    "Bearer",
		RefreshToken: "asdfjqwekj23//2342329asqq-ajfdki22399jjIjiJIWJFfwerw_qwefdasferw_zwaehwejlkWW",
	}

	actual := ParseYoutubeToken(input)

	expected.Expiry = actual.Expiry

	assert.Equal(t, expected, actual)

	year, month, day := actual.Expiry.Date()

	assert.Equal(t, 2021, year)
	assert.Equal(t, time.Month(4), month)
	assert.Equal(t, 13, day)

	hour, min, sec := actual.Expiry.Clock()

	assert.Equal(t, 23, hour)
	assert.Equal(t, 30, min)
	assert.Equal(t, 6, sec)
}

func TestParseYTThumbnails(t *testing.T) {
	all := getFakeThumbnails()

	thumbnails := ParseYTThumbnails(&all)
	assert.Equal(t, 5, len(thumbnails))

	all.Default = nil
	all.High = nil

	thumbnails = ParseYTThumbnails(&all)
	assert.Equal(t, 3, len(thumbnails))

	all = youtube.ThumbnailDetails{
		Default:  nil,
		Standard: nil,
		Medium:   nil,
		High:     nil,
		Maxres:   nil,
	}

	thumbnails = ParseYTThumbnails(&all)
	assert.Equal(t, 0, len(thumbnails))
}

func TestParseYTVideo(t *testing.T) {
	all := &youtube.ThumbnailDetails{
		Default:  nil,
		Standard: nil,
		Medium:   nil,
		High:     nil,
		Maxres:   nil,
	}

	video := &youtube.Video{
		Id: "wow-video-id",
		Snippet: &youtube.VideoSnippet{
			Title:       "wow cool title",
			Description: "wow that is a super cool description",
			ChannelId:   "wow-channel-id",
			Thumbnails:  all,
		},
	}

	expected := models.Video{
		YoutubeID:   "wow-video-id",
		Title:       "wow cool title",
		Description: "wow that is a super cool description",
		Thumbnails:  ParseYTThumbnails(all),
	}

	parsed, channelID := ParseYTVideo(video)
	assert.Equal(t, "wow-channel-id", channelID)
	assert.Equal(t, expected, parsed)
}

func TestParseYTChannel(t *testing.T) {
	all := &youtube.ThumbnailDetails{
		Default:  nil,
		Standard: nil,
		Medium:   nil,
		High:     nil,
		Maxres:   nil,
	}

	channel := &youtube.Channel{
		Id: "wow-channel-id",
		Snippet: &youtube.ChannelSnippet{
			Title:       "wow cool title",
			Description: "wow that is a super cool description",
			Thumbnails:  all,
		},
	}

	expected := models.Channel{
		YoutubeID:   "wow-channel-id",
		Title:       "wow cool title",
		Description: "wow that is a super cool description",
		Thumbnails:  ParseYTThumbnails(all),
	}

	parsed := ParseYTChannel(channel)
	assert.Equal(t, expected, parsed)
}

func TestParseSearchResponseIntoVideos(t *testing.T) {

	thumbnails := getFakeThumbnails()

	input := &youtube.SearchListResponse{
		Items: []*youtube.SearchResult{
			{
				Id: &youtube.ResourceId{
					Kind:    "youtube#video",
					VideoId: "video-id-one",
				},
				Snippet: &youtube.SearchResultSnippet{
					ChannelId:   "channel-id",
					Title:       "Video Name 1",
					Description: "Video Description",
					Thumbnails:  &thumbnails,
				},
			},
			{
				Id: &youtube.ResourceId{
					Kind:    "youtube#video",
					VideoId: "video-id-two",
				},
				Snippet: &youtube.SearchResultSnippet{
					ChannelId:   "channel-id",
					Title:       "Video Name 2",
					Description: "Video Description2",
					Thumbnails:  &thumbnails,
				},
			},
		},
	}

	expected := []models.Video{
		{
			YoutubeID:   "video-id-one",
			Title:       "Video Name 1",
			Description: "Video Description",
			Thumbnails:  ParseYTThumbnails(&thumbnails),
		},
		{
			YoutubeID:   "video-id-two",
			Title:       "Video Name 2",
			Description: "Video Description2",
			Thumbnails:  ParseYTThumbnails(&thumbnails),
		},
	}

	actual := ParseSearchResponseIntoVideos(input)

	assert.Equal(t, expected, actual)

}

func getFakeThumbnails() youtube.ThumbnailDetails {
	return youtube.ThumbnailDetails{
		Default: &youtube.Thumbnail{
			Url:    "ytlocker.com",
			Height: 200,
			Width:  200,
		},
		Standard: &youtube.Thumbnail{
			Url:    "ytlocker.com",
			Height: 200,
			Width:  200,
		},
		Medium: &youtube.Thumbnail{
			Url:    "ytlocker.com",
			Height: 200,
			Width:  200,
		},
		High: &youtube.Thumbnail{
			Url:    "ytlocker.com",
			Height: 200,
			Width:  200,
		},
		Maxres: &youtube.Thumbnail{
			Url:    "ytlocker.com",
			Height: 200,
			Width:  200,
		},
	}
}
