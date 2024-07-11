
package timestreamquery

import (
    "github.com/aws/aws-sdk-go-v2/service/timestreamquery"
)

// IClient defines the interface for timestreamquery
type IClient interface {
 Options() Options 
 CancelQuery(ctx context.Context, params *CancelQueryInput, optFns ...func(*Options)) (*CancelQueryOutput, error) 
 CreateScheduledQuery(ctx context.Context, params *CreateScheduledQueryInput, optFns ...func(*Options)) (*CreateScheduledQueryOutput, error) 
 DeleteScheduledQuery(ctx context.Context, params *DeleteScheduledQueryInput, optFns ...func(*Options)) (*DeleteScheduledQueryOutput, error) 
 DescribeAccountSettings(ctx context.Context, params *DescribeAccountSettingsInput, optFns ...func(*Options)) (*DescribeAccountSettingsOutput, error) 
 DescribeEndpoints(ctx context.Context, params *DescribeEndpointsInput, optFns ...func(*Options)) (*DescribeEndpointsOutput, error) 
 DescribeScheduledQuery(ctx context.Context, params *DescribeScheduledQueryInput, optFns ...func(*Options)) (*DescribeScheduledQueryOutput, error) 
 ExecuteScheduledQuery(ctx context.Context, params *ExecuteScheduledQueryInput, optFns ...func(*Options)) (*ExecuteScheduledQueryOutput, error) 
 ListScheduledQueries(ctx context.Context, params *ListScheduledQueriesInput, optFns ...func(*Options)) (*ListScheduledQueriesOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 PrepareQuery(ctx context.Context, params *PrepareQueryInput, optFns ...func(*Options)) (*PrepareQueryOutput, error) 
 Query(ctx context.Context, params *QueryInput, optFns ...func(*Options)) (*QueryOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateAccountSettings(ctx context.Context, params *UpdateAccountSettingsInput, optFns ...func(*Options)) (*UpdateAccountSettingsOutput, error) 
 UpdateScheduledQuery(ctx context.Context, params *UpdateScheduledQueryInput, optFns ...func(*Options)) (*UpdateScheduledQueryOutput, error) 
}
