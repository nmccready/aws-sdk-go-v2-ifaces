package kinesisanalyticsv2_test

// tests for the kinesisanalyticsv2 service interface for this ../../../service/kinesisanalyticsv2/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/kinesisanalyticsv2"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/kinesisanalyticsv2/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/kinesisanalyticsv2/kinesisanalyticsv2_iface"
	"github.com/stretchr/testify/assert"
)

func TestKinesisanalyticsv2ServiceCanBeMocked(t *testing.T) {
	var iface kinesisanalyticsv2_iface.IClient
	iface = &kinesisanalyticsv2.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := kinesisanalyticsv2.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAddApplicationCloudWatchLoggingOption", func(t *testing.T) {
        input := &kinesisanalyticsv2.AddApplicationCloudWatchLoggingOptionInput{}
        output := &kinesisanalyticsv2.AddApplicationCloudWatchLoggingOptionOutput{}

        mockClient.On("AddApplicationCloudWatchLoggingOption", ctx, input).Return(output, nil)

        result, err := mockClient.AddApplicationCloudWatchLoggingOption(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAddApplicationInput", func(t *testing.T) {
        input := &kinesisanalyticsv2.AddApplicationInputInput{}
        output := &kinesisanalyticsv2.AddApplicationInputOutput{}

        mockClient.On("AddApplicationInput", ctx, input).Return(output, nil)

        result, err := mockClient.AddApplicationInput(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAddApplicationInputProcessingConfiguration", func(t *testing.T) {
        input := &kinesisanalyticsv2.AddApplicationInputProcessingConfigurationInput{}
        output := &kinesisanalyticsv2.AddApplicationInputProcessingConfigurationOutput{}

        mockClient.On("AddApplicationInputProcessingConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.AddApplicationInputProcessingConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAddApplicationOutput", func(t *testing.T) {
        input := &kinesisanalyticsv2.AddApplicationOutputInput{}
        output := &kinesisanalyticsv2.AddApplicationOutputOutput{}

        mockClient.On("AddApplicationOutput", ctx, input).Return(output, nil)

        result, err := mockClient.AddApplicationOutput(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAddApplicationReferenceDataSource", func(t *testing.T) {
        input := &kinesisanalyticsv2.AddApplicationReferenceDataSourceInput{}
        output := &kinesisanalyticsv2.AddApplicationReferenceDataSourceOutput{}

        mockClient.On("AddApplicationReferenceDataSource", ctx, input).Return(output, nil)

        result, err := mockClient.AddApplicationReferenceDataSource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAddApplicationVpcConfiguration", func(t *testing.T) {
        input := &kinesisanalyticsv2.AddApplicationVpcConfigurationInput{}
        output := &kinesisanalyticsv2.AddApplicationVpcConfigurationOutput{}

        mockClient.On("AddApplicationVpcConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.AddApplicationVpcConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateApplication", func(t *testing.T) {
        input := &kinesisanalyticsv2.CreateApplicationInput{}
        output := &kinesisanalyticsv2.CreateApplicationOutput{}

        mockClient.On("CreateApplication", ctx, input).Return(output, nil)

        result, err := mockClient.CreateApplication(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateApplicationPresignedUrl", func(t *testing.T) {
        input := &kinesisanalyticsv2.CreateApplicationPresignedUrlInput{}
        output := &kinesisanalyticsv2.CreateApplicationPresignedUrlOutput{}

        mockClient.On("CreateApplicationPresignedUrl", ctx, input).Return(output, nil)

        result, err := mockClient.CreateApplicationPresignedUrl(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateApplicationSnapshot", func(t *testing.T) {
        input := &kinesisanalyticsv2.CreateApplicationSnapshotInput{}
        output := &kinesisanalyticsv2.CreateApplicationSnapshotOutput{}

        mockClient.On("CreateApplicationSnapshot", ctx, input).Return(output, nil)

        result, err := mockClient.CreateApplicationSnapshot(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteApplication", func(t *testing.T) {
        input := &kinesisanalyticsv2.DeleteApplicationInput{}
        output := &kinesisanalyticsv2.DeleteApplicationOutput{}

        mockClient.On("DeleteApplication", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteApplication(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteApplicationCloudWatchLoggingOption", func(t *testing.T) {
        input := &kinesisanalyticsv2.DeleteApplicationCloudWatchLoggingOptionInput{}
        output := &kinesisanalyticsv2.DeleteApplicationCloudWatchLoggingOptionOutput{}

        mockClient.On("DeleteApplicationCloudWatchLoggingOption", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteApplicationCloudWatchLoggingOption(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteApplicationInputProcessingConfiguration", func(t *testing.T) {
        input := &kinesisanalyticsv2.DeleteApplicationInputProcessingConfigurationInput{}
        output := &kinesisanalyticsv2.DeleteApplicationInputProcessingConfigurationOutput{}

        mockClient.On("DeleteApplicationInputProcessingConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteApplicationInputProcessingConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteApplicationOutput", func(t *testing.T) {
        input := &kinesisanalyticsv2.DeleteApplicationOutputInput{}
        output := &kinesisanalyticsv2.DeleteApplicationOutputOutput{}

        mockClient.On("DeleteApplicationOutput", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteApplicationOutput(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteApplicationReferenceDataSource", func(t *testing.T) {
        input := &kinesisanalyticsv2.DeleteApplicationReferenceDataSourceInput{}
        output := &kinesisanalyticsv2.DeleteApplicationReferenceDataSourceOutput{}

        mockClient.On("DeleteApplicationReferenceDataSource", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteApplicationReferenceDataSource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteApplicationSnapshot", func(t *testing.T) {
        input := &kinesisanalyticsv2.DeleteApplicationSnapshotInput{}
        output := &kinesisanalyticsv2.DeleteApplicationSnapshotOutput{}

        mockClient.On("DeleteApplicationSnapshot", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteApplicationSnapshot(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteApplicationVpcConfiguration", func(t *testing.T) {
        input := &kinesisanalyticsv2.DeleteApplicationVpcConfigurationInput{}
        output := &kinesisanalyticsv2.DeleteApplicationVpcConfigurationOutput{}

        mockClient.On("DeleteApplicationVpcConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteApplicationVpcConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeApplication", func(t *testing.T) {
        input := &kinesisanalyticsv2.DescribeApplicationInput{}
        output := &kinesisanalyticsv2.DescribeApplicationOutput{}

        mockClient.On("DescribeApplication", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeApplication(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeApplicationOperation", func(t *testing.T) {
        input := &kinesisanalyticsv2.DescribeApplicationOperationInput{}
        output := &kinesisanalyticsv2.DescribeApplicationOperationOutput{}

        mockClient.On("DescribeApplicationOperation", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeApplicationOperation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeApplicationSnapshot", func(t *testing.T) {
        input := &kinesisanalyticsv2.DescribeApplicationSnapshotInput{}
        output := &kinesisanalyticsv2.DescribeApplicationSnapshotOutput{}

        mockClient.On("DescribeApplicationSnapshot", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeApplicationSnapshot(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeApplicationVersion", func(t *testing.T) {
        input := &kinesisanalyticsv2.DescribeApplicationVersionInput{}
        output := &kinesisanalyticsv2.DescribeApplicationVersionOutput{}

        mockClient.On("DescribeApplicationVersion", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeApplicationVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDiscoverInputSchema", func(t *testing.T) {
        input := &kinesisanalyticsv2.DiscoverInputSchemaInput{}
        output := &kinesisanalyticsv2.DiscoverInputSchemaOutput{}

        mockClient.On("DiscoverInputSchema", ctx, input).Return(output, nil)

        result, err := mockClient.DiscoverInputSchema(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListApplicationOperations", func(t *testing.T) {
        input := &kinesisanalyticsv2.ListApplicationOperationsInput{}
        output := &kinesisanalyticsv2.ListApplicationOperationsOutput{}

        mockClient.On("ListApplicationOperations", ctx, input).Return(output, nil)

        result, err := mockClient.ListApplicationOperations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListApplicationSnapshots", func(t *testing.T) {
        input := &kinesisanalyticsv2.ListApplicationSnapshotsInput{}
        output := &kinesisanalyticsv2.ListApplicationSnapshotsOutput{}

        mockClient.On("ListApplicationSnapshots", ctx, input).Return(output, nil)

        result, err := mockClient.ListApplicationSnapshots(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListApplicationVersions", func(t *testing.T) {
        input := &kinesisanalyticsv2.ListApplicationVersionsInput{}
        output := &kinesisanalyticsv2.ListApplicationVersionsOutput{}

        mockClient.On("ListApplicationVersions", ctx, input).Return(output, nil)

        result, err := mockClient.ListApplicationVersions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListApplications", func(t *testing.T) {
        input := &kinesisanalyticsv2.ListApplicationsInput{}
        output := &kinesisanalyticsv2.ListApplicationsOutput{}

        mockClient.On("ListApplications", ctx, input).Return(output, nil)

        result, err := mockClient.ListApplications(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &kinesisanalyticsv2.ListTagsForResourceInput{}
        output := &kinesisanalyticsv2.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRollbackApplication", func(t *testing.T) {
        input := &kinesisanalyticsv2.RollbackApplicationInput{}
        output := &kinesisanalyticsv2.RollbackApplicationOutput{}

        mockClient.On("RollbackApplication", ctx, input).Return(output, nil)

        result, err := mockClient.RollbackApplication(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartApplication", func(t *testing.T) {
        input := &kinesisanalyticsv2.StartApplicationInput{}
        output := &kinesisanalyticsv2.StartApplicationOutput{}

        mockClient.On("StartApplication", ctx, input).Return(output, nil)

        result, err := mockClient.StartApplication(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopApplication", func(t *testing.T) {
        input := &kinesisanalyticsv2.StopApplicationInput{}
        output := &kinesisanalyticsv2.StopApplicationOutput{}

        mockClient.On("StopApplication", ctx, input).Return(output, nil)

        result, err := mockClient.StopApplication(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &kinesisanalyticsv2.TagResourceInput{}
        output := &kinesisanalyticsv2.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &kinesisanalyticsv2.UntagResourceInput{}
        output := &kinesisanalyticsv2.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateApplication", func(t *testing.T) {
        input := &kinesisanalyticsv2.UpdateApplicationInput{}
        output := &kinesisanalyticsv2.UpdateApplicationOutput{}

        mockClient.On("UpdateApplication", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateApplication(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateApplicationMaintenanceConfiguration", func(t *testing.T) {
        input := &kinesisanalyticsv2.UpdateApplicationMaintenanceConfigurationInput{}
        output := &kinesisanalyticsv2.UpdateApplicationMaintenanceConfigurationOutput{}

        mockClient.On("UpdateApplicationMaintenanceConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateApplicationMaintenanceConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
