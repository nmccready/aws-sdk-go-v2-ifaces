
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
 CreateEventRule(ctx context.Context, params *CreateEventRuleInput, optFns ...func(*Options)) (*CreateEventRuleOutput, error) 
 CreateNotificationConfiguration(ctx context.Context, params *CreateNotificationConfigurationInput, optFns ...func(*Options)) (*CreateNotificationConfigurationOutput, error) 
 DeleteEventRule(ctx context.Context, params *DeleteEventRuleInput, optFns ...func(*Options)) (*DeleteEventRuleOutput, error) 
 DeleteNotificationConfiguration(ctx context.Context, params *DeleteNotificationConfigurationInput, optFns ...func(*Options)) (*DeleteNotificationConfigurationOutput, error) 
 DeregisterNotificationHub(ctx context.Context, params *DeregisterNotificationHubInput, optFns ...func(*Options)) (*DeregisterNotificationHubOutput, error) 
 DisassociateChannel(ctx context.Context, params *DisassociateChannelInput, optFns ...func(*Options)) (*DisassociateChannelOutput, error) 
 GetEventRule(ctx context.Context, params *GetEventRuleInput, optFns ...func(*Options)) (*GetEventRuleOutput, error) 
 GetNotificationConfiguration(ctx context.Context, params *GetNotificationConfigurationInput, optFns ...func(*Options)) (*GetNotificationConfigurationOutput, error) 
 GetNotificationEvent(ctx context.Context, params *GetNotificationEventInput, optFns ...func(*Options)) (*GetNotificationEventOutput, error) 
 ListChannels(ctx context.Context, params *ListChannelsInput, optFns ...func(*Options)) (*ListChannelsOutput, error) 
 ListEventRules(ctx context.Context, params *ListEventRulesInput, optFns ...func(*Options)) (*ListEventRulesOutput, error) 
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