package dlm_test

// tests for the dlm service interface for this ../../../service/dlm/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/dlm"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/dlm/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/dlm/dlm_iface"
	"github.com/stretchr/testify/assert"
)

func TestDlmServiceCanBeMocked(t *testing.T) {
	var iface dlm_iface.IClient
	iface = &dlm.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := dlm.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateLifecyclePolicy", func(t *testing.T) {
        input := &dlm.CreateLifecyclePolicyInput{}
        output := &dlm.CreateLifecyclePolicyOutput{}

        mockClient.On("CreateLifecyclePolicy", ctx, input).Return(output, nil)

        result, err := mockClient.CreateLifecyclePolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteLifecyclePolicy", func(t *testing.T) {
        input := &dlm.DeleteLifecyclePolicyInput{}
        output := &dlm.DeleteLifecyclePolicyOutput{}

        mockClient.On("DeleteLifecyclePolicy", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteLifecyclePolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetLifecyclePolicies", func(t *testing.T) {
        input := &dlm.GetLifecyclePoliciesInput{}
        output := &dlm.GetLifecyclePoliciesOutput{}

        mockClient.On("GetLifecyclePolicies", ctx, input).Return(output, nil)

        result, err := mockClient.GetLifecyclePolicies(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetLifecyclePolicy", func(t *testing.T) {
        input := &dlm.GetLifecyclePolicyInput{}
        output := &dlm.GetLifecyclePolicyOutput{}

        mockClient.On("GetLifecyclePolicy", ctx, input).Return(output, nil)

        result, err := mockClient.GetLifecyclePolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &dlm.ListTagsForResourceInput{}
        output := &dlm.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &dlm.TagResourceInput{}
        output := &dlm.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &dlm.UntagResourceInput{}
        output := &dlm.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateLifecyclePolicy", func(t *testing.T) {
        input := &dlm.UpdateLifecyclePolicyInput{}
        output := &dlm.UpdateLifecyclePolicyOutput{}

        mockClient.On("UpdateLifecyclePolicy", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateLifecyclePolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
