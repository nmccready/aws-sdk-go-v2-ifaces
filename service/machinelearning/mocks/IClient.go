// Code generated by mockery v2.53.4. DO NOT EDIT.

package mocks

import (
	context "context"

	machinelearning "github.com/aws/aws-sdk-go-v2/service/machinelearning"

	mock "github.com/stretchr/testify/mock"
)

// IClient is an autogenerated mock type for the IClient type
type IClient struct {
	mock.Mock
}

// AddTags provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) AddTags(ctx context.Context, params *machinelearning.AddTagsInput, optFns ...func(*machinelearning.Options)) (*machinelearning.AddTagsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for AddTags")
	}

	var r0 *machinelearning.AddTagsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *machinelearning.AddTagsInput, ...func(*machinelearning.Options)) (*machinelearning.AddTagsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *machinelearning.AddTagsInput, ...func(*machinelearning.Options)) *machinelearning.AddTagsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*machinelearning.AddTagsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *machinelearning.AddTagsInput, ...func(*machinelearning.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateBatchPrediction provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) CreateBatchPrediction(ctx context.Context, params *machinelearning.CreateBatchPredictionInput, optFns ...func(*machinelearning.Options)) (*machinelearning.CreateBatchPredictionOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateBatchPrediction")
	}

	var r0 *machinelearning.CreateBatchPredictionOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *machinelearning.CreateBatchPredictionInput, ...func(*machinelearning.Options)) (*machinelearning.CreateBatchPredictionOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *machinelearning.CreateBatchPredictionInput, ...func(*machinelearning.Options)) *machinelearning.CreateBatchPredictionOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*machinelearning.CreateBatchPredictionOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *machinelearning.CreateBatchPredictionInput, ...func(*machinelearning.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateDataSourceFromRDS provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) CreateDataSourceFromRDS(ctx context.Context, params *machinelearning.CreateDataSourceFromRDSInput, optFns ...func(*machinelearning.Options)) (*machinelearning.CreateDataSourceFromRDSOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateDataSourceFromRDS")
	}

	var r0 *machinelearning.CreateDataSourceFromRDSOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *machinelearning.CreateDataSourceFromRDSInput, ...func(*machinelearning.Options)) (*machinelearning.CreateDataSourceFromRDSOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *machinelearning.CreateDataSourceFromRDSInput, ...func(*machinelearning.Options)) *machinelearning.CreateDataSourceFromRDSOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*machinelearning.CreateDataSourceFromRDSOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *machinelearning.CreateDataSourceFromRDSInput, ...func(*machinelearning.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateDataSourceFromRedshift provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) CreateDataSourceFromRedshift(ctx context.Context, params *machinelearning.CreateDataSourceFromRedshiftInput, optFns ...func(*machinelearning.Options)) (*machinelearning.CreateDataSourceFromRedshiftOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateDataSourceFromRedshift")
	}

	var r0 *machinelearning.CreateDataSourceFromRedshiftOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *machinelearning.CreateDataSourceFromRedshiftInput, ...func(*machinelearning.Options)) (*machinelearning.CreateDataSourceFromRedshiftOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *machinelearning.CreateDataSourceFromRedshiftInput, ...func(*machinelearning.Options)) *machinelearning.CreateDataSourceFromRedshiftOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*machinelearning.CreateDataSourceFromRedshiftOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *machinelearning.CreateDataSourceFromRedshiftInput, ...func(*machinelearning.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateDataSourceFromS3 provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) CreateDataSourceFromS3(ctx context.Context, params *machinelearning.CreateDataSourceFromS3Input, optFns ...func(*machinelearning.Options)) (*machinelearning.CreateDataSourceFromS3Output, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateDataSourceFromS3")
	}

	var r0 *machinelearning.CreateDataSourceFromS3Output
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *machinelearning.CreateDataSourceFromS3Input, ...func(*machinelearning.Options)) (*machinelearning.CreateDataSourceFromS3Output, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *machinelearning.CreateDataSourceFromS3Input, ...func(*machinelearning.Options)) *machinelearning.CreateDataSourceFromS3Output); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*machinelearning.CreateDataSourceFromS3Output)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *machinelearning.CreateDataSourceFromS3Input, ...func(*machinelearning.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateEvaluation provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) CreateEvaluation(ctx context.Context, params *machinelearning.CreateEvaluationInput, optFns ...func(*machinelearning.Options)) (*machinelearning.CreateEvaluationOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateEvaluation")
	}

	var r0 *machinelearning.CreateEvaluationOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *machinelearning.CreateEvaluationInput, ...func(*machinelearning.Options)) (*machinelearning.CreateEvaluationOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *machinelearning.CreateEvaluationInput, ...func(*machinelearning.Options)) *machinelearning.CreateEvaluationOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*machinelearning.CreateEvaluationOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *machinelearning.CreateEvaluationInput, ...func(*machinelearning.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateMLModel provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) CreateMLModel(ctx context.Context, params *machinelearning.CreateMLModelInput, optFns ...func(*machinelearning.Options)) (*machinelearning.CreateMLModelOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateMLModel")
	}

	var r0 *machinelearning.CreateMLModelOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *machinelearning.CreateMLModelInput, ...func(*machinelearning.Options)) (*machinelearning.CreateMLModelOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *machinelearning.CreateMLModelInput, ...func(*machinelearning.Options)) *machinelearning.CreateMLModelOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*machinelearning.CreateMLModelOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *machinelearning.CreateMLModelInput, ...func(*machinelearning.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateRealtimeEndpoint provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) CreateRealtimeEndpoint(ctx context.Context, params *machinelearning.CreateRealtimeEndpointInput, optFns ...func(*machinelearning.Options)) (*machinelearning.CreateRealtimeEndpointOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateRealtimeEndpoint")
	}

	var r0 *machinelearning.CreateRealtimeEndpointOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *machinelearning.CreateRealtimeEndpointInput, ...func(*machinelearning.Options)) (*machinelearning.CreateRealtimeEndpointOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *machinelearning.CreateRealtimeEndpointInput, ...func(*machinelearning.Options)) *machinelearning.CreateRealtimeEndpointOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*machinelearning.CreateRealtimeEndpointOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *machinelearning.CreateRealtimeEndpointInput, ...func(*machinelearning.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteBatchPrediction provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DeleteBatchPrediction(ctx context.Context, params *machinelearning.DeleteBatchPredictionInput, optFns ...func(*machinelearning.Options)) (*machinelearning.DeleteBatchPredictionOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteBatchPrediction")
	}

	var r0 *machinelearning.DeleteBatchPredictionOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *machinelearning.DeleteBatchPredictionInput, ...func(*machinelearning.Options)) (*machinelearning.DeleteBatchPredictionOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *machinelearning.DeleteBatchPredictionInput, ...func(*machinelearning.Options)) *machinelearning.DeleteBatchPredictionOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*machinelearning.DeleteBatchPredictionOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *machinelearning.DeleteBatchPredictionInput, ...func(*machinelearning.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteDataSource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DeleteDataSource(ctx context.Context, params *machinelearning.DeleteDataSourceInput, optFns ...func(*machinelearning.Options)) (*machinelearning.DeleteDataSourceOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteDataSource")
	}

	var r0 *machinelearning.DeleteDataSourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *machinelearning.DeleteDataSourceInput, ...func(*machinelearning.Options)) (*machinelearning.DeleteDataSourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *machinelearning.DeleteDataSourceInput, ...func(*machinelearning.Options)) *machinelearning.DeleteDataSourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*machinelearning.DeleteDataSourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *machinelearning.DeleteDataSourceInput, ...func(*machinelearning.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteEvaluation provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DeleteEvaluation(ctx context.Context, params *machinelearning.DeleteEvaluationInput, optFns ...func(*machinelearning.Options)) (*machinelearning.DeleteEvaluationOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteEvaluation")
	}

	var r0 *machinelearning.DeleteEvaluationOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *machinelearning.DeleteEvaluationInput, ...func(*machinelearning.Options)) (*machinelearning.DeleteEvaluationOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *machinelearning.DeleteEvaluationInput, ...func(*machinelearning.Options)) *machinelearning.DeleteEvaluationOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*machinelearning.DeleteEvaluationOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *machinelearning.DeleteEvaluationInput, ...func(*machinelearning.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteMLModel provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DeleteMLModel(ctx context.Context, params *machinelearning.DeleteMLModelInput, optFns ...func(*machinelearning.Options)) (*machinelearning.DeleteMLModelOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteMLModel")
	}

	var r0 *machinelearning.DeleteMLModelOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *machinelearning.DeleteMLModelInput, ...func(*machinelearning.Options)) (*machinelearning.DeleteMLModelOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *machinelearning.DeleteMLModelInput, ...func(*machinelearning.Options)) *machinelearning.DeleteMLModelOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*machinelearning.DeleteMLModelOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *machinelearning.DeleteMLModelInput, ...func(*machinelearning.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteRealtimeEndpoint provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DeleteRealtimeEndpoint(ctx context.Context, params *machinelearning.DeleteRealtimeEndpointInput, optFns ...func(*machinelearning.Options)) (*machinelearning.DeleteRealtimeEndpointOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteRealtimeEndpoint")
	}

	var r0 *machinelearning.DeleteRealtimeEndpointOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *machinelearning.DeleteRealtimeEndpointInput, ...func(*machinelearning.Options)) (*machinelearning.DeleteRealtimeEndpointOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *machinelearning.DeleteRealtimeEndpointInput, ...func(*machinelearning.Options)) *machinelearning.DeleteRealtimeEndpointOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*machinelearning.DeleteRealtimeEndpointOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *machinelearning.DeleteRealtimeEndpointInput, ...func(*machinelearning.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteTags provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DeleteTags(ctx context.Context, params *machinelearning.DeleteTagsInput, optFns ...func(*machinelearning.Options)) (*machinelearning.DeleteTagsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteTags")
	}

	var r0 *machinelearning.DeleteTagsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *machinelearning.DeleteTagsInput, ...func(*machinelearning.Options)) (*machinelearning.DeleteTagsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *machinelearning.DeleteTagsInput, ...func(*machinelearning.Options)) *machinelearning.DeleteTagsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*machinelearning.DeleteTagsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *machinelearning.DeleteTagsInput, ...func(*machinelearning.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeBatchPredictions provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DescribeBatchPredictions(ctx context.Context, params *machinelearning.DescribeBatchPredictionsInput, optFns ...func(*machinelearning.Options)) (*machinelearning.DescribeBatchPredictionsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DescribeBatchPredictions")
	}

	var r0 *machinelearning.DescribeBatchPredictionsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *machinelearning.DescribeBatchPredictionsInput, ...func(*machinelearning.Options)) (*machinelearning.DescribeBatchPredictionsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *machinelearning.DescribeBatchPredictionsInput, ...func(*machinelearning.Options)) *machinelearning.DescribeBatchPredictionsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*machinelearning.DescribeBatchPredictionsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *machinelearning.DescribeBatchPredictionsInput, ...func(*machinelearning.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeDataSources provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DescribeDataSources(ctx context.Context, params *machinelearning.DescribeDataSourcesInput, optFns ...func(*machinelearning.Options)) (*machinelearning.DescribeDataSourcesOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DescribeDataSources")
	}

	var r0 *machinelearning.DescribeDataSourcesOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *machinelearning.DescribeDataSourcesInput, ...func(*machinelearning.Options)) (*machinelearning.DescribeDataSourcesOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *machinelearning.DescribeDataSourcesInput, ...func(*machinelearning.Options)) *machinelearning.DescribeDataSourcesOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*machinelearning.DescribeDataSourcesOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *machinelearning.DescribeDataSourcesInput, ...func(*machinelearning.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeEvaluations provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DescribeEvaluations(ctx context.Context, params *machinelearning.DescribeEvaluationsInput, optFns ...func(*machinelearning.Options)) (*machinelearning.DescribeEvaluationsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DescribeEvaluations")
	}

	var r0 *machinelearning.DescribeEvaluationsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *machinelearning.DescribeEvaluationsInput, ...func(*machinelearning.Options)) (*machinelearning.DescribeEvaluationsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *machinelearning.DescribeEvaluationsInput, ...func(*machinelearning.Options)) *machinelearning.DescribeEvaluationsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*machinelearning.DescribeEvaluationsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *machinelearning.DescribeEvaluationsInput, ...func(*machinelearning.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeMLModels provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DescribeMLModels(ctx context.Context, params *machinelearning.DescribeMLModelsInput, optFns ...func(*machinelearning.Options)) (*machinelearning.DescribeMLModelsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DescribeMLModels")
	}

	var r0 *machinelearning.DescribeMLModelsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *machinelearning.DescribeMLModelsInput, ...func(*machinelearning.Options)) (*machinelearning.DescribeMLModelsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *machinelearning.DescribeMLModelsInput, ...func(*machinelearning.Options)) *machinelearning.DescribeMLModelsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*machinelearning.DescribeMLModelsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *machinelearning.DescribeMLModelsInput, ...func(*machinelearning.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeTags provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DescribeTags(ctx context.Context, params *machinelearning.DescribeTagsInput, optFns ...func(*machinelearning.Options)) (*machinelearning.DescribeTagsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DescribeTags")
	}

	var r0 *machinelearning.DescribeTagsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *machinelearning.DescribeTagsInput, ...func(*machinelearning.Options)) (*machinelearning.DescribeTagsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *machinelearning.DescribeTagsInput, ...func(*machinelearning.Options)) *machinelearning.DescribeTagsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*machinelearning.DescribeTagsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *machinelearning.DescribeTagsInput, ...func(*machinelearning.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBatchPrediction provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetBatchPrediction(ctx context.Context, params *machinelearning.GetBatchPredictionInput, optFns ...func(*machinelearning.Options)) (*machinelearning.GetBatchPredictionOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetBatchPrediction")
	}

	var r0 *machinelearning.GetBatchPredictionOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *machinelearning.GetBatchPredictionInput, ...func(*machinelearning.Options)) (*machinelearning.GetBatchPredictionOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *machinelearning.GetBatchPredictionInput, ...func(*machinelearning.Options)) *machinelearning.GetBatchPredictionOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*machinelearning.GetBatchPredictionOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *machinelearning.GetBatchPredictionInput, ...func(*machinelearning.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDataSource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetDataSource(ctx context.Context, params *machinelearning.GetDataSourceInput, optFns ...func(*machinelearning.Options)) (*machinelearning.GetDataSourceOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetDataSource")
	}

	var r0 *machinelearning.GetDataSourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *machinelearning.GetDataSourceInput, ...func(*machinelearning.Options)) (*machinelearning.GetDataSourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *machinelearning.GetDataSourceInput, ...func(*machinelearning.Options)) *machinelearning.GetDataSourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*machinelearning.GetDataSourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *machinelearning.GetDataSourceInput, ...func(*machinelearning.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetEvaluation provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetEvaluation(ctx context.Context, params *machinelearning.GetEvaluationInput, optFns ...func(*machinelearning.Options)) (*machinelearning.GetEvaluationOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetEvaluation")
	}

	var r0 *machinelearning.GetEvaluationOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *machinelearning.GetEvaluationInput, ...func(*machinelearning.Options)) (*machinelearning.GetEvaluationOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *machinelearning.GetEvaluationInput, ...func(*machinelearning.Options)) *machinelearning.GetEvaluationOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*machinelearning.GetEvaluationOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *machinelearning.GetEvaluationInput, ...func(*machinelearning.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMLModel provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetMLModel(ctx context.Context, params *machinelearning.GetMLModelInput, optFns ...func(*machinelearning.Options)) (*machinelearning.GetMLModelOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetMLModel")
	}

	var r0 *machinelearning.GetMLModelOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *machinelearning.GetMLModelInput, ...func(*machinelearning.Options)) (*machinelearning.GetMLModelOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *machinelearning.GetMLModelInput, ...func(*machinelearning.Options)) *machinelearning.GetMLModelOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*machinelearning.GetMLModelOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *machinelearning.GetMLModelInput, ...func(*machinelearning.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Options provides a mock function with no fields
func (_m *IClient) Options() machinelearning.Options {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Options")
	}

	var r0 machinelearning.Options
	if rf, ok := ret.Get(0).(func() machinelearning.Options); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(machinelearning.Options)
	}

	return r0
}

