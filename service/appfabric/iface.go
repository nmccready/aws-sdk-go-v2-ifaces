
package appfabric

import (
    "github.com/aws/aws-sdk-go-v2/service/appfabric"
)

// IAppfabric defines the interface for appfabric
type IAppfabric interface {
 Options() Options 
 BatchGetUserAccessTasks(ctx context.Context, params *BatchGetUserAccessTasksInput, optFns ...func(*Options)) (*BatchGetUserAccessTasksOutput, error) 
 ConnectAppAuthorization(ctx context.Context, params *ConnectAppAuthorizationInput, optFns ...func(*Options)) (*ConnectAppAuthorizationOutput, error) 
 CreateAppAuthorization(ctx context.Context, params *CreateAppAuthorizationInput, optFns ...func(*Options)) (*CreateAppAuthorizationOutput, error) 
 CreateAppBundle(ctx context.Context, params *CreateAppBundleInput, optFns ...func(*Options)) (*CreateAppBundleOutput, error) 
 CreateIngestion(ctx context.Context, params *CreateIngestionInput, optFns ...func(*Options)) (*CreateIngestionOutput, error) 
 CreateIngestionDestination(ctx context.Context, params *CreateIngestionDestinationInput, optFns ...func(*Options)) (*CreateIngestionDestinationOutput, error) 
 DeleteAppAuthorization(ctx context.Context, params *DeleteAppAuthorizationInput, optFns ...func(*Options)) (*DeleteAppAuthorizationOutput, error) 
 DeleteAppBundle(ctx context.Context, params *DeleteAppBundleInput, optFns ...func(*Options)) (*DeleteAppBundleOutput, error) 
 DeleteIngestion(ctx context.Context, params *DeleteIngestionInput, optFns ...func(*Options)) (*DeleteIngestionOutput, error) 
 DeleteIngestionDestination(ctx context.Context, params *DeleteIngestionDestinationInput, optFns ...func(*Options)) (*DeleteIngestionDestinationOutput, error) 
 GetAppAuthorization(ctx context.Context, params *GetAppAuthorizationInput, optFns ...func(*Options)) (*GetAppAuthorizationOutput, error) 
 GetAppBundle(ctx context.Context, params *GetAppBundleInput, optFns ...func(*Options)) (*GetAppBundleOutput, error) 
 GetIngestion(ctx context.Context, params *GetIngestionInput, optFns ...func(*Options)) (*GetIngestionOutput, error) 
 GetIngestionDestination(ctx context.Context, params *GetIngestionDestinationInput, optFns ...func(*Options)) (*GetIngestionDestinationOutput, error) 
 ListAppAuthorizations(ctx context.Context, params *ListAppAuthorizationsInput, optFns ...func(*Options)) (*ListAppAuthorizationsOutput, error) 
 ListAppBundles(ctx context.Context, params *ListAppBundlesInput, optFns ...func(*Options)) (*ListAppBundlesOutput, error) 
 ListIngestionDestinations(ctx context.Context, params *ListIngestionDestinationsInput, optFns ...func(*Options)) (*ListIngestionDestinationsOutput, error) 
 ListIngestions(ctx context.Context, params *ListIngestionsInput, optFns ...func(*Options)) (*ListIngestionsOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 StartIngestion(ctx context.Context, params *StartIngestionInput, optFns ...func(*Options)) (*StartIngestionOutput, error) 
 StartUserAccessTasks(ctx context.Context, params *StartUserAccessTasksInput, optFns ...func(*Options)) (*StartUserAccessTasksOutput, error) 
 StopIngestion(ctx context.Context, params *StopIngestionInput, optFns ...func(*Options)) (*StopIngestionOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateAppAuthorization(ctx context.Context, params *UpdateAppAuthorizationInput, optFns ...func(*Options)) (*UpdateAppAuthorizationOutput, error) 
 UpdateIngestionDestination(ctx context.Context, params *UpdateIngestionDestinationInput, optFns ...func(*Options)) (*UpdateIngestionDestinationOutput, error) 
}
