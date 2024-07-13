package lookoutmetrics_test

// tests for the lookoutmetrics service interface for this ../../../service/lookoutmetrics/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/lookoutmetrics"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/lookoutmetrics/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/lookoutmetrics/lookoutmetrics_iface"
	"github.com/stretchr/testify/assert"
)

func TestLookoutmetricsServiceCanBeMocked(t *testing.T) {
	var iface lookoutmetrics_iface.IClient
	iface = &lookoutmetrics.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := lookoutmetrics.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestActivateAnomalyDetector", func(t *testing.T) {
        input := &lookoutmetrics.ActivateAnomalyDetectorInput{}
        output := &lookoutmetrics.ActivateAnomalyDetectorOutput{}

        mockClient.On("ActivateAnomalyDetector", ctx, input).Return(output, nil)

        result, err := mockClient.ActivateAnomalyDetector(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBackTestAnomalyDetector", func(t *testing.T) {
        input := &lookoutmetrics.BackTestAnomalyDetectorInput{}
        output := &lookoutmetrics.BackTestAnomalyDetectorOutput{}

        mockClient.On("BackTestAnomalyDetector", ctx, input).Return(output, nil)

        result, err := mockClient.BackTestAnomalyDetector(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateAlert", func(t *testing.T) {
        input := &lookoutmetrics.CreateAlertInput{}
        output := &lookoutmetrics.CreateAlertOutput{}

        mockClient.On("CreateAlert", ctx, input).Return(output, nil)

        result, err := mockClient.CreateAlert(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateAnomalyDetector", func(t *testing.T) {
        input := &lookoutmetrics.CreateAnomalyDetectorInput{}
        output := &lookoutmetrics.CreateAnomalyDetectorOutput{}

        mockClient.On("CreateAnomalyDetector", ctx, input).Return(output, nil)

        result, err := mockClient.CreateAnomalyDetector(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateMetricSet", func(t *testing.T) {
        input := &lookoutmetrics.CreateMetricSetInput{}
        output := &lookoutmetrics.CreateMetricSetOutput{}

        mockClient.On("CreateMetricSet", ctx, input).Return(output, nil)

        result, err := mockClient.CreateMetricSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeactivateAnomalyDetector", func(t *testing.T) {
        input := &lookoutmetrics.DeactivateAnomalyDetectorInput{}
        output := &lookoutmetrics.DeactivateAnomalyDetectorOutput{}

        mockClient.On("DeactivateAnomalyDetector", ctx, input).Return(output, nil)

        result, err := mockClient.DeactivateAnomalyDetector(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteAlert", func(t *testing.T) {
        input := &lookoutmetrics.DeleteAlertInput{}
        output := &lookoutmetrics.DeleteAlertOutput{}

        mockClient.On("DeleteAlert", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteAlert(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteAnomalyDetector", func(t *testing.T) {
        input := &lookoutmetrics.DeleteAnomalyDetectorInput{}
        output := &lookoutmetrics.DeleteAnomalyDetectorOutput{}

        mockClient.On("DeleteAnomalyDetector", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteAnomalyDetector(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeAlert", func(t *testing.T) {
        input := &lookoutmetrics.DescribeAlertInput{}
        output := &lookoutmetrics.DescribeAlertOutput{}

        mockClient.On("DescribeAlert", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeAlert(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeAnomalyDetectionExecutions", func(t *testing.T) {
        input := &lookoutmetrics.DescribeAnomalyDetectionExecutionsInput{}
        output := &lookoutmetrics.DescribeAnomalyDetectionExecutionsOutput{}

        mockClient.On("DescribeAnomalyDetectionExecutions", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeAnomalyDetectionExecutions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeAnomalyDetector", func(t *testing.T) {
        input := &lookoutmetrics.DescribeAnomalyDetectorInput{}
        output := &lookoutmetrics.DescribeAnomalyDetectorOutput{}

        mockClient.On("DescribeAnomalyDetector", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeAnomalyDetector(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeMetricSet", func(t *testing.T) {
        input := &lookoutmetrics.DescribeMetricSetInput{}
        output := &lookoutmetrics.DescribeMetricSetOutput{}

        mockClient.On("DescribeMetricSet", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeMetricSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDetectMetricSetConfig", func(t *testing.T) {
        input := &lookoutmetrics.DetectMetricSetConfigInput{}
        output := &lookoutmetrics.DetectMetricSetConfigOutput{}

        mockClient.On("DetectMetricSetConfig", ctx, input).Return(output, nil)

        result, err := mockClient.DetectMetricSetConfig(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetAnomalyGroup", func(t *testing.T) {
        input := &lookoutmetrics.GetAnomalyGroupInput{}
        output := &lookoutmetrics.GetAnomalyGroupOutput{}

        mockClient.On("GetAnomalyGroup", ctx, input).Return(output, nil)

        result, err := mockClient.GetAnomalyGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDataQualityMetrics", func(t *testing.T) {
        input := &lookoutmetrics.GetDataQualityMetricsInput{}
        output := &lookoutmetrics.GetDataQualityMetricsOutput{}

        mockClient.On("GetDataQualityMetrics", ctx, input).Return(output, nil)

        result, err := mockClient.GetDataQualityMetrics(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetFeedback", func(t *testing.T) {
        input := &lookoutmetrics.GetFeedbackInput{}
        output := &lookoutmetrics.GetFeedbackOutput{}

        mockClient.On("GetFeedback", ctx, input).Return(output, nil)

        result, err := mockClient.GetFeedback(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSampleData", func(t *testing.T) {
        input := &lookoutmetrics.GetSampleDataInput{}
        output := &lookoutmetrics.GetSampleDataOutput{}

        mockClient.On("GetSampleData", ctx, input).Return(output, nil)

        result, err := mockClient.GetSampleData(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAlerts", func(t *testing.T) {
        input := &lookoutmetrics.ListAlertsInput{}
        output := &lookoutmetrics.ListAlertsOutput{}

        mockClient.On("ListAlerts", ctx, input).Return(output, nil)

        result, err := mockClient.ListAlerts(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAnomalyDetectors", func(t *testing.T) {
        input := &lookoutmetrics.ListAnomalyDetectorsInput{}
        output := &lookoutmetrics.ListAnomalyDetectorsOutput{}

        mockClient.On("ListAnomalyDetectors", ctx, input).Return(output, nil)

        result, err := mockClient.ListAnomalyDetectors(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAnomalyGroupRelatedMetrics", func(t *testing.T) {
        input := &lookoutmetrics.ListAnomalyGroupRelatedMetricsInput{}
        output := &lookoutmetrics.ListAnomalyGroupRelatedMetricsOutput{}

        mockClient.On("ListAnomalyGroupRelatedMetrics", ctx, input).Return(output, nil)

        result, err := mockClient.ListAnomalyGroupRelatedMetrics(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAnomalyGroupSummaries", func(t *testing.T) {
        input := &lookoutmetrics.ListAnomalyGroupSummariesInput{}
        output := &lookoutmetrics.ListAnomalyGroupSummariesOutput{}

        mockClient.On("ListAnomalyGroupSummaries", ctx, input).Return(output, nil)

        result, err := mockClient.ListAnomalyGroupSummaries(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAnomalyGroupTimeSeries", func(t *testing.T) {
        input := &lookoutmetrics.ListAnomalyGroupTimeSeriesInput{}
        output := &lookoutmetrics.ListAnomalyGroupTimeSeriesOutput{}

        mockClient.On("ListAnomalyGroupTimeSeries", ctx, input).Return(output, nil)

        result, err := mockClient.ListAnomalyGroupTimeSeries(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListMetricSets", func(t *testing.T) {
        input := &lookoutmetrics.ListMetricSetsInput{}
        output := &lookoutmetrics.ListMetricSetsOutput{}

        mockClient.On("ListMetricSets", ctx, input).Return(output, nil)

        result, err := mockClient.ListMetricSets(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &lookoutmetrics.ListTagsForResourceInput{}
        output := &lookoutmetrics.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutFeedback", func(t *testing.T) {
        input := &lookoutmetrics.PutFeedbackInput{}
        output := &lookoutmetrics.PutFeedbackOutput{}

        mockClient.On("PutFeedback", ctx, input).Return(output, nil)

        result, err := mockClient.PutFeedback(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &lookoutmetrics.TagResourceInput{}
        output := &lookoutmetrics.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &lookoutmetrics.UntagResourceInput{}
        output := &lookoutmetrics.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateAlert", func(t *testing.T) {
        input := &lookoutmetrics.UpdateAlertInput{}
        output := &lookoutmetrics.UpdateAlertOutput{}

        mockClient.On("UpdateAlert", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateAlert(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateAnomalyDetector", func(t *testing.T) {
        input := &lookoutmetrics.UpdateAnomalyDetectorInput{}
        output := &lookoutmetrics.UpdateAnomalyDetectorOutput{}

        mockClient.On("UpdateAnomalyDetector", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateAnomalyDetector(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateMetricSet", func(t *testing.T) {
        input := &lookoutmetrics.UpdateMetricSetInput{}
        output := &lookoutmetrics.UpdateMetricSetOutput{}

        mockClient.On("UpdateMetricSet", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateMetricSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
