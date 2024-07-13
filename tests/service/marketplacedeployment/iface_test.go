package marketplacedeployment_test

// tests for the marketplacedeployment service interface for this ../../../service/marketplacedeployment/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/marketplacedeployment"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/marketplacedeployment/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/marketplacedeployment/marketplacedeployment_iface"
	"github.com/stretchr/testify/assert"
)

func TestMarketplacedeploymentServiceCanBeMocked(t *testing.T) {
	var iface marketplacedeployment_iface.IClient
	iface = &marketplacedeployment.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := marketplacedeployment.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &marketplacedeployment.ListTagsForResourceInput{}
        output := &marketplacedeployment.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutDeploymentParameter", func(t *testing.T) {
        input := &marketplacedeployment.PutDeploymentParameterInput{}
        output := &marketplacedeployment.PutDeploymentParameterOutput{}

        mockClient.On("PutDeploymentParameter", ctx, input).Return(output, nil)

        result, err := mockClient.PutDeploymentParameter(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &marketplacedeployment.TagResourceInput{}
        output := &marketplacedeployment.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &marketplacedeployment.UntagResourceInput{}
        output := &marketplacedeployment.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
