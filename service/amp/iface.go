
package amp

import (
    "github.com/aws/aws-sdk-go-v2/service/amp"
)

// IAmp defines the interface for amp
type IAmp interface {
 Options() Options 
 CreateAlertManagerDefinition(ctx context.Context, params *CreateAlertManagerDefinitionInput, optFns ...func(*Options)) (*CreateAlertManagerDefinitionOutput, error) 
 CreateLoggingConfiguration(ctx context.Context, params *CreateLoggingConfigurationInput, optFns ...func(*Options)) (*CreateLoggingConfigurationOutput, error) 
 CreateRuleGroupsNamespace(ctx context.Context, params *CreateRuleGroupsNamespaceInput, optFns ...func(*Options)) (*CreateRuleGroupsNamespaceOutput, error) 
 CreateScraper(ctx context.Context, params *CreateScraperInput, optFns ...func(*Options)) (*CreateScraperOutput, error) 
 CreateWorkspace(ctx context.Context, params *CreateWorkspaceInput, optFns ...func(*Options)) (*CreateWorkspaceOutput, error) 
 DeleteAlertManagerDefinition(ctx context.Context, params *DeleteAlertManagerDefinitionInput, optFns ...func(*Options)) (*DeleteAlertManagerDefinitionOutput, error) 
 DeleteLoggingConfiguration(ctx context.Context, params *DeleteLoggingConfigurationInput, optFns ...func(*Options)) (*DeleteLoggingConfigurationOutput, error) 
 DeleteRuleGroupsNamespace(ctx context.Context, params *DeleteRuleGroupsNamespaceInput, optFns ...func(*Options)) (*DeleteRuleGroupsNamespaceOutput, error) 
 DeleteScraper(ctx context.Context, params *DeleteScraperInput, optFns ...func(*Options)) (*DeleteScraperOutput, error) 
 DeleteWorkspace(ctx context.Context, params *DeleteWorkspaceInput, optFns ...func(*Options)) (*DeleteWorkspaceOutput, error) 
 DescribeAlertManagerDefinition(ctx context.Context, params *DescribeAlertManagerDefinitionInput, optFns ...func(*Options)) (*DescribeAlertManagerDefinitionOutput, error) 
 DescribeLoggingConfiguration(ctx context.Context, params *DescribeLoggingConfigurationInput, optFns ...func(*Options)) (*DescribeLoggingConfigurationOutput, error) 
 DescribeRuleGroupsNamespace(ctx context.Context, params *DescribeRuleGroupsNamespaceInput, optFns ...func(*Options)) (*DescribeRuleGroupsNamespaceOutput, error) 
 DescribeScraper(ctx context.Context, params *DescribeScraperInput, optFns ...func(*Options)) (*DescribeScraperOutput, error) 
 DescribeWorkspace(ctx context.Context, params *DescribeWorkspaceInput, optFns ...func(*Options)) (*DescribeWorkspaceOutput, error) 
 GetDefaultScraperConfiguration(ctx context.Context, params *GetDefaultScraperConfigurationInput, optFns ...func(*Options)) (*GetDefaultScraperConfigurationOutput, error) 
 ListRuleGroupsNamespaces(ctx context.Context, params *ListRuleGroupsNamespacesInput, optFns ...func(*Options)) (*ListRuleGroupsNamespacesOutput, error) 
 ListScrapers(ctx context.Context, params *ListScrapersInput, optFns ...func(*Options)) (*ListScrapersOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 ListWorkspaces(ctx context.Context, params *ListWorkspacesInput, optFns ...func(*Options)) (*ListWorkspacesOutput, error) 
 PutAlertManagerDefinition(ctx context.Context, params *PutAlertManagerDefinitionInput, optFns ...func(*Options)) (*PutAlertManagerDefinitionOutput, error) 
 PutRuleGroupsNamespace(ctx context.Context, params *PutRuleGroupsNamespaceInput, optFns ...func(*Options)) (*PutRuleGroupsNamespaceOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateLoggingConfiguration(ctx context.Context, params *UpdateLoggingConfigurationInput, optFns ...func(*Options)) (*UpdateLoggingConfigurationOutput, error) 
 UpdateWorkspaceAlias(ctx context.Context, params *UpdateWorkspaceAliasInput, optFns ...func(*Options)) (*UpdateWorkspaceAliasOutput, error) 
}
