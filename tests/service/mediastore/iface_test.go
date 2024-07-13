package mediastore_test

// tests for the mediastore service interface for this ../../../service/mediastore/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/mediastore"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/mediastore/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/mediastore/mediastore_iface"
	"github.com/stretchr/testify/assert"
)

func TestMediastoreServiceCanBeMocked(t *testing.T) {
	var iface mediastore_iface.IClient
	iface = &mediastore.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := mediastore.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateContainer", func(t *testing.T) {
        input := &mediastore.CreateContainerInput{}
        output := &mediastore.CreateContainerOutput{}

        mockClient.On("CreateContainer", ctx, input).Return(output, nil)

        result, err := mockClient.CreateContainer(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteContainer", func(t *testing.T) {
        input := &mediastore.DeleteContainerInput{}
        output := &mediastore.DeleteContainerOutput{}

        mockClient.On("DeleteContainer", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteContainer(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteContainerPolicy", func(t *testing.T) {
        input := &mediastore.DeleteContainerPolicyInput{}
        output := &mediastore.DeleteContainerPolicyOutput{}

        mockClient.On("DeleteContainerPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteContainerPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteCorsPolicy", func(t *testing.T) {
        input := &mediastore.DeleteCorsPolicyInput{}
        output := &mediastore.DeleteCorsPolicyOutput{}

        mockClient.On("DeleteCorsPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteCorsPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteLifecyclePolicy", func(t *testing.T) {
        input := &mediastore.DeleteLifecyclePolicyInput{}
        output := &mediastore.DeleteLifecyclePolicyOutput{}

        mockClient.On("DeleteLifecyclePolicy", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteLifecyclePolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteMetricPolicy", func(t *testing.T) {
        input := &mediastore.DeleteMetricPolicyInput{}
        output := &mediastore.DeleteMetricPolicyOutput{}

        mockClient.On("DeleteMetricPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteMetricPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeContainer", func(t *testing.T) {
        input := &mediastore.DescribeContainerInput{}
        output := &mediastore.DescribeContainerOutput{}

        mockClient.On("DescribeContainer", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeContainer(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetContainerPolicy", func(t *testing.T) {
        input := &mediastore.GetContainerPolicyInput{}
        output := &mediastore.GetContainerPolicyOutput{}

        mockClient.On("GetContainerPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.GetContainerPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetCorsPolicy", func(t *testing.T) {
        input := &mediastore.GetCorsPolicyInput{}
        output := &mediastore.GetCorsPolicyOutput{}

        mockClient.On("GetCorsPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.GetCorsPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetLifecyclePolicy", func(t *testing.T) {
        input := &mediastore.GetLifecyclePolicyInput{}
        output := &mediastore.GetLifecyclePolicyOutput{}

        mockClient.On("GetLifecyclePolicy", ctx, input).Return(output, nil)

        result, err := mockClient.GetLifecyclePolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetMetricPolicy", func(t *testing.T) {
        input := &mediastore.GetMetricPolicyInput{}
        output := &mediastore.GetMetricPolicyOutput{}

        mockClient.On("GetMetricPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.GetMetricPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListContainers", func(t *testing.T) {
        input := &mediastore.ListContainersInput{}
        output := &mediastore.ListContainersOutput{}

        mockClient.On("ListContainers", ctx, input).Return(output, nil)

        result, err := mockClient.ListContainers(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &mediastore.ListTagsForResourceInput{}
        output := &mediastore.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutContainerPolicy", func(t *testing.T) {
        input := &mediastore.PutContainerPolicyInput{}
        output := &mediastore.PutContainerPolicyOutput{}

        mockClient.On("PutContainerPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.PutContainerPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutCorsPolicy", func(t *testing.T) {
        input := &mediastore.PutCorsPolicyInput{}
        output := &mediastore.PutCorsPolicyOutput{}

        mockClient.On("PutCorsPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.PutCorsPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutLifecyclePolicy", func(t *testing.T) {
        input := &mediastore.PutLifecyclePolicyInput{}
        output := &mediastore.PutLifecyclePolicyOutput{}

        mockClient.On("PutLifecyclePolicy", ctx, input).Return(output, nil)

        result, err := mockClient.PutLifecyclePolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutMetricPolicy", func(t *testing.T) {
        input := &mediastore.PutMetricPolicyInput{}
        output := &mediastore.PutMetricPolicyOutput{}

        mockClient.On("PutMetricPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.PutMetricPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartAccessLogging", func(t *testing.T) {
        input := &mediastore.StartAccessLoggingInput{}
        output := &mediastore.StartAccessLoggingOutput{}

        mockClient.On("StartAccessLogging", ctx, input).Return(output, nil)

        result, err := mockClient.StartAccessLogging(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopAccessLogging", func(t *testing.T) {
        input := &mediastore.StopAccessLoggingInput{}
        output := &mediastore.StopAccessLoggingOutput{}

        mockClient.On("StopAccessLogging", ctx, input).Return(output, nil)

        result, err := mockClient.StopAccessLogging(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &mediastore.TagResourceInput{}
        output := &mediastore.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &mediastore.UntagResourceInput{}
        output := &mediastore.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
