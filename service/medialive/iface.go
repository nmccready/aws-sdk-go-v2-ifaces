
package medialive

import (
    "github.com/aws/aws-sdk-go-v2/service/medialive"
)

// IClient defines the interface for medialive
type IClient interface {
 Options() Options 
 AcceptInputDeviceTransfer(ctx context.Context, params *AcceptInputDeviceTransferInput, optFns ...func(*Options)) (*AcceptInputDeviceTransferOutput, error) 
 BatchDelete(ctx context.Context, params *BatchDeleteInput, optFns ...func(*Options)) (*BatchDeleteOutput, error) 
 BatchStart(ctx context.Context, params *BatchStartInput, optFns ...func(*Options)) (*BatchStartOutput, error) 
 BatchStop(ctx context.Context, params *BatchStopInput, optFns ...func(*Options)) (*BatchStopOutput, error) 
 BatchUpdateSchedule(ctx context.Context, params *BatchUpdateScheduleInput, optFns ...func(*Options)) (*BatchUpdateScheduleOutput, error) 
 CancelInputDeviceTransfer(ctx context.Context, params *CancelInputDeviceTransferInput, optFns ...func(*Options)) (*CancelInputDeviceTransferOutput, error) 
 ClaimDevice(ctx context.Context, params *ClaimDeviceInput, optFns ...func(*Options)) (*ClaimDeviceOutput, error) 
 CreateChannel(ctx context.Context, params *CreateChannelInput, optFns ...func(*Options)) (*CreateChannelOutput, error) 
 CreateCloudWatchAlarmTemplate(ctx context.Context, params *CreateCloudWatchAlarmTemplateInput, optFns ...func(*Options)) (*CreateCloudWatchAlarmTemplateOutput, error) 
 CreateCloudWatchAlarmTemplateGroup(ctx context.Context, params *CreateCloudWatchAlarmTemplateGroupInput, optFns ...func(*Options)) (*CreateCloudWatchAlarmTemplateGroupOutput, error) 
 CreateEventBridgeRuleTemplate(ctx context.Context, params *CreateEventBridgeRuleTemplateInput, optFns ...func(*Options)) (*CreateEventBridgeRuleTemplateOutput, error) 
 CreateEventBridgeRuleTemplateGroup(ctx context.Context, params *CreateEventBridgeRuleTemplateGroupInput, optFns ...func(*Options)) (*CreateEventBridgeRuleTemplateGroupOutput, error) 
 CreateInput(ctx context.Context, params *CreateInputInput, optFns ...func(*Options)) (*CreateInputOutput, error) 
 CreateInputSecurityGroup(ctx context.Context, params *CreateInputSecurityGroupInput, optFns ...func(*Options)) (*CreateInputSecurityGroupOutput, error) 
 CreateMultiplex(ctx context.Context, params *CreateMultiplexInput, optFns ...func(*Options)) (*CreateMultiplexOutput, error) 
 CreateMultiplexProgram(ctx context.Context, params *CreateMultiplexProgramInput, optFns ...func(*Options)) (*CreateMultiplexProgramOutput, error) 
 CreatePartnerInput(ctx context.Context, params *CreatePartnerInputInput, optFns ...func(*Options)) (*CreatePartnerInputOutput, error) 
 CreateSignalMap(ctx context.Context, params *CreateSignalMapInput, optFns ...func(*Options)) (*CreateSignalMapOutput, error) 
 CreateTags(ctx context.Context, params *CreateTagsInput, optFns ...func(*Options)) (*CreateTagsOutput, error) 
 DeleteChannel(ctx context.Context, params *DeleteChannelInput, optFns ...func(*Options)) (*DeleteChannelOutput, error) 
 DeleteCloudWatchAlarmTemplate(ctx context.Context, params *DeleteCloudWatchAlarmTemplateInput, optFns ...func(*Options)) (*DeleteCloudWatchAlarmTemplateOutput, error) 
 DeleteCloudWatchAlarmTemplateGroup(ctx context.Context, params *DeleteCloudWatchAlarmTemplateGroupInput, optFns ...func(*Options)) (*DeleteCloudWatchAlarmTemplateGroupOutput, error) 
 DeleteEventBridgeRuleTemplate(ctx context.Context, params *DeleteEventBridgeRuleTemplateInput, optFns ...func(*Options)) (*DeleteEventBridgeRuleTemplateOutput, error) 
 DeleteEventBridgeRuleTemplateGroup(ctx context.Context, params *DeleteEventBridgeRuleTemplateGroupInput, optFns ...func(*Options)) (*DeleteEventBridgeRuleTemplateGroupOutput, error) 
 DeleteInput(ctx context.Context, params *DeleteInputInput, optFns ...func(*Options)) (*DeleteInputOutput, error) 
 DeleteInputSecurityGroup(ctx context.Context, params *DeleteInputSecurityGroupInput, optFns ...func(*Options)) (*DeleteInputSecurityGroupOutput, error) 
 DeleteMultiplex(ctx context.Context, params *DeleteMultiplexInput, optFns ...func(*Options)) (*DeleteMultiplexOutput, error) 
 DeleteMultiplexProgram(ctx context.Context, params *DeleteMultiplexProgramInput, optFns ...func(*Options)) (*DeleteMultiplexProgramOutput, error) 
 DeleteReservation(ctx context.Context, params *DeleteReservationInput, optFns ...func(*Options)) (*DeleteReservationOutput, error) 
 DeleteSchedule(ctx context.Context, params *DeleteScheduleInput, optFns ...func(*Options)) (*DeleteScheduleOutput, error) 
 DeleteSignalMap(ctx context.Context, params *DeleteSignalMapInput, optFns ...func(*Options)) (*DeleteSignalMapOutput, error) 
 DeleteTags(ctx context.Context, params *DeleteTagsInput, optFns ...func(*Options)) (*DeleteTagsOutput, error) 
 DescribeAccountConfiguration(ctx context.Context, params *DescribeAccountConfigurationInput, optFns ...func(*Options)) (*DescribeAccountConfigurationOutput, error) 
 DescribeChannel(ctx context.Context, params *DescribeChannelInput, optFns ...func(*Options)) (*DescribeChannelOutput, error) 
 DescribeInput(ctx context.Context, params *DescribeInputInput, optFns ...func(*Options)) (*DescribeInputOutput, error) 
 DescribeInputDevice(ctx context.Context, params *DescribeInputDeviceInput, optFns ...func(*Options)) (*DescribeInputDeviceOutput, error) 
 DescribeInputDeviceThumbnail(ctx context.Context, params *DescribeInputDeviceThumbnailInput, optFns ...func(*Options)) (*DescribeInputDeviceThumbnailOutput, error) 
 DescribeInputSecurityGroup(ctx context.Context, params *DescribeInputSecurityGroupInput, optFns ...func(*Options)) (*DescribeInputSecurityGroupOutput, error) 
 DescribeMultiplex(ctx context.Context, params *DescribeMultiplexInput, optFns ...func(*Options)) (*DescribeMultiplexOutput, error) 
 DescribeMultiplexProgram(ctx context.Context, params *DescribeMultiplexProgramInput, optFns ...func(*Options)) (*DescribeMultiplexProgramOutput, error) 
 DescribeOffering(ctx context.Context, params *DescribeOfferingInput, optFns ...func(*Options)) (*DescribeOfferingOutput, error) 
 DescribeReservation(ctx context.Context, params *DescribeReservationInput, optFns ...func(*Options)) (*DescribeReservationOutput, error) 
 DescribeSchedule(ctx context.Context, params *DescribeScheduleInput, optFns ...func(*Options)) (*DescribeScheduleOutput, error) 
 DescribeThumbnails(ctx context.Context, params *DescribeThumbnailsInput, optFns ...func(*Options)) (*DescribeThumbnailsOutput, error) 
 GetCloudWatchAlarmTemplate(ctx context.Context, params *GetCloudWatchAlarmTemplateInput, optFns ...func(*Options)) (*GetCloudWatchAlarmTemplateOutput, error) 
 GetCloudWatchAlarmTemplateGroup(ctx context.Context, params *GetCloudWatchAlarmTemplateGroupInput, optFns ...func(*Options)) (*GetCloudWatchAlarmTemplateGroupOutput, error) 
 GetEventBridgeRuleTemplate(ctx context.Context, params *GetEventBridgeRuleTemplateInput, optFns ...func(*Options)) (*GetEventBridgeRuleTemplateOutput, error) 
 GetEventBridgeRuleTemplateGroup(ctx context.Context, params *GetEventBridgeRuleTemplateGroupInput, optFns ...func(*Options)) (*GetEventBridgeRuleTemplateGroupOutput, error) 
 GetSignalMap(ctx context.Context, params *GetSignalMapInput, optFns ...func(*Options)) (*GetSignalMapOutput, error) 
 ListChannels(ctx context.Context, params *ListChannelsInput, optFns ...func(*Options)) (*ListChannelsOutput, error) 
 ListCloudWatchAlarmTemplateGroups(ctx context.Context, params *ListCloudWatchAlarmTemplateGroupsInput, optFns ...func(*Options)) (*ListCloudWatchAlarmTemplateGroupsOutput, error) 
 ListCloudWatchAlarmTemplates(ctx context.Context, params *ListCloudWatchAlarmTemplatesInput, optFns ...func(*Options)) (*ListCloudWatchAlarmTemplatesOutput, error) 
 ListEventBridgeRuleTemplateGroups(ctx context.Context, params *ListEventBridgeRuleTemplateGroupsInput, optFns ...func(*Options)) (*ListEventBridgeRuleTemplateGroupsOutput, error) 
 ListEventBridgeRuleTemplates(ctx context.Context, params *ListEventBridgeRuleTemplatesInput, optFns ...func(*Options)) (*ListEventBridgeRuleTemplatesOutput, error) 
 ListInputDeviceTransfers(ctx context.Context, params *ListInputDeviceTransfersInput, optFns ...func(*Options)) (*ListInputDeviceTransfersOutput, error) 
 ListInputDevices(ctx context.Context, params *ListInputDevicesInput, optFns ...func(*Options)) (*ListInputDevicesOutput, error) 
 ListInputSecurityGroups(ctx context.Context, params *ListInputSecurityGroupsInput, optFns ...func(*Options)) (*ListInputSecurityGroupsOutput, error) 
 ListInputs(ctx context.Context, params *ListInputsInput, optFns ...func(*Options)) (*ListInputsOutput, error) 
 ListMultiplexPrograms(ctx context.Context, params *ListMultiplexProgramsInput, optFns ...func(*Options)) (*ListMultiplexProgramsOutput, error) 
 ListMultiplexes(ctx context.Context, params *ListMultiplexesInput, optFns ...func(*Options)) (*ListMultiplexesOutput, error) 
 ListOfferings(ctx context.Context, params *ListOfferingsInput, optFns ...func(*Options)) (*ListOfferingsOutput, error) 
 ListReservations(ctx context.Context, params *ListReservationsInput, optFns ...func(*Options)) (*ListReservationsOutput, error) 
 ListSignalMaps(ctx context.Context, params *ListSignalMapsInput, optFns ...func(*Options)) (*ListSignalMapsOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 PurchaseOffering(ctx context.Context, params *PurchaseOfferingInput, optFns ...func(*Options)) (*PurchaseOfferingOutput, error) 
 RebootInputDevice(ctx context.Context, params *RebootInputDeviceInput, optFns ...func(*Options)) (*RebootInputDeviceOutput, error) 
 RejectInputDeviceTransfer(ctx context.Context, params *RejectInputDeviceTransferInput, optFns ...func(*Options)) (*RejectInputDeviceTransferOutput, error) 
 RestartChannelPipelines(ctx context.Context, params *RestartChannelPipelinesInput, optFns ...func(*Options)) (*RestartChannelPipelinesOutput, error) 
 StartChannel(ctx context.Context, params *StartChannelInput, optFns ...func(*Options)) (*StartChannelOutput, error) 
 StartDeleteMonitorDeployment(ctx context.Context, params *StartDeleteMonitorDeploymentInput, optFns ...func(*Options)) (*StartDeleteMonitorDeploymentOutput, error) 
 StartInputDevice(ctx context.Context, params *StartInputDeviceInput, optFns ...func(*Options)) (*StartInputDeviceOutput, error) 
 StartInputDeviceMaintenanceWindow(ctx context.Context, params *StartInputDeviceMaintenanceWindowInput, optFns ...func(*Options)) (*StartInputDeviceMaintenanceWindowOutput, error) 
 StartMonitorDeployment(ctx context.Context, params *StartMonitorDeploymentInput, optFns ...func(*Options)) (*StartMonitorDeploymentOutput, error) 
 StartMultiplex(ctx context.Context, params *StartMultiplexInput, optFns ...func(*Options)) (*StartMultiplexOutput, error) 
 StartUpdateSignalMap(ctx context.Context, params *StartUpdateSignalMapInput, optFns ...func(*Options)) (*StartUpdateSignalMapOutput, error) 
 StopChannel(ctx context.Context, params *StopChannelInput, optFns ...func(*Options)) (*StopChannelOutput, error) 
 StopInputDevice(ctx context.Context, params *StopInputDeviceInput, optFns ...func(*Options)) (*StopInputDeviceOutput, error) 
 StopMultiplex(ctx context.Context, params *StopMultiplexInput, optFns ...func(*Options)) (*StopMultiplexOutput, error) 
 TransferInputDevice(ctx context.Context, params *TransferInputDeviceInput, optFns ...func(*Options)) (*TransferInputDeviceOutput, error) 
 UpdateAccountConfiguration(ctx context.Context, params *UpdateAccountConfigurationInput, optFns ...func(*Options)) (*UpdateAccountConfigurationOutput, error) 
 UpdateChannel(ctx context.Context, params *UpdateChannelInput, optFns ...func(*Options)) (*UpdateChannelOutput, error) 
 UpdateChannelClass(ctx context.Context, params *UpdateChannelClassInput, optFns ...func(*Options)) (*UpdateChannelClassOutput, error) 
 UpdateCloudWatchAlarmTemplate(ctx context.Context, params *UpdateCloudWatchAlarmTemplateInput, optFns ...func(*Options)) (*UpdateCloudWatchAlarmTemplateOutput, error) 
 UpdateCloudWatchAlarmTemplateGroup(ctx context.Context, params *UpdateCloudWatchAlarmTemplateGroupInput, optFns ...func(*Options)) (*UpdateCloudWatchAlarmTemplateGroupOutput, error) 
 UpdateEventBridgeRuleTemplate(ctx context.Context, params *UpdateEventBridgeRuleTemplateInput, optFns ...func(*Options)) (*UpdateEventBridgeRuleTemplateOutput, error) 
 UpdateEventBridgeRuleTemplateGroup(ctx context.Context, params *UpdateEventBridgeRuleTemplateGroupInput, optFns ...func(*Options)) (*UpdateEventBridgeRuleTemplateGroupOutput, error) 
 UpdateInput(ctx context.Context, params *UpdateInputInput, optFns ...func(*Options)) (*UpdateInputOutput, error) 
 UpdateInputDevice(ctx context.Context, params *UpdateInputDeviceInput, optFns ...func(*Options)) (*UpdateInputDeviceOutput, error) 
 UpdateInputSecurityGroup(ctx context.Context, params *UpdateInputSecurityGroupInput, optFns ...func(*Options)) (*UpdateInputSecurityGroupOutput, error) 
 UpdateMultiplex(ctx context.Context, params *UpdateMultiplexInput, optFns ...func(*Options)) (*UpdateMultiplexOutput, error) 
 UpdateMultiplexProgram(ctx context.Context, params *UpdateMultiplexProgramInput, optFns ...func(*Options)) (*UpdateMultiplexProgramOutput, error) 
 UpdateReservation(ctx context.Context, params *UpdateReservationInput, optFns ...func(*Options)) (*UpdateReservationOutput, error) 
}
