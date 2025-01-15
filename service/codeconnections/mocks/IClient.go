// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	codeconnections "github.com/aws/aws-sdk-go-v2/service/codeconnections"

	context "context"

	mock "github.com/stretchr/testify/mock"
)

// IClient is an autogenerated mock type for the IClient type
type IClient struct {
	mock.Mock
}

// CreateConnection provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) CreateConnection(ctx context.Context, params *codeconnections.CreateConnectionInput, optFns ...func(*codeconnections.Options)) (*codeconnections.CreateConnectionOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateConnection")
	}

	var r0 *codeconnections.CreateConnectionOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *codeconnections.CreateConnectionInput, ...func(*codeconnections.Options)) (*codeconnections.CreateConnectionOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *codeconnections.CreateConnectionInput, ...func(*codeconnections.Options)) *codeconnections.CreateConnectionOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*codeconnections.CreateConnectionOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *codeconnections.CreateConnectionInput, ...func(*codeconnections.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateHost provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) CreateHost(ctx context.Context, params *codeconnections.CreateHostInput, optFns ...func(*codeconnections.Options)) (*codeconnections.CreateHostOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateHost")
	}

	var r0 *codeconnections.CreateHostOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *codeconnections.CreateHostInput, ...func(*codeconnections.Options)) (*codeconnections.CreateHostOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *codeconnections.CreateHostInput, ...func(*codeconnections.Options)) *codeconnections.CreateHostOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*codeconnections.CreateHostOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *codeconnections.CreateHostInput, ...func(*codeconnections.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateRepositoryLink provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) CreateRepositoryLink(ctx context.Context, params *codeconnections.CreateRepositoryLinkInput, optFns ...func(*codeconnections.Options)) (*codeconnections.CreateRepositoryLinkOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateRepositoryLink")
	}

	var r0 *codeconnections.CreateRepositoryLinkOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *codeconnections.CreateRepositoryLinkInput, ...func(*codeconnections.Options)) (*codeconnections.CreateRepositoryLinkOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *codeconnections.CreateRepositoryLinkInput, ...func(*codeconnections.Options)) *codeconnections.CreateRepositoryLinkOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*codeconnections.CreateRepositoryLinkOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *codeconnections.CreateRepositoryLinkInput, ...func(*codeconnections.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateSyncConfiguration provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) CreateSyncConfiguration(ctx context.Context, params *codeconnections.CreateSyncConfigurationInput, optFns ...func(*codeconnections.Options)) (*codeconnections.CreateSyncConfigurationOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateSyncConfiguration")
	}

	var r0 *codeconnections.CreateSyncConfigurationOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *codeconnections.CreateSyncConfigurationInput, ...func(*codeconnections.Options)) (*codeconnections.CreateSyncConfigurationOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *codeconnections.CreateSyncConfigurationInput, ...func(*codeconnections.Options)) *codeconnections.CreateSyncConfigurationOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*codeconnections.CreateSyncConfigurationOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *codeconnections.CreateSyncConfigurationInput, ...func(*codeconnections.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteConnection provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DeleteConnection(ctx context.Context, params *codeconnections.DeleteConnectionInput, optFns ...func(*codeconnections.Options)) (*codeconnections.DeleteConnectionOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteConnection")
	}

	var r0 *codeconnections.DeleteConnectionOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *codeconnections.DeleteConnectionInput, ...func(*codeconnections.Options)) (*codeconnections.DeleteConnectionOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *codeconnections.DeleteConnectionInput, ...func(*codeconnections.Options)) *codeconnections.DeleteConnectionOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*codeconnections.DeleteConnectionOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *codeconnections.DeleteConnectionInput, ...func(*codeconnections.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteHost provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DeleteHost(ctx context.Context, params *codeconnections.DeleteHostInput, optFns ...func(*codeconnections.Options)) (*codeconnections.DeleteHostOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteHost")
	}

	var r0 *codeconnections.DeleteHostOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *codeconnections.DeleteHostInput, ...func(*codeconnections.Options)) (*codeconnections.DeleteHostOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *codeconnections.DeleteHostInput, ...func(*codeconnections.Options)) *codeconnections.DeleteHostOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*codeconnections.DeleteHostOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *codeconnections.DeleteHostInput, ...func(*codeconnections.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteRepositoryLink provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DeleteRepositoryLink(ctx context.Context, params *codeconnections.DeleteRepositoryLinkInput, optFns ...func(*codeconnections.Options)) (*codeconnections.DeleteRepositoryLinkOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteRepositoryLink")
	}

	var r0 *codeconnections.DeleteRepositoryLinkOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *codeconnections.DeleteRepositoryLinkInput, ...func(*codeconnections.Options)) (*codeconnections.DeleteRepositoryLinkOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *codeconnections.DeleteRepositoryLinkInput, ...func(*codeconnections.Options)) *codeconnections.DeleteRepositoryLinkOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*codeconnections.DeleteRepositoryLinkOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *codeconnections.DeleteRepositoryLinkInput, ...func(*codeconnections.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteSyncConfiguration provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DeleteSyncConfiguration(ctx context.Context, params *codeconnections.DeleteSyncConfigurationInput, optFns ...func(*codeconnections.Options)) (*codeconnections.DeleteSyncConfigurationOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteSyncConfiguration")
	}

	var r0 *codeconnections.DeleteSyncConfigurationOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *codeconnections.DeleteSyncConfigurationInput, ...func(*codeconnections.Options)) (*codeconnections.DeleteSyncConfigurationOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *codeconnections.DeleteSyncConfigurationInput, ...func(*codeconnections.Options)) *codeconnections.DeleteSyncConfigurationOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*codeconnections.DeleteSyncConfigurationOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *codeconnections.DeleteSyncConfigurationInput, ...func(*codeconnections.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetConnection provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetConnection(ctx context.Context, params *codeconnections.GetConnectionInput, optFns ...func(*codeconnections.Options)) (*codeconnections.GetConnectionOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetConnection")
	}

	var r0 *codeconnections.GetConnectionOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *codeconnections.GetConnectionInput, ...func(*codeconnections.Options)) (*codeconnections.GetConnectionOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *codeconnections.GetConnectionInput, ...func(*codeconnections.Options)) *codeconnections.GetConnectionOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*codeconnections.GetConnectionOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *codeconnections.GetConnectionInput, ...func(*codeconnections.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetHost provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetHost(ctx context.Context, params *codeconnections.GetHostInput, optFns ...func(*codeconnections.Options)) (*codeconnections.GetHostOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetHost")
	}

	var r0 *codeconnections.GetHostOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *codeconnections.GetHostInput, ...func(*codeconnections.Options)) (*codeconnections.GetHostOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *codeconnections.GetHostInput, ...func(*codeconnections.Options)) *codeconnections.GetHostOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*codeconnections.GetHostOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *codeconnections.GetHostInput, ...func(*codeconnections.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRepositoryLink provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetRepositoryLink(ctx context.Context, params *codeconnections.GetRepositoryLinkInput, optFns ...func(*codeconnections.Options)) (*codeconnections.GetRepositoryLinkOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetRepositoryLink")
	}

	var r0 *codeconnections.GetRepositoryLinkOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *codeconnections.GetRepositoryLinkInput, ...func(*codeconnections.Options)) (*codeconnections.GetRepositoryLinkOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *codeconnections.GetRepositoryLinkInput, ...func(*codeconnections.Options)) *codeconnections.GetRepositoryLinkOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*codeconnections.GetRepositoryLinkOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *codeconnections.GetRepositoryLinkInput, ...func(*codeconnections.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRepositorySyncStatus provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetRepositorySyncStatus(ctx context.Context, params *codeconnections.GetRepositorySyncStatusInput, optFns ...func(*codeconnections.Options)) (*codeconnections.GetRepositorySyncStatusOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetRepositorySyncStatus")
	}

	var r0 *codeconnections.GetRepositorySyncStatusOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *codeconnections.GetRepositorySyncStatusInput, ...func(*codeconnections.Options)) (*codeconnections.GetRepositorySyncStatusOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *codeconnections.GetRepositorySyncStatusInput, ...func(*codeconnections.Options)) *codeconnections.GetRepositorySyncStatusOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*codeconnections.GetRepositorySyncStatusOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *codeconnections.GetRepositorySyncStatusInput, ...func(*codeconnections.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetResourceSyncStatus provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetResourceSyncStatus(ctx context.Context, params *codeconnections.GetResourceSyncStatusInput, optFns ...func(*codeconnections.Options)) (*codeconnections.GetResourceSyncStatusOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetResourceSyncStatus")
	}

	var r0 *codeconnections.GetResourceSyncStatusOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *codeconnections.GetResourceSyncStatusInput, ...func(*codeconnections.Options)) (*codeconnections.GetResourceSyncStatusOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *codeconnections.GetResourceSyncStatusInput, ...func(*codeconnections.Options)) *codeconnections.GetResourceSyncStatusOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*codeconnections.GetResourceSyncStatusOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *codeconnections.GetResourceSyncStatusInput, ...func(*codeconnections.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSyncBlockerSummary provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetSyncBlockerSummary(ctx context.Context, params *codeconnections.GetSyncBlockerSummaryInput, optFns ...func(*codeconnections.Options)) (*codeconnections.GetSyncBlockerSummaryOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetSyncBlockerSummary")
	}

	var r0 *codeconnections.GetSyncBlockerSummaryOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *codeconnections.GetSyncBlockerSummaryInput, ...func(*codeconnections.Options)) (*codeconnections.GetSyncBlockerSummaryOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *codeconnections.GetSyncBlockerSummaryInput, ...func(*codeconnections.Options)) *codeconnections.GetSyncBlockerSummaryOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*codeconnections.GetSyncBlockerSummaryOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *codeconnections.GetSyncBlockerSummaryInput, ...func(*codeconnections.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSyncConfiguration provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetSyncConfiguration(ctx context.Context, params *codeconnections.GetSyncConfigurationInput, optFns ...func(*codeconnections.Options)) (*codeconnections.GetSyncConfigurationOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetSyncConfiguration")
	}

	var r0 *codeconnections.GetSyncConfigurationOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *codeconnections.GetSyncConfigurationInput, ...func(*codeconnections.Options)) (*codeconnections.GetSyncConfigurationOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *codeconnections.GetSyncConfigurationInput, ...func(*codeconnections.Options)) *codeconnections.GetSyncConfigurationOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*codeconnections.GetSyncConfigurationOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *codeconnections.GetSyncConfigurationInput, ...func(*codeconnections.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListConnections provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListConnections(ctx context.Context, params *codeconnections.ListConnectionsInput, optFns ...func(*codeconnections.Options)) (*codeconnections.ListConnectionsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListConnections")
	}

	var r0 *codeconnections.ListConnectionsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *codeconnections.ListConnectionsInput, ...func(*codeconnections.Options)) (*codeconnections.ListConnectionsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *codeconnections.ListConnectionsInput, ...func(*codeconnections.Options)) *codeconnections.ListConnectionsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*codeconnections.ListConnectionsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *codeconnections.ListConnectionsInput, ...func(*codeconnections.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListHosts provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListHosts(ctx context.Context, params *codeconnections.ListHostsInput, optFns ...func(*codeconnections.Options)) (*codeconnections.ListHostsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListHosts")
	}

	var r0 *codeconnections.ListHostsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *codeconnections.ListHostsInput, ...func(*codeconnections.Options)) (*codeconnections.ListHostsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *codeconnections.ListHostsInput, ...func(*codeconnections.Options)) *codeconnections.ListHostsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*codeconnections.ListHostsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *codeconnections.ListHostsInput, ...func(*codeconnections.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListRepositoryLinks provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListRepositoryLinks(ctx context.Context, params *codeconnections.ListRepositoryLinksInput, optFns ...func(*codeconnections.Options)) (*codeconnections.ListRepositoryLinksOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListRepositoryLinks")
	}

	var r0 *codeconnections.ListRepositoryLinksOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *codeconnections.ListRepositoryLinksInput, ...func(*codeconnections.Options)) (*codeconnections.ListRepositoryLinksOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *codeconnections.ListRepositoryLinksInput, ...func(*codeconnections.Options)) *codeconnections.ListRepositoryLinksOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*codeconnections.ListRepositoryLinksOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *codeconnections.ListRepositoryLinksInput, ...func(*codeconnections.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListRepositorySyncDefinitions provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListRepositorySyncDefinitions(ctx context.Context, params *codeconnections.ListRepositorySyncDefinitionsInput, optFns ...func(*codeconnections.Options)) (*codeconnections.ListRepositorySyncDefinitionsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListRepositorySyncDefinitions")
	}

	var r0 *codeconnections.ListRepositorySyncDefinitionsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *codeconnections.ListRepositorySyncDefinitionsInput, ...func(*codeconnections.Options)) (*codeconnections.ListRepositorySyncDefinitionsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *codeconnections.ListRepositorySyncDefinitionsInput, ...func(*codeconnections.Options)) *codeconnections.ListRepositorySyncDefinitionsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*codeconnections.ListRepositorySyncDefinitionsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *codeconnections.ListRepositorySyncDefinitionsInput, ...func(*codeconnections.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListSyncConfigurations provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListSyncConfigurations(ctx context.Context, params *codeconnections.ListSyncConfigurationsInput, optFns ...func(*codeconnections.Options)) (*codeconnections.ListSyncConfigurationsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListSyncConfigurations")
	}

	var r0 *codeconnections.ListSyncConfigurationsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *codeconnections.ListSyncConfigurationsInput, ...func(*codeconnections.Options)) (*codeconnections.ListSyncConfigurationsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *codeconnections.ListSyncConfigurationsInput, ...func(*codeconnections.Options)) *codeconnections.ListSyncConfigurationsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*codeconnections.ListSyncConfigurationsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *codeconnections.ListSyncConfigurationsInput, ...func(*codeconnections.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListTagsForResource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListTagsForResource(ctx context.Context, params *codeconnections.ListTagsForResourceInput, optFns ...func(*codeconnections.Options)) (*codeconnections.ListTagsForResourceOutput, error) {
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

	var r0 *codeconnections.ListTagsForResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *codeconnections.ListTagsForResourceInput, ...func(*codeconnections.Options)) (*codeconnections.ListTagsForResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *codeconnections.ListTagsForResourceInput, ...func(*codeconnections.Options)) *codeconnections.ListTagsForResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*codeconnections.ListTagsForResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *codeconnections.ListTagsForResourceInput, ...func(*codeconnections.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Options provides a mock function with given fields:
func (_m *IClient) Options() codeconnections.Options {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Options")
	}

	var r0 codeconnections.Options
	if rf, ok := ret.Get(0).(func() codeconnections.Options); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(codeconnections.Options)
	}

	return r0
}

