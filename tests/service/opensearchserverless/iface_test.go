package opensearchserverless_test

// tests for the opensearchserverless service interface for this ../../../service/opensearchserverless/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/opensearchserverless"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/opensearchserverless/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/opensearchserverless/opensearchserverless_iface"
	"github.com/stretchr/testify/assert"
)

func TestOpensearchserverlessServiceCanBeMocked(t *testing.T) {
	var iface opensearchserverless_iface.IClient
	iface = &opensearchserverless.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := opensearchserverless.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchGetCollection", func(t *testing.T) {
        input := &opensearchserverless.BatchGetCollectionInput{}
        output := &opensearchserverless.BatchGetCollectionOutput{}

        mockClient.On("BatchGetCollection", ctx, input).Return(output, nil)

        result, err := mockClient.BatchGetCollection(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchGetEffectiveLifecyclePolicy", func(t *testing.T) {
        input := &opensearchserverless.BatchGetEffectiveLifecyclePolicyInput{}
        output := &opensearchserverless.BatchGetEffectiveLifecyclePolicyOutput{}

        mockClient.On("BatchGetEffectiveLifecyclePolicy", ctx, input).Return(output, nil)

        result, err := mockClient.BatchGetEffectiveLifecyclePolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchGetLifecyclePolicy", func(t *testing.T) {
        input := &opensearchserverless.BatchGetLifecyclePolicyInput{}
        output := &opensearchserverless.BatchGetLifecyclePolicyOutput{}

        mockClient.On("BatchGetLifecyclePolicy", ctx, input).Return(output, nil)

        result, err := mockClient.BatchGetLifecyclePolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchGetVpcEndpoint", func(t *testing.T) {
        input := &opensearchserverless.BatchGetVpcEndpointInput{}
        output := &opensearchserverless.BatchGetVpcEndpointOutput{}

        mockClient.On("BatchGetVpcEndpoint", ctx, input).Return(output, nil)

        result, err := mockClient.BatchGetVpcEndpoint(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateAccessPolicy", func(t *testing.T) {
        input := &opensearchserverless.CreateAccessPolicyInput{}
        output := &opensearchserverless.CreateAccessPolicyOutput{}

        mockClient.On("CreateAccessPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.CreateAccessPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateCollection", func(t *testing.T) {
        input := &opensearchserverless.CreateCollectionInput{}
        output := &opensearchserverless.CreateCollectionOutput{}

        mockClient.On("CreateCollection", ctx, input).Return(output, nil)

        result, err := mockClient.CreateCollection(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateLifecyclePolicy", func(t *testing.T) {
        input := &opensearchserverless.CreateLifecyclePolicyInput{}
        output := &opensearchserverless.CreateLifecyclePolicyOutput{}

        mockClient.On("CreateLifecyclePolicy", ctx, input).Return(output, nil)

        result, err := mockClient.CreateLifecyclePolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateSecurityConfig", func(t *testing.T) {
        input := &opensearchserverless.CreateSecurityConfigInput{}
        output := &opensearchserverless.CreateSecurityConfigOutput{}

        mockClient.On("CreateSecurityConfig", ctx, input).Return(output, nil)

        result, err := mockClient.CreateSecurityConfig(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateSecurityPolicy", func(t *testing.T) {
        input := &opensearchserverless.CreateSecurityPolicyInput{}
        output := &opensearchserverless.CreateSecurityPolicyOutput{}

        mockClient.On("CreateSecurityPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.CreateSecurityPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateVpcEndpoint", func(t *testing.T) {
        input := &opensearchserverless.CreateVpcEndpointInput{}
        output := &opensearchserverless.CreateVpcEndpointOutput{}

        mockClient.On("CreateVpcEndpoint", ctx, input).Return(output, nil)

        result, err := mockClient.CreateVpcEndpoint(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteAccessPolicy", func(t *testing.T) {
        input := &opensearchserverless.DeleteAccessPolicyInput{}
        output := &opensearchserverless.DeleteAccessPolicyOutput{}

        mockClient.On("DeleteAccessPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteAccessPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteCollection", func(t *testing.T) {
        input := &opensearchserverless.DeleteCollectionInput{}
        output := &opensearchserverless.DeleteCollectionOutput{}

        mockClient.On("DeleteCollection", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteCollection(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteLifecyclePolicy", func(t *testing.T) {
        input := &opensearchserverless.DeleteLifecyclePolicyInput{}
        output := &opensearchserverless.DeleteLifecyclePolicyOutput{}

        mockClient.On("DeleteLifecyclePolicy", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteLifecyclePolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteSecurityConfig", func(t *testing.T) {
        input := &opensearchserverless.DeleteSecurityConfigInput{}
        output := &opensearchserverless.DeleteSecurityConfigOutput{}

        mockClient.On("DeleteSecurityConfig", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteSecurityConfig(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteSecurityPolicy", func(t *testing.T) {
        input := &opensearchserverless.DeleteSecurityPolicyInput{}
        output := &opensearchserverless.DeleteSecurityPolicyOutput{}

        mockClient.On("DeleteSecurityPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteSecurityPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteVpcEndpoint", func(t *testing.T) {
        input := &opensearchserverless.DeleteVpcEndpointInput{}
        output := &opensearchserverless.DeleteVpcEndpointOutput{}

        mockClient.On("DeleteVpcEndpoint", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteVpcEndpoint(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetAccessPolicy", func(t *testing.T) {
        input := &opensearchserverless.GetAccessPolicyInput{}
        output := &opensearchserverless.GetAccessPolicyOutput{}

        mockClient.On("GetAccessPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.GetAccessPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetAccountSettings", func(t *testing.T) {
        input := &opensearchserverless.GetAccountSettingsInput{}
        output := &opensearchserverless.GetAccountSettingsOutput{}

        mockClient.On("GetAccountSettings", ctx, input).Return(output, nil)

        result, err := mockClient.GetAccountSettings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetPoliciesStats", func(t *testing.T) {
        input := &opensearchserverless.GetPoliciesStatsInput{}
        output := &opensearchserverless.GetPoliciesStatsOutput{}

        mockClient.On("GetPoliciesStats", ctx, input).Return(output, nil)

        result, err := mockClient.GetPoliciesStats(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSecurityConfig", func(t *testing.T) {
        input := &opensearchserverless.GetSecurityConfigInput{}
        output := &opensearchserverless.GetSecurityConfigOutput{}

        mockClient.On("GetSecurityConfig", ctx, input).Return(output, nil)

        result, err := mockClient.GetSecurityConfig(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSecurityPolicy", func(t *testing.T) {
        input := &opensearchserverless.GetSecurityPolicyInput{}
        output := &opensearchserverless.GetSecurityPolicyOutput{}

        mockClient.On("GetSecurityPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.GetSecurityPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAccessPolicies", func(t *testing.T) {
        input := &opensearchserverless.ListAccessPoliciesInput{}
        output := &opensearchserverless.ListAccessPoliciesOutput{}

        mockClient.On("ListAccessPolicies", ctx, input).Return(output, nil)

        result, err := mockClient.ListAccessPolicies(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListCollections", func(t *testing.T) {
        input := &opensearchserverless.ListCollectionsInput{}
        output := &opensearchserverless.ListCollectionsOutput{}

        mockClient.On("ListCollections", ctx, input).Return(output, nil)

        result, err := mockClient.ListCollections(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListLifecyclePolicies", func(t *testing.T) {
        input := &opensearchserverless.ListLifecyclePoliciesInput{}
        output := &opensearchserverless.ListLifecyclePoliciesOutput{}

        mockClient.On("ListLifecyclePolicies", ctx, input).Return(output, nil)

        result, err := mockClient.ListLifecyclePolicies(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListSecurityConfigs", func(t *testing.T) {
        input := &opensearchserverless.ListSecurityConfigsInput{}
        output := &opensearchserverless.ListSecurityConfigsOutput{}

        mockClient.On("ListSecurityConfigs", ctx, input).Return(output, nil)

        result, err := mockClient.ListSecurityConfigs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListSecurityPolicies", func(t *testing.T) {
        input := &opensearchserverless.ListSecurityPoliciesInput{}
        output := &opensearchserverless.ListSecurityPoliciesOutput{}

        mockClient.On("ListSecurityPolicies", ctx, input).Return(output, nil)

        result, err := mockClient.ListSecurityPolicies(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &opensearchserverless.ListTagsForResourceInput{}
        output := &opensearchserverless.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListVpcEndpoints", func(t *testing.T) {
        input := &opensearchserverless.ListVpcEndpointsInput{}
        output := &opensearchserverless.ListVpcEndpointsOutput{}

        mockClient.On("ListVpcEndpoints", ctx, input).Return(output, nil)

        result, err := mockClient.ListVpcEndpoints(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &opensearchserverless.TagResourceInput{}
        output := &opensearchserverless.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &opensearchserverless.UntagResourceInput{}
        output := &opensearchserverless.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateAccessPolicy", func(t *testing.T) {
        input := &opensearchserverless.UpdateAccessPolicyInput{}
        output := &opensearchserverless.UpdateAccessPolicyOutput{}

        mockClient.On("UpdateAccessPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateAccessPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateAccountSettings", func(t *testing.T) {
        input := &opensearchserverless.UpdateAccountSettingsInput{}
        output := &opensearchserverless.UpdateAccountSettingsOutput{}

        mockClient.On("UpdateAccountSettings", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateAccountSettings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateCollection", func(t *testing.T) {
        input := &opensearchserverless.UpdateCollectionInput{}
        output := &opensearchserverless.UpdateCollectionOutput{}

        mockClient.On("UpdateCollection", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateCollection(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateLifecyclePolicy", func(t *testing.T) {
        input := &opensearchserverless.UpdateLifecyclePolicyInput{}
        output := &opensearchserverless.UpdateLifecyclePolicyOutput{}

        mockClient.On("UpdateLifecyclePolicy", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateLifecyclePolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateSecurityConfig", func(t *testing.T) {
        input := &opensearchserverless.UpdateSecurityConfigInput{}
        output := &opensearchserverless.UpdateSecurityConfigOutput{}

        mockClient.On("UpdateSecurityConfig", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateSecurityConfig(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateSecurityPolicy", func(t *testing.T) {
        input := &opensearchserverless.UpdateSecurityPolicyInput{}
        output := &opensearchserverless.UpdateSecurityPolicyOutput{}

        mockClient.On("UpdateSecurityPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateSecurityPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateVpcEndpoint", func(t *testing.T) {
        input := &opensearchserverless.UpdateVpcEndpointInput{}
        output := &opensearchserverless.UpdateVpcEndpointOutput{}

        mockClient.On("UpdateVpcEndpoint", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateVpcEndpoint(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
