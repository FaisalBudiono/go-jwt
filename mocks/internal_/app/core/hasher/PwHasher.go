// Code generated by mockery v2.45.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// PwHasher is an autogenerated mock type for the PwHasher type
type PwHasher struct {
	mock.Mock
}

type PwHasher_Expecter struct {
	mock *mock.Mock
}

func (_m *PwHasher) EXPECT() *PwHasher_Expecter {
	return &PwHasher_Expecter{mock: &_m.Mock}
}

// Hash provides a mock function with given fields: plain
func (_m *PwHasher) Hash(plain string) (string, error) {
	ret := _m.Called(plain)

	if len(ret) == 0 {
		panic("no return value specified for Hash")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (string, error)); ok {
		return rf(plain)
	}
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(plain)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(plain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PwHasher_Hash_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Hash'
type PwHasher_Hash_Call struct {
	*mock.Call
}

// Hash is a helper method to define mock.On call
//   - plain string
func (_e *PwHasher_Expecter) Hash(plain interface{}) *PwHasher_Hash_Call {
	return &PwHasher_Hash_Call{Call: _e.mock.On("Hash", plain)}
}

func (_c *PwHasher_Hash_Call) Run(run func(plain string)) *PwHasher_Hash_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *PwHasher_Hash_Call) Return(_a0 string, _a1 error) *PwHasher_Hash_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *PwHasher_Hash_Call) RunAndReturn(run func(string) (string, error)) *PwHasher_Hash_Call {
	_c.Call.Return(run)
	return _c
}

// Verify provides a mock function with given fields: plain, hashed
func (_m *PwHasher) Verify(plain string, hashed string) (bool, error) {
	ret := _m.Called(plain, hashed)

	if len(ret) == 0 {
		panic("no return value specified for Verify")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (bool, error)); ok {
		return rf(plain, hashed)
	}
	if rf, ok := ret.Get(0).(func(string, string) bool); ok {
		r0 = rf(plain, hashed)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(plain, hashed)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PwHasher_Verify_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Verify'
type PwHasher_Verify_Call struct {
	*mock.Call
}

// Verify is a helper method to define mock.On call
//   - plain string
//   - hashed string
func (_e *PwHasher_Expecter) Verify(plain interface{}, hashed interface{}) *PwHasher_Verify_Call {
	return &PwHasher_Verify_Call{Call: _e.mock.On("Verify", plain, hashed)}
}

func (_c *PwHasher_Verify_Call) Run(run func(plain string, hashed string)) *PwHasher_Verify_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *PwHasher_Verify_Call) Return(_a0 bool, _a1 error) *PwHasher_Verify_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *PwHasher_Verify_Call) RunAndReturn(run func(string, string) (bool, error)) *PwHasher_Verify_Call {
	_c.Call.Return(run)
	return _c
}

// NewPwHasher creates a new instance of PwHasher. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewPwHasher(t interface {
	mock.TestingT
	Cleanup(func())
}) *PwHasher {
	mock := &PwHasher{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
