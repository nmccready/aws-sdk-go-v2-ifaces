
package chimesdkvoice_iface

// This file was generated by scripts/gen_ifaces.sh do not edit it manually
import (
    "context"
    . "github.com/aws/aws-sdk-go-v2/service/chimesdkvoice"
)

// IClient defines the interface for chimesdkvoice
type IClient interface {
 Options() Options 
 AssociatePhoneNumbersWithVoiceConnector(ctx context.Context, params *AssociatePhoneNumbersWithVoiceConnectorInput, optFns ...func(*Options)) (*AssociatePhoneNumbersWithVoiceConnectorOutput, error) 
 AssociatePhoneNumbersWithVoiceConnectorGroup(ctx context.Context, params *AssociatePhoneNumbersWithVoiceConnectorGroupInput, optFns ...func(*Options)) (*AssociatePhoneNumbersWithVoiceConnectorGroupOutput, error) 
 BatchDeletePhoneNumber(ctx context.Context, params *BatchDeletePhoneNumberInput, optFns ...func(*Options)) (*BatchDeletePhoneNumberOutput, error) 
 BatchUpdatePhoneNumber(ctx context.Context, params *BatchUpdatePhoneNumberInput, optFns ...func(*Options)) (*BatchUpdatePhoneNumberOutput, error) 
 CreatePhoneNumberOrder(ctx context.Context, params *CreatePhoneNumberOrderInput, optFns ...func(*Options)) (*CreatePhoneNumberOrderOutput, error) 
 CreateProxySession(ctx context.Context, params *CreateProxySessionInput, optFns ...func(*Options)) (*CreateProxySessionOutput, error) 
 CreateSipMediaApplication(ctx context.Context, params *CreateSipMediaApplicationInput, optFns ...func(*Options)) (*CreateSipMediaApplicationOutput, error) 
 CreateSipMediaApplicationCall(ctx context.Context, params *CreateSipMediaApplicationCallInput, optFns ...func(*Options)) (*CreateSipMediaApplicationCallOutput, error) 
 CreateSipRule(ctx context.Context, params *CreateSipRuleInput, optFns ...func(*Options)) (*CreateSipRuleOutput, error) 
 CreateVoiceConnector(ctx context.Context, params *CreateVoiceConnectorInput, optFns ...func(*Options)) (*CreateVoiceConnectorOutput, error) 
 CreateVoiceConnectorGroup(ctx context.Context, params *CreateVoiceConnectorGroupInput, optFns ...func(*Options)) (*CreateVoiceConnectorGroupOutput, error) 
 CreateVoiceProfile(ctx context.Context, params *CreateVoiceProfileInput, optFns ...func(*Options)) (*CreateVoiceProfileOutput, error) 
 CreateVoiceProfileDomain(ctx context.Context, params *CreateVoiceProfileDomainInput, optFns ...func(*Options)) (*CreateVoiceProfileDomainOutput, error) 
 DeletePhoneNumber(ctx context.Context, params *DeletePhoneNumberInput, optFns ...func(*Options)) (*DeletePhoneNumberOutput, error) 
 DeleteProxySession(ctx context.Context, params *DeleteProxySessionInput, optFns ...func(*Options)) (*DeleteProxySessionOutput, error) 
 DeleteSipMediaApplication(ctx context.Context, params *DeleteSipMediaApplicationInput, optFns ...func(*Options)) (*DeleteSipMediaApplicationOutput, error) 
 DeleteSipRule(ctx context.Context, params *DeleteSipRuleInput, optFns ...func(*Options)) (*DeleteSipRuleOutput, error) 
 DeleteVoiceConnector(ctx context.Context, params *DeleteVoiceConnectorInput, optFns ...func(*Options)) (*DeleteVoiceConnectorOutput, error) 
 DeleteVoiceConnectorEmergencyCallingConfiguration(ctx context.Context, params *DeleteVoiceConnectorEmergencyCallingConfigurationInput, optFns ...func(*Options)) (*DeleteVoiceConnectorEmergencyCallingConfigurationOutput, error) 
 DeleteVoiceConnectorExternalSystemsConfiguration(ctx context.Context, params *DeleteVoiceConnectorExternalSystemsConfigurationInput, optFns ...func(*Options)) (*DeleteVoiceConnectorExternalSystemsConfigurationOutput, error) 
 DeleteVoiceConnectorGroup(ctx context.Context, params *DeleteVoiceConnectorGroupInput, optFns ...func(*Options)) (*DeleteVoiceConnectorGroupOutput, error) 
 DeleteVoiceConnectorOrigination(ctx context.Context, params *DeleteVoiceConnectorOriginationInput, optFns ...func(*Options)) (*DeleteVoiceConnectorOriginationOutput, error) 
 DeleteVoiceConnectorProxy(ctx context.Context, params *DeleteVoiceConnectorProxyInput, optFns ...func(*Options)) (*DeleteVoiceConnectorProxyOutput, error) 
 DeleteVoiceConnectorStreamingConfiguration(ctx context.Context, params *DeleteVoiceConnectorStreamingConfigurationInput, optFns ...func(*Options)) (*DeleteVoiceConnectorStreamingConfigurationOutput, error) 
 DeleteVoiceConnectorTermination(ctx context.Context, params *DeleteVoiceConnectorTerminationInput, optFns ...func(*Options)) (*DeleteVoiceConnectorTerminationOutput, error) 
 DeleteVoiceConnectorTerminationCredentials(ctx context.Context, params *DeleteVoiceConnectorTerminationCredentialsInput, optFns ...func(*Options)) (*DeleteVoiceConnectorTerminationCredentialsOutput, error) 
 DeleteVoiceProfile(ctx context.Context, params *DeleteVoiceProfileInput, optFns ...func(*Options)) (*DeleteVoiceProfileOutput, error) 
 DeleteVoiceProfileDomain(ctx context.Context, params *DeleteVoiceProfileDomainInput, optFns ...func(*Options)) (*DeleteVoiceProfileDomainOutput, error) 
 DisassociatePhoneNumbersFromVoiceConnector(ctx context.Context, params *DisassociatePhoneNumbersFromVoiceConnectorInput, optFns ...func(*Options)) (*DisassociatePhoneNumbersFromVoiceConnectorOutput, error) 
 DisassociatePhoneNumbersFromVoiceConnectorGroup(ctx context.Context, params *DisassociatePhoneNumbersFromVoiceConnectorGroupInput, optFns ...func(*Options)) (*DisassociatePhoneNumbersFromVoiceConnectorGroupOutput, error) 
 GetGlobalSettings(ctx context.Context, params *GetGlobalSettingsInput, optFns ...func(*Options)) (*GetGlobalSettingsOutput, error) 
 GetPhoneNumber(ctx context.Context, params *GetPhoneNumberInput, optFns ...func(*Options)) (*GetPhoneNumberOutput, error) 
 GetPhoneNumberOrder(ctx context.Context, params *GetPhoneNumberOrderInput, optFns ...func(*Options)) (*GetPhoneNumberOrderOutput, error) 
 GetPhoneNumberSettings(ctx context.Context, params *GetPhoneNumberSettingsInput, optFns ...func(*Options)) (*GetPhoneNumberSettingsOutput, error) 
 GetProxySession(ctx context.Context, params *GetProxySessionInput, optFns ...func(*Options)) (*GetProxySessionOutput, error) 
 GetSipMediaApplication(ctx context.Context, params *GetSipMediaApplicationInput, optFns ...func(*Options)) (*GetSipMediaApplicationOutput, error) 
 GetSipMediaApplicationAlexaSkillConfiguration(ctx context.Context, params *GetSipMediaApplicationAlexaSkillConfigurationInput, optFns ...func(*Options)) (*GetSipMediaApplicationAlexaSkillConfigurationOutput, error) 
 GetSipMediaApplicationLoggingConfiguration(ctx context.Context, params *GetSipMediaApplicationLoggingConfigurationInput, optFns ...func(*Options)) (*GetSipMediaApplicationLoggingConfigurationOutput, error) 
 GetSipRule(ctx context.Context, params *GetSipRuleInput, optFns ...func(*Options)) (*GetSipRuleOutput, error) 
 GetSpeakerSearchTask(ctx context.Context, params *GetSpeakerSearchTaskInput, optFns ...func(*Options)) (*GetSpeakerSearchTaskOutput, error) 
 GetVoiceConnector(ctx context.Context, params *GetVoiceConnectorInput, optFns ...func(*Options)) (*GetVoiceConnectorOutput, error) 
 GetVoiceConnectorEmergencyCallingConfiguration(ctx context.Context, params *GetVoiceConnectorEmergencyCallingConfigurationInput, optFns ...func(*Options)) (*GetVoiceConnectorEmergencyCallingConfigurationOutput, error) 
 GetVoiceConnectorExternalSystemsConfiguration(ctx context.Context, params *GetVoiceConnectorExternalSystemsConfigurationInput, optFns ...func(*Options)) (*GetVoiceConnectorExternalSystemsConfigurationOutput, error) 
 GetVoiceConnectorGroup(ctx context.Context, params *GetVoiceConnectorGroupInput, optFns ...func(*Options)) (*GetVoiceConnectorGroupOutput, error) 
 GetVoiceConnectorLoggingConfiguration(ctx context.Context, params *GetVoiceConnectorLoggingConfigurationInput, optFns ...func(*Options)) (*GetVoiceConnectorLoggingConfigurationOutput, error) 
 GetVoiceConnectorOrigination(ctx context.Context, params *GetVoiceConnectorOriginationInput, optFns ...func(*Options)) (*GetVoiceConnectorOriginationOutput, error) 
 GetVoiceConnectorProxy(ctx context.Context, params *GetVoiceConnectorProxyInput, optFns ...func(*Options)) (*GetVoiceConnectorProxyOutput, error) 
 GetVoiceConnectorStreamingConfiguration(ctx context.Context, params *GetVoiceConnectorStreamingConfigurationInput, optFns ...func(*Options)) (*GetVoiceConnectorStreamingConfigurationOutput, error) 
 GetVoiceConnectorTermination(ctx context.Context, params *GetVoiceConnectorTerminationInput, optFns ...func(*Options)) (*GetVoiceConnectorTerminationOutput, error) 
 GetVoiceConnectorTerminationHealth(ctx context.Context, params *GetVoiceConnectorTerminationHealthInput, optFns ...func(*Options)) (*GetVoiceConnectorTerminationHealthOutput, error) 
 GetVoiceProfile(ctx context.Context, params *GetVoiceProfileInput, optFns ...func(*Options)) (*GetVoiceProfileOutput, error) 
 GetVoiceProfileDomain(ctx context.Context, params *GetVoiceProfileDomainInput, optFns ...func(*Options)) (*GetVoiceProfileDomainOutput, error) 
 GetVoiceToneAnalysisTask(ctx context.Context, params *GetVoiceToneAnalysisTaskInput, optFns ...func(*Options)) (*GetVoiceToneAnalysisTaskOutput, error) 
 ListAvailableVoiceConnectorRegions(ctx context.Context, params *ListAvailableVoiceConnectorRegionsInput, optFns ...func(*Options)) (*ListAvailableVoiceConnectorRegionsOutput, error) 
 ListPhoneNumberOrders(ctx context.Context, params *ListPhoneNumberOrdersInput, optFns ...func(*Options)) (*ListPhoneNumberOrdersOutput, error) 
 ListPhoneNumbers(ctx context.Context, params *ListPhoneNumbersInput, optFns ...func(*Options)) (*ListPhoneNumbersOutput, error) 
 ListProxySessions(ctx context.Context, params *ListProxySessionsInput, optFns ...func(*Options)) (*ListProxySessionsOutput, error) 
 ListSipMediaApplications(ctx context.Context, params *ListSipMediaApplicationsInput, optFns ...func(*Options)) (*ListSipMediaApplicationsOutput, error) 
 ListSipRules(ctx context.Context, params *ListSipRulesInput, optFns ...func(*Options)) (*ListSipRulesOutput, error) 
 ListSupportedPhoneNumberCountries(ctx context.Context, params *ListSupportedPhoneNumberCountriesInput, optFns ...func(*Options)) (*ListSupportedPhoneNumberCountriesOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 ListVoiceConnectorGroups(ctx context.Context, params *ListVoiceConnectorGroupsInput, optFns ...func(*Options)) (*ListVoiceConnectorGroupsOutput, error) 
 ListVoiceConnectorTerminationCredentials(ctx context.Context, params *ListVoiceConnectorTerminationCredentialsInput, optFns ...func(*Options)) (*ListVoiceConnectorTerminationCredentialsOutput, error) 
 ListVoiceConnectors(ctx context.Context, params *ListVoiceConnectorsInput, optFns ...func(*Options)) (*ListVoiceConnectorsOutput, error) 
 ListVoiceProfileDomains(ctx context.Context, params *ListVoiceProfileDomainsInput, optFns ...func(*Options)) (*ListVoiceProfileDomainsOutput, error) 
 ListVoiceProfiles(ctx context.Context, params *ListVoiceProfilesInput, optFns ...func(*Options)) (*ListVoiceProfilesOutput, error) 
 PutSipMediaApplicationAlexaSkillConfiguration(ctx context.Context, params *PutSipMediaApplicationAlexaSkillConfigurationInput, optFns ...func(*Options)) (*PutSipMediaApplicationAlexaSkillConfigurationOutput, error) 
 PutSipMediaApplicationLoggingConfiguration(ctx context.Context, params *PutSipMediaApplicationLoggingConfigurationInput, optFns ...func(*Options)) (*PutSipMediaApplicationLoggingConfigurationOutput, error) 
 PutVoiceConnectorEmergencyCallingConfiguration(ctx context.Context, params *PutVoiceConnectorEmergencyCallingConfigurationInput, optFns ...func(*Options)) (*PutVoiceConnectorEmergencyCallingConfigurationOutput, error) 
 PutVoiceConnectorExternalSystemsConfiguration(ctx context.Context, params *PutVoiceConnectorExternalSystemsConfigurationInput, optFns ...func(*Options)) (*PutVoiceConnectorExternalSystemsConfigurationOutput, error) 
 PutVoiceConnectorLoggingConfiguration(ctx context.Context, params *PutVoiceConnectorLoggingConfigurationInput, optFns ...func(*Options)) (*PutVoiceConnectorLoggingConfigurationOutput, error) 
 PutVoiceConnectorOrigination(ctx context.Context, params *PutVoiceConnectorOriginationInput, optFns ...func(*Options)) (*PutVoiceConnectorOriginationOutput, error) 
 PutVoiceConnectorProxy(ctx context.Context, params *PutVoiceConnectorProxyInput, optFns ...func(*Options)) (*PutVoiceConnectorProxyOutput, error) 
 PutVoiceConnectorStreamingConfiguration(ctx context.Context, params *PutVoiceConnectorStreamingConfigurationInput, optFns ...func(*Options)) (*PutVoiceConnectorStreamingConfigurationOutput, error) 
 PutVoiceConnectorTermination(ctx context.Context, params *PutVoiceConnectorTerminationInput, optFns ...func(*Options)) (*PutVoiceConnectorTerminationOutput, error) 
 PutVoiceConnectorTerminationCredentials(ctx context.Context, params *PutVoiceConnectorTerminationCredentialsInput, optFns ...func(*Options)) (*PutVoiceConnectorTerminationCredentialsOutput, error) 
 RestorePhoneNumber(ctx context.Context, params *RestorePhoneNumberInput, optFns ...func(*Options)) (*RestorePhoneNumberOutput, error) 
 SearchAvailablePhoneNumbers(ctx context.Context, params *SearchAvailablePhoneNumbersInput, optFns ...func(*Options)) (*SearchAvailablePhoneNumbersOutput, error) 
 StartSpeakerSearchTask(ctx context.Context, params *StartSpeakerSearchTaskInput, optFns ...func(*Options)) (*StartSpeakerSearchTaskOutput, error) 
 StartVoiceToneAnalysisTask(ctx context.Context, params *StartVoiceToneAnalysisTaskInput, optFns ...func(*Options)) (*StartVoiceToneAnalysisTaskOutput, error) 
 StopSpeakerSearchTask(ctx context.Context, params *StopSpeakerSearchTaskInput, optFns ...func(*Options)) (*StopSpeakerSearchTaskOutput, error) 
 StopVoiceToneAnalysisTask(ctx context.Context, params *StopVoiceToneAnalysisTaskInput, optFns ...func(*Options)) (*StopVoiceToneAnalysisTaskOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateGlobalSettings(ctx context.Context, params *UpdateGlobalSettingsInput, optFns ...func(*Options)) (*UpdateGlobalSettingsOutput, error) 
 UpdatePhoneNumber(ctx context.Context, params *UpdatePhoneNumberInput, optFns ...func(*Options)) (*UpdatePhoneNumberOutput, error) 
 UpdatePhoneNumberSettings(ctx context.Context, params *UpdatePhoneNumberSettingsInput, optFns ...func(*Options)) (*UpdatePhoneNumberSettingsOutput, error) 
 UpdateProxySession(ctx context.Context, params *UpdateProxySessionInput, optFns ...func(*Options)) (*UpdateProxySessionOutput, error) 
 UpdateSipMediaApplication(ctx context.Context, params *UpdateSipMediaApplicationInput, optFns ...func(*Options)) (*UpdateSipMediaApplicationOutput, error) 
 UpdateSipMediaApplicationCall(ctx context.Context, params *UpdateSipMediaApplicationCallInput, optFns ...func(*Options)) (*UpdateSipMediaApplicationCallOutput, error) 
 UpdateSipRule(ctx context.Context, params *UpdateSipRuleInput, optFns ...func(*Options)) (*UpdateSipRuleOutput, error) 
 UpdateVoiceConnector(ctx context.Context, params *UpdateVoiceConnectorInput, optFns ...func(*Options)) (*UpdateVoiceConnectorOutput, error) 
 UpdateVoiceConnectorGroup(ctx context.Context, params *UpdateVoiceConnectorGroupInput, optFns ...func(*Options)) (*UpdateVoiceConnectorGroupOutput, error) 
 UpdateVoiceProfile(ctx context.Context, params *UpdateVoiceProfileInput, optFns ...func(*Options)) (*UpdateVoiceProfileOutput, error) 
 UpdateVoiceProfileDomain(ctx context.Context, params *UpdateVoiceProfileDomainInput, optFns ...func(*Options)) (*UpdateVoiceProfileDomainOutput, error) 
 ValidateE911Address(ctx context.Context, params *ValidateE911AddressInput, optFns ...func(*Options)) (*ValidateE911AddressOutput, error) 
}
