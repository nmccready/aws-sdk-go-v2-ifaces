
package applicationdiscoveryservice

import (
    "github.com/aws/aws-sdk-go-v2/service/applicationdiscoveryservice"
)

// IApplicationdiscoveryservice defines the interface for applicationdiscoveryservice
type IApplicationdiscoveryservice interface {
 Options() Options 
 AssociateConfigurationItemsToApplication(ctx context.Context, params *AssociateConfigurationItemsToApplicationInput, optFns ...func(*Options)) (*AssociateConfigurationItemsToApplicationOutput, error) 
 BatchDeleteAgents(ctx context.Context, params *BatchDeleteAgentsInput, optFns ...func(*Options)) (*BatchDeleteAgentsOutput, error) 
 BatchDeleteImportData(ctx context.Context, params *BatchDeleteImportDataInput, optFns ...func(*Options)) (*BatchDeleteImportDataOutput, error) 
 CreateApplication(ctx context.Context, params *CreateApplicationInput, optFns ...func(*Options)) (*CreateApplicationOutput, error) 
 CreateTags(ctx context.Context, params *CreateTagsInput, optFns ...func(*Options)) (*CreateTagsOutput, error) 
 DeleteApplications(ctx context.Context, params *DeleteApplicationsInput, optFns ...func(*Options)) (*DeleteApplicationsOutput, error) 
 DeleteTags(ctx context.Context, params *DeleteTagsInput, optFns ...func(*Options)) (*DeleteTagsOutput, error) 
 DescribeAgents(ctx context.Context, params *DescribeAgentsInput, optFns ...func(*Options)) (*DescribeAgentsOutput, error) 
 DescribeBatchDeleteConfigurationTask(ctx context.Context, params *DescribeBatchDeleteConfigurationTaskInput, optFns ...func(*Options)) (*DescribeBatchDeleteConfigurationTaskOutput, error) 
 DescribeConfigurations(ctx context.Context, params *DescribeConfigurationsInput, optFns ...func(*Options)) (*DescribeConfigurationsOutput, error) 
 DescribeContinuousExports(ctx context.Context, params *DescribeContinuousExportsInput, optFns ...func(*Options)) (*DescribeContinuousExportsOutput, error) 
 DescribeExportConfigurations(ctx context.Context, params *DescribeExportConfigurationsInput, optFns ...func(*Options)) (*DescribeExportConfigurationsOutput, error) 
 DescribeExportTasks(ctx context.Context, params *DescribeExportTasksInput, optFns ...func(*Options)) (*DescribeExportTasksOutput, error) 
 DescribeImportTasks(ctx context.Context, params *DescribeImportTasksInput, optFns ...func(*Options)) (*DescribeImportTasksOutput, error) 
 DescribeTags(ctx context.Context, params *DescribeTagsInput, optFns ...func(*Options)) (*DescribeTagsOutput, error) 
 DisassociateConfigurationItemsFromApplication(ctx context.Context, params *DisassociateConfigurationItemsFromApplicationInput, optFns ...func(*Options)) (*DisassociateConfigurationItemsFromApplicationOutput, error) 
 ExportConfigurations(ctx context.Context, params *ExportConfigurationsInput, optFns ...func(*Options)) (*ExportConfigurationsOutput, error) 
 GetDiscoverySummary(ctx context.Context, params *GetDiscoverySummaryInput, optFns ...func(*Options)) (*GetDiscoverySummaryOutput, error) 
 ListConfigurations(ctx context.Context, params *ListConfigurationsInput, optFns ...func(*Options)) (*ListConfigurationsOutput, error) 
 ListServerNeighbors(ctx context.Context, params *ListServerNeighborsInput, optFns ...func(*Options)) (*ListServerNeighborsOutput, error) 
 StartBatchDeleteConfigurationTask(ctx context.Context, params *StartBatchDeleteConfigurationTaskInput, optFns ...func(*Options)) (*StartBatchDeleteConfigurationTaskOutput, error) 
 StartContinuousExport(ctx context.Context, params *StartContinuousExportInput, optFns ...func(*Options)) (*StartContinuousExportOutput, error) 
 StartDataCollectionByAgentIds(ctx context.Context, params *StartDataCollectionByAgentIdsInput, optFns ...func(*Options)) (*StartDataCollectionByAgentIdsOutput, error) 
 StartExportTask(ctx context.Context, params *StartExportTaskInput, optFns ...func(*Options)) (*StartExportTaskOutput, error) 
 StartImportTask(ctx context.Context, params *StartImportTaskInput, optFns ...func(*Options)) (*StartImportTaskOutput, error) 
 StopContinuousExport(ctx context.Context, params *StopContinuousExportInput, optFns ...func(*Options)) (*StopContinuousExportOutput, error) 
 StopDataCollectionByAgentIds(ctx context.Context, params *StopDataCollectionByAgentIdsInput, optFns ...func(*Options)) (*StopDataCollectionByAgentIdsOutput, error) 
 UpdateApplication(ctx context.Context, params *UpdateApplicationInput, optFns ...func(*Options)) (*UpdateApplicationOutput, error) 
}
