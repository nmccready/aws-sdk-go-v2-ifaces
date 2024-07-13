package mediatailor_test

// tests for the mediatailor service interface for this ../../../service/mediatailor/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/mediatailor"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/mediatailor/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/mediatailor/mediatailor_iface"
	"github.com/stretchr/testify/assert"
)

func TestMediatailorServiceCanBeMocked(t *testing.T) {
	var iface mediatailor_iface.IClient
	iface = &mediatailor.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := mediatailor.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestConfigureLogsForChannel", func(t *testing.T) {
        input := &mediatailor.ConfigureLogsForChannelInput{}
        output := &mediatailor.ConfigureLogsForChannelOutput{}

        mockClient.On("ConfigureLogsForChannel", ctx, input).Return(output, nil)

        result, err := mockClient.ConfigureLogsForChannel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestConfigureLogsForPlaybackConfiguration", func(t *testing.T) {
        input := &mediatailor.ConfigureLogsForPlaybackConfigurationInput{}
        output := &mediatailor.ConfigureLogsForPlaybackConfigurationOutput{}

        mockClient.On("ConfigureLogsForPlaybackConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.ConfigureLogsForPlaybackConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateChannel", func(t *testing.T) {
        input := &mediatailor.CreateChannelInput{}
        output := &mediatailor.CreateChannelOutput{}

        mockClient.On("CreateChannel", ctx, input).Return(output, nil)

        result, err := mockClient.CreateChannel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateLiveSource", func(t *testing.T) {
        input := &mediatailor.CreateLiveSourceInput{}
        output := &mediatailor.CreateLiveSourceOutput{}

        mockClient.On("CreateLiveSource", ctx, input).Return(output, nil)

        result, err := mockClient.CreateLiveSource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreatePrefetchSchedule", func(t *testing.T) {
        input := &mediatailor.CreatePrefetchScheduleInput{}
        output := &mediatailor.CreatePrefetchScheduleOutput{}

        mockClient.On("CreatePrefetchSchedule", ctx, input).Return(output, nil)

        result, err := mockClient.CreatePrefetchSchedule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateProgram", func(t *testing.T) {
        input := &mediatailor.CreateProgramInput{}
        output := &mediatailor.CreateProgramOutput{}

        mockClient.On("CreateProgram", ctx, input).Return(output, nil)

        result, err := mockClient.CreateProgram(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateSourceLocation", func(t *testing.T) {
        input := &mediatailor.CreateSourceLocationInput{}
        output := &mediatailor.CreateSourceLocationOutput{}

        mockClient.On("CreateSourceLocation", ctx, input).Return(output, nil)

        result, err := mockClient.CreateSourceLocation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateVodSource", func(t *testing.T) {
        input := &mediatailor.CreateVodSourceInput{}
        output := &mediatailor.CreateVodSourceOutput{}

        mockClient.On("CreateVodSource", ctx, input).Return(output, nil)

        result, err := mockClient.CreateVodSource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteChannel", func(t *testing.T) {
        input := &mediatailor.DeleteChannelInput{}
        output := &mediatailor.DeleteChannelOutput{}

        mockClient.On("DeleteChannel", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteChannel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteChannelPolicy", func(t *testing.T) {
        input := &mediatailor.DeleteChannelPolicyInput{}
        output := &mediatailor.DeleteChannelPolicyOutput{}

        mockClient.On("DeleteChannelPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteChannelPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteLiveSource", func(t *testing.T) {
        input := &mediatailor.DeleteLiveSourceInput{}
        output := &mediatailor.DeleteLiveSourceOutput{}

        mockClient.On("DeleteLiveSource", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteLiveSource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeletePlaybackConfiguration", func(t *testing.T) {
        input := &mediatailor.DeletePlaybackConfigurationInput{}
        output := &mediatailor.DeletePlaybackConfigurationOutput{}

        mockClient.On("DeletePlaybackConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.DeletePlaybackConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeletePrefetchSchedule", func(t *testing.T) {
        input := &mediatailor.DeletePrefetchScheduleInput{}
        output := &mediatailor.DeletePrefetchScheduleOutput{}

        mockClient.On("DeletePrefetchSchedule", ctx, input).Return(output, nil)

        result, err := mockClient.DeletePrefetchSchedule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteProgram", func(t *testing.T) {
        input := &mediatailor.DeleteProgramInput{}
        output := &mediatailor.DeleteProgramOutput{}

        mockClient.On("DeleteProgram", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteProgram(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteSourceLocation", func(t *testing.T) {
        input := &mediatailor.DeleteSourceLocationInput{}
        output := &mediatailor.DeleteSourceLocationOutput{}

        mockClient.On("DeleteSourceLocation", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteSourceLocation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteVodSource", func(t *testing.T) {
        input := &mediatailor.DeleteVodSourceInput{}
        output := &mediatailor.DeleteVodSourceOutput{}

        mockClient.On("DeleteVodSource", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteVodSource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeChannel", func(t *testing.T) {
        input := &mediatailor.DescribeChannelInput{}
        output := &mediatailor.DescribeChannelOutput{}

        mockClient.On("DescribeChannel", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeChannel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeLiveSource", func(t *testing.T) {
        input := &mediatailor.DescribeLiveSourceInput{}
        output := &mediatailor.DescribeLiveSourceOutput{}

        mockClient.On("DescribeLiveSource", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeLiveSource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeProgram", func(t *testing.T) {
        input := &mediatailor.DescribeProgramInput{}
        output := &mediatailor.DescribeProgramOutput{}

        mockClient.On("DescribeProgram", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeProgram(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeSourceLocation", func(t *testing.T) {
        input := &mediatailor.DescribeSourceLocationInput{}
        output := &mediatailor.DescribeSourceLocationOutput{}

        mockClient.On("DescribeSourceLocation", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeSourceLocation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeVodSource", func(t *testing.T) {
        input := &mediatailor.DescribeVodSourceInput{}
        output := &mediatailor.DescribeVodSourceOutput{}

        mockClient.On("DescribeVodSource", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeVodSource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetChannelPolicy", func(t *testing.T) {
        input := &mediatailor.GetChannelPolicyInput{}
        output := &mediatailor.GetChannelPolicyOutput{}

        mockClient.On("GetChannelPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.GetChannelPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetChannelSchedule", func(t *testing.T) {
        input := &mediatailor.GetChannelScheduleInput{}
        output := &mediatailor.GetChannelScheduleOutput{}

        mockClient.On("GetChannelSchedule", ctx, input).Return(output, nil)

        result, err := mockClient.GetChannelSchedule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetPlaybackConfiguration", func(t *testing.T) {
        input := &mediatailor.GetPlaybackConfigurationInput{}
        output := &mediatailor.GetPlaybackConfigurationOutput{}

        mockClient.On("GetPlaybackConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.GetPlaybackConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetPrefetchSchedule", func(t *testing.T) {
        input := &mediatailor.GetPrefetchScheduleInput{}
        output := &mediatailor.GetPrefetchScheduleOutput{}

        mockClient.On("GetPrefetchSchedule", ctx, input).Return(output, nil)

        result, err := mockClient.GetPrefetchSchedule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAlerts", func(t *testing.T) {
        input := &mediatailor.ListAlertsInput{}
        output := &mediatailor.ListAlertsOutput{}

        mockClient.On("ListAlerts", ctx, input).Return(output, nil)

        result, err := mockClient.ListAlerts(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListChannels", func(t *testing.T) {
        input := &mediatailor.ListChannelsInput{}
        output := &mediatailor.ListChannelsOutput{}

        mockClient.On("ListChannels", ctx, input).Return(output, nil)

        result, err := mockClient.ListChannels(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListLiveSources", func(t *testing.T) {
        input := &mediatailor.ListLiveSourcesInput{}
        output := &mediatailor.ListLiveSourcesOutput{}

        mockClient.On("ListLiveSources", ctx, input).Return(output, nil)

        result, err := mockClient.ListLiveSources(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListPlaybackConfigurations", func(t *testing.T) {
        input := &mediatailor.ListPlaybackConfigurationsInput{}
        output := &mediatailor.ListPlaybackConfigurationsOutput{}

        mockClient.On("ListPlaybackConfigurations", ctx, input).Return(output, nil)

        result, err := mockClient.ListPlaybackConfigurations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListPrefetchSchedules", func(t *testing.T) {
        input := &mediatailor.ListPrefetchSchedulesInput{}
        output := &mediatailor.ListPrefetchSchedulesOutput{}

        mockClient.On("ListPrefetchSchedules", ctx, input).Return(output, nil)

        result, err := mockClient.ListPrefetchSchedules(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListSourceLocations", func(t *testing.T) {
        input := &mediatailor.ListSourceLocationsInput{}
        output := &mediatailor.ListSourceLocationsOutput{}

        mockClient.On("ListSourceLocations", ctx, input).Return(output, nil)

        result, err := mockClient.ListSourceLocations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &mediatailor.ListTagsForResourceInput{}
        output := &mediatailor.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListVodSources", func(t *testing.T) {
        input := &mediatailor.ListVodSourcesInput{}
        output := &mediatailor.ListVodSourcesOutput{}

        mockClient.On("ListVodSources", ctx, input).Return(output, nil)

        result, err := mockClient.ListVodSources(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutChannelPolicy", func(t *testing.T) {
        input := &mediatailor.PutChannelPolicyInput{}
        output := &mediatailor.PutChannelPolicyOutput{}

        mockClient.On("PutChannelPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.PutChannelPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutPlaybackConfiguration", func(t *testing.T) {
        input := &mediatailor.PutPlaybackConfigurationInput{}
        output := &mediatailor.PutPlaybackConfigurationOutput{}

        mockClient.On("PutPlaybackConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.PutPlaybackConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartChannel", func(t *testing.T) {
        input := &mediatailor.StartChannelInput{}
        output := &mediatailor.StartChannelOutput{}

        mockClient.On("StartChannel", ctx, input).Return(output, nil)

        result, err := mockClient.StartChannel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopChannel", func(t *testing.T) {
        input := &mediatailor.StopChannelInput{}
        output := &mediatailor.StopChannelOutput{}

        mockClient.On("StopChannel", ctx, input).Return(output, nil)

        result, err := mockClient.StopChannel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &mediatailor.TagResourceInput{}
        output := &mediatailor.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &mediatailor.UntagResourceInput{}
        output := &mediatailor.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateChannel", func(t *testing.T) {
        input := &mediatailor.UpdateChannelInput{}
        output := &mediatailor.UpdateChannelOutput{}

        mockClient.On("UpdateChannel", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateChannel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateLiveSource", func(t *testing.T) {
        input := &mediatailor.UpdateLiveSourceInput{}
        output := &mediatailor.UpdateLiveSourceOutput{}

        mockClient.On("UpdateLiveSource", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateLiveSource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateProgram", func(t *testing.T) {
        input := &mediatailor.UpdateProgramInput{}
        output := &mediatailor.UpdateProgramOutput{}

        mockClient.On("UpdateProgram", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateProgram(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateSourceLocation", func(t *testing.T) {
        input := &mediatailor.UpdateSourceLocationInput{}
        output := &mediatailor.UpdateSourceLocationOutput{}

        mockClient.On("UpdateSourceLocation", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateSourceLocation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateVodSource", func(t *testing.T) {
        input := &mediatailor.UpdateVodSourceInput{}
        output := &mediatailor.UpdateVodSourceOutput{}

        mockClient.On("UpdateVodSource", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateVodSource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
