package cloudwatch_test

// tests for the cloudwatch service interface for this ../../../service/cloudwatch/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/cloudwatch/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/cloudwatch/cloudwatch_iface"
	"github.com/stretchr/testify/assert"
)

func TestCloudwatchServiceCanBeMocked(t *testing.T) {
	var iface cloudwatch_iface.IClient
	iface = &cloudwatch.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := cloudwatch.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteAlarms", func(t *testing.T) {
        input := &cloudwatch.DeleteAlarmsInput{}
        output := &cloudwatch.DeleteAlarmsOutput{}

        mockClient.On("DeleteAlarms", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteAlarms(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteAnomalyDetector", func(t *testing.T) {
        input := &cloudwatch.DeleteAnomalyDetectorInput{}
        output := &cloudwatch.DeleteAnomalyDetectorOutput{}

        mockClient.On("DeleteAnomalyDetector", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteAnomalyDetector(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteDashboards", func(t *testing.T) {
        input := &cloudwatch.DeleteDashboardsInput{}
        output := &cloudwatch.DeleteDashboardsOutput{}

        mockClient.On("DeleteDashboards", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteDashboards(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteInsightRules", func(t *testing.T) {
        input := &cloudwatch.DeleteInsightRulesInput{}
        output := &cloudwatch.DeleteInsightRulesOutput{}

        mockClient.On("DeleteInsightRules", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteInsightRules(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteMetricStream", func(t *testing.T) {
        input := &cloudwatch.DeleteMetricStreamInput{}
        output := &cloudwatch.DeleteMetricStreamOutput{}

        mockClient.On("DeleteMetricStream", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteMetricStream(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeAlarmHistory", func(t *testing.T) {
        input := &cloudwatch.DescribeAlarmHistoryInput{}
        output := &cloudwatch.DescribeAlarmHistoryOutput{}

        mockClient.On("DescribeAlarmHistory", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeAlarmHistory(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeAlarms", func(t *testing.T) {
        input := &cloudwatch.DescribeAlarmsInput{}
        output := &cloudwatch.DescribeAlarmsOutput{}

        mockClient.On("DescribeAlarms", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeAlarms(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeAlarmsForMetric", func(t *testing.T) {
        input := &cloudwatch.DescribeAlarmsForMetricInput{}
        output := &cloudwatch.DescribeAlarmsForMetricOutput{}

        mockClient.On("DescribeAlarmsForMetric", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeAlarmsForMetric(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeAnomalyDetectors", func(t *testing.T) {
        input := &cloudwatch.DescribeAnomalyDetectorsInput{}
        output := &cloudwatch.DescribeAnomalyDetectorsOutput{}

        mockClient.On("DescribeAnomalyDetectors", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeAnomalyDetectors(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeInsightRules", func(t *testing.T) {
        input := &cloudwatch.DescribeInsightRulesInput{}
        output := &cloudwatch.DescribeInsightRulesOutput{}

        mockClient.On("DescribeInsightRules", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeInsightRules(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisableAlarmActions", func(t *testing.T) {
        input := &cloudwatch.DisableAlarmActionsInput{}
        output := &cloudwatch.DisableAlarmActionsOutput{}

        mockClient.On("DisableAlarmActions", ctx, input).Return(output, nil)

        result, err := mockClient.DisableAlarmActions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisableInsightRules", func(t *testing.T) {
        input := &cloudwatch.DisableInsightRulesInput{}
        output := &cloudwatch.DisableInsightRulesOutput{}

        mockClient.On("DisableInsightRules", ctx, input).Return(output, nil)

        result, err := mockClient.DisableInsightRules(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestEnableAlarmActions", func(t *testing.T) {
        input := &cloudwatch.EnableAlarmActionsInput{}
        output := &cloudwatch.EnableAlarmActionsOutput{}

        mockClient.On("EnableAlarmActions", ctx, input).Return(output, nil)

        result, err := mockClient.EnableAlarmActions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestEnableInsightRules", func(t *testing.T) {
        input := &cloudwatch.EnableInsightRulesInput{}
        output := &cloudwatch.EnableInsightRulesOutput{}

        mockClient.On("EnableInsightRules", ctx, input).Return(output, nil)

        result, err := mockClient.EnableInsightRules(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDashboard", func(t *testing.T) {
        input := &cloudwatch.GetDashboardInput{}
        output := &cloudwatch.GetDashboardOutput{}

        mockClient.On("GetDashboard", ctx, input).Return(output, nil)

        result, err := mockClient.GetDashboard(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetInsightRuleReport", func(t *testing.T) {
        input := &cloudwatch.GetInsightRuleReportInput{}
        output := &cloudwatch.GetInsightRuleReportOutput{}

        mockClient.On("GetInsightRuleReport", ctx, input).Return(output, nil)

        result, err := mockClient.GetInsightRuleReport(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetMetricData", func(t *testing.T) {
        input := &cloudwatch.GetMetricDataInput{}
        output := &cloudwatch.GetMetricDataOutput{}

        mockClient.On("GetMetricData", ctx, input).Return(output, nil)

        result, err := mockClient.GetMetricData(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetMetricStatistics", func(t *testing.T) {
        input := &cloudwatch.GetMetricStatisticsInput{}
        output := &cloudwatch.GetMetricStatisticsOutput{}

        mockClient.On("GetMetricStatistics", ctx, input).Return(output, nil)

        result, err := mockClient.GetMetricStatistics(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetMetricStream", func(t *testing.T) {
        input := &cloudwatch.GetMetricStreamInput{}
        output := &cloudwatch.GetMetricStreamOutput{}

        mockClient.On("GetMetricStream", ctx, input).Return(output, nil)

        result, err := mockClient.GetMetricStream(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetMetricWidgetImage", func(t *testing.T) {
        input := &cloudwatch.GetMetricWidgetImageInput{}
        output := &cloudwatch.GetMetricWidgetImageOutput{}

        mockClient.On("GetMetricWidgetImage", ctx, input).Return(output, nil)

        result, err := mockClient.GetMetricWidgetImage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDashboards", func(t *testing.T) {
        input := &cloudwatch.ListDashboardsInput{}
        output := &cloudwatch.ListDashboardsOutput{}

        mockClient.On("ListDashboards", ctx, input).Return(output, nil)

        result, err := mockClient.ListDashboards(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListManagedInsightRules", func(t *testing.T) {
        input := &cloudwatch.ListManagedInsightRulesInput{}
        output := &cloudwatch.ListManagedInsightRulesOutput{}

        mockClient.On("ListManagedInsightRules", ctx, input).Return(output, nil)

        result, err := mockClient.ListManagedInsightRules(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListMetricStreams", func(t *testing.T) {
        input := &cloudwatch.ListMetricStreamsInput{}
        output := &cloudwatch.ListMetricStreamsOutput{}

        mockClient.On("ListMetricStreams", ctx, input).Return(output, nil)

        result, err := mockClient.ListMetricStreams(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListMetrics", func(t *testing.T) {
        input := &cloudwatch.ListMetricsInput{}
        output := &cloudwatch.ListMetricsOutput{}

        mockClient.On("ListMetrics", ctx, input).Return(output, nil)

        result, err := mockClient.ListMetrics(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &cloudwatch.ListTagsForResourceInput{}
        output := &cloudwatch.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutAnomalyDetector", func(t *testing.T) {
        input := &cloudwatch.PutAnomalyDetectorInput{}
        output := &cloudwatch.PutAnomalyDetectorOutput{}

        mockClient.On("PutAnomalyDetector", ctx, input).Return(output, nil)

        result, err := mockClient.PutAnomalyDetector(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutCompositeAlarm", func(t *testing.T) {
        input := &cloudwatch.PutCompositeAlarmInput{}
        output := &cloudwatch.PutCompositeAlarmOutput{}

        mockClient.On("PutCompositeAlarm", ctx, input).Return(output, nil)

        result, err := mockClient.PutCompositeAlarm(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutDashboard", func(t *testing.T) {
        input := &cloudwatch.PutDashboardInput{}
        output := &cloudwatch.PutDashboardOutput{}

        mockClient.On("PutDashboard", ctx, input).Return(output, nil)

        result, err := mockClient.PutDashboard(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutInsightRule", func(t *testing.T) {
        input := &cloudwatch.PutInsightRuleInput{}
        output := &cloudwatch.PutInsightRuleOutput{}

        mockClient.On("PutInsightRule", ctx, input).Return(output, nil)

        result, err := mockClient.PutInsightRule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutManagedInsightRules", func(t *testing.T) {
        input := &cloudwatch.PutManagedInsightRulesInput{}
        output := &cloudwatch.PutManagedInsightRulesOutput{}

        mockClient.On("PutManagedInsightRules", ctx, input).Return(output, nil)

        result, err := mockClient.PutManagedInsightRules(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutMetricAlarm", func(t *testing.T) {
        input := &cloudwatch.PutMetricAlarmInput{}
        output := &cloudwatch.PutMetricAlarmOutput{}

        mockClient.On("PutMetricAlarm", ctx, input).Return(output, nil)

        result, err := mockClient.PutMetricAlarm(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutMetricData", func(t *testing.T) {
        input := &cloudwatch.PutMetricDataInput{}
        output := &cloudwatch.PutMetricDataOutput{}

        mockClient.On("PutMetricData", ctx, input).Return(output, nil)

        result, err := mockClient.PutMetricData(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutMetricStream", func(t *testing.T) {
        input := &cloudwatch.PutMetricStreamInput{}
        output := &cloudwatch.PutMetricStreamOutput{}

        mockClient.On("PutMetricStream", ctx, input).Return(output, nil)

        result, err := mockClient.PutMetricStream(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSetAlarmState", func(t *testing.T) {
        input := &cloudwatch.SetAlarmStateInput{}
        output := &cloudwatch.SetAlarmStateOutput{}

        mockClient.On("SetAlarmState", ctx, input).Return(output, nil)

        result, err := mockClient.SetAlarmState(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartMetricStreams", func(t *testing.T) {
        input := &cloudwatch.StartMetricStreamsInput{}
        output := &cloudwatch.StartMetricStreamsOutput{}

        mockClient.On("StartMetricStreams", ctx, input).Return(output, nil)

        result, err := mockClient.StartMetricStreams(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopMetricStreams", func(t *testing.T) {
        input := &cloudwatch.StopMetricStreamsInput{}
        output := &cloudwatch.StopMetricStreamsOutput{}

        mockClient.On("StopMetricStreams", ctx, input).Return(output, nil)

        result, err := mockClient.StopMetricStreams(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &cloudwatch.TagResourceInput{}
        output := &cloudwatch.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &cloudwatch.UntagResourceInput{}
        output := &cloudwatch.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
