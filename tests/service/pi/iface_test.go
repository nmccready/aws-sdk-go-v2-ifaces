package pi_test

// tests for the pi service interface for this ../../../service/pi/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/pi"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/pi/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/pi/pi_iface"
	"github.com/stretchr/testify/assert"
)

func TestPiServiceCanBeMocked(t *testing.T) {
	var iface pi_iface.IClient
	iface = &pi.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := pi.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreatePerformanceAnalysisReport", func(t *testing.T) {
        input := &pi.CreatePerformanceAnalysisReportInput{}
        output := &pi.CreatePerformanceAnalysisReportOutput{}

        mockClient.On("CreatePerformanceAnalysisReport", ctx, input).Return(output, nil)

        result, err := mockClient.CreatePerformanceAnalysisReport(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeletePerformanceAnalysisReport", func(t *testing.T) {
        input := &pi.DeletePerformanceAnalysisReportInput{}
        output := &pi.DeletePerformanceAnalysisReportOutput{}

        mockClient.On("DeletePerformanceAnalysisReport", ctx, input).Return(output, nil)

        result, err := mockClient.DeletePerformanceAnalysisReport(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeDimensionKeys", func(t *testing.T) {
        input := &pi.DescribeDimensionKeysInput{}
        output := &pi.DescribeDimensionKeysOutput{}

        mockClient.On("DescribeDimensionKeys", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeDimensionKeys(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDimensionKeyDetails", func(t *testing.T) {
        input := &pi.GetDimensionKeyDetailsInput{}
        output := &pi.GetDimensionKeyDetailsOutput{}

        mockClient.On("GetDimensionKeyDetails", ctx, input).Return(output, nil)

        result, err := mockClient.GetDimensionKeyDetails(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetPerformanceAnalysisReport", func(t *testing.T) {
        input := &pi.GetPerformanceAnalysisReportInput{}
        output := &pi.GetPerformanceAnalysisReportOutput{}

        mockClient.On("GetPerformanceAnalysisReport", ctx, input).Return(output, nil)

        result, err := mockClient.GetPerformanceAnalysisReport(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetResourceMetadata", func(t *testing.T) {
        input := &pi.GetResourceMetadataInput{}
        output := &pi.GetResourceMetadataOutput{}

        mockClient.On("GetResourceMetadata", ctx, input).Return(output, nil)

        result, err := mockClient.GetResourceMetadata(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetResourceMetrics", func(t *testing.T) {
        input := &pi.GetResourceMetricsInput{}
        output := &pi.GetResourceMetricsOutput{}

        mockClient.On("GetResourceMetrics", ctx, input).Return(output, nil)

        result, err := mockClient.GetResourceMetrics(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAvailableResourceDimensions", func(t *testing.T) {
        input := &pi.ListAvailableResourceDimensionsInput{}
        output := &pi.ListAvailableResourceDimensionsOutput{}

        mockClient.On("ListAvailableResourceDimensions", ctx, input).Return(output, nil)

        result, err := mockClient.ListAvailableResourceDimensions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAvailableResourceMetrics", func(t *testing.T) {
        input := &pi.ListAvailableResourceMetricsInput{}
        output := &pi.ListAvailableResourceMetricsOutput{}

        mockClient.On("ListAvailableResourceMetrics", ctx, input).Return(output, nil)

        result, err := mockClient.ListAvailableResourceMetrics(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListPerformanceAnalysisReports", func(t *testing.T) {
        input := &pi.ListPerformanceAnalysisReportsInput{}
        output := &pi.ListPerformanceAnalysisReportsOutput{}

        mockClient.On("ListPerformanceAnalysisReports", ctx, input).Return(output, nil)

        result, err := mockClient.ListPerformanceAnalysisReports(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &pi.ListTagsForResourceInput{}
        output := &pi.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &pi.TagResourceInput{}
        output := &pi.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &pi.UntagResourceInput{}
        output := &pi.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
