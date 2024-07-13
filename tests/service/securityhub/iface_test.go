package securityhub_test

// tests for the securityhub service interface for this ../../../service/securityhub/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/securityhub"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/securityhub/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/securityhub/securityhub_iface"
	"github.com/stretchr/testify/assert"
)

func TestSecurityhubServiceCanBeMocked(t *testing.T) {
	var iface securityhub_iface.IClient
	iface = &securityhub.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := securityhub.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAcceptAdministratorInvitation", func(t *testing.T) {
        input := &securityhub.AcceptAdministratorInvitationInput{}
        output := &securityhub.AcceptAdministratorInvitationOutput{}

        mockClient.On("AcceptAdministratorInvitation", ctx, input).Return(output, nil)

        result, err := mockClient.AcceptAdministratorInvitation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAcceptInvitation", func(t *testing.T) {
        input := &securityhub.AcceptInvitationInput{}
        output := &securityhub.AcceptInvitationOutput{}

        mockClient.On("AcceptInvitation", ctx, input).Return(output, nil)

        result, err := mockClient.AcceptInvitation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchDeleteAutomationRules", func(t *testing.T) {
        input := &securityhub.BatchDeleteAutomationRulesInput{}
        output := &securityhub.BatchDeleteAutomationRulesOutput{}

        mockClient.On("BatchDeleteAutomationRules", ctx, input).Return(output, nil)

        result, err := mockClient.BatchDeleteAutomationRules(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchDisableStandards", func(t *testing.T) {
        input := &securityhub.BatchDisableStandardsInput{}
        output := &securityhub.BatchDisableStandardsOutput{}

        mockClient.On("BatchDisableStandards", ctx, input).Return(output, nil)

        result, err := mockClient.BatchDisableStandards(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchEnableStandards", func(t *testing.T) {
        input := &securityhub.BatchEnableStandardsInput{}
        output := &securityhub.BatchEnableStandardsOutput{}

        mockClient.On("BatchEnableStandards", ctx, input).Return(output, nil)

        result, err := mockClient.BatchEnableStandards(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchGetAutomationRules", func(t *testing.T) {
        input := &securityhub.BatchGetAutomationRulesInput{}
        output := &securityhub.BatchGetAutomationRulesOutput{}

        mockClient.On("BatchGetAutomationRules", ctx, input).Return(output, nil)

        result, err := mockClient.BatchGetAutomationRules(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchGetConfigurationPolicyAssociations", func(t *testing.T) {
        input := &securityhub.BatchGetConfigurationPolicyAssociationsInput{}
        output := &securityhub.BatchGetConfigurationPolicyAssociationsOutput{}

        mockClient.On("BatchGetConfigurationPolicyAssociations", ctx, input).Return(output, nil)

        result, err := mockClient.BatchGetConfigurationPolicyAssociations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchGetSecurityControls", func(t *testing.T) {
        input := &securityhub.BatchGetSecurityControlsInput{}
        output := &securityhub.BatchGetSecurityControlsOutput{}

        mockClient.On("BatchGetSecurityControls", ctx, input).Return(output, nil)

        result, err := mockClient.BatchGetSecurityControls(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchGetStandardsControlAssociations", func(t *testing.T) {
        input := &securityhub.BatchGetStandardsControlAssociationsInput{}
        output := &securityhub.BatchGetStandardsControlAssociationsOutput{}

        mockClient.On("BatchGetStandardsControlAssociations", ctx, input).Return(output, nil)

        result, err := mockClient.BatchGetStandardsControlAssociations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchImportFindings", func(t *testing.T) {
        input := &securityhub.BatchImportFindingsInput{}
        output := &securityhub.BatchImportFindingsOutput{}

        mockClient.On("BatchImportFindings", ctx, input).Return(output, nil)

        result, err := mockClient.BatchImportFindings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchUpdateAutomationRules", func(t *testing.T) {
        input := &securityhub.BatchUpdateAutomationRulesInput{}
        output := &securityhub.BatchUpdateAutomationRulesOutput{}

        mockClient.On("BatchUpdateAutomationRules", ctx, input).Return(output, nil)

        result, err := mockClient.BatchUpdateAutomationRules(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchUpdateFindings", func(t *testing.T) {
        input := &securityhub.BatchUpdateFindingsInput{}
        output := &securityhub.BatchUpdateFindingsOutput{}

        mockClient.On("BatchUpdateFindings", ctx, input).Return(output, nil)

        result, err := mockClient.BatchUpdateFindings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchUpdateStandardsControlAssociations", func(t *testing.T) {
        input := &securityhub.BatchUpdateStandardsControlAssociationsInput{}
        output := &securityhub.BatchUpdateStandardsControlAssociationsOutput{}

        mockClient.On("BatchUpdateStandardsControlAssociations", ctx, input).Return(output, nil)

        result, err := mockClient.BatchUpdateStandardsControlAssociations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateActionTarget", func(t *testing.T) {
        input := &securityhub.CreateActionTargetInput{}
        output := &securityhub.CreateActionTargetOutput{}

        mockClient.On("CreateActionTarget", ctx, input).Return(output, nil)

        result, err := mockClient.CreateActionTarget(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateAutomationRule", func(t *testing.T) {
        input := &securityhub.CreateAutomationRuleInput{}
        output := &securityhub.CreateAutomationRuleOutput{}

        mockClient.On("CreateAutomationRule", ctx, input).Return(output, nil)

        result, err := mockClient.CreateAutomationRule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateConfigurationPolicy", func(t *testing.T) {
        input := &securityhub.CreateConfigurationPolicyInput{}
        output := &securityhub.CreateConfigurationPolicyOutput{}

        mockClient.On("CreateConfigurationPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.CreateConfigurationPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateFindingAggregator", func(t *testing.T) {
        input := &securityhub.CreateFindingAggregatorInput{}
        output := &securityhub.CreateFindingAggregatorOutput{}

        mockClient.On("CreateFindingAggregator", ctx, input).Return(output, nil)

        result, err := mockClient.CreateFindingAggregator(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateInsight", func(t *testing.T) {
        input := &securityhub.CreateInsightInput{}
        output := &securityhub.CreateInsightOutput{}

        mockClient.On("CreateInsight", ctx, input).Return(output, nil)

        result, err := mockClient.CreateInsight(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateMembers", func(t *testing.T) {
        input := &securityhub.CreateMembersInput{}
        output := &securityhub.CreateMembersOutput{}

        mockClient.On("CreateMembers", ctx, input).Return(output, nil)

        result, err := mockClient.CreateMembers(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeclineInvitations", func(t *testing.T) {
        input := &securityhub.DeclineInvitationsInput{}
        output := &securityhub.DeclineInvitationsOutput{}

        mockClient.On("DeclineInvitations", ctx, input).Return(output, nil)

        result, err := mockClient.DeclineInvitations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteActionTarget", func(t *testing.T) {
        input := &securityhub.DeleteActionTargetInput{}
        output := &securityhub.DeleteActionTargetOutput{}

        mockClient.On("DeleteActionTarget", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteActionTarget(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteConfigurationPolicy", func(t *testing.T) {
        input := &securityhub.DeleteConfigurationPolicyInput{}
        output := &securityhub.DeleteConfigurationPolicyOutput{}

        mockClient.On("DeleteConfigurationPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteConfigurationPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteFindingAggregator", func(t *testing.T) {
        input := &securityhub.DeleteFindingAggregatorInput{}
        output := &securityhub.DeleteFindingAggregatorOutput{}

        mockClient.On("DeleteFindingAggregator", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteFindingAggregator(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteInsight", func(t *testing.T) {
        input := &securityhub.DeleteInsightInput{}
        output := &securityhub.DeleteInsightOutput{}

        mockClient.On("DeleteInsight", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteInsight(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteInvitations", func(t *testing.T) {
        input := &securityhub.DeleteInvitationsInput{}
        output := &securityhub.DeleteInvitationsOutput{}

        mockClient.On("DeleteInvitations", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteInvitations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteMembers", func(t *testing.T) {
        input := &securityhub.DeleteMembersInput{}
        output := &securityhub.DeleteMembersOutput{}

        mockClient.On("DeleteMembers", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteMembers(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeActionTargets", func(t *testing.T) {
        input := &securityhub.DescribeActionTargetsInput{}
        output := &securityhub.DescribeActionTargetsOutput{}

        mockClient.On("DescribeActionTargets", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeActionTargets(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeHub", func(t *testing.T) {
        input := &securityhub.DescribeHubInput{}
        output := &securityhub.DescribeHubOutput{}

        mockClient.On("DescribeHub", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeHub(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeOrganizationConfiguration", func(t *testing.T) {
        input := &securityhub.DescribeOrganizationConfigurationInput{}
        output := &securityhub.DescribeOrganizationConfigurationOutput{}

        mockClient.On("DescribeOrganizationConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeOrganizationConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeProducts", func(t *testing.T) {
        input := &securityhub.DescribeProductsInput{}
        output := &securityhub.DescribeProductsOutput{}

        mockClient.On("DescribeProducts", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeProducts(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeStandards", func(t *testing.T) {
        input := &securityhub.DescribeStandardsInput{}
        output := &securityhub.DescribeStandardsOutput{}

        mockClient.On("DescribeStandards", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeStandards(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeStandardsControls", func(t *testing.T) {
        input := &securityhub.DescribeStandardsControlsInput{}
        output := &securityhub.DescribeStandardsControlsOutput{}

        mockClient.On("DescribeStandardsControls", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeStandardsControls(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisableImportFindingsForProduct", func(t *testing.T) {
        input := &securityhub.DisableImportFindingsForProductInput{}
        output := &securityhub.DisableImportFindingsForProductOutput{}

        mockClient.On("DisableImportFindingsForProduct", ctx, input).Return(output, nil)

        result, err := mockClient.DisableImportFindingsForProduct(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisableOrganizationAdminAccount", func(t *testing.T) {
        input := &securityhub.DisableOrganizationAdminAccountInput{}
        output := &securityhub.DisableOrganizationAdminAccountOutput{}

        mockClient.On("DisableOrganizationAdminAccount", ctx, input).Return(output, nil)

        result, err := mockClient.DisableOrganizationAdminAccount(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisableSecurityHub", func(t *testing.T) {
        input := &securityhub.DisableSecurityHubInput{}
        output := &securityhub.DisableSecurityHubOutput{}

        mockClient.On("DisableSecurityHub", ctx, input).Return(output, nil)

        result, err := mockClient.DisableSecurityHub(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateFromAdministratorAccount", func(t *testing.T) {
        input := &securityhub.DisassociateFromAdministratorAccountInput{}
        output := &securityhub.DisassociateFromAdministratorAccountOutput{}

        mockClient.On("DisassociateFromAdministratorAccount", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateFromAdministratorAccount(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateFromMasterAccount", func(t *testing.T) {
        input := &securityhub.DisassociateFromMasterAccountInput{}
        output := &securityhub.DisassociateFromMasterAccountOutput{}

        mockClient.On("DisassociateFromMasterAccount", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateFromMasterAccount(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateMembers", func(t *testing.T) {
        input := &securityhub.DisassociateMembersInput{}
        output := &securityhub.DisassociateMembersOutput{}

        mockClient.On("DisassociateMembers", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateMembers(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestEnableImportFindingsForProduct", func(t *testing.T) {
        input := &securityhub.EnableImportFindingsForProductInput{}
        output := &securityhub.EnableImportFindingsForProductOutput{}

        mockClient.On("EnableImportFindingsForProduct", ctx, input).Return(output, nil)

        result, err := mockClient.EnableImportFindingsForProduct(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestEnableOrganizationAdminAccount", func(t *testing.T) {
        input := &securityhub.EnableOrganizationAdminAccountInput{}
        output := &securityhub.EnableOrganizationAdminAccountOutput{}

        mockClient.On("EnableOrganizationAdminAccount", ctx, input).Return(output, nil)

        result, err := mockClient.EnableOrganizationAdminAccount(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestEnableSecurityHub", func(t *testing.T) {
        input := &securityhub.EnableSecurityHubInput{}
        output := &securityhub.EnableSecurityHubOutput{}

        mockClient.On("EnableSecurityHub", ctx, input).Return(output, nil)

        result, err := mockClient.EnableSecurityHub(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetAdministratorAccount", func(t *testing.T) {
        input := &securityhub.GetAdministratorAccountInput{}
        output := &securityhub.GetAdministratorAccountOutput{}

        mockClient.On("GetAdministratorAccount", ctx, input).Return(output, nil)

        result, err := mockClient.GetAdministratorAccount(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetConfigurationPolicy", func(t *testing.T) {
        input := &securityhub.GetConfigurationPolicyInput{}
        output := &securityhub.GetConfigurationPolicyOutput{}

        mockClient.On("GetConfigurationPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.GetConfigurationPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetConfigurationPolicyAssociation", func(t *testing.T) {
        input := &securityhub.GetConfigurationPolicyAssociationInput{}
        output := &securityhub.GetConfigurationPolicyAssociationOutput{}

        mockClient.On("GetConfigurationPolicyAssociation", ctx, input).Return(output, nil)

        result, err := mockClient.GetConfigurationPolicyAssociation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetEnabledStandards", func(t *testing.T) {
        input := &securityhub.GetEnabledStandardsInput{}
        output := &securityhub.GetEnabledStandardsOutput{}

        mockClient.On("GetEnabledStandards", ctx, input).Return(output, nil)

        result, err := mockClient.GetEnabledStandards(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetFindingAggregator", func(t *testing.T) {
        input := &securityhub.GetFindingAggregatorInput{}
        output := &securityhub.GetFindingAggregatorOutput{}

        mockClient.On("GetFindingAggregator", ctx, input).Return(output, nil)

        result, err := mockClient.GetFindingAggregator(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetFindingHistory", func(t *testing.T) {
        input := &securityhub.GetFindingHistoryInput{}
        output := &securityhub.GetFindingHistoryOutput{}

        mockClient.On("GetFindingHistory", ctx, input).Return(output, nil)

        result, err := mockClient.GetFindingHistory(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetFindings", func(t *testing.T) {
        input := &securityhub.GetFindingsInput{}
        output := &securityhub.GetFindingsOutput{}

        mockClient.On("GetFindings", ctx, input).Return(output, nil)

        result, err := mockClient.GetFindings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetInsightResults", func(t *testing.T) {
        input := &securityhub.GetInsightResultsInput{}
        output := &securityhub.GetInsightResultsOutput{}

        mockClient.On("GetInsightResults", ctx, input).Return(output, nil)

        result, err := mockClient.GetInsightResults(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetInsights", func(t *testing.T) {
        input := &securityhub.GetInsightsInput{}
        output := &securityhub.GetInsightsOutput{}

        mockClient.On("GetInsights", ctx, input).Return(output, nil)

        result, err := mockClient.GetInsights(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetInvitationsCount", func(t *testing.T) {
        input := &securityhub.GetInvitationsCountInput{}
        output := &securityhub.GetInvitationsCountOutput{}

        mockClient.On("GetInvitationsCount", ctx, input).Return(output, nil)

        result, err := mockClient.GetInvitationsCount(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetMasterAccount", func(t *testing.T) {
        input := &securityhub.GetMasterAccountInput{}
        output := &securityhub.GetMasterAccountOutput{}

        mockClient.On("GetMasterAccount", ctx, input).Return(output, nil)

        result, err := mockClient.GetMasterAccount(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetMembers", func(t *testing.T) {
        input := &securityhub.GetMembersInput{}
        output := &securityhub.GetMembersOutput{}

        mockClient.On("GetMembers", ctx, input).Return(output, nil)

        result, err := mockClient.GetMembers(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSecurityControlDefinition", func(t *testing.T) {
        input := &securityhub.GetSecurityControlDefinitionInput{}
        output := &securityhub.GetSecurityControlDefinitionOutput{}

        mockClient.On("GetSecurityControlDefinition", ctx, input).Return(output, nil)

        result, err := mockClient.GetSecurityControlDefinition(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestInviteMembers", func(t *testing.T) {
        input := &securityhub.InviteMembersInput{}
        output := &securityhub.InviteMembersOutput{}

        mockClient.On("InviteMembers", ctx, input).Return(output, nil)

        result, err := mockClient.InviteMembers(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAutomationRules", func(t *testing.T) {
        input := &securityhub.ListAutomationRulesInput{}
        output := &securityhub.ListAutomationRulesOutput{}

        mockClient.On("ListAutomationRules", ctx, input).Return(output, nil)

        result, err := mockClient.ListAutomationRules(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListConfigurationPolicies", func(t *testing.T) {
        input := &securityhub.ListConfigurationPoliciesInput{}
        output := &securityhub.ListConfigurationPoliciesOutput{}

        mockClient.On("ListConfigurationPolicies", ctx, input).Return(output, nil)

        result, err := mockClient.ListConfigurationPolicies(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListConfigurationPolicyAssociations", func(t *testing.T) {
        input := &securityhub.ListConfigurationPolicyAssociationsInput{}
        output := &securityhub.ListConfigurationPolicyAssociationsOutput{}

        mockClient.On("ListConfigurationPolicyAssociations", ctx, input).Return(output, nil)

        result, err := mockClient.ListConfigurationPolicyAssociations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListEnabledProductsForImport", func(t *testing.T) {
        input := &securityhub.ListEnabledProductsForImportInput{}
        output := &securityhub.ListEnabledProductsForImportOutput{}

        mockClient.On("ListEnabledProductsForImport", ctx, input).Return(output, nil)

        result, err := mockClient.ListEnabledProductsForImport(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListFindingAggregators", func(t *testing.T) {
        input := &securityhub.ListFindingAggregatorsInput{}
        output := &securityhub.ListFindingAggregatorsOutput{}

        mockClient.On("ListFindingAggregators", ctx, input).Return(output, nil)

        result, err := mockClient.ListFindingAggregators(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListInvitations", func(t *testing.T) {
        input := &securityhub.ListInvitationsInput{}
        output := &securityhub.ListInvitationsOutput{}

        mockClient.On("ListInvitations", ctx, input).Return(output, nil)

        result, err := mockClient.ListInvitations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListMembers", func(t *testing.T) {
        input := &securityhub.ListMembersInput{}
        output := &securityhub.ListMembersOutput{}

        mockClient.On("ListMembers", ctx, input).Return(output, nil)

        result, err := mockClient.ListMembers(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListOrganizationAdminAccounts", func(t *testing.T) {
        input := &securityhub.ListOrganizationAdminAccountsInput{}
        output := &securityhub.ListOrganizationAdminAccountsOutput{}

        mockClient.On("ListOrganizationAdminAccounts", ctx, input).Return(output, nil)

        result, err := mockClient.ListOrganizationAdminAccounts(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListSecurityControlDefinitions", func(t *testing.T) {
        input := &securityhub.ListSecurityControlDefinitionsInput{}
        output := &securityhub.ListSecurityControlDefinitionsOutput{}

        mockClient.On("ListSecurityControlDefinitions", ctx, input).Return(output, nil)

        result, err := mockClient.ListSecurityControlDefinitions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListStandardsControlAssociations", func(t *testing.T) {
        input := &securityhub.ListStandardsControlAssociationsInput{}
        output := &securityhub.ListStandardsControlAssociationsOutput{}

        mockClient.On("ListStandardsControlAssociations", ctx, input).Return(output, nil)

        result, err := mockClient.ListStandardsControlAssociations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &securityhub.ListTagsForResourceInput{}
        output := &securityhub.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartConfigurationPolicyAssociation", func(t *testing.T) {
        input := &securityhub.StartConfigurationPolicyAssociationInput{}
        output := &securityhub.StartConfigurationPolicyAssociationOutput{}

        mockClient.On("StartConfigurationPolicyAssociation", ctx, input).Return(output, nil)

        result, err := mockClient.StartConfigurationPolicyAssociation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartConfigurationPolicyDisassociation", func(t *testing.T) {
        input := &securityhub.StartConfigurationPolicyDisassociationInput{}
        output := &securityhub.StartConfigurationPolicyDisassociationOutput{}

        mockClient.On("StartConfigurationPolicyDisassociation", ctx, input).Return(output, nil)

        result, err := mockClient.StartConfigurationPolicyDisassociation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &securityhub.TagResourceInput{}
        output := &securityhub.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &securityhub.UntagResourceInput{}
        output := &securityhub.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateActionTarget", func(t *testing.T) {
        input := &securityhub.UpdateActionTargetInput{}
        output := &securityhub.UpdateActionTargetOutput{}

        mockClient.On("UpdateActionTarget", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateActionTarget(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateConfigurationPolicy", func(t *testing.T) {
        input := &securityhub.UpdateConfigurationPolicyInput{}
        output := &securityhub.UpdateConfigurationPolicyOutput{}

        mockClient.On("UpdateConfigurationPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateConfigurationPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateFindingAggregator", func(t *testing.T) {
        input := &securityhub.UpdateFindingAggregatorInput{}
        output := &securityhub.UpdateFindingAggregatorOutput{}

        mockClient.On("UpdateFindingAggregator", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateFindingAggregator(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateFindings", func(t *testing.T) {
        input := &securityhub.UpdateFindingsInput{}
        output := &securityhub.UpdateFindingsOutput{}

        mockClient.On("UpdateFindings", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateFindings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateInsight", func(t *testing.T) {
        input := &securityhub.UpdateInsightInput{}
        output := &securityhub.UpdateInsightOutput{}

        mockClient.On("UpdateInsight", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateInsight(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateOrganizationConfiguration", func(t *testing.T) {
        input := &securityhub.UpdateOrganizationConfigurationInput{}
        output := &securityhub.UpdateOrganizationConfigurationOutput{}

        mockClient.On("UpdateOrganizationConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateOrganizationConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateSecurityControl", func(t *testing.T) {
        input := &securityhub.UpdateSecurityControlInput{}
        output := &securityhub.UpdateSecurityControlOutput{}

        mockClient.On("UpdateSecurityControl", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateSecurityControl(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateSecurityHubConfiguration", func(t *testing.T) {
        input := &securityhub.UpdateSecurityHubConfigurationInput{}
        output := &securityhub.UpdateSecurityHubConfigurationOutput{}

        mockClient.On("UpdateSecurityHubConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateSecurityHubConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateStandardsControl", func(t *testing.T) {
        input := &securityhub.UpdateStandardsControlInput{}
        output := &securityhub.UpdateStandardsControlOutput{}

        mockClient.On("UpdateStandardsControl", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateStandardsControl(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
