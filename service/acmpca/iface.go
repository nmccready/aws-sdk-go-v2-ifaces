
package acmpca

import (
    "github.com/aws/aws-sdk-go-v2/service/acmpca"
)

// IClient defines the interface for acmpca
type IClient interface {
 Options() Options 
 CreateCertificateAuthority(ctx context.Context, params *CreateCertificateAuthorityInput, optFns ...func(*Options)) (*CreateCertificateAuthorityOutput, error) 
 CreateCertificateAuthorityAuditReport(ctx context.Context, params *CreateCertificateAuthorityAuditReportInput, optFns ...func(*Options)) (*CreateCertificateAuthorityAuditReportOutput, error) 
 CreatePermission(ctx context.Context, params *CreatePermissionInput, optFns ...func(*Options)) (*CreatePermissionOutput, error) 
 DeleteCertificateAuthority(ctx context.Context, params *DeleteCertificateAuthorityInput, optFns ...func(*Options)) (*DeleteCertificateAuthorityOutput, error) 
 DeletePermission(ctx context.Context, params *DeletePermissionInput, optFns ...func(*Options)) (*DeletePermissionOutput, error) 
 DeletePolicy(ctx context.Context, params *DeletePolicyInput, optFns ...func(*Options)) (*DeletePolicyOutput, error) 
 DescribeCertificateAuthority(ctx context.Context, params *DescribeCertificateAuthorityInput, optFns ...func(*Options)) (*DescribeCertificateAuthorityOutput, error) 
 DescribeCertificateAuthorityAuditReport(ctx context.Context, params *DescribeCertificateAuthorityAuditReportInput, optFns ...func(*Options)) (*DescribeCertificateAuthorityAuditReportOutput, error) 
 GetCertificate(ctx context.Context, params *GetCertificateInput, optFns ...func(*Options)) (*GetCertificateOutput, error) 
 GetCertificateAuthorityCertificate(ctx context.Context, params *GetCertificateAuthorityCertificateInput, optFns ...func(*Options)) (*GetCertificateAuthorityCertificateOutput, error) 
 GetCertificateAuthorityCsr(ctx context.Context, params *GetCertificateAuthorityCsrInput, optFns ...func(*Options)) (*GetCertificateAuthorityCsrOutput, error) 
 GetPolicy(ctx context.Context, params *GetPolicyInput, optFns ...func(*Options)) (*GetPolicyOutput, error) 
 ImportCertificateAuthorityCertificate(ctx context.Context, params *ImportCertificateAuthorityCertificateInput, optFns ...func(*Options)) (*ImportCertificateAuthorityCertificateOutput, error) 
 IssueCertificate(ctx context.Context, params *IssueCertificateInput, optFns ...func(*Options)) (*IssueCertificateOutput, error) 
 ListCertificateAuthorities(ctx context.Context, params *ListCertificateAuthoritiesInput, optFns ...func(*Options)) (*ListCertificateAuthoritiesOutput, error) 
 ListPermissions(ctx context.Context, params *ListPermissionsInput, optFns ...func(*Options)) (*ListPermissionsOutput, error) 
 ListTags(ctx context.Context, params *ListTagsInput, optFns ...func(*Options)) (*ListTagsOutput, error) 
 PutPolicy(ctx context.Context, params *PutPolicyInput, optFns ...func(*Options)) (*PutPolicyOutput, error) 
 RestoreCertificateAuthority(ctx context.Context, params *RestoreCertificateAuthorityInput, optFns ...func(*Options)) (*RestoreCertificateAuthorityOutput, error) 
 RevokeCertificate(ctx context.Context, params *RevokeCertificateInput, optFns ...func(*Options)) (*RevokeCertificateOutput, error) 
 TagCertificateAuthority(ctx context.Context, params *TagCertificateAuthorityInput, optFns ...func(*Options)) (*TagCertificateAuthorityOutput, error) 
 UntagCertificateAuthority(ctx context.Context, params *UntagCertificateAuthorityInput, optFns ...func(*Options)) (*UntagCertificateAuthorityOutput, error) 
 UpdateCertificateAuthority(ctx context.Context, params *UpdateCertificateAuthorityInput, optFns ...func(*Options)) (*UpdateCertificateAuthorityOutput, error) 
}
