// Code generated by mockery v2.39.2. DO NOT EDIT.

package mocks

import (
	context "context"

	migrationhub "github.com/aws/aws-sdk-go-v2/service/migrationhub"
	mock "github.com/stretchr/testify/mock"
)

// IClient is an autogenerated mock type for the IClient type
type IClient struct {
	mock.Mock
}

// AssociateCreatedArtifact provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) AssociateCreatedArtifact(ctx context.Context, params *migrationhub.AssociateCreatedArtifactInput, optFns ...func(*migrationhub.Options)) (*migrationhub.AssociateCreatedArtifactOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for AssociateCreatedArtifact")
	}

	var r0 *migrationhub.AssociateCreatedArtifactOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *migrationhub.AssociateCreatedArtifactInput, ...func(*migrationhub.Options)) (*migrationhub.AssociateCreatedArtifactOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *migrationhub.AssociateCreatedArtifactInput, ...func(*migrationhub.Options)) *migrationhub.AssociateCreatedArtifactOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*migrationhub.AssociateCreatedArtifactOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *migrationhub.AssociateCreatedArtifactInput, ...func(*migrationhub.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AssociateDiscoveredResource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) AssociateDiscoveredResource(ctx context.Context, params *migrationhub.AssociateDiscoveredResourceInput, optFns ...func(*migrationhub.Options)) (*migrationhub.AssociateDiscoveredResourceOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for AssociateDiscoveredResource")
	}

	var r0 *migrationhub.AssociateDiscoveredResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *migrationhub.AssociateDiscoveredResourceInput, ...func(*migrationhub.Options)) (*migrationhub.AssociateDiscoveredResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *migrationhub.AssociateDiscoveredResourceInput, ...func(*migrationhub.Options)) *migrationhub.AssociateDiscoveredResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*migrationhub.AssociateDiscoveredResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *migrationhub.AssociateDiscoveredResourceInput, ...func(*migrationhub.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateProgressUpdateStream provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) CreateProgressUpdateStream(ctx context.Context, params *migrationhub.CreateProgressUpdateStreamInput, optFns ...func(*migrationhub.Options)) (*migrationhub.CreateProgressUpdateStreamOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateProgressUpdateStream")
	}

	var r0 *migrationhub.CreateProgressUpdateStreamOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *migrationhub.CreateProgressUpdateStreamInput, ...func(*migrationhub.Options)) (*migrationhub.CreateProgressUpdateStreamOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *migrationhub.CreateProgressUpdateStreamInput, ...func(*migrationhub.Options)) *migrationhub.CreateProgressUpdateStreamOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*migrationhub.CreateProgressUpdateStreamOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *migrationhub.CreateProgressUpdateStreamInput, ...func(*migrationhub.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteProgressUpdateStream provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DeleteProgressUpdateStream(ctx context.Context, params *migrationhub.DeleteProgressUpdateStreamInput, optFns ...func(*migrationhub.Options)) (*migrationhub.DeleteProgressUpdateStreamOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteProgressUpdateStream")
	}

	var r0 *migrationhub.DeleteProgressUpdateStreamOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *migrationhub.DeleteProgressUpdateStreamInput, ...func(*migrationhub.Options)) (*migrationhub.DeleteProgressUpdateStreamOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *migrationhub.DeleteProgressUpdateStreamInput, ...func(*migrationhub.Options)) *migrationhub.DeleteProgressUpdateStreamOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*migrationhub.DeleteProgressUpdateStreamOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *migrationhub.DeleteProgressUpdateStreamInput, ...func(*migrationhub.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeApplicationState provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DescribeApplicationState(ctx context.Context, params *migrationhub.DescribeApplicationStateInput, optFns ...func(*migrationhub.Options)) (*migrationhub.DescribeApplicationStateOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DescribeApplicationState")
	}

	var r0 *migrationhub.DescribeApplicationStateOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *migrationhub.DescribeApplicationStateInput, ...func(*migrationhub.Options)) (*migrationhub.DescribeApplicationStateOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *migrationhub.DescribeApplicationStateInput, ...func(*migrationhub.Options)) *migrationhub.DescribeApplicationStateOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*migrationhub.DescribeApplicationStateOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *migrationhub.DescribeApplicationStateInput, ...func(*migrationhub.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeMigrationTask provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DescribeMigrationTask(ctx context.Context, params *migrationhub.DescribeMigrationTaskInput, optFns ...func(*migrationhub.Options)) (*migrationhub.DescribeMigrationTaskOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DescribeMigrationTask")
	}

	var r0 *migrationhub.DescribeMigrationTaskOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *migrationhub.DescribeMigrationTaskInput, ...func(*migrationhub.Options)) (*migrationhub.DescribeMigrationTaskOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *migrationhub.DescribeMigrationTaskInput, ...func(*migrationhub.Options)) *migrationhub.DescribeMigrationTaskOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*migrationhub.DescribeMigrationTaskOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *migrationhub.DescribeMigrationTaskInput, ...func(*migrationhub.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DisassociateCreatedArtifact provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DisassociateCreatedArtifact(ctx context.Context, params *migrationhub.DisassociateCreatedArtifactInput, optFns ...func(*migrationhub.Options)) (*migrationhub.DisassociateCreatedArtifactOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DisassociateCreatedArtifact")
	}

	var r0 *migrationhub.DisassociateCreatedArtifactOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *migrationhub.DisassociateCreatedArtifactInput, ...func(*migrationhub.Options)) (*migrationhub.DisassociateCreatedArtifactOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *migrationhub.DisassociateCreatedArtifactInput, ...func(*migrationhub.Options)) *migrationhub.DisassociateCreatedArtifactOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*migrationhub.DisassociateCreatedArtifactOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *migrationhub.DisassociateCreatedArtifactInput, ...func(*migrationhub.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DisassociateDiscoveredResource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DisassociateDiscoveredResource(ctx context.Context, params *migrationhub.DisassociateDiscoveredResourceInput, optFns ...func(*migrationhub.Options)) (*migrationhub.DisassociateDiscoveredResourceOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DisassociateDiscoveredResource")
	}

	var r0 *migrationhub.DisassociateDiscoveredResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *migrationhub.DisassociateDiscoveredResourceInput, ...func(*migrationhub.Options)) (*migrationhub.DisassociateDiscoveredResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *migrationhub.DisassociateDiscoveredResourceInput, ...func(*migrationhub.Options)) *migrationhub.DisassociateDiscoveredResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*migrationhub.DisassociateDiscoveredResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *migrationhub.DisassociateDiscoveredResourceInput, ...func(*migrationhub.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ImportMigrationTask provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ImportMigrationTask(ctx context.Context, params *migrationhub.ImportMigrationTaskInput, optFns ...func(*migrationhub.Options)) (*migrationhub.ImportMigrationTaskOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ImportMigrationTask")
	}

	var r0 *migrationhub.ImportMigrationTaskOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *migrationhub.ImportMigrationTaskInput, ...func(*migrationhub.Options)) (*migrationhub.ImportMigrationTaskOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *migrationhub.ImportMigrationTaskInput, ...func(*migrationhub.Options)) *migrationhub.ImportMigrationTaskOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*migrationhub.ImportMigrationTaskOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *migrationhub.ImportMigrationTaskInput, ...func(*migrationhub.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListApplicationStates provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListApplicationStates(ctx context.Context, params *migrationhub.ListApplicationStatesInput, optFns ...func(*migrationhub.Options)) (*migrationhub.ListApplicationStatesOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListApplicationStates")
	}

	var r0 *migrationhub.ListApplicationStatesOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *migrationhub.ListApplicationStatesInput, ...func(*migrationhub.Options)) (*migrationhub.ListApplicationStatesOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *migrationhub.ListApplicationStatesInput, ...func(*migrationhub.Options)) *migrationhub.ListApplicationStatesOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*migrationhub.ListApplicationStatesOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *migrationhub.ListApplicationStatesInput, ...func(*migrationhub.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListCreatedArtifacts provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListCreatedArtifacts(ctx context.Context, params *migrationhub.ListCreatedArtifactsInput, optFns ...func(*migrationhub.Options)) (*migrationhub.ListCreatedArtifactsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListCreatedArtifacts")
	}

	var r0 *migrationhub.ListCreatedArtifactsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *migrationhub.ListCreatedArtifactsInput, ...func(*migrationhub.Options)) (*migrationhub.ListCreatedArtifactsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *migrationhub.ListCreatedArtifactsInput, ...func(*migrationhub.Options)) *migrationhub.ListCreatedArtifactsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*migrationhub.ListCreatedArtifactsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *migrationhub.ListCreatedArtifactsInput, ...func(*migrationhub.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListDiscoveredResources provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListDiscoveredResources(ctx context.Context, params *migrationhub.ListDiscoveredResourcesInput, optFns ...func(*migrationhub.Options)) (*migrationhub.ListDiscoveredResourcesOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListDiscoveredResources")
	}

	var r0 *migrationhub.ListDiscoveredResourcesOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *migrationhub.ListDiscoveredResourcesInput, ...func(*migrationhub.Options)) (*migrationhub.ListDiscoveredResourcesOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *migrationhub.ListDiscoveredResourcesInput, ...func(*migrationhub.Options)) *migrationhub.ListDiscoveredResourcesOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*migrationhub.ListDiscoveredResourcesOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *migrationhub.ListDiscoveredResourcesInput, ...func(*migrationhub.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListMigrationTasks provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListMigrationTasks(ctx context.Context, params *migrationhub.ListMigrationTasksInput, optFns ...func(*migrationhub.Options)) (*migrationhub.ListMigrationTasksOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListMigrationTasks")
	}

	var r0 *migrationhub.ListMigrationTasksOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *migrationhub.ListMigrationTasksInput, ...func(*migrationhub.Options)) (*migrationhub.ListMigrationTasksOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *migrationhub.ListMigrationTasksInput, ...func(*migrationhub.Options)) *migrationhub.ListMigrationTasksOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*migrationhub.ListMigrationTasksOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *migrationhub.ListMigrationTasksInput, ...func(*migrationhub.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListProgressUpdateStreams provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListProgressUpdateStreams(ctx context.Context, params *migrationhub.ListProgressUpdateStreamsInput, optFns ...func(*migrationhub.Options)) (*migrationhub.ListProgressUpdateStreamsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListProgressUpdateStreams")
	}

	var r0 *migrationhub.ListProgressUpdateStreamsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *migrationhub.ListProgressUpdateStreamsInput, ...func(*migrationhub.Options)) (*migrationhub.ListProgressUpdateStreamsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *migrationhub.ListProgressUpdateStreamsInput, ...func(*migrationhub.Options)) *migrationhub.ListProgressUpdateStreamsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*migrationhub.ListProgressUpdateStreamsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *migrationhub.ListProgressUpdateStreamsInput, ...func(*migrationhub.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NotifyApplicationState provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) NotifyApplicationState(ctx context.Context, params *migrationhub.NotifyApplicationStateInput, optFns ...func(*migrationhub.Options)) (*migrationhub.NotifyApplicationStateOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for NotifyApplicationState")
	}

	var r0 *migrationhub.NotifyApplicationStateOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *migrationhub.NotifyApplicationStateInput, ...func(*migrationhub.Options)) (*migrationhub.NotifyApplicationStateOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *migrationhub.NotifyApplicationStateInput, ...func(*migrationhub.Options)) *migrationhub.NotifyApplicationStateOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*migrationhub.NotifyApplicationStateOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *migrationhub.NotifyApplicationStateInput, ...func(*migrationhub.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NotifyMigrationTaskState provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) NotifyMigrationTaskState(ctx context.Context, params *migrationhub.NotifyMigrationTaskStateInput, optFns ...func(*migrationhub.Options)) (*migrationhub.NotifyMigrationTaskStateOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for NotifyMigrationTaskState")
	}

	var r0 *migrationhub.NotifyMigrationTaskStateOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *migrationhub.NotifyMigrationTaskStateInput, ...func(*migrationhub.Options)) (*migrationhub.NotifyMigrationTaskStateOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *migrationhub.NotifyMigrationTaskStateInput, ...func(*migrationhub.Options)) *migrationhub.NotifyMigrationTaskStateOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*migrationhub.NotifyMigrationTaskStateOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *migrationhub.NotifyMigrationTaskStateInput, ...func(*migrationhub.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Options provides a mock function with given fields:
func (_m *IClient) Options() migrationhub.Options {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Options")
	}

	var r0 migrationhub.Options
	if rf, ok := ret.Get(0).(func() migrationhub.Options); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(migrationhub.Options)
	}

	return r0
}

// PutResourceAttributes provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) PutResourceAttributes(ctx context.Context, params *migrationhub.PutResourceAttributesInput, optFns ...func(*migrationhub.Options)) (*migrationhub.PutResourceAttributesOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for PutResourceAttributes")
	}

	var r0 *migrationhub.PutResourceAttributesOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *migrationhub.PutResourceAttributesInput, ...func(*migrationhub.Options)) (*migrationhub.PutResourceAttributesOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *migrationhub.PutResourceAttributesInput, ...func(*migrationhub.Options)) *migrationhub.PutResourceAttributesOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*migrationhub.PutResourceAttributesOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *migrationhub.PutResourceAttributesInput, ...func(*migrationhub.Options)) error); ok {
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
