
package backupgateway

import (
    "github.com/aws/aws-sdk-go-v2/service/backupgateway"
)

// IBackupgateway defines the interface for backupgateway
type IBackupgateway interface {
 Options() Options 
 AssociateGatewayToServer(ctx context.Context, params *AssociateGatewayToServerInput, optFns ...func(*Options)) (*AssociateGatewayToServerOutput, error) 
 CreateGateway(ctx context.Context, params *CreateGatewayInput, optFns ...func(*Options)) (*CreateGatewayOutput, error) 
 DeleteGateway(ctx context.Context, params *DeleteGatewayInput, optFns ...func(*Options)) (*DeleteGatewayOutput, error) 
 DeleteHypervisor(ctx context.Context, params *DeleteHypervisorInput, optFns ...func(*Options)) (*DeleteHypervisorOutput, error) 
 DisassociateGatewayFromServer(ctx context.Context, params *DisassociateGatewayFromServerInput, optFns ...func(*Options)) (*DisassociateGatewayFromServerOutput, error) 
 GetBandwidthRateLimitSchedule(ctx context.Context, params *GetBandwidthRateLimitScheduleInput, optFns ...func(*Options)) (*GetBandwidthRateLimitScheduleOutput, error) 
 GetGateway(ctx context.Context, params *GetGatewayInput, optFns ...func(*Options)) (*GetGatewayOutput, error) 
 GetHypervisor(ctx context.Context, params *GetHypervisorInput, optFns ...func(*Options)) (*GetHypervisorOutput, error) 
 GetHypervisorPropertyMappings(ctx context.Context, params *GetHypervisorPropertyMappingsInput, optFns ...func(*Options)) (*GetHypervisorPropertyMappingsOutput, error) 
 GetVirtualMachine(ctx context.Context, params *GetVirtualMachineInput, optFns ...func(*Options)) (*GetVirtualMachineOutput, error) 
 ImportHypervisorConfiguration(ctx context.Context, params *ImportHypervisorConfigurationInput, optFns ...func(*Options)) (*ImportHypervisorConfigurationOutput, error) 
 ListGateways(ctx context.Context, params *ListGatewaysInput, optFns ...func(*Options)) (*ListGatewaysOutput, error) 
 ListHypervisors(ctx context.Context, params *ListHypervisorsInput, optFns ...func(*Options)) (*ListHypervisorsOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 ListVirtualMachines(ctx context.Context, params *ListVirtualMachinesInput, optFns ...func(*Options)) (*ListVirtualMachinesOutput, error) 
 PutBandwidthRateLimitSchedule(ctx context.Context, params *PutBandwidthRateLimitScheduleInput, optFns ...func(*Options)) (*PutBandwidthRateLimitScheduleOutput, error) 
 PutHypervisorPropertyMappings(ctx context.Context, params *PutHypervisorPropertyMappingsInput, optFns ...func(*Options)) (*PutHypervisorPropertyMappingsOutput, error) 
 PutMaintenanceStartTime(ctx context.Context, params *PutMaintenanceStartTimeInput, optFns ...func(*Options)) (*PutMaintenanceStartTimeOutput, error) 
 StartVirtualMachinesMetadataSync(ctx context.Context, params *StartVirtualMachinesMetadataSyncInput, optFns ...func(*Options)) (*StartVirtualMachinesMetadataSyncOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 TestHypervisorConfiguration(ctx context.Context, params *TestHypervisorConfigurationInput, optFns ...func(*Options)) (*TestHypervisorConfigurationOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateGatewayInformation(ctx context.Context, params *UpdateGatewayInformationInput, optFns ...func(*Options)) (*UpdateGatewayInformationOutput, error) 
 UpdateGatewaySoftwareNow(ctx context.Context, params *UpdateGatewaySoftwareNowInput, optFns ...func(*Options)) (*UpdateGatewaySoftwareNowOutput, error) 
 UpdateHypervisor(ctx context.Context, params *UpdateHypervisorInput, optFns ...func(*Options)) (*UpdateHypervisorOutput, error) 
}
