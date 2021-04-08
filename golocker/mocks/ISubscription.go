// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	models "github.com/Killian264/YTLocker/golocker/models"
	mock "github.com/stretchr/testify/mock"
)

// ISubscription is an autogenerated mock type for the ISubscription type
type ISubscription struct {
	mock.Mock
}

// CreateSubscription provides a mock function with given fields: channelID
func (_m *ISubscription) CreateSubscription(channelID string) (*models.SubscriptionRequest, error) {
	ret := _m.Called(channelID)

	var r0 *models.SubscriptionRequest
	if rf, ok := ret.Get(0).(func(string) *models.SubscriptionRequest); ok {
		r0 = rf(channelID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.SubscriptionRequest)
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

// HandleChallenge provides a mock function with given fields: request
func (_m *ISubscription) HandleChallenge(request *models.SubscriptionRequest) (bool, error) {
	ret := _m.Called(request)

	var r0 bool
	if rf, ok := ret.Get(0).(func(*models.SubscriptionRequest) bool); ok {
		r0 = rf(request)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*models.SubscriptionRequest) error); ok {
		r1 = rf(request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HandleVideoPush provides a mock function with given fields: push, secret
func (_m *ISubscription) HandleVideoPush(push *models.YTHookPush, secret string) error {
	ret := _m.Called(push, secret)

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.YTHookPush, string) error); ok {
		r0 = rf(push, secret)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ResubscribeAll provides a mock function with given fields:
func (_m *ISubscription) ResubscribeAll() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetSubscribeUrl provides a mock function with given fields: base, path
func (_m *ISubscription) SetSubscribeUrl(base string, path string) {
	_m.Called(base, path)
}

// SetYTPubSubUrl provides a mock function with given fields: url
func (_m *ISubscription) SetYTPubSubUrl(url string) {
	_m.Called(url)
}

// Subscribe provides a mock function with given fields: request
func (_m *ISubscription) Subscribe(request *models.SubscriptionRequest) error {
	ret := _m.Called(request)

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.SubscriptionRequest) error); ok {
		r0 = rf(request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}