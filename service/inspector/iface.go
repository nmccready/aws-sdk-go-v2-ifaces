
package inspector

import (
    "github.com/aws/aws-sdk-go-v2/service/inspector"
)

// IClient defines the interface for inspector
type IClient interface {
 Options() Options 
 AddAttributesToFindings(ctx context.Context, params *AddAttributesToFindingsInput, optFns ...func(*Options)) (*AddAttributesToFindingsOutput, error) 
 CreateAssessmentTarget(ctx context.Context, params *CreateAssessmentTargetInput, optFns ...func(*Options)) (*CreateAssessmentTargetOutput, error) 
 CreateAssessmentTemplate(ctx context.Context, params *CreateAssessmentTemplateInput, optFns ...func(*Options)) (*CreateAssessmentTemplateOutput, error) 
 CreateExclusionsPreview(ctx context.Context, params *CreateExclusionsPreviewInput, optFns ...func(*Options)) (*CreateExclusionsPreviewOutput, error) 
 CreateResourceGroup(ctx context.Context, params *CreateResourceGroupInput, optFns ...func(*Options)) (*CreateResourceGroupOutput, error) 
 DeleteAssessmentRun(ctx context.Context, params *DeleteAssessmentRunInput, optFns ...func(*Options)) (*DeleteAssessmentRunOutput, error) 
 DeleteAssessmentTarget(ctx context.Context, params *DeleteAssessmentTargetInput, optFns ...func(*Options)) (*DeleteAssessmentTargetOutput, error) 
 DeleteAssessmentTemplate(ctx context.Context, params *DeleteAssessmentTemplateInput, optFns ...func(*Options)) (*DeleteAssessmentTemplateOutput, error) 
 DescribeAssessmentRuns(ctx context.Context, params *DescribeAssessmentRunsInput, optFns ...func(*Options)) (*DescribeAssessmentRunsOutput, error) 
 DescribeAssessmentTargets(ctx context.Context, params *DescribeAssessmentTargetsInput, optFns ...func(*Options)) (*DescribeAssessmentTargetsOutput, error) 
 DescribeAssessmentTemplates(ctx context.Context, params *DescribeAssessmentTemplatesInput, optFns ...func(*Options)) (*DescribeAssessmentTemplatesOutput, error) 
 DescribeCrossAccountAccessRole(ctx context.Context, params *DescribeCrossAccountAccessRoleInput, optFns ...func(*Options)) (*DescribeCrossAccountAccessRoleOutput, error) 
 DescribeExclusions(ctx context.Context, params *DescribeExclusionsInput, optFns ...func(*Options)) (*DescribeExclusionsOutput, error) 
 DescribeFindings(ctx context.Context, params *DescribeFindingsInput, optFns ...func(*Options)) (*DescribeFindingsOutput, error) 
 DescribeResourceGroups(ctx context.Context, params *DescribeResourceGroupsInput, optFns ...func(*Options)) (*DescribeResourceGroupsOutput, error) 
 DescribeRulesPackages(ctx context.Context, params *DescribeRulesPackagesInput, optFns ...func(*Options)) (*DescribeRulesPackagesOutput, error) 
 GetAssessmentReport(ctx context.Context, params *GetAssessmentReportInput, optFns ...func(*Options)) (*GetAssessmentReportOutput, error) 
 GetExclusionsPreview(ctx context.Context, params *GetExclusionsPreviewInput, optFns ...func(*Options)) (*GetExclusionsPreviewOutput, error) 
 GetTelemetryMetadata(ctx context.Context, params *GetTelemetryMetadataInput, optFns ...func(*Options)) (*GetTelemetryMetadataOutput, error) 
 ListAssessmentRunAgents(ctx context.Context, params *ListAssessmentRunAgentsInput, optFns ...func(*Options)) (*ListAssessmentRunAgentsOutput, error) 
 ListAssessmentRuns(ctx context.Context, params *ListAssessmentRunsInput, optFns ...func(*Options)) (*ListAssessmentRunsOutput, error) 
 ListAssessmentTargets(ctx context.Context, params *ListAssessmentTargetsInput, optFns ...func(*Options)) (*ListAssessmentTargetsOutput, error) 
 ListAssessmentTemplates(ctx context.Context, params *ListAssessmentTemplatesInput, optFns ...func(*Options)) (*ListAssessmentTemplatesOutput, error) 
 ListEventSubscriptions(ctx context.Context, params *ListEventSubscriptionsInput, optFns ...func(*Options)) (*ListEventSubscriptionsOutput, error) 
 ListExclusions(ctx context.Context, params *ListExclusionsInput, optFns ...func(*Options)) (*ListExclusionsOutput, error) 
 ListFindings(ctx context.Context, params *ListFindingsInput, optFns ...func(*Options)) (*ListFindingsOutput, error) 
 ListRulesPackages(ctx context.Context, params *ListRulesPackagesInput, optFns ...func(*Options)) (*ListRulesPackagesOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 PreviewAgents(ctx context.Context, params *PreviewAgentsInput, optFns ...func(*Options)) (*PreviewAgentsOutput, error) 
 RegisterCrossAccountAccessRole(ctx context.Context, params *RegisterCrossAccountAccessRoleInput, optFns ...func(*Options)) (*RegisterCrossAccountAccessRoleOutput, error) 
 RemoveAttributesFromFindings(ctx context.Context, params *RemoveAttributesFromFindingsInput, optFns ...func(*Options)) (*RemoveAttributesFromFindingsOutput, error) 
 SetTagsForResource(ctx context.Context, params *SetTagsForResourceInput, optFns ...func(*Options)) (*SetTagsForResourceOutput, error) 
 StartAssessmentRun(ctx context.Context, params *StartAssessmentRunInput, optFns ...func(*Options)) (*StartAssessmentRunOutput, error) 
 StopAssessmentRun(ctx context.Context, params *StopAssessmentRunInput, optFns ...func(*Options)) (*StopAssessmentRunOutput, error) 
 SubscribeToEvent(ctx context.Context, params *SubscribeToEventInput, optFns ...func(*Options)) (*SubscribeToEventOutput, error) 
 UnsubscribeFromEvent(ctx context.Context, params *UnsubscribeFromEventInput, optFns ...func(*Options)) (*UnsubscribeFromEventOutput, error) 
 UpdateAssessmentTarget(ctx context.Context, params *UpdateAssessmentTargetInput, optFns ...func(*Options)) (*UpdateAssessmentTargetOutput, error) 
}
