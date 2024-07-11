
package paymentcryptography

import (
    "github.com/aws/aws-sdk-go-v2/service/paymentcryptography"
)

// IPaymentcryptography defines the interface for paymentcryptography
type IPaymentcryptography interface {
 Options() Options 
 CreateAlias(ctx context.Context, params *CreateAliasInput, optFns ...func(*Options)) (*CreateAliasOutput, error) 
 CreateKey(ctx context.Context, params *CreateKeyInput, optFns ...func(*Options)) (*CreateKeyOutput, error) 
 DeleteAlias(ctx context.Context, params *DeleteAliasInput, optFns ...func(*Options)) (*DeleteAliasOutput, error) 
 DeleteKey(ctx context.Context, params *DeleteKeyInput, optFns ...func(*Options)) (*DeleteKeyOutput, error) 
 ExportKey(ctx context.Context, params *ExportKeyInput, optFns ...func(*Options)) (*ExportKeyOutput, error) 
 GetAlias(ctx context.Context, params *GetAliasInput, optFns ...func(*Options)) (*GetAliasOutput, error) 
 GetKey(ctx context.Context, params *GetKeyInput, optFns ...func(*Options)) (*GetKeyOutput, error) 
 GetParametersForExport(ctx context.Context, params *GetParametersForExportInput, optFns ...func(*Options)) (*GetParametersForExportOutput, error) 
 GetParametersForImport(ctx context.Context, params *GetParametersForImportInput, optFns ...func(*Options)) (*GetParametersForImportOutput, error) 
 GetPublicKeyCertificate(ctx context.Context, params *GetPublicKeyCertificateInput, optFns ...func(*Options)) (*GetPublicKeyCertificateOutput, error) 
 ImportKey(ctx context.Context, params *ImportKeyInput, optFns ...func(*Options)) (*ImportKeyOutput, error) 
 ListAliases(ctx context.Context, params *ListAliasesInput, optFns ...func(*Options)) (*ListAliasesOutput, error) 
 ListKeys(ctx context.Context, params *ListKeysInput, optFns ...func(*Options)) (*ListKeysOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 RestoreKey(ctx context.Context, params *RestoreKeyInput, optFns ...func(*Options)) (*RestoreKeyOutput, error) 
 StartKeyUsage(ctx context.Context, params *StartKeyUsageInput, optFns ...func(*Options)) (*StartKeyUsageOutput, error) 
 StopKeyUsage(ctx context.Context, params *StopKeyUsageInput, optFns ...func(*Options)) (*StopKeyUsageOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateAlias(ctx context.Context, params *UpdateAliasInput, optFns ...func(*Options)) (*UpdateAliasOutput, error) 
}
