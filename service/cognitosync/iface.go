
package cognitosync

import (
    "github.com/aws/aws-sdk-go-v2/service/cognitosync"
)

// ICognitosync defines the interface for cognitosync
type ICognitosync interface {
 Options() Options 
 BulkPublish(ctx context.Context, params *BulkPublishInput, optFns ...func(*Options)) (*BulkPublishOutput, error) 
 DeleteDataset(ctx context.Context, params *DeleteDatasetInput, optFns ...func(*Options)) (*DeleteDatasetOutput, error) 
 DescribeDataset(ctx context.Context, params *DescribeDatasetInput, optFns ...func(*Options)) (*DescribeDatasetOutput, error) 
 DescribeIdentityPoolUsage(ctx context.Context, params *DescribeIdentityPoolUsageInput, optFns ...func(*Options)) (*DescribeIdentityPoolUsageOutput, error) 
 DescribeIdentityUsage(ctx context.Context, params *DescribeIdentityUsageInput, optFns ...func(*Options)) (*DescribeIdentityUsageOutput, error) 
 GetBulkPublishDetails(ctx context.Context, params *GetBulkPublishDetailsInput, optFns ...func(*Options)) (*GetBulkPublishDetailsOutput, error) 
 GetCognitoEvents(ctx context.Context, params *GetCognitoEventsInput, optFns ...func(*Options)) (*GetCognitoEventsOutput, error) 
 GetIdentityPoolConfiguration(ctx context.Context, params *GetIdentityPoolConfigurationInput, optFns ...func(*Options)) (*GetIdentityPoolConfigurationOutput, error) 
 ListDatasets(ctx context.Context, params *ListDatasetsInput, optFns ...func(*Options)) (*ListDatasetsOutput, error) 
 ListIdentityPoolUsage(ctx context.Context, params *ListIdentityPoolUsageInput, optFns ...func(*Options)) (*ListIdentityPoolUsageOutput, error) 
 ListRecords(ctx context.Context, params *ListRecordsInput, optFns ...func(*Options)) (*ListRecordsOutput, error) 
 RegisterDevice(ctx context.Context, params *RegisterDeviceInput, optFns ...func(*Options)) (*RegisterDeviceOutput, error) 
 SetCognitoEvents(ctx context.Context, params *SetCognitoEventsInput, optFns ...func(*Options)) (*SetCognitoEventsOutput, error) 
 SetIdentityPoolConfiguration(ctx context.Context, params *SetIdentityPoolConfigurationInput, optFns ...func(*Options)) (*SetIdentityPoolConfigurationOutput, error) 
 SubscribeToDataset(ctx context.Context, params *SubscribeToDatasetInput, optFns ...func(*Options)) (*SubscribeToDatasetOutput, error) 
 UnsubscribeFromDataset(ctx context.Context, params *UnsubscribeFromDatasetInput, optFns ...func(*Options)) (*UnsubscribeFromDatasetOutput, error) 
 UpdateRecords(ctx context.Context, params *UpdateRecordsInput, optFns ...func(*Options)) (*UpdateRecordsOutput, error) 
}
