package rbin_test

// tests for the rbin service interface for this ../../../service/rbin/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/rbin"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/rbin/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/rbin/rbin_iface"
	"github.com/stretchr/testify/assert"
)

func TestRbinServiceCanBeMocked(t *testing.T) {
	var iface rbin_iface.IClient
	iface = &rbin.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := rbin.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateRule", func(t *testing.T) {
        input := &rbin.CreateRuleInput{}
        output := &rbin.CreateRuleOutput{}

        mockClient.On("CreateRule", ctx, input).Return(output, nil)

        result, err := mockClient.CreateRule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteRule", func(t *testing.T) {
        input := &rbin.DeleteRuleInput{}
        output := &rbin.DeleteRuleOutput{}

        mockClient.On("DeleteRule", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteRule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetRule", func(t *testing.T) {
        input := &rbin.GetRuleInput{}
        output := &rbin.GetRuleOutput{}

        mockClient.On("GetRule", ctx, input).Return(output, nil)

        result, err := mockClient.GetRule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListRules", func(t *testing.T) {
        input := &rbin.ListRulesInput{}
        output := &rbin.ListRulesOutput{}

        mockClient.On("ListRules", ctx, input).Return(output, nil)

        result, err := mockClient.ListRules(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &rbin.ListTagsForResourceInput{}
        output := &rbin.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestLockRule", func(t *testing.T) {
        input := &rbin.LockRuleInput{}
        output := &rbin.LockRuleOutput{}

        mockClient.On("LockRule", ctx, input).Return(output, nil)

        result, err := mockClient.LockRule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &rbin.TagResourceInput{}
        output := &rbin.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUnlockRule", func(t *testing.T) {
        input := &rbin.UnlockRuleInput{}
        output := &rbin.UnlockRuleOutput{}

        mockClient.On("UnlockRule", ctx, input).Return(output, nil)

        result, err := mockClient.UnlockRule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &rbin.UntagResourceInput{}
        output := &rbin.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateRule", func(t *testing.T) {
        input := &rbin.UpdateRuleInput{}
        output := &rbin.UpdateRuleOutput{}

        mockClient.On("UpdateRule", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateRule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
