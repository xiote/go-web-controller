// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// WebDriver is an autogenerated mock type for the WebDriver type
type WebDriver struct {
	mock.Mock
}

// Get provides a mock function with given fields: url
func (_m *WebDriver) Get(url string) error {
	ret := _m.Called(url)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(url)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}