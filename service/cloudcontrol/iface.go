
package cloudcontrol

import (
    "github.com/aws/aws-sdk-go-v2/service/cloudcontrol"
)

// ICloudcontrol defines the interface for cloudcontrol
type ICloudcontrol interface {
 Options() Options 
 CancelResourceRequest(ctx context.Context, params *CancelResourceRequestInput, optFns ...func(*Options)) (*CancelResourceRequestOutput, error) 
 CreateResource(ctx context.Context, params *CreateResourceInput, optFns ...func(*Options)) (*CreateResourceOutput, error) 
 DeleteResource(ctx context.Context, params *DeleteResourceInput, optFns ...func(*Options)) (*DeleteResourceOutput, error) 
 GetResource(ctx context.Context, params *GetResourceInput, optFns ...func(*Options)) (*GetResourceOutput, error) 
 GetResourceRequestStatus(ctx context.Context, params *GetResourceRequestStatusInput, optFns ...func(*Options)) (*GetResourceRequestStatusOutput, error) 
 ListResourceRequests(ctx context.Context, params *ListResourceRequestsInput, optFns ...func(*Options)) (*ListResourceRequestsOutput, error) 
 ListResources(ctx context.Context, params *ListResourcesInput, optFns ...func(*Options)) (*ListResourcesOutput, error) 
 UpdateResource(ctx context.Context, params *UpdateResourceInput, optFns ...func(*Options)) (*UpdateResourceOutput, error) 
}
