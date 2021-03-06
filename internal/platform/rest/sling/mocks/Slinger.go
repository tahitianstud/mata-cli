// Code generated by mockery v1.0.0. DO NOT EDIT.
package mocks

import http "net/http"
import mock "github.com/stretchr/testify/mock"

import sling "github.com/dghubble/sling"

// Slinger is an autogenerated mock type for the Slinger type
type Slinger struct {
	mock.Mock
}

// BodyJSON provides a mock function with given fields: bodyJSON
func (_m *Slinger) BodyJSON(bodyJSON interface{}) *sling.Sling {
	ret := _m.Called(bodyJSON)

	var r0 *sling.Sling
	if rf, ok := ret.Get(0).(func(interface{}) *sling.Sling); ok {
		r0 = rf(bodyJSON)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sling.Sling)
		}
	}

	return r0
}

// Get provides a mock function with given fields: pathURL
func (_m *Slinger) Get(pathURL string) *sling.Sling {
	ret := _m.Called(pathURL)

	var r0 *sling.Sling
	if rf, ok := ret.Get(0).(func(string) *sling.Sling); ok {
		r0 = rf(pathURL)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sling.Sling)
		}
	}

	return r0
}

// Post provides a mock function with given fields: pathURL
func (_m *Slinger) Post(pathURL string) *sling.Sling {
	ret := _m.Called(pathURL)

	var r0 *sling.Sling
	if rf, ok := ret.Get(0).(func(string) *sling.Sling); ok {
		r0 = rf(pathURL)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sling.Sling)
		}
	}

	return r0
}

// ReceiveSuccess provides a mock function with given fields: _a0
func (_m *Slinger) ReceiveSuccess(_a0 interface{}) (*http.Response, error) {
	ret := _m.Called(_a0)

	var r0 *http.Response
	if rf, ok := ret.Get(0).(func(interface{}) *http.Response); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*http.Response)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(interface{}) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Set provides a mock function with given fields: key, value
func (_m *Slinger) Set(key string, value string) *sling.Sling {
	ret := _m.Called(key, value)

	var r0 *sling.Sling
	if rf, ok := ret.Get(0).(func(string, string) *sling.Sling); ok {
		r0 = rf(key, value)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sling.Sling)
		}
	}

	return r0
}