// TagResource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) TagResource(ctx context.Context, params *codeconnections.TagResourceInput, optFns ...func(*codeconnections.Options)) (*codeconnections.TagResourceOutput, error) {
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

	var r0 *codeconnections.TagResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *codeconnections.TagResourceInput, ...func(*codeconnections.Options)) (*codeconnections.TagResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *codeconnections.TagResourceInput, ...func(*codeconnections.Options)) *codeconnections.TagResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*codeconnections.TagResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *codeconnections.TagResourceInput, ...func(*codeconnections.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UntagResource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) UntagResource(ctx context.Context, params *codeconnections.UntagResourceInput, optFns ...func(*codeconnections.Options)) (*codeconnections.UntagResourceOutput, error) {
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

	var r0 *codeconnections.UntagResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *codeconnections.UntagResourceInput, ...func(*codeconnections.Options)) (*codeconnections.UntagResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *codeconnections.UntagResourceInput, ...func(*codeconnections.Options)) *codeconnections.UntagResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*codeconnections.UntagResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *codeconnections.UntagResourceInput, ...func(*codeconnections.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateHost provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) UpdateHost(ctx context.Context, params *codeconnections.UpdateHostInput, optFns ...func(*codeconnections.Options)) (*codeconnections.UpdateHostOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UpdateHost")
	}

	var r0 *codeconnections.UpdateHostOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *codeconnections.UpdateHostInput, ...func(*codeconnections.Options)) (*codeconnections.UpdateHostOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *codeconnections.UpdateHostInput, ...func(*codeconnections.Options)) *codeconnections.UpdateHostOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*codeconnections.UpdateHostOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *codeconnections.UpdateHostInput, ...func(*codeconnections.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateRepositoryLink provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) UpdateRepositoryLink(ctx context.Context, params *codeconnections.UpdateRepositoryLinkInput, optFns ...func(*codeconnections.Options)) (*codeconnections.UpdateRepositoryLinkOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UpdateRepositoryLink")
	}

	var r0 *codeconnections.UpdateRepositoryLinkOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *codeconnections.UpdateRepositoryLinkInput, ...func(*codeconnections.Options)) (*codeconnections.UpdateRepositoryLinkOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *codeconnections.UpdateRepositoryLinkInput, ...func(*codeconnections.Options)) *codeconnections.UpdateRepositoryLinkOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*codeconnections.UpdateRepositoryLinkOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *codeconnections.UpdateRepositoryLinkInput, ...func(*codeconnections.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateSyncBlocker provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) UpdateSyncBlocker(ctx context.Context, params *codeconnections.UpdateSyncBlockerInput, optFns ...func(*codeconnections.Options)) (*codeconnections.UpdateSyncBlockerOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UpdateSyncBlocker")
	}

	var r0 *codeconnections.UpdateSyncBlockerOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *codeconnections.UpdateSyncBlockerInput, ...func(*codeconnections.Options)) (*codeconnections.UpdateSyncBlockerOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *codeconnections.UpdateSyncBlockerInput, ...func(*codeconnections.Options)) *codeconnections.UpdateSyncBlockerOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*codeconnections.UpdateSyncBlockerOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *codeconnections.UpdateSyncBlockerInput, ...func(*codeconnections.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateSyncConfiguration provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) UpdateSyncConfiguration(ctx context.Context, params *codeconnections.UpdateSyncConfigurationInput, optFns ...func(*codeconnections.Options)) (*codeconnections.UpdateSyncConfigurationOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UpdateSyncConfiguration")
	}

	var r0 *codeconnections.UpdateSyncConfigurationOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *codeconnections.UpdateSyncConfigurationInput, ...func(*codeconnections.Options)) (*codeconnections.UpdateSyncConfigurationOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *codeconnections.UpdateSyncConfigurationInput, ...func(*codeconnections.Options)) *codeconnections.UpdateSyncConfigurationOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*codeconnections.UpdateSyncConfigurationOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *codeconnections.UpdateSyncConfigurationInput, ...func(*codeconnections.Options)) error); ok {
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
