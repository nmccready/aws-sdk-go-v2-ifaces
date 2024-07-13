package wafregional_test

// tests for the wafregional service interface for this ../../../service/wafregional/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/wafregional"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/wafregional/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/wafregional/wafregional_iface"
	"github.com/stretchr/testify/assert"
)

func TestWafregionalServiceCanBeMocked(t *testing.T) {
	var iface wafregional_iface.IClient
	iface = &wafregional.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := wafregional.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateWebACL", func(t *testing.T) {
        input := &wafregional.AssociateWebACLInput{}
        output := &wafregional.AssociateWebACLOutput{}

        mockClient.On("AssociateWebACL", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateWebACL(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateByteMatchSet", func(t *testing.T) {
        input := &wafregional.CreateByteMatchSetInput{}
        output := &wafregional.CreateByteMatchSetOutput{}

        mockClient.On("CreateByteMatchSet", ctx, input).Return(output, nil)

        result, err := mockClient.CreateByteMatchSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateGeoMatchSet", func(t *testing.T) {
        input := &wafregional.CreateGeoMatchSetInput{}
        output := &wafregional.CreateGeoMatchSetOutput{}

        mockClient.On("CreateGeoMatchSet", ctx, input).Return(output, nil)

        result, err := mockClient.CreateGeoMatchSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateIPSet", func(t *testing.T) {
        input := &wafregional.CreateIPSetInput{}
        output := &wafregional.CreateIPSetOutput{}

        mockClient.On("CreateIPSet", ctx, input).Return(output, nil)

        result, err := mockClient.CreateIPSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateRateBasedRule", func(t *testing.T) {
        input := &wafregional.CreateRateBasedRuleInput{}
        output := &wafregional.CreateRateBasedRuleOutput{}

        mockClient.On("CreateRateBasedRule", ctx, input).Return(output, nil)

        result, err := mockClient.CreateRateBasedRule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateRegexMatchSet", func(t *testing.T) {
        input := &wafregional.CreateRegexMatchSetInput{}
        output := &wafregional.CreateRegexMatchSetOutput{}

        mockClient.On("CreateRegexMatchSet", ctx, input).Return(output, nil)

        result, err := mockClient.CreateRegexMatchSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateRegexPatternSet", func(t *testing.T) {
        input := &wafregional.CreateRegexPatternSetInput{}
        output := &wafregional.CreateRegexPatternSetOutput{}

        mockClient.On("CreateRegexPatternSet", ctx, input).Return(output, nil)

        result, err := mockClient.CreateRegexPatternSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateRule", func(t *testing.T) {
        input := &wafregional.CreateRuleInput{}
        output := &wafregional.CreateRuleOutput{}

        mockClient.On("CreateRule", ctx, input).Return(output, nil)

        result, err := mockClient.CreateRule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateRuleGroup", func(t *testing.T) {
        input := &wafregional.CreateRuleGroupInput{}
        output := &wafregional.CreateRuleGroupOutput{}

        mockClient.On("CreateRuleGroup", ctx, input).Return(output, nil)

        result, err := mockClient.CreateRuleGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateSizeConstraintSet", func(t *testing.T) {
        input := &wafregional.CreateSizeConstraintSetInput{}
        output := &wafregional.CreateSizeConstraintSetOutput{}

        mockClient.On("CreateSizeConstraintSet", ctx, input).Return(output, nil)

        result, err := mockClient.CreateSizeConstraintSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateSqlInjectionMatchSet", func(t *testing.T) {
        input := &wafregional.CreateSqlInjectionMatchSetInput{}
        output := &wafregional.CreateSqlInjectionMatchSetOutput{}

        mockClient.On("CreateSqlInjectionMatchSet", ctx, input).Return(output, nil)

        result, err := mockClient.CreateSqlInjectionMatchSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateWebACL", func(t *testing.T) {
        input := &wafregional.CreateWebACLInput{}
        output := &wafregional.CreateWebACLOutput{}

        mockClient.On("CreateWebACL", ctx, input).Return(output, nil)

        result, err := mockClient.CreateWebACL(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateWebACLMigrationStack", func(t *testing.T) {
        input := &wafregional.CreateWebACLMigrationStackInput{}
        output := &wafregional.CreateWebACLMigrationStackOutput{}

        mockClient.On("CreateWebACLMigrationStack", ctx, input).Return(output, nil)

        result, err := mockClient.CreateWebACLMigrationStack(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateXssMatchSet", func(t *testing.T) {
        input := &wafregional.CreateXssMatchSetInput{}
        output := &wafregional.CreateXssMatchSetOutput{}

        mockClient.On("CreateXssMatchSet", ctx, input).Return(output, nil)

        result, err := mockClient.CreateXssMatchSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteByteMatchSet", func(t *testing.T) {
        input := &wafregional.DeleteByteMatchSetInput{}
        output := &wafregional.DeleteByteMatchSetOutput{}

        mockClient.On("DeleteByteMatchSet", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteByteMatchSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteGeoMatchSet", func(t *testing.T) {
        input := &wafregional.DeleteGeoMatchSetInput{}
        output := &wafregional.DeleteGeoMatchSetOutput{}

        mockClient.On("DeleteGeoMatchSet", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteGeoMatchSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteIPSet", func(t *testing.T) {
        input := &wafregional.DeleteIPSetInput{}
        output := &wafregional.DeleteIPSetOutput{}

        mockClient.On("DeleteIPSet", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteIPSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteLoggingConfiguration", func(t *testing.T) {
        input := &wafregional.DeleteLoggingConfigurationInput{}
        output := &wafregional.DeleteLoggingConfigurationOutput{}

        mockClient.On("DeleteLoggingConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteLoggingConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeletePermissionPolicy", func(t *testing.T) {
        input := &wafregional.DeletePermissionPolicyInput{}
        output := &wafregional.DeletePermissionPolicyOutput{}

        mockClient.On("DeletePermissionPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.DeletePermissionPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteRateBasedRule", func(t *testing.T) {
        input := &wafregional.DeleteRateBasedRuleInput{}
        output := &wafregional.DeleteRateBasedRuleOutput{}

        mockClient.On("DeleteRateBasedRule", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteRateBasedRule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteRegexMatchSet", func(t *testing.T) {
        input := &wafregional.DeleteRegexMatchSetInput{}
        output := &wafregional.DeleteRegexMatchSetOutput{}

        mockClient.On("DeleteRegexMatchSet", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteRegexMatchSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteRegexPatternSet", func(t *testing.T) {
        input := &wafregional.DeleteRegexPatternSetInput{}
        output := &wafregional.DeleteRegexPatternSetOutput{}

        mockClient.On("DeleteRegexPatternSet", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteRegexPatternSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteRule", func(t *testing.T) {
        input := &wafregional.DeleteRuleInput{}
        output := &wafregional.DeleteRuleOutput{}

        mockClient.On("DeleteRule", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteRule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteRuleGroup", func(t *testing.T) {
        input := &wafregional.DeleteRuleGroupInput{}
        output := &wafregional.DeleteRuleGroupOutput{}

        mockClient.On("DeleteRuleGroup", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteRuleGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteSizeConstraintSet", func(t *testing.T) {
        input := &wafregional.DeleteSizeConstraintSetInput{}
        output := &wafregional.DeleteSizeConstraintSetOutput{}

        mockClient.On("DeleteSizeConstraintSet", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteSizeConstraintSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteSqlInjectionMatchSet", func(t *testing.T) {
        input := &wafregional.DeleteSqlInjectionMatchSetInput{}
        output := &wafregional.DeleteSqlInjectionMatchSetOutput{}

        mockClient.On("DeleteSqlInjectionMatchSet", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteSqlInjectionMatchSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteWebACL", func(t *testing.T) {
        input := &wafregional.DeleteWebACLInput{}
        output := &wafregional.DeleteWebACLOutput{}

        mockClient.On("DeleteWebACL", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteWebACL(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteXssMatchSet", func(t *testing.T) {
        input := &wafregional.DeleteXssMatchSetInput{}
        output := &wafregional.DeleteXssMatchSetOutput{}

        mockClient.On("DeleteXssMatchSet", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteXssMatchSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateWebACL", func(t *testing.T) {
        input := &wafregional.DisassociateWebACLInput{}
        output := &wafregional.DisassociateWebACLOutput{}

        mockClient.On("DisassociateWebACL", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateWebACL(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetByteMatchSet", func(t *testing.T) {
        input := &wafregional.GetByteMatchSetInput{}
        output := &wafregional.GetByteMatchSetOutput{}

        mockClient.On("GetByteMatchSet", ctx, input).Return(output, nil)

        result, err := mockClient.GetByteMatchSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetChangeToken", func(t *testing.T) {
        input := &wafregional.GetChangeTokenInput{}
        output := &wafregional.GetChangeTokenOutput{}

        mockClient.On("GetChangeToken", ctx, input).Return(output, nil)

        result, err := mockClient.GetChangeToken(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetChangeTokenStatus", func(t *testing.T) {
        input := &wafregional.GetChangeTokenStatusInput{}
        output := &wafregional.GetChangeTokenStatusOutput{}

        mockClient.On("GetChangeTokenStatus", ctx, input).Return(output, nil)

        result, err := mockClient.GetChangeTokenStatus(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetGeoMatchSet", func(t *testing.T) {
        input := &wafregional.GetGeoMatchSetInput{}
        output := &wafregional.GetGeoMatchSetOutput{}

        mockClient.On("GetGeoMatchSet", ctx, input).Return(output, nil)

        result, err := mockClient.GetGeoMatchSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetIPSet", func(t *testing.T) {
        input := &wafregional.GetIPSetInput{}
        output := &wafregional.GetIPSetOutput{}

        mockClient.On("GetIPSet", ctx, input).Return(output, nil)

        result, err := mockClient.GetIPSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetLoggingConfiguration", func(t *testing.T) {
        input := &wafregional.GetLoggingConfigurationInput{}
        output := &wafregional.GetLoggingConfigurationOutput{}

        mockClient.On("GetLoggingConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.GetLoggingConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetPermissionPolicy", func(t *testing.T) {
        input := &wafregional.GetPermissionPolicyInput{}
        output := &wafregional.GetPermissionPolicyOutput{}

        mockClient.On("GetPermissionPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.GetPermissionPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetRateBasedRule", func(t *testing.T) {
        input := &wafregional.GetRateBasedRuleInput{}
        output := &wafregional.GetRateBasedRuleOutput{}

        mockClient.On("GetRateBasedRule", ctx, input).Return(output, nil)

        result, err := mockClient.GetRateBasedRule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetRateBasedRuleManagedKeys", func(t *testing.T) {
        input := &wafregional.GetRateBasedRuleManagedKeysInput{}
        output := &wafregional.GetRateBasedRuleManagedKeysOutput{}

        mockClient.On("GetRateBasedRuleManagedKeys", ctx, input).Return(output, nil)

        result, err := mockClient.GetRateBasedRuleManagedKeys(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetRegexMatchSet", func(t *testing.T) {
        input := &wafregional.GetRegexMatchSetInput{}
        output := &wafregional.GetRegexMatchSetOutput{}

        mockClient.On("GetRegexMatchSet", ctx, input).Return(output, nil)

        result, err := mockClient.GetRegexMatchSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetRegexPatternSet", func(t *testing.T) {
        input := &wafregional.GetRegexPatternSetInput{}
        output := &wafregional.GetRegexPatternSetOutput{}

        mockClient.On("GetRegexPatternSet", ctx, input).Return(output, nil)

        result, err := mockClient.GetRegexPatternSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetRule", func(t *testing.T) {
        input := &wafregional.GetRuleInput{}
        output := &wafregional.GetRuleOutput{}

        mockClient.On("GetRule", ctx, input).Return(output, nil)

        result, err := mockClient.GetRule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetRuleGroup", func(t *testing.T) {
        input := &wafregional.GetRuleGroupInput{}
        output := &wafregional.GetRuleGroupOutput{}

        mockClient.On("GetRuleGroup", ctx, input).Return(output, nil)

        result, err := mockClient.GetRuleGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSampledRequests", func(t *testing.T) {
        input := &wafregional.GetSampledRequestsInput{}
        output := &wafregional.GetSampledRequestsOutput{}

        mockClient.On("GetSampledRequests", ctx, input).Return(output, nil)

        result, err := mockClient.GetSampledRequests(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSizeConstraintSet", func(t *testing.T) {
        input := &wafregional.GetSizeConstraintSetInput{}
        output := &wafregional.GetSizeConstraintSetOutput{}

        mockClient.On("GetSizeConstraintSet", ctx, input).Return(output, nil)

        result, err := mockClient.GetSizeConstraintSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSqlInjectionMatchSet", func(t *testing.T) {
        input := &wafregional.GetSqlInjectionMatchSetInput{}
        output := &wafregional.GetSqlInjectionMatchSetOutput{}

        mockClient.On("GetSqlInjectionMatchSet", ctx, input).Return(output, nil)

        result, err := mockClient.GetSqlInjectionMatchSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetWebACL", func(t *testing.T) {
        input := &wafregional.GetWebACLInput{}
        output := &wafregional.GetWebACLOutput{}

        mockClient.On("GetWebACL", ctx, input).Return(output, nil)

        result, err := mockClient.GetWebACL(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetWebACLForResource", func(t *testing.T) {
        input := &wafregional.GetWebACLForResourceInput{}
        output := &wafregional.GetWebACLForResourceOutput{}

        mockClient.On("GetWebACLForResource", ctx, input).Return(output, nil)

        result, err := mockClient.GetWebACLForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetXssMatchSet", func(t *testing.T) {
        input := &wafregional.GetXssMatchSetInput{}
        output := &wafregional.GetXssMatchSetOutput{}

        mockClient.On("GetXssMatchSet", ctx, input).Return(output, nil)

        result, err := mockClient.GetXssMatchSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListActivatedRulesInRuleGroup", func(t *testing.T) {
        input := &wafregional.ListActivatedRulesInRuleGroupInput{}
        output := &wafregional.ListActivatedRulesInRuleGroupOutput{}

        mockClient.On("ListActivatedRulesInRuleGroup", ctx, input).Return(output, nil)

        result, err := mockClient.ListActivatedRulesInRuleGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListByteMatchSets", func(t *testing.T) {
        input := &wafregional.ListByteMatchSetsInput{}
        output := &wafregional.ListByteMatchSetsOutput{}

        mockClient.On("ListByteMatchSets", ctx, input).Return(output, nil)

        result, err := mockClient.ListByteMatchSets(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListGeoMatchSets", func(t *testing.T) {
        input := &wafregional.ListGeoMatchSetsInput{}
        output := &wafregional.ListGeoMatchSetsOutput{}

        mockClient.On("ListGeoMatchSets", ctx, input).Return(output, nil)

        result, err := mockClient.ListGeoMatchSets(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListIPSets", func(t *testing.T) {
        input := &wafregional.ListIPSetsInput{}
        output := &wafregional.ListIPSetsOutput{}

        mockClient.On("ListIPSets", ctx, input).Return(output, nil)

        result, err := mockClient.ListIPSets(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListLoggingConfigurations", func(t *testing.T) {
        input := &wafregional.ListLoggingConfigurationsInput{}
        output := &wafregional.ListLoggingConfigurationsOutput{}

        mockClient.On("ListLoggingConfigurations", ctx, input).Return(output, nil)

        result, err := mockClient.ListLoggingConfigurations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListRateBasedRules", func(t *testing.T) {
        input := &wafregional.ListRateBasedRulesInput{}
        output := &wafregional.ListRateBasedRulesOutput{}

        mockClient.On("ListRateBasedRules", ctx, input).Return(output, nil)

        result, err := mockClient.ListRateBasedRules(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListRegexMatchSets", func(t *testing.T) {
        input := &wafregional.ListRegexMatchSetsInput{}
        output := &wafregional.ListRegexMatchSetsOutput{}

        mockClient.On("ListRegexMatchSets", ctx, input).Return(output, nil)

        result, err := mockClient.ListRegexMatchSets(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListRegexPatternSets", func(t *testing.T) {
        input := &wafregional.ListRegexPatternSetsInput{}
        output := &wafregional.ListRegexPatternSetsOutput{}

        mockClient.On("ListRegexPatternSets", ctx, input).Return(output, nil)

        result, err := mockClient.ListRegexPatternSets(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListResourcesForWebACL", func(t *testing.T) {
        input := &wafregional.ListResourcesForWebACLInput{}
        output := &wafregional.ListResourcesForWebACLOutput{}

        mockClient.On("ListResourcesForWebACL", ctx, input).Return(output, nil)

        result, err := mockClient.ListResourcesForWebACL(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListRuleGroups", func(t *testing.T) {
        input := &wafregional.ListRuleGroupsInput{}
        output := &wafregional.ListRuleGroupsOutput{}

        mockClient.On("ListRuleGroups", ctx, input).Return(output, nil)

        result, err := mockClient.ListRuleGroups(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListRules", func(t *testing.T) {
        input := &wafregional.ListRulesInput{}
        output := &wafregional.ListRulesOutput{}

        mockClient.On("ListRules", ctx, input).Return(output, nil)

        result, err := mockClient.ListRules(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListSizeConstraintSets", func(t *testing.T) {
        input := &wafregional.ListSizeConstraintSetsInput{}
        output := &wafregional.ListSizeConstraintSetsOutput{}

        mockClient.On("ListSizeConstraintSets", ctx, input).Return(output, nil)

        result, err := mockClient.ListSizeConstraintSets(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListSqlInjectionMatchSets", func(t *testing.T) {
        input := &wafregional.ListSqlInjectionMatchSetsInput{}
        output := &wafregional.ListSqlInjectionMatchSetsOutput{}

        mockClient.On("ListSqlInjectionMatchSets", ctx, input).Return(output, nil)

        result, err := mockClient.ListSqlInjectionMatchSets(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListSubscribedRuleGroups", func(t *testing.T) {
        input := &wafregional.ListSubscribedRuleGroupsInput{}
        output := &wafregional.ListSubscribedRuleGroupsOutput{}

        mockClient.On("ListSubscribedRuleGroups", ctx, input).Return(output, nil)

        result, err := mockClient.ListSubscribedRuleGroups(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &wafregional.ListTagsForResourceInput{}
        output := &wafregional.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListWebACLs", func(t *testing.T) {
        input := &wafregional.ListWebACLsInput{}
        output := &wafregional.ListWebACLsOutput{}

        mockClient.On("ListWebACLs", ctx, input).Return(output, nil)

        result, err := mockClient.ListWebACLs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListXssMatchSets", func(t *testing.T) {
        input := &wafregional.ListXssMatchSetsInput{}
        output := &wafregional.ListXssMatchSetsOutput{}

        mockClient.On("ListXssMatchSets", ctx, input).Return(output, nil)

        result, err := mockClient.ListXssMatchSets(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutLoggingConfiguration", func(t *testing.T) {
        input := &wafregional.PutLoggingConfigurationInput{}
        output := &wafregional.PutLoggingConfigurationOutput{}

        mockClient.On("PutLoggingConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.PutLoggingConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutPermissionPolicy", func(t *testing.T) {
        input := &wafregional.PutPermissionPolicyInput{}
        output := &wafregional.PutPermissionPolicyOutput{}

        mockClient.On("PutPermissionPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.PutPermissionPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &wafregional.TagResourceInput{}
        output := &wafregional.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &wafregional.UntagResourceInput{}
        output := &wafregional.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateByteMatchSet", func(t *testing.T) {
        input := &wafregional.UpdateByteMatchSetInput{}
        output := &wafregional.UpdateByteMatchSetOutput{}

        mockClient.On("UpdateByteMatchSet", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateByteMatchSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateGeoMatchSet", func(t *testing.T) {
        input := &wafregional.UpdateGeoMatchSetInput{}
        output := &wafregional.UpdateGeoMatchSetOutput{}

        mockClient.On("UpdateGeoMatchSet", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateGeoMatchSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateIPSet", func(t *testing.T) {
        input := &wafregional.UpdateIPSetInput{}
        output := &wafregional.UpdateIPSetOutput{}

        mockClient.On("UpdateIPSet", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateIPSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateRateBasedRule", func(t *testing.T) {
        input := &wafregional.UpdateRateBasedRuleInput{}
        output := &wafregional.UpdateRateBasedRuleOutput{}

        mockClient.On("UpdateRateBasedRule", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateRateBasedRule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateRegexMatchSet", func(t *testing.T) {
        input := &wafregional.UpdateRegexMatchSetInput{}
        output := &wafregional.UpdateRegexMatchSetOutput{}

        mockClient.On("UpdateRegexMatchSet", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateRegexMatchSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateRegexPatternSet", func(t *testing.T) {
        input := &wafregional.UpdateRegexPatternSetInput{}
        output := &wafregional.UpdateRegexPatternSetOutput{}

        mockClient.On("UpdateRegexPatternSet", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateRegexPatternSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateRule", func(t *testing.T) {
        input := &wafregional.UpdateRuleInput{}
        output := &wafregional.UpdateRuleOutput{}

        mockClient.On("UpdateRule", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateRule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateRuleGroup", func(t *testing.T) {
        input := &wafregional.UpdateRuleGroupInput{}
        output := &wafregional.UpdateRuleGroupOutput{}

        mockClient.On("UpdateRuleGroup", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateRuleGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateSizeConstraintSet", func(t *testing.T) {
        input := &wafregional.UpdateSizeConstraintSetInput{}
        output := &wafregional.UpdateSizeConstraintSetOutput{}

        mockClient.On("UpdateSizeConstraintSet", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateSizeConstraintSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateSqlInjectionMatchSet", func(t *testing.T) {
        input := &wafregional.UpdateSqlInjectionMatchSetInput{}
        output := &wafregional.UpdateSqlInjectionMatchSetOutput{}

        mockClient.On("UpdateSqlInjectionMatchSet", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateSqlInjectionMatchSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateWebACL", func(t *testing.T) {
        input := &wafregional.UpdateWebACLInput{}
        output := &wafregional.UpdateWebACLOutput{}

        mockClient.On("UpdateWebACL", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateWebACL(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateXssMatchSet", func(t *testing.T) {
        input := &wafregional.UpdateXssMatchSetInput{}
        output := &wafregional.UpdateXssMatchSetOutput{}

        mockClient.On("UpdateXssMatchSet", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateXssMatchSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
