
package b2bi

import (
    "github.com/aws/aws-sdk-go-v2/service/b2bi"
)

// IB2bi defines the interface for b2bi
type IB2bi interface {
 Options() Options 
 CreateCapability(ctx context.Context, params *CreateCapabilityInput, optFns ...func(*Options)) (*CreateCapabilityOutput, error) 
 CreatePartnership(ctx context.Context, params *CreatePartnershipInput, optFns ...func(*Options)) (*CreatePartnershipOutput, error) 
 CreateProfile(ctx context.Context, params *CreateProfileInput, optFns ...func(*Options)) (*CreateProfileOutput, error) 
 CreateTransformer(ctx context.Context, params *CreateTransformerInput, optFns ...func(*Options)) (*CreateTransformerOutput, error) 
 DeleteCapability(ctx context.Context, params *DeleteCapabilityInput, optFns ...func(*Options)) (*DeleteCapabilityOutput, error) 
 DeletePartnership(ctx context.Context, params *DeletePartnershipInput, optFns ...func(*Options)) (*DeletePartnershipOutput, error) 
 DeleteProfile(ctx context.Context, params *DeleteProfileInput, optFns ...func(*Options)) (*DeleteProfileOutput, error) 
 DeleteTransformer(ctx context.Context, params *DeleteTransformerInput, optFns ...func(*Options)) (*DeleteTransformerOutput, error) 
 GetCapability(ctx context.Context, params *GetCapabilityInput, optFns ...func(*Options)) (*GetCapabilityOutput, error) 
 GetPartnership(ctx context.Context, params *GetPartnershipInput, optFns ...func(*Options)) (*GetPartnershipOutput, error) 
 GetProfile(ctx context.Context, params *GetProfileInput, optFns ...func(*Options)) (*GetProfileOutput, error) 
 GetTransformer(ctx context.Context, params *GetTransformerInput, optFns ...func(*Options)) (*GetTransformerOutput, error) 
 GetTransformerJob(ctx context.Context, params *GetTransformerJobInput, optFns ...func(*Options)) (*GetTransformerJobOutput, error) 
 ListCapabilities(ctx context.Context, params *ListCapabilitiesInput, optFns ...func(*Options)) (*ListCapabilitiesOutput, error) 
 ListPartnerships(ctx context.Context, params *ListPartnershipsInput, optFns ...func(*Options)) (*ListPartnershipsOutput, error) 
 ListProfiles(ctx context.Context, params *ListProfilesInput, optFns ...func(*Options)) (*ListProfilesOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 ListTransformers(ctx context.Context, params *ListTransformersInput, optFns ...func(*Options)) (*ListTransformersOutput, error) 
 StartTransformerJob(ctx context.Context, params *StartTransformerJobInput, optFns ...func(*Options)) (*StartTransformerJobOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 TestMapping(ctx context.Context, params *TestMappingInput, optFns ...func(*Options)) (*TestMappingOutput, error) 
 TestParsing(ctx context.Context, params *TestParsingInput, optFns ...func(*Options)) (*TestParsingOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateCapability(ctx context.Context, params *UpdateCapabilityInput, optFns ...func(*Options)) (*UpdateCapabilityOutput, error) 
 UpdatePartnership(ctx context.Context, params *UpdatePartnershipInput, optFns ...func(*Options)) (*UpdatePartnershipOutput, error) 
 UpdateProfile(ctx context.Context, params *UpdateProfileInput, optFns ...func(*Options)) (*UpdateProfileOutput, error) 
 UpdateTransformer(ctx context.Context, params *UpdateTransformerInput, optFns ...func(*Options)) (*UpdateTransformerOutput, error) 
}
