package shield_test

// tests for the shield service interface for this ../../../service/shield/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/shield"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/shield/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/shield/shield_iface"
	"github.com/stretchr/testify/assert"
)

func TestShieldServiceCanBeMocked(t *testing.T) {
	var iface shield_iface.IClient
	iface = &shield.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := shield.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateDRTLogBucket", func(t *testing.T) {
        input := &shield.AssociateDRTLogBucketInput{}
        output := &shield.AssociateDRTLogBucketOutput{}

        mockClient.On("AssociateDRTLogBucket", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateDRTLogBucket(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateDRTRole", func(t *testing.T) {
        input := &shield.AssociateDRTRoleInput{}
        output := &shield.AssociateDRTRoleOutput{}

        mockClient.On("AssociateDRTRole", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateDRTRole(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateHealthCheck", func(t *testing.T) {
        input := &shield.AssociateHealthCheckInput{}
        output := &shield.AssociateHealthCheckOutput{}

        mockClient.On("AssociateHealthCheck", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateHealthCheck(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateProactiveEngagementDetails", func(t *testing.T) {
        input := &shield.AssociateProactiveEngagementDetailsInput{}
        output := &shield.AssociateProactiveEngagementDetailsOutput{}

        mockClient.On("AssociateProactiveEngagementDetails", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateProactiveEngagementDetails(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateProtection", func(t *testing.T) {
        input := &shield.CreateProtectionInput{}
        output := &shield.CreateProtectionOutput{}

        mockClient.On("CreateProtection", ctx, input).Return(output, nil)

        result, err := mockClient.CreateProtection(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateProtectionGroup", func(t *testing.T) {
        input := &shield.CreateProtectionGroupInput{}
        output := &shield.CreateProtectionGroupOutput{}

        mockClient.On("CreateProtectionGroup", ctx, input).Return(output, nil)

        result, err := mockClient.CreateProtectionGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateSubscription", func(t *testing.T) {
        input := &shield.CreateSubscriptionInput{}
        output := &shield.CreateSubscriptionOutput{}

        mockClient.On("CreateSubscription", ctx, input).Return(output, nil)

        result, err := mockClient.CreateSubscription(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteProtection", func(t *testing.T) {
        input := &shield.DeleteProtectionInput{}
        output := &shield.DeleteProtectionOutput{}

        mockClient.On("DeleteProtection", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteProtection(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteProtectionGroup", func(t *testing.T) {
        input := &shield.DeleteProtectionGroupInput{}
        output := &shield.DeleteProtectionGroupOutput{}

        mockClient.On("DeleteProtectionGroup", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteProtectionGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteSubscription", func(t *testing.T) {
        input := &shield.DeleteSubscriptionInput{}
        output := &shield.DeleteSubscriptionOutput{}

        mockClient.On("DeleteSubscription", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteSubscription(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeAttack", func(t *testing.T) {
        input := &shield.DescribeAttackInput{}
        output := &shield.DescribeAttackOutput{}

        mockClient.On("DescribeAttack", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeAttack(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeAttackStatistics", func(t *testing.T) {
        input := &shield.DescribeAttackStatisticsInput{}
        output := &shield.DescribeAttackStatisticsOutput{}

        mockClient.On("DescribeAttackStatistics", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeAttackStatistics(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeDRTAccess", func(t *testing.T) {
        input := &shield.DescribeDRTAccessInput{}
        output := &shield.DescribeDRTAccessOutput{}

        mockClient.On("DescribeDRTAccess", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeDRTAccess(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeEmergencyContactSettings", func(t *testing.T) {
        input := &shield.DescribeEmergencyContactSettingsInput{}
        output := &shield.DescribeEmergencyContactSettingsOutput{}

        mockClient.On("DescribeEmergencyContactSettings", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeEmergencyContactSettings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeProtection", func(t *testing.T) {
        input := &shield.DescribeProtectionInput{}
        output := &shield.DescribeProtectionOutput{}

        mockClient.On("DescribeProtection", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeProtection(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeProtectionGroup", func(t *testing.T) {
        input := &shield.DescribeProtectionGroupInput{}
        output := &shield.DescribeProtectionGroupOutput{}

        mockClient.On("DescribeProtectionGroup", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeProtectionGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeSubscription", func(t *testing.T) {
        input := &shield.DescribeSubscriptionInput{}
        output := &shield.DescribeSubscriptionOutput{}

        mockClient.On("DescribeSubscription", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeSubscription(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisableApplicationLayerAutomaticResponse", func(t *testing.T) {
        input := &shield.DisableApplicationLayerAutomaticResponseInput{}
        output := &shield.DisableApplicationLayerAutomaticResponseOutput{}

        mockClient.On("DisableApplicationLayerAutomaticResponse", ctx, input).Return(output, nil)

        result, err := mockClient.DisableApplicationLayerAutomaticResponse(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisableProactiveEngagement", func(t *testing.T) {
        input := &shield.DisableProactiveEngagementInput{}
        output := &shield.DisableProactiveEngagementOutput{}

        mockClient.On("DisableProactiveEngagement", ctx, input).Return(output, nil)

        result, err := mockClient.DisableProactiveEngagement(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateDRTLogBucket", func(t *testing.T) {
        input := &shield.DisassociateDRTLogBucketInput{}
        output := &shield.DisassociateDRTLogBucketOutput{}

        mockClient.On("DisassociateDRTLogBucket", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateDRTLogBucket(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateDRTRole", func(t *testing.T) {
        input := &shield.DisassociateDRTRoleInput{}
        output := &shield.DisassociateDRTRoleOutput{}

        mockClient.On("DisassociateDRTRole", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateDRTRole(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateHealthCheck", func(t *testing.T) {
        input := &shield.DisassociateHealthCheckInput{}
        output := &shield.DisassociateHealthCheckOutput{}

        mockClient.On("DisassociateHealthCheck", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateHealthCheck(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestEnableApplicationLayerAutomaticResponse", func(t *testing.T) {
        input := &shield.EnableApplicationLayerAutomaticResponseInput{}
        output := &shield.EnableApplicationLayerAutomaticResponseOutput{}

        mockClient.On("EnableApplicationLayerAutomaticResponse", ctx, input).Return(output, nil)

        result, err := mockClient.EnableApplicationLayerAutomaticResponse(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestEnableProactiveEngagement", func(t *testing.T) {
        input := &shield.EnableProactiveEngagementInput{}
        output := &shield.EnableProactiveEngagementOutput{}

        mockClient.On("EnableProactiveEngagement", ctx, input).Return(output, nil)

        result, err := mockClient.EnableProactiveEngagement(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSubscriptionState", func(t *testing.T) {
        input := &shield.GetSubscriptionStateInput{}
        output := &shield.GetSubscriptionStateOutput{}

        mockClient.On("GetSubscriptionState", ctx, input).Return(output, nil)

        result, err := mockClient.GetSubscriptionState(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAttacks", func(t *testing.T) {
        input := &shield.ListAttacksInput{}
        output := &shield.ListAttacksOutput{}

        mockClient.On("ListAttacks", ctx, input).Return(output, nil)

        result, err := mockClient.ListAttacks(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListProtectionGroups", func(t *testing.T) {
        input := &shield.ListProtectionGroupsInput{}
        output := &shield.ListProtectionGroupsOutput{}

        mockClient.On("ListProtectionGroups", ctx, input).Return(output, nil)

        result, err := mockClient.ListProtectionGroups(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListProtections", func(t *testing.T) {
        input := &shield.ListProtectionsInput{}
        output := &shield.ListProtectionsOutput{}

        mockClient.On("ListProtections", ctx, input).Return(output, nil)

        result, err := mockClient.ListProtections(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListResourcesInProtectionGroup", func(t *testing.T) {
        input := &shield.ListResourcesInProtectionGroupInput{}
        output := &shield.ListResourcesInProtectionGroupOutput{}

        mockClient.On("ListResourcesInProtectionGroup", ctx, input).Return(output, nil)

        result, err := mockClient.ListResourcesInProtectionGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &shield.ListTagsForResourceInput{}
        output := &shield.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &shield.TagResourceInput{}
        output := &shield.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &shield.UntagResourceInput{}
        output := &shield.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateApplicationLayerAutomaticResponse", func(t *testing.T) {
        input := &shield.UpdateApplicationLayerAutomaticResponseInput{}
        output := &shield.UpdateApplicationLayerAutomaticResponseOutput{}

        mockClient.On("UpdateApplicationLayerAutomaticResponse", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateApplicationLayerAutomaticResponse(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateEmergencyContactSettings", func(t *testing.T) {
        input := &shield.UpdateEmergencyContactSettingsInput{}
        output := &shield.UpdateEmergencyContactSettingsOutput{}

        mockClient.On("UpdateEmergencyContactSettings", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateEmergencyContactSettings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateProtectionGroup", func(t *testing.T) {
        input := &shield.UpdateProtectionGroupInput{}
        output := &shield.UpdateProtectionGroupOutput{}

        mockClient.On("UpdateProtectionGroup", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateProtectionGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateSubscription", func(t *testing.T) {
        input := &shield.UpdateSubscriptionInput{}
        output := &shield.UpdateSubscriptionOutput{}

        mockClient.On("UpdateSubscription", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateSubscription(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
