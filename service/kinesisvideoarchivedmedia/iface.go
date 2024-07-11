
package kinesisvideoarchivedmedia

import (
    "github.com/aws/aws-sdk-go-v2/service/kinesisvideoarchivedmedia"
)

// IKinesisvideoarchivedmedia defines the interface for kinesisvideoarchivedmedia
type IKinesisvideoarchivedmedia interface {
 Options() Options 
 GetClip(ctx context.Context, params *GetClipInput, optFns ...func(*Options)) (*GetClipOutput, error) 
 GetDASHStreamingSessionURL(ctx context.Context, params *GetDASHStreamingSessionURLInput, optFns ...func(*Options)) (*GetDASHStreamingSessionURLOutput, error) 
 GetHLSStreamingSessionURL(ctx context.Context, params *GetHLSStreamingSessionURLInput, optFns ...func(*Options)) (*GetHLSStreamingSessionURLOutput, error) 
 GetImages(ctx context.Context, params *GetImagesInput, optFns ...func(*Options)) (*GetImagesOutput, error) 
 GetMediaForFragmentList(ctx context.Context, params *GetMediaForFragmentListInput, optFns ...func(*Options)) (*GetMediaForFragmentListOutput, error) 
 ListFragments(ctx context.Context, params *ListFragmentsInput, optFns ...func(*Options)) (*ListFragmentsOutput, error) 
}
