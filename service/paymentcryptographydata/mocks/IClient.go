// Code generated by mockery v2.53.4. DO NOT EDIT.

package mocks

import (
	context "context"

	paymentcryptographydata "github.com/aws/aws-sdk-go-v2/service/paymentcryptographydata"
	mock "github.com/stretchr/testify/mock"
)

// IClient is an autogenerated mock type for the IClient type
type IClient struct {
	mock.Mock
}

// DecryptData provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DecryptData(ctx context.Context, params *paymentcryptographydata.DecryptDataInput, optFns ...func(*paymentcryptographydata.Options)) (*paymentcryptographydata.DecryptDataOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DecryptData")
	}

	var r0 *paymentcryptographydata.DecryptDataOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *paymentcryptographydata.DecryptDataInput, ...func(*paymentcryptographydata.Options)) (*paymentcryptographydata.DecryptDataOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *paymentcryptographydata.DecryptDataInput, ...func(*paymentcryptographydata.Options)) *paymentcryptographydata.DecryptDataOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*paymentcryptographydata.DecryptDataOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *paymentcryptographydata.DecryptDataInput, ...func(*paymentcryptographydata.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EncryptData provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) EncryptData(ctx context.Context, params *paymentcryptographydata.EncryptDataInput, optFns ...func(*paymentcryptographydata.Options)) (*paymentcryptographydata.EncryptDataOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for EncryptData")
	}

	var r0 *paymentcryptographydata.EncryptDataOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *paymentcryptographydata.EncryptDataInput, ...func(*paymentcryptographydata.Options)) (*paymentcryptographydata.EncryptDataOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *paymentcryptographydata.EncryptDataInput, ...func(*paymentcryptographydata.Options)) *paymentcryptographydata.EncryptDataOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*paymentcryptographydata.EncryptDataOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *paymentcryptographydata.EncryptDataInput, ...func(*paymentcryptographydata.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GenerateCardValidationData provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GenerateCardValidationData(ctx context.Context, params *paymentcryptographydata.GenerateCardValidationDataInput, optFns ...func(*paymentcryptographydata.Options)) (*paymentcryptographydata.GenerateCardValidationDataOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GenerateCardValidationData")
	}

	var r0 *paymentcryptographydata.GenerateCardValidationDataOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *paymentcryptographydata.GenerateCardValidationDataInput, ...func(*paymentcryptographydata.Options)) (*paymentcryptographydata.GenerateCardValidationDataOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *paymentcryptographydata.GenerateCardValidationDataInput, ...func(*paymentcryptographydata.Options)) *paymentcryptographydata.GenerateCardValidationDataOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*paymentcryptographydata.GenerateCardValidationDataOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *paymentcryptographydata.GenerateCardValidationDataInput, ...func(*paymentcryptographydata.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GenerateMac provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GenerateMac(ctx context.Context, params *paymentcryptographydata.GenerateMacInput, optFns ...func(*paymentcryptographydata.Options)) (*paymentcryptographydata.GenerateMacOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GenerateMac")
	}

	var r0 *paymentcryptographydata.GenerateMacOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *paymentcryptographydata.GenerateMacInput, ...func(*paymentcryptographydata.Options)) (*paymentcryptographydata.GenerateMacOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *paymentcryptographydata.GenerateMacInput, ...func(*paymentcryptographydata.Options)) *paymentcryptographydata.GenerateMacOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*paymentcryptographydata.GenerateMacOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *paymentcryptographydata.GenerateMacInput, ...func(*paymentcryptographydata.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GenerateMacEmvPinChange provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GenerateMacEmvPinChange(ctx context.Context, params *paymentcryptographydata.GenerateMacEmvPinChangeInput, optFns ...func(*paymentcryptographydata.Options)) (*paymentcryptographydata.GenerateMacEmvPinChangeOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GenerateMacEmvPinChange")
	}

	var r0 *paymentcryptographydata.GenerateMacEmvPinChangeOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *paymentcryptographydata.GenerateMacEmvPinChangeInput, ...func(*paymentcryptographydata.Options)) (*paymentcryptographydata.GenerateMacEmvPinChangeOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *paymentcryptographydata.GenerateMacEmvPinChangeInput, ...func(*paymentcryptographydata.Options)) *paymentcryptographydata.GenerateMacEmvPinChangeOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*paymentcryptographydata.GenerateMacEmvPinChangeOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *paymentcryptographydata.GenerateMacEmvPinChangeInput, ...func(*paymentcryptographydata.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GeneratePinData provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GeneratePinData(ctx context.Context, params *paymentcryptographydata.GeneratePinDataInput, optFns ...func(*paymentcryptographydata.Options)) (*paymentcryptographydata.GeneratePinDataOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GeneratePinData")
	}

	var r0 *paymentcryptographydata.GeneratePinDataOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *paymentcryptographydata.GeneratePinDataInput, ...func(*paymentcryptographydata.Options)) (*paymentcryptographydata.GeneratePinDataOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *paymentcryptographydata.GeneratePinDataInput, ...func(*paymentcryptographydata.Options)) *paymentcryptographydata.GeneratePinDataOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*paymentcryptographydata.GeneratePinDataOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *paymentcryptographydata.GeneratePinDataInput, ...func(*paymentcryptographydata.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Options provides a mock function with no fields
func (_m *IClient) Options() paymentcryptographydata.Options {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Options")
	}

	var r0 paymentcryptographydata.Options
	if rf, ok := ret.Get(0).(func() paymentcryptographydata.Options); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(paymentcryptographydata.Options)
	}

	return r0
}

