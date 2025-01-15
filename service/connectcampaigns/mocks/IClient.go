// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	connectcampaigns "github.com/aws/aws-sdk-go-v2/service/connectcampaigns"

	context "context"

	mock "github.com/stretchr/testify/mock"
)

// IClient is an autogenerated mock type for the IClient type
type IClient struct {
	mock.Mock
}

// CreateCampaign provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) CreateCampaign(ctx context.Context, params *connectcampaigns.CreateCampaignInput, optFns ...func(*connectcampaigns.Options)) (*connectcampaigns.CreateCampaignOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateCampaign")
	}

	var r0 *connectcampaigns.CreateCampaignOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *connectcampaigns.CreateCampaignInput, ...func(*connectcampaigns.Options)) (*connectcampaigns.CreateCampaignOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *connectcampaigns.CreateCampaignInput, ...func(*connectcampaigns.Options)) *connectcampaigns.CreateCampaignOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*connectcampaigns.CreateCampaignOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *connectcampaigns.CreateCampaignInput, ...func(*connectcampaigns.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteCampaign provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DeleteCampaign(ctx context.Context, params *connectcampaigns.DeleteCampaignInput, optFns ...func(*connectcampaigns.Options)) (*connectcampaigns.DeleteCampaignOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteCampaign")
	}

	var r0 *connectcampaigns.DeleteCampaignOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *connectcampaigns.DeleteCampaignInput, ...func(*connectcampaigns.Options)) (*connectcampaigns.DeleteCampaignOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *connectcampaigns.DeleteCampaignInput, ...func(*connectcampaigns.Options)) *connectcampaigns.DeleteCampaignOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*connectcampaigns.DeleteCampaignOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *connectcampaigns.DeleteCampaignInput, ...func(*connectcampaigns.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteConnectInstanceConfig provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DeleteConnectInstanceConfig(ctx context.Context, params *connectcampaigns.DeleteConnectInstanceConfigInput, optFns ...func(*connectcampaigns.Options)) (*connectcampaigns.DeleteConnectInstanceConfigOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteConnectInstanceConfig")
	}

	var r0 *connectcampaigns.DeleteConnectInstanceConfigOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *connectcampaigns.DeleteConnectInstanceConfigInput, ...func(*connectcampaigns.Options)) (*connectcampaigns.DeleteConnectInstanceConfigOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *connectcampaigns.DeleteConnectInstanceConfigInput, ...func(*connectcampaigns.Options)) *connectcampaigns.DeleteConnectInstanceConfigOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*connectcampaigns.DeleteConnectInstanceConfigOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *connectcampaigns.DeleteConnectInstanceConfigInput, ...func(*connectcampaigns.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteInstanceOnboardingJob provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DeleteInstanceOnboardingJob(ctx context.Context, params *connectcampaigns.DeleteInstanceOnboardingJobInput, optFns ...func(*connectcampaigns.Options)) (*connectcampaigns.DeleteInstanceOnboardingJobOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteInstanceOnboardingJob")
	}

	var r0 *connectcampaigns.DeleteInstanceOnboardingJobOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *connectcampaigns.DeleteInstanceOnboardingJobInput, ...func(*connectcampaigns.Options)) (*connectcampaigns.DeleteInstanceOnboardingJobOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *connectcampaigns.DeleteInstanceOnboardingJobInput, ...func(*connectcampaigns.Options)) *connectcampaigns.DeleteInstanceOnboardingJobOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*connectcampaigns.DeleteInstanceOnboardingJobOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *connectcampaigns.DeleteInstanceOnboardingJobInput, ...func(*connectcampaigns.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeCampaign provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DescribeCampaign(ctx context.Context, params *connectcampaigns.DescribeCampaignInput, optFns ...func(*connectcampaigns.Options)) (*connectcampaigns.DescribeCampaignOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DescribeCampaign")
	}

	var r0 *connectcampaigns.DescribeCampaignOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *connectcampaigns.DescribeCampaignInput, ...func(*connectcampaigns.Options)) (*connectcampaigns.DescribeCampaignOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *connectcampaigns.DescribeCampaignInput, ...func(*connectcampaigns.Options)) *connectcampaigns.DescribeCampaignOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*connectcampaigns.DescribeCampaignOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *connectcampaigns.DescribeCampaignInput, ...func(*connectcampaigns.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCampaignState provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetCampaignState(ctx context.Context, params *connectcampaigns.GetCampaignStateInput, optFns ...func(*connectcampaigns.Options)) (*connectcampaigns.GetCampaignStateOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetCampaignState")
	}

	var r0 *connectcampaigns.GetCampaignStateOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *connectcampaigns.GetCampaignStateInput, ...func(*connectcampaigns.Options)) (*connectcampaigns.GetCampaignStateOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *connectcampaigns.GetCampaignStateInput, ...func(*connectcampaigns.Options)) *connectcampaigns.GetCampaignStateOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*connectcampaigns.GetCampaignStateOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *connectcampaigns.GetCampaignStateInput, ...func(*connectcampaigns.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCampaignStateBatch provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetCampaignStateBatch(ctx context.Context, params *connectcampaigns.GetCampaignStateBatchInput, optFns ...func(*connectcampaigns.Options)) (*connectcampaigns.GetCampaignStateBatchOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetCampaignStateBatch")
	}

	var r0 *connectcampaigns.GetCampaignStateBatchOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *connectcampaigns.GetCampaignStateBatchInput, ...func(*connectcampaigns.Options)) (*connectcampaigns.GetCampaignStateBatchOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *connectcampaigns.GetCampaignStateBatchInput, ...func(*connectcampaigns.Options)) *connectcampaigns.GetCampaignStateBatchOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*connectcampaigns.GetCampaignStateBatchOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *connectcampaigns.GetCampaignStateBatchInput, ...func(*connectcampaigns.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetConnectInstanceConfig provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetConnectInstanceConfig(ctx context.Context, params *connectcampaigns.GetConnectInstanceConfigInput, optFns ...func(*connectcampaigns.Options)) (*connectcampaigns.GetConnectInstanceConfigOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetConnectInstanceConfig")
	}

	var r0 *connectcampaigns.GetConnectInstanceConfigOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *connectcampaigns.GetConnectInstanceConfigInput, ...func(*connectcampaigns.Options)) (*connectcampaigns.GetConnectInstanceConfigOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *connectcampaigns.GetConnectInstanceConfigInput, ...func(*connectcampaigns.Options)) *connectcampaigns.GetConnectInstanceConfigOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*connectcampaigns.GetConnectInstanceConfigOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *connectcampaigns.GetConnectInstanceConfigInput, ...func(*connectcampaigns.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetInstanceOnboardingJobStatus provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetInstanceOnboardingJobStatus(ctx context.Context, params *connectcampaigns.GetInstanceOnboardingJobStatusInput, optFns ...func(*connectcampaigns.Options)) (*connectcampaigns.GetInstanceOnboardingJobStatusOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetInstanceOnboardingJobStatus")
	}

	var r0 *connectcampaigns.GetInstanceOnboardingJobStatusOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *connectcampaigns.GetInstanceOnboardingJobStatusInput, ...func(*connectcampaigns.Options)) (*connectcampaigns.GetInstanceOnboardingJobStatusOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *connectcampaigns.GetInstanceOnboardingJobStatusInput, ...func(*connectcampaigns.Options)) *connectcampaigns.GetInstanceOnboardingJobStatusOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*connectcampaigns.GetInstanceOnboardingJobStatusOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *connectcampaigns.GetInstanceOnboardingJobStatusInput, ...func(*connectcampaigns.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListCampaigns provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListCampaigns(ctx context.Context, params *connectcampaigns.ListCampaignsInput, optFns ...func(*connectcampaigns.Options)) (*connectcampaigns.ListCampaignsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListCampaigns")
	}

	var r0 *connectcampaigns.ListCampaignsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *connectcampaigns.ListCampaignsInput, ...func(*connectcampaigns.Options)) (*connectcampaigns.ListCampaignsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *connectcampaigns.ListCampaignsInput, ...func(*connectcampaigns.Options)) *connectcampaigns.ListCampaignsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*connectcampaigns.ListCampaignsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *connectcampaigns.ListCampaignsInput, ...func(*connectcampaigns.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListTagsForResource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListTagsForResource(ctx context.Context, params *connectcampaigns.ListTagsForResourceInput, optFns ...func(*connectcampaigns.Options)) (*connectcampaigns.ListTagsForResourceOutput, error) {
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

	var r0 *connectcampaigns.ListTagsForResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *connectcampaigns.ListTagsForResourceInput, ...func(*connectcampaigns.Options)) (*connectcampaigns.ListTagsForResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *connectcampaigns.ListTagsForResourceInput, ...func(*connectcampaigns.Options)) *connectcampaigns.ListTagsForResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*connectcampaigns.ListTagsForResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *connectcampaigns.ListTagsForResourceInput, ...func(*connectcampaigns.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Options provides a mock function with given fields:
func (_m *IClient) Options() connectcampaigns.Options {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Options")
	}

	var r0 connectcampaigns.Options
	if rf, ok := ret.Get(0).(func() connectcampaigns.Options); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(connectcampaigns.Options)
	}

	return r0
}

