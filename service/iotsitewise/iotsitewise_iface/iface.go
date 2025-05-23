
package iotsitewise_iface

// This file was generated by scripts/gen_ifaces.sh do not edit it manually
import (
    "context"
    . "github.com/aws/aws-sdk-go-v2/service/iotsitewise"
)

// IClient defines the interface for iotsitewise
type IClient interface {
 Options() Options 
 AssociateAssets(ctx context.Context, params *AssociateAssetsInput, optFns ...func(*Options)) (*AssociateAssetsOutput, error) 
 AssociateTimeSeriesToAssetProperty(ctx context.Context, params *AssociateTimeSeriesToAssetPropertyInput, optFns ...func(*Options)) (*AssociateTimeSeriesToAssetPropertyOutput, error) 
 BatchAssociateProjectAssets(ctx context.Context, params *BatchAssociateProjectAssetsInput, optFns ...func(*Options)) (*BatchAssociateProjectAssetsOutput, error) 
 BatchDisassociateProjectAssets(ctx context.Context, params *BatchDisassociateProjectAssetsInput, optFns ...func(*Options)) (*BatchDisassociateProjectAssetsOutput, error) 
 BatchGetAssetPropertyAggregates(ctx context.Context, params *BatchGetAssetPropertyAggregatesInput, optFns ...func(*Options)) (*BatchGetAssetPropertyAggregatesOutput, error) 
 BatchGetAssetPropertyValue(ctx context.Context, params *BatchGetAssetPropertyValueInput, optFns ...func(*Options)) (*BatchGetAssetPropertyValueOutput, error) 
 BatchGetAssetPropertyValueHistory(ctx context.Context, params *BatchGetAssetPropertyValueHistoryInput, optFns ...func(*Options)) (*BatchGetAssetPropertyValueHistoryOutput, error) 
 BatchPutAssetPropertyValue(ctx context.Context, params *BatchPutAssetPropertyValueInput, optFns ...func(*Options)) (*BatchPutAssetPropertyValueOutput, error) 
 CreateAccessPolicy(ctx context.Context, params *CreateAccessPolicyInput, optFns ...func(*Options)) (*CreateAccessPolicyOutput, error) 
 CreateAsset(ctx context.Context, params *CreateAssetInput, optFns ...func(*Options)) (*CreateAssetOutput, error) 
 CreateAssetModel(ctx context.Context, params *CreateAssetModelInput, optFns ...func(*Options)) (*CreateAssetModelOutput, error) 
 CreateAssetModelCompositeModel(ctx context.Context, params *CreateAssetModelCompositeModelInput, optFns ...func(*Options)) (*CreateAssetModelCompositeModelOutput, error) 
 CreateBulkImportJob(ctx context.Context, params *CreateBulkImportJobInput, optFns ...func(*Options)) (*CreateBulkImportJobOutput, error) 
 CreateDashboard(ctx context.Context, params *CreateDashboardInput, optFns ...func(*Options)) (*CreateDashboardOutput, error) 
 CreateDataset(ctx context.Context, params *CreateDatasetInput, optFns ...func(*Options)) (*CreateDatasetOutput, error) 
 CreateGateway(ctx context.Context, params *CreateGatewayInput, optFns ...func(*Options)) (*CreateGatewayOutput, error) 
 CreatePortal(ctx context.Context, params *CreatePortalInput, optFns ...func(*Options)) (*CreatePortalOutput, error) 
 CreateProject(ctx context.Context, params *CreateProjectInput, optFns ...func(*Options)) (*CreateProjectOutput, error) 
 DeleteAccessPolicy(ctx context.Context, params *DeleteAccessPolicyInput, optFns ...func(*Options)) (*DeleteAccessPolicyOutput, error) 
 DeleteAsset(ctx context.Context, params *DeleteAssetInput, optFns ...func(*Options)) (*DeleteAssetOutput, error) 
 DeleteAssetModel(ctx context.Context, params *DeleteAssetModelInput, optFns ...func(*Options)) (*DeleteAssetModelOutput, error) 
 DeleteAssetModelCompositeModel(ctx context.Context, params *DeleteAssetModelCompositeModelInput, optFns ...func(*Options)) (*DeleteAssetModelCompositeModelOutput, error) 
 DeleteDashboard(ctx context.Context, params *DeleteDashboardInput, optFns ...func(*Options)) (*DeleteDashboardOutput, error) 
 DeleteDataset(ctx context.Context, params *DeleteDatasetInput, optFns ...func(*Options)) (*DeleteDatasetOutput, error) 
 DeleteGateway(ctx context.Context, params *DeleteGatewayInput, optFns ...func(*Options)) (*DeleteGatewayOutput, error) 
 DeletePortal(ctx context.Context, params *DeletePortalInput, optFns ...func(*Options)) (*DeletePortalOutput, error) 
 DeleteProject(ctx context.Context, params *DeleteProjectInput, optFns ...func(*Options)) (*DeleteProjectOutput, error) 
 DeleteTimeSeries(ctx context.Context, params *DeleteTimeSeriesInput, optFns ...func(*Options)) (*DeleteTimeSeriesOutput, error) 
 DescribeAccessPolicy(ctx context.Context, params *DescribeAccessPolicyInput, optFns ...func(*Options)) (*DescribeAccessPolicyOutput, error) 
 DescribeAction(ctx context.Context, params *DescribeActionInput, optFns ...func(*Options)) (*DescribeActionOutput, error) 
 DescribeAsset(ctx context.Context, params *DescribeAssetInput, optFns ...func(*Options)) (*DescribeAssetOutput, error) 
 DescribeAssetCompositeModel(ctx context.Context, params *DescribeAssetCompositeModelInput, optFns ...func(*Options)) (*DescribeAssetCompositeModelOutput, error) 
 DescribeAssetModel(ctx context.Context, params *DescribeAssetModelInput, optFns ...func(*Options)) (*DescribeAssetModelOutput, error) 
 DescribeAssetModelCompositeModel(ctx context.Context, params *DescribeAssetModelCompositeModelInput, optFns ...func(*Options)) (*DescribeAssetModelCompositeModelOutput, error) 
 DescribeAssetProperty(ctx context.Context, params *DescribeAssetPropertyInput, optFns ...func(*Options)) (*DescribeAssetPropertyOutput, error) 
 DescribeBulkImportJob(ctx context.Context, params *DescribeBulkImportJobInput, optFns ...func(*Options)) (*DescribeBulkImportJobOutput, error) 
 DescribeDashboard(ctx context.Context, params *DescribeDashboardInput, optFns ...func(*Options)) (*DescribeDashboardOutput, error) 
 DescribeDataset(ctx context.Context, params *DescribeDatasetInput, optFns ...func(*Options)) (*DescribeDatasetOutput, error) 
 DescribeDefaultEncryptionConfiguration(ctx context.Context, params *DescribeDefaultEncryptionConfigurationInput, optFns ...func(*Options)) (*DescribeDefaultEncryptionConfigurationOutput, error) 
 DescribeGateway(ctx context.Context, params *DescribeGatewayInput, optFns ...func(*Options)) (*DescribeGatewayOutput, error) 
 DescribeGatewayCapabilityConfiguration(ctx context.Context, params *DescribeGatewayCapabilityConfigurationInput, optFns ...func(*Options)) (*DescribeGatewayCapabilityConfigurationOutput, error) 
 DescribeLoggingOptions(ctx context.Context, params *DescribeLoggingOptionsInput, optFns ...func(*Options)) (*DescribeLoggingOptionsOutput, error) 
 DescribePortal(ctx context.Context, params *DescribePortalInput, optFns ...func(*Options)) (*DescribePortalOutput, error) 
 DescribeProject(ctx context.Context, params *DescribeProjectInput, optFns ...func(*Options)) (*DescribeProjectOutput, error) 
 DescribeStorageConfiguration(ctx context.Context, params *DescribeStorageConfigurationInput, optFns ...func(*Options)) (*DescribeStorageConfigurationOutput, error) 
 DescribeTimeSeries(ctx context.Context, params *DescribeTimeSeriesInput, optFns ...func(*Options)) (*DescribeTimeSeriesOutput, error) 
 DisassociateAssets(ctx context.Context, params *DisassociateAssetsInput, optFns ...func(*Options)) (*DisassociateAssetsOutput, error) 
 DisassociateTimeSeriesFromAssetProperty(ctx context.Context, params *DisassociateTimeSeriesFromAssetPropertyInput, optFns ...func(*Options)) (*DisassociateTimeSeriesFromAssetPropertyOutput, error) 
 ExecuteAction(ctx context.Context, params *ExecuteActionInput, optFns ...func(*Options)) (*ExecuteActionOutput, error) 
 ExecuteQuery(ctx context.Context, params *ExecuteQueryInput, optFns ...func(*Options)) (*ExecuteQueryOutput, error) 
 GetAssetPropertyAggregates(ctx context.Context, params *GetAssetPropertyAggregatesInput, optFns ...func(*Options)) (*GetAssetPropertyAggregatesOutput, error) 
 GetAssetPropertyValue(ctx context.Context, params *GetAssetPropertyValueInput, optFns ...func(*Options)) (*GetAssetPropertyValueOutput, error) 
 GetAssetPropertyValueHistory(ctx context.Context, params *GetAssetPropertyValueHistoryInput, optFns ...func(*Options)) (*GetAssetPropertyValueHistoryOutput, error) 
 GetInterpolatedAssetPropertyValues(ctx context.Context, params *GetInterpolatedAssetPropertyValuesInput, optFns ...func(*Options)) (*GetInterpolatedAssetPropertyValuesOutput, error) 
 InvokeAssistant(ctx context.Context, params *InvokeAssistantInput, optFns ...func(*Options)) (*InvokeAssistantOutput, error) 
 ListAccessPolicies(ctx context.Context, params *ListAccessPoliciesInput, optFns ...func(*Options)) (*ListAccessPoliciesOutput, error) 
 ListActions(ctx context.Context, params *ListActionsInput, optFns ...func(*Options)) (*ListActionsOutput, error) 
 ListAssetModelCompositeModels(ctx context.Context, params *ListAssetModelCompositeModelsInput, optFns ...func(*Options)) (*ListAssetModelCompositeModelsOutput, error) 
 ListAssetModelProperties(ctx context.Context, params *ListAssetModelPropertiesInput, optFns ...func(*Options)) (*ListAssetModelPropertiesOutput, error) 
 ListAssetModels(ctx context.Context, params *ListAssetModelsInput, optFns ...func(*Options)) (*ListAssetModelsOutput, error) 
 ListAssetProperties(ctx context.Context, params *ListAssetPropertiesInput, optFns ...func(*Options)) (*ListAssetPropertiesOutput, error) 
 ListAssetRelationships(ctx context.Context, params *ListAssetRelationshipsInput, optFns ...func(*Options)) (*ListAssetRelationshipsOutput, error) 
 ListAssets(ctx context.Context, params *ListAssetsInput, optFns ...func(*Options)) (*ListAssetsOutput, error) 
 ListAssociatedAssets(ctx context.Context, params *ListAssociatedAssetsInput, optFns ...func(*Options)) (*ListAssociatedAssetsOutput, error) 
 ListBulkImportJobs(ctx context.Context, params *ListBulkImportJobsInput, optFns ...func(*Options)) (*ListBulkImportJobsOutput, error) 
 ListCompositionRelationships(ctx context.Context, params *ListCompositionRelationshipsInput, optFns ...func(*Options)) (*ListCompositionRelationshipsOutput, error) 
 ListDashboards(ctx context.Context, params *ListDashboardsInput, optFns ...func(*Options)) (*ListDashboardsOutput, error) 
 ListDatasets(ctx context.Context, params *ListDatasetsInput, optFns ...func(*Options)) (*ListDatasetsOutput, error) 
 ListGateways(ctx context.Context, params *ListGatewaysInput, optFns ...func(*Options)) (*ListGatewaysOutput, error) 
 ListPortals(ctx context.Context, params *ListPortalsInput, optFns ...func(*Options)) (*ListPortalsOutput, error) 
 ListProjectAssets(ctx context.Context, params *ListProjectAssetsInput, optFns ...func(*Options)) (*ListProjectAssetsOutput, error) 
 ListProjects(ctx context.Context, params *ListProjectsInput, optFns ...func(*Options)) (*ListProjectsOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 ListTimeSeries(ctx context.Context, params *ListTimeSeriesInput, optFns ...func(*Options)) (*ListTimeSeriesOutput, error) 
 PutDefaultEncryptionConfiguration(ctx context.Context, params *PutDefaultEncryptionConfigurationInput, optFns ...func(*Options)) (*PutDefaultEncryptionConfigurationOutput, error) 
 PutLoggingOptions(ctx context.Context, params *PutLoggingOptionsInput, optFns ...func(*Options)) (*PutLoggingOptionsOutput, error) 
 PutStorageConfiguration(ctx context.Context, params *PutStorageConfigurationInput, optFns ...func(*Options)) (*PutStorageConfigurationOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateAccessPolicy(ctx context.Context, params *UpdateAccessPolicyInput, optFns ...func(*Options)) (*UpdateAccessPolicyOutput, error) 
 UpdateAsset(ctx context.Context, params *UpdateAssetInput, optFns ...func(*Options)) (*UpdateAssetOutput, error) 
 UpdateAssetModel(ctx context.Context, params *UpdateAssetModelInput, optFns ...func(*Options)) (*UpdateAssetModelOutput, error) 
 UpdateAssetModelCompositeModel(ctx context.Context, params *UpdateAssetModelCompositeModelInput, optFns ...func(*Options)) (*UpdateAssetModelCompositeModelOutput, error) 
 UpdateAssetProperty(ctx context.Context, params *UpdateAssetPropertyInput, optFns ...func(*Options)) (*UpdateAssetPropertyOutput, error) 
 UpdateDashboard(ctx context.Context, params *UpdateDashboardInput, optFns ...func(*Options)) (*UpdateDashboardOutput, error) 
 UpdateDataset(ctx context.Context, params *UpdateDatasetInput, optFns ...func(*Options)) (*UpdateDatasetOutput, error) 
 UpdateGateway(ctx context.Context, params *UpdateGatewayInput, optFns ...func(*Options)) (*UpdateGatewayOutput, error) 
 UpdateGatewayCapabilityConfiguration(ctx context.Context, params *UpdateGatewayCapabilityConfigurationInput, optFns ...func(*Options)) (*UpdateGatewayCapabilityConfigurationOutput, error) 
 UpdatePortal(ctx context.Context, params *UpdatePortalInput, optFns ...func(*Options)) (*UpdatePortalOutput, error) 
 UpdateProject(ctx context.Context, params *UpdateProjectInput, optFns ...func(*Options)) (*UpdateProjectOutput, error) 
}
