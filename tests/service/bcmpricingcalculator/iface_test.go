// Code generated by ./scripts/gen_tests.sh DO NOT EDIT.
package bcmpricingcalculator_test

// tests for the bcmpricingcalculator service interface for 
// this ../../../service/bcmpricingcalculator/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/bcmpricingcalculator"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/bcmpricingcalculator/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/bcmpricingcalculator/bcmpricingcalculator_iface"
	"github.com/stretchr/testify/assert"
)

func TestBcmpricingcalculatorServiceCanBeMocked(t *testing.T) {
	var iface bcmpricingcalculator_iface.IClient
	iface = &bcmpricingcalculator.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := bcmpricingcalculator.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchCreateBillScenarioCommitmentModification", func(t *testing.T) {
        input := &bcmpricingcalculator.BatchCreateBillScenarioCommitmentModificationInput{}
        output := &bcmpricingcalculator.BatchCreateBillScenarioCommitmentModificationOutput{}

        mockClient.On("BatchCreateBillScenarioCommitmentModification", ctx, input).Return(output, nil)

        result, err := mockClient.BatchCreateBillScenarioCommitmentModification(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchCreateBillScenarioUsageModification", func(t *testing.T) {
        input := &bcmpricingcalculator.BatchCreateBillScenarioUsageModificationInput{}
        output := &bcmpricingcalculator.BatchCreateBillScenarioUsageModificationOutput{}

        mockClient.On("BatchCreateBillScenarioUsageModification", ctx, input).Return(output, nil)

        result, err := mockClient.BatchCreateBillScenarioUsageModification(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchCreateWorkloadEstimateUsage", func(t *testing.T) {
        input := &bcmpricingcalculator.BatchCreateWorkloadEstimateUsageInput{}
        output := &bcmpricingcalculator.BatchCreateWorkloadEstimateUsageOutput{}

        mockClient.On("BatchCreateWorkloadEstimateUsage", ctx, input).Return(output, nil)

        result, err := mockClient.BatchCreateWorkloadEstimateUsage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchDeleteBillScenarioCommitmentModification", func(t *testing.T) {
        input := &bcmpricingcalculator.BatchDeleteBillScenarioCommitmentModificationInput{}
        output := &bcmpricingcalculator.BatchDeleteBillScenarioCommitmentModificationOutput{}

        mockClient.On("BatchDeleteBillScenarioCommitmentModification", ctx, input).Return(output, nil)

        result, err := mockClient.BatchDeleteBillScenarioCommitmentModification(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchDeleteBillScenarioUsageModification", func(t *testing.T) {
        input := &bcmpricingcalculator.BatchDeleteBillScenarioUsageModificationInput{}
        output := &bcmpricingcalculator.BatchDeleteBillScenarioUsageModificationOutput{}

        mockClient.On("BatchDeleteBillScenarioUsageModification", ctx, input).Return(output, nil)

        result, err := mockClient.BatchDeleteBillScenarioUsageModification(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchDeleteWorkloadEstimateUsage", func(t *testing.T) {
        input := &bcmpricingcalculator.BatchDeleteWorkloadEstimateUsageInput{}
        output := &bcmpricingcalculator.BatchDeleteWorkloadEstimateUsageOutput{}

        mockClient.On("BatchDeleteWorkloadEstimateUsage", ctx, input).Return(output, nil)

        result, err := mockClient.BatchDeleteWorkloadEstimateUsage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchUpdateBillScenarioCommitmentModification", func(t *testing.T) {
        input := &bcmpricingcalculator.BatchUpdateBillScenarioCommitmentModificationInput{}
        output := &bcmpricingcalculator.BatchUpdateBillScenarioCommitmentModificationOutput{}

        mockClient.On("BatchUpdateBillScenarioCommitmentModification", ctx, input).Return(output, nil)

        result, err := mockClient.BatchUpdateBillScenarioCommitmentModification(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchUpdateBillScenarioUsageModification", func(t *testing.T) {
        input := &bcmpricingcalculator.BatchUpdateBillScenarioUsageModificationInput{}
        output := &bcmpricingcalculator.BatchUpdateBillScenarioUsageModificationOutput{}

        mockClient.On("BatchUpdateBillScenarioUsageModification", ctx, input).Return(output, nil)

        result, err := mockClient.BatchUpdateBillScenarioUsageModification(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchUpdateWorkloadEstimateUsage", func(t *testing.T) {
        input := &bcmpricingcalculator.BatchUpdateWorkloadEstimateUsageInput{}
        output := &bcmpricingcalculator.BatchUpdateWorkloadEstimateUsageOutput{}

        mockClient.On("BatchUpdateWorkloadEstimateUsage", ctx, input).Return(output, nil)

        result, err := mockClient.BatchUpdateWorkloadEstimateUsage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateBillEstimate", func(t *testing.T) {
        input := &bcmpricingcalculator.CreateBillEstimateInput{}
        output := &bcmpricingcalculator.CreateBillEstimateOutput{}

        mockClient.On("CreateBillEstimate", ctx, input).Return(output, nil)

        result, err := mockClient.CreateBillEstimate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateBillScenario", func(t *testing.T) {
        input := &bcmpricingcalculator.CreateBillScenarioInput{}
        output := &bcmpricingcalculator.CreateBillScenarioOutput{}

        mockClient.On("CreateBillScenario", ctx, input).Return(output, nil)

        result, err := mockClient.CreateBillScenario(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateWorkloadEstimate", func(t *testing.T) {
        input := &bcmpricingcalculator.CreateWorkloadEstimateInput{}
        output := &bcmpricingcalculator.CreateWorkloadEstimateOutput{}

        mockClient.On("CreateWorkloadEstimate", ctx, input).Return(output, nil)

        result, err := mockClient.CreateWorkloadEstimate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteBillEstimate", func(t *testing.T) {
        input := &bcmpricingcalculator.DeleteBillEstimateInput{}
        output := &bcmpricingcalculator.DeleteBillEstimateOutput{}

        mockClient.On("DeleteBillEstimate", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteBillEstimate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteBillScenario", func(t *testing.T) {
        input := &bcmpricingcalculator.DeleteBillScenarioInput{}
        output := &bcmpricingcalculator.DeleteBillScenarioOutput{}

        mockClient.On("DeleteBillScenario", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteBillScenario(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteWorkloadEstimate", func(t *testing.T) {
        input := &bcmpricingcalculator.DeleteWorkloadEstimateInput{}
        output := &bcmpricingcalculator.DeleteWorkloadEstimateOutput{}

        mockClient.On("DeleteWorkloadEstimate", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteWorkloadEstimate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetBillEstimate", func(t *testing.T) {
        input := &bcmpricingcalculator.GetBillEstimateInput{}
        output := &bcmpricingcalculator.GetBillEstimateOutput{}

        mockClient.On("GetBillEstimate", ctx, input).Return(output, nil)

        result, err := mockClient.GetBillEstimate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetBillScenario", func(t *testing.T) {
        input := &bcmpricingcalculator.GetBillScenarioInput{}
        output := &bcmpricingcalculator.GetBillScenarioOutput{}

        mockClient.On("GetBillScenario", ctx, input).Return(output, nil)

        result, err := mockClient.GetBillScenario(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetPreferences", func(t *testing.T) {
        input := &bcmpricingcalculator.GetPreferencesInput{}
        output := &bcmpricingcalculator.GetPreferencesOutput{}

        mockClient.On("GetPreferences", ctx, input).Return(output, nil)

        result, err := mockClient.GetPreferences(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetWorkloadEstimate", func(t *testing.T) {
        input := &bcmpricingcalculator.GetWorkloadEstimateInput{}
        output := &bcmpricingcalculator.GetWorkloadEstimateOutput{}

        mockClient.On("GetWorkloadEstimate", ctx, input).Return(output, nil)

        result, err := mockClient.GetWorkloadEstimate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListBillEstimateCommitments", func(t *testing.T) {
        input := &bcmpricingcalculator.ListBillEstimateCommitmentsInput{}
        output := &bcmpricingcalculator.ListBillEstimateCommitmentsOutput{}

        mockClient.On("ListBillEstimateCommitments", ctx, input).Return(output, nil)

        result, err := mockClient.ListBillEstimateCommitments(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListBillEstimateInputCommitmentModifications", func(t *testing.T) {
        input := &bcmpricingcalculator.ListBillEstimateInputCommitmentModificationsInput{}
        output := &bcmpricingcalculator.ListBillEstimateInputCommitmentModificationsOutput{}

        mockClient.On("ListBillEstimateInputCommitmentModifications", ctx, input).Return(output, nil)

        result, err := mockClient.ListBillEstimateInputCommitmentModifications(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListBillEstimateInputUsageModifications", func(t *testing.T) {
        input := &bcmpricingcalculator.ListBillEstimateInputUsageModificationsInput{}
        output := &bcmpricingcalculator.ListBillEstimateInputUsageModificationsOutput{}

        mockClient.On("ListBillEstimateInputUsageModifications", ctx, input).Return(output, nil)

        result, err := mockClient.ListBillEstimateInputUsageModifications(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListBillEstimateLineItems", func(t *testing.T) {
        input := &bcmpricingcalculator.ListBillEstimateLineItemsInput{}
        output := &bcmpricingcalculator.ListBillEstimateLineItemsOutput{}

        mockClient.On("ListBillEstimateLineItems", ctx, input).Return(output, nil)

        result, err := mockClient.ListBillEstimateLineItems(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListBillEstimates", func(t *testing.T) {
        input := &bcmpricingcalculator.ListBillEstimatesInput{}
        output := &bcmpricingcalculator.ListBillEstimatesOutput{}

        mockClient.On("ListBillEstimates", ctx, input).Return(output, nil)

        result, err := mockClient.ListBillEstimates(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListBillScenarioCommitmentModifications", func(t *testing.T) {
        input := &bcmpricingcalculator.ListBillScenarioCommitmentModificationsInput{}
        output := &bcmpricingcalculator.ListBillScenarioCommitmentModificationsOutput{}

        mockClient.On("ListBillScenarioCommitmentModifications", ctx, input).Return(output, nil)

        result, err := mockClient.ListBillScenarioCommitmentModifications(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListBillScenarioUsageModifications", func(t *testing.T) {
        input := &bcmpricingcalculator.ListBillScenarioUsageModificationsInput{}
        output := &bcmpricingcalculator.ListBillScenarioUsageModificationsOutput{}

        mockClient.On("ListBillScenarioUsageModifications", ctx, input).Return(output, nil)

        result, err := mockClient.ListBillScenarioUsageModifications(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListBillScenarios", func(t *testing.T) {
        input := &bcmpricingcalculator.ListBillScenariosInput{}
        output := &bcmpricingcalculator.ListBillScenariosOutput{}

        mockClient.On("ListBillScenarios", ctx, input).Return(output, nil)

        result, err := mockClient.ListBillScenarios(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &bcmpricingcalculator.ListTagsForResourceInput{}
        output := &bcmpricingcalculator.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListWorkloadEstimateUsage", func(t *testing.T) {
        input := &bcmpricingcalculator.ListWorkloadEstimateUsageInput{}
        output := &bcmpricingcalculator.ListWorkloadEstimateUsageOutput{}

        mockClient.On("ListWorkloadEstimateUsage", ctx, input).Return(output, nil)

        result, err := mockClient.ListWorkloadEstimateUsage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListWorkloadEstimates", func(t *testing.T) {
        input := &bcmpricingcalculator.ListWorkloadEstimatesInput{}
        output := &bcmpricingcalculator.ListWorkloadEstimatesOutput{}

        mockClient.On("ListWorkloadEstimates", ctx, input).Return(output, nil)

        result, err := mockClient.ListWorkloadEstimates(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &bcmpricingcalculator.TagResourceInput{}
        output := &bcmpricingcalculator.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &bcmpricingcalculator.UntagResourceInput{}
        output := &bcmpricingcalculator.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateBillEstimate", func(t *testing.T) {
        input := &bcmpricingcalculator.UpdateBillEstimateInput{}
        output := &bcmpricingcalculator.UpdateBillEstimateOutput{}

        mockClient.On("UpdateBillEstimate", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateBillEstimate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateBillScenario", func(t *testing.T) {
        input := &bcmpricingcalculator.UpdateBillScenarioInput{}
        output := &bcmpricingcalculator.UpdateBillScenarioOutput{}

        mockClient.On("UpdateBillScenario", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateBillScenario(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdatePreferences", func(t *testing.T) {
        input := &bcmpricingcalculator.UpdatePreferencesInput{}
        output := &bcmpricingcalculator.UpdatePreferencesOutput{}

        mockClient.On("UpdatePreferences", ctx, input).Return(output, nil)

        result, err := mockClient.UpdatePreferences(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateWorkloadEstimate", func(t *testing.T) {
        input := &bcmpricingcalculator.UpdateWorkloadEstimateInput{}
        output := &bcmpricingcalculator.UpdateWorkloadEstimateOutput{}

        mockClient.On("UpdateWorkloadEstimate", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateWorkloadEstimate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}