package savingsplans_test

// tests for the savingsplans service interface for this ../../../service/savingsplans/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/savingsplans"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/savingsplans/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/savingsplans/savingsplans_iface"
	"github.com/stretchr/testify/assert"
)

func TestSavingsplansServiceCanBeMocked(t *testing.T) {
	var iface savingsplans_iface.IClient
	iface = &savingsplans.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := savingsplans.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateSavingsPlan", func(t *testing.T) {
        input := &savingsplans.CreateSavingsPlanInput{}
        output := &savingsplans.CreateSavingsPlanOutput{}

        mockClient.On("CreateSavingsPlan", ctx, input).Return(output, nil)

        result, err := mockClient.CreateSavingsPlan(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteQueuedSavingsPlan", func(t *testing.T) {
        input := &savingsplans.DeleteQueuedSavingsPlanInput{}
        output := &savingsplans.DeleteQueuedSavingsPlanOutput{}

        mockClient.On("DeleteQueuedSavingsPlan", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteQueuedSavingsPlan(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeSavingsPlanRates", func(t *testing.T) {
        input := &savingsplans.DescribeSavingsPlanRatesInput{}
        output := &savingsplans.DescribeSavingsPlanRatesOutput{}

        mockClient.On("DescribeSavingsPlanRates", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeSavingsPlanRates(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeSavingsPlans", func(t *testing.T) {
        input := &savingsplans.DescribeSavingsPlansInput{}
        output := &savingsplans.DescribeSavingsPlansOutput{}

        mockClient.On("DescribeSavingsPlans", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeSavingsPlans(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeSavingsPlansOfferingRates", func(t *testing.T) {
        input := &savingsplans.DescribeSavingsPlansOfferingRatesInput{}
        output := &savingsplans.DescribeSavingsPlansOfferingRatesOutput{}

        mockClient.On("DescribeSavingsPlansOfferingRates", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeSavingsPlansOfferingRates(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeSavingsPlansOfferings", func(t *testing.T) {
        input := &savingsplans.DescribeSavingsPlansOfferingsInput{}
        output := &savingsplans.DescribeSavingsPlansOfferingsOutput{}

        mockClient.On("DescribeSavingsPlansOfferings", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeSavingsPlansOfferings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &savingsplans.ListTagsForResourceInput{}
        output := &savingsplans.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestReturnSavingsPlan", func(t *testing.T) {
        input := &savingsplans.ReturnSavingsPlanInput{}
        output := &savingsplans.ReturnSavingsPlanOutput{}

        mockClient.On("ReturnSavingsPlan", ctx, input).Return(output, nil)

        result, err := mockClient.ReturnSavingsPlan(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &savingsplans.TagResourceInput{}
        output := &savingsplans.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &savingsplans.UntagResourceInput{}
        output := &savingsplans.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
