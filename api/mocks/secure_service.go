// Code generated by mockery v2.52.1. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// SecureService is an autogenerated mock type for the SecureService type
type SecureService struct {
	mock.Mock
}

// CheckPassword provides a mock function with given fields: ctx, hashedPassword, password
func (_m *SecureService) CheckPassword(ctx context.Context, hashedPassword string, password string) error {
	ret := _m.Called(ctx, hashedPassword, password)

	if len(ret) == 0 {
		panic("no return value specified for CheckPassword")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, hashedPassword, password)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreatePassword provides a mock function with given fields: ctx
func (_m *SecureService) CreatePassword(ctx context.Context) (string, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for CreatePassword")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (string, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) string); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HashPassword provides a mock function with given fields: ctx, password
func (_m *SecureService) HashPassword(ctx context.Context, password string) (string, error) {
	ret := _m.Called(ctx, password)

	if len(ret) == 0 {
		panic("no return value specified for HashPassword")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (string, error)); ok {
		return rf(ctx, password)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) string); ok {
		r0 = rf(ctx, password)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewSecureService creates a new instance of SecureService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewSecureService(t interface {
	mock.TestingT
	Cleanup(func())
}) *SecureService {
	mock := &SecureService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
