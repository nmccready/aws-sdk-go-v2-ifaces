package inspector_test

// tests for the inspector service interface for this ../../../service/inspector/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/inspector"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/inspector/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/inspector/inspector_iface"
	"github.com/stretchr/testify/assert"
)

func TestInspectorServiceCanBeMocked(t *testing.T) {
	var iface inspector_iface.IClient
	iface = &inspector.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := inspector.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAddAttributesToFindings", func(t *testing.T) {
        input := &inspector.AddAttributesToFindingsInput{}
        output := &inspector.AddAttributesToFindingsOutput{}

        mockClient.On("AddAttributesToFindings", ctx, input).Return(output, nil)

        result, err := mockClient.AddAttributesToFindings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateAssessmentTarget", func(t *testing.T) {
        input := &inspector.CreateAssessmentTargetInput{}
        output := &inspector.CreateAssessmentTargetOutput{}

        mockClient.On("CreateAssessmentTarget", ctx, input).Return(output, nil)

        result, err := mockClient.CreateAssessmentTarget(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateAssessmentTemplate", func(t *testing.T) {
        input := &inspector.CreateAssessmentTemplateInput{}
        output := &inspector.CreateAssessmentTemplateOutput{}

        mockClient.On("CreateAssessmentTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.CreateAssessmentTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateExclusionsPreview", func(t *testing.T) {
        input := &inspector.CreateExclusionsPreviewInput{}
        output := &inspector.CreateExclusionsPreviewOutput{}

        mockClient.On("CreateExclusionsPreview", ctx, input).Return(output, nil)

        result, err := mockClient.CreateExclusionsPreview(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateResourceGroup", func(t *testing.T) {
        input := &inspector.CreateResourceGroupInput{}
        output := &inspector.CreateResourceGroupOutput{}

        mockClient.On("CreateResourceGroup", ctx, input).Return(output, nil)

        result, err := mockClient.CreateResourceGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteAssessmentRun", func(t *testing.T) {
        input := &inspector.DeleteAssessmentRunInput{}
        output := &inspector.DeleteAssessmentRunOutput{}

        mockClient.On("DeleteAssessmentRun", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteAssessmentRun(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteAssessmentTarget", func(t *testing.T) {
        input := &inspector.DeleteAssessmentTargetInput{}
        output := &inspector.DeleteAssessmentTargetOutput{}

        mockClient.On("DeleteAssessmentTarget", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteAssessmentTarget(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteAssessmentTemplate", func(t *testing.T) {
        input := &inspector.DeleteAssessmentTemplateInput{}
        output := &inspector.DeleteAssessmentTemplateOutput{}

        mockClient.On("DeleteAssessmentTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteAssessmentTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeAssessmentRuns", func(t *testing.T) {
        input := &inspector.DescribeAssessmentRunsInput{}
        output := &inspector.DescribeAssessmentRunsOutput{}

        mockClient.On("DescribeAssessmentRuns", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeAssessmentRuns(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeAssessmentTargets", func(t *testing.T) {
        input := &inspector.DescribeAssessmentTargetsInput{}
        output := &inspector.DescribeAssessmentTargetsOutput{}

        mockClient.On("DescribeAssessmentTargets", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeAssessmentTargets(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeAssessmentTemplates", func(t *testing.T) {
        input := &inspector.DescribeAssessmentTemplatesInput{}
        output := &inspector.DescribeAssessmentTemplatesOutput{}

        mockClient.On("DescribeAssessmentTemplates", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeAssessmentTemplates(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeCrossAccountAccessRole", func(t *testing.T) {
        input := &inspector.DescribeCrossAccountAccessRoleInput{}
        output := &inspector.DescribeCrossAccountAccessRoleOutput{}

        mockClient.On("DescribeCrossAccountAccessRole", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeCrossAccountAccessRole(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeExclusions", func(t *testing.T) {
        input := &inspector.DescribeExclusionsInput{}
        output := &inspector.DescribeExclusionsOutput{}

        mockClient.On("DescribeExclusions", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeExclusions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeFindings", func(t *testing.T) {
        input := &inspector.DescribeFindingsInput{}
        output := &inspector.DescribeFindingsOutput{}

        mockClient.On("DescribeFindings", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeFindings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeResourceGroups", func(t *testing.T) {
        input := &inspector.DescribeResourceGroupsInput{}
        output := &inspector.DescribeResourceGroupsOutput{}

        mockClient.On("DescribeResourceGroups", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeResourceGroups(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeRulesPackages", func(t *testing.T) {
        input := &inspector.DescribeRulesPackagesInput{}
        output := &inspector.DescribeRulesPackagesOutput{}

        mockClient.On("DescribeRulesPackages", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeRulesPackages(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetAssessmentReport", func(t *testing.T) {
        input := &inspector.GetAssessmentReportInput{}
        output := &inspector.GetAssessmentReportOutput{}

        mockClient.On("GetAssessmentReport", ctx, input).Return(output, nil)

        result, err := mockClient.GetAssessmentReport(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetExclusionsPreview", func(t *testing.T) {
        input := &inspector.GetExclusionsPreviewInput{}
        output := &inspector.GetExclusionsPreviewOutput{}

        mockClient.On("GetExclusionsPreview", ctx, input).Return(output, nil)

        result, err := mockClient.GetExclusionsPreview(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetTelemetryMetadata", func(t *testing.T) {
        input := &inspector.GetTelemetryMetadataInput{}
        output := &inspector.GetTelemetryMetadataOutput{}

        mockClient.On("GetTelemetryMetadata", ctx, input).Return(output, nil)

        result, err := mockClient.GetTelemetryMetadata(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAssessmentRunAgents", func(t *testing.T) {
        input := &inspector.ListAssessmentRunAgentsInput{}
        output := &inspector.ListAssessmentRunAgentsOutput{}

        mockClient.On("ListAssessmentRunAgents", ctx, input).Return(output, nil)

        result, err := mockClient.ListAssessmentRunAgents(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAssessmentRuns", func(t *testing.T) {
        input := &inspector.ListAssessmentRunsInput{}
        output := &inspector.ListAssessmentRunsOutput{}

        mockClient.On("ListAssessmentRuns", ctx, input).Return(output, nil)

        result, err := mockClient.ListAssessmentRuns(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAssessmentTargets", func(t *testing.T) {
        input := &inspector.ListAssessmentTargetsInput{}
        output := &inspector.ListAssessmentTargetsOutput{}

        mockClient.On("ListAssessmentTargets", ctx, input).Return(output, nil)

        result, err := mockClient.ListAssessmentTargets(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAssessmentTemplates", func(t *testing.T) {
        input := &inspector.ListAssessmentTemplatesInput{}
        output := &inspector.ListAssessmentTemplatesOutput{}

        mockClient.On("ListAssessmentTemplates", ctx, input).Return(output, nil)

        result, err := mockClient.ListAssessmentTemplates(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListEventSubscriptions", func(t *testing.T) {
        input := &inspector.ListEventSubscriptionsInput{}
        output := &inspector.ListEventSubscriptionsOutput{}

        mockClient.On("ListEventSubscriptions", ctx, input).Return(output, nil)

        result, err := mockClient.ListEventSubscriptions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListExclusions", func(t *testing.T) {
        input := &inspector.ListExclusionsInput{}
        output := &inspector.ListExclusionsOutput{}

        mockClient.On("ListExclusions", ctx, input).Return(output, nil)

        result, err := mockClient.ListExclusions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListFindings", func(t *testing.T) {
        input := &inspector.ListFindingsInput{}
        output := &inspector.ListFindingsOutput{}

        mockClient.On("ListFindings", ctx, input).Return(output, nil)

        result, err := mockClient.ListFindings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListRulesPackages", func(t *testing.T) {
        input := &inspector.ListRulesPackagesInput{}
        output := &inspector.ListRulesPackagesOutput{}

        mockClient.On("ListRulesPackages", ctx, input).Return(output, nil)

        result, err := mockClient.ListRulesPackages(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &inspector.ListTagsForResourceInput{}
        output := &inspector.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPreviewAgents", func(t *testing.T) {
        input := &inspector.PreviewAgentsInput{}
        output := &inspector.PreviewAgentsOutput{}

        mockClient.On("PreviewAgents", ctx, input).Return(output, nil)

        result, err := mockClient.PreviewAgents(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRegisterCrossAccountAccessRole", func(t *testing.T) {
        input := &inspector.RegisterCrossAccountAccessRoleInput{}
        output := &inspector.RegisterCrossAccountAccessRoleOutput{}

        mockClient.On("RegisterCrossAccountAccessRole", ctx, input).Return(output, nil)

        result, err := mockClient.RegisterCrossAccountAccessRole(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRemoveAttributesFromFindings", func(t *testing.T) {
        input := &inspector.RemoveAttributesFromFindingsInput{}
        output := &inspector.RemoveAttributesFromFindingsOutput{}

        mockClient.On("RemoveAttributesFromFindings", ctx, input).Return(output, nil)

        result, err := mockClient.RemoveAttributesFromFindings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSetTagsForResource", func(t *testing.T) {
        input := &inspector.SetTagsForResourceInput{}
        output := &inspector.SetTagsForResourceOutput{}

        mockClient.On("SetTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.SetTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartAssessmentRun", func(t *testing.T) {
        input := &inspector.StartAssessmentRunInput{}
        output := &inspector.StartAssessmentRunOutput{}

        mockClient.On("StartAssessmentRun", ctx, input).Return(output, nil)

        result, err := mockClient.StartAssessmentRun(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopAssessmentRun", func(t *testing.T) {
        input := &inspector.StopAssessmentRunInput{}
        output := &inspector.StopAssessmentRunOutput{}

        mockClient.On("StopAssessmentRun", ctx, input).Return(output, nil)

        result, err := mockClient.StopAssessmentRun(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSubscribeToEvent", func(t *testing.T) {
        input := &inspector.SubscribeToEventInput{}
        output := &inspector.SubscribeToEventOutput{}

        mockClient.On("SubscribeToEvent", ctx, input).Return(output, nil)

        result, err := mockClient.SubscribeToEvent(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUnsubscribeFromEvent", func(t *testing.T) {
        input := &inspector.UnsubscribeFromEventInput{}
        output := &inspector.UnsubscribeFromEventOutput{}

        mockClient.On("UnsubscribeFromEvent", ctx, input).Return(output, nil)

        result, err := mockClient.UnsubscribeFromEvent(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateAssessmentTarget", func(t *testing.T) {
        input := &inspector.UpdateAssessmentTargetInput{}
        output := &inspector.UpdateAssessmentTargetOutput{}

        mockClient.On("UpdateAssessmentTarget", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateAssessmentTarget(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
