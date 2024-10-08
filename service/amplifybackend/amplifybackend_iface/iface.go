
package amplifybackend_iface

// This file was generated by scripts/gen_ifaces.sh do not edit it manually
import (
    "context"
    . "github.com/aws/aws-sdk-go-v2/service/amplifybackend"
)

// IClient defines the interface for amplifybackend
type IClient interface {
 Options() Options 
 CloneBackend(ctx context.Context, params *CloneBackendInput, optFns ...func(*Options)) (*CloneBackendOutput, error) 
 CreateBackend(ctx context.Context, params *CreateBackendInput, optFns ...func(*Options)) (*CreateBackendOutput, error) 
 CreateBackendAPI(ctx context.Context, params *CreateBackendAPIInput, optFns ...func(*Options)) (*CreateBackendAPIOutput, error) 
 CreateBackendAuth(ctx context.Context, params *CreateBackendAuthInput, optFns ...func(*Options)) (*CreateBackendAuthOutput, error) 
 CreateBackendConfig(ctx context.Context, params *CreateBackendConfigInput, optFns ...func(*Options)) (*CreateBackendConfigOutput, error) 
 CreateBackendStorage(ctx context.Context, params *CreateBackendStorageInput, optFns ...func(*Options)) (*CreateBackendStorageOutput, error) 
 CreateToken(ctx context.Context, params *CreateTokenInput, optFns ...func(*Options)) (*CreateTokenOutput, error) 
 DeleteBackend(ctx context.Context, params *DeleteBackendInput, optFns ...func(*Options)) (*DeleteBackendOutput, error) 
 DeleteBackendAPI(ctx context.Context, params *DeleteBackendAPIInput, optFns ...func(*Options)) (*DeleteBackendAPIOutput, error) 
 DeleteBackendAuth(ctx context.Context, params *DeleteBackendAuthInput, optFns ...func(*Options)) (*DeleteBackendAuthOutput, error) 
 DeleteBackendStorage(ctx context.Context, params *DeleteBackendStorageInput, optFns ...func(*Options)) (*DeleteBackendStorageOutput, error) 
 DeleteToken(ctx context.Context, params *DeleteTokenInput, optFns ...func(*Options)) (*DeleteTokenOutput, error) 
 GenerateBackendAPIModels(ctx context.Context, params *GenerateBackendAPIModelsInput, optFns ...func(*Options)) (*GenerateBackendAPIModelsOutput, error) 
 GetBackend(ctx context.Context, params *GetBackendInput, optFns ...func(*Options)) (*GetBackendOutput, error) 
 GetBackendAPI(ctx context.Context, params *GetBackendAPIInput, optFns ...func(*Options)) (*GetBackendAPIOutput, error) 
 GetBackendAPIModels(ctx context.Context, params *GetBackendAPIModelsInput, optFns ...func(*Options)) (*GetBackendAPIModelsOutput, error) 
 GetBackendAuth(ctx context.Context, params *GetBackendAuthInput, optFns ...func(*Options)) (*GetBackendAuthOutput, error) 
 GetBackendJob(ctx context.Context, params *GetBackendJobInput, optFns ...func(*Options)) (*GetBackendJobOutput, error) 
 GetBackendStorage(ctx context.Context, params *GetBackendStorageInput, optFns ...func(*Options)) (*GetBackendStorageOutput, error) 
 GetToken(ctx context.Context, params *GetTokenInput, optFns ...func(*Options)) (*GetTokenOutput, error) 
 ImportBackendAuth(ctx context.Context, params *ImportBackendAuthInput, optFns ...func(*Options)) (*ImportBackendAuthOutput, error) 
 ImportBackendStorage(ctx context.Context, params *ImportBackendStorageInput, optFns ...func(*Options)) (*ImportBackendStorageOutput, error) 
 ListBackendJobs(ctx context.Context, params *ListBackendJobsInput, optFns ...func(*Options)) (*ListBackendJobsOutput, error) 
 ListS3Buckets(ctx context.Context, params *ListS3BucketsInput, optFns ...func(*Options)) (*ListS3BucketsOutput, error) 
 RemoveAllBackends(ctx context.Context, params *RemoveAllBackendsInput, optFns ...func(*Options)) (*RemoveAllBackendsOutput, error) 
 RemoveBackendConfig(ctx context.Context, params *RemoveBackendConfigInput, optFns ...func(*Options)) (*RemoveBackendConfigOutput, error) 
 UpdateBackendAPI(ctx context.Context, params *UpdateBackendAPIInput, optFns ...func(*Options)) (*UpdateBackendAPIOutput, error) 
 UpdateBackendAuth(ctx context.Context, params *UpdateBackendAuthInput, optFns ...func(*Options)) (*UpdateBackendAuthOutput, error) 
 UpdateBackendConfig(ctx context.Context, params *UpdateBackendConfigInput, optFns ...func(*Options)) (*UpdateBackendConfigOutput, error) 
 UpdateBackendJob(ctx context.Context, params *UpdateBackendJobInput, optFns ...func(*Options)) (*UpdateBackendJobOutput, error) 
 UpdateBackendStorage(ctx context.Context, params *UpdateBackendStorageInput, optFns ...func(*Options)) (*UpdateBackendStorageOutput, error) 
}
