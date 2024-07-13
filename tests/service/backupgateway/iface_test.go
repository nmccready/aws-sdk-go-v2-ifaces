package backupgateway_test

// tests for the backupgateway service interface for this ../../../service/backupgateway/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/backupgateway"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/backupgateway/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/backupgateway/backupgateway_iface"
	"github.com/stretchr/testify/assert"
)

func TestBackupgatewayServiceCanBeMocked(t *testing.T) {
	var iface backupgateway_iface.IClient
	iface = &backupgateway.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := backupgateway.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateGatewayToServer", func(t *testing.T) {
        input := &backupgateway.AssociateGatewayToServerInput{}
        output := &backupgateway.AssociateGatewayToServerOutput{}

        mockClient.On("AssociateGatewayToServer", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateGatewayToServer(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateGateway", func(t *testing.T) {
        input := &backupgateway.CreateGatewayInput{}
        output := &backupgateway.CreateGatewayOutput{}

        mockClient.On("CreateGateway", ctx, input).Return(output, nil)

        result, err := mockClient.CreateGateway(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteGateway", func(t *testing.T) {
        input := &backupgateway.DeleteGatewayInput{}
        output := &backupgateway.DeleteGatewayOutput{}

        mockClient.On("DeleteGateway", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteGateway(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteHypervisor", func(t *testing.T) {
        input := &backupgateway.DeleteHypervisorInput{}
        output := &backupgateway.DeleteHypervisorOutput{}

        mockClient.On("DeleteHypervisor", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteHypervisor(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateGatewayFromServer", func(t *testing.T) {
        input := &backupgateway.DisassociateGatewayFromServerInput{}
        output := &backupgateway.DisassociateGatewayFromServerOutput{}

        mockClient.On("DisassociateGatewayFromServer", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateGatewayFromServer(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetBandwidthRateLimitSchedule", func(t *testing.T) {
        input := &backupgateway.GetBandwidthRateLimitScheduleInput{}
        output := &backupgateway.GetBandwidthRateLimitScheduleOutput{}

        mockClient.On("GetBandwidthRateLimitSchedule", ctx, input).Return(output, nil)

        result, err := mockClient.GetBandwidthRateLimitSchedule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetGateway", func(t *testing.T) {
        input := &backupgateway.GetGatewayInput{}
        output := &backupgateway.GetGatewayOutput{}

        mockClient.On("GetGateway", ctx, input).Return(output, nil)

        result, err := mockClient.GetGateway(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetHypervisor", func(t *testing.T) {
        input := &backupgateway.GetHypervisorInput{}
        output := &backupgateway.GetHypervisorOutput{}

        mockClient.On("GetHypervisor", ctx, input).Return(output, nil)

        result, err := mockClient.GetHypervisor(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetHypervisorPropertyMappings", func(t *testing.T) {
        input := &backupgateway.GetHypervisorPropertyMappingsInput{}
        output := &backupgateway.GetHypervisorPropertyMappingsOutput{}

        mockClient.On("GetHypervisorPropertyMappings", ctx, input).Return(output, nil)

        result, err := mockClient.GetHypervisorPropertyMappings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetVirtualMachine", func(t *testing.T) {
        input := &backupgateway.GetVirtualMachineInput{}
        output := &backupgateway.GetVirtualMachineOutput{}

        mockClient.On("GetVirtualMachine", ctx, input).Return(output, nil)

        result, err := mockClient.GetVirtualMachine(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestImportHypervisorConfiguration", func(t *testing.T) {
        input := &backupgateway.ImportHypervisorConfigurationInput{}
        output := &backupgateway.ImportHypervisorConfigurationOutput{}

        mockClient.On("ImportHypervisorConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.ImportHypervisorConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListGateways", func(t *testing.T) {
        input := &backupgateway.ListGatewaysInput{}
        output := &backupgateway.ListGatewaysOutput{}

        mockClient.On("ListGateways", ctx, input).Return(output, nil)

        result, err := mockClient.ListGateways(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListHypervisors", func(t *testing.T) {
        input := &backupgateway.ListHypervisorsInput{}
        output := &backupgateway.ListHypervisorsOutput{}

        mockClient.On("ListHypervisors", ctx, input).Return(output, nil)

        result, err := mockClient.ListHypervisors(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &backupgateway.ListTagsForResourceInput{}
        output := &backupgateway.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListVirtualMachines", func(t *testing.T) {
        input := &backupgateway.ListVirtualMachinesInput{}
        output := &backupgateway.ListVirtualMachinesOutput{}

        mockClient.On("ListVirtualMachines", ctx, input).Return(output, nil)

        result, err := mockClient.ListVirtualMachines(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutBandwidthRateLimitSchedule", func(t *testing.T) {
        input := &backupgateway.PutBandwidthRateLimitScheduleInput{}
        output := &backupgateway.PutBandwidthRateLimitScheduleOutput{}

        mockClient.On("PutBandwidthRateLimitSchedule", ctx, input).Return(output, nil)

        result, err := mockClient.PutBandwidthRateLimitSchedule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutHypervisorPropertyMappings", func(t *testing.T) {
        input := &backupgateway.PutHypervisorPropertyMappingsInput{}
        output := &backupgateway.PutHypervisorPropertyMappingsOutput{}

        mockClient.On("PutHypervisorPropertyMappings", ctx, input).Return(output, nil)

        result, err := mockClient.PutHypervisorPropertyMappings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutMaintenanceStartTime", func(t *testing.T) {
        input := &backupgateway.PutMaintenanceStartTimeInput{}
        output := &backupgateway.PutMaintenanceStartTimeOutput{}

        mockClient.On("PutMaintenanceStartTime", ctx, input).Return(output, nil)

        result, err := mockClient.PutMaintenanceStartTime(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartVirtualMachinesMetadataSync", func(t *testing.T) {
        input := &backupgateway.StartVirtualMachinesMetadataSyncInput{}
        output := &backupgateway.StartVirtualMachinesMetadataSyncOutput{}

        mockClient.On("StartVirtualMachinesMetadataSync", ctx, input).Return(output, nil)

        result, err := mockClient.StartVirtualMachinesMetadataSync(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &backupgateway.TagResourceInput{}
        output := &backupgateway.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTestHypervisorConfiguration", func(t *testing.T) {
        input := &backupgateway.TestHypervisorConfigurationInput{}
        output := &backupgateway.TestHypervisorConfigurationOutput{}

        mockClient.On("TestHypervisorConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.TestHypervisorConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &backupgateway.UntagResourceInput{}
        output := &backupgateway.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateGatewayInformation", func(t *testing.T) {
        input := &backupgateway.UpdateGatewayInformationInput{}
        output := &backupgateway.UpdateGatewayInformationOutput{}

        mockClient.On("UpdateGatewayInformation", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateGatewayInformation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateGatewaySoftwareNow", func(t *testing.T) {
        input := &backupgateway.UpdateGatewaySoftwareNowInput{}
        output := &backupgateway.UpdateGatewaySoftwareNowOutput{}

        mockClient.On("UpdateGatewaySoftwareNow", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateGatewaySoftwareNow(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateHypervisor", func(t *testing.T) {
        input := &backupgateway.UpdateHypervisorInput{}
        output := &backupgateway.UpdateHypervisorOutput{}

        mockClient.On("UpdateHypervisor", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateHypervisor(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
