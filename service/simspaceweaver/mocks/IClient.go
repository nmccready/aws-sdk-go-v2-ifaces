// Code generated by mockery v2.39.2. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	simspaceweaver "github.com/aws/aws-sdk-go-v2/service/simspaceweaver"
)

// IClient is an autogenerated mock type for the IClient type
type IClient struct {
	mock.Mock
}

// CreateSnapshot provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) CreateSnapshot(ctx context.Context, params *simspaceweaver.CreateSnapshotInput, optFns ...func(*simspaceweaver.Options)) (*simspaceweaver.CreateSnapshotOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateSnapshot")
	}

	var r0 *simspaceweaver.CreateSnapshotOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *simspaceweaver.CreateSnapshotInput, ...func(*simspaceweaver.Options)) (*simspaceweaver.CreateSnapshotOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *simspaceweaver.CreateSnapshotInput, ...func(*simspaceweaver.Options)) *simspaceweaver.CreateSnapshotOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*simspaceweaver.CreateSnapshotOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *simspaceweaver.CreateSnapshotInput, ...func(*simspaceweaver.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteApp provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DeleteApp(ctx context.Context, params *simspaceweaver.DeleteAppInput, optFns ...func(*simspaceweaver.Options)) (*simspaceweaver.DeleteAppOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteApp")
	}

	var r0 *simspaceweaver.DeleteAppOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *simspaceweaver.DeleteAppInput, ...func(*simspaceweaver.Options)) (*simspaceweaver.DeleteAppOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *simspaceweaver.DeleteAppInput, ...func(*simspaceweaver.Options)) *simspaceweaver.DeleteAppOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*simspaceweaver.DeleteAppOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *simspaceweaver.DeleteAppInput, ...func(*simspaceweaver.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteSimulation provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DeleteSimulation(ctx context.Context, params *simspaceweaver.DeleteSimulationInput, optFns ...func(*simspaceweaver.Options)) (*simspaceweaver.DeleteSimulationOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteSimulation")
	}

	var r0 *simspaceweaver.DeleteSimulationOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *simspaceweaver.DeleteSimulationInput, ...func(*simspaceweaver.Options)) (*simspaceweaver.DeleteSimulationOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *simspaceweaver.DeleteSimulationInput, ...func(*simspaceweaver.Options)) *simspaceweaver.DeleteSimulationOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*simspaceweaver.DeleteSimulationOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *simspaceweaver.DeleteSimulationInput, ...func(*simspaceweaver.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeApp provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DescribeApp(ctx context.Context, params *simspaceweaver.DescribeAppInput, optFns ...func(*simspaceweaver.Options)) (*simspaceweaver.DescribeAppOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DescribeApp")
	}

	var r0 *simspaceweaver.DescribeAppOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *simspaceweaver.DescribeAppInput, ...func(*simspaceweaver.Options)) (*simspaceweaver.DescribeAppOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *simspaceweaver.DescribeAppInput, ...func(*simspaceweaver.Options)) *simspaceweaver.DescribeAppOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*simspaceweaver.DescribeAppOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *simspaceweaver.DescribeAppInput, ...func(*simspaceweaver.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeSimulation provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DescribeSimulation(ctx context.Context, params *simspaceweaver.DescribeSimulationInput, optFns ...func(*simspaceweaver.Options)) (*simspaceweaver.DescribeSimulationOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DescribeSimulation")
	}

	var r0 *simspaceweaver.DescribeSimulationOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *simspaceweaver.DescribeSimulationInput, ...func(*simspaceweaver.Options)) (*simspaceweaver.DescribeSimulationOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *simspaceweaver.DescribeSimulationInput, ...func(*simspaceweaver.Options)) *simspaceweaver.DescribeSimulationOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*simspaceweaver.DescribeSimulationOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *simspaceweaver.DescribeSimulationInput, ...func(*simspaceweaver.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListApps provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListApps(ctx context.Context, params *simspaceweaver.ListAppsInput, optFns ...func(*simspaceweaver.Options)) (*simspaceweaver.ListAppsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListApps")
	}

	var r0 *simspaceweaver.ListAppsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *simspaceweaver.ListAppsInput, ...func(*simspaceweaver.Options)) (*simspaceweaver.ListAppsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *simspaceweaver.ListAppsInput, ...func(*simspaceweaver.Options)) *simspaceweaver.ListAppsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*simspaceweaver.ListAppsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *simspaceweaver.ListAppsInput, ...func(*simspaceweaver.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListSimulations provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListSimulations(ctx context.Context, params *simspaceweaver.ListSimulationsInput, optFns ...func(*simspaceweaver.Options)) (*simspaceweaver.ListSimulationsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListSimulations")
	}

	var r0 *simspaceweaver.ListSimulationsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *simspaceweaver.ListSimulationsInput, ...func(*simspaceweaver.Options)) (*simspaceweaver.ListSimulationsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *simspaceweaver.ListSimulationsInput, ...func(*simspaceweaver.Options)) *simspaceweaver.ListSimulationsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*simspaceweaver.ListSimulationsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *simspaceweaver.ListSimulationsInput, ...func(*simspaceweaver.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListTagsForResource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListTagsForResource(ctx context.Context, params *simspaceweaver.ListTagsForResourceInput, optFns ...func(*simspaceweaver.Options)) (*simspaceweaver.ListTagsForResourceOutput, error) {
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

	var r0 *simspaceweaver.ListTagsForResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *simspaceweaver.ListTagsForResourceInput, ...func(*simspaceweaver.Options)) (*simspaceweaver.ListTagsForResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *simspaceweaver.ListTagsForResourceInput, ...func(*simspaceweaver.Options)) *simspaceweaver.ListTagsForResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*simspaceweaver.ListTagsForResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *simspaceweaver.ListTagsForResourceInput, ...func(*simspaceweaver.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Options provides a mock function with given fields:
func (_m *IClient) Options() simspaceweaver.Options {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Options")
	}

	var r0 simspaceweaver.Options
	if rf, ok := ret.Get(0).(func() simspaceweaver.Options); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(simspaceweaver.Options)
	}

	return r0
}

