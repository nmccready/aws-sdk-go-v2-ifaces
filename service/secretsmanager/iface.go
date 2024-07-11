
package secretsmanager

import (
    "github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

// ISecretsmanager defines the interface for secretsmanager
type ISecretsmanager interface {
 Options() Options 
 BatchGetSecretValue(ctx context.Context, params *BatchGetSecretValueInput, optFns ...func(*Options)) (*BatchGetSecretValueOutput, error) 
 CancelRotateSecret(ctx context.Context, params *CancelRotateSecretInput, optFns ...func(*Options)) (*CancelRotateSecretOutput, error) 
 CreateSecret(ctx context.Context, params *CreateSecretInput, optFns ...func(*Options)) (*CreateSecretOutput, error) 
 DeleteResourcePolicy(ctx context.Context, params *DeleteResourcePolicyInput, optFns ...func(*Options)) (*DeleteResourcePolicyOutput, error) 
 DeleteSecret(ctx context.Context, params *DeleteSecretInput, optFns ...func(*Options)) (*DeleteSecretOutput, error) 
 DescribeSecret(ctx context.Context, params *DescribeSecretInput, optFns ...func(*Options)) (*DescribeSecretOutput, error) 
 GetRandomPassword(ctx context.Context, params *GetRandomPasswordInput, optFns ...func(*Options)) (*GetRandomPasswordOutput, error) 
 GetResourcePolicy(ctx context.Context, params *GetResourcePolicyInput, optFns ...func(*Options)) (*GetResourcePolicyOutput, error) 
 GetSecretValue(ctx context.Context, params *GetSecretValueInput, optFns ...func(*Options)) (*GetSecretValueOutput, error) 
 ListSecretVersionIds(ctx context.Context, params *ListSecretVersionIdsInput, optFns ...func(*Options)) (*ListSecretVersionIdsOutput, error) 
 ListSecrets(ctx context.Context, params *ListSecretsInput, optFns ...func(*Options)) (*ListSecretsOutput, error) 
 PutResourcePolicy(ctx context.Context, params *PutResourcePolicyInput, optFns ...func(*Options)) (*PutResourcePolicyOutput, error) 
 PutSecretValue(ctx context.Context, params *PutSecretValueInput, optFns ...func(*Options)) (*PutSecretValueOutput, error) 
 RemoveRegionsFromReplication(ctx context.Context, params *RemoveRegionsFromReplicationInput, optFns ...func(*Options)) (*RemoveRegionsFromReplicationOutput, error) 
 ReplicateSecretToRegions(ctx context.Context, params *ReplicateSecretToRegionsInput, optFns ...func(*Options)) (*ReplicateSecretToRegionsOutput, error) 
 RestoreSecret(ctx context.Context, params *RestoreSecretInput, optFns ...func(*Options)) (*RestoreSecretOutput, error) 
 RotateSecret(ctx context.Context, params *RotateSecretInput, optFns ...func(*Options)) (*RotateSecretOutput, error) 
 StopReplicationToReplica(ctx context.Context, params *StopReplicationToReplicaInput, optFns ...func(*Options)) (*StopReplicationToReplicaOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateSecret(ctx context.Context, params *UpdateSecretInput, optFns ...func(*Options)) (*UpdateSecretOutput, error) 
 UpdateSecretVersionStage(ctx context.Context, params *UpdateSecretVersionStageInput, optFns ...func(*Options)) (*UpdateSecretVersionStageOutput, error) 
 ValidateResourcePolicy(ctx context.Context, params *ValidateResourcePolicyInput, optFns ...func(*Options)) (*ValidateResourcePolicyOutput, error) 
}
