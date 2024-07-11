
package keyspaces

import (
    "github.com/aws/aws-sdk-go-v2/service/keyspaces"
)

// IKeyspaces defines the interface for keyspaces
type IKeyspaces interface {
 Options() Options 
 CreateKeyspace(ctx context.Context, params *CreateKeyspaceInput, optFns ...func(*Options)) (*CreateKeyspaceOutput, error) 
 CreateTable(ctx context.Context, params *CreateTableInput, optFns ...func(*Options)) (*CreateTableOutput, error) 
 DeleteKeyspace(ctx context.Context, params *DeleteKeyspaceInput, optFns ...func(*Options)) (*DeleteKeyspaceOutput, error) 
 DeleteTable(ctx context.Context, params *DeleteTableInput, optFns ...func(*Options)) (*DeleteTableOutput, error) 
 GetKeyspace(ctx context.Context, params *GetKeyspaceInput, optFns ...func(*Options)) (*GetKeyspaceOutput, error) 
 GetTable(ctx context.Context, params *GetTableInput, optFns ...func(*Options)) (*GetTableOutput, error) 
 GetTableAutoScalingSettings(ctx context.Context, params *GetTableAutoScalingSettingsInput, optFns ...func(*Options)) (*GetTableAutoScalingSettingsOutput, error) 
 ListKeyspaces(ctx context.Context, params *ListKeyspacesInput, optFns ...func(*Options)) (*ListKeyspacesOutput, error) 
 ListTables(ctx context.Context, params *ListTablesInput, optFns ...func(*Options)) (*ListTablesOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 RestoreTable(ctx context.Context, params *RestoreTableInput, optFns ...func(*Options)) (*RestoreTableOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateTable(ctx context.Context, params *UpdateTableInput, optFns ...func(*Options)) (*UpdateTableOutput, error) 
}
