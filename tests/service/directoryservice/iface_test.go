package directoryservice_test

// tests for the directoryservice service interface for this ../../../service/directoryservice/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/directoryservice"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/directoryservice/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/directoryservice/directoryservice_iface"
	"github.com/stretchr/testify/assert"
)

func TestDirectoryserviceServiceCanBeMocked(t *testing.T) {
	var iface directoryservice_iface.IClient
	iface = &directoryservice.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := directoryservice.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAcceptSharedDirectory", func(t *testing.T) {
        input := &directoryservice.AcceptSharedDirectoryInput{}
        output := &directoryservice.AcceptSharedDirectoryOutput{}

        mockClient.On("AcceptSharedDirectory", ctx, input).Return(output, nil)

        result, err := mockClient.AcceptSharedDirectory(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAddIpRoutes", func(t *testing.T) {
        input := &directoryservice.AddIpRoutesInput{}
        output := &directoryservice.AddIpRoutesOutput{}

        mockClient.On("AddIpRoutes", ctx, input).Return(output, nil)

        result, err := mockClient.AddIpRoutes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAddRegion", func(t *testing.T) {
        input := &directoryservice.AddRegionInput{}
        output := &directoryservice.AddRegionOutput{}

        mockClient.On("AddRegion", ctx, input).Return(output, nil)

        result, err := mockClient.AddRegion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAddTagsToResource", func(t *testing.T) {
        input := &directoryservice.AddTagsToResourceInput{}
        output := &directoryservice.AddTagsToResourceOutput{}

        mockClient.On("AddTagsToResource", ctx, input).Return(output, nil)

        result, err := mockClient.AddTagsToResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCancelSchemaExtension", func(t *testing.T) {
        input := &directoryservice.CancelSchemaExtensionInput{}
        output := &directoryservice.CancelSchemaExtensionOutput{}

        mockClient.On("CancelSchemaExtension", ctx, input).Return(output, nil)

        result, err := mockClient.CancelSchemaExtension(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestConnectDirectory", func(t *testing.T) {
        input := &directoryservice.ConnectDirectoryInput{}
        output := &directoryservice.ConnectDirectoryOutput{}

        mockClient.On("ConnectDirectory", ctx, input).Return(output, nil)

        result, err := mockClient.ConnectDirectory(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateAlias", func(t *testing.T) {
        input := &directoryservice.CreateAliasInput{}
        output := &directoryservice.CreateAliasOutput{}

        mockClient.On("CreateAlias", ctx, input).Return(output, nil)

        result, err := mockClient.CreateAlias(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateComputer", func(t *testing.T) {
        input := &directoryservice.CreateComputerInput{}
        output := &directoryservice.CreateComputerOutput{}

        mockClient.On("CreateComputer", ctx, input).Return(output, nil)

        result, err := mockClient.CreateComputer(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateConditionalForwarder", func(t *testing.T) {
        input := &directoryservice.CreateConditionalForwarderInput{}
        output := &directoryservice.CreateConditionalForwarderOutput{}

        mockClient.On("CreateConditionalForwarder", ctx, input).Return(output, nil)

        result, err := mockClient.CreateConditionalForwarder(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateDirectory", func(t *testing.T) {
        input := &directoryservice.CreateDirectoryInput{}
        output := &directoryservice.CreateDirectoryOutput{}

        mockClient.On("CreateDirectory", ctx, input).Return(output, nil)

        result, err := mockClient.CreateDirectory(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateLogSubscription", func(t *testing.T) {
        input := &directoryservice.CreateLogSubscriptionInput{}
        output := &directoryservice.CreateLogSubscriptionOutput{}

        mockClient.On("CreateLogSubscription", ctx, input).Return(output, nil)

        result, err := mockClient.CreateLogSubscription(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateMicrosoftAD", func(t *testing.T) {
        input := &directoryservice.CreateMicrosoftADInput{}
        output := &directoryservice.CreateMicrosoftADOutput{}

        mockClient.On("CreateMicrosoftAD", ctx, input).Return(output, nil)

        result, err := mockClient.CreateMicrosoftAD(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateSnapshot", func(t *testing.T) {
        input := &directoryservice.CreateSnapshotInput{}
        output := &directoryservice.CreateSnapshotOutput{}

        mockClient.On("CreateSnapshot", ctx, input).Return(output, nil)

        result, err := mockClient.CreateSnapshot(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateTrust", func(t *testing.T) {
        input := &directoryservice.CreateTrustInput{}
        output := &directoryservice.CreateTrustOutput{}

        mockClient.On("CreateTrust", ctx, input).Return(output, nil)

        result, err := mockClient.CreateTrust(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteConditionalForwarder", func(t *testing.T) {
        input := &directoryservice.DeleteConditionalForwarderInput{}
        output := &directoryservice.DeleteConditionalForwarderOutput{}

        mockClient.On("DeleteConditionalForwarder", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteConditionalForwarder(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteDirectory", func(t *testing.T) {
        input := &directoryservice.DeleteDirectoryInput{}
        output := &directoryservice.DeleteDirectoryOutput{}

        mockClient.On("DeleteDirectory", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteDirectory(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteLogSubscription", func(t *testing.T) {
        input := &directoryservice.DeleteLogSubscriptionInput{}
        output := &directoryservice.DeleteLogSubscriptionOutput{}

        mockClient.On("DeleteLogSubscription", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteLogSubscription(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteSnapshot", func(t *testing.T) {
        input := &directoryservice.DeleteSnapshotInput{}
        output := &directoryservice.DeleteSnapshotOutput{}

        mockClient.On("DeleteSnapshot", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteSnapshot(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteTrust", func(t *testing.T) {
        input := &directoryservice.DeleteTrustInput{}
        output := &directoryservice.DeleteTrustOutput{}

        mockClient.On("DeleteTrust", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteTrust(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeregisterCertificate", func(t *testing.T) {
        input := &directoryservice.DeregisterCertificateInput{}
        output := &directoryservice.DeregisterCertificateOutput{}

        mockClient.On("DeregisterCertificate", ctx, input).Return(output, nil)

        result, err := mockClient.DeregisterCertificate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeregisterEventTopic", func(t *testing.T) {
        input := &directoryservice.DeregisterEventTopicInput{}
        output := &directoryservice.DeregisterEventTopicOutput{}

        mockClient.On("DeregisterEventTopic", ctx, input).Return(output, nil)

        result, err := mockClient.DeregisterEventTopic(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeCertificate", func(t *testing.T) {
        input := &directoryservice.DescribeCertificateInput{}
        output := &directoryservice.DescribeCertificateOutput{}

        mockClient.On("DescribeCertificate", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeCertificate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeClientAuthenticationSettings", func(t *testing.T) {
        input := &directoryservice.DescribeClientAuthenticationSettingsInput{}
        output := &directoryservice.DescribeClientAuthenticationSettingsOutput{}

        mockClient.On("DescribeClientAuthenticationSettings", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeClientAuthenticationSettings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeConditionalForwarders", func(t *testing.T) {
        input := &directoryservice.DescribeConditionalForwardersInput{}
        output := &directoryservice.DescribeConditionalForwardersOutput{}

        mockClient.On("DescribeConditionalForwarders", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeConditionalForwarders(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeDirectories", func(t *testing.T) {
        input := &directoryservice.DescribeDirectoriesInput{}
        output := &directoryservice.DescribeDirectoriesOutput{}

        mockClient.On("DescribeDirectories", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeDirectories(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeDomainControllers", func(t *testing.T) {
        input := &directoryservice.DescribeDomainControllersInput{}
        output := &directoryservice.DescribeDomainControllersOutput{}

        mockClient.On("DescribeDomainControllers", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeDomainControllers(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeEventTopics", func(t *testing.T) {
        input := &directoryservice.DescribeEventTopicsInput{}
        output := &directoryservice.DescribeEventTopicsOutput{}

        mockClient.On("DescribeEventTopics", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeEventTopics(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeLDAPSSettings", func(t *testing.T) {
        input := &directoryservice.DescribeLDAPSSettingsInput{}
        output := &directoryservice.DescribeLDAPSSettingsOutput{}

        mockClient.On("DescribeLDAPSSettings", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeLDAPSSettings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeRegions", func(t *testing.T) {
        input := &directoryservice.DescribeRegionsInput{}
        output := &directoryservice.DescribeRegionsOutput{}

        mockClient.On("DescribeRegions", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeRegions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeSettings", func(t *testing.T) {
        input := &directoryservice.DescribeSettingsInput{}
        output := &directoryservice.DescribeSettingsOutput{}

        mockClient.On("DescribeSettings", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeSettings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeSharedDirectories", func(t *testing.T) {
        input := &directoryservice.DescribeSharedDirectoriesInput{}
        output := &directoryservice.DescribeSharedDirectoriesOutput{}

        mockClient.On("DescribeSharedDirectories", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeSharedDirectories(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeSnapshots", func(t *testing.T) {
        input := &directoryservice.DescribeSnapshotsInput{}
        output := &directoryservice.DescribeSnapshotsOutput{}

        mockClient.On("DescribeSnapshots", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeSnapshots(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeTrusts", func(t *testing.T) {
        input := &directoryservice.DescribeTrustsInput{}
        output := &directoryservice.DescribeTrustsOutput{}

        mockClient.On("DescribeTrusts", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeTrusts(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeUpdateDirectory", func(t *testing.T) {
        input := &directoryservice.DescribeUpdateDirectoryInput{}
        output := &directoryservice.DescribeUpdateDirectoryOutput{}

        mockClient.On("DescribeUpdateDirectory", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeUpdateDirectory(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisableClientAuthentication", func(t *testing.T) {
        input := &directoryservice.DisableClientAuthenticationInput{}
        output := &directoryservice.DisableClientAuthenticationOutput{}

        mockClient.On("DisableClientAuthentication", ctx, input).Return(output, nil)

        result, err := mockClient.DisableClientAuthentication(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisableLDAPS", func(t *testing.T) {
        input := &directoryservice.DisableLDAPSInput{}
        output := &directoryservice.DisableLDAPSOutput{}

        mockClient.On("DisableLDAPS", ctx, input).Return(output, nil)

        result, err := mockClient.DisableLDAPS(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisableRadius", func(t *testing.T) {
        input := &directoryservice.DisableRadiusInput{}
        output := &directoryservice.DisableRadiusOutput{}

        mockClient.On("DisableRadius", ctx, input).Return(output, nil)

        result, err := mockClient.DisableRadius(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisableSso", func(t *testing.T) {
        input := &directoryservice.DisableSsoInput{}
        output := &directoryservice.DisableSsoOutput{}

        mockClient.On("DisableSso", ctx, input).Return(output, nil)

        result, err := mockClient.DisableSso(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestEnableClientAuthentication", func(t *testing.T) {
        input := &directoryservice.EnableClientAuthenticationInput{}
        output := &directoryservice.EnableClientAuthenticationOutput{}

        mockClient.On("EnableClientAuthentication", ctx, input).Return(output, nil)

        result, err := mockClient.EnableClientAuthentication(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestEnableLDAPS", func(t *testing.T) {
        input := &directoryservice.EnableLDAPSInput{}
        output := &directoryservice.EnableLDAPSOutput{}

        mockClient.On("EnableLDAPS", ctx, input).Return(output, nil)

        result, err := mockClient.EnableLDAPS(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestEnableRadius", func(t *testing.T) {
        input := &directoryservice.EnableRadiusInput{}
        output := &directoryservice.EnableRadiusOutput{}

        mockClient.On("EnableRadius", ctx, input).Return(output, nil)

        result, err := mockClient.EnableRadius(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestEnableSso", func(t *testing.T) {
        input := &directoryservice.EnableSsoInput{}
        output := &directoryservice.EnableSsoOutput{}

        mockClient.On("EnableSso", ctx, input).Return(output, nil)

        result, err := mockClient.EnableSso(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDirectoryLimits", func(t *testing.T) {
        input := &directoryservice.GetDirectoryLimitsInput{}
        output := &directoryservice.GetDirectoryLimitsOutput{}

        mockClient.On("GetDirectoryLimits", ctx, input).Return(output, nil)

        result, err := mockClient.GetDirectoryLimits(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSnapshotLimits", func(t *testing.T) {
        input := &directoryservice.GetSnapshotLimitsInput{}
        output := &directoryservice.GetSnapshotLimitsOutput{}

        mockClient.On("GetSnapshotLimits", ctx, input).Return(output, nil)

        result, err := mockClient.GetSnapshotLimits(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListCertificates", func(t *testing.T) {
        input := &directoryservice.ListCertificatesInput{}
        output := &directoryservice.ListCertificatesOutput{}

        mockClient.On("ListCertificates", ctx, input).Return(output, nil)

        result, err := mockClient.ListCertificates(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListIpRoutes", func(t *testing.T) {
        input := &directoryservice.ListIpRoutesInput{}
        output := &directoryservice.ListIpRoutesOutput{}

        mockClient.On("ListIpRoutes", ctx, input).Return(output, nil)

        result, err := mockClient.ListIpRoutes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListLogSubscriptions", func(t *testing.T) {
        input := &directoryservice.ListLogSubscriptionsInput{}
        output := &directoryservice.ListLogSubscriptionsOutput{}

        mockClient.On("ListLogSubscriptions", ctx, input).Return(output, nil)

        result, err := mockClient.ListLogSubscriptions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListSchemaExtensions", func(t *testing.T) {
        input := &directoryservice.ListSchemaExtensionsInput{}
        output := &directoryservice.ListSchemaExtensionsOutput{}

        mockClient.On("ListSchemaExtensions", ctx, input).Return(output, nil)

        result, err := mockClient.ListSchemaExtensions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &directoryservice.ListTagsForResourceInput{}
        output := &directoryservice.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRegisterCertificate", func(t *testing.T) {
        input := &directoryservice.RegisterCertificateInput{}
        output := &directoryservice.RegisterCertificateOutput{}

        mockClient.On("RegisterCertificate", ctx, input).Return(output, nil)

        result, err := mockClient.RegisterCertificate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRegisterEventTopic", func(t *testing.T) {
        input := &directoryservice.RegisterEventTopicInput{}
        output := &directoryservice.RegisterEventTopicOutput{}

        mockClient.On("RegisterEventTopic", ctx, input).Return(output, nil)

        result, err := mockClient.RegisterEventTopic(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRejectSharedDirectory", func(t *testing.T) {
        input := &directoryservice.RejectSharedDirectoryInput{}
        output := &directoryservice.RejectSharedDirectoryOutput{}

        mockClient.On("RejectSharedDirectory", ctx, input).Return(output, nil)

        result, err := mockClient.RejectSharedDirectory(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRemoveIpRoutes", func(t *testing.T) {
        input := &directoryservice.RemoveIpRoutesInput{}
        output := &directoryservice.RemoveIpRoutesOutput{}

        mockClient.On("RemoveIpRoutes", ctx, input).Return(output, nil)

        result, err := mockClient.RemoveIpRoutes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRemoveRegion", func(t *testing.T) {
        input := &directoryservice.RemoveRegionInput{}
        output := &directoryservice.RemoveRegionOutput{}

        mockClient.On("RemoveRegion", ctx, input).Return(output, nil)

        result, err := mockClient.RemoveRegion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRemoveTagsFromResource", func(t *testing.T) {
        input := &directoryservice.RemoveTagsFromResourceInput{}
        output := &directoryservice.RemoveTagsFromResourceOutput{}

        mockClient.On("RemoveTagsFromResource", ctx, input).Return(output, nil)

        result, err := mockClient.RemoveTagsFromResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestResetUserPassword", func(t *testing.T) {
        input := &directoryservice.ResetUserPasswordInput{}
        output := &directoryservice.ResetUserPasswordOutput{}

        mockClient.On("ResetUserPassword", ctx, input).Return(output, nil)

        result, err := mockClient.ResetUserPassword(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRestoreFromSnapshot", func(t *testing.T) {
        input := &directoryservice.RestoreFromSnapshotInput{}
        output := &directoryservice.RestoreFromSnapshotOutput{}

        mockClient.On("RestoreFromSnapshot", ctx, input).Return(output, nil)

        result, err := mockClient.RestoreFromSnapshot(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestShareDirectory", func(t *testing.T) {
        input := &directoryservice.ShareDirectoryInput{}
        output := &directoryservice.ShareDirectoryOutput{}

        mockClient.On("ShareDirectory", ctx, input).Return(output, nil)

        result, err := mockClient.ShareDirectory(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartSchemaExtension", func(t *testing.T) {
        input := &directoryservice.StartSchemaExtensionInput{}
        output := &directoryservice.StartSchemaExtensionOutput{}

        mockClient.On("StartSchemaExtension", ctx, input).Return(output, nil)

        result, err := mockClient.StartSchemaExtension(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUnshareDirectory", func(t *testing.T) {
        input := &directoryservice.UnshareDirectoryInput{}
        output := &directoryservice.UnshareDirectoryOutput{}

        mockClient.On("UnshareDirectory", ctx, input).Return(output, nil)

        result, err := mockClient.UnshareDirectory(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateConditionalForwarder", func(t *testing.T) {
        input := &directoryservice.UpdateConditionalForwarderInput{}
        output := &directoryservice.UpdateConditionalForwarderOutput{}

        mockClient.On("UpdateConditionalForwarder", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateConditionalForwarder(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateDirectorySetup", func(t *testing.T) {
        input := &directoryservice.UpdateDirectorySetupInput{}
        output := &directoryservice.UpdateDirectorySetupOutput{}

        mockClient.On("UpdateDirectorySetup", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateDirectorySetup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateNumberOfDomainControllers", func(t *testing.T) {
        input := &directoryservice.UpdateNumberOfDomainControllersInput{}
        output := &directoryservice.UpdateNumberOfDomainControllersOutput{}

        mockClient.On("UpdateNumberOfDomainControllers", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateNumberOfDomainControllers(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateRadius", func(t *testing.T) {
        input := &directoryservice.UpdateRadiusInput{}
        output := &directoryservice.UpdateRadiusOutput{}

        mockClient.On("UpdateRadius", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateRadius(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateSettings", func(t *testing.T) {
        input := &directoryservice.UpdateSettingsInput{}
        output := &directoryservice.UpdateSettingsOutput{}

        mockClient.On("UpdateSettings", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateSettings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateTrust", func(t *testing.T) {
        input := &directoryservice.UpdateTrustInput{}
        output := &directoryservice.UpdateTrustOutput{}

        mockClient.On("UpdateTrust", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateTrust(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestVerifyTrust", func(t *testing.T) {
        input := &directoryservice.VerifyTrustInput{}
        output := &directoryservice.VerifyTrustOutput{}

        mockClient.On("VerifyTrust", ctx, input).Return(output, nil)

        result, err := mockClient.VerifyTrust(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
