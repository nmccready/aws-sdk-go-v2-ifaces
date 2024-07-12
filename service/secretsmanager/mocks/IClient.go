// Code generated by mockery v2.39.2. DO NOT EDIT.

package mocks

import (
	context "context"

	secretsmanager "github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	mock "github.com/stretchr/testify/mock"
)

// IClient is an autogenerated mock type for the IClient type
type IClient struct {
	mock.Mock
}

// BatchGetSecretValue provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) BatchGetSecretValue(ctx context.Context, params *secretsmanager.BatchGetSecretValueInput, optFns ...func(*secretsmanager.Options)) (*secretsmanager.BatchGetSecretValueOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for BatchGetSecretValue")
	}

	var r0 *secretsmanager.BatchGetSecretValueOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *secretsmanager.BatchGetSecretValueInput, ...func(*secretsmanager.Options)) (*secretsmanager.BatchGetSecretValueOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *secretsmanager.BatchGetSecretValueInput, ...func(*secretsmanager.Options)) *secretsmanager.BatchGetSecretValueOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*secretsmanager.BatchGetSecretValueOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *secretsmanager.BatchGetSecretValueInput, ...func(*secretsmanager.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CancelRotateSecret provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) CancelRotateSecret(ctx context.Context, params *secretsmanager.CancelRotateSecretInput, optFns ...func(*secretsmanager.Options)) (*secretsmanager.CancelRotateSecretOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CancelRotateSecret")
	}

	var r0 *secretsmanager.CancelRotateSecretOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *secretsmanager.CancelRotateSecretInput, ...func(*secretsmanager.Options)) (*secretsmanager.CancelRotateSecretOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *secretsmanager.CancelRotateSecretInput, ...func(*secretsmanager.Options)) *secretsmanager.CancelRotateSecretOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*secretsmanager.CancelRotateSecretOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *secretsmanager.CancelRotateSecretInput, ...func(*secretsmanager.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateSecret provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) CreateSecret(ctx context.Context, params *secretsmanager.CreateSecretInput, optFns ...func(*secretsmanager.Options)) (*secretsmanager.CreateSecretOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateSecret")
	}

	var r0 *secretsmanager.CreateSecretOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *secretsmanager.CreateSecretInput, ...func(*secretsmanager.Options)) (*secretsmanager.CreateSecretOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *secretsmanager.CreateSecretInput, ...func(*secretsmanager.Options)) *secretsmanager.CreateSecretOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*secretsmanager.CreateSecretOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *secretsmanager.CreateSecretInput, ...func(*secretsmanager.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteResourcePolicy provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DeleteResourcePolicy(ctx context.Context, params *secretsmanager.DeleteResourcePolicyInput, optFns ...func(*secretsmanager.Options)) (*secretsmanager.DeleteResourcePolicyOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteResourcePolicy")
	}

	var r0 *secretsmanager.DeleteResourcePolicyOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *secretsmanager.DeleteResourcePolicyInput, ...func(*secretsmanager.Options)) (*secretsmanager.DeleteResourcePolicyOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *secretsmanager.DeleteResourcePolicyInput, ...func(*secretsmanager.Options)) *secretsmanager.DeleteResourcePolicyOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*secretsmanager.DeleteResourcePolicyOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *secretsmanager.DeleteResourcePolicyInput, ...func(*secretsmanager.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteSecret provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DeleteSecret(ctx context.Context, params *secretsmanager.DeleteSecretInput, optFns ...func(*secretsmanager.Options)) (*secretsmanager.DeleteSecretOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteSecret")
	}

	var r0 *secretsmanager.DeleteSecretOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *secretsmanager.DeleteSecretInput, ...func(*secretsmanager.Options)) (*secretsmanager.DeleteSecretOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *secretsmanager.DeleteSecretInput, ...func(*secretsmanager.Options)) *secretsmanager.DeleteSecretOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*secretsmanager.DeleteSecretOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *secretsmanager.DeleteSecretInput, ...func(*secretsmanager.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeSecret provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DescribeSecret(ctx context.Context, params *secretsmanager.DescribeSecretInput, optFns ...func(*secretsmanager.Options)) (*secretsmanager.DescribeSecretOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DescribeSecret")
	}

	var r0 *secretsmanager.DescribeSecretOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *secretsmanager.DescribeSecretInput, ...func(*secretsmanager.Options)) (*secretsmanager.DescribeSecretOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *secretsmanager.DescribeSecretInput, ...func(*secretsmanager.Options)) *secretsmanager.DescribeSecretOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*secretsmanager.DescribeSecretOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *secretsmanager.DescribeSecretInput, ...func(*secretsmanager.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRandomPassword provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetRandomPassword(ctx context.Context, params *secretsmanager.GetRandomPasswordInput, optFns ...func(*secretsmanager.Options)) (*secretsmanager.GetRandomPasswordOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetRandomPassword")
	}

	var r0 *secretsmanager.GetRandomPasswordOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *secretsmanager.GetRandomPasswordInput, ...func(*secretsmanager.Options)) (*secretsmanager.GetRandomPasswordOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *secretsmanager.GetRandomPasswordInput, ...func(*secretsmanager.Options)) *secretsmanager.GetRandomPasswordOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*secretsmanager.GetRandomPasswordOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *secretsmanager.GetRandomPasswordInput, ...func(*secretsmanager.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetResourcePolicy provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetResourcePolicy(ctx context.Context, params *secretsmanager.GetResourcePolicyInput, optFns ...func(*secretsmanager.Options)) (*secretsmanager.GetResourcePolicyOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetResourcePolicy")
	}

	var r0 *secretsmanager.GetResourcePolicyOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *secretsmanager.GetResourcePolicyInput, ...func(*secretsmanager.Options)) (*secretsmanager.GetResourcePolicyOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *secretsmanager.GetResourcePolicyInput, ...func(*secretsmanager.Options)) *secretsmanager.GetResourcePolicyOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*secretsmanager.GetResourcePolicyOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *secretsmanager.GetResourcePolicyInput, ...func(*secretsmanager.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSecretValue provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetSecretValue(ctx context.Context, params *secretsmanager.GetSecretValueInput, optFns ...func(*secretsmanager.Options)) (*secretsmanager.GetSecretValueOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetSecretValue")
	}

	var r0 *secretsmanager.GetSecretValueOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *secretsmanager.GetSecretValueInput, ...func(*secretsmanager.Options)) (*secretsmanager.GetSecretValueOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *secretsmanager.GetSecretValueInput, ...func(*secretsmanager.Options)) *secretsmanager.GetSecretValueOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*secretsmanager.GetSecretValueOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *secretsmanager.GetSecretValueInput, ...func(*secretsmanager.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListSecretVersionIds provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListSecretVersionIds(ctx context.Context, params *secretsmanager.ListSecretVersionIdsInput, optFns ...func(*secretsmanager.Options)) (*secretsmanager.ListSecretVersionIdsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListSecretVersionIds")
	}

	var r0 *secretsmanager.ListSecretVersionIdsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *secretsmanager.ListSecretVersionIdsInput, ...func(*secretsmanager.Options)) (*secretsmanager.ListSecretVersionIdsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *secretsmanager.ListSecretVersionIdsInput, ...func(*secretsmanager.Options)) *secretsmanager.ListSecretVersionIdsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*secretsmanager.ListSecretVersionIdsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *secretsmanager.ListSecretVersionIdsInput, ...func(*secretsmanager.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListSecrets provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListSecrets(ctx context.Context, params *secretsmanager.ListSecretsInput, optFns ...func(*secretsmanager.Options)) (*secretsmanager.ListSecretsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListSecrets")
	}

	var r0 *secretsmanager.ListSecretsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *secretsmanager.ListSecretsInput, ...func(*secretsmanager.Options)) (*secretsmanager.ListSecretsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *secretsmanager.ListSecretsInput, ...func(*secretsmanager.Options)) *secretsmanager.ListSecretsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*secretsmanager.ListSecretsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *secretsmanager.ListSecretsInput, ...func(*secretsmanager.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Options provides a mock function with given fields:
func (_m *IClient) Options() secretsmanager.Options {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Options")
	}

	var r0 secretsmanager.Options
	if rf, ok := ret.Get(0).(func() secretsmanager.Options); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(secretsmanager.Options)
	}

	return r0
}

