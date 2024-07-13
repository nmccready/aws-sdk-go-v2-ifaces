package amp_test

// tests for the amp service interface for this ../../../service/amp/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/amp"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/amp/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/amp/amp_iface"
	"github.com/stretchr/testify/assert"
)

func TestAmpServiceCanBeMocked(t *testing.T) {
	var iface amp_iface.IClient
	iface = &amp.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := amp.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateAlertManagerDefinition", func(t *testing.T) {
        input := &amp.CreateAlertManagerDefinitionInput{}
        output := &amp.CreateAlertManagerDefinitionOutput{}

        mockClient.On("CreateAlertManagerDefinition", ctx, input).Return(output, nil)

        result, err := mockClient.CreateAlertManagerDefinition(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateLoggingConfiguration", func(t *testing.T) {
        input := &amp.CreateLoggingConfigurationInput{}
        output := &amp.CreateLoggingConfigurationOutput{}

        mockClient.On("CreateLoggingConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.CreateLoggingConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateRuleGroupsNamespace", func(t *testing.T) {
        input := &amp.CreateRuleGroupsNamespaceInput{}
        output := &amp.CreateRuleGroupsNamespaceOutput{}

        mockClient.On("CreateRuleGroupsNamespace", ctx, input).Return(output, nil)

        result, err := mockClient.CreateRuleGroupsNamespace(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateScraper", func(t *testing.T) {
        input := &amp.CreateScraperInput{}
        output := &amp.CreateScraperOutput{}

        mockClient.On("CreateScraper", ctx, input).Return(output, nil)

        result, err := mockClient.CreateScraper(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateWorkspace", func(t *testing.T) {
        input := &amp.CreateWorkspaceInput{}
        output := &amp.CreateWorkspaceOutput{}

        mockClient.On("CreateWorkspace", ctx, input).Return(output, nil)

        result, err := mockClient.CreateWorkspace(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteAlertManagerDefinition", func(t *testing.T) {
        input := &amp.DeleteAlertManagerDefinitionInput{}
        output := &amp.DeleteAlertManagerDefinitionOutput{}

        mockClient.On("DeleteAlertManagerDefinition", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteAlertManagerDefinition(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteLoggingConfiguration", func(t *testing.T) {
        input := &amp.DeleteLoggingConfigurationInput{}
        output := &amp.DeleteLoggingConfigurationOutput{}

        mockClient.On("DeleteLoggingConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteLoggingConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteRuleGroupsNamespace", func(t *testing.T) {
        input := &amp.DeleteRuleGroupsNamespaceInput{}
        output := &amp.DeleteRuleGroupsNamespaceOutput{}

        mockClient.On("DeleteRuleGroupsNamespace", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteRuleGroupsNamespace(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteScraper", func(t *testing.T) {
        input := &amp.DeleteScraperInput{}
        output := &amp.DeleteScraperOutput{}

        mockClient.On("DeleteScraper", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteScraper(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteWorkspace", func(t *testing.T) {
        input := &amp.DeleteWorkspaceInput{}
        output := &amp.DeleteWorkspaceOutput{}

        mockClient.On("DeleteWorkspace", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteWorkspace(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeAlertManagerDefinition", func(t *testing.T) {
        input := &amp.DescribeAlertManagerDefinitionInput{}
        output := &amp.DescribeAlertManagerDefinitionOutput{}

        mockClient.On("DescribeAlertManagerDefinition", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeAlertManagerDefinition(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeLoggingConfiguration", func(t *testing.T) {
        input := &amp.DescribeLoggingConfigurationInput{}
        output := &amp.DescribeLoggingConfigurationOutput{}

        mockClient.On("DescribeLoggingConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeLoggingConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeRuleGroupsNamespace", func(t *testing.T) {
        input := &amp.DescribeRuleGroupsNamespaceInput{}
        output := &amp.DescribeRuleGroupsNamespaceOutput{}

        mockClient.On("DescribeRuleGroupsNamespace", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeRuleGroupsNamespace(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeScraper", func(t *testing.T) {
        input := &amp.DescribeScraperInput{}
        output := &amp.DescribeScraperOutput{}

        mockClient.On("DescribeScraper", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeScraper(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeWorkspace", func(t *testing.T) {
        input := &amp.DescribeWorkspaceInput{}
        output := &amp.DescribeWorkspaceOutput{}

        mockClient.On("DescribeWorkspace", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeWorkspace(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDefaultScraperConfiguration", func(t *testing.T) {
        input := &amp.GetDefaultScraperConfigurationInput{}
        output := &amp.GetDefaultScraperConfigurationOutput{}

        mockClient.On("GetDefaultScraperConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.GetDefaultScraperConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListRuleGroupsNamespaces", func(t *testing.T) {
        input := &amp.ListRuleGroupsNamespacesInput{}
        output := &amp.ListRuleGroupsNamespacesOutput{}

        mockClient.On("ListRuleGroupsNamespaces", ctx, input).Return(output, nil)

        result, err := mockClient.ListRuleGroupsNamespaces(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListScrapers", func(t *testing.T) {
        input := &amp.ListScrapersInput{}
        output := &amp.ListScrapersOutput{}

        mockClient.On("ListScrapers", ctx, input).Return(output, nil)

        result, err := mockClient.ListScrapers(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &amp.ListTagsForResourceInput{}
        output := &amp.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListWorkspaces", func(t *testing.T) {
        input := &amp.ListWorkspacesInput{}
        output := &amp.ListWorkspacesOutput{}

        mockClient.On("ListWorkspaces", ctx, input).Return(output, nil)

        result, err := mockClient.ListWorkspaces(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutAlertManagerDefinition", func(t *testing.T) {
        input := &amp.PutAlertManagerDefinitionInput{}
        output := &amp.PutAlertManagerDefinitionOutput{}

        mockClient.On("PutAlertManagerDefinition", ctx, input).Return(output, nil)

        result, err := mockClient.PutAlertManagerDefinition(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutRuleGroupsNamespace", func(t *testing.T) {
        input := &amp.PutRuleGroupsNamespaceInput{}
        output := &amp.PutRuleGroupsNamespaceOutput{}

        mockClient.On("PutRuleGroupsNamespace", ctx, input).Return(output, nil)

        result, err := mockClient.PutRuleGroupsNamespace(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &amp.TagResourceInput{}
        output := &amp.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &amp.UntagResourceInput{}
        output := &amp.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateLoggingConfiguration", func(t *testing.T) {
        input := &amp.UpdateLoggingConfigurationInput{}
        output := &amp.UpdateLoggingConfigurationOutput{}

        mockClient.On("UpdateLoggingConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateLoggingConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateWorkspaceAlias", func(t *testing.T) {
        input := &amp.UpdateWorkspaceAliasInput{}
        output := &amp.UpdateWorkspaceAliasOutput{}

        mockClient.On("UpdateWorkspaceAlias", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateWorkspaceAlias(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
