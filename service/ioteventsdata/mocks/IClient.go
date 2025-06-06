// Code generated by mockery v2.53.4. DO NOT EDIT.

package mocks

import (
	context "context"

	ioteventsdata "github.com/aws/aws-sdk-go-v2/service/ioteventsdata"

	mock "github.com/stretchr/testify/mock"
)

// IClient is an autogenerated mock type for the IClient type
type IClient struct {
	mock.Mock
}

// BatchAcknowledgeAlarm provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) BatchAcknowledgeAlarm(ctx context.Context, params *ioteventsdata.BatchAcknowledgeAlarmInput, optFns ...func(*ioteventsdata.Options)) (*ioteventsdata.BatchAcknowledgeAlarmOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for BatchAcknowledgeAlarm")
	}

	var r0 *ioteventsdata.BatchAcknowledgeAlarmOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *ioteventsdata.BatchAcknowledgeAlarmInput, ...func(*ioteventsdata.Options)) (*ioteventsdata.BatchAcknowledgeAlarmOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *ioteventsdata.BatchAcknowledgeAlarmInput, ...func(*ioteventsdata.Options)) *ioteventsdata.BatchAcknowledgeAlarmOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ioteventsdata.BatchAcknowledgeAlarmOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *ioteventsdata.BatchAcknowledgeAlarmInput, ...func(*ioteventsdata.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BatchDeleteDetector provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) BatchDeleteDetector(ctx context.Context, params *ioteventsdata.BatchDeleteDetectorInput, optFns ...func(*ioteventsdata.Options)) (*ioteventsdata.BatchDeleteDetectorOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for BatchDeleteDetector")
	}

	var r0 *ioteventsdata.BatchDeleteDetectorOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *ioteventsdata.BatchDeleteDetectorInput, ...func(*ioteventsdata.Options)) (*ioteventsdata.BatchDeleteDetectorOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *ioteventsdata.BatchDeleteDetectorInput, ...func(*ioteventsdata.Options)) *ioteventsdata.BatchDeleteDetectorOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ioteventsdata.BatchDeleteDetectorOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *ioteventsdata.BatchDeleteDetectorInput, ...func(*ioteventsdata.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BatchDisableAlarm provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) BatchDisableAlarm(ctx context.Context, params *ioteventsdata.BatchDisableAlarmInput, optFns ...func(*ioteventsdata.Options)) (*ioteventsdata.BatchDisableAlarmOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for BatchDisableAlarm")
	}

	var r0 *ioteventsdata.BatchDisableAlarmOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *ioteventsdata.BatchDisableAlarmInput, ...func(*ioteventsdata.Options)) (*ioteventsdata.BatchDisableAlarmOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *ioteventsdata.BatchDisableAlarmInput, ...func(*ioteventsdata.Options)) *ioteventsdata.BatchDisableAlarmOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ioteventsdata.BatchDisableAlarmOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *ioteventsdata.BatchDisableAlarmInput, ...func(*ioteventsdata.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BatchEnableAlarm provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) BatchEnableAlarm(ctx context.Context, params *ioteventsdata.BatchEnableAlarmInput, optFns ...func(*ioteventsdata.Options)) (*ioteventsdata.BatchEnableAlarmOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for BatchEnableAlarm")
	}

	var r0 *ioteventsdata.BatchEnableAlarmOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *ioteventsdata.BatchEnableAlarmInput, ...func(*ioteventsdata.Options)) (*ioteventsdata.BatchEnableAlarmOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *ioteventsdata.BatchEnableAlarmInput, ...func(*ioteventsdata.Options)) *ioteventsdata.BatchEnableAlarmOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ioteventsdata.BatchEnableAlarmOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *ioteventsdata.BatchEnableAlarmInput, ...func(*ioteventsdata.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BatchPutMessage provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) BatchPutMessage(ctx context.Context, params *ioteventsdata.BatchPutMessageInput, optFns ...func(*ioteventsdata.Options)) (*ioteventsdata.BatchPutMessageOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for BatchPutMessage")
	}

	var r0 *ioteventsdata.BatchPutMessageOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *ioteventsdata.BatchPutMessageInput, ...func(*ioteventsdata.Options)) (*ioteventsdata.BatchPutMessageOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *ioteventsdata.BatchPutMessageInput, ...func(*ioteventsdata.Options)) *ioteventsdata.BatchPutMessageOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ioteventsdata.BatchPutMessageOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *ioteventsdata.BatchPutMessageInput, ...func(*ioteventsdata.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BatchResetAlarm provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) BatchResetAlarm(ctx context.Context, params *ioteventsdata.BatchResetAlarmInput, optFns ...func(*ioteventsdata.Options)) (*ioteventsdata.BatchResetAlarmOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for BatchResetAlarm")
	}

	var r0 *ioteventsdata.BatchResetAlarmOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *ioteventsdata.BatchResetAlarmInput, ...func(*ioteventsdata.Options)) (*ioteventsdata.BatchResetAlarmOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *ioteventsdata.BatchResetAlarmInput, ...func(*ioteventsdata.Options)) *ioteventsdata.BatchResetAlarmOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ioteventsdata.BatchResetAlarmOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *ioteventsdata.BatchResetAlarmInput, ...func(*ioteventsdata.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BatchSnoozeAlarm provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) BatchSnoozeAlarm(ctx context.Context, params *ioteventsdata.BatchSnoozeAlarmInput, optFns ...func(*ioteventsdata.Options)) (*ioteventsdata.BatchSnoozeAlarmOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for BatchSnoozeAlarm")
	}

	var r0 *ioteventsdata.BatchSnoozeAlarmOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *ioteventsdata.BatchSnoozeAlarmInput, ...func(*ioteventsdata.Options)) (*ioteventsdata.BatchSnoozeAlarmOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *ioteventsdata.BatchSnoozeAlarmInput, ...func(*ioteventsdata.Options)) *ioteventsdata.BatchSnoozeAlarmOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ioteventsdata.BatchSnoozeAlarmOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *ioteventsdata.BatchSnoozeAlarmInput, ...func(*ioteventsdata.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BatchUpdateDetector provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) BatchUpdateDetector(ctx context.Context, params *ioteventsdata.BatchUpdateDetectorInput, optFns ...func(*ioteventsdata.Options)) (*ioteventsdata.BatchUpdateDetectorOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for BatchUpdateDetector")
	}

	var r0 *ioteventsdata.BatchUpdateDetectorOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *ioteventsdata.BatchUpdateDetectorInput, ...func(*ioteventsdata.Options)) (*ioteventsdata.BatchUpdateDetectorOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *ioteventsdata.BatchUpdateDetectorInput, ...func(*ioteventsdata.Options)) *ioteventsdata.BatchUpdateDetectorOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ioteventsdata.BatchUpdateDetectorOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *ioteventsdata.BatchUpdateDetectorInput, ...func(*ioteventsdata.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeAlarm provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DescribeAlarm(ctx context.Context, params *ioteventsdata.DescribeAlarmInput, optFns ...func(*ioteventsdata.Options)) (*ioteventsdata.DescribeAlarmOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DescribeAlarm")
	}

	var r0 *ioteventsdata.DescribeAlarmOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *ioteventsdata.DescribeAlarmInput, ...func(*ioteventsdata.Options)) (*ioteventsdata.DescribeAlarmOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *ioteventsdata.DescribeAlarmInput, ...func(*ioteventsdata.Options)) *ioteventsdata.DescribeAlarmOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ioteventsdata.DescribeAlarmOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *ioteventsdata.DescribeAlarmInput, ...func(*ioteventsdata.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeDetector provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DescribeDetector(ctx context.Context, params *ioteventsdata.DescribeDetectorInput, optFns ...func(*ioteventsdata.Options)) (*ioteventsdata.DescribeDetectorOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DescribeDetector")
	}

	var r0 *ioteventsdata.DescribeDetectorOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *ioteventsdata.DescribeDetectorInput, ...func(*ioteventsdata.Options)) (*ioteventsdata.DescribeDetectorOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *ioteventsdata.DescribeDetectorInput, ...func(*ioteventsdata.Options)) *ioteventsdata.DescribeDetectorOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ioteventsdata.DescribeDetectorOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *ioteventsdata.DescribeDetectorInput, ...func(*ioteventsdata.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListAlarms provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListAlarms(ctx context.Context, params *ioteventsdata.ListAlarmsInput, optFns ...func(*ioteventsdata.Options)) (*ioteventsdata.ListAlarmsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListAlarms")
	}

	var r0 *ioteventsdata.ListAlarmsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *ioteventsdata.ListAlarmsInput, ...func(*ioteventsdata.Options)) (*ioteventsdata.ListAlarmsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *ioteventsdata.ListAlarmsInput, ...func(*ioteventsdata.Options)) *ioteventsdata.ListAlarmsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ioteventsdata.ListAlarmsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *ioteventsdata.ListAlarmsInput, ...func(*ioteventsdata.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListDetectors provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListDetectors(ctx context.Context, params *ioteventsdata.ListDetectorsInput, optFns ...func(*ioteventsdata.Options)) (*ioteventsdata.ListDetectorsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListDetectors")
	}

	var r0 *ioteventsdata.ListDetectorsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *ioteventsdata.ListDetectorsInput, ...func(*ioteventsdata.Options)) (*ioteventsdata.ListDetectorsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *ioteventsdata.ListDetectorsInput, ...func(*ioteventsdata.Options)) *ioteventsdata.ListDetectorsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ioteventsdata.ListDetectorsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *ioteventsdata.ListDetectorsInput, ...func(*ioteventsdata.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Options provides a mock function with no fields
func (_m *IClient) Options() ioteventsdata.Options {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Options")
	}

	var r0 ioteventsdata.Options
	if rf, ok := ret.Get(0).(func() ioteventsdata.Options); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(ioteventsdata.Options)
	}

	return r0
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
