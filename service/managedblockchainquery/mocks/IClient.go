// Code generated by mockery v2.39.2. DO NOT EDIT.

package mocks

import (
	context "context"

	managedblockchainquery "github.com/aws/aws-sdk-go-v2/service/managedblockchainquery"

	mock "github.com/stretchr/testify/mock"
)

// IClient is an autogenerated mock type for the IClient type
type IClient struct {
	mock.Mock
}

// BatchGetTokenBalance provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) BatchGetTokenBalance(ctx context.Context, params *managedblockchainquery.BatchGetTokenBalanceInput, optFns ...func(*managedblockchainquery.Options)) (*managedblockchainquery.BatchGetTokenBalanceOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for BatchGetTokenBalance")
	}

	var r0 *managedblockchainquery.BatchGetTokenBalanceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *managedblockchainquery.BatchGetTokenBalanceInput, ...func(*managedblockchainquery.Options)) (*managedblockchainquery.BatchGetTokenBalanceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *managedblockchainquery.BatchGetTokenBalanceInput, ...func(*managedblockchainquery.Options)) *managedblockchainquery.BatchGetTokenBalanceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*managedblockchainquery.BatchGetTokenBalanceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *managedblockchainquery.BatchGetTokenBalanceInput, ...func(*managedblockchainquery.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAssetContract provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetAssetContract(ctx context.Context, params *managedblockchainquery.GetAssetContractInput, optFns ...func(*managedblockchainquery.Options)) (*managedblockchainquery.GetAssetContractOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetAssetContract")
	}

	var r0 *managedblockchainquery.GetAssetContractOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *managedblockchainquery.GetAssetContractInput, ...func(*managedblockchainquery.Options)) (*managedblockchainquery.GetAssetContractOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *managedblockchainquery.GetAssetContractInput, ...func(*managedblockchainquery.Options)) *managedblockchainquery.GetAssetContractOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*managedblockchainquery.GetAssetContractOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *managedblockchainquery.GetAssetContractInput, ...func(*managedblockchainquery.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTokenBalance provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetTokenBalance(ctx context.Context, params *managedblockchainquery.GetTokenBalanceInput, optFns ...func(*managedblockchainquery.Options)) (*managedblockchainquery.GetTokenBalanceOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetTokenBalance")
	}

	var r0 *managedblockchainquery.GetTokenBalanceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *managedblockchainquery.GetTokenBalanceInput, ...func(*managedblockchainquery.Options)) (*managedblockchainquery.GetTokenBalanceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *managedblockchainquery.GetTokenBalanceInput, ...func(*managedblockchainquery.Options)) *managedblockchainquery.GetTokenBalanceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*managedblockchainquery.GetTokenBalanceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *managedblockchainquery.GetTokenBalanceInput, ...func(*managedblockchainquery.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTransaction provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) GetTransaction(ctx context.Context, params *managedblockchainquery.GetTransactionInput, optFns ...func(*managedblockchainquery.Options)) (*managedblockchainquery.GetTransactionOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetTransaction")
	}

	var r0 *managedblockchainquery.GetTransactionOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *managedblockchainquery.GetTransactionInput, ...func(*managedblockchainquery.Options)) (*managedblockchainquery.GetTransactionOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *managedblockchainquery.GetTransactionInput, ...func(*managedblockchainquery.Options)) *managedblockchainquery.GetTransactionOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*managedblockchainquery.GetTransactionOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *managedblockchainquery.GetTransactionInput, ...func(*managedblockchainquery.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListAssetContracts provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListAssetContracts(ctx context.Context, params *managedblockchainquery.ListAssetContractsInput, optFns ...func(*managedblockchainquery.Options)) (*managedblockchainquery.ListAssetContractsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListAssetContracts")
	}

	var r0 *managedblockchainquery.ListAssetContractsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *managedblockchainquery.ListAssetContractsInput, ...func(*managedblockchainquery.Options)) (*managedblockchainquery.ListAssetContractsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *managedblockchainquery.ListAssetContractsInput, ...func(*managedblockchainquery.Options)) *managedblockchainquery.ListAssetContractsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*managedblockchainquery.ListAssetContractsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *managedblockchainquery.ListAssetContractsInput, ...func(*managedblockchainquery.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListFilteredTransactionEvents provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListFilteredTransactionEvents(ctx context.Context, params *managedblockchainquery.ListFilteredTransactionEventsInput, optFns ...func(*managedblockchainquery.Options)) (*managedblockchainquery.ListFilteredTransactionEventsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListFilteredTransactionEvents")
	}

	var r0 *managedblockchainquery.ListFilteredTransactionEventsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *managedblockchainquery.ListFilteredTransactionEventsInput, ...func(*managedblockchainquery.Options)) (*managedblockchainquery.ListFilteredTransactionEventsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *managedblockchainquery.ListFilteredTransactionEventsInput, ...func(*managedblockchainquery.Options)) *managedblockchainquery.ListFilteredTransactionEventsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*managedblockchainquery.ListFilteredTransactionEventsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *managedblockchainquery.ListFilteredTransactionEventsInput, ...func(*managedblockchainquery.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListTokenBalances provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListTokenBalances(ctx context.Context, params *managedblockchainquery.ListTokenBalancesInput, optFns ...func(*managedblockchainquery.Options)) (*managedblockchainquery.ListTokenBalancesOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListTokenBalances")
	}

	var r0 *managedblockchainquery.ListTokenBalancesOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *managedblockchainquery.ListTokenBalancesInput, ...func(*managedblockchainquery.Options)) (*managedblockchainquery.ListTokenBalancesOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *managedblockchainquery.ListTokenBalancesInput, ...func(*managedblockchainquery.Options)) *managedblockchainquery.ListTokenBalancesOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*managedblockchainquery.ListTokenBalancesOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *managedblockchainquery.ListTokenBalancesInput, ...func(*managedblockchainquery.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListTransactionEvents provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListTransactionEvents(ctx context.Context, params *managedblockchainquery.ListTransactionEventsInput, optFns ...func(*managedblockchainquery.Options)) (*managedblockchainquery.ListTransactionEventsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListTransactionEvents")
	}

	var r0 *managedblockchainquery.ListTransactionEventsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *managedblockchainquery.ListTransactionEventsInput, ...func(*managedblockchainquery.Options)) (*managedblockchainquery.ListTransactionEventsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *managedblockchainquery.ListTransactionEventsInput, ...func(*managedblockchainquery.Options)) *managedblockchainquery.ListTransactionEventsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*managedblockchainquery.ListTransactionEventsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *managedblockchainquery.ListTransactionEventsInput, ...func(*managedblockchainquery.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListTransactions provides a mock function with given fields: ctx, params, optFns
func (_m *IClient) ListTransactions(ctx context.Context, params *managedblockchainquery.ListTransactionsInput, optFns ...func(*managedblockchainquery.Options)) (*managedblockchainquery.ListTransactionsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListTransactions")
	}

	var r0 *managedblockchainquery.ListTransactionsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *managedblockchainquery.ListTransactionsInput, ...func(*managedblockchainquery.Options)) (*managedblockchainquery.ListTransactionsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *managedblockchainquery.ListTransactionsInput, ...func(*managedblockchainquery.Options)) *managedblockchainquery.ListTransactionsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*managedblockchainquery.ListTransactionsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *managedblockchainquery.ListTransactionsInput, ...func(*managedblockchainquery.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Options provides a mock function with given fields:
func (_m *IClient) Options() managedblockchainquery.Options {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Options")
	}

	var r0 managedblockchainquery.Options
	if rf, ok := ret.Get(0).(func() managedblockchainquery.Options); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(managedblockchainquery.Options)
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
