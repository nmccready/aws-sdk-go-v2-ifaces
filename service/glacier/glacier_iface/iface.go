
package glacier_iface

// This file was generated by scripts/gen_ifaces.sh do not edit it manually
import (
    "context"
    . "github.com/aws/aws-sdk-go-v2/service/glacier"
)

// IClient defines the interface for glacier
type IClient interface {
 Options() Options 
 AbortMultipartUpload(ctx context.Context, params *AbortMultipartUploadInput, optFns ...func(*Options)) (*AbortMultipartUploadOutput, error) 
 AbortVaultLock(ctx context.Context, params *AbortVaultLockInput, optFns ...func(*Options)) (*AbortVaultLockOutput, error) 
 AddTagsToVault(ctx context.Context, params *AddTagsToVaultInput, optFns ...func(*Options)) (*AddTagsToVaultOutput, error) 
 CompleteMultipartUpload(ctx context.Context, params *CompleteMultipartUploadInput, optFns ...func(*Options)) (*CompleteMultipartUploadOutput, error) 
 CompleteVaultLock(ctx context.Context, params *CompleteVaultLockInput, optFns ...func(*Options)) (*CompleteVaultLockOutput, error) 
 CreateVault(ctx context.Context, params *CreateVaultInput, optFns ...func(*Options)) (*CreateVaultOutput, error) 
 DeleteArchive(ctx context.Context, params *DeleteArchiveInput, optFns ...func(*Options)) (*DeleteArchiveOutput, error) 
 DeleteVault(ctx context.Context, params *DeleteVaultInput, optFns ...func(*Options)) (*DeleteVaultOutput, error) 
 DeleteVaultAccessPolicy(ctx context.Context, params *DeleteVaultAccessPolicyInput, optFns ...func(*Options)) (*DeleteVaultAccessPolicyOutput, error) 
 DeleteVaultNotifications(ctx context.Context, params *DeleteVaultNotificationsInput, optFns ...func(*Options)) (*DeleteVaultNotificationsOutput, error) 
 DescribeJob(ctx context.Context, params *DescribeJobInput, optFns ...func(*Options)) (*DescribeJobOutput, error) 
 DescribeVault(ctx context.Context, params *DescribeVaultInput, optFns ...func(*Options)) (*DescribeVaultOutput, error) 
 GetDataRetrievalPolicy(ctx context.Context, params *GetDataRetrievalPolicyInput, optFns ...func(*Options)) (*GetDataRetrievalPolicyOutput, error) 
 GetJobOutput(ctx context.Context, params *GetJobOutputInput, optFns ...func(*Options)) (*GetJobOutputOutput, error) 
 GetVaultAccessPolicy(ctx context.Context, params *GetVaultAccessPolicyInput, optFns ...func(*Options)) (*GetVaultAccessPolicyOutput, error) 
 GetVaultLock(ctx context.Context, params *GetVaultLockInput, optFns ...func(*Options)) (*GetVaultLockOutput, error) 
 GetVaultNotifications(ctx context.Context, params *GetVaultNotificationsInput, optFns ...func(*Options)) (*GetVaultNotificationsOutput, error) 
 InitiateJob(ctx context.Context, params *InitiateJobInput, optFns ...func(*Options)) (*InitiateJobOutput, error) 
 InitiateMultipartUpload(ctx context.Context, params *InitiateMultipartUploadInput, optFns ...func(*Options)) (*InitiateMultipartUploadOutput, error) 
 InitiateVaultLock(ctx context.Context, params *InitiateVaultLockInput, optFns ...func(*Options)) (*InitiateVaultLockOutput, error) 
 ListJobs(ctx context.Context, params *ListJobsInput, optFns ...func(*Options)) (*ListJobsOutput, error) 
 ListMultipartUploads(ctx context.Context, params *ListMultipartUploadsInput, optFns ...func(*Options)) (*ListMultipartUploadsOutput, error) 
 ListParts(ctx context.Context, params *ListPartsInput, optFns ...func(*Options)) (*ListPartsOutput, error) 
 ListProvisionedCapacity(ctx context.Context, params *ListProvisionedCapacityInput, optFns ...func(*Options)) (*ListProvisionedCapacityOutput, error) 
 ListTagsForVault(ctx context.Context, params *ListTagsForVaultInput, optFns ...func(*Options)) (*ListTagsForVaultOutput, error) 
 ListVaults(ctx context.Context, params *ListVaultsInput, optFns ...func(*Options)) (*ListVaultsOutput, error) 
 PurchaseProvisionedCapacity(ctx context.Context, params *PurchaseProvisionedCapacityInput, optFns ...func(*Options)) (*PurchaseProvisionedCapacityOutput, error) 
 RemoveTagsFromVault(ctx context.Context, params *RemoveTagsFromVaultInput, optFns ...func(*Options)) (*RemoveTagsFromVaultOutput, error) 
 SetDataRetrievalPolicy(ctx context.Context, params *SetDataRetrievalPolicyInput, optFns ...func(*Options)) (*SetDataRetrievalPolicyOutput, error) 
 SetVaultAccessPolicy(ctx context.Context, params *SetVaultAccessPolicyInput, optFns ...func(*Options)) (*SetVaultAccessPolicyOutput, error) 
 SetVaultNotifications(ctx context.Context, params *SetVaultNotificationsInput, optFns ...func(*Options)) (*SetVaultNotificationsOutput, error) 
 UploadArchive(ctx context.Context, params *UploadArchiveInput, optFns ...func(*Options)) (*UploadArchiveOutput, error) 
 UploadMultipartPart(ctx context.Context, params *UploadMultipartPartInput, optFns ...func(*Options)) (*UploadMultipartPartOutput, error) 
}