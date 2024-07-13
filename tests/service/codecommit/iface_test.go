package codecommit_test

// tests for the codecommit service interface for this ../../../service/codecommit/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/codecommit"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/codecommit/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/codecommit/codecommit_iface"
	"github.com/stretchr/testify/assert"
)

func TestCodecommitServiceCanBeMocked(t *testing.T) {
	var iface codecommit_iface.IClient
	iface = &codecommit.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := codecommit.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateApprovalRuleTemplateWithRepository", func(t *testing.T) {
        input := &codecommit.AssociateApprovalRuleTemplateWithRepositoryInput{}
        output := &codecommit.AssociateApprovalRuleTemplateWithRepositoryOutput{}

        mockClient.On("AssociateApprovalRuleTemplateWithRepository", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateApprovalRuleTemplateWithRepository(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchAssociateApprovalRuleTemplateWithRepositories", func(t *testing.T) {
        input := &codecommit.BatchAssociateApprovalRuleTemplateWithRepositoriesInput{}
        output := &codecommit.BatchAssociateApprovalRuleTemplateWithRepositoriesOutput{}

        mockClient.On("BatchAssociateApprovalRuleTemplateWithRepositories", ctx, input).Return(output, nil)

        result, err := mockClient.BatchAssociateApprovalRuleTemplateWithRepositories(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchDescribeMergeConflicts", func(t *testing.T) {
        input := &codecommit.BatchDescribeMergeConflictsInput{}
        output := &codecommit.BatchDescribeMergeConflictsOutput{}

        mockClient.On("BatchDescribeMergeConflicts", ctx, input).Return(output, nil)

        result, err := mockClient.BatchDescribeMergeConflicts(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchDisassociateApprovalRuleTemplateFromRepositories", func(t *testing.T) {
        input := &codecommit.BatchDisassociateApprovalRuleTemplateFromRepositoriesInput{}
        output := &codecommit.BatchDisassociateApprovalRuleTemplateFromRepositoriesOutput{}

        mockClient.On("BatchDisassociateApprovalRuleTemplateFromRepositories", ctx, input).Return(output, nil)

        result, err := mockClient.BatchDisassociateApprovalRuleTemplateFromRepositories(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchGetCommits", func(t *testing.T) {
        input := &codecommit.BatchGetCommitsInput{}
        output := &codecommit.BatchGetCommitsOutput{}

        mockClient.On("BatchGetCommits", ctx, input).Return(output, nil)

        result, err := mockClient.BatchGetCommits(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchGetRepositories", func(t *testing.T) {
        input := &codecommit.BatchGetRepositoriesInput{}
        output := &codecommit.BatchGetRepositoriesOutput{}

        mockClient.On("BatchGetRepositories", ctx, input).Return(output, nil)

        result, err := mockClient.BatchGetRepositories(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateApprovalRuleTemplate", func(t *testing.T) {
        input := &codecommit.CreateApprovalRuleTemplateInput{}
        output := &codecommit.CreateApprovalRuleTemplateOutput{}

        mockClient.On("CreateApprovalRuleTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.CreateApprovalRuleTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateBranch", func(t *testing.T) {
        input := &codecommit.CreateBranchInput{}
        output := &codecommit.CreateBranchOutput{}

        mockClient.On("CreateBranch", ctx, input).Return(output, nil)

        result, err := mockClient.CreateBranch(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateCommit", func(t *testing.T) {
        input := &codecommit.CreateCommitInput{}
        output := &codecommit.CreateCommitOutput{}

        mockClient.On("CreateCommit", ctx, input).Return(output, nil)

        result, err := mockClient.CreateCommit(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreatePullRequest", func(t *testing.T) {
        input := &codecommit.CreatePullRequestInput{}
        output := &codecommit.CreatePullRequestOutput{}

        mockClient.On("CreatePullRequest", ctx, input).Return(output, nil)

        result, err := mockClient.CreatePullRequest(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreatePullRequestApprovalRule", func(t *testing.T) {
        input := &codecommit.CreatePullRequestApprovalRuleInput{}
        output := &codecommit.CreatePullRequestApprovalRuleOutput{}

        mockClient.On("CreatePullRequestApprovalRule", ctx, input).Return(output, nil)

        result, err := mockClient.CreatePullRequestApprovalRule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateRepository", func(t *testing.T) {
        input := &codecommit.CreateRepositoryInput{}
        output := &codecommit.CreateRepositoryOutput{}

        mockClient.On("CreateRepository", ctx, input).Return(output, nil)

        result, err := mockClient.CreateRepository(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateUnreferencedMergeCommit", func(t *testing.T) {
        input := &codecommit.CreateUnreferencedMergeCommitInput{}
        output := &codecommit.CreateUnreferencedMergeCommitOutput{}

        mockClient.On("CreateUnreferencedMergeCommit", ctx, input).Return(output, nil)

        result, err := mockClient.CreateUnreferencedMergeCommit(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteApprovalRuleTemplate", func(t *testing.T) {
        input := &codecommit.DeleteApprovalRuleTemplateInput{}
        output := &codecommit.DeleteApprovalRuleTemplateOutput{}

        mockClient.On("DeleteApprovalRuleTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteApprovalRuleTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteBranch", func(t *testing.T) {
        input := &codecommit.DeleteBranchInput{}
        output := &codecommit.DeleteBranchOutput{}

        mockClient.On("DeleteBranch", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteBranch(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteCommentContent", func(t *testing.T) {
        input := &codecommit.DeleteCommentContentInput{}
        output := &codecommit.DeleteCommentContentOutput{}

        mockClient.On("DeleteCommentContent", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteCommentContent(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteFile", func(t *testing.T) {
        input := &codecommit.DeleteFileInput{}
        output := &codecommit.DeleteFileOutput{}

        mockClient.On("DeleteFile", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteFile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeletePullRequestApprovalRule", func(t *testing.T) {
        input := &codecommit.DeletePullRequestApprovalRuleInput{}
        output := &codecommit.DeletePullRequestApprovalRuleOutput{}

        mockClient.On("DeletePullRequestApprovalRule", ctx, input).Return(output, nil)

        result, err := mockClient.DeletePullRequestApprovalRule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteRepository", func(t *testing.T) {
        input := &codecommit.DeleteRepositoryInput{}
        output := &codecommit.DeleteRepositoryOutput{}

        mockClient.On("DeleteRepository", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteRepository(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeMergeConflicts", func(t *testing.T) {
        input := &codecommit.DescribeMergeConflictsInput{}
        output := &codecommit.DescribeMergeConflictsOutput{}

        mockClient.On("DescribeMergeConflicts", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeMergeConflicts(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribePullRequestEvents", func(t *testing.T) {
        input := &codecommit.DescribePullRequestEventsInput{}
        output := &codecommit.DescribePullRequestEventsOutput{}

        mockClient.On("DescribePullRequestEvents", ctx, input).Return(output, nil)

        result, err := mockClient.DescribePullRequestEvents(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateApprovalRuleTemplateFromRepository", func(t *testing.T) {
        input := &codecommit.DisassociateApprovalRuleTemplateFromRepositoryInput{}
        output := &codecommit.DisassociateApprovalRuleTemplateFromRepositoryOutput{}

        mockClient.On("DisassociateApprovalRuleTemplateFromRepository", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateApprovalRuleTemplateFromRepository(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestEvaluatePullRequestApprovalRules", func(t *testing.T) {
        input := &codecommit.EvaluatePullRequestApprovalRulesInput{}
        output := &codecommit.EvaluatePullRequestApprovalRulesOutput{}

        mockClient.On("EvaluatePullRequestApprovalRules", ctx, input).Return(output, nil)

        result, err := mockClient.EvaluatePullRequestApprovalRules(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetApprovalRuleTemplate", func(t *testing.T) {
        input := &codecommit.GetApprovalRuleTemplateInput{}
        output := &codecommit.GetApprovalRuleTemplateOutput{}

        mockClient.On("GetApprovalRuleTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.GetApprovalRuleTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetBlob", func(t *testing.T) {
        input := &codecommit.GetBlobInput{}
        output := &codecommit.GetBlobOutput{}

        mockClient.On("GetBlob", ctx, input).Return(output, nil)

        result, err := mockClient.GetBlob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetBranch", func(t *testing.T) {
        input := &codecommit.GetBranchInput{}
        output := &codecommit.GetBranchOutput{}

        mockClient.On("GetBranch", ctx, input).Return(output, nil)

        result, err := mockClient.GetBranch(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetComment", func(t *testing.T) {
        input := &codecommit.GetCommentInput{}
        output := &codecommit.GetCommentOutput{}

        mockClient.On("GetComment", ctx, input).Return(output, nil)

        result, err := mockClient.GetComment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetCommentReactions", func(t *testing.T) {
        input := &codecommit.GetCommentReactionsInput{}
        output := &codecommit.GetCommentReactionsOutput{}

        mockClient.On("GetCommentReactions", ctx, input).Return(output, nil)

        result, err := mockClient.GetCommentReactions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetCommentsForComparedCommit", func(t *testing.T) {
        input := &codecommit.GetCommentsForComparedCommitInput{}
        output := &codecommit.GetCommentsForComparedCommitOutput{}

        mockClient.On("GetCommentsForComparedCommit", ctx, input).Return(output, nil)

        result, err := mockClient.GetCommentsForComparedCommit(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetCommentsForPullRequest", func(t *testing.T) {
        input := &codecommit.GetCommentsForPullRequestInput{}
        output := &codecommit.GetCommentsForPullRequestOutput{}

        mockClient.On("GetCommentsForPullRequest", ctx, input).Return(output, nil)

        result, err := mockClient.GetCommentsForPullRequest(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetCommit", func(t *testing.T) {
        input := &codecommit.GetCommitInput{}
        output := &codecommit.GetCommitOutput{}

        mockClient.On("GetCommit", ctx, input).Return(output, nil)

        result, err := mockClient.GetCommit(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDifferences", func(t *testing.T) {
        input := &codecommit.GetDifferencesInput{}
        output := &codecommit.GetDifferencesOutput{}

        mockClient.On("GetDifferences", ctx, input).Return(output, nil)

        result, err := mockClient.GetDifferences(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetFile", func(t *testing.T) {
        input := &codecommit.GetFileInput{}
        output := &codecommit.GetFileOutput{}

        mockClient.On("GetFile", ctx, input).Return(output, nil)

        result, err := mockClient.GetFile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetFolder", func(t *testing.T) {
        input := &codecommit.GetFolderInput{}
        output := &codecommit.GetFolderOutput{}

        mockClient.On("GetFolder", ctx, input).Return(output, nil)

        result, err := mockClient.GetFolder(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetMergeCommit", func(t *testing.T) {
        input := &codecommit.GetMergeCommitInput{}
        output := &codecommit.GetMergeCommitOutput{}

        mockClient.On("GetMergeCommit", ctx, input).Return(output, nil)

        result, err := mockClient.GetMergeCommit(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetMergeConflicts", func(t *testing.T) {
        input := &codecommit.GetMergeConflictsInput{}
        output := &codecommit.GetMergeConflictsOutput{}

        mockClient.On("GetMergeConflicts", ctx, input).Return(output, nil)

        result, err := mockClient.GetMergeConflicts(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetMergeOptions", func(t *testing.T) {
        input := &codecommit.GetMergeOptionsInput{}
        output := &codecommit.GetMergeOptionsOutput{}

        mockClient.On("GetMergeOptions", ctx, input).Return(output, nil)

        result, err := mockClient.GetMergeOptions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetPullRequest", func(t *testing.T) {
        input := &codecommit.GetPullRequestInput{}
        output := &codecommit.GetPullRequestOutput{}

        mockClient.On("GetPullRequest", ctx, input).Return(output, nil)

        result, err := mockClient.GetPullRequest(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetPullRequestApprovalStates", func(t *testing.T) {
        input := &codecommit.GetPullRequestApprovalStatesInput{}
        output := &codecommit.GetPullRequestApprovalStatesOutput{}

        mockClient.On("GetPullRequestApprovalStates", ctx, input).Return(output, nil)

        result, err := mockClient.GetPullRequestApprovalStates(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetPullRequestOverrideState", func(t *testing.T) {
        input := &codecommit.GetPullRequestOverrideStateInput{}
        output := &codecommit.GetPullRequestOverrideStateOutput{}

        mockClient.On("GetPullRequestOverrideState", ctx, input).Return(output, nil)

        result, err := mockClient.GetPullRequestOverrideState(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetRepository", func(t *testing.T) {
        input := &codecommit.GetRepositoryInput{}
        output := &codecommit.GetRepositoryOutput{}

        mockClient.On("GetRepository", ctx, input).Return(output, nil)

        result, err := mockClient.GetRepository(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetRepositoryTriggers", func(t *testing.T) {
        input := &codecommit.GetRepositoryTriggersInput{}
        output := &codecommit.GetRepositoryTriggersOutput{}

        mockClient.On("GetRepositoryTriggers", ctx, input).Return(output, nil)

        result, err := mockClient.GetRepositoryTriggers(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListApprovalRuleTemplates", func(t *testing.T) {
        input := &codecommit.ListApprovalRuleTemplatesInput{}
        output := &codecommit.ListApprovalRuleTemplatesOutput{}

        mockClient.On("ListApprovalRuleTemplates", ctx, input).Return(output, nil)

        result, err := mockClient.ListApprovalRuleTemplates(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAssociatedApprovalRuleTemplatesForRepository", func(t *testing.T) {
        input := &codecommit.ListAssociatedApprovalRuleTemplatesForRepositoryInput{}
        output := &codecommit.ListAssociatedApprovalRuleTemplatesForRepositoryOutput{}

        mockClient.On("ListAssociatedApprovalRuleTemplatesForRepository", ctx, input).Return(output, nil)

        result, err := mockClient.ListAssociatedApprovalRuleTemplatesForRepository(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListBranches", func(t *testing.T) {
        input := &codecommit.ListBranchesInput{}
        output := &codecommit.ListBranchesOutput{}

        mockClient.On("ListBranches", ctx, input).Return(output, nil)

        result, err := mockClient.ListBranches(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListFileCommitHistory", func(t *testing.T) {
        input := &codecommit.ListFileCommitHistoryInput{}
        output := &codecommit.ListFileCommitHistoryOutput{}

        mockClient.On("ListFileCommitHistory", ctx, input).Return(output, nil)

        result, err := mockClient.ListFileCommitHistory(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListPullRequests", func(t *testing.T) {
        input := &codecommit.ListPullRequestsInput{}
        output := &codecommit.ListPullRequestsOutput{}

        mockClient.On("ListPullRequests", ctx, input).Return(output, nil)

        result, err := mockClient.ListPullRequests(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListRepositories", func(t *testing.T) {
        input := &codecommit.ListRepositoriesInput{}
        output := &codecommit.ListRepositoriesOutput{}

        mockClient.On("ListRepositories", ctx, input).Return(output, nil)

        result, err := mockClient.ListRepositories(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListRepositoriesForApprovalRuleTemplate", func(t *testing.T) {
        input := &codecommit.ListRepositoriesForApprovalRuleTemplateInput{}
        output := &codecommit.ListRepositoriesForApprovalRuleTemplateOutput{}

        mockClient.On("ListRepositoriesForApprovalRuleTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.ListRepositoriesForApprovalRuleTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &codecommit.ListTagsForResourceInput{}
        output := &codecommit.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestMergeBranchesByFastForward", func(t *testing.T) {
        input := &codecommit.MergeBranchesByFastForwardInput{}
        output := &codecommit.MergeBranchesByFastForwardOutput{}

        mockClient.On("MergeBranchesByFastForward", ctx, input).Return(output, nil)

        result, err := mockClient.MergeBranchesByFastForward(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestMergeBranchesBySquash", func(t *testing.T) {
        input := &codecommit.MergeBranchesBySquashInput{}
        output := &codecommit.MergeBranchesBySquashOutput{}

        mockClient.On("MergeBranchesBySquash", ctx, input).Return(output, nil)

        result, err := mockClient.MergeBranchesBySquash(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestMergeBranchesByThreeWay", func(t *testing.T) {
        input := &codecommit.MergeBranchesByThreeWayInput{}
        output := &codecommit.MergeBranchesByThreeWayOutput{}

        mockClient.On("MergeBranchesByThreeWay", ctx, input).Return(output, nil)

        result, err := mockClient.MergeBranchesByThreeWay(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestMergePullRequestByFastForward", func(t *testing.T) {
        input := &codecommit.MergePullRequestByFastForwardInput{}
        output := &codecommit.MergePullRequestByFastForwardOutput{}

        mockClient.On("MergePullRequestByFastForward", ctx, input).Return(output, nil)

        result, err := mockClient.MergePullRequestByFastForward(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestMergePullRequestBySquash", func(t *testing.T) {
        input := &codecommit.MergePullRequestBySquashInput{}
        output := &codecommit.MergePullRequestBySquashOutput{}

        mockClient.On("MergePullRequestBySquash", ctx, input).Return(output, nil)

        result, err := mockClient.MergePullRequestBySquash(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestMergePullRequestByThreeWay", func(t *testing.T) {
        input := &codecommit.MergePullRequestByThreeWayInput{}
        output := &codecommit.MergePullRequestByThreeWayOutput{}

        mockClient.On("MergePullRequestByThreeWay", ctx, input).Return(output, nil)

        result, err := mockClient.MergePullRequestByThreeWay(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestOverridePullRequestApprovalRules", func(t *testing.T) {
        input := &codecommit.OverridePullRequestApprovalRulesInput{}
        output := &codecommit.OverridePullRequestApprovalRulesOutput{}

        mockClient.On("OverridePullRequestApprovalRules", ctx, input).Return(output, nil)

        result, err := mockClient.OverridePullRequestApprovalRules(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPostCommentForComparedCommit", func(t *testing.T) {
        input := &codecommit.PostCommentForComparedCommitInput{}
        output := &codecommit.PostCommentForComparedCommitOutput{}

        mockClient.On("PostCommentForComparedCommit", ctx, input).Return(output, nil)

        result, err := mockClient.PostCommentForComparedCommit(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPostCommentForPullRequest", func(t *testing.T) {
        input := &codecommit.PostCommentForPullRequestInput{}
        output := &codecommit.PostCommentForPullRequestOutput{}

        mockClient.On("PostCommentForPullRequest", ctx, input).Return(output, nil)

        result, err := mockClient.PostCommentForPullRequest(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPostCommentReply", func(t *testing.T) {
        input := &codecommit.PostCommentReplyInput{}
        output := &codecommit.PostCommentReplyOutput{}

        mockClient.On("PostCommentReply", ctx, input).Return(output, nil)

        result, err := mockClient.PostCommentReply(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutCommentReaction", func(t *testing.T) {
        input := &codecommit.PutCommentReactionInput{}
        output := &codecommit.PutCommentReactionOutput{}

        mockClient.On("PutCommentReaction", ctx, input).Return(output, nil)

        result, err := mockClient.PutCommentReaction(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutFile", func(t *testing.T) {
        input := &codecommit.PutFileInput{}
        output := &codecommit.PutFileOutput{}

        mockClient.On("PutFile", ctx, input).Return(output, nil)

        result, err := mockClient.PutFile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutRepositoryTriggers", func(t *testing.T) {
        input := &codecommit.PutRepositoryTriggersInput{}
        output := &codecommit.PutRepositoryTriggersOutput{}

        mockClient.On("PutRepositoryTriggers", ctx, input).Return(output, nil)

        result, err := mockClient.PutRepositoryTriggers(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &codecommit.TagResourceInput{}
        output := &codecommit.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTestRepositoryTriggers", func(t *testing.T) {
        input := &codecommit.TestRepositoryTriggersInput{}
        output := &codecommit.TestRepositoryTriggersOutput{}

        mockClient.On("TestRepositoryTriggers", ctx, input).Return(output, nil)

        result, err := mockClient.TestRepositoryTriggers(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &codecommit.UntagResourceInput{}
        output := &codecommit.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateApprovalRuleTemplateContent", func(t *testing.T) {
        input := &codecommit.UpdateApprovalRuleTemplateContentInput{}
        output := &codecommit.UpdateApprovalRuleTemplateContentOutput{}

        mockClient.On("UpdateApprovalRuleTemplateContent", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateApprovalRuleTemplateContent(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateApprovalRuleTemplateDescription", func(t *testing.T) {
        input := &codecommit.UpdateApprovalRuleTemplateDescriptionInput{}
        output := &codecommit.UpdateApprovalRuleTemplateDescriptionOutput{}

        mockClient.On("UpdateApprovalRuleTemplateDescription", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateApprovalRuleTemplateDescription(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateApprovalRuleTemplateName", func(t *testing.T) {
        input := &codecommit.UpdateApprovalRuleTemplateNameInput{}
        output := &codecommit.UpdateApprovalRuleTemplateNameOutput{}

        mockClient.On("UpdateApprovalRuleTemplateName", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateApprovalRuleTemplateName(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateComment", func(t *testing.T) {
        input := &codecommit.UpdateCommentInput{}
        output := &codecommit.UpdateCommentOutput{}

        mockClient.On("UpdateComment", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateComment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateDefaultBranch", func(t *testing.T) {
        input := &codecommit.UpdateDefaultBranchInput{}
        output := &codecommit.UpdateDefaultBranchOutput{}

        mockClient.On("UpdateDefaultBranch", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateDefaultBranch(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdatePullRequestApprovalRuleContent", func(t *testing.T) {
        input := &codecommit.UpdatePullRequestApprovalRuleContentInput{}
        output := &codecommit.UpdatePullRequestApprovalRuleContentOutput{}

        mockClient.On("UpdatePullRequestApprovalRuleContent", ctx, input).Return(output, nil)

        result, err := mockClient.UpdatePullRequestApprovalRuleContent(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdatePullRequestApprovalState", func(t *testing.T) {
        input := &codecommit.UpdatePullRequestApprovalStateInput{}
        output := &codecommit.UpdatePullRequestApprovalStateOutput{}

        mockClient.On("UpdatePullRequestApprovalState", ctx, input).Return(output, nil)

        result, err := mockClient.UpdatePullRequestApprovalState(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdatePullRequestDescription", func(t *testing.T) {
        input := &codecommit.UpdatePullRequestDescriptionInput{}
        output := &codecommit.UpdatePullRequestDescriptionOutput{}

        mockClient.On("UpdatePullRequestDescription", ctx, input).Return(output, nil)

        result, err := mockClient.UpdatePullRequestDescription(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdatePullRequestStatus", func(t *testing.T) {
        input := &codecommit.UpdatePullRequestStatusInput{}
        output := &codecommit.UpdatePullRequestStatusOutput{}

        mockClient.On("UpdatePullRequestStatus", ctx, input).Return(output, nil)

        result, err := mockClient.UpdatePullRequestStatus(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdatePullRequestTitle", func(t *testing.T) {
        input := &codecommit.UpdatePullRequestTitleInput{}
        output := &codecommit.UpdatePullRequestTitleOutput{}

        mockClient.On("UpdatePullRequestTitle", ctx, input).Return(output, nil)

        result, err := mockClient.UpdatePullRequestTitle(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateRepositoryDescription", func(t *testing.T) {
        input := &codecommit.UpdateRepositoryDescriptionInput{}
        output := &codecommit.UpdateRepositoryDescriptionOutput{}

        mockClient.On("UpdateRepositoryDescription", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateRepositoryDescription(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateRepositoryEncryptionKey", func(t *testing.T) {
        input := &codecommit.UpdateRepositoryEncryptionKeyInput{}
        output := &codecommit.UpdateRepositoryEncryptionKeyOutput{}

        mockClient.On("UpdateRepositoryEncryptionKey", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateRepositoryEncryptionKey(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateRepositoryName", func(t *testing.T) {
        input := &codecommit.UpdateRepositoryNameInput{}
        output := &codecommit.UpdateRepositoryNameOutput{}

        mockClient.On("UpdateRepositoryName", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateRepositoryName(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
