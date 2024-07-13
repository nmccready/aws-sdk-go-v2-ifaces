package mwaa_test

// tests for the mwaa service interface for this ../../../service/mwaa/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/mwaa"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/mwaa/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/mwaa/mwaa_iface"
	"github.com/stretchr/testify/assert"
)

func TestMwaaServiceCanBeMocked(t *testing.T) {
	var iface mwaa_iface.IClient
	iface = &mwaa.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := mwaa.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateCliToken", func(t *testing.T) {
        input := &mwaa.CreateCliTokenInput{}
        output := &mwaa.CreateCliTokenOutput{}

        mockClient.On("CreateCliToken", ctx, input).Return(output, nil)

        result, err := mockClient.CreateCliToken(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateEnvironment", func(t *testing.T) {
        input := &mwaa.CreateEnvironmentInput{}
        output := &mwaa.CreateEnvironmentOutput{}

        mockClient.On("CreateEnvironment", ctx, input).Return(output, nil)

        result, err := mockClient.CreateEnvironment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateWebLoginToken", func(t *testing.T) {
        input := &mwaa.CreateWebLoginTokenInput{}
        output := &mwaa.CreateWebLoginTokenOutput{}

        mockClient.On("CreateWebLoginToken", ctx, input).Return(output, nil)

        result, err := mockClient.CreateWebLoginToken(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteEnvironment", func(t *testing.T) {
        input := &mwaa.DeleteEnvironmentInput{}
        output := &mwaa.DeleteEnvironmentOutput{}

        mockClient.On("DeleteEnvironment", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteEnvironment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetEnvironment", func(t *testing.T) {
        input := &mwaa.GetEnvironmentInput{}
        output := &mwaa.GetEnvironmentOutput{}

        mockClient.On("GetEnvironment", ctx, input).Return(output, nil)

        result, err := mockClient.GetEnvironment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListEnvironments", func(t *testing.T) {
        input := &mwaa.ListEnvironmentsInput{}
        output := &mwaa.ListEnvironmentsOutput{}

        mockClient.On("ListEnvironments", ctx, input).Return(output, nil)

        result, err := mockClient.ListEnvironments(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &mwaa.ListTagsForResourceInput{}
        output := &mwaa.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPublishMetrics", func(t *testing.T) {
        input := &mwaa.PublishMetricsInput{}
        output := &mwaa.PublishMetricsOutput{}

        mockClient.On("PublishMetrics", ctx, input).Return(output, nil)

        result, err := mockClient.PublishMetrics(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &mwaa.TagResourceInput{}
        output := &mwaa.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &mwaa.UntagResourceInput{}
        output := &mwaa.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateEnvironment", func(t *testing.T) {
        input := &mwaa.UpdateEnvironmentInput{}
        output := &mwaa.UpdateEnvironmentOutput{}

        mockClient.On("UpdateEnvironment", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateEnvironment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
