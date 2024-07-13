// Code generated by mockery v2.39.2. DO NOT EDIT.

package mocks

import (
	context "context"

	iotevents "github.com/aws/aws-sdk-go-v2/service/iotevents"

	mock "github.com/stretchr/testify/mock"
)

// IClient is an autogenerated mock type for the IClient type
type IClient struct {
	mock.Mock
}

// CreateAlarmModel provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) CreateAlarmModel(ctx context.Context, params *iotevents.CreateAlarmModelInput, optFns ...func(*iotevents.Options)) (*iotevents.CreateAlarmModelOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateAlarmModel")
	}

	var r0 *iotevents.CreateAlarmModelOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *iotevents.CreateAlarmModelInput, ...func(*iotevents.Options)) (*iotevents.CreateAlarmModelOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *iotevents.CreateAlarmModelInput, ...func(*iotevents.Options)) *iotevents.CreateAlarmModelOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*iotevents.CreateAlarmModelOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *iotevents.CreateAlarmModelInput, ...func(*iotevents.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateDetectorModel provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) CreateDetectorModel(ctx context.Context, params *iotevents.CreateDetectorModelInput, optFns ...func(*iotevents.Options)) (*iotevents.CreateDetectorModelOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateDetectorModel")
	}

	var r0 *iotevents.CreateDetectorModelOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *iotevents.CreateDetectorModelInput, ...func(*iotevents.Options)) (*iotevents.CreateDetectorModelOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *iotevents.CreateDetectorModelInput, ...func(*iotevents.Options)) *iotevents.CreateDetectorModelOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*iotevents.CreateDetectorModelOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *iotevents.CreateDetectorModelInput, ...func(*iotevents.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateInput provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) CreateInput(ctx context.Context, params *iotevents.CreateInputInput, optFns ...func(*iotevents.Options)) (*iotevents.CreateInputOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateInput")
	}

	var r0 *iotevents.CreateInputOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *iotevents.CreateInputInput, ...func(*iotevents.Options)) (*iotevents.CreateInputOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *iotevents.CreateInputInput, ...func(*iotevents.Options)) *iotevents.CreateInputOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*iotevents.CreateInputOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *iotevents.CreateInputInput, ...func(*iotevents.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteAlarmModel provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DeleteAlarmModel(ctx context.Context, params *iotevents.DeleteAlarmModelInput, optFns ...func(*iotevents.Options)) (*iotevents.DeleteAlarmModelOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteAlarmModel")
	}

	var r0 *iotevents.DeleteAlarmModelOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *iotevents.DeleteAlarmModelInput, ...func(*iotevents.Options)) (*iotevents.DeleteAlarmModelOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *iotevents.DeleteAlarmModelInput, ...func(*iotevents.Options)) *iotevents.DeleteAlarmModelOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*iotevents.DeleteAlarmModelOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *iotevents.DeleteAlarmModelInput, ...func(*iotevents.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteDetectorModel provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DeleteDetectorModel(ctx context.Context, params *iotevents.DeleteDetectorModelInput, optFns ...func(*iotevents.Options)) (*iotevents.DeleteDetectorModelOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteDetectorModel")
	}

	var r0 *iotevents.DeleteDetectorModelOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *iotevents.DeleteDetectorModelInput, ...func(*iotevents.Options)) (*iotevents.DeleteDetectorModelOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *iotevents.DeleteDetectorModelInput, ...func(*iotevents.Options)) *iotevents.DeleteDetectorModelOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*iotevents.DeleteDetectorModelOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *iotevents.DeleteDetectorModelInput, ...func(*iotevents.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteInput provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DeleteInput(ctx context.Context, params *iotevents.DeleteInputInput, optFns ...func(*iotevents.Options)) (*iotevents.DeleteInputOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteInput")
	}

	var r0 *iotevents.DeleteInputOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *iotevents.DeleteInputInput, ...func(*iotevents.Options)) (*iotevents.DeleteInputOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *iotevents.DeleteInputInput, ...func(*iotevents.Options)) *iotevents.DeleteInputOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*iotevents.DeleteInputOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *iotevents.DeleteInputInput, ...func(*iotevents.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeAlarmModel provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DescribeAlarmModel(ctx context.Context, params *iotevents.DescribeAlarmModelInput, optFns ...func(*iotevents.Options)) (*iotevents.DescribeAlarmModelOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DescribeAlarmModel")
	}

	var r0 *iotevents.DescribeAlarmModelOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *iotevents.DescribeAlarmModelInput, ...func(*iotevents.Options)) (*iotevents.DescribeAlarmModelOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *iotevents.DescribeAlarmModelInput, ...func(*iotevents.Options)) *iotevents.DescribeAlarmModelOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*iotevents.DescribeAlarmModelOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *iotevents.DescribeAlarmModelInput, ...func(*iotevents.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeDetectorModel provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DescribeDetectorModel(ctx context.Context, params *iotevents.DescribeDetectorModelInput, optFns ...func(*iotevents.Options)) (*iotevents.DescribeDetectorModelOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DescribeDetectorModel")
	}

	var r0 *iotevents.DescribeDetectorModelOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *iotevents.DescribeDetectorModelInput, ...func(*iotevents.Options)) (*iotevents.DescribeDetectorModelOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *iotevents.DescribeDetectorModelInput, ...func(*iotevents.Options)) *iotevents.DescribeDetectorModelOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*iotevents.DescribeDetectorModelOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *iotevents.DescribeDetectorModelInput, ...func(*iotevents.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeDetectorModelAnalysis provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DescribeDetectorModelAnalysis(ctx context.Context, params *iotevents.DescribeDetectorModelAnalysisInput, optFns ...func(*iotevents.Options)) (*iotevents.DescribeDetectorModelAnalysisOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DescribeDetectorModelAnalysis")
	}

	var r0 *iotevents.DescribeDetectorModelAnalysisOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *iotevents.DescribeDetectorModelAnalysisInput, ...func(*iotevents.Options)) (*iotevents.DescribeDetectorModelAnalysisOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *iotevents.DescribeDetectorModelAnalysisInput, ...func(*iotevents.Options)) *iotevents.DescribeDetectorModelAnalysisOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*iotevents.DescribeDetectorModelAnalysisOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *iotevents.DescribeDetectorModelAnalysisInput, ...func(*iotevents.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeInput provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DescribeInput(ctx context.Context, params *iotevents.DescribeInputInput, optFns ...func(*iotevents.Options)) (*iotevents.DescribeInputOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DescribeInput")
	}

	var r0 *iotevents.DescribeInputOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *iotevents.DescribeInputInput, ...func(*iotevents.Options)) (*iotevents.DescribeInputOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *iotevents.DescribeInputInput, ...func(*iotevents.Options)) *iotevents.DescribeInputOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*iotevents.DescribeInputOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *iotevents.DescribeInputInput, ...func(*iotevents.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeLoggingOptions provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DescribeLoggingOptions(ctx context.Context, params *iotevents.DescribeLoggingOptionsInput, optFns ...func(*iotevents.Options)) (*iotevents.DescribeLoggingOptionsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DescribeLoggingOptions")
	}

	var r0 *iotevents.DescribeLoggingOptionsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *iotevents.DescribeLoggingOptionsInput, ...func(*iotevents.Options)) (*iotevents.DescribeLoggingOptionsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *iotevents.DescribeLoggingOptionsInput, ...func(*iotevents.Options)) *iotevents.DescribeLoggingOptionsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*iotevents.DescribeLoggingOptionsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *iotevents.DescribeLoggingOptionsInput, ...func(*iotevents.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDetectorModelAnalysisResults provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetDetectorModelAnalysisResults(ctx context.Context, params *iotevents.GetDetectorModelAnalysisResultsInput, optFns ...func(*iotevents.Options)) (*iotevents.GetDetectorModelAnalysisResultsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetDetectorModelAnalysisResults")
	}

	var r0 *iotevents.GetDetectorModelAnalysisResultsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *iotevents.GetDetectorModelAnalysisResultsInput, ...func(*iotevents.Options)) (*iotevents.GetDetectorModelAnalysisResultsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *iotevents.GetDetectorModelAnalysisResultsInput, ...func(*iotevents.Options)) *iotevents.GetDetectorModelAnalysisResultsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*iotevents.GetDetectorModelAnalysisResultsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *iotevents.GetDetectorModelAnalysisResultsInput, ...func(*iotevents.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListAlarmModelVersions provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListAlarmModelVersions(ctx context.Context, params *iotevents.ListAlarmModelVersionsInput, optFns ...func(*iotevents.Options)) (*iotevents.ListAlarmModelVersionsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListAlarmModelVersions")
	}

	var r0 *iotevents.ListAlarmModelVersionsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *iotevents.ListAlarmModelVersionsInput, ...func(*iotevents.Options)) (*iotevents.ListAlarmModelVersionsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *iotevents.ListAlarmModelVersionsInput, ...func(*iotevents.Options)) *iotevents.ListAlarmModelVersionsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*iotevents.ListAlarmModelVersionsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *iotevents.ListAlarmModelVersionsInput, ...func(*iotevents.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListAlarmModels provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListAlarmModels(ctx context.Context, params *iotevents.ListAlarmModelsInput, optFns ...func(*iotevents.Options)) (*iotevents.ListAlarmModelsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListAlarmModels")
	}

	var r0 *iotevents.ListAlarmModelsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *iotevents.ListAlarmModelsInput, ...func(*iotevents.Options)) (*iotevents.ListAlarmModelsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *iotevents.ListAlarmModelsInput, ...func(*iotevents.Options)) *iotevents.ListAlarmModelsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*iotevents.ListAlarmModelsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *iotevents.ListAlarmModelsInput, ...func(*iotevents.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListDetectorModelVersions provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListDetectorModelVersions(ctx context.Context, params *iotevents.ListDetectorModelVersionsInput, optFns ...func(*iotevents.Options)) (*iotevents.ListDetectorModelVersionsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListDetectorModelVersions")
	}

	var r0 *iotevents.ListDetectorModelVersionsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *iotevents.ListDetectorModelVersionsInput, ...func(*iotevents.Options)) (*iotevents.ListDetectorModelVersionsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *iotevents.ListDetectorModelVersionsInput, ...func(*iotevents.Options)) *iotevents.ListDetectorModelVersionsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*iotevents.ListDetectorModelVersionsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *iotevents.ListDetectorModelVersionsInput, ...func(*iotevents.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListDetectorModels provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListDetectorModels(ctx context.Context, params *iotevents.ListDetectorModelsInput, optFns ...func(*iotevents.Options)) (*iotevents.ListDetectorModelsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListDetectorModels")
	}

	var r0 *iotevents.ListDetectorModelsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *iotevents.ListDetectorModelsInput, ...func(*iotevents.Options)) (*iotevents.ListDetectorModelsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *iotevents.ListDetectorModelsInput, ...func(*iotevents.Options)) *iotevents.ListDetectorModelsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*iotevents.ListDetectorModelsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *iotevents.ListDetectorModelsInput, ...func(*iotevents.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListInputRoutings provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListInputRoutings(ctx context.Context, params *iotevents.ListInputRoutingsInput, optFns ...func(*iotevents.Options)) (*iotevents.ListInputRoutingsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListInputRoutings")
	}

	var r0 *iotevents.ListInputRoutingsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *iotevents.ListInputRoutingsInput, ...func(*iotevents.Options)) (*iotevents.ListInputRoutingsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *iotevents.ListInputRoutingsInput, ...func(*iotevents.Options)) *iotevents.ListInputRoutingsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*iotevents.ListInputRoutingsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *iotevents.ListInputRoutingsInput, ...func(*iotevents.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListInputs provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListInputs(ctx context.Context, params *iotevents.ListInputsInput, optFns ...func(*iotevents.Options)) (*iotevents.ListInputsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListInputs")
	}

	var r0 *iotevents.ListInputsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *iotevents.ListInputsInput, ...func(*iotevents.Options)) (*iotevents.ListInputsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *iotevents.ListInputsInput, ...func(*iotevents.Options)) *iotevents.ListInputsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*iotevents.ListInputsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *iotevents.ListInputsInput, ...func(*iotevents.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListTagsForResource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListTagsForResource(ctx context.Context, params *iotevents.ListTagsForResourceInput, optFns ...func(*iotevents.Options)) (*iotevents.ListTagsForResourceOutput, error) {
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

	var r0 *iotevents.ListTagsForResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *iotevents.ListTagsForResourceInput, ...func(*iotevents.Options)) (*iotevents.ListTagsForResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *iotevents.ListTagsForResourceInput, ...func(*iotevents.Options)) *iotevents.ListTagsForResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*iotevents.ListTagsForResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *iotevents.ListTagsForResourceInput, ...func(*iotevents.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Options provides a mock function with given fields:
func (_m *IClient) Options() iotevents.Options {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Options")
	}

	var r0 iotevents.Options
	if rf, ok := ret.Get(0).(func() iotevents.Options); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(iotevents.Options)
	}

	return r0
}

