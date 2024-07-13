package trustedadvisor_test

// tests for the trustedadvisor service interface for this ../../../service/trustedadvisor/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/trustedadvisor"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/trustedadvisor/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/trustedadvisor/trustedadvisor_iface"
	"github.com/stretchr/testify/assert"
)

func TestTrustedadvisorServiceCanBeMocked(t *testing.T) {
	var iface trustedadvisor_iface.IClient
	iface = &trustedadvisor.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := trustedadvisor.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchUpdateRecommendationResourceExclusion", func(t *testing.T) {
        input := &trustedadvisor.BatchUpdateRecommendationResourceExclusionInput{}
        output := &trustedadvisor.BatchUpdateRecommendationResourceExclusionOutput{}

        mockClient.On("BatchUpdateRecommendationResourceExclusion", ctx, input).Return(output, nil)

        result, err := mockClient.BatchUpdateRecommendationResourceExclusion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetOrganizationRecommendation", func(t *testing.T) {
        input := &trustedadvisor.GetOrganizationRecommendationInput{}
        output := &trustedadvisor.GetOrganizationRecommendationOutput{}

        mockClient.On("GetOrganizationRecommendation", ctx, input).Return(output, nil)

        result, err := mockClient.GetOrganizationRecommendation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetRecommendation", func(t *testing.T) {
        input := &trustedadvisor.GetRecommendationInput{}
        output := &trustedadvisor.GetRecommendationOutput{}

        mockClient.On("GetRecommendation", ctx, input).Return(output, nil)

        result, err := mockClient.GetRecommendation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListChecks", func(t *testing.T) {
        input := &trustedadvisor.ListChecksInput{}
        output := &trustedadvisor.ListChecksOutput{}

        mockClient.On("ListChecks", ctx, input).Return(output, nil)

        result, err := mockClient.ListChecks(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListOrganizationRecommendationAccounts", func(t *testing.T) {
        input := &trustedadvisor.ListOrganizationRecommendationAccountsInput{}
        output := &trustedadvisor.ListOrganizationRecommendationAccountsOutput{}

        mockClient.On("ListOrganizationRecommendationAccounts", ctx, input).Return(output, nil)

        result, err := mockClient.ListOrganizationRecommendationAccounts(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListOrganizationRecommendationResources", func(t *testing.T) {
        input := &trustedadvisor.ListOrganizationRecommendationResourcesInput{}
        output := &trustedadvisor.ListOrganizationRecommendationResourcesOutput{}

        mockClient.On("ListOrganizationRecommendationResources", ctx, input).Return(output, nil)

        result, err := mockClient.ListOrganizationRecommendationResources(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListOrganizationRecommendations", func(t *testing.T) {
        input := &trustedadvisor.ListOrganizationRecommendationsInput{}
        output := &trustedadvisor.ListOrganizationRecommendationsOutput{}

        mockClient.On("ListOrganizationRecommendations", ctx, input).Return(output, nil)

        result, err := mockClient.ListOrganizationRecommendations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListRecommendationResources", func(t *testing.T) {
        input := &trustedadvisor.ListRecommendationResourcesInput{}
        output := &trustedadvisor.ListRecommendationResourcesOutput{}

        mockClient.On("ListRecommendationResources", ctx, input).Return(output, nil)

        result, err := mockClient.ListRecommendationResources(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListRecommendations", func(t *testing.T) {
        input := &trustedadvisor.ListRecommendationsInput{}
        output := &trustedadvisor.ListRecommendationsOutput{}

        mockClient.On("ListRecommendations", ctx, input).Return(output, nil)

        result, err := mockClient.ListRecommendations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateOrganizationRecommendationLifecycle", func(t *testing.T) {
        input := &trustedadvisor.UpdateOrganizationRecommendationLifecycleInput{}
        output := &trustedadvisor.UpdateOrganizationRecommendationLifecycleOutput{}

        mockClient.On("UpdateOrganizationRecommendationLifecycle", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateOrganizationRecommendationLifecycle(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateRecommendationLifecycle", func(t *testing.T) {
        input := &trustedadvisor.UpdateRecommendationLifecycleInput{}
        output := &trustedadvisor.UpdateRecommendationLifecycleOutput{}

        mockClient.On("UpdateRecommendationLifecycle", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateRecommendationLifecycle(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
