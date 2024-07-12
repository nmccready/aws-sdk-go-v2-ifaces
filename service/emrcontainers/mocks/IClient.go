// Code generated by mockery v2.39.2. DO NOT EDIT.

package mocks

import (
	context "context"

	emrcontainers "github.com/aws/aws-sdk-go-v2/service/emrcontainers"
	mock "github.com/stretchr/testify/mock"
)

// IClient is an autogenerated mock type for the IClient type
type IClient struct {
	mock.Mock
}

// CancelJobRun provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) CancelJobRun(ctx context.Context, params *emrcontainers.CancelJobRunInput, optFns ...func(*emrcontainers.Options)) (*emrcontainers.CancelJobRunOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CancelJobRun")
	}

	var r0 *emrcontainers.CancelJobRunOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *emrcontainers.CancelJobRunInput, ...func(*emrcontainers.Options)) (*emrcontainers.CancelJobRunOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *emrcontainers.CancelJobRunInput, ...func(*emrcontainers.Options)) *emrcontainers.CancelJobRunOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*emrcontainers.CancelJobRunOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *emrcontainers.CancelJobRunInput, ...func(*emrcontainers.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateJobTemplate provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) CreateJobTemplate(ctx context.Context, params *emrcontainers.CreateJobTemplateInput, optFns ...func(*emrcontainers.Options)) (*emrcontainers.CreateJobTemplateOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateJobTemplate")
	}

	var r0 *emrcontainers.CreateJobTemplateOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *emrcontainers.CreateJobTemplateInput, ...func(*emrcontainers.Options)) (*emrcontainers.CreateJobTemplateOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *emrcontainers.CreateJobTemplateInput, ...func(*emrcontainers.Options)) *emrcontainers.CreateJobTemplateOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*emrcontainers.CreateJobTemplateOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *emrcontainers.CreateJobTemplateInput, ...func(*emrcontainers.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateManagedEndpoint provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) CreateManagedEndpoint(ctx context.Context, params *emrcontainers.CreateManagedEndpointInput, optFns ...func(*emrcontainers.Options)) (*emrcontainers.CreateManagedEndpointOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateManagedEndpoint")
	}

	var r0 *emrcontainers.CreateManagedEndpointOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *emrcontainers.CreateManagedEndpointInput, ...func(*emrcontainers.Options)) (*emrcontainers.CreateManagedEndpointOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *emrcontainers.CreateManagedEndpointInput, ...func(*emrcontainers.Options)) *emrcontainers.CreateManagedEndpointOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*emrcontainers.CreateManagedEndpointOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *emrcontainers.CreateManagedEndpointInput, ...func(*emrcontainers.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateSecurityConfiguration provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) CreateSecurityConfiguration(ctx context.Context, params *emrcontainers.CreateSecurityConfigurationInput, optFns ...func(*emrcontainers.Options)) (*emrcontainers.CreateSecurityConfigurationOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateSecurityConfiguration")
	}

	var r0 *emrcontainers.CreateSecurityConfigurationOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *emrcontainers.CreateSecurityConfigurationInput, ...func(*emrcontainers.Options)) (*emrcontainers.CreateSecurityConfigurationOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *emrcontainers.CreateSecurityConfigurationInput, ...func(*emrcontainers.Options)) *emrcontainers.CreateSecurityConfigurationOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*emrcontainers.CreateSecurityConfigurationOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *emrcontainers.CreateSecurityConfigurationInput, ...func(*emrcontainers.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateVirtualCluster provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) CreateVirtualCluster(ctx context.Context, params *emrcontainers.CreateVirtualClusterInput, optFns ...func(*emrcontainers.Options)) (*emrcontainers.CreateVirtualClusterOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateVirtualCluster")
	}

	var r0 *emrcontainers.CreateVirtualClusterOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *emrcontainers.CreateVirtualClusterInput, ...func(*emrcontainers.Options)) (*emrcontainers.CreateVirtualClusterOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *emrcontainers.CreateVirtualClusterInput, ...func(*emrcontainers.Options)) *emrcontainers.CreateVirtualClusterOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*emrcontainers.CreateVirtualClusterOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *emrcontainers.CreateVirtualClusterInput, ...func(*emrcontainers.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteJobTemplate provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DeleteJobTemplate(ctx context.Context, params *emrcontainers.DeleteJobTemplateInput, optFns ...func(*emrcontainers.Options)) (*emrcontainers.DeleteJobTemplateOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteJobTemplate")
	}

	var r0 *emrcontainers.DeleteJobTemplateOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *emrcontainers.DeleteJobTemplateInput, ...func(*emrcontainers.Options)) (*emrcontainers.DeleteJobTemplateOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *emrcontainers.DeleteJobTemplateInput, ...func(*emrcontainers.Options)) *emrcontainers.DeleteJobTemplateOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*emrcontainers.DeleteJobTemplateOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *emrcontainers.DeleteJobTemplateInput, ...func(*emrcontainers.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteManagedEndpoint provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DeleteManagedEndpoint(ctx context.Context, params *emrcontainers.DeleteManagedEndpointInput, optFns ...func(*emrcontainers.Options)) (*emrcontainers.DeleteManagedEndpointOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteManagedEndpoint")
	}

	var r0 *emrcontainers.DeleteManagedEndpointOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *emrcontainers.DeleteManagedEndpointInput, ...func(*emrcontainers.Options)) (*emrcontainers.DeleteManagedEndpointOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *emrcontainers.DeleteManagedEndpointInput, ...func(*emrcontainers.Options)) *emrcontainers.DeleteManagedEndpointOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*emrcontainers.DeleteManagedEndpointOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *emrcontainers.DeleteManagedEndpointInput, ...func(*emrcontainers.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteVirtualCluster provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DeleteVirtualCluster(ctx context.Context, params *emrcontainers.DeleteVirtualClusterInput, optFns ...func(*emrcontainers.Options)) (*emrcontainers.DeleteVirtualClusterOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteVirtualCluster")
	}

	var r0 *emrcontainers.DeleteVirtualClusterOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *emrcontainers.DeleteVirtualClusterInput, ...func(*emrcontainers.Options)) (*emrcontainers.DeleteVirtualClusterOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *emrcontainers.DeleteVirtualClusterInput, ...func(*emrcontainers.Options)) *emrcontainers.DeleteVirtualClusterOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*emrcontainers.DeleteVirtualClusterOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *emrcontainers.DeleteVirtualClusterInput, ...func(*emrcontainers.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeJobRun provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DescribeJobRun(ctx context.Context, params *emrcontainers.DescribeJobRunInput, optFns ...func(*emrcontainers.Options)) (*emrcontainers.DescribeJobRunOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DescribeJobRun")
	}

	var r0 *emrcontainers.DescribeJobRunOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *emrcontainers.DescribeJobRunInput, ...func(*emrcontainers.Options)) (*emrcontainers.DescribeJobRunOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *emrcontainers.DescribeJobRunInput, ...func(*emrcontainers.Options)) *emrcontainers.DescribeJobRunOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*emrcontainers.DescribeJobRunOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *emrcontainers.DescribeJobRunInput, ...func(*emrcontainers.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeJobTemplate provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DescribeJobTemplate(ctx context.Context, params *emrcontainers.DescribeJobTemplateInput, optFns ...func(*emrcontainers.Options)) (*emrcontainers.DescribeJobTemplateOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DescribeJobTemplate")
	}

	var r0 *emrcontainers.DescribeJobTemplateOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *emrcontainers.DescribeJobTemplateInput, ...func(*emrcontainers.Options)) (*emrcontainers.DescribeJobTemplateOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *emrcontainers.DescribeJobTemplateInput, ...func(*emrcontainers.Options)) *emrcontainers.DescribeJobTemplateOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*emrcontainers.DescribeJobTemplateOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *emrcontainers.DescribeJobTemplateInput, ...func(*emrcontainers.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeManagedEndpoint provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DescribeManagedEndpoint(ctx context.Context, params *emrcontainers.DescribeManagedEndpointInput, optFns ...func(*emrcontainers.Options)) (*emrcontainers.DescribeManagedEndpointOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DescribeManagedEndpoint")
	}

	var r0 *emrcontainers.DescribeManagedEndpointOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *emrcontainers.DescribeManagedEndpointInput, ...func(*emrcontainers.Options)) (*emrcontainers.DescribeManagedEndpointOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *emrcontainers.DescribeManagedEndpointInput, ...func(*emrcontainers.Options)) *emrcontainers.DescribeManagedEndpointOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*emrcontainers.DescribeManagedEndpointOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *emrcontainers.DescribeManagedEndpointInput, ...func(*emrcontainers.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeSecurityConfiguration provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DescribeSecurityConfiguration(ctx context.Context, params *emrcontainers.DescribeSecurityConfigurationInput, optFns ...func(*emrcontainers.Options)) (*emrcontainers.DescribeSecurityConfigurationOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DescribeSecurityConfiguration")
	}

	var r0 *emrcontainers.DescribeSecurityConfigurationOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *emrcontainers.DescribeSecurityConfigurationInput, ...func(*emrcontainers.Options)) (*emrcontainers.DescribeSecurityConfigurationOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *emrcontainers.DescribeSecurityConfigurationInput, ...func(*emrcontainers.Options)) *emrcontainers.DescribeSecurityConfigurationOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*emrcontainers.DescribeSecurityConfigurationOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *emrcontainers.DescribeSecurityConfigurationInput, ...func(*emrcontainers.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeVirtualCluster provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DescribeVirtualCluster(ctx context.Context, params *emrcontainers.DescribeVirtualClusterInput, optFns ...func(*emrcontainers.Options)) (*emrcontainers.DescribeVirtualClusterOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DescribeVirtualCluster")
	}

	var r0 *emrcontainers.DescribeVirtualClusterOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *emrcontainers.DescribeVirtualClusterInput, ...func(*emrcontainers.Options)) (*emrcontainers.DescribeVirtualClusterOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *emrcontainers.DescribeVirtualClusterInput, ...func(*emrcontainers.Options)) *emrcontainers.DescribeVirtualClusterOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*emrcontainers.DescribeVirtualClusterOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *emrcontainers.DescribeVirtualClusterInput, ...func(*emrcontainers.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetManagedEndpointSessionCredentials provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetManagedEndpointSessionCredentials(ctx context.Context, params *emrcontainers.GetManagedEndpointSessionCredentialsInput, optFns ...func(*emrcontainers.Options)) (*emrcontainers.GetManagedEndpointSessionCredentialsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetManagedEndpointSessionCredentials")
	}

	var r0 *emrcontainers.GetManagedEndpointSessionCredentialsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *emrcontainers.GetManagedEndpointSessionCredentialsInput, ...func(*emrcontainers.Options)) (*emrcontainers.GetManagedEndpointSessionCredentialsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *emrcontainers.GetManagedEndpointSessionCredentialsInput, ...func(*emrcontainers.Options)) *emrcontainers.GetManagedEndpointSessionCredentialsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*emrcontainers.GetManagedEndpointSessionCredentialsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *emrcontainers.GetManagedEndpointSessionCredentialsInput, ...func(*emrcontainers.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListJobRuns provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListJobRuns(ctx context.Context, params *emrcontainers.ListJobRunsInput, optFns ...func(*emrcontainers.Options)) (*emrcontainers.ListJobRunsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListJobRuns")
	}

	var r0 *emrcontainers.ListJobRunsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *emrcontainers.ListJobRunsInput, ...func(*emrcontainers.Options)) (*emrcontainers.ListJobRunsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *emrcontainers.ListJobRunsInput, ...func(*emrcontainers.Options)) *emrcontainers.ListJobRunsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*emrcontainers.ListJobRunsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *emrcontainers.ListJobRunsInput, ...func(*emrcontainers.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListJobTemplates provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListJobTemplates(ctx context.Context, params *emrcontainers.ListJobTemplatesInput, optFns ...func(*emrcontainers.Options)) (*emrcontainers.ListJobTemplatesOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListJobTemplates")
	}

	var r0 *emrcontainers.ListJobTemplatesOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *emrcontainers.ListJobTemplatesInput, ...func(*emrcontainers.Options)) (*emrcontainers.ListJobTemplatesOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *emrcontainers.ListJobTemplatesInput, ...func(*emrcontainers.Options)) *emrcontainers.ListJobTemplatesOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*emrcontainers.ListJobTemplatesOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *emrcontainers.ListJobTemplatesInput, ...func(*emrcontainers.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListManagedEndpoints provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListManagedEndpoints(ctx context.Context, params *emrcontainers.ListManagedEndpointsInput, optFns ...func(*emrcontainers.Options)) (*emrcontainers.ListManagedEndpointsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListManagedEndpoints")
	}

	var r0 *emrcontainers.ListManagedEndpointsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *emrcontainers.ListManagedEndpointsInput, ...func(*emrcontainers.Options)) (*emrcontainers.ListManagedEndpointsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *emrcontainers.ListManagedEndpointsInput, ...func(*emrcontainers.Options)) *emrcontainers.ListManagedEndpointsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*emrcontainers.ListManagedEndpointsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *emrcontainers.ListManagedEndpointsInput, ...func(*emrcontainers.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListSecurityConfigurations provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListSecurityConfigurations(ctx context.Context, params *emrcontainers.ListSecurityConfigurationsInput, optFns ...func(*emrcontainers.Options)) (*emrcontainers.ListSecurityConfigurationsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListSecurityConfigurations")
	}

	var r0 *emrcontainers.ListSecurityConfigurationsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *emrcontainers.ListSecurityConfigurationsInput, ...func(*emrcontainers.Options)) (*emrcontainers.ListSecurityConfigurationsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *emrcontainers.ListSecurityConfigurationsInput, ...func(*emrcontainers.Options)) *emrcontainers.ListSecurityConfigurationsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*emrcontainers.ListSecurityConfigurationsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *emrcontainers.ListSecurityConfigurationsInput, ...func(*emrcontainers.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListTagsForResource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListTagsForResource(ctx context.Context, params *emrcontainers.ListTagsForResourceInput, optFns ...func(*emrcontainers.Options)) (*emrcontainers.ListTagsForResourceOutput, error) {
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

	var r0 *emrcontainers.ListTagsForResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *emrcontainers.ListTagsForResourceInput, ...func(*emrcontainers.Options)) (*emrcontainers.ListTagsForResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *emrcontainers.ListTagsForResourceInput, ...func(*emrcontainers.Options)) *emrcontainers.ListTagsForResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*emrcontainers.ListTagsForResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *emrcontainers.ListTagsForResourceInput, ...func(*emrcontainers.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListVirtualClusters provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListVirtualClusters(ctx context.Context, params *emrcontainers.ListVirtualClustersInput, optFns ...func(*emrcontainers.Options)) (*emrcontainers.ListVirtualClustersOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListVirtualClusters")
	}

	var r0 *emrcontainers.ListVirtualClustersOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *emrcontainers.ListVirtualClustersInput, ...func(*emrcontainers.Options)) (*emrcontainers.ListVirtualClustersOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *emrcontainers.ListVirtualClustersInput, ...func(*emrcontainers.Options)) *emrcontainers.ListVirtualClustersOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*emrcontainers.ListVirtualClustersOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *emrcontainers.ListVirtualClustersInput, ...func(*emrcontainers.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Options provides a mock function with given fields:
func (_m *IClient) Options() emrcontainers.Options {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Options")
	}

	var r0 emrcontainers.Options
	if rf, ok := ret.Get(0).(func() emrcontainers.Options); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(emrcontainers.Options)
	}

	return r0
}

