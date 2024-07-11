
package accessanalyzer

import (
    "github.com/aws/aws-sdk-go-v2/service/accessanalyzer"
)

// IClient defines the interface for accessanalyzer
type IClient interface {
 Options() Options 
 ApplyArchiveRule(ctx context.Context, params *ApplyArchiveRuleInput, optFns ...func(*Options)) (*ApplyArchiveRuleOutput, error) 
 CancelPolicyGeneration(ctx context.Context, params *CancelPolicyGenerationInput, optFns ...func(*Options)) (*CancelPolicyGenerationOutput, error) 
 CheckAccessNotGranted(ctx context.Context, params *CheckAccessNotGrantedInput, optFns ...func(*Options)) (*CheckAccessNotGrantedOutput, error) 
 CheckNoNewAccess(ctx context.Context, params *CheckNoNewAccessInput, optFns ...func(*Options)) (*CheckNoNewAccessOutput, error) 
 CheckNoPublicAccess(ctx context.Context, params *CheckNoPublicAccessInput, optFns ...func(*Options)) (*CheckNoPublicAccessOutput, error) 
 CreateAccessPreview(ctx context.Context, params *CreateAccessPreviewInput, optFns ...func(*Options)) (*CreateAccessPreviewOutput, error) 
 CreateAnalyzer(ctx context.Context, params *CreateAnalyzerInput, optFns ...func(*Options)) (*CreateAnalyzerOutput, error) 
 CreateArchiveRule(ctx context.Context, params *CreateArchiveRuleInput, optFns ...func(*Options)) (*CreateArchiveRuleOutput, error) 
 DeleteAnalyzer(ctx context.Context, params *DeleteAnalyzerInput, optFns ...func(*Options)) (*DeleteAnalyzerOutput, error) 
 DeleteArchiveRule(ctx context.Context, params *DeleteArchiveRuleInput, optFns ...func(*Options)) (*DeleteArchiveRuleOutput, error) 
 GenerateFindingRecommendation(ctx context.Context, params *GenerateFindingRecommendationInput, optFns ...func(*Options)) (*GenerateFindingRecommendationOutput, error) 
 GetAccessPreview(ctx context.Context, params *GetAccessPreviewInput, optFns ...func(*Options)) (*GetAccessPreviewOutput, error) 
 GetAnalyzedResource(ctx context.Context, params *GetAnalyzedResourceInput, optFns ...func(*Options)) (*GetAnalyzedResourceOutput, error) 
 GetAnalyzer(ctx context.Context, params *GetAnalyzerInput, optFns ...func(*Options)) (*GetAnalyzerOutput, error) 
 GetArchiveRule(ctx context.Context, params *GetArchiveRuleInput, optFns ...func(*Options)) (*GetArchiveRuleOutput, error) 
 GetFinding(ctx context.Context, params *GetFindingInput, optFns ...func(*Options)) (*GetFindingOutput, error) 
 GetFindingRecommendation(ctx context.Context, params *GetFindingRecommendationInput, optFns ...func(*Options)) (*GetFindingRecommendationOutput, error) 
 GetFindingV2(ctx context.Context, params *GetFindingV2Input, optFns ...func(*Options)) (*GetFindingV2Output, error) 
 GetGeneratedPolicy(ctx context.Context, params *GetGeneratedPolicyInput, optFns ...func(*Options)) (*GetGeneratedPolicyOutput, error) 
 ListAccessPreviewFindings(ctx context.Context, params *ListAccessPreviewFindingsInput, optFns ...func(*Options)) (*ListAccessPreviewFindingsOutput, error) 
 ListAccessPreviews(ctx context.Context, params *ListAccessPreviewsInput, optFns ...func(*Options)) (*ListAccessPreviewsOutput, error) 
 ListAnalyzedResources(ctx context.Context, params *ListAnalyzedResourcesInput, optFns ...func(*Options)) (*ListAnalyzedResourcesOutput, error) 
 ListAnalyzers(ctx context.Context, params *ListAnalyzersInput, optFns ...func(*Options)) (*ListAnalyzersOutput, error) 
 ListArchiveRules(ctx context.Context, params *ListArchiveRulesInput, optFns ...func(*Options)) (*ListArchiveRulesOutput, error) 
 ListFindings(ctx context.Context, params *ListFindingsInput, optFns ...func(*Options)) (*ListFindingsOutput, error) 
 ListFindingsV2(ctx context.Context, params *ListFindingsV2Input, optFns ...func(*Options)) (*ListFindingsV2Output, error) 
 ListPolicyGenerations(ctx context.Context, params *ListPolicyGenerationsInput, optFns ...func(*Options)) (*ListPolicyGenerationsOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 StartPolicyGeneration(ctx context.Context, params *StartPolicyGenerationInput, optFns ...func(*Options)) (*StartPolicyGenerationOutput, error) 
 StartResourceScan(ctx context.Context, params *StartResourceScanInput, optFns ...func(*Options)) (*StartResourceScanOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateArchiveRule(ctx context.Context, params *UpdateArchiveRuleInput, optFns ...func(*Options)) (*UpdateArchiveRuleOutput, error) 
 UpdateFindings(ctx context.Context, params *UpdateFindingsInput, optFns ...func(*Options)) (*UpdateFindingsOutput, error) 
 ValidatePolicy(ctx context.Context, params *ValidatePolicyInput, optFns ...func(*Options)) (*ValidatePolicyOutput, error) 
}
