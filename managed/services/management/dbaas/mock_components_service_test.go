// Code generated by mockery v1.0.0. DO NOT EDIT.

package dbaas

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	dbaasv1beta1 "github.com/percona/pmm/api/managementpb/dbaas"
)

// mockComponentsService is an autogenerated mock type for the ComponentsService type
type mockComponentsService struct {
	mock.Mock
}

// ChangePSMDBComponents provides a mock function with given fields: _a0, _a1
func (_m *mockComponentsService) ChangePSMDBComponents(_a0 context.Context, _a1 *dbaasv1beta1.ChangePSMDBComponentsRequest) (*dbaasv1beta1.ChangePSMDBComponentsResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *dbaasv1beta1.ChangePSMDBComponentsResponse
	if rf, ok := ret.Get(0).(func(context.Context, *dbaasv1beta1.ChangePSMDBComponentsRequest) *dbaasv1beta1.ChangePSMDBComponentsResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dbaasv1beta1.ChangePSMDBComponentsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *dbaasv1beta1.ChangePSMDBComponentsRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ChangePXCComponents provides a mock function with given fields: _a0, _a1
func (_m *mockComponentsService) ChangePXCComponents(_a0 context.Context, _a1 *dbaasv1beta1.ChangePXCComponentsRequest) (*dbaasv1beta1.ChangePXCComponentsResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *dbaasv1beta1.ChangePXCComponentsResponse
	if rf, ok := ret.Get(0).(func(context.Context, *dbaasv1beta1.ChangePXCComponentsRequest) *dbaasv1beta1.ChangePXCComponentsResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dbaasv1beta1.ChangePXCComponentsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *dbaasv1beta1.ChangePXCComponentsRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CheckForOperatorUpdate provides a mock function with given fields: _a0, _a1
func (_m *mockComponentsService) CheckForOperatorUpdate(_a0 context.Context, _a1 *dbaasv1beta1.CheckForOperatorUpdateRequest) (*dbaasv1beta1.CheckForOperatorUpdateResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *dbaasv1beta1.CheckForOperatorUpdateResponse
	if rf, ok := ret.Get(0).(func(context.Context, *dbaasv1beta1.CheckForOperatorUpdateRequest) *dbaasv1beta1.CheckForOperatorUpdateResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dbaasv1beta1.CheckForOperatorUpdateResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *dbaasv1beta1.CheckForOperatorUpdateRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPSMDBComponents provides a mock function with given fields: _a0, _a1
func (_m *mockComponentsService) GetPSMDBComponents(_a0 context.Context, _a1 *dbaasv1beta1.GetPSMDBComponentsRequest) (*dbaasv1beta1.GetPSMDBComponentsResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *dbaasv1beta1.GetPSMDBComponentsResponse
	if rf, ok := ret.Get(0).(func(context.Context, *dbaasv1beta1.GetPSMDBComponentsRequest) *dbaasv1beta1.GetPSMDBComponentsResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dbaasv1beta1.GetPSMDBComponentsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *dbaasv1beta1.GetPSMDBComponentsRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPXCComponents provides a mock function with given fields: _a0, _a1
func (_m *mockComponentsService) GetPXCComponents(_a0 context.Context, _a1 *dbaasv1beta1.GetPXCComponentsRequest) (*dbaasv1beta1.GetPXCComponentsResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *dbaasv1beta1.GetPXCComponentsResponse
	if rf, ok := ret.Get(0).(func(context.Context, *dbaasv1beta1.GetPXCComponentsRequest) *dbaasv1beta1.GetPXCComponentsResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dbaasv1beta1.GetPXCComponentsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *dbaasv1beta1.GetPXCComponentsRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InstallOperator provides a mock function with given fields: _a0, _a1
func (_m *mockComponentsService) InstallOperator(_a0 context.Context, _a1 *dbaasv1beta1.InstallOperatorRequest) (*dbaasv1beta1.InstallOperatorResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *dbaasv1beta1.InstallOperatorResponse
	if rf, ok := ret.Get(0).(func(context.Context, *dbaasv1beta1.InstallOperatorRequest) *dbaasv1beta1.InstallOperatorResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dbaasv1beta1.InstallOperatorResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *dbaasv1beta1.InstallOperatorRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