// PauseCampaign provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) PauseCampaign(ctx context.Context, params *connectcampaigns.PauseCampaignInput, optFns ...func(*connectcampaigns.Options)) (*connectcampaigns.PauseCampaignOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for PauseCampaign")
	}

	var r0 *connectcampaigns.PauseCampaignOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *connectcampaigns.PauseCampaignInput, ...func(*connectcampaigns.Options)) (*connectcampaigns.PauseCampaignOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *connectcampaigns.PauseCampaignInput, ...func(*connectcampaigns.Options)) *connectcampaigns.PauseCampaignOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*connectcampaigns.PauseCampaignOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *connectcampaigns.PauseCampaignInput, ...func(*connectcampaigns.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PutDialRequestBatch provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) PutDialRequestBatch(ctx context.Context, params *connectcampaigns.PutDialRequestBatchInput, optFns ...func(*connectcampaigns.Options)) (*connectcampaigns.PutDialRequestBatchOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for PutDialRequestBatch")
	}

	var r0 *connectcampaigns.PutDialRequestBatchOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *connectcampaigns.PutDialRequestBatchInput, ...func(*connectcampaigns.Options)) (*connectcampaigns.PutDialRequestBatchOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *connectcampaigns.PutDialRequestBatchInput, ...func(*connectcampaigns.Options)) *connectcampaigns.PutDialRequestBatchOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*connectcampaigns.PutDialRequestBatchOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *connectcampaigns.PutDialRequestBatchInput, ...func(*connectcampaigns.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ResumeCampaign provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ResumeCampaign(ctx context.Context, params *connectcampaigns.ResumeCampaignInput, optFns ...func(*connectcampaigns.Options)) (*connectcampaigns.ResumeCampaignOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ResumeCampaign")
	}

	var r0 *connectcampaigns.ResumeCampaignOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *connectcampaigns.ResumeCampaignInput, ...func(*connectcampaigns.Options)) (*connectcampaigns.ResumeCampaignOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *connectcampaigns.ResumeCampaignInput, ...func(*connectcampaigns.Options)) *connectcampaigns.ResumeCampaignOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*connectcampaigns.ResumeCampaignOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *connectcampaigns.ResumeCampaignInput, ...func(*connectcampaigns.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StartCampaign provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) StartCampaign(ctx context.Context, params *connectcampaigns.StartCampaignInput, optFns ...func(*connectcampaigns.Options)) (*connectcampaigns.StartCampaignOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for StartCampaign")
	}

	var r0 *connectcampaigns.StartCampaignOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *connectcampaigns.StartCampaignInput, ...func(*connectcampaigns.Options)) (*connectcampaigns.StartCampaignOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *connectcampaigns.StartCampaignInput, ...func(*connectcampaigns.Options)) *connectcampaigns.StartCampaignOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*connectcampaigns.StartCampaignOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *connectcampaigns.StartCampaignInput, ...func(*connectcampaigns.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StartInstanceOnboardingJob provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) StartInstanceOnboardingJob(ctx context.Context, params *connectcampaigns.StartInstanceOnboardingJobInput, optFns ...func(*connectcampaigns.Options)) (*connectcampaigns.StartInstanceOnboardingJobOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for StartInstanceOnboardingJob")
	}

	var r0 *connectcampaigns.StartInstanceOnboardingJobOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *connectcampaigns.StartInstanceOnboardingJobInput, ...func(*connectcampaigns.Options)) (*connectcampaigns.StartInstanceOnboardingJobOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *connectcampaigns.StartInstanceOnboardingJobInput, ...func(*connectcampaigns.Options)) *connectcampaigns.StartInstanceOnboardingJobOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*connectcampaigns.StartInstanceOnboardingJobOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *connectcampaigns.StartInstanceOnboardingJobInput, ...func(*connectcampaigns.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StopCampaign provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) StopCampaign(ctx context.Context, params *connectcampaigns.StopCampaignInput, optFns ...func(*connectcampaigns.Options)) (*connectcampaigns.StopCampaignOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for StopCampaign")
	}

	var r0 *connectcampaigns.StopCampaignOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *connectcampaigns.StopCampaignInput, ...func(*connectcampaigns.Options)) (*connectcampaigns.StopCampaignOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *connectcampaigns.StopCampaignInput, ...func(*connectcampaigns.Options)) *connectcampaigns.StopCampaignOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*connectcampaigns.StopCampaignOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *connectcampaigns.StopCampaignInput, ...func(*connectcampaigns.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TagResource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) TagResource(ctx context.Context, params *connectcampaigns.TagResourceInput, optFns ...func(*connectcampaigns.Options)) (*connectcampaigns.TagResourceOutput, error) {
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

	var r0 *connectcampaigns.TagResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *connectcampaigns.TagResourceInput, ...func(*connectcampaigns.Options)) (*connectcampaigns.TagResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *connectcampaigns.TagResourceInput, ...func(*connectcampaigns.Options)) *connectcampaigns.TagResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*connectcampaigns.TagResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *connectcampaigns.TagResourceInput, ...func(*connectcampaigns.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UntagResource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) UntagResource(ctx context.Context, params *connectcampaigns.UntagResourceInput, optFns ...func(*connectcampaigns.Options)) (*connectcampaigns.UntagResourceOutput, error) {
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

	var r0 *connectcampaigns.UntagResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *connectcampaigns.UntagResourceInput, ...func(*connectcampaigns.Options)) (*connectcampaigns.UntagResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *connectcampaigns.UntagResourceInput, ...func(*connectcampaigns.Options)) *connectcampaigns.UntagResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*connectcampaigns.UntagResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *connectcampaigns.UntagResourceInput, ...func(*connectcampaigns.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateCampaignDialerConfig provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) UpdateCampaignDialerConfig(ctx context.Context, params *connectcampaigns.UpdateCampaignDialerConfigInput, optFns ...func(*connectcampaigns.Options)) (*connectcampaigns.UpdateCampaignDialerConfigOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UpdateCampaignDialerConfig")
	}

	var r0 *connectcampaigns.UpdateCampaignDialerConfigOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *connectcampaigns.UpdateCampaignDialerConfigInput, ...func(*connectcampaigns.Options)) (*connectcampaigns.UpdateCampaignDialerConfigOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *connectcampaigns.UpdateCampaignDialerConfigInput, ...func(*connectcampaigns.Options)) *connectcampaigns.UpdateCampaignDialerConfigOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*connectcampaigns.UpdateCampaignDialerConfigOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *connectcampaigns.UpdateCampaignDialerConfigInput, ...func(*connectcampaigns.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateCampaignName provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) UpdateCampaignName(ctx context.Context, params *connectcampaigns.UpdateCampaignNameInput, optFns ...func(*connectcampaigns.Options)) (*connectcampaigns.UpdateCampaignNameOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UpdateCampaignName")
	}

	var r0 *connectcampaigns.UpdateCampaignNameOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *connectcampaigns.UpdateCampaignNameInput, ...func(*connectcampaigns.Options)) (*connectcampaigns.UpdateCampaignNameOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *connectcampaigns.UpdateCampaignNameInput, ...func(*connectcampaigns.Options)) *connectcampaigns.UpdateCampaignNameOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*connectcampaigns.UpdateCampaignNameOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *connectcampaigns.UpdateCampaignNameInput, ...func(*connectcampaigns.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateCampaignOutboundCallConfig provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) UpdateCampaignOutboundCallConfig(ctx context.Context, params *connectcampaigns.UpdateCampaignOutboundCallConfigInput, optFns ...func(*connectcampaigns.Options)) (*connectcampaigns.UpdateCampaignOutboundCallConfigOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UpdateCampaignOutboundCallConfig")
	}

	var r0 *connectcampaigns.UpdateCampaignOutboundCallConfigOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *connectcampaigns.UpdateCampaignOutboundCallConfigInput, ...func(*connectcampaigns.Options)) (*connectcampaigns.UpdateCampaignOutboundCallConfigOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *connectcampaigns.UpdateCampaignOutboundCallConfigInput, ...func(*connectcampaigns.Options)) *connectcampaigns.UpdateCampaignOutboundCallConfigOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*connectcampaigns.UpdateCampaignOutboundCallConfigOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *connectcampaigns.UpdateCampaignOutboundCallConfigInput, ...func(*connectcampaigns.Options)) error); ok {
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
