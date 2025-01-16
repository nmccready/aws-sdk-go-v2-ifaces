// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	context "context"

	taxsettings "github.com/aws/aws-sdk-go-v2/service/taxsettings"
	mock "github.com/stretchr/testify/mock"
)

// IClient is an autogenerated mock type for the IClient type
type IClient struct {
	mock.Mock
}

// BatchDeleteTaxRegistration provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) BatchDeleteTaxRegistration(ctx context.Context, params *taxsettings.BatchDeleteTaxRegistrationInput, optFns ...func(*taxsettings.Options)) (*taxsettings.BatchDeleteTaxRegistrationOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for BatchDeleteTaxRegistration")
	}

	var r0 *taxsettings.BatchDeleteTaxRegistrationOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *taxsettings.BatchDeleteTaxRegistrationInput, ...func(*taxsettings.Options)) (*taxsettings.BatchDeleteTaxRegistrationOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *taxsettings.BatchDeleteTaxRegistrationInput, ...func(*taxsettings.Options)) *taxsettings.BatchDeleteTaxRegistrationOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*taxsettings.BatchDeleteTaxRegistrationOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *taxsettings.BatchDeleteTaxRegistrationInput, ...func(*taxsettings.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BatchGetTaxExemptions provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) BatchGetTaxExemptions(ctx context.Context, params *taxsettings.BatchGetTaxExemptionsInput, optFns ...func(*taxsettings.Options)) (*taxsettings.BatchGetTaxExemptionsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for BatchGetTaxExemptions")
	}

	var r0 *taxsettings.BatchGetTaxExemptionsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *taxsettings.BatchGetTaxExemptionsInput, ...func(*taxsettings.Options)) (*taxsettings.BatchGetTaxExemptionsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *taxsettings.BatchGetTaxExemptionsInput, ...func(*taxsettings.Options)) *taxsettings.BatchGetTaxExemptionsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*taxsettings.BatchGetTaxExemptionsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *taxsettings.BatchGetTaxExemptionsInput, ...func(*taxsettings.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BatchPutTaxRegistration provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) BatchPutTaxRegistration(ctx context.Context, params *taxsettings.BatchPutTaxRegistrationInput, optFns ...func(*taxsettings.Options)) (*taxsettings.BatchPutTaxRegistrationOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for BatchPutTaxRegistration")
	}

	var r0 *taxsettings.BatchPutTaxRegistrationOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *taxsettings.BatchPutTaxRegistrationInput, ...func(*taxsettings.Options)) (*taxsettings.BatchPutTaxRegistrationOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *taxsettings.BatchPutTaxRegistrationInput, ...func(*taxsettings.Options)) *taxsettings.BatchPutTaxRegistrationOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*taxsettings.BatchPutTaxRegistrationOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *taxsettings.BatchPutTaxRegistrationInput, ...func(*taxsettings.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteSupplementalTaxRegistration provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DeleteSupplementalTaxRegistration(ctx context.Context, params *taxsettings.DeleteSupplementalTaxRegistrationInput, optFns ...func(*taxsettings.Options)) (*taxsettings.DeleteSupplementalTaxRegistrationOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteSupplementalTaxRegistration")
	}

	var r0 *taxsettings.DeleteSupplementalTaxRegistrationOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *taxsettings.DeleteSupplementalTaxRegistrationInput, ...func(*taxsettings.Options)) (*taxsettings.DeleteSupplementalTaxRegistrationOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *taxsettings.DeleteSupplementalTaxRegistrationInput, ...func(*taxsettings.Options)) *taxsettings.DeleteSupplementalTaxRegistrationOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*taxsettings.DeleteSupplementalTaxRegistrationOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *taxsettings.DeleteSupplementalTaxRegistrationInput, ...func(*taxsettings.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteTaxRegistration provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) DeleteTaxRegistration(ctx context.Context, params *taxsettings.DeleteTaxRegistrationInput, optFns ...func(*taxsettings.Options)) (*taxsettings.DeleteTaxRegistrationOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteTaxRegistration")
	}

	var r0 *taxsettings.DeleteTaxRegistrationOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *taxsettings.DeleteTaxRegistrationInput, ...func(*taxsettings.Options)) (*taxsettings.DeleteTaxRegistrationOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *taxsettings.DeleteTaxRegistrationInput, ...func(*taxsettings.Options)) *taxsettings.DeleteTaxRegistrationOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*taxsettings.DeleteTaxRegistrationOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *taxsettings.DeleteTaxRegistrationInput, ...func(*taxsettings.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTaxExemptionTypes provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetTaxExemptionTypes(ctx context.Context, params *taxsettings.GetTaxExemptionTypesInput, optFns ...func(*taxsettings.Options)) (*taxsettings.GetTaxExemptionTypesOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetTaxExemptionTypes")
	}

	var r0 *taxsettings.GetTaxExemptionTypesOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *taxsettings.GetTaxExemptionTypesInput, ...func(*taxsettings.Options)) (*taxsettings.GetTaxExemptionTypesOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *taxsettings.GetTaxExemptionTypesInput, ...func(*taxsettings.Options)) *taxsettings.GetTaxExemptionTypesOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*taxsettings.GetTaxExemptionTypesOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *taxsettings.GetTaxExemptionTypesInput, ...func(*taxsettings.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTaxInheritance provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetTaxInheritance(ctx context.Context, params *taxsettings.GetTaxInheritanceInput, optFns ...func(*taxsettings.Options)) (*taxsettings.GetTaxInheritanceOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetTaxInheritance")
	}

	var r0 *taxsettings.GetTaxInheritanceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *taxsettings.GetTaxInheritanceInput, ...func(*taxsettings.Options)) (*taxsettings.GetTaxInheritanceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *taxsettings.GetTaxInheritanceInput, ...func(*taxsettings.Options)) *taxsettings.GetTaxInheritanceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*taxsettings.GetTaxInheritanceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *taxsettings.GetTaxInheritanceInput, ...func(*taxsettings.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTaxRegistration provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetTaxRegistration(ctx context.Context, params *taxsettings.GetTaxRegistrationInput, optFns ...func(*taxsettings.Options)) (*taxsettings.GetTaxRegistrationOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetTaxRegistration")
	}

	var r0 *taxsettings.GetTaxRegistrationOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *taxsettings.GetTaxRegistrationInput, ...func(*taxsettings.Options)) (*taxsettings.GetTaxRegistrationOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *taxsettings.GetTaxRegistrationInput, ...func(*taxsettings.Options)) *taxsettings.GetTaxRegistrationOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*taxsettings.GetTaxRegistrationOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *taxsettings.GetTaxRegistrationInput, ...func(*taxsettings.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTaxRegistrationDocument provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetTaxRegistrationDocument(ctx context.Context, params *taxsettings.GetTaxRegistrationDocumentInput, optFns ...func(*taxsettings.Options)) (*taxsettings.GetTaxRegistrationDocumentOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetTaxRegistrationDocument")
	}

	var r0 *taxsettings.GetTaxRegistrationDocumentOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *taxsettings.GetTaxRegistrationDocumentInput, ...func(*taxsettings.Options)) (*taxsettings.GetTaxRegistrationDocumentOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *taxsettings.GetTaxRegistrationDocumentInput, ...func(*taxsettings.Options)) *taxsettings.GetTaxRegistrationDocumentOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*taxsettings.GetTaxRegistrationDocumentOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *taxsettings.GetTaxRegistrationDocumentInput, ...func(*taxsettings.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListSupplementalTaxRegistrations provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListSupplementalTaxRegistrations(ctx context.Context, params *taxsettings.ListSupplementalTaxRegistrationsInput, optFns ...func(*taxsettings.Options)) (*taxsettings.ListSupplementalTaxRegistrationsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListSupplementalTaxRegistrations")
	}

	var r0 *taxsettings.ListSupplementalTaxRegistrationsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *taxsettings.ListSupplementalTaxRegistrationsInput, ...func(*taxsettings.Options)) (*taxsettings.ListSupplementalTaxRegistrationsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *taxsettings.ListSupplementalTaxRegistrationsInput, ...func(*taxsettings.Options)) *taxsettings.ListSupplementalTaxRegistrationsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*taxsettings.ListSupplementalTaxRegistrationsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *taxsettings.ListSupplementalTaxRegistrationsInput, ...func(*taxsettings.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListTaxExemptions provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListTaxExemptions(ctx context.Context, params *taxsettings.ListTaxExemptionsInput, optFns ...func(*taxsettings.Options)) (*taxsettings.ListTaxExemptionsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListTaxExemptions")
	}

	var r0 *taxsettings.ListTaxExemptionsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *taxsettings.ListTaxExemptionsInput, ...func(*taxsettings.Options)) (*taxsettings.ListTaxExemptionsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *taxsettings.ListTaxExemptionsInput, ...func(*taxsettings.Options)) *taxsettings.ListTaxExemptionsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*taxsettings.ListTaxExemptionsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *taxsettings.ListTaxExemptionsInput, ...func(*taxsettings.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListTaxRegistrations provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListTaxRegistrations(ctx context.Context, params *taxsettings.ListTaxRegistrationsInput, optFns ...func(*taxsettings.Options)) (*taxsettings.ListTaxRegistrationsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListTaxRegistrations")
	}

	var r0 *taxsettings.ListTaxRegistrationsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *taxsettings.ListTaxRegistrationsInput, ...func(*taxsettings.Options)) (*taxsettings.ListTaxRegistrationsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *taxsettings.ListTaxRegistrationsInput, ...func(*taxsettings.Options)) *taxsettings.ListTaxRegistrationsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*taxsettings.ListTaxRegistrationsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *taxsettings.ListTaxRegistrationsInput, ...func(*taxsettings.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Options provides a mock function with given fields:
func (_m *IClient) Options() taxsettings.Options {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Options")
	}

	var r0 taxsettings.Options
	if rf, ok := ret.Get(0).(func() taxsettings.Options); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(taxsettings.Options)
	}

	return r0
}

