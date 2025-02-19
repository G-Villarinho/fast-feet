// Code generated by mockery v2.52.1. DO NOT EDIT.

package mocks

import (
	context "context"

	models "github.com/G-Villarinho/fast-feet-api/models"
	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// RecipientRepository is an autogenerated mock type for the RecipientRepository type
type RecipientRepository struct {
	mock.Mock
}

// CreateRecipient provides a mock function with given fields: ctx, recipient
func (_m *RecipientRepository) CreateRecipient(ctx context.Context, recipient models.Recipient) error {
	ret := _m.Called(ctx, recipient)

	if len(ret) == 0 {
		panic("no return value specified for CreateRecipient")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, models.Recipient) error); ok {
		r0 = rf(ctx, recipient)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteRecipient provides a mock function with given fields: ctx, ID
func (_m *RecipientRepository) DeleteRecipient(ctx context.Context, ID uuid.UUID) error {
	ret := _m.Called(ctx, ID)

	if len(ret) == 0 {
		panic("no return value specified for DeleteRecipient")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) error); ok {
		r0 = rf(ctx, ID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetRecipientByEmail provides a mock function with given fields: ctx, email
func (_m *RecipientRepository) GetRecipientByEmail(ctx context.Context, email string) (*models.Recipient, error) {
	ret := _m.Called(ctx, email)

	if len(ret) == 0 {
		panic("no return value specified for GetRecipientByEmail")
	}

	var r0 *models.Recipient
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*models.Recipient, error)); ok {
		return rf(ctx, email)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *models.Recipient); ok {
		r0 = rf(ctx, email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Recipient)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRecipientByID provides a mock function with given fields: ctx, ID
func (_m *RecipientRepository) GetRecipientByID(ctx context.Context, ID uuid.UUID) (*models.Recipient, error) {
	ret := _m.Called(ctx, ID)

	if len(ret) == 0 {
		panic("no return value specified for GetRecipientByID")
	}

	var r0 *models.Recipient
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) (*models.Recipient, error)); ok {
		return rf(ctx, ID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) *models.Recipient); ok {
		r0 = rf(ctx, ID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Recipient)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, ID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRecipientLitePagedList provides a mock function with given fields: ctx, pagination
func (_m *RecipientRepository) GetRecipientLitePagedList(ctx context.Context, pagination *models.RecipientBasicInfoPagination) (*models.PaginatedResponse[models.Recipient], error) {
	ret := _m.Called(ctx, pagination)

	if len(ret) == 0 {
		panic("no return value specified for GetRecipientLitePagedList")
	}

	var r0 *models.PaginatedResponse[models.Recipient]
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.RecipientBasicInfoPagination) (*models.PaginatedResponse[models.Recipient], error)); ok {
		return rf(ctx, pagination)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *models.RecipientBasicInfoPagination) *models.PaginatedResponse[models.Recipient]); ok {
		r0 = rf(ctx, pagination)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.PaginatedResponse[models.Recipient])
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *models.RecipientBasicInfoPagination) error); ok {
		r1 = rf(ctx, pagination)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateRecipient provides a mock function with given fields: ctx, recipient
func (_m *RecipientRepository) UpdateRecipient(ctx context.Context, recipient models.Recipient) error {
	ret := _m.Called(ctx, recipient)

	if len(ret) == 0 {
		panic("no return value specified for UpdateRecipient")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, models.Recipient) error); ok {
		r0 = rf(ctx, recipient)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewRecipientRepository creates a new instance of RecipientRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRecipientRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *RecipientRepository {
	mock := &RecipientRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
