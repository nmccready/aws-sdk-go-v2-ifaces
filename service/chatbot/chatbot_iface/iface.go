
package chatbot_iface

// This file was generated by scripts/gen_ifaces.sh do not edit it manually
import (
    "context"
    . "github.com/aws/aws-sdk-go-v2/service/chatbot"
)

// IClient defines the interface for chatbot
type IClient interface {
 Options() Options 
 AssociateToConfiguration(ctx context.Context, params *AssociateToConfigurationInput, optFns ...func(*Options)) (*AssociateToConfigurationOutput, error) 
 CreateChimeWebhookConfiguration(ctx context.Context, params *CreateChimeWebhookConfigurationInput, optFns ...func(*Options)) (*CreateChimeWebhookConfigurationOutput, error) 
 CreateCustomAction(ctx context.Context, params *CreateCustomActionInput, optFns ...func(*Options)) (*CreateCustomActionOutput, error) 
 CreateMicrosoftTeamsChannelConfiguration(ctx context.Context, params *CreateMicrosoftTeamsChannelConfigurationInput, optFns ...func(*Options)) (*CreateMicrosoftTeamsChannelConfigurationOutput, error) 
 CreateSlackChannelConfiguration(ctx context.Context, params *CreateSlackChannelConfigurationInput, optFns ...func(*Options)) (*CreateSlackChannelConfigurationOutput, error) 
 DeleteChimeWebhookConfiguration(ctx context.Context, params *DeleteChimeWebhookConfigurationInput, optFns ...func(*Options)) (*DeleteChimeWebhookConfigurationOutput, error) 
 DeleteCustomAction(ctx context.Context, params *DeleteCustomActionInput, optFns ...func(*Options)) (*DeleteCustomActionOutput, error) 
 DeleteMicrosoftTeamsChannelConfiguration(ctx context.Context, params *DeleteMicrosoftTeamsChannelConfigurationInput, optFns ...func(*Options)) (*DeleteMicrosoftTeamsChannelConfigurationOutput, error) 
 DeleteMicrosoftTeamsConfiguredTeam(ctx context.Context, params *DeleteMicrosoftTeamsConfiguredTeamInput, optFns ...func(*Options)) (*DeleteMicrosoftTeamsConfiguredTeamOutput, error) 
 DeleteMicrosoftTeamsUserIdentity(ctx context.Context, params *DeleteMicrosoftTeamsUserIdentityInput, optFns ...func(*Options)) (*DeleteMicrosoftTeamsUserIdentityOutput, error) 
 DeleteSlackChannelConfiguration(ctx context.Context, params *DeleteSlackChannelConfigurationInput, optFns ...func(*Options)) (*DeleteSlackChannelConfigurationOutput, error) 
 DeleteSlackUserIdentity(ctx context.Context, params *DeleteSlackUserIdentityInput, optFns ...func(*Options)) (*DeleteSlackUserIdentityOutput, error) 
 DeleteSlackWorkspaceAuthorization(ctx context.Context, params *DeleteSlackWorkspaceAuthorizationInput, optFns ...func(*Options)) (*DeleteSlackWorkspaceAuthorizationOutput, error) 
 DescribeChimeWebhookConfigurations(ctx context.Context, params *DescribeChimeWebhookConfigurationsInput, optFns ...func(*Options)) (*DescribeChimeWebhookConfigurationsOutput, error) 
 DescribeSlackChannelConfigurations(ctx context.Context, params *DescribeSlackChannelConfigurationsInput, optFns ...func(*Options)) (*DescribeSlackChannelConfigurationsOutput, error) 
 DescribeSlackUserIdentities(ctx context.Context, params *DescribeSlackUserIdentitiesInput, optFns ...func(*Options)) (*DescribeSlackUserIdentitiesOutput, error) 
 DescribeSlackWorkspaces(ctx context.Context, params *DescribeSlackWorkspacesInput, optFns ...func(*Options)) (*DescribeSlackWorkspacesOutput, error) 
 DisassociateFromConfiguration(ctx context.Context, params *DisassociateFromConfigurationInput, optFns ...func(*Options)) (*DisassociateFromConfigurationOutput, error) 
 GetAccountPreferences(ctx context.Context, params *GetAccountPreferencesInput, optFns ...func(*Options)) (*GetAccountPreferencesOutput, error) 
 GetCustomAction(ctx context.Context, params *GetCustomActionInput, optFns ...func(*Options)) (*GetCustomActionOutput, error) 
 GetMicrosoftTeamsChannelConfiguration(ctx context.Context, params *GetMicrosoftTeamsChannelConfigurationInput, optFns ...func(*Options)) (*GetMicrosoftTeamsChannelConfigurationOutput, error) 
 ListAssociations(ctx context.Context, params *ListAssociationsInput, optFns ...func(*Options)) (*ListAssociationsOutput, error) 
 ListCustomActions(ctx context.Context, params *ListCustomActionsInput, optFns ...func(*Options)) (*ListCustomActionsOutput, error) 
 ListMicrosoftTeamsChannelConfigurations(ctx context.Context, params *ListMicrosoftTeamsChannelConfigurationsInput, optFns ...func(*Options)) (*ListMicrosoftTeamsChannelConfigurationsOutput, error) 
 ListMicrosoftTeamsConfiguredTeams(ctx context.Context, params *ListMicrosoftTeamsConfiguredTeamsInput, optFns ...func(*Options)) (*ListMicrosoftTeamsConfiguredTeamsOutput, error) 
 ListMicrosoftTeamsUserIdentities(ctx context.Context, params *ListMicrosoftTeamsUserIdentitiesInput, optFns ...func(*Options)) (*ListMicrosoftTeamsUserIdentitiesOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateAccountPreferences(ctx context.Context, params *UpdateAccountPreferencesInput, optFns ...func(*Options)) (*UpdateAccountPreferencesOutput, error) 
 UpdateChimeWebhookConfiguration(ctx context.Context, params *UpdateChimeWebhookConfigurationInput, optFns ...func(*Options)) (*UpdateChimeWebhookConfigurationOutput, error) 
 UpdateCustomAction(ctx context.Context, params *UpdateCustomActionInput, optFns ...func(*Options)) (*UpdateCustomActionOutput, error) 
 UpdateMicrosoftTeamsChannelConfiguration(ctx context.Context, params *UpdateMicrosoftTeamsChannelConfigurationInput, optFns ...func(*Options)) (*UpdateMicrosoftTeamsChannelConfigurationOutput, error) 
 UpdateSlackChannelConfiguration(ctx context.Context, params *UpdateSlackChannelConfigurationInput, optFns ...func(*Options)) (*UpdateSlackChannelConfigurationOutput, error) 
}