// PutSupplementalTaxRegistration provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) PutSupplementalTaxRegistration(ctx context.Context, params *taxsettings.PutSupplementalTaxRegistrationInput, optFns ...func(*taxsettings.Options)) (*taxsettings.PutSupplementalTaxRegistrationOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for PutSupplementalTaxRegistration")
	}

	var r0 *taxsettings.PutSupplementalTaxRegistrationOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *taxsettings.PutSupplementalTaxRegistrationInput, ...func(*taxsettings.Options)) (*taxsettings.PutSupplementalTaxRegistrationOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *taxsettings.PutSupplementalTaxRegistrationInput, ...func(*taxsettings.Options)) *taxsettings.PutSupplementalTaxRegistrationOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*taxsettings.PutSupplementalTaxRegistrationOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *taxsettings.PutSupplementalTaxRegistrationInput, ...func(*taxsettings.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PutTaxExemption provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) PutTaxExemption(ctx context.Context, params *taxsettings.PutTaxExemptionInput, optFns ...func(*taxsettings.Options)) (*taxsettings.PutTaxExemptionOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for PutTaxExemption")
	}

	var r0 *taxsettings.PutTaxExemptionOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *taxsettings.PutTaxExemptionInput, ...func(*taxsettings.Options)) (*taxsettings.PutTaxExemptionOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *taxsettings.PutTaxExemptionInput, ...func(*taxsettings.Options)) *taxsettings.PutTaxExemptionOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*taxsettings.PutTaxExemptionOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *taxsettings.PutTaxExemptionInput, ...func(*taxsettings.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PutTaxInheritance provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) PutTaxInheritance(ctx context.Context, params *taxsettings.PutTaxInheritanceInput, optFns ...func(*taxsettings.Options)) (*taxsettings.PutTaxInheritanceOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for PutTaxInheritance")
	}

	var r0 *taxsettings.PutTaxInheritanceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *taxsettings.PutTaxInheritanceInput, ...func(*taxsettings.Options)) (*taxsettings.PutTaxInheritanceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *taxsettings.PutTaxInheritanceInput, ...func(*taxsettings.Options)) *taxsettings.PutTaxInheritanceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*taxsettings.PutTaxInheritanceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *taxsettings.PutTaxInheritanceInput, ...func(*taxsettings.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PutTaxRegistration provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) PutTaxRegistration(ctx context.Context, params *taxsettings.PutTaxRegistrationInput, optFns ...func(*taxsettings.Options)) (*taxsettings.PutTaxRegistrationOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for PutTaxRegistration")
	}

	var r0 *taxsettings.PutTaxRegistrationOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *taxsettings.PutTaxRegistrationInput, ...func(*taxsettings.Options)) (*taxsettings.PutTaxRegistrationOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *taxsettings.PutTaxRegistrationInput, ...func(*taxsettings.Options)) *taxsettings.PutTaxRegistrationOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*taxsettings.PutTaxRegistrationOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *taxsettings.PutTaxRegistrationInput, ...func(*taxsettings.Options)) error); ok {
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
