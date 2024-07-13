package cloudfront_test

// tests for the cloudfront service interface for this ../../../service/cloudfront/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/cloudfront"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/cloudfront/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/cloudfront/cloudfront_iface"
	"github.com/stretchr/testify/assert"
)

func TestCloudfrontServiceCanBeMocked(t *testing.T) {
	var iface cloudfront_iface.IClient
	iface = &cloudfront.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := cloudfront.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateAlias", func(t *testing.T) {
        input := &cloudfront.AssociateAliasInput{}
        output := &cloudfront.AssociateAliasOutput{}

        mockClient.On("AssociateAlias", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateAlias(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCopyDistribution", func(t *testing.T) {
        input := &cloudfront.CopyDistributionInput{}
        output := &cloudfront.CopyDistributionOutput{}

        mockClient.On("CopyDistribution", ctx, input).Return(output, nil)

        result, err := mockClient.CopyDistribution(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateCachePolicy", func(t *testing.T) {
        input := &cloudfront.CreateCachePolicyInput{}
        output := &cloudfront.CreateCachePolicyOutput{}

        mockClient.On("CreateCachePolicy", ctx, input).Return(output, nil)

        result, err := mockClient.CreateCachePolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateCloudFrontOriginAccessIdentity", func(t *testing.T) {
        input := &cloudfront.CreateCloudFrontOriginAccessIdentityInput{}
        output := &cloudfront.CreateCloudFrontOriginAccessIdentityOutput{}

        mockClient.On("CreateCloudFrontOriginAccessIdentity", ctx, input).Return(output, nil)

        result, err := mockClient.CreateCloudFrontOriginAccessIdentity(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateContinuousDeploymentPolicy", func(t *testing.T) {
        input := &cloudfront.CreateContinuousDeploymentPolicyInput{}
        output := &cloudfront.CreateContinuousDeploymentPolicyOutput{}

        mockClient.On("CreateContinuousDeploymentPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.CreateContinuousDeploymentPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateDistribution", func(t *testing.T) {
        input := &cloudfront.CreateDistributionInput{}
        output := &cloudfront.CreateDistributionOutput{}

        mockClient.On("CreateDistribution", ctx, input).Return(output, nil)

        result, err := mockClient.CreateDistribution(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateDistributionWithTags", func(t *testing.T) {
        input := &cloudfront.CreateDistributionWithTagsInput{}
        output := &cloudfront.CreateDistributionWithTagsOutput{}

        mockClient.On("CreateDistributionWithTags", ctx, input).Return(output, nil)

        result, err := mockClient.CreateDistributionWithTags(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateFieldLevelEncryptionConfig", func(t *testing.T) {
        input := &cloudfront.CreateFieldLevelEncryptionConfigInput{}
        output := &cloudfront.CreateFieldLevelEncryptionConfigOutput{}

        mockClient.On("CreateFieldLevelEncryptionConfig", ctx, input).Return(output, nil)

        result, err := mockClient.CreateFieldLevelEncryptionConfig(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateFieldLevelEncryptionProfile", func(t *testing.T) {
        input := &cloudfront.CreateFieldLevelEncryptionProfileInput{}
        output := &cloudfront.CreateFieldLevelEncryptionProfileOutput{}

        mockClient.On("CreateFieldLevelEncryptionProfile", ctx, input).Return(output, nil)

        result, err := mockClient.CreateFieldLevelEncryptionProfile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateFunction", func(t *testing.T) {
        input := &cloudfront.CreateFunctionInput{}
        output := &cloudfront.CreateFunctionOutput{}

        mockClient.On("CreateFunction", ctx, input).Return(output, nil)

        result, err := mockClient.CreateFunction(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateInvalidation", func(t *testing.T) {
        input := &cloudfront.CreateInvalidationInput{}
        output := &cloudfront.CreateInvalidationOutput{}

        mockClient.On("CreateInvalidation", ctx, input).Return(output, nil)

        result, err := mockClient.CreateInvalidation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateKeyGroup", func(t *testing.T) {
        input := &cloudfront.CreateKeyGroupInput{}
        output := &cloudfront.CreateKeyGroupOutput{}

        mockClient.On("CreateKeyGroup", ctx, input).Return(output, nil)

        result, err := mockClient.CreateKeyGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateKeyValueStore", func(t *testing.T) {
        input := &cloudfront.CreateKeyValueStoreInput{}
        output := &cloudfront.CreateKeyValueStoreOutput{}

        mockClient.On("CreateKeyValueStore", ctx, input).Return(output, nil)

        result, err := mockClient.CreateKeyValueStore(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateMonitoringSubscription", func(t *testing.T) {
        input := &cloudfront.CreateMonitoringSubscriptionInput{}
        output := &cloudfront.CreateMonitoringSubscriptionOutput{}

        mockClient.On("CreateMonitoringSubscription", ctx, input).Return(output, nil)

        result, err := mockClient.CreateMonitoringSubscription(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateOriginAccessControl", func(t *testing.T) {
        input := &cloudfront.CreateOriginAccessControlInput{}
        output := &cloudfront.CreateOriginAccessControlOutput{}

        mockClient.On("CreateOriginAccessControl", ctx, input).Return(output, nil)

        result, err := mockClient.CreateOriginAccessControl(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateOriginRequestPolicy", func(t *testing.T) {
        input := &cloudfront.CreateOriginRequestPolicyInput{}
        output := &cloudfront.CreateOriginRequestPolicyOutput{}

        mockClient.On("CreateOriginRequestPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.CreateOriginRequestPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreatePublicKey", func(t *testing.T) {
        input := &cloudfront.CreatePublicKeyInput{}
        output := &cloudfront.CreatePublicKeyOutput{}

        mockClient.On("CreatePublicKey", ctx, input).Return(output, nil)

        result, err := mockClient.CreatePublicKey(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateRealtimeLogConfig", func(t *testing.T) {
        input := &cloudfront.CreateRealtimeLogConfigInput{}
        output := &cloudfront.CreateRealtimeLogConfigOutput{}

        mockClient.On("CreateRealtimeLogConfig", ctx, input).Return(output, nil)

        result, err := mockClient.CreateRealtimeLogConfig(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateResponseHeadersPolicy", func(t *testing.T) {
        input := &cloudfront.CreateResponseHeadersPolicyInput{}
        output := &cloudfront.CreateResponseHeadersPolicyOutput{}

        mockClient.On("CreateResponseHeadersPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.CreateResponseHeadersPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateStreamingDistribution", func(t *testing.T) {
        input := &cloudfront.CreateStreamingDistributionInput{}
        output := &cloudfront.CreateStreamingDistributionOutput{}

        mockClient.On("CreateStreamingDistribution", ctx, input).Return(output, nil)

        result, err := mockClient.CreateStreamingDistribution(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateStreamingDistributionWithTags", func(t *testing.T) {
        input := &cloudfront.CreateStreamingDistributionWithTagsInput{}
        output := &cloudfront.CreateStreamingDistributionWithTagsOutput{}

        mockClient.On("CreateStreamingDistributionWithTags", ctx, input).Return(output, nil)

        result, err := mockClient.CreateStreamingDistributionWithTags(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteCachePolicy", func(t *testing.T) {
        input := &cloudfront.DeleteCachePolicyInput{}
        output := &cloudfront.DeleteCachePolicyOutput{}

        mockClient.On("DeleteCachePolicy", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteCachePolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteCloudFrontOriginAccessIdentity", func(t *testing.T) {
        input := &cloudfront.DeleteCloudFrontOriginAccessIdentityInput{}
        output := &cloudfront.DeleteCloudFrontOriginAccessIdentityOutput{}

        mockClient.On("DeleteCloudFrontOriginAccessIdentity", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteCloudFrontOriginAccessIdentity(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteContinuousDeploymentPolicy", func(t *testing.T) {
        input := &cloudfront.DeleteContinuousDeploymentPolicyInput{}
        output := &cloudfront.DeleteContinuousDeploymentPolicyOutput{}

        mockClient.On("DeleteContinuousDeploymentPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteContinuousDeploymentPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteDistribution", func(t *testing.T) {
        input := &cloudfront.DeleteDistributionInput{}
        output := &cloudfront.DeleteDistributionOutput{}

        mockClient.On("DeleteDistribution", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteDistribution(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteFieldLevelEncryptionConfig", func(t *testing.T) {
        input := &cloudfront.DeleteFieldLevelEncryptionConfigInput{}
        output := &cloudfront.DeleteFieldLevelEncryptionConfigOutput{}

        mockClient.On("DeleteFieldLevelEncryptionConfig", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteFieldLevelEncryptionConfig(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteFieldLevelEncryptionProfile", func(t *testing.T) {
        input := &cloudfront.DeleteFieldLevelEncryptionProfileInput{}
        output := &cloudfront.DeleteFieldLevelEncryptionProfileOutput{}

        mockClient.On("DeleteFieldLevelEncryptionProfile", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteFieldLevelEncryptionProfile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteFunction", func(t *testing.T) {
        input := &cloudfront.DeleteFunctionInput{}
        output := &cloudfront.DeleteFunctionOutput{}

        mockClient.On("DeleteFunction", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteFunction(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteKeyGroup", func(t *testing.T) {
        input := &cloudfront.DeleteKeyGroupInput{}
        output := &cloudfront.DeleteKeyGroupOutput{}

        mockClient.On("DeleteKeyGroup", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteKeyGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteKeyValueStore", func(t *testing.T) {
        input := &cloudfront.DeleteKeyValueStoreInput{}
        output := &cloudfront.DeleteKeyValueStoreOutput{}

        mockClient.On("DeleteKeyValueStore", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteKeyValueStore(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteMonitoringSubscription", func(t *testing.T) {
        input := &cloudfront.DeleteMonitoringSubscriptionInput{}
        output := &cloudfront.DeleteMonitoringSubscriptionOutput{}

        mockClient.On("DeleteMonitoringSubscription", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteMonitoringSubscription(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteOriginAccessControl", func(t *testing.T) {
        input := &cloudfront.DeleteOriginAccessControlInput{}
        output := &cloudfront.DeleteOriginAccessControlOutput{}

        mockClient.On("DeleteOriginAccessControl", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteOriginAccessControl(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteOriginRequestPolicy", func(t *testing.T) {
        input := &cloudfront.DeleteOriginRequestPolicyInput{}
        output := &cloudfront.DeleteOriginRequestPolicyOutput{}

        mockClient.On("DeleteOriginRequestPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteOriginRequestPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeletePublicKey", func(t *testing.T) {
        input := &cloudfront.DeletePublicKeyInput{}
        output := &cloudfront.DeletePublicKeyOutput{}

        mockClient.On("DeletePublicKey", ctx, input).Return(output, nil)

        result, err := mockClient.DeletePublicKey(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteRealtimeLogConfig", func(t *testing.T) {
        input := &cloudfront.DeleteRealtimeLogConfigInput{}
        output := &cloudfront.DeleteRealtimeLogConfigOutput{}

        mockClient.On("DeleteRealtimeLogConfig", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteRealtimeLogConfig(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteResponseHeadersPolicy", func(t *testing.T) {
        input := &cloudfront.DeleteResponseHeadersPolicyInput{}
        output := &cloudfront.DeleteResponseHeadersPolicyOutput{}

        mockClient.On("DeleteResponseHeadersPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteResponseHeadersPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteStreamingDistribution", func(t *testing.T) {
        input := &cloudfront.DeleteStreamingDistributionInput{}
        output := &cloudfront.DeleteStreamingDistributionOutput{}

        mockClient.On("DeleteStreamingDistribution", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteStreamingDistribution(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeFunction", func(t *testing.T) {
        input := &cloudfront.DescribeFunctionInput{}
        output := &cloudfront.DescribeFunctionOutput{}

        mockClient.On("DescribeFunction", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeFunction(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeKeyValueStore", func(t *testing.T) {
        input := &cloudfront.DescribeKeyValueStoreInput{}
        output := &cloudfront.DescribeKeyValueStoreOutput{}

        mockClient.On("DescribeKeyValueStore", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeKeyValueStore(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetCachePolicy", func(t *testing.T) {
        input := &cloudfront.GetCachePolicyInput{}
        output := &cloudfront.GetCachePolicyOutput{}

        mockClient.On("GetCachePolicy", ctx, input).Return(output, nil)

        result, err := mockClient.GetCachePolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetCachePolicyConfig", func(t *testing.T) {
        input := &cloudfront.GetCachePolicyConfigInput{}
        output := &cloudfront.GetCachePolicyConfigOutput{}

        mockClient.On("GetCachePolicyConfig", ctx, input).Return(output, nil)

        result, err := mockClient.GetCachePolicyConfig(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetCloudFrontOriginAccessIdentity", func(t *testing.T) {
        input := &cloudfront.GetCloudFrontOriginAccessIdentityInput{}
        output := &cloudfront.GetCloudFrontOriginAccessIdentityOutput{}

        mockClient.On("GetCloudFrontOriginAccessIdentity", ctx, input).Return(output, nil)

        result, err := mockClient.GetCloudFrontOriginAccessIdentity(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetCloudFrontOriginAccessIdentityConfig", func(t *testing.T) {
        input := &cloudfront.GetCloudFrontOriginAccessIdentityConfigInput{}
        output := &cloudfront.GetCloudFrontOriginAccessIdentityConfigOutput{}

        mockClient.On("GetCloudFrontOriginAccessIdentityConfig", ctx, input).Return(output, nil)

        result, err := mockClient.GetCloudFrontOriginAccessIdentityConfig(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetContinuousDeploymentPolicy", func(t *testing.T) {
        input := &cloudfront.GetContinuousDeploymentPolicyInput{}
        output := &cloudfront.GetContinuousDeploymentPolicyOutput{}

        mockClient.On("GetContinuousDeploymentPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.GetContinuousDeploymentPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetContinuousDeploymentPolicyConfig", func(t *testing.T) {
        input := &cloudfront.GetContinuousDeploymentPolicyConfigInput{}
        output := &cloudfront.GetContinuousDeploymentPolicyConfigOutput{}

        mockClient.On("GetContinuousDeploymentPolicyConfig", ctx, input).Return(output, nil)

        result, err := mockClient.GetContinuousDeploymentPolicyConfig(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDistribution", func(t *testing.T) {
        input := &cloudfront.GetDistributionInput{}
        output := &cloudfront.GetDistributionOutput{}

        mockClient.On("GetDistribution", ctx, input).Return(output, nil)

        result, err := mockClient.GetDistribution(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDistributionConfig", func(t *testing.T) {
        input := &cloudfront.GetDistributionConfigInput{}
        output := &cloudfront.GetDistributionConfigOutput{}

        mockClient.On("GetDistributionConfig", ctx, input).Return(output, nil)

        result, err := mockClient.GetDistributionConfig(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetFieldLevelEncryption", func(t *testing.T) {
        input := &cloudfront.GetFieldLevelEncryptionInput{}
        output := &cloudfront.GetFieldLevelEncryptionOutput{}

        mockClient.On("GetFieldLevelEncryption", ctx, input).Return(output, nil)

        result, err := mockClient.GetFieldLevelEncryption(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetFieldLevelEncryptionConfig", func(t *testing.T) {
        input := &cloudfront.GetFieldLevelEncryptionConfigInput{}
        output := &cloudfront.GetFieldLevelEncryptionConfigOutput{}

        mockClient.On("GetFieldLevelEncryptionConfig", ctx, input).Return(output, nil)

        result, err := mockClient.GetFieldLevelEncryptionConfig(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetFieldLevelEncryptionProfile", func(t *testing.T) {
        input := &cloudfront.GetFieldLevelEncryptionProfileInput{}
        output := &cloudfront.GetFieldLevelEncryptionProfileOutput{}

        mockClient.On("GetFieldLevelEncryptionProfile", ctx, input).Return(output, nil)

        result, err := mockClient.GetFieldLevelEncryptionProfile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetFieldLevelEncryptionProfileConfig", func(t *testing.T) {
        input := &cloudfront.GetFieldLevelEncryptionProfileConfigInput{}
        output := &cloudfront.GetFieldLevelEncryptionProfileConfigOutput{}

        mockClient.On("GetFieldLevelEncryptionProfileConfig", ctx, input).Return(output, nil)

        result, err := mockClient.GetFieldLevelEncryptionProfileConfig(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetFunction", func(t *testing.T) {
        input := &cloudfront.GetFunctionInput{}
        output := &cloudfront.GetFunctionOutput{}

        mockClient.On("GetFunction", ctx, input).Return(output, nil)

        result, err := mockClient.GetFunction(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetInvalidation", func(t *testing.T) {
        input := &cloudfront.GetInvalidationInput{}
        output := &cloudfront.GetInvalidationOutput{}

        mockClient.On("GetInvalidation", ctx, input).Return(output, nil)

        result, err := mockClient.GetInvalidation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetKeyGroup", func(t *testing.T) {
        input := &cloudfront.GetKeyGroupInput{}
        output := &cloudfront.GetKeyGroupOutput{}

        mockClient.On("GetKeyGroup", ctx, input).Return(output, nil)

        result, err := mockClient.GetKeyGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetKeyGroupConfig", func(t *testing.T) {
        input := &cloudfront.GetKeyGroupConfigInput{}
        output := &cloudfront.GetKeyGroupConfigOutput{}

        mockClient.On("GetKeyGroupConfig", ctx, input).Return(output, nil)

        result, err := mockClient.GetKeyGroupConfig(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetMonitoringSubscription", func(t *testing.T) {
        input := &cloudfront.GetMonitoringSubscriptionInput{}
        output := &cloudfront.GetMonitoringSubscriptionOutput{}

        mockClient.On("GetMonitoringSubscription", ctx, input).Return(output, nil)

        result, err := mockClient.GetMonitoringSubscription(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetOriginAccessControl", func(t *testing.T) {
        input := &cloudfront.GetOriginAccessControlInput{}
        output := &cloudfront.GetOriginAccessControlOutput{}

        mockClient.On("GetOriginAccessControl", ctx, input).Return(output, nil)

        result, err := mockClient.GetOriginAccessControl(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetOriginAccessControlConfig", func(t *testing.T) {
        input := &cloudfront.GetOriginAccessControlConfigInput{}
        output := &cloudfront.GetOriginAccessControlConfigOutput{}

        mockClient.On("GetOriginAccessControlConfig", ctx, input).Return(output, nil)

        result, err := mockClient.GetOriginAccessControlConfig(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetOriginRequestPolicy", func(t *testing.T) {
        input := &cloudfront.GetOriginRequestPolicyInput{}
        output := &cloudfront.GetOriginRequestPolicyOutput{}

        mockClient.On("GetOriginRequestPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.GetOriginRequestPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetOriginRequestPolicyConfig", func(t *testing.T) {
        input := &cloudfront.GetOriginRequestPolicyConfigInput{}
        output := &cloudfront.GetOriginRequestPolicyConfigOutput{}

        mockClient.On("GetOriginRequestPolicyConfig", ctx, input).Return(output, nil)

        result, err := mockClient.GetOriginRequestPolicyConfig(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetPublicKey", func(t *testing.T) {
        input := &cloudfront.GetPublicKeyInput{}
        output := &cloudfront.GetPublicKeyOutput{}

        mockClient.On("GetPublicKey", ctx, input).Return(output, nil)

        result, err := mockClient.GetPublicKey(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetPublicKeyConfig", func(t *testing.T) {
        input := &cloudfront.GetPublicKeyConfigInput{}
        output := &cloudfront.GetPublicKeyConfigOutput{}

        mockClient.On("GetPublicKeyConfig", ctx, input).Return(output, nil)

        result, err := mockClient.GetPublicKeyConfig(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetRealtimeLogConfig", func(t *testing.T) {
        input := &cloudfront.GetRealtimeLogConfigInput{}
        output := &cloudfront.GetRealtimeLogConfigOutput{}

        mockClient.On("GetRealtimeLogConfig", ctx, input).Return(output, nil)

        result, err := mockClient.GetRealtimeLogConfig(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetResponseHeadersPolicy", func(t *testing.T) {
        input := &cloudfront.GetResponseHeadersPolicyInput{}
        output := &cloudfront.GetResponseHeadersPolicyOutput{}

        mockClient.On("GetResponseHeadersPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.GetResponseHeadersPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetResponseHeadersPolicyConfig", func(t *testing.T) {
        input := &cloudfront.GetResponseHeadersPolicyConfigInput{}
        output := &cloudfront.GetResponseHeadersPolicyConfigOutput{}

        mockClient.On("GetResponseHeadersPolicyConfig", ctx, input).Return(output, nil)

        result, err := mockClient.GetResponseHeadersPolicyConfig(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetStreamingDistribution", func(t *testing.T) {
        input := &cloudfront.GetStreamingDistributionInput{}
        output := &cloudfront.GetStreamingDistributionOutput{}

        mockClient.On("GetStreamingDistribution", ctx, input).Return(output, nil)

        result, err := mockClient.GetStreamingDistribution(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetStreamingDistributionConfig", func(t *testing.T) {
        input := &cloudfront.GetStreamingDistributionConfigInput{}
        output := &cloudfront.GetStreamingDistributionConfigOutput{}

        mockClient.On("GetStreamingDistributionConfig", ctx, input).Return(output, nil)

        result, err := mockClient.GetStreamingDistributionConfig(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListCachePolicies", func(t *testing.T) {
        input := &cloudfront.ListCachePoliciesInput{}
        output := &cloudfront.ListCachePoliciesOutput{}

        mockClient.On("ListCachePolicies", ctx, input).Return(output, nil)

        result, err := mockClient.ListCachePolicies(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListCloudFrontOriginAccessIdentities", func(t *testing.T) {
        input := &cloudfront.ListCloudFrontOriginAccessIdentitiesInput{}
        output := &cloudfront.ListCloudFrontOriginAccessIdentitiesOutput{}

        mockClient.On("ListCloudFrontOriginAccessIdentities", ctx, input).Return(output, nil)

        result, err := mockClient.ListCloudFrontOriginAccessIdentities(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListConflictingAliases", func(t *testing.T) {
        input := &cloudfront.ListConflictingAliasesInput{}
        output := &cloudfront.ListConflictingAliasesOutput{}

        mockClient.On("ListConflictingAliases", ctx, input).Return(output, nil)

        result, err := mockClient.ListConflictingAliases(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListContinuousDeploymentPolicies", func(t *testing.T) {
        input := &cloudfront.ListContinuousDeploymentPoliciesInput{}
        output := &cloudfront.ListContinuousDeploymentPoliciesOutput{}

        mockClient.On("ListContinuousDeploymentPolicies", ctx, input).Return(output, nil)

        result, err := mockClient.ListContinuousDeploymentPolicies(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDistributions", func(t *testing.T) {
        input := &cloudfront.ListDistributionsInput{}
        output := &cloudfront.ListDistributionsOutput{}

        mockClient.On("ListDistributions", ctx, input).Return(output, nil)

        result, err := mockClient.ListDistributions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDistributionsByCachePolicyId", func(t *testing.T) {
        input := &cloudfront.ListDistributionsByCachePolicyIdInput{}
        output := &cloudfront.ListDistributionsByCachePolicyIdOutput{}

        mockClient.On("ListDistributionsByCachePolicyId", ctx, input).Return(output, nil)

        result, err := mockClient.ListDistributionsByCachePolicyId(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDistributionsByKeyGroup", func(t *testing.T) {
        input := &cloudfront.ListDistributionsByKeyGroupInput{}
        output := &cloudfront.ListDistributionsByKeyGroupOutput{}

        mockClient.On("ListDistributionsByKeyGroup", ctx, input).Return(output, nil)

        result, err := mockClient.ListDistributionsByKeyGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDistributionsByOriginRequestPolicyId", func(t *testing.T) {
        input := &cloudfront.ListDistributionsByOriginRequestPolicyIdInput{}
        output := &cloudfront.ListDistributionsByOriginRequestPolicyIdOutput{}

        mockClient.On("ListDistributionsByOriginRequestPolicyId", ctx, input).Return(output, nil)

        result, err := mockClient.ListDistributionsByOriginRequestPolicyId(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDistributionsByRealtimeLogConfig", func(t *testing.T) {
        input := &cloudfront.ListDistributionsByRealtimeLogConfigInput{}
        output := &cloudfront.ListDistributionsByRealtimeLogConfigOutput{}

        mockClient.On("ListDistributionsByRealtimeLogConfig", ctx, input).Return(output, nil)

        result, err := mockClient.ListDistributionsByRealtimeLogConfig(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDistributionsByResponseHeadersPolicyId", func(t *testing.T) {
        input := &cloudfront.ListDistributionsByResponseHeadersPolicyIdInput{}
        output := &cloudfront.ListDistributionsByResponseHeadersPolicyIdOutput{}

        mockClient.On("ListDistributionsByResponseHeadersPolicyId", ctx, input).Return(output, nil)

        result, err := mockClient.ListDistributionsByResponseHeadersPolicyId(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDistributionsByWebACLId", func(t *testing.T) {
        input := &cloudfront.ListDistributionsByWebACLIdInput{}
        output := &cloudfront.ListDistributionsByWebACLIdOutput{}

        mockClient.On("ListDistributionsByWebACLId", ctx, input).Return(output, nil)

        result, err := mockClient.ListDistributionsByWebACLId(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListFieldLevelEncryptionConfigs", func(t *testing.T) {
        input := &cloudfront.ListFieldLevelEncryptionConfigsInput{}
        output := &cloudfront.ListFieldLevelEncryptionConfigsOutput{}

        mockClient.On("ListFieldLevelEncryptionConfigs", ctx, input).Return(output, nil)

        result, err := mockClient.ListFieldLevelEncryptionConfigs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListFieldLevelEncryptionProfiles", func(t *testing.T) {
        input := &cloudfront.ListFieldLevelEncryptionProfilesInput{}
        output := &cloudfront.ListFieldLevelEncryptionProfilesOutput{}

        mockClient.On("ListFieldLevelEncryptionProfiles", ctx, input).Return(output, nil)

        result, err := mockClient.ListFieldLevelEncryptionProfiles(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListFunctions", func(t *testing.T) {
        input := &cloudfront.ListFunctionsInput{}
        output := &cloudfront.ListFunctionsOutput{}

        mockClient.On("ListFunctions", ctx, input).Return(output, nil)

        result, err := mockClient.ListFunctions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListInvalidations", func(t *testing.T) {
        input := &cloudfront.ListInvalidationsInput{}
        output := &cloudfront.ListInvalidationsOutput{}

        mockClient.On("ListInvalidations", ctx, input).Return(output, nil)

        result, err := mockClient.ListInvalidations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListKeyGroups", func(t *testing.T) {
        input := &cloudfront.ListKeyGroupsInput{}
        output := &cloudfront.ListKeyGroupsOutput{}

        mockClient.On("ListKeyGroups", ctx, input).Return(output, nil)

        result, err := mockClient.ListKeyGroups(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListKeyValueStores", func(t *testing.T) {
        input := &cloudfront.ListKeyValueStoresInput{}
        output := &cloudfront.ListKeyValueStoresOutput{}

        mockClient.On("ListKeyValueStores", ctx, input).Return(output, nil)

        result, err := mockClient.ListKeyValueStores(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListOriginAccessControls", func(t *testing.T) {
        input := &cloudfront.ListOriginAccessControlsInput{}
        output := &cloudfront.ListOriginAccessControlsOutput{}

        mockClient.On("ListOriginAccessControls", ctx, input).Return(output, nil)

        result, err := mockClient.ListOriginAccessControls(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListOriginRequestPolicies", func(t *testing.T) {
        input := &cloudfront.ListOriginRequestPoliciesInput{}
        output := &cloudfront.ListOriginRequestPoliciesOutput{}

        mockClient.On("ListOriginRequestPolicies", ctx, input).Return(output, nil)

        result, err := mockClient.ListOriginRequestPolicies(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListPublicKeys", func(t *testing.T) {
        input := &cloudfront.ListPublicKeysInput{}
        output := &cloudfront.ListPublicKeysOutput{}

        mockClient.On("ListPublicKeys", ctx, input).Return(output, nil)

        result, err := mockClient.ListPublicKeys(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListRealtimeLogConfigs", func(t *testing.T) {
        input := &cloudfront.ListRealtimeLogConfigsInput{}
        output := &cloudfront.ListRealtimeLogConfigsOutput{}

        mockClient.On("ListRealtimeLogConfigs", ctx, input).Return(output, nil)

        result, err := mockClient.ListRealtimeLogConfigs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListResponseHeadersPolicies", func(t *testing.T) {
        input := &cloudfront.ListResponseHeadersPoliciesInput{}
        output := &cloudfront.ListResponseHeadersPoliciesOutput{}

        mockClient.On("ListResponseHeadersPolicies", ctx, input).Return(output, nil)

        result, err := mockClient.ListResponseHeadersPolicies(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListStreamingDistributions", func(t *testing.T) {
        input := &cloudfront.ListStreamingDistributionsInput{}
        output := &cloudfront.ListStreamingDistributionsOutput{}

        mockClient.On("ListStreamingDistributions", ctx, input).Return(output, nil)

        result, err := mockClient.ListStreamingDistributions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &cloudfront.ListTagsForResourceInput{}
        output := &cloudfront.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPublishFunction", func(t *testing.T) {
        input := &cloudfront.PublishFunctionInput{}
        output := &cloudfront.PublishFunctionOutput{}

        mockClient.On("PublishFunction", ctx, input).Return(output, nil)

        result, err := mockClient.PublishFunction(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &cloudfront.TagResourceInput{}
        output := &cloudfront.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTestFunction", func(t *testing.T) {
        input := &cloudfront.TestFunctionInput{}
        output := &cloudfront.TestFunctionOutput{}

        mockClient.On("TestFunction", ctx, input).Return(output, nil)

        result, err := mockClient.TestFunction(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &cloudfront.UntagResourceInput{}
        output := &cloudfront.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateCachePolicy", func(t *testing.T) {
        input := &cloudfront.UpdateCachePolicyInput{}
        output := &cloudfront.UpdateCachePolicyOutput{}

        mockClient.On("UpdateCachePolicy", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateCachePolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateCloudFrontOriginAccessIdentity", func(t *testing.T) {
        input := &cloudfront.UpdateCloudFrontOriginAccessIdentityInput{}
        output := &cloudfront.UpdateCloudFrontOriginAccessIdentityOutput{}

        mockClient.On("UpdateCloudFrontOriginAccessIdentity", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateCloudFrontOriginAccessIdentity(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateContinuousDeploymentPolicy", func(t *testing.T) {
        input := &cloudfront.UpdateContinuousDeploymentPolicyInput{}
        output := &cloudfront.UpdateContinuousDeploymentPolicyOutput{}

        mockClient.On("UpdateContinuousDeploymentPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateContinuousDeploymentPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateDistribution", func(t *testing.T) {
        input := &cloudfront.UpdateDistributionInput{}
        output := &cloudfront.UpdateDistributionOutput{}

        mockClient.On("UpdateDistribution", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateDistribution(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateDistributionWithStagingConfig", func(t *testing.T) {
        input := &cloudfront.UpdateDistributionWithStagingConfigInput{}
        output := &cloudfront.UpdateDistributionWithStagingConfigOutput{}

        mockClient.On("UpdateDistributionWithStagingConfig", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateDistributionWithStagingConfig(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateFieldLevelEncryptionConfig", func(t *testing.T) {
        input := &cloudfront.UpdateFieldLevelEncryptionConfigInput{}
        output := &cloudfront.UpdateFieldLevelEncryptionConfigOutput{}

        mockClient.On("UpdateFieldLevelEncryptionConfig", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateFieldLevelEncryptionConfig(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateFieldLevelEncryptionProfile", func(t *testing.T) {
        input := &cloudfront.UpdateFieldLevelEncryptionProfileInput{}
        output := &cloudfront.UpdateFieldLevelEncryptionProfileOutput{}

        mockClient.On("UpdateFieldLevelEncryptionProfile", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateFieldLevelEncryptionProfile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateFunction", func(t *testing.T) {
        input := &cloudfront.UpdateFunctionInput{}
        output := &cloudfront.UpdateFunctionOutput{}

        mockClient.On("UpdateFunction", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateFunction(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateKeyGroup", func(t *testing.T) {
        input := &cloudfront.UpdateKeyGroupInput{}
        output := &cloudfront.UpdateKeyGroupOutput{}

        mockClient.On("UpdateKeyGroup", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateKeyGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateKeyValueStore", func(t *testing.T) {
        input := &cloudfront.UpdateKeyValueStoreInput{}
        output := &cloudfront.UpdateKeyValueStoreOutput{}

        mockClient.On("UpdateKeyValueStore", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateKeyValueStore(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateOriginAccessControl", func(t *testing.T) {
        input := &cloudfront.UpdateOriginAccessControlInput{}
        output := &cloudfront.UpdateOriginAccessControlOutput{}

        mockClient.On("UpdateOriginAccessControl", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateOriginAccessControl(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateOriginRequestPolicy", func(t *testing.T) {
        input := &cloudfront.UpdateOriginRequestPolicyInput{}
        output := &cloudfront.UpdateOriginRequestPolicyOutput{}

        mockClient.On("UpdateOriginRequestPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateOriginRequestPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdatePublicKey", func(t *testing.T) {
        input := &cloudfront.UpdatePublicKeyInput{}
        output := &cloudfront.UpdatePublicKeyOutput{}

        mockClient.On("UpdatePublicKey", ctx, input).Return(output, nil)

        result, err := mockClient.UpdatePublicKey(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateRealtimeLogConfig", func(t *testing.T) {
        input := &cloudfront.UpdateRealtimeLogConfigInput{}
        output := &cloudfront.UpdateRealtimeLogConfigOutput{}

        mockClient.On("UpdateRealtimeLogConfig", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateRealtimeLogConfig(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateResponseHeadersPolicy", func(t *testing.T) {
        input := &cloudfront.UpdateResponseHeadersPolicyInput{}
        output := &cloudfront.UpdateResponseHeadersPolicyOutput{}

        mockClient.On("UpdateResponseHeadersPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateResponseHeadersPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateStreamingDistribution", func(t *testing.T) {
        input := &cloudfront.UpdateStreamingDistributionInput{}
        output := &cloudfront.UpdateStreamingDistributionOutput{}

        mockClient.On("UpdateStreamingDistribution", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateStreamingDistribution(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
