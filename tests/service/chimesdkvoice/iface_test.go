package chimesdkvoice_test

// tests for the chimesdkvoice service interface for this ../../../service/chimesdkvoice/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/chimesdkvoice"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/chimesdkvoice/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/chimesdkvoice/chimesdkvoice_iface"
	"github.com/stretchr/testify/assert"
)

func TestChimesdkvoiceServiceCanBeMocked(t *testing.T) {
	var iface chimesdkvoice_iface.IClient
	iface = &chimesdkvoice.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := chimesdkvoice.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociatePhoneNumbersWithVoiceConnector", func(t *testing.T) {
        input := &chimesdkvoice.AssociatePhoneNumbersWithVoiceConnectorInput{}
        output := &chimesdkvoice.AssociatePhoneNumbersWithVoiceConnectorOutput{}

        mockClient.On("AssociatePhoneNumbersWithVoiceConnector", ctx, input).Return(output, nil)

        result, err := mockClient.AssociatePhoneNumbersWithVoiceConnector(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociatePhoneNumbersWithVoiceConnectorGroup", func(t *testing.T) {
        input := &chimesdkvoice.AssociatePhoneNumbersWithVoiceConnectorGroupInput{}
        output := &chimesdkvoice.AssociatePhoneNumbersWithVoiceConnectorGroupOutput{}

        mockClient.On("AssociatePhoneNumbersWithVoiceConnectorGroup", ctx, input).Return(output, nil)

        result, err := mockClient.AssociatePhoneNumbersWithVoiceConnectorGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchDeletePhoneNumber", func(t *testing.T) {
        input := &chimesdkvoice.BatchDeletePhoneNumberInput{}
        output := &chimesdkvoice.BatchDeletePhoneNumberOutput{}

        mockClient.On("BatchDeletePhoneNumber", ctx, input).Return(output, nil)

        result, err := mockClient.BatchDeletePhoneNumber(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchUpdatePhoneNumber", func(t *testing.T) {
        input := &chimesdkvoice.BatchUpdatePhoneNumberInput{}
        output := &chimesdkvoice.BatchUpdatePhoneNumberOutput{}

        mockClient.On("BatchUpdatePhoneNumber", ctx, input).Return(output, nil)

        result, err := mockClient.BatchUpdatePhoneNumber(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreatePhoneNumberOrder", func(t *testing.T) {
        input := &chimesdkvoice.CreatePhoneNumberOrderInput{}
        output := &chimesdkvoice.CreatePhoneNumberOrderOutput{}

        mockClient.On("CreatePhoneNumberOrder", ctx, input).Return(output, nil)

        result, err := mockClient.CreatePhoneNumberOrder(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateProxySession", func(t *testing.T) {
        input := &chimesdkvoice.CreateProxySessionInput{}
        output := &chimesdkvoice.CreateProxySessionOutput{}

        mockClient.On("CreateProxySession", ctx, input).Return(output, nil)

        result, err := mockClient.CreateProxySession(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateSipMediaApplication", func(t *testing.T) {
        input := &chimesdkvoice.CreateSipMediaApplicationInput{}
        output := &chimesdkvoice.CreateSipMediaApplicationOutput{}

        mockClient.On("CreateSipMediaApplication", ctx, input).Return(output, nil)

        result, err := mockClient.CreateSipMediaApplication(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateSipMediaApplicationCall", func(t *testing.T) {
        input := &chimesdkvoice.CreateSipMediaApplicationCallInput{}
        output := &chimesdkvoice.CreateSipMediaApplicationCallOutput{}

        mockClient.On("CreateSipMediaApplicationCall", ctx, input).Return(output, nil)

        result, err := mockClient.CreateSipMediaApplicationCall(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateSipRule", func(t *testing.T) {
        input := &chimesdkvoice.CreateSipRuleInput{}
        output := &chimesdkvoice.CreateSipRuleOutput{}

        mockClient.On("CreateSipRule", ctx, input).Return(output, nil)

        result, err := mockClient.CreateSipRule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateVoiceConnector", func(t *testing.T) {
        input := &chimesdkvoice.CreateVoiceConnectorInput{}
        output := &chimesdkvoice.CreateVoiceConnectorOutput{}

        mockClient.On("CreateVoiceConnector", ctx, input).Return(output, nil)

        result, err := mockClient.CreateVoiceConnector(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateVoiceConnectorGroup", func(t *testing.T) {
        input := &chimesdkvoice.CreateVoiceConnectorGroupInput{}
        output := &chimesdkvoice.CreateVoiceConnectorGroupOutput{}

        mockClient.On("CreateVoiceConnectorGroup", ctx, input).Return(output, nil)

        result, err := mockClient.CreateVoiceConnectorGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateVoiceProfile", func(t *testing.T) {
        input := &chimesdkvoice.CreateVoiceProfileInput{}
        output := &chimesdkvoice.CreateVoiceProfileOutput{}

        mockClient.On("CreateVoiceProfile", ctx, input).Return(output, nil)

        result, err := mockClient.CreateVoiceProfile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateVoiceProfileDomain", func(t *testing.T) {
        input := &chimesdkvoice.CreateVoiceProfileDomainInput{}
        output := &chimesdkvoice.CreateVoiceProfileDomainOutput{}

        mockClient.On("CreateVoiceProfileDomain", ctx, input).Return(output, nil)

        result, err := mockClient.CreateVoiceProfileDomain(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeletePhoneNumber", func(t *testing.T) {
        input := &chimesdkvoice.DeletePhoneNumberInput{}
        output := &chimesdkvoice.DeletePhoneNumberOutput{}

        mockClient.On("DeletePhoneNumber", ctx, input).Return(output, nil)

        result, err := mockClient.DeletePhoneNumber(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteProxySession", func(t *testing.T) {
        input := &chimesdkvoice.DeleteProxySessionInput{}
        output := &chimesdkvoice.DeleteProxySessionOutput{}

        mockClient.On("DeleteProxySession", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteProxySession(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteSipMediaApplication", func(t *testing.T) {
        input := &chimesdkvoice.DeleteSipMediaApplicationInput{}
        output := &chimesdkvoice.DeleteSipMediaApplicationOutput{}

        mockClient.On("DeleteSipMediaApplication", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteSipMediaApplication(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteSipRule", func(t *testing.T) {
        input := &chimesdkvoice.DeleteSipRuleInput{}
        output := &chimesdkvoice.DeleteSipRuleOutput{}

        mockClient.On("DeleteSipRule", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteSipRule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteVoiceConnector", func(t *testing.T) {
        input := &chimesdkvoice.DeleteVoiceConnectorInput{}
        output := &chimesdkvoice.DeleteVoiceConnectorOutput{}

        mockClient.On("DeleteVoiceConnector", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteVoiceConnector(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteVoiceConnectorEmergencyCallingConfiguration", func(t *testing.T) {
        input := &chimesdkvoice.DeleteVoiceConnectorEmergencyCallingConfigurationInput{}
        output := &chimesdkvoice.DeleteVoiceConnectorEmergencyCallingConfigurationOutput{}

        mockClient.On("DeleteVoiceConnectorEmergencyCallingConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteVoiceConnectorEmergencyCallingConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteVoiceConnectorGroup", func(t *testing.T) {
        input := &chimesdkvoice.DeleteVoiceConnectorGroupInput{}
        output := &chimesdkvoice.DeleteVoiceConnectorGroupOutput{}

        mockClient.On("DeleteVoiceConnectorGroup", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteVoiceConnectorGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteVoiceConnectorOrigination", func(t *testing.T) {
        input := &chimesdkvoice.DeleteVoiceConnectorOriginationInput{}
        output := &chimesdkvoice.DeleteVoiceConnectorOriginationOutput{}

        mockClient.On("DeleteVoiceConnectorOrigination", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteVoiceConnectorOrigination(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteVoiceConnectorProxy", func(t *testing.T) {
        input := &chimesdkvoice.DeleteVoiceConnectorProxyInput{}
        output := &chimesdkvoice.DeleteVoiceConnectorProxyOutput{}

        mockClient.On("DeleteVoiceConnectorProxy", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteVoiceConnectorProxy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteVoiceConnectorStreamingConfiguration", func(t *testing.T) {
        input := &chimesdkvoice.DeleteVoiceConnectorStreamingConfigurationInput{}
        output := &chimesdkvoice.DeleteVoiceConnectorStreamingConfigurationOutput{}

        mockClient.On("DeleteVoiceConnectorStreamingConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteVoiceConnectorStreamingConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteVoiceConnectorTermination", func(t *testing.T) {
        input := &chimesdkvoice.DeleteVoiceConnectorTerminationInput{}
        output := &chimesdkvoice.DeleteVoiceConnectorTerminationOutput{}

        mockClient.On("DeleteVoiceConnectorTermination", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteVoiceConnectorTermination(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteVoiceConnectorTerminationCredentials", func(t *testing.T) {
        input := &chimesdkvoice.DeleteVoiceConnectorTerminationCredentialsInput{}
        output := &chimesdkvoice.DeleteVoiceConnectorTerminationCredentialsOutput{}

        mockClient.On("DeleteVoiceConnectorTerminationCredentials", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteVoiceConnectorTerminationCredentials(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteVoiceProfile", func(t *testing.T) {
        input := &chimesdkvoice.DeleteVoiceProfileInput{}
        output := &chimesdkvoice.DeleteVoiceProfileOutput{}

        mockClient.On("DeleteVoiceProfile", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteVoiceProfile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteVoiceProfileDomain", func(t *testing.T) {
        input := &chimesdkvoice.DeleteVoiceProfileDomainInput{}
        output := &chimesdkvoice.DeleteVoiceProfileDomainOutput{}

        mockClient.On("DeleteVoiceProfileDomain", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteVoiceProfileDomain(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociatePhoneNumbersFromVoiceConnector", func(t *testing.T) {
        input := &chimesdkvoice.DisassociatePhoneNumbersFromVoiceConnectorInput{}
        output := &chimesdkvoice.DisassociatePhoneNumbersFromVoiceConnectorOutput{}

        mockClient.On("DisassociatePhoneNumbersFromVoiceConnector", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociatePhoneNumbersFromVoiceConnector(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociatePhoneNumbersFromVoiceConnectorGroup", func(t *testing.T) {
        input := &chimesdkvoice.DisassociatePhoneNumbersFromVoiceConnectorGroupInput{}
        output := &chimesdkvoice.DisassociatePhoneNumbersFromVoiceConnectorGroupOutput{}

        mockClient.On("DisassociatePhoneNumbersFromVoiceConnectorGroup", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociatePhoneNumbersFromVoiceConnectorGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetGlobalSettings", func(t *testing.T) {
        input := &chimesdkvoice.GetGlobalSettingsInput{}
        output := &chimesdkvoice.GetGlobalSettingsOutput{}

        mockClient.On("GetGlobalSettings", ctx, input).Return(output, nil)

        result, err := mockClient.GetGlobalSettings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetPhoneNumber", func(t *testing.T) {
        input := &chimesdkvoice.GetPhoneNumberInput{}
        output := &chimesdkvoice.GetPhoneNumberOutput{}

        mockClient.On("GetPhoneNumber", ctx, input).Return(output, nil)

        result, err := mockClient.GetPhoneNumber(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetPhoneNumberOrder", func(t *testing.T) {
        input := &chimesdkvoice.GetPhoneNumberOrderInput{}
        output := &chimesdkvoice.GetPhoneNumberOrderOutput{}

        mockClient.On("GetPhoneNumberOrder", ctx, input).Return(output, nil)

        result, err := mockClient.GetPhoneNumberOrder(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetPhoneNumberSettings", func(t *testing.T) {
        input := &chimesdkvoice.GetPhoneNumberSettingsInput{}
        output := &chimesdkvoice.GetPhoneNumberSettingsOutput{}

        mockClient.On("GetPhoneNumberSettings", ctx, input).Return(output, nil)

        result, err := mockClient.GetPhoneNumberSettings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetProxySession", func(t *testing.T) {
        input := &chimesdkvoice.GetProxySessionInput{}
        output := &chimesdkvoice.GetProxySessionOutput{}

        mockClient.On("GetProxySession", ctx, input).Return(output, nil)

        result, err := mockClient.GetProxySession(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSipMediaApplication", func(t *testing.T) {
        input := &chimesdkvoice.GetSipMediaApplicationInput{}
        output := &chimesdkvoice.GetSipMediaApplicationOutput{}

        mockClient.On("GetSipMediaApplication", ctx, input).Return(output, nil)

        result, err := mockClient.GetSipMediaApplication(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSipMediaApplicationAlexaSkillConfiguration", func(t *testing.T) {
        input := &chimesdkvoice.GetSipMediaApplicationAlexaSkillConfigurationInput{}
        output := &chimesdkvoice.GetSipMediaApplicationAlexaSkillConfigurationOutput{}

        mockClient.On("GetSipMediaApplicationAlexaSkillConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.GetSipMediaApplicationAlexaSkillConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSipMediaApplicationLoggingConfiguration", func(t *testing.T) {
        input := &chimesdkvoice.GetSipMediaApplicationLoggingConfigurationInput{}
        output := &chimesdkvoice.GetSipMediaApplicationLoggingConfigurationOutput{}

        mockClient.On("GetSipMediaApplicationLoggingConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.GetSipMediaApplicationLoggingConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSipRule", func(t *testing.T) {
        input := &chimesdkvoice.GetSipRuleInput{}
        output := &chimesdkvoice.GetSipRuleOutput{}

        mockClient.On("GetSipRule", ctx, input).Return(output, nil)

        result, err := mockClient.GetSipRule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSpeakerSearchTask", func(t *testing.T) {
        input := &chimesdkvoice.GetSpeakerSearchTaskInput{}
        output := &chimesdkvoice.GetSpeakerSearchTaskOutput{}

        mockClient.On("GetSpeakerSearchTask", ctx, input).Return(output, nil)

        result, err := mockClient.GetSpeakerSearchTask(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetVoiceConnector", func(t *testing.T) {
        input := &chimesdkvoice.GetVoiceConnectorInput{}
        output := &chimesdkvoice.GetVoiceConnectorOutput{}

        mockClient.On("GetVoiceConnector", ctx, input).Return(output, nil)

        result, err := mockClient.GetVoiceConnector(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetVoiceConnectorEmergencyCallingConfiguration", func(t *testing.T) {
        input := &chimesdkvoice.GetVoiceConnectorEmergencyCallingConfigurationInput{}
        output := &chimesdkvoice.GetVoiceConnectorEmergencyCallingConfigurationOutput{}

        mockClient.On("GetVoiceConnectorEmergencyCallingConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.GetVoiceConnectorEmergencyCallingConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetVoiceConnectorGroup", func(t *testing.T) {
        input := &chimesdkvoice.GetVoiceConnectorGroupInput{}
        output := &chimesdkvoice.GetVoiceConnectorGroupOutput{}

        mockClient.On("GetVoiceConnectorGroup", ctx, input).Return(output, nil)

        result, err := mockClient.GetVoiceConnectorGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetVoiceConnectorLoggingConfiguration", func(t *testing.T) {
        input := &chimesdkvoice.GetVoiceConnectorLoggingConfigurationInput{}
        output := &chimesdkvoice.GetVoiceConnectorLoggingConfigurationOutput{}

        mockClient.On("GetVoiceConnectorLoggingConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.GetVoiceConnectorLoggingConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetVoiceConnectorOrigination", func(t *testing.T) {
        input := &chimesdkvoice.GetVoiceConnectorOriginationInput{}
        output := &chimesdkvoice.GetVoiceConnectorOriginationOutput{}

        mockClient.On("GetVoiceConnectorOrigination", ctx, input).Return(output, nil)

        result, err := mockClient.GetVoiceConnectorOrigination(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetVoiceConnectorProxy", func(t *testing.T) {
        input := &chimesdkvoice.GetVoiceConnectorProxyInput{}
        output := &chimesdkvoice.GetVoiceConnectorProxyOutput{}

        mockClient.On("GetVoiceConnectorProxy", ctx, input).Return(output, nil)

        result, err := mockClient.GetVoiceConnectorProxy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetVoiceConnectorStreamingConfiguration", func(t *testing.T) {
        input := &chimesdkvoice.GetVoiceConnectorStreamingConfigurationInput{}
        output := &chimesdkvoice.GetVoiceConnectorStreamingConfigurationOutput{}

        mockClient.On("GetVoiceConnectorStreamingConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.GetVoiceConnectorStreamingConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetVoiceConnectorTermination", func(t *testing.T) {
        input := &chimesdkvoice.GetVoiceConnectorTerminationInput{}
        output := &chimesdkvoice.GetVoiceConnectorTerminationOutput{}

        mockClient.On("GetVoiceConnectorTermination", ctx, input).Return(output, nil)

        result, err := mockClient.GetVoiceConnectorTermination(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetVoiceConnectorTerminationHealth", func(t *testing.T) {
        input := &chimesdkvoice.GetVoiceConnectorTerminationHealthInput{}
        output := &chimesdkvoice.GetVoiceConnectorTerminationHealthOutput{}

        mockClient.On("GetVoiceConnectorTerminationHealth", ctx, input).Return(output, nil)

        result, err := mockClient.GetVoiceConnectorTerminationHealth(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetVoiceProfile", func(t *testing.T) {
        input := &chimesdkvoice.GetVoiceProfileInput{}
        output := &chimesdkvoice.GetVoiceProfileOutput{}

        mockClient.On("GetVoiceProfile", ctx, input).Return(output, nil)

        result, err := mockClient.GetVoiceProfile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetVoiceProfileDomain", func(t *testing.T) {
        input := &chimesdkvoice.GetVoiceProfileDomainInput{}
        output := &chimesdkvoice.GetVoiceProfileDomainOutput{}

        mockClient.On("GetVoiceProfileDomain", ctx, input).Return(output, nil)

        result, err := mockClient.GetVoiceProfileDomain(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetVoiceToneAnalysisTask", func(t *testing.T) {
        input := &chimesdkvoice.GetVoiceToneAnalysisTaskInput{}
        output := &chimesdkvoice.GetVoiceToneAnalysisTaskOutput{}

        mockClient.On("GetVoiceToneAnalysisTask", ctx, input).Return(output, nil)

        result, err := mockClient.GetVoiceToneAnalysisTask(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAvailableVoiceConnectorRegions", func(t *testing.T) {
        input := &chimesdkvoice.ListAvailableVoiceConnectorRegionsInput{}
        output := &chimesdkvoice.ListAvailableVoiceConnectorRegionsOutput{}

        mockClient.On("ListAvailableVoiceConnectorRegions", ctx, input).Return(output, nil)

        result, err := mockClient.ListAvailableVoiceConnectorRegions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListPhoneNumberOrders", func(t *testing.T) {
        input := &chimesdkvoice.ListPhoneNumberOrdersInput{}
        output := &chimesdkvoice.ListPhoneNumberOrdersOutput{}

        mockClient.On("ListPhoneNumberOrders", ctx, input).Return(output, nil)

        result, err := mockClient.ListPhoneNumberOrders(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListPhoneNumbers", func(t *testing.T) {
        input := &chimesdkvoice.ListPhoneNumbersInput{}
        output := &chimesdkvoice.ListPhoneNumbersOutput{}

        mockClient.On("ListPhoneNumbers", ctx, input).Return(output, nil)

        result, err := mockClient.ListPhoneNumbers(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListProxySessions", func(t *testing.T) {
        input := &chimesdkvoice.ListProxySessionsInput{}
        output := &chimesdkvoice.ListProxySessionsOutput{}

        mockClient.On("ListProxySessions", ctx, input).Return(output, nil)

        result, err := mockClient.ListProxySessions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListSipMediaApplications", func(t *testing.T) {
        input := &chimesdkvoice.ListSipMediaApplicationsInput{}
        output := &chimesdkvoice.ListSipMediaApplicationsOutput{}

        mockClient.On("ListSipMediaApplications", ctx, input).Return(output, nil)

        result, err := mockClient.ListSipMediaApplications(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListSipRules", func(t *testing.T) {
        input := &chimesdkvoice.ListSipRulesInput{}
        output := &chimesdkvoice.ListSipRulesOutput{}

        mockClient.On("ListSipRules", ctx, input).Return(output, nil)

        result, err := mockClient.ListSipRules(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListSupportedPhoneNumberCountries", func(t *testing.T) {
        input := &chimesdkvoice.ListSupportedPhoneNumberCountriesInput{}
        output := &chimesdkvoice.ListSupportedPhoneNumberCountriesOutput{}

        mockClient.On("ListSupportedPhoneNumberCountries", ctx, input).Return(output, nil)

        result, err := mockClient.ListSupportedPhoneNumberCountries(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &chimesdkvoice.ListTagsForResourceInput{}
        output := &chimesdkvoice.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListVoiceConnectorGroups", func(t *testing.T) {
        input := &chimesdkvoice.ListVoiceConnectorGroupsInput{}
        output := &chimesdkvoice.ListVoiceConnectorGroupsOutput{}

        mockClient.On("ListVoiceConnectorGroups", ctx, input).Return(output, nil)

        result, err := mockClient.ListVoiceConnectorGroups(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListVoiceConnectorTerminationCredentials", func(t *testing.T) {
        input := &chimesdkvoice.ListVoiceConnectorTerminationCredentialsInput{}
        output := &chimesdkvoice.ListVoiceConnectorTerminationCredentialsOutput{}

        mockClient.On("ListVoiceConnectorTerminationCredentials", ctx, input).Return(output, nil)

        result, err := mockClient.ListVoiceConnectorTerminationCredentials(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListVoiceConnectors", func(t *testing.T) {
        input := &chimesdkvoice.ListVoiceConnectorsInput{}
        output := &chimesdkvoice.ListVoiceConnectorsOutput{}

        mockClient.On("ListVoiceConnectors", ctx, input).Return(output, nil)

        result, err := mockClient.ListVoiceConnectors(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListVoiceProfileDomains", func(t *testing.T) {
        input := &chimesdkvoice.ListVoiceProfileDomainsInput{}
        output := &chimesdkvoice.ListVoiceProfileDomainsOutput{}

        mockClient.On("ListVoiceProfileDomains", ctx, input).Return(output, nil)

        result, err := mockClient.ListVoiceProfileDomains(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListVoiceProfiles", func(t *testing.T) {
        input := &chimesdkvoice.ListVoiceProfilesInput{}
        output := &chimesdkvoice.ListVoiceProfilesOutput{}

        mockClient.On("ListVoiceProfiles", ctx, input).Return(output, nil)

        result, err := mockClient.ListVoiceProfiles(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutSipMediaApplicationAlexaSkillConfiguration", func(t *testing.T) {
        input := &chimesdkvoice.PutSipMediaApplicationAlexaSkillConfigurationInput{}
        output := &chimesdkvoice.PutSipMediaApplicationAlexaSkillConfigurationOutput{}

        mockClient.On("PutSipMediaApplicationAlexaSkillConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.PutSipMediaApplicationAlexaSkillConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutSipMediaApplicationLoggingConfiguration", func(t *testing.T) {
        input := &chimesdkvoice.PutSipMediaApplicationLoggingConfigurationInput{}
        output := &chimesdkvoice.PutSipMediaApplicationLoggingConfigurationOutput{}

        mockClient.On("PutSipMediaApplicationLoggingConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.PutSipMediaApplicationLoggingConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutVoiceConnectorEmergencyCallingConfiguration", func(t *testing.T) {
        input := &chimesdkvoice.PutVoiceConnectorEmergencyCallingConfigurationInput{}
        output := &chimesdkvoice.PutVoiceConnectorEmergencyCallingConfigurationOutput{}

        mockClient.On("PutVoiceConnectorEmergencyCallingConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.PutVoiceConnectorEmergencyCallingConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutVoiceConnectorLoggingConfiguration", func(t *testing.T) {
        input := &chimesdkvoice.PutVoiceConnectorLoggingConfigurationInput{}
        output := &chimesdkvoice.PutVoiceConnectorLoggingConfigurationOutput{}

        mockClient.On("PutVoiceConnectorLoggingConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.PutVoiceConnectorLoggingConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutVoiceConnectorOrigination", func(t *testing.T) {
        input := &chimesdkvoice.PutVoiceConnectorOriginationInput{}
        output := &chimesdkvoice.PutVoiceConnectorOriginationOutput{}

        mockClient.On("PutVoiceConnectorOrigination", ctx, input).Return(output, nil)

        result, err := mockClient.PutVoiceConnectorOrigination(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutVoiceConnectorProxy", func(t *testing.T) {
        input := &chimesdkvoice.PutVoiceConnectorProxyInput{}
        output := &chimesdkvoice.PutVoiceConnectorProxyOutput{}

        mockClient.On("PutVoiceConnectorProxy", ctx, input).Return(output, nil)

        result, err := mockClient.PutVoiceConnectorProxy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutVoiceConnectorStreamingConfiguration", func(t *testing.T) {
        input := &chimesdkvoice.PutVoiceConnectorStreamingConfigurationInput{}
        output := &chimesdkvoice.PutVoiceConnectorStreamingConfigurationOutput{}

        mockClient.On("PutVoiceConnectorStreamingConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.PutVoiceConnectorStreamingConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutVoiceConnectorTermination", func(t *testing.T) {
        input := &chimesdkvoice.PutVoiceConnectorTerminationInput{}
        output := &chimesdkvoice.PutVoiceConnectorTerminationOutput{}

        mockClient.On("PutVoiceConnectorTermination", ctx, input).Return(output, nil)

        result, err := mockClient.PutVoiceConnectorTermination(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutVoiceConnectorTerminationCredentials", func(t *testing.T) {
        input := &chimesdkvoice.PutVoiceConnectorTerminationCredentialsInput{}
        output := &chimesdkvoice.PutVoiceConnectorTerminationCredentialsOutput{}

        mockClient.On("PutVoiceConnectorTerminationCredentials", ctx, input).Return(output, nil)

        result, err := mockClient.PutVoiceConnectorTerminationCredentials(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRestorePhoneNumber", func(t *testing.T) {
        input := &chimesdkvoice.RestorePhoneNumberInput{}
        output := &chimesdkvoice.RestorePhoneNumberOutput{}

        mockClient.On("RestorePhoneNumber", ctx, input).Return(output, nil)

        result, err := mockClient.RestorePhoneNumber(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSearchAvailablePhoneNumbers", func(t *testing.T) {
        input := &chimesdkvoice.SearchAvailablePhoneNumbersInput{}
        output := &chimesdkvoice.SearchAvailablePhoneNumbersOutput{}

        mockClient.On("SearchAvailablePhoneNumbers", ctx, input).Return(output, nil)

        result, err := mockClient.SearchAvailablePhoneNumbers(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartSpeakerSearchTask", func(t *testing.T) {
        input := &chimesdkvoice.StartSpeakerSearchTaskInput{}
        output := &chimesdkvoice.StartSpeakerSearchTaskOutput{}

        mockClient.On("StartSpeakerSearchTask", ctx, input).Return(output, nil)

        result, err := mockClient.StartSpeakerSearchTask(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartVoiceToneAnalysisTask", func(t *testing.T) {
        input := &chimesdkvoice.StartVoiceToneAnalysisTaskInput{}
        output := &chimesdkvoice.StartVoiceToneAnalysisTaskOutput{}

        mockClient.On("StartVoiceToneAnalysisTask", ctx, input).Return(output, nil)

        result, err := mockClient.StartVoiceToneAnalysisTask(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopSpeakerSearchTask", func(t *testing.T) {
        input := &chimesdkvoice.StopSpeakerSearchTaskInput{}
        output := &chimesdkvoice.StopSpeakerSearchTaskOutput{}

        mockClient.On("StopSpeakerSearchTask", ctx, input).Return(output, nil)

        result, err := mockClient.StopSpeakerSearchTask(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopVoiceToneAnalysisTask", func(t *testing.T) {
        input := &chimesdkvoice.StopVoiceToneAnalysisTaskInput{}
        output := &chimesdkvoice.StopVoiceToneAnalysisTaskOutput{}

        mockClient.On("StopVoiceToneAnalysisTask", ctx, input).Return(output, nil)

        result, err := mockClient.StopVoiceToneAnalysisTask(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &chimesdkvoice.TagResourceInput{}
        output := &chimesdkvoice.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &chimesdkvoice.UntagResourceInput{}
        output := &chimesdkvoice.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateGlobalSettings", func(t *testing.T) {
        input := &chimesdkvoice.UpdateGlobalSettingsInput{}
        output := &chimesdkvoice.UpdateGlobalSettingsOutput{}

        mockClient.On("UpdateGlobalSettings", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateGlobalSettings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdatePhoneNumber", func(t *testing.T) {
        input := &chimesdkvoice.UpdatePhoneNumberInput{}
        output := &chimesdkvoice.UpdatePhoneNumberOutput{}

        mockClient.On("UpdatePhoneNumber", ctx, input).Return(output, nil)

        result, err := mockClient.UpdatePhoneNumber(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdatePhoneNumberSettings", func(t *testing.T) {
        input := &chimesdkvoice.UpdatePhoneNumberSettingsInput{}
        output := &chimesdkvoice.UpdatePhoneNumberSettingsOutput{}

        mockClient.On("UpdatePhoneNumberSettings", ctx, input).Return(output, nil)

        result, err := mockClient.UpdatePhoneNumberSettings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateProxySession", func(t *testing.T) {
        input := &chimesdkvoice.UpdateProxySessionInput{}
        output := &chimesdkvoice.UpdateProxySessionOutput{}

        mockClient.On("UpdateProxySession", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateProxySession(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateSipMediaApplication", func(t *testing.T) {
        input := &chimesdkvoice.UpdateSipMediaApplicationInput{}
        output := &chimesdkvoice.UpdateSipMediaApplicationOutput{}

        mockClient.On("UpdateSipMediaApplication", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateSipMediaApplication(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateSipMediaApplicationCall", func(t *testing.T) {
        input := &chimesdkvoice.UpdateSipMediaApplicationCallInput{}
        output := &chimesdkvoice.UpdateSipMediaApplicationCallOutput{}

        mockClient.On("UpdateSipMediaApplicationCall", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateSipMediaApplicationCall(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateSipRule", func(t *testing.T) {
        input := &chimesdkvoice.UpdateSipRuleInput{}
        output := &chimesdkvoice.UpdateSipRuleOutput{}

        mockClient.On("UpdateSipRule", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateSipRule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateVoiceConnector", func(t *testing.T) {
        input := &chimesdkvoice.UpdateVoiceConnectorInput{}
        output := &chimesdkvoice.UpdateVoiceConnectorOutput{}

        mockClient.On("UpdateVoiceConnector", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateVoiceConnector(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateVoiceConnectorGroup", func(t *testing.T) {
        input := &chimesdkvoice.UpdateVoiceConnectorGroupInput{}
        output := &chimesdkvoice.UpdateVoiceConnectorGroupOutput{}

        mockClient.On("UpdateVoiceConnectorGroup", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateVoiceConnectorGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateVoiceProfile", func(t *testing.T) {
        input := &chimesdkvoice.UpdateVoiceProfileInput{}
        output := &chimesdkvoice.UpdateVoiceProfileOutput{}

        mockClient.On("UpdateVoiceProfile", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateVoiceProfile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateVoiceProfileDomain", func(t *testing.T) {
        input := &chimesdkvoice.UpdateVoiceProfileDomainInput{}
        output := &chimesdkvoice.UpdateVoiceProfileDomainOutput{}

        mockClient.On("UpdateVoiceProfileDomain", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateVoiceProfileDomain(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestValidateE911Address", func(t *testing.T) {
        input := &chimesdkvoice.ValidateE911AddressInput{}
        output := &chimesdkvoice.ValidateE911AddressOutput{}

        mockClient.On("ValidateE911Address", ctx, input).Return(output, nil)

        result, err := mockClient.ValidateE911Address(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
