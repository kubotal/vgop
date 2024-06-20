// Code generated by mockery v2.40.3. DO NOT EDIT.

package lsblk

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	lsblk "vgop/internal/controllers/vgmanager/lsblk"
)

// MockLSBLK is an autogenerated mock type for the LSBLK type
type MockLSBLK struct {
	mock.Mock
}

type MockLSBLK_Expecter struct {
	mock *mock.Mock
}

func (_m *MockLSBLK) EXPECT() *MockLSBLK_Expecter {
	return &MockLSBLK_Expecter{mock: &_m.Mock}
}

// BlockDeviceInfos provides a mock function with given fields: ctx, bs
func (_m *MockLSBLK) BlockDeviceInfos(ctx context.Context, bs []lsblk.BlockDevice) (lsblk.BlockDeviceInfos, error) {
	ret := _m.Called(ctx, bs)

	if len(ret) == 0 {
		panic("no return value specified for BlockDeviceInfos")
	}

	var r0 lsblk.BlockDeviceInfos
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []lsblk.BlockDevice) (lsblk.BlockDeviceInfos, error)); ok {
		return rf(ctx, bs)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []lsblk.BlockDevice) lsblk.BlockDeviceInfos); ok {
		r0 = rf(ctx, bs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(lsblk.BlockDeviceInfos)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []lsblk.BlockDevice) error); ok {
		r1 = rf(ctx, bs)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockLSBLK_BlockDeviceInfos_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'BlockDeviceInfos'
type MockLSBLK_BlockDeviceInfos_Call struct {
	*mock.Call
}

// BlockDeviceInfos is a helper method to define mock.On call
//   - ctx context.Context
//   - bs []lsblk.BlockDevice
func (_e *MockLSBLK_Expecter) BlockDeviceInfos(ctx interface{}, bs interface{}) *MockLSBLK_BlockDeviceInfos_Call {
	return &MockLSBLK_BlockDeviceInfos_Call{Call: _e.mock.On("BlockDeviceInfos", ctx, bs)}
}

func (_c *MockLSBLK_BlockDeviceInfos_Call) Run(run func(ctx context.Context, bs []lsblk.BlockDevice)) *MockLSBLK_BlockDeviceInfos_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]lsblk.BlockDevice))
	})
	return _c
}

func (_c *MockLSBLK_BlockDeviceInfos_Call) Return(_a0 lsblk.BlockDeviceInfos, _a1 error) *MockLSBLK_BlockDeviceInfos_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockLSBLK_BlockDeviceInfos_Call) RunAndReturn(run func(context.Context, []lsblk.BlockDevice) (lsblk.BlockDeviceInfos, error)) *MockLSBLK_BlockDeviceInfos_Call {
	_c.Call.Return(run)
	return _c
}

// IsUsableLoopDev provides a mock function with given fields: ctx, b
func (_m *MockLSBLK) IsUsableLoopDev(ctx context.Context, b lsblk.BlockDevice) (bool, error) {
	ret := _m.Called(ctx, b)

	if len(ret) == 0 {
		panic("no return value specified for IsUsableLoopDev")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, lsblk.BlockDevice) (bool, error)); ok {
		return rf(ctx, b)
	}
	if rf, ok := ret.Get(0).(func(context.Context, lsblk.BlockDevice) bool); ok {
		r0 = rf(ctx, b)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context, lsblk.BlockDevice) error); ok {
		r1 = rf(ctx, b)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockLSBLK_IsUsableLoopDev_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsUsableLoopDev'
type MockLSBLK_IsUsableLoopDev_Call struct {
	*mock.Call
}

// IsUsableLoopDev is a helper method to define mock.On call
//   - ctx context.Context
//   - b lsblk.BlockDevice
func (_e *MockLSBLK_Expecter) IsUsableLoopDev(ctx interface{}, b interface{}) *MockLSBLK_IsUsableLoopDev_Call {
	return &MockLSBLK_IsUsableLoopDev_Call{Call: _e.mock.On("IsUsableLoopDev", ctx, b)}
}

func (_c *MockLSBLK_IsUsableLoopDev_Call) Run(run func(ctx context.Context, b lsblk.BlockDevice)) *MockLSBLK_IsUsableLoopDev_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(lsblk.BlockDevice))
	})
	return _c
}

func (_c *MockLSBLK_IsUsableLoopDev_Call) Return(_a0 bool, _a1 error) *MockLSBLK_IsUsableLoopDev_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockLSBLK_IsUsableLoopDev_Call) RunAndReturn(run func(context.Context, lsblk.BlockDevice) (bool, error)) *MockLSBLK_IsUsableLoopDev_Call {
	_c.Call.Return(run)
	return _c
}

// ListBlockDevices provides a mock function with given fields: ctx
func (_m *MockLSBLK) ListBlockDevices(ctx context.Context) ([]lsblk.BlockDevice, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for ListBlockDevices")
	}

	var r0 []lsblk.BlockDevice
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]lsblk.BlockDevice, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []lsblk.BlockDevice); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]lsblk.BlockDevice)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockLSBLK_ListBlockDevices_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListBlockDevices'
type MockLSBLK_ListBlockDevices_Call struct {
	*mock.Call
}

// ListBlockDevices is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockLSBLK_Expecter) ListBlockDevices(ctx interface{}) *MockLSBLK_ListBlockDevices_Call {
	return &MockLSBLK_ListBlockDevices_Call{Call: _e.mock.On("ListBlockDevices", ctx)}
}

func (_c *MockLSBLK_ListBlockDevices_Call) Run(run func(ctx context.Context)) *MockLSBLK_ListBlockDevices_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockLSBLK_ListBlockDevices_Call) Return(_a0 []lsblk.BlockDevice, _a1 error) *MockLSBLK_ListBlockDevices_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockLSBLK_ListBlockDevices_Call) RunAndReturn(run func(context.Context) ([]lsblk.BlockDevice, error)) *MockLSBLK_ListBlockDevices_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockLSBLK creates a new instance of MockLSBLK. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockLSBLK(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockLSBLK {
	mock := &MockLSBLK{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
