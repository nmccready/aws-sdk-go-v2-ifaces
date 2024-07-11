
package supportapp

import (
    "github.com/aws/aws-sdk-go-v2/service/supportapp"
)

// ISupportapp defines the interface for supportapp
type ISupportapp interface {
 Options() Options 
 CreateSlackChannelConfiguration(ctx context.Context, params *CreateSlackChannelConfigurationInput, optFns ...func(*Options)) (*CreateSlackChannelConfigurationOutput, error) 
 DeleteAccountAlias(ctx context.Context, params *DeleteAccountAliasInput, optFns ...func(*Options)) (*DeleteAccountAliasOutput, error) 
 DeleteSlackChannelConfiguration(ctx context.Context, params *DeleteSlackChannelConfigurationInput, optFns ...func(*Options)) (*DeleteSlackChannelConfigurationOutput, error) 
 DeleteSlackWorkspaceConfiguration(ctx context.Context, params *DeleteSlackWorkspaceConfigurationInput, optFns ...func(*Options)) (*DeleteSlackWorkspaceConfigurationOutput, error) 
 GetAccountAlias(ctx context.Context, params *GetAccountAliasInput, optFns ...func(*Options)) (*GetAccountAliasOutput, error) 
 ListSlackChannelConfigurations(ctx context.Context, params *ListSlackChannelConfigurationsInput, optFns ...func(*Options)) (*ListSlackChannelConfigurationsOutput, error) 
 ListSlackWorkspaceConfigurations(ctx context.Context, params *ListSlackWorkspaceConfigurationsInput, optFns ...func(*Options)) (*ListSlackWorkspaceConfigurationsOutput, error) 
 PutAccountAlias(ctx context.Context, params *PutAccountAliasInput, optFns ...func(*Options)) (*PutAccountAliasOutput, error) 
 RegisterSlackWorkspaceForOrganization(ctx context.Context, params *RegisterSlackWorkspaceForOrganizationInput, optFns ...func(*Options)) (*RegisterSlackWorkspaceForOrganizationOutput, error) 
 UpdateSlackChannelConfiguration(ctx context.Context, params *UpdateSlackChannelConfigurationInput, optFns ...func(*Options)) (*UpdateSlackChannelConfigurationOutput, error) 
}
