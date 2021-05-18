package ytmanager

import (
	"log"
	"os"
	"testing"

	"github.com/Killian264/YTLocker/golocker/data"
	"github.com/Killian264/YTLocker/golocker/models"
	"github.com/stretchr/testify/assert"
)

func Test_New_Channel_Valid_Channel(t *testing.T) {
	service := createMockServices(t)

	channel, err := service.NewChannel("valid-id")
	assert.NotNil(t, channel)
	assert.Nil(t, err)

	saved, err := service.GetChannel(channel.ID)
	assert.NotNil(t, saved)
	assert.Nil(t, err)

	saved2, err := service.GetChannelByID(channel.YoutubeID)
	assert.NotNil(t, saved)
	assert.Nil(t, err)

	ChannelsAreEqualTest(t, channel, saved)
	ChannelsAreEqualTest(t, channel, saved2)
}

func Test_New_Channel_InValid_Channel(t *testing.T) {
	service := createMockServices(t)

	channel, err := service.NewChannel("fake-channel-id")
	assert.Equal(t, models.Channel{}, channel)
	assert.NotNil(t, err)
}

func Test_New_Channel_Ignore_Duplicates(t *testing.T) {
	service := createMockServices(t)

	channel, err := service.NewChannel("valid-id")

	channel2, err := service.NewChannel("valid-id")
	assert.Nil(t, err)

	assert.Equal(t, channel.ID, channel2.ID)
}

func Test_New_Video_Valid_Video(t *testing.T) {
	service := createMockServices(t)

	channel, err := service.NewChannel("valid-id")

	video, err := service.NewVideo(channel, "valid-id")
	assert.NotNil(t, video)
	assert.Nil(t, err)

	saved, err := service.GetVideo(video.ID)
	assert.NotNil(t, saved)
	assert.Nil(t, err)

	saved2, err := service.GetVideoByID(video.YoutubeID)
	assert.NotNil(t, saved)
	assert.Nil(t, err)

	VideosAreEqualTest(t, video, saved)
	VideosAreEqualTest(t, video, saved2)
}

func Test_New_Video_InValid_Channel(t *testing.T) {
	service := createMockServices(t)

	channel, err := service.NewChannel("valid-id")

	video, err := service.NewVideo(channel, "fake-video-id")
	assert.NotNil(t, err)
	assert.Equal(t, models.Video{}, video)
}

func Test_New_Video_Ignore_Duplicates(t *testing.T) {
	service := createMockServices(t)

	channel, err := service.NewChannel("valid-id")

	video, err := service.NewVideo(channel, "valid-id")

	video2, err := service.NewVideo(channel, "valid-id")
	assert.Nil(t, err)

	assert.Equal(t, video.ID, video2.ID)
}

func Test_Get_All_Videos(t *testing.T) {
	service := createMockServices(t)

	channel, _ := service.NewChannel("valid-id")

	expected := []models.Video{}

	video, _ := service.NewVideo(channel, "valid-id")
	expected = append(expected, video)

	video, _ = service.NewVideo(channel, "valid-id2")
	expected = append(expected, video)

	actual, err := service.GetAllVideosFromLast24Hours()
	assert.Nil(t, err)

	assert.Equal(t, len(actual), len(expected))
}

func Test_CheckForMissedUploads(t *testing.T) {
	service := createMockServices(t)
	logger := log.New(os.Stdout, "Cron: ", log.Lshortfile)

	channel, _ := service.NewChannel("valid-id")
	service.NewVideo(channel, "video-id-one") // valid id one is specified in ytservicefake getlastvideosfromchannel

	err := service.CheckForMissedUploads(logger)
	assert.Nil(t, err)

	saved, err := service.GetVideoByID("video-id-two")
	assert.NotNil(t, saved)
	assert.Nil(t, err)
}

// New Video Wrong Channel is not tested

func VideosAreEqualTest(t *testing.T, video1 models.Video, video2 models.Video) {
	assert.Equal(t, len(video1.Thumbnails), len(video2.Thumbnails))

	// Encoding decoding to database loses some information for datetimes
	video1.CreatedAt = video2.CreatedAt
	video1.UpdatedAt = video2.UpdatedAt
	video1.Thumbnails = video2.Thumbnails

	assert.Equal(t, video1, video2)
}

func ChannelsAreEqualTest(t *testing.T, channel1 models.Channel, channel2 models.Channel) {
	assert.Equal(t, len(channel1.Thumbnails), len(channel2.Thumbnails))

	// Encoding decoding to database loses some information for datetimes
	channel1.CreatedAt = channel2.CreatedAt
	channel1.UpdatedAt = channel2.UpdatedAt
	channel1.Thumbnails = channel2.Thumbnails

	assert.Equal(t, channel1, channel2)
}

func createMockServices(t *testing.T) *YoutubeManager {
	data := data.InMemorySQLiteConnect()

	return FakeNewYoutubeManager(data)
}
