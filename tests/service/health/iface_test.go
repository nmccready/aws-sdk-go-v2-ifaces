// Code generated by ./scripts/gen_tests.sh DO NOT EDIT.
package health_test

// tests for the health service interface for 
// this ../../../service/health/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/health"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/health/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/health/health_iface"
	"github.com/stretchr/testify/assert"
)

func TestHealthServiceCanBeMocked(t *testing.T) {
	var iface health_iface.IClient
	iface = &health.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := health.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeAffectedAccountsForOrganization", func(t *testing.T) {
        input := &health.DescribeAffectedAccountsForOrganizationInput{}
        output := &health.DescribeAffectedAccountsForOrganizationOutput{}

        mockClient.On("DescribeAffectedAccountsForOrganization", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeAffectedAccountsForOrganization(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeAffectedEntities", func(t *testing.T) {
        input := &health.DescribeAffectedEntitiesInput{}
        output := &health.DescribeAffectedEntitiesOutput{}

        mockClient.On("DescribeAffectedEntities", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeAffectedEntities(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeAffectedEntitiesForOrganization", func(t *testing.T) {
        input := &health.DescribeAffectedEntitiesForOrganizationInput{}
        output := &health.DescribeAffectedEntitiesForOrganizationOutput{}

        mockClient.On("DescribeAffectedEntitiesForOrganization", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeAffectedEntitiesForOrganization(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeEntityAggregates", func(t *testing.T) {
        input := &health.DescribeEntityAggregatesInput{}
        output := &health.DescribeEntityAggregatesOutput{}

        mockClient.On("DescribeEntityAggregates", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeEntityAggregates(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeEntityAggregatesForOrganization", func(t *testing.T) {
        input := &health.DescribeEntityAggregatesForOrganizationInput{}
        output := &health.DescribeEntityAggregatesForOrganizationOutput{}

        mockClient.On("DescribeEntityAggregatesForOrganization", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeEntityAggregatesForOrganization(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeEventAggregates", func(t *testing.T) {
        input := &health.DescribeEventAggregatesInput{}
        output := &health.DescribeEventAggregatesOutput{}

        mockClient.On("DescribeEventAggregates", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeEventAggregates(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeEventDetails", func(t *testing.T) {
        input := &health.DescribeEventDetailsInput{}
        output := &health.DescribeEventDetailsOutput{}

        mockClient.On("DescribeEventDetails", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeEventDetails(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeEventDetailsForOrganization", func(t *testing.T) {
        input := &health.DescribeEventDetailsForOrganizationInput{}
        output := &health.DescribeEventDetailsForOrganizationOutput{}

        mockClient.On("DescribeEventDetailsForOrganization", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeEventDetailsForOrganization(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeEventTypes", func(t *testing.T) {
        input := &health.DescribeEventTypesInput{}
        output := &health.DescribeEventTypesOutput{}

        mockClient.On("DescribeEventTypes", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeEventTypes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeEvents", func(t *testing.T) {
        input := &health.DescribeEventsInput{}
        output := &health.DescribeEventsOutput{}

        mockClient.On("DescribeEvents", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeEvents(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeEventsForOrganization", func(t *testing.T) {
        input := &health.DescribeEventsForOrganizationInput{}
        output := &health.DescribeEventsForOrganizationOutput{}

        mockClient.On("DescribeEventsForOrganization", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeEventsForOrganization(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeHealthServiceStatusForOrganization", func(t *testing.T) {
        input := &health.DescribeHealthServiceStatusForOrganizationInput{}
        output := &health.DescribeHealthServiceStatusForOrganizationOutput{}

        mockClient.On("DescribeHealthServiceStatusForOrganization", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeHealthServiceStatusForOrganization(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisableHealthServiceAccessForOrganization", func(t *testing.T) {
        input := &health.DisableHealthServiceAccessForOrganizationInput{}
        output := &health.DisableHealthServiceAccessForOrganizationOutput{}

        mockClient.On("DisableHealthServiceAccessForOrganization", ctx, input).Return(output, nil)

        result, err := mockClient.DisableHealthServiceAccessForOrganization(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestEnableHealthServiceAccessForOrganization", func(t *testing.T) {
        input := &health.EnableHealthServiceAccessForOrganizationInput{}
        output := &health.EnableHealthServiceAccessForOrganizationOutput{}

        mockClient.On("EnableHealthServiceAccessForOrganization", ctx, input).Return(output, nil)

        result, err := mockClient.EnableHealthServiceAccessForOrganization(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
