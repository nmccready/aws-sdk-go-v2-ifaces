
package deadline

// This file was generated by scripts/gen_ifaces.sh do not edit it manually
import (
    "github.com/aws/aws-sdk-go-v2/service/deadline"
)

// IClient defines the interface for deadline
type IClient interface {
 Options() Options 
 AssociateMemberToFarm(ctx context.Context, params *AssociateMemberToFarmInput, optFns ...func(*Options)) (*AssociateMemberToFarmOutput, error) 
 AssociateMemberToFleet(ctx context.Context, params *AssociateMemberToFleetInput, optFns ...func(*Options)) (*AssociateMemberToFleetOutput, error) 
 AssociateMemberToJob(ctx context.Context, params *AssociateMemberToJobInput, optFns ...func(*Options)) (*AssociateMemberToJobOutput, error) 
 AssociateMemberToQueue(ctx context.Context, params *AssociateMemberToQueueInput, optFns ...func(*Options)) (*AssociateMemberToQueueOutput, error) 
 AssumeFleetRoleForRead(ctx context.Context, params *AssumeFleetRoleForReadInput, optFns ...func(*Options)) (*AssumeFleetRoleForReadOutput, error) 
 AssumeFleetRoleForWorker(ctx context.Context, params *AssumeFleetRoleForWorkerInput, optFns ...func(*Options)) (*AssumeFleetRoleForWorkerOutput, error) 
 AssumeQueueRoleForRead(ctx context.Context, params *AssumeQueueRoleForReadInput, optFns ...func(*Options)) (*AssumeQueueRoleForReadOutput, error) 
 AssumeQueueRoleForUser(ctx context.Context, params *AssumeQueueRoleForUserInput, optFns ...func(*Options)) (*AssumeQueueRoleForUserOutput, error) 
 AssumeQueueRoleForWorker(ctx context.Context, params *AssumeQueueRoleForWorkerInput, optFns ...func(*Options)) (*AssumeQueueRoleForWorkerOutput, error) 
 BatchGetJobEntity(ctx context.Context, params *BatchGetJobEntityInput, optFns ...func(*Options)) (*BatchGetJobEntityOutput, error) 
 CopyJobTemplate(ctx context.Context, params *CopyJobTemplateInput, optFns ...func(*Options)) (*CopyJobTemplateOutput, error) 
 CreateBudget(ctx context.Context, params *CreateBudgetInput, optFns ...func(*Options)) (*CreateBudgetOutput, error) 
 CreateFarm(ctx context.Context, params *CreateFarmInput, optFns ...func(*Options)) (*CreateFarmOutput, error) 
 CreateFleet(ctx context.Context, params *CreateFleetInput, optFns ...func(*Options)) (*CreateFleetOutput, error) 
 CreateJob(ctx context.Context, params *CreateJobInput, optFns ...func(*Options)) (*CreateJobOutput, error) 
 CreateLicenseEndpoint(ctx context.Context, params *CreateLicenseEndpointInput, optFns ...func(*Options)) (*CreateLicenseEndpointOutput, error) 
 CreateMonitor(ctx context.Context, params *CreateMonitorInput, optFns ...func(*Options)) (*CreateMonitorOutput, error) 
 CreateQueue(ctx context.Context, params *CreateQueueInput, optFns ...func(*Options)) (*CreateQueueOutput, error) 
 CreateQueueEnvironment(ctx context.Context, params *CreateQueueEnvironmentInput, optFns ...func(*Options)) (*CreateQueueEnvironmentOutput, error) 
 CreateQueueFleetAssociation(ctx context.Context, params *CreateQueueFleetAssociationInput, optFns ...func(*Options)) (*CreateQueueFleetAssociationOutput, error) 
 CreateStorageProfile(ctx context.Context, params *CreateStorageProfileInput, optFns ...func(*Options)) (*CreateStorageProfileOutput, error) 
 CreateWorker(ctx context.Context, params *CreateWorkerInput, optFns ...func(*Options)) (*CreateWorkerOutput, error) 
 DeleteBudget(ctx context.Context, params *DeleteBudgetInput, optFns ...func(*Options)) (*DeleteBudgetOutput, error) 
 DeleteFarm(ctx context.Context, params *DeleteFarmInput, optFns ...func(*Options)) (*DeleteFarmOutput, error) 
 DeleteFleet(ctx context.Context, params *DeleteFleetInput, optFns ...func(*Options)) (*DeleteFleetOutput, error) 
 DeleteLicenseEndpoint(ctx context.Context, params *DeleteLicenseEndpointInput, optFns ...func(*Options)) (*DeleteLicenseEndpointOutput, error) 
 DeleteMeteredProduct(ctx context.Context, params *DeleteMeteredProductInput, optFns ...func(*Options)) (*DeleteMeteredProductOutput, error) 
 DeleteMonitor(ctx context.Context, params *DeleteMonitorInput, optFns ...func(*Options)) (*DeleteMonitorOutput, error) 
 DeleteQueue(ctx context.Context, params *DeleteQueueInput, optFns ...func(*Options)) (*DeleteQueueOutput, error) 
 DeleteQueueEnvironment(ctx context.Context, params *DeleteQueueEnvironmentInput, optFns ...func(*Options)) (*DeleteQueueEnvironmentOutput, error) 
 DeleteQueueFleetAssociation(ctx context.Context, params *DeleteQueueFleetAssociationInput, optFns ...func(*Options)) (*DeleteQueueFleetAssociationOutput, error) 
 DeleteStorageProfile(ctx context.Context, params *DeleteStorageProfileInput, optFns ...func(*Options)) (*DeleteStorageProfileOutput, error) 
 DeleteWorker(ctx context.Context, params *DeleteWorkerInput, optFns ...func(*Options)) (*DeleteWorkerOutput, error) 
 DisassociateMemberFromFarm(ctx context.Context, params *DisassociateMemberFromFarmInput, optFns ...func(*Options)) (*DisassociateMemberFromFarmOutput, error) 
 DisassociateMemberFromFleet(ctx context.Context, params *DisassociateMemberFromFleetInput, optFns ...func(*Options)) (*DisassociateMemberFromFleetOutput, error) 
 DisassociateMemberFromJob(ctx context.Context, params *DisassociateMemberFromJobInput, optFns ...func(*Options)) (*DisassociateMemberFromJobOutput, error) 
 DisassociateMemberFromQueue(ctx context.Context, params *DisassociateMemberFromQueueInput, optFns ...func(*Options)) (*DisassociateMemberFromQueueOutput, error) 
 GetBudget(ctx context.Context, params *GetBudgetInput, optFns ...func(*Options)) (*GetBudgetOutput, error) 
 GetFarm(ctx context.Context, params *GetFarmInput, optFns ...func(*Options)) (*GetFarmOutput, error) 
 GetFleet(ctx context.Context, params *GetFleetInput, optFns ...func(*Options)) (*GetFleetOutput, error) 
 GetJob(ctx context.Context, params *GetJobInput, optFns ...func(*Options)) (*GetJobOutput, error) 
 GetLicenseEndpoint(ctx context.Context, params *GetLicenseEndpointInput, optFns ...func(*Options)) (*GetLicenseEndpointOutput, error) 
 GetMonitor(ctx context.Context, params *GetMonitorInput, optFns ...func(*Options)) (*GetMonitorOutput, error) 
 GetQueue(ctx context.Context, params *GetQueueInput, optFns ...func(*Options)) (*GetQueueOutput, error) 
 GetQueueEnvironment(ctx context.Context, params *GetQueueEnvironmentInput, optFns ...func(*Options)) (*GetQueueEnvironmentOutput, error) 
 GetQueueFleetAssociation(ctx context.Context, params *GetQueueFleetAssociationInput, optFns ...func(*Options)) (*GetQueueFleetAssociationOutput, error) 
 GetSession(ctx context.Context, params *GetSessionInput, optFns ...func(*Options)) (*GetSessionOutput, error) 
 GetSessionAction(ctx context.Context, params *GetSessionActionInput, optFns ...func(*Options)) (*GetSessionActionOutput, error) 
 GetSessionsStatisticsAggregation(ctx context.Context, params *GetSessionsStatisticsAggregationInput, optFns ...func(*Options)) (*GetSessionsStatisticsAggregationOutput, error) 
 GetStep(ctx context.Context, params *GetStepInput, optFns ...func(*Options)) (*GetStepOutput, error) 
 GetStorageProfile(ctx context.Context, params *GetStorageProfileInput, optFns ...func(*Options)) (*GetStorageProfileOutput, error) 
 GetStorageProfileForQueue(ctx context.Context, params *GetStorageProfileForQueueInput, optFns ...func(*Options)) (*GetStorageProfileForQueueOutput, error) 
 GetTask(ctx context.Context, params *GetTaskInput, optFns ...func(*Options)) (*GetTaskOutput, error) 
 GetWorker(ctx context.Context, params *GetWorkerInput, optFns ...func(*Options)) (*GetWorkerOutput, error) 
 ListAvailableMeteredProducts(ctx context.Context, params *ListAvailableMeteredProductsInput, optFns ...func(*Options)) (*ListAvailableMeteredProductsOutput, error) 
 ListBudgets(ctx context.Context, params *ListBudgetsInput, optFns ...func(*Options)) (*ListBudgetsOutput, error) 
 ListFarmMembers(ctx context.Context, params *ListFarmMembersInput, optFns ...func(*Options)) (*ListFarmMembersOutput, error) 
 ListFarms(ctx context.Context, params *ListFarmsInput, optFns ...func(*Options)) (*ListFarmsOutput, error) 
 ListFleetMembers(ctx context.Context, params *ListFleetMembersInput, optFns ...func(*Options)) (*ListFleetMembersOutput, error) 
 ListFleets(ctx context.Context, params *ListFleetsInput, optFns ...func(*Options)) (*ListFleetsOutput, error) 
 ListJobMembers(ctx context.Context, params *ListJobMembersInput, optFns ...func(*Options)) (*ListJobMembersOutput, error) 
 ListJobs(ctx context.Context, params *ListJobsInput, optFns ...func(*Options)) (*ListJobsOutput, error) 
 ListLicenseEndpoints(ctx context.Context, params *ListLicenseEndpointsInput, optFns ...func(*Options)) (*ListLicenseEndpointsOutput, error) 
 ListMeteredProducts(ctx context.Context, params *ListMeteredProductsInput, optFns ...func(*Options)) (*ListMeteredProductsOutput, error) 
 ListMonitors(ctx context.Context, params *ListMonitorsInput, optFns ...func(*Options)) (*ListMonitorsOutput, error) 
 ListQueueEnvironments(ctx context.Context, params *ListQueueEnvironmentsInput, optFns ...func(*Options)) (*ListQueueEnvironmentsOutput, error) 
 ListQueueFleetAssociations(ctx context.Context, params *ListQueueFleetAssociationsInput, optFns ...func(*Options)) (*ListQueueFleetAssociationsOutput, error) 
 ListQueueMembers(ctx context.Context, params *ListQueueMembersInput, optFns ...func(*Options)) (*ListQueueMembersOutput, error) 
 ListQueues(ctx context.Context, params *ListQueuesInput, optFns ...func(*Options)) (*ListQueuesOutput, error) 
 ListSessionActions(ctx context.Context, params *ListSessionActionsInput, optFns ...func(*Options)) (*ListSessionActionsOutput, error) 
 ListSessions(ctx context.Context, params *ListSessionsInput, optFns ...func(*Options)) (*ListSessionsOutput, error) 
 ListSessionsForWorker(ctx context.Context, params *ListSessionsForWorkerInput, optFns ...func(*Options)) (*ListSessionsForWorkerOutput, error) 
 ListStepConsumers(ctx context.Context, params *ListStepConsumersInput, optFns ...func(*Options)) (*ListStepConsumersOutput, error) 
 ListStepDependencies(ctx context.Context, params *ListStepDependenciesInput, optFns ...func(*Options)) (*ListStepDependenciesOutput, error) 
 ListSteps(ctx context.Context, params *ListStepsInput, optFns ...func(*Options)) (*ListStepsOutput, error) 
 ListStorageProfiles(ctx context.Context, params *ListStorageProfilesInput, optFns ...func(*Options)) (*ListStorageProfilesOutput, error) 
 ListStorageProfilesForQueue(ctx context.Context, params *ListStorageProfilesForQueueInput, optFns ...func(*Options)) (*ListStorageProfilesForQueueOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 ListTasks(ctx context.Context, params *ListTasksInput, optFns ...func(*Options)) (*ListTasksOutput, error) 
 ListWorkers(ctx context.Context, params *ListWorkersInput, optFns ...func(*Options)) (*ListWorkersOutput, error) 
 PutMeteredProduct(ctx context.Context, params *PutMeteredProductInput, optFns ...func(*Options)) (*PutMeteredProductOutput, error) 
 SearchJobs(ctx context.Context, params *SearchJobsInput, optFns ...func(*Options)) (*SearchJobsOutput, error) 
 SearchSteps(ctx context.Context, params *SearchStepsInput, optFns ...func(*Options)) (*SearchStepsOutput, error) 
 SearchTasks(ctx context.Context, params *SearchTasksInput, optFns ...func(*Options)) (*SearchTasksOutput, error) 
 SearchWorkers(ctx context.Context, params *SearchWorkersInput, optFns ...func(*Options)) (*SearchWorkersOutput, error) 
 StartSessionsStatisticsAggregation(ctx context.Context, params *StartSessionsStatisticsAggregationInput, optFns ...func(*Options)) (*StartSessionsStatisticsAggregationOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateBudget(ctx context.Context, params *UpdateBudgetInput, optFns ...func(*Options)) (*UpdateBudgetOutput, error) 
 UpdateFarm(ctx context.Context, params *UpdateFarmInput, optFns ...func(*Options)) (*UpdateFarmOutput, error) 
 UpdateFleet(ctx context.Context, params *UpdateFleetInput, optFns ...func(*Options)) (*UpdateFleetOutput, error) 
 UpdateJob(ctx context.Context, params *UpdateJobInput, optFns ...func(*Options)) (*UpdateJobOutput, error) 
 UpdateMonitor(ctx context.Context, params *UpdateMonitorInput, optFns ...func(*Options)) (*UpdateMonitorOutput, error) 
 UpdateQueue(ctx context.Context, params *UpdateQueueInput, optFns ...func(*Options)) (*UpdateQueueOutput, error) 
 UpdateQueueEnvironment(ctx context.Context, params *UpdateQueueEnvironmentInput, optFns ...func(*Options)) (*UpdateQueueEnvironmentOutput, error) 
 UpdateQueueFleetAssociation(ctx context.Context, params *UpdateQueueFleetAssociationInput, optFns ...func(*Options)) (*UpdateQueueFleetAssociationOutput, error) 
 UpdateSession(ctx context.Context, params *UpdateSessionInput, optFns ...func(*Options)) (*UpdateSessionOutput, error) 
 UpdateStep(ctx context.Context, params *UpdateStepInput, optFns ...func(*Options)) (*UpdateStepOutput, error) 
 UpdateStorageProfile(ctx context.Context, params *UpdateStorageProfileInput, optFns ...func(*Options)) (*UpdateStorageProfileOutput, error) 
 UpdateTask(ctx context.Context, params *UpdateTaskInput, optFns ...func(*Options)) (*UpdateTaskOutput, error) 
 UpdateWorker(ctx context.Context, params *UpdateWorkerInput, optFns ...func(*Options)) (*UpdateWorkerOutput, error) 
 UpdateWorkerSchedule(ctx context.Context, params *UpdateWorkerScheduleInput, optFns ...func(*Options)) (*UpdateWorkerScheduleOutput, error) 
}
