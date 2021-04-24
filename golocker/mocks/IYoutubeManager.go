// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	models "github.com/Killian264/YTLocker/golocker/models"
	mock "github.com/stretchr/testify/mock"
)

// IYoutubeManager is an autogenerated mock type for the IYoutubeManager type
type IYoutubeManager struct {
	mock.Mock
}

// GetChannel provides a mock function with given fields: ID
func (_m *IYoutubeManager) GetChannel(ID uint64) (*models.Channel, error) {
	ret := _m.Called(ID)

	var r0 *models.Channel
	if rf, ok := ret.Get(0).(func(uint64) *models.Channel); ok {
		r0 = rf(ID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Channel)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint64) error); ok {
		r1 = rf(ID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetChannelByID provides a mock function with given fields: youtubeID
func (_m *IYoutubeManager) GetChannelByID(youtubeID string) (*models.Channel, error) {
	ret := _m.Called(youtubeID)

	var r0 *models.Channel
	if rf, ok := ret.Get(0).(func(string) *models.Channel); ok {
		r0 = rf(youtubeID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Channel)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(youtubeID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetVideo provides a mock function with given fields: ID
func (_m *IYoutubeManager) GetVideo(ID uint64) (*models.Video, error) {
	ret := _m.Called(ID)

	var r0 *models.Video
	if rf, ok := ret.Get(0).(func(uint64) *models.Video); ok {
		r0 = rf(ID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Video)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint64) error); ok {
		r1 = rf(ID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetVideoByID provides a mock function with given fields: youtubeID
func (_m *IYoutubeManager) GetVideoByID(youtubeID string) (*models.Video, error) {
	ret := _m.Called(youtubeID)

	var r0 *models.Video
	if rf, ok := ret.Get(0).(func(string) *models.Video); ok {
		r0 = rf(youtubeID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Video)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(youtubeID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewChannel provides a mock function with given fields: channelID
func (_m *IYoutubeManager) NewChannel(channelID string) (*models.Channel, error) {
	ret := _m.Called(channelID)

	var r0 *models.Channel
	if rf, ok := ret.Get(0).(func(string) *models.Channel); ok {
		r0 = rf(channelID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Channel)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(channelID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewVideo provides a mock function with given fields: channel, videoID
func (_m *IYoutubeManager) NewVideo(channel *models.Channel, videoID string) (*models.Video, error) {
	ret := _m.Called(channel, videoID)

	var r0 *models.Video
	if rf, ok := ret.Get(0).(func(*models.Channel, string) *models.Video); ok {
		r0 = rf(channel, videoID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Video)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*models.Channel, string) error); ok {
		r1 = rf(channel, videoID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
