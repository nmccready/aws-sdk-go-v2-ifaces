// Code generated by ./scripts/gen_tests.sh DO NOT EDIT.
package securityir_test

// tests for the securityir service interface for 
// this ../../../service/securityir/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/securityir"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/securityir/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/securityir/securityir_iface"
	"github.com/stretchr/testify/assert"
)

func TestSecurityirServiceCanBeMocked(t *testing.T) {
	var iface securityir_iface.IClient
	iface = &securityir.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := securityir.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchGetMemberAccountDetails", func(t *testing.T) {
        input := &securityir.BatchGetMemberAccountDetailsInput{}
        output := &securityir.BatchGetMemberAccountDetailsOutput{}

        mockClient.On("BatchGetMemberAccountDetails", ctx, input).Return(output, nil)

        result, err := mockClient.BatchGetMemberAccountDetails(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCancelMembership", func(t *testing.T) {
        input := &securityir.CancelMembershipInput{}
        output := &securityir.CancelMembershipOutput{}

        mockClient.On("CancelMembership", ctx, input).Return(output, nil)

        result, err := mockClient.CancelMembership(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCloseCase", func(t *testing.T) {
        input := &securityir.CloseCaseInput{}
        output := &securityir.CloseCaseOutput{}

        mockClient.On("CloseCase", ctx, input).Return(output, nil)

        result, err := mockClient.CloseCase(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateCase", func(t *testing.T) {
        input := &securityir.CreateCaseInput{}
        output := &securityir.CreateCaseOutput{}

        mockClient.On("CreateCase", ctx, input).Return(output, nil)

        result, err := mockClient.CreateCase(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateCaseComment", func(t *testing.T) {
        input := &securityir.CreateCaseCommentInput{}
        output := &securityir.CreateCaseCommentOutput{}

        mockClient.On("CreateCaseComment", ctx, input).Return(output, nil)

        result, err := mockClient.CreateCaseComment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateMembership", func(t *testing.T) {
        input := &securityir.CreateMembershipInput{}
        output := &securityir.CreateMembershipOutput{}

        mockClient.On("CreateMembership", ctx, input).Return(output, nil)

        result, err := mockClient.CreateMembership(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetCase", func(t *testing.T) {
        input := &securityir.GetCaseInput{}
        output := &securityir.GetCaseOutput{}

        mockClient.On("GetCase", ctx, input).Return(output, nil)

        result, err := mockClient.GetCase(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetCaseAttachmentDownloadUrl", func(t *testing.T) {
        input := &securityir.GetCaseAttachmentDownloadUrlInput{}
        output := &securityir.GetCaseAttachmentDownloadUrlOutput{}

        mockClient.On("GetCaseAttachmentDownloadUrl", ctx, input).Return(output, nil)

        result, err := mockClient.GetCaseAttachmentDownloadUrl(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetCaseAttachmentUploadUrl", func(t *testing.T) {
        input := &securityir.GetCaseAttachmentUploadUrlInput{}
        output := &securityir.GetCaseAttachmentUploadUrlOutput{}

        mockClient.On("GetCaseAttachmentUploadUrl", ctx, input).Return(output, nil)

        result, err := mockClient.GetCaseAttachmentUploadUrl(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetMembership", func(t *testing.T) {
        input := &securityir.GetMembershipInput{}
        output := &securityir.GetMembershipOutput{}

        mockClient.On("GetMembership", ctx, input).Return(output, nil)

        result, err := mockClient.GetMembership(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListCaseEdits", func(t *testing.T) {
        input := &securityir.ListCaseEditsInput{}
        output := &securityir.ListCaseEditsOutput{}

        mockClient.On("ListCaseEdits", ctx, input).Return(output, nil)

        result, err := mockClient.ListCaseEdits(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListCases", func(t *testing.T) {
        input := &securityir.ListCasesInput{}
        output := &securityir.ListCasesOutput{}

        mockClient.On("ListCases", ctx, input).Return(output, nil)

        result, err := mockClient.ListCases(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListComments", func(t *testing.T) {
        input := &securityir.ListCommentsInput{}
        output := &securityir.ListCommentsOutput{}

        mockClient.On("ListComments", ctx, input).Return(output, nil)

        result, err := mockClient.ListComments(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListMemberships", func(t *testing.T) {
        input := &securityir.ListMembershipsInput{}
        output := &securityir.ListMembershipsOutput{}

        mockClient.On("ListMemberships", ctx, input).Return(output, nil)

        result, err := mockClient.ListMemberships(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &securityir.ListTagsForResourceInput{}
        output := &securityir.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &securityir.TagResourceInput{}
        output := &securityir.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &securityir.UntagResourceInput{}
        output := &securityir.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateCase", func(t *testing.T) {
        input := &securityir.UpdateCaseInput{}
        output := &securityir.UpdateCaseOutput{}

        mockClient.On("UpdateCase", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateCase(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateCaseComment", func(t *testing.T) {
        input := &securityir.UpdateCaseCommentInput{}
        output := &securityir.UpdateCaseCommentOutput{}

        mockClient.On("UpdateCaseComment", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateCaseComment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateCaseStatus", func(t *testing.T) {
        input := &securityir.UpdateCaseStatusInput{}
        output := &securityir.UpdateCaseStatusOutput{}

        mockClient.On("UpdateCaseStatus", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateCaseStatus(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateMembership", func(t *testing.T) {
        input := &securityir.UpdateMembershipInput{}
        output := &securityir.UpdateMembershipOutput{}

        mockClient.On("UpdateMembership", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateMembership(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateResolverType", func(t *testing.T) {
        input := &securityir.UpdateResolverTypeInput{}
        output := &securityir.UpdateResolverTypeOutput{}

        mockClient.On("UpdateResolverType", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateResolverType(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
