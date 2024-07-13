package costoptimizationhub_test

// tests for the costoptimizationhub service interface for this ../../../service/costoptimizationhub/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/costoptimizationhub"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/costoptimizationhub/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/costoptimizationhub/costoptimizationhub_iface"
	"github.com/stretchr/testify/assert"
)

func TestCostoptimizationhubServiceCanBeMocked(t *testing.T) {
	var iface costoptimizationhub_iface.IClient
	iface = &costoptimizationhub.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := costoptimizationhub.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetPreferences", func(t *testing.T) {
        input := &costoptimizationhub.GetPreferencesInput{}
        output := &costoptimizationhub.GetPreferencesOutput{}

        mockClient.On("GetPreferences", ctx, input).Return(output, nil)

        result, err := mockClient.GetPreferences(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetRecommendation", func(t *testing.T) {
        input := &costoptimizationhub.GetRecommendationInput{}
        output := &costoptimizationhub.GetRecommendationOutput{}

        mockClient.On("GetRecommendation", ctx, input).Return(output, nil)

        result, err := mockClient.GetRecommendation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListEnrollmentStatuses", func(t *testing.T) {
        input := &costoptimizationhub.ListEnrollmentStatusesInput{}
        output := &costoptimizationhub.ListEnrollmentStatusesOutput{}

        mockClient.On("ListEnrollmentStatuses", ctx, input).Return(output, nil)

        result, err := mockClient.ListEnrollmentStatuses(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListRecommendationSummaries", func(t *testing.T) {
        input := &costoptimizationhub.ListRecommendationSummariesInput{}
        output := &costoptimizationhub.ListRecommendationSummariesOutput{}

        mockClient.On("ListRecommendationSummaries", ctx, input).Return(output, nil)

        result, err := mockClient.ListRecommendationSummaries(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListRecommendations", func(t *testing.T) {
        input := &costoptimizationhub.ListRecommendationsInput{}
        output := &costoptimizationhub.ListRecommendationsOutput{}

        mockClient.On("ListRecommendations", ctx, input).Return(output, nil)

        result, err := mockClient.ListRecommendations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateEnrollmentStatus", func(t *testing.T) {
        input := &costoptimizationhub.UpdateEnrollmentStatusInput{}
        output := &costoptimizationhub.UpdateEnrollmentStatusOutput{}

        mockClient.On("UpdateEnrollmentStatus", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateEnrollmentStatus(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdatePreferences", func(t *testing.T) {
        input := &costoptimizationhub.UpdatePreferencesInput{}
        output := &costoptimizationhub.UpdatePreferencesOutput{}

        mockClient.On("UpdatePreferences", ctx, input).Return(output, nil)

        result, err := mockClient.UpdatePreferences(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
