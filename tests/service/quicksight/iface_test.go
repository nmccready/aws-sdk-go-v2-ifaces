package quicksight_test

// tests for the quicksight service interface for this ../../../service/quicksight/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/quicksight"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/quicksight/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/quicksight/quicksight_iface"
	"github.com/stretchr/testify/assert"
)

func TestQuicksightServiceCanBeMocked(t *testing.T) {
	var iface quicksight_iface.IClient
	iface = &quicksight.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := quicksight.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCancelIngestion", func(t *testing.T) {
        input := &quicksight.CancelIngestionInput{}
        output := &quicksight.CancelIngestionOutput{}

        mockClient.On("CancelIngestion", ctx, input).Return(output, nil)

        result, err := mockClient.CancelIngestion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateAccountCustomization", func(t *testing.T) {
        input := &quicksight.CreateAccountCustomizationInput{}
        output := &quicksight.CreateAccountCustomizationOutput{}

        mockClient.On("CreateAccountCustomization", ctx, input).Return(output, nil)

        result, err := mockClient.CreateAccountCustomization(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateAccountSubscription", func(t *testing.T) {
        input := &quicksight.CreateAccountSubscriptionInput{}
        output := &quicksight.CreateAccountSubscriptionOutput{}

        mockClient.On("CreateAccountSubscription", ctx, input).Return(output, nil)

        result, err := mockClient.CreateAccountSubscription(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateAnalysis", func(t *testing.T) {
        input := &quicksight.CreateAnalysisInput{}
        output := &quicksight.CreateAnalysisOutput{}

        mockClient.On("CreateAnalysis", ctx, input).Return(output, nil)

        result, err := mockClient.CreateAnalysis(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateDashboard", func(t *testing.T) {
        input := &quicksight.CreateDashboardInput{}
        output := &quicksight.CreateDashboardOutput{}

        mockClient.On("CreateDashboard", ctx, input).Return(output, nil)

        result, err := mockClient.CreateDashboard(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateDataSet", func(t *testing.T) {
        input := &quicksight.CreateDataSetInput{}
        output := &quicksight.CreateDataSetOutput{}

        mockClient.On("CreateDataSet", ctx, input).Return(output, nil)

        result, err := mockClient.CreateDataSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateDataSource", func(t *testing.T) {
        input := &quicksight.CreateDataSourceInput{}
        output := &quicksight.CreateDataSourceOutput{}

        mockClient.On("CreateDataSource", ctx, input).Return(output, nil)

        result, err := mockClient.CreateDataSource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateFolder", func(t *testing.T) {
        input := &quicksight.CreateFolderInput{}
        output := &quicksight.CreateFolderOutput{}

        mockClient.On("CreateFolder", ctx, input).Return(output, nil)

        result, err := mockClient.CreateFolder(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateFolderMembership", func(t *testing.T) {
        input := &quicksight.CreateFolderMembershipInput{}
        output := &quicksight.CreateFolderMembershipOutput{}

        mockClient.On("CreateFolderMembership", ctx, input).Return(output, nil)

        result, err := mockClient.CreateFolderMembership(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateGroup", func(t *testing.T) {
        input := &quicksight.CreateGroupInput{}
        output := &quicksight.CreateGroupOutput{}

        mockClient.On("CreateGroup", ctx, input).Return(output, nil)

        result, err := mockClient.CreateGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateGroupMembership", func(t *testing.T) {
        input := &quicksight.CreateGroupMembershipInput{}
        output := &quicksight.CreateGroupMembershipOutput{}

        mockClient.On("CreateGroupMembership", ctx, input).Return(output, nil)

        result, err := mockClient.CreateGroupMembership(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateIAMPolicyAssignment", func(t *testing.T) {
        input := &quicksight.CreateIAMPolicyAssignmentInput{}
        output := &quicksight.CreateIAMPolicyAssignmentOutput{}

        mockClient.On("CreateIAMPolicyAssignment", ctx, input).Return(output, nil)

        result, err := mockClient.CreateIAMPolicyAssignment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateIngestion", func(t *testing.T) {
        input := &quicksight.CreateIngestionInput{}
        output := &quicksight.CreateIngestionOutput{}

        mockClient.On("CreateIngestion", ctx, input).Return(output, nil)

        result, err := mockClient.CreateIngestion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateNamespace", func(t *testing.T) {
        input := &quicksight.CreateNamespaceInput{}
        output := &quicksight.CreateNamespaceOutput{}

        mockClient.On("CreateNamespace", ctx, input).Return(output, nil)

        result, err := mockClient.CreateNamespace(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateRefreshSchedule", func(t *testing.T) {
        input := &quicksight.CreateRefreshScheduleInput{}
        output := &quicksight.CreateRefreshScheduleOutput{}

        mockClient.On("CreateRefreshSchedule", ctx, input).Return(output, nil)

        result, err := mockClient.CreateRefreshSchedule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateRoleMembership", func(t *testing.T) {
        input := &quicksight.CreateRoleMembershipInput{}
        output := &quicksight.CreateRoleMembershipOutput{}

        mockClient.On("CreateRoleMembership", ctx, input).Return(output, nil)

        result, err := mockClient.CreateRoleMembership(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateTemplate", func(t *testing.T) {
        input := &quicksight.CreateTemplateInput{}
        output := &quicksight.CreateTemplateOutput{}

        mockClient.On("CreateTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.CreateTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateTemplateAlias", func(t *testing.T) {
        input := &quicksight.CreateTemplateAliasInput{}
        output := &quicksight.CreateTemplateAliasOutput{}

        mockClient.On("CreateTemplateAlias", ctx, input).Return(output, nil)

        result, err := mockClient.CreateTemplateAlias(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateTheme", func(t *testing.T) {
        input := &quicksight.CreateThemeInput{}
        output := &quicksight.CreateThemeOutput{}

        mockClient.On("CreateTheme", ctx, input).Return(output, nil)

        result, err := mockClient.CreateTheme(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateThemeAlias", func(t *testing.T) {
        input := &quicksight.CreateThemeAliasInput{}
        output := &quicksight.CreateThemeAliasOutput{}

        mockClient.On("CreateThemeAlias", ctx, input).Return(output, nil)

        result, err := mockClient.CreateThemeAlias(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateTopic", func(t *testing.T) {
        input := &quicksight.CreateTopicInput{}
        output := &quicksight.CreateTopicOutput{}

        mockClient.On("CreateTopic", ctx, input).Return(output, nil)

        result, err := mockClient.CreateTopic(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateTopicRefreshSchedule", func(t *testing.T) {
        input := &quicksight.CreateTopicRefreshScheduleInput{}
        output := &quicksight.CreateTopicRefreshScheduleOutput{}

        mockClient.On("CreateTopicRefreshSchedule", ctx, input).Return(output, nil)

        result, err := mockClient.CreateTopicRefreshSchedule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateVPCConnection", func(t *testing.T) {
        input := &quicksight.CreateVPCConnectionInput{}
        output := &quicksight.CreateVPCConnectionOutput{}

        mockClient.On("CreateVPCConnection", ctx, input).Return(output, nil)

        result, err := mockClient.CreateVPCConnection(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteAccountCustomization", func(t *testing.T) {
        input := &quicksight.DeleteAccountCustomizationInput{}
        output := &quicksight.DeleteAccountCustomizationOutput{}

        mockClient.On("DeleteAccountCustomization", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteAccountCustomization(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteAccountSubscription", func(t *testing.T) {
        input := &quicksight.DeleteAccountSubscriptionInput{}
        output := &quicksight.DeleteAccountSubscriptionOutput{}

        mockClient.On("DeleteAccountSubscription", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteAccountSubscription(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteAnalysis", func(t *testing.T) {
        input := &quicksight.DeleteAnalysisInput{}
        output := &quicksight.DeleteAnalysisOutput{}

        mockClient.On("DeleteAnalysis", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteAnalysis(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteDashboard", func(t *testing.T) {
        input := &quicksight.DeleteDashboardInput{}
        output := &quicksight.DeleteDashboardOutput{}

        mockClient.On("DeleteDashboard", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteDashboard(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteDataSet", func(t *testing.T) {
        input := &quicksight.DeleteDataSetInput{}
        output := &quicksight.DeleteDataSetOutput{}

        mockClient.On("DeleteDataSet", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteDataSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteDataSetRefreshProperties", func(t *testing.T) {
        input := &quicksight.DeleteDataSetRefreshPropertiesInput{}
        output := &quicksight.DeleteDataSetRefreshPropertiesOutput{}

        mockClient.On("DeleteDataSetRefreshProperties", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteDataSetRefreshProperties(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteDataSource", func(t *testing.T) {
        input := &quicksight.DeleteDataSourceInput{}
        output := &quicksight.DeleteDataSourceOutput{}

        mockClient.On("DeleteDataSource", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteDataSource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteFolder", func(t *testing.T) {
        input := &quicksight.DeleteFolderInput{}
        output := &quicksight.DeleteFolderOutput{}

        mockClient.On("DeleteFolder", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteFolder(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteFolderMembership", func(t *testing.T) {
        input := &quicksight.DeleteFolderMembershipInput{}
        output := &quicksight.DeleteFolderMembershipOutput{}

        mockClient.On("DeleteFolderMembership", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteFolderMembership(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteGroup", func(t *testing.T) {
        input := &quicksight.DeleteGroupInput{}
        output := &quicksight.DeleteGroupOutput{}

        mockClient.On("DeleteGroup", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteGroupMembership", func(t *testing.T) {
        input := &quicksight.DeleteGroupMembershipInput{}
        output := &quicksight.DeleteGroupMembershipOutput{}

        mockClient.On("DeleteGroupMembership", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteGroupMembership(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteIAMPolicyAssignment", func(t *testing.T) {
        input := &quicksight.DeleteIAMPolicyAssignmentInput{}
        output := &quicksight.DeleteIAMPolicyAssignmentOutput{}

        mockClient.On("DeleteIAMPolicyAssignment", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteIAMPolicyAssignment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteIdentityPropagationConfig", func(t *testing.T) {
        input := &quicksight.DeleteIdentityPropagationConfigInput{}
        output := &quicksight.DeleteIdentityPropagationConfigOutput{}

        mockClient.On("DeleteIdentityPropagationConfig", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteIdentityPropagationConfig(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteNamespace", func(t *testing.T) {
        input := &quicksight.DeleteNamespaceInput{}
        output := &quicksight.DeleteNamespaceOutput{}

        mockClient.On("DeleteNamespace", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteNamespace(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteRefreshSchedule", func(t *testing.T) {
        input := &quicksight.DeleteRefreshScheduleInput{}
        output := &quicksight.DeleteRefreshScheduleOutput{}

        mockClient.On("DeleteRefreshSchedule", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteRefreshSchedule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteRoleCustomPermission", func(t *testing.T) {
        input := &quicksight.DeleteRoleCustomPermissionInput{}
        output := &quicksight.DeleteRoleCustomPermissionOutput{}

        mockClient.On("DeleteRoleCustomPermission", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteRoleCustomPermission(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteRoleMembership", func(t *testing.T) {
        input := &quicksight.DeleteRoleMembershipInput{}
        output := &quicksight.DeleteRoleMembershipOutput{}

        mockClient.On("DeleteRoleMembership", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteRoleMembership(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteTemplate", func(t *testing.T) {
        input := &quicksight.DeleteTemplateInput{}
        output := &quicksight.DeleteTemplateOutput{}

        mockClient.On("DeleteTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteTemplateAlias", func(t *testing.T) {
        input := &quicksight.DeleteTemplateAliasInput{}
        output := &quicksight.DeleteTemplateAliasOutput{}

        mockClient.On("DeleteTemplateAlias", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteTemplateAlias(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteTheme", func(t *testing.T) {
        input := &quicksight.DeleteThemeInput{}
        output := &quicksight.DeleteThemeOutput{}

        mockClient.On("DeleteTheme", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteTheme(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteThemeAlias", func(t *testing.T) {
        input := &quicksight.DeleteThemeAliasInput{}
        output := &quicksight.DeleteThemeAliasOutput{}

        mockClient.On("DeleteThemeAlias", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteThemeAlias(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteTopic", func(t *testing.T) {
        input := &quicksight.DeleteTopicInput{}
        output := &quicksight.DeleteTopicOutput{}

        mockClient.On("DeleteTopic", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteTopic(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteTopicRefreshSchedule", func(t *testing.T) {
        input := &quicksight.DeleteTopicRefreshScheduleInput{}
        output := &quicksight.DeleteTopicRefreshScheduleOutput{}

        mockClient.On("DeleteTopicRefreshSchedule", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteTopicRefreshSchedule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteUser", func(t *testing.T) {
        input := &quicksight.DeleteUserInput{}
        output := &quicksight.DeleteUserOutput{}

        mockClient.On("DeleteUser", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteUser(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteUserByPrincipalId", func(t *testing.T) {
        input := &quicksight.DeleteUserByPrincipalIdInput{}
        output := &quicksight.DeleteUserByPrincipalIdOutput{}

        mockClient.On("DeleteUserByPrincipalId", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteUserByPrincipalId(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteVPCConnection", func(t *testing.T) {
        input := &quicksight.DeleteVPCConnectionInput{}
        output := &quicksight.DeleteVPCConnectionOutput{}

        mockClient.On("DeleteVPCConnection", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteVPCConnection(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeAccountCustomization", func(t *testing.T) {
        input := &quicksight.DescribeAccountCustomizationInput{}
        output := &quicksight.DescribeAccountCustomizationOutput{}

        mockClient.On("DescribeAccountCustomization", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeAccountCustomization(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeAccountSettings", func(t *testing.T) {
        input := &quicksight.DescribeAccountSettingsInput{}
        output := &quicksight.DescribeAccountSettingsOutput{}

        mockClient.On("DescribeAccountSettings", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeAccountSettings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeAccountSubscription", func(t *testing.T) {
        input := &quicksight.DescribeAccountSubscriptionInput{}
        output := &quicksight.DescribeAccountSubscriptionOutput{}

        mockClient.On("DescribeAccountSubscription", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeAccountSubscription(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeAnalysis", func(t *testing.T) {
        input := &quicksight.DescribeAnalysisInput{}
        output := &quicksight.DescribeAnalysisOutput{}

        mockClient.On("DescribeAnalysis", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeAnalysis(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeAnalysisDefinition", func(t *testing.T) {
        input := &quicksight.DescribeAnalysisDefinitionInput{}
        output := &quicksight.DescribeAnalysisDefinitionOutput{}

        mockClient.On("DescribeAnalysisDefinition", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeAnalysisDefinition(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeAnalysisPermissions", func(t *testing.T) {
        input := &quicksight.DescribeAnalysisPermissionsInput{}
        output := &quicksight.DescribeAnalysisPermissionsOutput{}

        mockClient.On("DescribeAnalysisPermissions", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeAnalysisPermissions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeAssetBundleExportJob", func(t *testing.T) {
        input := &quicksight.DescribeAssetBundleExportJobInput{}
        output := &quicksight.DescribeAssetBundleExportJobOutput{}

        mockClient.On("DescribeAssetBundleExportJob", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeAssetBundleExportJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeAssetBundleImportJob", func(t *testing.T) {
        input := &quicksight.DescribeAssetBundleImportJobInput{}
        output := &quicksight.DescribeAssetBundleImportJobOutput{}

        mockClient.On("DescribeAssetBundleImportJob", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeAssetBundleImportJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeDashboard", func(t *testing.T) {
        input := &quicksight.DescribeDashboardInput{}
        output := &quicksight.DescribeDashboardOutput{}

        mockClient.On("DescribeDashboard", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeDashboard(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeDashboardDefinition", func(t *testing.T) {
        input := &quicksight.DescribeDashboardDefinitionInput{}
        output := &quicksight.DescribeDashboardDefinitionOutput{}

        mockClient.On("DescribeDashboardDefinition", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeDashboardDefinition(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeDashboardPermissions", func(t *testing.T) {
        input := &quicksight.DescribeDashboardPermissionsInput{}
        output := &quicksight.DescribeDashboardPermissionsOutput{}

        mockClient.On("DescribeDashboardPermissions", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeDashboardPermissions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeDashboardSnapshotJob", func(t *testing.T) {
        input := &quicksight.DescribeDashboardSnapshotJobInput{}
        output := &quicksight.DescribeDashboardSnapshotJobOutput{}

        mockClient.On("DescribeDashboardSnapshotJob", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeDashboardSnapshotJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeDashboardSnapshotJobResult", func(t *testing.T) {
        input := &quicksight.DescribeDashboardSnapshotJobResultInput{}
        output := &quicksight.DescribeDashboardSnapshotJobResultOutput{}

        mockClient.On("DescribeDashboardSnapshotJobResult", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeDashboardSnapshotJobResult(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeDataSet", func(t *testing.T) {
        input := &quicksight.DescribeDataSetInput{}
        output := &quicksight.DescribeDataSetOutput{}

        mockClient.On("DescribeDataSet", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeDataSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeDataSetPermissions", func(t *testing.T) {
        input := &quicksight.DescribeDataSetPermissionsInput{}
        output := &quicksight.DescribeDataSetPermissionsOutput{}

        mockClient.On("DescribeDataSetPermissions", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeDataSetPermissions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeDataSetRefreshProperties", func(t *testing.T) {
        input := &quicksight.DescribeDataSetRefreshPropertiesInput{}
        output := &quicksight.DescribeDataSetRefreshPropertiesOutput{}

        mockClient.On("DescribeDataSetRefreshProperties", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeDataSetRefreshProperties(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeDataSource", func(t *testing.T) {
        input := &quicksight.DescribeDataSourceInput{}
        output := &quicksight.DescribeDataSourceOutput{}

        mockClient.On("DescribeDataSource", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeDataSource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeDataSourcePermissions", func(t *testing.T) {
        input := &quicksight.DescribeDataSourcePermissionsInput{}
        output := &quicksight.DescribeDataSourcePermissionsOutput{}

        mockClient.On("DescribeDataSourcePermissions", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeDataSourcePermissions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeFolder", func(t *testing.T) {
        input := &quicksight.DescribeFolderInput{}
        output := &quicksight.DescribeFolderOutput{}

        mockClient.On("DescribeFolder", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeFolder(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeFolderPermissions", func(t *testing.T) {
        input := &quicksight.DescribeFolderPermissionsInput{}
        output := &quicksight.DescribeFolderPermissionsOutput{}

        mockClient.On("DescribeFolderPermissions", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeFolderPermissions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeFolderResolvedPermissions", func(t *testing.T) {
        input := &quicksight.DescribeFolderResolvedPermissionsInput{}
        output := &quicksight.DescribeFolderResolvedPermissionsOutput{}

        mockClient.On("DescribeFolderResolvedPermissions", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeFolderResolvedPermissions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeGroup", func(t *testing.T) {
        input := &quicksight.DescribeGroupInput{}
        output := &quicksight.DescribeGroupOutput{}

        mockClient.On("DescribeGroup", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeGroupMembership", func(t *testing.T) {
        input := &quicksight.DescribeGroupMembershipInput{}
        output := &quicksight.DescribeGroupMembershipOutput{}

        mockClient.On("DescribeGroupMembership", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeGroupMembership(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeIAMPolicyAssignment", func(t *testing.T) {
        input := &quicksight.DescribeIAMPolicyAssignmentInput{}
        output := &quicksight.DescribeIAMPolicyAssignmentOutput{}

        mockClient.On("DescribeIAMPolicyAssignment", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeIAMPolicyAssignment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeIngestion", func(t *testing.T) {
        input := &quicksight.DescribeIngestionInput{}
        output := &quicksight.DescribeIngestionOutput{}

        mockClient.On("DescribeIngestion", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeIngestion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeIpRestriction", func(t *testing.T) {
        input := &quicksight.DescribeIpRestrictionInput{}
        output := &quicksight.DescribeIpRestrictionOutput{}

        mockClient.On("DescribeIpRestriction", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeIpRestriction(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeKeyRegistration", func(t *testing.T) {
        input := &quicksight.DescribeKeyRegistrationInput{}
        output := &quicksight.DescribeKeyRegistrationOutput{}

        mockClient.On("DescribeKeyRegistration", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeKeyRegistration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeNamespace", func(t *testing.T) {
        input := &quicksight.DescribeNamespaceInput{}
        output := &quicksight.DescribeNamespaceOutput{}

        mockClient.On("DescribeNamespace", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeNamespace(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeRefreshSchedule", func(t *testing.T) {
        input := &quicksight.DescribeRefreshScheduleInput{}
        output := &quicksight.DescribeRefreshScheduleOutput{}

        mockClient.On("DescribeRefreshSchedule", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeRefreshSchedule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeRoleCustomPermission", func(t *testing.T) {
        input := &quicksight.DescribeRoleCustomPermissionInput{}
        output := &quicksight.DescribeRoleCustomPermissionOutput{}

        mockClient.On("DescribeRoleCustomPermission", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeRoleCustomPermission(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeTemplate", func(t *testing.T) {
        input := &quicksight.DescribeTemplateInput{}
        output := &quicksight.DescribeTemplateOutput{}

        mockClient.On("DescribeTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeTemplateAlias", func(t *testing.T) {
        input := &quicksight.DescribeTemplateAliasInput{}
        output := &quicksight.DescribeTemplateAliasOutput{}

        mockClient.On("DescribeTemplateAlias", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeTemplateAlias(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeTemplateDefinition", func(t *testing.T) {
        input := &quicksight.DescribeTemplateDefinitionInput{}
        output := &quicksight.DescribeTemplateDefinitionOutput{}

        mockClient.On("DescribeTemplateDefinition", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeTemplateDefinition(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeTemplatePermissions", func(t *testing.T) {
        input := &quicksight.DescribeTemplatePermissionsInput{}
        output := &quicksight.DescribeTemplatePermissionsOutput{}

        mockClient.On("DescribeTemplatePermissions", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeTemplatePermissions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeTheme", func(t *testing.T) {
        input := &quicksight.DescribeThemeInput{}
        output := &quicksight.DescribeThemeOutput{}

        mockClient.On("DescribeTheme", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeTheme(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeThemeAlias", func(t *testing.T) {
        input := &quicksight.DescribeThemeAliasInput{}
        output := &quicksight.DescribeThemeAliasOutput{}

        mockClient.On("DescribeThemeAlias", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeThemeAlias(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeThemePermissions", func(t *testing.T) {
        input := &quicksight.DescribeThemePermissionsInput{}
        output := &quicksight.DescribeThemePermissionsOutput{}

        mockClient.On("DescribeThemePermissions", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeThemePermissions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeTopic", func(t *testing.T) {
        input := &quicksight.DescribeTopicInput{}
        output := &quicksight.DescribeTopicOutput{}

        mockClient.On("DescribeTopic", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeTopic(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeTopicPermissions", func(t *testing.T) {
        input := &quicksight.DescribeTopicPermissionsInput{}
        output := &quicksight.DescribeTopicPermissionsOutput{}

        mockClient.On("DescribeTopicPermissions", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeTopicPermissions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeTopicRefresh", func(t *testing.T) {
        input := &quicksight.DescribeTopicRefreshInput{}
        output := &quicksight.DescribeTopicRefreshOutput{}

        mockClient.On("DescribeTopicRefresh", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeTopicRefresh(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeTopicRefreshSchedule", func(t *testing.T) {
        input := &quicksight.DescribeTopicRefreshScheduleInput{}
        output := &quicksight.DescribeTopicRefreshScheduleOutput{}

        mockClient.On("DescribeTopicRefreshSchedule", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeTopicRefreshSchedule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeUser", func(t *testing.T) {
        input := &quicksight.DescribeUserInput{}
        output := &quicksight.DescribeUserOutput{}

        mockClient.On("DescribeUser", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeUser(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeVPCConnection", func(t *testing.T) {
        input := &quicksight.DescribeVPCConnectionInput{}
        output := &quicksight.DescribeVPCConnectionOutput{}

        mockClient.On("DescribeVPCConnection", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeVPCConnection(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGenerateEmbedUrlForAnonymousUser", func(t *testing.T) {
        input := &quicksight.GenerateEmbedUrlForAnonymousUserInput{}
        output := &quicksight.GenerateEmbedUrlForAnonymousUserOutput{}

        mockClient.On("GenerateEmbedUrlForAnonymousUser", ctx, input).Return(output, nil)

        result, err := mockClient.GenerateEmbedUrlForAnonymousUser(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGenerateEmbedUrlForRegisteredUser", func(t *testing.T) {
        input := &quicksight.GenerateEmbedUrlForRegisteredUserInput{}
        output := &quicksight.GenerateEmbedUrlForRegisteredUserOutput{}

        mockClient.On("GenerateEmbedUrlForRegisteredUser", ctx, input).Return(output, nil)

        result, err := mockClient.GenerateEmbedUrlForRegisteredUser(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDashboardEmbedUrl", func(t *testing.T) {
        input := &quicksight.GetDashboardEmbedUrlInput{}
        output := &quicksight.GetDashboardEmbedUrlOutput{}

        mockClient.On("GetDashboardEmbedUrl", ctx, input).Return(output, nil)

        result, err := mockClient.GetDashboardEmbedUrl(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSessionEmbedUrl", func(t *testing.T) {
        input := &quicksight.GetSessionEmbedUrlInput{}
        output := &quicksight.GetSessionEmbedUrlOutput{}

        mockClient.On("GetSessionEmbedUrl", ctx, input).Return(output, nil)

        result, err := mockClient.GetSessionEmbedUrl(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAnalyses", func(t *testing.T) {
        input := &quicksight.ListAnalysesInput{}
        output := &quicksight.ListAnalysesOutput{}

        mockClient.On("ListAnalyses", ctx, input).Return(output, nil)

        result, err := mockClient.ListAnalyses(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAssetBundleExportJobs", func(t *testing.T) {
        input := &quicksight.ListAssetBundleExportJobsInput{}
        output := &quicksight.ListAssetBundleExportJobsOutput{}

        mockClient.On("ListAssetBundleExportJobs", ctx, input).Return(output, nil)

        result, err := mockClient.ListAssetBundleExportJobs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAssetBundleImportJobs", func(t *testing.T) {
        input := &quicksight.ListAssetBundleImportJobsInput{}
        output := &quicksight.ListAssetBundleImportJobsOutput{}

        mockClient.On("ListAssetBundleImportJobs", ctx, input).Return(output, nil)

        result, err := mockClient.ListAssetBundleImportJobs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDashboardVersions", func(t *testing.T) {
        input := &quicksight.ListDashboardVersionsInput{}
        output := &quicksight.ListDashboardVersionsOutput{}

        mockClient.On("ListDashboardVersions", ctx, input).Return(output, nil)

        result, err := mockClient.ListDashboardVersions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDashboards", func(t *testing.T) {
        input := &quicksight.ListDashboardsInput{}
        output := &quicksight.ListDashboardsOutput{}

        mockClient.On("ListDashboards", ctx, input).Return(output, nil)

        result, err := mockClient.ListDashboards(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDataSets", func(t *testing.T) {
        input := &quicksight.ListDataSetsInput{}
        output := &quicksight.ListDataSetsOutput{}

        mockClient.On("ListDataSets", ctx, input).Return(output, nil)

        result, err := mockClient.ListDataSets(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDataSources", func(t *testing.T) {
        input := &quicksight.ListDataSourcesInput{}
        output := &quicksight.ListDataSourcesOutput{}

        mockClient.On("ListDataSources", ctx, input).Return(output, nil)

        result, err := mockClient.ListDataSources(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListFolderMembers", func(t *testing.T) {
        input := &quicksight.ListFolderMembersInput{}
        output := &quicksight.ListFolderMembersOutput{}

        mockClient.On("ListFolderMembers", ctx, input).Return(output, nil)

        result, err := mockClient.ListFolderMembers(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListFolders", func(t *testing.T) {
        input := &quicksight.ListFoldersInput{}
        output := &quicksight.ListFoldersOutput{}

        mockClient.On("ListFolders", ctx, input).Return(output, nil)

        result, err := mockClient.ListFolders(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListGroupMemberships", func(t *testing.T) {
        input := &quicksight.ListGroupMembershipsInput{}
        output := &quicksight.ListGroupMembershipsOutput{}

        mockClient.On("ListGroupMemberships", ctx, input).Return(output, nil)

        result, err := mockClient.ListGroupMemberships(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListGroups", func(t *testing.T) {
        input := &quicksight.ListGroupsInput{}
        output := &quicksight.ListGroupsOutput{}

        mockClient.On("ListGroups", ctx, input).Return(output, nil)

        result, err := mockClient.ListGroups(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListIAMPolicyAssignments", func(t *testing.T) {
        input := &quicksight.ListIAMPolicyAssignmentsInput{}
        output := &quicksight.ListIAMPolicyAssignmentsOutput{}

        mockClient.On("ListIAMPolicyAssignments", ctx, input).Return(output, nil)

        result, err := mockClient.ListIAMPolicyAssignments(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListIAMPolicyAssignmentsForUser", func(t *testing.T) {
        input := &quicksight.ListIAMPolicyAssignmentsForUserInput{}
        output := &quicksight.ListIAMPolicyAssignmentsForUserOutput{}

        mockClient.On("ListIAMPolicyAssignmentsForUser", ctx, input).Return(output, nil)

        result, err := mockClient.ListIAMPolicyAssignmentsForUser(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListIdentityPropagationConfigs", func(t *testing.T) {
        input := &quicksight.ListIdentityPropagationConfigsInput{}
        output := &quicksight.ListIdentityPropagationConfigsOutput{}

        mockClient.On("ListIdentityPropagationConfigs", ctx, input).Return(output, nil)

        result, err := mockClient.ListIdentityPropagationConfigs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListIngestions", func(t *testing.T) {
        input := &quicksight.ListIngestionsInput{}
        output := &quicksight.ListIngestionsOutput{}

        mockClient.On("ListIngestions", ctx, input).Return(output, nil)

        result, err := mockClient.ListIngestions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListNamespaces", func(t *testing.T) {
        input := &quicksight.ListNamespacesInput{}
        output := &quicksight.ListNamespacesOutput{}

        mockClient.On("ListNamespaces", ctx, input).Return(output, nil)

        result, err := mockClient.ListNamespaces(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListRefreshSchedules", func(t *testing.T) {
        input := &quicksight.ListRefreshSchedulesInput{}
        output := &quicksight.ListRefreshSchedulesOutput{}

        mockClient.On("ListRefreshSchedules", ctx, input).Return(output, nil)

        result, err := mockClient.ListRefreshSchedules(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListRoleMemberships", func(t *testing.T) {
        input := &quicksight.ListRoleMembershipsInput{}
        output := &quicksight.ListRoleMembershipsOutput{}

        mockClient.On("ListRoleMemberships", ctx, input).Return(output, nil)

        result, err := mockClient.ListRoleMemberships(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &quicksight.ListTagsForResourceInput{}
        output := &quicksight.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTemplateAliases", func(t *testing.T) {
        input := &quicksight.ListTemplateAliasesInput{}
        output := &quicksight.ListTemplateAliasesOutput{}

        mockClient.On("ListTemplateAliases", ctx, input).Return(output, nil)

        result, err := mockClient.ListTemplateAliases(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTemplateVersions", func(t *testing.T) {
        input := &quicksight.ListTemplateVersionsInput{}
        output := &quicksight.ListTemplateVersionsOutput{}

        mockClient.On("ListTemplateVersions", ctx, input).Return(output, nil)

        result, err := mockClient.ListTemplateVersions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTemplates", func(t *testing.T) {
        input := &quicksight.ListTemplatesInput{}
        output := &quicksight.ListTemplatesOutput{}

        mockClient.On("ListTemplates", ctx, input).Return(output, nil)

        result, err := mockClient.ListTemplates(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListThemeAliases", func(t *testing.T) {
        input := &quicksight.ListThemeAliasesInput{}
        output := &quicksight.ListThemeAliasesOutput{}

        mockClient.On("ListThemeAliases", ctx, input).Return(output, nil)

        result, err := mockClient.ListThemeAliases(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListThemeVersions", func(t *testing.T) {
        input := &quicksight.ListThemeVersionsInput{}
        output := &quicksight.ListThemeVersionsOutput{}

        mockClient.On("ListThemeVersions", ctx, input).Return(output, nil)

        result, err := mockClient.ListThemeVersions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListThemes", func(t *testing.T) {
        input := &quicksight.ListThemesInput{}
        output := &quicksight.ListThemesOutput{}

        mockClient.On("ListThemes", ctx, input).Return(output, nil)

        result, err := mockClient.ListThemes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTopicRefreshSchedules", func(t *testing.T) {
        input := &quicksight.ListTopicRefreshSchedulesInput{}
        output := &quicksight.ListTopicRefreshSchedulesOutput{}

        mockClient.On("ListTopicRefreshSchedules", ctx, input).Return(output, nil)

        result, err := mockClient.ListTopicRefreshSchedules(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTopics", func(t *testing.T) {
        input := &quicksight.ListTopicsInput{}
        output := &quicksight.ListTopicsOutput{}

        mockClient.On("ListTopics", ctx, input).Return(output, nil)

        result, err := mockClient.ListTopics(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListUserGroups", func(t *testing.T) {
        input := &quicksight.ListUserGroupsInput{}
        output := &quicksight.ListUserGroupsOutput{}

        mockClient.On("ListUserGroups", ctx, input).Return(output, nil)

        result, err := mockClient.ListUserGroups(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListUsers", func(t *testing.T) {
        input := &quicksight.ListUsersInput{}
        output := &quicksight.ListUsersOutput{}

        mockClient.On("ListUsers", ctx, input).Return(output, nil)

        result, err := mockClient.ListUsers(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListVPCConnections", func(t *testing.T) {
        input := &quicksight.ListVPCConnectionsInput{}
        output := &quicksight.ListVPCConnectionsOutput{}

        mockClient.On("ListVPCConnections", ctx, input).Return(output, nil)

        result, err := mockClient.ListVPCConnections(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutDataSetRefreshProperties", func(t *testing.T) {
        input := &quicksight.PutDataSetRefreshPropertiesInput{}
        output := &quicksight.PutDataSetRefreshPropertiesOutput{}

        mockClient.On("PutDataSetRefreshProperties", ctx, input).Return(output, nil)

        result, err := mockClient.PutDataSetRefreshProperties(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRegisterUser", func(t *testing.T) {
        input := &quicksight.RegisterUserInput{}
        output := &quicksight.RegisterUserOutput{}

        mockClient.On("RegisterUser", ctx, input).Return(output, nil)

        result, err := mockClient.RegisterUser(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRestoreAnalysis", func(t *testing.T) {
        input := &quicksight.RestoreAnalysisInput{}
        output := &quicksight.RestoreAnalysisOutput{}

        mockClient.On("RestoreAnalysis", ctx, input).Return(output, nil)

        result, err := mockClient.RestoreAnalysis(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSearchAnalyses", func(t *testing.T) {
        input := &quicksight.SearchAnalysesInput{}
        output := &quicksight.SearchAnalysesOutput{}

        mockClient.On("SearchAnalyses", ctx, input).Return(output, nil)

        result, err := mockClient.SearchAnalyses(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSearchDashboards", func(t *testing.T) {
        input := &quicksight.SearchDashboardsInput{}
        output := &quicksight.SearchDashboardsOutput{}

        mockClient.On("SearchDashboards", ctx, input).Return(output, nil)

        result, err := mockClient.SearchDashboards(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSearchDataSets", func(t *testing.T) {
        input := &quicksight.SearchDataSetsInput{}
        output := &quicksight.SearchDataSetsOutput{}

        mockClient.On("SearchDataSets", ctx, input).Return(output, nil)

        result, err := mockClient.SearchDataSets(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSearchDataSources", func(t *testing.T) {
        input := &quicksight.SearchDataSourcesInput{}
        output := &quicksight.SearchDataSourcesOutput{}

        mockClient.On("SearchDataSources", ctx, input).Return(output, nil)

        result, err := mockClient.SearchDataSources(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSearchFolders", func(t *testing.T) {
        input := &quicksight.SearchFoldersInput{}
        output := &quicksight.SearchFoldersOutput{}

        mockClient.On("SearchFolders", ctx, input).Return(output, nil)

        result, err := mockClient.SearchFolders(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSearchGroups", func(t *testing.T) {
        input := &quicksight.SearchGroupsInput{}
        output := &quicksight.SearchGroupsOutput{}

        mockClient.On("SearchGroups", ctx, input).Return(output, nil)

        result, err := mockClient.SearchGroups(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartAssetBundleExportJob", func(t *testing.T) {
        input := &quicksight.StartAssetBundleExportJobInput{}
        output := &quicksight.StartAssetBundleExportJobOutput{}

        mockClient.On("StartAssetBundleExportJob", ctx, input).Return(output, nil)

        result, err := mockClient.StartAssetBundleExportJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartAssetBundleImportJob", func(t *testing.T) {
        input := &quicksight.StartAssetBundleImportJobInput{}
        output := &quicksight.StartAssetBundleImportJobOutput{}

        mockClient.On("StartAssetBundleImportJob", ctx, input).Return(output, nil)

        result, err := mockClient.StartAssetBundleImportJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartDashboardSnapshotJob", func(t *testing.T) {
        input := &quicksight.StartDashboardSnapshotJobInput{}
        output := &quicksight.StartDashboardSnapshotJobOutput{}

        mockClient.On("StartDashboardSnapshotJob", ctx, input).Return(output, nil)

        result, err := mockClient.StartDashboardSnapshotJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &quicksight.TagResourceInput{}
        output := &quicksight.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &quicksight.UntagResourceInput{}
        output := &quicksight.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateAccountCustomization", func(t *testing.T) {
        input := &quicksight.UpdateAccountCustomizationInput{}
        output := &quicksight.UpdateAccountCustomizationOutput{}

        mockClient.On("UpdateAccountCustomization", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateAccountCustomization(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateAccountSettings", func(t *testing.T) {
        input := &quicksight.UpdateAccountSettingsInput{}
        output := &quicksight.UpdateAccountSettingsOutput{}

        mockClient.On("UpdateAccountSettings", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateAccountSettings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateAnalysis", func(t *testing.T) {
        input := &quicksight.UpdateAnalysisInput{}
        output := &quicksight.UpdateAnalysisOutput{}

        mockClient.On("UpdateAnalysis", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateAnalysis(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateAnalysisPermissions", func(t *testing.T) {
        input := &quicksight.UpdateAnalysisPermissionsInput{}
        output := &quicksight.UpdateAnalysisPermissionsOutput{}

        mockClient.On("UpdateAnalysisPermissions", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateAnalysisPermissions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateDashboard", func(t *testing.T) {
        input := &quicksight.UpdateDashboardInput{}
        output := &quicksight.UpdateDashboardOutput{}

        mockClient.On("UpdateDashboard", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateDashboard(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateDashboardLinks", func(t *testing.T) {
        input := &quicksight.UpdateDashboardLinksInput{}
        output := &quicksight.UpdateDashboardLinksOutput{}

        mockClient.On("UpdateDashboardLinks", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateDashboardLinks(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateDashboardPermissions", func(t *testing.T) {
        input := &quicksight.UpdateDashboardPermissionsInput{}
        output := &quicksight.UpdateDashboardPermissionsOutput{}

        mockClient.On("UpdateDashboardPermissions", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateDashboardPermissions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateDashboardPublishedVersion", func(t *testing.T) {
        input := &quicksight.UpdateDashboardPublishedVersionInput{}
        output := &quicksight.UpdateDashboardPublishedVersionOutput{}

        mockClient.On("UpdateDashboardPublishedVersion", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateDashboardPublishedVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateDataSet", func(t *testing.T) {
        input := &quicksight.UpdateDataSetInput{}
        output := &quicksight.UpdateDataSetOutput{}

        mockClient.On("UpdateDataSet", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateDataSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateDataSetPermissions", func(t *testing.T) {
        input := &quicksight.UpdateDataSetPermissionsInput{}
        output := &quicksight.UpdateDataSetPermissionsOutput{}

        mockClient.On("UpdateDataSetPermissions", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateDataSetPermissions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateDataSource", func(t *testing.T) {
        input := &quicksight.UpdateDataSourceInput{}
        output := &quicksight.UpdateDataSourceOutput{}

        mockClient.On("UpdateDataSource", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateDataSource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateDataSourcePermissions", func(t *testing.T) {
        input := &quicksight.UpdateDataSourcePermissionsInput{}
        output := &quicksight.UpdateDataSourcePermissionsOutput{}

        mockClient.On("UpdateDataSourcePermissions", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateDataSourcePermissions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateFolder", func(t *testing.T) {
        input := &quicksight.UpdateFolderInput{}
        output := &quicksight.UpdateFolderOutput{}

        mockClient.On("UpdateFolder", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateFolder(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateFolderPermissions", func(t *testing.T) {
        input := &quicksight.UpdateFolderPermissionsInput{}
        output := &quicksight.UpdateFolderPermissionsOutput{}

        mockClient.On("UpdateFolderPermissions", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateFolderPermissions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateGroup", func(t *testing.T) {
        input := &quicksight.UpdateGroupInput{}
        output := &quicksight.UpdateGroupOutput{}

        mockClient.On("UpdateGroup", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateIAMPolicyAssignment", func(t *testing.T) {
        input := &quicksight.UpdateIAMPolicyAssignmentInput{}
        output := &quicksight.UpdateIAMPolicyAssignmentOutput{}

        mockClient.On("UpdateIAMPolicyAssignment", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateIAMPolicyAssignment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateIdentityPropagationConfig", func(t *testing.T) {
        input := &quicksight.UpdateIdentityPropagationConfigInput{}
        output := &quicksight.UpdateIdentityPropagationConfigOutput{}

        mockClient.On("UpdateIdentityPropagationConfig", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateIdentityPropagationConfig(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateIpRestriction", func(t *testing.T) {
        input := &quicksight.UpdateIpRestrictionInput{}
        output := &quicksight.UpdateIpRestrictionOutput{}

        mockClient.On("UpdateIpRestriction", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateIpRestriction(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateKeyRegistration", func(t *testing.T) {
        input := &quicksight.UpdateKeyRegistrationInput{}
        output := &quicksight.UpdateKeyRegistrationOutput{}

        mockClient.On("UpdateKeyRegistration", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateKeyRegistration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdatePublicSharingSettings", func(t *testing.T) {
        input := &quicksight.UpdatePublicSharingSettingsInput{}
        output := &quicksight.UpdatePublicSharingSettingsOutput{}

        mockClient.On("UpdatePublicSharingSettings", ctx, input).Return(output, nil)

        result, err := mockClient.UpdatePublicSharingSettings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateRefreshSchedule", func(t *testing.T) {
        input := &quicksight.UpdateRefreshScheduleInput{}
        output := &quicksight.UpdateRefreshScheduleOutput{}

        mockClient.On("UpdateRefreshSchedule", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateRefreshSchedule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateRoleCustomPermission", func(t *testing.T) {
        input := &quicksight.UpdateRoleCustomPermissionInput{}
        output := &quicksight.UpdateRoleCustomPermissionOutput{}

        mockClient.On("UpdateRoleCustomPermission", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateRoleCustomPermission(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateSPICECapacityConfiguration", func(t *testing.T) {
        input := &quicksight.UpdateSPICECapacityConfigurationInput{}
        output := &quicksight.UpdateSPICECapacityConfigurationOutput{}

        mockClient.On("UpdateSPICECapacityConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateSPICECapacityConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateTemplate", func(t *testing.T) {
        input := &quicksight.UpdateTemplateInput{}
        output := &quicksight.UpdateTemplateOutput{}

        mockClient.On("UpdateTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateTemplateAlias", func(t *testing.T) {
        input := &quicksight.UpdateTemplateAliasInput{}
        output := &quicksight.UpdateTemplateAliasOutput{}

        mockClient.On("UpdateTemplateAlias", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateTemplateAlias(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateTemplatePermissions", func(t *testing.T) {
        input := &quicksight.UpdateTemplatePermissionsInput{}
        output := &quicksight.UpdateTemplatePermissionsOutput{}

        mockClient.On("UpdateTemplatePermissions", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateTemplatePermissions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateTheme", func(t *testing.T) {
        input := &quicksight.UpdateThemeInput{}
        output := &quicksight.UpdateThemeOutput{}

        mockClient.On("UpdateTheme", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateTheme(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateThemeAlias", func(t *testing.T) {
        input := &quicksight.UpdateThemeAliasInput{}
        output := &quicksight.UpdateThemeAliasOutput{}

        mockClient.On("UpdateThemeAlias", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateThemeAlias(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateThemePermissions", func(t *testing.T) {
        input := &quicksight.UpdateThemePermissionsInput{}
        output := &quicksight.UpdateThemePermissionsOutput{}

        mockClient.On("UpdateThemePermissions", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateThemePermissions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateTopic", func(t *testing.T) {
        input := &quicksight.UpdateTopicInput{}
        output := &quicksight.UpdateTopicOutput{}

        mockClient.On("UpdateTopic", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateTopic(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateTopicPermissions", func(t *testing.T) {
        input := &quicksight.UpdateTopicPermissionsInput{}
        output := &quicksight.UpdateTopicPermissionsOutput{}

        mockClient.On("UpdateTopicPermissions", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateTopicPermissions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateTopicRefreshSchedule", func(t *testing.T) {
        input := &quicksight.UpdateTopicRefreshScheduleInput{}
        output := &quicksight.UpdateTopicRefreshScheduleOutput{}

        mockClient.On("UpdateTopicRefreshSchedule", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateTopicRefreshSchedule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateUser", func(t *testing.T) {
        input := &quicksight.UpdateUserInput{}
        output := &quicksight.UpdateUserOutput{}

        mockClient.On("UpdateUser", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateUser(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateVPCConnection", func(t *testing.T) {
        input := &quicksight.UpdateVPCConnectionInput{}
        output := &quicksight.UpdateVPCConnectionOutput{}

        mockClient.On("UpdateVPCConnection", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateVPCConnection(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
