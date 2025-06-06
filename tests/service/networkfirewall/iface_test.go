// Code generated by ./scripts/gen_tests.sh DO NOT EDIT.
package networkfirewall_test

// tests for the networkfirewall service interface for 
// this ../../../service/networkfirewall/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/networkfirewall"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/networkfirewall/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/networkfirewall/networkfirewall_iface"
	"github.com/stretchr/testify/assert"
)

func TestNetworkfirewallServiceCanBeMocked(t *testing.T) {
	var iface networkfirewall_iface.IClient
	iface = &networkfirewall.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := networkfirewall.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateFirewallPolicy", func(t *testing.T) {
        input := &networkfirewall.AssociateFirewallPolicyInput{}
        output := &networkfirewall.AssociateFirewallPolicyOutput{}

        mockClient.On("AssociateFirewallPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateFirewallPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateSubnets", func(t *testing.T) {
        input := &networkfirewall.AssociateSubnetsInput{}
        output := &networkfirewall.AssociateSubnetsOutput{}

        mockClient.On("AssociateSubnets", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateSubnets(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateFirewall", func(t *testing.T) {
        input := &networkfirewall.CreateFirewallInput{}
        output := &networkfirewall.CreateFirewallOutput{}

        mockClient.On("CreateFirewall", ctx, input).Return(output, nil)

        result, err := mockClient.CreateFirewall(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateFirewallPolicy", func(t *testing.T) {
        input := &networkfirewall.CreateFirewallPolicyInput{}
        output := &networkfirewall.CreateFirewallPolicyOutput{}

        mockClient.On("CreateFirewallPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.CreateFirewallPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateRuleGroup", func(t *testing.T) {
        input := &networkfirewall.CreateRuleGroupInput{}
        output := &networkfirewall.CreateRuleGroupOutput{}

        mockClient.On("CreateRuleGroup", ctx, input).Return(output, nil)

        result, err := mockClient.CreateRuleGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateTLSInspectionConfiguration", func(t *testing.T) {
        input := &networkfirewall.CreateTLSInspectionConfigurationInput{}
        output := &networkfirewall.CreateTLSInspectionConfigurationOutput{}

        mockClient.On("CreateTLSInspectionConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.CreateTLSInspectionConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteFirewall", func(t *testing.T) {
        input := &networkfirewall.DeleteFirewallInput{}
        output := &networkfirewall.DeleteFirewallOutput{}

        mockClient.On("DeleteFirewall", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteFirewall(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteFirewallPolicy", func(t *testing.T) {
        input := &networkfirewall.DeleteFirewallPolicyInput{}
        output := &networkfirewall.DeleteFirewallPolicyOutput{}

        mockClient.On("DeleteFirewallPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteFirewallPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteResourcePolicy", func(t *testing.T) {
        input := &networkfirewall.DeleteResourcePolicyInput{}
        output := &networkfirewall.DeleteResourcePolicyOutput{}

        mockClient.On("DeleteResourcePolicy", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteResourcePolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteRuleGroup", func(t *testing.T) {
        input := &networkfirewall.DeleteRuleGroupInput{}
        output := &networkfirewall.DeleteRuleGroupOutput{}

        mockClient.On("DeleteRuleGroup", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteRuleGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteTLSInspectionConfiguration", func(t *testing.T) {
        input := &networkfirewall.DeleteTLSInspectionConfigurationInput{}
        output := &networkfirewall.DeleteTLSInspectionConfigurationOutput{}

        mockClient.On("DeleteTLSInspectionConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteTLSInspectionConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeFirewall", func(t *testing.T) {
        input := &networkfirewall.DescribeFirewallInput{}
        output := &networkfirewall.DescribeFirewallOutput{}

        mockClient.On("DescribeFirewall", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeFirewall(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeFirewallPolicy", func(t *testing.T) {
        input := &networkfirewall.DescribeFirewallPolicyInput{}
        output := &networkfirewall.DescribeFirewallPolicyOutput{}

        mockClient.On("DescribeFirewallPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeFirewallPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeFlowOperation", func(t *testing.T) {
        input := &networkfirewall.DescribeFlowOperationInput{}
        output := &networkfirewall.DescribeFlowOperationOutput{}

        mockClient.On("DescribeFlowOperation", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeFlowOperation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeLoggingConfiguration", func(t *testing.T) {
        input := &networkfirewall.DescribeLoggingConfigurationInput{}
        output := &networkfirewall.DescribeLoggingConfigurationOutput{}

        mockClient.On("DescribeLoggingConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeLoggingConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeResourcePolicy", func(t *testing.T) {
        input := &networkfirewall.DescribeResourcePolicyInput{}
        output := &networkfirewall.DescribeResourcePolicyOutput{}

        mockClient.On("DescribeResourcePolicy", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeResourcePolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeRuleGroup", func(t *testing.T) {
        input := &networkfirewall.DescribeRuleGroupInput{}
        output := &networkfirewall.DescribeRuleGroupOutput{}

        mockClient.On("DescribeRuleGroup", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeRuleGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeRuleGroupMetadata", func(t *testing.T) {
        input := &networkfirewall.DescribeRuleGroupMetadataInput{}
        output := &networkfirewall.DescribeRuleGroupMetadataOutput{}

        mockClient.On("DescribeRuleGroupMetadata", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeRuleGroupMetadata(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeTLSInspectionConfiguration", func(t *testing.T) {
        input := &networkfirewall.DescribeTLSInspectionConfigurationInput{}
        output := &networkfirewall.DescribeTLSInspectionConfigurationOutput{}

        mockClient.On("DescribeTLSInspectionConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeTLSInspectionConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateSubnets", func(t *testing.T) {
        input := &networkfirewall.DisassociateSubnetsInput{}
        output := &networkfirewall.DisassociateSubnetsOutput{}

        mockClient.On("DisassociateSubnets", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateSubnets(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetAnalysisReportResults", func(t *testing.T) {
        input := &networkfirewall.GetAnalysisReportResultsInput{}
        output := &networkfirewall.GetAnalysisReportResultsOutput{}

        mockClient.On("GetAnalysisReportResults", ctx, input).Return(output, nil)

        result, err := mockClient.GetAnalysisReportResults(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAnalysisReports", func(t *testing.T) {
        input := &networkfirewall.ListAnalysisReportsInput{}
        output := &networkfirewall.ListAnalysisReportsOutput{}

        mockClient.On("ListAnalysisReports", ctx, input).Return(output, nil)

        result, err := mockClient.ListAnalysisReports(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListFirewallPolicies", func(t *testing.T) {
        input := &networkfirewall.ListFirewallPoliciesInput{}
        output := &networkfirewall.ListFirewallPoliciesOutput{}

        mockClient.On("ListFirewallPolicies", ctx, input).Return(output, nil)

        result, err := mockClient.ListFirewallPolicies(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListFirewalls", func(t *testing.T) {
        input := &networkfirewall.ListFirewallsInput{}
        output := &networkfirewall.ListFirewallsOutput{}

        mockClient.On("ListFirewalls", ctx, input).Return(output, nil)

        result, err := mockClient.ListFirewalls(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListFlowOperationResults", func(t *testing.T) {
        input := &networkfirewall.ListFlowOperationResultsInput{}
        output := &networkfirewall.ListFlowOperationResultsOutput{}

        mockClient.On("ListFlowOperationResults", ctx, input).Return(output, nil)

        result, err := mockClient.ListFlowOperationResults(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListFlowOperations", func(t *testing.T) {
        input := &networkfirewall.ListFlowOperationsInput{}
        output := &networkfirewall.ListFlowOperationsOutput{}

        mockClient.On("ListFlowOperations", ctx, input).Return(output, nil)

        result, err := mockClient.ListFlowOperations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListRuleGroups", func(t *testing.T) {
        input := &networkfirewall.ListRuleGroupsInput{}
        output := &networkfirewall.ListRuleGroupsOutput{}

        mockClient.On("ListRuleGroups", ctx, input).Return(output, nil)

        result, err := mockClient.ListRuleGroups(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTLSInspectionConfigurations", func(t *testing.T) {
        input := &networkfirewall.ListTLSInspectionConfigurationsInput{}
        output := &networkfirewall.ListTLSInspectionConfigurationsOutput{}

        mockClient.On("ListTLSInspectionConfigurations", ctx, input).Return(output, nil)

        result, err := mockClient.ListTLSInspectionConfigurations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &networkfirewall.ListTagsForResourceInput{}
        output := &networkfirewall.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutResourcePolicy", func(t *testing.T) {
        input := &networkfirewall.PutResourcePolicyInput{}
        output := &networkfirewall.PutResourcePolicyOutput{}

        mockClient.On("PutResourcePolicy", ctx, input).Return(output, nil)

        result, err := mockClient.PutResourcePolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartAnalysisReport", func(t *testing.T) {
        input := &networkfirewall.StartAnalysisReportInput{}
        output := &networkfirewall.StartAnalysisReportOutput{}

        mockClient.On("StartAnalysisReport", ctx, input).Return(output, nil)

        result, err := mockClient.StartAnalysisReport(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartFlowCapture", func(t *testing.T) {
        input := &networkfirewall.StartFlowCaptureInput{}
        output := &networkfirewall.StartFlowCaptureOutput{}

        mockClient.On("StartFlowCapture", ctx, input).Return(output, nil)

        result, err := mockClient.StartFlowCapture(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartFlowFlush", func(t *testing.T) {
        input := &networkfirewall.StartFlowFlushInput{}
        output := &networkfirewall.StartFlowFlushOutput{}

        mockClient.On("StartFlowFlush", ctx, input).Return(output, nil)

        result, err := mockClient.StartFlowFlush(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &networkfirewall.TagResourceInput{}
        output := &networkfirewall.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &networkfirewall.UntagResourceInput{}
        output := &networkfirewall.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateFirewallAnalysisSettings", func(t *testing.T) {
        input := &networkfirewall.UpdateFirewallAnalysisSettingsInput{}
        output := &networkfirewall.UpdateFirewallAnalysisSettingsOutput{}

        mockClient.On("UpdateFirewallAnalysisSettings", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateFirewallAnalysisSettings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateFirewallDeleteProtection", func(t *testing.T) {
        input := &networkfirewall.UpdateFirewallDeleteProtectionInput{}
        output := &networkfirewall.UpdateFirewallDeleteProtectionOutput{}

        mockClient.On("UpdateFirewallDeleteProtection", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateFirewallDeleteProtection(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateFirewallDescription", func(t *testing.T) {
        input := &networkfirewall.UpdateFirewallDescriptionInput{}
        output := &networkfirewall.UpdateFirewallDescriptionOutput{}

        mockClient.On("UpdateFirewallDescription", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateFirewallDescription(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateFirewallEncryptionConfiguration", func(t *testing.T) {
        input := &networkfirewall.UpdateFirewallEncryptionConfigurationInput{}
        output := &networkfirewall.UpdateFirewallEncryptionConfigurationOutput{}

        mockClient.On("UpdateFirewallEncryptionConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateFirewallEncryptionConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateFirewallPolicy", func(t *testing.T) {
        input := &networkfirewall.UpdateFirewallPolicyInput{}
        output := &networkfirewall.UpdateFirewallPolicyOutput{}

        mockClient.On("UpdateFirewallPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateFirewallPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateFirewallPolicyChangeProtection", func(t *testing.T) {
        input := &networkfirewall.UpdateFirewallPolicyChangeProtectionInput{}
        output := &networkfirewall.UpdateFirewallPolicyChangeProtectionOutput{}

        mockClient.On("UpdateFirewallPolicyChangeProtection", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateFirewallPolicyChangeProtection(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateLoggingConfiguration", func(t *testing.T) {
        input := &networkfirewall.UpdateLoggingConfigurationInput{}
        output := &networkfirewall.UpdateLoggingConfigurationOutput{}

        mockClient.On("UpdateLoggingConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateLoggingConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateRuleGroup", func(t *testing.T) {
        input := &networkfirewall.UpdateRuleGroupInput{}
        output := &networkfirewall.UpdateRuleGroupOutput{}

        mockClient.On("UpdateRuleGroup", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateRuleGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateSubnetChangeProtection", func(t *testing.T) {
        input := &networkfirewall.UpdateSubnetChangeProtectionInput{}
        output := &networkfirewall.UpdateSubnetChangeProtectionOutput{}

        mockClient.On("UpdateSubnetChangeProtection", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateSubnetChangeProtection(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateTLSInspectionConfiguration", func(t *testing.T) {
        input := &networkfirewall.UpdateTLSInspectionConfigurationInput{}
        output := &networkfirewall.UpdateTLSInspectionConfigurationOutput{}

        mockClient.On("UpdateTLSInspectionConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateTLSInspectionConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
