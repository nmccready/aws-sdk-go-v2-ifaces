// Code generated by mockery v2.53.4. DO NOT EDIT.

package mocks

import (
	context "context"

	marketplaceagreement "github.com/aws/aws-sdk-go-v2/service/marketplaceagreement"

	mock "github.com/stretchr/testify/mock"
)

// IClient is an autogenerated mock type for the IClient type
type IClient struct {
	mock.Mock
}

// DescribeAgreement provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DescribeAgreement(ctx context.Context, params *marketplaceagreement.DescribeAgreementInput, optFns ...func(*marketplaceagreement.Options)) (*marketplaceagreement.DescribeAgreementOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DescribeAgreement")
	}

	var r0 *marketplaceagreement.DescribeAgreementOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *marketplaceagreement.DescribeAgreementInput, ...func(*marketplaceagreement.Options)) (*marketplaceagreement.DescribeAgreementOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *marketplaceagreement.DescribeAgreementInput, ...func(*marketplaceagreement.Options)) *marketplaceagreement.DescribeAgreementOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*marketplaceagreement.DescribeAgreementOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *marketplaceagreement.DescribeAgreementInput, ...func(*marketplaceagreement.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAgreementTerms provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetAgreementTerms(ctx context.Context, params *marketplaceagreement.GetAgreementTermsInput, optFns ...func(*marketplaceagreement.Options)) (*marketplaceagreement.GetAgreementTermsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetAgreementTerms")
	}

	var r0 *marketplaceagreement.GetAgreementTermsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *marketplaceagreement.GetAgreementTermsInput, ...func(*marketplaceagreement.Options)) (*marketplaceagreement.GetAgreementTermsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *marketplaceagreement.GetAgreementTermsInput, ...func(*marketplaceagreement.Options)) *marketplaceagreement.GetAgreementTermsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*marketplaceagreement.GetAgreementTermsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *marketplaceagreement.GetAgreementTermsInput, ...func(*marketplaceagreement.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Options provides a mock function with no fields
func (_m *IClient) Options() marketplaceagreement.Options {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Options")
	}

	var r0 marketplaceagreement.Options
	if rf, ok := ret.Get(0).(func() marketplaceagreement.Options); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(marketplaceagreement.Options)
	}

	return r0
}

// SearchAgreements provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) SearchAgreements(ctx context.Context, params *marketplaceagreement.SearchAgreementsInput, optFns ...func(*marketplaceagreement.Options)) (*marketplaceagreement.SearchAgreementsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for SearchAgreements")
	}

	var r0 *marketplaceagreement.SearchAgreementsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *marketplaceagreement.SearchAgreementsInput, ...func(*marketplaceagreement.Options)) (*marketplaceagreement.SearchAgreementsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *marketplaceagreement.SearchAgreementsInput, ...func(*marketplaceagreement.Options)) *marketplaceagreement.SearchAgreementsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*marketplaceagreement.SearchAgreementsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *marketplaceagreement.SearchAgreementsInput, ...func(*marketplaceagreement.Options)) error); ok {
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