// StartJobRun provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) StartJobRun(ctx context.Context, params *emrcontainers.StartJobRunInput, optFns ...func(*emrcontainers.Options)) (*emrcontainers.StartJobRunOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for StartJobRun")
	}

	var r0 *emrcontainers.StartJobRunOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *emrcontainers.StartJobRunInput, ...func(*emrcontainers.Options)) (*emrcontainers.StartJobRunOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *emrcontainers.StartJobRunInput, ...func(*emrcontainers.Options)) *emrcontainers.StartJobRunOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*emrcontainers.StartJobRunOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *emrcontainers.StartJobRunInput, ...func(*emrcontainers.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TagResource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) TagResource(ctx context.Context, params *emrcontainers.TagResourceInput, optFns ...func(*emrcontainers.Options)) (*emrcontainers.TagResourceOutput, error) {
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

	var r0 *emrcontainers.TagResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *emrcontainers.TagResourceInput, ...func(*emrcontainers.Options)) (*emrcontainers.TagResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *emrcontainers.TagResourceInput, ...func(*emrcontainers.Options)) *emrcontainers.TagResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*emrcontainers.TagResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *emrcontainers.TagResourceInput, ...func(*emrcontainers.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UntagResource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) UntagResource(ctx context.Context, params *emrcontainers.UntagResourceInput, optFns ...func(*emrcontainers.Options)) (*emrcontainers.UntagResourceOutput, error) {
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

	var r0 *emrcontainers.UntagResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *emrcontainers.UntagResourceInput, ...func(*emrcontainers.Options)) (*emrcontainers.UntagResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *emrcontainers.UntagResourceInput, ...func(*emrcontainers.Options)) *emrcontainers.UntagResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*emrcontainers.UntagResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *emrcontainers.UntagResourceInput, ...func(*emrcontainers.Options)) error); ok {
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