// PutResourcePolicy provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) PutResourcePolicy(ctx context.Context, params *secretsmanager.PutResourcePolicyInput, optFns ...func(*secretsmanager.Options)) (*secretsmanager.PutResourcePolicyOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for PutResourcePolicy")
	}

	var r0 *secretsmanager.PutResourcePolicyOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *secretsmanager.PutResourcePolicyInput, ...func(*secretsmanager.Options)) (*secretsmanager.PutResourcePolicyOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *secretsmanager.PutResourcePolicyInput, ...func(*secretsmanager.Options)) *secretsmanager.PutResourcePolicyOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*secretsmanager.PutResourcePolicyOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *secretsmanager.PutResourcePolicyInput, ...func(*secretsmanager.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PutSecretValue provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) PutSecretValue(ctx context.Context, params *secretsmanager.PutSecretValueInput, optFns ...func(*secretsmanager.Options)) (*secretsmanager.PutSecretValueOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for PutSecretValue")
	}

	var r0 *secretsmanager.PutSecretValueOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *secretsmanager.PutSecretValueInput, ...func(*secretsmanager.Options)) (*secretsmanager.PutSecretValueOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *secretsmanager.PutSecretValueInput, ...func(*secretsmanager.Options)) *secretsmanager.PutSecretValueOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*secretsmanager.PutSecretValueOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *secretsmanager.PutSecretValueInput, ...func(*secretsmanager.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RemoveRegionsFromReplication provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) RemoveRegionsFromReplication(ctx context.Context, params *secretsmanager.RemoveRegionsFromReplicationInput, optFns ...func(*secretsmanager.Options)) (*secretsmanager.RemoveRegionsFromReplicationOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for RemoveRegionsFromReplication")
	}

	var r0 *secretsmanager.RemoveRegionsFromReplicationOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *secretsmanager.RemoveRegionsFromReplicationInput, ...func(*secretsmanager.Options)) (*secretsmanager.RemoveRegionsFromReplicationOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *secretsmanager.RemoveRegionsFromReplicationInput, ...func(*secretsmanager.Options)) *secretsmanager.RemoveRegionsFromReplicationOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*secretsmanager.RemoveRegionsFromReplicationOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *secretsmanager.RemoveRegionsFromReplicationInput, ...func(*secretsmanager.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReplicateSecretToRegions provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ReplicateSecretToRegions(ctx context.Context, params *secretsmanager.ReplicateSecretToRegionsInput, optFns ...func(*secretsmanager.Options)) (*secretsmanager.ReplicateSecretToRegionsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ReplicateSecretToRegions")
	}

	var r0 *secretsmanager.ReplicateSecretToRegionsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *secretsmanager.ReplicateSecretToRegionsInput, ...func(*secretsmanager.Options)) (*secretsmanager.ReplicateSecretToRegionsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *secretsmanager.ReplicateSecretToRegionsInput, ...func(*secretsmanager.Options)) *secretsmanager.ReplicateSecretToRegionsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*secretsmanager.ReplicateSecretToRegionsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *secretsmanager.ReplicateSecretToRegionsInput, ...func(*secretsmanager.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RestoreSecret provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) RestoreSecret(ctx context.Context, params *secretsmanager.RestoreSecretInput, optFns ...func(*secretsmanager.Options)) (*secretsmanager.RestoreSecretOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for RestoreSecret")
	}

	var r0 *secretsmanager.RestoreSecretOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *secretsmanager.RestoreSecretInput, ...func(*secretsmanager.Options)) (*secretsmanager.RestoreSecretOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *secretsmanager.RestoreSecretInput, ...func(*secretsmanager.Options)) *secretsmanager.RestoreSecretOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*secretsmanager.RestoreSecretOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *secretsmanager.RestoreSecretInput, ...func(*secretsmanager.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RotateSecret provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) RotateSecret(ctx context.Context, params *secretsmanager.RotateSecretInput, optFns ...func(*secretsmanager.Options)) (*secretsmanager.RotateSecretOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for RotateSecret")
	}

	var r0 *secretsmanager.RotateSecretOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *secretsmanager.RotateSecretInput, ...func(*secretsmanager.Options)) (*secretsmanager.RotateSecretOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *secretsmanager.RotateSecretInput, ...func(*secretsmanager.Options)) *secretsmanager.RotateSecretOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*secretsmanager.RotateSecretOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *secretsmanager.RotateSecretInput, ...func(*secretsmanager.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StopReplicationToReplica provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) StopReplicationToReplica(ctx context.Context, params *secretsmanager.StopReplicationToReplicaInput, optFns ...func(*secretsmanager.Options)) (*secretsmanager.StopReplicationToReplicaOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for StopReplicationToReplica")
	}

	var r0 *secretsmanager.StopReplicationToReplicaOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *secretsmanager.StopReplicationToReplicaInput, ...func(*secretsmanager.Options)) (*secretsmanager.StopReplicationToReplicaOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *secretsmanager.StopReplicationToReplicaInput, ...func(*secretsmanager.Options)) *secretsmanager.StopReplicationToReplicaOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*secretsmanager.StopReplicationToReplicaOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *secretsmanager.StopReplicationToReplicaInput, ...func(*secretsmanager.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TagResource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) TagResource(ctx context.Context, params *secretsmanager.TagResourceInput, optFns ...func(*secretsmanager.Options)) (*secretsmanager.TagResourceOutput, error) {
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

	var r0 *secretsmanager.TagResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *secretsmanager.TagResourceInput, ...func(*secretsmanager.Options)) (*secretsmanager.TagResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *secretsmanager.TagResourceInput, ...func(*secretsmanager.Options)) *secretsmanager.TagResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*secretsmanager.TagResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *secretsmanager.TagResourceInput, ...func(*secretsmanager.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UntagResource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) UntagResource(ctx context.Context, params *secretsmanager.UntagResourceInput, optFns ...func(*secretsmanager.Options)) (*secretsmanager.UntagResourceOutput, error) {
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

	var r0 *secretsmanager.UntagResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *secretsmanager.UntagResourceInput, ...func(*secretsmanager.Options)) (*secretsmanager.UntagResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *secretsmanager.UntagResourceInput, ...func(*secretsmanager.Options)) *secretsmanager.UntagResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*secretsmanager.UntagResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *secretsmanager.UntagResourceInput, ...func(*secretsmanager.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateSecret provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) UpdateSecret(ctx context.Context, params *secretsmanager.UpdateSecretInput, optFns ...func(*secretsmanager.Options)) (*secretsmanager.UpdateSecretOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UpdateSecret")
	}

	var r0 *secretsmanager.UpdateSecretOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *secretsmanager.UpdateSecretInput, ...func(*secretsmanager.Options)) (*secretsmanager.UpdateSecretOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *secretsmanager.UpdateSecretInput, ...func(*secretsmanager.Options)) *secretsmanager.UpdateSecretOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*secretsmanager.UpdateSecretOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *secretsmanager.UpdateSecretInput, ...func(*secretsmanager.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateSecretVersionStage provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) UpdateSecretVersionStage(ctx context.Context, params *secretsmanager.UpdateSecretVersionStageInput, optFns ...func(*secretsmanager.Options)) (*secretsmanager.UpdateSecretVersionStageOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UpdateSecretVersionStage")
	}

	var r0 *secretsmanager.UpdateSecretVersionStageOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *secretsmanager.UpdateSecretVersionStageInput, ...func(*secretsmanager.Options)) (*secretsmanager.UpdateSecretVersionStageOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *secretsmanager.UpdateSecretVersionStageInput, ...func(*secretsmanager.Options)) *secretsmanager.UpdateSecretVersionStageOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*secretsmanager.UpdateSecretVersionStageOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *secretsmanager.UpdateSecretVersionStageInput, ...func(*secretsmanager.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ValidateResourcePolicy provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ValidateResourcePolicy(ctx context.Context, params *secretsmanager.ValidateResourcePolicyInput, optFns ...func(*secretsmanager.Options)) (*secretsmanager.ValidateResourcePolicyOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ValidateResourcePolicy")
	}

	var r0 *secretsmanager.ValidateResourcePolicyOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *secretsmanager.ValidateResourcePolicyInput, ...func(*secretsmanager.Options)) (*secretsmanager.ValidateResourcePolicyOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *secretsmanager.ValidateResourcePolicyInput, ...func(*secretsmanager.Options)) *secretsmanager.ValidateResourcePolicyOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*secretsmanager.ValidateResourcePolicyOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *secretsmanager.ValidateResourcePolicyInput, ...func(*secretsmanager.Options)) error); ok {
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
