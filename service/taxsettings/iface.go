
package taxsettings

import (
    "github.com/aws/aws-sdk-go-v2/service/taxsettings"
)

// ITaxsettings defines the interface for taxsettings
type ITaxsettings interface {
 Options() Options 
 BatchDeleteTaxRegistration(ctx context.Context, params *BatchDeleteTaxRegistrationInput, optFns ...func(*Options)) (*BatchDeleteTaxRegistrationOutput, error) 
 BatchPutTaxRegistration(ctx context.Context, params *BatchPutTaxRegistrationInput, optFns ...func(*Options)) (*BatchPutTaxRegistrationOutput, error) 
 DeleteTaxRegistration(ctx context.Context, params *DeleteTaxRegistrationInput, optFns ...func(*Options)) (*DeleteTaxRegistrationOutput, error) 
 GetTaxRegistration(ctx context.Context, params *GetTaxRegistrationInput, optFns ...func(*Options)) (*GetTaxRegistrationOutput, error) 
 GetTaxRegistrationDocument(ctx context.Context, params *GetTaxRegistrationDocumentInput, optFns ...func(*Options)) (*GetTaxRegistrationDocumentOutput, error) 
 ListTaxRegistrations(ctx context.Context, params *ListTaxRegistrationsInput, optFns ...func(*Options)) (*ListTaxRegistrationsOutput, error) 
 PutTaxRegistration(ctx context.Context, params *PutTaxRegistrationInput, optFns ...func(*Options)) (*PutTaxRegistrationOutput, error) 
}
