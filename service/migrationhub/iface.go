
package migrationhub

// This file was generated by scripts/gen_ifaces.sh do not edit it manually
import (
    "context"
    . "github.com/aws/aws-sdk-go-v2/service/migrationhub"
)

// IClient defines the interface for migrationhub
type IClient interface {
 Options() Options 
 AssociateCreatedArtifact(ctx context.Context, params *AssociateCreatedArtifactInput, optFns ...func(*Options)) (*AssociateCreatedArtifactOutput, error) 
 AssociateDiscoveredResource(ctx context.Context, params *AssociateDiscoveredResourceInput, optFns ...func(*Options)) (*AssociateDiscoveredResourceOutput, error) 
 CreateProgressUpdateStream(ctx context.Context, params *CreateProgressUpdateStreamInput, optFns ...func(*Options)) (*CreateProgressUpdateStreamOutput, error) 
 DeleteProgressUpdateStream(ctx context.Context, params *DeleteProgressUpdateStreamInput, optFns ...func(*Options)) (*DeleteProgressUpdateStreamOutput, error) 
 DescribeApplicationState(ctx context.Context, params *DescribeApplicationStateInput, optFns ...func(*Options)) (*DescribeApplicationStateOutput, error) 
 DescribeMigrationTask(ctx context.Context, params *DescribeMigrationTaskInput, optFns ...func(*Options)) (*DescribeMigrationTaskOutput, error) 
 DisassociateCreatedArtifact(ctx context.Context, params *DisassociateCreatedArtifactInput, optFns ...func(*Options)) (*DisassociateCreatedArtifactOutput, error) 
 DisassociateDiscoveredResource(ctx context.Context, params *DisassociateDiscoveredResourceInput, optFns ...func(*Options)) (*DisassociateDiscoveredResourceOutput, error) 
 ImportMigrationTask(ctx context.Context, params *ImportMigrationTaskInput, optFns ...func(*Options)) (*ImportMigrationTaskOutput, error) 
 ListApplicationStates(ctx context.Context, params *ListApplicationStatesInput, optFns ...func(*Options)) (*ListApplicationStatesOutput, error) 
 ListCreatedArtifacts(ctx context.Context, params *ListCreatedArtifactsInput, optFns ...func(*Options)) (*ListCreatedArtifactsOutput, error) 
 ListDiscoveredResources(ctx context.Context, params *ListDiscoveredResourcesInput, optFns ...func(*Options)) (*ListDiscoveredResourcesOutput, error) 
 ListMigrationTasks(ctx context.Context, params *ListMigrationTasksInput, optFns ...func(*Options)) (*ListMigrationTasksOutput, error) 
 ListProgressUpdateStreams(ctx context.Context, params *ListProgressUpdateStreamsInput, optFns ...func(*Options)) (*ListProgressUpdateStreamsOutput, error) 
 NotifyApplicationState(ctx context.Context, params *NotifyApplicationStateInput, optFns ...func(*Options)) (*NotifyApplicationStateOutput, error) 
 NotifyMigrationTaskState(ctx context.Context, params *NotifyMigrationTaskStateInput, optFns ...func(*Options)) (*NotifyMigrationTaskStateOutput, error) 
 PutResourceAttributes(ctx context.Context, params *PutResourceAttributesInput, optFns ...func(*Options)) (*PutResourceAttributesOutput, error) 
}
