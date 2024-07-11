
package cloudsearch

import (
    "github.com/aws/aws-sdk-go-v2/service/cloudsearch"
)

// IClient defines the interface for cloudsearch
type IClient interface {
 Options() Options 
 BuildSuggesters(ctx context.Context, params *BuildSuggestersInput, optFns ...func(*Options)) (*BuildSuggestersOutput, error) 
 CreateDomain(ctx context.Context, params *CreateDomainInput, optFns ...func(*Options)) (*CreateDomainOutput, error) 
 DefineAnalysisScheme(ctx context.Context, params *DefineAnalysisSchemeInput, optFns ...func(*Options)) (*DefineAnalysisSchemeOutput, error) 
 DefineExpression(ctx context.Context, params *DefineExpressionInput, optFns ...func(*Options)) (*DefineExpressionOutput, error) 
 DefineIndexField(ctx context.Context, params *DefineIndexFieldInput, optFns ...func(*Options)) (*DefineIndexFieldOutput, error) 
 DefineSuggester(ctx context.Context, params *DefineSuggesterInput, optFns ...func(*Options)) (*DefineSuggesterOutput, error) 
 DeleteAnalysisScheme(ctx context.Context, params *DeleteAnalysisSchemeInput, optFns ...func(*Options)) (*DeleteAnalysisSchemeOutput, error) 
 DeleteDomain(ctx context.Context, params *DeleteDomainInput, optFns ...func(*Options)) (*DeleteDomainOutput, error) 
 DeleteExpression(ctx context.Context, params *DeleteExpressionInput, optFns ...func(*Options)) (*DeleteExpressionOutput, error) 
 DeleteIndexField(ctx context.Context, params *DeleteIndexFieldInput, optFns ...func(*Options)) (*DeleteIndexFieldOutput, error) 
 DeleteSuggester(ctx context.Context, params *DeleteSuggesterInput, optFns ...func(*Options)) (*DeleteSuggesterOutput, error) 
 DescribeAnalysisSchemes(ctx context.Context, params *DescribeAnalysisSchemesInput, optFns ...func(*Options)) (*DescribeAnalysisSchemesOutput, error) 
 DescribeAvailabilityOptions(ctx context.Context, params *DescribeAvailabilityOptionsInput, optFns ...func(*Options)) (*DescribeAvailabilityOptionsOutput, error) 
 DescribeDomainEndpointOptions(ctx context.Context, params *DescribeDomainEndpointOptionsInput, optFns ...func(*Options)) (*DescribeDomainEndpointOptionsOutput, error) 
 DescribeDomains(ctx context.Context, params *DescribeDomainsInput, optFns ...func(*Options)) (*DescribeDomainsOutput, error) 
 DescribeExpressions(ctx context.Context, params *DescribeExpressionsInput, optFns ...func(*Options)) (*DescribeExpressionsOutput, error) 
 DescribeIndexFields(ctx context.Context, params *DescribeIndexFieldsInput, optFns ...func(*Options)) (*DescribeIndexFieldsOutput, error) 
 DescribeScalingParameters(ctx context.Context, params *DescribeScalingParametersInput, optFns ...func(*Options)) (*DescribeScalingParametersOutput, error) 
 DescribeServiceAccessPolicies(ctx context.Context, params *DescribeServiceAccessPoliciesInput, optFns ...func(*Options)) (*DescribeServiceAccessPoliciesOutput, error) 
 DescribeSuggesters(ctx context.Context, params *DescribeSuggestersInput, optFns ...func(*Options)) (*DescribeSuggestersOutput, error) 
 IndexDocuments(ctx context.Context, params *IndexDocumentsInput, optFns ...func(*Options)) (*IndexDocumentsOutput, error) 
 ListDomainNames(ctx context.Context, params *ListDomainNamesInput, optFns ...func(*Options)) (*ListDomainNamesOutput, error) 
 UpdateAvailabilityOptions(ctx context.Context, params *UpdateAvailabilityOptionsInput, optFns ...func(*Options)) (*UpdateAvailabilityOptionsOutput, error) 
 UpdateDomainEndpointOptions(ctx context.Context, params *UpdateDomainEndpointOptionsInput, optFns ...func(*Options)) (*UpdateDomainEndpointOptionsOutput, error) 
 UpdateScalingParameters(ctx context.Context, params *UpdateScalingParametersInput, optFns ...func(*Options)) (*UpdateScalingParametersOutput, error) 
 UpdateServiceAccessPolicies(ctx context.Context, params *UpdateServiceAccessPoliciesInput, optFns ...func(*Options)) (*UpdateServiceAccessPoliciesOutput, error) 
}
