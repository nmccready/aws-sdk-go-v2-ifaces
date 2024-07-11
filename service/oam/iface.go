
package oam

import (
    "github.com/aws/aws-sdk-go-v2/service/oam"
)

// IClient defines the interface for oam
type IClient interface {
 Options() Options 
 CreateLink(ctx context.Context, params *CreateLinkInput, optFns ...func(*Options)) (*CreateLinkOutput, error) 
 CreateSink(ctx context.Context, params *CreateSinkInput, optFns ...func(*Options)) (*CreateSinkOutput, error) 
 DeleteLink(ctx context.Context, params *DeleteLinkInput, optFns ...func(*Options)) (*DeleteLinkOutput, error) 
 DeleteSink(ctx context.Context, params *DeleteSinkInput, optFns ...func(*Options)) (*DeleteSinkOutput, error) 
 GetLink(ctx context.Context, params *GetLinkInput, optFns ...func(*Options)) (*GetLinkOutput, error) 
 GetSink(ctx context.Context, params *GetSinkInput, optFns ...func(*Options)) (*GetSinkOutput, error) 
 GetSinkPolicy(ctx context.Context, params *GetSinkPolicyInput, optFns ...func(*Options)) (*GetSinkPolicyOutput, error) 
 ListAttachedLinks(ctx context.Context, params *ListAttachedLinksInput, optFns ...func(*Options)) (*ListAttachedLinksOutput, error) 
 ListLinks(ctx context.Context, params *ListLinksInput, optFns ...func(*Options)) (*ListLinksOutput, error) 
 ListSinks(ctx context.Context, params *ListSinksInput, optFns ...func(*Options)) (*ListSinksOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 PutSinkPolicy(ctx context.Context, params *PutSinkPolicyInput, optFns ...func(*Options)) (*PutSinkPolicyOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateLink(ctx context.Context, params *UpdateLinkInput, optFns ...func(*Options)) (*UpdateLinkOutput, error) 
}
