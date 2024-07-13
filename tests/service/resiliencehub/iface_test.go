package resiliencehub_test

// tests for the resiliencehub service interface for this ../../../service/resiliencehub/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/resiliencehub"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/resiliencehub/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/resiliencehub/resiliencehub_iface"
	"github.com/stretchr/testify/assert"
)

func TestResiliencehubServiceCanBeMocked(t *testing.T) {
	var iface resiliencehub_iface.IClient
	iface = &resiliencehub.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := resiliencehub.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAddDraftAppVersionResourceMappings", func(t *testing.T) {
        input := &resiliencehub.AddDraftAppVersionResourceMappingsInput{}
        output := &resiliencehub.AddDraftAppVersionResourceMappingsOutput{}

        mockClient.On("AddDraftAppVersionResourceMappings", ctx, input).Return(output, nil)

        result, err := mockClient.AddDraftAppVersionResourceMappings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchUpdateRecommendationStatus", func(t *testing.T) {
        input := &resiliencehub.BatchUpdateRecommendationStatusInput{}
        output := &resiliencehub.BatchUpdateRecommendationStatusOutput{}

        mockClient.On("BatchUpdateRecommendationStatus", ctx, input).Return(output, nil)

        result, err := mockClient.BatchUpdateRecommendationStatus(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateApp", func(t *testing.T) {
        input := &resiliencehub.CreateAppInput{}
        output := &resiliencehub.CreateAppOutput{}

        mockClient.On("CreateApp", ctx, input).Return(output, nil)

        result, err := mockClient.CreateApp(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateAppVersionAppComponent", func(t *testing.T) {
        input := &resiliencehub.CreateAppVersionAppComponentInput{}
        output := &resiliencehub.CreateAppVersionAppComponentOutput{}

        mockClient.On("CreateAppVersionAppComponent", ctx, input).Return(output, nil)

        result, err := mockClient.CreateAppVersionAppComponent(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateAppVersionResource", func(t *testing.T) {
        input := &resiliencehub.CreateAppVersionResourceInput{}
        output := &resiliencehub.CreateAppVersionResourceOutput{}

        mockClient.On("CreateAppVersionResource", ctx, input).Return(output, nil)

        result, err := mockClient.CreateAppVersionResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateRecommendationTemplate", func(t *testing.T) {
        input := &resiliencehub.CreateRecommendationTemplateInput{}
        output := &resiliencehub.CreateRecommendationTemplateOutput{}

        mockClient.On("CreateRecommendationTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.CreateRecommendationTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateResiliencyPolicy", func(t *testing.T) {
        input := &resiliencehub.CreateResiliencyPolicyInput{}
        output := &resiliencehub.CreateResiliencyPolicyOutput{}

        mockClient.On("CreateResiliencyPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.CreateResiliencyPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteApp", func(t *testing.T) {
        input := &resiliencehub.DeleteAppInput{}
        output := &resiliencehub.DeleteAppOutput{}

        mockClient.On("DeleteApp", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteApp(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteAppAssessment", func(t *testing.T) {
        input := &resiliencehub.DeleteAppAssessmentInput{}
        output := &resiliencehub.DeleteAppAssessmentOutput{}

        mockClient.On("DeleteAppAssessment", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteAppAssessment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteAppInputSource", func(t *testing.T) {
        input := &resiliencehub.DeleteAppInputSourceInput{}
        output := &resiliencehub.DeleteAppInputSourceOutput{}

        mockClient.On("DeleteAppInputSource", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteAppInputSource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteAppVersionAppComponent", func(t *testing.T) {
        input := &resiliencehub.DeleteAppVersionAppComponentInput{}
        output := &resiliencehub.DeleteAppVersionAppComponentOutput{}

        mockClient.On("DeleteAppVersionAppComponent", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteAppVersionAppComponent(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteAppVersionResource", func(t *testing.T) {
        input := &resiliencehub.DeleteAppVersionResourceInput{}
        output := &resiliencehub.DeleteAppVersionResourceOutput{}

        mockClient.On("DeleteAppVersionResource", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteAppVersionResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteRecommendationTemplate", func(t *testing.T) {
        input := &resiliencehub.DeleteRecommendationTemplateInput{}
        output := &resiliencehub.DeleteRecommendationTemplateOutput{}

        mockClient.On("DeleteRecommendationTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteRecommendationTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteResiliencyPolicy", func(t *testing.T) {
        input := &resiliencehub.DeleteResiliencyPolicyInput{}
        output := &resiliencehub.DeleteResiliencyPolicyOutput{}

        mockClient.On("DeleteResiliencyPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteResiliencyPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeApp", func(t *testing.T) {
        input := &resiliencehub.DescribeAppInput{}
        output := &resiliencehub.DescribeAppOutput{}

        mockClient.On("DescribeApp", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeApp(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeAppAssessment", func(t *testing.T) {
        input := &resiliencehub.DescribeAppAssessmentInput{}
        output := &resiliencehub.DescribeAppAssessmentOutput{}

        mockClient.On("DescribeAppAssessment", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeAppAssessment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeAppVersion", func(t *testing.T) {
        input := &resiliencehub.DescribeAppVersionInput{}
        output := &resiliencehub.DescribeAppVersionOutput{}

        mockClient.On("DescribeAppVersion", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeAppVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeAppVersionAppComponent", func(t *testing.T) {
        input := &resiliencehub.DescribeAppVersionAppComponentInput{}
        output := &resiliencehub.DescribeAppVersionAppComponentOutput{}

        mockClient.On("DescribeAppVersionAppComponent", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeAppVersionAppComponent(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeAppVersionResource", func(t *testing.T) {
        input := &resiliencehub.DescribeAppVersionResourceInput{}
        output := &resiliencehub.DescribeAppVersionResourceOutput{}

        mockClient.On("DescribeAppVersionResource", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeAppVersionResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeAppVersionResourcesResolutionStatus", func(t *testing.T) {
        input := &resiliencehub.DescribeAppVersionResourcesResolutionStatusInput{}
        output := &resiliencehub.DescribeAppVersionResourcesResolutionStatusOutput{}

        mockClient.On("DescribeAppVersionResourcesResolutionStatus", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeAppVersionResourcesResolutionStatus(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeAppVersionTemplate", func(t *testing.T) {
        input := &resiliencehub.DescribeAppVersionTemplateInput{}
        output := &resiliencehub.DescribeAppVersionTemplateOutput{}

        mockClient.On("DescribeAppVersionTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeAppVersionTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeDraftAppVersionResourcesImportStatus", func(t *testing.T) {
        input := &resiliencehub.DescribeDraftAppVersionResourcesImportStatusInput{}
        output := &resiliencehub.DescribeDraftAppVersionResourcesImportStatusOutput{}

        mockClient.On("DescribeDraftAppVersionResourcesImportStatus", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeDraftAppVersionResourcesImportStatus(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeResiliencyPolicy", func(t *testing.T) {
        input := &resiliencehub.DescribeResiliencyPolicyInput{}
        output := &resiliencehub.DescribeResiliencyPolicyOutput{}

        mockClient.On("DescribeResiliencyPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeResiliencyPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestImportResourcesToDraftAppVersion", func(t *testing.T) {
        input := &resiliencehub.ImportResourcesToDraftAppVersionInput{}
        output := &resiliencehub.ImportResourcesToDraftAppVersionOutput{}

        mockClient.On("ImportResourcesToDraftAppVersion", ctx, input).Return(output, nil)

        result, err := mockClient.ImportResourcesToDraftAppVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAlarmRecommendations", func(t *testing.T) {
        input := &resiliencehub.ListAlarmRecommendationsInput{}
        output := &resiliencehub.ListAlarmRecommendationsOutput{}

        mockClient.On("ListAlarmRecommendations", ctx, input).Return(output, nil)

        result, err := mockClient.ListAlarmRecommendations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAppAssessmentComplianceDrifts", func(t *testing.T) {
        input := &resiliencehub.ListAppAssessmentComplianceDriftsInput{}
        output := &resiliencehub.ListAppAssessmentComplianceDriftsOutput{}

        mockClient.On("ListAppAssessmentComplianceDrifts", ctx, input).Return(output, nil)

        result, err := mockClient.ListAppAssessmentComplianceDrifts(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAppAssessmentResourceDrifts", func(t *testing.T) {
        input := &resiliencehub.ListAppAssessmentResourceDriftsInput{}
        output := &resiliencehub.ListAppAssessmentResourceDriftsOutput{}

        mockClient.On("ListAppAssessmentResourceDrifts", ctx, input).Return(output, nil)

        result, err := mockClient.ListAppAssessmentResourceDrifts(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAppAssessments", func(t *testing.T) {
        input := &resiliencehub.ListAppAssessmentsInput{}
        output := &resiliencehub.ListAppAssessmentsOutput{}

        mockClient.On("ListAppAssessments", ctx, input).Return(output, nil)

        result, err := mockClient.ListAppAssessments(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAppComponentCompliances", func(t *testing.T) {
        input := &resiliencehub.ListAppComponentCompliancesInput{}
        output := &resiliencehub.ListAppComponentCompliancesOutput{}

        mockClient.On("ListAppComponentCompliances", ctx, input).Return(output, nil)

        result, err := mockClient.ListAppComponentCompliances(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAppComponentRecommendations", func(t *testing.T) {
        input := &resiliencehub.ListAppComponentRecommendationsInput{}
        output := &resiliencehub.ListAppComponentRecommendationsOutput{}

        mockClient.On("ListAppComponentRecommendations", ctx, input).Return(output, nil)

        result, err := mockClient.ListAppComponentRecommendations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAppInputSources", func(t *testing.T) {
        input := &resiliencehub.ListAppInputSourcesInput{}
        output := &resiliencehub.ListAppInputSourcesOutput{}

        mockClient.On("ListAppInputSources", ctx, input).Return(output, nil)

        result, err := mockClient.ListAppInputSources(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAppVersionAppComponents", func(t *testing.T) {
        input := &resiliencehub.ListAppVersionAppComponentsInput{}
        output := &resiliencehub.ListAppVersionAppComponentsOutput{}

        mockClient.On("ListAppVersionAppComponents", ctx, input).Return(output, nil)

        result, err := mockClient.ListAppVersionAppComponents(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAppVersionResourceMappings", func(t *testing.T) {
        input := &resiliencehub.ListAppVersionResourceMappingsInput{}
        output := &resiliencehub.ListAppVersionResourceMappingsOutput{}

        mockClient.On("ListAppVersionResourceMappings", ctx, input).Return(output, nil)

        result, err := mockClient.ListAppVersionResourceMappings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAppVersionResources", func(t *testing.T) {
        input := &resiliencehub.ListAppVersionResourcesInput{}
        output := &resiliencehub.ListAppVersionResourcesOutput{}

        mockClient.On("ListAppVersionResources", ctx, input).Return(output, nil)

        result, err := mockClient.ListAppVersionResources(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAppVersions", func(t *testing.T) {
        input := &resiliencehub.ListAppVersionsInput{}
        output := &resiliencehub.ListAppVersionsOutput{}

        mockClient.On("ListAppVersions", ctx, input).Return(output, nil)

        result, err := mockClient.ListAppVersions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListApps", func(t *testing.T) {
        input := &resiliencehub.ListAppsInput{}
        output := &resiliencehub.ListAppsOutput{}

        mockClient.On("ListApps", ctx, input).Return(output, nil)

        result, err := mockClient.ListApps(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListRecommendationTemplates", func(t *testing.T) {
        input := &resiliencehub.ListRecommendationTemplatesInput{}
        output := &resiliencehub.ListRecommendationTemplatesOutput{}

        mockClient.On("ListRecommendationTemplates", ctx, input).Return(output, nil)

        result, err := mockClient.ListRecommendationTemplates(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListResiliencyPolicies", func(t *testing.T) {
        input := &resiliencehub.ListResiliencyPoliciesInput{}
        output := &resiliencehub.ListResiliencyPoliciesOutput{}

        mockClient.On("ListResiliencyPolicies", ctx, input).Return(output, nil)

        result, err := mockClient.ListResiliencyPolicies(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListSopRecommendations", func(t *testing.T) {
        input := &resiliencehub.ListSopRecommendationsInput{}
        output := &resiliencehub.ListSopRecommendationsOutput{}

        mockClient.On("ListSopRecommendations", ctx, input).Return(output, nil)

        result, err := mockClient.ListSopRecommendations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListSuggestedResiliencyPolicies", func(t *testing.T) {
        input := &resiliencehub.ListSuggestedResiliencyPoliciesInput{}
        output := &resiliencehub.ListSuggestedResiliencyPoliciesOutput{}

        mockClient.On("ListSuggestedResiliencyPolicies", ctx, input).Return(output, nil)

        result, err := mockClient.ListSuggestedResiliencyPolicies(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &resiliencehub.ListTagsForResourceInput{}
        output := &resiliencehub.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTestRecommendations", func(t *testing.T) {
        input := &resiliencehub.ListTestRecommendationsInput{}
        output := &resiliencehub.ListTestRecommendationsOutput{}

        mockClient.On("ListTestRecommendations", ctx, input).Return(output, nil)

        result, err := mockClient.ListTestRecommendations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListUnsupportedAppVersionResources", func(t *testing.T) {
        input := &resiliencehub.ListUnsupportedAppVersionResourcesInput{}
        output := &resiliencehub.ListUnsupportedAppVersionResourcesOutput{}

        mockClient.On("ListUnsupportedAppVersionResources", ctx, input).Return(output, nil)

        result, err := mockClient.ListUnsupportedAppVersionResources(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPublishAppVersion", func(t *testing.T) {
        input := &resiliencehub.PublishAppVersionInput{}
        output := &resiliencehub.PublishAppVersionOutput{}

        mockClient.On("PublishAppVersion", ctx, input).Return(output, nil)

        result, err := mockClient.PublishAppVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutDraftAppVersionTemplate", func(t *testing.T) {
        input := &resiliencehub.PutDraftAppVersionTemplateInput{}
        output := &resiliencehub.PutDraftAppVersionTemplateOutput{}

        mockClient.On("PutDraftAppVersionTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.PutDraftAppVersionTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRemoveDraftAppVersionResourceMappings", func(t *testing.T) {
        input := &resiliencehub.RemoveDraftAppVersionResourceMappingsInput{}
        output := &resiliencehub.RemoveDraftAppVersionResourceMappingsOutput{}

        mockClient.On("RemoveDraftAppVersionResourceMappings", ctx, input).Return(output, nil)

        result, err := mockClient.RemoveDraftAppVersionResourceMappings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestResolveAppVersionResources", func(t *testing.T) {
        input := &resiliencehub.ResolveAppVersionResourcesInput{}
        output := &resiliencehub.ResolveAppVersionResourcesOutput{}

        mockClient.On("ResolveAppVersionResources", ctx, input).Return(output, nil)

        result, err := mockClient.ResolveAppVersionResources(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartAppAssessment", func(t *testing.T) {
        input := &resiliencehub.StartAppAssessmentInput{}
        output := &resiliencehub.StartAppAssessmentOutput{}

        mockClient.On("StartAppAssessment", ctx, input).Return(output, nil)

        result, err := mockClient.StartAppAssessment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &resiliencehub.TagResourceInput{}
        output := &resiliencehub.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &resiliencehub.UntagResourceInput{}
        output := &resiliencehub.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateApp", func(t *testing.T) {
        input := &resiliencehub.UpdateAppInput{}
        output := &resiliencehub.UpdateAppOutput{}

        mockClient.On("UpdateApp", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateApp(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateAppVersion", func(t *testing.T) {
        input := &resiliencehub.UpdateAppVersionInput{}
        output := &resiliencehub.UpdateAppVersionOutput{}

        mockClient.On("UpdateAppVersion", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateAppVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateAppVersionAppComponent", func(t *testing.T) {
        input := &resiliencehub.UpdateAppVersionAppComponentInput{}
        output := &resiliencehub.UpdateAppVersionAppComponentOutput{}

        mockClient.On("UpdateAppVersionAppComponent", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateAppVersionAppComponent(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateAppVersionResource", func(t *testing.T) {
        input := &resiliencehub.UpdateAppVersionResourceInput{}
        output := &resiliencehub.UpdateAppVersionResourceOutput{}

        mockClient.On("UpdateAppVersionResource", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateAppVersionResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateResiliencyPolicy", func(t *testing.T) {
        input := &resiliencehub.UpdateResiliencyPolicyInput{}
        output := &resiliencehub.UpdateResiliencyPolicyOutput{}

        mockClient.On("UpdateResiliencyPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateResiliencyPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
