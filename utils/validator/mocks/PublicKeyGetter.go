// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// PublicKeyGetter is an autogenerated mock type for the PublicKeyGetter type
type PublicKeyGetter struct {
	mock.Mock
}

// GetPublicKey provides a mock function with given fields: keyID
func (_m *PublicKeyGetter) GetPublicKey(keyID string) (interface{}, error) {
	ret := _m.Called(keyID)

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(string) interface{}); ok {
		r0 = rf(keyID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(keyID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
