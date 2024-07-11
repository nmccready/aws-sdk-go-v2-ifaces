
package billingconductor

import (
    "github.com/aws/aws-sdk-go-v2/service/billingconductor"
)

// IClient defines the interface for billingconductor
type IClient interface {
 Options() Options 
 AssociateAccounts(ctx context.Context, params *AssociateAccountsInput, optFns ...func(*Options)) (*AssociateAccountsOutput, error) 
 AssociatePricingRules(ctx context.Context, params *AssociatePricingRulesInput, optFns ...func(*Options)) (*AssociatePricingRulesOutput, error) 
 BatchAssociateResourcesToCustomLineItem(ctx context.Context, params *BatchAssociateResourcesToCustomLineItemInput, optFns ...func(*Options)) (*BatchAssociateResourcesToCustomLineItemOutput, error) 
 BatchDisassociateResourcesFromCustomLineItem(ctx context.Context, params *BatchDisassociateResourcesFromCustomLineItemInput, optFns ...func(*Options)) (*BatchDisassociateResourcesFromCustomLineItemOutput, error) 
 CreateBillingGroup(ctx context.Context, params *CreateBillingGroupInput, optFns ...func(*Options)) (*CreateBillingGroupOutput, error) 
 CreateCustomLineItem(ctx context.Context, params *CreateCustomLineItemInput, optFns ...func(*Options)) (*CreateCustomLineItemOutput, error) 
 CreatePricingPlan(ctx context.Context, params *CreatePricingPlanInput, optFns ...func(*Options)) (*CreatePricingPlanOutput, error) 
 CreatePricingRule(ctx context.Context, params *CreatePricingRuleInput, optFns ...func(*Options)) (*CreatePricingRuleOutput, error) 
 DeleteBillingGroup(ctx context.Context, params *DeleteBillingGroupInput, optFns ...func(*Options)) (*DeleteBillingGroupOutput, error) 
 DeleteCustomLineItem(ctx context.Context, params *DeleteCustomLineItemInput, optFns ...func(*Options)) (*DeleteCustomLineItemOutput, error) 
 DeletePricingPlan(ctx context.Context, params *DeletePricingPlanInput, optFns ...func(*Options)) (*DeletePricingPlanOutput, error) 
 DeletePricingRule(ctx context.Context, params *DeletePricingRuleInput, optFns ...func(*Options)) (*DeletePricingRuleOutput, error) 
 DisassociateAccounts(ctx context.Context, params *DisassociateAccountsInput, optFns ...func(*Options)) (*DisassociateAccountsOutput, error) 
 DisassociatePricingRules(ctx context.Context, params *DisassociatePricingRulesInput, optFns ...func(*Options)) (*DisassociatePricingRulesOutput, error) 
 GetBillingGroupCostReport(ctx context.Context, params *GetBillingGroupCostReportInput, optFns ...func(*Options)) (*GetBillingGroupCostReportOutput, error) 
 ListAccountAssociations(ctx context.Context, params *ListAccountAssociationsInput, optFns ...func(*Options)) (*ListAccountAssociationsOutput, error) 
 ListBillingGroupCostReports(ctx context.Context, params *ListBillingGroupCostReportsInput, optFns ...func(*Options)) (*ListBillingGroupCostReportsOutput, error) 
 ListBillingGroups(ctx context.Context, params *ListBillingGroupsInput, optFns ...func(*Options)) (*ListBillingGroupsOutput, error) 
 ListCustomLineItemVersions(ctx context.Context, params *ListCustomLineItemVersionsInput, optFns ...func(*Options)) (*ListCustomLineItemVersionsOutput, error) 
 ListCustomLineItems(ctx context.Context, params *ListCustomLineItemsInput, optFns ...func(*Options)) (*ListCustomLineItemsOutput, error) 
 ListPricingPlans(ctx context.Context, params *ListPricingPlansInput, optFns ...func(*Options)) (*ListPricingPlansOutput, error) 
 ListPricingPlansAssociatedWithPricingRule(ctx context.Context, params *ListPricingPlansAssociatedWithPricingRuleInput, optFns ...func(*Options)) (*ListPricingPlansAssociatedWithPricingRuleOutput, error) 
 ListPricingRules(ctx context.Context, params *ListPricingRulesInput, optFns ...func(*Options)) (*ListPricingRulesOutput, error) 
 ListPricingRulesAssociatedToPricingPlan(ctx context.Context, params *ListPricingRulesAssociatedToPricingPlanInput, optFns ...func(*Options)) (*ListPricingRulesAssociatedToPricingPlanOutput, error) 
 ListResourcesAssociatedToCustomLineItem(ctx context.Context, params *ListResourcesAssociatedToCustomLineItemInput, optFns ...func(*Options)) (*ListResourcesAssociatedToCustomLineItemOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateBillingGroup(ctx context.Context, params *UpdateBillingGroupInput, optFns ...func(*Options)) (*UpdateBillingGroupOutput, error) 
 UpdateCustomLineItem(ctx context.Context, params *UpdateCustomLineItemInput, optFns ...func(*Options)) (*UpdateCustomLineItemOutput, error) 
 UpdatePricingPlan(ctx context.Context, params *UpdatePricingPlanInput, optFns ...func(*Options)) (*UpdatePricingPlanOutput, error) 
 UpdatePricingRule(ctx context.Context, params *UpdatePricingRuleInput, optFns ...func(*Options)) (*UpdatePricingRuleOutput, error) 
}
