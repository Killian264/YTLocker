// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// IChannel is an autogenerated mock type for the IChannel type
type IChannel struct {
	mock.Mock
}

// GetOrCreateChannel provides a mock function with given fields: channelID
func (_m *IChannel) GetOrCreateChannel(channelID string) {
	_m.Called(channelID)
}