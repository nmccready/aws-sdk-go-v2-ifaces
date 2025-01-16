// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	context "context"

	textract "github.com/aws/aws-sdk-go-v2/service/textract"
	mock "github.com/stretchr/testify/mock"
)

// IClient is an autogenerated mock type for the IClient type
type IClient struct {
	mock.Mock
}

// AnalyzeDocument provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) AnalyzeDocument(ctx context.Context, params *textract.AnalyzeDocumentInput, optFns ...func(*textract.Options)) (*textract.AnalyzeDocumentOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for AnalyzeDocument")
	}

	var r0 *textract.AnalyzeDocumentOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *textract.AnalyzeDocumentInput, ...func(*textract.Options)) (*textract.AnalyzeDocumentOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *textract.AnalyzeDocumentInput, ...func(*textract.Options)) *textract.AnalyzeDocumentOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*textract.AnalyzeDocumentOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *textract.AnalyzeDocumentInput, ...func(*textract.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AnalyzeExpense provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) AnalyzeExpense(ctx context.Context, params *textract.AnalyzeExpenseInput, optFns ...func(*textract.Options)) (*textract.AnalyzeExpenseOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for AnalyzeExpense")
	}

	var r0 *textract.AnalyzeExpenseOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *textract.AnalyzeExpenseInput, ...func(*textract.Options)) (*textract.AnalyzeExpenseOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *textract.AnalyzeExpenseInput, ...func(*textract.Options)) *textract.AnalyzeExpenseOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*textract.AnalyzeExpenseOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *textract.AnalyzeExpenseInput, ...func(*textract.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AnalyzeID provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) AnalyzeID(ctx context.Context, params *textract.AnalyzeIDInput, optFns ...func(*textract.Options)) (*textract.AnalyzeIDOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for AnalyzeID")
	}

	var r0 *textract.AnalyzeIDOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *textract.AnalyzeIDInput, ...func(*textract.Options)) (*textract.AnalyzeIDOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *textract.AnalyzeIDInput, ...func(*textract.Options)) *textract.AnalyzeIDOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*textract.AnalyzeIDOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *textract.AnalyzeIDInput, ...func(*textract.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateAdapter provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) CreateAdapter(ctx context.Context, params *textract.CreateAdapterInput, optFns ...func(*textract.Options)) (*textract.CreateAdapterOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateAdapter")
	}

	var r0 *textract.CreateAdapterOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *textract.CreateAdapterInput, ...func(*textract.Options)) (*textract.CreateAdapterOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *textract.CreateAdapterInput, ...func(*textract.Options)) *textract.CreateAdapterOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*textract.CreateAdapterOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *textract.CreateAdapterInput, ...func(*textract.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateAdapterVersion provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) CreateAdapterVersion(ctx context.Context, params *textract.CreateAdapterVersionInput, optFns ...func(*textract.Options)) (*textract.CreateAdapterVersionOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateAdapterVersion")
	}

	var r0 *textract.CreateAdapterVersionOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *textract.CreateAdapterVersionInput, ...func(*textract.Options)) (*textract.CreateAdapterVersionOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *textract.CreateAdapterVersionInput, ...func(*textract.Options)) *textract.CreateAdapterVersionOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*textract.CreateAdapterVersionOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *textract.CreateAdapterVersionInput, ...func(*textract.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteAdapter provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DeleteAdapter(ctx context.Context, params *textract.DeleteAdapterInput, optFns ...func(*textract.Options)) (*textract.DeleteAdapterOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteAdapter")
	}

	var r0 *textract.DeleteAdapterOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *textract.DeleteAdapterInput, ...func(*textract.Options)) (*textract.DeleteAdapterOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *textract.DeleteAdapterInput, ...func(*textract.Options)) *textract.DeleteAdapterOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*textract.DeleteAdapterOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *textract.DeleteAdapterInput, ...func(*textract.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteAdapterVersion provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DeleteAdapterVersion(ctx context.Context, params *textract.DeleteAdapterVersionInput, optFns ...func(*textract.Options)) (*textract.DeleteAdapterVersionOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteAdapterVersion")
	}

	var r0 *textract.DeleteAdapterVersionOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *textract.DeleteAdapterVersionInput, ...func(*textract.Options)) (*textract.DeleteAdapterVersionOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *textract.DeleteAdapterVersionInput, ...func(*textract.Options)) *textract.DeleteAdapterVersionOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*textract.DeleteAdapterVersionOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *textract.DeleteAdapterVersionInput, ...func(*textract.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DetectDocumentText provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DetectDocumentText(ctx context.Context, params *textract.DetectDocumentTextInput, optFns ...func(*textract.Options)) (*textract.DetectDocumentTextOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DetectDocumentText")
	}

	var r0 *textract.DetectDocumentTextOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *textract.DetectDocumentTextInput, ...func(*textract.Options)) (*textract.DetectDocumentTextOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *textract.DetectDocumentTextInput, ...func(*textract.Options)) *textract.DetectDocumentTextOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*textract.DetectDocumentTextOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *textract.DetectDocumentTextInput, ...func(*textract.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAdapter provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetAdapter(ctx context.Context, params *textract.GetAdapterInput, optFns ...func(*textract.Options)) (*textract.GetAdapterOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetAdapter")
	}

	var r0 *textract.GetAdapterOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *textract.GetAdapterInput, ...func(*textract.Options)) (*textract.GetAdapterOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *textract.GetAdapterInput, ...func(*textract.Options)) *textract.GetAdapterOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*textract.GetAdapterOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *textract.GetAdapterInput, ...func(*textract.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAdapterVersion provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetAdapterVersion(ctx context.Context, params *textract.GetAdapterVersionInput, optFns ...func(*textract.Options)) (*textract.GetAdapterVersionOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetAdapterVersion")
	}

	var r0 *textract.GetAdapterVersionOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *textract.GetAdapterVersionInput, ...func(*textract.Options)) (*textract.GetAdapterVersionOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *textract.GetAdapterVersionInput, ...func(*textract.Options)) *textract.GetAdapterVersionOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*textract.GetAdapterVersionOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *textract.GetAdapterVersionInput, ...func(*textract.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDocumentAnalysis provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetDocumentAnalysis(ctx context.Context, params *textract.GetDocumentAnalysisInput, optFns ...func(*textract.Options)) (*textract.GetDocumentAnalysisOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetDocumentAnalysis")
	}

	var r0 *textract.GetDocumentAnalysisOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *textract.GetDocumentAnalysisInput, ...func(*textract.Options)) (*textract.GetDocumentAnalysisOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *textract.GetDocumentAnalysisInput, ...func(*textract.Options)) *textract.GetDocumentAnalysisOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*textract.GetDocumentAnalysisOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *textract.GetDocumentAnalysisInput, ...func(*textract.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDocumentTextDetection provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetDocumentTextDetection(ctx context.Context, params *textract.GetDocumentTextDetectionInput, optFns ...func(*textract.Options)) (*textract.GetDocumentTextDetectionOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetDocumentTextDetection")
	}

	var r0 *textract.GetDocumentTextDetectionOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *textract.GetDocumentTextDetectionInput, ...func(*textract.Options)) (*textract.GetDocumentTextDetectionOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *textract.GetDocumentTextDetectionInput, ...func(*textract.Options)) *textract.GetDocumentTextDetectionOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*textract.GetDocumentTextDetectionOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *textract.GetDocumentTextDetectionInput, ...func(*textract.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetExpenseAnalysis provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetExpenseAnalysis(ctx context.Context, params *textract.GetExpenseAnalysisInput, optFns ...func(*textract.Options)) (*textract.GetExpenseAnalysisOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetExpenseAnalysis")
	}

	var r0 *textract.GetExpenseAnalysisOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *textract.GetExpenseAnalysisInput, ...func(*textract.Options)) (*textract.GetExpenseAnalysisOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *textract.GetExpenseAnalysisInput, ...func(*textract.Options)) *textract.GetExpenseAnalysisOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*textract.GetExpenseAnalysisOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *textract.GetExpenseAnalysisInput, ...func(*textract.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLendingAnalysis provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetLendingAnalysis(ctx context.Context, params *textract.GetLendingAnalysisInput, optFns ...func(*textract.Options)) (*textract.GetLendingAnalysisOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetLendingAnalysis")
	}

	var r0 *textract.GetLendingAnalysisOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *textract.GetLendingAnalysisInput, ...func(*textract.Options)) (*textract.GetLendingAnalysisOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *textract.GetLendingAnalysisInput, ...func(*textract.Options)) *textract.GetLendingAnalysisOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*textract.GetLendingAnalysisOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *textract.GetLendingAnalysisInput, ...func(*textract.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLendingAnalysisSummary provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetLendingAnalysisSummary(ctx context.Context, params *textract.GetLendingAnalysisSummaryInput, optFns ...func(*textract.Options)) (*textract.GetLendingAnalysisSummaryOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetLendingAnalysisSummary")
	}

	var r0 *textract.GetLendingAnalysisSummaryOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *textract.GetLendingAnalysisSummaryInput, ...func(*textract.Options)) (*textract.GetLendingAnalysisSummaryOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *textract.GetLendingAnalysisSummaryInput, ...func(*textract.Options)) *textract.GetLendingAnalysisSummaryOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*textract.GetLendingAnalysisSummaryOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *textract.GetLendingAnalysisSummaryInput, ...func(*textract.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListAdapterVersions provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListAdapterVersions(ctx context.Context, params *textract.ListAdapterVersionsInput, optFns ...func(*textract.Options)) (*textract.ListAdapterVersionsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListAdapterVersions")
	}

	var r0 *textract.ListAdapterVersionsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *textract.ListAdapterVersionsInput, ...func(*textract.Options)) (*textract.ListAdapterVersionsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *textract.ListAdapterVersionsInput, ...func(*textract.Options)) *textract.ListAdapterVersionsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*textract.ListAdapterVersionsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *textract.ListAdapterVersionsInput, ...func(*textract.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListAdapters provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListAdapters(ctx context.Context, params *textract.ListAdaptersInput, optFns ...func(*textract.Options)) (*textract.ListAdaptersOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListAdapters")
	}

	var r0 *textract.ListAdaptersOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *textract.ListAdaptersInput, ...func(*textract.Options)) (*textract.ListAdaptersOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *textract.ListAdaptersInput, ...func(*textract.Options)) *textract.ListAdaptersOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*textract.ListAdaptersOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *textract.ListAdaptersInput, ...func(*textract.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListTagsForResource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListTagsForResource(ctx context.Context, params *textract.ListTagsForResourceInput, optFns ...func(*textract.Options)) (*textract.ListTagsForResourceOutput, error) {
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

	var r0 *textract.ListTagsForResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *textract.ListTagsForResourceInput, ...func(*textract.Options)) (*textract.ListTagsForResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *textract.ListTagsForResourceInput, ...func(*textract.Options)) *textract.ListTagsForResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*textract.ListTagsForResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *textract.ListTagsForResourceInput, ...func(*textract.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Options provides a mock function with given fields:
func (_m *IClient) Options() textract.Options {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Options")
	}

	var r0 textract.Options
	if rf, ok := ret.Get(0).(func() textract.Options); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(textract.Options)
	}

	return r0
}

