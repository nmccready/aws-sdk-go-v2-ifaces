
package robomaker_iface

// This file was generated by scripts/gen_ifaces.sh do not edit it manually
import (
    "context"
    . "github.com/aws/aws-sdk-go-v2/service/robomaker"
)

// IClient defines the interface for robomaker
type IClient interface {
 Options() Options 
 BatchDeleteWorlds(ctx context.Context, params *BatchDeleteWorldsInput, optFns ...func(*Options)) (*BatchDeleteWorldsOutput, error) 
 BatchDescribeSimulationJob(ctx context.Context, params *BatchDescribeSimulationJobInput, optFns ...func(*Options)) (*BatchDescribeSimulationJobOutput, error) 
 CancelDeploymentJob(ctx context.Context, params *CancelDeploymentJobInput, optFns ...func(*Options)) (*CancelDeploymentJobOutput, error) 
 CancelSimulationJob(ctx context.Context, params *CancelSimulationJobInput, optFns ...func(*Options)) (*CancelSimulationJobOutput, error) 
 CancelSimulationJobBatch(ctx context.Context, params *CancelSimulationJobBatchInput, optFns ...func(*Options)) (*CancelSimulationJobBatchOutput, error) 
 CancelWorldExportJob(ctx context.Context, params *CancelWorldExportJobInput, optFns ...func(*Options)) (*CancelWorldExportJobOutput, error) 
 CancelWorldGenerationJob(ctx context.Context, params *CancelWorldGenerationJobInput, optFns ...func(*Options)) (*CancelWorldGenerationJobOutput, error) 
 CreateDeploymentJob(ctx context.Context, params *CreateDeploymentJobInput, optFns ...func(*Options)) (*CreateDeploymentJobOutput, error) 
 CreateFleet(ctx context.Context, params *CreateFleetInput, optFns ...func(*Options)) (*CreateFleetOutput, error) 
 CreateRobot(ctx context.Context, params *CreateRobotInput, optFns ...func(*Options)) (*CreateRobotOutput, error) 
 CreateRobotApplication(ctx context.Context, params *CreateRobotApplicationInput, optFns ...func(*Options)) (*CreateRobotApplicationOutput, error) 
 CreateRobotApplicationVersion(ctx context.Context, params *CreateRobotApplicationVersionInput, optFns ...func(*Options)) (*CreateRobotApplicationVersionOutput, error) 
 CreateSimulationApplication(ctx context.Context, params *CreateSimulationApplicationInput, optFns ...func(*Options)) (*CreateSimulationApplicationOutput, error) 
 CreateSimulationApplicationVersion(ctx context.Context, params *CreateSimulationApplicationVersionInput, optFns ...func(*Options)) (*CreateSimulationApplicationVersionOutput, error) 
 CreateSimulationJob(ctx context.Context, params *CreateSimulationJobInput, optFns ...func(*Options)) (*CreateSimulationJobOutput, error) 
 CreateWorldExportJob(ctx context.Context, params *CreateWorldExportJobInput, optFns ...func(*Options)) (*CreateWorldExportJobOutput, error) 
 CreateWorldGenerationJob(ctx context.Context, params *CreateWorldGenerationJobInput, optFns ...func(*Options)) (*CreateWorldGenerationJobOutput, error) 
 CreateWorldTemplate(ctx context.Context, params *CreateWorldTemplateInput, optFns ...func(*Options)) (*CreateWorldTemplateOutput, error) 
 DeleteFleet(ctx context.Context, params *DeleteFleetInput, optFns ...func(*Options)) (*DeleteFleetOutput, error) 
 DeleteRobot(ctx context.Context, params *DeleteRobotInput, optFns ...func(*Options)) (*DeleteRobotOutput, error) 
 DeleteRobotApplication(ctx context.Context, params *DeleteRobotApplicationInput, optFns ...func(*Options)) (*DeleteRobotApplicationOutput, error) 
 DeleteSimulationApplication(ctx context.Context, params *DeleteSimulationApplicationInput, optFns ...func(*Options)) (*DeleteSimulationApplicationOutput, error) 
 DeleteWorldTemplate(ctx context.Context, params *DeleteWorldTemplateInput, optFns ...func(*Options)) (*DeleteWorldTemplateOutput, error) 
 DeregisterRobot(ctx context.Context, params *DeregisterRobotInput, optFns ...func(*Options)) (*DeregisterRobotOutput, error) 
 DescribeDeploymentJob(ctx context.Context, params *DescribeDeploymentJobInput, optFns ...func(*Options)) (*DescribeDeploymentJobOutput, error) 
 DescribeFleet(ctx context.Context, params *DescribeFleetInput, optFns ...func(*Options)) (*DescribeFleetOutput, error) 
 DescribeRobot(ctx context.Context, params *DescribeRobotInput, optFns ...func(*Options)) (*DescribeRobotOutput, error) 
 DescribeRobotApplication(ctx context.Context, params *DescribeRobotApplicationInput, optFns ...func(*Options)) (*DescribeRobotApplicationOutput, error) 
 DescribeSimulationApplication(ctx context.Context, params *DescribeSimulationApplicationInput, optFns ...func(*Options)) (*DescribeSimulationApplicationOutput, error) 
 DescribeSimulationJob(ctx context.Context, params *DescribeSimulationJobInput, optFns ...func(*Options)) (*DescribeSimulationJobOutput, error) 
 DescribeSimulationJobBatch(ctx context.Context, params *DescribeSimulationJobBatchInput, optFns ...func(*Options)) (*DescribeSimulationJobBatchOutput, error) 
 DescribeWorld(ctx context.Context, params *DescribeWorldInput, optFns ...func(*Options)) (*DescribeWorldOutput, error) 
 DescribeWorldExportJob(ctx context.Context, params *DescribeWorldExportJobInput, optFns ...func(*Options)) (*DescribeWorldExportJobOutput, error) 
 DescribeWorldGenerationJob(ctx context.Context, params *DescribeWorldGenerationJobInput, optFns ...func(*Options)) (*DescribeWorldGenerationJobOutput, error) 
 DescribeWorldTemplate(ctx context.Context, params *DescribeWorldTemplateInput, optFns ...func(*Options)) (*DescribeWorldTemplateOutput, error) 
 GetWorldTemplateBody(ctx context.Context, params *GetWorldTemplateBodyInput, optFns ...func(*Options)) (*GetWorldTemplateBodyOutput, error) 
 ListDeploymentJobs(ctx context.Context, params *ListDeploymentJobsInput, optFns ...func(*Options)) (*ListDeploymentJobsOutput, error) 
 ListFleets(ctx context.Context, params *ListFleetsInput, optFns ...func(*Options)) (*ListFleetsOutput, error) 
 ListRobotApplications(ctx context.Context, params *ListRobotApplicationsInput, optFns ...func(*Options)) (*ListRobotApplicationsOutput, error) 
 ListRobots(ctx context.Context, params *ListRobotsInput, optFns ...func(*Options)) (*ListRobotsOutput, error) 
 ListSimulationApplications(ctx context.Context, params *ListSimulationApplicationsInput, optFns ...func(*Options)) (*ListSimulationApplicationsOutput, error) 
 ListSimulationJobBatches(ctx context.Context, params *ListSimulationJobBatchesInput, optFns ...func(*Options)) (*ListSimulationJobBatchesOutput, error) 
 ListSimulationJobs(ctx context.Context, params *ListSimulationJobsInput, optFns ...func(*Options)) (*ListSimulationJobsOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 ListWorldExportJobs(ctx context.Context, params *ListWorldExportJobsInput, optFns ...func(*Options)) (*ListWorldExportJobsOutput, error) 
 ListWorldGenerationJobs(ctx context.Context, params *ListWorldGenerationJobsInput, optFns ...func(*Options)) (*ListWorldGenerationJobsOutput, error) 
 ListWorldTemplates(ctx context.Context, params *ListWorldTemplatesInput, optFns ...func(*Options)) (*ListWorldTemplatesOutput, error) 
 ListWorlds(ctx context.Context, params *ListWorldsInput, optFns ...func(*Options)) (*ListWorldsOutput, error) 
 RegisterRobot(ctx context.Context, params *RegisterRobotInput, optFns ...func(*Options)) (*RegisterRobotOutput, error) 
 RestartSimulationJob(ctx context.Context, params *RestartSimulationJobInput, optFns ...func(*Options)) (*RestartSimulationJobOutput, error) 
 StartSimulationJobBatch(ctx context.Context, params *StartSimulationJobBatchInput, optFns ...func(*Options)) (*StartSimulationJobBatchOutput, error) 
 SyncDeploymentJob(ctx context.Context, params *SyncDeploymentJobInput, optFns ...func(*Options)) (*SyncDeploymentJobOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateRobotApplication(ctx context.Context, params *UpdateRobotApplicationInput, optFns ...func(*Options)) (*UpdateRobotApplicationOutput, error) 
 UpdateSimulationApplication(ctx context.Context, params *UpdateSimulationApplicationInput, optFns ...func(*Options)) (*UpdateSimulationApplicationOutput, error) 
 UpdateWorldTemplate(ctx context.Context, params *UpdateWorldTemplateInput, optFns ...func(*Options)) (*UpdateWorldTemplateOutput, error) 
}
