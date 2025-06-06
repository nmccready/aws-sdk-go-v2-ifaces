// Code generated by mockery v2.53.4. DO NOT EDIT.

package mocks

import (
	context "context"

	snowball "github.com/aws/aws-sdk-go-v2/service/snowball"
	mock "github.com/stretchr/testify/mock"
)

// IClient is an autogenerated mock type for the IClient type
type IClient struct {
	mock.Mock
}

// CancelCluster provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) CancelCluster(ctx context.Context, params *snowball.CancelClusterInput, optFns ...func(*snowball.Options)) (*snowball.CancelClusterOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CancelCluster")
	}

	var r0 *snowball.CancelClusterOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *snowball.CancelClusterInput, ...func(*snowball.Options)) (*snowball.CancelClusterOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *snowball.CancelClusterInput, ...func(*snowball.Options)) *snowball.CancelClusterOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*snowball.CancelClusterOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *snowball.CancelClusterInput, ...func(*snowball.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CancelJob provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) CancelJob(ctx context.Context, params *snowball.CancelJobInput, optFns ...func(*snowball.Options)) (*snowball.CancelJobOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CancelJob")
	}

	var r0 *snowball.CancelJobOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *snowball.CancelJobInput, ...func(*snowball.Options)) (*snowball.CancelJobOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *snowball.CancelJobInput, ...func(*snowball.Options)) *snowball.CancelJobOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*snowball.CancelJobOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *snowball.CancelJobInput, ...func(*snowball.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateAddress provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) CreateAddress(ctx context.Context, params *snowball.CreateAddressInput, optFns ...func(*snowball.Options)) (*snowball.CreateAddressOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateAddress")
	}

	var r0 *snowball.CreateAddressOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *snowball.CreateAddressInput, ...func(*snowball.Options)) (*snowball.CreateAddressOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *snowball.CreateAddressInput, ...func(*snowball.Options)) *snowball.CreateAddressOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*snowball.CreateAddressOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *snowball.CreateAddressInput, ...func(*snowball.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateCluster provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) CreateCluster(ctx context.Context, params *snowball.CreateClusterInput, optFns ...func(*snowball.Options)) (*snowball.CreateClusterOutput, error) {
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

	var r0 *snowball.CreateClusterOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *snowball.CreateClusterInput, ...func(*snowball.Options)) (*snowball.CreateClusterOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *snowball.CreateClusterInput, ...func(*snowball.Options)) *snowball.CreateClusterOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*snowball.CreateClusterOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *snowball.CreateClusterInput, ...func(*snowball.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateJob provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) CreateJob(ctx context.Context, params *snowball.CreateJobInput, optFns ...func(*snowball.Options)) (*snowball.CreateJobOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateJob")
	}

	var r0 *snowball.CreateJobOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *snowball.CreateJobInput, ...func(*snowball.Options)) (*snowball.CreateJobOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *snowball.CreateJobInput, ...func(*snowball.Options)) *snowball.CreateJobOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*snowball.CreateJobOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *snowball.CreateJobInput, ...func(*snowball.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateLongTermPricing provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) CreateLongTermPricing(ctx context.Context, params *snowball.CreateLongTermPricingInput, optFns ...func(*snowball.Options)) (*snowball.CreateLongTermPricingOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateLongTermPricing")
	}

	var r0 *snowball.CreateLongTermPricingOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *snowball.CreateLongTermPricingInput, ...func(*snowball.Options)) (*snowball.CreateLongTermPricingOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *snowball.CreateLongTermPricingInput, ...func(*snowball.Options)) *snowball.CreateLongTermPricingOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*snowball.CreateLongTermPricingOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *snowball.CreateLongTermPricingInput, ...func(*snowball.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateReturnShippingLabel provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) CreateReturnShippingLabel(ctx context.Context, params *snowball.CreateReturnShippingLabelInput, optFns ...func(*snowball.Options)) (*snowball.CreateReturnShippingLabelOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateReturnShippingLabel")
	}

	var r0 *snowball.CreateReturnShippingLabelOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *snowball.CreateReturnShippingLabelInput, ...func(*snowball.Options)) (*snowball.CreateReturnShippingLabelOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *snowball.CreateReturnShippingLabelInput, ...func(*snowball.Options)) *snowball.CreateReturnShippingLabelOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*snowball.CreateReturnShippingLabelOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *snowball.CreateReturnShippingLabelInput, ...func(*snowball.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeAddress provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DescribeAddress(ctx context.Context, params *snowball.DescribeAddressInput, optFns ...func(*snowball.Options)) (*snowball.DescribeAddressOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DescribeAddress")
	}

	var r0 *snowball.DescribeAddressOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *snowball.DescribeAddressInput, ...func(*snowball.Options)) (*snowball.DescribeAddressOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *snowball.DescribeAddressInput, ...func(*snowball.Options)) *snowball.DescribeAddressOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*snowball.DescribeAddressOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *snowball.DescribeAddressInput, ...func(*snowball.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeAddresses provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DescribeAddresses(ctx context.Context, params *snowball.DescribeAddressesInput, optFns ...func(*snowball.Options)) (*snowball.DescribeAddressesOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DescribeAddresses")
	}

	var r0 *snowball.DescribeAddressesOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *snowball.DescribeAddressesInput, ...func(*snowball.Options)) (*snowball.DescribeAddressesOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *snowball.DescribeAddressesInput, ...func(*snowball.Options)) *snowball.DescribeAddressesOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*snowball.DescribeAddressesOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *snowball.DescribeAddressesInput, ...func(*snowball.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeCluster provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DescribeCluster(ctx context.Context, params *snowball.DescribeClusterInput, optFns ...func(*snowball.Options)) (*snowball.DescribeClusterOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DescribeCluster")
	}

	var r0 *snowball.DescribeClusterOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *snowball.DescribeClusterInput, ...func(*snowball.Options)) (*snowball.DescribeClusterOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *snowball.DescribeClusterInput, ...func(*snowball.Options)) *snowball.DescribeClusterOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*snowball.DescribeClusterOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *snowball.DescribeClusterInput, ...func(*snowball.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeJob provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DescribeJob(ctx context.Context, params *snowball.DescribeJobInput, optFns ...func(*snowball.Options)) (*snowball.DescribeJobOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DescribeJob")
	}

	var r0 *snowball.DescribeJobOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *snowball.DescribeJobInput, ...func(*snowball.Options)) (*snowball.DescribeJobOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *snowball.DescribeJobInput, ...func(*snowball.Options)) *snowball.DescribeJobOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*snowball.DescribeJobOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *snowball.DescribeJobInput, ...func(*snowball.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeReturnShippingLabel provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DescribeReturnShippingLabel(ctx context.Context, params *snowball.DescribeReturnShippingLabelInput, optFns ...func(*snowball.Options)) (*snowball.DescribeReturnShippingLabelOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DescribeReturnShippingLabel")
	}

	var r0 *snowball.DescribeReturnShippingLabelOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *snowball.DescribeReturnShippingLabelInput, ...func(*snowball.Options)) (*snowball.DescribeReturnShippingLabelOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *snowball.DescribeReturnShippingLabelInput, ...func(*snowball.Options)) *snowball.DescribeReturnShippingLabelOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*snowball.DescribeReturnShippingLabelOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *snowball.DescribeReturnShippingLabelInput, ...func(*snowball.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetJobManifest provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetJobManifest(ctx context.Context, params *snowball.GetJobManifestInput, optFns ...func(*snowball.Options)) (*snowball.GetJobManifestOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetJobManifest")
	}

	var r0 *snowball.GetJobManifestOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *snowball.GetJobManifestInput, ...func(*snowball.Options)) (*snowball.GetJobManifestOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *snowball.GetJobManifestInput, ...func(*snowball.Options)) *snowball.GetJobManifestOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*snowball.GetJobManifestOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *snowball.GetJobManifestInput, ...func(*snowball.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetJobUnlockCode provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetJobUnlockCode(ctx context.Context, params *snowball.GetJobUnlockCodeInput, optFns ...func(*snowball.Options)) (*snowball.GetJobUnlockCodeOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetJobUnlockCode")
	}

	var r0 *snowball.GetJobUnlockCodeOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *snowball.GetJobUnlockCodeInput, ...func(*snowball.Options)) (*snowball.GetJobUnlockCodeOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *snowball.GetJobUnlockCodeInput, ...func(*snowball.Options)) *snowball.GetJobUnlockCodeOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*snowball.GetJobUnlockCodeOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *snowball.GetJobUnlockCodeInput, ...func(*snowball.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSnowballUsage provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetSnowballUsage(ctx context.Context, params *snowball.GetSnowballUsageInput, optFns ...func(*snowball.Options)) (*snowball.GetSnowballUsageOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetSnowballUsage")
	}

	var r0 *snowball.GetSnowballUsageOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *snowball.GetSnowballUsageInput, ...func(*snowball.Options)) (*snowball.GetSnowballUsageOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *snowball.GetSnowballUsageInput, ...func(*snowball.Options)) *snowball.GetSnowballUsageOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*snowball.GetSnowballUsageOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *snowball.GetSnowballUsageInput, ...func(*snowball.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSoftwareUpdates provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetSoftwareUpdates(ctx context.Context, params *snowball.GetSoftwareUpdatesInput, optFns ...func(*snowball.Options)) (*snowball.GetSoftwareUpdatesOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetSoftwareUpdates")
	}

	var r0 *snowball.GetSoftwareUpdatesOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *snowball.GetSoftwareUpdatesInput, ...func(*snowball.Options)) (*snowball.GetSoftwareUpdatesOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *snowball.GetSoftwareUpdatesInput, ...func(*snowball.Options)) *snowball.GetSoftwareUpdatesOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*snowball.GetSoftwareUpdatesOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *snowball.GetSoftwareUpdatesInput, ...func(*snowball.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListClusterJobs provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListClusterJobs(ctx context.Context, params *snowball.ListClusterJobsInput, optFns ...func(*snowball.Options)) (*snowball.ListClusterJobsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListClusterJobs")
	}

	var r0 *snowball.ListClusterJobsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *snowball.ListClusterJobsInput, ...func(*snowball.Options)) (*snowball.ListClusterJobsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *snowball.ListClusterJobsInput, ...func(*snowball.Options)) *snowball.ListClusterJobsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*snowball.ListClusterJobsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *snowball.ListClusterJobsInput, ...func(*snowball.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListClusters provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListClusters(ctx context.Context, params *snowball.ListClustersInput, optFns ...func(*snowball.Options)) (*snowball.ListClustersOutput, error) {
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

	var r0 *snowball.ListClustersOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *snowball.ListClustersInput, ...func(*snowball.Options)) (*snowball.ListClustersOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *snowball.ListClustersInput, ...func(*snowball.Options)) *snowball.ListClustersOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*snowball.ListClustersOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *snowball.ListClustersInput, ...func(*snowball.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListCompatibleImages provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListCompatibleImages(ctx context.Context, params *snowball.ListCompatibleImagesInput, optFns ...func(*snowball.Options)) (*snowball.ListCompatibleImagesOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListCompatibleImages")
	}

	var r0 *snowball.ListCompatibleImagesOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *snowball.ListCompatibleImagesInput, ...func(*snowball.Options)) (*snowball.ListCompatibleImagesOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *snowball.ListCompatibleImagesInput, ...func(*snowball.Options)) *snowball.ListCompatibleImagesOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*snowball.ListCompatibleImagesOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *snowball.ListCompatibleImagesInput, ...func(*snowball.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListJobs provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListJobs(ctx context.Context, params *snowball.ListJobsInput, optFns ...func(*snowball.Options)) (*snowball.ListJobsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListJobs")
	}

	var r0 *snowball.ListJobsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *snowball.ListJobsInput, ...func(*snowball.Options)) (*snowball.ListJobsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *snowball.ListJobsInput, ...func(*snowball.Options)) *snowball.ListJobsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*snowball.ListJobsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *snowball.ListJobsInput, ...func(*snowball.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListLongTermPricing provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListLongTermPricing(ctx context.Context, params *snowball.ListLongTermPricingInput, optFns ...func(*snowball.Options)) (*snowball.ListLongTermPricingOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListLongTermPricing")
	}

	var r0 *snowball.ListLongTermPricingOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *snowball.ListLongTermPricingInput, ...func(*snowball.Options)) (*snowball.ListLongTermPricingOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *snowball.ListLongTermPricingInput, ...func(*snowball.Options)) *snowball.ListLongTermPricingOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*snowball.ListLongTermPricingOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *snowball.ListLongTermPricingInput, ...func(*snowball.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListPickupLocations provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListPickupLocations(ctx context.Context, params *snowball.ListPickupLocationsInput, optFns ...func(*snowball.Options)) (*snowball.ListPickupLocationsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListPickupLocations")
	}

	var r0 *snowball.ListPickupLocationsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *snowball.ListPickupLocationsInput, ...func(*snowball.Options)) (*snowball.ListPickupLocationsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *snowball.ListPickupLocationsInput, ...func(*snowball.Options)) *snowball.ListPickupLocationsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*snowball.ListPickupLocationsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *snowball.ListPickupLocationsInput, ...func(*snowball.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListServiceVersions provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListServiceVersions(ctx context.Context, params *snowball.ListServiceVersionsInput, optFns ...func(*snowball.Options)) (*snowball.ListServiceVersionsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListServiceVersions")
	}

	var r0 *snowball.ListServiceVersionsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *snowball.ListServiceVersionsInput, ...func(*snowball.Options)) (*snowball.ListServiceVersionsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *snowball.ListServiceVersionsInput, ...func(*snowball.Options)) *snowball.ListServiceVersionsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*snowball.ListServiceVersionsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *snowball.ListServiceVersionsInput, ...func(*snowball.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Options provides a mock function with no fields
func (_m *IClient) Options() snowball.Options {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Options")
	}

	var r0 snowball.Options
	if rf, ok := ret.Get(0).(func() snowball.Options); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(snowball.Options)
	}

	return r0
}

