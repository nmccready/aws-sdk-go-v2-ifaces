
package account

import (
    "github.com/aws/aws-sdk-go-v2/service/account"
)

// IAccount defines the interface for account
type IAccount interface {
 Options() Options 
 AcceptPrimaryEmailUpdate(ctx context.Context, params *AcceptPrimaryEmailUpdateInput, optFns ...func(*Options)) (*AcceptPrimaryEmailUpdateOutput, error) 
 DeleteAlternateContact(ctx context.Context, params *DeleteAlternateContactInput, optFns ...func(*Options)) (*DeleteAlternateContactOutput, error) 
 DisableRegion(ctx context.Context, params *DisableRegionInput, optFns ...func(*Options)) (*DisableRegionOutput, error) 
 EnableRegion(ctx context.Context, params *EnableRegionInput, optFns ...func(*Options)) (*EnableRegionOutput, error) 
 GetAlternateContact(ctx context.Context, params *GetAlternateContactInput, optFns ...func(*Options)) (*GetAlternateContactOutput, error) 
 GetContactInformation(ctx context.Context, params *GetContactInformationInput, optFns ...func(*Options)) (*GetContactInformationOutput, error) 
 GetPrimaryEmail(ctx context.Context, params *GetPrimaryEmailInput, optFns ...func(*Options)) (*GetPrimaryEmailOutput, error) 
 GetRegionOptStatus(ctx context.Context, params *GetRegionOptStatusInput, optFns ...func(*Options)) (*GetRegionOptStatusOutput, error) 
 ListRegions(ctx context.Context, params *ListRegionsInput, optFns ...func(*Options)) (*ListRegionsOutput, error) 
 PutAlternateContact(ctx context.Context, params *PutAlternateContactInput, optFns ...func(*Options)) (*PutAlternateContactOutput, error) 
 PutContactInformation(ctx context.Context, params *PutContactInformationInput, optFns ...func(*Options)) (*PutContactInformationOutput, error) 
 StartPrimaryEmailUpdate(ctx context.Context, params *StartPrimaryEmailUpdateInput, optFns ...func(*Options)) (*StartPrimaryEmailUpdateOutput, error) 
}
