
package repostspace

import (
    "github.com/aws/aws-sdk-go-v2/service/repostspace"
)

// IRepostspace defines the interface for repostspace
type IRepostspace interface {
 Options() Options 
 CreateSpace(ctx context.Context, params *CreateSpaceInput, optFns ...func(*Options)) (*CreateSpaceOutput, error) 
 DeleteSpace(ctx context.Context, params *DeleteSpaceInput, optFns ...func(*Options)) (*DeleteSpaceOutput, error) 
 DeregisterAdmin(ctx context.Context, params *DeregisterAdminInput, optFns ...func(*Options)) (*DeregisterAdminOutput, error) 
 GetSpace(ctx context.Context, params *GetSpaceInput, optFns ...func(*Options)) (*GetSpaceOutput, error) 
 ListSpaces(ctx context.Context, params *ListSpacesInput, optFns ...func(*Options)) (*ListSpacesOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 RegisterAdmin(ctx context.Context, params *RegisterAdminInput, optFns ...func(*Options)) (*RegisterAdminOutput, error) 
 SendInvites(ctx context.Context, params *SendInvitesInput, optFns ...func(*Options)) (*SendInvitesOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateSpace(ctx context.Context, params *UpdateSpaceInput, optFns ...func(*Options)) (*UpdateSpaceOutput, error) 
}
