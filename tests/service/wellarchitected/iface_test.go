package wellarchitected_test

// tests for the wellarchitected service interface for this ../../../service/wellarchitected/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/wellarchitected"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/wellarchitected/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/wellarchitected/wellarchitected_iface"
	"github.com/stretchr/testify/assert"
)

func TestWellarchitectedServiceCanBeMocked(t *testing.T) {
	var iface wellarchitected_iface.IClient
	iface = &wellarchitected.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := wellarchitected.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateLenses", func(t *testing.T) {
        input := &wellarchitected.AssociateLensesInput{}
        output := &wellarchitected.AssociateLensesOutput{}

        mockClient.On("AssociateLenses", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateLenses(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateProfiles", func(t *testing.T) {
        input := &wellarchitected.AssociateProfilesInput{}
        output := &wellarchitected.AssociateProfilesOutput{}

        mockClient.On("AssociateProfiles", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateProfiles(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateLensShare", func(t *testing.T) {
        input := &wellarchitected.CreateLensShareInput{}
        output := &wellarchitected.CreateLensShareOutput{}

        mockClient.On("CreateLensShare", ctx, input).Return(output, nil)

        result, err := mockClient.CreateLensShare(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateLensVersion", func(t *testing.T) {
        input := &wellarchitected.CreateLensVersionInput{}
        output := &wellarchitected.CreateLensVersionOutput{}

        mockClient.On("CreateLensVersion", ctx, input).Return(output, nil)

        result, err := mockClient.CreateLensVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateMilestone", func(t *testing.T) {
        input := &wellarchitected.CreateMilestoneInput{}
        output := &wellarchitected.CreateMilestoneOutput{}

        mockClient.On("CreateMilestone", ctx, input).Return(output, nil)

        result, err := mockClient.CreateMilestone(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateProfile", func(t *testing.T) {
        input := &wellarchitected.CreateProfileInput{}
        output := &wellarchitected.CreateProfileOutput{}

        mockClient.On("CreateProfile", ctx, input).Return(output, nil)

        result, err := mockClient.CreateProfile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateProfileShare", func(t *testing.T) {
        input := &wellarchitected.CreateProfileShareInput{}
        output := &wellarchitected.CreateProfileShareOutput{}

        mockClient.On("CreateProfileShare", ctx, input).Return(output, nil)

        result, err := mockClient.CreateProfileShare(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateReviewTemplate", func(t *testing.T) {
        input := &wellarchitected.CreateReviewTemplateInput{}
        output := &wellarchitected.CreateReviewTemplateOutput{}

        mockClient.On("CreateReviewTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.CreateReviewTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateTemplateShare", func(t *testing.T) {
        input := &wellarchitected.CreateTemplateShareInput{}
        output := &wellarchitected.CreateTemplateShareOutput{}

        mockClient.On("CreateTemplateShare", ctx, input).Return(output, nil)

        result, err := mockClient.CreateTemplateShare(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateWorkload", func(t *testing.T) {
        input := &wellarchitected.CreateWorkloadInput{}
        output := &wellarchitected.CreateWorkloadOutput{}

        mockClient.On("CreateWorkload", ctx, input).Return(output, nil)

        result, err := mockClient.CreateWorkload(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateWorkloadShare", func(t *testing.T) {
        input := &wellarchitected.CreateWorkloadShareInput{}
        output := &wellarchitected.CreateWorkloadShareOutput{}

        mockClient.On("CreateWorkloadShare", ctx, input).Return(output, nil)

        result, err := mockClient.CreateWorkloadShare(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteLens", func(t *testing.T) {
        input := &wellarchitected.DeleteLensInput{}
        output := &wellarchitected.DeleteLensOutput{}

        mockClient.On("DeleteLens", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteLens(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteLensShare", func(t *testing.T) {
        input := &wellarchitected.DeleteLensShareInput{}
        output := &wellarchitected.DeleteLensShareOutput{}

        mockClient.On("DeleteLensShare", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteLensShare(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteProfile", func(t *testing.T) {
        input := &wellarchitected.DeleteProfileInput{}
        output := &wellarchitected.DeleteProfileOutput{}

        mockClient.On("DeleteProfile", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteProfile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteProfileShare", func(t *testing.T) {
        input := &wellarchitected.DeleteProfileShareInput{}
        output := &wellarchitected.DeleteProfileShareOutput{}

        mockClient.On("DeleteProfileShare", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteProfileShare(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteReviewTemplate", func(t *testing.T) {
        input := &wellarchitected.DeleteReviewTemplateInput{}
        output := &wellarchitected.DeleteReviewTemplateOutput{}

        mockClient.On("DeleteReviewTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteReviewTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteTemplateShare", func(t *testing.T) {
        input := &wellarchitected.DeleteTemplateShareInput{}
        output := &wellarchitected.DeleteTemplateShareOutput{}

        mockClient.On("DeleteTemplateShare", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteTemplateShare(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteWorkload", func(t *testing.T) {
        input := &wellarchitected.DeleteWorkloadInput{}
        output := &wellarchitected.DeleteWorkloadOutput{}

        mockClient.On("DeleteWorkload", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteWorkload(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteWorkloadShare", func(t *testing.T) {
        input := &wellarchitected.DeleteWorkloadShareInput{}
        output := &wellarchitected.DeleteWorkloadShareOutput{}

        mockClient.On("DeleteWorkloadShare", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteWorkloadShare(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateLenses", func(t *testing.T) {
        input := &wellarchitected.DisassociateLensesInput{}
        output := &wellarchitected.DisassociateLensesOutput{}

        mockClient.On("DisassociateLenses", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateLenses(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateProfiles", func(t *testing.T) {
        input := &wellarchitected.DisassociateProfilesInput{}
        output := &wellarchitected.DisassociateProfilesOutput{}

        mockClient.On("DisassociateProfiles", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateProfiles(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestExportLens", func(t *testing.T) {
        input := &wellarchitected.ExportLensInput{}
        output := &wellarchitected.ExportLensOutput{}

        mockClient.On("ExportLens", ctx, input).Return(output, nil)

        result, err := mockClient.ExportLens(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetAnswer", func(t *testing.T) {
        input := &wellarchitected.GetAnswerInput{}
        output := &wellarchitected.GetAnswerOutput{}

        mockClient.On("GetAnswer", ctx, input).Return(output, nil)

        result, err := mockClient.GetAnswer(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetConsolidatedReport", func(t *testing.T) {
        input := &wellarchitected.GetConsolidatedReportInput{}
        output := &wellarchitected.GetConsolidatedReportOutput{}

        mockClient.On("GetConsolidatedReport", ctx, input).Return(output, nil)

        result, err := mockClient.GetConsolidatedReport(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetGlobalSettings", func(t *testing.T) {
        input := &wellarchitected.GetGlobalSettingsInput{}
        output := &wellarchitected.GetGlobalSettingsOutput{}

        mockClient.On("GetGlobalSettings", ctx, input).Return(output, nil)

        result, err := mockClient.GetGlobalSettings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetLens", func(t *testing.T) {
        input := &wellarchitected.GetLensInput{}
        output := &wellarchitected.GetLensOutput{}

        mockClient.On("GetLens", ctx, input).Return(output, nil)

        result, err := mockClient.GetLens(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetLensReview", func(t *testing.T) {
        input := &wellarchitected.GetLensReviewInput{}
        output := &wellarchitected.GetLensReviewOutput{}

        mockClient.On("GetLensReview", ctx, input).Return(output, nil)

        result, err := mockClient.GetLensReview(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetLensReviewReport", func(t *testing.T) {
        input := &wellarchitected.GetLensReviewReportInput{}
        output := &wellarchitected.GetLensReviewReportOutput{}

        mockClient.On("GetLensReviewReport", ctx, input).Return(output, nil)

        result, err := mockClient.GetLensReviewReport(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetLensVersionDifference", func(t *testing.T) {
        input := &wellarchitected.GetLensVersionDifferenceInput{}
        output := &wellarchitected.GetLensVersionDifferenceOutput{}

        mockClient.On("GetLensVersionDifference", ctx, input).Return(output, nil)

        result, err := mockClient.GetLensVersionDifference(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetMilestone", func(t *testing.T) {
        input := &wellarchitected.GetMilestoneInput{}
        output := &wellarchitected.GetMilestoneOutput{}

        mockClient.On("GetMilestone", ctx, input).Return(output, nil)

        result, err := mockClient.GetMilestone(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetProfile", func(t *testing.T) {
        input := &wellarchitected.GetProfileInput{}
        output := &wellarchitected.GetProfileOutput{}

        mockClient.On("GetProfile", ctx, input).Return(output, nil)

        result, err := mockClient.GetProfile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetProfileTemplate", func(t *testing.T) {
        input := &wellarchitected.GetProfileTemplateInput{}
        output := &wellarchitected.GetProfileTemplateOutput{}

        mockClient.On("GetProfileTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.GetProfileTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetReviewTemplate", func(t *testing.T) {
        input := &wellarchitected.GetReviewTemplateInput{}
        output := &wellarchitected.GetReviewTemplateOutput{}

        mockClient.On("GetReviewTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.GetReviewTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetReviewTemplateAnswer", func(t *testing.T) {
        input := &wellarchitected.GetReviewTemplateAnswerInput{}
        output := &wellarchitected.GetReviewTemplateAnswerOutput{}

        mockClient.On("GetReviewTemplateAnswer", ctx, input).Return(output, nil)

        result, err := mockClient.GetReviewTemplateAnswer(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetReviewTemplateLensReview", func(t *testing.T) {
        input := &wellarchitected.GetReviewTemplateLensReviewInput{}
        output := &wellarchitected.GetReviewTemplateLensReviewOutput{}

        mockClient.On("GetReviewTemplateLensReview", ctx, input).Return(output, nil)

        result, err := mockClient.GetReviewTemplateLensReview(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetWorkload", func(t *testing.T) {
        input := &wellarchitected.GetWorkloadInput{}
        output := &wellarchitected.GetWorkloadOutput{}

        mockClient.On("GetWorkload", ctx, input).Return(output, nil)

        result, err := mockClient.GetWorkload(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestImportLens", func(t *testing.T) {
        input := &wellarchitected.ImportLensInput{}
        output := &wellarchitected.ImportLensOutput{}

        mockClient.On("ImportLens", ctx, input).Return(output, nil)

        result, err := mockClient.ImportLens(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAnswers", func(t *testing.T) {
        input := &wellarchitected.ListAnswersInput{}
        output := &wellarchitected.ListAnswersOutput{}

        mockClient.On("ListAnswers", ctx, input).Return(output, nil)

        result, err := mockClient.ListAnswers(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListCheckDetails", func(t *testing.T) {
        input := &wellarchitected.ListCheckDetailsInput{}
        output := &wellarchitected.ListCheckDetailsOutput{}

        mockClient.On("ListCheckDetails", ctx, input).Return(output, nil)

        result, err := mockClient.ListCheckDetails(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListCheckSummaries", func(t *testing.T) {
        input := &wellarchitected.ListCheckSummariesInput{}
        output := &wellarchitected.ListCheckSummariesOutput{}

        mockClient.On("ListCheckSummaries", ctx, input).Return(output, nil)

        result, err := mockClient.ListCheckSummaries(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListLensReviewImprovements", func(t *testing.T) {
        input := &wellarchitected.ListLensReviewImprovementsInput{}
        output := &wellarchitected.ListLensReviewImprovementsOutput{}

        mockClient.On("ListLensReviewImprovements", ctx, input).Return(output, nil)

        result, err := mockClient.ListLensReviewImprovements(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListLensReviews", func(t *testing.T) {
        input := &wellarchitected.ListLensReviewsInput{}
        output := &wellarchitected.ListLensReviewsOutput{}

        mockClient.On("ListLensReviews", ctx, input).Return(output, nil)

        result, err := mockClient.ListLensReviews(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListLensShares", func(t *testing.T) {
        input := &wellarchitected.ListLensSharesInput{}
        output := &wellarchitected.ListLensSharesOutput{}

        mockClient.On("ListLensShares", ctx, input).Return(output, nil)

        result, err := mockClient.ListLensShares(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListLenses", func(t *testing.T) {
        input := &wellarchitected.ListLensesInput{}
        output := &wellarchitected.ListLensesOutput{}

        mockClient.On("ListLenses", ctx, input).Return(output, nil)

        result, err := mockClient.ListLenses(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListMilestones", func(t *testing.T) {
        input := &wellarchitected.ListMilestonesInput{}
        output := &wellarchitected.ListMilestonesOutput{}

        mockClient.On("ListMilestones", ctx, input).Return(output, nil)

        result, err := mockClient.ListMilestones(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListNotifications", func(t *testing.T) {
        input := &wellarchitected.ListNotificationsInput{}
        output := &wellarchitected.ListNotificationsOutput{}

        mockClient.On("ListNotifications", ctx, input).Return(output, nil)

        result, err := mockClient.ListNotifications(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListProfileNotifications", func(t *testing.T) {
        input := &wellarchitected.ListProfileNotificationsInput{}
        output := &wellarchitected.ListProfileNotificationsOutput{}

        mockClient.On("ListProfileNotifications", ctx, input).Return(output, nil)

        result, err := mockClient.ListProfileNotifications(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListProfileShares", func(t *testing.T) {
        input := &wellarchitected.ListProfileSharesInput{}
        output := &wellarchitected.ListProfileSharesOutput{}

        mockClient.On("ListProfileShares", ctx, input).Return(output, nil)

        result, err := mockClient.ListProfileShares(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListProfiles", func(t *testing.T) {
        input := &wellarchitected.ListProfilesInput{}
        output := &wellarchitected.ListProfilesOutput{}

        mockClient.On("ListProfiles", ctx, input).Return(output, nil)

        result, err := mockClient.ListProfiles(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListReviewTemplateAnswers", func(t *testing.T) {
        input := &wellarchitected.ListReviewTemplateAnswersInput{}
        output := &wellarchitected.ListReviewTemplateAnswersOutput{}

        mockClient.On("ListReviewTemplateAnswers", ctx, input).Return(output, nil)

        result, err := mockClient.ListReviewTemplateAnswers(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListReviewTemplates", func(t *testing.T) {
        input := &wellarchitected.ListReviewTemplatesInput{}
        output := &wellarchitected.ListReviewTemplatesOutput{}

        mockClient.On("ListReviewTemplates", ctx, input).Return(output, nil)

        result, err := mockClient.ListReviewTemplates(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListShareInvitations", func(t *testing.T) {
        input := &wellarchitected.ListShareInvitationsInput{}
        output := &wellarchitected.ListShareInvitationsOutput{}

        mockClient.On("ListShareInvitations", ctx, input).Return(output, nil)

        result, err := mockClient.ListShareInvitations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &wellarchitected.ListTagsForResourceInput{}
        output := &wellarchitected.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTemplateShares", func(t *testing.T) {
        input := &wellarchitected.ListTemplateSharesInput{}
        output := &wellarchitected.ListTemplateSharesOutput{}

        mockClient.On("ListTemplateShares", ctx, input).Return(output, nil)

        result, err := mockClient.ListTemplateShares(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListWorkloadShares", func(t *testing.T) {
        input := &wellarchitected.ListWorkloadSharesInput{}
        output := &wellarchitected.ListWorkloadSharesOutput{}

        mockClient.On("ListWorkloadShares", ctx, input).Return(output, nil)

        result, err := mockClient.ListWorkloadShares(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListWorkloads", func(t *testing.T) {
        input := &wellarchitected.ListWorkloadsInput{}
        output := &wellarchitected.ListWorkloadsOutput{}

        mockClient.On("ListWorkloads", ctx, input).Return(output, nil)

        result, err := mockClient.ListWorkloads(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &wellarchitected.TagResourceInput{}
        output := &wellarchitected.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &wellarchitected.UntagResourceInput{}
        output := &wellarchitected.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateAnswer", func(t *testing.T) {
        input := &wellarchitected.UpdateAnswerInput{}
        output := &wellarchitected.UpdateAnswerOutput{}

        mockClient.On("UpdateAnswer", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateAnswer(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateGlobalSettings", func(t *testing.T) {
        input := &wellarchitected.UpdateGlobalSettingsInput{}
        output := &wellarchitected.UpdateGlobalSettingsOutput{}

        mockClient.On("UpdateGlobalSettings", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateGlobalSettings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateIntegration", func(t *testing.T) {
        input := &wellarchitected.UpdateIntegrationInput{}
        output := &wellarchitected.UpdateIntegrationOutput{}

        mockClient.On("UpdateIntegration", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateIntegration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateLensReview", func(t *testing.T) {
        input := &wellarchitected.UpdateLensReviewInput{}
        output := &wellarchitected.UpdateLensReviewOutput{}

        mockClient.On("UpdateLensReview", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateLensReview(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateProfile", func(t *testing.T) {
        input := &wellarchitected.UpdateProfileInput{}
        output := &wellarchitected.UpdateProfileOutput{}

        mockClient.On("UpdateProfile", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateProfile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateReviewTemplate", func(t *testing.T) {
        input := &wellarchitected.UpdateReviewTemplateInput{}
        output := &wellarchitected.UpdateReviewTemplateOutput{}

        mockClient.On("UpdateReviewTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateReviewTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateReviewTemplateAnswer", func(t *testing.T) {
        input := &wellarchitected.UpdateReviewTemplateAnswerInput{}
        output := &wellarchitected.UpdateReviewTemplateAnswerOutput{}

        mockClient.On("UpdateReviewTemplateAnswer", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateReviewTemplateAnswer(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateReviewTemplateLensReview", func(t *testing.T) {
        input := &wellarchitected.UpdateReviewTemplateLensReviewInput{}
        output := &wellarchitected.UpdateReviewTemplateLensReviewOutput{}

        mockClient.On("UpdateReviewTemplateLensReview", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateReviewTemplateLensReview(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateShareInvitation", func(t *testing.T) {
        input := &wellarchitected.UpdateShareInvitationInput{}
        output := &wellarchitected.UpdateShareInvitationOutput{}

        mockClient.On("UpdateShareInvitation", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateShareInvitation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateWorkload", func(t *testing.T) {
        input := &wellarchitected.UpdateWorkloadInput{}
        output := &wellarchitected.UpdateWorkloadOutput{}

        mockClient.On("UpdateWorkload", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateWorkload(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateWorkloadShare", func(t *testing.T) {
        input := &wellarchitected.UpdateWorkloadShareInput{}
        output := &wellarchitected.UpdateWorkloadShareOutput{}

        mockClient.On("UpdateWorkloadShare", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateWorkloadShare(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpgradeLensReview", func(t *testing.T) {
        input := &wellarchitected.UpgradeLensReviewInput{}
        output := &wellarchitected.UpgradeLensReviewOutput{}

        mockClient.On("UpgradeLensReview", ctx, input).Return(output, nil)

        result, err := mockClient.UpgradeLensReview(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpgradeProfileVersion", func(t *testing.T) {
        input := &wellarchitected.UpgradeProfileVersionInput{}
        output := &wellarchitected.UpgradeProfileVersionOutput{}

        mockClient.On("UpgradeProfileVersion", ctx, input).Return(output, nil)

        result, err := mockClient.UpgradeProfileVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpgradeReviewTemplateLensReview", func(t *testing.T) {
        input := &wellarchitected.UpgradeReviewTemplateLensReviewInput{}
        output := &wellarchitected.UpgradeReviewTemplateLensReviewOutput{}

        mockClient.On("UpgradeReviewTemplateLensReview", ctx, input).Return(output, nil)

        result, err := mockClient.UpgradeReviewTemplateLensReview(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
