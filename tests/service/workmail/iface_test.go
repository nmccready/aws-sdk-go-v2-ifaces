package workmail_test

// tests for the workmail service interface for this ../../../service/workmail/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/workmail"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/workmail/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/workmail/workmail_iface"
	"github.com/stretchr/testify/assert"
)

func TestWorkmailServiceCanBeMocked(t *testing.T) {
	var iface workmail_iface.IClient
	iface = &workmail.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := workmail.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateDelegateToResource", func(t *testing.T) {
        input := &workmail.AssociateDelegateToResourceInput{}
        output := &workmail.AssociateDelegateToResourceOutput{}

        mockClient.On("AssociateDelegateToResource", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateDelegateToResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateMemberToGroup", func(t *testing.T) {
        input := &workmail.AssociateMemberToGroupInput{}
        output := &workmail.AssociateMemberToGroupOutput{}

        mockClient.On("AssociateMemberToGroup", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateMemberToGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssumeImpersonationRole", func(t *testing.T) {
        input := &workmail.AssumeImpersonationRoleInput{}
        output := &workmail.AssumeImpersonationRoleOutput{}

        mockClient.On("AssumeImpersonationRole", ctx, input).Return(output, nil)

        result, err := mockClient.AssumeImpersonationRole(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCancelMailboxExportJob", func(t *testing.T) {
        input := &workmail.CancelMailboxExportJobInput{}
        output := &workmail.CancelMailboxExportJobOutput{}

        mockClient.On("CancelMailboxExportJob", ctx, input).Return(output, nil)

        result, err := mockClient.CancelMailboxExportJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateAlias", func(t *testing.T) {
        input := &workmail.CreateAliasInput{}
        output := &workmail.CreateAliasOutput{}

        mockClient.On("CreateAlias", ctx, input).Return(output, nil)

        result, err := mockClient.CreateAlias(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateAvailabilityConfiguration", func(t *testing.T) {
        input := &workmail.CreateAvailabilityConfigurationInput{}
        output := &workmail.CreateAvailabilityConfigurationOutput{}

        mockClient.On("CreateAvailabilityConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.CreateAvailabilityConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateGroup", func(t *testing.T) {
        input := &workmail.CreateGroupInput{}
        output := &workmail.CreateGroupOutput{}

        mockClient.On("CreateGroup", ctx, input).Return(output, nil)

        result, err := mockClient.CreateGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateImpersonationRole", func(t *testing.T) {
        input := &workmail.CreateImpersonationRoleInput{}
        output := &workmail.CreateImpersonationRoleOutput{}

        mockClient.On("CreateImpersonationRole", ctx, input).Return(output, nil)

        result, err := mockClient.CreateImpersonationRole(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateMobileDeviceAccessRule", func(t *testing.T) {
        input := &workmail.CreateMobileDeviceAccessRuleInput{}
        output := &workmail.CreateMobileDeviceAccessRuleOutput{}

        mockClient.On("CreateMobileDeviceAccessRule", ctx, input).Return(output, nil)

        result, err := mockClient.CreateMobileDeviceAccessRule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateOrganization", func(t *testing.T) {
        input := &workmail.CreateOrganizationInput{}
        output := &workmail.CreateOrganizationOutput{}

        mockClient.On("CreateOrganization", ctx, input).Return(output, nil)

        result, err := mockClient.CreateOrganization(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateResource", func(t *testing.T) {
        input := &workmail.CreateResourceInput{}
        output := &workmail.CreateResourceOutput{}

        mockClient.On("CreateResource", ctx, input).Return(output, nil)

        result, err := mockClient.CreateResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateUser", func(t *testing.T) {
        input := &workmail.CreateUserInput{}
        output := &workmail.CreateUserOutput{}

        mockClient.On("CreateUser", ctx, input).Return(output, nil)

        result, err := mockClient.CreateUser(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteAccessControlRule", func(t *testing.T) {
        input := &workmail.DeleteAccessControlRuleInput{}
        output := &workmail.DeleteAccessControlRuleOutput{}

        mockClient.On("DeleteAccessControlRule", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteAccessControlRule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteAlias", func(t *testing.T) {
        input := &workmail.DeleteAliasInput{}
        output := &workmail.DeleteAliasOutput{}

        mockClient.On("DeleteAlias", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteAlias(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteAvailabilityConfiguration", func(t *testing.T) {
        input := &workmail.DeleteAvailabilityConfigurationInput{}
        output := &workmail.DeleteAvailabilityConfigurationOutput{}

        mockClient.On("DeleteAvailabilityConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteAvailabilityConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteEmailMonitoringConfiguration", func(t *testing.T) {
        input := &workmail.DeleteEmailMonitoringConfigurationInput{}
        output := &workmail.DeleteEmailMonitoringConfigurationOutput{}

        mockClient.On("DeleteEmailMonitoringConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteEmailMonitoringConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteGroup", func(t *testing.T) {
        input := &workmail.DeleteGroupInput{}
        output := &workmail.DeleteGroupOutput{}

        mockClient.On("DeleteGroup", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteImpersonationRole", func(t *testing.T) {
        input := &workmail.DeleteImpersonationRoleInput{}
        output := &workmail.DeleteImpersonationRoleOutput{}

        mockClient.On("DeleteImpersonationRole", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteImpersonationRole(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteMailboxPermissions", func(t *testing.T) {
        input := &workmail.DeleteMailboxPermissionsInput{}
        output := &workmail.DeleteMailboxPermissionsOutput{}

        mockClient.On("DeleteMailboxPermissions", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteMailboxPermissions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteMobileDeviceAccessOverride", func(t *testing.T) {
        input := &workmail.DeleteMobileDeviceAccessOverrideInput{}
        output := &workmail.DeleteMobileDeviceAccessOverrideOutput{}

        mockClient.On("DeleteMobileDeviceAccessOverride", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteMobileDeviceAccessOverride(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteMobileDeviceAccessRule", func(t *testing.T) {
        input := &workmail.DeleteMobileDeviceAccessRuleInput{}
        output := &workmail.DeleteMobileDeviceAccessRuleOutput{}

        mockClient.On("DeleteMobileDeviceAccessRule", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteMobileDeviceAccessRule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteOrganization", func(t *testing.T) {
        input := &workmail.DeleteOrganizationInput{}
        output := &workmail.DeleteOrganizationOutput{}

        mockClient.On("DeleteOrganization", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteOrganization(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteResource", func(t *testing.T) {
        input := &workmail.DeleteResourceInput{}
        output := &workmail.DeleteResourceOutput{}

        mockClient.On("DeleteResource", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteRetentionPolicy", func(t *testing.T) {
        input := &workmail.DeleteRetentionPolicyInput{}
        output := &workmail.DeleteRetentionPolicyOutput{}

        mockClient.On("DeleteRetentionPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteRetentionPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteUser", func(t *testing.T) {
        input := &workmail.DeleteUserInput{}
        output := &workmail.DeleteUserOutput{}

        mockClient.On("DeleteUser", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteUser(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeregisterFromWorkMail", func(t *testing.T) {
        input := &workmail.DeregisterFromWorkMailInput{}
        output := &workmail.DeregisterFromWorkMailOutput{}

        mockClient.On("DeregisterFromWorkMail", ctx, input).Return(output, nil)

        result, err := mockClient.DeregisterFromWorkMail(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeregisterMailDomain", func(t *testing.T) {
        input := &workmail.DeregisterMailDomainInput{}
        output := &workmail.DeregisterMailDomainOutput{}

        mockClient.On("DeregisterMailDomain", ctx, input).Return(output, nil)

        result, err := mockClient.DeregisterMailDomain(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeEmailMonitoringConfiguration", func(t *testing.T) {
        input := &workmail.DescribeEmailMonitoringConfigurationInput{}
        output := &workmail.DescribeEmailMonitoringConfigurationOutput{}

        mockClient.On("DescribeEmailMonitoringConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeEmailMonitoringConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeEntity", func(t *testing.T) {
        input := &workmail.DescribeEntityInput{}
        output := &workmail.DescribeEntityOutput{}

        mockClient.On("DescribeEntity", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeEntity(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeGroup", func(t *testing.T) {
        input := &workmail.DescribeGroupInput{}
        output := &workmail.DescribeGroupOutput{}

        mockClient.On("DescribeGroup", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeInboundDmarcSettings", func(t *testing.T) {
        input := &workmail.DescribeInboundDmarcSettingsInput{}
        output := &workmail.DescribeInboundDmarcSettingsOutput{}

        mockClient.On("DescribeInboundDmarcSettings", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeInboundDmarcSettings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeMailboxExportJob", func(t *testing.T) {
        input := &workmail.DescribeMailboxExportJobInput{}
        output := &workmail.DescribeMailboxExportJobOutput{}

        mockClient.On("DescribeMailboxExportJob", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeMailboxExportJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeOrganization", func(t *testing.T) {
        input := &workmail.DescribeOrganizationInput{}
        output := &workmail.DescribeOrganizationOutput{}

        mockClient.On("DescribeOrganization", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeOrganization(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeResource", func(t *testing.T) {
        input := &workmail.DescribeResourceInput{}
        output := &workmail.DescribeResourceOutput{}

        mockClient.On("DescribeResource", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeUser", func(t *testing.T) {
        input := &workmail.DescribeUserInput{}
        output := &workmail.DescribeUserOutput{}

        mockClient.On("DescribeUser", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeUser(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateDelegateFromResource", func(t *testing.T) {
        input := &workmail.DisassociateDelegateFromResourceInput{}
        output := &workmail.DisassociateDelegateFromResourceOutput{}

        mockClient.On("DisassociateDelegateFromResource", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateDelegateFromResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateMemberFromGroup", func(t *testing.T) {
        input := &workmail.DisassociateMemberFromGroupInput{}
        output := &workmail.DisassociateMemberFromGroupOutput{}

        mockClient.On("DisassociateMemberFromGroup", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateMemberFromGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetAccessControlEffect", func(t *testing.T) {
        input := &workmail.GetAccessControlEffectInput{}
        output := &workmail.GetAccessControlEffectOutput{}

        mockClient.On("GetAccessControlEffect", ctx, input).Return(output, nil)

        result, err := mockClient.GetAccessControlEffect(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDefaultRetentionPolicy", func(t *testing.T) {
        input := &workmail.GetDefaultRetentionPolicyInput{}
        output := &workmail.GetDefaultRetentionPolicyOutput{}

        mockClient.On("GetDefaultRetentionPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.GetDefaultRetentionPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetImpersonationRole", func(t *testing.T) {
        input := &workmail.GetImpersonationRoleInput{}
        output := &workmail.GetImpersonationRoleOutput{}

        mockClient.On("GetImpersonationRole", ctx, input).Return(output, nil)

        result, err := mockClient.GetImpersonationRole(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetImpersonationRoleEffect", func(t *testing.T) {
        input := &workmail.GetImpersonationRoleEffectInput{}
        output := &workmail.GetImpersonationRoleEffectOutput{}

        mockClient.On("GetImpersonationRoleEffect", ctx, input).Return(output, nil)

        result, err := mockClient.GetImpersonationRoleEffect(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetMailDomain", func(t *testing.T) {
        input := &workmail.GetMailDomainInput{}
        output := &workmail.GetMailDomainOutput{}

        mockClient.On("GetMailDomain", ctx, input).Return(output, nil)

        result, err := mockClient.GetMailDomain(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetMailboxDetails", func(t *testing.T) {
        input := &workmail.GetMailboxDetailsInput{}
        output := &workmail.GetMailboxDetailsOutput{}

        mockClient.On("GetMailboxDetails", ctx, input).Return(output, nil)

        result, err := mockClient.GetMailboxDetails(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetMobileDeviceAccessEffect", func(t *testing.T) {
        input := &workmail.GetMobileDeviceAccessEffectInput{}
        output := &workmail.GetMobileDeviceAccessEffectOutput{}

        mockClient.On("GetMobileDeviceAccessEffect", ctx, input).Return(output, nil)

        result, err := mockClient.GetMobileDeviceAccessEffect(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetMobileDeviceAccessOverride", func(t *testing.T) {
        input := &workmail.GetMobileDeviceAccessOverrideInput{}
        output := &workmail.GetMobileDeviceAccessOverrideOutput{}

        mockClient.On("GetMobileDeviceAccessOverride", ctx, input).Return(output, nil)

        result, err := mockClient.GetMobileDeviceAccessOverride(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAccessControlRules", func(t *testing.T) {
        input := &workmail.ListAccessControlRulesInput{}
        output := &workmail.ListAccessControlRulesOutput{}

        mockClient.On("ListAccessControlRules", ctx, input).Return(output, nil)

        result, err := mockClient.ListAccessControlRules(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAliases", func(t *testing.T) {
        input := &workmail.ListAliasesInput{}
        output := &workmail.ListAliasesOutput{}

        mockClient.On("ListAliases", ctx, input).Return(output, nil)

        result, err := mockClient.ListAliases(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAvailabilityConfigurations", func(t *testing.T) {
        input := &workmail.ListAvailabilityConfigurationsInput{}
        output := &workmail.ListAvailabilityConfigurationsOutput{}

        mockClient.On("ListAvailabilityConfigurations", ctx, input).Return(output, nil)

        result, err := mockClient.ListAvailabilityConfigurations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListGroupMembers", func(t *testing.T) {
        input := &workmail.ListGroupMembersInput{}
        output := &workmail.ListGroupMembersOutput{}

        mockClient.On("ListGroupMembers", ctx, input).Return(output, nil)

        result, err := mockClient.ListGroupMembers(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListGroups", func(t *testing.T) {
        input := &workmail.ListGroupsInput{}
        output := &workmail.ListGroupsOutput{}

        mockClient.On("ListGroups", ctx, input).Return(output, nil)

        result, err := mockClient.ListGroups(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListGroupsForEntity", func(t *testing.T) {
        input := &workmail.ListGroupsForEntityInput{}
        output := &workmail.ListGroupsForEntityOutput{}

        mockClient.On("ListGroupsForEntity", ctx, input).Return(output, nil)

        result, err := mockClient.ListGroupsForEntity(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListImpersonationRoles", func(t *testing.T) {
        input := &workmail.ListImpersonationRolesInput{}
        output := &workmail.ListImpersonationRolesOutput{}

        mockClient.On("ListImpersonationRoles", ctx, input).Return(output, nil)

        result, err := mockClient.ListImpersonationRoles(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListMailDomains", func(t *testing.T) {
        input := &workmail.ListMailDomainsInput{}
        output := &workmail.ListMailDomainsOutput{}

        mockClient.On("ListMailDomains", ctx, input).Return(output, nil)

        result, err := mockClient.ListMailDomains(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListMailboxExportJobs", func(t *testing.T) {
        input := &workmail.ListMailboxExportJobsInput{}
        output := &workmail.ListMailboxExportJobsOutput{}

        mockClient.On("ListMailboxExportJobs", ctx, input).Return(output, nil)

        result, err := mockClient.ListMailboxExportJobs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListMailboxPermissions", func(t *testing.T) {
        input := &workmail.ListMailboxPermissionsInput{}
        output := &workmail.ListMailboxPermissionsOutput{}

        mockClient.On("ListMailboxPermissions", ctx, input).Return(output, nil)

        result, err := mockClient.ListMailboxPermissions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListMobileDeviceAccessOverrides", func(t *testing.T) {
        input := &workmail.ListMobileDeviceAccessOverridesInput{}
        output := &workmail.ListMobileDeviceAccessOverridesOutput{}

        mockClient.On("ListMobileDeviceAccessOverrides", ctx, input).Return(output, nil)

        result, err := mockClient.ListMobileDeviceAccessOverrides(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListMobileDeviceAccessRules", func(t *testing.T) {
        input := &workmail.ListMobileDeviceAccessRulesInput{}
        output := &workmail.ListMobileDeviceAccessRulesOutput{}

        mockClient.On("ListMobileDeviceAccessRules", ctx, input).Return(output, nil)

        result, err := mockClient.ListMobileDeviceAccessRules(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListOrganizations", func(t *testing.T) {
        input := &workmail.ListOrganizationsInput{}
        output := &workmail.ListOrganizationsOutput{}

        mockClient.On("ListOrganizations", ctx, input).Return(output, nil)

        result, err := mockClient.ListOrganizations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListResourceDelegates", func(t *testing.T) {
        input := &workmail.ListResourceDelegatesInput{}
        output := &workmail.ListResourceDelegatesOutput{}

        mockClient.On("ListResourceDelegates", ctx, input).Return(output, nil)

        result, err := mockClient.ListResourceDelegates(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListResources", func(t *testing.T) {
        input := &workmail.ListResourcesInput{}
        output := &workmail.ListResourcesOutput{}

        mockClient.On("ListResources", ctx, input).Return(output, nil)

        result, err := mockClient.ListResources(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &workmail.ListTagsForResourceInput{}
        output := &workmail.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListUsers", func(t *testing.T) {
        input := &workmail.ListUsersInput{}
        output := &workmail.ListUsersOutput{}

        mockClient.On("ListUsers", ctx, input).Return(output, nil)

        result, err := mockClient.ListUsers(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutAccessControlRule", func(t *testing.T) {
        input := &workmail.PutAccessControlRuleInput{}
        output := &workmail.PutAccessControlRuleOutput{}

        mockClient.On("PutAccessControlRule", ctx, input).Return(output, nil)

        result, err := mockClient.PutAccessControlRule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutEmailMonitoringConfiguration", func(t *testing.T) {
        input := &workmail.PutEmailMonitoringConfigurationInput{}
        output := &workmail.PutEmailMonitoringConfigurationOutput{}

        mockClient.On("PutEmailMonitoringConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.PutEmailMonitoringConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutInboundDmarcSettings", func(t *testing.T) {
        input := &workmail.PutInboundDmarcSettingsInput{}
        output := &workmail.PutInboundDmarcSettingsOutput{}

        mockClient.On("PutInboundDmarcSettings", ctx, input).Return(output, nil)

        result, err := mockClient.PutInboundDmarcSettings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutMailboxPermissions", func(t *testing.T) {
        input := &workmail.PutMailboxPermissionsInput{}
        output := &workmail.PutMailboxPermissionsOutput{}

        mockClient.On("PutMailboxPermissions", ctx, input).Return(output, nil)

        result, err := mockClient.PutMailboxPermissions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutMobileDeviceAccessOverride", func(t *testing.T) {
        input := &workmail.PutMobileDeviceAccessOverrideInput{}
        output := &workmail.PutMobileDeviceAccessOverrideOutput{}

        mockClient.On("PutMobileDeviceAccessOverride", ctx, input).Return(output, nil)

        result, err := mockClient.PutMobileDeviceAccessOverride(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutRetentionPolicy", func(t *testing.T) {
        input := &workmail.PutRetentionPolicyInput{}
        output := &workmail.PutRetentionPolicyOutput{}

        mockClient.On("PutRetentionPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.PutRetentionPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRegisterMailDomain", func(t *testing.T) {
        input := &workmail.RegisterMailDomainInput{}
        output := &workmail.RegisterMailDomainOutput{}

        mockClient.On("RegisterMailDomain", ctx, input).Return(output, nil)

        result, err := mockClient.RegisterMailDomain(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRegisterToWorkMail", func(t *testing.T) {
        input := &workmail.RegisterToWorkMailInput{}
        output := &workmail.RegisterToWorkMailOutput{}

        mockClient.On("RegisterToWorkMail", ctx, input).Return(output, nil)

        result, err := mockClient.RegisterToWorkMail(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestResetPassword", func(t *testing.T) {
        input := &workmail.ResetPasswordInput{}
        output := &workmail.ResetPasswordOutput{}

        mockClient.On("ResetPassword", ctx, input).Return(output, nil)

        result, err := mockClient.ResetPassword(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartMailboxExportJob", func(t *testing.T) {
        input := &workmail.StartMailboxExportJobInput{}
        output := &workmail.StartMailboxExportJobOutput{}

        mockClient.On("StartMailboxExportJob", ctx, input).Return(output, nil)

        result, err := mockClient.StartMailboxExportJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &workmail.TagResourceInput{}
        output := &workmail.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTestAvailabilityConfiguration", func(t *testing.T) {
        input := &workmail.TestAvailabilityConfigurationInput{}
        output := &workmail.TestAvailabilityConfigurationOutput{}

        mockClient.On("TestAvailabilityConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.TestAvailabilityConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &workmail.UntagResourceInput{}
        output := &workmail.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateAvailabilityConfiguration", func(t *testing.T) {
        input := &workmail.UpdateAvailabilityConfigurationInput{}
        output := &workmail.UpdateAvailabilityConfigurationOutput{}

        mockClient.On("UpdateAvailabilityConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateAvailabilityConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateDefaultMailDomain", func(t *testing.T) {
        input := &workmail.UpdateDefaultMailDomainInput{}
        output := &workmail.UpdateDefaultMailDomainOutput{}

        mockClient.On("UpdateDefaultMailDomain", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateDefaultMailDomain(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateGroup", func(t *testing.T) {
        input := &workmail.UpdateGroupInput{}
        output := &workmail.UpdateGroupOutput{}

        mockClient.On("UpdateGroup", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateImpersonationRole", func(t *testing.T) {
        input := &workmail.UpdateImpersonationRoleInput{}
        output := &workmail.UpdateImpersonationRoleOutput{}

        mockClient.On("UpdateImpersonationRole", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateImpersonationRole(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateMailboxQuota", func(t *testing.T) {
        input := &workmail.UpdateMailboxQuotaInput{}
        output := &workmail.UpdateMailboxQuotaOutput{}

        mockClient.On("UpdateMailboxQuota", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateMailboxQuota(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateMobileDeviceAccessRule", func(t *testing.T) {
        input := &workmail.UpdateMobileDeviceAccessRuleInput{}
        output := &workmail.UpdateMobileDeviceAccessRuleOutput{}

        mockClient.On("UpdateMobileDeviceAccessRule", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateMobileDeviceAccessRule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdatePrimaryEmailAddress", func(t *testing.T) {
        input := &workmail.UpdatePrimaryEmailAddressInput{}
        output := &workmail.UpdatePrimaryEmailAddressOutput{}

        mockClient.On("UpdatePrimaryEmailAddress", ctx, input).Return(output, nil)

        result, err := mockClient.UpdatePrimaryEmailAddress(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateResource", func(t *testing.T) {
        input := &workmail.UpdateResourceInput{}
        output := &workmail.UpdateResourceOutput{}

        mockClient.On("UpdateResource", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateUser", func(t *testing.T) {
        input := &workmail.UpdateUserInput{}
        output := &workmail.UpdateUserOutput{}

        mockClient.On("UpdateUser", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateUser(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
