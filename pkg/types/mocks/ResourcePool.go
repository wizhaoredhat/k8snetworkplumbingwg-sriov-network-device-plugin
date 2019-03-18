// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"
import types "github.com/intel/sriov-network-device-plugin/pkg/types"
import v1beta1 "k8s.io/kubernetes/pkg/kubelet/apis/deviceplugin/v1beta1"

// ResourcePool is an autogenerated mock type for the ResourcePool type
type ResourcePool struct {
	mock.Mock
}

// DiscoverDevices provides a mock function with given fields:
func (_m *ResourcePool) DiscoverDevices() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetConfig provides a mock function with given fields:
func (_m *ResourcePool) GetConfig() *types.ResourceConfig {
	ret := _m.Called()

	var r0 *types.ResourceConfig
	if rf, ok := ret.Get(0).(func() *types.ResourceConfig); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.ResourceConfig)
		}
	}

	return r0
}

// GetDeviceFile provides a mock function with given fields: dev
func (_m *ResourcePool) GetDeviceFile(dev string) (string, error) {
	ret := _m.Called(dev)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(dev)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(dev)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDeviceFiles provides a mock function with given fields:
func (_m *ResourcePool) GetDeviceFiles() map[string]string {
	ret := _m.Called()

	var r0 map[string]string
	if rf, ok := ret.Get(0).(func() map[string]string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]string)
		}
	}

	return r0
}

// GetDeviceSpecs provides a mock function with given fields: deviceFiles, deviceIDs
func (_m *ResourcePool) GetDeviceSpecs(deviceFiles map[string]string, deviceIDs []string) []*v1beta1.DeviceSpec {
	ret := _m.Called(deviceFiles, deviceIDs)

	var r0 []*v1beta1.DeviceSpec
	if rf, ok := ret.Get(0).(func(map[string]string, []string) []*v1beta1.DeviceSpec); ok {
		r0 = rf(deviceFiles, deviceIDs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*v1beta1.DeviceSpec)
		}
	}

	return r0
}

// GetDevices provides a mock function with given fields:
func (_m *ResourcePool) GetDevices() map[string]*v1beta1.Device {
	ret := _m.Called()

	var r0 map[string]*v1beta1.Device
	if rf, ok := ret.Get(0).(func() map[string]*v1beta1.Device); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]*v1beta1.Device)
		}
	}

	return r0
}

// GetEnvs provides a mock function with given fields: deviceIDs
func (_m *ResourcePool) GetEnvs(deviceIDs []string) map[string]string {
	ret := _m.Called(deviceIDs)

	var r0 map[string]string
	if rf, ok := ret.Get(0).(func([]string) map[string]string); ok {
		r0 = rf(deviceIDs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]string)
		}
	}

	return r0
}

// GetMounts provides a mock function with given fields:
func (_m *ResourcePool) GetMounts() []*v1beta1.Mount {
	ret := _m.Called()

	var r0 []*v1beta1.Mount
	if rf, ok := ret.Get(0).(func() []*v1beta1.Mount); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*v1beta1.Mount)
		}
	}

	return r0
}

// GetResourceName provides a mock function with given fields:
func (_m *ResourcePool) GetResourceName() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// InitDevice provides a mock function with given fields:
func (_m *ResourcePool) InitDevice() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Probe provides a mock function with given fields: _a0, _a1
func (_m *ResourcePool) Probe(_a0 *types.ResourceConfig, _a1 map[string]*v1beta1.Device) bool {
	ret := _m.Called(_a0, _a1)

	var r0 bool
	if rf, ok := ret.Get(0).(func(*types.ResourceConfig, map[string]*v1beta1.Device) bool); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}