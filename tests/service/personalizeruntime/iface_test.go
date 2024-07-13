package personalizeruntime_test

// tests for the personalizeruntime service interface for this ../../../service/personalizeruntime/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/personalizeruntime"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/personalizeruntime/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/personalizeruntime/personalizeruntime_iface"
	"github.com/stretchr/testify/assert"
)

func TestPersonalizeruntimeServiceCanBeMocked(t *testing.T) {
	var iface personalizeruntime_iface.IClient
	iface = &personalizeruntime.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := personalizeruntime.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetActionRecommendations", func(t *testing.T) {
        input := &personalizeruntime.GetActionRecommendationsInput{}
        output := &personalizeruntime.GetActionRecommendationsOutput{}

        mockClient.On("GetActionRecommendations", ctx, input).Return(output, nil)

        result, err := mockClient.GetActionRecommendations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetPersonalizedRanking", func(t *testing.T) {
        input := &personalizeruntime.GetPersonalizedRankingInput{}
        output := &personalizeruntime.GetPersonalizedRankingOutput{}

        mockClient.On("GetPersonalizedRanking", ctx, input).Return(output, nil)

        result, err := mockClient.GetPersonalizedRanking(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetRecommendations", func(t *testing.T) {
        input := &personalizeruntime.GetRecommendationsInput{}
        output := &personalizeruntime.GetRecommendationsOutput{}

        mockClient.On("GetRecommendations", ctx, input).Return(output, nil)

        result, err := mockClient.GetRecommendations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
