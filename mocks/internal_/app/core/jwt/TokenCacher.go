// Code generated by mockery v2.45.0. DO NOT EDIT.

package mocks

import (
	domain "FaisalBudiono/go-jwt/internal/app/domain"
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// TokenCacher is an autogenerated mock type for the TokenCacher type
type TokenCacher struct {
	mock.Mock
}

type TokenCacher_Expecter struct {
	mock *mock.Mock
}

func (_m *TokenCacher) EXPECT() *TokenCacher_Expecter {
	return &TokenCacher_Expecter{mock: &_m.Mock}
}

// Cache provides a mock function with given fields: ctx, t
func (_m *TokenCacher) Cache(ctx context.Context, t domain.Token) error {
	ret := _m.Called(ctx, t)

	if len(ret) == 0 {
		panic("no return value specified for Cache")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.Token) error); ok {
		r0 = rf(ctx, t)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// TokenCacher_Cache_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Cache'
type TokenCacher_Cache_Call struct {
	*mock.Call
}

// Cache is a helper method to define mock.On call
//   - ctx context.Context
//   - t domain.Token
func (_e *TokenCacher_Expecter) Cache(ctx interface{}, t interface{}) *TokenCacher_Cache_Call {
	return &TokenCacher_Cache_Call{Call: _e.mock.On("Cache", ctx, t)}
}

func (_c *TokenCacher_Cache_Call) Run(run func(ctx context.Context, t domain.Token)) *TokenCacher_Cache_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(domain.Token))
	})
	return _c
}

func (_c *TokenCacher_Cache_Call) Return(_a0 error) *TokenCacher_Cache_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *TokenCacher_Cache_Call) RunAndReturn(run func(context.Context, domain.Token) error) *TokenCacher_Cache_Call {
	_c.Call.Return(run)
	return _c
}

// NewTokenCacher creates a new instance of TokenCacher. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTokenCacher(t interface {
	mock.TestingT
	Cleanup(func())
}) *TokenCacher {
	mock := &TokenCacher{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
