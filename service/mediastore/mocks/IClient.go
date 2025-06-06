// Code generated by mockery v2.53.4. DO NOT EDIT.

package mocks

import (
	context "context"

	mediastore "github.com/aws/aws-sdk-go-v2/service/mediastore"

	mock "github.com/stretchr/testify/mock"
)

// IClient is an autogenerated mock type for the IClient type
type IClient struct {
	mock.Mock
}

// CreateContainer provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) CreateContainer(ctx context.Context, params *mediastore.CreateContainerInput, optFns ...func(*mediastore.Options)) (*mediastore.CreateContainerOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateContainer")
	}

	var r0 *mediastore.CreateContainerOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *mediastore.CreateContainerInput, ...func(*mediastore.Options)) (*mediastore.CreateContainerOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *mediastore.CreateContainerInput, ...func(*mediastore.Options)) *mediastore.CreateContainerOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mediastore.CreateContainerOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *mediastore.CreateContainerInput, ...func(*mediastore.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteContainer provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DeleteContainer(ctx context.Context, params *mediastore.DeleteContainerInput, optFns ...func(*mediastore.Options)) (*mediastore.DeleteContainerOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteContainer")
	}

	var r0 *mediastore.DeleteContainerOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *mediastore.DeleteContainerInput, ...func(*mediastore.Options)) (*mediastore.DeleteContainerOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *mediastore.DeleteContainerInput, ...func(*mediastore.Options)) *mediastore.DeleteContainerOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mediastore.DeleteContainerOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *mediastore.DeleteContainerInput, ...func(*mediastore.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteContainerPolicy provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DeleteContainerPolicy(ctx context.Context, params *mediastore.DeleteContainerPolicyInput, optFns ...func(*mediastore.Options)) (*mediastore.DeleteContainerPolicyOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteContainerPolicy")
	}

	var r0 *mediastore.DeleteContainerPolicyOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *mediastore.DeleteContainerPolicyInput, ...func(*mediastore.Options)) (*mediastore.DeleteContainerPolicyOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *mediastore.DeleteContainerPolicyInput, ...func(*mediastore.Options)) *mediastore.DeleteContainerPolicyOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mediastore.DeleteContainerPolicyOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *mediastore.DeleteContainerPolicyInput, ...func(*mediastore.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteCorsPolicy provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DeleteCorsPolicy(ctx context.Context, params *mediastore.DeleteCorsPolicyInput, optFns ...func(*mediastore.Options)) (*mediastore.DeleteCorsPolicyOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteCorsPolicy")
	}

	var r0 *mediastore.DeleteCorsPolicyOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *mediastore.DeleteCorsPolicyInput, ...func(*mediastore.Options)) (*mediastore.DeleteCorsPolicyOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *mediastore.DeleteCorsPolicyInput, ...func(*mediastore.Options)) *mediastore.DeleteCorsPolicyOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mediastore.DeleteCorsPolicyOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *mediastore.DeleteCorsPolicyInput, ...func(*mediastore.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteLifecyclePolicy provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DeleteLifecyclePolicy(ctx context.Context, params *mediastore.DeleteLifecyclePolicyInput, optFns ...func(*mediastore.Options)) (*mediastore.DeleteLifecyclePolicyOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteLifecyclePolicy")
	}

	var r0 *mediastore.DeleteLifecyclePolicyOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *mediastore.DeleteLifecyclePolicyInput, ...func(*mediastore.Options)) (*mediastore.DeleteLifecyclePolicyOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *mediastore.DeleteLifecyclePolicyInput, ...func(*mediastore.Options)) *mediastore.DeleteLifecyclePolicyOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mediastore.DeleteLifecyclePolicyOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *mediastore.DeleteLifecyclePolicyInput, ...func(*mediastore.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteMetricPolicy provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DeleteMetricPolicy(ctx context.Context, params *mediastore.DeleteMetricPolicyInput, optFns ...func(*mediastore.Options)) (*mediastore.DeleteMetricPolicyOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteMetricPolicy")
	}

	var r0 *mediastore.DeleteMetricPolicyOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *mediastore.DeleteMetricPolicyInput, ...func(*mediastore.Options)) (*mediastore.DeleteMetricPolicyOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *mediastore.DeleteMetricPolicyInput, ...func(*mediastore.Options)) *mediastore.DeleteMetricPolicyOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mediastore.DeleteMetricPolicyOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *mediastore.DeleteMetricPolicyInput, ...func(*mediastore.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeContainer provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DescribeContainer(ctx context.Context, params *mediastore.DescribeContainerInput, optFns ...func(*mediastore.Options)) (*mediastore.DescribeContainerOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DescribeContainer")
	}

	var r0 *mediastore.DescribeContainerOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *mediastore.DescribeContainerInput, ...func(*mediastore.Options)) (*mediastore.DescribeContainerOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *mediastore.DescribeContainerInput, ...func(*mediastore.Options)) *mediastore.DescribeContainerOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mediastore.DescribeContainerOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *mediastore.DescribeContainerInput, ...func(*mediastore.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetContainerPolicy provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetContainerPolicy(ctx context.Context, params *mediastore.GetContainerPolicyInput, optFns ...func(*mediastore.Options)) (*mediastore.GetContainerPolicyOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetContainerPolicy")
	}

	var r0 *mediastore.GetContainerPolicyOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *mediastore.GetContainerPolicyInput, ...func(*mediastore.Options)) (*mediastore.GetContainerPolicyOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *mediastore.GetContainerPolicyInput, ...func(*mediastore.Options)) *mediastore.GetContainerPolicyOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mediastore.GetContainerPolicyOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *mediastore.GetContainerPolicyInput, ...func(*mediastore.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCorsPolicy provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetCorsPolicy(ctx context.Context, params *mediastore.GetCorsPolicyInput, optFns ...func(*mediastore.Options)) (*mediastore.GetCorsPolicyOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetCorsPolicy")
	}

	var r0 *mediastore.GetCorsPolicyOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *mediastore.GetCorsPolicyInput, ...func(*mediastore.Options)) (*mediastore.GetCorsPolicyOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *mediastore.GetCorsPolicyInput, ...func(*mediastore.Options)) *mediastore.GetCorsPolicyOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mediastore.GetCorsPolicyOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *mediastore.GetCorsPolicyInput, ...func(*mediastore.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLifecyclePolicy provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetLifecyclePolicy(ctx context.Context, params *mediastore.GetLifecyclePolicyInput, optFns ...func(*mediastore.Options)) (*mediastore.GetLifecyclePolicyOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetLifecyclePolicy")
	}

	var r0 *mediastore.GetLifecyclePolicyOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *mediastore.GetLifecyclePolicyInput, ...func(*mediastore.Options)) (*mediastore.GetLifecyclePolicyOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *mediastore.GetLifecyclePolicyInput, ...func(*mediastore.Options)) *mediastore.GetLifecyclePolicyOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mediastore.GetLifecyclePolicyOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *mediastore.GetLifecyclePolicyInput, ...func(*mediastore.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMetricPolicy provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetMetricPolicy(ctx context.Context, params *mediastore.GetMetricPolicyInput, optFns ...func(*mediastore.Options)) (*mediastore.GetMetricPolicyOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetMetricPolicy")
	}

	var r0 *mediastore.GetMetricPolicyOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *mediastore.GetMetricPolicyInput, ...func(*mediastore.Options)) (*mediastore.GetMetricPolicyOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *mediastore.GetMetricPolicyInput, ...func(*mediastore.Options)) *mediastore.GetMetricPolicyOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mediastore.GetMetricPolicyOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *mediastore.GetMetricPolicyInput, ...func(*mediastore.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListContainers provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListContainers(ctx context.Context, params *mediastore.ListContainersInput, optFns ...func(*mediastore.Options)) (*mediastore.ListContainersOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListContainers")
	}

	var r0 *mediastore.ListContainersOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *mediastore.ListContainersInput, ...func(*mediastore.Options)) (*mediastore.ListContainersOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *mediastore.ListContainersInput, ...func(*mediastore.Options)) *mediastore.ListContainersOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mediastore.ListContainersOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *mediastore.ListContainersInput, ...func(*mediastore.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListTagsForResource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListTagsForResource(ctx context.Context, params *mediastore.ListTagsForResourceInput, optFns ...func(*mediastore.Options)) (*mediastore.ListTagsForResourceOutput, error) {
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

	var r0 *mediastore.ListTagsForResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *mediastore.ListTagsForResourceInput, ...func(*mediastore.Options)) (*mediastore.ListTagsForResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *mediastore.ListTagsForResourceInput, ...func(*mediastore.Options)) *mediastore.ListTagsForResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mediastore.ListTagsForResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *mediastore.ListTagsForResourceInput, ...func(*mediastore.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Options provides a mock function with no fields
func (_m *IClient) Options() mediastore.Options {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Options")
	}

	var r0 mediastore.Options
	if rf, ok := ret.Get(0).(func() mediastore.Options); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(mediastore.Options)
	}

	return r0
}

