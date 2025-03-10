
package taxsettings_iface

// This file was generated by scripts/gen_ifaces.sh do not edit it manually
import (
    "context"
    . "github.com/aws/aws-sdk-go-v2/service/taxsettings"
)

// IClient defines the interface for taxsettings
type IClient interface {
 Options() Options 
 BatchDeleteTaxRegistration(ctx context.Context, params *BatchDeleteTaxRegistrationInput, optFns ...func(*Options)) (*BatchDeleteTaxRegistrationOutput, error) 
 BatchGetTaxExemptions(ctx context.Context, params *BatchGetTaxExemptionsInput, optFns ...func(*Options)) (*BatchGetTaxExemptionsOutput, error) 
 BatchPutTaxRegistration(ctx context.Context, params *BatchPutTaxRegistrationInput, optFns ...func(*Options)) (*BatchPutTaxRegistrationOutput, error) 
 DeleteSupplementalTaxRegistration(ctx context.Context, params *DeleteSupplementalTaxRegistrationInput, optFns ...func(*Options)) (*DeleteSupplementalTaxRegistrationOutput, error) 
 DeleteTaxRegistration(ctx context.Context, params *DeleteTaxRegistrationInput, optFns ...func(*Options)) (*DeleteTaxRegistrationOutput, error) 
 GetTaxExemptionTypes(ctx context.Context, params *GetTaxExemptionTypesInput, optFns ...func(*Options)) (*GetTaxExemptionTypesOutput, error) 
 GetTaxInheritance(ctx context.Context, params *GetTaxInheritanceInput, optFns ...func(*Options)) (*GetTaxInheritanceOutput, error) 
 GetTaxRegistration(ctx context.Context, params *GetTaxRegistrationInput, optFns ...func(*Options)) (*GetTaxRegistrationOutput, error) 
 GetTaxRegistrationDocument(ctx context.Context, params *GetTaxRegistrationDocumentInput, optFns ...func(*Options)) (*GetTaxRegistrationDocumentOutput, error) 
 ListSupplementalTaxRegistrations(ctx context.Context, params *ListSupplementalTaxRegistrationsInput, optFns ...func(*Options)) (*ListSupplementalTaxRegistrationsOutput, error) 
 ListTaxExemptions(ctx context.Context, params *ListTaxExemptionsInput, optFns ...func(*Options)) (*ListTaxExemptionsOutput, error) 
 ListTaxRegistrations(ctx context.Context, params *ListTaxRegistrationsInput, optFns ...func(*Options)) (*ListTaxRegistrationsOutput, error) 
 PutSupplementalTaxRegistration(ctx context.Context, params *PutSupplementalTaxRegistrationInput, optFns ...func(*Options)) (*PutSupplementalTaxRegistrationOutput, error) 
 PutTaxExemption(ctx context.Context, params *PutTaxExemptionInput, optFns ...func(*Options)) (*PutTaxExemptionOutput, error) 
 PutTaxInheritance(ctx context.Context, params *PutTaxInheritanceInput, optFns ...func(*Options)) (*PutTaxInheritanceOutput, error) 
 PutTaxRegistration(ctx context.Context, params *PutTaxRegistrationInput, optFns ...func(*Options)) (*PutTaxRegistrationOutput, error) 
}
