package grafana_test

// tests for the grafana service interface for this ../../../service/grafana/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/grafana"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/grafana/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/grafana/grafana_iface"
	"github.com/stretchr/testify/assert"
)

func TestGrafanaServiceCanBeMocked(t *testing.T) {
	var iface grafana_iface.IClient
	iface = &grafana.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := grafana.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateLicense", func(t *testing.T) {
        input := &grafana.AssociateLicenseInput{}
        output := &grafana.AssociateLicenseOutput{}

        mockClient.On("AssociateLicense", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateLicense(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateWorkspace", func(t *testing.T) {
        input := &grafana.CreateWorkspaceInput{}
        output := &grafana.CreateWorkspaceOutput{}

        mockClient.On("CreateWorkspace", ctx, input).Return(output, nil)

        result, err := mockClient.CreateWorkspace(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateWorkspaceApiKey", func(t *testing.T) {
        input := &grafana.CreateWorkspaceApiKeyInput{}
        output := &grafana.CreateWorkspaceApiKeyOutput{}

        mockClient.On("CreateWorkspaceApiKey", ctx, input).Return(output, nil)

        result, err := mockClient.CreateWorkspaceApiKey(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateWorkspaceServiceAccount", func(t *testing.T) {
        input := &grafana.CreateWorkspaceServiceAccountInput{}
        output := &grafana.CreateWorkspaceServiceAccountOutput{}

        mockClient.On("CreateWorkspaceServiceAccount", ctx, input).Return(output, nil)

        result, err := mockClient.CreateWorkspaceServiceAccount(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateWorkspaceServiceAccountToken", func(t *testing.T) {
        input := &grafana.CreateWorkspaceServiceAccountTokenInput{}
        output := &grafana.CreateWorkspaceServiceAccountTokenOutput{}

        mockClient.On("CreateWorkspaceServiceAccountToken", ctx, input).Return(output, nil)

        result, err := mockClient.CreateWorkspaceServiceAccountToken(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteWorkspace", func(t *testing.T) {
        input := &grafana.DeleteWorkspaceInput{}
        output := &grafana.DeleteWorkspaceOutput{}

        mockClient.On("DeleteWorkspace", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteWorkspace(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteWorkspaceApiKey", func(t *testing.T) {
        input := &grafana.DeleteWorkspaceApiKeyInput{}
        output := &grafana.DeleteWorkspaceApiKeyOutput{}

        mockClient.On("DeleteWorkspaceApiKey", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteWorkspaceApiKey(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteWorkspaceServiceAccount", func(t *testing.T) {
        input := &grafana.DeleteWorkspaceServiceAccountInput{}
        output := &grafana.DeleteWorkspaceServiceAccountOutput{}

        mockClient.On("DeleteWorkspaceServiceAccount", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteWorkspaceServiceAccount(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteWorkspaceServiceAccountToken", func(t *testing.T) {
        input := &grafana.DeleteWorkspaceServiceAccountTokenInput{}
        output := &grafana.DeleteWorkspaceServiceAccountTokenOutput{}

        mockClient.On("DeleteWorkspaceServiceAccountToken", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteWorkspaceServiceAccountToken(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeWorkspace", func(t *testing.T) {
        input := &grafana.DescribeWorkspaceInput{}
        output := &grafana.DescribeWorkspaceOutput{}

        mockClient.On("DescribeWorkspace", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeWorkspace(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeWorkspaceAuthentication", func(t *testing.T) {
        input := &grafana.DescribeWorkspaceAuthenticationInput{}
        output := &grafana.DescribeWorkspaceAuthenticationOutput{}

        mockClient.On("DescribeWorkspaceAuthentication", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeWorkspaceAuthentication(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeWorkspaceConfiguration", func(t *testing.T) {
        input := &grafana.DescribeWorkspaceConfigurationInput{}
        output := &grafana.DescribeWorkspaceConfigurationOutput{}

        mockClient.On("DescribeWorkspaceConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeWorkspaceConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateLicense", func(t *testing.T) {
        input := &grafana.DisassociateLicenseInput{}
        output := &grafana.DisassociateLicenseOutput{}

        mockClient.On("DisassociateLicense", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateLicense(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListPermissions", func(t *testing.T) {
        input := &grafana.ListPermissionsInput{}
        output := &grafana.ListPermissionsOutput{}

        mockClient.On("ListPermissions", ctx, input).Return(output, nil)

        result, err := mockClient.ListPermissions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &grafana.ListTagsForResourceInput{}
        output := &grafana.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListVersions", func(t *testing.T) {
        input := &grafana.ListVersionsInput{}
        output := &grafana.ListVersionsOutput{}

        mockClient.On("ListVersions", ctx, input).Return(output, nil)

        result, err := mockClient.ListVersions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListWorkspaceServiceAccountTokens", func(t *testing.T) {
        input := &grafana.ListWorkspaceServiceAccountTokensInput{}
        output := &grafana.ListWorkspaceServiceAccountTokensOutput{}

        mockClient.On("ListWorkspaceServiceAccountTokens", ctx, input).Return(output, nil)

        result, err := mockClient.ListWorkspaceServiceAccountTokens(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListWorkspaceServiceAccounts", func(t *testing.T) {
        input := &grafana.ListWorkspaceServiceAccountsInput{}
        output := &grafana.ListWorkspaceServiceAccountsOutput{}

        mockClient.On("ListWorkspaceServiceAccounts", ctx, input).Return(output, nil)

        result, err := mockClient.ListWorkspaceServiceAccounts(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListWorkspaces", func(t *testing.T) {
        input := &grafana.ListWorkspacesInput{}
        output := &grafana.ListWorkspacesOutput{}

        mockClient.On("ListWorkspaces", ctx, input).Return(output, nil)

        result, err := mockClient.ListWorkspaces(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &grafana.TagResourceInput{}
        output := &grafana.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &grafana.UntagResourceInput{}
        output := &grafana.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdatePermissions", func(t *testing.T) {
        input := &grafana.UpdatePermissionsInput{}
        output := &grafana.UpdatePermissionsOutput{}

        mockClient.On("UpdatePermissions", ctx, input).Return(output, nil)

        result, err := mockClient.UpdatePermissions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateWorkspace", func(t *testing.T) {
        input := &grafana.UpdateWorkspaceInput{}
        output := &grafana.UpdateWorkspaceOutput{}

        mockClient.On("UpdateWorkspace", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateWorkspace(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateWorkspaceAuthentication", func(t *testing.T) {
        input := &grafana.UpdateWorkspaceAuthenticationInput{}
        output := &grafana.UpdateWorkspaceAuthenticationOutput{}

        mockClient.On("UpdateWorkspaceAuthentication", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateWorkspaceAuthentication(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateWorkspaceConfiguration", func(t *testing.T) {
        input := &grafana.UpdateWorkspaceConfigurationInput{}
        output := &grafana.UpdateWorkspaceConfigurationOutput{}

        mockClient.On("UpdateWorkspaceConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateWorkspaceConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