// UpdateCluster provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) UpdateCluster(ctx context.Context, params *snowball.UpdateClusterInput, optFns ...func(*snowball.Options)) (*snowball.UpdateClusterOutput, error) {
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

	var r0 *snowball.UpdateClusterOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *snowball.UpdateClusterInput, ...func(*snowball.Options)) (*snowball.UpdateClusterOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *snowball.UpdateClusterInput, ...func(*snowball.Options)) *snowball.UpdateClusterOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*snowball.UpdateClusterOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *snowball.UpdateClusterInput, ...func(*snowball.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateJob provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) UpdateJob(ctx context.Context, params *snowball.UpdateJobInput, optFns ...func(*snowball.Options)) (*snowball.UpdateJobOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UpdateJob")
	}

	var r0 *snowball.UpdateJobOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *snowball.UpdateJobInput, ...func(*snowball.Options)) (*snowball.UpdateJobOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *snowball.UpdateJobInput, ...func(*snowball.Options)) *snowball.UpdateJobOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*snowball.UpdateJobOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *snowball.UpdateJobInput, ...func(*snowball.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateJobShipmentState provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) UpdateJobShipmentState(ctx context.Context, params *snowball.UpdateJobShipmentStateInput, optFns ...func(*snowball.Options)) (*snowball.UpdateJobShipmentStateOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UpdateJobShipmentState")
	}

	var r0 *snowball.UpdateJobShipmentStateOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *snowball.UpdateJobShipmentStateInput, ...func(*snowball.Options)) (*snowball.UpdateJobShipmentStateOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *snowball.UpdateJobShipmentStateInput, ...func(*snowball.Options)) *snowball.UpdateJobShipmentStateOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*snowball.UpdateJobShipmentStateOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *snowball.UpdateJobShipmentStateInput, ...func(*snowball.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateLongTermPricing provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) UpdateLongTermPricing(ctx context.Context, params *snowball.UpdateLongTermPricingInput, optFns ...func(*snowball.Options)) (*snowball.UpdateLongTermPricingOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UpdateLongTermPricing")
	}

	var r0 *snowball.UpdateLongTermPricingOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *snowball.UpdateLongTermPricingInput, ...func(*snowball.Options)) (*snowball.UpdateLongTermPricingOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *snowball.UpdateLongTermPricingInput, ...func(*snowball.Options)) *snowball.UpdateLongTermPricingOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*snowball.UpdateLongTermPricingOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *snowball.UpdateLongTermPricingInput, ...func(*snowball.Options)) error); ok {
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
