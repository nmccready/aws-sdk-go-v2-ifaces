package route53domains_test

// tests for the route53domains service interface for this ../../../service/route53domains/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/route53domains"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/route53domains/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/route53domains/route53domains_iface"
	"github.com/stretchr/testify/assert"
)

func TestRoute53domainsServiceCanBeMocked(t *testing.T) {
	var iface route53domains_iface.IClient
	iface = &route53domains.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := route53domains.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAcceptDomainTransferFromAnotherAwsAccount", func(t *testing.T) {
        input := &route53domains.AcceptDomainTransferFromAnotherAwsAccountInput{}
        output := &route53domains.AcceptDomainTransferFromAnotherAwsAccountOutput{}

        mockClient.On("AcceptDomainTransferFromAnotherAwsAccount", ctx, input).Return(output, nil)

        result, err := mockClient.AcceptDomainTransferFromAnotherAwsAccount(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateDelegationSignerToDomain", func(t *testing.T) {
        input := &route53domains.AssociateDelegationSignerToDomainInput{}
        output := &route53domains.AssociateDelegationSignerToDomainOutput{}

        mockClient.On("AssociateDelegationSignerToDomain", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateDelegationSignerToDomain(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCancelDomainTransferToAnotherAwsAccount", func(t *testing.T) {
        input := &route53domains.CancelDomainTransferToAnotherAwsAccountInput{}
        output := &route53domains.CancelDomainTransferToAnotherAwsAccountOutput{}

        mockClient.On("CancelDomainTransferToAnotherAwsAccount", ctx, input).Return(output, nil)

        result, err := mockClient.CancelDomainTransferToAnotherAwsAccount(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCheckDomainAvailability", func(t *testing.T) {
        input := &route53domains.CheckDomainAvailabilityInput{}
        output := &route53domains.CheckDomainAvailabilityOutput{}

        mockClient.On("CheckDomainAvailability", ctx, input).Return(output, nil)

        result, err := mockClient.CheckDomainAvailability(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCheckDomainTransferability", func(t *testing.T) {
        input := &route53domains.CheckDomainTransferabilityInput{}
        output := &route53domains.CheckDomainTransferabilityOutput{}

        mockClient.On("CheckDomainTransferability", ctx, input).Return(output, nil)

        result, err := mockClient.CheckDomainTransferability(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteDomain", func(t *testing.T) {
        input := &route53domains.DeleteDomainInput{}
        output := &route53domains.DeleteDomainOutput{}

        mockClient.On("DeleteDomain", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteDomain(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteTagsForDomain", func(t *testing.T) {
        input := &route53domains.DeleteTagsForDomainInput{}
        output := &route53domains.DeleteTagsForDomainOutput{}

        mockClient.On("DeleteTagsForDomain", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteTagsForDomain(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisableDomainAutoRenew", func(t *testing.T) {
        input := &route53domains.DisableDomainAutoRenewInput{}
        output := &route53domains.DisableDomainAutoRenewOutput{}

        mockClient.On("DisableDomainAutoRenew", ctx, input).Return(output, nil)

        result, err := mockClient.DisableDomainAutoRenew(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisableDomainTransferLock", func(t *testing.T) {
        input := &route53domains.DisableDomainTransferLockInput{}
        output := &route53domains.DisableDomainTransferLockOutput{}

        mockClient.On("DisableDomainTransferLock", ctx, input).Return(output, nil)

        result, err := mockClient.DisableDomainTransferLock(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateDelegationSignerFromDomain", func(t *testing.T) {
        input := &route53domains.DisassociateDelegationSignerFromDomainInput{}
        output := &route53domains.DisassociateDelegationSignerFromDomainOutput{}

        mockClient.On("DisassociateDelegationSignerFromDomain", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateDelegationSignerFromDomain(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestEnableDomainAutoRenew", func(t *testing.T) {
        input := &route53domains.EnableDomainAutoRenewInput{}
        output := &route53domains.EnableDomainAutoRenewOutput{}

        mockClient.On("EnableDomainAutoRenew", ctx, input).Return(output, nil)

        result, err := mockClient.EnableDomainAutoRenew(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestEnableDomainTransferLock", func(t *testing.T) {
        input := &route53domains.EnableDomainTransferLockInput{}
        output := &route53domains.EnableDomainTransferLockOutput{}

        mockClient.On("EnableDomainTransferLock", ctx, input).Return(output, nil)

        result, err := mockClient.EnableDomainTransferLock(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetContactReachabilityStatus", func(t *testing.T) {
        input := &route53domains.GetContactReachabilityStatusInput{}
        output := &route53domains.GetContactReachabilityStatusOutput{}

        mockClient.On("GetContactReachabilityStatus", ctx, input).Return(output, nil)

        result, err := mockClient.GetContactReachabilityStatus(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDomainDetail", func(t *testing.T) {
        input := &route53domains.GetDomainDetailInput{}
        output := &route53domains.GetDomainDetailOutput{}

        mockClient.On("GetDomainDetail", ctx, input).Return(output, nil)

        result, err := mockClient.GetDomainDetail(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDomainSuggestions", func(t *testing.T) {
        input := &route53domains.GetDomainSuggestionsInput{}
        output := &route53domains.GetDomainSuggestionsOutput{}

        mockClient.On("GetDomainSuggestions", ctx, input).Return(output, nil)

        result, err := mockClient.GetDomainSuggestions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetOperationDetail", func(t *testing.T) {
        input := &route53domains.GetOperationDetailInput{}
        output := &route53domains.GetOperationDetailOutput{}

        mockClient.On("GetOperationDetail", ctx, input).Return(output, nil)

        result, err := mockClient.GetOperationDetail(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDomains", func(t *testing.T) {
        input := &route53domains.ListDomainsInput{}
        output := &route53domains.ListDomainsOutput{}

        mockClient.On("ListDomains", ctx, input).Return(output, nil)

        result, err := mockClient.ListDomains(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListOperations", func(t *testing.T) {
        input := &route53domains.ListOperationsInput{}
        output := &route53domains.ListOperationsOutput{}

        mockClient.On("ListOperations", ctx, input).Return(output, nil)

        result, err := mockClient.ListOperations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListPrices", func(t *testing.T) {
        input := &route53domains.ListPricesInput{}
        output := &route53domains.ListPricesOutput{}

        mockClient.On("ListPrices", ctx, input).Return(output, nil)

        result, err := mockClient.ListPrices(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForDomain", func(t *testing.T) {
        input := &route53domains.ListTagsForDomainInput{}
        output := &route53domains.ListTagsForDomainOutput{}

        mockClient.On("ListTagsForDomain", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForDomain(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPushDomain", func(t *testing.T) {
        input := &route53domains.PushDomainInput{}
        output := &route53domains.PushDomainOutput{}

        mockClient.On("PushDomain", ctx, input).Return(output, nil)

        result, err := mockClient.PushDomain(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRegisterDomain", func(t *testing.T) {
        input := &route53domains.RegisterDomainInput{}
        output := &route53domains.RegisterDomainOutput{}

        mockClient.On("RegisterDomain", ctx, input).Return(output, nil)

        result, err := mockClient.RegisterDomain(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRejectDomainTransferFromAnotherAwsAccount", func(t *testing.T) {
        input := &route53domains.RejectDomainTransferFromAnotherAwsAccountInput{}
        output := &route53domains.RejectDomainTransferFromAnotherAwsAccountOutput{}

        mockClient.On("RejectDomainTransferFromAnotherAwsAccount", ctx, input).Return(output, nil)

        result, err := mockClient.RejectDomainTransferFromAnotherAwsAccount(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRenewDomain", func(t *testing.T) {
        input := &route53domains.RenewDomainInput{}
        output := &route53domains.RenewDomainOutput{}

        mockClient.On("RenewDomain", ctx, input).Return(output, nil)

        result, err := mockClient.RenewDomain(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestResendContactReachabilityEmail", func(t *testing.T) {
        input := &route53domains.ResendContactReachabilityEmailInput{}
        output := &route53domains.ResendContactReachabilityEmailOutput{}

        mockClient.On("ResendContactReachabilityEmail", ctx, input).Return(output, nil)

        result, err := mockClient.ResendContactReachabilityEmail(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestResendOperationAuthorization", func(t *testing.T) {
        input := &route53domains.ResendOperationAuthorizationInput{}
        output := &route53domains.ResendOperationAuthorizationOutput{}

        mockClient.On("ResendOperationAuthorization", ctx, input).Return(output, nil)

        result, err := mockClient.ResendOperationAuthorization(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRetrieveDomainAuthCode", func(t *testing.T) {
        input := &route53domains.RetrieveDomainAuthCodeInput{}
        output := &route53domains.RetrieveDomainAuthCodeOutput{}

        mockClient.On("RetrieveDomainAuthCode", ctx, input).Return(output, nil)

        result, err := mockClient.RetrieveDomainAuthCode(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTransferDomain", func(t *testing.T) {
        input := &route53domains.TransferDomainInput{}
        output := &route53domains.TransferDomainOutput{}

        mockClient.On("TransferDomain", ctx, input).Return(output, nil)

        result, err := mockClient.TransferDomain(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTransferDomainToAnotherAwsAccount", func(t *testing.T) {
        input := &route53domains.TransferDomainToAnotherAwsAccountInput{}
        output := &route53domains.TransferDomainToAnotherAwsAccountOutput{}

        mockClient.On("TransferDomainToAnotherAwsAccount", ctx, input).Return(output, nil)

        result, err := mockClient.TransferDomainToAnotherAwsAccount(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateDomainContact", func(t *testing.T) {
        input := &route53domains.UpdateDomainContactInput{}
        output := &route53domains.UpdateDomainContactOutput{}

        mockClient.On("UpdateDomainContact", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateDomainContact(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateDomainContactPrivacy", func(t *testing.T) {
        input := &route53domains.UpdateDomainContactPrivacyInput{}
        output := &route53domains.UpdateDomainContactPrivacyOutput{}

        mockClient.On("UpdateDomainContactPrivacy", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateDomainContactPrivacy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateDomainNameservers", func(t *testing.T) {
        input := &route53domains.UpdateDomainNameserversInput{}
        output := &route53domains.UpdateDomainNameserversOutput{}

        mockClient.On("UpdateDomainNameservers", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateDomainNameservers(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateTagsForDomain", func(t *testing.T) {
        input := &route53domains.UpdateTagsForDomainInput{}
        output := &route53domains.UpdateTagsForDomainOutput{}

        mockClient.On("UpdateTagsForDomain", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateTagsForDomain(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestViewBilling", func(t *testing.T) {
        input := &route53domains.ViewBillingInput{}
        output := &route53domains.ViewBillingOutput{}

        mockClient.On("ViewBilling", ctx, input).Return(output, nil)

        result, err := mockClient.ViewBilling(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
