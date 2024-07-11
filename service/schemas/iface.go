
package schemas

import (
    "github.com/aws/aws-sdk-go-v2/service/schemas"
)

// ISchemas defines the interface for schemas
type ISchemas interface {
 Options() Options 
 CreateDiscoverer(ctx context.Context, params *CreateDiscovererInput, optFns ...func(*Options)) (*CreateDiscovererOutput, error) 
 CreateRegistry(ctx context.Context, params *CreateRegistryInput, optFns ...func(*Options)) (*CreateRegistryOutput, error) 
 CreateSchema(ctx context.Context, params *CreateSchemaInput, optFns ...func(*Options)) (*CreateSchemaOutput, error) 
 DeleteDiscoverer(ctx context.Context, params *DeleteDiscovererInput, optFns ...func(*Options)) (*DeleteDiscovererOutput, error) 
 DeleteRegistry(ctx context.Context, params *DeleteRegistryInput, optFns ...func(*Options)) (*DeleteRegistryOutput, error) 
 DeleteResourcePolicy(ctx context.Context, params *DeleteResourcePolicyInput, optFns ...func(*Options)) (*DeleteResourcePolicyOutput, error) 
 DeleteSchema(ctx context.Context, params *DeleteSchemaInput, optFns ...func(*Options)) (*DeleteSchemaOutput, error) 
 DeleteSchemaVersion(ctx context.Context, params *DeleteSchemaVersionInput, optFns ...func(*Options)) (*DeleteSchemaVersionOutput, error) 
 DescribeCodeBinding(ctx context.Context, params *DescribeCodeBindingInput, optFns ...func(*Options)) (*DescribeCodeBindingOutput, error) 
 DescribeDiscoverer(ctx context.Context, params *DescribeDiscovererInput, optFns ...func(*Options)) (*DescribeDiscovererOutput, error) 
 DescribeRegistry(ctx context.Context, params *DescribeRegistryInput, optFns ...func(*Options)) (*DescribeRegistryOutput, error) 
 DescribeSchema(ctx context.Context, params *DescribeSchemaInput, optFns ...func(*Options)) (*DescribeSchemaOutput, error) 
 ExportSchema(ctx context.Context, params *ExportSchemaInput, optFns ...func(*Options)) (*ExportSchemaOutput, error) 
 GetCodeBindingSource(ctx context.Context, params *GetCodeBindingSourceInput, optFns ...func(*Options)) (*GetCodeBindingSourceOutput, error) 
 GetDiscoveredSchema(ctx context.Context, params *GetDiscoveredSchemaInput, optFns ...func(*Options)) (*GetDiscoveredSchemaOutput, error) 
 GetResourcePolicy(ctx context.Context, params *GetResourcePolicyInput, optFns ...func(*Options)) (*GetResourcePolicyOutput, error) 
 ListDiscoverers(ctx context.Context, params *ListDiscoverersInput, optFns ...func(*Options)) (*ListDiscoverersOutput, error) 
 ListRegistries(ctx context.Context, params *ListRegistriesInput, optFns ...func(*Options)) (*ListRegistriesOutput, error) 
 ListSchemaVersions(ctx context.Context, params *ListSchemaVersionsInput, optFns ...func(*Options)) (*ListSchemaVersionsOutput, error) 
 ListSchemas(ctx context.Context, params *ListSchemasInput, optFns ...func(*Options)) (*ListSchemasOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 PutCodeBinding(ctx context.Context, params *PutCodeBindingInput, optFns ...func(*Options)) (*PutCodeBindingOutput, error) 
 PutResourcePolicy(ctx context.Context, params *PutResourcePolicyInput, optFns ...func(*Options)) (*PutResourcePolicyOutput, error) 
 SearchSchemas(ctx context.Context, params *SearchSchemasInput, optFns ...func(*Options)) (*SearchSchemasOutput, error) 
 StartDiscoverer(ctx context.Context, params *StartDiscovererInput, optFns ...func(*Options)) (*StartDiscovererOutput, error) 
 StopDiscoverer(ctx context.Context, params *StopDiscovererInput, optFns ...func(*Options)) (*StopDiscovererOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateDiscoverer(ctx context.Context, params *UpdateDiscovererInput, optFns ...func(*Options)) (*UpdateDiscovererOutput, error) 
 UpdateRegistry(ctx context.Context, params *UpdateRegistryInput, optFns ...func(*Options)) (*UpdateRegistryOutput, error) 
 UpdateSchema(ctx context.Context, params *UpdateSchemaInput, optFns ...func(*Options)) (*UpdateSchemaOutput, error) 
}
