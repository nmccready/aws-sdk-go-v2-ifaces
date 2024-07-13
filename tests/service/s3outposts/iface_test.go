package s3outposts_test

// tests for the s3outposts service interface for this ../../../service/s3outposts/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/s3outposts"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/s3outposts/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/s3outposts/s3outposts_iface"
	"github.com/stretchr/testify/assert"
)

func TestS3outpostsServiceCanBeMocked(t *testing.T) {
	var iface s3outposts_iface.IClient
	iface = &s3outposts.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := s3outposts.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateEndpoint", func(t *testing.T) {
        input := &s3outposts.CreateEndpointInput{}
        output := &s3outposts.CreateEndpointOutput{}

        mockClient.On("CreateEndpoint", ctx, input).Return(output, nil)

        result, err := mockClient.CreateEndpoint(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteEndpoint", func(t *testing.T) {
        input := &s3outposts.DeleteEndpointInput{}
        output := &s3outposts.DeleteEndpointOutput{}

        mockClient.On("DeleteEndpoint", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteEndpoint(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListEndpoints", func(t *testing.T) {
        input := &s3outposts.ListEndpointsInput{}
        output := &s3outposts.ListEndpointsOutput{}

        mockClient.On("ListEndpoints", ctx, input).Return(output, nil)

        result, err := mockClient.ListEndpoints(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListOutpostsWithS3", func(t *testing.T) {
        input := &s3outposts.ListOutpostsWithS3Input{}
        output := &s3outposts.ListOutpostsWithS3Output{}

        mockClient.On("ListOutpostsWithS3", ctx, input).Return(output, nil)

        result, err := mockClient.ListOutpostsWithS3(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListSharedEndpoints", func(t *testing.T) {
        input := &s3outposts.ListSharedEndpointsInput{}
        output := &s3outposts.ListSharedEndpointsOutput{}

        mockClient.On("ListSharedEndpoints", ctx, input).Return(output, nil)

        result, err := mockClient.ListSharedEndpoints(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
