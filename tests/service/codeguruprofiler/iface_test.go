package codeguruprofiler_test

// tests for the codeguruprofiler service interface for this ../../../service/codeguruprofiler/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/codeguruprofiler"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/codeguruprofiler/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/codeguruprofiler/codeguruprofiler_iface"
	"github.com/stretchr/testify/assert"
)

func TestCodeguruprofilerServiceCanBeMocked(t *testing.T) {
	var iface codeguruprofiler_iface.IClient
	iface = &codeguruprofiler.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := codeguruprofiler.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAddNotificationChannels", func(t *testing.T) {
        input := &codeguruprofiler.AddNotificationChannelsInput{}
        output := &codeguruprofiler.AddNotificationChannelsOutput{}

        mockClient.On("AddNotificationChannels", ctx, input).Return(output, nil)

        result, err := mockClient.AddNotificationChannels(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchGetFrameMetricData", func(t *testing.T) {
        input := &codeguruprofiler.BatchGetFrameMetricDataInput{}
        output := &codeguruprofiler.BatchGetFrameMetricDataOutput{}

        mockClient.On("BatchGetFrameMetricData", ctx, input).Return(output, nil)

        result, err := mockClient.BatchGetFrameMetricData(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestConfigureAgent", func(t *testing.T) {
        input := &codeguruprofiler.ConfigureAgentInput{}
        output := &codeguruprofiler.ConfigureAgentOutput{}

        mockClient.On("ConfigureAgent", ctx, input).Return(output, nil)

        result, err := mockClient.ConfigureAgent(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateProfilingGroup", func(t *testing.T) {
        input := &codeguruprofiler.CreateProfilingGroupInput{}
        output := &codeguruprofiler.CreateProfilingGroupOutput{}

        mockClient.On("CreateProfilingGroup", ctx, input).Return(output, nil)

        result, err := mockClient.CreateProfilingGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteProfilingGroup", func(t *testing.T) {
        input := &codeguruprofiler.DeleteProfilingGroupInput{}
        output := &codeguruprofiler.DeleteProfilingGroupOutput{}

        mockClient.On("DeleteProfilingGroup", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteProfilingGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeProfilingGroup", func(t *testing.T) {
        input := &codeguruprofiler.DescribeProfilingGroupInput{}
        output := &codeguruprofiler.DescribeProfilingGroupOutput{}

        mockClient.On("DescribeProfilingGroup", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeProfilingGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetFindingsReportAccountSummary", func(t *testing.T) {
        input := &codeguruprofiler.GetFindingsReportAccountSummaryInput{}
        output := &codeguruprofiler.GetFindingsReportAccountSummaryOutput{}

        mockClient.On("GetFindingsReportAccountSummary", ctx, input).Return(output, nil)

        result, err := mockClient.GetFindingsReportAccountSummary(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetNotificationConfiguration", func(t *testing.T) {
        input := &codeguruprofiler.GetNotificationConfigurationInput{}
        output := &codeguruprofiler.GetNotificationConfigurationOutput{}

        mockClient.On("GetNotificationConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.GetNotificationConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetPolicy", func(t *testing.T) {
        input := &codeguruprofiler.GetPolicyInput{}
        output := &codeguruprofiler.GetPolicyOutput{}

        mockClient.On("GetPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.GetPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetProfile", func(t *testing.T) {
        input := &codeguruprofiler.GetProfileInput{}
        output := &codeguruprofiler.GetProfileOutput{}

        mockClient.On("GetProfile", ctx, input).Return(output, nil)

        result, err := mockClient.GetProfile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetRecommendations", func(t *testing.T) {
        input := &codeguruprofiler.GetRecommendationsInput{}
        output := &codeguruprofiler.GetRecommendationsOutput{}

        mockClient.On("GetRecommendations", ctx, input).Return(output, nil)

        result, err := mockClient.GetRecommendations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListFindingsReports", func(t *testing.T) {
        input := &codeguruprofiler.ListFindingsReportsInput{}
        output := &codeguruprofiler.ListFindingsReportsOutput{}

        mockClient.On("ListFindingsReports", ctx, input).Return(output, nil)

        result, err := mockClient.ListFindingsReports(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListProfileTimes", func(t *testing.T) {
        input := &codeguruprofiler.ListProfileTimesInput{}
        output := &codeguruprofiler.ListProfileTimesOutput{}

        mockClient.On("ListProfileTimes", ctx, input).Return(output, nil)

        result, err := mockClient.ListProfileTimes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListProfilingGroups", func(t *testing.T) {
        input := &codeguruprofiler.ListProfilingGroupsInput{}
        output := &codeguruprofiler.ListProfilingGroupsOutput{}

        mockClient.On("ListProfilingGroups", ctx, input).Return(output, nil)

        result, err := mockClient.ListProfilingGroups(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &codeguruprofiler.ListTagsForResourceInput{}
        output := &codeguruprofiler.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPostAgentProfile", func(t *testing.T) {
        input := &codeguruprofiler.PostAgentProfileInput{}
        output := &codeguruprofiler.PostAgentProfileOutput{}

        mockClient.On("PostAgentProfile", ctx, input).Return(output, nil)

        result, err := mockClient.PostAgentProfile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutPermission", func(t *testing.T) {
        input := &codeguruprofiler.PutPermissionInput{}
        output := &codeguruprofiler.PutPermissionOutput{}

        mockClient.On("PutPermission", ctx, input).Return(output, nil)

        result, err := mockClient.PutPermission(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRemoveNotificationChannel", func(t *testing.T) {
        input := &codeguruprofiler.RemoveNotificationChannelInput{}
        output := &codeguruprofiler.RemoveNotificationChannelOutput{}

        mockClient.On("RemoveNotificationChannel", ctx, input).Return(output, nil)

        result, err := mockClient.RemoveNotificationChannel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRemovePermission", func(t *testing.T) {
        input := &codeguruprofiler.RemovePermissionInput{}
        output := &codeguruprofiler.RemovePermissionOutput{}

        mockClient.On("RemovePermission", ctx, input).Return(output, nil)

        result, err := mockClient.RemovePermission(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSubmitFeedback", func(t *testing.T) {
        input := &codeguruprofiler.SubmitFeedbackInput{}
        output := &codeguruprofiler.SubmitFeedbackOutput{}

        mockClient.On("SubmitFeedback", ctx, input).Return(output, nil)

        result, err := mockClient.SubmitFeedback(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &codeguruprofiler.TagResourceInput{}
        output := &codeguruprofiler.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &codeguruprofiler.UntagResourceInput{}
        output := &codeguruprofiler.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateProfilingGroup", func(t *testing.T) {
        input := &codeguruprofiler.UpdateProfilingGroupInput{}
        output := &codeguruprofiler.UpdateProfilingGroupOutput{}

        mockClient.On("UpdateProfilingGroup", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateProfilingGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
