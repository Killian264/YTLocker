// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	models "github.com/Killian264/YTLocker/golocker/models"
	mock "github.com/stretchr/testify/mock"
)

// IYoutubePlaylistHelper is an autogenerated mock type for the IYoutubePlaylistHelper type
type IYoutubePlaylistHelper struct {
	mock.Mock
}

// Create provides a mock function with given fields: playlist
func (_m *IYoutubePlaylistHelper) Create(playlist models.Playlist) (models.Playlist, error) {
	ret := _m.Called(playlist)

	var r0 models.Playlist
	if rf, ok := ret.Get(0).(func(models.Playlist) models.Playlist); ok {
		r0 = rf(playlist)
	} else {
		r0 = ret.Get(0).(models.Playlist)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(models.Playlist) error); ok {
		r1 = rf(playlist)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Initalize provides a mock function with given fields: configData, tokenData
func (_m *IYoutubePlaylistHelper) Initalize(configData models.YoutubeClientConfig, tokenData models.YoutubeToken) {
	_m.Called(configData, tokenData)
}

// Insert provides a mock function with given fields: playlist, video
func (_m *IYoutubePlaylistHelper) Insert(playlist models.Playlist, video models.Video) error {
	ret := _m.Called(playlist, video)

	var r0 error
	if rf, ok := ret.Get(0).(func(models.Playlist, models.Video) error); ok {
		r0 = rf(playlist, video)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}