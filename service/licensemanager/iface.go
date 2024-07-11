
package licensemanager

// This file was generated by scripts/gen_ifaces.sh do not edit it manually
import (
    "github.com/aws/aws-sdk-go-v2/service/licensemanager"
)

// IClient defines the interface for licensemanager
type IClient interface {
 Options() Options 
 AcceptGrant(ctx context.Context, params *AcceptGrantInput, optFns ...func(*Options)) (*AcceptGrantOutput, error) 
 CheckInLicense(ctx context.Context, params *CheckInLicenseInput, optFns ...func(*Options)) (*CheckInLicenseOutput, error) 
 CheckoutBorrowLicense(ctx context.Context, params *CheckoutBorrowLicenseInput, optFns ...func(*Options)) (*CheckoutBorrowLicenseOutput, error) 
 CheckoutLicense(ctx context.Context, params *CheckoutLicenseInput, optFns ...func(*Options)) (*CheckoutLicenseOutput, error) 
 CreateGrant(ctx context.Context, params *CreateGrantInput, optFns ...func(*Options)) (*CreateGrantOutput, error) 
 CreateGrantVersion(ctx context.Context, params *CreateGrantVersionInput, optFns ...func(*Options)) (*CreateGrantVersionOutput, error) 
 CreateLicense(ctx context.Context, params *CreateLicenseInput, optFns ...func(*Options)) (*CreateLicenseOutput, error) 
 CreateLicenseConfiguration(ctx context.Context, params *CreateLicenseConfigurationInput, optFns ...func(*Options)) (*CreateLicenseConfigurationOutput, error) 
 CreateLicenseConversionTaskForResource(ctx context.Context, params *CreateLicenseConversionTaskForResourceInput, optFns ...func(*Options)) (*CreateLicenseConversionTaskForResourceOutput, error) 
 CreateLicenseManagerReportGenerator(ctx context.Context, params *CreateLicenseManagerReportGeneratorInput, optFns ...func(*Options)) (*CreateLicenseManagerReportGeneratorOutput, error) 
 CreateLicenseVersion(ctx context.Context, params *CreateLicenseVersionInput, optFns ...func(*Options)) (*CreateLicenseVersionOutput, error) 
 CreateToken(ctx context.Context, params *CreateTokenInput, optFns ...func(*Options)) (*CreateTokenOutput, error) 
 DeleteGrant(ctx context.Context, params *DeleteGrantInput, optFns ...func(*Options)) (*DeleteGrantOutput, error) 
 DeleteLicense(ctx context.Context, params *DeleteLicenseInput, optFns ...func(*Options)) (*DeleteLicenseOutput, error) 
 DeleteLicenseConfiguration(ctx context.Context, params *DeleteLicenseConfigurationInput, optFns ...func(*Options)) (*DeleteLicenseConfigurationOutput, error) 
 DeleteLicenseManagerReportGenerator(ctx context.Context, params *DeleteLicenseManagerReportGeneratorInput, optFns ...func(*Options)) (*DeleteLicenseManagerReportGeneratorOutput, error) 
 DeleteToken(ctx context.Context, params *DeleteTokenInput, optFns ...func(*Options)) (*DeleteTokenOutput, error) 
 ExtendLicenseConsumption(ctx context.Context, params *ExtendLicenseConsumptionInput, optFns ...func(*Options)) (*ExtendLicenseConsumptionOutput, error) 
 GetAccessToken(ctx context.Context, params *GetAccessTokenInput, optFns ...func(*Options)) (*GetAccessTokenOutput, error) 
 GetGrant(ctx context.Context, params *GetGrantInput, optFns ...func(*Options)) (*GetGrantOutput, error) 
 GetLicense(ctx context.Context, params *GetLicenseInput, optFns ...func(*Options)) (*GetLicenseOutput, error) 
 GetLicenseConfiguration(ctx context.Context, params *GetLicenseConfigurationInput, optFns ...func(*Options)) (*GetLicenseConfigurationOutput, error) 
 GetLicenseConversionTask(ctx context.Context, params *GetLicenseConversionTaskInput, optFns ...func(*Options)) (*GetLicenseConversionTaskOutput, error) 
 GetLicenseManagerReportGenerator(ctx context.Context, params *GetLicenseManagerReportGeneratorInput, optFns ...func(*Options)) (*GetLicenseManagerReportGeneratorOutput, error) 
 GetLicenseUsage(ctx context.Context, params *GetLicenseUsageInput, optFns ...func(*Options)) (*GetLicenseUsageOutput, error) 
 GetServiceSettings(ctx context.Context, params *GetServiceSettingsInput, optFns ...func(*Options)) (*GetServiceSettingsOutput, error) 
 ListAssociationsForLicenseConfiguration(ctx context.Context, params *ListAssociationsForLicenseConfigurationInput, optFns ...func(*Options)) (*ListAssociationsForLicenseConfigurationOutput, error) 
 ListDistributedGrants(ctx context.Context, params *ListDistributedGrantsInput, optFns ...func(*Options)) (*ListDistributedGrantsOutput, error) 
 ListFailuresForLicenseConfigurationOperations(ctx context.Context, params *ListFailuresForLicenseConfigurationOperationsInput, optFns ...func(*Options)) (*ListFailuresForLicenseConfigurationOperationsOutput, error) 
 ListLicenseConfigurations(ctx context.Context, params *ListLicenseConfigurationsInput, optFns ...func(*Options)) (*ListLicenseConfigurationsOutput, error) 
 ListLicenseConversionTasks(ctx context.Context, params *ListLicenseConversionTasksInput, optFns ...func(*Options)) (*ListLicenseConversionTasksOutput, error) 
 ListLicenseManagerReportGenerators(ctx context.Context, params *ListLicenseManagerReportGeneratorsInput, optFns ...func(*Options)) (*ListLicenseManagerReportGeneratorsOutput, error) 
 ListLicenseSpecificationsForResource(ctx context.Context, params *ListLicenseSpecificationsForResourceInput, optFns ...func(*Options)) (*ListLicenseSpecificationsForResourceOutput, error) 
 ListLicenseVersions(ctx context.Context, params *ListLicenseVersionsInput, optFns ...func(*Options)) (*ListLicenseVersionsOutput, error) 
 ListLicenses(ctx context.Context, params *ListLicensesInput, optFns ...func(*Options)) (*ListLicensesOutput, error) 
 ListReceivedGrants(ctx context.Context, params *ListReceivedGrantsInput, optFns ...func(*Options)) (*ListReceivedGrantsOutput, error) 
 ListReceivedGrantsForOrganization(ctx context.Context, params *ListReceivedGrantsForOrganizationInput, optFns ...func(*Options)) (*ListReceivedGrantsForOrganizationOutput, error) 
 ListReceivedLicenses(ctx context.Context, params *ListReceivedLicensesInput, optFns ...func(*Options)) (*ListReceivedLicensesOutput, error) 
 ListReceivedLicensesForOrganization(ctx context.Context, params *ListReceivedLicensesForOrganizationInput, optFns ...func(*Options)) (*ListReceivedLicensesForOrganizationOutput, error) 
 ListResourceInventory(ctx context.Context, params *ListResourceInventoryInput, optFns ...func(*Options)) (*ListResourceInventoryOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 ListTokens(ctx context.Context, params *ListTokensInput, optFns ...func(*Options)) (*ListTokensOutput, error) 
 ListUsageForLicenseConfiguration(ctx context.Context, params *ListUsageForLicenseConfigurationInput, optFns ...func(*Options)) (*ListUsageForLicenseConfigurationOutput, error) 
 RejectGrant(ctx context.Context, params *RejectGrantInput, optFns ...func(*Options)) (*RejectGrantOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateLicenseConfiguration(ctx context.Context, params *UpdateLicenseConfigurationInput, optFns ...func(*Options)) (*UpdateLicenseConfigurationOutput, error) 
 UpdateLicenseManagerReportGenerator(ctx context.Context, params *UpdateLicenseManagerReportGeneratorInput, optFns ...func(*Options)) (*UpdateLicenseManagerReportGeneratorOutput, error) 
 UpdateLicenseSpecificationsForResource(ctx context.Context, params *UpdateLicenseSpecificationsForResourceInput, optFns ...func(*Options)) (*UpdateLicenseSpecificationsForResourceOutput, error) 
 UpdateServiceSettings(ctx context.Context, params *UpdateServiceSettingsInput, optFns ...func(*Options)) (*UpdateServiceSettingsOutput, error) 
}
