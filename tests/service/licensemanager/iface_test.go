package licensemanager_test

// tests for the licensemanager service interface for this ../../../service/licensemanager/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/licensemanager"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/licensemanager/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/licensemanager/licensemanager_iface"
	"github.com/stretchr/testify/assert"
)

func TestLicensemanagerServiceCanBeMocked(t *testing.T) {
	var iface licensemanager_iface.IClient
	iface = &licensemanager.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := licensemanager.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAcceptGrant", func(t *testing.T) {
        input := &licensemanager.AcceptGrantInput{}
        output := &licensemanager.AcceptGrantOutput{}

        mockClient.On("AcceptGrant", ctx, input).Return(output, nil)

        result, err := mockClient.AcceptGrant(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCheckInLicense", func(t *testing.T) {
        input := &licensemanager.CheckInLicenseInput{}
        output := &licensemanager.CheckInLicenseOutput{}

        mockClient.On("CheckInLicense", ctx, input).Return(output, nil)

        result, err := mockClient.CheckInLicense(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCheckoutBorrowLicense", func(t *testing.T) {
        input := &licensemanager.CheckoutBorrowLicenseInput{}
        output := &licensemanager.CheckoutBorrowLicenseOutput{}

        mockClient.On("CheckoutBorrowLicense", ctx, input).Return(output, nil)

        result, err := mockClient.CheckoutBorrowLicense(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCheckoutLicense", func(t *testing.T) {
        input := &licensemanager.CheckoutLicenseInput{}
        output := &licensemanager.CheckoutLicenseOutput{}

        mockClient.On("CheckoutLicense", ctx, input).Return(output, nil)

        result, err := mockClient.CheckoutLicense(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateGrant", func(t *testing.T) {
        input := &licensemanager.CreateGrantInput{}
        output := &licensemanager.CreateGrantOutput{}

        mockClient.On("CreateGrant", ctx, input).Return(output, nil)

        result, err := mockClient.CreateGrant(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateGrantVersion", func(t *testing.T) {
        input := &licensemanager.CreateGrantVersionInput{}
        output := &licensemanager.CreateGrantVersionOutput{}

        mockClient.On("CreateGrantVersion", ctx, input).Return(output, nil)

        result, err := mockClient.CreateGrantVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateLicense", func(t *testing.T) {
        input := &licensemanager.CreateLicenseInput{}
        output := &licensemanager.CreateLicenseOutput{}

        mockClient.On("CreateLicense", ctx, input).Return(output, nil)

        result, err := mockClient.CreateLicense(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateLicenseConfiguration", func(t *testing.T) {
        input := &licensemanager.CreateLicenseConfigurationInput{}
        output := &licensemanager.CreateLicenseConfigurationOutput{}

        mockClient.On("CreateLicenseConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.CreateLicenseConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateLicenseConversionTaskForResource", func(t *testing.T) {
        input := &licensemanager.CreateLicenseConversionTaskForResourceInput{}
        output := &licensemanager.CreateLicenseConversionTaskForResourceOutput{}

        mockClient.On("CreateLicenseConversionTaskForResource", ctx, input).Return(output, nil)

        result, err := mockClient.CreateLicenseConversionTaskForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateLicenseManagerReportGenerator", func(t *testing.T) {
        input := &licensemanager.CreateLicenseManagerReportGeneratorInput{}
        output := &licensemanager.CreateLicenseManagerReportGeneratorOutput{}

        mockClient.On("CreateLicenseManagerReportGenerator", ctx, input).Return(output, nil)

        result, err := mockClient.CreateLicenseManagerReportGenerator(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateLicenseVersion", func(t *testing.T) {
        input := &licensemanager.CreateLicenseVersionInput{}
        output := &licensemanager.CreateLicenseVersionOutput{}

        mockClient.On("CreateLicenseVersion", ctx, input).Return(output, nil)

        result, err := mockClient.CreateLicenseVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateToken", func(t *testing.T) {
        input := &licensemanager.CreateTokenInput{}
        output := &licensemanager.CreateTokenOutput{}

        mockClient.On("CreateToken", ctx, input).Return(output, nil)

        result, err := mockClient.CreateToken(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteGrant", func(t *testing.T) {
        input := &licensemanager.DeleteGrantInput{}
        output := &licensemanager.DeleteGrantOutput{}

        mockClient.On("DeleteGrant", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteGrant(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteLicense", func(t *testing.T) {
        input := &licensemanager.DeleteLicenseInput{}
        output := &licensemanager.DeleteLicenseOutput{}

        mockClient.On("DeleteLicense", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteLicense(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteLicenseConfiguration", func(t *testing.T) {
        input := &licensemanager.DeleteLicenseConfigurationInput{}
        output := &licensemanager.DeleteLicenseConfigurationOutput{}

        mockClient.On("DeleteLicenseConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteLicenseConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteLicenseManagerReportGenerator", func(t *testing.T) {
        input := &licensemanager.DeleteLicenseManagerReportGeneratorInput{}
        output := &licensemanager.DeleteLicenseManagerReportGeneratorOutput{}

        mockClient.On("DeleteLicenseManagerReportGenerator", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteLicenseManagerReportGenerator(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteToken", func(t *testing.T) {
        input := &licensemanager.DeleteTokenInput{}
        output := &licensemanager.DeleteTokenOutput{}

        mockClient.On("DeleteToken", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteToken(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestExtendLicenseConsumption", func(t *testing.T) {
        input := &licensemanager.ExtendLicenseConsumptionInput{}
        output := &licensemanager.ExtendLicenseConsumptionOutput{}

        mockClient.On("ExtendLicenseConsumption", ctx, input).Return(output, nil)

        result, err := mockClient.ExtendLicenseConsumption(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetAccessToken", func(t *testing.T) {
        input := &licensemanager.GetAccessTokenInput{}
        output := &licensemanager.GetAccessTokenOutput{}

        mockClient.On("GetAccessToken", ctx, input).Return(output, nil)

        result, err := mockClient.GetAccessToken(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetGrant", func(t *testing.T) {
        input := &licensemanager.GetGrantInput{}
        output := &licensemanager.GetGrantOutput{}

        mockClient.On("GetGrant", ctx, input).Return(output, nil)

        result, err := mockClient.GetGrant(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetLicense", func(t *testing.T) {
        input := &licensemanager.GetLicenseInput{}
        output := &licensemanager.GetLicenseOutput{}

        mockClient.On("GetLicense", ctx, input).Return(output, nil)

        result, err := mockClient.GetLicense(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetLicenseConfiguration", func(t *testing.T) {
        input := &licensemanager.GetLicenseConfigurationInput{}
        output := &licensemanager.GetLicenseConfigurationOutput{}

        mockClient.On("GetLicenseConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.GetLicenseConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetLicenseConversionTask", func(t *testing.T) {
        input := &licensemanager.GetLicenseConversionTaskInput{}
        output := &licensemanager.GetLicenseConversionTaskOutput{}

        mockClient.On("GetLicenseConversionTask", ctx, input).Return(output, nil)

        result, err := mockClient.GetLicenseConversionTask(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetLicenseManagerReportGenerator", func(t *testing.T) {
        input := &licensemanager.GetLicenseManagerReportGeneratorInput{}
        output := &licensemanager.GetLicenseManagerReportGeneratorOutput{}

        mockClient.On("GetLicenseManagerReportGenerator", ctx, input).Return(output, nil)

        result, err := mockClient.GetLicenseManagerReportGenerator(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetLicenseUsage", func(t *testing.T) {
        input := &licensemanager.GetLicenseUsageInput{}
        output := &licensemanager.GetLicenseUsageOutput{}

        mockClient.On("GetLicenseUsage", ctx, input).Return(output, nil)

        result, err := mockClient.GetLicenseUsage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetServiceSettings", func(t *testing.T) {
        input := &licensemanager.GetServiceSettingsInput{}
        output := &licensemanager.GetServiceSettingsOutput{}

        mockClient.On("GetServiceSettings", ctx, input).Return(output, nil)

        result, err := mockClient.GetServiceSettings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAssociationsForLicenseConfiguration", func(t *testing.T) {
        input := &licensemanager.ListAssociationsForLicenseConfigurationInput{}
        output := &licensemanager.ListAssociationsForLicenseConfigurationOutput{}

        mockClient.On("ListAssociationsForLicenseConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.ListAssociationsForLicenseConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDistributedGrants", func(t *testing.T) {
        input := &licensemanager.ListDistributedGrantsInput{}
        output := &licensemanager.ListDistributedGrantsOutput{}

        mockClient.On("ListDistributedGrants", ctx, input).Return(output, nil)

        result, err := mockClient.ListDistributedGrants(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListFailuresForLicenseConfigurationOperations", func(t *testing.T) {
        input := &licensemanager.ListFailuresForLicenseConfigurationOperationsInput{}
        output := &licensemanager.ListFailuresForLicenseConfigurationOperationsOutput{}

        mockClient.On("ListFailuresForLicenseConfigurationOperations", ctx, input).Return(output, nil)

        result, err := mockClient.ListFailuresForLicenseConfigurationOperations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListLicenseConfigurations", func(t *testing.T) {
        input := &licensemanager.ListLicenseConfigurationsInput{}
        output := &licensemanager.ListLicenseConfigurationsOutput{}

        mockClient.On("ListLicenseConfigurations", ctx, input).Return(output, nil)

        result, err := mockClient.ListLicenseConfigurations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListLicenseConversionTasks", func(t *testing.T) {
        input := &licensemanager.ListLicenseConversionTasksInput{}
        output := &licensemanager.ListLicenseConversionTasksOutput{}

        mockClient.On("ListLicenseConversionTasks", ctx, input).Return(output, nil)

        result, err := mockClient.ListLicenseConversionTasks(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListLicenseManagerReportGenerators", func(t *testing.T) {
        input := &licensemanager.ListLicenseManagerReportGeneratorsInput{}
        output := &licensemanager.ListLicenseManagerReportGeneratorsOutput{}

        mockClient.On("ListLicenseManagerReportGenerators", ctx, input).Return(output, nil)

        result, err := mockClient.ListLicenseManagerReportGenerators(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListLicenseSpecificationsForResource", func(t *testing.T) {
        input := &licensemanager.ListLicenseSpecificationsForResourceInput{}
        output := &licensemanager.ListLicenseSpecificationsForResourceOutput{}

        mockClient.On("ListLicenseSpecificationsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListLicenseSpecificationsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListLicenseVersions", func(t *testing.T) {
        input := &licensemanager.ListLicenseVersionsInput{}
        output := &licensemanager.ListLicenseVersionsOutput{}

        mockClient.On("ListLicenseVersions", ctx, input).Return(output, nil)

        result, err := mockClient.ListLicenseVersions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListLicenses", func(t *testing.T) {
        input := &licensemanager.ListLicensesInput{}
        output := &licensemanager.ListLicensesOutput{}

        mockClient.On("ListLicenses", ctx, input).Return(output, nil)

        result, err := mockClient.ListLicenses(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListReceivedGrants", func(t *testing.T) {
        input := &licensemanager.ListReceivedGrantsInput{}
        output := &licensemanager.ListReceivedGrantsOutput{}

        mockClient.On("ListReceivedGrants", ctx, input).Return(output, nil)

        result, err := mockClient.ListReceivedGrants(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListReceivedGrantsForOrganization", func(t *testing.T) {
        input := &licensemanager.ListReceivedGrantsForOrganizationInput{}
        output := &licensemanager.ListReceivedGrantsForOrganizationOutput{}

        mockClient.On("ListReceivedGrantsForOrganization", ctx, input).Return(output, nil)

        result, err := mockClient.ListReceivedGrantsForOrganization(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListReceivedLicenses", func(t *testing.T) {
        input := &licensemanager.ListReceivedLicensesInput{}
        output := &licensemanager.ListReceivedLicensesOutput{}

        mockClient.On("ListReceivedLicenses", ctx, input).Return(output, nil)

        result, err := mockClient.ListReceivedLicenses(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListReceivedLicensesForOrganization", func(t *testing.T) {
        input := &licensemanager.ListReceivedLicensesForOrganizationInput{}
        output := &licensemanager.ListReceivedLicensesForOrganizationOutput{}

        mockClient.On("ListReceivedLicensesForOrganization", ctx, input).Return(output, nil)

        result, err := mockClient.ListReceivedLicensesForOrganization(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListResourceInventory", func(t *testing.T) {
        input := &licensemanager.ListResourceInventoryInput{}
        output := &licensemanager.ListResourceInventoryOutput{}

        mockClient.On("ListResourceInventory", ctx, input).Return(output, nil)

        result, err := mockClient.ListResourceInventory(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &licensemanager.ListTagsForResourceInput{}
        output := &licensemanager.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTokens", func(t *testing.T) {
        input := &licensemanager.ListTokensInput{}
        output := &licensemanager.ListTokensOutput{}

        mockClient.On("ListTokens", ctx, input).Return(output, nil)

        result, err := mockClient.ListTokens(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListUsageForLicenseConfiguration", func(t *testing.T) {
        input := &licensemanager.ListUsageForLicenseConfigurationInput{}
        output := &licensemanager.ListUsageForLicenseConfigurationOutput{}

        mockClient.On("ListUsageForLicenseConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.ListUsageForLicenseConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRejectGrant", func(t *testing.T) {
        input := &licensemanager.RejectGrantInput{}
        output := &licensemanager.RejectGrantOutput{}

        mockClient.On("RejectGrant", ctx, input).Return(output, nil)

        result, err := mockClient.RejectGrant(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &licensemanager.TagResourceInput{}
        output := &licensemanager.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &licensemanager.UntagResourceInput{}
        output := &licensemanager.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateLicenseConfiguration", func(t *testing.T) {
        input := &licensemanager.UpdateLicenseConfigurationInput{}
        output := &licensemanager.UpdateLicenseConfigurationOutput{}

        mockClient.On("UpdateLicenseConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateLicenseConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateLicenseManagerReportGenerator", func(t *testing.T) {
        input := &licensemanager.UpdateLicenseManagerReportGeneratorInput{}
        output := &licensemanager.UpdateLicenseManagerReportGeneratorOutput{}

        mockClient.On("UpdateLicenseManagerReportGenerator", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateLicenseManagerReportGenerator(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateLicenseSpecificationsForResource", func(t *testing.T) {
        input := &licensemanager.UpdateLicenseSpecificationsForResourceInput{}
        output := &licensemanager.UpdateLicenseSpecificationsForResourceOutput{}

        mockClient.On("UpdateLicenseSpecificationsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateLicenseSpecificationsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateServiceSettings", func(t *testing.T) {
        input := &licensemanager.UpdateServiceSettingsInput{}
        output := &licensemanager.UpdateServiceSettingsOutput{}

        mockClient.On("UpdateServiceSettings", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateServiceSettings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