// StartDocumentAnalysis provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) StartDocumentAnalysis(ctx context.Context, params *textract.StartDocumentAnalysisInput, optFns ...func(*textract.Options)) (*textract.StartDocumentAnalysisOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for StartDocumentAnalysis")
	}

	var r0 *textract.StartDocumentAnalysisOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *textract.StartDocumentAnalysisInput, ...func(*textract.Options)) (*textract.StartDocumentAnalysisOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *textract.StartDocumentAnalysisInput, ...func(*textract.Options)) *textract.StartDocumentAnalysisOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*textract.StartDocumentAnalysisOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *textract.StartDocumentAnalysisInput, ...func(*textract.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StartDocumentTextDetection provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) StartDocumentTextDetection(ctx context.Context, params *textract.StartDocumentTextDetectionInput, optFns ...func(*textract.Options)) (*textract.StartDocumentTextDetectionOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for StartDocumentTextDetection")
	}

	var r0 *textract.StartDocumentTextDetectionOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *textract.StartDocumentTextDetectionInput, ...func(*textract.Options)) (*textract.StartDocumentTextDetectionOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *textract.StartDocumentTextDetectionInput, ...func(*textract.Options)) *textract.StartDocumentTextDetectionOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*textract.StartDocumentTextDetectionOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *textract.StartDocumentTextDetectionInput, ...func(*textract.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StartExpenseAnalysis provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) StartExpenseAnalysis(ctx context.Context, params *textract.StartExpenseAnalysisInput, optFns ...func(*textract.Options)) (*textract.StartExpenseAnalysisOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for StartExpenseAnalysis")
	}

	var r0 *textract.StartExpenseAnalysisOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *textract.StartExpenseAnalysisInput, ...func(*textract.Options)) (*textract.StartExpenseAnalysisOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *textract.StartExpenseAnalysisInput, ...func(*textract.Options)) *textract.StartExpenseAnalysisOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*textract.StartExpenseAnalysisOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *textract.StartExpenseAnalysisInput, ...func(*textract.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StartLendingAnalysis provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) StartLendingAnalysis(ctx context.Context, params *textract.StartLendingAnalysisInput, optFns ...func(*textract.Options)) (*textract.StartLendingAnalysisOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for StartLendingAnalysis")
	}

	var r0 *textract.StartLendingAnalysisOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *textract.StartLendingAnalysisInput, ...func(*textract.Options)) (*textract.StartLendingAnalysisOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *textract.StartLendingAnalysisInput, ...func(*textract.Options)) *textract.StartLendingAnalysisOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*textract.StartLendingAnalysisOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *textract.StartLendingAnalysisInput, ...func(*textract.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TagResource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) TagResource(ctx context.Context, params *textract.TagResourceInput, optFns ...func(*textract.Options)) (*textract.TagResourceOutput, error) {
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

	var r0 *textract.TagResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *textract.TagResourceInput, ...func(*textract.Options)) (*textract.TagResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *textract.TagResourceInput, ...func(*textract.Options)) *textract.TagResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*textract.TagResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *textract.TagResourceInput, ...func(*textract.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UntagResource provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) UntagResource(ctx context.Context, params *textract.UntagResourceInput, optFns ...func(*textract.Options)) (*textract.UntagResourceOutput, error) {
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

	var r0 *textract.UntagResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *textract.UntagResourceInput, ...func(*textract.Options)) (*textract.UntagResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *textract.UntagResourceInput, ...func(*textract.Options)) *textract.UntagResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*textract.UntagResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *textract.UntagResourceInput, ...func(*textract.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateAdapter provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) UpdateAdapter(ctx context.Context, params *textract.UpdateAdapterInput, optFns ...func(*textract.Options)) (*textract.UpdateAdapterOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UpdateAdapter")
	}

	var r0 *textract.UpdateAdapterOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *textract.UpdateAdapterInput, ...func(*textract.Options)) (*textract.UpdateAdapterOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *textract.UpdateAdapterInput, ...func(*textract.Options)) *textract.UpdateAdapterOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*textract.UpdateAdapterOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *textract.UpdateAdapterInput, ...func(*textract.Options)) error); ok {
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
