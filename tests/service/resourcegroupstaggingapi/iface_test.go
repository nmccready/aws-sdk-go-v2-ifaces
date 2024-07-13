package resourcegroupstaggingapi_test

// tests for the resourcegroupstaggingapi service interface for this ../../../service/resourcegroupstaggingapi/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/resourcegroupstaggingapi"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/resourcegroupstaggingapi/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/resourcegroupstaggingapi/resourcegroupstaggingapi_iface"
	"github.com/stretchr/testify/assert"
)

func TestResourcegroupstaggingapiServiceCanBeMocked(t *testing.T) {
	var iface resourcegroupstaggingapi_iface.IClient
	iface = &resourcegroupstaggingapi.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := resourcegroupstaggingapi.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeReportCreation", func(t *testing.T) {
        input := &resourcegroupstaggingapi.DescribeReportCreationInput{}
        output := &resourcegroupstaggingapi.DescribeReportCreationOutput{}

        mockClient.On("DescribeReportCreation", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeReportCreation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetComplianceSummary", func(t *testing.T) {
        input := &resourcegroupstaggingapi.GetComplianceSummaryInput{}
        output := &resourcegroupstaggingapi.GetComplianceSummaryOutput{}

        mockClient.On("GetComplianceSummary", ctx, input).Return(output, nil)

        result, err := mockClient.GetComplianceSummary(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetResources", func(t *testing.T) {
        input := &resourcegroupstaggingapi.GetResourcesInput{}
        output := &resourcegroupstaggingapi.GetResourcesOutput{}

        mockClient.On("GetResources", ctx, input).Return(output, nil)

        result, err := mockClient.GetResources(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetTagKeys", func(t *testing.T) {
        input := &resourcegroupstaggingapi.GetTagKeysInput{}
        output := &resourcegroupstaggingapi.GetTagKeysOutput{}

        mockClient.On("GetTagKeys", ctx, input).Return(output, nil)

        result, err := mockClient.GetTagKeys(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetTagValues", func(t *testing.T) {
        input := &resourcegroupstaggingapi.GetTagValuesInput{}
        output := &resourcegroupstaggingapi.GetTagValuesOutput{}

        mockClient.On("GetTagValues", ctx, input).Return(output, nil)

        result, err := mockClient.GetTagValues(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartReportCreation", func(t *testing.T) {
        input := &resourcegroupstaggingapi.StartReportCreationInput{}
        output := &resourcegroupstaggingapi.StartReportCreationOutput{}

        mockClient.On("StartReportCreation", ctx, input).Return(output, nil)

        result, err := mockClient.StartReportCreation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResources", func(t *testing.T) {
        input := &resourcegroupstaggingapi.TagResourcesInput{}
        output := &resourcegroupstaggingapi.TagResourcesOutput{}

        mockClient.On("TagResources", ctx, input).Return(output, nil)

        result, err := mockClient.TagResources(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResources", func(t *testing.T) {
        input := &resourcegroupstaggingapi.UntagResourcesInput{}
        output := &resourcegroupstaggingapi.UntagResourcesOutput{}

        mockClient.On("UntagResources", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResources(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
