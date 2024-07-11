
package iotwireless

import (
    "github.com/aws/aws-sdk-go-v2/service/iotwireless"
)

// IClient defines the interface for iotwireless
type IClient interface {
 Options() Options 
 AssociateAwsAccountWithPartnerAccount(ctx context.Context, params *AssociateAwsAccountWithPartnerAccountInput, optFns ...func(*Options)) (*AssociateAwsAccountWithPartnerAccountOutput, error) 
 AssociateMulticastGroupWithFuotaTask(ctx context.Context, params *AssociateMulticastGroupWithFuotaTaskInput, optFns ...func(*Options)) (*AssociateMulticastGroupWithFuotaTaskOutput, error) 
 AssociateWirelessDeviceWithFuotaTask(ctx context.Context, params *AssociateWirelessDeviceWithFuotaTaskInput, optFns ...func(*Options)) (*AssociateWirelessDeviceWithFuotaTaskOutput, error) 
 AssociateWirelessDeviceWithMulticastGroup(ctx context.Context, params *AssociateWirelessDeviceWithMulticastGroupInput, optFns ...func(*Options)) (*AssociateWirelessDeviceWithMulticastGroupOutput, error) 
 AssociateWirelessDeviceWithThing(ctx context.Context, params *AssociateWirelessDeviceWithThingInput, optFns ...func(*Options)) (*AssociateWirelessDeviceWithThingOutput, error) 
 AssociateWirelessGatewayWithCertificate(ctx context.Context, params *AssociateWirelessGatewayWithCertificateInput, optFns ...func(*Options)) (*AssociateWirelessGatewayWithCertificateOutput, error) 
 AssociateWirelessGatewayWithThing(ctx context.Context, params *AssociateWirelessGatewayWithThingInput, optFns ...func(*Options)) (*AssociateWirelessGatewayWithThingOutput, error) 
 CancelMulticastGroupSession(ctx context.Context, params *CancelMulticastGroupSessionInput, optFns ...func(*Options)) (*CancelMulticastGroupSessionOutput, error) 
 CreateDestination(ctx context.Context, params *CreateDestinationInput, optFns ...func(*Options)) (*CreateDestinationOutput, error) 
 CreateDeviceProfile(ctx context.Context, params *CreateDeviceProfileInput, optFns ...func(*Options)) (*CreateDeviceProfileOutput, error) 
 CreateFuotaTask(ctx context.Context, params *CreateFuotaTaskInput, optFns ...func(*Options)) (*CreateFuotaTaskOutput, error) 
 CreateMulticastGroup(ctx context.Context, params *CreateMulticastGroupInput, optFns ...func(*Options)) (*CreateMulticastGroupOutput, error) 
 CreateNetworkAnalyzerConfiguration(ctx context.Context, params *CreateNetworkAnalyzerConfigurationInput, optFns ...func(*Options)) (*CreateNetworkAnalyzerConfigurationOutput, error) 
 CreateServiceProfile(ctx context.Context, params *CreateServiceProfileInput, optFns ...func(*Options)) (*CreateServiceProfileOutput, error) 
 CreateWirelessDevice(ctx context.Context, params *CreateWirelessDeviceInput, optFns ...func(*Options)) (*CreateWirelessDeviceOutput, error) 
 CreateWirelessGateway(ctx context.Context, params *CreateWirelessGatewayInput, optFns ...func(*Options)) (*CreateWirelessGatewayOutput, error) 
 CreateWirelessGatewayTask(ctx context.Context, params *CreateWirelessGatewayTaskInput, optFns ...func(*Options)) (*CreateWirelessGatewayTaskOutput, error) 
 CreateWirelessGatewayTaskDefinition(ctx context.Context, params *CreateWirelessGatewayTaskDefinitionInput, optFns ...func(*Options)) (*CreateWirelessGatewayTaskDefinitionOutput, error) 
 DeleteDestination(ctx context.Context, params *DeleteDestinationInput, optFns ...func(*Options)) (*DeleteDestinationOutput, error) 
 DeleteDeviceProfile(ctx context.Context, params *DeleteDeviceProfileInput, optFns ...func(*Options)) (*DeleteDeviceProfileOutput, error) 
 DeleteFuotaTask(ctx context.Context, params *DeleteFuotaTaskInput, optFns ...func(*Options)) (*DeleteFuotaTaskOutput, error) 
 DeleteMulticastGroup(ctx context.Context, params *DeleteMulticastGroupInput, optFns ...func(*Options)) (*DeleteMulticastGroupOutput, error) 
 DeleteNetworkAnalyzerConfiguration(ctx context.Context, params *DeleteNetworkAnalyzerConfigurationInput, optFns ...func(*Options)) (*DeleteNetworkAnalyzerConfigurationOutput, error) 
 DeleteQueuedMessages(ctx context.Context, params *DeleteQueuedMessagesInput, optFns ...func(*Options)) (*DeleteQueuedMessagesOutput, error) 
 DeleteServiceProfile(ctx context.Context, params *DeleteServiceProfileInput, optFns ...func(*Options)) (*DeleteServiceProfileOutput, error) 
 DeleteWirelessDevice(ctx context.Context, params *DeleteWirelessDeviceInput, optFns ...func(*Options)) (*DeleteWirelessDeviceOutput, error) 
 DeleteWirelessDeviceImportTask(ctx context.Context, params *DeleteWirelessDeviceImportTaskInput, optFns ...func(*Options)) (*DeleteWirelessDeviceImportTaskOutput, error) 
 DeleteWirelessGateway(ctx context.Context, params *DeleteWirelessGatewayInput, optFns ...func(*Options)) (*DeleteWirelessGatewayOutput, error) 
 DeleteWirelessGatewayTask(ctx context.Context, params *DeleteWirelessGatewayTaskInput, optFns ...func(*Options)) (*DeleteWirelessGatewayTaskOutput, error) 
 DeleteWirelessGatewayTaskDefinition(ctx context.Context, params *DeleteWirelessGatewayTaskDefinitionInput, optFns ...func(*Options)) (*DeleteWirelessGatewayTaskDefinitionOutput, error) 
 DeregisterWirelessDevice(ctx context.Context, params *DeregisterWirelessDeviceInput, optFns ...func(*Options)) (*DeregisterWirelessDeviceOutput, error) 
 DisassociateAwsAccountFromPartnerAccount(ctx context.Context, params *DisassociateAwsAccountFromPartnerAccountInput, optFns ...func(*Options)) (*DisassociateAwsAccountFromPartnerAccountOutput, error) 
 DisassociateMulticastGroupFromFuotaTask(ctx context.Context, params *DisassociateMulticastGroupFromFuotaTaskInput, optFns ...func(*Options)) (*DisassociateMulticastGroupFromFuotaTaskOutput, error) 
 DisassociateWirelessDeviceFromFuotaTask(ctx context.Context, params *DisassociateWirelessDeviceFromFuotaTaskInput, optFns ...func(*Options)) (*DisassociateWirelessDeviceFromFuotaTaskOutput, error) 
 DisassociateWirelessDeviceFromMulticastGroup(ctx context.Context, params *DisassociateWirelessDeviceFromMulticastGroupInput, optFns ...func(*Options)) (*DisassociateWirelessDeviceFromMulticastGroupOutput, error) 
 DisassociateWirelessDeviceFromThing(ctx context.Context, params *DisassociateWirelessDeviceFromThingInput, optFns ...func(*Options)) (*DisassociateWirelessDeviceFromThingOutput, error) 
 DisassociateWirelessGatewayFromCertificate(ctx context.Context, params *DisassociateWirelessGatewayFromCertificateInput, optFns ...func(*Options)) (*DisassociateWirelessGatewayFromCertificateOutput, error) 
 DisassociateWirelessGatewayFromThing(ctx context.Context, params *DisassociateWirelessGatewayFromThingInput, optFns ...func(*Options)) (*DisassociateWirelessGatewayFromThingOutput, error) 
 GetDestination(ctx context.Context, params *GetDestinationInput, optFns ...func(*Options)) (*GetDestinationOutput, error) 
 GetDeviceProfile(ctx context.Context, params *GetDeviceProfileInput, optFns ...func(*Options)) (*GetDeviceProfileOutput, error) 
 GetEventConfigurationByResourceTypes(ctx context.Context, params *GetEventConfigurationByResourceTypesInput, optFns ...func(*Options)) (*GetEventConfigurationByResourceTypesOutput, error) 
 GetFuotaTask(ctx context.Context, params *GetFuotaTaskInput, optFns ...func(*Options)) (*GetFuotaTaskOutput, error) 
 GetLogLevelsByResourceTypes(ctx context.Context, params *GetLogLevelsByResourceTypesInput, optFns ...func(*Options)) (*GetLogLevelsByResourceTypesOutput, error) 
 GetMetricConfiguration(ctx context.Context, params *GetMetricConfigurationInput, optFns ...func(*Options)) (*GetMetricConfigurationOutput, error) 
 GetMetrics(ctx context.Context, params *GetMetricsInput, optFns ...func(*Options)) (*GetMetricsOutput, error) 
 GetMulticastGroup(ctx context.Context, params *GetMulticastGroupInput, optFns ...func(*Options)) (*GetMulticastGroupOutput, error) 
 GetMulticastGroupSession(ctx context.Context, params *GetMulticastGroupSessionInput, optFns ...func(*Options)) (*GetMulticastGroupSessionOutput, error) 
 GetNetworkAnalyzerConfiguration(ctx context.Context, params *GetNetworkAnalyzerConfigurationInput, optFns ...func(*Options)) (*GetNetworkAnalyzerConfigurationOutput, error) 
 GetPartnerAccount(ctx context.Context, params *GetPartnerAccountInput, optFns ...func(*Options)) (*GetPartnerAccountOutput, error) 
 GetPosition(ctx context.Context, params *GetPositionInput, optFns ...func(*Options)) (*GetPositionOutput, error) 
 GetPositionConfiguration(ctx context.Context, params *GetPositionConfigurationInput, optFns ...func(*Options)) (*GetPositionConfigurationOutput, error) 
 GetPositionEstimate(ctx context.Context, params *GetPositionEstimateInput, optFns ...func(*Options)) (*GetPositionEstimateOutput, error) 
 GetResourceEventConfiguration(ctx context.Context, params *GetResourceEventConfigurationInput, optFns ...func(*Options)) (*GetResourceEventConfigurationOutput, error) 
 GetResourceLogLevel(ctx context.Context, params *GetResourceLogLevelInput, optFns ...func(*Options)) (*GetResourceLogLevelOutput, error) 
 GetResourcePosition(ctx context.Context, params *GetResourcePositionInput, optFns ...func(*Options)) (*GetResourcePositionOutput, error) 
 GetServiceEndpoint(ctx context.Context, params *GetServiceEndpointInput, optFns ...func(*Options)) (*GetServiceEndpointOutput, error) 
 GetServiceProfile(ctx context.Context, params *GetServiceProfileInput, optFns ...func(*Options)) (*GetServiceProfileOutput, error) 
 GetWirelessDevice(ctx context.Context, params *GetWirelessDeviceInput, optFns ...func(*Options)) (*GetWirelessDeviceOutput, error) 
 GetWirelessDeviceImportTask(ctx context.Context, params *GetWirelessDeviceImportTaskInput, optFns ...func(*Options)) (*GetWirelessDeviceImportTaskOutput, error) 
 GetWirelessDeviceStatistics(ctx context.Context, params *GetWirelessDeviceStatisticsInput, optFns ...func(*Options)) (*GetWirelessDeviceStatisticsOutput, error) 
 GetWirelessGateway(ctx context.Context, params *GetWirelessGatewayInput, optFns ...func(*Options)) (*GetWirelessGatewayOutput, error) 
 GetWirelessGatewayCertificate(ctx context.Context, params *GetWirelessGatewayCertificateInput, optFns ...func(*Options)) (*GetWirelessGatewayCertificateOutput, error) 
 GetWirelessGatewayFirmwareInformation(ctx context.Context, params *GetWirelessGatewayFirmwareInformationInput, optFns ...func(*Options)) (*GetWirelessGatewayFirmwareInformationOutput, error) 
 GetWirelessGatewayStatistics(ctx context.Context, params *GetWirelessGatewayStatisticsInput, optFns ...func(*Options)) (*GetWirelessGatewayStatisticsOutput, error) 
 GetWirelessGatewayTask(ctx context.Context, params *GetWirelessGatewayTaskInput, optFns ...func(*Options)) (*GetWirelessGatewayTaskOutput, error) 
 GetWirelessGatewayTaskDefinition(ctx context.Context, params *GetWirelessGatewayTaskDefinitionInput, optFns ...func(*Options)) (*GetWirelessGatewayTaskDefinitionOutput, error) 
 ListDestinations(ctx context.Context, params *ListDestinationsInput, optFns ...func(*Options)) (*ListDestinationsOutput, error) 
 ListDeviceProfiles(ctx context.Context, params *ListDeviceProfilesInput, optFns ...func(*Options)) (*ListDeviceProfilesOutput, error) 
 ListDevicesForWirelessDeviceImportTask(ctx context.Context, params *ListDevicesForWirelessDeviceImportTaskInput, optFns ...func(*Options)) (*ListDevicesForWirelessDeviceImportTaskOutput, error) 
 ListEventConfigurations(ctx context.Context, params *ListEventConfigurationsInput, optFns ...func(*Options)) (*ListEventConfigurationsOutput, error) 
 ListFuotaTasks(ctx context.Context, params *ListFuotaTasksInput, optFns ...func(*Options)) (*ListFuotaTasksOutput, error) 
 ListMulticastGroups(ctx context.Context, params *ListMulticastGroupsInput, optFns ...func(*Options)) (*ListMulticastGroupsOutput, error) 
 ListMulticastGroupsByFuotaTask(ctx context.Context, params *ListMulticastGroupsByFuotaTaskInput, optFns ...func(*Options)) (*ListMulticastGroupsByFuotaTaskOutput, error) 
 ListNetworkAnalyzerConfigurations(ctx context.Context, params *ListNetworkAnalyzerConfigurationsInput, optFns ...func(*Options)) (*ListNetworkAnalyzerConfigurationsOutput, error) 
 ListPartnerAccounts(ctx context.Context, params *ListPartnerAccountsInput, optFns ...func(*Options)) (*ListPartnerAccountsOutput, error) 
 ListPositionConfigurations(ctx context.Context, params *ListPositionConfigurationsInput, optFns ...func(*Options)) (*ListPositionConfigurationsOutput, error) 
 ListQueuedMessages(ctx context.Context, params *ListQueuedMessagesInput, optFns ...func(*Options)) (*ListQueuedMessagesOutput, error) 
 ListServiceProfiles(ctx context.Context, params *ListServiceProfilesInput, optFns ...func(*Options)) (*ListServiceProfilesOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 ListWirelessDeviceImportTasks(ctx context.Context, params *ListWirelessDeviceImportTasksInput, optFns ...func(*Options)) (*ListWirelessDeviceImportTasksOutput, error) 
 ListWirelessDevices(ctx context.Context, params *ListWirelessDevicesInput, optFns ...func(*Options)) (*ListWirelessDevicesOutput, error) 
 ListWirelessGatewayTaskDefinitions(ctx context.Context, params *ListWirelessGatewayTaskDefinitionsInput, optFns ...func(*Options)) (*ListWirelessGatewayTaskDefinitionsOutput, error) 
 ListWirelessGateways(ctx context.Context, params *ListWirelessGatewaysInput, optFns ...func(*Options)) (*ListWirelessGatewaysOutput, error) 
 PutPositionConfiguration(ctx context.Context, params *PutPositionConfigurationInput, optFns ...func(*Options)) (*PutPositionConfigurationOutput, error) 
 PutResourceLogLevel(ctx context.Context, params *PutResourceLogLevelInput, optFns ...func(*Options)) (*PutResourceLogLevelOutput, error) 
 ResetAllResourceLogLevels(ctx context.Context, params *ResetAllResourceLogLevelsInput, optFns ...func(*Options)) (*ResetAllResourceLogLevelsOutput, error) 
 ResetResourceLogLevel(ctx context.Context, params *ResetResourceLogLevelInput, optFns ...func(*Options)) (*ResetResourceLogLevelOutput, error) 
 SendDataToMulticastGroup(ctx context.Context, params *SendDataToMulticastGroupInput, optFns ...func(*Options)) (*SendDataToMulticastGroupOutput, error) 
 SendDataToWirelessDevice(ctx context.Context, params *SendDataToWirelessDeviceInput, optFns ...func(*Options)) (*SendDataToWirelessDeviceOutput, error) 
 StartBulkAssociateWirelessDeviceWithMulticastGroup(ctx context.Context, params *StartBulkAssociateWirelessDeviceWithMulticastGroupInput, optFns ...func(*Options)) (*StartBulkAssociateWirelessDeviceWithMulticastGroupOutput, error) 
 StartBulkDisassociateWirelessDeviceFromMulticastGroup(ctx context.Context, params *StartBulkDisassociateWirelessDeviceFromMulticastGroupInput, optFns ...func(*Options)) (*StartBulkDisassociateWirelessDeviceFromMulticastGroupOutput, error) 
 StartFuotaTask(ctx context.Context, params *StartFuotaTaskInput, optFns ...func(*Options)) (*StartFuotaTaskOutput, error) 
 StartMulticastGroupSession(ctx context.Context, params *StartMulticastGroupSessionInput, optFns ...func(*Options)) (*StartMulticastGroupSessionOutput, error) 
 StartSingleWirelessDeviceImportTask(ctx context.Context, params *StartSingleWirelessDeviceImportTaskInput, optFns ...func(*Options)) (*StartSingleWirelessDeviceImportTaskOutput, error) 
 StartWirelessDeviceImportTask(ctx context.Context, params *StartWirelessDeviceImportTaskInput, optFns ...func(*Options)) (*StartWirelessDeviceImportTaskOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 TestWirelessDevice(ctx context.Context, params *TestWirelessDeviceInput, optFns ...func(*Options)) (*TestWirelessDeviceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateDestination(ctx context.Context, params *UpdateDestinationInput, optFns ...func(*Options)) (*UpdateDestinationOutput, error) 
 UpdateEventConfigurationByResourceTypes(ctx context.Context, params *UpdateEventConfigurationByResourceTypesInput, optFns ...func(*Options)) (*UpdateEventConfigurationByResourceTypesOutput, error) 
 UpdateFuotaTask(ctx context.Context, params *UpdateFuotaTaskInput, optFns ...func(*Options)) (*UpdateFuotaTaskOutput, error) 
 UpdateLogLevelsByResourceTypes(ctx context.Context, params *UpdateLogLevelsByResourceTypesInput, optFns ...func(*Options)) (*UpdateLogLevelsByResourceTypesOutput, error) 
 UpdateMetricConfiguration(ctx context.Context, params *UpdateMetricConfigurationInput, optFns ...func(*Options)) (*UpdateMetricConfigurationOutput, error) 
 UpdateMulticastGroup(ctx context.Context, params *UpdateMulticastGroupInput, optFns ...func(*Options)) (*UpdateMulticastGroupOutput, error) 
 UpdateNetworkAnalyzerConfiguration(ctx context.Context, params *UpdateNetworkAnalyzerConfigurationInput, optFns ...func(*Options)) (*UpdateNetworkAnalyzerConfigurationOutput, error) 
 UpdatePartnerAccount(ctx context.Context, params *UpdatePartnerAccountInput, optFns ...func(*Options)) (*UpdatePartnerAccountOutput, error) 
 UpdatePosition(ctx context.Context, params *UpdatePositionInput, optFns ...func(*Options)) (*UpdatePositionOutput, error) 
 UpdateResourceEventConfiguration(ctx context.Context, params *UpdateResourceEventConfigurationInput, optFns ...func(*Options)) (*UpdateResourceEventConfigurationOutput, error) 
 UpdateResourcePosition(ctx context.Context, params *UpdateResourcePositionInput, optFns ...func(*Options)) (*UpdateResourcePositionOutput, error) 
 UpdateWirelessDevice(ctx context.Context, params *UpdateWirelessDeviceInput, optFns ...func(*Options)) (*UpdateWirelessDeviceOutput, error) 
 UpdateWirelessDeviceImportTask(ctx context.Context, params *UpdateWirelessDeviceImportTaskInput, optFns ...func(*Options)) (*UpdateWirelessDeviceImportTaskOutput, error) 
 UpdateWirelessGateway(ctx context.Context, params *UpdateWirelessGatewayInput, optFns ...func(*Options)) (*UpdateWirelessGatewayOutput, error) 
}