// PutLoggingOptions provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) PutLoggingOptions(ctx context.Context, params *iotevents.PutLoggingOptionsInput, optFns ...func(*iotevents.Options)) (*iotevents.PutLoggingOptionsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for PutLoggingOptions")
	}

	var r0 *iotevents.PutLoggingOptionsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *iotevents.PutLoggingOptionsInput, ...func(*iotevents.Options)) (*iotevents.PutLoggingOptionsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *iotevents.PutLoggingOptionsInput, ...func(*iotevents.Options)) *iotevents.PutLoggingOptionsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*iotevents.PutLoggingOptionsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *iotevents.PutLoggingOptionsInput, ...func(*iotevents.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StartDetectorModelAnalysis provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) StartDetectorModelAnalysis(ctx context.Context, params *iotevents.StartDetectorModelAnalysisInput, optFns ...func(*iotevents.Options)) (*iotevents.StartDetectorModelAnalysisOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for StartDetectorModelAnalysis")
	}

	var r0 *iotevents.StartDetectorModelAnalysisOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *iotevents.StartDetectorModelAnalysisInput, ...func(*iotevents.Options)) (*iotevents.StartDetectorModelAnalysisOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *iotevents.StartDetectorModelAnalysisInput, ...func(*iotevents.Options)) *iotevents.StartDetectorModelAnalysisOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*iotevents.StartDetectorModelAnalysisOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *iotevents.StartDetectorModelAnalysisInput, ...func(*iotevents.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TagResource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) TagResource(ctx context.Context, params *iotevents.TagResourceInput, optFns ...func(*iotevents.Options)) (*iotevents.TagResourceOutput, error) {
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

	var r0 *iotevents.TagResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *iotevents.TagResourceInput, ...func(*iotevents.Options)) (*iotevents.TagResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *iotevents.TagResourceInput, ...func(*iotevents.Options)) *iotevents.TagResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*iotevents.TagResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *iotevents.TagResourceInput, ...func(*iotevents.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UntagResource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) UntagResource(ctx context.Context, params *iotevents.UntagResourceInput, optFns ...func(*iotevents.Options)) (*iotevents.UntagResourceOutput, error) {
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

	var r0 *iotevents.UntagResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *iotevents.UntagResourceInput, ...func(*iotevents.Options)) (*iotevents.UntagResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *iotevents.UntagResourceInput, ...func(*iotevents.Options)) *iotevents.UntagResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*iotevents.UntagResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *iotevents.UntagResourceInput, ...func(*iotevents.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateAlarmModel provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) UpdateAlarmModel(ctx context.Context, params *iotevents.UpdateAlarmModelInput, optFns ...func(*iotevents.Options)) (*iotevents.UpdateAlarmModelOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UpdateAlarmModel")
	}

	var r0 *iotevents.UpdateAlarmModelOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *iotevents.UpdateAlarmModelInput, ...func(*iotevents.Options)) (*iotevents.UpdateAlarmModelOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *iotevents.UpdateAlarmModelInput, ...func(*iotevents.Options)) *iotevents.UpdateAlarmModelOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*iotevents.UpdateAlarmModelOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *iotevents.UpdateAlarmModelInput, ...func(*iotevents.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateDetectorModel provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) UpdateDetectorModel(ctx context.Context, params *iotevents.UpdateDetectorModelInput, optFns ...func(*iotevents.Options)) (*iotevents.UpdateDetectorModelOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UpdateDetectorModel")
	}

	var r0 *iotevents.UpdateDetectorModelOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *iotevents.UpdateDetectorModelInput, ...func(*iotevents.Options)) (*iotevents.UpdateDetectorModelOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *iotevents.UpdateDetectorModelInput, ...func(*iotevents.Options)) *iotevents.UpdateDetectorModelOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*iotevents.UpdateDetectorModelOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *iotevents.UpdateDetectorModelInput, ...func(*iotevents.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateInput provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) UpdateInput(ctx context.Context, params *iotevents.UpdateInputInput, optFns ...func(*iotevents.Options)) (*iotevents.UpdateInputOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UpdateInput")
	}

	var r0 *iotevents.UpdateInputOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *iotevents.UpdateInputInput, ...func(*iotevents.Options)) (*iotevents.UpdateInputOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *iotevents.UpdateInputInput, ...func(*iotevents.Options)) *iotevents.UpdateInputOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*iotevents.UpdateInputOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *iotevents.UpdateInputInput, ...func(*iotevents.Options)) error); ok {
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
