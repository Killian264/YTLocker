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

	ChannelsAreEqualTest(t, service, channel)
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

func Test_GetChannelIdFromUsername(t *testing.T) {
	service := createMockServices(t)

	id, err := service.GetChannelIdFromUsername("any-id-works-here")
	assert.Nil(t, err)
	assert.NotEmpty(t, id)
}

func Test_GetChannelIdFromUsername_Invalid_Id(t *testing.T) {
	service := createMockServices(t)

	id, err := service.GetChannelIdFromUsername("fake-channel-username")
	assert.NotNil(t, err)
	assert.Empty(t, id)
}

func Test_New_Video_Valid_Video(t *testing.T) {
	service := createMockServices(t)

	channel, err := service.NewChannel("valid-id")

	video, err := service.NewVideo(channel, "valid-id")
	assert.NotNil(t, video)
	assert.Nil(t, err)

	VideosAreEqualTest(t, service, video)
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

func Test_Channel_Get_Videos(t *testing.T) {
	service := createMockServices(t)

	channel, _ := service.NewChannel("valid-id")

	videos, err := service.GetAllChannelVideos(channel)
	assert.Equal(t, 0, len(videos))
	assert.Nil(t, err)

	service.NewVideo(channel, "valid-id")

	videos, err = service.GetAllChannelVideos(channel)
	assert.Equal(t, 1, len(videos))
	assert.Nil(t, err)
}


// New Video Wrong Channel is not tested
func VideosAreEqualTest(t *testing.T, s *YoutubeManager, video models.Video) {
	saved, err := s.GetVideo(video.ID)
	assert.Nil(t, err)

	saved2, err := s.GetVideoByID(video.YoutubeID)
	assert.Nil(t, err)

	thumbnails, err := s.GetAllVideoThumbnails(video)
	assert.Nil(t, err)

	assert.Equal(t, len(video.Thumbnails), len(thumbnails))

	// Encoding decoding to database loses some information for datetimes
	video.CreatedAt = saved.CreatedAt
	video.UpdatedAt = saved.UpdatedAt
	video.Thumbnails = saved.Thumbnails

	assert.Equal(t, video, saved)
	assert.Equal(t, saved, saved2)
}

func ChannelsAreEqualTest(t *testing.T, s *YoutubeManager, channel models.Channel) {
	saved, err := s.GetChannel(channel.ID)
	assert.Nil(t, err)

	saved2, err := s.GetChannelByID(channel.YoutubeID)
	assert.Nil(t, err)

	thumbnails, err := s.GetAllChannelThumbnails(channel)
	assert.Nil(t, err)

	videos, err := s.GetAllChannelVideos(channel)
	assert.Nil(t, err)

	assert.Equal(t, len(channel.Thumbnails), len(thumbnails))
	assert.Equal(t, len(channel.Videos), len(videos))

	// Encoding decoding to database loses some information for datetimes
	channel.CreatedAt = saved.CreatedAt
	channel.UpdatedAt = saved.UpdatedAt
	channel.Thumbnails = saved.Thumbnails

	assert.Equal(t, channel, saved)
	assert.Equal(t, channel, saved2)
}

func createMockServices(t *testing.T) *YoutubeManager {
	data := data.InMemorySQLiteConnect()

	return FakeNewYoutubeManager(data)
}
