package iotfleethub_test

// tests for the iotfleethub service interface for this ../../../service/iotfleethub/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/iotfleethub"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/iotfleethub/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/iotfleethub/iotfleethub_iface"
	"github.com/stretchr/testify/assert"
)

func TestIotfleethubServiceCanBeMocked(t *testing.T) {
	var iface iotfleethub_iface.IClient
	iface = &iotfleethub.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := iotfleethub.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateApplication", func(t *testing.T) {
        input := &iotfleethub.CreateApplicationInput{}
        output := &iotfleethub.CreateApplicationOutput{}

        mockClient.On("CreateApplication", ctx, input).Return(output, nil)

        result, err := mockClient.CreateApplication(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteApplication", func(t *testing.T) {
        input := &iotfleethub.DeleteApplicationInput{}
        output := &iotfleethub.DeleteApplicationOutput{}

        mockClient.On("DeleteApplication", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteApplication(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeApplication", func(t *testing.T) {
        input := &iotfleethub.DescribeApplicationInput{}
        output := &iotfleethub.DescribeApplicationOutput{}

        mockClient.On("DescribeApplication", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeApplication(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListApplications", func(t *testing.T) {
        input := &iotfleethub.ListApplicationsInput{}
        output := &iotfleethub.ListApplicationsOutput{}

        mockClient.On("ListApplications", ctx, input).Return(output, nil)

        result, err := mockClient.ListApplications(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &iotfleethub.ListTagsForResourceInput{}
        output := &iotfleethub.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &iotfleethub.TagResourceInput{}
        output := &iotfleethub.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &iotfleethub.UntagResourceInput{}
        output := &iotfleethub.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateApplication", func(t *testing.T) {
        input := &iotfleethub.UpdateApplicationInput{}
        output := &iotfleethub.UpdateApplicationOutput{}

        mockClient.On("UpdateApplication", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateApplication(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
