package xray_test

// tests for the xray service interface for this ../../../service/xray/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/xray"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/xray/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/xray/xray_iface"
	"github.com/stretchr/testify/assert"
)

func TestXrayServiceCanBeMocked(t *testing.T) {
	var iface xray_iface.IClient
	iface = &xray.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := xray.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchGetTraces", func(t *testing.T) {
        input := &xray.BatchGetTracesInput{}
        output := &xray.BatchGetTracesOutput{}

        mockClient.On("BatchGetTraces", ctx, input).Return(output, nil)

        result, err := mockClient.BatchGetTraces(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateGroup", func(t *testing.T) {
        input := &xray.CreateGroupInput{}
        output := &xray.CreateGroupOutput{}

        mockClient.On("CreateGroup", ctx, input).Return(output, nil)

        result, err := mockClient.CreateGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateSamplingRule", func(t *testing.T) {
        input := &xray.CreateSamplingRuleInput{}
        output := &xray.CreateSamplingRuleOutput{}

        mockClient.On("CreateSamplingRule", ctx, input).Return(output, nil)

        result, err := mockClient.CreateSamplingRule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteGroup", func(t *testing.T) {
        input := &xray.DeleteGroupInput{}
        output := &xray.DeleteGroupOutput{}

        mockClient.On("DeleteGroup", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteResourcePolicy", func(t *testing.T) {
        input := &xray.DeleteResourcePolicyInput{}
        output := &xray.DeleteResourcePolicyOutput{}

        mockClient.On("DeleteResourcePolicy", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteResourcePolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteSamplingRule", func(t *testing.T) {
        input := &xray.DeleteSamplingRuleInput{}
        output := &xray.DeleteSamplingRuleOutput{}

        mockClient.On("DeleteSamplingRule", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteSamplingRule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetEncryptionConfig", func(t *testing.T) {
        input := &xray.GetEncryptionConfigInput{}
        output := &xray.GetEncryptionConfigOutput{}

        mockClient.On("GetEncryptionConfig", ctx, input).Return(output, nil)

        result, err := mockClient.GetEncryptionConfig(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetGroup", func(t *testing.T) {
        input := &xray.GetGroupInput{}
        output := &xray.GetGroupOutput{}

        mockClient.On("GetGroup", ctx, input).Return(output, nil)

        result, err := mockClient.GetGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetGroups", func(t *testing.T) {
        input := &xray.GetGroupsInput{}
        output := &xray.GetGroupsOutput{}

        mockClient.On("GetGroups", ctx, input).Return(output, nil)

        result, err := mockClient.GetGroups(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetInsight", func(t *testing.T) {
        input := &xray.GetInsightInput{}
        output := &xray.GetInsightOutput{}

        mockClient.On("GetInsight", ctx, input).Return(output, nil)

        result, err := mockClient.GetInsight(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetInsightEvents", func(t *testing.T) {
        input := &xray.GetInsightEventsInput{}
        output := &xray.GetInsightEventsOutput{}

        mockClient.On("GetInsightEvents", ctx, input).Return(output, nil)

        result, err := mockClient.GetInsightEvents(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetInsightImpactGraph", func(t *testing.T) {
        input := &xray.GetInsightImpactGraphInput{}
        output := &xray.GetInsightImpactGraphOutput{}

        mockClient.On("GetInsightImpactGraph", ctx, input).Return(output, nil)

        result, err := mockClient.GetInsightImpactGraph(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetInsightSummaries", func(t *testing.T) {
        input := &xray.GetInsightSummariesInput{}
        output := &xray.GetInsightSummariesOutput{}

        mockClient.On("GetInsightSummaries", ctx, input).Return(output, nil)

        result, err := mockClient.GetInsightSummaries(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSamplingRules", func(t *testing.T) {
        input := &xray.GetSamplingRulesInput{}
        output := &xray.GetSamplingRulesOutput{}

        mockClient.On("GetSamplingRules", ctx, input).Return(output, nil)

        result, err := mockClient.GetSamplingRules(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSamplingStatisticSummaries", func(t *testing.T) {
        input := &xray.GetSamplingStatisticSummariesInput{}
        output := &xray.GetSamplingStatisticSummariesOutput{}

        mockClient.On("GetSamplingStatisticSummaries", ctx, input).Return(output, nil)

        result, err := mockClient.GetSamplingStatisticSummaries(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSamplingTargets", func(t *testing.T) {
        input := &xray.GetSamplingTargetsInput{}
        output := &xray.GetSamplingTargetsOutput{}

        mockClient.On("GetSamplingTargets", ctx, input).Return(output, nil)

        result, err := mockClient.GetSamplingTargets(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetServiceGraph", func(t *testing.T) {
        input := &xray.GetServiceGraphInput{}
        output := &xray.GetServiceGraphOutput{}

        mockClient.On("GetServiceGraph", ctx, input).Return(output, nil)

        result, err := mockClient.GetServiceGraph(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetTimeSeriesServiceStatistics", func(t *testing.T) {
        input := &xray.GetTimeSeriesServiceStatisticsInput{}
        output := &xray.GetTimeSeriesServiceStatisticsOutput{}

        mockClient.On("GetTimeSeriesServiceStatistics", ctx, input).Return(output, nil)

        result, err := mockClient.GetTimeSeriesServiceStatistics(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetTraceGraph", func(t *testing.T) {
        input := &xray.GetTraceGraphInput{}
        output := &xray.GetTraceGraphOutput{}

        mockClient.On("GetTraceGraph", ctx, input).Return(output, nil)

        result, err := mockClient.GetTraceGraph(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetTraceSummaries", func(t *testing.T) {
        input := &xray.GetTraceSummariesInput{}
        output := &xray.GetTraceSummariesOutput{}

        mockClient.On("GetTraceSummaries", ctx, input).Return(output, nil)

        result, err := mockClient.GetTraceSummaries(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListResourcePolicies", func(t *testing.T) {
        input := &xray.ListResourcePoliciesInput{}
        output := &xray.ListResourcePoliciesOutput{}

        mockClient.On("ListResourcePolicies", ctx, input).Return(output, nil)

        result, err := mockClient.ListResourcePolicies(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &xray.ListTagsForResourceInput{}
        output := &xray.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutEncryptionConfig", func(t *testing.T) {
        input := &xray.PutEncryptionConfigInput{}
        output := &xray.PutEncryptionConfigOutput{}

        mockClient.On("PutEncryptionConfig", ctx, input).Return(output, nil)

        result, err := mockClient.PutEncryptionConfig(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutResourcePolicy", func(t *testing.T) {
        input := &xray.PutResourcePolicyInput{}
        output := &xray.PutResourcePolicyOutput{}

        mockClient.On("PutResourcePolicy", ctx, input).Return(output, nil)

        result, err := mockClient.PutResourcePolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutTelemetryRecords", func(t *testing.T) {
        input := &xray.PutTelemetryRecordsInput{}
        output := &xray.PutTelemetryRecordsOutput{}

        mockClient.On("PutTelemetryRecords", ctx, input).Return(output, nil)

        result, err := mockClient.PutTelemetryRecords(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutTraceSegments", func(t *testing.T) {
        input := &xray.PutTraceSegmentsInput{}
        output := &xray.PutTraceSegmentsOutput{}

        mockClient.On("PutTraceSegments", ctx, input).Return(output, nil)

        result, err := mockClient.PutTraceSegments(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &xray.TagResourceInput{}
        output := &xray.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &xray.UntagResourceInput{}
        output := &xray.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateGroup", func(t *testing.T) {
        input := &xray.UpdateGroupInput{}
        output := &xray.UpdateGroupOutput{}

        mockClient.On("UpdateGroup", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateSamplingRule", func(t *testing.T) {
        input := &xray.UpdateSamplingRuleInput{}
        output := &xray.UpdateSamplingRuleOutput{}

        mockClient.On("UpdateSamplingRule", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateSamplingRule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
