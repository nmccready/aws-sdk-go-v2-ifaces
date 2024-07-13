package iotsitewise_test

// tests for the iotsitewise service interface for this ../../../service/iotsitewise/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/iotsitewise"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/iotsitewise/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/iotsitewise/iotsitewise_iface"
	"github.com/stretchr/testify/assert"
)

func TestIotsitewiseServiceCanBeMocked(t *testing.T) {
	var iface iotsitewise_iface.IClient
	iface = &iotsitewise.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := iotsitewise.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateAssets", func(t *testing.T) {
        input := &iotsitewise.AssociateAssetsInput{}
        output := &iotsitewise.AssociateAssetsOutput{}

        mockClient.On("AssociateAssets", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateAssets(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateTimeSeriesToAssetProperty", func(t *testing.T) {
        input := &iotsitewise.AssociateTimeSeriesToAssetPropertyInput{}
        output := &iotsitewise.AssociateTimeSeriesToAssetPropertyOutput{}

        mockClient.On("AssociateTimeSeriesToAssetProperty", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateTimeSeriesToAssetProperty(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchAssociateProjectAssets", func(t *testing.T) {
        input := &iotsitewise.BatchAssociateProjectAssetsInput{}
        output := &iotsitewise.BatchAssociateProjectAssetsOutput{}

        mockClient.On("BatchAssociateProjectAssets", ctx, input).Return(output, nil)

        result, err := mockClient.BatchAssociateProjectAssets(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchDisassociateProjectAssets", func(t *testing.T) {
        input := &iotsitewise.BatchDisassociateProjectAssetsInput{}
        output := &iotsitewise.BatchDisassociateProjectAssetsOutput{}

        mockClient.On("BatchDisassociateProjectAssets", ctx, input).Return(output, nil)

        result, err := mockClient.BatchDisassociateProjectAssets(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchGetAssetPropertyAggregates", func(t *testing.T) {
        input := &iotsitewise.BatchGetAssetPropertyAggregatesInput{}
        output := &iotsitewise.BatchGetAssetPropertyAggregatesOutput{}

        mockClient.On("BatchGetAssetPropertyAggregates", ctx, input).Return(output, nil)

        result, err := mockClient.BatchGetAssetPropertyAggregates(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchGetAssetPropertyValue", func(t *testing.T) {
        input := &iotsitewise.BatchGetAssetPropertyValueInput{}
        output := &iotsitewise.BatchGetAssetPropertyValueOutput{}

        mockClient.On("BatchGetAssetPropertyValue", ctx, input).Return(output, nil)

        result, err := mockClient.BatchGetAssetPropertyValue(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchGetAssetPropertyValueHistory", func(t *testing.T) {
        input := &iotsitewise.BatchGetAssetPropertyValueHistoryInput{}
        output := &iotsitewise.BatchGetAssetPropertyValueHistoryOutput{}

        mockClient.On("BatchGetAssetPropertyValueHistory", ctx, input).Return(output, nil)

        result, err := mockClient.BatchGetAssetPropertyValueHistory(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchPutAssetPropertyValue", func(t *testing.T) {
        input := &iotsitewise.BatchPutAssetPropertyValueInput{}
        output := &iotsitewise.BatchPutAssetPropertyValueOutput{}

        mockClient.On("BatchPutAssetPropertyValue", ctx, input).Return(output, nil)

        result, err := mockClient.BatchPutAssetPropertyValue(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateAccessPolicy", func(t *testing.T) {
        input := &iotsitewise.CreateAccessPolicyInput{}
        output := &iotsitewise.CreateAccessPolicyOutput{}

        mockClient.On("CreateAccessPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.CreateAccessPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateAsset", func(t *testing.T) {
        input := &iotsitewise.CreateAssetInput{}
        output := &iotsitewise.CreateAssetOutput{}

        mockClient.On("CreateAsset", ctx, input).Return(output, nil)

        result, err := mockClient.CreateAsset(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateAssetModel", func(t *testing.T) {
        input := &iotsitewise.CreateAssetModelInput{}
        output := &iotsitewise.CreateAssetModelOutput{}

        mockClient.On("CreateAssetModel", ctx, input).Return(output, nil)

        result, err := mockClient.CreateAssetModel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateAssetModelCompositeModel", func(t *testing.T) {
        input := &iotsitewise.CreateAssetModelCompositeModelInput{}
        output := &iotsitewise.CreateAssetModelCompositeModelOutput{}

        mockClient.On("CreateAssetModelCompositeModel", ctx, input).Return(output, nil)

        result, err := mockClient.CreateAssetModelCompositeModel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateBulkImportJob", func(t *testing.T) {
        input := &iotsitewise.CreateBulkImportJobInput{}
        output := &iotsitewise.CreateBulkImportJobOutput{}

        mockClient.On("CreateBulkImportJob", ctx, input).Return(output, nil)

        result, err := mockClient.CreateBulkImportJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateDashboard", func(t *testing.T) {
        input := &iotsitewise.CreateDashboardInput{}
        output := &iotsitewise.CreateDashboardOutput{}

        mockClient.On("CreateDashboard", ctx, input).Return(output, nil)

        result, err := mockClient.CreateDashboard(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateGateway", func(t *testing.T) {
        input := &iotsitewise.CreateGatewayInput{}
        output := &iotsitewise.CreateGatewayOutput{}

        mockClient.On("CreateGateway", ctx, input).Return(output, nil)

        result, err := mockClient.CreateGateway(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreatePortal", func(t *testing.T) {
        input := &iotsitewise.CreatePortalInput{}
        output := &iotsitewise.CreatePortalOutput{}

        mockClient.On("CreatePortal", ctx, input).Return(output, nil)

        result, err := mockClient.CreatePortal(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateProject", func(t *testing.T) {
        input := &iotsitewise.CreateProjectInput{}
        output := &iotsitewise.CreateProjectOutput{}

        mockClient.On("CreateProject", ctx, input).Return(output, nil)

        result, err := mockClient.CreateProject(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteAccessPolicy", func(t *testing.T) {
        input := &iotsitewise.DeleteAccessPolicyInput{}
        output := &iotsitewise.DeleteAccessPolicyOutput{}

        mockClient.On("DeleteAccessPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteAccessPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteAsset", func(t *testing.T) {
        input := &iotsitewise.DeleteAssetInput{}
        output := &iotsitewise.DeleteAssetOutput{}

        mockClient.On("DeleteAsset", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteAsset(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteAssetModel", func(t *testing.T) {
        input := &iotsitewise.DeleteAssetModelInput{}
        output := &iotsitewise.DeleteAssetModelOutput{}

        mockClient.On("DeleteAssetModel", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteAssetModel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteAssetModelCompositeModel", func(t *testing.T) {
        input := &iotsitewise.DeleteAssetModelCompositeModelInput{}
        output := &iotsitewise.DeleteAssetModelCompositeModelOutput{}

        mockClient.On("DeleteAssetModelCompositeModel", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteAssetModelCompositeModel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteDashboard", func(t *testing.T) {
        input := &iotsitewise.DeleteDashboardInput{}
        output := &iotsitewise.DeleteDashboardOutput{}

        mockClient.On("DeleteDashboard", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteDashboard(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteGateway", func(t *testing.T) {
        input := &iotsitewise.DeleteGatewayInput{}
        output := &iotsitewise.DeleteGatewayOutput{}

        mockClient.On("DeleteGateway", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteGateway(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeletePortal", func(t *testing.T) {
        input := &iotsitewise.DeletePortalInput{}
        output := &iotsitewise.DeletePortalOutput{}

        mockClient.On("DeletePortal", ctx, input).Return(output, nil)

        result, err := mockClient.DeletePortal(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteProject", func(t *testing.T) {
        input := &iotsitewise.DeleteProjectInput{}
        output := &iotsitewise.DeleteProjectOutput{}

        mockClient.On("DeleteProject", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteProject(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteTimeSeries", func(t *testing.T) {
        input := &iotsitewise.DeleteTimeSeriesInput{}
        output := &iotsitewise.DeleteTimeSeriesOutput{}

        mockClient.On("DeleteTimeSeries", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteTimeSeries(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeAccessPolicy", func(t *testing.T) {
        input := &iotsitewise.DescribeAccessPolicyInput{}
        output := &iotsitewise.DescribeAccessPolicyOutput{}

        mockClient.On("DescribeAccessPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeAccessPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeAction", func(t *testing.T) {
        input := &iotsitewise.DescribeActionInput{}
        output := &iotsitewise.DescribeActionOutput{}

        mockClient.On("DescribeAction", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeAction(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeAsset", func(t *testing.T) {
        input := &iotsitewise.DescribeAssetInput{}
        output := &iotsitewise.DescribeAssetOutput{}

        mockClient.On("DescribeAsset", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeAsset(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeAssetCompositeModel", func(t *testing.T) {
        input := &iotsitewise.DescribeAssetCompositeModelInput{}
        output := &iotsitewise.DescribeAssetCompositeModelOutput{}

        mockClient.On("DescribeAssetCompositeModel", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeAssetCompositeModel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeAssetModel", func(t *testing.T) {
        input := &iotsitewise.DescribeAssetModelInput{}
        output := &iotsitewise.DescribeAssetModelOutput{}

        mockClient.On("DescribeAssetModel", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeAssetModel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeAssetModelCompositeModel", func(t *testing.T) {
        input := &iotsitewise.DescribeAssetModelCompositeModelInput{}
        output := &iotsitewise.DescribeAssetModelCompositeModelOutput{}

        mockClient.On("DescribeAssetModelCompositeModel", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeAssetModelCompositeModel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeAssetProperty", func(t *testing.T) {
        input := &iotsitewise.DescribeAssetPropertyInput{}
        output := &iotsitewise.DescribeAssetPropertyOutput{}

        mockClient.On("DescribeAssetProperty", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeAssetProperty(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeBulkImportJob", func(t *testing.T) {
        input := &iotsitewise.DescribeBulkImportJobInput{}
        output := &iotsitewise.DescribeBulkImportJobOutput{}

        mockClient.On("DescribeBulkImportJob", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeBulkImportJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeDashboard", func(t *testing.T) {
        input := &iotsitewise.DescribeDashboardInput{}
        output := &iotsitewise.DescribeDashboardOutput{}

        mockClient.On("DescribeDashboard", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeDashboard(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeDefaultEncryptionConfiguration", func(t *testing.T) {
        input := &iotsitewise.DescribeDefaultEncryptionConfigurationInput{}
        output := &iotsitewise.DescribeDefaultEncryptionConfigurationOutput{}

        mockClient.On("DescribeDefaultEncryptionConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeDefaultEncryptionConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeGateway", func(t *testing.T) {
        input := &iotsitewise.DescribeGatewayInput{}
        output := &iotsitewise.DescribeGatewayOutput{}

        mockClient.On("DescribeGateway", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeGateway(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeGatewayCapabilityConfiguration", func(t *testing.T) {
        input := &iotsitewise.DescribeGatewayCapabilityConfigurationInput{}
        output := &iotsitewise.DescribeGatewayCapabilityConfigurationOutput{}

        mockClient.On("DescribeGatewayCapabilityConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeGatewayCapabilityConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeLoggingOptions", func(t *testing.T) {
        input := &iotsitewise.DescribeLoggingOptionsInput{}
        output := &iotsitewise.DescribeLoggingOptionsOutput{}

        mockClient.On("DescribeLoggingOptions", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeLoggingOptions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribePortal", func(t *testing.T) {
        input := &iotsitewise.DescribePortalInput{}
        output := &iotsitewise.DescribePortalOutput{}

        mockClient.On("DescribePortal", ctx, input).Return(output, nil)

        result, err := mockClient.DescribePortal(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeProject", func(t *testing.T) {
        input := &iotsitewise.DescribeProjectInput{}
        output := &iotsitewise.DescribeProjectOutput{}

        mockClient.On("DescribeProject", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeProject(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeStorageConfiguration", func(t *testing.T) {
        input := &iotsitewise.DescribeStorageConfigurationInput{}
        output := &iotsitewise.DescribeStorageConfigurationOutput{}

        mockClient.On("DescribeStorageConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeStorageConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeTimeSeries", func(t *testing.T) {
        input := &iotsitewise.DescribeTimeSeriesInput{}
        output := &iotsitewise.DescribeTimeSeriesOutput{}

        mockClient.On("DescribeTimeSeries", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeTimeSeries(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateAssets", func(t *testing.T) {
        input := &iotsitewise.DisassociateAssetsInput{}
        output := &iotsitewise.DisassociateAssetsOutput{}

        mockClient.On("DisassociateAssets", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateAssets(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateTimeSeriesFromAssetProperty", func(t *testing.T) {
        input := &iotsitewise.DisassociateTimeSeriesFromAssetPropertyInput{}
        output := &iotsitewise.DisassociateTimeSeriesFromAssetPropertyOutput{}

        mockClient.On("DisassociateTimeSeriesFromAssetProperty", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateTimeSeriesFromAssetProperty(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestExecuteAction", func(t *testing.T) {
        input := &iotsitewise.ExecuteActionInput{}
        output := &iotsitewise.ExecuteActionOutput{}

        mockClient.On("ExecuteAction", ctx, input).Return(output, nil)

        result, err := mockClient.ExecuteAction(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestExecuteQuery", func(t *testing.T) {
        input := &iotsitewise.ExecuteQueryInput{}
        output := &iotsitewise.ExecuteQueryOutput{}

        mockClient.On("ExecuteQuery", ctx, input).Return(output, nil)

        result, err := mockClient.ExecuteQuery(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetAssetPropertyAggregates", func(t *testing.T) {
        input := &iotsitewise.GetAssetPropertyAggregatesInput{}
        output := &iotsitewise.GetAssetPropertyAggregatesOutput{}

        mockClient.On("GetAssetPropertyAggregates", ctx, input).Return(output, nil)

        result, err := mockClient.GetAssetPropertyAggregates(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetAssetPropertyValue", func(t *testing.T) {
        input := &iotsitewise.GetAssetPropertyValueInput{}
        output := &iotsitewise.GetAssetPropertyValueOutput{}

        mockClient.On("GetAssetPropertyValue", ctx, input).Return(output, nil)

        result, err := mockClient.GetAssetPropertyValue(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetAssetPropertyValueHistory", func(t *testing.T) {
        input := &iotsitewise.GetAssetPropertyValueHistoryInput{}
        output := &iotsitewise.GetAssetPropertyValueHistoryOutput{}

        mockClient.On("GetAssetPropertyValueHistory", ctx, input).Return(output, nil)

        result, err := mockClient.GetAssetPropertyValueHistory(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetInterpolatedAssetPropertyValues", func(t *testing.T) {
        input := &iotsitewise.GetInterpolatedAssetPropertyValuesInput{}
        output := &iotsitewise.GetInterpolatedAssetPropertyValuesOutput{}

        mockClient.On("GetInterpolatedAssetPropertyValues", ctx, input).Return(output, nil)

        result, err := mockClient.GetInterpolatedAssetPropertyValues(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAccessPolicies", func(t *testing.T) {
        input := &iotsitewise.ListAccessPoliciesInput{}
        output := &iotsitewise.ListAccessPoliciesOutput{}

        mockClient.On("ListAccessPolicies", ctx, input).Return(output, nil)

        result, err := mockClient.ListAccessPolicies(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListActions", func(t *testing.T) {
        input := &iotsitewise.ListActionsInput{}
        output := &iotsitewise.ListActionsOutput{}

        mockClient.On("ListActions", ctx, input).Return(output, nil)

        result, err := mockClient.ListActions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAssetModelCompositeModels", func(t *testing.T) {
        input := &iotsitewise.ListAssetModelCompositeModelsInput{}
        output := &iotsitewise.ListAssetModelCompositeModelsOutput{}

        mockClient.On("ListAssetModelCompositeModels", ctx, input).Return(output, nil)

        result, err := mockClient.ListAssetModelCompositeModels(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAssetModelProperties", func(t *testing.T) {
        input := &iotsitewise.ListAssetModelPropertiesInput{}
        output := &iotsitewise.ListAssetModelPropertiesOutput{}

        mockClient.On("ListAssetModelProperties", ctx, input).Return(output, nil)

        result, err := mockClient.ListAssetModelProperties(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAssetModels", func(t *testing.T) {
        input := &iotsitewise.ListAssetModelsInput{}
        output := &iotsitewise.ListAssetModelsOutput{}

        mockClient.On("ListAssetModels", ctx, input).Return(output, nil)

        result, err := mockClient.ListAssetModels(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAssetProperties", func(t *testing.T) {
        input := &iotsitewise.ListAssetPropertiesInput{}
        output := &iotsitewise.ListAssetPropertiesOutput{}

        mockClient.On("ListAssetProperties", ctx, input).Return(output, nil)

        result, err := mockClient.ListAssetProperties(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAssetRelationships", func(t *testing.T) {
        input := &iotsitewise.ListAssetRelationshipsInput{}
        output := &iotsitewise.ListAssetRelationshipsOutput{}

        mockClient.On("ListAssetRelationships", ctx, input).Return(output, nil)

        result, err := mockClient.ListAssetRelationships(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAssets", func(t *testing.T) {
        input := &iotsitewise.ListAssetsInput{}
        output := &iotsitewise.ListAssetsOutput{}

        mockClient.On("ListAssets", ctx, input).Return(output, nil)

        result, err := mockClient.ListAssets(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAssociatedAssets", func(t *testing.T) {
        input := &iotsitewise.ListAssociatedAssetsInput{}
        output := &iotsitewise.ListAssociatedAssetsOutput{}

        mockClient.On("ListAssociatedAssets", ctx, input).Return(output, nil)

        result, err := mockClient.ListAssociatedAssets(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListBulkImportJobs", func(t *testing.T) {
        input := &iotsitewise.ListBulkImportJobsInput{}
        output := &iotsitewise.ListBulkImportJobsOutput{}

        mockClient.On("ListBulkImportJobs", ctx, input).Return(output, nil)

        result, err := mockClient.ListBulkImportJobs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListCompositionRelationships", func(t *testing.T) {
        input := &iotsitewise.ListCompositionRelationshipsInput{}
        output := &iotsitewise.ListCompositionRelationshipsOutput{}

        mockClient.On("ListCompositionRelationships", ctx, input).Return(output, nil)

        result, err := mockClient.ListCompositionRelationships(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDashboards", func(t *testing.T) {
        input := &iotsitewise.ListDashboardsInput{}
        output := &iotsitewise.ListDashboardsOutput{}

        mockClient.On("ListDashboards", ctx, input).Return(output, nil)

        result, err := mockClient.ListDashboards(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListGateways", func(t *testing.T) {
        input := &iotsitewise.ListGatewaysInput{}
        output := &iotsitewise.ListGatewaysOutput{}

        mockClient.On("ListGateways", ctx, input).Return(output, nil)

        result, err := mockClient.ListGateways(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListPortals", func(t *testing.T) {
        input := &iotsitewise.ListPortalsInput{}
        output := &iotsitewise.ListPortalsOutput{}

        mockClient.On("ListPortals", ctx, input).Return(output, nil)

        result, err := mockClient.ListPortals(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListProjectAssets", func(t *testing.T) {
        input := &iotsitewise.ListProjectAssetsInput{}
        output := &iotsitewise.ListProjectAssetsOutput{}

        mockClient.On("ListProjectAssets", ctx, input).Return(output, nil)

        result, err := mockClient.ListProjectAssets(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListProjects", func(t *testing.T) {
        input := &iotsitewise.ListProjectsInput{}
        output := &iotsitewise.ListProjectsOutput{}

        mockClient.On("ListProjects", ctx, input).Return(output, nil)

        result, err := mockClient.ListProjects(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &iotsitewise.ListTagsForResourceInput{}
        output := &iotsitewise.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTimeSeries", func(t *testing.T) {
        input := &iotsitewise.ListTimeSeriesInput{}
        output := &iotsitewise.ListTimeSeriesOutput{}

        mockClient.On("ListTimeSeries", ctx, input).Return(output, nil)

        result, err := mockClient.ListTimeSeries(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutDefaultEncryptionConfiguration", func(t *testing.T) {
        input := &iotsitewise.PutDefaultEncryptionConfigurationInput{}
        output := &iotsitewise.PutDefaultEncryptionConfigurationOutput{}

        mockClient.On("PutDefaultEncryptionConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.PutDefaultEncryptionConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutLoggingOptions", func(t *testing.T) {
        input := &iotsitewise.PutLoggingOptionsInput{}
        output := &iotsitewise.PutLoggingOptionsOutput{}

        mockClient.On("PutLoggingOptions", ctx, input).Return(output, nil)

        result, err := mockClient.PutLoggingOptions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutStorageConfiguration", func(t *testing.T) {
        input := &iotsitewise.PutStorageConfigurationInput{}
        output := &iotsitewise.PutStorageConfigurationOutput{}

        mockClient.On("PutStorageConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.PutStorageConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &iotsitewise.TagResourceInput{}
        output := &iotsitewise.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &iotsitewise.UntagResourceInput{}
        output := &iotsitewise.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateAccessPolicy", func(t *testing.T) {
        input := &iotsitewise.UpdateAccessPolicyInput{}
        output := &iotsitewise.UpdateAccessPolicyOutput{}

        mockClient.On("UpdateAccessPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateAccessPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateAsset", func(t *testing.T) {
        input := &iotsitewise.UpdateAssetInput{}
        output := &iotsitewise.UpdateAssetOutput{}

        mockClient.On("UpdateAsset", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateAsset(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateAssetModel", func(t *testing.T) {
        input := &iotsitewise.UpdateAssetModelInput{}
        output := &iotsitewise.UpdateAssetModelOutput{}

        mockClient.On("UpdateAssetModel", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateAssetModel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateAssetModelCompositeModel", func(t *testing.T) {
        input := &iotsitewise.UpdateAssetModelCompositeModelInput{}
        output := &iotsitewise.UpdateAssetModelCompositeModelOutput{}

        mockClient.On("UpdateAssetModelCompositeModel", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateAssetModelCompositeModel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateAssetProperty", func(t *testing.T) {
        input := &iotsitewise.UpdateAssetPropertyInput{}
        output := &iotsitewise.UpdateAssetPropertyOutput{}

        mockClient.On("UpdateAssetProperty", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateAssetProperty(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateDashboard", func(t *testing.T) {
        input := &iotsitewise.UpdateDashboardInput{}
        output := &iotsitewise.UpdateDashboardOutput{}

        mockClient.On("UpdateDashboard", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateDashboard(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateGateway", func(t *testing.T) {
        input := &iotsitewise.UpdateGatewayInput{}
        output := &iotsitewise.UpdateGatewayOutput{}

        mockClient.On("UpdateGateway", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateGateway(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateGatewayCapabilityConfiguration", func(t *testing.T) {
        input := &iotsitewise.UpdateGatewayCapabilityConfigurationInput{}
        output := &iotsitewise.UpdateGatewayCapabilityConfigurationOutput{}

        mockClient.On("UpdateGatewayCapabilityConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateGatewayCapabilityConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdatePortal", func(t *testing.T) {
        input := &iotsitewise.UpdatePortalInput{}
        output := &iotsitewise.UpdatePortalOutput{}

        mockClient.On("UpdatePortal", ctx, input).Return(output, nil)

        result, err := mockClient.UpdatePortal(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateProject", func(t *testing.T) {
        input := &iotsitewise.UpdateProjectInput{}
        output := &iotsitewise.UpdateProjectOutput{}

        mockClient.On("UpdateProject", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateProject(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
