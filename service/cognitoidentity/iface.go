
package cognitoidentity

import (
    "github.com/aws/aws-sdk-go-v2/service/cognitoidentity"
)

// IClient defines the interface for cognitoidentity
type IClient interface {
 Options() Options 
 CreateIdentityPool(ctx context.Context, params *CreateIdentityPoolInput, optFns ...func(*Options)) (*CreateIdentityPoolOutput, error) 
 DeleteIdentities(ctx context.Context, params *DeleteIdentitiesInput, optFns ...func(*Options)) (*DeleteIdentitiesOutput, error) 
 DeleteIdentityPool(ctx context.Context, params *DeleteIdentityPoolInput, optFns ...func(*Options)) (*DeleteIdentityPoolOutput, error) 
 DescribeIdentity(ctx context.Context, params *DescribeIdentityInput, optFns ...func(*Options)) (*DescribeIdentityOutput, error) 
 DescribeIdentityPool(ctx context.Context, params *DescribeIdentityPoolInput, optFns ...func(*Options)) (*DescribeIdentityPoolOutput, error) 
 GetCredentialsForIdentity(ctx context.Context, params *GetCredentialsForIdentityInput, optFns ...func(*Options)) (*GetCredentialsForIdentityOutput, error) 
 GetId(ctx context.Context, params *GetIdInput, optFns ...func(*Options)) (*GetIdOutput, error) 
 GetIdentityPoolRoles(ctx context.Context, params *GetIdentityPoolRolesInput, optFns ...func(*Options)) (*GetIdentityPoolRolesOutput, error) 
 GetOpenIdToken(ctx context.Context, params *GetOpenIdTokenInput, optFns ...func(*Options)) (*GetOpenIdTokenOutput, error) 
 GetOpenIdTokenForDeveloperIdentity(ctx context.Context, params *GetOpenIdTokenForDeveloperIdentityInput, optFns ...func(*Options)) (*GetOpenIdTokenForDeveloperIdentityOutput, error) 
 GetPrincipalTagAttributeMap(ctx context.Context, params *GetPrincipalTagAttributeMapInput, optFns ...func(*Options)) (*GetPrincipalTagAttributeMapOutput, error) 
 ListIdentities(ctx context.Context, params *ListIdentitiesInput, optFns ...func(*Options)) (*ListIdentitiesOutput, error) 
 ListIdentityPools(ctx context.Context, params *ListIdentityPoolsInput, optFns ...func(*Options)) (*ListIdentityPoolsOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 LookupDeveloperIdentity(ctx context.Context, params *LookupDeveloperIdentityInput, optFns ...func(*Options)) (*LookupDeveloperIdentityOutput, error) 
 MergeDeveloperIdentities(ctx context.Context, params *MergeDeveloperIdentitiesInput, optFns ...func(*Options)) (*MergeDeveloperIdentitiesOutput, error) 
 SetIdentityPoolRoles(ctx context.Context, params *SetIdentityPoolRolesInput, optFns ...func(*Options)) (*SetIdentityPoolRolesOutput, error) 
 SetPrincipalTagAttributeMap(ctx context.Context, params *SetPrincipalTagAttributeMapInput, optFns ...func(*Options)) (*SetPrincipalTagAttributeMapOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UnlinkDeveloperIdentity(ctx context.Context, params *UnlinkDeveloperIdentityInput, optFns ...func(*Options)) (*UnlinkDeveloperIdentityOutput, error) 
 UnlinkIdentity(ctx context.Context, params *UnlinkIdentityInput, optFns ...func(*Options)) (*UnlinkIdentityOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateIdentityPool(ctx context.Context, params *UpdateIdentityPoolInput, optFns ...func(*Options)) (*UpdateIdentityPoolOutput, error) 
}
