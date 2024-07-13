package groundstation_test

// tests for the groundstation service interface for this ../../../service/groundstation/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/groundstation"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/groundstation/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/groundstation/groundstation_iface"
	"github.com/stretchr/testify/assert"
)

func TestGroundstationServiceCanBeMocked(t *testing.T) {
	var iface groundstation_iface.IClient
	iface = &groundstation.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := groundstation.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCancelContact", func(t *testing.T) {
        input := &groundstation.CancelContactInput{}
        output := &groundstation.CancelContactOutput{}

        mockClient.On("CancelContact", ctx, input).Return(output, nil)

        result, err := mockClient.CancelContact(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateConfig", func(t *testing.T) {
        input := &groundstation.CreateConfigInput{}
        output := &groundstation.CreateConfigOutput{}

        mockClient.On("CreateConfig", ctx, input).Return(output, nil)

        result, err := mockClient.CreateConfig(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateDataflowEndpointGroup", func(t *testing.T) {
        input := &groundstation.CreateDataflowEndpointGroupInput{}
        output := &groundstation.CreateDataflowEndpointGroupOutput{}

        mockClient.On("CreateDataflowEndpointGroup", ctx, input).Return(output, nil)

        result, err := mockClient.CreateDataflowEndpointGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateEphemeris", func(t *testing.T) {
        input := &groundstation.CreateEphemerisInput{}
        output := &groundstation.CreateEphemerisOutput{}

        mockClient.On("CreateEphemeris", ctx, input).Return(output, nil)

        result, err := mockClient.CreateEphemeris(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateMissionProfile", func(t *testing.T) {
        input := &groundstation.CreateMissionProfileInput{}
        output := &groundstation.CreateMissionProfileOutput{}

        mockClient.On("CreateMissionProfile", ctx, input).Return(output, nil)

        result, err := mockClient.CreateMissionProfile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteConfig", func(t *testing.T) {
        input := &groundstation.DeleteConfigInput{}
        output := &groundstation.DeleteConfigOutput{}

        mockClient.On("DeleteConfig", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteConfig(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteDataflowEndpointGroup", func(t *testing.T) {
        input := &groundstation.DeleteDataflowEndpointGroupInput{}
        output := &groundstation.DeleteDataflowEndpointGroupOutput{}

        mockClient.On("DeleteDataflowEndpointGroup", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteDataflowEndpointGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteEphemeris", func(t *testing.T) {
        input := &groundstation.DeleteEphemerisInput{}
        output := &groundstation.DeleteEphemerisOutput{}

        mockClient.On("DeleteEphemeris", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteEphemeris(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteMissionProfile", func(t *testing.T) {
        input := &groundstation.DeleteMissionProfileInput{}
        output := &groundstation.DeleteMissionProfileOutput{}

        mockClient.On("DeleteMissionProfile", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteMissionProfile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeContact", func(t *testing.T) {
        input := &groundstation.DescribeContactInput{}
        output := &groundstation.DescribeContactOutput{}

        mockClient.On("DescribeContact", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeContact(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeEphemeris", func(t *testing.T) {
        input := &groundstation.DescribeEphemerisInput{}
        output := &groundstation.DescribeEphemerisOutput{}

        mockClient.On("DescribeEphemeris", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeEphemeris(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetAgentConfiguration", func(t *testing.T) {
        input := &groundstation.GetAgentConfigurationInput{}
        output := &groundstation.GetAgentConfigurationOutput{}

        mockClient.On("GetAgentConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.GetAgentConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetConfig", func(t *testing.T) {
        input := &groundstation.GetConfigInput{}
        output := &groundstation.GetConfigOutput{}

        mockClient.On("GetConfig", ctx, input).Return(output, nil)

        result, err := mockClient.GetConfig(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDataflowEndpointGroup", func(t *testing.T) {
        input := &groundstation.GetDataflowEndpointGroupInput{}
        output := &groundstation.GetDataflowEndpointGroupOutput{}

        mockClient.On("GetDataflowEndpointGroup", ctx, input).Return(output, nil)

        result, err := mockClient.GetDataflowEndpointGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetMinuteUsage", func(t *testing.T) {
        input := &groundstation.GetMinuteUsageInput{}
        output := &groundstation.GetMinuteUsageOutput{}

        mockClient.On("GetMinuteUsage", ctx, input).Return(output, nil)

        result, err := mockClient.GetMinuteUsage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetMissionProfile", func(t *testing.T) {
        input := &groundstation.GetMissionProfileInput{}
        output := &groundstation.GetMissionProfileOutput{}

        mockClient.On("GetMissionProfile", ctx, input).Return(output, nil)

        result, err := mockClient.GetMissionProfile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSatellite", func(t *testing.T) {
        input := &groundstation.GetSatelliteInput{}
        output := &groundstation.GetSatelliteOutput{}

        mockClient.On("GetSatellite", ctx, input).Return(output, nil)

        result, err := mockClient.GetSatellite(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListConfigs", func(t *testing.T) {
        input := &groundstation.ListConfigsInput{}
        output := &groundstation.ListConfigsOutput{}

        mockClient.On("ListConfigs", ctx, input).Return(output, nil)

        result, err := mockClient.ListConfigs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListContacts", func(t *testing.T) {
        input := &groundstation.ListContactsInput{}
        output := &groundstation.ListContactsOutput{}

        mockClient.On("ListContacts", ctx, input).Return(output, nil)

        result, err := mockClient.ListContacts(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDataflowEndpointGroups", func(t *testing.T) {
        input := &groundstation.ListDataflowEndpointGroupsInput{}
        output := &groundstation.ListDataflowEndpointGroupsOutput{}

        mockClient.On("ListDataflowEndpointGroups", ctx, input).Return(output, nil)

        result, err := mockClient.ListDataflowEndpointGroups(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListEphemerides", func(t *testing.T) {
        input := &groundstation.ListEphemeridesInput{}
        output := &groundstation.ListEphemeridesOutput{}

        mockClient.On("ListEphemerides", ctx, input).Return(output, nil)

        result, err := mockClient.ListEphemerides(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListGroundStations", func(t *testing.T) {
        input := &groundstation.ListGroundStationsInput{}
        output := &groundstation.ListGroundStationsOutput{}

        mockClient.On("ListGroundStations", ctx, input).Return(output, nil)

        result, err := mockClient.ListGroundStations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListMissionProfiles", func(t *testing.T) {
        input := &groundstation.ListMissionProfilesInput{}
        output := &groundstation.ListMissionProfilesOutput{}

        mockClient.On("ListMissionProfiles", ctx, input).Return(output, nil)

        result, err := mockClient.ListMissionProfiles(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListSatellites", func(t *testing.T) {
        input := &groundstation.ListSatellitesInput{}
        output := &groundstation.ListSatellitesOutput{}

        mockClient.On("ListSatellites", ctx, input).Return(output, nil)

        result, err := mockClient.ListSatellites(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &groundstation.ListTagsForResourceInput{}
        output := &groundstation.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRegisterAgent", func(t *testing.T) {
        input := &groundstation.RegisterAgentInput{}
        output := &groundstation.RegisterAgentOutput{}

        mockClient.On("RegisterAgent", ctx, input).Return(output, nil)

        result, err := mockClient.RegisterAgent(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestReserveContact", func(t *testing.T) {
        input := &groundstation.ReserveContactInput{}
        output := &groundstation.ReserveContactOutput{}

        mockClient.On("ReserveContact", ctx, input).Return(output, nil)

        result, err := mockClient.ReserveContact(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &groundstation.TagResourceInput{}
        output := &groundstation.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &groundstation.UntagResourceInput{}
        output := &groundstation.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateAgentStatus", func(t *testing.T) {
        input := &groundstation.UpdateAgentStatusInput{}
        output := &groundstation.UpdateAgentStatusOutput{}

        mockClient.On("UpdateAgentStatus", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateAgentStatus(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateConfig", func(t *testing.T) {
        input := &groundstation.UpdateConfigInput{}
        output := &groundstation.UpdateConfigOutput{}

        mockClient.On("UpdateConfig", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateConfig(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateEphemeris", func(t *testing.T) {
        input := &groundstation.UpdateEphemerisInput{}
        output := &groundstation.UpdateEphemerisOutput{}

        mockClient.On("UpdateEphemeris", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateEphemeris(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateMissionProfile", func(t *testing.T) {
        input := &groundstation.UpdateMissionProfileInput{}
        output := &groundstation.UpdateMissionProfileOutput{}

        mockClient.On("UpdateMissionProfile", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateMissionProfile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