// ReEncryptData provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ReEncryptData(ctx context.Context, params *paymentcryptographydata.ReEncryptDataInput, optFns ...func(*paymentcryptographydata.Options)) (*paymentcryptographydata.ReEncryptDataOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ReEncryptData")
	}

	var r0 *paymentcryptographydata.ReEncryptDataOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *paymentcryptographydata.ReEncryptDataInput, ...func(*paymentcryptographydata.Options)) (*paymentcryptographydata.ReEncryptDataOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *paymentcryptographydata.ReEncryptDataInput, ...func(*paymentcryptographydata.Options)) *paymentcryptographydata.ReEncryptDataOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*paymentcryptographydata.ReEncryptDataOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *paymentcryptographydata.ReEncryptDataInput, ...func(*paymentcryptographydata.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TranslatePinData provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) TranslatePinData(ctx context.Context, params *paymentcryptographydata.TranslatePinDataInput, optFns ...func(*paymentcryptographydata.Options)) (*paymentcryptographydata.TranslatePinDataOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for TranslatePinData")
	}

	var r0 *paymentcryptographydata.TranslatePinDataOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *paymentcryptographydata.TranslatePinDataInput, ...func(*paymentcryptographydata.Options)) (*paymentcryptographydata.TranslatePinDataOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *paymentcryptographydata.TranslatePinDataInput, ...func(*paymentcryptographydata.Options)) *paymentcryptographydata.TranslatePinDataOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*paymentcryptographydata.TranslatePinDataOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *paymentcryptographydata.TranslatePinDataInput, ...func(*paymentcryptographydata.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// VerifyAuthRequestCryptogram provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) VerifyAuthRequestCryptogram(ctx context.Context, params *paymentcryptographydata.VerifyAuthRequestCryptogramInput, optFns ...func(*paymentcryptographydata.Options)) (*paymentcryptographydata.VerifyAuthRequestCryptogramOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for VerifyAuthRequestCryptogram")
	}

	var r0 *paymentcryptographydata.VerifyAuthRequestCryptogramOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *paymentcryptographydata.VerifyAuthRequestCryptogramInput, ...func(*paymentcryptographydata.Options)) (*paymentcryptographydata.VerifyAuthRequestCryptogramOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *paymentcryptographydata.VerifyAuthRequestCryptogramInput, ...func(*paymentcryptographydata.Options)) *paymentcryptographydata.VerifyAuthRequestCryptogramOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*paymentcryptographydata.VerifyAuthRequestCryptogramOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *paymentcryptographydata.VerifyAuthRequestCryptogramInput, ...func(*paymentcryptographydata.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// VerifyCardValidationData provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) VerifyCardValidationData(ctx context.Context, params *paymentcryptographydata.VerifyCardValidationDataInput, optFns ...func(*paymentcryptographydata.Options)) (*paymentcryptographydata.VerifyCardValidationDataOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for VerifyCardValidationData")
	}

	var r0 *paymentcryptographydata.VerifyCardValidationDataOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *paymentcryptographydata.VerifyCardValidationDataInput, ...func(*paymentcryptographydata.Options)) (*paymentcryptographydata.VerifyCardValidationDataOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *paymentcryptographydata.VerifyCardValidationDataInput, ...func(*paymentcryptographydata.Options)) *paymentcryptographydata.VerifyCardValidationDataOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*paymentcryptographydata.VerifyCardValidationDataOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *paymentcryptographydata.VerifyCardValidationDataInput, ...func(*paymentcryptographydata.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// VerifyMac provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) VerifyMac(ctx context.Context, params *paymentcryptographydata.VerifyMacInput, optFns ...func(*paymentcryptographydata.Options)) (*paymentcryptographydata.VerifyMacOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for VerifyMac")
	}

	var r0 *paymentcryptographydata.VerifyMacOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *paymentcryptographydata.VerifyMacInput, ...func(*paymentcryptographydata.Options)) (*paymentcryptographydata.VerifyMacOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *paymentcryptographydata.VerifyMacInput, ...func(*paymentcryptographydata.Options)) *paymentcryptographydata.VerifyMacOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*paymentcryptographydata.VerifyMacOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *paymentcryptographydata.VerifyMacInput, ...func(*paymentcryptographydata.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// VerifyPinData provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) VerifyPinData(ctx context.Context, params *paymentcryptographydata.VerifyPinDataInput, optFns ...func(*paymentcryptographydata.Options)) (*paymentcryptographydata.VerifyPinDataOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for VerifyPinData")
	}

	var r0 *paymentcryptographydata.VerifyPinDataOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *paymentcryptographydata.VerifyPinDataInput, ...func(*paymentcryptographydata.Options)) (*paymentcryptographydata.VerifyPinDataOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *paymentcryptographydata.VerifyPinDataInput, ...func(*paymentcryptographydata.Options)) *paymentcryptographydata.VerifyPinDataOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*paymentcryptographydata.VerifyPinDataOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *paymentcryptographydata.VerifyPinDataInput, ...func(*paymentcryptographydata.Options)) error); ok {
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
