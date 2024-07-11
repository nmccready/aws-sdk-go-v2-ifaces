
package codeconnections

import (
    "github.com/aws/aws-sdk-go-v2/service/codeconnections"
)

// ICodeconnections defines the interface for codeconnections
type ICodeconnections interface {
 Options() Options 
 CreateConnection(ctx context.Context, params *CreateConnectionInput, optFns ...func(*Options)) (*CreateConnectionOutput, error) 
 CreateHost(ctx context.Context, params *CreateHostInput, optFns ...func(*Options)) (*CreateHostOutput, error) 
 CreateRepositoryLink(ctx context.Context, params *CreateRepositoryLinkInput, optFns ...func(*Options)) (*CreateRepositoryLinkOutput, error) 
 CreateSyncConfiguration(ctx context.Context, params *CreateSyncConfigurationInput, optFns ...func(*Options)) (*CreateSyncConfigurationOutput, error) 
 DeleteConnection(ctx context.Context, params *DeleteConnectionInput, optFns ...func(*Options)) (*DeleteConnectionOutput, error) 
 DeleteHost(ctx context.Context, params *DeleteHostInput, optFns ...func(*Options)) (*DeleteHostOutput, error) 
 DeleteRepositoryLink(ctx context.Context, params *DeleteRepositoryLinkInput, optFns ...func(*Options)) (*DeleteRepositoryLinkOutput, error) 
 DeleteSyncConfiguration(ctx context.Context, params *DeleteSyncConfigurationInput, optFns ...func(*Options)) (*DeleteSyncConfigurationOutput, error) 
 GetConnection(ctx context.Context, params *GetConnectionInput, optFns ...func(*Options)) (*GetConnectionOutput, error) 
 GetHost(ctx context.Context, params *GetHostInput, optFns ...func(*Options)) (*GetHostOutput, error) 
 GetRepositoryLink(ctx context.Context, params *GetRepositoryLinkInput, optFns ...func(*Options)) (*GetRepositoryLinkOutput, error) 
 GetRepositorySyncStatus(ctx context.Context, params *GetRepositorySyncStatusInput, optFns ...func(*Options)) (*GetRepositorySyncStatusOutput, error) 
 GetResourceSyncStatus(ctx context.Context, params *GetResourceSyncStatusInput, optFns ...func(*Options)) (*GetResourceSyncStatusOutput, error) 
 GetSyncBlockerSummary(ctx context.Context, params *GetSyncBlockerSummaryInput, optFns ...func(*Options)) (*GetSyncBlockerSummaryOutput, error) 
 GetSyncConfiguration(ctx context.Context, params *GetSyncConfigurationInput, optFns ...func(*Options)) (*GetSyncConfigurationOutput, error) 
 ListConnections(ctx context.Context, params *ListConnectionsInput, optFns ...func(*Options)) (*ListConnectionsOutput, error) 
 ListHosts(ctx context.Context, params *ListHostsInput, optFns ...func(*Options)) (*ListHostsOutput, error) 
 ListRepositoryLinks(ctx context.Context, params *ListRepositoryLinksInput, optFns ...func(*Options)) (*ListRepositoryLinksOutput, error) 
 ListRepositorySyncDefinitions(ctx context.Context, params *ListRepositorySyncDefinitionsInput, optFns ...func(*Options)) (*ListRepositorySyncDefinitionsOutput, error) 
 ListSyncConfigurations(ctx context.Context, params *ListSyncConfigurationsInput, optFns ...func(*Options)) (*ListSyncConfigurationsOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateHost(ctx context.Context, params *UpdateHostInput, optFns ...func(*Options)) (*UpdateHostOutput, error) 
 UpdateRepositoryLink(ctx context.Context, params *UpdateRepositoryLinkInput, optFns ...func(*Options)) (*UpdateRepositoryLinkOutput, error) 
 UpdateSyncBlocker(ctx context.Context, params *UpdateSyncBlockerInput, optFns ...func(*Options)) (*UpdateSyncBlockerOutput, error) 
 UpdateSyncConfiguration(ctx context.Context, params *UpdateSyncConfigurationInput, optFns ...func(*Options)) (*UpdateSyncConfigurationOutput, error) 
}
