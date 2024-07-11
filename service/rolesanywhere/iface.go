
package rolesanywhere

import (
    "github.com/aws/aws-sdk-go-v2/service/rolesanywhere"
)

// IRolesanywhere defines the interface for rolesanywhere
type IRolesanywhere interface {
 Options() Options 
 CreateProfile(ctx context.Context, params *CreateProfileInput, optFns ...func(*Options)) (*CreateProfileOutput, error) 
 CreateTrustAnchor(ctx context.Context, params *CreateTrustAnchorInput, optFns ...func(*Options)) (*CreateTrustAnchorOutput, error) 
 DeleteAttributeMapping(ctx context.Context, params *DeleteAttributeMappingInput, optFns ...func(*Options)) (*DeleteAttributeMappingOutput, error) 
 DeleteCrl(ctx context.Context, params *DeleteCrlInput, optFns ...func(*Options)) (*DeleteCrlOutput, error) 
 DeleteProfile(ctx context.Context, params *DeleteProfileInput, optFns ...func(*Options)) (*DeleteProfileOutput, error) 
 DeleteTrustAnchor(ctx context.Context, params *DeleteTrustAnchorInput, optFns ...func(*Options)) (*DeleteTrustAnchorOutput, error) 
 DisableCrl(ctx context.Context, params *DisableCrlInput, optFns ...func(*Options)) (*DisableCrlOutput, error) 
 DisableProfile(ctx context.Context, params *DisableProfileInput, optFns ...func(*Options)) (*DisableProfileOutput, error) 
 DisableTrustAnchor(ctx context.Context, params *DisableTrustAnchorInput, optFns ...func(*Options)) (*DisableTrustAnchorOutput, error) 
 EnableCrl(ctx context.Context, params *EnableCrlInput, optFns ...func(*Options)) (*EnableCrlOutput, error) 
 EnableProfile(ctx context.Context, params *EnableProfileInput, optFns ...func(*Options)) (*EnableProfileOutput, error) 
 EnableTrustAnchor(ctx context.Context, params *EnableTrustAnchorInput, optFns ...func(*Options)) (*EnableTrustAnchorOutput, error) 
 GetCrl(ctx context.Context, params *GetCrlInput, optFns ...func(*Options)) (*GetCrlOutput, error) 
 GetProfile(ctx context.Context, params *GetProfileInput, optFns ...func(*Options)) (*GetProfileOutput, error) 
 GetSubject(ctx context.Context, params *GetSubjectInput, optFns ...func(*Options)) (*GetSubjectOutput, error) 
 GetTrustAnchor(ctx context.Context, params *GetTrustAnchorInput, optFns ...func(*Options)) (*GetTrustAnchorOutput, error) 
 ImportCrl(ctx context.Context, params *ImportCrlInput, optFns ...func(*Options)) (*ImportCrlOutput, error) 
 ListCrls(ctx context.Context, params *ListCrlsInput, optFns ...func(*Options)) (*ListCrlsOutput, error) 
 ListProfiles(ctx context.Context, params *ListProfilesInput, optFns ...func(*Options)) (*ListProfilesOutput, error) 
 ListSubjects(ctx context.Context, params *ListSubjectsInput, optFns ...func(*Options)) (*ListSubjectsOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 ListTrustAnchors(ctx context.Context, params *ListTrustAnchorsInput, optFns ...func(*Options)) (*ListTrustAnchorsOutput, error) 
 PutAttributeMapping(ctx context.Context, params *PutAttributeMappingInput, optFns ...func(*Options)) (*PutAttributeMappingOutput, error) 
 PutNotificationSettings(ctx context.Context, params *PutNotificationSettingsInput, optFns ...func(*Options)) (*PutNotificationSettingsOutput, error) 
 ResetNotificationSettings(ctx context.Context, params *ResetNotificationSettingsInput, optFns ...func(*Options)) (*ResetNotificationSettingsOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateCrl(ctx context.Context, params *UpdateCrlInput, optFns ...func(*Options)) (*UpdateCrlOutput, error) 
 UpdateProfile(ctx context.Context, params *UpdateProfileInput, optFns ...func(*Options)) (*UpdateProfileOutput, error) 
 UpdateTrustAnchor(ctx context.Context, params *UpdateTrustAnchorInput, optFns ...func(*Options)) (*UpdateTrustAnchorOutput, error) 
}
