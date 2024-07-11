
package verifiedpermissions

import (
    "github.com/aws/aws-sdk-go-v2/service/verifiedpermissions"
)

// IClient defines the interface for verifiedpermissions
type IClient interface {
 Options() Options 
 BatchIsAuthorized(ctx context.Context, params *BatchIsAuthorizedInput, optFns ...func(*Options)) (*BatchIsAuthorizedOutput, error) 
 BatchIsAuthorizedWithToken(ctx context.Context, params *BatchIsAuthorizedWithTokenInput, optFns ...func(*Options)) (*BatchIsAuthorizedWithTokenOutput, error) 
 CreateIdentitySource(ctx context.Context, params *CreateIdentitySourceInput, optFns ...func(*Options)) (*CreateIdentitySourceOutput, error) 
 CreatePolicy(ctx context.Context, params *CreatePolicyInput, optFns ...func(*Options)) (*CreatePolicyOutput, error) 
 CreatePolicyStore(ctx context.Context, params *CreatePolicyStoreInput, optFns ...func(*Options)) (*CreatePolicyStoreOutput, error) 
 CreatePolicyTemplate(ctx context.Context, params *CreatePolicyTemplateInput, optFns ...func(*Options)) (*CreatePolicyTemplateOutput, error) 
 DeleteIdentitySource(ctx context.Context, params *DeleteIdentitySourceInput, optFns ...func(*Options)) (*DeleteIdentitySourceOutput, error) 
 DeletePolicy(ctx context.Context, params *DeletePolicyInput, optFns ...func(*Options)) (*DeletePolicyOutput, error) 
 DeletePolicyStore(ctx context.Context, params *DeletePolicyStoreInput, optFns ...func(*Options)) (*DeletePolicyStoreOutput, error) 
 DeletePolicyTemplate(ctx context.Context, params *DeletePolicyTemplateInput, optFns ...func(*Options)) (*DeletePolicyTemplateOutput, error) 
 GetIdentitySource(ctx context.Context, params *GetIdentitySourceInput, optFns ...func(*Options)) (*GetIdentitySourceOutput, error) 
 GetPolicy(ctx context.Context, params *GetPolicyInput, optFns ...func(*Options)) (*GetPolicyOutput, error) 
 GetPolicyStore(ctx context.Context, params *GetPolicyStoreInput, optFns ...func(*Options)) (*GetPolicyStoreOutput, error) 
 GetPolicyTemplate(ctx context.Context, params *GetPolicyTemplateInput, optFns ...func(*Options)) (*GetPolicyTemplateOutput, error) 
 GetSchema(ctx context.Context, params *GetSchemaInput, optFns ...func(*Options)) (*GetSchemaOutput, error) 
 IsAuthorized(ctx context.Context, params *IsAuthorizedInput, optFns ...func(*Options)) (*IsAuthorizedOutput, error) 
 IsAuthorizedWithToken(ctx context.Context, params *IsAuthorizedWithTokenInput, optFns ...func(*Options)) (*IsAuthorizedWithTokenOutput, error) 
 ListIdentitySources(ctx context.Context, params *ListIdentitySourcesInput, optFns ...func(*Options)) (*ListIdentitySourcesOutput, error) 
 ListPolicies(ctx context.Context, params *ListPoliciesInput, optFns ...func(*Options)) (*ListPoliciesOutput, error) 
 ListPolicyStores(ctx context.Context, params *ListPolicyStoresInput, optFns ...func(*Options)) (*ListPolicyStoresOutput, error) 
 ListPolicyTemplates(ctx context.Context, params *ListPolicyTemplatesInput, optFns ...func(*Options)) (*ListPolicyTemplatesOutput, error) 
 PutSchema(ctx context.Context, params *PutSchemaInput, optFns ...func(*Options)) (*PutSchemaOutput, error) 
 UpdateIdentitySource(ctx context.Context, params *UpdateIdentitySourceInput, optFns ...func(*Options)) (*UpdateIdentitySourceOutput, error) 
 UpdatePolicy(ctx context.Context, params *UpdatePolicyInput, optFns ...func(*Options)) (*UpdatePolicyOutput, error) 
 UpdatePolicyStore(ctx context.Context, params *UpdatePolicyStoreInput, optFns ...func(*Options)) (*UpdatePolicyStoreOutput, error) 
 UpdatePolicyTemplate(ctx context.Context, params *UpdatePolicyTemplateInput, optFns ...func(*Options)) (*UpdatePolicyTemplateOutput, error) 
}
