
package paymentcryptographydata

import (
    "github.com/aws/aws-sdk-go-v2/service/paymentcryptographydata"
)

// IClient defines the interface for paymentcryptographydata
type IClient interface {
 Options() Options 
 DecryptData(ctx context.Context, params *DecryptDataInput, optFns ...func(*Options)) (*DecryptDataOutput, error) 
 EncryptData(ctx context.Context, params *EncryptDataInput, optFns ...func(*Options)) (*EncryptDataOutput, error) 
 GenerateCardValidationData(ctx context.Context, params *GenerateCardValidationDataInput, optFns ...func(*Options)) (*GenerateCardValidationDataOutput, error) 
 GenerateMac(ctx context.Context, params *GenerateMacInput, optFns ...func(*Options)) (*GenerateMacOutput, error) 
 GeneratePinData(ctx context.Context, params *GeneratePinDataInput, optFns ...func(*Options)) (*GeneratePinDataOutput, error) 
 ReEncryptData(ctx context.Context, params *ReEncryptDataInput, optFns ...func(*Options)) (*ReEncryptDataOutput, error) 
 TranslatePinData(ctx context.Context, params *TranslatePinDataInput, optFns ...func(*Options)) (*TranslatePinDataOutput, error) 
 VerifyAuthRequestCryptogram(ctx context.Context, params *VerifyAuthRequestCryptogramInput, optFns ...func(*Options)) (*VerifyAuthRequestCryptogramOutput, error) 
 VerifyCardValidationData(ctx context.Context, params *VerifyCardValidationDataInput, optFns ...func(*Options)) (*VerifyCardValidationDataOutput, error) 
 VerifyMac(ctx context.Context, params *VerifyMacInput, optFns ...func(*Options)) (*VerifyMacOutput, error) 
 VerifyPinData(ctx context.Context, params *VerifyPinDataInput, optFns ...func(*Options)) (*VerifyPinDataOutput, error) 
}
