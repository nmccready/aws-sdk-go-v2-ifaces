
package cloudhsm

import (
    "github.com/aws/aws-sdk-go-v2/service/cloudhsm"
)

// IClient defines the interface for cloudhsm
type IClient interface {
 Options() Options 
 AddTagsToResource(ctx context.Context, params *AddTagsToResourceInput, optFns ...func(*Options)) (*AddTagsToResourceOutput, error) 
 CreateHapg(ctx context.Context, params *CreateHapgInput, optFns ...func(*Options)) (*CreateHapgOutput, error) 
 CreateHsm(ctx context.Context, params *CreateHsmInput, optFns ...func(*Options)) (*CreateHsmOutput, error) 
 CreateLunaClient(ctx context.Context, params *CreateLunaClientInput, optFns ...func(*Options)) (*CreateLunaClientOutput, error) 
 DeleteHapg(ctx context.Context, params *DeleteHapgInput, optFns ...func(*Options)) (*DeleteHapgOutput, error) 
 DeleteHsm(ctx context.Context, params *DeleteHsmInput, optFns ...func(*Options)) (*DeleteHsmOutput, error) 
 DeleteLunaClient(ctx context.Context, params *DeleteLunaClientInput, optFns ...func(*Options)) (*DeleteLunaClientOutput, error) 
 DescribeHapg(ctx context.Context, params *DescribeHapgInput, optFns ...func(*Options)) (*DescribeHapgOutput, error) 
 DescribeHsm(ctx context.Context, params *DescribeHsmInput, optFns ...func(*Options)) (*DescribeHsmOutput, error) 
 DescribeLunaClient(ctx context.Context, params *DescribeLunaClientInput, optFns ...func(*Options)) (*DescribeLunaClientOutput, error) 
 GetConfig(ctx context.Context, params *GetConfigInput, optFns ...func(*Options)) (*GetConfigOutput, error) 
 ListAvailableZones(ctx context.Context, params *ListAvailableZonesInput, optFns ...func(*Options)) (*ListAvailableZonesOutput, error) 
 ListHapgs(ctx context.Context, params *ListHapgsInput, optFns ...func(*Options)) (*ListHapgsOutput, error) 
 ListHsms(ctx context.Context, params *ListHsmsInput, optFns ...func(*Options)) (*ListHsmsOutput, error) 
 ListLunaClients(ctx context.Context, params *ListLunaClientsInput, optFns ...func(*Options)) (*ListLunaClientsOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 ModifyHapg(ctx context.Context, params *ModifyHapgInput, optFns ...func(*Options)) (*ModifyHapgOutput, error) 
 ModifyHsm(ctx context.Context, params *ModifyHsmInput, optFns ...func(*Options)) (*ModifyHsmOutput, error) 
 ModifyLunaClient(ctx context.Context, params *ModifyLunaClientInput, optFns ...func(*Options)) (*ModifyLunaClientOutput, error) 
 RemoveTagsFromResource(ctx context.Context, params *RemoveTagsFromResourceInput, optFns ...func(*Options)) (*RemoveTagsFromResourceOutput, error) 
 ID() string 
 HandleDeserialize(ctx context.Context, in middleware.DeserializeInput, next middleware.DeserializeHandler) (
 ID() string 
 HandleDeserialize(ctx context.Context, in middleware.DeserializeInput, next middleware.DeserializeHandler) (
 ID() string 
 HandleDeserialize(ctx context.Context, in middleware.DeserializeInput, next middleware.DeserializeHandler) (
 ID() string 
 HandleDeserialize(ctx context.Context, in middleware.DeserializeInput, next middleware.DeserializeHandler) (
 ID() string 
 HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 ID() string 
 HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 ID() string 
 HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 ID() string 
 HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 ID() string 
 HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
 ID() string 
 HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
 ID() string 
 HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
}
