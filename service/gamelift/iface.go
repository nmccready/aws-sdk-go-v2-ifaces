
package gamelift

// This file was generated by scripts/gen_ifaces.sh do not edit it manually
import (
    "github.com/aws/aws-sdk-go-v2/service/gamelift"
)

// IClient defines the interface for gamelift
type IClient interface {
 Options() Options 
 AcceptMatch(ctx context.Context, params *AcceptMatchInput, optFns ...func(*Options)) (*AcceptMatchOutput, error) 
 ClaimGameServer(ctx context.Context, params *ClaimGameServerInput, optFns ...func(*Options)) (*ClaimGameServerOutput, error) 
 CreateAlias(ctx context.Context, params *CreateAliasInput, optFns ...func(*Options)) (*CreateAliasOutput, error) 
 CreateBuild(ctx context.Context, params *CreateBuildInput, optFns ...func(*Options)) (*CreateBuildOutput, error) 
 CreateContainerGroupDefinition(ctx context.Context, params *CreateContainerGroupDefinitionInput, optFns ...func(*Options)) (*CreateContainerGroupDefinitionOutput, error) 
 CreateFleet(ctx context.Context, params *CreateFleetInput, optFns ...func(*Options)) (*CreateFleetOutput, error) 
 CreateFleetLocations(ctx context.Context, params *CreateFleetLocationsInput, optFns ...func(*Options)) (*CreateFleetLocationsOutput, error) 
 CreateGameServerGroup(ctx context.Context, params *CreateGameServerGroupInput, optFns ...func(*Options)) (*CreateGameServerGroupOutput, error) 
 CreateGameSession(ctx context.Context, params *CreateGameSessionInput, optFns ...func(*Options)) (*CreateGameSessionOutput, error) 
 CreateGameSessionQueue(ctx context.Context, params *CreateGameSessionQueueInput, optFns ...func(*Options)) (*CreateGameSessionQueueOutput, error) 
 CreateLocation(ctx context.Context, params *CreateLocationInput, optFns ...func(*Options)) (*CreateLocationOutput, error) 
 CreateMatchmakingConfiguration(ctx context.Context, params *CreateMatchmakingConfigurationInput, optFns ...func(*Options)) (*CreateMatchmakingConfigurationOutput, error) 
 CreateMatchmakingRuleSet(ctx context.Context, params *CreateMatchmakingRuleSetInput, optFns ...func(*Options)) (*CreateMatchmakingRuleSetOutput, error) 
 CreatePlayerSession(ctx context.Context, params *CreatePlayerSessionInput, optFns ...func(*Options)) (*CreatePlayerSessionOutput, error) 
 CreatePlayerSessions(ctx context.Context, params *CreatePlayerSessionsInput, optFns ...func(*Options)) (*CreatePlayerSessionsOutput, error) 
 CreateScript(ctx context.Context, params *CreateScriptInput, optFns ...func(*Options)) (*CreateScriptOutput, error) 
 CreateVpcPeeringAuthorization(ctx context.Context, params *CreateVpcPeeringAuthorizationInput, optFns ...func(*Options)) (*CreateVpcPeeringAuthorizationOutput, error) 
 CreateVpcPeeringConnection(ctx context.Context, params *CreateVpcPeeringConnectionInput, optFns ...func(*Options)) (*CreateVpcPeeringConnectionOutput, error) 
 DeleteAlias(ctx context.Context, params *DeleteAliasInput, optFns ...func(*Options)) (*DeleteAliasOutput, error) 
 DeleteBuild(ctx context.Context, params *DeleteBuildInput, optFns ...func(*Options)) (*DeleteBuildOutput, error) 
 DeleteContainerGroupDefinition(ctx context.Context, params *DeleteContainerGroupDefinitionInput, optFns ...func(*Options)) (*DeleteContainerGroupDefinitionOutput, error) 
 DeleteFleet(ctx context.Context, params *DeleteFleetInput, optFns ...func(*Options)) (*DeleteFleetOutput, error) 
 DeleteFleetLocations(ctx context.Context, params *DeleteFleetLocationsInput, optFns ...func(*Options)) (*DeleteFleetLocationsOutput, error) 
 DeleteGameServerGroup(ctx context.Context, params *DeleteGameServerGroupInput, optFns ...func(*Options)) (*DeleteGameServerGroupOutput, error) 
 DeleteGameSessionQueue(ctx context.Context, params *DeleteGameSessionQueueInput, optFns ...func(*Options)) (*DeleteGameSessionQueueOutput, error) 
 DeleteLocation(ctx context.Context, params *DeleteLocationInput, optFns ...func(*Options)) (*DeleteLocationOutput, error) 
 DeleteMatchmakingConfiguration(ctx context.Context, params *DeleteMatchmakingConfigurationInput, optFns ...func(*Options)) (*DeleteMatchmakingConfigurationOutput, error) 
 DeleteMatchmakingRuleSet(ctx context.Context, params *DeleteMatchmakingRuleSetInput, optFns ...func(*Options)) (*DeleteMatchmakingRuleSetOutput, error) 
 DeleteScalingPolicy(ctx context.Context, params *DeleteScalingPolicyInput, optFns ...func(*Options)) (*DeleteScalingPolicyOutput, error) 
 DeleteScript(ctx context.Context, params *DeleteScriptInput, optFns ...func(*Options)) (*DeleteScriptOutput, error) 
 DeleteVpcPeeringAuthorization(ctx context.Context, params *DeleteVpcPeeringAuthorizationInput, optFns ...func(*Options)) (*DeleteVpcPeeringAuthorizationOutput, error) 
 DeleteVpcPeeringConnection(ctx context.Context, params *DeleteVpcPeeringConnectionInput, optFns ...func(*Options)) (*DeleteVpcPeeringConnectionOutput, error) 
 DeregisterCompute(ctx context.Context, params *DeregisterComputeInput, optFns ...func(*Options)) (*DeregisterComputeOutput, error) 
 DeregisterGameServer(ctx context.Context, params *DeregisterGameServerInput, optFns ...func(*Options)) (*DeregisterGameServerOutput, error) 
 DescribeAlias(ctx context.Context, params *DescribeAliasInput, optFns ...func(*Options)) (*DescribeAliasOutput, error) 
 DescribeBuild(ctx context.Context, params *DescribeBuildInput, optFns ...func(*Options)) (*DescribeBuildOutput, error) 
 DescribeCompute(ctx context.Context, params *DescribeComputeInput, optFns ...func(*Options)) (*DescribeComputeOutput, error) 
 DescribeContainerGroupDefinition(ctx context.Context, params *DescribeContainerGroupDefinitionInput, optFns ...func(*Options)) (*DescribeContainerGroupDefinitionOutput, error) 
 DescribeEC2InstanceLimits(ctx context.Context, params *DescribeEC2InstanceLimitsInput, optFns ...func(*Options)) (*DescribeEC2InstanceLimitsOutput, error) 
 DescribeFleetAttributes(ctx context.Context, params *DescribeFleetAttributesInput, optFns ...func(*Options)) (*DescribeFleetAttributesOutput, error) 
 DescribeFleetCapacity(ctx context.Context, params *DescribeFleetCapacityInput, optFns ...func(*Options)) (*DescribeFleetCapacityOutput, error) 
 DescribeFleetEvents(ctx context.Context, params *DescribeFleetEventsInput, optFns ...func(*Options)) (*DescribeFleetEventsOutput, error) 
 DescribeFleetLocationAttributes(ctx context.Context, params *DescribeFleetLocationAttributesInput, optFns ...func(*Options)) (*DescribeFleetLocationAttributesOutput, error) 
 DescribeFleetLocationCapacity(ctx context.Context, params *DescribeFleetLocationCapacityInput, optFns ...func(*Options)) (*DescribeFleetLocationCapacityOutput, error) 
 DescribeFleetLocationUtilization(ctx context.Context, params *DescribeFleetLocationUtilizationInput, optFns ...func(*Options)) (*DescribeFleetLocationUtilizationOutput, error) 
 DescribeFleetPortSettings(ctx context.Context, params *DescribeFleetPortSettingsInput, optFns ...func(*Options)) (*DescribeFleetPortSettingsOutput, error) 
 DescribeFleetUtilization(ctx context.Context, params *DescribeFleetUtilizationInput, optFns ...func(*Options)) (*DescribeFleetUtilizationOutput, error) 
 DescribeGameServer(ctx context.Context, params *DescribeGameServerInput, optFns ...func(*Options)) (*DescribeGameServerOutput, error) 
 DescribeGameServerGroup(ctx context.Context, params *DescribeGameServerGroupInput, optFns ...func(*Options)) (*DescribeGameServerGroupOutput, error) 
 DescribeGameServerInstances(ctx context.Context, params *DescribeGameServerInstancesInput, optFns ...func(*Options)) (*DescribeGameServerInstancesOutput, error) 
 DescribeGameSessionDetails(ctx context.Context, params *DescribeGameSessionDetailsInput, optFns ...func(*Options)) (*DescribeGameSessionDetailsOutput, error) 
 DescribeGameSessionPlacement(ctx context.Context, params *DescribeGameSessionPlacementInput, optFns ...func(*Options)) (*DescribeGameSessionPlacementOutput, error) 
 DescribeGameSessionQueues(ctx context.Context, params *DescribeGameSessionQueuesInput, optFns ...func(*Options)) (*DescribeGameSessionQueuesOutput, error) 
 DescribeGameSessions(ctx context.Context, params *DescribeGameSessionsInput, optFns ...func(*Options)) (*DescribeGameSessionsOutput, error) 
 DescribeInstances(ctx context.Context, params *DescribeInstancesInput, optFns ...func(*Options)) (*DescribeInstancesOutput, error) 
 DescribeMatchmaking(ctx context.Context, params *DescribeMatchmakingInput, optFns ...func(*Options)) (*DescribeMatchmakingOutput, error) 
 DescribeMatchmakingConfigurations(ctx context.Context, params *DescribeMatchmakingConfigurationsInput, optFns ...func(*Options)) (*DescribeMatchmakingConfigurationsOutput, error) 
 DescribeMatchmakingRuleSets(ctx context.Context, params *DescribeMatchmakingRuleSetsInput, optFns ...func(*Options)) (*DescribeMatchmakingRuleSetsOutput, error) 
 DescribePlayerSessions(ctx context.Context, params *DescribePlayerSessionsInput, optFns ...func(*Options)) (*DescribePlayerSessionsOutput, error) 
 DescribeRuntimeConfiguration(ctx context.Context, params *DescribeRuntimeConfigurationInput, optFns ...func(*Options)) (*DescribeRuntimeConfigurationOutput, error) 
 DescribeScalingPolicies(ctx context.Context, params *DescribeScalingPoliciesInput, optFns ...func(*Options)) (*DescribeScalingPoliciesOutput, error) 
 DescribeScript(ctx context.Context, params *DescribeScriptInput, optFns ...func(*Options)) (*DescribeScriptOutput, error) 
 DescribeVpcPeeringAuthorizations(ctx context.Context, params *DescribeVpcPeeringAuthorizationsInput, optFns ...func(*Options)) (*DescribeVpcPeeringAuthorizationsOutput, error) 
 DescribeVpcPeeringConnections(ctx context.Context, params *DescribeVpcPeeringConnectionsInput, optFns ...func(*Options)) (*DescribeVpcPeeringConnectionsOutput, error) 
 GetComputeAccess(ctx context.Context, params *GetComputeAccessInput, optFns ...func(*Options)) (*GetComputeAccessOutput, error) 
 GetComputeAuthToken(ctx context.Context, params *GetComputeAuthTokenInput, optFns ...func(*Options)) (*GetComputeAuthTokenOutput, error) 
 GetGameSessionLogUrl(ctx context.Context, params *GetGameSessionLogUrlInput, optFns ...func(*Options)) (*GetGameSessionLogUrlOutput, error) 
 GetInstanceAccess(ctx context.Context, params *GetInstanceAccessInput, optFns ...func(*Options)) (*GetInstanceAccessOutput, error) 
 ListAliases(ctx context.Context, params *ListAliasesInput, optFns ...func(*Options)) (*ListAliasesOutput, error) 
 ListBuilds(ctx context.Context, params *ListBuildsInput, optFns ...func(*Options)) (*ListBuildsOutput, error) 
 ListCompute(ctx context.Context, params *ListComputeInput, optFns ...func(*Options)) (*ListComputeOutput, error) 
 ListContainerGroupDefinitions(ctx context.Context, params *ListContainerGroupDefinitionsInput, optFns ...func(*Options)) (*ListContainerGroupDefinitionsOutput, error) 
 ListFleets(ctx context.Context, params *ListFleetsInput, optFns ...func(*Options)) (*ListFleetsOutput, error) 
 ListGameServerGroups(ctx context.Context, params *ListGameServerGroupsInput, optFns ...func(*Options)) (*ListGameServerGroupsOutput, error) 
 ListGameServers(ctx context.Context, params *ListGameServersInput, optFns ...func(*Options)) (*ListGameServersOutput, error) 
 ListLocations(ctx context.Context, params *ListLocationsInput, optFns ...func(*Options)) (*ListLocationsOutput, error) 
 ListScripts(ctx context.Context, params *ListScriptsInput, optFns ...func(*Options)) (*ListScriptsOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 PutScalingPolicy(ctx context.Context, params *PutScalingPolicyInput, optFns ...func(*Options)) (*PutScalingPolicyOutput, error) 
 RegisterCompute(ctx context.Context, params *RegisterComputeInput, optFns ...func(*Options)) (*RegisterComputeOutput, error) 
 RegisterGameServer(ctx context.Context, params *RegisterGameServerInput, optFns ...func(*Options)) (*RegisterGameServerOutput, error) 
 RequestUploadCredentials(ctx context.Context, params *RequestUploadCredentialsInput, optFns ...func(*Options)) (*RequestUploadCredentialsOutput, error) 
 ResolveAlias(ctx context.Context, params *ResolveAliasInput, optFns ...func(*Options)) (*ResolveAliasOutput, error) 
 ResumeGameServerGroup(ctx context.Context, params *ResumeGameServerGroupInput, optFns ...func(*Options)) (*ResumeGameServerGroupOutput, error) 
 SearchGameSessions(ctx context.Context, params *SearchGameSessionsInput, optFns ...func(*Options)) (*SearchGameSessionsOutput, error) 
 StartFleetActions(ctx context.Context, params *StartFleetActionsInput, optFns ...func(*Options)) (*StartFleetActionsOutput, error) 
 StartGameSessionPlacement(ctx context.Context, params *StartGameSessionPlacementInput, optFns ...func(*Options)) (*StartGameSessionPlacementOutput, error) 
 StartMatchBackfill(ctx context.Context, params *StartMatchBackfillInput, optFns ...func(*Options)) (*StartMatchBackfillOutput, error) 
 StartMatchmaking(ctx context.Context, params *StartMatchmakingInput, optFns ...func(*Options)) (*StartMatchmakingOutput, error) 
 StopFleetActions(ctx context.Context, params *StopFleetActionsInput, optFns ...func(*Options)) (*StopFleetActionsOutput, error) 
 StopGameSessionPlacement(ctx context.Context, params *StopGameSessionPlacementInput, optFns ...func(*Options)) (*StopGameSessionPlacementOutput, error) 
 StopMatchmaking(ctx context.Context, params *StopMatchmakingInput, optFns ...func(*Options)) (*StopMatchmakingOutput, error) 
 SuspendGameServerGroup(ctx context.Context, params *SuspendGameServerGroupInput, optFns ...func(*Options)) (*SuspendGameServerGroupOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateAlias(ctx context.Context, params *UpdateAliasInput, optFns ...func(*Options)) (*UpdateAliasOutput, error) 
 UpdateBuild(ctx context.Context, params *UpdateBuildInput, optFns ...func(*Options)) (*UpdateBuildOutput, error) 
 UpdateFleetAttributes(ctx context.Context, params *UpdateFleetAttributesInput, optFns ...func(*Options)) (*UpdateFleetAttributesOutput, error) 
 UpdateFleetCapacity(ctx context.Context, params *UpdateFleetCapacityInput, optFns ...func(*Options)) (*UpdateFleetCapacityOutput, error) 
 UpdateFleetPortSettings(ctx context.Context, params *UpdateFleetPortSettingsInput, optFns ...func(*Options)) (*UpdateFleetPortSettingsOutput, error) 
 UpdateGameServer(ctx context.Context, params *UpdateGameServerInput, optFns ...func(*Options)) (*UpdateGameServerOutput, error) 
 UpdateGameServerGroup(ctx context.Context, params *UpdateGameServerGroupInput, optFns ...func(*Options)) (*UpdateGameServerGroupOutput, error) 
 UpdateGameSession(ctx context.Context, params *UpdateGameSessionInput, optFns ...func(*Options)) (*UpdateGameSessionOutput, error) 
 UpdateGameSessionQueue(ctx context.Context, params *UpdateGameSessionQueueInput, optFns ...func(*Options)) (*UpdateGameSessionQueueOutput, error) 
 UpdateMatchmakingConfiguration(ctx context.Context, params *UpdateMatchmakingConfigurationInput, optFns ...func(*Options)) (*UpdateMatchmakingConfigurationOutput, error) 
 UpdateRuntimeConfiguration(ctx context.Context, params *UpdateRuntimeConfigurationInput, optFns ...func(*Options)) (*UpdateRuntimeConfigurationOutput, error) 
 UpdateScript(ctx context.Context, params *UpdateScriptInput, optFns ...func(*Options)) (*UpdateScriptOutput, error) 
 ValidateMatchmakingRuleSet(ctx context.Context, params *ValidateMatchmakingRuleSetInput, optFns ...func(*Options)) (*ValidateMatchmakingRuleSetOutput, error) 
}
