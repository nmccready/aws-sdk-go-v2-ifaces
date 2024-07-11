
package support

// This file was generated by scripts/gen_ifaces.sh do not edit it manually
import (
    "github.com/aws/aws-sdk-go-v2/service/support"
)

// IClient defines the interface for support
type IClient interface {
 Options() Options 
 AddAttachmentsToSet(ctx context.Context, params *AddAttachmentsToSetInput, optFns ...func(*Options)) (*AddAttachmentsToSetOutput, error) 
 AddCommunicationToCase(ctx context.Context, params *AddCommunicationToCaseInput, optFns ...func(*Options)) (*AddCommunicationToCaseOutput, error) 
 CreateCase(ctx context.Context, params *CreateCaseInput, optFns ...func(*Options)) (*CreateCaseOutput, error) 
 DescribeAttachment(ctx context.Context, params *DescribeAttachmentInput, optFns ...func(*Options)) (*DescribeAttachmentOutput, error) 
 DescribeCases(ctx context.Context, params *DescribeCasesInput, optFns ...func(*Options)) (*DescribeCasesOutput, error) 
 DescribeCommunications(ctx context.Context, params *DescribeCommunicationsInput, optFns ...func(*Options)) (*DescribeCommunicationsOutput, error) 
 DescribeCreateCaseOptions(ctx context.Context, params *DescribeCreateCaseOptionsInput, optFns ...func(*Options)) (*DescribeCreateCaseOptionsOutput, error) 
 DescribeServices(ctx context.Context, params *DescribeServicesInput, optFns ...func(*Options)) (*DescribeServicesOutput, error) 
 DescribeSeverityLevels(ctx context.Context, params *DescribeSeverityLevelsInput, optFns ...func(*Options)) (*DescribeSeverityLevelsOutput, error) 
 DescribeSupportedLanguages(ctx context.Context, params *DescribeSupportedLanguagesInput, optFns ...func(*Options)) (*DescribeSupportedLanguagesOutput, error) 
 DescribeTrustedAdvisorCheckRefreshStatuses(ctx context.Context, params *DescribeTrustedAdvisorCheckRefreshStatusesInput, optFns ...func(*Options)) (*DescribeTrustedAdvisorCheckRefreshStatusesOutput, error) 
 DescribeTrustedAdvisorCheckResult(ctx context.Context, params *DescribeTrustedAdvisorCheckResultInput, optFns ...func(*Options)) (*DescribeTrustedAdvisorCheckResultOutput, error) 
 DescribeTrustedAdvisorCheckSummaries(ctx context.Context, params *DescribeTrustedAdvisorCheckSummariesInput, optFns ...func(*Options)) (*DescribeTrustedAdvisorCheckSummariesOutput, error) 
 DescribeTrustedAdvisorChecks(ctx context.Context, params *DescribeTrustedAdvisorChecksInput, optFns ...func(*Options)) (*DescribeTrustedAdvisorChecksOutput, error) 
 RefreshTrustedAdvisorCheck(ctx context.Context, params *RefreshTrustedAdvisorCheckInput, optFns ...func(*Options)) (*RefreshTrustedAdvisorCheckOutput, error) 
 ResolveCase(ctx context.Context, params *ResolveCaseInput, optFns ...func(*Options)) (*ResolveCaseOutput, error) 
}
