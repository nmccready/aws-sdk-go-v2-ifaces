package waf_test

// tests for the waf service interface for this ../../../service/waf/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/waf"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/waf/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/waf/waf_iface"
	"github.com/stretchr/testify/assert"
)

func TestWafServiceCanBeMocked(t *testing.T) {
	var iface waf_iface.IClient
	iface = &waf.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := waf.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateByteMatchSet", func(t *testing.T) {
        input := &waf.CreateByteMatchSetInput{}
        output := &waf.CreateByteMatchSetOutput{}

        mockClient.On("CreateByteMatchSet", ctx, input).Return(output, nil)

        result, err := mockClient.CreateByteMatchSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateGeoMatchSet", func(t *testing.T) {
        input := &waf.CreateGeoMatchSetInput{}
        output := &waf.CreateGeoMatchSetOutput{}

        mockClient.On("CreateGeoMatchSet", ctx, input).Return(output, nil)

        result, err := mockClient.CreateGeoMatchSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateIPSet", func(t *testing.T) {
        input := &waf.CreateIPSetInput{}
        output := &waf.CreateIPSetOutput{}

        mockClient.On("CreateIPSet", ctx, input).Return(output, nil)

        result, err := mockClient.CreateIPSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateRateBasedRule", func(t *testing.T) {
        input := &waf.CreateRateBasedRuleInput{}
        output := &waf.CreateRateBasedRuleOutput{}

        mockClient.On("CreateRateBasedRule", ctx, input).Return(output, nil)

        result, err := mockClient.CreateRateBasedRule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateRegexMatchSet", func(t *testing.T) {
        input := &waf.CreateRegexMatchSetInput{}
        output := &waf.CreateRegexMatchSetOutput{}

        mockClient.On("CreateRegexMatchSet", ctx, input).Return(output, nil)

        result, err := mockClient.CreateRegexMatchSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateRegexPatternSet", func(t *testing.T) {
        input := &waf.CreateRegexPatternSetInput{}
        output := &waf.CreateRegexPatternSetOutput{}

        mockClient.On("CreateRegexPatternSet", ctx, input).Return(output, nil)

        result, err := mockClient.CreateRegexPatternSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateRule", func(t *testing.T) {
        input := &waf.CreateRuleInput{}
        output := &waf.CreateRuleOutput{}

        mockClient.On("CreateRule", ctx, input).Return(output, nil)

        result, err := mockClient.CreateRule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateRuleGroup", func(t *testing.T) {
        input := &waf.CreateRuleGroupInput{}
        output := &waf.CreateRuleGroupOutput{}

        mockClient.On("CreateRuleGroup", ctx, input).Return(output, nil)

        result, err := mockClient.CreateRuleGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateSizeConstraintSet", func(t *testing.T) {
        input := &waf.CreateSizeConstraintSetInput{}
        output := &waf.CreateSizeConstraintSetOutput{}

        mockClient.On("CreateSizeConstraintSet", ctx, input).Return(output, nil)

        result, err := mockClient.CreateSizeConstraintSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateSqlInjectionMatchSet", func(t *testing.T) {
        input := &waf.CreateSqlInjectionMatchSetInput{}
        output := &waf.CreateSqlInjectionMatchSetOutput{}

        mockClient.On("CreateSqlInjectionMatchSet", ctx, input).Return(output, nil)

        result, err := mockClient.CreateSqlInjectionMatchSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateWebACL", func(t *testing.T) {
        input := &waf.CreateWebACLInput{}
        output := &waf.CreateWebACLOutput{}

        mockClient.On("CreateWebACL", ctx, input).Return(output, nil)

        result, err := mockClient.CreateWebACL(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateWebACLMigrationStack", func(t *testing.T) {
        input := &waf.CreateWebACLMigrationStackInput{}
        output := &waf.CreateWebACLMigrationStackOutput{}

        mockClient.On("CreateWebACLMigrationStack", ctx, input).Return(output, nil)

        result, err := mockClient.CreateWebACLMigrationStack(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateXssMatchSet", func(t *testing.T) {
        input := &waf.CreateXssMatchSetInput{}
        output := &waf.CreateXssMatchSetOutput{}

        mockClient.On("CreateXssMatchSet", ctx, input).Return(output, nil)

        result, err := mockClient.CreateXssMatchSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteByteMatchSet", func(t *testing.T) {
        input := &waf.DeleteByteMatchSetInput{}
        output := &waf.DeleteByteMatchSetOutput{}

        mockClient.On("DeleteByteMatchSet", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteByteMatchSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteGeoMatchSet", func(t *testing.T) {
        input := &waf.DeleteGeoMatchSetInput{}
        output := &waf.DeleteGeoMatchSetOutput{}

        mockClient.On("DeleteGeoMatchSet", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteGeoMatchSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteIPSet", func(t *testing.T) {
        input := &waf.DeleteIPSetInput{}
        output := &waf.DeleteIPSetOutput{}

        mockClient.On("DeleteIPSet", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteIPSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteLoggingConfiguration", func(t *testing.T) {
        input := &waf.DeleteLoggingConfigurationInput{}
        output := &waf.DeleteLoggingConfigurationOutput{}

        mockClient.On("DeleteLoggingConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteLoggingConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeletePermissionPolicy", func(t *testing.T) {
        input := &waf.DeletePermissionPolicyInput{}
        output := &waf.DeletePermissionPolicyOutput{}

        mockClient.On("DeletePermissionPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.DeletePermissionPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteRateBasedRule", func(t *testing.T) {
        input := &waf.DeleteRateBasedRuleInput{}
        output := &waf.DeleteRateBasedRuleOutput{}

        mockClient.On("DeleteRateBasedRule", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteRateBasedRule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteRegexMatchSet", func(t *testing.T) {
        input := &waf.DeleteRegexMatchSetInput{}
        output := &waf.DeleteRegexMatchSetOutput{}

        mockClient.On("DeleteRegexMatchSet", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteRegexMatchSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteRegexPatternSet", func(t *testing.T) {
        input := &waf.DeleteRegexPatternSetInput{}
        output := &waf.DeleteRegexPatternSetOutput{}

        mockClient.On("DeleteRegexPatternSet", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteRegexPatternSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteRule", func(t *testing.T) {
        input := &waf.DeleteRuleInput{}
        output := &waf.DeleteRuleOutput{}

        mockClient.On("DeleteRule", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteRule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteRuleGroup", func(t *testing.T) {
        input := &waf.DeleteRuleGroupInput{}
        output := &waf.DeleteRuleGroupOutput{}

        mockClient.On("DeleteRuleGroup", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteRuleGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteSizeConstraintSet", func(t *testing.T) {
        input := &waf.DeleteSizeConstraintSetInput{}
        output := &waf.DeleteSizeConstraintSetOutput{}

        mockClient.On("DeleteSizeConstraintSet", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteSizeConstraintSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteSqlInjectionMatchSet", func(t *testing.T) {
        input := &waf.DeleteSqlInjectionMatchSetInput{}
        output := &waf.DeleteSqlInjectionMatchSetOutput{}

        mockClient.On("DeleteSqlInjectionMatchSet", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteSqlInjectionMatchSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteWebACL", func(t *testing.T) {
        input := &waf.DeleteWebACLInput{}
        output := &waf.DeleteWebACLOutput{}

        mockClient.On("DeleteWebACL", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteWebACL(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteXssMatchSet", func(t *testing.T) {
        input := &waf.DeleteXssMatchSetInput{}
        output := &waf.DeleteXssMatchSetOutput{}

        mockClient.On("DeleteXssMatchSet", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteXssMatchSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetByteMatchSet", func(t *testing.T) {
        input := &waf.GetByteMatchSetInput{}
        output := &waf.GetByteMatchSetOutput{}

        mockClient.On("GetByteMatchSet", ctx, input).Return(output, nil)

        result, err := mockClient.GetByteMatchSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetChangeToken", func(t *testing.T) {
        input := &waf.GetChangeTokenInput{}
        output := &waf.GetChangeTokenOutput{}

        mockClient.On("GetChangeToken", ctx, input).Return(output, nil)

        result, err := mockClient.GetChangeToken(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetChangeTokenStatus", func(t *testing.T) {
        input := &waf.GetChangeTokenStatusInput{}
        output := &waf.GetChangeTokenStatusOutput{}

        mockClient.On("GetChangeTokenStatus", ctx, input).Return(output, nil)

        result, err := mockClient.GetChangeTokenStatus(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetGeoMatchSet", func(t *testing.T) {
        input := &waf.GetGeoMatchSetInput{}
        output := &waf.GetGeoMatchSetOutput{}

        mockClient.On("GetGeoMatchSet", ctx, input).Return(output, nil)

        result, err := mockClient.GetGeoMatchSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetIPSet", func(t *testing.T) {
        input := &waf.GetIPSetInput{}
        output := &waf.GetIPSetOutput{}

        mockClient.On("GetIPSet", ctx, input).Return(output, nil)

        result, err := mockClient.GetIPSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetLoggingConfiguration", func(t *testing.T) {
        input := &waf.GetLoggingConfigurationInput{}
        output := &waf.GetLoggingConfigurationOutput{}

        mockClient.On("GetLoggingConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.GetLoggingConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetPermissionPolicy", func(t *testing.T) {
        input := &waf.GetPermissionPolicyInput{}
        output := &waf.GetPermissionPolicyOutput{}

        mockClient.On("GetPermissionPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.GetPermissionPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetRateBasedRule", func(t *testing.T) {
        input := &waf.GetRateBasedRuleInput{}
        output := &waf.GetRateBasedRuleOutput{}

        mockClient.On("GetRateBasedRule", ctx, input).Return(output, nil)

        result, err := mockClient.GetRateBasedRule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetRateBasedRuleManagedKeys", func(t *testing.T) {
        input := &waf.GetRateBasedRuleManagedKeysInput{}
        output := &waf.GetRateBasedRuleManagedKeysOutput{}

        mockClient.On("GetRateBasedRuleManagedKeys", ctx, input).Return(output, nil)

        result, err := mockClient.GetRateBasedRuleManagedKeys(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetRegexMatchSet", func(t *testing.T) {
        input := &waf.GetRegexMatchSetInput{}
        output := &waf.GetRegexMatchSetOutput{}

        mockClient.On("GetRegexMatchSet", ctx, input).Return(output, nil)

        result, err := mockClient.GetRegexMatchSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetRegexPatternSet", func(t *testing.T) {
        input := &waf.GetRegexPatternSetInput{}
        output := &waf.GetRegexPatternSetOutput{}

        mockClient.On("GetRegexPatternSet", ctx, input).Return(output, nil)

        result, err := mockClient.GetRegexPatternSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetRule", func(t *testing.T) {
        input := &waf.GetRuleInput{}
        output := &waf.GetRuleOutput{}

        mockClient.On("GetRule", ctx, input).Return(output, nil)

        result, err := mockClient.GetRule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetRuleGroup", func(t *testing.T) {
        input := &waf.GetRuleGroupInput{}
        output := &waf.GetRuleGroupOutput{}

        mockClient.On("GetRuleGroup", ctx, input).Return(output, nil)

        result, err := mockClient.GetRuleGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSampledRequests", func(t *testing.T) {
        input := &waf.GetSampledRequestsInput{}
        output := &waf.GetSampledRequestsOutput{}

        mockClient.On("GetSampledRequests", ctx, input).Return(output, nil)

        result, err := mockClient.GetSampledRequests(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSizeConstraintSet", func(t *testing.T) {
        input := &waf.GetSizeConstraintSetInput{}
        output := &waf.GetSizeConstraintSetOutput{}

        mockClient.On("GetSizeConstraintSet", ctx, input).Return(output, nil)

        result, err := mockClient.GetSizeConstraintSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSqlInjectionMatchSet", func(t *testing.T) {
        input := &waf.GetSqlInjectionMatchSetInput{}
        output := &waf.GetSqlInjectionMatchSetOutput{}

        mockClient.On("GetSqlInjectionMatchSet", ctx, input).Return(output, nil)

        result, err := mockClient.GetSqlInjectionMatchSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetWebACL", func(t *testing.T) {
        input := &waf.GetWebACLInput{}
        output := &waf.GetWebACLOutput{}

        mockClient.On("GetWebACL", ctx, input).Return(output, nil)

        result, err := mockClient.GetWebACL(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetXssMatchSet", func(t *testing.T) {
        input := &waf.GetXssMatchSetInput{}
        output := &waf.GetXssMatchSetOutput{}

        mockClient.On("GetXssMatchSet", ctx, input).Return(output, nil)

        result, err := mockClient.GetXssMatchSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListActivatedRulesInRuleGroup", func(t *testing.T) {
        input := &waf.ListActivatedRulesInRuleGroupInput{}
        output := &waf.ListActivatedRulesInRuleGroupOutput{}

        mockClient.On("ListActivatedRulesInRuleGroup", ctx, input).Return(output, nil)

        result, err := mockClient.ListActivatedRulesInRuleGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListByteMatchSets", func(t *testing.T) {
        input := &waf.ListByteMatchSetsInput{}
        output := &waf.ListByteMatchSetsOutput{}

        mockClient.On("ListByteMatchSets", ctx, input).Return(output, nil)

        result, err := mockClient.ListByteMatchSets(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListGeoMatchSets", func(t *testing.T) {
        input := &waf.ListGeoMatchSetsInput{}
        output := &waf.ListGeoMatchSetsOutput{}

        mockClient.On("ListGeoMatchSets", ctx, input).Return(output, nil)

        result, err := mockClient.ListGeoMatchSets(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListIPSets", func(t *testing.T) {
        input := &waf.ListIPSetsInput{}
        output := &waf.ListIPSetsOutput{}

        mockClient.On("ListIPSets", ctx, input).Return(output, nil)

        result, err := mockClient.ListIPSets(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListLoggingConfigurations", func(t *testing.T) {
        input := &waf.ListLoggingConfigurationsInput{}
        output := &waf.ListLoggingConfigurationsOutput{}

        mockClient.On("ListLoggingConfigurations", ctx, input).Return(output, nil)

        result, err := mockClient.ListLoggingConfigurations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListRateBasedRules", func(t *testing.T) {
        input := &waf.ListRateBasedRulesInput{}
        output := &waf.ListRateBasedRulesOutput{}

        mockClient.On("ListRateBasedRules", ctx, input).Return(output, nil)

        result, err := mockClient.ListRateBasedRules(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListRegexMatchSets", func(t *testing.T) {
        input := &waf.ListRegexMatchSetsInput{}
        output := &waf.ListRegexMatchSetsOutput{}

        mockClient.On("ListRegexMatchSets", ctx, input).Return(output, nil)

        result, err := mockClient.ListRegexMatchSets(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListRegexPatternSets", func(t *testing.T) {
        input := &waf.ListRegexPatternSetsInput{}
        output := &waf.ListRegexPatternSetsOutput{}

        mockClient.On("ListRegexPatternSets", ctx, input).Return(output, nil)

        result, err := mockClient.ListRegexPatternSets(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListRuleGroups", func(t *testing.T) {
        input := &waf.ListRuleGroupsInput{}
        output := &waf.ListRuleGroupsOutput{}

        mockClient.On("ListRuleGroups", ctx, input).Return(output, nil)

        result, err := mockClient.ListRuleGroups(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListRules", func(t *testing.T) {
        input := &waf.ListRulesInput{}
        output := &waf.ListRulesOutput{}

        mockClient.On("ListRules", ctx, input).Return(output, nil)

        result, err := mockClient.ListRules(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListSizeConstraintSets", func(t *testing.T) {
        input := &waf.ListSizeConstraintSetsInput{}
        output := &waf.ListSizeConstraintSetsOutput{}

        mockClient.On("ListSizeConstraintSets", ctx, input).Return(output, nil)

        result, err := mockClient.ListSizeConstraintSets(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListSqlInjectionMatchSets", func(t *testing.T) {
        input := &waf.ListSqlInjectionMatchSetsInput{}
        output := &waf.ListSqlInjectionMatchSetsOutput{}

        mockClient.On("ListSqlInjectionMatchSets", ctx, input).Return(output, nil)

        result, err := mockClient.ListSqlInjectionMatchSets(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListSubscribedRuleGroups", func(t *testing.T) {
        input := &waf.ListSubscribedRuleGroupsInput{}
        output := &waf.ListSubscribedRuleGroupsOutput{}

        mockClient.On("ListSubscribedRuleGroups", ctx, input).Return(output, nil)

        result, err := mockClient.ListSubscribedRuleGroups(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &waf.ListTagsForResourceInput{}
        output := &waf.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListWebACLs", func(t *testing.T) {
        input := &waf.ListWebACLsInput{}
        output := &waf.ListWebACLsOutput{}

        mockClient.On("ListWebACLs", ctx, input).Return(output, nil)

        result, err := mockClient.ListWebACLs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListXssMatchSets", func(t *testing.T) {
        input := &waf.ListXssMatchSetsInput{}
        output := &waf.ListXssMatchSetsOutput{}

        mockClient.On("ListXssMatchSets", ctx, input).Return(output, nil)

        result, err := mockClient.ListXssMatchSets(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutLoggingConfiguration", func(t *testing.T) {
        input := &waf.PutLoggingConfigurationInput{}
        output := &waf.PutLoggingConfigurationOutput{}

        mockClient.On("PutLoggingConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.PutLoggingConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutPermissionPolicy", func(t *testing.T) {
        input := &waf.PutPermissionPolicyInput{}
        output := &waf.PutPermissionPolicyOutput{}

        mockClient.On("PutPermissionPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.PutPermissionPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &waf.TagResourceInput{}
        output := &waf.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &waf.UntagResourceInput{}
        output := &waf.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateByteMatchSet", func(t *testing.T) {
        input := &waf.UpdateByteMatchSetInput{}
        output := &waf.UpdateByteMatchSetOutput{}

        mockClient.On("UpdateByteMatchSet", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateByteMatchSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateGeoMatchSet", func(t *testing.T) {
        input := &waf.UpdateGeoMatchSetInput{}
        output := &waf.UpdateGeoMatchSetOutput{}

        mockClient.On("UpdateGeoMatchSet", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateGeoMatchSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateIPSet", func(t *testing.T) {
        input := &waf.UpdateIPSetInput{}
        output := &waf.UpdateIPSetOutput{}

        mockClient.On("UpdateIPSet", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateIPSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateRateBasedRule", func(t *testing.T) {
        input := &waf.UpdateRateBasedRuleInput{}
        output := &waf.UpdateRateBasedRuleOutput{}

        mockClient.On("UpdateRateBasedRule", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateRateBasedRule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateRegexMatchSet", func(t *testing.T) {
        input := &waf.UpdateRegexMatchSetInput{}
        output := &waf.UpdateRegexMatchSetOutput{}

        mockClient.On("UpdateRegexMatchSet", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateRegexMatchSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateRegexPatternSet", func(t *testing.T) {
        input := &waf.UpdateRegexPatternSetInput{}
        output := &waf.UpdateRegexPatternSetOutput{}

        mockClient.On("UpdateRegexPatternSet", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateRegexPatternSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateRule", func(t *testing.T) {
        input := &waf.UpdateRuleInput{}
        output := &waf.UpdateRuleOutput{}

        mockClient.On("UpdateRule", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateRule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateRuleGroup", func(t *testing.T) {
        input := &waf.UpdateRuleGroupInput{}
        output := &waf.UpdateRuleGroupOutput{}

        mockClient.On("UpdateRuleGroup", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateRuleGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateSizeConstraintSet", func(t *testing.T) {
        input := &waf.UpdateSizeConstraintSetInput{}
        output := &waf.UpdateSizeConstraintSetOutput{}

        mockClient.On("UpdateSizeConstraintSet", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateSizeConstraintSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateSqlInjectionMatchSet", func(t *testing.T) {
        input := &waf.UpdateSqlInjectionMatchSetInput{}
        output := &waf.UpdateSqlInjectionMatchSetOutput{}

        mockClient.On("UpdateSqlInjectionMatchSet", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateSqlInjectionMatchSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateWebACL", func(t *testing.T) {
        input := &waf.UpdateWebACLInput{}
        output := &waf.UpdateWebACLOutput{}

        mockClient.On("UpdateWebACL", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateWebACL(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateXssMatchSet", func(t *testing.T) {
        input := &waf.UpdateXssMatchSetInput{}
        output := &waf.UpdateXssMatchSetOutput{}

        mockClient.On("UpdateXssMatchSet", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateXssMatchSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
