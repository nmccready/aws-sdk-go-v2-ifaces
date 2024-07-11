
package signer

import (
    "github.com/aws/aws-sdk-go-v2/service/signer"
)

// IClient defines the interface for signer
type IClient interface {
 Options() Options 
 AddProfilePermission(ctx context.Context, params *AddProfilePermissionInput, optFns ...func(*Options)) (*AddProfilePermissionOutput, error) 
 CancelSigningProfile(ctx context.Context, params *CancelSigningProfileInput, optFns ...func(*Options)) (*CancelSigningProfileOutput, error) 
 DescribeSigningJob(ctx context.Context, params *DescribeSigningJobInput, optFns ...func(*Options)) (*DescribeSigningJobOutput, error) 
 GetRevocationStatus(ctx context.Context, params *GetRevocationStatusInput, optFns ...func(*Options)) (*GetRevocationStatusOutput, error) 
 GetSigningPlatform(ctx context.Context, params *GetSigningPlatformInput, optFns ...func(*Options)) (*GetSigningPlatformOutput, error) 
 GetSigningProfile(ctx context.Context, params *GetSigningProfileInput, optFns ...func(*Options)) (*GetSigningProfileOutput, error) 
 ListProfilePermissions(ctx context.Context, params *ListProfilePermissionsInput, optFns ...func(*Options)) (*ListProfilePermissionsOutput, error) 
 ListSigningJobs(ctx context.Context, params *ListSigningJobsInput, optFns ...func(*Options)) (*ListSigningJobsOutput, error) 
 ListSigningPlatforms(ctx context.Context, params *ListSigningPlatformsInput, optFns ...func(*Options)) (*ListSigningPlatformsOutput, error) 
 ListSigningProfiles(ctx context.Context, params *ListSigningProfilesInput, optFns ...func(*Options)) (*ListSigningProfilesOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 PutSigningProfile(ctx context.Context, params *PutSigningProfileInput, optFns ...func(*Options)) (*PutSigningProfileOutput, error) 
 RemoveProfilePermission(ctx context.Context, params *RemoveProfilePermissionInput, optFns ...func(*Options)) (*RemoveProfilePermissionOutput, error) 
 RevokeSignature(ctx context.Context, params *RevokeSignatureInput, optFns ...func(*Options)) (*RevokeSignatureOutput, error) 
 RevokeSigningProfile(ctx context.Context, params *RevokeSigningProfileInput, optFns ...func(*Options)) (*RevokeSigningProfileOutput, error) 
 SignPayload(ctx context.Context, params *SignPayloadInput, optFns ...func(*Options)) (*SignPayloadOutput, error) 
 StartSigningJob(ctx context.Context, params *StartSigningJobInput, optFns ...func(*Options)) (*StartSigningJobOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
}
