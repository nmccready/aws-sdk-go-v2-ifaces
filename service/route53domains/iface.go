
package route53domains

import (
    "github.com/aws/aws-sdk-go-v2/service/route53domains"
)

// IRoute53domains defines the interface for route53domains
type IRoute53domains interface {
 Options() Options 
 AcceptDomainTransferFromAnotherAwsAccount(ctx context.Context, params *AcceptDomainTransferFromAnotherAwsAccountInput, optFns ...func(*Options)) (*AcceptDomainTransferFromAnotherAwsAccountOutput, error) 
 AssociateDelegationSignerToDomain(ctx context.Context, params *AssociateDelegationSignerToDomainInput, optFns ...func(*Options)) (*AssociateDelegationSignerToDomainOutput, error) 
 CancelDomainTransferToAnotherAwsAccount(ctx context.Context, params *CancelDomainTransferToAnotherAwsAccountInput, optFns ...func(*Options)) (*CancelDomainTransferToAnotherAwsAccountOutput, error) 
 CheckDomainAvailability(ctx context.Context, params *CheckDomainAvailabilityInput, optFns ...func(*Options)) (*CheckDomainAvailabilityOutput, error) 
 CheckDomainTransferability(ctx context.Context, params *CheckDomainTransferabilityInput, optFns ...func(*Options)) (*CheckDomainTransferabilityOutput, error) 
 DeleteDomain(ctx context.Context, params *DeleteDomainInput, optFns ...func(*Options)) (*DeleteDomainOutput, error) 
 DeleteTagsForDomain(ctx context.Context, params *DeleteTagsForDomainInput, optFns ...func(*Options)) (*DeleteTagsForDomainOutput, error) 
 DisableDomainAutoRenew(ctx context.Context, params *DisableDomainAutoRenewInput, optFns ...func(*Options)) (*DisableDomainAutoRenewOutput, error) 
 DisableDomainTransferLock(ctx context.Context, params *DisableDomainTransferLockInput, optFns ...func(*Options)) (*DisableDomainTransferLockOutput, error) 
 DisassociateDelegationSignerFromDomain(ctx context.Context, params *DisassociateDelegationSignerFromDomainInput, optFns ...func(*Options)) (*DisassociateDelegationSignerFromDomainOutput, error) 
 EnableDomainAutoRenew(ctx context.Context, params *EnableDomainAutoRenewInput, optFns ...func(*Options)) (*EnableDomainAutoRenewOutput, error) 
 EnableDomainTransferLock(ctx context.Context, params *EnableDomainTransferLockInput, optFns ...func(*Options)) (*EnableDomainTransferLockOutput, error) 
 GetContactReachabilityStatus(ctx context.Context, params *GetContactReachabilityStatusInput, optFns ...func(*Options)) (*GetContactReachabilityStatusOutput, error) 
 GetDomainDetail(ctx context.Context, params *GetDomainDetailInput, optFns ...func(*Options)) (*GetDomainDetailOutput, error) 
 GetDomainSuggestions(ctx context.Context, params *GetDomainSuggestionsInput, optFns ...func(*Options)) (*GetDomainSuggestionsOutput, error) 
 GetOperationDetail(ctx context.Context, params *GetOperationDetailInput, optFns ...func(*Options)) (*GetOperationDetailOutput, error) 
 ListDomains(ctx context.Context, params *ListDomainsInput, optFns ...func(*Options)) (*ListDomainsOutput, error) 
 ListOperations(ctx context.Context, params *ListOperationsInput, optFns ...func(*Options)) (*ListOperationsOutput, error) 
 ListPrices(ctx context.Context, params *ListPricesInput, optFns ...func(*Options)) (*ListPricesOutput, error) 
 ListTagsForDomain(ctx context.Context, params *ListTagsForDomainInput, optFns ...func(*Options)) (*ListTagsForDomainOutput, error) 
 PushDomain(ctx context.Context, params *PushDomainInput, optFns ...func(*Options)) (*PushDomainOutput, error) 
 RegisterDomain(ctx context.Context, params *RegisterDomainInput, optFns ...func(*Options)) (*RegisterDomainOutput, error) 
 RejectDomainTransferFromAnotherAwsAccount(ctx context.Context, params *RejectDomainTransferFromAnotherAwsAccountInput, optFns ...func(*Options)) (*RejectDomainTransferFromAnotherAwsAccountOutput, error) 
 RenewDomain(ctx context.Context, params *RenewDomainInput, optFns ...func(*Options)) (*RenewDomainOutput, error) 
 ResendContactReachabilityEmail(ctx context.Context, params *ResendContactReachabilityEmailInput, optFns ...func(*Options)) (*ResendContactReachabilityEmailOutput, error) 
 ResendOperationAuthorization(ctx context.Context, params *ResendOperationAuthorizationInput, optFns ...func(*Options)) (*ResendOperationAuthorizationOutput, error) 
 RetrieveDomainAuthCode(ctx context.Context, params *RetrieveDomainAuthCodeInput, optFns ...func(*Options)) (*RetrieveDomainAuthCodeOutput, error) 
 TransferDomain(ctx context.Context, params *TransferDomainInput, optFns ...func(*Options)) (*TransferDomainOutput, error) 
 TransferDomainToAnotherAwsAccount(ctx context.Context, params *TransferDomainToAnotherAwsAccountInput, optFns ...func(*Options)) (*TransferDomainToAnotherAwsAccountOutput, error) 
 UpdateDomainContact(ctx context.Context, params *UpdateDomainContactInput, optFns ...func(*Options)) (*UpdateDomainContactOutput, error) 
 UpdateDomainContactPrivacy(ctx context.Context, params *UpdateDomainContactPrivacyInput, optFns ...func(*Options)) (*UpdateDomainContactPrivacyOutput, error) 
 UpdateDomainNameservers(ctx context.Context, params *UpdateDomainNameserversInput, optFns ...func(*Options)) (*UpdateDomainNameserversOutput, error) 
 UpdateTagsForDomain(ctx context.Context, params *UpdateTagsForDomainInput, optFns ...func(*Options)) (*UpdateTagsForDomainOutput, error) 
 ViewBilling(ctx context.Context, params *ViewBillingInput, optFns ...func(*Options)) (*ViewBillingOutput, error) 
}
