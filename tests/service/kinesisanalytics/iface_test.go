package kinesisanalytics_test

// tests for the kinesisanalytics service interface for this ../../../service/kinesisanalytics/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/kinesisanalytics"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/kinesisanalytics/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/kinesisanalytics/kinesisanalytics_iface"
	"github.com/stretchr/testify/assert"
)

func TestKinesisanalyticsServiceCanBeMocked(t *testing.T) {
	var iface kinesisanalytics_iface.IClient
	iface = &kinesisanalytics.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := kinesisanalytics.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAddApplicationCloudWatchLoggingOption", func(t *testing.T) {
        input := &kinesisanalytics.AddApplicationCloudWatchLoggingOptionInput{}
        output := &kinesisanalytics.AddApplicationCloudWatchLoggingOptionOutput{}

        mockClient.On("AddApplicationCloudWatchLoggingOption", ctx, input).Return(output, nil)

        result, err := mockClient.AddApplicationCloudWatchLoggingOption(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAddApplicationInput", func(t *testing.T) {
        input := &kinesisanalytics.AddApplicationInputInput{}
        output := &kinesisanalytics.AddApplicationInputOutput{}

        mockClient.On("AddApplicationInput", ctx, input).Return(output, nil)

        result, err := mockClient.AddApplicationInput(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAddApplicationInputProcessingConfiguration", func(t *testing.T) {
        input := &kinesisanalytics.AddApplicationInputProcessingConfigurationInput{}
        output := &kinesisanalytics.AddApplicationInputProcessingConfigurationOutput{}

        mockClient.On("AddApplicationInputProcessingConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.AddApplicationInputProcessingConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAddApplicationOutput", func(t *testing.T) {
        input := &kinesisanalytics.AddApplicationOutputInput{}
        output := &kinesisanalytics.AddApplicationOutputOutput{}

        mockClient.On("AddApplicationOutput", ctx, input).Return(output, nil)

        result, err := mockClient.AddApplicationOutput(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAddApplicationReferenceDataSource", func(t *testing.T) {
        input := &kinesisanalytics.AddApplicationReferenceDataSourceInput{}
        output := &kinesisanalytics.AddApplicationReferenceDataSourceOutput{}

        mockClient.On("AddApplicationReferenceDataSource", ctx, input).Return(output, nil)

        result, err := mockClient.AddApplicationReferenceDataSource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateApplication", func(t *testing.T) {
        input := &kinesisanalytics.CreateApplicationInput{}
        output := &kinesisanalytics.CreateApplicationOutput{}

        mockClient.On("CreateApplication", ctx, input).Return(output, nil)

        result, err := mockClient.CreateApplication(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteApplication", func(t *testing.T) {
        input := &kinesisanalytics.DeleteApplicationInput{}
        output := &kinesisanalytics.DeleteApplicationOutput{}

        mockClient.On("DeleteApplication", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteApplication(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteApplicationCloudWatchLoggingOption", func(t *testing.T) {
        input := &kinesisanalytics.DeleteApplicationCloudWatchLoggingOptionInput{}
        output := &kinesisanalytics.DeleteApplicationCloudWatchLoggingOptionOutput{}

        mockClient.On("DeleteApplicationCloudWatchLoggingOption", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteApplicationCloudWatchLoggingOption(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteApplicationInputProcessingConfiguration", func(t *testing.T) {
        input := &kinesisanalytics.DeleteApplicationInputProcessingConfigurationInput{}
        output := &kinesisanalytics.DeleteApplicationInputProcessingConfigurationOutput{}

        mockClient.On("DeleteApplicationInputProcessingConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteApplicationInputProcessingConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteApplicationOutput", func(t *testing.T) {
        input := &kinesisanalytics.DeleteApplicationOutputInput{}
        output := &kinesisanalytics.DeleteApplicationOutputOutput{}

        mockClient.On("DeleteApplicationOutput", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteApplicationOutput(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteApplicationReferenceDataSource", func(t *testing.T) {
        input := &kinesisanalytics.DeleteApplicationReferenceDataSourceInput{}
        output := &kinesisanalytics.DeleteApplicationReferenceDataSourceOutput{}

        mockClient.On("DeleteApplicationReferenceDataSource", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteApplicationReferenceDataSource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeApplication", func(t *testing.T) {
        input := &kinesisanalytics.DescribeApplicationInput{}
        output := &kinesisanalytics.DescribeApplicationOutput{}

        mockClient.On("DescribeApplication", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeApplication(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDiscoverInputSchema", func(t *testing.T) {
        input := &kinesisanalytics.DiscoverInputSchemaInput{}
        output := &kinesisanalytics.DiscoverInputSchemaOutput{}

        mockClient.On("DiscoverInputSchema", ctx, input).Return(output, nil)

        result, err := mockClient.DiscoverInputSchema(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListApplications", func(t *testing.T) {
        input := &kinesisanalytics.ListApplicationsInput{}
        output := &kinesisanalytics.ListApplicationsOutput{}

        mockClient.On("ListApplications", ctx, input).Return(output, nil)

        result, err := mockClient.ListApplications(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &kinesisanalytics.ListTagsForResourceInput{}
        output := &kinesisanalytics.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartApplication", func(t *testing.T) {
        input := &kinesisanalytics.StartApplicationInput{}
        output := &kinesisanalytics.StartApplicationOutput{}

        mockClient.On("StartApplication", ctx, input).Return(output, nil)

        result, err := mockClient.StartApplication(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopApplication", func(t *testing.T) {
        input := &kinesisanalytics.StopApplicationInput{}
        output := &kinesisanalytics.StopApplicationOutput{}

        mockClient.On("StopApplication", ctx, input).Return(output, nil)

        result, err := mockClient.StopApplication(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &kinesisanalytics.TagResourceInput{}
        output := &kinesisanalytics.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &kinesisanalytics.UntagResourceInput{}
        output := &kinesisanalytics.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateApplication", func(t *testing.T) {
        input := &kinesisanalytics.UpdateApplicationInput{}
        output := &kinesisanalytics.UpdateApplicationOutput{}

        mockClient.On("UpdateApplication", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateApplication(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
