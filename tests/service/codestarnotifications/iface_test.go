package codestarnotifications_test

// tests for the codestarnotifications service interface for this ../../../service/codestarnotifications/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/codestarnotifications"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/codestarnotifications/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/codestarnotifications/codestarnotifications_iface"
	"github.com/stretchr/testify/assert"
)

func TestCodestarnotificationsServiceCanBeMocked(t *testing.T) {
	var iface codestarnotifications_iface.IClient
	iface = &codestarnotifications.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := codestarnotifications.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateNotificationRule", func(t *testing.T) {
        input := &codestarnotifications.CreateNotificationRuleInput{}
        output := &codestarnotifications.CreateNotificationRuleOutput{}

        mockClient.On("CreateNotificationRule", ctx, input).Return(output, nil)

        result, err := mockClient.CreateNotificationRule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteNotificationRule", func(t *testing.T) {
        input := &codestarnotifications.DeleteNotificationRuleInput{}
        output := &codestarnotifications.DeleteNotificationRuleOutput{}

        mockClient.On("DeleteNotificationRule", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteNotificationRule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteTarget", func(t *testing.T) {
        input := &codestarnotifications.DeleteTargetInput{}
        output := &codestarnotifications.DeleteTargetOutput{}

        mockClient.On("DeleteTarget", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteTarget(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeNotificationRule", func(t *testing.T) {
        input := &codestarnotifications.DescribeNotificationRuleInput{}
        output := &codestarnotifications.DescribeNotificationRuleOutput{}

        mockClient.On("DescribeNotificationRule", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeNotificationRule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListEventTypes", func(t *testing.T) {
        input := &codestarnotifications.ListEventTypesInput{}
        output := &codestarnotifications.ListEventTypesOutput{}

        mockClient.On("ListEventTypes", ctx, input).Return(output, nil)

        result, err := mockClient.ListEventTypes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListNotificationRules", func(t *testing.T) {
        input := &codestarnotifications.ListNotificationRulesInput{}
        output := &codestarnotifications.ListNotificationRulesOutput{}

        mockClient.On("ListNotificationRules", ctx, input).Return(output, nil)

        result, err := mockClient.ListNotificationRules(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &codestarnotifications.ListTagsForResourceInput{}
        output := &codestarnotifications.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTargets", func(t *testing.T) {
        input := &codestarnotifications.ListTargetsInput{}
        output := &codestarnotifications.ListTargetsOutput{}

        mockClient.On("ListTargets", ctx, input).Return(output, nil)

        result, err := mockClient.ListTargets(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSubscribe", func(t *testing.T) {
        input := &codestarnotifications.SubscribeInput{}
        output := &codestarnotifications.SubscribeOutput{}

        mockClient.On("Subscribe", ctx, input).Return(output, nil)

        result, err := mockClient.Subscribe(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &codestarnotifications.TagResourceInput{}
        output := &codestarnotifications.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUnsubscribe", func(t *testing.T) {
        input := &codestarnotifications.UnsubscribeInput{}
        output := &codestarnotifications.UnsubscribeOutput{}

        mockClient.On("Unsubscribe", ctx, input).Return(output, nil)

        result, err := mockClient.Unsubscribe(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &codestarnotifications.UntagResourceInput{}
        output := &codestarnotifications.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateNotificationRule", func(t *testing.T) {
        input := &codestarnotifications.UpdateNotificationRuleInput{}
        output := &codestarnotifications.UpdateNotificationRuleOutput{}

        mockClient.On("UpdateNotificationRule", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateNotificationRule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
