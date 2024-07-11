
package groundstation

import (
    "github.com/aws/aws-sdk-go-v2/service/groundstation"
)

// IGroundstation defines the interface for groundstation
type IGroundstation interface {
 Options() Options 
 CancelContact(ctx context.Context, params *CancelContactInput, optFns ...func(*Options)) (*CancelContactOutput, error) 
 CreateConfig(ctx context.Context, params *CreateConfigInput, optFns ...func(*Options)) (*CreateConfigOutput, error) 
 CreateDataflowEndpointGroup(ctx context.Context, params *CreateDataflowEndpointGroupInput, optFns ...func(*Options)) (*CreateDataflowEndpointGroupOutput, error) 
 CreateEphemeris(ctx context.Context, params *CreateEphemerisInput, optFns ...func(*Options)) (*CreateEphemerisOutput, error) 
 CreateMissionProfile(ctx context.Context, params *CreateMissionProfileInput, optFns ...func(*Options)) (*CreateMissionProfileOutput, error) 
 DeleteConfig(ctx context.Context, params *DeleteConfigInput, optFns ...func(*Options)) (*DeleteConfigOutput, error) 
 DeleteDataflowEndpointGroup(ctx context.Context, params *DeleteDataflowEndpointGroupInput, optFns ...func(*Options)) (*DeleteDataflowEndpointGroupOutput, error) 
 DeleteEphemeris(ctx context.Context, params *DeleteEphemerisInput, optFns ...func(*Options)) (*DeleteEphemerisOutput, error) 
 DeleteMissionProfile(ctx context.Context, params *DeleteMissionProfileInput, optFns ...func(*Options)) (*DeleteMissionProfileOutput, error) 
 DescribeContact(ctx context.Context, params *DescribeContactInput, optFns ...func(*Options)) (*DescribeContactOutput, error) 
 DescribeEphemeris(ctx context.Context, params *DescribeEphemerisInput, optFns ...func(*Options)) (*DescribeEphemerisOutput, error) 
 GetAgentConfiguration(ctx context.Context, params *GetAgentConfigurationInput, optFns ...func(*Options)) (*GetAgentConfigurationOutput, error) 
 GetConfig(ctx context.Context, params *GetConfigInput, optFns ...func(*Options)) (*GetConfigOutput, error) 
 GetDataflowEndpointGroup(ctx context.Context, params *GetDataflowEndpointGroupInput, optFns ...func(*Options)) (*GetDataflowEndpointGroupOutput, error) 
 GetMinuteUsage(ctx context.Context, params *GetMinuteUsageInput, optFns ...func(*Options)) (*GetMinuteUsageOutput, error) 
 GetMissionProfile(ctx context.Context, params *GetMissionProfileInput, optFns ...func(*Options)) (*GetMissionProfileOutput, error) 
 GetSatellite(ctx context.Context, params *GetSatelliteInput, optFns ...func(*Options)) (*GetSatelliteOutput, error) 
 ListConfigs(ctx context.Context, params *ListConfigsInput, optFns ...func(*Options)) (*ListConfigsOutput, error) 
 ListContacts(ctx context.Context, params *ListContactsInput, optFns ...func(*Options)) (*ListContactsOutput, error) 
 ListDataflowEndpointGroups(ctx context.Context, params *ListDataflowEndpointGroupsInput, optFns ...func(*Options)) (*ListDataflowEndpointGroupsOutput, error) 
 ListEphemerides(ctx context.Context, params *ListEphemeridesInput, optFns ...func(*Options)) (*ListEphemeridesOutput, error) 
 ListGroundStations(ctx context.Context, params *ListGroundStationsInput, optFns ...func(*Options)) (*ListGroundStationsOutput, error) 
 ListMissionProfiles(ctx context.Context, params *ListMissionProfilesInput, optFns ...func(*Options)) (*ListMissionProfilesOutput, error) 
 ListSatellites(ctx context.Context, params *ListSatellitesInput, optFns ...func(*Options)) (*ListSatellitesOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 RegisterAgent(ctx context.Context, params *RegisterAgentInput, optFns ...func(*Options)) (*RegisterAgentOutput, error) 
 ReserveContact(ctx context.Context, params *ReserveContactInput, optFns ...func(*Options)) (*ReserveContactOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateAgentStatus(ctx context.Context, params *UpdateAgentStatusInput, optFns ...func(*Options)) (*UpdateAgentStatusOutput, error) 
 UpdateConfig(ctx context.Context, params *UpdateConfigInput, optFns ...func(*Options)) (*UpdateConfigOutput, error) 
 UpdateEphemeris(ctx context.Context, params *UpdateEphemerisInput, optFns ...func(*Options)) (*UpdateEphemerisOutput, error) 
 UpdateMissionProfile(ctx context.Context, params *UpdateMissionProfileInput, optFns ...func(*Options)) (*UpdateMissionProfileOutput, error) 
}
