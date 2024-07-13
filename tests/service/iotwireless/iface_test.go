package iotwireless_test

// tests for the iotwireless service interface for this ../../../service/iotwireless/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/iotwireless"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/iotwireless/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/iotwireless/iotwireless_iface"
	"github.com/stretchr/testify/assert"
)

func TestIotwirelessServiceCanBeMocked(t *testing.T) {
	var iface iotwireless_iface.IClient
	iface = &iotwireless.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := iotwireless.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateAwsAccountWithPartnerAccount", func(t *testing.T) {
        input := &iotwireless.AssociateAwsAccountWithPartnerAccountInput{}
        output := &iotwireless.AssociateAwsAccountWithPartnerAccountOutput{}

        mockClient.On("AssociateAwsAccountWithPartnerAccount", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateAwsAccountWithPartnerAccount(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateMulticastGroupWithFuotaTask", func(t *testing.T) {
        input := &iotwireless.AssociateMulticastGroupWithFuotaTaskInput{}
        output := &iotwireless.AssociateMulticastGroupWithFuotaTaskOutput{}

        mockClient.On("AssociateMulticastGroupWithFuotaTask", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateMulticastGroupWithFuotaTask(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateWirelessDeviceWithFuotaTask", func(t *testing.T) {
        input := &iotwireless.AssociateWirelessDeviceWithFuotaTaskInput{}
        output := &iotwireless.AssociateWirelessDeviceWithFuotaTaskOutput{}

        mockClient.On("AssociateWirelessDeviceWithFuotaTask", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateWirelessDeviceWithFuotaTask(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateWirelessDeviceWithMulticastGroup", func(t *testing.T) {
        input := &iotwireless.AssociateWirelessDeviceWithMulticastGroupInput{}
        output := &iotwireless.AssociateWirelessDeviceWithMulticastGroupOutput{}

        mockClient.On("AssociateWirelessDeviceWithMulticastGroup", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateWirelessDeviceWithMulticastGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateWirelessDeviceWithThing", func(t *testing.T) {
        input := &iotwireless.AssociateWirelessDeviceWithThingInput{}
        output := &iotwireless.AssociateWirelessDeviceWithThingOutput{}

        mockClient.On("AssociateWirelessDeviceWithThing", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateWirelessDeviceWithThing(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateWirelessGatewayWithCertificate", func(t *testing.T) {
        input := &iotwireless.AssociateWirelessGatewayWithCertificateInput{}
        output := &iotwireless.AssociateWirelessGatewayWithCertificateOutput{}

        mockClient.On("AssociateWirelessGatewayWithCertificate", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateWirelessGatewayWithCertificate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateWirelessGatewayWithThing", func(t *testing.T) {
        input := &iotwireless.AssociateWirelessGatewayWithThingInput{}
        output := &iotwireless.AssociateWirelessGatewayWithThingOutput{}

        mockClient.On("AssociateWirelessGatewayWithThing", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateWirelessGatewayWithThing(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCancelMulticastGroupSession", func(t *testing.T) {
        input := &iotwireless.CancelMulticastGroupSessionInput{}
        output := &iotwireless.CancelMulticastGroupSessionOutput{}

        mockClient.On("CancelMulticastGroupSession", ctx, input).Return(output, nil)

        result, err := mockClient.CancelMulticastGroupSession(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateDestination", func(t *testing.T) {
        input := &iotwireless.CreateDestinationInput{}
        output := &iotwireless.CreateDestinationOutput{}

        mockClient.On("CreateDestination", ctx, input).Return(output, nil)

        result, err := mockClient.CreateDestination(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateDeviceProfile", func(t *testing.T) {
        input := &iotwireless.CreateDeviceProfileInput{}
        output := &iotwireless.CreateDeviceProfileOutput{}

        mockClient.On("CreateDeviceProfile", ctx, input).Return(output, nil)

        result, err := mockClient.CreateDeviceProfile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateFuotaTask", func(t *testing.T) {
        input := &iotwireless.CreateFuotaTaskInput{}
        output := &iotwireless.CreateFuotaTaskOutput{}

        mockClient.On("CreateFuotaTask", ctx, input).Return(output, nil)

        result, err := mockClient.CreateFuotaTask(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateMulticastGroup", func(t *testing.T) {
        input := &iotwireless.CreateMulticastGroupInput{}
        output := &iotwireless.CreateMulticastGroupOutput{}

        mockClient.On("CreateMulticastGroup", ctx, input).Return(output, nil)

        result, err := mockClient.CreateMulticastGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateNetworkAnalyzerConfiguration", func(t *testing.T) {
        input := &iotwireless.CreateNetworkAnalyzerConfigurationInput{}
        output := &iotwireless.CreateNetworkAnalyzerConfigurationOutput{}

        mockClient.On("CreateNetworkAnalyzerConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.CreateNetworkAnalyzerConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateServiceProfile", func(t *testing.T) {
        input := &iotwireless.CreateServiceProfileInput{}
        output := &iotwireless.CreateServiceProfileOutput{}

        mockClient.On("CreateServiceProfile", ctx, input).Return(output, nil)

        result, err := mockClient.CreateServiceProfile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateWirelessDevice", func(t *testing.T) {
        input := &iotwireless.CreateWirelessDeviceInput{}
        output := &iotwireless.CreateWirelessDeviceOutput{}

        mockClient.On("CreateWirelessDevice", ctx, input).Return(output, nil)

        result, err := mockClient.CreateWirelessDevice(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateWirelessGateway", func(t *testing.T) {
        input := &iotwireless.CreateWirelessGatewayInput{}
        output := &iotwireless.CreateWirelessGatewayOutput{}

        mockClient.On("CreateWirelessGateway", ctx, input).Return(output, nil)

        result, err := mockClient.CreateWirelessGateway(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateWirelessGatewayTask", func(t *testing.T) {
        input := &iotwireless.CreateWirelessGatewayTaskInput{}
        output := &iotwireless.CreateWirelessGatewayTaskOutput{}

        mockClient.On("CreateWirelessGatewayTask", ctx, input).Return(output, nil)

        result, err := mockClient.CreateWirelessGatewayTask(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateWirelessGatewayTaskDefinition", func(t *testing.T) {
        input := &iotwireless.CreateWirelessGatewayTaskDefinitionInput{}
        output := &iotwireless.CreateWirelessGatewayTaskDefinitionOutput{}

        mockClient.On("CreateWirelessGatewayTaskDefinition", ctx, input).Return(output, nil)

        result, err := mockClient.CreateWirelessGatewayTaskDefinition(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteDestination", func(t *testing.T) {
        input := &iotwireless.DeleteDestinationInput{}
        output := &iotwireless.DeleteDestinationOutput{}

        mockClient.On("DeleteDestination", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteDestination(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteDeviceProfile", func(t *testing.T) {
        input := &iotwireless.DeleteDeviceProfileInput{}
        output := &iotwireless.DeleteDeviceProfileOutput{}

        mockClient.On("DeleteDeviceProfile", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteDeviceProfile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteFuotaTask", func(t *testing.T) {
        input := &iotwireless.DeleteFuotaTaskInput{}
        output := &iotwireless.DeleteFuotaTaskOutput{}

        mockClient.On("DeleteFuotaTask", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteFuotaTask(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteMulticastGroup", func(t *testing.T) {
        input := &iotwireless.DeleteMulticastGroupInput{}
        output := &iotwireless.DeleteMulticastGroupOutput{}

        mockClient.On("DeleteMulticastGroup", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteMulticastGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteNetworkAnalyzerConfiguration", func(t *testing.T) {
        input := &iotwireless.DeleteNetworkAnalyzerConfigurationInput{}
        output := &iotwireless.DeleteNetworkAnalyzerConfigurationOutput{}

        mockClient.On("DeleteNetworkAnalyzerConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteNetworkAnalyzerConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteQueuedMessages", func(t *testing.T) {
        input := &iotwireless.DeleteQueuedMessagesInput{}
        output := &iotwireless.DeleteQueuedMessagesOutput{}

        mockClient.On("DeleteQueuedMessages", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteQueuedMessages(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteServiceProfile", func(t *testing.T) {
        input := &iotwireless.DeleteServiceProfileInput{}
        output := &iotwireless.DeleteServiceProfileOutput{}

        mockClient.On("DeleteServiceProfile", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteServiceProfile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteWirelessDevice", func(t *testing.T) {
        input := &iotwireless.DeleteWirelessDeviceInput{}
        output := &iotwireless.DeleteWirelessDeviceOutput{}

        mockClient.On("DeleteWirelessDevice", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteWirelessDevice(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteWirelessDeviceImportTask", func(t *testing.T) {
        input := &iotwireless.DeleteWirelessDeviceImportTaskInput{}
        output := &iotwireless.DeleteWirelessDeviceImportTaskOutput{}

        mockClient.On("DeleteWirelessDeviceImportTask", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteWirelessDeviceImportTask(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteWirelessGateway", func(t *testing.T) {
        input := &iotwireless.DeleteWirelessGatewayInput{}
        output := &iotwireless.DeleteWirelessGatewayOutput{}

        mockClient.On("DeleteWirelessGateway", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteWirelessGateway(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteWirelessGatewayTask", func(t *testing.T) {
        input := &iotwireless.DeleteWirelessGatewayTaskInput{}
        output := &iotwireless.DeleteWirelessGatewayTaskOutput{}

        mockClient.On("DeleteWirelessGatewayTask", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteWirelessGatewayTask(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteWirelessGatewayTaskDefinition", func(t *testing.T) {
        input := &iotwireless.DeleteWirelessGatewayTaskDefinitionInput{}
        output := &iotwireless.DeleteWirelessGatewayTaskDefinitionOutput{}

        mockClient.On("DeleteWirelessGatewayTaskDefinition", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteWirelessGatewayTaskDefinition(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeregisterWirelessDevice", func(t *testing.T) {
        input := &iotwireless.DeregisterWirelessDeviceInput{}
        output := &iotwireless.DeregisterWirelessDeviceOutput{}

        mockClient.On("DeregisterWirelessDevice", ctx, input).Return(output, nil)

        result, err := mockClient.DeregisterWirelessDevice(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateAwsAccountFromPartnerAccount", func(t *testing.T) {
        input := &iotwireless.DisassociateAwsAccountFromPartnerAccountInput{}
        output := &iotwireless.DisassociateAwsAccountFromPartnerAccountOutput{}

        mockClient.On("DisassociateAwsAccountFromPartnerAccount", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateAwsAccountFromPartnerAccount(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateMulticastGroupFromFuotaTask", func(t *testing.T) {
        input := &iotwireless.DisassociateMulticastGroupFromFuotaTaskInput{}
        output := &iotwireless.DisassociateMulticastGroupFromFuotaTaskOutput{}

        mockClient.On("DisassociateMulticastGroupFromFuotaTask", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateMulticastGroupFromFuotaTask(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateWirelessDeviceFromFuotaTask", func(t *testing.T) {
        input := &iotwireless.DisassociateWirelessDeviceFromFuotaTaskInput{}
        output := &iotwireless.DisassociateWirelessDeviceFromFuotaTaskOutput{}

        mockClient.On("DisassociateWirelessDeviceFromFuotaTask", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateWirelessDeviceFromFuotaTask(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateWirelessDeviceFromMulticastGroup", func(t *testing.T) {
        input := &iotwireless.DisassociateWirelessDeviceFromMulticastGroupInput{}
        output := &iotwireless.DisassociateWirelessDeviceFromMulticastGroupOutput{}

        mockClient.On("DisassociateWirelessDeviceFromMulticastGroup", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateWirelessDeviceFromMulticastGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateWirelessDeviceFromThing", func(t *testing.T) {
        input := &iotwireless.DisassociateWirelessDeviceFromThingInput{}
        output := &iotwireless.DisassociateWirelessDeviceFromThingOutput{}

        mockClient.On("DisassociateWirelessDeviceFromThing", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateWirelessDeviceFromThing(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateWirelessGatewayFromCertificate", func(t *testing.T) {
        input := &iotwireless.DisassociateWirelessGatewayFromCertificateInput{}
        output := &iotwireless.DisassociateWirelessGatewayFromCertificateOutput{}

        mockClient.On("DisassociateWirelessGatewayFromCertificate", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateWirelessGatewayFromCertificate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateWirelessGatewayFromThing", func(t *testing.T) {
        input := &iotwireless.DisassociateWirelessGatewayFromThingInput{}
        output := &iotwireless.DisassociateWirelessGatewayFromThingOutput{}

        mockClient.On("DisassociateWirelessGatewayFromThing", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateWirelessGatewayFromThing(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDestination", func(t *testing.T) {
        input := &iotwireless.GetDestinationInput{}
        output := &iotwireless.GetDestinationOutput{}

        mockClient.On("GetDestination", ctx, input).Return(output, nil)

        result, err := mockClient.GetDestination(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDeviceProfile", func(t *testing.T) {
        input := &iotwireless.GetDeviceProfileInput{}
        output := &iotwireless.GetDeviceProfileOutput{}

        mockClient.On("GetDeviceProfile", ctx, input).Return(output, nil)

        result, err := mockClient.GetDeviceProfile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetEventConfigurationByResourceTypes", func(t *testing.T) {
        input := &iotwireless.GetEventConfigurationByResourceTypesInput{}
        output := &iotwireless.GetEventConfigurationByResourceTypesOutput{}

        mockClient.On("GetEventConfigurationByResourceTypes", ctx, input).Return(output, nil)

        result, err := mockClient.GetEventConfigurationByResourceTypes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetFuotaTask", func(t *testing.T) {
        input := &iotwireless.GetFuotaTaskInput{}
        output := &iotwireless.GetFuotaTaskOutput{}

        mockClient.On("GetFuotaTask", ctx, input).Return(output, nil)

        result, err := mockClient.GetFuotaTask(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetLogLevelsByResourceTypes", func(t *testing.T) {
        input := &iotwireless.GetLogLevelsByResourceTypesInput{}
        output := &iotwireless.GetLogLevelsByResourceTypesOutput{}

        mockClient.On("GetLogLevelsByResourceTypes", ctx, input).Return(output, nil)

        result, err := mockClient.GetLogLevelsByResourceTypes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetMetricConfiguration", func(t *testing.T) {
        input := &iotwireless.GetMetricConfigurationInput{}
        output := &iotwireless.GetMetricConfigurationOutput{}

        mockClient.On("GetMetricConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.GetMetricConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetMetrics", func(t *testing.T) {
        input := &iotwireless.GetMetricsInput{}
        output := &iotwireless.GetMetricsOutput{}

        mockClient.On("GetMetrics", ctx, input).Return(output, nil)

        result, err := mockClient.GetMetrics(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetMulticastGroup", func(t *testing.T) {
        input := &iotwireless.GetMulticastGroupInput{}
        output := &iotwireless.GetMulticastGroupOutput{}

        mockClient.On("GetMulticastGroup", ctx, input).Return(output, nil)

        result, err := mockClient.GetMulticastGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetMulticastGroupSession", func(t *testing.T) {
        input := &iotwireless.GetMulticastGroupSessionInput{}
        output := &iotwireless.GetMulticastGroupSessionOutput{}

        mockClient.On("GetMulticastGroupSession", ctx, input).Return(output, nil)

        result, err := mockClient.GetMulticastGroupSession(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetNetworkAnalyzerConfiguration", func(t *testing.T) {
        input := &iotwireless.GetNetworkAnalyzerConfigurationInput{}
        output := &iotwireless.GetNetworkAnalyzerConfigurationOutput{}

        mockClient.On("GetNetworkAnalyzerConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.GetNetworkAnalyzerConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetPartnerAccount", func(t *testing.T) {
        input := &iotwireless.GetPartnerAccountInput{}
        output := &iotwireless.GetPartnerAccountOutput{}

        mockClient.On("GetPartnerAccount", ctx, input).Return(output, nil)

        result, err := mockClient.GetPartnerAccount(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetPosition", func(t *testing.T) {
        input := &iotwireless.GetPositionInput{}
        output := &iotwireless.GetPositionOutput{}

        mockClient.On("GetPosition", ctx, input).Return(output, nil)

        result, err := mockClient.GetPosition(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetPositionConfiguration", func(t *testing.T) {
        input := &iotwireless.GetPositionConfigurationInput{}
        output := &iotwireless.GetPositionConfigurationOutput{}

        mockClient.On("GetPositionConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.GetPositionConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetPositionEstimate", func(t *testing.T) {
        input := &iotwireless.GetPositionEstimateInput{}
        output := &iotwireless.GetPositionEstimateOutput{}

        mockClient.On("GetPositionEstimate", ctx, input).Return(output, nil)

        result, err := mockClient.GetPositionEstimate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetResourceEventConfiguration", func(t *testing.T) {
        input := &iotwireless.GetResourceEventConfigurationInput{}
        output := &iotwireless.GetResourceEventConfigurationOutput{}

        mockClient.On("GetResourceEventConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.GetResourceEventConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetResourceLogLevel", func(t *testing.T) {
        input := &iotwireless.GetResourceLogLevelInput{}
        output := &iotwireless.GetResourceLogLevelOutput{}

        mockClient.On("GetResourceLogLevel", ctx, input).Return(output, nil)

        result, err := mockClient.GetResourceLogLevel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetResourcePosition", func(t *testing.T) {
        input := &iotwireless.GetResourcePositionInput{}
        output := &iotwireless.GetResourcePositionOutput{}

        mockClient.On("GetResourcePosition", ctx, input).Return(output, nil)

        result, err := mockClient.GetResourcePosition(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetServiceEndpoint", func(t *testing.T) {
        input := &iotwireless.GetServiceEndpointInput{}
        output := &iotwireless.GetServiceEndpointOutput{}

        mockClient.On("GetServiceEndpoint", ctx, input).Return(output, nil)

        result, err := mockClient.GetServiceEndpoint(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetServiceProfile", func(t *testing.T) {
        input := &iotwireless.GetServiceProfileInput{}
        output := &iotwireless.GetServiceProfileOutput{}

        mockClient.On("GetServiceProfile", ctx, input).Return(output, nil)

        result, err := mockClient.GetServiceProfile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetWirelessDevice", func(t *testing.T) {
        input := &iotwireless.GetWirelessDeviceInput{}
        output := &iotwireless.GetWirelessDeviceOutput{}

        mockClient.On("GetWirelessDevice", ctx, input).Return(output, nil)

        result, err := mockClient.GetWirelessDevice(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetWirelessDeviceImportTask", func(t *testing.T) {
        input := &iotwireless.GetWirelessDeviceImportTaskInput{}
        output := &iotwireless.GetWirelessDeviceImportTaskOutput{}

        mockClient.On("GetWirelessDeviceImportTask", ctx, input).Return(output, nil)

        result, err := mockClient.GetWirelessDeviceImportTask(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetWirelessDeviceStatistics", func(t *testing.T) {
        input := &iotwireless.GetWirelessDeviceStatisticsInput{}
        output := &iotwireless.GetWirelessDeviceStatisticsOutput{}

        mockClient.On("GetWirelessDeviceStatistics", ctx, input).Return(output, nil)

        result, err := mockClient.GetWirelessDeviceStatistics(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetWirelessGateway", func(t *testing.T) {
        input := &iotwireless.GetWirelessGatewayInput{}
        output := &iotwireless.GetWirelessGatewayOutput{}

        mockClient.On("GetWirelessGateway", ctx, input).Return(output, nil)

        result, err := mockClient.GetWirelessGateway(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetWirelessGatewayCertificate", func(t *testing.T) {
        input := &iotwireless.GetWirelessGatewayCertificateInput{}
        output := &iotwireless.GetWirelessGatewayCertificateOutput{}

        mockClient.On("GetWirelessGatewayCertificate", ctx, input).Return(output, nil)

        result, err := mockClient.GetWirelessGatewayCertificate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetWirelessGatewayFirmwareInformation", func(t *testing.T) {
        input := &iotwireless.GetWirelessGatewayFirmwareInformationInput{}
        output := &iotwireless.GetWirelessGatewayFirmwareInformationOutput{}

        mockClient.On("GetWirelessGatewayFirmwareInformation", ctx, input).Return(output, nil)

        result, err := mockClient.GetWirelessGatewayFirmwareInformation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetWirelessGatewayStatistics", func(t *testing.T) {
        input := &iotwireless.GetWirelessGatewayStatisticsInput{}
        output := &iotwireless.GetWirelessGatewayStatisticsOutput{}

        mockClient.On("GetWirelessGatewayStatistics", ctx, input).Return(output, nil)

        result, err := mockClient.GetWirelessGatewayStatistics(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetWirelessGatewayTask", func(t *testing.T) {
        input := &iotwireless.GetWirelessGatewayTaskInput{}
        output := &iotwireless.GetWirelessGatewayTaskOutput{}

        mockClient.On("GetWirelessGatewayTask", ctx, input).Return(output, nil)

        result, err := mockClient.GetWirelessGatewayTask(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetWirelessGatewayTaskDefinition", func(t *testing.T) {
        input := &iotwireless.GetWirelessGatewayTaskDefinitionInput{}
        output := &iotwireless.GetWirelessGatewayTaskDefinitionOutput{}

        mockClient.On("GetWirelessGatewayTaskDefinition", ctx, input).Return(output, nil)

        result, err := mockClient.GetWirelessGatewayTaskDefinition(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDestinations", func(t *testing.T) {
        input := &iotwireless.ListDestinationsInput{}
        output := &iotwireless.ListDestinationsOutput{}

        mockClient.On("ListDestinations", ctx, input).Return(output, nil)

        result, err := mockClient.ListDestinations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDeviceProfiles", func(t *testing.T) {
        input := &iotwireless.ListDeviceProfilesInput{}
        output := &iotwireless.ListDeviceProfilesOutput{}

        mockClient.On("ListDeviceProfiles", ctx, input).Return(output, nil)

        result, err := mockClient.ListDeviceProfiles(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDevicesForWirelessDeviceImportTask", func(t *testing.T) {
        input := &iotwireless.ListDevicesForWirelessDeviceImportTaskInput{}
        output := &iotwireless.ListDevicesForWirelessDeviceImportTaskOutput{}

        mockClient.On("ListDevicesForWirelessDeviceImportTask", ctx, input).Return(output, nil)

        result, err := mockClient.ListDevicesForWirelessDeviceImportTask(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListEventConfigurations", func(t *testing.T) {
        input := &iotwireless.ListEventConfigurationsInput{}
        output := &iotwireless.ListEventConfigurationsOutput{}

        mockClient.On("ListEventConfigurations", ctx, input).Return(output, nil)

        result, err := mockClient.ListEventConfigurations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListFuotaTasks", func(t *testing.T) {
        input := &iotwireless.ListFuotaTasksInput{}
        output := &iotwireless.ListFuotaTasksOutput{}

        mockClient.On("ListFuotaTasks", ctx, input).Return(output, nil)

        result, err := mockClient.ListFuotaTasks(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListMulticastGroups", func(t *testing.T) {
        input := &iotwireless.ListMulticastGroupsInput{}
        output := &iotwireless.ListMulticastGroupsOutput{}

        mockClient.On("ListMulticastGroups", ctx, input).Return(output, nil)

        result, err := mockClient.ListMulticastGroups(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListMulticastGroupsByFuotaTask", func(t *testing.T) {
        input := &iotwireless.ListMulticastGroupsByFuotaTaskInput{}
        output := &iotwireless.ListMulticastGroupsByFuotaTaskOutput{}

        mockClient.On("ListMulticastGroupsByFuotaTask", ctx, input).Return(output, nil)

        result, err := mockClient.ListMulticastGroupsByFuotaTask(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListNetworkAnalyzerConfigurations", func(t *testing.T) {
        input := &iotwireless.ListNetworkAnalyzerConfigurationsInput{}
        output := &iotwireless.ListNetworkAnalyzerConfigurationsOutput{}

        mockClient.On("ListNetworkAnalyzerConfigurations", ctx, input).Return(output, nil)

        result, err := mockClient.ListNetworkAnalyzerConfigurations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListPartnerAccounts", func(t *testing.T) {
        input := &iotwireless.ListPartnerAccountsInput{}
        output := &iotwireless.ListPartnerAccountsOutput{}

        mockClient.On("ListPartnerAccounts", ctx, input).Return(output, nil)

        result, err := mockClient.ListPartnerAccounts(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListPositionConfigurations", func(t *testing.T) {
        input := &iotwireless.ListPositionConfigurationsInput{}
        output := &iotwireless.ListPositionConfigurationsOutput{}

        mockClient.On("ListPositionConfigurations", ctx, input).Return(output, nil)

        result, err := mockClient.ListPositionConfigurations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListQueuedMessages", func(t *testing.T) {
        input := &iotwireless.ListQueuedMessagesInput{}
        output := &iotwireless.ListQueuedMessagesOutput{}

        mockClient.On("ListQueuedMessages", ctx, input).Return(output, nil)

        result, err := mockClient.ListQueuedMessages(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListServiceProfiles", func(t *testing.T) {
        input := &iotwireless.ListServiceProfilesInput{}
        output := &iotwireless.ListServiceProfilesOutput{}

        mockClient.On("ListServiceProfiles", ctx, input).Return(output, nil)

        result, err := mockClient.ListServiceProfiles(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &iotwireless.ListTagsForResourceInput{}
        output := &iotwireless.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListWirelessDeviceImportTasks", func(t *testing.T) {
        input := &iotwireless.ListWirelessDeviceImportTasksInput{}
        output := &iotwireless.ListWirelessDeviceImportTasksOutput{}

        mockClient.On("ListWirelessDeviceImportTasks", ctx, input).Return(output, nil)

        result, err := mockClient.ListWirelessDeviceImportTasks(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListWirelessDevices", func(t *testing.T) {
        input := &iotwireless.ListWirelessDevicesInput{}
        output := &iotwireless.ListWirelessDevicesOutput{}

        mockClient.On("ListWirelessDevices", ctx, input).Return(output, nil)

        result, err := mockClient.ListWirelessDevices(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListWirelessGatewayTaskDefinitions", func(t *testing.T) {
        input := &iotwireless.ListWirelessGatewayTaskDefinitionsInput{}
        output := &iotwireless.ListWirelessGatewayTaskDefinitionsOutput{}

        mockClient.On("ListWirelessGatewayTaskDefinitions", ctx, input).Return(output, nil)

        result, err := mockClient.ListWirelessGatewayTaskDefinitions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListWirelessGateways", func(t *testing.T) {
        input := &iotwireless.ListWirelessGatewaysInput{}
        output := &iotwireless.ListWirelessGatewaysOutput{}

        mockClient.On("ListWirelessGateways", ctx, input).Return(output, nil)

        result, err := mockClient.ListWirelessGateways(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutPositionConfiguration", func(t *testing.T) {
        input := &iotwireless.PutPositionConfigurationInput{}
        output := &iotwireless.PutPositionConfigurationOutput{}

        mockClient.On("PutPositionConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.PutPositionConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutResourceLogLevel", func(t *testing.T) {
        input := &iotwireless.PutResourceLogLevelInput{}
        output := &iotwireless.PutResourceLogLevelOutput{}

        mockClient.On("PutResourceLogLevel", ctx, input).Return(output, nil)

        result, err := mockClient.PutResourceLogLevel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestResetAllResourceLogLevels", func(t *testing.T) {
        input := &iotwireless.ResetAllResourceLogLevelsInput{}
        output := &iotwireless.ResetAllResourceLogLevelsOutput{}

        mockClient.On("ResetAllResourceLogLevels", ctx, input).Return(output, nil)

        result, err := mockClient.ResetAllResourceLogLevels(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestResetResourceLogLevel", func(t *testing.T) {
        input := &iotwireless.ResetResourceLogLevelInput{}
        output := &iotwireless.ResetResourceLogLevelOutput{}

        mockClient.On("ResetResourceLogLevel", ctx, input).Return(output, nil)

        result, err := mockClient.ResetResourceLogLevel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSendDataToMulticastGroup", func(t *testing.T) {
        input := &iotwireless.SendDataToMulticastGroupInput{}
        output := &iotwireless.SendDataToMulticastGroupOutput{}

        mockClient.On("SendDataToMulticastGroup", ctx, input).Return(output, nil)

        result, err := mockClient.SendDataToMulticastGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSendDataToWirelessDevice", func(t *testing.T) {
        input := &iotwireless.SendDataToWirelessDeviceInput{}
        output := &iotwireless.SendDataToWirelessDeviceOutput{}

        mockClient.On("SendDataToWirelessDevice", ctx, input).Return(output, nil)

        result, err := mockClient.SendDataToWirelessDevice(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartBulkAssociateWirelessDeviceWithMulticastGroup", func(t *testing.T) {
        input := &iotwireless.StartBulkAssociateWirelessDeviceWithMulticastGroupInput{}
        output := &iotwireless.StartBulkAssociateWirelessDeviceWithMulticastGroupOutput{}

        mockClient.On("StartBulkAssociateWirelessDeviceWithMulticastGroup", ctx, input).Return(output, nil)

        result, err := mockClient.StartBulkAssociateWirelessDeviceWithMulticastGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartBulkDisassociateWirelessDeviceFromMulticastGroup", func(t *testing.T) {
        input := &iotwireless.StartBulkDisassociateWirelessDeviceFromMulticastGroupInput{}
        output := &iotwireless.StartBulkDisassociateWirelessDeviceFromMulticastGroupOutput{}

        mockClient.On("StartBulkDisassociateWirelessDeviceFromMulticastGroup", ctx, input).Return(output, nil)

        result, err := mockClient.StartBulkDisassociateWirelessDeviceFromMulticastGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartFuotaTask", func(t *testing.T) {
        input := &iotwireless.StartFuotaTaskInput{}
        output := &iotwireless.StartFuotaTaskOutput{}

        mockClient.On("StartFuotaTask", ctx, input).Return(output, nil)

        result, err := mockClient.StartFuotaTask(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartMulticastGroupSession", func(t *testing.T) {
        input := &iotwireless.StartMulticastGroupSessionInput{}
        output := &iotwireless.StartMulticastGroupSessionOutput{}

        mockClient.On("StartMulticastGroupSession", ctx, input).Return(output, nil)

        result, err := mockClient.StartMulticastGroupSession(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartSingleWirelessDeviceImportTask", func(t *testing.T) {
        input := &iotwireless.StartSingleWirelessDeviceImportTaskInput{}
        output := &iotwireless.StartSingleWirelessDeviceImportTaskOutput{}

        mockClient.On("StartSingleWirelessDeviceImportTask", ctx, input).Return(output, nil)

        result, err := mockClient.StartSingleWirelessDeviceImportTask(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartWirelessDeviceImportTask", func(t *testing.T) {
        input := &iotwireless.StartWirelessDeviceImportTaskInput{}
        output := &iotwireless.StartWirelessDeviceImportTaskOutput{}

        mockClient.On("StartWirelessDeviceImportTask", ctx, input).Return(output, nil)

        result, err := mockClient.StartWirelessDeviceImportTask(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &iotwireless.TagResourceInput{}
        output := &iotwireless.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTestWirelessDevice", func(t *testing.T) {
        input := &iotwireless.TestWirelessDeviceInput{}
        output := &iotwireless.TestWirelessDeviceOutput{}

        mockClient.On("TestWirelessDevice", ctx, input).Return(output, nil)

        result, err := mockClient.TestWirelessDevice(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &iotwireless.UntagResourceInput{}
        output := &iotwireless.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateDestination", func(t *testing.T) {
        input := &iotwireless.UpdateDestinationInput{}
        output := &iotwireless.UpdateDestinationOutput{}

        mockClient.On("UpdateDestination", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateDestination(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateEventConfigurationByResourceTypes", func(t *testing.T) {
        input := &iotwireless.UpdateEventConfigurationByResourceTypesInput{}
        output := &iotwireless.UpdateEventConfigurationByResourceTypesOutput{}

        mockClient.On("UpdateEventConfigurationByResourceTypes", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateEventConfigurationByResourceTypes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateFuotaTask", func(t *testing.T) {
        input := &iotwireless.UpdateFuotaTaskInput{}
        output := &iotwireless.UpdateFuotaTaskOutput{}

        mockClient.On("UpdateFuotaTask", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateFuotaTask(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateLogLevelsByResourceTypes", func(t *testing.T) {
        input := &iotwireless.UpdateLogLevelsByResourceTypesInput{}
        output := &iotwireless.UpdateLogLevelsByResourceTypesOutput{}

        mockClient.On("UpdateLogLevelsByResourceTypes", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateLogLevelsByResourceTypes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateMetricConfiguration", func(t *testing.T) {
        input := &iotwireless.UpdateMetricConfigurationInput{}
        output := &iotwireless.UpdateMetricConfigurationOutput{}

        mockClient.On("UpdateMetricConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateMetricConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateMulticastGroup", func(t *testing.T) {
        input := &iotwireless.UpdateMulticastGroupInput{}
        output := &iotwireless.UpdateMulticastGroupOutput{}

        mockClient.On("UpdateMulticastGroup", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateMulticastGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateNetworkAnalyzerConfiguration", func(t *testing.T) {
        input := &iotwireless.UpdateNetworkAnalyzerConfigurationInput{}
        output := &iotwireless.UpdateNetworkAnalyzerConfigurationOutput{}

        mockClient.On("UpdateNetworkAnalyzerConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateNetworkAnalyzerConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdatePartnerAccount", func(t *testing.T) {
        input := &iotwireless.UpdatePartnerAccountInput{}
        output := &iotwireless.UpdatePartnerAccountOutput{}

        mockClient.On("UpdatePartnerAccount", ctx, input).Return(output, nil)

        result, err := mockClient.UpdatePartnerAccount(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdatePosition", func(t *testing.T) {
        input := &iotwireless.UpdatePositionInput{}
        output := &iotwireless.UpdatePositionOutput{}

        mockClient.On("UpdatePosition", ctx, input).Return(output, nil)

        result, err := mockClient.UpdatePosition(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateResourceEventConfiguration", func(t *testing.T) {
        input := &iotwireless.UpdateResourceEventConfigurationInput{}
        output := &iotwireless.UpdateResourceEventConfigurationOutput{}

        mockClient.On("UpdateResourceEventConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateResourceEventConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateResourcePosition", func(t *testing.T) {
        input := &iotwireless.UpdateResourcePositionInput{}
        output := &iotwireless.UpdateResourcePositionOutput{}

        mockClient.On("UpdateResourcePosition", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateResourcePosition(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateWirelessDevice", func(t *testing.T) {
        input := &iotwireless.UpdateWirelessDeviceInput{}
        output := &iotwireless.UpdateWirelessDeviceOutput{}

        mockClient.On("UpdateWirelessDevice", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateWirelessDevice(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateWirelessDeviceImportTask", func(t *testing.T) {
        input := &iotwireless.UpdateWirelessDeviceImportTaskInput{}
        output := &iotwireless.UpdateWirelessDeviceImportTaskOutput{}

        mockClient.On("UpdateWirelessDeviceImportTask", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateWirelessDeviceImportTask(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateWirelessGateway", func(t *testing.T) {
        input := &iotwireless.UpdateWirelessGatewayInput{}
        output := &iotwireless.UpdateWirelessGatewayOutput{}

        mockClient.On("UpdateWirelessGateway", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateWirelessGateway(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
