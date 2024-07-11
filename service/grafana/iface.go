
package grafana

import (
    "github.com/aws/aws-sdk-go-v2/service/grafana"
)

// IClient defines the interface for grafana
type IClient interface {
 Options() Options 
 AssociateLicense(ctx context.Context, params *AssociateLicenseInput, optFns ...func(*Options)) (*AssociateLicenseOutput, error) 
 CreateWorkspace(ctx context.Context, params *CreateWorkspaceInput, optFns ...func(*Options)) (*CreateWorkspaceOutput, error) 
 CreateWorkspaceApiKey(ctx context.Context, params *CreateWorkspaceApiKeyInput, optFns ...func(*Options)) (*CreateWorkspaceApiKeyOutput, error) 
 CreateWorkspaceServiceAccount(ctx context.Context, params *CreateWorkspaceServiceAccountInput, optFns ...func(*Options)) (*CreateWorkspaceServiceAccountOutput, error) 
 CreateWorkspaceServiceAccountToken(ctx context.Context, params *CreateWorkspaceServiceAccountTokenInput, optFns ...func(*Options)) (*CreateWorkspaceServiceAccountTokenOutput, error) 
 DeleteWorkspace(ctx context.Context, params *DeleteWorkspaceInput, optFns ...func(*Options)) (*DeleteWorkspaceOutput, error) 
 DeleteWorkspaceApiKey(ctx context.Context, params *DeleteWorkspaceApiKeyInput, optFns ...func(*Options)) (*DeleteWorkspaceApiKeyOutput, error) 
 DeleteWorkspaceServiceAccount(ctx context.Context, params *DeleteWorkspaceServiceAccountInput, optFns ...func(*Options)) (*DeleteWorkspaceServiceAccountOutput, error) 
 DeleteWorkspaceServiceAccountToken(ctx context.Context, params *DeleteWorkspaceServiceAccountTokenInput, optFns ...func(*Options)) (*DeleteWorkspaceServiceAccountTokenOutput, error) 
 DescribeWorkspace(ctx context.Context, params *DescribeWorkspaceInput, optFns ...func(*Options)) (*DescribeWorkspaceOutput, error) 
 DescribeWorkspaceAuthentication(ctx context.Context, params *DescribeWorkspaceAuthenticationInput, optFns ...func(*Options)) (*DescribeWorkspaceAuthenticationOutput, error) 
 DescribeWorkspaceConfiguration(ctx context.Context, params *DescribeWorkspaceConfigurationInput, optFns ...func(*Options)) (*DescribeWorkspaceConfigurationOutput, error) 
 DisassociateLicense(ctx context.Context, params *DisassociateLicenseInput, optFns ...func(*Options)) (*DisassociateLicenseOutput, error) 
 ListPermissions(ctx context.Context, params *ListPermissionsInput, optFns ...func(*Options)) (*ListPermissionsOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 ListVersions(ctx context.Context, params *ListVersionsInput, optFns ...func(*Options)) (*ListVersionsOutput, error) 
 ListWorkspaceServiceAccountTokens(ctx context.Context, params *ListWorkspaceServiceAccountTokensInput, optFns ...func(*Options)) (*ListWorkspaceServiceAccountTokensOutput, error) 
 ListWorkspaceServiceAccounts(ctx context.Context, params *ListWorkspaceServiceAccountsInput, optFns ...func(*Options)) (*ListWorkspaceServiceAccountsOutput, error) 
 ListWorkspaces(ctx context.Context, params *ListWorkspacesInput, optFns ...func(*Options)) (*ListWorkspacesOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdatePermissions(ctx context.Context, params *UpdatePermissionsInput, optFns ...func(*Options)) (*UpdatePermissionsOutput, error) 
 UpdateWorkspace(ctx context.Context, params *UpdateWorkspaceInput, optFns ...func(*Options)) (*UpdateWorkspaceOutput, error) 
 UpdateWorkspaceAuthentication(ctx context.Context, params *UpdateWorkspaceAuthenticationInput, optFns ...func(*Options)) (*UpdateWorkspaceAuthenticationOutput, error) 
 UpdateWorkspaceConfiguration(ctx context.Context, params *UpdateWorkspaceConfigurationInput, optFns ...func(*Options)) (*UpdateWorkspaceConfigurationOutput, error) 
}
