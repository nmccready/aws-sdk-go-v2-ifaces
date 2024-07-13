package applicationdiscoveryservice_test

// tests for the applicationdiscoveryservice service interface for this ../../../service/applicationdiscoveryservice/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/applicationdiscoveryservice"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/applicationdiscoveryservice/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/applicationdiscoveryservice/applicationdiscoveryservice_iface"
	"github.com/stretchr/testify/assert"
)

func TestApplicationdiscoveryserviceServiceCanBeMocked(t *testing.T) {
	var iface applicationdiscoveryservice_iface.IClient
	iface = &applicationdiscoveryservice.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := applicationdiscoveryservice.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateConfigurationItemsToApplication", func(t *testing.T) {
        input := &applicationdiscoveryservice.AssociateConfigurationItemsToApplicationInput{}
        output := &applicationdiscoveryservice.AssociateConfigurationItemsToApplicationOutput{}

        mockClient.On("AssociateConfigurationItemsToApplication", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateConfigurationItemsToApplication(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchDeleteAgents", func(t *testing.T) {
        input := &applicationdiscoveryservice.BatchDeleteAgentsInput{}
        output := &applicationdiscoveryservice.BatchDeleteAgentsOutput{}

        mockClient.On("BatchDeleteAgents", ctx, input).Return(output, nil)

        result, err := mockClient.BatchDeleteAgents(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchDeleteImportData", func(t *testing.T) {
        input := &applicationdiscoveryservice.BatchDeleteImportDataInput{}
        output := &applicationdiscoveryservice.BatchDeleteImportDataOutput{}

        mockClient.On("BatchDeleteImportData", ctx, input).Return(output, nil)

        result, err := mockClient.BatchDeleteImportData(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateApplication", func(t *testing.T) {
        input := &applicationdiscoveryservice.CreateApplicationInput{}
        output := &applicationdiscoveryservice.CreateApplicationOutput{}

        mockClient.On("CreateApplication", ctx, input).Return(output, nil)

        result, err := mockClient.CreateApplication(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateTags", func(t *testing.T) {
        input := &applicationdiscoveryservice.CreateTagsInput{}
        output := &applicationdiscoveryservice.CreateTagsOutput{}

        mockClient.On("CreateTags", ctx, input).Return(output, nil)

        result, err := mockClient.CreateTags(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteApplications", func(t *testing.T) {
        input := &applicationdiscoveryservice.DeleteApplicationsInput{}
        output := &applicationdiscoveryservice.DeleteApplicationsOutput{}

        mockClient.On("DeleteApplications", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteApplications(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteTags", func(t *testing.T) {
        input := &applicationdiscoveryservice.DeleteTagsInput{}
        output := &applicationdiscoveryservice.DeleteTagsOutput{}

        mockClient.On("DeleteTags", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteTags(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeAgents", func(t *testing.T) {
        input := &applicationdiscoveryservice.DescribeAgentsInput{}
        output := &applicationdiscoveryservice.DescribeAgentsOutput{}

        mockClient.On("DescribeAgents", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeAgents(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeBatchDeleteConfigurationTask", func(t *testing.T) {
        input := &applicationdiscoveryservice.DescribeBatchDeleteConfigurationTaskInput{}
        output := &applicationdiscoveryservice.DescribeBatchDeleteConfigurationTaskOutput{}

        mockClient.On("DescribeBatchDeleteConfigurationTask", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeBatchDeleteConfigurationTask(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeConfigurations", func(t *testing.T) {
        input := &applicationdiscoveryservice.DescribeConfigurationsInput{}
        output := &applicationdiscoveryservice.DescribeConfigurationsOutput{}

        mockClient.On("DescribeConfigurations", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeConfigurations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeContinuousExports", func(t *testing.T) {
        input := &applicationdiscoveryservice.DescribeContinuousExportsInput{}
        output := &applicationdiscoveryservice.DescribeContinuousExportsOutput{}

        mockClient.On("DescribeContinuousExports", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeContinuousExports(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeExportConfigurations", func(t *testing.T) {
        input := &applicationdiscoveryservice.DescribeExportConfigurationsInput{}
        output := &applicationdiscoveryservice.DescribeExportConfigurationsOutput{}

        mockClient.On("DescribeExportConfigurations", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeExportConfigurations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeExportTasks", func(t *testing.T) {
        input := &applicationdiscoveryservice.DescribeExportTasksInput{}
        output := &applicationdiscoveryservice.DescribeExportTasksOutput{}

        mockClient.On("DescribeExportTasks", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeExportTasks(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeImportTasks", func(t *testing.T) {
        input := &applicationdiscoveryservice.DescribeImportTasksInput{}
        output := &applicationdiscoveryservice.DescribeImportTasksOutput{}

        mockClient.On("DescribeImportTasks", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeImportTasks(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeTags", func(t *testing.T) {
        input := &applicationdiscoveryservice.DescribeTagsInput{}
        output := &applicationdiscoveryservice.DescribeTagsOutput{}

        mockClient.On("DescribeTags", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeTags(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateConfigurationItemsFromApplication", func(t *testing.T) {
        input := &applicationdiscoveryservice.DisassociateConfigurationItemsFromApplicationInput{}
        output := &applicationdiscoveryservice.DisassociateConfigurationItemsFromApplicationOutput{}

        mockClient.On("DisassociateConfigurationItemsFromApplication", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateConfigurationItemsFromApplication(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestExportConfigurations", func(t *testing.T) {
        input := &applicationdiscoveryservice.ExportConfigurationsInput{}
        output := &applicationdiscoveryservice.ExportConfigurationsOutput{}

        mockClient.On("ExportConfigurations", ctx, input).Return(output, nil)

        result, err := mockClient.ExportConfigurations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDiscoverySummary", func(t *testing.T) {
        input := &applicationdiscoveryservice.GetDiscoverySummaryInput{}
        output := &applicationdiscoveryservice.GetDiscoverySummaryOutput{}

        mockClient.On("GetDiscoverySummary", ctx, input).Return(output, nil)

        result, err := mockClient.GetDiscoverySummary(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListConfigurations", func(t *testing.T) {
        input := &applicationdiscoveryservice.ListConfigurationsInput{}
        output := &applicationdiscoveryservice.ListConfigurationsOutput{}

        mockClient.On("ListConfigurations", ctx, input).Return(output, nil)

        result, err := mockClient.ListConfigurations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListServerNeighbors", func(t *testing.T) {
        input := &applicationdiscoveryservice.ListServerNeighborsInput{}
        output := &applicationdiscoveryservice.ListServerNeighborsOutput{}

        mockClient.On("ListServerNeighbors", ctx, input).Return(output, nil)

        result, err := mockClient.ListServerNeighbors(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartBatchDeleteConfigurationTask", func(t *testing.T) {
        input := &applicationdiscoveryservice.StartBatchDeleteConfigurationTaskInput{}
        output := &applicationdiscoveryservice.StartBatchDeleteConfigurationTaskOutput{}

        mockClient.On("StartBatchDeleteConfigurationTask", ctx, input).Return(output, nil)

        result, err := mockClient.StartBatchDeleteConfigurationTask(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartContinuousExport", func(t *testing.T) {
        input := &applicationdiscoveryservice.StartContinuousExportInput{}
        output := &applicationdiscoveryservice.StartContinuousExportOutput{}

        mockClient.On("StartContinuousExport", ctx, input).Return(output, nil)

        result, err := mockClient.StartContinuousExport(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartDataCollectionByAgentIds", func(t *testing.T) {
        input := &applicationdiscoveryservice.StartDataCollectionByAgentIdsInput{}
        output := &applicationdiscoveryservice.StartDataCollectionByAgentIdsOutput{}

        mockClient.On("StartDataCollectionByAgentIds", ctx, input).Return(output, nil)

        result, err := mockClient.StartDataCollectionByAgentIds(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartExportTask", func(t *testing.T) {
        input := &applicationdiscoveryservice.StartExportTaskInput{}
        output := &applicationdiscoveryservice.StartExportTaskOutput{}

        mockClient.On("StartExportTask", ctx, input).Return(output, nil)

        result, err := mockClient.StartExportTask(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartImportTask", func(t *testing.T) {
        input := &applicationdiscoveryservice.StartImportTaskInput{}
        output := &applicationdiscoveryservice.StartImportTaskOutput{}

        mockClient.On("StartImportTask", ctx, input).Return(output, nil)

        result, err := mockClient.StartImportTask(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopContinuousExport", func(t *testing.T) {
        input := &applicationdiscoveryservice.StopContinuousExportInput{}
        output := &applicationdiscoveryservice.StopContinuousExportOutput{}

        mockClient.On("StopContinuousExport", ctx, input).Return(output, nil)

        result, err := mockClient.StopContinuousExport(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopDataCollectionByAgentIds", func(t *testing.T) {
        input := &applicationdiscoveryservice.StopDataCollectionByAgentIdsInput{}
        output := &applicationdiscoveryservice.StopDataCollectionByAgentIdsOutput{}

        mockClient.On("StopDataCollectionByAgentIds", ctx, input).Return(output, nil)

        result, err := mockClient.StopDataCollectionByAgentIds(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateApplication", func(t *testing.T) {
        input := &applicationdiscoveryservice.UpdateApplicationInput{}
        output := &applicationdiscoveryservice.UpdateApplicationOutput{}

        mockClient.On("UpdateApplication", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateApplication(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
