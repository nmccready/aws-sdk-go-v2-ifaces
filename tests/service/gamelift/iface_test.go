package gamelift_test

// tests for the gamelift service interface for this ../../../service/gamelift/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/gamelift"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/gamelift/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/gamelift/gamelift_iface"
	"github.com/stretchr/testify/assert"
)

func TestGameliftServiceCanBeMocked(t *testing.T) {
	var iface gamelift_iface.IClient
	iface = &gamelift.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := gamelift.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAcceptMatch", func(t *testing.T) {
        input := &gamelift.AcceptMatchInput{}
        output := &gamelift.AcceptMatchOutput{}

        mockClient.On("AcceptMatch", ctx, input).Return(output, nil)

        result, err := mockClient.AcceptMatch(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestClaimGameServer", func(t *testing.T) {
        input := &gamelift.ClaimGameServerInput{}
        output := &gamelift.ClaimGameServerOutput{}

        mockClient.On("ClaimGameServer", ctx, input).Return(output, nil)

        result, err := mockClient.ClaimGameServer(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateAlias", func(t *testing.T) {
        input := &gamelift.CreateAliasInput{}
        output := &gamelift.CreateAliasOutput{}

        mockClient.On("CreateAlias", ctx, input).Return(output, nil)

        result, err := mockClient.CreateAlias(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateBuild", func(t *testing.T) {
        input := &gamelift.CreateBuildInput{}
        output := &gamelift.CreateBuildOutput{}

        mockClient.On("CreateBuild", ctx, input).Return(output, nil)

        result, err := mockClient.CreateBuild(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateContainerGroupDefinition", func(t *testing.T) {
        input := &gamelift.CreateContainerGroupDefinitionInput{}
        output := &gamelift.CreateContainerGroupDefinitionOutput{}

        mockClient.On("CreateContainerGroupDefinition", ctx, input).Return(output, nil)

        result, err := mockClient.CreateContainerGroupDefinition(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateFleet", func(t *testing.T) {
        input := &gamelift.CreateFleetInput{}
        output := &gamelift.CreateFleetOutput{}

        mockClient.On("CreateFleet", ctx, input).Return(output, nil)

        result, err := mockClient.CreateFleet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateFleetLocations", func(t *testing.T) {
        input := &gamelift.CreateFleetLocationsInput{}
        output := &gamelift.CreateFleetLocationsOutput{}

        mockClient.On("CreateFleetLocations", ctx, input).Return(output, nil)

        result, err := mockClient.CreateFleetLocations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateGameServerGroup", func(t *testing.T) {
        input := &gamelift.CreateGameServerGroupInput{}
        output := &gamelift.CreateGameServerGroupOutput{}

        mockClient.On("CreateGameServerGroup", ctx, input).Return(output, nil)

        result, err := mockClient.CreateGameServerGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateGameSession", func(t *testing.T) {
        input := &gamelift.CreateGameSessionInput{}
        output := &gamelift.CreateGameSessionOutput{}

        mockClient.On("CreateGameSession", ctx, input).Return(output, nil)

        result, err := mockClient.CreateGameSession(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateGameSessionQueue", func(t *testing.T) {
        input := &gamelift.CreateGameSessionQueueInput{}
        output := &gamelift.CreateGameSessionQueueOutput{}

        mockClient.On("CreateGameSessionQueue", ctx, input).Return(output, nil)

        result, err := mockClient.CreateGameSessionQueue(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateLocation", func(t *testing.T) {
        input := &gamelift.CreateLocationInput{}
        output := &gamelift.CreateLocationOutput{}

        mockClient.On("CreateLocation", ctx, input).Return(output, nil)

        result, err := mockClient.CreateLocation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateMatchmakingConfiguration", func(t *testing.T) {
        input := &gamelift.CreateMatchmakingConfigurationInput{}
        output := &gamelift.CreateMatchmakingConfigurationOutput{}

        mockClient.On("CreateMatchmakingConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.CreateMatchmakingConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateMatchmakingRuleSet", func(t *testing.T) {
        input := &gamelift.CreateMatchmakingRuleSetInput{}
        output := &gamelift.CreateMatchmakingRuleSetOutput{}

        mockClient.On("CreateMatchmakingRuleSet", ctx, input).Return(output, nil)

        result, err := mockClient.CreateMatchmakingRuleSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreatePlayerSession", func(t *testing.T) {
        input := &gamelift.CreatePlayerSessionInput{}
        output := &gamelift.CreatePlayerSessionOutput{}

        mockClient.On("CreatePlayerSession", ctx, input).Return(output, nil)

        result, err := mockClient.CreatePlayerSession(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreatePlayerSessions", func(t *testing.T) {
        input := &gamelift.CreatePlayerSessionsInput{}
        output := &gamelift.CreatePlayerSessionsOutput{}

        mockClient.On("CreatePlayerSessions", ctx, input).Return(output, nil)

        result, err := mockClient.CreatePlayerSessions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateScript", func(t *testing.T) {
        input := &gamelift.CreateScriptInput{}
        output := &gamelift.CreateScriptOutput{}

        mockClient.On("CreateScript", ctx, input).Return(output, nil)

        result, err := mockClient.CreateScript(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateVpcPeeringAuthorization", func(t *testing.T) {
        input := &gamelift.CreateVpcPeeringAuthorizationInput{}
        output := &gamelift.CreateVpcPeeringAuthorizationOutput{}

        mockClient.On("CreateVpcPeeringAuthorization", ctx, input).Return(output, nil)

        result, err := mockClient.CreateVpcPeeringAuthorization(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateVpcPeeringConnection", func(t *testing.T) {
        input := &gamelift.CreateVpcPeeringConnectionInput{}
        output := &gamelift.CreateVpcPeeringConnectionOutput{}

        mockClient.On("CreateVpcPeeringConnection", ctx, input).Return(output, nil)

        result, err := mockClient.CreateVpcPeeringConnection(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteAlias", func(t *testing.T) {
        input := &gamelift.DeleteAliasInput{}
        output := &gamelift.DeleteAliasOutput{}

        mockClient.On("DeleteAlias", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteAlias(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteBuild", func(t *testing.T) {
        input := &gamelift.DeleteBuildInput{}
        output := &gamelift.DeleteBuildOutput{}

        mockClient.On("DeleteBuild", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteBuild(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteContainerGroupDefinition", func(t *testing.T) {
        input := &gamelift.DeleteContainerGroupDefinitionInput{}
        output := &gamelift.DeleteContainerGroupDefinitionOutput{}

        mockClient.On("DeleteContainerGroupDefinition", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteContainerGroupDefinition(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteFleet", func(t *testing.T) {
        input := &gamelift.DeleteFleetInput{}
        output := &gamelift.DeleteFleetOutput{}

        mockClient.On("DeleteFleet", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteFleet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteFleetLocations", func(t *testing.T) {
        input := &gamelift.DeleteFleetLocationsInput{}
        output := &gamelift.DeleteFleetLocationsOutput{}

        mockClient.On("DeleteFleetLocations", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteFleetLocations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteGameServerGroup", func(t *testing.T) {
        input := &gamelift.DeleteGameServerGroupInput{}
        output := &gamelift.DeleteGameServerGroupOutput{}

        mockClient.On("DeleteGameServerGroup", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteGameServerGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteGameSessionQueue", func(t *testing.T) {
        input := &gamelift.DeleteGameSessionQueueInput{}
        output := &gamelift.DeleteGameSessionQueueOutput{}

        mockClient.On("DeleteGameSessionQueue", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteGameSessionQueue(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteLocation", func(t *testing.T) {
        input := &gamelift.DeleteLocationInput{}
        output := &gamelift.DeleteLocationOutput{}

        mockClient.On("DeleteLocation", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteLocation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteMatchmakingConfiguration", func(t *testing.T) {
        input := &gamelift.DeleteMatchmakingConfigurationInput{}
        output := &gamelift.DeleteMatchmakingConfigurationOutput{}

        mockClient.On("DeleteMatchmakingConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteMatchmakingConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteMatchmakingRuleSet", func(t *testing.T) {
        input := &gamelift.DeleteMatchmakingRuleSetInput{}
        output := &gamelift.DeleteMatchmakingRuleSetOutput{}

        mockClient.On("DeleteMatchmakingRuleSet", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteMatchmakingRuleSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteScalingPolicy", func(t *testing.T) {
        input := &gamelift.DeleteScalingPolicyInput{}
        output := &gamelift.DeleteScalingPolicyOutput{}

        mockClient.On("DeleteScalingPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteScalingPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteScript", func(t *testing.T) {
        input := &gamelift.DeleteScriptInput{}
        output := &gamelift.DeleteScriptOutput{}

        mockClient.On("DeleteScript", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteScript(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteVpcPeeringAuthorization", func(t *testing.T) {
        input := &gamelift.DeleteVpcPeeringAuthorizationInput{}
        output := &gamelift.DeleteVpcPeeringAuthorizationOutput{}

        mockClient.On("DeleteVpcPeeringAuthorization", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteVpcPeeringAuthorization(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteVpcPeeringConnection", func(t *testing.T) {
        input := &gamelift.DeleteVpcPeeringConnectionInput{}
        output := &gamelift.DeleteVpcPeeringConnectionOutput{}

        mockClient.On("DeleteVpcPeeringConnection", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteVpcPeeringConnection(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeregisterCompute", func(t *testing.T) {
        input := &gamelift.DeregisterComputeInput{}
        output := &gamelift.DeregisterComputeOutput{}

        mockClient.On("DeregisterCompute", ctx, input).Return(output, nil)

        result, err := mockClient.DeregisterCompute(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeregisterGameServer", func(t *testing.T) {
        input := &gamelift.DeregisterGameServerInput{}
        output := &gamelift.DeregisterGameServerOutput{}

        mockClient.On("DeregisterGameServer", ctx, input).Return(output, nil)

        result, err := mockClient.DeregisterGameServer(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeAlias", func(t *testing.T) {
        input := &gamelift.DescribeAliasInput{}
        output := &gamelift.DescribeAliasOutput{}

        mockClient.On("DescribeAlias", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeAlias(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeBuild", func(t *testing.T) {
        input := &gamelift.DescribeBuildInput{}
        output := &gamelift.DescribeBuildOutput{}

        mockClient.On("DescribeBuild", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeBuild(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeCompute", func(t *testing.T) {
        input := &gamelift.DescribeComputeInput{}
        output := &gamelift.DescribeComputeOutput{}

        mockClient.On("DescribeCompute", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeCompute(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeContainerGroupDefinition", func(t *testing.T) {
        input := &gamelift.DescribeContainerGroupDefinitionInput{}
        output := &gamelift.DescribeContainerGroupDefinitionOutput{}

        mockClient.On("DescribeContainerGroupDefinition", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeContainerGroupDefinition(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeEC2InstanceLimits", func(t *testing.T) {
        input := &gamelift.DescribeEC2InstanceLimitsInput{}
        output := &gamelift.DescribeEC2InstanceLimitsOutput{}

        mockClient.On("DescribeEC2InstanceLimits", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeEC2InstanceLimits(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeFleetAttributes", func(t *testing.T) {
        input := &gamelift.DescribeFleetAttributesInput{}
        output := &gamelift.DescribeFleetAttributesOutput{}

        mockClient.On("DescribeFleetAttributes", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeFleetAttributes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeFleetCapacity", func(t *testing.T) {
        input := &gamelift.DescribeFleetCapacityInput{}
        output := &gamelift.DescribeFleetCapacityOutput{}

        mockClient.On("DescribeFleetCapacity", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeFleetCapacity(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeFleetEvents", func(t *testing.T) {
        input := &gamelift.DescribeFleetEventsInput{}
        output := &gamelift.DescribeFleetEventsOutput{}

        mockClient.On("DescribeFleetEvents", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeFleetEvents(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeFleetLocationAttributes", func(t *testing.T) {
        input := &gamelift.DescribeFleetLocationAttributesInput{}
        output := &gamelift.DescribeFleetLocationAttributesOutput{}

        mockClient.On("DescribeFleetLocationAttributes", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeFleetLocationAttributes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeFleetLocationCapacity", func(t *testing.T) {
        input := &gamelift.DescribeFleetLocationCapacityInput{}
        output := &gamelift.DescribeFleetLocationCapacityOutput{}

        mockClient.On("DescribeFleetLocationCapacity", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeFleetLocationCapacity(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeFleetLocationUtilization", func(t *testing.T) {
        input := &gamelift.DescribeFleetLocationUtilizationInput{}
        output := &gamelift.DescribeFleetLocationUtilizationOutput{}

        mockClient.On("DescribeFleetLocationUtilization", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeFleetLocationUtilization(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeFleetPortSettings", func(t *testing.T) {
        input := &gamelift.DescribeFleetPortSettingsInput{}
        output := &gamelift.DescribeFleetPortSettingsOutput{}

        mockClient.On("DescribeFleetPortSettings", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeFleetPortSettings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeFleetUtilization", func(t *testing.T) {
        input := &gamelift.DescribeFleetUtilizationInput{}
        output := &gamelift.DescribeFleetUtilizationOutput{}

        mockClient.On("DescribeFleetUtilization", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeFleetUtilization(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeGameServer", func(t *testing.T) {
        input := &gamelift.DescribeGameServerInput{}
        output := &gamelift.DescribeGameServerOutput{}

        mockClient.On("DescribeGameServer", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeGameServer(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeGameServerGroup", func(t *testing.T) {
        input := &gamelift.DescribeGameServerGroupInput{}
        output := &gamelift.DescribeGameServerGroupOutput{}

        mockClient.On("DescribeGameServerGroup", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeGameServerGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeGameServerInstances", func(t *testing.T) {
        input := &gamelift.DescribeGameServerInstancesInput{}
        output := &gamelift.DescribeGameServerInstancesOutput{}

        mockClient.On("DescribeGameServerInstances", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeGameServerInstances(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeGameSessionDetails", func(t *testing.T) {
        input := &gamelift.DescribeGameSessionDetailsInput{}
        output := &gamelift.DescribeGameSessionDetailsOutput{}

        mockClient.On("DescribeGameSessionDetails", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeGameSessionDetails(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeGameSessionPlacement", func(t *testing.T) {
        input := &gamelift.DescribeGameSessionPlacementInput{}
        output := &gamelift.DescribeGameSessionPlacementOutput{}

        mockClient.On("DescribeGameSessionPlacement", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeGameSessionPlacement(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeGameSessionQueues", func(t *testing.T) {
        input := &gamelift.DescribeGameSessionQueuesInput{}
        output := &gamelift.DescribeGameSessionQueuesOutput{}

        mockClient.On("DescribeGameSessionQueues", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeGameSessionQueues(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeGameSessions", func(t *testing.T) {
        input := &gamelift.DescribeGameSessionsInput{}
        output := &gamelift.DescribeGameSessionsOutput{}

        mockClient.On("DescribeGameSessions", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeGameSessions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeInstances", func(t *testing.T) {
        input := &gamelift.DescribeInstancesInput{}
        output := &gamelift.DescribeInstancesOutput{}

        mockClient.On("DescribeInstances", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeInstances(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeMatchmaking", func(t *testing.T) {
        input := &gamelift.DescribeMatchmakingInput{}
        output := &gamelift.DescribeMatchmakingOutput{}

        mockClient.On("DescribeMatchmaking", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeMatchmaking(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeMatchmakingConfigurations", func(t *testing.T) {
        input := &gamelift.DescribeMatchmakingConfigurationsInput{}
        output := &gamelift.DescribeMatchmakingConfigurationsOutput{}

        mockClient.On("DescribeMatchmakingConfigurations", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeMatchmakingConfigurations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeMatchmakingRuleSets", func(t *testing.T) {
        input := &gamelift.DescribeMatchmakingRuleSetsInput{}
        output := &gamelift.DescribeMatchmakingRuleSetsOutput{}

        mockClient.On("DescribeMatchmakingRuleSets", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeMatchmakingRuleSets(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribePlayerSessions", func(t *testing.T) {
        input := &gamelift.DescribePlayerSessionsInput{}
        output := &gamelift.DescribePlayerSessionsOutput{}

        mockClient.On("DescribePlayerSessions", ctx, input).Return(output, nil)

        result, err := mockClient.DescribePlayerSessions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeRuntimeConfiguration", func(t *testing.T) {
        input := &gamelift.DescribeRuntimeConfigurationInput{}
        output := &gamelift.DescribeRuntimeConfigurationOutput{}

        mockClient.On("DescribeRuntimeConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeRuntimeConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeScalingPolicies", func(t *testing.T) {
        input := &gamelift.DescribeScalingPoliciesInput{}
        output := &gamelift.DescribeScalingPoliciesOutput{}

        mockClient.On("DescribeScalingPolicies", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeScalingPolicies(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeScript", func(t *testing.T) {
        input := &gamelift.DescribeScriptInput{}
        output := &gamelift.DescribeScriptOutput{}

        mockClient.On("DescribeScript", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeScript(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeVpcPeeringAuthorizations", func(t *testing.T) {
        input := &gamelift.DescribeVpcPeeringAuthorizationsInput{}
        output := &gamelift.DescribeVpcPeeringAuthorizationsOutput{}

        mockClient.On("DescribeVpcPeeringAuthorizations", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeVpcPeeringAuthorizations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeVpcPeeringConnections", func(t *testing.T) {
        input := &gamelift.DescribeVpcPeeringConnectionsInput{}
        output := &gamelift.DescribeVpcPeeringConnectionsOutput{}

        mockClient.On("DescribeVpcPeeringConnections", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeVpcPeeringConnections(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetComputeAccess", func(t *testing.T) {
        input := &gamelift.GetComputeAccessInput{}
        output := &gamelift.GetComputeAccessOutput{}

        mockClient.On("GetComputeAccess", ctx, input).Return(output, nil)

        result, err := mockClient.GetComputeAccess(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetComputeAuthToken", func(t *testing.T) {
        input := &gamelift.GetComputeAuthTokenInput{}
        output := &gamelift.GetComputeAuthTokenOutput{}

        mockClient.On("GetComputeAuthToken", ctx, input).Return(output, nil)

        result, err := mockClient.GetComputeAuthToken(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetGameSessionLogUrl", func(t *testing.T) {
        input := &gamelift.GetGameSessionLogUrlInput{}
        output := &gamelift.GetGameSessionLogUrlOutput{}

        mockClient.On("GetGameSessionLogUrl", ctx, input).Return(output, nil)

        result, err := mockClient.GetGameSessionLogUrl(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetInstanceAccess", func(t *testing.T) {
        input := &gamelift.GetInstanceAccessInput{}
        output := &gamelift.GetInstanceAccessOutput{}

        mockClient.On("GetInstanceAccess", ctx, input).Return(output, nil)

        result, err := mockClient.GetInstanceAccess(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAliases", func(t *testing.T) {
        input := &gamelift.ListAliasesInput{}
        output := &gamelift.ListAliasesOutput{}

        mockClient.On("ListAliases", ctx, input).Return(output, nil)

        result, err := mockClient.ListAliases(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListBuilds", func(t *testing.T) {
        input := &gamelift.ListBuildsInput{}
        output := &gamelift.ListBuildsOutput{}

        mockClient.On("ListBuilds", ctx, input).Return(output, nil)

        result, err := mockClient.ListBuilds(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListCompute", func(t *testing.T) {
        input := &gamelift.ListComputeInput{}
        output := &gamelift.ListComputeOutput{}

        mockClient.On("ListCompute", ctx, input).Return(output, nil)

        result, err := mockClient.ListCompute(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListContainerGroupDefinitions", func(t *testing.T) {
        input := &gamelift.ListContainerGroupDefinitionsInput{}
        output := &gamelift.ListContainerGroupDefinitionsOutput{}

        mockClient.On("ListContainerGroupDefinitions", ctx, input).Return(output, nil)

        result, err := mockClient.ListContainerGroupDefinitions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListFleets", func(t *testing.T) {
        input := &gamelift.ListFleetsInput{}
        output := &gamelift.ListFleetsOutput{}

        mockClient.On("ListFleets", ctx, input).Return(output, nil)

        result, err := mockClient.ListFleets(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListGameServerGroups", func(t *testing.T) {
        input := &gamelift.ListGameServerGroupsInput{}
        output := &gamelift.ListGameServerGroupsOutput{}

        mockClient.On("ListGameServerGroups", ctx, input).Return(output, nil)

        result, err := mockClient.ListGameServerGroups(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListGameServers", func(t *testing.T) {
        input := &gamelift.ListGameServersInput{}
        output := &gamelift.ListGameServersOutput{}

        mockClient.On("ListGameServers", ctx, input).Return(output, nil)

        result, err := mockClient.ListGameServers(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListLocations", func(t *testing.T) {
        input := &gamelift.ListLocationsInput{}
        output := &gamelift.ListLocationsOutput{}

        mockClient.On("ListLocations", ctx, input).Return(output, nil)

        result, err := mockClient.ListLocations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListScripts", func(t *testing.T) {
        input := &gamelift.ListScriptsInput{}
        output := &gamelift.ListScriptsOutput{}

        mockClient.On("ListScripts", ctx, input).Return(output, nil)

        result, err := mockClient.ListScripts(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &gamelift.ListTagsForResourceInput{}
        output := &gamelift.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutScalingPolicy", func(t *testing.T) {
        input := &gamelift.PutScalingPolicyInput{}
        output := &gamelift.PutScalingPolicyOutput{}

        mockClient.On("PutScalingPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.PutScalingPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRegisterCompute", func(t *testing.T) {
        input := &gamelift.RegisterComputeInput{}
        output := &gamelift.RegisterComputeOutput{}

        mockClient.On("RegisterCompute", ctx, input).Return(output, nil)

        result, err := mockClient.RegisterCompute(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRegisterGameServer", func(t *testing.T) {
        input := &gamelift.RegisterGameServerInput{}
        output := &gamelift.RegisterGameServerOutput{}

        mockClient.On("RegisterGameServer", ctx, input).Return(output, nil)

        result, err := mockClient.RegisterGameServer(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRequestUploadCredentials", func(t *testing.T) {
        input := &gamelift.RequestUploadCredentialsInput{}
        output := &gamelift.RequestUploadCredentialsOutput{}

        mockClient.On("RequestUploadCredentials", ctx, input).Return(output, nil)

        result, err := mockClient.RequestUploadCredentials(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestResolveAlias", func(t *testing.T) {
        input := &gamelift.ResolveAliasInput{}
        output := &gamelift.ResolveAliasOutput{}

        mockClient.On("ResolveAlias", ctx, input).Return(output, nil)

        result, err := mockClient.ResolveAlias(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestResumeGameServerGroup", func(t *testing.T) {
        input := &gamelift.ResumeGameServerGroupInput{}
        output := &gamelift.ResumeGameServerGroupOutput{}

        mockClient.On("ResumeGameServerGroup", ctx, input).Return(output, nil)

        result, err := mockClient.ResumeGameServerGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSearchGameSessions", func(t *testing.T) {
        input := &gamelift.SearchGameSessionsInput{}
        output := &gamelift.SearchGameSessionsOutput{}

        mockClient.On("SearchGameSessions", ctx, input).Return(output, nil)

        result, err := mockClient.SearchGameSessions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartFleetActions", func(t *testing.T) {
        input := &gamelift.StartFleetActionsInput{}
        output := &gamelift.StartFleetActionsOutput{}

        mockClient.On("StartFleetActions", ctx, input).Return(output, nil)

        result, err := mockClient.StartFleetActions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartGameSessionPlacement", func(t *testing.T) {
        input := &gamelift.StartGameSessionPlacementInput{}
        output := &gamelift.StartGameSessionPlacementOutput{}

        mockClient.On("StartGameSessionPlacement", ctx, input).Return(output, nil)

        result, err := mockClient.StartGameSessionPlacement(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartMatchBackfill", func(t *testing.T) {
        input := &gamelift.StartMatchBackfillInput{}
        output := &gamelift.StartMatchBackfillOutput{}

        mockClient.On("StartMatchBackfill", ctx, input).Return(output, nil)

        result, err := mockClient.StartMatchBackfill(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartMatchmaking", func(t *testing.T) {
        input := &gamelift.StartMatchmakingInput{}
        output := &gamelift.StartMatchmakingOutput{}

        mockClient.On("StartMatchmaking", ctx, input).Return(output, nil)

        result, err := mockClient.StartMatchmaking(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopFleetActions", func(t *testing.T) {
        input := &gamelift.StopFleetActionsInput{}
        output := &gamelift.StopFleetActionsOutput{}

        mockClient.On("StopFleetActions", ctx, input).Return(output, nil)

        result, err := mockClient.StopFleetActions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopGameSessionPlacement", func(t *testing.T) {
        input := &gamelift.StopGameSessionPlacementInput{}
        output := &gamelift.StopGameSessionPlacementOutput{}

        mockClient.On("StopGameSessionPlacement", ctx, input).Return(output, nil)

        result, err := mockClient.StopGameSessionPlacement(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopMatchmaking", func(t *testing.T) {
        input := &gamelift.StopMatchmakingInput{}
        output := &gamelift.StopMatchmakingOutput{}

        mockClient.On("StopMatchmaking", ctx, input).Return(output, nil)

        result, err := mockClient.StopMatchmaking(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSuspendGameServerGroup", func(t *testing.T) {
        input := &gamelift.SuspendGameServerGroupInput{}
        output := &gamelift.SuspendGameServerGroupOutput{}

        mockClient.On("SuspendGameServerGroup", ctx, input).Return(output, nil)

        result, err := mockClient.SuspendGameServerGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &gamelift.TagResourceInput{}
        output := &gamelift.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &gamelift.UntagResourceInput{}
        output := &gamelift.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateAlias", func(t *testing.T) {
        input := &gamelift.UpdateAliasInput{}
        output := &gamelift.UpdateAliasOutput{}

        mockClient.On("UpdateAlias", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateAlias(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateBuild", func(t *testing.T) {
        input := &gamelift.UpdateBuildInput{}
        output := &gamelift.UpdateBuildOutput{}

        mockClient.On("UpdateBuild", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateBuild(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateFleetAttributes", func(t *testing.T) {
        input := &gamelift.UpdateFleetAttributesInput{}
        output := &gamelift.UpdateFleetAttributesOutput{}

        mockClient.On("UpdateFleetAttributes", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateFleetAttributes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateFleetCapacity", func(t *testing.T) {
        input := &gamelift.UpdateFleetCapacityInput{}
        output := &gamelift.UpdateFleetCapacityOutput{}

        mockClient.On("UpdateFleetCapacity", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateFleetCapacity(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateFleetPortSettings", func(t *testing.T) {
        input := &gamelift.UpdateFleetPortSettingsInput{}
        output := &gamelift.UpdateFleetPortSettingsOutput{}

        mockClient.On("UpdateFleetPortSettings", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateFleetPortSettings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateGameServer", func(t *testing.T) {
        input := &gamelift.UpdateGameServerInput{}
        output := &gamelift.UpdateGameServerOutput{}

        mockClient.On("UpdateGameServer", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateGameServer(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateGameServerGroup", func(t *testing.T) {
        input := &gamelift.UpdateGameServerGroupInput{}
        output := &gamelift.UpdateGameServerGroupOutput{}

        mockClient.On("UpdateGameServerGroup", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateGameServerGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateGameSession", func(t *testing.T) {
        input := &gamelift.UpdateGameSessionInput{}
        output := &gamelift.UpdateGameSessionOutput{}

        mockClient.On("UpdateGameSession", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateGameSession(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateGameSessionQueue", func(t *testing.T) {
        input := &gamelift.UpdateGameSessionQueueInput{}
        output := &gamelift.UpdateGameSessionQueueOutput{}

        mockClient.On("UpdateGameSessionQueue", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateGameSessionQueue(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateMatchmakingConfiguration", func(t *testing.T) {
        input := &gamelift.UpdateMatchmakingConfigurationInput{}
        output := &gamelift.UpdateMatchmakingConfigurationOutput{}

        mockClient.On("UpdateMatchmakingConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateMatchmakingConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateRuntimeConfiguration", func(t *testing.T) {
        input := &gamelift.UpdateRuntimeConfigurationInput{}
        output := &gamelift.UpdateRuntimeConfigurationOutput{}

        mockClient.On("UpdateRuntimeConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateRuntimeConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateScript", func(t *testing.T) {
        input := &gamelift.UpdateScriptInput{}
        output := &gamelift.UpdateScriptOutput{}

        mockClient.On("UpdateScript", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateScript(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestValidateMatchmakingRuleSet", func(t *testing.T) {
        input := &gamelift.ValidateMatchmakingRuleSetInput{}
        output := &gamelift.ValidateMatchmakingRuleSetOutput{}

        mockClient.On("ValidateMatchmakingRuleSet", ctx, input).Return(output, nil)

        result, err := mockClient.ValidateMatchmakingRuleSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
