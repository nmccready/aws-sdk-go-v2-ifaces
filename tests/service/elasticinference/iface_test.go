package elasticinference_test

// tests for the elasticinference service interface for this ../../../service/elasticinference/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/elasticinference"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/elasticinference/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/elasticinference/elasticinference_iface"
	"github.com/stretchr/testify/assert"
)

func TestElasticinferenceServiceCanBeMocked(t *testing.T) {
	var iface elasticinference_iface.IClient
	iface = &elasticinference.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := elasticinference.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeAcceleratorOfferings", func(t *testing.T) {
        input := &elasticinference.DescribeAcceleratorOfferingsInput{}
        output := &elasticinference.DescribeAcceleratorOfferingsOutput{}

        mockClient.On("DescribeAcceleratorOfferings", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeAcceleratorOfferings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeAcceleratorTypes", func(t *testing.T) {
        input := &elasticinference.DescribeAcceleratorTypesInput{}
        output := &elasticinference.DescribeAcceleratorTypesOutput{}

        mockClient.On("DescribeAcceleratorTypes", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeAcceleratorTypes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeAccelerators", func(t *testing.T) {
        input := &elasticinference.DescribeAcceleratorsInput{}
        output := &elasticinference.DescribeAcceleratorsOutput{}

        mockClient.On("DescribeAccelerators", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeAccelerators(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &elasticinference.ListTagsForResourceInput{}
        output := &elasticinference.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &elasticinference.TagResourceInput{}
        output := &elasticinference.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &elasticinference.UntagResourceInput{}
        output := &elasticinference.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
