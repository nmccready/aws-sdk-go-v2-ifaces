
package health

import (
    "github.com/aws/aws-sdk-go-v2/service/health"
)

// IHealth defines the interface for health
type IHealth interface {
 Options() Options 
 DescribeAffectedAccountsForOrganization(ctx context.Context, params *DescribeAffectedAccountsForOrganizationInput, optFns ...func(*Options)) (*DescribeAffectedAccountsForOrganizationOutput, error) 
 DescribeAffectedEntities(ctx context.Context, params *DescribeAffectedEntitiesInput, optFns ...func(*Options)) (*DescribeAffectedEntitiesOutput, error) 
 DescribeAffectedEntitiesForOrganization(ctx context.Context, params *DescribeAffectedEntitiesForOrganizationInput, optFns ...func(*Options)) (*DescribeAffectedEntitiesForOrganizationOutput, error) 
 DescribeEntityAggregates(ctx context.Context, params *DescribeEntityAggregatesInput, optFns ...func(*Options)) (*DescribeEntityAggregatesOutput, error) 
 DescribeEntityAggregatesForOrganization(ctx context.Context, params *DescribeEntityAggregatesForOrganizationInput, optFns ...func(*Options)) (*DescribeEntityAggregatesForOrganizationOutput, error) 
 DescribeEventAggregates(ctx context.Context, params *DescribeEventAggregatesInput, optFns ...func(*Options)) (*DescribeEventAggregatesOutput, error) 
 DescribeEventDetails(ctx context.Context, params *DescribeEventDetailsInput, optFns ...func(*Options)) (*DescribeEventDetailsOutput, error) 
 DescribeEventDetailsForOrganization(ctx context.Context, params *DescribeEventDetailsForOrganizationInput, optFns ...func(*Options)) (*DescribeEventDetailsForOrganizationOutput, error) 
 DescribeEventTypes(ctx context.Context, params *DescribeEventTypesInput, optFns ...func(*Options)) (*DescribeEventTypesOutput, error) 
 DescribeEvents(ctx context.Context, params *DescribeEventsInput, optFns ...func(*Options)) (*DescribeEventsOutput, error) 
 DescribeEventsForOrganization(ctx context.Context, params *DescribeEventsForOrganizationInput, optFns ...func(*Options)) (*DescribeEventsForOrganizationOutput, error) 
 DescribeHealthServiceStatusForOrganization(ctx context.Context, params *DescribeHealthServiceStatusForOrganizationInput, optFns ...func(*Options)) (*DescribeHealthServiceStatusForOrganizationOutput, error) 
 DisableHealthServiceAccessForOrganization(ctx context.Context, params *DisableHealthServiceAccessForOrganizationInput, optFns ...func(*Options)) (*DisableHealthServiceAccessForOrganizationOutput, error) 
 EnableHealthServiceAccessForOrganization(ctx context.Context, params *EnableHealthServiceAccessForOrganizationInput, optFns ...func(*Options)) (*EnableHealthServiceAccessForOrganizationOutput, error) 
}
