package mediapackagevod_test

// tests for the mediapackagevod service interface for this ../../../service/mediapackagevod/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/mediapackagevod"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/mediapackagevod/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/mediapackagevod/mediapackagevod_iface"
	"github.com/stretchr/testify/assert"
)

func TestMediapackagevodServiceCanBeMocked(t *testing.T) {
	var iface mediapackagevod_iface.IClient
	iface = &mediapackagevod.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := mediapackagevod.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestConfigureLogs", func(t *testing.T) {
        input := &mediapackagevod.ConfigureLogsInput{}
        output := &mediapackagevod.ConfigureLogsOutput{}

        mockClient.On("ConfigureLogs", ctx, input).Return(output, nil)

        result, err := mockClient.ConfigureLogs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateAsset", func(t *testing.T) {
        input := &mediapackagevod.CreateAssetInput{}
        output := &mediapackagevod.CreateAssetOutput{}

        mockClient.On("CreateAsset", ctx, input).Return(output, nil)

        result, err := mockClient.CreateAsset(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreatePackagingConfiguration", func(t *testing.T) {
        input := &mediapackagevod.CreatePackagingConfigurationInput{}
        output := &mediapackagevod.CreatePackagingConfigurationOutput{}

        mockClient.On("CreatePackagingConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.CreatePackagingConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreatePackagingGroup", func(t *testing.T) {
        input := &mediapackagevod.CreatePackagingGroupInput{}
        output := &mediapackagevod.CreatePackagingGroupOutput{}

        mockClient.On("CreatePackagingGroup", ctx, input).Return(output, nil)

        result, err := mockClient.CreatePackagingGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteAsset", func(t *testing.T) {
        input := &mediapackagevod.DeleteAssetInput{}
        output := &mediapackagevod.DeleteAssetOutput{}

        mockClient.On("DeleteAsset", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteAsset(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeletePackagingConfiguration", func(t *testing.T) {
        input := &mediapackagevod.DeletePackagingConfigurationInput{}
        output := &mediapackagevod.DeletePackagingConfigurationOutput{}

        mockClient.On("DeletePackagingConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.DeletePackagingConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeletePackagingGroup", func(t *testing.T) {
        input := &mediapackagevod.DeletePackagingGroupInput{}
        output := &mediapackagevod.DeletePackagingGroupOutput{}

        mockClient.On("DeletePackagingGroup", ctx, input).Return(output, nil)

        result, err := mockClient.DeletePackagingGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeAsset", func(t *testing.T) {
        input := &mediapackagevod.DescribeAssetInput{}
        output := &mediapackagevod.DescribeAssetOutput{}

        mockClient.On("DescribeAsset", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeAsset(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribePackagingConfiguration", func(t *testing.T) {
        input := &mediapackagevod.DescribePackagingConfigurationInput{}
        output := &mediapackagevod.DescribePackagingConfigurationOutput{}

        mockClient.On("DescribePackagingConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.DescribePackagingConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribePackagingGroup", func(t *testing.T) {
        input := &mediapackagevod.DescribePackagingGroupInput{}
        output := &mediapackagevod.DescribePackagingGroupOutput{}

        mockClient.On("DescribePackagingGroup", ctx, input).Return(output, nil)

        result, err := mockClient.DescribePackagingGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAssets", func(t *testing.T) {
        input := &mediapackagevod.ListAssetsInput{}
        output := &mediapackagevod.ListAssetsOutput{}

        mockClient.On("ListAssets", ctx, input).Return(output, nil)

        result, err := mockClient.ListAssets(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListPackagingConfigurations", func(t *testing.T) {
        input := &mediapackagevod.ListPackagingConfigurationsInput{}
        output := &mediapackagevod.ListPackagingConfigurationsOutput{}

        mockClient.On("ListPackagingConfigurations", ctx, input).Return(output, nil)

        result, err := mockClient.ListPackagingConfigurations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListPackagingGroups", func(t *testing.T) {
        input := &mediapackagevod.ListPackagingGroupsInput{}
        output := &mediapackagevod.ListPackagingGroupsOutput{}

        mockClient.On("ListPackagingGroups", ctx, input).Return(output, nil)

        result, err := mockClient.ListPackagingGroups(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &mediapackagevod.ListTagsForResourceInput{}
        output := &mediapackagevod.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &mediapackagevod.TagResourceInput{}
        output := &mediapackagevod.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &mediapackagevod.UntagResourceInput{}
        output := &mediapackagevod.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdatePackagingGroup", func(t *testing.T) {
        input := &mediapackagevod.UpdatePackagingGroupInput{}
        output := &mediapackagevod.UpdatePackagingGroupOutput{}

        mockClient.On("UpdatePackagingGroup", ctx, input).Return(output, nil)

        result, err := mockClient.UpdatePackagingGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
