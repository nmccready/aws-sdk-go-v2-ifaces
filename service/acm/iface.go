
package acm

import (
    "github.com/aws/aws-sdk-go-v2/service/acm"
)

// IClient defines the interface for acm
type IClient interface {
 Options() Options 
 AddTagsToCertificate(ctx context.Context, params *AddTagsToCertificateInput, optFns ...func(*Options)) (*AddTagsToCertificateOutput, error) 
 DeleteCertificate(ctx context.Context, params *DeleteCertificateInput, optFns ...func(*Options)) (*DeleteCertificateOutput, error) 
 DescribeCertificate(ctx context.Context, params *DescribeCertificateInput, optFns ...func(*Options)) (*DescribeCertificateOutput, error) 
 ExportCertificate(ctx context.Context, params *ExportCertificateInput, optFns ...func(*Options)) (*ExportCertificateOutput, error) 
 GetAccountConfiguration(ctx context.Context, params *GetAccountConfigurationInput, optFns ...func(*Options)) (*GetAccountConfigurationOutput, error) 
 GetCertificate(ctx context.Context, params *GetCertificateInput, optFns ...func(*Options)) (*GetCertificateOutput, error) 
 ImportCertificate(ctx context.Context, params *ImportCertificateInput, optFns ...func(*Options)) (*ImportCertificateOutput, error) 
 ListCertificates(ctx context.Context, params *ListCertificatesInput, optFns ...func(*Options)) (*ListCertificatesOutput, error) 
 ListTagsForCertificate(ctx context.Context, params *ListTagsForCertificateInput, optFns ...func(*Options)) (*ListTagsForCertificateOutput, error) 
 PutAccountConfiguration(ctx context.Context, params *PutAccountConfigurationInput, optFns ...func(*Options)) (*PutAccountConfigurationOutput, error) 
 RemoveTagsFromCertificate(ctx context.Context, params *RemoveTagsFromCertificateInput, optFns ...func(*Options)) (*RemoveTagsFromCertificateOutput, error) 
 RenewCertificate(ctx context.Context, params *RenewCertificateInput, optFns ...func(*Options)) (*RenewCertificateOutput, error) 
 RequestCertificate(ctx context.Context, params *RequestCertificateInput, optFns ...func(*Options)) (*RequestCertificateOutput, error) 
 ResendValidationEmail(ctx context.Context, params *ResendValidationEmailInput, optFns ...func(*Options)) (*ResendValidationEmailOutput, error) 
 UpdateCertificateOptions(ctx context.Context, params *UpdateCertificateOptionsInput, optFns ...func(*Options)) (*UpdateCertificateOptionsOutput, error) 
}