// StartApp provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) StartApp(ctx context.Context, params *simspaceweaver.StartAppInput, optFns ...func(*simspaceweaver.Options)) (*simspaceweaver.StartAppOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for StartApp")
	}

	var r0 *simspaceweaver.StartAppOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *simspaceweaver.StartAppInput, ...func(*simspaceweaver.Options)) (*simspaceweaver.StartAppOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *simspaceweaver.StartAppInput, ...func(*simspaceweaver.Options)) *simspaceweaver.StartAppOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*simspaceweaver.StartAppOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *simspaceweaver.StartAppInput, ...func(*simspaceweaver.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StartClock provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) StartClock(ctx context.Context, params *simspaceweaver.StartClockInput, optFns ...func(*simspaceweaver.Options)) (*simspaceweaver.StartClockOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for StartClock")
	}

	var r0 *simspaceweaver.StartClockOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *simspaceweaver.StartClockInput, ...func(*simspaceweaver.Options)) (*simspaceweaver.StartClockOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *simspaceweaver.StartClockInput, ...func(*simspaceweaver.Options)) *simspaceweaver.StartClockOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*simspaceweaver.StartClockOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *simspaceweaver.StartClockInput, ...func(*simspaceweaver.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StartSimulation provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) StartSimulation(ctx context.Context, params *simspaceweaver.StartSimulationInput, optFns ...func(*simspaceweaver.Options)) (*simspaceweaver.StartSimulationOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for StartSimulation")
	}

	var r0 *simspaceweaver.StartSimulationOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *simspaceweaver.StartSimulationInput, ...func(*simspaceweaver.Options)) (*simspaceweaver.StartSimulationOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *simspaceweaver.StartSimulationInput, ...func(*simspaceweaver.Options)) *simspaceweaver.StartSimulationOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*simspaceweaver.StartSimulationOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *simspaceweaver.StartSimulationInput, ...func(*simspaceweaver.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StopApp provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) StopApp(ctx context.Context, params *simspaceweaver.StopAppInput, optFns ...func(*simspaceweaver.Options)) (*simspaceweaver.StopAppOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for StopApp")
	}

	var r0 *simspaceweaver.StopAppOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *simspaceweaver.StopAppInput, ...func(*simspaceweaver.Options)) (*simspaceweaver.StopAppOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *simspaceweaver.StopAppInput, ...func(*simspaceweaver.Options)) *simspaceweaver.StopAppOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*simspaceweaver.StopAppOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *simspaceweaver.StopAppInput, ...func(*simspaceweaver.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StopClock provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) StopClock(ctx context.Context, params *simspaceweaver.StopClockInput, optFns ...func(*simspaceweaver.Options)) (*simspaceweaver.StopClockOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for StopClock")
	}

	var r0 *simspaceweaver.StopClockOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *simspaceweaver.StopClockInput, ...func(*simspaceweaver.Options)) (*simspaceweaver.StopClockOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *simspaceweaver.StopClockInput, ...func(*simspaceweaver.Options)) *simspaceweaver.StopClockOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*simspaceweaver.StopClockOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *simspaceweaver.StopClockInput, ...func(*simspaceweaver.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StopSimulation provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) StopSimulation(ctx context.Context, params *simspaceweaver.StopSimulationInput, optFns ...func(*simspaceweaver.Options)) (*simspaceweaver.StopSimulationOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for StopSimulation")
	}

	var r0 *simspaceweaver.StopSimulationOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *simspaceweaver.StopSimulationInput, ...func(*simspaceweaver.Options)) (*simspaceweaver.StopSimulationOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *simspaceweaver.StopSimulationInput, ...func(*simspaceweaver.Options)) *simspaceweaver.StopSimulationOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*simspaceweaver.StopSimulationOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *simspaceweaver.StopSimulationInput, ...func(*simspaceweaver.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TagResource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) TagResource(ctx context.Context, params *simspaceweaver.TagResourceInput, optFns ...func(*simspaceweaver.Options)) (*simspaceweaver.TagResourceOutput, error) {
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

	var r0 *simspaceweaver.TagResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *simspaceweaver.TagResourceInput, ...func(*simspaceweaver.Options)) (*simspaceweaver.TagResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *simspaceweaver.TagResourceInput, ...func(*simspaceweaver.Options)) *simspaceweaver.TagResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*simspaceweaver.TagResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *simspaceweaver.TagResourceInput, ...func(*simspaceweaver.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UntagResource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) UntagResource(ctx context.Context, params *simspaceweaver.UntagResourceInput, optFns ...func(*simspaceweaver.Options)) (*simspaceweaver.UntagResourceOutput, error) {
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

	var r0 *simspaceweaver.UntagResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *simspaceweaver.UntagResourceInput, ...func(*simspaceweaver.Options)) (*simspaceweaver.UntagResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *simspaceweaver.UntagResourceInput, ...func(*simspaceweaver.Options)) *simspaceweaver.UntagResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*simspaceweaver.UntagResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *simspaceweaver.UntagResourceInput, ...func(*simspaceweaver.Options)) error); ok {
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