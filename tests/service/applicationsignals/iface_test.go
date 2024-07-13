package applicationsignals_test

// tests for the applicationsignals service interface for this ../../../service/applicationsignals/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/applicationsignals"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/applicationsignals/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/applicationsignals/applicationsignals_iface"
	"github.com/stretchr/testify/assert"
)

func TestApplicationsignalsServiceCanBeMocked(t *testing.T) {
	var iface applicationsignals_iface.IClient
	iface = &applicationsignals.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := applicationsignals.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchGetServiceLevelObjectiveBudgetReport", func(t *testing.T) {
        input := &applicationsignals.BatchGetServiceLevelObjectiveBudgetReportInput{}
        output := &applicationsignals.BatchGetServiceLevelObjectiveBudgetReportOutput{}

        mockClient.On("BatchGetServiceLevelObjectiveBudgetReport", ctx, input).Return(output, nil)

        result, err := mockClient.BatchGetServiceLevelObjectiveBudgetReport(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateServiceLevelObjective", func(t *testing.T) {
        input := &applicationsignals.CreateServiceLevelObjectiveInput{}
        output := &applicationsignals.CreateServiceLevelObjectiveOutput{}

        mockClient.On("CreateServiceLevelObjective", ctx, input).Return(output, nil)

        result, err := mockClient.CreateServiceLevelObjective(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteServiceLevelObjective", func(t *testing.T) {
        input := &applicationsignals.DeleteServiceLevelObjectiveInput{}
        output := &applicationsignals.DeleteServiceLevelObjectiveOutput{}

        mockClient.On("DeleteServiceLevelObjective", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteServiceLevelObjective(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetService", func(t *testing.T) {
        input := &applicationsignals.GetServiceInput{}
        output := &applicationsignals.GetServiceOutput{}

        mockClient.On("GetService", ctx, input).Return(output, nil)

        result, err := mockClient.GetService(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetServiceLevelObjective", func(t *testing.T) {
        input := &applicationsignals.GetServiceLevelObjectiveInput{}
        output := &applicationsignals.GetServiceLevelObjectiveOutput{}

        mockClient.On("GetServiceLevelObjective", ctx, input).Return(output, nil)

        result, err := mockClient.GetServiceLevelObjective(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListServiceDependencies", func(t *testing.T) {
        input := &applicationsignals.ListServiceDependenciesInput{}
        output := &applicationsignals.ListServiceDependenciesOutput{}

        mockClient.On("ListServiceDependencies", ctx, input).Return(output, nil)

        result, err := mockClient.ListServiceDependencies(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListServiceDependents", func(t *testing.T) {
        input := &applicationsignals.ListServiceDependentsInput{}
        output := &applicationsignals.ListServiceDependentsOutput{}

        mockClient.On("ListServiceDependents", ctx, input).Return(output, nil)

        result, err := mockClient.ListServiceDependents(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListServiceLevelObjectives", func(t *testing.T) {
        input := &applicationsignals.ListServiceLevelObjectivesInput{}
        output := &applicationsignals.ListServiceLevelObjectivesOutput{}

        mockClient.On("ListServiceLevelObjectives", ctx, input).Return(output, nil)

        result, err := mockClient.ListServiceLevelObjectives(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListServiceOperations", func(t *testing.T) {
        input := &applicationsignals.ListServiceOperationsInput{}
        output := &applicationsignals.ListServiceOperationsOutput{}

        mockClient.On("ListServiceOperations", ctx, input).Return(output, nil)

        result, err := mockClient.ListServiceOperations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListServices", func(t *testing.T) {
        input := &applicationsignals.ListServicesInput{}
        output := &applicationsignals.ListServicesOutput{}

        mockClient.On("ListServices", ctx, input).Return(output, nil)

        result, err := mockClient.ListServices(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &applicationsignals.ListTagsForResourceInput{}
        output := &applicationsignals.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartDiscovery", func(t *testing.T) {
        input := &applicationsignals.StartDiscoveryInput{}
        output := &applicationsignals.StartDiscoveryOutput{}

        mockClient.On("StartDiscovery", ctx, input).Return(output, nil)

        result, err := mockClient.StartDiscovery(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &applicationsignals.TagResourceInput{}
        output := &applicationsignals.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &applicationsignals.UntagResourceInput{}
        output := &applicationsignals.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateServiceLevelObjective", func(t *testing.T) {
        input := &applicationsignals.UpdateServiceLevelObjectiveInput{}
        output := &applicationsignals.UpdateServiceLevelObjectiveOutput{}

        mockClient.On("UpdateServiceLevelObjective", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateServiceLevelObjective(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
