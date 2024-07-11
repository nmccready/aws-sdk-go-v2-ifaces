
package mgn

// This file was generated by scripts/gen_ifaces.sh do not edit it manually
import (
    "github.com/aws/aws-sdk-go-v2/service/mgn"
)

// IClient defines the interface for mgn
type IClient interface {
 Options() Options 
 ArchiveApplication(ctx context.Context, params *ArchiveApplicationInput, optFns ...func(*Options)) (*ArchiveApplicationOutput, error) 
 ArchiveWave(ctx context.Context, params *ArchiveWaveInput, optFns ...func(*Options)) (*ArchiveWaveOutput, error) 
 AssociateApplications(ctx context.Context, params *AssociateApplicationsInput, optFns ...func(*Options)) (*AssociateApplicationsOutput, error) 
 AssociateSourceServers(ctx context.Context, params *AssociateSourceServersInput, optFns ...func(*Options)) (*AssociateSourceServersOutput, error) 
 ChangeServerLifeCycleState(ctx context.Context, params *ChangeServerLifeCycleStateInput, optFns ...func(*Options)) (*ChangeServerLifeCycleStateOutput, error) 
 CreateApplication(ctx context.Context, params *CreateApplicationInput, optFns ...func(*Options)) (*CreateApplicationOutput, error) 
 CreateConnector(ctx context.Context, params *CreateConnectorInput, optFns ...func(*Options)) (*CreateConnectorOutput, error) 
 CreateLaunchConfigurationTemplate(ctx context.Context, params *CreateLaunchConfigurationTemplateInput, optFns ...func(*Options)) (*CreateLaunchConfigurationTemplateOutput, error) 
 CreateReplicationConfigurationTemplate(ctx context.Context, params *CreateReplicationConfigurationTemplateInput, optFns ...func(*Options)) (*CreateReplicationConfigurationTemplateOutput, error) 
 CreateWave(ctx context.Context, params *CreateWaveInput, optFns ...func(*Options)) (*CreateWaveOutput, error) 
 DeleteApplication(ctx context.Context, params *DeleteApplicationInput, optFns ...func(*Options)) (*DeleteApplicationOutput, error) 
 DeleteConnector(ctx context.Context, params *DeleteConnectorInput, optFns ...func(*Options)) (*DeleteConnectorOutput, error) 
 DeleteJob(ctx context.Context, params *DeleteJobInput, optFns ...func(*Options)) (*DeleteJobOutput, error) 
 DeleteLaunchConfigurationTemplate(ctx context.Context, params *DeleteLaunchConfigurationTemplateInput, optFns ...func(*Options)) (*DeleteLaunchConfigurationTemplateOutput, error) 
 DeleteReplicationConfigurationTemplate(ctx context.Context, params *DeleteReplicationConfigurationTemplateInput, optFns ...func(*Options)) (*DeleteReplicationConfigurationTemplateOutput, error) 
 DeleteSourceServer(ctx context.Context, params *DeleteSourceServerInput, optFns ...func(*Options)) (*DeleteSourceServerOutput, error) 
 DeleteVcenterClient(ctx context.Context, params *DeleteVcenterClientInput, optFns ...func(*Options)) (*DeleteVcenterClientOutput, error) 
 DeleteWave(ctx context.Context, params *DeleteWaveInput, optFns ...func(*Options)) (*DeleteWaveOutput, error) 
 DescribeJobLogItems(ctx context.Context, params *DescribeJobLogItemsInput, optFns ...func(*Options)) (*DescribeJobLogItemsOutput, error) 
 DescribeJobs(ctx context.Context, params *DescribeJobsInput, optFns ...func(*Options)) (*DescribeJobsOutput, error) 
 DescribeLaunchConfigurationTemplates(ctx context.Context, params *DescribeLaunchConfigurationTemplatesInput, optFns ...func(*Options)) (*DescribeLaunchConfigurationTemplatesOutput, error) 
 DescribeReplicationConfigurationTemplates(ctx context.Context, params *DescribeReplicationConfigurationTemplatesInput, optFns ...func(*Options)) (*DescribeReplicationConfigurationTemplatesOutput, error) 
 DescribeSourceServers(ctx context.Context, params *DescribeSourceServersInput, optFns ...func(*Options)) (*DescribeSourceServersOutput, error) 
 DescribeVcenterClients(ctx context.Context, params *DescribeVcenterClientsInput, optFns ...func(*Options)) (*DescribeVcenterClientsOutput, error) 
 DisassociateApplications(ctx context.Context, params *DisassociateApplicationsInput, optFns ...func(*Options)) (*DisassociateApplicationsOutput, error) 
 DisassociateSourceServers(ctx context.Context, params *DisassociateSourceServersInput, optFns ...func(*Options)) (*DisassociateSourceServersOutput, error) 
 DisconnectFromService(ctx context.Context, params *DisconnectFromServiceInput, optFns ...func(*Options)) (*DisconnectFromServiceOutput, error) 
 FinalizeCutover(ctx context.Context, params *FinalizeCutoverInput, optFns ...func(*Options)) (*FinalizeCutoverOutput, error) 
 GetLaunchConfiguration(ctx context.Context, params *GetLaunchConfigurationInput, optFns ...func(*Options)) (*GetLaunchConfigurationOutput, error) 
 GetReplicationConfiguration(ctx context.Context, params *GetReplicationConfigurationInput, optFns ...func(*Options)) (*GetReplicationConfigurationOutput, error) 
 InitializeService(ctx context.Context, params *InitializeServiceInput, optFns ...func(*Options)) (*InitializeServiceOutput, error) 
 ListApplications(ctx context.Context, params *ListApplicationsInput, optFns ...func(*Options)) (*ListApplicationsOutput, error) 
 ListConnectors(ctx context.Context, params *ListConnectorsInput, optFns ...func(*Options)) (*ListConnectorsOutput, error) 
 ListExportErrors(ctx context.Context, params *ListExportErrorsInput, optFns ...func(*Options)) (*ListExportErrorsOutput, error) 
 ListExports(ctx context.Context, params *ListExportsInput, optFns ...func(*Options)) (*ListExportsOutput, error) 
 ListImportErrors(ctx context.Context, params *ListImportErrorsInput, optFns ...func(*Options)) (*ListImportErrorsOutput, error) 
 ListImports(ctx context.Context, params *ListImportsInput, optFns ...func(*Options)) (*ListImportsOutput, error) 
 ListManagedAccounts(ctx context.Context, params *ListManagedAccountsInput, optFns ...func(*Options)) (*ListManagedAccountsOutput, error) 
 ListSourceServerActions(ctx context.Context, params *ListSourceServerActionsInput, optFns ...func(*Options)) (*ListSourceServerActionsOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 ListTemplateActions(ctx context.Context, params *ListTemplateActionsInput, optFns ...func(*Options)) (*ListTemplateActionsOutput, error) 
 ListWaves(ctx context.Context, params *ListWavesInput, optFns ...func(*Options)) (*ListWavesOutput, error) 
 MarkAsArchived(ctx context.Context, params *MarkAsArchivedInput, optFns ...func(*Options)) (*MarkAsArchivedOutput, error) 
 PauseReplication(ctx context.Context, params *PauseReplicationInput, optFns ...func(*Options)) (*PauseReplicationOutput, error) 
 PutSourceServerAction(ctx context.Context, params *PutSourceServerActionInput, optFns ...func(*Options)) (*PutSourceServerActionOutput, error) 
 PutTemplateAction(ctx context.Context, params *PutTemplateActionInput, optFns ...func(*Options)) (*PutTemplateActionOutput, error) 
 RemoveSourceServerAction(ctx context.Context, params *RemoveSourceServerActionInput, optFns ...func(*Options)) (*RemoveSourceServerActionOutput, error) 
 RemoveTemplateAction(ctx context.Context, params *RemoveTemplateActionInput, optFns ...func(*Options)) (*RemoveTemplateActionOutput, error) 
 ResumeReplication(ctx context.Context, params *ResumeReplicationInput, optFns ...func(*Options)) (*ResumeReplicationOutput, error) 
 RetryDataReplication(ctx context.Context, params *RetryDataReplicationInput, optFns ...func(*Options)) (*RetryDataReplicationOutput, error) 
 StartCutover(ctx context.Context, params *StartCutoverInput, optFns ...func(*Options)) (*StartCutoverOutput, error) 
 StartExport(ctx context.Context, params *StartExportInput, optFns ...func(*Options)) (*StartExportOutput, error) 
 StartImport(ctx context.Context, params *StartImportInput, optFns ...func(*Options)) (*StartImportOutput, error) 
 StartReplication(ctx context.Context, params *StartReplicationInput, optFns ...func(*Options)) (*StartReplicationOutput, error) 
 StartTest(ctx context.Context, params *StartTestInput, optFns ...func(*Options)) (*StartTestOutput, error) 
 StopReplication(ctx context.Context, params *StopReplicationInput, optFns ...func(*Options)) (*StopReplicationOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 TerminateTargetInstances(ctx context.Context, params *TerminateTargetInstancesInput, optFns ...func(*Options)) (*TerminateTargetInstancesOutput, error) 
 UnarchiveApplication(ctx context.Context, params *UnarchiveApplicationInput, optFns ...func(*Options)) (*UnarchiveApplicationOutput, error) 
 UnarchiveWave(ctx context.Context, params *UnarchiveWaveInput, optFns ...func(*Options)) (*UnarchiveWaveOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateApplication(ctx context.Context, params *UpdateApplicationInput, optFns ...func(*Options)) (*UpdateApplicationOutput, error) 
 UpdateConnector(ctx context.Context, params *UpdateConnectorInput, optFns ...func(*Options)) (*UpdateConnectorOutput, error) 
 UpdateLaunchConfiguration(ctx context.Context, params *UpdateLaunchConfigurationInput, optFns ...func(*Options)) (*UpdateLaunchConfigurationOutput, error) 
 UpdateLaunchConfigurationTemplate(ctx context.Context, params *UpdateLaunchConfigurationTemplateInput, optFns ...func(*Options)) (*UpdateLaunchConfigurationTemplateOutput, error) 
 UpdateReplicationConfiguration(ctx context.Context, params *UpdateReplicationConfigurationInput, optFns ...func(*Options)) (*UpdateReplicationConfigurationOutput, error) 
 UpdateReplicationConfigurationTemplate(ctx context.Context, params *UpdateReplicationConfigurationTemplateInput, optFns ...func(*Options)) (*UpdateReplicationConfigurationTemplateOutput, error) 
 UpdateSourceServer(ctx context.Context, params *UpdateSourceServerInput, optFns ...func(*Options)) (*UpdateSourceServerOutput, error) 
 UpdateSourceServerReplicationType(ctx context.Context, params *UpdateSourceServerReplicationTypeInput, optFns ...func(*Options)) (*UpdateSourceServerReplicationTypeOutput, error) 
 UpdateWave(ctx context.Context, params *UpdateWaveInput, optFns ...func(*Options)) (*UpdateWaveOutput, error) 
 ID() string 
 HandleDeserialize(ctx context.Context, in middleware.DeserializeInput, next middleware.DeserializeHandler) (
 ID() string 
 HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 ID() string 
 HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
}
