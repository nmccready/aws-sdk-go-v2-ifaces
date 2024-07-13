package billingconductor_test

// tests for the billingconductor service interface for this ../../../service/billingconductor/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/billingconductor"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/billingconductor/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/billingconductor/billingconductor_iface"
	"github.com/stretchr/testify/assert"
)

func TestBillingconductorServiceCanBeMocked(t *testing.T) {
	var iface billingconductor_iface.IClient
	iface = &billingconductor.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := billingconductor.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateAccounts", func(t *testing.T) {
        input := &billingconductor.AssociateAccountsInput{}
        output := &billingconductor.AssociateAccountsOutput{}

        mockClient.On("AssociateAccounts", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateAccounts(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociatePricingRules", func(t *testing.T) {
        input := &billingconductor.AssociatePricingRulesInput{}
        output := &billingconductor.AssociatePricingRulesOutput{}

        mockClient.On("AssociatePricingRules", ctx, input).Return(output, nil)

        result, err := mockClient.AssociatePricingRules(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchAssociateResourcesToCustomLineItem", func(t *testing.T) {
        input := &billingconductor.BatchAssociateResourcesToCustomLineItemInput{}
        output := &billingconductor.BatchAssociateResourcesToCustomLineItemOutput{}

        mockClient.On("BatchAssociateResourcesToCustomLineItem", ctx, input).Return(output, nil)

        result, err := mockClient.BatchAssociateResourcesToCustomLineItem(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchDisassociateResourcesFromCustomLineItem", func(t *testing.T) {
        input := &billingconductor.BatchDisassociateResourcesFromCustomLineItemInput{}
        output := &billingconductor.BatchDisassociateResourcesFromCustomLineItemOutput{}

        mockClient.On("BatchDisassociateResourcesFromCustomLineItem", ctx, input).Return(output, nil)

        result, err := mockClient.BatchDisassociateResourcesFromCustomLineItem(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateBillingGroup", func(t *testing.T) {
        input := &billingconductor.CreateBillingGroupInput{}
        output := &billingconductor.CreateBillingGroupOutput{}

        mockClient.On("CreateBillingGroup", ctx, input).Return(output, nil)

        result, err := mockClient.CreateBillingGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateCustomLineItem", func(t *testing.T) {
        input := &billingconductor.CreateCustomLineItemInput{}
        output := &billingconductor.CreateCustomLineItemOutput{}

        mockClient.On("CreateCustomLineItem", ctx, input).Return(output, nil)

        result, err := mockClient.CreateCustomLineItem(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreatePricingPlan", func(t *testing.T) {
        input := &billingconductor.CreatePricingPlanInput{}
        output := &billingconductor.CreatePricingPlanOutput{}

        mockClient.On("CreatePricingPlan", ctx, input).Return(output, nil)

        result, err := mockClient.CreatePricingPlan(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreatePricingRule", func(t *testing.T) {
        input := &billingconductor.CreatePricingRuleInput{}
        output := &billingconductor.CreatePricingRuleOutput{}

        mockClient.On("CreatePricingRule", ctx, input).Return(output, nil)

        result, err := mockClient.CreatePricingRule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteBillingGroup", func(t *testing.T) {
        input := &billingconductor.DeleteBillingGroupInput{}
        output := &billingconductor.DeleteBillingGroupOutput{}

        mockClient.On("DeleteBillingGroup", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteBillingGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteCustomLineItem", func(t *testing.T) {
        input := &billingconductor.DeleteCustomLineItemInput{}
        output := &billingconductor.DeleteCustomLineItemOutput{}

        mockClient.On("DeleteCustomLineItem", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteCustomLineItem(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeletePricingPlan", func(t *testing.T) {
        input := &billingconductor.DeletePricingPlanInput{}
        output := &billingconductor.DeletePricingPlanOutput{}

        mockClient.On("DeletePricingPlan", ctx, input).Return(output, nil)

        result, err := mockClient.DeletePricingPlan(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeletePricingRule", func(t *testing.T) {
        input := &billingconductor.DeletePricingRuleInput{}
        output := &billingconductor.DeletePricingRuleOutput{}

        mockClient.On("DeletePricingRule", ctx, input).Return(output, nil)

        result, err := mockClient.DeletePricingRule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateAccounts", func(t *testing.T) {
        input := &billingconductor.DisassociateAccountsInput{}
        output := &billingconductor.DisassociateAccountsOutput{}

        mockClient.On("DisassociateAccounts", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateAccounts(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociatePricingRules", func(t *testing.T) {
        input := &billingconductor.DisassociatePricingRulesInput{}
        output := &billingconductor.DisassociatePricingRulesOutput{}

        mockClient.On("DisassociatePricingRules", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociatePricingRules(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetBillingGroupCostReport", func(t *testing.T) {
        input := &billingconductor.GetBillingGroupCostReportInput{}
        output := &billingconductor.GetBillingGroupCostReportOutput{}

        mockClient.On("GetBillingGroupCostReport", ctx, input).Return(output, nil)

        result, err := mockClient.GetBillingGroupCostReport(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAccountAssociations", func(t *testing.T) {
        input := &billingconductor.ListAccountAssociationsInput{}
        output := &billingconductor.ListAccountAssociationsOutput{}

        mockClient.On("ListAccountAssociations", ctx, input).Return(output, nil)

        result, err := mockClient.ListAccountAssociations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListBillingGroupCostReports", func(t *testing.T) {
        input := &billingconductor.ListBillingGroupCostReportsInput{}
        output := &billingconductor.ListBillingGroupCostReportsOutput{}

        mockClient.On("ListBillingGroupCostReports", ctx, input).Return(output, nil)

        result, err := mockClient.ListBillingGroupCostReports(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListBillingGroups", func(t *testing.T) {
        input := &billingconductor.ListBillingGroupsInput{}
        output := &billingconductor.ListBillingGroupsOutput{}

        mockClient.On("ListBillingGroups", ctx, input).Return(output, nil)

        result, err := mockClient.ListBillingGroups(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListCustomLineItemVersions", func(t *testing.T) {
        input := &billingconductor.ListCustomLineItemVersionsInput{}
        output := &billingconductor.ListCustomLineItemVersionsOutput{}

        mockClient.On("ListCustomLineItemVersions", ctx, input).Return(output, nil)

        result, err := mockClient.ListCustomLineItemVersions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListCustomLineItems", func(t *testing.T) {
        input := &billingconductor.ListCustomLineItemsInput{}
        output := &billingconductor.ListCustomLineItemsOutput{}

        mockClient.On("ListCustomLineItems", ctx, input).Return(output, nil)

        result, err := mockClient.ListCustomLineItems(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListPricingPlans", func(t *testing.T) {
        input := &billingconductor.ListPricingPlansInput{}
        output := &billingconductor.ListPricingPlansOutput{}

        mockClient.On("ListPricingPlans", ctx, input).Return(output, nil)

        result, err := mockClient.ListPricingPlans(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListPricingPlansAssociatedWithPricingRule", func(t *testing.T) {
        input := &billingconductor.ListPricingPlansAssociatedWithPricingRuleInput{}
        output := &billingconductor.ListPricingPlansAssociatedWithPricingRuleOutput{}

        mockClient.On("ListPricingPlansAssociatedWithPricingRule", ctx, input).Return(output, nil)

        result, err := mockClient.ListPricingPlansAssociatedWithPricingRule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListPricingRules", func(t *testing.T) {
        input := &billingconductor.ListPricingRulesInput{}
        output := &billingconductor.ListPricingRulesOutput{}

        mockClient.On("ListPricingRules", ctx, input).Return(output, nil)

        result, err := mockClient.ListPricingRules(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListPricingRulesAssociatedToPricingPlan", func(t *testing.T) {
        input := &billingconductor.ListPricingRulesAssociatedToPricingPlanInput{}
        output := &billingconductor.ListPricingRulesAssociatedToPricingPlanOutput{}

        mockClient.On("ListPricingRulesAssociatedToPricingPlan", ctx, input).Return(output, nil)

        result, err := mockClient.ListPricingRulesAssociatedToPricingPlan(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListResourcesAssociatedToCustomLineItem", func(t *testing.T) {
        input := &billingconductor.ListResourcesAssociatedToCustomLineItemInput{}
        output := &billingconductor.ListResourcesAssociatedToCustomLineItemOutput{}

        mockClient.On("ListResourcesAssociatedToCustomLineItem", ctx, input).Return(output, nil)

        result, err := mockClient.ListResourcesAssociatedToCustomLineItem(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &billingconductor.ListTagsForResourceInput{}
        output := &billingconductor.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &billingconductor.TagResourceInput{}
        output := &billingconductor.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &billingconductor.UntagResourceInput{}
        output := &billingconductor.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateBillingGroup", func(t *testing.T) {
        input := &billingconductor.UpdateBillingGroupInput{}
        output := &billingconductor.UpdateBillingGroupOutput{}

        mockClient.On("UpdateBillingGroup", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateBillingGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateCustomLineItem", func(t *testing.T) {
        input := &billingconductor.UpdateCustomLineItemInput{}
        output := &billingconductor.UpdateCustomLineItemOutput{}

        mockClient.On("UpdateCustomLineItem", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateCustomLineItem(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdatePricingPlan", func(t *testing.T) {
        input := &billingconductor.UpdatePricingPlanInput{}
        output := &billingconductor.UpdatePricingPlanOutput{}

        mockClient.On("UpdatePricingPlan", ctx, input).Return(output, nil)

        result, err := mockClient.UpdatePricingPlan(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdatePricingRule", func(t *testing.T) {
        input := &billingconductor.UpdatePricingRuleInput{}
        output := &billingconductor.UpdatePricingRuleOutput{}

        mockClient.On("UpdatePricingRule", ctx, input).Return(output, nil)

        result, err := mockClient.UpdatePricingRule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
