package organizations_test

// tests for the organizations service interface for this ../../../service/organizations/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/organizations"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/organizations/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/organizations/organizations_iface"
	"github.com/stretchr/testify/assert"
)

func TestOrganizationsServiceCanBeMocked(t *testing.T) {
	var iface organizations_iface.IClient
	iface = &organizations.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := organizations.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAcceptHandshake", func(t *testing.T) {
        input := &organizations.AcceptHandshakeInput{}
        output := &organizations.AcceptHandshakeOutput{}

        mockClient.On("AcceptHandshake", ctx, input).Return(output, nil)

        result, err := mockClient.AcceptHandshake(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAttachPolicy", func(t *testing.T) {
        input := &organizations.AttachPolicyInput{}
        output := &organizations.AttachPolicyOutput{}

        mockClient.On("AttachPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.AttachPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCancelHandshake", func(t *testing.T) {
        input := &organizations.CancelHandshakeInput{}
        output := &organizations.CancelHandshakeOutput{}

        mockClient.On("CancelHandshake", ctx, input).Return(output, nil)

        result, err := mockClient.CancelHandshake(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCloseAccount", func(t *testing.T) {
        input := &organizations.CloseAccountInput{}
        output := &organizations.CloseAccountOutput{}

        mockClient.On("CloseAccount", ctx, input).Return(output, nil)

        result, err := mockClient.CloseAccount(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateAccount", func(t *testing.T) {
        input := &organizations.CreateAccountInput{}
        output := &organizations.CreateAccountOutput{}

        mockClient.On("CreateAccount", ctx, input).Return(output, nil)

        result, err := mockClient.CreateAccount(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateGovCloudAccount", func(t *testing.T) {
        input := &organizations.CreateGovCloudAccountInput{}
        output := &organizations.CreateGovCloudAccountOutput{}

        mockClient.On("CreateGovCloudAccount", ctx, input).Return(output, nil)

        result, err := mockClient.CreateGovCloudAccount(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateOrganization", func(t *testing.T) {
        input := &organizations.CreateOrganizationInput{}
        output := &organizations.CreateOrganizationOutput{}

        mockClient.On("CreateOrganization", ctx, input).Return(output, nil)

        result, err := mockClient.CreateOrganization(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateOrganizationalUnit", func(t *testing.T) {
        input := &organizations.CreateOrganizationalUnitInput{}
        output := &organizations.CreateOrganizationalUnitOutput{}

        mockClient.On("CreateOrganizationalUnit", ctx, input).Return(output, nil)

        result, err := mockClient.CreateOrganizationalUnit(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreatePolicy", func(t *testing.T) {
        input := &organizations.CreatePolicyInput{}
        output := &organizations.CreatePolicyOutput{}

        mockClient.On("CreatePolicy", ctx, input).Return(output, nil)

        result, err := mockClient.CreatePolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeclineHandshake", func(t *testing.T) {
        input := &organizations.DeclineHandshakeInput{}
        output := &organizations.DeclineHandshakeOutput{}

        mockClient.On("DeclineHandshake", ctx, input).Return(output, nil)

        result, err := mockClient.DeclineHandshake(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteOrganization", func(t *testing.T) {
        input := &organizations.DeleteOrganizationInput{}
        output := &organizations.DeleteOrganizationOutput{}

        mockClient.On("DeleteOrganization", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteOrganization(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteOrganizationalUnit", func(t *testing.T) {
        input := &organizations.DeleteOrganizationalUnitInput{}
        output := &organizations.DeleteOrganizationalUnitOutput{}

        mockClient.On("DeleteOrganizationalUnit", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteOrganizationalUnit(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeletePolicy", func(t *testing.T) {
        input := &organizations.DeletePolicyInput{}
        output := &organizations.DeletePolicyOutput{}

        mockClient.On("DeletePolicy", ctx, input).Return(output, nil)

        result, err := mockClient.DeletePolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteResourcePolicy", func(t *testing.T) {
        input := &organizations.DeleteResourcePolicyInput{}
        output := &organizations.DeleteResourcePolicyOutput{}

        mockClient.On("DeleteResourcePolicy", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteResourcePolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeregisterDelegatedAdministrator", func(t *testing.T) {
        input := &organizations.DeregisterDelegatedAdministratorInput{}
        output := &organizations.DeregisterDelegatedAdministratorOutput{}

        mockClient.On("DeregisterDelegatedAdministrator", ctx, input).Return(output, nil)

        result, err := mockClient.DeregisterDelegatedAdministrator(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeAccount", func(t *testing.T) {
        input := &organizations.DescribeAccountInput{}
        output := &organizations.DescribeAccountOutput{}

        mockClient.On("DescribeAccount", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeAccount(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeCreateAccountStatus", func(t *testing.T) {
        input := &organizations.DescribeCreateAccountStatusInput{}
        output := &organizations.DescribeCreateAccountStatusOutput{}

        mockClient.On("DescribeCreateAccountStatus", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeCreateAccountStatus(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeEffectivePolicy", func(t *testing.T) {
        input := &organizations.DescribeEffectivePolicyInput{}
        output := &organizations.DescribeEffectivePolicyOutput{}

        mockClient.On("DescribeEffectivePolicy", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeEffectivePolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeHandshake", func(t *testing.T) {
        input := &organizations.DescribeHandshakeInput{}
        output := &organizations.DescribeHandshakeOutput{}

        mockClient.On("DescribeHandshake", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeHandshake(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeOrganization", func(t *testing.T) {
        input := &organizations.DescribeOrganizationInput{}
        output := &organizations.DescribeOrganizationOutput{}

        mockClient.On("DescribeOrganization", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeOrganization(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeOrganizationalUnit", func(t *testing.T) {
        input := &organizations.DescribeOrganizationalUnitInput{}
        output := &organizations.DescribeOrganizationalUnitOutput{}

        mockClient.On("DescribeOrganizationalUnit", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeOrganizationalUnit(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribePolicy", func(t *testing.T) {
        input := &organizations.DescribePolicyInput{}
        output := &organizations.DescribePolicyOutput{}

        mockClient.On("DescribePolicy", ctx, input).Return(output, nil)

        result, err := mockClient.DescribePolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeResourcePolicy", func(t *testing.T) {
        input := &organizations.DescribeResourcePolicyInput{}
        output := &organizations.DescribeResourcePolicyOutput{}

        mockClient.On("DescribeResourcePolicy", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeResourcePolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDetachPolicy", func(t *testing.T) {
        input := &organizations.DetachPolicyInput{}
        output := &organizations.DetachPolicyOutput{}

        mockClient.On("DetachPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.DetachPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisableAWSServiceAccess", func(t *testing.T) {
        input := &organizations.DisableAWSServiceAccessInput{}
        output := &organizations.DisableAWSServiceAccessOutput{}

        mockClient.On("DisableAWSServiceAccess", ctx, input).Return(output, nil)

        result, err := mockClient.DisableAWSServiceAccess(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisablePolicyType", func(t *testing.T) {
        input := &organizations.DisablePolicyTypeInput{}
        output := &organizations.DisablePolicyTypeOutput{}

        mockClient.On("DisablePolicyType", ctx, input).Return(output, nil)

        result, err := mockClient.DisablePolicyType(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestEnableAWSServiceAccess", func(t *testing.T) {
        input := &organizations.EnableAWSServiceAccessInput{}
        output := &organizations.EnableAWSServiceAccessOutput{}

        mockClient.On("EnableAWSServiceAccess", ctx, input).Return(output, nil)

        result, err := mockClient.EnableAWSServiceAccess(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestEnableAllFeatures", func(t *testing.T) {
        input := &organizations.EnableAllFeaturesInput{}
        output := &organizations.EnableAllFeaturesOutput{}

        mockClient.On("EnableAllFeatures", ctx, input).Return(output, nil)

        result, err := mockClient.EnableAllFeatures(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestEnablePolicyType", func(t *testing.T) {
        input := &organizations.EnablePolicyTypeInput{}
        output := &organizations.EnablePolicyTypeOutput{}

        mockClient.On("EnablePolicyType", ctx, input).Return(output, nil)

        result, err := mockClient.EnablePolicyType(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestInviteAccountToOrganization", func(t *testing.T) {
        input := &organizations.InviteAccountToOrganizationInput{}
        output := &organizations.InviteAccountToOrganizationOutput{}

        mockClient.On("InviteAccountToOrganization", ctx, input).Return(output, nil)

        result, err := mockClient.InviteAccountToOrganization(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestLeaveOrganization", func(t *testing.T) {
        input := &organizations.LeaveOrganizationInput{}
        output := &organizations.LeaveOrganizationOutput{}

        mockClient.On("LeaveOrganization", ctx, input).Return(output, nil)

        result, err := mockClient.LeaveOrganization(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAWSServiceAccessForOrganization", func(t *testing.T) {
        input := &organizations.ListAWSServiceAccessForOrganizationInput{}
        output := &organizations.ListAWSServiceAccessForOrganizationOutput{}

        mockClient.On("ListAWSServiceAccessForOrganization", ctx, input).Return(output, nil)

        result, err := mockClient.ListAWSServiceAccessForOrganization(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAccounts", func(t *testing.T) {
        input := &organizations.ListAccountsInput{}
        output := &organizations.ListAccountsOutput{}

        mockClient.On("ListAccounts", ctx, input).Return(output, nil)

        result, err := mockClient.ListAccounts(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAccountsForParent", func(t *testing.T) {
        input := &organizations.ListAccountsForParentInput{}
        output := &organizations.ListAccountsForParentOutput{}

        mockClient.On("ListAccountsForParent", ctx, input).Return(output, nil)

        result, err := mockClient.ListAccountsForParent(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListChildren", func(t *testing.T) {
        input := &organizations.ListChildrenInput{}
        output := &organizations.ListChildrenOutput{}

        mockClient.On("ListChildren", ctx, input).Return(output, nil)

        result, err := mockClient.ListChildren(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListCreateAccountStatus", func(t *testing.T) {
        input := &organizations.ListCreateAccountStatusInput{}
        output := &organizations.ListCreateAccountStatusOutput{}

        mockClient.On("ListCreateAccountStatus", ctx, input).Return(output, nil)

        result, err := mockClient.ListCreateAccountStatus(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDelegatedAdministrators", func(t *testing.T) {
        input := &organizations.ListDelegatedAdministratorsInput{}
        output := &organizations.ListDelegatedAdministratorsOutput{}

        mockClient.On("ListDelegatedAdministrators", ctx, input).Return(output, nil)

        result, err := mockClient.ListDelegatedAdministrators(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDelegatedServicesForAccount", func(t *testing.T) {
        input := &organizations.ListDelegatedServicesForAccountInput{}
        output := &organizations.ListDelegatedServicesForAccountOutput{}

        mockClient.On("ListDelegatedServicesForAccount", ctx, input).Return(output, nil)

        result, err := mockClient.ListDelegatedServicesForAccount(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListHandshakesForAccount", func(t *testing.T) {
        input := &organizations.ListHandshakesForAccountInput{}
        output := &organizations.ListHandshakesForAccountOutput{}

        mockClient.On("ListHandshakesForAccount", ctx, input).Return(output, nil)

        result, err := mockClient.ListHandshakesForAccount(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListHandshakesForOrganization", func(t *testing.T) {
        input := &organizations.ListHandshakesForOrganizationInput{}
        output := &organizations.ListHandshakesForOrganizationOutput{}

        mockClient.On("ListHandshakesForOrganization", ctx, input).Return(output, nil)

        result, err := mockClient.ListHandshakesForOrganization(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListOrganizationalUnitsForParent", func(t *testing.T) {
        input := &organizations.ListOrganizationalUnitsForParentInput{}
        output := &organizations.ListOrganizationalUnitsForParentOutput{}

        mockClient.On("ListOrganizationalUnitsForParent", ctx, input).Return(output, nil)

        result, err := mockClient.ListOrganizationalUnitsForParent(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListParents", func(t *testing.T) {
        input := &organizations.ListParentsInput{}
        output := &organizations.ListParentsOutput{}

        mockClient.On("ListParents", ctx, input).Return(output, nil)

        result, err := mockClient.ListParents(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListPolicies", func(t *testing.T) {
        input := &organizations.ListPoliciesInput{}
        output := &organizations.ListPoliciesOutput{}

        mockClient.On("ListPolicies", ctx, input).Return(output, nil)

        result, err := mockClient.ListPolicies(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListPoliciesForTarget", func(t *testing.T) {
        input := &organizations.ListPoliciesForTargetInput{}
        output := &organizations.ListPoliciesForTargetOutput{}

        mockClient.On("ListPoliciesForTarget", ctx, input).Return(output, nil)

        result, err := mockClient.ListPoliciesForTarget(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListRoots", func(t *testing.T) {
        input := &organizations.ListRootsInput{}
        output := &organizations.ListRootsOutput{}

        mockClient.On("ListRoots", ctx, input).Return(output, nil)

        result, err := mockClient.ListRoots(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &organizations.ListTagsForResourceInput{}
        output := &organizations.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTargetsForPolicy", func(t *testing.T) {
        input := &organizations.ListTargetsForPolicyInput{}
        output := &organizations.ListTargetsForPolicyOutput{}

        mockClient.On("ListTargetsForPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.ListTargetsForPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestMoveAccount", func(t *testing.T) {
        input := &organizations.MoveAccountInput{}
        output := &organizations.MoveAccountOutput{}

        mockClient.On("MoveAccount", ctx, input).Return(output, nil)

        result, err := mockClient.MoveAccount(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutResourcePolicy", func(t *testing.T) {
        input := &organizations.PutResourcePolicyInput{}
        output := &organizations.PutResourcePolicyOutput{}

        mockClient.On("PutResourcePolicy", ctx, input).Return(output, nil)

        result, err := mockClient.PutResourcePolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRegisterDelegatedAdministrator", func(t *testing.T) {
        input := &organizations.RegisterDelegatedAdministratorInput{}
        output := &organizations.RegisterDelegatedAdministratorOutput{}

        mockClient.On("RegisterDelegatedAdministrator", ctx, input).Return(output, nil)

        result, err := mockClient.RegisterDelegatedAdministrator(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRemoveAccountFromOrganization", func(t *testing.T) {
        input := &organizations.RemoveAccountFromOrganizationInput{}
        output := &organizations.RemoveAccountFromOrganizationOutput{}

        mockClient.On("RemoveAccountFromOrganization", ctx, input).Return(output, nil)

        result, err := mockClient.RemoveAccountFromOrganization(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &organizations.TagResourceInput{}
        output := &organizations.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &organizations.UntagResourceInput{}
        output := &organizations.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateOrganizationalUnit", func(t *testing.T) {
        input := &organizations.UpdateOrganizationalUnitInput{}
        output := &organizations.UpdateOrganizationalUnitOutput{}

        mockClient.On("UpdateOrganizationalUnit", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateOrganizationalUnit(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdatePolicy", func(t *testing.T) {
        input := &organizations.UpdatePolicyInput{}
        output := &organizations.UpdatePolicyOutput{}

        mockClient.On("UpdatePolicy", ctx, input).Return(output, nil)

        result, err := mockClient.UpdatePolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
