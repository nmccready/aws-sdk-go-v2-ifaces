// Code generated by ./scripts/gen_tests.sh DO NOT EDIT.
package appstream_test

// tests for the appstream service interface for 
// this ../../../service/appstream/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/appstream"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/appstream/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/appstream/appstream_iface"
	"github.com/stretchr/testify/assert"
)

func TestAppstreamServiceCanBeMocked(t *testing.T) {
	var iface appstream_iface.IClient
	iface = &appstream.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := appstream.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateAppBlockBuilderAppBlock", func(t *testing.T) {
        input := &appstream.AssociateAppBlockBuilderAppBlockInput{}
        output := &appstream.AssociateAppBlockBuilderAppBlockOutput{}

        mockClient.On("AssociateAppBlockBuilderAppBlock", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateAppBlockBuilderAppBlock(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateApplicationFleet", func(t *testing.T) {
        input := &appstream.AssociateApplicationFleetInput{}
        output := &appstream.AssociateApplicationFleetOutput{}

        mockClient.On("AssociateApplicationFleet", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateApplicationFleet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateApplicationToEntitlement", func(t *testing.T) {
        input := &appstream.AssociateApplicationToEntitlementInput{}
        output := &appstream.AssociateApplicationToEntitlementOutput{}

        mockClient.On("AssociateApplicationToEntitlement", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateApplicationToEntitlement(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateFleet", func(t *testing.T) {
        input := &appstream.AssociateFleetInput{}
        output := &appstream.AssociateFleetOutput{}

        mockClient.On("AssociateFleet", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateFleet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchAssociateUserStack", func(t *testing.T) {
        input := &appstream.BatchAssociateUserStackInput{}
        output := &appstream.BatchAssociateUserStackOutput{}

        mockClient.On("BatchAssociateUserStack", ctx, input).Return(output, nil)

        result, err := mockClient.BatchAssociateUserStack(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchDisassociateUserStack", func(t *testing.T) {
        input := &appstream.BatchDisassociateUserStackInput{}
        output := &appstream.BatchDisassociateUserStackOutput{}

        mockClient.On("BatchDisassociateUserStack", ctx, input).Return(output, nil)

        result, err := mockClient.BatchDisassociateUserStack(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCopyImage", func(t *testing.T) {
        input := &appstream.CopyImageInput{}
        output := &appstream.CopyImageOutput{}

        mockClient.On("CopyImage", ctx, input).Return(output, nil)

        result, err := mockClient.CopyImage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateAppBlock", func(t *testing.T) {
        input := &appstream.CreateAppBlockInput{}
        output := &appstream.CreateAppBlockOutput{}

        mockClient.On("CreateAppBlock", ctx, input).Return(output, nil)

        result, err := mockClient.CreateAppBlock(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateAppBlockBuilder", func(t *testing.T) {
        input := &appstream.CreateAppBlockBuilderInput{}
        output := &appstream.CreateAppBlockBuilderOutput{}

        mockClient.On("CreateAppBlockBuilder", ctx, input).Return(output, nil)

        result, err := mockClient.CreateAppBlockBuilder(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateAppBlockBuilderStreamingURL", func(t *testing.T) {
        input := &appstream.CreateAppBlockBuilderStreamingURLInput{}
        output := &appstream.CreateAppBlockBuilderStreamingURLOutput{}

        mockClient.On("CreateAppBlockBuilderStreamingURL", ctx, input).Return(output, nil)

        result, err := mockClient.CreateAppBlockBuilderStreamingURL(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateApplication", func(t *testing.T) {
        input := &appstream.CreateApplicationInput{}
        output := &appstream.CreateApplicationOutput{}

        mockClient.On("CreateApplication", ctx, input).Return(output, nil)

        result, err := mockClient.CreateApplication(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateDirectoryConfig", func(t *testing.T) {
        input := &appstream.CreateDirectoryConfigInput{}
        output := &appstream.CreateDirectoryConfigOutput{}

        mockClient.On("CreateDirectoryConfig", ctx, input).Return(output, nil)

        result, err := mockClient.CreateDirectoryConfig(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateEntitlement", func(t *testing.T) {
        input := &appstream.CreateEntitlementInput{}
        output := &appstream.CreateEntitlementOutput{}

        mockClient.On("CreateEntitlement", ctx, input).Return(output, nil)

        result, err := mockClient.CreateEntitlement(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateFleet", func(t *testing.T) {
        input := &appstream.CreateFleetInput{}
        output := &appstream.CreateFleetOutput{}

        mockClient.On("CreateFleet", ctx, input).Return(output, nil)

        result, err := mockClient.CreateFleet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateImageBuilder", func(t *testing.T) {
        input := &appstream.CreateImageBuilderInput{}
        output := &appstream.CreateImageBuilderOutput{}

        mockClient.On("CreateImageBuilder", ctx, input).Return(output, nil)

        result, err := mockClient.CreateImageBuilder(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateImageBuilderStreamingURL", func(t *testing.T) {
        input := &appstream.CreateImageBuilderStreamingURLInput{}
        output := &appstream.CreateImageBuilderStreamingURLOutput{}

        mockClient.On("CreateImageBuilderStreamingURL", ctx, input).Return(output, nil)

        result, err := mockClient.CreateImageBuilderStreamingURL(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateStack", func(t *testing.T) {
        input := &appstream.CreateStackInput{}
        output := &appstream.CreateStackOutput{}

        mockClient.On("CreateStack", ctx, input).Return(output, nil)

        result, err := mockClient.CreateStack(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateStreamingURL", func(t *testing.T) {
        input := &appstream.CreateStreamingURLInput{}
        output := &appstream.CreateStreamingURLOutput{}

        mockClient.On("CreateStreamingURL", ctx, input).Return(output, nil)

        result, err := mockClient.CreateStreamingURL(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateUpdatedImage", func(t *testing.T) {
        input := &appstream.CreateUpdatedImageInput{}
        output := &appstream.CreateUpdatedImageOutput{}

        mockClient.On("CreateUpdatedImage", ctx, input).Return(output, nil)

        result, err := mockClient.CreateUpdatedImage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateUsageReportSubscription", func(t *testing.T) {
        input := &appstream.CreateUsageReportSubscriptionInput{}
        output := &appstream.CreateUsageReportSubscriptionOutput{}

        mockClient.On("CreateUsageReportSubscription", ctx, input).Return(output, nil)

        result, err := mockClient.CreateUsageReportSubscription(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateUser", func(t *testing.T) {
        input := &appstream.CreateUserInput{}
        output := &appstream.CreateUserOutput{}

        mockClient.On("CreateUser", ctx, input).Return(output, nil)

        result, err := mockClient.CreateUser(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteAppBlock", func(t *testing.T) {
        input := &appstream.DeleteAppBlockInput{}
        output := &appstream.DeleteAppBlockOutput{}

        mockClient.On("DeleteAppBlock", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteAppBlock(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteAppBlockBuilder", func(t *testing.T) {
        input := &appstream.DeleteAppBlockBuilderInput{}
        output := &appstream.DeleteAppBlockBuilderOutput{}

        mockClient.On("DeleteAppBlockBuilder", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteAppBlockBuilder(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteApplication", func(t *testing.T) {
        input := &appstream.DeleteApplicationInput{}
        output := &appstream.DeleteApplicationOutput{}

        mockClient.On("DeleteApplication", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteApplication(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteDirectoryConfig", func(t *testing.T) {
        input := &appstream.DeleteDirectoryConfigInput{}
        output := &appstream.DeleteDirectoryConfigOutput{}

        mockClient.On("DeleteDirectoryConfig", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteDirectoryConfig(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteEntitlement", func(t *testing.T) {
        input := &appstream.DeleteEntitlementInput{}
        output := &appstream.DeleteEntitlementOutput{}

        mockClient.On("DeleteEntitlement", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteEntitlement(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteFleet", func(t *testing.T) {
        input := &appstream.DeleteFleetInput{}
        output := &appstream.DeleteFleetOutput{}

        mockClient.On("DeleteFleet", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteFleet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteImage", func(t *testing.T) {
        input := &appstream.DeleteImageInput{}
        output := &appstream.DeleteImageOutput{}

        mockClient.On("DeleteImage", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteImage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteImageBuilder", func(t *testing.T) {
        input := &appstream.DeleteImageBuilderInput{}
        output := &appstream.DeleteImageBuilderOutput{}

        mockClient.On("DeleteImageBuilder", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteImageBuilder(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteImagePermissions", func(t *testing.T) {
        input := &appstream.DeleteImagePermissionsInput{}
        output := &appstream.DeleteImagePermissionsOutput{}

        mockClient.On("DeleteImagePermissions", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteImagePermissions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteStack", func(t *testing.T) {
        input := &appstream.DeleteStackInput{}
        output := &appstream.DeleteStackOutput{}

        mockClient.On("DeleteStack", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteStack(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteUsageReportSubscription", func(t *testing.T) {
        input := &appstream.DeleteUsageReportSubscriptionInput{}
        output := &appstream.DeleteUsageReportSubscriptionOutput{}

        mockClient.On("DeleteUsageReportSubscription", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteUsageReportSubscription(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteUser", func(t *testing.T) {
        input := &appstream.DeleteUserInput{}
        output := &appstream.DeleteUserOutput{}

        mockClient.On("DeleteUser", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteUser(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeAppBlockBuilderAppBlockAssociations", func(t *testing.T) {
        input := &appstream.DescribeAppBlockBuilderAppBlockAssociationsInput{}
        output := &appstream.DescribeAppBlockBuilderAppBlockAssociationsOutput{}

        mockClient.On("DescribeAppBlockBuilderAppBlockAssociations", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeAppBlockBuilderAppBlockAssociations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeAppBlockBuilders", func(t *testing.T) {
        input := &appstream.DescribeAppBlockBuildersInput{}
        output := &appstream.DescribeAppBlockBuildersOutput{}

        mockClient.On("DescribeAppBlockBuilders", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeAppBlockBuilders(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeAppBlocks", func(t *testing.T) {
        input := &appstream.DescribeAppBlocksInput{}
        output := &appstream.DescribeAppBlocksOutput{}

        mockClient.On("DescribeAppBlocks", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeAppBlocks(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeApplicationFleetAssociations", func(t *testing.T) {
        input := &appstream.DescribeApplicationFleetAssociationsInput{}
        output := &appstream.DescribeApplicationFleetAssociationsOutput{}

        mockClient.On("DescribeApplicationFleetAssociations", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeApplicationFleetAssociations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeApplications", func(t *testing.T) {
        input := &appstream.DescribeApplicationsInput{}
        output := &appstream.DescribeApplicationsOutput{}

        mockClient.On("DescribeApplications", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeApplications(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeDirectoryConfigs", func(t *testing.T) {
        input := &appstream.DescribeDirectoryConfigsInput{}
        output := &appstream.DescribeDirectoryConfigsOutput{}

        mockClient.On("DescribeDirectoryConfigs", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeDirectoryConfigs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeEntitlements", func(t *testing.T) {
        input := &appstream.DescribeEntitlementsInput{}
        output := &appstream.DescribeEntitlementsOutput{}

        mockClient.On("DescribeEntitlements", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeEntitlements(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeFleets", func(t *testing.T) {
        input := &appstream.DescribeFleetsInput{}
        output := &appstream.DescribeFleetsOutput{}

        mockClient.On("DescribeFleets", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeFleets(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeImageBuilders", func(t *testing.T) {
        input := &appstream.DescribeImageBuildersInput{}
        output := &appstream.DescribeImageBuildersOutput{}

        mockClient.On("DescribeImageBuilders", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeImageBuilders(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeImagePermissions", func(t *testing.T) {
        input := &appstream.DescribeImagePermissionsInput{}
        output := &appstream.DescribeImagePermissionsOutput{}

        mockClient.On("DescribeImagePermissions", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeImagePermissions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeImages", func(t *testing.T) {
        input := &appstream.DescribeImagesInput{}
        output := &appstream.DescribeImagesOutput{}

        mockClient.On("DescribeImages", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeImages(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeSessions", func(t *testing.T) {
        input := &appstream.DescribeSessionsInput{}
        output := &appstream.DescribeSessionsOutput{}

        mockClient.On("DescribeSessions", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeSessions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeStacks", func(t *testing.T) {
        input := &appstream.DescribeStacksInput{}
        output := &appstream.DescribeStacksOutput{}

        mockClient.On("DescribeStacks", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeStacks(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeUsageReportSubscriptions", func(t *testing.T) {
        input := &appstream.DescribeUsageReportSubscriptionsInput{}
        output := &appstream.DescribeUsageReportSubscriptionsOutput{}

        mockClient.On("DescribeUsageReportSubscriptions", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeUsageReportSubscriptions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeUserStackAssociations", func(t *testing.T) {
        input := &appstream.DescribeUserStackAssociationsInput{}
        output := &appstream.DescribeUserStackAssociationsOutput{}

        mockClient.On("DescribeUserStackAssociations", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeUserStackAssociations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeUsers", func(t *testing.T) {
        input := &appstream.DescribeUsersInput{}
        output := &appstream.DescribeUsersOutput{}

        mockClient.On("DescribeUsers", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeUsers(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisableUser", func(t *testing.T) {
        input := &appstream.DisableUserInput{}
        output := &appstream.DisableUserOutput{}

        mockClient.On("DisableUser", ctx, input).Return(output, nil)

        result, err := mockClient.DisableUser(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateAppBlockBuilderAppBlock", func(t *testing.T) {
        input := &appstream.DisassociateAppBlockBuilderAppBlockInput{}
        output := &appstream.DisassociateAppBlockBuilderAppBlockOutput{}

        mockClient.On("DisassociateAppBlockBuilderAppBlock", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateAppBlockBuilderAppBlock(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateApplicationFleet", func(t *testing.T) {
        input := &appstream.DisassociateApplicationFleetInput{}
        output := &appstream.DisassociateApplicationFleetOutput{}

        mockClient.On("DisassociateApplicationFleet", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateApplicationFleet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateApplicationFromEntitlement", func(t *testing.T) {
        input := &appstream.DisassociateApplicationFromEntitlementInput{}
        output := &appstream.DisassociateApplicationFromEntitlementOutput{}

        mockClient.On("DisassociateApplicationFromEntitlement", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateApplicationFromEntitlement(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateFleet", func(t *testing.T) {
        input := &appstream.DisassociateFleetInput{}
        output := &appstream.DisassociateFleetOutput{}

        mockClient.On("DisassociateFleet", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateFleet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestEnableUser", func(t *testing.T) {
        input := &appstream.EnableUserInput{}
        output := &appstream.EnableUserOutput{}

        mockClient.On("EnableUser", ctx, input).Return(output, nil)

        result, err := mockClient.EnableUser(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestExpireSession", func(t *testing.T) {
        input := &appstream.ExpireSessionInput{}
        output := &appstream.ExpireSessionOutput{}

        mockClient.On("ExpireSession", ctx, input).Return(output, nil)

        result, err := mockClient.ExpireSession(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAssociatedFleets", func(t *testing.T) {
        input := &appstream.ListAssociatedFleetsInput{}
        output := &appstream.ListAssociatedFleetsOutput{}

        mockClient.On("ListAssociatedFleets", ctx, input).Return(output, nil)

        result, err := mockClient.ListAssociatedFleets(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAssociatedStacks", func(t *testing.T) {
        input := &appstream.ListAssociatedStacksInput{}
        output := &appstream.ListAssociatedStacksOutput{}

        mockClient.On("ListAssociatedStacks", ctx, input).Return(output, nil)

        result, err := mockClient.ListAssociatedStacks(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListEntitledApplications", func(t *testing.T) {
        input := &appstream.ListEntitledApplicationsInput{}
        output := &appstream.ListEntitledApplicationsOutput{}

        mockClient.On("ListEntitledApplications", ctx, input).Return(output, nil)

        result, err := mockClient.ListEntitledApplications(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &appstream.ListTagsForResourceInput{}
        output := &appstream.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartAppBlockBuilder", func(t *testing.T) {
        input := &appstream.StartAppBlockBuilderInput{}
        output := &appstream.StartAppBlockBuilderOutput{}

        mockClient.On("StartAppBlockBuilder", ctx, input).Return(output, nil)

        result, err := mockClient.StartAppBlockBuilder(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartFleet", func(t *testing.T) {
        input := &appstream.StartFleetInput{}
        output := &appstream.StartFleetOutput{}

        mockClient.On("StartFleet", ctx, input).Return(output, nil)

        result, err := mockClient.StartFleet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartImageBuilder", func(t *testing.T) {
        input := &appstream.StartImageBuilderInput{}
        output := &appstream.StartImageBuilderOutput{}

        mockClient.On("StartImageBuilder", ctx, input).Return(output, nil)

        result, err := mockClient.StartImageBuilder(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopAppBlockBuilder", func(t *testing.T) {
        input := &appstream.StopAppBlockBuilderInput{}
        output := &appstream.StopAppBlockBuilderOutput{}

        mockClient.On("StopAppBlockBuilder", ctx, input).Return(output, nil)

        result, err := mockClient.StopAppBlockBuilder(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopFleet", func(t *testing.T) {
        input := &appstream.StopFleetInput{}
        output := &appstream.StopFleetOutput{}

        mockClient.On("StopFleet", ctx, input).Return(output, nil)

        result, err := mockClient.StopFleet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopImageBuilder", func(t *testing.T) {
        input := &appstream.StopImageBuilderInput{}
        output := &appstream.StopImageBuilderOutput{}

        mockClient.On("StopImageBuilder", ctx, input).Return(output, nil)

        result, err := mockClient.StopImageBuilder(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &appstream.TagResourceInput{}
        output := &appstream.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &appstream.UntagResourceInput{}
        output := &appstream.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateAppBlockBuilder", func(t *testing.T) {
        input := &appstream.UpdateAppBlockBuilderInput{}
        output := &appstream.UpdateAppBlockBuilderOutput{}

        mockClient.On("UpdateAppBlockBuilder", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateAppBlockBuilder(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateApplication", func(t *testing.T) {
        input := &appstream.UpdateApplicationInput{}
        output := &appstream.UpdateApplicationOutput{}

        mockClient.On("UpdateApplication", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateApplication(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateDirectoryConfig", func(t *testing.T) {
        input := &appstream.UpdateDirectoryConfigInput{}
        output := &appstream.UpdateDirectoryConfigOutput{}

        mockClient.On("UpdateDirectoryConfig", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateDirectoryConfig(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateEntitlement", func(t *testing.T) {
        input := &appstream.UpdateEntitlementInput{}
        output := &appstream.UpdateEntitlementOutput{}

        mockClient.On("UpdateEntitlement", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateEntitlement(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateFleet", func(t *testing.T) {
        input := &appstream.UpdateFleetInput{}
        output := &appstream.UpdateFleetOutput{}

        mockClient.On("UpdateFleet", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateFleet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateImagePermissions", func(t *testing.T) {
        input := &appstream.UpdateImagePermissionsInput{}
        output := &appstream.UpdateImagePermissionsOutput{}

        mockClient.On("UpdateImagePermissions", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateImagePermissions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateStack", func(t *testing.T) {
        input := &appstream.UpdateStackInput{}
        output := &appstream.UpdateStackOutput{}

        mockClient.On("UpdateStack", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateStack(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}