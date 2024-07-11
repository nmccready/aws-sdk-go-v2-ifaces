
package ssmincidents

import (
    "github.com/aws/aws-sdk-go-v2/service/ssmincidents"
)

// ISsmincidents defines the interface for ssmincidents
type ISsmincidents interface {
 Options() Options 
 BatchGetIncidentFindings(ctx context.Context, params *BatchGetIncidentFindingsInput, optFns ...func(*Options)) (*BatchGetIncidentFindingsOutput, error) 
 CreateReplicationSet(ctx context.Context, params *CreateReplicationSetInput, optFns ...func(*Options)) (*CreateReplicationSetOutput, error) 
 CreateResponsePlan(ctx context.Context, params *CreateResponsePlanInput, optFns ...func(*Options)) (*CreateResponsePlanOutput, error) 
 CreateTimelineEvent(ctx context.Context, params *CreateTimelineEventInput, optFns ...func(*Options)) (*CreateTimelineEventOutput, error) 
 DeleteIncidentRecord(ctx context.Context, params *DeleteIncidentRecordInput, optFns ...func(*Options)) (*DeleteIncidentRecordOutput, error) 
 DeleteReplicationSet(ctx context.Context, params *DeleteReplicationSetInput, optFns ...func(*Options)) (*DeleteReplicationSetOutput, error) 
 DeleteResourcePolicy(ctx context.Context, params *DeleteResourcePolicyInput, optFns ...func(*Options)) (*DeleteResourcePolicyOutput, error) 
 DeleteResponsePlan(ctx context.Context, params *DeleteResponsePlanInput, optFns ...func(*Options)) (*DeleteResponsePlanOutput, error) 
 DeleteTimelineEvent(ctx context.Context, params *DeleteTimelineEventInput, optFns ...func(*Options)) (*DeleteTimelineEventOutput, error) 
 GetIncidentRecord(ctx context.Context, params *GetIncidentRecordInput, optFns ...func(*Options)) (*GetIncidentRecordOutput, error) 
 GetReplicationSet(ctx context.Context, params *GetReplicationSetInput, optFns ...func(*Options)) (*GetReplicationSetOutput, error) 
 GetResourcePolicies(ctx context.Context, params *GetResourcePoliciesInput, optFns ...func(*Options)) (*GetResourcePoliciesOutput, error) 
 GetResponsePlan(ctx context.Context, params *GetResponsePlanInput, optFns ...func(*Options)) (*GetResponsePlanOutput, error) 
 GetTimelineEvent(ctx context.Context, params *GetTimelineEventInput, optFns ...func(*Options)) (*GetTimelineEventOutput, error) 
 ListIncidentFindings(ctx context.Context, params *ListIncidentFindingsInput, optFns ...func(*Options)) (*ListIncidentFindingsOutput, error) 
 ListIncidentRecords(ctx context.Context, params *ListIncidentRecordsInput, optFns ...func(*Options)) (*ListIncidentRecordsOutput, error) 
 ListRelatedItems(ctx context.Context, params *ListRelatedItemsInput, optFns ...func(*Options)) (*ListRelatedItemsOutput, error) 
 ListReplicationSets(ctx context.Context, params *ListReplicationSetsInput, optFns ...func(*Options)) (*ListReplicationSetsOutput, error) 
 ListResponsePlans(ctx context.Context, params *ListResponsePlansInput, optFns ...func(*Options)) (*ListResponsePlansOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 ListTimelineEvents(ctx context.Context, params *ListTimelineEventsInput, optFns ...func(*Options)) (*ListTimelineEventsOutput, error) 
 PutResourcePolicy(ctx context.Context, params *PutResourcePolicyInput, optFns ...func(*Options)) (*PutResourcePolicyOutput, error) 
 StartIncident(ctx context.Context, params *StartIncidentInput, optFns ...func(*Options)) (*StartIncidentOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateDeletionProtection(ctx context.Context, params *UpdateDeletionProtectionInput, optFns ...func(*Options)) (*UpdateDeletionProtectionOutput, error) 
 UpdateIncidentRecord(ctx context.Context, params *UpdateIncidentRecordInput, optFns ...func(*Options)) (*UpdateIncidentRecordOutput, error) 
 UpdateRelatedItems(ctx context.Context, params *UpdateRelatedItemsInput, optFns ...func(*Options)) (*UpdateRelatedItemsOutput, error) 
 UpdateReplicationSet(ctx context.Context, params *UpdateReplicationSetInput, optFns ...func(*Options)) (*UpdateReplicationSetOutput, error) 
 UpdateResponsePlan(ctx context.Context, params *UpdateResponsePlanInput, optFns ...func(*Options)) (*UpdateResponsePlanOutput, error) 
 UpdateTimelineEvent(ctx context.Context, params *UpdateTimelineEventInput, optFns ...func(*Options)) (*UpdateTimelineEventOutput, error) 
}
