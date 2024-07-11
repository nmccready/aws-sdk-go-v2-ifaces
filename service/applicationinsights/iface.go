
package applicationinsights

import (
    "github.com/aws/aws-sdk-go-v2/service/applicationinsights"
)

// IClient defines the interface for applicationinsights
type IClient interface {
 Options() Options 
 AddWorkload(ctx context.Context, params *AddWorkloadInput, optFns ...func(*Options)) (*AddWorkloadOutput, error) 
 CreateApplication(ctx context.Context, params *CreateApplicationInput, optFns ...func(*Options)) (*CreateApplicationOutput, error) 
 CreateComponent(ctx context.Context, params *CreateComponentInput, optFns ...func(*Options)) (*CreateComponentOutput, error) 
 CreateLogPattern(ctx context.Context, params *CreateLogPatternInput, optFns ...func(*Options)) (*CreateLogPatternOutput, error) 
 DeleteApplication(ctx context.Context, params *DeleteApplicationInput, optFns ...func(*Options)) (*DeleteApplicationOutput, error) 
 DeleteComponent(ctx context.Context, params *DeleteComponentInput, optFns ...func(*Options)) (*DeleteComponentOutput, error) 
 DeleteLogPattern(ctx context.Context, params *DeleteLogPatternInput, optFns ...func(*Options)) (*DeleteLogPatternOutput, error) 
 DescribeApplication(ctx context.Context, params *DescribeApplicationInput, optFns ...func(*Options)) (*DescribeApplicationOutput, error) 
 DescribeComponent(ctx context.Context, params *DescribeComponentInput, optFns ...func(*Options)) (*DescribeComponentOutput, error) 
 DescribeComponentConfiguration(ctx context.Context, params *DescribeComponentConfigurationInput, optFns ...func(*Options)) (*DescribeComponentConfigurationOutput, error) 
 DescribeComponentConfigurationRecommendation(ctx context.Context, params *DescribeComponentConfigurationRecommendationInput, optFns ...func(*Options)) (*DescribeComponentConfigurationRecommendationOutput, error) 
 DescribeLogPattern(ctx context.Context, params *DescribeLogPatternInput, optFns ...func(*Options)) (*DescribeLogPatternOutput, error) 
 DescribeObservation(ctx context.Context, params *DescribeObservationInput, optFns ...func(*Options)) (*DescribeObservationOutput, error) 
 DescribeProblem(ctx context.Context, params *DescribeProblemInput, optFns ...func(*Options)) (*DescribeProblemOutput, error) 
 DescribeProblemObservations(ctx context.Context, params *DescribeProblemObservationsInput, optFns ...func(*Options)) (*DescribeProblemObservationsOutput, error) 
 DescribeWorkload(ctx context.Context, params *DescribeWorkloadInput, optFns ...func(*Options)) (*DescribeWorkloadOutput, error) 
 ListApplications(ctx context.Context, params *ListApplicationsInput, optFns ...func(*Options)) (*ListApplicationsOutput, error) 
 ListComponents(ctx context.Context, params *ListComponentsInput, optFns ...func(*Options)) (*ListComponentsOutput, error) 
 ListConfigurationHistory(ctx context.Context, params *ListConfigurationHistoryInput, optFns ...func(*Options)) (*ListConfigurationHistoryOutput, error) 
 ListLogPatternSets(ctx context.Context, params *ListLogPatternSetsInput, optFns ...func(*Options)) (*ListLogPatternSetsOutput, error) 
 ListLogPatterns(ctx context.Context, params *ListLogPatternsInput, optFns ...func(*Options)) (*ListLogPatternsOutput, error) 
 ListProblems(ctx context.Context, params *ListProblemsInput, optFns ...func(*Options)) (*ListProblemsOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 ListWorkloads(ctx context.Context, params *ListWorkloadsInput, optFns ...func(*Options)) (*ListWorkloadsOutput, error) 
 RemoveWorkload(ctx context.Context, params *RemoveWorkloadInput, optFns ...func(*Options)) (*RemoveWorkloadOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateApplication(ctx context.Context, params *UpdateApplicationInput, optFns ...func(*Options)) (*UpdateApplicationOutput, error) 
 UpdateComponent(ctx context.Context, params *UpdateComponentInput, optFns ...func(*Options)) (*UpdateComponentOutput, error) 
 UpdateComponentConfiguration(ctx context.Context, params *UpdateComponentConfigurationInput, optFns ...func(*Options)) (*UpdateComponentConfigurationOutput, error) 
 UpdateLogPattern(ctx context.Context, params *UpdateLogPatternInput, optFns ...func(*Options)) (*UpdateLogPatternOutput, error) 
 UpdateProblem(ctx context.Context, params *UpdateProblemInput, optFns ...func(*Options)) (*UpdateProblemOutput, error) 
 UpdateWorkload(ctx context.Context, params *UpdateWorkloadInput, optFns ...func(*Options)) (*UpdateWorkloadOutput, error) 
}
