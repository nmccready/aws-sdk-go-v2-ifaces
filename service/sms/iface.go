
package sms

import (
    "github.com/aws/aws-sdk-go-v2/service/sms"
)

// ISms defines the interface for sms
type ISms interface {
 Options() Options 
 CreateApp(ctx context.Context, params *CreateAppInput, optFns ...func(*Options)) (*CreateAppOutput, error) 
 CreateReplicationJob(ctx context.Context, params *CreateReplicationJobInput, optFns ...func(*Options)) (*CreateReplicationJobOutput, error) 
 DeleteApp(ctx context.Context, params *DeleteAppInput, optFns ...func(*Options)) (*DeleteAppOutput, error) 
 DeleteAppLaunchConfiguration(ctx context.Context, params *DeleteAppLaunchConfigurationInput, optFns ...func(*Options)) (*DeleteAppLaunchConfigurationOutput, error) 
 DeleteAppReplicationConfiguration(ctx context.Context, params *DeleteAppReplicationConfigurationInput, optFns ...func(*Options)) (*DeleteAppReplicationConfigurationOutput, error) 
 DeleteAppValidationConfiguration(ctx context.Context, params *DeleteAppValidationConfigurationInput, optFns ...func(*Options)) (*DeleteAppValidationConfigurationOutput, error) 
 DeleteReplicationJob(ctx context.Context, params *DeleteReplicationJobInput, optFns ...func(*Options)) (*DeleteReplicationJobOutput, error) 
 DeleteServerCatalog(ctx context.Context, params *DeleteServerCatalogInput, optFns ...func(*Options)) (*DeleteServerCatalogOutput, error) 
 DisassociateConnector(ctx context.Context, params *DisassociateConnectorInput, optFns ...func(*Options)) (*DisassociateConnectorOutput, error) 
 GenerateChangeSet(ctx context.Context, params *GenerateChangeSetInput, optFns ...func(*Options)) (*GenerateChangeSetOutput, error) 
 GenerateTemplate(ctx context.Context, params *GenerateTemplateInput, optFns ...func(*Options)) (*GenerateTemplateOutput, error) 
 GetApp(ctx context.Context, params *GetAppInput, optFns ...func(*Options)) (*GetAppOutput, error) 
 GetAppLaunchConfiguration(ctx context.Context, params *GetAppLaunchConfigurationInput, optFns ...func(*Options)) (*GetAppLaunchConfigurationOutput, error) 
 GetAppReplicationConfiguration(ctx context.Context, params *GetAppReplicationConfigurationInput, optFns ...func(*Options)) (*GetAppReplicationConfigurationOutput, error) 
 GetAppValidationConfiguration(ctx context.Context, params *GetAppValidationConfigurationInput, optFns ...func(*Options)) (*GetAppValidationConfigurationOutput, error) 
 GetAppValidationOutput(ctx context.Context, params *GetAppValidationOutputInput, optFns ...func(*Options)) (*GetAppValidationOutputOutput, error) 
 GetConnectors(ctx context.Context, params *GetConnectorsInput, optFns ...func(*Options)) (*GetConnectorsOutput, error) 
 GetReplicationJobs(ctx context.Context, params *GetReplicationJobsInput, optFns ...func(*Options)) (*GetReplicationJobsOutput, error) 
 GetReplicationRuns(ctx context.Context, params *GetReplicationRunsInput, optFns ...func(*Options)) (*GetReplicationRunsOutput, error) 
 GetServers(ctx context.Context, params *GetServersInput, optFns ...func(*Options)) (*GetServersOutput, error) 
 ImportAppCatalog(ctx context.Context, params *ImportAppCatalogInput, optFns ...func(*Options)) (*ImportAppCatalogOutput, error) 
 ImportServerCatalog(ctx context.Context, params *ImportServerCatalogInput, optFns ...func(*Options)) (*ImportServerCatalogOutput, error) 
 LaunchApp(ctx context.Context, params *LaunchAppInput, optFns ...func(*Options)) (*LaunchAppOutput, error) 
 ListApps(ctx context.Context, params *ListAppsInput, optFns ...func(*Options)) (*ListAppsOutput, error) 
 NotifyAppValidationOutput(ctx context.Context, params *NotifyAppValidationOutputInput, optFns ...func(*Options)) (*NotifyAppValidationOutputOutput, error) 
 PutAppLaunchConfiguration(ctx context.Context, params *PutAppLaunchConfigurationInput, optFns ...func(*Options)) (*PutAppLaunchConfigurationOutput, error) 
 PutAppReplicationConfiguration(ctx context.Context, params *PutAppReplicationConfigurationInput, optFns ...func(*Options)) (*PutAppReplicationConfigurationOutput, error) 
 PutAppValidationConfiguration(ctx context.Context, params *PutAppValidationConfigurationInput, optFns ...func(*Options)) (*PutAppValidationConfigurationOutput, error) 
 StartAppReplication(ctx context.Context, params *StartAppReplicationInput, optFns ...func(*Options)) (*StartAppReplicationOutput, error) 
 StartOnDemandAppReplication(ctx context.Context, params *StartOnDemandAppReplicationInput, optFns ...func(*Options)) (*StartOnDemandAppReplicationOutput, error) 
 StartOnDemandReplicationRun(ctx context.Context, params *StartOnDemandReplicationRunInput, optFns ...func(*Options)) (*StartOnDemandReplicationRunOutput, error) 
 StopAppReplication(ctx context.Context, params *StopAppReplicationInput, optFns ...func(*Options)) (*StopAppReplicationOutput, error) 
 TerminateApp(ctx context.Context, params *TerminateAppInput, optFns ...func(*Options)) (*TerminateAppOutput, error) 
 UpdateApp(ctx context.Context, params *UpdateAppInput, optFns ...func(*Options)) (*UpdateAppOutput, error) 
 UpdateReplicationJob(ctx context.Context, params *UpdateReplicationJobInput, optFns ...func(*Options)) (*UpdateReplicationJobOutput, error) 
}
