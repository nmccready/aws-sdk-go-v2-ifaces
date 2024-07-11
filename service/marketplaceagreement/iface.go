
package marketplaceagreement

import (
    "github.com/aws/aws-sdk-go-v2/service/marketplaceagreement"
)

// IMarketplaceagreement defines the interface for marketplaceagreement
type IMarketplaceagreement interface {
 Options() Options 
 DescribeAgreement(ctx context.Context, params *DescribeAgreementInput, optFns ...func(*Options)) (*DescribeAgreementOutput, error) 
 GetAgreementTerms(ctx context.Context, params *GetAgreementTermsInput, optFns ...func(*Options)) (*GetAgreementTermsOutput, error) 
 SearchAgreements(ctx context.Context, params *SearchAgreementsInput, optFns ...func(*Options)) (*SearchAgreementsOutput, error) 
}
