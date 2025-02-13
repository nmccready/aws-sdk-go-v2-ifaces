
package notifications_iface

// This file was generated by scripts/gen_ifaces.sh do not edit it manually
import (
    "context"
    . "github.com/aws/aws-sdk-go-v2/service/notifications"
)

// IClient defines the interface for notifications
type IClient interface {
 Options() Options 
 AssociateChannel(ctx context.Context, params *AssociateChannelInput, optFns ...func(*Options)) (*AssociateChannelOutput, error) 
 AssociateManagedNotificationAccountContact(ctx context.Context, params *AssociateManagedNotificationAccountContactInput, optFns ...func(*Options)) (*AssociateManagedNotificationAccountContactOutput, error) 
 AssociateManagedNotificationAdditionalChannel(ctx context.Context, params *AssociateManagedNotificationAdditionalChannelInput, optFns ...func(*Options)) (*AssociateManagedNotificationAdditionalChannelOutput, error) 
 CreateEventRule(ctx context.Context, params *CreateEventRuleInput, optFns ...func(*Options)) (*CreateEventRuleOutput, error) 
 CreateNotificationConfiguration(ctx context.Context, params *CreateNotificationConfigurationInput, optFns ...func(*Options)) (*CreateNotificationConfigurationOutput, error) 
 DeleteEventRule(ctx context.Context, params *DeleteEventRuleInput, optFns ...func(*Options)) (*DeleteEventRuleOutput, error) 
 DeleteNotificationConfiguration(ctx context.Context, params *DeleteNotificationConfigurationInput, optFns ...func(*Options)) (*DeleteNotificationConfigurationOutput, error) 
 DeregisterNotificationHub(ctx context.Context, params *DeregisterNotificationHubInput, optFns ...func(*Options)) (*DeregisterNotificationHubOutput, error) 
 DisableNotificationsAccessForOrganization(ctx context.Context, params *DisableNotificationsAccessForOrganizationInput, optFns ...func(*Options)) (*DisableNotificationsAccessForOrganizationOutput, error) 
 DisassociateChannel(ctx context.Context, params *DisassociateChannelInput, optFns ...func(*Options)) (*DisassociateChannelOutput, error) 
 DisassociateManagedNotificationAccountContact(ctx context.Context, params *DisassociateManagedNotificationAccountContactInput, optFns ...func(*Options)) (*DisassociateManagedNotificationAccountContactOutput, error) 
 DisassociateManagedNotificationAdditionalChannel(ctx context.Context, params *DisassociateManagedNotificationAdditionalChannelInput, optFns ...func(*Options)) (*DisassociateManagedNotificationAdditionalChannelOutput, error) 
 EnableNotificationsAccessForOrganization(ctx context.Context, params *EnableNotificationsAccessForOrganizationInput, optFns ...func(*Options)) (*EnableNotificationsAccessForOrganizationOutput, error) 
 GetEventRule(ctx context.Context, params *GetEventRuleInput, optFns ...func(*Options)) (*GetEventRuleOutput, error) 
 GetManagedNotificationChildEvent(ctx context.Context, params *GetManagedNotificationChildEventInput, optFns ...func(*Options)) (*GetManagedNotificationChildEventOutput, error) 
 GetManagedNotificationConfiguration(ctx context.Context, params *GetManagedNotificationConfigurationInput, optFns ...func(*Options)) (*GetManagedNotificationConfigurationOutput, error) 
 GetManagedNotificationEvent(ctx context.Context, params *GetManagedNotificationEventInput, optFns ...func(*Options)) (*GetManagedNotificationEventOutput, error) 
 GetNotificationConfiguration(ctx context.Context, params *GetNotificationConfigurationInput, optFns ...func(*Options)) (*GetNotificationConfigurationOutput, error) 
 GetNotificationEvent(ctx context.Context, params *GetNotificationEventInput, optFns ...func(*Options)) (*GetNotificationEventOutput, error) 
 GetNotificationsAccessForOrganization(ctx context.Context, params *GetNotificationsAccessForOrganizationInput, optFns ...func(*Options)) (*GetNotificationsAccessForOrganizationOutput, error) 
 ListChannels(ctx context.Context, params *ListChannelsInput, optFns ...func(*Options)) (*ListChannelsOutput, error) 
 ListEventRules(ctx context.Context, params *ListEventRulesInput, optFns ...func(*Options)) (*ListEventRulesOutput, error) 
 ListManagedNotificationChannelAssociations(ctx context.Context, params *ListManagedNotificationChannelAssociationsInput, optFns ...func(*Options)) (*ListManagedNotificationChannelAssociationsOutput, error) 
 ListManagedNotificationChildEvents(ctx context.Context, params *ListManagedNotificationChildEventsInput, optFns ...func(*Options)) (*ListManagedNotificationChildEventsOutput, error) 
 ListManagedNotificationConfigurations(ctx context.Context, params *ListManagedNotificationConfigurationsInput, optFns ...func(*Options)) (*ListManagedNotificationConfigurationsOutput, error) 
 ListManagedNotificationEvents(ctx context.Context, params *ListManagedNotificationEventsInput, optFns ...func(*Options)) (*ListManagedNotificationEventsOutput, error) 
 ListNotificationConfigurations(ctx context.Context, params *ListNotificationConfigurationsInput, optFns ...func(*Options)) (*ListNotificationConfigurationsOutput, error) 
 ListNotificationEvents(ctx context.Context, params *ListNotificationEventsInput, optFns ...func(*Options)) (*ListNotificationEventsOutput, error) 
 ListNotificationHubs(ctx context.Context, params *ListNotificationHubsInput, optFns ...func(*Options)) (*ListNotificationHubsOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 RegisterNotificationHub(ctx context.Context, params *RegisterNotificationHubInput, optFns ...func(*Options)) (*RegisterNotificationHubOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateEventRule(ctx context.Context, params *UpdateEventRuleInput, optFns ...func(*Options)) (*UpdateEventRuleOutput, error) 
 UpdateNotificationConfiguration(ctx context.Context, params *UpdateNotificationConfigurationInput, optFns ...func(*Options)) (*UpdateNotificationConfigurationOutput, error) 
}
