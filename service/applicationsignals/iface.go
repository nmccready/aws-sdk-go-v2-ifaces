
package applicationsignals

import (
    "github.com/aws/aws-sdk-go-v2/service/applicationsignals"
)

// IApplicationsignals defines the interface for applicationsignals
type IApplicationsignals interface {
 Options() Options 
 BatchGetServiceLevelObjectiveBudgetReport(ctx context.Context, params *BatchGetServiceLevelObjectiveBudgetReportInput, optFns ...func(*Options)) (*BatchGetServiceLevelObjectiveBudgetReportOutput, error) 
 CreateServiceLevelObjective(ctx context.Context, params *CreateServiceLevelObjectiveInput, optFns ...func(*Options)) (*CreateServiceLevelObjectiveOutput, error) 
 DeleteServiceLevelObjective(ctx context.Context, params *DeleteServiceLevelObjectiveInput, optFns ...func(*Options)) (*DeleteServiceLevelObjectiveOutput, error) 
 GetService(ctx context.Context, params *GetServiceInput, optFns ...func(*Options)) (*GetServiceOutput, error) 
 GetServiceLevelObjective(ctx context.Context, params *GetServiceLevelObjectiveInput, optFns ...func(*Options)) (*GetServiceLevelObjectiveOutput, error) 
 ListServiceDependencies(ctx context.Context, params *ListServiceDependenciesInput, optFns ...func(*Options)) (*ListServiceDependenciesOutput, error) 
 ListServiceDependents(ctx context.Context, params *ListServiceDependentsInput, optFns ...func(*Options)) (*ListServiceDependentsOutput, error) 
 ListServiceLevelObjectives(ctx context.Context, params *ListServiceLevelObjectivesInput, optFns ...func(*Options)) (*ListServiceLevelObjectivesOutput, error) 
 ListServiceOperations(ctx context.Context, params *ListServiceOperationsInput, optFns ...func(*Options)) (*ListServiceOperationsOutput, error) 
 ListServices(ctx context.Context, params *ListServicesInput, optFns ...func(*Options)) (*ListServicesOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 StartDiscovery(ctx context.Context, params *StartDiscoveryInput, optFns ...func(*Options)) (*StartDiscoveryOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateServiceLevelObjective(ctx context.Context, params *UpdateServiceLevelObjectiveInput, optFns ...func(*Options)) (*UpdateServiceLevelObjectiveOutput, error) 
}
