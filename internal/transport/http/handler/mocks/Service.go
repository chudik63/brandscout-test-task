// Code generated by mockery v2.53.3. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	models "brandscout-test-task/internal/models"
)

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

// AddQuote provides a mock function with given fields: ctx, quote
func (_m *Service) AddQuote(ctx context.Context, quote *models.Quote) {
	_m.Called(ctx, quote)
}

// DeleteQuote provides a mock function with given fields: ctx, id
func (_m *Service) DeleteQuote(ctx context.Context, id uint64) error {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for DeleteQuote")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllQuotes provides a mock function with given fields: ctx, author
func (_m *Service) GetAllQuotes(ctx context.Context, author string) ([]*models.Quote, error) {
	ret := _m.Called(ctx, author)

	if len(ret) == 0 {
		panic("no return value specified for GetAllQuotes")
	}

	var r0 []*models.Quote
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]*models.Quote, error)); ok {
		return rf(ctx, author)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []*models.Quote); ok {
		r0 = rf(ctx, author)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Quote)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, author)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetQuote provides a mock function with given fields: ctx, id
func (_m *Service) GetQuote(ctx context.Context, id uint64) (*models.Quote, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetQuote")
	}

	var r0 *models.Quote
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64) (*models.Quote, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64) *models.Quote); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Quote)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRandomQuote provides a mock function with given fields: ctx
func (_m *Service) GetRandomQuote(ctx context.Context) (*models.Quote, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetRandomQuote")
	}

	var r0 *models.Quote
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (*models.Quote, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) *models.Quote); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Quote)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewService creates a new instance of Service. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewService(t interface {
	mock.TestingT
	Cleanup(func())
}) *Service {
	mock := &Service{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