// PutContainerPolicy provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) PutContainerPolicy(ctx context.Context, params *mediastore.PutContainerPolicyInput, optFns ...func(*mediastore.Options)) (*mediastore.PutContainerPolicyOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for PutContainerPolicy")
	}

	var r0 *mediastore.PutContainerPolicyOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *mediastore.PutContainerPolicyInput, ...func(*mediastore.Options)) (*mediastore.PutContainerPolicyOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *mediastore.PutContainerPolicyInput, ...func(*mediastore.Options)) *mediastore.PutContainerPolicyOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mediastore.PutContainerPolicyOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *mediastore.PutContainerPolicyInput, ...func(*mediastore.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PutCorsPolicy provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) PutCorsPolicy(ctx context.Context, params *mediastore.PutCorsPolicyInput, optFns ...func(*mediastore.Options)) (*mediastore.PutCorsPolicyOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for PutCorsPolicy")
	}

	var r0 *mediastore.PutCorsPolicyOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *mediastore.PutCorsPolicyInput, ...func(*mediastore.Options)) (*mediastore.PutCorsPolicyOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *mediastore.PutCorsPolicyInput, ...func(*mediastore.Options)) *mediastore.PutCorsPolicyOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mediastore.PutCorsPolicyOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *mediastore.PutCorsPolicyInput, ...func(*mediastore.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PutLifecyclePolicy provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) PutLifecyclePolicy(ctx context.Context, params *mediastore.PutLifecyclePolicyInput, optFns ...func(*mediastore.Options)) (*mediastore.PutLifecyclePolicyOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for PutLifecyclePolicy")
	}

	var r0 *mediastore.PutLifecyclePolicyOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *mediastore.PutLifecyclePolicyInput, ...func(*mediastore.Options)) (*mediastore.PutLifecyclePolicyOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *mediastore.PutLifecyclePolicyInput, ...func(*mediastore.Options)) *mediastore.PutLifecyclePolicyOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mediastore.PutLifecyclePolicyOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *mediastore.PutLifecyclePolicyInput, ...func(*mediastore.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PutMetricPolicy provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) PutMetricPolicy(ctx context.Context, params *mediastore.PutMetricPolicyInput, optFns ...func(*mediastore.Options)) (*mediastore.PutMetricPolicyOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for PutMetricPolicy")
	}

	var r0 *mediastore.PutMetricPolicyOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *mediastore.PutMetricPolicyInput, ...func(*mediastore.Options)) (*mediastore.PutMetricPolicyOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *mediastore.PutMetricPolicyInput, ...func(*mediastore.Options)) *mediastore.PutMetricPolicyOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mediastore.PutMetricPolicyOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *mediastore.PutMetricPolicyInput, ...func(*mediastore.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StartAccessLogging provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) StartAccessLogging(ctx context.Context, params *mediastore.StartAccessLoggingInput, optFns ...func(*mediastore.Options)) (*mediastore.StartAccessLoggingOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for StartAccessLogging")
	}

	var r0 *mediastore.StartAccessLoggingOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *mediastore.StartAccessLoggingInput, ...func(*mediastore.Options)) (*mediastore.StartAccessLoggingOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *mediastore.StartAccessLoggingInput, ...func(*mediastore.Options)) *mediastore.StartAccessLoggingOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mediastore.StartAccessLoggingOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *mediastore.StartAccessLoggingInput, ...func(*mediastore.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StopAccessLogging provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) StopAccessLogging(ctx context.Context, params *mediastore.StopAccessLoggingInput, optFns ...func(*mediastore.Options)) (*mediastore.StopAccessLoggingOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for StopAccessLogging")
	}

	var r0 *mediastore.StopAccessLoggingOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *mediastore.StopAccessLoggingInput, ...func(*mediastore.Options)) (*mediastore.StopAccessLoggingOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *mediastore.StopAccessLoggingInput, ...func(*mediastore.Options)) *mediastore.StopAccessLoggingOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mediastore.StopAccessLoggingOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *mediastore.StopAccessLoggingInput, ...func(*mediastore.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TagResource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) TagResource(ctx context.Context, params *mediastore.TagResourceInput, optFns ...func(*mediastore.Options)) (*mediastore.TagResourceOutput, error) {
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

	var r0 *mediastore.TagResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *mediastore.TagResourceInput, ...func(*mediastore.Options)) (*mediastore.TagResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *mediastore.TagResourceInput, ...func(*mediastore.Options)) *mediastore.TagResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mediastore.TagResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *mediastore.TagResourceInput, ...func(*mediastore.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UntagResource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) UntagResource(ctx context.Context, params *mediastore.UntagResourceInput, optFns ...func(*mediastore.Options)) (*mediastore.UntagResourceOutput, error) {
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

	var r0 *mediastore.UntagResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *mediastore.UntagResourceInput, ...func(*mediastore.Options)) (*mediastore.UntagResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *mediastore.UntagResourceInput, ...func(*mediastore.Options)) *mediastore.UntagResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mediastore.UntagResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *mediastore.UntagResourceInput, ...func(*mediastore.Options)) error); ok {
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
