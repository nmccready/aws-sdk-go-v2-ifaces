// Code generated by mockery v2.53.4. DO NOT EDIT.

package mocks

import (
	context "context"

	docdbelastic "github.com/aws/aws-sdk-go-v2/service/docdbelastic"

	mock "github.com/stretchr/testify/mock"
)

// IClient is an autogenerated mock type for the IClient type
type IClient struct {
	mock.Mock
}

// ApplyPendingMaintenanceAction provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ApplyPendingMaintenanceAction(ctx context.Context, params *docdbelastic.ApplyPendingMaintenanceActionInput, optFns ...func(*docdbelastic.Options)) (*docdbelastic.ApplyPendingMaintenanceActionOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ApplyPendingMaintenanceAction")
	}

	var r0 *docdbelastic.ApplyPendingMaintenanceActionOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *docdbelastic.ApplyPendingMaintenanceActionInput, ...func(*docdbelastic.Options)) (*docdbelastic.ApplyPendingMaintenanceActionOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *docdbelastic.ApplyPendingMaintenanceActionInput, ...func(*docdbelastic.Options)) *docdbelastic.ApplyPendingMaintenanceActionOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*docdbelastic.ApplyPendingMaintenanceActionOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *docdbelastic.ApplyPendingMaintenanceActionInput, ...func(*docdbelastic.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CopyClusterSnapshot provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) CopyClusterSnapshot(ctx context.Context, params *docdbelastic.CopyClusterSnapshotInput, optFns ...func(*docdbelastic.Options)) (*docdbelastic.CopyClusterSnapshotOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CopyClusterSnapshot")
	}

	var r0 *docdbelastic.CopyClusterSnapshotOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *docdbelastic.CopyClusterSnapshotInput, ...func(*docdbelastic.Options)) (*docdbelastic.CopyClusterSnapshotOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *docdbelastic.CopyClusterSnapshotInput, ...func(*docdbelastic.Options)) *docdbelastic.CopyClusterSnapshotOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*docdbelastic.CopyClusterSnapshotOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *docdbelastic.CopyClusterSnapshotInput, ...func(*docdbelastic.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateCluster provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) CreateCluster(ctx context.Context, params *docdbelastic.CreateClusterInput, optFns ...func(*docdbelastic.Options)) (*docdbelastic.CreateClusterOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateCluster")
	}

	var r0 *docdbelastic.CreateClusterOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *docdbelastic.CreateClusterInput, ...func(*docdbelastic.Options)) (*docdbelastic.CreateClusterOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *docdbelastic.CreateClusterInput, ...func(*docdbelastic.Options)) *docdbelastic.CreateClusterOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*docdbelastic.CreateClusterOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *docdbelastic.CreateClusterInput, ...func(*docdbelastic.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateClusterSnapshot provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) CreateClusterSnapshot(ctx context.Context, params *docdbelastic.CreateClusterSnapshotInput, optFns ...func(*docdbelastic.Options)) (*docdbelastic.CreateClusterSnapshotOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateClusterSnapshot")
	}

	var r0 *docdbelastic.CreateClusterSnapshotOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *docdbelastic.CreateClusterSnapshotInput, ...func(*docdbelastic.Options)) (*docdbelastic.CreateClusterSnapshotOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *docdbelastic.CreateClusterSnapshotInput, ...func(*docdbelastic.Options)) *docdbelastic.CreateClusterSnapshotOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*docdbelastic.CreateClusterSnapshotOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *docdbelastic.CreateClusterSnapshotInput, ...func(*docdbelastic.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteCluster provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DeleteCluster(ctx context.Context, params *docdbelastic.DeleteClusterInput, optFns ...func(*docdbelastic.Options)) (*docdbelastic.DeleteClusterOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteCluster")
	}

	var r0 *docdbelastic.DeleteClusterOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *docdbelastic.DeleteClusterInput, ...func(*docdbelastic.Options)) (*docdbelastic.DeleteClusterOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *docdbelastic.DeleteClusterInput, ...func(*docdbelastic.Options)) *docdbelastic.DeleteClusterOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*docdbelastic.DeleteClusterOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *docdbelastic.DeleteClusterInput, ...func(*docdbelastic.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteClusterSnapshot provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DeleteClusterSnapshot(ctx context.Context, params *docdbelastic.DeleteClusterSnapshotInput, optFns ...func(*docdbelastic.Options)) (*docdbelastic.DeleteClusterSnapshotOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteClusterSnapshot")
	}

	var r0 *docdbelastic.DeleteClusterSnapshotOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *docdbelastic.DeleteClusterSnapshotInput, ...func(*docdbelastic.Options)) (*docdbelastic.DeleteClusterSnapshotOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *docdbelastic.DeleteClusterSnapshotInput, ...func(*docdbelastic.Options)) *docdbelastic.DeleteClusterSnapshotOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*docdbelastic.DeleteClusterSnapshotOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *docdbelastic.DeleteClusterSnapshotInput, ...func(*docdbelastic.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCluster provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetCluster(ctx context.Context, params *docdbelastic.GetClusterInput, optFns ...func(*docdbelastic.Options)) (*docdbelastic.GetClusterOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetCluster")
	}

	var r0 *docdbelastic.GetClusterOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *docdbelastic.GetClusterInput, ...func(*docdbelastic.Options)) (*docdbelastic.GetClusterOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *docdbelastic.GetClusterInput, ...func(*docdbelastic.Options)) *docdbelastic.GetClusterOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*docdbelastic.GetClusterOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *docdbelastic.GetClusterInput, ...func(*docdbelastic.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetClusterSnapshot provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetClusterSnapshot(ctx context.Context, params *docdbelastic.GetClusterSnapshotInput, optFns ...func(*docdbelastic.Options)) (*docdbelastic.GetClusterSnapshotOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetClusterSnapshot")
	}

	var r0 *docdbelastic.GetClusterSnapshotOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *docdbelastic.GetClusterSnapshotInput, ...func(*docdbelastic.Options)) (*docdbelastic.GetClusterSnapshotOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *docdbelastic.GetClusterSnapshotInput, ...func(*docdbelastic.Options)) *docdbelastic.GetClusterSnapshotOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*docdbelastic.GetClusterSnapshotOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *docdbelastic.GetClusterSnapshotInput, ...func(*docdbelastic.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPendingMaintenanceAction provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetPendingMaintenanceAction(ctx context.Context, params *docdbelastic.GetPendingMaintenanceActionInput, optFns ...func(*docdbelastic.Options)) (*docdbelastic.GetPendingMaintenanceActionOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetPendingMaintenanceAction")
	}

	var r0 *docdbelastic.GetPendingMaintenanceActionOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *docdbelastic.GetPendingMaintenanceActionInput, ...func(*docdbelastic.Options)) (*docdbelastic.GetPendingMaintenanceActionOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *docdbelastic.GetPendingMaintenanceActionInput, ...func(*docdbelastic.Options)) *docdbelastic.GetPendingMaintenanceActionOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*docdbelastic.GetPendingMaintenanceActionOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *docdbelastic.GetPendingMaintenanceActionInput, ...func(*docdbelastic.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListClusterSnapshots provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListClusterSnapshots(ctx context.Context, params *docdbelastic.ListClusterSnapshotsInput, optFns ...func(*docdbelastic.Options)) (*docdbelastic.ListClusterSnapshotsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListClusterSnapshots")
	}

	var r0 *docdbelastic.ListClusterSnapshotsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *docdbelastic.ListClusterSnapshotsInput, ...func(*docdbelastic.Options)) (*docdbelastic.ListClusterSnapshotsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *docdbelastic.ListClusterSnapshotsInput, ...func(*docdbelastic.Options)) *docdbelastic.ListClusterSnapshotsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*docdbelastic.ListClusterSnapshotsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *docdbelastic.ListClusterSnapshotsInput, ...func(*docdbelastic.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListClusters provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListClusters(ctx context.Context, params *docdbelastic.ListClustersInput, optFns ...func(*docdbelastic.Options)) (*docdbelastic.ListClustersOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListClusters")
	}

	var r0 *docdbelastic.ListClustersOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *docdbelastic.ListClustersInput, ...func(*docdbelastic.Options)) (*docdbelastic.ListClustersOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *docdbelastic.ListClustersInput, ...func(*docdbelastic.Options)) *docdbelastic.ListClustersOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*docdbelastic.ListClustersOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *docdbelastic.ListClustersInput, ...func(*docdbelastic.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListPendingMaintenanceActions provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListPendingMaintenanceActions(ctx context.Context, params *docdbelastic.ListPendingMaintenanceActionsInput, optFns ...func(*docdbelastic.Options)) (*docdbelastic.ListPendingMaintenanceActionsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListPendingMaintenanceActions")
	}

	var r0 *docdbelastic.ListPendingMaintenanceActionsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *docdbelastic.ListPendingMaintenanceActionsInput, ...func(*docdbelastic.Options)) (*docdbelastic.ListPendingMaintenanceActionsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *docdbelastic.ListPendingMaintenanceActionsInput, ...func(*docdbelastic.Options)) *docdbelastic.ListPendingMaintenanceActionsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*docdbelastic.ListPendingMaintenanceActionsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *docdbelastic.ListPendingMaintenanceActionsInput, ...func(*docdbelastic.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListTagsForResource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListTagsForResource(ctx context.Context, params *docdbelastic.ListTagsForResourceInput, optFns ...func(*docdbelastic.Options)) (*docdbelastic.ListTagsForResourceOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListTagsForResource")
	}

	var r0 *docdbelastic.ListTagsForResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *docdbelastic.ListTagsForResourceInput, ...func(*docdbelastic.Options)) (*docdbelastic.ListTagsForResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *docdbelastic.ListTagsForResourceInput, ...func(*docdbelastic.Options)) *docdbelastic.ListTagsForResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*docdbelastic.ListTagsForResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *docdbelastic.ListTagsForResourceInput, ...func(*docdbelastic.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Options provides a mock function with no fields
func (_m *IClient) Options() docdbelastic.Options {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Options")
	}

	var r0 docdbelastic.Options
	if rf, ok := ret.Get(0).(func() docdbelastic.Options); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(docdbelastic.Options)
	}

	return r0
}

