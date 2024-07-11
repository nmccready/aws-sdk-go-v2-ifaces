
package pcaconnectorscep

import (
    "github.com/aws/aws-sdk-go-v2/service/pcaconnectorscep"
)

// IPcaconnectorscep defines the interface for pcaconnectorscep
type IPcaconnectorscep interface {
 Options() Options 
 CreateChallenge(ctx context.Context, params *CreateChallengeInput, optFns ...func(*Options)) (*CreateChallengeOutput, error) 
 CreateConnector(ctx context.Context, params *CreateConnectorInput, optFns ...func(*Options)) (*CreateConnectorOutput, error) 
 DeleteChallenge(ctx context.Context, params *DeleteChallengeInput, optFns ...func(*Options)) (*DeleteChallengeOutput, error) 
 DeleteConnector(ctx context.Context, params *DeleteConnectorInput, optFns ...func(*Options)) (*DeleteConnectorOutput, error) 
 GetChallengeMetadata(ctx context.Context, params *GetChallengeMetadataInput, optFns ...func(*Options)) (*GetChallengeMetadataOutput, error) 
 GetChallengePassword(ctx context.Context, params *GetChallengePasswordInput, optFns ...func(*Options)) (*GetChallengePasswordOutput, error) 
 GetConnector(ctx context.Context, params *GetConnectorInput, optFns ...func(*Options)) (*GetConnectorOutput, error) 
 ListChallengeMetadata(ctx context.Context, params *ListChallengeMetadataInput, optFns ...func(*Options)) (*ListChallengeMetadataOutput, error) 
 ListConnectors(ctx context.Context, params *ListConnectorsInput, optFns ...func(*Options)) (*ListConnectorsOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
}
