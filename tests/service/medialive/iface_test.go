package medialive_test

// tests for the medialive service interface for this ../../../service/medialive/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/medialive"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/medialive/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/medialive/medialive_iface"
	"github.com/stretchr/testify/assert"
)

func TestMedialiveServiceCanBeMocked(t *testing.T) {
	var iface medialive_iface.IClient
	iface = &medialive.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := medialive.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAcceptInputDeviceTransfer", func(t *testing.T) {
        input := &medialive.AcceptInputDeviceTransferInput{}
        output := &medialive.AcceptInputDeviceTransferOutput{}

        mockClient.On("AcceptInputDeviceTransfer", ctx, input).Return(output, nil)

        result, err := mockClient.AcceptInputDeviceTransfer(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchDelete", func(t *testing.T) {
        input := &medialive.BatchDeleteInput{}
        output := &medialive.BatchDeleteOutput{}

        mockClient.On("BatchDelete", ctx, input).Return(output, nil)

        result, err := mockClient.BatchDelete(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchStart", func(t *testing.T) {
        input := &medialive.BatchStartInput{}
        output := &medialive.BatchStartOutput{}

        mockClient.On("BatchStart", ctx, input).Return(output, nil)

        result, err := mockClient.BatchStart(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchStop", func(t *testing.T) {
        input := &medialive.BatchStopInput{}
        output := &medialive.BatchStopOutput{}

        mockClient.On("BatchStop", ctx, input).Return(output, nil)

        result, err := mockClient.BatchStop(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchUpdateSchedule", func(t *testing.T) {
        input := &medialive.BatchUpdateScheduleInput{}
        output := &medialive.BatchUpdateScheduleOutput{}

        mockClient.On("BatchUpdateSchedule", ctx, input).Return(output, nil)

        result, err := mockClient.BatchUpdateSchedule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCancelInputDeviceTransfer", func(t *testing.T) {
        input := &medialive.CancelInputDeviceTransferInput{}
        output := &medialive.CancelInputDeviceTransferOutput{}

        mockClient.On("CancelInputDeviceTransfer", ctx, input).Return(output, nil)

        result, err := mockClient.CancelInputDeviceTransfer(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestClaimDevice", func(t *testing.T) {
        input := &medialive.ClaimDeviceInput{}
        output := &medialive.ClaimDeviceOutput{}

        mockClient.On("ClaimDevice", ctx, input).Return(output, nil)

        result, err := mockClient.ClaimDevice(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateChannel", func(t *testing.T) {
        input := &medialive.CreateChannelInput{}
        output := &medialive.CreateChannelOutput{}

        mockClient.On("CreateChannel", ctx, input).Return(output, nil)

        result, err := mockClient.CreateChannel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateCloudWatchAlarmTemplate", func(t *testing.T) {
        input := &medialive.CreateCloudWatchAlarmTemplateInput{}
        output := &medialive.CreateCloudWatchAlarmTemplateOutput{}

        mockClient.On("CreateCloudWatchAlarmTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.CreateCloudWatchAlarmTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateCloudWatchAlarmTemplateGroup", func(t *testing.T) {
        input := &medialive.CreateCloudWatchAlarmTemplateGroupInput{}
        output := &medialive.CreateCloudWatchAlarmTemplateGroupOutput{}

        mockClient.On("CreateCloudWatchAlarmTemplateGroup", ctx, input).Return(output, nil)

        result, err := mockClient.CreateCloudWatchAlarmTemplateGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateEventBridgeRuleTemplate", func(t *testing.T) {
        input := &medialive.CreateEventBridgeRuleTemplateInput{}
        output := &medialive.CreateEventBridgeRuleTemplateOutput{}

        mockClient.On("CreateEventBridgeRuleTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.CreateEventBridgeRuleTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateEventBridgeRuleTemplateGroup", func(t *testing.T) {
        input := &medialive.CreateEventBridgeRuleTemplateGroupInput{}
        output := &medialive.CreateEventBridgeRuleTemplateGroupOutput{}

        mockClient.On("CreateEventBridgeRuleTemplateGroup", ctx, input).Return(output, nil)

        result, err := mockClient.CreateEventBridgeRuleTemplateGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateInput", func(t *testing.T) {
        input := &medialive.CreateInputInput{}
        output := &medialive.CreateInputOutput{}

        mockClient.On("CreateInput", ctx, input).Return(output, nil)

        result, err := mockClient.CreateInput(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateInputSecurityGroup", func(t *testing.T) {
        input := &medialive.CreateInputSecurityGroupInput{}
        output := &medialive.CreateInputSecurityGroupOutput{}

        mockClient.On("CreateInputSecurityGroup", ctx, input).Return(output, nil)

        result, err := mockClient.CreateInputSecurityGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateMultiplex", func(t *testing.T) {
        input := &medialive.CreateMultiplexInput{}
        output := &medialive.CreateMultiplexOutput{}

        mockClient.On("CreateMultiplex", ctx, input).Return(output, nil)

        result, err := mockClient.CreateMultiplex(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateMultiplexProgram", func(t *testing.T) {
        input := &medialive.CreateMultiplexProgramInput{}
        output := &medialive.CreateMultiplexProgramOutput{}

        mockClient.On("CreateMultiplexProgram", ctx, input).Return(output, nil)

        result, err := mockClient.CreateMultiplexProgram(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreatePartnerInput", func(t *testing.T) {
        input := &medialive.CreatePartnerInputInput{}
        output := &medialive.CreatePartnerInputOutput{}

        mockClient.On("CreatePartnerInput", ctx, input).Return(output, nil)

        result, err := mockClient.CreatePartnerInput(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateSignalMap", func(t *testing.T) {
        input := &medialive.CreateSignalMapInput{}
        output := &medialive.CreateSignalMapOutput{}

        mockClient.On("CreateSignalMap", ctx, input).Return(output, nil)

        result, err := mockClient.CreateSignalMap(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateTags", func(t *testing.T) {
        input := &medialive.CreateTagsInput{}
        output := &medialive.CreateTagsOutput{}

        mockClient.On("CreateTags", ctx, input).Return(output, nil)

        result, err := mockClient.CreateTags(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteChannel", func(t *testing.T) {
        input := &medialive.DeleteChannelInput{}
        output := &medialive.DeleteChannelOutput{}

        mockClient.On("DeleteChannel", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteChannel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteCloudWatchAlarmTemplate", func(t *testing.T) {
        input := &medialive.DeleteCloudWatchAlarmTemplateInput{}
        output := &medialive.DeleteCloudWatchAlarmTemplateOutput{}

        mockClient.On("DeleteCloudWatchAlarmTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteCloudWatchAlarmTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteCloudWatchAlarmTemplateGroup", func(t *testing.T) {
        input := &medialive.DeleteCloudWatchAlarmTemplateGroupInput{}
        output := &medialive.DeleteCloudWatchAlarmTemplateGroupOutput{}

        mockClient.On("DeleteCloudWatchAlarmTemplateGroup", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteCloudWatchAlarmTemplateGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteEventBridgeRuleTemplate", func(t *testing.T) {
        input := &medialive.DeleteEventBridgeRuleTemplateInput{}
        output := &medialive.DeleteEventBridgeRuleTemplateOutput{}

        mockClient.On("DeleteEventBridgeRuleTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteEventBridgeRuleTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteEventBridgeRuleTemplateGroup", func(t *testing.T) {
        input := &medialive.DeleteEventBridgeRuleTemplateGroupInput{}
        output := &medialive.DeleteEventBridgeRuleTemplateGroupOutput{}

        mockClient.On("DeleteEventBridgeRuleTemplateGroup", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteEventBridgeRuleTemplateGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteInput", func(t *testing.T) {
        input := &medialive.DeleteInputInput{}
        output := &medialive.DeleteInputOutput{}

        mockClient.On("DeleteInput", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteInput(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteInputSecurityGroup", func(t *testing.T) {
        input := &medialive.DeleteInputSecurityGroupInput{}
        output := &medialive.DeleteInputSecurityGroupOutput{}

        mockClient.On("DeleteInputSecurityGroup", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteInputSecurityGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteMultiplex", func(t *testing.T) {
        input := &medialive.DeleteMultiplexInput{}
        output := &medialive.DeleteMultiplexOutput{}

        mockClient.On("DeleteMultiplex", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteMultiplex(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteMultiplexProgram", func(t *testing.T) {
        input := &medialive.DeleteMultiplexProgramInput{}
        output := &medialive.DeleteMultiplexProgramOutput{}

        mockClient.On("DeleteMultiplexProgram", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteMultiplexProgram(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteReservation", func(t *testing.T) {
        input := &medialive.DeleteReservationInput{}
        output := &medialive.DeleteReservationOutput{}

        mockClient.On("DeleteReservation", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteReservation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteSchedule", func(t *testing.T) {
        input := &medialive.DeleteScheduleInput{}
        output := &medialive.DeleteScheduleOutput{}

        mockClient.On("DeleteSchedule", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteSchedule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteSignalMap", func(t *testing.T) {
        input := &medialive.DeleteSignalMapInput{}
        output := &medialive.DeleteSignalMapOutput{}

        mockClient.On("DeleteSignalMap", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteSignalMap(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteTags", func(t *testing.T) {
        input := &medialive.DeleteTagsInput{}
        output := &medialive.DeleteTagsOutput{}

        mockClient.On("DeleteTags", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteTags(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeAccountConfiguration", func(t *testing.T) {
        input := &medialive.DescribeAccountConfigurationInput{}
        output := &medialive.DescribeAccountConfigurationOutput{}

        mockClient.On("DescribeAccountConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeAccountConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeChannel", func(t *testing.T) {
        input := &medialive.DescribeChannelInput{}
        output := &medialive.DescribeChannelOutput{}

        mockClient.On("DescribeChannel", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeChannel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeInput", func(t *testing.T) {
        input := &medialive.DescribeInputInput{}
        output := &medialive.DescribeInputOutput{}

        mockClient.On("DescribeInput", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeInput(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeInputDevice", func(t *testing.T) {
        input := &medialive.DescribeInputDeviceInput{}
        output := &medialive.DescribeInputDeviceOutput{}

        mockClient.On("DescribeInputDevice", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeInputDevice(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeInputDeviceThumbnail", func(t *testing.T) {
        input := &medialive.DescribeInputDeviceThumbnailInput{}
        output := &medialive.DescribeInputDeviceThumbnailOutput{}

        mockClient.On("DescribeInputDeviceThumbnail", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeInputDeviceThumbnail(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeInputSecurityGroup", func(t *testing.T) {
        input := &medialive.DescribeInputSecurityGroupInput{}
        output := &medialive.DescribeInputSecurityGroupOutput{}

        mockClient.On("DescribeInputSecurityGroup", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeInputSecurityGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeMultiplex", func(t *testing.T) {
        input := &medialive.DescribeMultiplexInput{}
        output := &medialive.DescribeMultiplexOutput{}

        mockClient.On("DescribeMultiplex", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeMultiplex(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeMultiplexProgram", func(t *testing.T) {
        input := &medialive.DescribeMultiplexProgramInput{}
        output := &medialive.DescribeMultiplexProgramOutput{}

        mockClient.On("DescribeMultiplexProgram", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeMultiplexProgram(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeOffering", func(t *testing.T) {
        input := &medialive.DescribeOfferingInput{}
        output := &medialive.DescribeOfferingOutput{}

        mockClient.On("DescribeOffering", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeOffering(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeReservation", func(t *testing.T) {
        input := &medialive.DescribeReservationInput{}
        output := &medialive.DescribeReservationOutput{}

        mockClient.On("DescribeReservation", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeReservation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeSchedule", func(t *testing.T) {
        input := &medialive.DescribeScheduleInput{}
        output := &medialive.DescribeScheduleOutput{}

        mockClient.On("DescribeSchedule", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeSchedule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeThumbnails", func(t *testing.T) {
        input := &medialive.DescribeThumbnailsInput{}
        output := &medialive.DescribeThumbnailsOutput{}

        mockClient.On("DescribeThumbnails", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeThumbnails(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetCloudWatchAlarmTemplate", func(t *testing.T) {
        input := &medialive.GetCloudWatchAlarmTemplateInput{}
        output := &medialive.GetCloudWatchAlarmTemplateOutput{}

        mockClient.On("GetCloudWatchAlarmTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.GetCloudWatchAlarmTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetCloudWatchAlarmTemplateGroup", func(t *testing.T) {
        input := &medialive.GetCloudWatchAlarmTemplateGroupInput{}
        output := &medialive.GetCloudWatchAlarmTemplateGroupOutput{}

        mockClient.On("GetCloudWatchAlarmTemplateGroup", ctx, input).Return(output, nil)

        result, err := mockClient.GetCloudWatchAlarmTemplateGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetEventBridgeRuleTemplate", func(t *testing.T) {
        input := &medialive.GetEventBridgeRuleTemplateInput{}
        output := &medialive.GetEventBridgeRuleTemplateOutput{}

        mockClient.On("GetEventBridgeRuleTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.GetEventBridgeRuleTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetEventBridgeRuleTemplateGroup", func(t *testing.T) {
        input := &medialive.GetEventBridgeRuleTemplateGroupInput{}
        output := &medialive.GetEventBridgeRuleTemplateGroupOutput{}

        mockClient.On("GetEventBridgeRuleTemplateGroup", ctx, input).Return(output, nil)

        result, err := mockClient.GetEventBridgeRuleTemplateGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSignalMap", func(t *testing.T) {
        input := &medialive.GetSignalMapInput{}
        output := &medialive.GetSignalMapOutput{}

        mockClient.On("GetSignalMap", ctx, input).Return(output, nil)

        result, err := mockClient.GetSignalMap(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListChannels", func(t *testing.T) {
        input := &medialive.ListChannelsInput{}
        output := &medialive.ListChannelsOutput{}

        mockClient.On("ListChannels", ctx, input).Return(output, nil)

        result, err := mockClient.ListChannels(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListCloudWatchAlarmTemplateGroups", func(t *testing.T) {
        input := &medialive.ListCloudWatchAlarmTemplateGroupsInput{}
        output := &medialive.ListCloudWatchAlarmTemplateGroupsOutput{}

        mockClient.On("ListCloudWatchAlarmTemplateGroups", ctx, input).Return(output, nil)

        result, err := mockClient.ListCloudWatchAlarmTemplateGroups(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListCloudWatchAlarmTemplates", func(t *testing.T) {
        input := &medialive.ListCloudWatchAlarmTemplatesInput{}
        output := &medialive.ListCloudWatchAlarmTemplatesOutput{}

        mockClient.On("ListCloudWatchAlarmTemplates", ctx, input).Return(output, nil)

        result, err := mockClient.ListCloudWatchAlarmTemplates(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListEventBridgeRuleTemplateGroups", func(t *testing.T) {
        input := &medialive.ListEventBridgeRuleTemplateGroupsInput{}
        output := &medialive.ListEventBridgeRuleTemplateGroupsOutput{}

        mockClient.On("ListEventBridgeRuleTemplateGroups", ctx, input).Return(output, nil)

        result, err := mockClient.ListEventBridgeRuleTemplateGroups(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListEventBridgeRuleTemplates", func(t *testing.T) {
        input := &medialive.ListEventBridgeRuleTemplatesInput{}
        output := &medialive.ListEventBridgeRuleTemplatesOutput{}

        mockClient.On("ListEventBridgeRuleTemplates", ctx, input).Return(output, nil)

        result, err := mockClient.ListEventBridgeRuleTemplates(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListInputDeviceTransfers", func(t *testing.T) {
        input := &medialive.ListInputDeviceTransfersInput{}
        output := &medialive.ListInputDeviceTransfersOutput{}

        mockClient.On("ListInputDeviceTransfers", ctx, input).Return(output, nil)

        result, err := mockClient.ListInputDeviceTransfers(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListInputDevices", func(t *testing.T) {
        input := &medialive.ListInputDevicesInput{}
        output := &medialive.ListInputDevicesOutput{}

        mockClient.On("ListInputDevices", ctx, input).Return(output, nil)

        result, err := mockClient.ListInputDevices(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListInputSecurityGroups", func(t *testing.T) {
        input := &medialive.ListInputSecurityGroupsInput{}
        output := &medialive.ListInputSecurityGroupsOutput{}

        mockClient.On("ListInputSecurityGroups", ctx, input).Return(output, nil)

        result, err := mockClient.ListInputSecurityGroups(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListInputs", func(t *testing.T) {
        input := &medialive.ListInputsInput{}
        output := &medialive.ListInputsOutput{}

        mockClient.On("ListInputs", ctx, input).Return(output, nil)

        result, err := mockClient.ListInputs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListMultiplexPrograms", func(t *testing.T) {
        input := &medialive.ListMultiplexProgramsInput{}
        output := &medialive.ListMultiplexProgramsOutput{}

        mockClient.On("ListMultiplexPrograms", ctx, input).Return(output, nil)

        result, err := mockClient.ListMultiplexPrograms(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListMultiplexes", func(t *testing.T) {
        input := &medialive.ListMultiplexesInput{}
        output := &medialive.ListMultiplexesOutput{}

        mockClient.On("ListMultiplexes", ctx, input).Return(output, nil)

        result, err := mockClient.ListMultiplexes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListOfferings", func(t *testing.T) {
        input := &medialive.ListOfferingsInput{}
        output := &medialive.ListOfferingsOutput{}

        mockClient.On("ListOfferings", ctx, input).Return(output, nil)

        result, err := mockClient.ListOfferings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListReservations", func(t *testing.T) {
        input := &medialive.ListReservationsInput{}
        output := &medialive.ListReservationsOutput{}

        mockClient.On("ListReservations", ctx, input).Return(output, nil)

        result, err := mockClient.ListReservations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListSignalMaps", func(t *testing.T) {
        input := &medialive.ListSignalMapsInput{}
        output := &medialive.ListSignalMapsOutput{}

        mockClient.On("ListSignalMaps", ctx, input).Return(output, nil)

        result, err := mockClient.ListSignalMaps(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &medialive.ListTagsForResourceInput{}
        output := &medialive.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPurchaseOffering", func(t *testing.T) {
        input := &medialive.PurchaseOfferingInput{}
        output := &medialive.PurchaseOfferingOutput{}

        mockClient.On("PurchaseOffering", ctx, input).Return(output, nil)

        result, err := mockClient.PurchaseOffering(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRebootInputDevice", func(t *testing.T) {
        input := &medialive.RebootInputDeviceInput{}
        output := &medialive.RebootInputDeviceOutput{}

        mockClient.On("RebootInputDevice", ctx, input).Return(output, nil)

        result, err := mockClient.RebootInputDevice(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRejectInputDeviceTransfer", func(t *testing.T) {
        input := &medialive.RejectInputDeviceTransferInput{}
        output := &medialive.RejectInputDeviceTransferOutput{}

        mockClient.On("RejectInputDeviceTransfer", ctx, input).Return(output, nil)

        result, err := mockClient.RejectInputDeviceTransfer(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRestartChannelPipelines", func(t *testing.T) {
        input := &medialive.RestartChannelPipelinesInput{}
        output := &medialive.RestartChannelPipelinesOutput{}

        mockClient.On("RestartChannelPipelines", ctx, input).Return(output, nil)

        result, err := mockClient.RestartChannelPipelines(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartChannel", func(t *testing.T) {
        input := &medialive.StartChannelInput{}
        output := &medialive.StartChannelOutput{}

        mockClient.On("StartChannel", ctx, input).Return(output, nil)

        result, err := mockClient.StartChannel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartDeleteMonitorDeployment", func(t *testing.T) {
        input := &medialive.StartDeleteMonitorDeploymentInput{}
        output := &medialive.StartDeleteMonitorDeploymentOutput{}

        mockClient.On("StartDeleteMonitorDeployment", ctx, input).Return(output, nil)

        result, err := mockClient.StartDeleteMonitorDeployment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartInputDevice", func(t *testing.T) {
        input := &medialive.StartInputDeviceInput{}
        output := &medialive.StartInputDeviceOutput{}

        mockClient.On("StartInputDevice", ctx, input).Return(output, nil)

        result, err := mockClient.StartInputDevice(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartInputDeviceMaintenanceWindow", func(t *testing.T) {
        input := &medialive.StartInputDeviceMaintenanceWindowInput{}
        output := &medialive.StartInputDeviceMaintenanceWindowOutput{}

        mockClient.On("StartInputDeviceMaintenanceWindow", ctx, input).Return(output, nil)

        result, err := mockClient.StartInputDeviceMaintenanceWindow(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartMonitorDeployment", func(t *testing.T) {
        input := &medialive.StartMonitorDeploymentInput{}
        output := &medialive.StartMonitorDeploymentOutput{}

        mockClient.On("StartMonitorDeployment", ctx, input).Return(output, nil)

        result, err := mockClient.StartMonitorDeployment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartMultiplex", func(t *testing.T) {
        input := &medialive.StartMultiplexInput{}
        output := &medialive.StartMultiplexOutput{}

        mockClient.On("StartMultiplex", ctx, input).Return(output, nil)

        result, err := mockClient.StartMultiplex(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartUpdateSignalMap", func(t *testing.T) {
        input := &medialive.StartUpdateSignalMapInput{}
        output := &medialive.StartUpdateSignalMapOutput{}

        mockClient.On("StartUpdateSignalMap", ctx, input).Return(output, nil)

        result, err := mockClient.StartUpdateSignalMap(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopChannel", func(t *testing.T) {
        input := &medialive.StopChannelInput{}
        output := &medialive.StopChannelOutput{}

        mockClient.On("StopChannel", ctx, input).Return(output, nil)

        result, err := mockClient.StopChannel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopInputDevice", func(t *testing.T) {
        input := &medialive.StopInputDeviceInput{}
        output := &medialive.StopInputDeviceOutput{}

        mockClient.On("StopInputDevice", ctx, input).Return(output, nil)

        result, err := mockClient.StopInputDevice(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopMultiplex", func(t *testing.T) {
        input := &medialive.StopMultiplexInput{}
        output := &medialive.StopMultiplexOutput{}

        mockClient.On("StopMultiplex", ctx, input).Return(output, nil)

        result, err := mockClient.StopMultiplex(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTransferInputDevice", func(t *testing.T) {
        input := &medialive.TransferInputDeviceInput{}
        output := &medialive.TransferInputDeviceOutput{}

        mockClient.On("TransferInputDevice", ctx, input).Return(output, nil)

        result, err := mockClient.TransferInputDevice(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateAccountConfiguration", func(t *testing.T) {
        input := &medialive.UpdateAccountConfigurationInput{}
        output := &medialive.UpdateAccountConfigurationOutput{}

        mockClient.On("UpdateAccountConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateAccountConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateChannel", func(t *testing.T) {
        input := &medialive.UpdateChannelInput{}
        output := &medialive.UpdateChannelOutput{}

        mockClient.On("UpdateChannel", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateChannel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateChannelClass", func(t *testing.T) {
        input := &medialive.UpdateChannelClassInput{}
        output := &medialive.UpdateChannelClassOutput{}

        mockClient.On("UpdateChannelClass", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateChannelClass(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateCloudWatchAlarmTemplate", func(t *testing.T) {
        input := &medialive.UpdateCloudWatchAlarmTemplateInput{}
        output := &medialive.UpdateCloudWatchAlarmTemplateOutput{}

        mockClient.On("UpdateCloudWatchAlarmTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateCloudWatchAlarmTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateCloudWatchAlarmTemplateGroup", func(t *testing.T) {
        input := &medialive.UpdateCloudWatchAlarmTemplateGroupInput{}
        output := &medialive.UpdateCloudWatchAlarmTemplateGroupOutput{}

        mockClient.On("UpdateCloudWatchAlarmTemplateGroup", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateCloudWatchAlarmTemplateGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateEventBridgeRuleTemplate", func(t *testing.T) {
        input := &medialive.UpdateEventBridgeRuleTemplateInput{}
        output := &medialive.UpdateEventBridgeRuleTemplateOutput{}

        mockClient.On("UpdateEventBridgeRuleTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateEventBridgeRuleTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateEventBridgeRuleTemplateGroup", func(t *testing.T) {
        input := &medialive.UpdateEventBridgeRuleTemplateGroupInput{}
        output := &medialive.UpdateEventBridgeRuleTemplateGroupOutput{}

        mockClient.On("UpdateEventBridgeRuleTemplateGroup", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateEventBridgeRuleTemplateGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateInput", func(t *testing.T) {
        input := &medialive.UpdateInputInput{}
        output := &medialive.UpdateInputOutput{}

        mockClient.On("UpdateInput", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateInput(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateInputDevice", func(t *testing.T) {
        input := &medialive.UpdateInputDeviceInput{}
        output := &medialive.UpdateInputDeviceOutput{}

        mockClient.On("UpdateInputDevice", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateInputDevice(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateInputSecurityGroup", func(t *testing.T) {
        input := &medialive.UpdateInputSecurityGroupInput{}
        output := &medialive.UpdateInputSecurityGroupOutput{}

        mockClient.On("UpdateInputSecurityGroup", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateInputSecurityGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateMultiplex", func(t *testing.T) {
        input := &medialive.UpdateMultiplexInput{}
        output := &medialive.UpdateMultiplexOutput{}

        mockClient.On("UpdateMultiplex", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateMultiplex(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateMultiplexProgram", func(t *testing.T) {
        input := &medialive.UpdateMultiplexProgramInput{}
        output := &medialive.UpdateMultiplexProgramOutput{}

        mockClient.On("UpdateMultiplexProgram", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateMultiplexProgram(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateReservation", func(t *testing.T) {
        input := &medialive.UpdateReservationInput{}
        output := &medialive.UpdateReservationOutput{}

        mockClient.On("UpdateReservation", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateReservation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
