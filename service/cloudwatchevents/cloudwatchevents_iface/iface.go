
package cloudwatchevents_iface

// This file was generated by scripts/gen_ifaces.sh do not edit it manually
import (
    "context"
    . "github.com/aws/aws-sdk-go-v2/service/cloudwatchevents"
)

// IClient defines the interface for cloudwatchevents
type IClient interface {
 Options() Options 
 ActivateEventSource(ctx context.Context, params *ActivateEventSourceInput, optFns ...func(*Options)) (*ActivateEventSourceOutput, error) 
 CancelReplay(ctx context.Context, params *CancelReplayInput, optFns ...func(*Options)) (*CancelReplayOutput, error) 
 CreateApiDestination(ctx context.Context, params *CreateApiDestinationInput, optFns ...func(*Options)) (*CreateApiDestinationOutput, error) 
 CreateArchive(ctx context.Context, params *CreateArchiveInput, optFns ...func(*Options)) (*CreateArchiveOutput, error) 
 CreateConnection(ctx context.Context, params *CreateConnectionInput, optFns ...func(*Options)) (*CreateConnectionOutput, error) 
 CreateEventBus(ctx context.Context, params *CreateEventBusInput, optFns ...func(*Options)) (*CreateEventBusOutput, error) 
 CreatePartnerEventSource(ctx context.Context, params *CreatePartnerEventSourceInput, optFns ...func(*Options)) (*CreatePartnerEventSourceOutput, error) 
 DeactivateEventSource(ctx context.Context, params *DeactivateEventSourceInput, optFns ...func(*Options)) (*DeactivateEventSourceOutput, error) 
 DeauthorizeConnection(ctx context.Context, params *DeauthorizeConnectionInput, optFns ...func(*Options)) (*DeauthorizeConnectionOutput, error) 
 DeleteApiDestination(ctx context.Context, params *DeleteApiDestinationInput, optFns ...func(*Options)) (*DeleteApiDestinationOutput, error) 
 DeleteArchive(ctx context.Context, params *DeleteArchiveInput, optFns ...func(*Options)) (*DeleteArchiveOutput, error) 
 DeleteConnection(ctx context.Context, params *DeleteConnectionInput, optFns ...func(*Options)) (*DeleteConnectionOutput, error) 
 DeleteEventBus(ctx context.Context, params *DeleteEventBusInput, optFns ...func(*Options)) (*DeleteEventBusOutput, error) 
 DeletePartnerEventSource(ctx context.Context, params *DeletePartnerEventSourceInput, optFns ...func(*Options)) (*DeletePartnerEventSourceOutput, error) 
 DeleteRule(ctx context.Context, params *DeleteRuleInput, optFns ...func(*Options)) (*DeleteRuleOutput, error) 
 DescribeApiDestination(ctx context.Context, params *DescribeApiDestinationInput, optFns ...func(*Options)) (*DescribeApiDestinationOutput, error) 
 DescribeArchive(ctx context.Context, params *DescribeArchiveInput, optFns ...func(*Options)) (*DescribeArchiveOutput, error) 
 DescribeConnection(ctx context.Context, params *DescribeConnectionInput, optFns ...func(*Options)) (*DescribeConnectionOutput, error) 
 DescribeEventBus(ctx context.Context, params *DescribeEventBusInput, optFns ...func(*Options)) (*DescribeEventBusOutput, error) 
 DescribeEventSource(ctx context.Context, params *DescribeEventSourceInput, optFns ...func(*Options)) (*DescribeEventSourceOutput, error) 
 DescribePartnerEventSource(ctx context.Context, params *DescribePartnerEventSourceInput, optFns ...func(*Options)) (*DescribePartnerEventSourceOutput, error) 
 DescribeReplay(ctx context.Context, params *DescribeReplayInput, optFns ...func(*Options)) (*DescribeReplayOutput, error) 
 DescribeRule(ctx context.Context, params *DescribeRuleInput, optFns ...func(*Options)) (*DescribeRuleOutput, error) 
 DisableRule(ctx context.Context, params *DisableRuleInput, optFns ...func(*Options)) (*DisableRuleOutput, error) 
 EnableRule(ctx context.Context, params *EnableRuleInput, optFns ...func(*Options)) (*EnableRuleOutput, error) 
 ListApiDestinations(ctx context.Context, params *ListApiDestinationsInput, optFns ...func(*Options)) (*ListApiDestinationsOutput, error) 
 ListArchives(ctx context.Context, params *ListArchivesInput, optFns ...func(*Options)) (*ListArchivesOutput, error) 
 ListConnections(ctx context.Context, params *ListConnectionsInput, optFns ...func(*Options)) (*ListConnectionsOutput, error) 
 ListEventBuses(ctx context.Context, params *ListEventBusesInput, optFns ...func(*Options)) (*ListEventBusesOutput, error) 
 ListEventSources(ctx context.Context, params *ListEventSourcesInput, optFns ...func(*Options)) (*ListEventSourcesOutput, error) 
 ListPartnerEventSourceAccounts(ctx context.Context, params *ListPartnerEventSourceAccountsInput, optFns ...func(*Options)) (*ListPartnerEventSourceAccountsOutput, error) 
 ListPartnerEventSources(ctx context.Context, params *ListPartnerEventSourcesInput, optFns ...func(*Options)) (*ListPartnerEventSourcesOutput, error) 
 ListReplays(ctx context.Context, params *ListReplaysInput, optFns ...func(*Options)) (*ListReplaysOutput, error) 
 ListRuleNamesByTarget(ctx context.Context, params *ListRuleNamesByTargetInput, optFns ...func(*Options)) (*ListRuleNamesByTargetOutput, error) 
 ListRules(ctx context.Context, params *ListRulesInput, optFns ...func(*Options)) (*ListRulesOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 ListTargetsByRule(ctx context.Context, params *ListTargetsByRuleInput, optFns ...func(*Options)) (*ListTargetsByRuleOutput, error) 
 PutEvents(ctx context.Context, params *PutEventsInput, optFns ...func(*Options)) (*PutEventsOutput, error) 
 PutPartnerEvents(ctx context.Context, params *PutPartnerEventsInput, optFns ...func(*Options)) (*PutPartnerEventsOutput, error) 
 PutPermission(ctx context.Context, params *PutPermissionInput, optFns ...func(*Options)) (*PutPermissionOutput, error) 
 PutRule(ctx context.Context, params *PutRuleInput, optFns ...func(*Options)) (*PutRuleOutput, error) 
 PutTargets(ctx context.Context, params *PutTargetsInput, optFns ...func(*Options)) (*PutTargetsOutput, error) 
 RemovePermission(ctx context.Context, params *RemovePermissionInput, optFns ...func(*Options)) (*RemovePermissionOutput, error) 
 RemoveTargets(ctx context.Context, params *RemoveTargetsInput, optFns ...func(*Options)) (*RemoveTargetsOutput, error) 
 StartReplay(ctx context.Context, params *StartReplayInput, optFns ...func(*Options)) (*StartReplayOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 TestEventPattern(ctx context.Context, params *TestEventPatternInput, optFns ...func(*Options)) (*TestEventPatternOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateApiDestination(ctx context.Context, params *UpdateApiDestinationInput, optFns ...func(*Options)) (*UpdateApiDestinationOutput, error) 
 UpdateArchive(ctx context.Context, params *UpdateArchiveInput, optFns ...func(*Options)) (*UpdateArchiveOutput, error) 
 UpdateConnection(ctx context.Context, params *UpdateConnectionInput, optFns ...func(*Options)) (*UpdateConnectionOutput, error) 
}