package mturk_test

// tests for the mturk service interface for this ../../../service/mturk/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/mturk"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/mturk/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/mturk/mturk_iface"
	"github.com/stretchr/testify/assert"
)

func TestMturkServiceCanBeMocked(t *testing.T) {
	var iface mturk_iface.IClient
	iface = &mturk.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := mturk.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAcceptQualificationRequest", func(t *testing.T) {
        input := &mturk.AcceptQualificationRequestInput{}
        output := &mturk.AcceptQualificationRequestOutput{}

        mockClient.On("AcceptQualificationRequest", ctx, input).Return(output, nil)

        result, err := mockClient.AcceptQualificationRequest(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestApproveAssignment", func(t *testing.T) {
        input := &mturk.ApproveAssignmentInput{}
        output := &mturk.ApproveAssignmentOutput{}

        mockClient.On("ApproveAssignment", ctx, input).Return(output, nil)

        result, err := mockClient.ApproveAssignment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateQualificationWithWorker", func(t *testing.T) {
        input := &mturk.AssociateQualificationWithWorkerInput{}
        output := &mturk.AssociateQualificationWithWorkerOutput{}

        mockClient.On("AssociateQualificationWithWorker", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateQualificationWithWorker(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateAdditionalAssignmentsForHIT", func(t *testing.T) {
        input := &mturk.CreateAdditionalAssignmentsForHITInput{}
        output := &mturk.CreateAdditionalAssignmentsForHITOutput{}

        mockClient.On("CreateAdditionalAssignmentsForHIT", ctx, input).Return(output, nil)

        result, err := mockClient.CreateAdditionalAssignmentsForHIT(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateHIT", func(t *testing.T) {
        input := &mturk.CreateHITInput{}
        output := &mturk.CreateHITOutput{}

        mockClient.On("CreateHIT", ctx, input).Return(output, nil)

        result, err := mockClient.CreateHIT(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateHITType", func(t *testing.T) {
        input := &mturk.CreateHITTypeInput{}
        output := &mturk.CreateHITTypeOutput{}

        mockClient.On("CreateHITType", ctx, input).Return(output, nil)

        result, err := mockClient.CreateHITType(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateHITWithHITType", func(t *testing.T) {
        input := &mturk.CreateHITWithHITTypeInput{}
        output := &mturk.CreateHITWithHITTypeOutput{}

        mockClient.On("CreateHITWithHITType", ctx, input).Return(output, nil)

        result, err := mockClient.CreateHITWithHITType(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateQualificationType", func(t *testing.T) {
        input := &mturk.CreateQualificationTypeInput{}
        output := &mturk.CreateQualificationTypeOutput{}

        mockClient.On("CreateQualificationType", ctx, input).Return(output, nil)

        result, err := mockClient.CreateQualificationType(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateWorkerBlock", func(t *testing.T) {
        input := &mturk.CreateWorkerBlockInput{}
        output := &mturk.CreateWorkerBlockOutput{}

        mockClient.On("CreateWorkerBlock", ctx, input).Return(output, nil)

        result, err := mockClient.CreateWorkerBlock(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteHIT", func(t *testing.T) {
        input := &mturk.DeleteHITInput{}
        output := &mturk.DeleteHITOutput{}

        mockClient.On("DeleteHIT", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteHIT(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteQualificationType", func(t *testing.T) {
        input := &mturk.DeleteQualificationTypeInput{}
        output := &mturk.DeleteQualificationTypeOutput{}

        mockClient.On("DeleteQualificationType", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteQualificationType(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteWorkerBlock", func(t *testing.T) {
        input := &mturk.DeleteWorkerBlockInput{}
        output := &mturk.DeleteWorkerBlockOutput{}

        mockClient.On("DeleteWorkerBlock", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteWorkerBlock(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateQualificationFromWorker", func(t *testing.T) {
        input := &mturk.DisassociateQualificationFromWorkerInput{}
        output := &mturk.DisassociateQualificationFromWorkerOutput{}

        mockClient.On("DisassociateQualificationFromWorker", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateQualificationFromWorker(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetAccountBalance", func(t *testing.T) {
        input := &mturk.GetAccountBalanceInput{}
        output := &mturk.GetAccountBalanceOutput{}

        mockClient.On("GetAccountBalance", ctx, input).Return(output, nil)

        result, err := mockClient.GetAccountBalance(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetAssignment", func(t *testing.T) {
        input := &mturk.GetAssignmentInput{}
        output := &mturk.GetAssignmentOutput{}

        mockClient.On("GetAssignment", ctx, input).Return(output, nil)

        result, err := mockClient.GetAssignment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetFileUploadURL", func(t *testing.T) {
        input := &mturk.GetFileUploadURLInput{}
        output := &mturk.GetFileUploadURLOutput{}

        mockClient.On("GetFileUploadURL", ctx, input).Return(output, nil)

        result, err := mockClient.GetFileUploadURL(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetHIT", func(t *testing.T) {
        input := &mturk.GetHITInput{}
        output := &mturk.GetHITOutput{}

        mockClient.On("GetHIT", ctx, input).Return(output, nil)

        result, err := mockClient.GetHIT(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetQualificationScore", func(t *testing.T) {
        input := &mturk.GetQualificationScoreInput{}
        output := &mturk.GetQualificationScoreOutput{}

        mockClient.On("GetQualificationScore", ctx, input).Return(output, nil)

        result, err := mockClient.GetQualificationScore(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetQualificationType", func(t *testing.T) {
        input := &mturk.GetQualificationTypeInput{}
        output := &mturk.GetQualificationTypeOutput{}

        mockClient.On("GetQualificationType", ctx, input).Return(output, nil)

        result, err := mockClient.GetQualificationType(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAssignmentsForHIT", func(t *testing.T) {
        input := &mturk.ListAssignmentsForHITInput{}
        output := &mturk.ListAssignmentsForHITOutput{}

        mockClient.On("ListAssignmentsForHIT", ctx, input).Return(output, nil)

        result, err := mockClient.ListAssignmentsForHIT(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListBonusPayments", func(t *testing.T) {
        input := &mturk.ListBonusPaymentsInput{}
        output := &mturk.ListBonusPaymentsOutput{}

        mockClient.On("ListBonusPayments", ctx, input).Return(output, nil)

        result, err := mockClient.ListBonusPayments(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListHITs", func(t *testing.T) {
        input := &mturk.ListHITsInput{}
        output := &mturk.ListHITsOutput{}

        mockClient.On("ListHITs", ctx, input).Return(output, nil)

        result, err := mockClient.ListHITs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListHITsForQualificationType", func(t *testing.T) {
        input := &mturk.ListHITsForQualificationTypeInput{}
        output := &mturk.ListHITsForQualificationTypeOutput{}

        mockClient.On("ListHITsForQualificationType", ctx, input).Return(output, nil)

        result, err := mockClient.ListHITsForQualificationType(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListQualificationRequests", func(t *testing.T) {
        input := &mturk.ListQualificationRequestsInput{}
        output := &mturk.ListQualificationRequestsOutput{}

        mockClient.On("ListQualificationRequests", ctx, input).Return(output, nil)

        result, err := mockClient.ListQualificationRequests(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListQualificationTypes", func(t *testing.T) {
        input := &mturk.ListQualificationTypesInput{}
        output := &mturk.ListQualificationTypesOutput{}

        mockClient.On("ListQualificationTypes", ctx, input).Return(output, nil)

        result, err := mockClient.ListQualificationTypes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListReviewPolicyResultsForHIT", func(t *testing.T) {
        input := &mturk.ListReviewPolicyResultsForHITInput{}
        output := &mturk.ListReviewPolicyResultsForHITOutput{}

        mockClient.On("ListReviewPolicyResultsForHIT", ctx, input).Return(output, nil)

        result, err := mockClient.ListReviewPolicyResultsForHIT(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListReviewableHITs", func(t *testing.T) {
        input := &mturk.ListReviewableHITsInput{}
        output := &mturk.ListReviewableHITsOutput{}

        mockClient.On("ListReviewableHITs", ctx, input).Return(output, nil)

        result, err := mockClient.ListReviewableHITs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListWorkerBlocks", func(t *testing.T) {
        input := &mturk.ListWorkerBlocksInput{}
        output := &mturk.ListWorkerBlocksOutput{}

        mockClient.On("ListWorkerBlocks", ctx, input).Return(output, nil)

        result, err := mockClient.ListWorkerBlocks(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListWorkersWithQualificationType", func(t *testing.T) {
        input := &mturk.ListWorkersWithQualificationTypeInput{}
        output := &mturk.ListWorkersWithQualificationTypeOutput{}

        mockClient.On("ListWorkersWithQualificationType", ctx, input).Return(output, nil)

        result, err := mockClient.ListWorkersWithQualificationType(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestNotifyWorkers", func(t *testing.T) {
        input := &mturk.NotifyWorkersInput{}
        output := &mturk.NotifyWorkersOutput{}

        mockClient.On("NotifyWorkers", ctx, input).Return(output, nil)

        result, err := mockClient.NotifyWorkers(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRejectAssignment", func(t *testing.T) {
        input := &mturk.RejectAssignmentInput{}
        output := &mturk.RejectAssignmentOutput{}

        mockClient.On("RejectAssignment", ctx, input).Return(output, nil)

        result, err := mockClient.RejectAssignment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRejectQualificationRequest", func(t *testing.T) {
        input := &mturk.RejectQualificationRequestInput{}
        output := &mturk.RejectQualificationRequestOutput{}

        mockClient.On("RejectQualificationRequest", ctx, input).Return(output, nil)

        result, err := mockClient.RejectQualificationRequest(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSendBonus", func(t *testing.T) {
        input := &mturk.SendBonusInput{}
        output := &mturk.SendBonusOutput{}

        mockClient.On("SendBonus", ctx, input).Return(output, nil)

        result, err := mockClient.SendBonus(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSendTestEventNotification", func(t *testing.T) {
        input := &mturk.SendTestEventNotificationInput{}
        output := &mturk.SendTestEventNotificationOutput{}

        mockClient.On("SendTestEventNotification", ctx, input).Return(output, nil)

        result, err := mockClient.SendTestEventNotification(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateExpirationForHIT", func(t *testing.T) {
        input := &mturk.UpdateExpirationForHITInput{}
        output := &mturk.UpdateExpirationForHITOutput{}

        mockClient.On("UpdateExpirationForHIT", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateExpirationForHIT(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateHITReviewStatus", func(t *testing.T) {
        input := &mturk.UpdateHITReviewStatusInput{}
        output := &mturk.UpdateHITReviewStatusOutput{}

        mockClient.On("UpdateHITReviewStatus", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateHITReviewStatus(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateHITTypeOfHIT", func(t *testing.T) {
        input := &mturk.UpdateHITTypeOfHITInput{}
        output := &mturk.UpdateHITTypeOfHITOutput{}

        mockClient.On("UpdateHITTypeOfHIT", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateHITTypeOfHIT(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateNotificationSettings", func(t *testing.T) {
        input := &mturk.UpdateNotificationSettingsInput{}
        output := &mturk.UpdateNotificationSettingsOutput{}

        mockClient.On("UpdateNotificationSettings", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateNotificationSettings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateQualificationType", func(t *testing.T) {
        input := &mturk.UpdateQualificationTypeInput{}
        output := &mturk.UpdateQualificationTypeOutput{}

        mockClient.On("UpdateQualificationType", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateQualificationType(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
