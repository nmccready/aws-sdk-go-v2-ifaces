package devopsguru_test

// tests for the devopsguru service interface for this ../../../service/devopsguru/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/devopsguru"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/devopsguru/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/devopsguru/devopsguru_iface"
	"github.com/stretchr/testify/assert"
)

func TestDevopsguruServiceCanBeMocked(t *testing.T) {
	var iface devopsguru_iface.IClient
	iface = &devopsguru.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := devopsguru.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAddNotificationChannel", func(t *testing.T) {
        input := &devopsguru.AddNotificationChannelInput{}
        output := &devopsguru.AddNotificationChannelOutput{}

        mockClient.On("AddNotificationChannel", ctx, input).Return(output, nil)

        result, err := mockClient.AddNotificationChannel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteInsight", func(t *testing.T) {
        input := &devopsguru.DeleteInsightInput{}
        output := &devopsguru.DeleteInsightOutput{}

        mockClient.On("DeleteInsight", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteInsight(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeAccountHealth", func(t *testing.T) {
        input := &devopsguru.DescribeAccountHealthInput{}
        output := &devopsguru.DescribeAccountHealthOutput{}

        mockClient.On("DescribeAccountHealth", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeAccountHealth(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeAccountOverview", func(t *testing.T) {
        input := &devopsguru.DescribeAccountOverviewInput{}
        output := &devopsguru.DescribeAccountOverviewOutput{}

        mockClient.On("DescribeAccountOverview", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeAccountOverview(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeAnomaly", func(t *testing.T) {
        input := &devopsguru.DescribeAnomalyInput{}
        output := &devopsguru.DescribeAnomalyOutput{}

        mockClient.On("DescribeAnomaly", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeAnomaly(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeEventSourcesConfig", func(t *testing.T) {
        input := &devopsguru.DescribeEventSourcesConfigInput{}
        output := &devopsguru.DescribeEventSourcesConfigOutput{}

        mockClient.On("DescribeEventSourcesConfig", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeEventSourcesConfig(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeFeedback", func(t *testing.T) {
        input := &devopsguru.DescribeFeedbackInput{}
        output := &devopsguru.DescribeFeedbackOutput{}

        mockClient.On("DescribeFeedback", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeFeedback(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeInsight", func(t *testing.T) {
        input := &devopsguru.DescribeInsightInput{}
        output := &devopsguru.DescribeInsightOutput{}

        mockClient.On("DescribeInsight", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeInsight(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeOrganizationHealth", func(t *testing.T) {
        input := &devopsguru.DescribeOrganizationHealthInput{}
        output := &devopsguru.DescribeOrganizationHealthOutput{}

        mockClient.On("DescribeOrganizationHealth", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeOrganizationHealth(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeOrganizationOverview", func(t *testing.T) {
        input := &devopsguru.DescribeOrganizationOverviewInput{}
        output := &devopsguru.DescribeOrganizationOverviewOutput{}

        mockClient.On("DescribeOrganizationOverview", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeOrganizationOverview(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeOrganizationResourceCollectionHealth", func(t *testing.T) {
        input := &devopsguru.DescribeOrganizationResourceCollectionHealthInput{}
        output := &devopsguru.DescribeOrganizationResourceCollectionHealthOutput{}

        mockClient.On("DescribeOrganizationResourceCollectionHealth", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeOrganizationResourceCollectionHealth(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeResourceCollectionHealth", func(t *testing.T) {
        input := &devopsguru.DescribeResourceCollectionHealthInput{}
        output := &devopsguru.DescribeResourceCollectionHealthOutput{}

        mockClient.On("DescribeResourceCollectionHealth", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeResourceCollectionHealth(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeServiceIntegration", func(t *testing.T) {
        input := &devopsguru.DescribeServiceIntegrationInput{}
        output := &devopsguru.DescribeServiceIntegrationOutput{}

        mockClient.On("DescribeServiceIntegration", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeServiceIntegration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetCostEstimation", func(t *testing.T) {
        input := &devopsguru.GetCostEstimationInput{}
        output := &devopsguru.GetCostEstimationOutput{}

        mockClient.On("GetCostEstimation", ctx, input).Return(output, nil)

        result, err := mockClient.GetCostEstimation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetResourceCollection", func(t *testing.T) {
        input := &devopsguru.GetResourceCollectionInput{}
        output := &devopsguru.GetResourceCollectionOutput{}

        mockClient.On("GetResourceCollection", ctx, input).Return(output, nil)

        result, err := mockClient.GetResourceCollection(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAnomaliesForInsight", func(t *testing.T) {
        input := &devopsguru.ListAnomaliesForInsightInput{}
        output := &devopsguru.ListAnomaliesForInsightOutput{}

        mockClient.On("ListAnomaliesForInsight", ctx, input).Return(output, nil)

        result, err := mockClient.ListAnomaliesForInsight(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAnomalousLogGroups", func(t *testing.T) {
        input := &devopsguru.ListAnomalousLogGroupsInput{}
        output := &devopsguru.ListAnomalousLogGroupsOutput{}

        mockClient.On("ListAnomalousLogGroups", ctx, input).Return(output, nil)

        result, err := mockClient.ListAnomalousLogGroups(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListEvents", func(t *testing.T) {
        input := &devopsguru.ListEventsInput{}
        output := &devopsguru.ListEventsOutput{}

        mockClient.On("ListEvents", ctx, input).Return(output, nil)

        result, err := mockClient.ListEvents(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListInsights", func(t *testing.T) {
        input := &devopsguru.ListInsightsInput{}
        output := &devopsguru.ListInsightsOutput{}

        mockClient.On("ListInsights", ctx, input).Return(output, nil)

        result, err := mockClient.ListInsights(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListMonitoredResources", func(t *testing.T) {
        input := &devopsguru.ListMonitoredResourcesInput{}
        output := &devopsguru.ListMonitoredResourcesOutput{}

        mockClient.On("ListMonitoredResources", ctx, input).Return(output, nil)

        result, err := mockClient.ListMonitoredResources(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListNotificationChannels", func(t *testing.T) {
        input := &devopsguru.ListNotificationChannelsInput{}
        output := &devopsguru.ListNotificationChannelsOutput{}

        mockClient.On("ListNotificationChannels", ctx, input).Return(output, nil)

        result, err := mockClient.ListNotificationChannels(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListOrganizationInsights", func(t *testing.T) {
        input := &devopsguru.ListOrganizationInsightsInput{}
        output := &devopsguru.ListOrganizationInsightsOutput{}

        mockClient.On("ListOrganizationInsights", ctx, input).Return(output, nil)

        result, err := mockClient.ListOrganizationInsights(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListRecommendations", func(t *testing.T) {
        input := &devopsguru.ListRecommendationsInput{}
        output := &devopsguru.ListRecommendationsOutput{}

        mockClient.On("ListRecommendations", ctx, input).Return(output, nil)

        result, err := mockClient.ListRecommendations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutFeedback", func(t *testing.T) {
        input := &devopsguru.PutFeedbackInput{}
        output := &devopsguru.PutFeedbackOutput{}

        mockClient.On("PutFeedback", ctx, input).Return(output, nil)

        result, err := mockClient.PutFeedback(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRemoveNotificationChannel", func(t *testing.T) {
        input := &devopsguru.RemoveNotificationChannelInput{}
        output := &devopsguru.RemoveNotificationChannelOutput{}

        mockClient.On("RemoveNotificationChannel", ctx, input).Return(output, nil)

        result, err := mockClient.RemoveNotificationChannel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSearchInsights", func(t *testing.T) {
        input := &devopsguru.SearchInsightsInput{}
        output := &devopsguru.SearchInsightsOutput{}

        mockClient.On("SearchInsights", ctx, input).Return(output, nil)

        result, err := mockClient.SearchInsights(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSearchOrganizationInsights", func(t *testing.T) {
        input := &devopsguru.SearchOrganizationInsightsInput{}
        output := &devopsguru.SearchOrganizationInsightsOutput{}

        mockClient.On("SearchOrganizationInsights", ctx, input).Return(output, nil)

        result, err := mockClient.SearchOrganizationInsights(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartCostEstimation", func(t *testing.T) {
        input := &devopsguru.StartCostEstimationInput{}
        output := &devopsguru.StartCostEstimationOutput{}

        mockClient.On("StartCostEstimation", ctx, input).Return(output, nil)

        result, err := mockClient.StartCostEstimation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateEventSourcesConfig", func(t *testing.T) {
        input := &devopsguru.UpdateEventSourcesConfigInput{}
        output := &devopsguru.UpdateEventSourcesConfigOutput{}

        mockClient.On("UpdateEventSourcesConfig", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateEventSourcesConfig(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateResourceCollection", func(t *testing.T) {
        input := &devopsguru.UpdateResourceCollectionInput{}
        output := &devopsguru.UpdateResourceCollectionOutput{}

        mockClient.On("UpdateResourceCollection", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateResourceCollection(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateServiceIntegration", func(t *testing.T) {
        input := &devopsguru.UpdateServiceIntegrationInput{}
        output := &devopsguru.UpdateServiceIntegrationOutput{}

        mockClient.On("UpdateServiceIntegration", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateServiceIntegration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
