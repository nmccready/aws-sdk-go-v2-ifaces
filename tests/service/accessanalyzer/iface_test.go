// Code generated by ./scripts/gen_tests.sh DO NOT EDIT.
package accessanalyzer_test

// tests for the accessanalyzer service interface for 
// this ../../../service/accessanalyzer/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/accessanalyzer"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/accessanalyzer/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/accessanalyzer/accessanalyzer_iface"
	"github.com/stretchr/testify/assert"
)

func TestAccessanalyzerServiceCanBeMocked(t *testing.T) {
	var iface accessanalyzer_iface.IClient
	iface = &accessanalyzer.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := accessanalyzer.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestApplyArchiveRule", func(t *testing.T) {
        input := &accessanalyzer.ApplyArchiveRuleInput{}
        output := &accessanalyzer.ApplyArchiveRuleOutput{}

        mockClient.On("ApplyArchiveRule", ctx, input).Return(output, nil)

        result, err := mockClient.ApplyArchiveRule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCancelPolicyGeneration", func(t *testing.T) {
        input := &accessanalyzer.CancelPolicyGenerationInput{}
        output := &accessanalyzer.CancelPolicyGenerationOutput{}

        mockClient.On("CancelPolicyGeneration", ctx, input).Return(output, nil)

        result, err := mockClient.CancelPolicyGeneration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCheckAccessNotGranted", func(t *testing.T) {
        input := &accessanalyzer.CheckAccessNotGrantedInput{}
        output := &accessanalyzer.CheckAccessNotGrantedOutput{}

        mockClient.On("CheckAccessNotGranted", ctx, input).Return(output, nil)

        result, err := mockClient.CheckAccessNotGranted(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCheckNoNewAccess", func(t *testing.T) {
        input := &accessanalyzer.CheckNoNewAccessInput{}
        output := &accessanalyzer.CheckNoNewAccessOutput{}

        mockClient.On("CheckNoNewAccess", ctx, input).Return(output, nil)

        result, err := mockClient.CheckNoNewAccess(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCheckNoPublicAccess", func(t *testing.T) {
        input := &accessanalyzer.CheckNoPublicAccessInput{}
        output := &accessanalyzer.CheckNoPublicAccessOutput{}

        mockClient.On("CheckNoPublicAccess", ctx, input).Return(output, nil)

        result, err := mockClient.CheckNoPublicAccess(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateAccessPreview", func(t *testing.T) {
        input := &accessanalyzer.CreateAccessPreviewInput{}
        output := &accessanalyzer.CreateAccessPreviewOutput{}

        mockClient.On("CreateAccessPreview", ctx, input).Return(output, nil)

        result, err := mockClient.CreateAccessPreview(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateAnalyzer", func(t *testing.T) {
        input := &accessanalyzer.CreateAnalyzerInput{}
        output := &accessanalyzer.CreateAnalyzerOutput{}

        mockClient.On("CreateAnalyzer", ctx, input).Return(output, nil)

        result, err := mockClient.CreateAnalyzer(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateArchiveRule", func(t *testing.T) {
        input := &accessanalyzer.CreateArchiveRuleInput{}
        output := &accessanalyzer.CreateArchiveRuleOutput{}

        mockClient.On("CreateArchiveRule", ctx, input).Return(output, nil)

        result, err := mockClient.CreateArchiveRule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteAnalyzer", func(t *testing.T) {
        input := &accessanalyzer.DeleteAnalyzerInput{}
        output := &accessanalyzer.DeleteAnalyzerOutput{}

        mockClient.On("DeleteAnalyzer", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteAnalyzer(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteArchiveRule", func(t *testing.T) {
        input := &accessanalyzer.DeleteArchiveRuleInput{}
        output := &accessanalyzer.DeleteArchiveRuleOutput{}

        mockClient.On("DeleteArchiveRule", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteArchiveRule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGenerateFindingRecommendation", func(t *testing.T) {
        input := &accessanalyzer.GenerateFindingRecommendationInput{}
        output := &accessanalyzer.GenerateFindingRecommendationOutput{}

        mockClient.On("GenerateFindingRecommendation", ctx, input).Return(output, nil)

        result, err := mockClient.GenerateFindingRecommendation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetAccessPreview", func(t *testing.T) {
        input := &accessanalyzer.GetAccessPreviewInput{}
        output := &accessanalyzer.GetAccessPreviewOutput{}

        mockClient.On("GetAccessPreview", ctx, input).Return(output, nil)

        result, err := mockClient.GetAccessPreview(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetAnalyzedResource", func(t *testing.T) {
        input := &accessanalyzer.GetAnalyzedResourceInput{}
        output := &accessanalyzer.GetAnalyzedResourceOutput{}

        mockClient.On("GetAnalyzedResource", ctx, input).Return(output, nil)

        result, err := mockClient.GetAnalyzedResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetAnalyzer", func(t *testing.T) {
        input := &accessanalyzer.GetAnalyzerInput{}
        output := &accessanalyzer.GetAnalyzerOutput{}

        mockClient.On("GetAnalyzer", ctx, input).Return(output, nil)

        result, err := mockClient.GetAnalyzer(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetArchiveRule", func(t *testing.T) {
        input := &accessanalyzer.GetArchiveRuleInput{}
        output := &accessanalyzer.GetArchiveRuleOutput{}

        mockClient.On("GetArchiveRule", ctx, input).Return(output, nil)

        result, err := mockClient.GetArchiveRule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetFinding", func(t *testing.T) {
        input := &accessanalyzer.GetFindingInput{}
        output := &accessanalyzer.GetFindingOutput{}

        mockClient.On("GetFinding", ctx, input).Return(output, nil)

        result, err := mockClient.GetFinding(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetFindingRecommendation", func(t *testing.T) {
        input := &accessanalyzer.GetFindingRecommendationInput{}
        output := &accessanalyzer.GetFindingRecommendationOutput{}

        mockClient.On("GetFindingRecommendation", ctx, input).Return(output, nil)

        result, err := mockClient.GetFindingRecommendation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetFindingV2", func(t *testing.T) {
        input := &accessanalyzer.GetFindingV2Input{}
        output := &accessanalyzer.GetFindingV2Output{}

        mockClient.On("GetFindingV2", ctx, input).Return(output, nil)

        result, err := mockClient.GetFindingV2(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetGeneratedPolicy", func(t *testing.T) {
        input := &accessanalyzer.GetGeneratedPolicyInput{}
        output := &accessanalyzer.GetGeneratedPolicyOutput{}

        mockClient.On("GetGeneratedPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.GetGeneratedPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAccessPreviewFindings", func(t *testing.T) {
        input := &accessanalyzer.ListAccessPreviewFindingsInput{}
        output := &accessanalyzer.ListAccessPreviewFindingsOutput{}

        mockClient.On("ListAccessPreviewFindings", ctx, input).Return(output, nil)

        result, err := mockClient.ListAccessPreviewFindings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAccessPreviews", func(t *testing.T) {
        input := &accessanalyzer.ListAccessPreviewsInput{}
        output := &accessanalyzer.ListAccessPreviewsOutput{}

        mockClient.On("ListAccessPreviews", ctx, input).Return(output, nil)

        result, err := mockClient.ListAccessPreviews(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAnalyzedResources", func(t *testing.T) {
        input := &accessanalyzer.ListAnalyzedResourcesInput{}
        output := &accessanalyzer.ListAnalyzedResourcesOutput{}

        mockClient.On("ListAnalyzedResources", ctx, input).Return(output, nil)

        result, err := mockClient.ListAnalyzedResources(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAnalyzers", func(t *testing.T) {
        input := &accessanalyzer.ListAnalyzersInput{}
        output := &accessanalyzer.ListAnalyzersOutput{}

        mockClient.On("ListAnalyzers", ctx, input).Return(output, nil)

        result, err := mockClient.ListAnalyzers(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListArchiveRules", func(t *testing.T) {
        input := &accessanalyzer.ListArchiveRulesInput{}
        output := &accessanalyzer.ListArchiveRulesOutput{}

        mockClient.On("ListArchiveRules", ctx, input).Return(output, nil)

        result, err := mockClient.ListArchiveRules(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListFindings", func(t *testing.T) {
        input := &accessanalyzer.ListFindingsInput{}
        output := &accessanalyzer.ListFindingsOutput{}

        mockClient.On("ListFindings", ctx, input).Return(output, nil)

        result, err := mockClient.ListFindings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListFindingsV2", func(t *testing.T) {
        input := &accessanalyzer.ListFindingsV2Input{}
        output := &accessanalyzer.ListFindingsV2Output{}

        mockClient.On("ListFindingsV2", ctx, input).Return(output, nil)

        result, err := mockClient.ListFindingsV2(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListPolicyGenerations", func(t *testing.T) {
        input := &accessanalyzer.ListPolicyGenerationsInput{}
        output := &accessanalyzer.ListPolicyGenerationsOutput{}

        mockClient.On("ListPolicyGenerations", ctx, input).Return(output, nil)

        result, err := mockClient.ListPolicyGenerations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &accessanalyzer.ListTagsForResourceInput{}
        output := &accessanalyzer.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartPolicyGeneration", func(t *testing.T) {
        input := &accessanalyzer.StartPolicyGenerationInput{}
        output := &accessanalyzer.StartPolicyGenerationOutput{}

        mockClient.On("StartPolicyGeneration", ctx, input).Return(output, nil)

        result, err := mockClient.StartPolicyGeneration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartResourceScan", func(t *testing.T) {
        input := &accessanalyzer.StartResourceScanInput{}
        output := &accessanalyzer.StartResourceScanOutput{}

        mockClient.On("StartResourceScan", ctx, input).Return(output, nil)

        result, err := mockClient.StartResourceScan(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &accessanalyzer.TagResourceInput{}
        output := &accessanalyzer.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &accessanalyzer.UntagResourceInput{}
        output := &accessanalyzer.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateArchiveRule", func(t *testing.T) {
        input := &accessanalyzer.UpdateArchiveRuleInput{}
        output := &accessanalyzer.UpdateArchiveRuleOutput{}

        mockClient.On("UpdateArchiveRule", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateArchiveRule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateFindings", func(t *testing.T) {
        input := &accessanalyzer.UpdateFindingsInput{}
        output := &accessanalyzer.UpdateFindingsOutput{}

        mockClient.On("UpdateFindings", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateFindings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestValidatePolicy", func(t *testing.T) {
        input := &accessanalyzer.ValidatePolicyInput{}
        output := &accessanalyzer.ValidatePolicyOutput{}

        mockClient.On("ValidatePolicy", ctx, input).Return(output, nil)

        result, err := mockClient.ValidatePolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