// Predict provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) Predict(ctx context.Context, params *machinelearning.PredictInput, optFns ...func(*machinelearning.Options)) (*machinelearning.PredictOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Predict")
	}

	var r0 *machinelearning.PredictOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *machinelearning.PredictInput, ...func(*machinelearning.Options)) (*machinelearning.PredictOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *machinelearning.PredictInput, ...func(*machinelearning.Options)) *machinelearning.PredictOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*machinelearning.PredictOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *machinelearning.PredictInput, ...func(*machinelearning.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateBatchPrediction provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) UpdateBatchPrediction(ctx context.Context, params *machinelearning.UpdateBatchPredictionInput, optFns ...func(*machinelearning.Options)) (*machinelearning.UpdateBatchPredictionOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UpdateBatchPrediction")
	}

	var r0 *machinelearning.UpdateBatchPredictionOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *machinelearning.UpdateBatchPredictionInput, ...func(*machinelearning.Options)) (*machinelearning.UpdateBatchPredictionOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *machinelearning.UpdateBatchPredictionInput, ...func(*machinelearning.Options)) *machinelearning.UpdateBatchPredictionOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*machinelearning.UpdateBatchPredictionOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *machinelearning.UpdateBatchPredictionInput, ...func(*machinelearning.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateDataSource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) UpdateDataSource(ctx context.Context, params *machinelearning.UpdateDataSourceInput, optFns ...func(*machinelearning.Options)) (*machinelearning.UpdateDataSourceOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UpdateDataSource")
	}

	var r0 *machinelearning.UpdateDataSourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *machinelearning.UpdateDataSourceInput, ...func(*machinelearning.Options)) (*machinelearning.UpdateDataSourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *machinelearning.UpdateDataSourceInput, ...func(*machinelearning.Options)) *machinelearning.UpdateDataSourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*machinelearning.UpdateDataSourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *machinelearning.UpdateDataSourceInput, ...func(*machinelearning.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateEvaluation provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) UpdateEvaluation(ctx context.Context, params *machinelearning.UpdateEvaluationInput, optFns ...func(*machinelearning.Options)) (*machinelearning.UpdateEvaluationOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UpdateEvaluation")
	}

	var r0 *machinelearning.UpdateEvaluationOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *machinelearning.UpdateEvaluationInput, ...func(*machinelearning.Options)) (*machinelearning.UpdateEvaluationOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *machinelearning.UpdateEvaluationInput, ...func(*machinelearning.Options)) *machinelearning.UpdateEvaluationOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*machinelearning.UpdateEvaluationOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *machinelearning.UpdateEvaluationInput, ...func(*machinelearning.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateMLModel provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) UpdateMLModel(ctx context.Context, params *machinelearning.UpdateMLModelInput, optFns ...func(*machinelearning.Options)) (*machinelearning.UpdateMLModelOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UpdateMLModel")
	}

	var r0 *machinelearning.UpdateMLModelOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *machinelearning.UpdateMLModelInput, ...func(*machinelearning.Options)) (*machinelearning.UpdateMLModelOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *machinelearning.UpdateMLModelInput, ...func(*machinelearning.Options)) *machinelearning.UpdateMLModelOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*machinelearning.UpdateMLModelOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *machinelearning.UpdateMLModelInput, ...func(*machinelearning.Options)) error); ok {
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