// RestoreClusterFromSnapshot provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) RestoreClusterFromSnapshot(ctx context.Context, params *docdbelastic.RestoreClusterFromSnapshotInput, optFns ...func(*docdbelastic.Options)) (*docdbelastic.RestoreClusterFromSnapshotOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for RestoreClusterFromSnapshot")
	}

	var r0 *docdbelastic.RestoreClusterFromSnapshotOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *docdbelastic.RestoreClusterFromSnapshotInput, ...func(*docdbelastic.Options)) (*docdbelastic.RestoreClusterFromSnapshotOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *docdbelastic.RestoreClusterFromSnapshotInput, ...func(*docdbelastic.Options)) *docdbelastic.RestoreClusterFromSnapshotOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*docdbelastic.RestoreClusterFromSnapshotOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *docdbelastic.RestoreClusterFromSnapshotInput, ...func(*docdbelastic.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StartCluster provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) StartCluster(ctx context.Context, params *docdbelastic.StartClusterInput, optFns ...func(*docdbelastic.Options)) (*docdbelastic.StartClusterOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for StartCluster")
	}

	var r0 *docdbelastic.StartClusterOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *docdbelastic.StartClusterInput, ...func(*docdbelastic.Options)) (*docdbelastic.StartClusterOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *docdbelastic.StartClusterInput, ...func(*docdbelastic.Options)) *docdbelastic.StartClusterOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*docdbelastic.StartClusterOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *docdbelastic.StartClusterInput, ...func(*docdbelastic.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StopCluster provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) StopCluster(ctx context.Context, params *docdbelastic.StopClusterInput, optFns ...func(*docdbelastic.Options)) (*docdbelastic.StopClusterOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for StopCluster")
	}

	var r0 *docdbelastic.StopClusterOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *docdbelastic.StopClusterInput, ...func(*docdbelastic.Options)) (*docdbelastic.StopClusterOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *docdbelastic.StopClusterInput, ...func(*docdbelastic.Options)) *docdbelastic.StopClusterOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*docdbelastic.StopClusterOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *docdbelastic.StopClusterInput, ...func(*docdbelastic.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TagResource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) TagResource(ctx context.Context, params *docdbelastic.TagResourceInput, optFns ...func(*docdbelastic.Options)) (*docdbelastic.TagResourceOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for TagResource")
	}

	var r0 *docdbelastic.TagResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *docdbelastic.TagResourceInput, ...func(*docdbelastic.Options)) (*docdbelastic.TagResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *docdbelastic.TagResourceInput, ...func(*docdbelastic.Options)) *docdbelastic.TagResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*docdbelastic.TagResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *docdbelastic.TagResourceInput, ...func(*docdbelastic.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UntagResource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) UntagResource(ctx context.Context, params *docdbelastic.UntagResourceInput, optFns ...func(*docdbelastic.Options)) (*docdbelastic.UntagResourceOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UntagResource")
	}

	var r0 *docdbelastic.UntagResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *docdbelastic.UntagResourceInput, ...func(*docdbelastic.Options)) (*docdbelastic.UntagResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *docdbelastic.UntagResourceInput, ...func(*docdbelastic.Options)) *docdbelastic.UntagResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*docdbelastic.UntagResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *docdbelastic.UntagResourceInput, ...func(*docdbelastic.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateCluster provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) UpdateCluster(ctx context.Context, params *docdbelastic.UpdateClusterInput, optFns ...func(*docdbelastic.Options)) (*docdbelastic.UpdateClusterOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UpdateCluster")
	}

	var r0 *docdbelastic.UpdateClusterOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *docdbelastic.UpdateClusterInput, ...func(*docdbelastic.Options)) (*docdbelastic.UpdateClusterOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *docdbelastic.UpdateClusterInput, ...func(*docdbelastic.Options)) *docdbelastic.UpdateClusterOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*docdbelastic.UpdateClusterOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *docdbelastic.UpdateClusterInput, ...func(*docdbelastic.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewIClient creates a new instance of IClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *IClient {
	mock := &IClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
