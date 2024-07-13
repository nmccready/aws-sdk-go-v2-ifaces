package codegurusecurity_test

// tests for the codegurusecurity service interface for this ../../../service/codegurusecurity/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/codegurusecurity"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/codegurusecurity/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/codegurusecurity/codegurusecurity_iface"
	"github.com/stretchr/testify/assert"
)

func TestCodegurusecurityServiceCanBeMocked(t *testing.T) {
	var iface codegurusecurity_iface.IClient
	iface = &codegurusecurity.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := codegurusecurity.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchGetFindings", func(t *testing.T) {
        input := &codegurusecurity.BatchGetFindingsInput{}
        output := &codegurusecurity.BatchGetFindingsOutput{}

        mockClient.On("BatchGetFindings", ctx, input).Return(output, nil)

        result, err := mockClient.BatchGetFindings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateScan", func(t *testing.T) {
        input := &codegurusecurity.CreateScanInput{}
        output := &codegurusecurity.CreateScanOutput{}

        mockClient.On("CreateScan", ctx, input).Return(output, nil)

        result, err := mockClient.CreateScan(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateUploadUrl", func(t *testing.T) {
        input := &codegurusecurity.CreateUploadUrlInput{}
        output := &codegurusecurity.CreateUploadUrlOutput{}

        mockClient.On("CreateUploadUrl", ctx, input).Return(output, nil)

        result, err := mockClient.CreateUploadUrl(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetAccountConfiguration", func(t *testing.T) {
        input := &codegurusecurity.GetAccountConfigurationInput{}
        output := &codegurusecurity.GetAccountConfigurationOutput{}

        mockClient.On("GetAccountConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.GetAccountConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetFindings", func(t *testing.T) {
        input := &codegurusecurity.GetFindingsInput{}
        output := &codegurusecurity.GetFindingsOutput{}

        mockClient.On("GetFindings", ctx, input).Return(output, nil)

        result, err := mockClient.GetFindings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetMetricsSummary", func(t *testing.T) {
        input := &codegurusecurity.GetMetricsSummaryInput{}
        output := &codegurusecurity.GetMetricsSummaryOutput{}

        mockClient.On("GetMetricsSummary", ctx, input).Return(output, nil)

        result, err := mockClient.GetMetricsSummary(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetScan", func(t *testing.T) {
        input := &codegurusecurity.GetScanInput{}
        output := &codegurusecurity.GetScanOutput{}

        mockClient.On("GetScan", ctx, input).Return(output, nil)

        result, err := mockClient.GetScan(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListFindingsMetrics", func(t *testing.T) {
        input := &codegurusecurity.ListFindingsMetricsInput{}
        output := &codegurusecurity.ListFindingsMetricsOutput{}

        mockClient.On("ListFindingsMetrics", ctx, input).Return(output, nil)

        result, err := mockClient.ListFindingsMetrics(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListScans", func(t *testing.T) {
        input := &codegurusecurity.ListScansInput{}
        output := &codegurusecurity.ListScansOutput{}

        mockClient.On("ListScans", ctx, input).Return(output, nil)

        result, err := mockClient.ListScans(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &codegurusecurity.ListTagsForResourceInput{}
        output := &codegurusecurity.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &codegurusecurity.TagResourceInput{}
        output := &codegurusecurity.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &codegurusecurity.UntagResourceInput{}
        output := &codegurusecurity.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateAccountConfiguration", func(t *testing.T) {
        input := &codegurusecurity.UpdateAccountConfigurationInput{}
        output := &codegurusecurity.UpdateAccountConfigurationOutput{}

        mockClient.On("UpdateAccountConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateAccountConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
