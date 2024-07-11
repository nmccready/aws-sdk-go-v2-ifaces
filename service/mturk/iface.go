
package mturk

import (
    "github.com/aws/aws-sdk-go-v2/service/mturk"
)

// IClient defines the interface for mturk
type IClient interface {
 Options() Options 
 AcceptQualificationRequest(ctx context.Context, params *AcceptQualificationRequestInput, optFns ...func(*Options)) (*AcceptQualificationRequestOutput, error) 
 ApproveAssignment(ctx context.Context, params *ApproveAssignmentInput, optFns ...func(*Options)) (*ApproveAssignmentOutput, error) 
 AssociateQualificationWithWorker(ctx context.Context, params *AssociateQualificationWithWorkerInput, optFns ...func(*Options)) (*AssociateQualificationWithWorkerOutput, error) 
 CreateAdditionalAssignmentsForHIT(ctx context.Context, params *CreateAdditionalAssignmentsForHITInput, optFns ...func(*Options)) (*CreateAdditionalAssignmentsForHITOutput, error) 
 CreateHIT(ctx context.Context, params *CreateHITInput, optFns ...func(*Options)) (*CreateHITOutput, error) 
 CreateHITType(ctx context.Context, params *CreateHITTypeInput, optFns ...func(*Options)) (*CreateHITTypeOutput, error) 
 CreateHITWithHITType(ctx context.Context, params *CreateHITWithHITTypeInput, optFns ...func(*Options)) (*CreateHITWithHITTypeOutput, error) 
 CreateQualificationType(ctx context.Context, params *CreateQualificationTypeInput, optFns ...func(*Options)) (*CreateQualificationTypeOutput, error) 
 CreateWorkerBlock(ctx context.Context, params *CreateWorkerBlockInput, optFns ...func(*Options)) (*CreateWorkerBlockOutput, error) 
 DeleteHIT(ctx context.Context, params *DeleteHITInput, optFns ...func(*Options)) (*DeleteHITOutput, error) 
 DeleteQualificationType(ctx context.Context, params *DeleteQualificationTypeInput, optFns ...func(*Options)) (*DeleteQualificationTypeOutput, error) 
 DeleteWorkerBlock(ctx context.Context, params *DeleteWorkerBlockInput, optFns ...func(*Options)) (*DeleteWorkerBlockOutput, error) 
 DisassociateQualificationFromWorker(ctx context.Context, params *DisassociateQualificationFromWorkerInput, optFns ...func(*Options)) (*DisassociateQualificationFromWorkerOutput, error) 
 GetAccountBalance(ctx context.Context, params *GetAccountBalanceInput, optFns ...func(*Options)) (*GetAccountBalanceOutput, error) 
 GetAssignment(ctx context.Context, params *GetAssignmentInput, optFns ...func(*Options)) (*GetAssignmentOutput, error) 
 GetFileUploadURL(ctx context.Context, params *GetFileUploadURLInput, optFns ...func(*Options)) (*GetFileUploadURLOutput, error) 
 GetHIT(ctx context.Context, params *GetHITInput, optFns ...func(*Options)) (*GetHITOutput, error) 
 GetQualificationScore(ctx context.Context, params *GetQualificationScoreInput, optFns ...func(*Options)) (*GetQualificationScoreOutput, error) 
 GetQualificationType(ctx context.Context, params *GetQualificationTypeInput, optFns ...func(*Options)) (*GetQualificationTypeOutput, error) 
 ListAssignmentsForHIT(ctx context.Context, params *ListAssignmentsForHITInput, optFns ...func(*Options)) (*ListAssignmentsForHITOutput, error) 
 ListBonusPayments(ctx context.Context, params *ListBonusPaymentsInput, optFns ...func(*Options)) (*ListBonusPaymentsOutput, error) 
 ListHITs(ctx context.Context, params *ListHITsInput, optFns ...func(*Options)) (*ListHITsOutput, error) 
 ListHITsForQualificationType(ctx context.Context, params *ListHITsForQualificationTypeInput, optFns ...func(*Options)) (*ListHITsForQualificationTypeOutput, error) 
 ListQualificationRequests(ctx context.Context, params *ListQualificationRequestsInput, optFns ...func(*Options)) (*ListQualificationRequestsOutput, error) 
 ListQualificationTypes(ctx context.Context, params *ListQualificationTypesInput, optFns ...func(*Options)) (*ListQualificationTypesOutput, error) 
 ListReviewPolicyResultsForHIT(ctx context.Context, params *ListReviewPolicyResultsForHITInput, optFns ...func(*Options)) (*ListReviewPolicyResultsForHITOutput, error) 
 ListReviewableHITs(ctx context.Context, params *ListReviewableHITsInput, optFns ...func(*Options)) (*ListReviewableHITsOutput, error) 
 ListWorkerBlocks(ctx context.Context, params *ListWorkerBlocksInput, optFns ...func(*Options)) (*ListWorkerBlocksOutput, error) 
 ListWorkersWithQualificationType(ctx context.Context, params *ListWorkersWithQualificationTypeInput, optFns ...func(*Options)) (*ListWorkersWithQualificationTypeOutput, error) 
 NotifyWorkers(ctx context.Context, params *NotifyWorkersInput, optFns ...func(*Options)) (*NotifyWorkersOutput, error) 
 RejectAssignment(ctx context.Context, params *RejectAssignmentInput, optFns ...func(*Options)) (*RejectAssignmentOutput, error) 
 RejectQualificationRequest(ctx context.Context, params *RejectQualificationRequestInput, optFns ...func(*Options)) (*RejectQualificationRequestOutput, error) 
 SendBonus(ctx context.Context, params *SendBonusInput, optFns ...func(*Options)) (*SendBonusOutput, error) 
 SendTestEventNotification(ctx context.Context, params *SendTestEventNotificationInput, optFns ...func(*Options)) (*SendTestEventNotificationOutput, error) 
 UpdateExpirationForHIT(ctx context.Context, params *UpdateExpirationForHITInput, optFns ...func(*Options)) (*UpdateExpirationForHITOutput, error) 
 UpdateHITReviewStatus(ctx context.Context, params *UpdateHITReviewStatusInput, optFns ...func(*Options)) (*UpdateHITReviewStatusOutput, error) 
 UpdateHITTypeOfHIT(ctx context.Context, params *UpdateHITTypeOfHITInput, optFns ...func(*Options)) (*UpdateHITTypeOfHITOutput, error) 
 UpdateNotificationSettings(ctx context.Context, params *UpdateNotificationSettingsInput, optFns ...func(*Options)) (*UpdateNotificationSettingsOutput, error) 
 UpdateQualificationType(ctx context.Context, params *UpdateQualificationTypeInput, optFns ...func(*Options)) (*UpdateQualificationTypeOutput, error) 
}
