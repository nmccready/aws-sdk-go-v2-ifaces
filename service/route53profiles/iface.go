
package route53profiles

import (
    "github.com/aws/aws-sdk-go-v2/service/route53profiles"
)

// IRoute53profiles defines the interface for route53profiles
type IRoute53profiles interface {
 Options() Options 
 AssociateProfile(ctx context.Context, params *AssociateProfileInput, optFns ...func(*Options)) (*AssociateProfileOutput, error) 
 AssociateResourceToProfile(ctx context.Context, params *AssociateResourceToProfileInput, optFns ...func(*Options)) (*AssociateResourceToProfileOutput, error) 
 CreateProfile(ctx context.Context, params *CreateProfileInput, optFns ...func(*Options)) (*CreateProfileOutput, error) 
 DeleteProfile(ctx context.Context, params *DeleteProfileInput, optFns ...func(*Options)) (*DeleteProfileOutput, error) 
 DisassociateProfile(ctx context.Context, params *DisassociateProfileInput, optFns ...func(*Options)) (*DisassociateProfileOutput, error) 
 DisassociateResourceFromProfile(ctx context.Context, params *DisassociateResourceFromProfileInput, optFns ...func(*Options)) (*DisassociateResourceFromProfileOutput, error) 
 GetProfile(ctx context.Context, params *GetProfileInput, optFns ...func(*Options)) (*GetProfileOutput, error) 
 GetProfileAssociation(ctx context.Context, params *GetProfileAssociationInput, optFns ...func(*Options)) (*GetProfileAssociationOutput, error) 
 GetProfileResourceAssociation(ctx context.Context, params *GetProfileResourceAssociationInput, optFns ...func(*Options)) (*GetProfileResourceAssociationOutput, error) 
 ListProfileAssociations(ctx context.Context, params *ListProfileAssociationsInput, optFns ...func(*Options)) (*ListProfileAssociationsOutput, error) 
 ListProfileResourceAssociations(ctx context.Context, params *ListProfileResourceAssociationsInput, optFns ...func(*Options)) (*ListProfileResourceAssociationsOutput, error) 
 ListProfiles(ctx context.Context, params *ListProfilesInput, optFns ...func(*Options)) (*ListProfilesOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateProfileResourceAssociation(ctx context.Context, params *UpdateProfileResourceAssociationInput, optFns ...func(*Options)) (*UpdateProfileResourceAssociationOutput, error) 
}
