
package route53recoveryreadiness

import (
    "github.com/aws/aws-sdk-go-v2/service/route53recoveryreadiness"
)

// IClient defines the interface for route53recoveryreadiness
type IClient interface {
 Options() Options 
 CreateCell(ctx context.Context, params *CreateCellInput, optFns ...func(*Options)) (*CreateCellOutput, error) 
 CreateCrossAccountAuthorization(ctx context.Context, params *CreateCrossAccountAuthorizationInput, optFns ...func(*Options)) (*CreateCrossAccountAuthorizationOutput, error) 
 CreateReadinessCheck(ctx context.Context, params *CreateReadinessCheckInput, optFns ...func(*Options)) (*CreateReadinessCheckOutput, error) 
 CreateRecoveryGroup(ctx context.Context, params *CreateRecoveryGroupInput, optFns ...func(*Options)) (*CreateRecoveryGroupOutput, error) 
 CreateResourceSet(ctx context.Context, params *CreateResourceSetInput, optFns ...func(*Options)) (*CreateResourceSetOutput, error) 
 DeleteCell(ctx context.Context, params *DeleteCellInput, optFns ...func(*Options)) (*DeleteCellOutput, error) 
 DeleteCrossAccountAuthorization(ctx context.Context, params *DeleteCrossAccountAuthorizationInput, optFns ...func(*Options)) (*DeleteCrossAccountAuthorizationOutput, error) 
 DeleteReadinessCheck(ctx context.Context, params *DeleteReadinessCheckInput, optFns ...func(*Options)) (*DeleteReadinessCheckOutput, error) 
 DeleteRecoveryGroup(ctx context.Context, params *DeleteRecoveryGroupInput, optFns ...func(*Options)) (*DeleteRecoveryGroupOutput, error) 
 DeleteResourceSet(ctx context.Context, params *DeleteResourceSetInput, optFns ...func(*Options)) (*DeleteResourceSetOutput, error) 
 GetArchitectureRecommendations(ctx context.Context, params *GetArchitectureRecommendationsInput, optFns ...func(*Options)) (*GetArchitectureRecommendationsOutput, error) 
 GetCell(ctx context.Context, params *GetCellInput, optFns ...func(*Options)) (*GetCellOutput, error) 
 GetCellReadinessSummary(ctx context.Context, params *GetCellReadinessSummaryInput, optFns ...func(*Options)) (*GetCellReadinessSummaryOutput, error) 
 GetReadinessCheck(ctx context.Context, params *GetReadinessCheckInput, optFns ...func(*Options)) (*GetReadinessCheckOutput, error) 
 GetReadinessCheckResourceStatus(ctx context.Context, params *GetReadinessCheckResourceStatusInput, optFns ...func(*Options)) (*GetReadinessCheckResourceStatusOutput, error) 
 GetReadinessCheckStatus(ctx context.Context, params *GetReadinessCheckStatusInput, optFns ...func(*Options)) (*GetReadinessCheckStatusOutput, error) 
 GetRecoveryGroup(ctx context.Context, params *GetRecoveryGroupInput, optFns ...func(*Options)) (*GetRecoveryGroupOutput, error) 
 GetRecoveryGroupReadinessSummary(ctx context.Context, params *GetRecoveryGroupReadinessSummaryInput, optFns ...func(*Options)) (*GetRecoveryGroupReadinessSummaryOutput, error) 
 GetResourceSet(ctx context.Context, params *GetResourceSetInput, optFns ...func(*Options)) (*GetResourceSetOutput, error) 
 ListCells(ctx context.Context, params *ListCellsInput, optFns ...func(*Options)) (*ListCellsOutput, error) 
 ListCrossAccountAuthorizations(ctx context.Context, params *ListCrossAccountAuthorizationsInput, optFns ...func(*Options)) (*ListCrossAccountAuthorizationsOutput, error) 
 ListReadinessChecks(ctx context.Context, params *ListReadinessChecksInput, optFns ...func(*Options)) (*ListReadinessChecksOutput, error) 
 ListRecoveryGroups(ctx context.Context, params *ListRecoveryGroupsInput, optFns ...func(*Options)) (*ListRecoveryGroupsOutput, error) 
 ListResourceSets(ctx context.Context, params *ListResourceSetsInput, optFns ...func(*Options)) (*ListResourceSetsOutput, error) 
 ListRules(ctx context.Context, params *ListRulesInput, optFns ...func(*Options)) (*ListRulesOutput, error) 
 ListTagsForResources(ctx context.Context, params *ListTagsForResourcesInput, optFns ...func(*Options)) (*ListTagsForResourcesOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateCell(ctx context.Context, params *UpdateCellInput, optFns ...func(*Options)) (*UpdateCellOutput, error) 
 UpdateReadinessCheck(ctx context.Context, params *UpdateReadinessCheckInput, optFns ...func(*Options)) (*UpdateReadinessCheckOutput, error) 
 UpdateRecoveryGroup(ctx context.Context, params *UpdateRecoveryGroupInput, optFns ...func(*Options)) (*UpdateRecoveryGroupOutput, error) 
 UpdateResourceSet(ctx context.Context, params *UpdateResourceSetInput, optFns ...func(*Options)) (*UpdateResourceSetOutput, error) 
}
