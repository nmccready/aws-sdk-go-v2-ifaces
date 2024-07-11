
package licensemanagerusersubscriptions

import (
    "github.com/aws/aws-sdk-go-v2/service/licensemanagerusersubscriptions"
)

// ILicensemanagerusersubscriptions defines the interface for licensemanagerusersubscriptions
type ILicensemanagerusersubscriptions interface {
 Options() Options 
 AssociateUser(ctx context.Context, params *AssociateUserInput, optFns ...func(*Options)) (*AssociateUserOutput, error) 
 DeregisterIdentityProvider(ctx context.Context, params *DeregisterIdentityProviderInput, optFns ...func(*Options)) (*DeregisterIdentityProviderOutput, error) 
 DisassociateUser(ctx context.Context, params *DisassociateUserInput, optFns ...func(*Options)) (*DisassociateUserOutput, error) 
 ListIdentityProviders(ctx context.Context, params *ListIdentityProvidersInput, optFns ...func(*Options)) (*ListIdentityProvidersOutput, error) 
 ListInstances(ctx context.Context, params *ListInstancesInput, optFns ...func(*Options)) (*ListInstancesOutput, error) 
 ListProductSubscriptions(ctx context.Context, params *ListProductSubscriptionsInput, optFns ...func(*Options)) (*ListProductSubscriptionsOutput, error) 
 ListUserAssociations(ctx context.Context, params *ListUserAssociationsInput, optFns ...func(*Options)) (*ListUserAssociationsOutput, error) 
 RegisterIdentityProvider(ctx context.Context, params *RegisterIdentityProviderInput, optFns ...func(*Options)) (*RegisterIdentityProviderOutput, error) 
 StartProductSubscription(ctx context.Context, params *StartProductSubscriptionInput, optFns ...func(*Options)) (*StartProductSubscriptionOutput, error) 
 StopProductSubscription(ctx context.Context, params *StopProductSubscriptionInput, optFns ...func(*Options)) (*StopProductSubscriptionOutput, error) 
 UpdateIdentityProviderSettings(ctx context.Context, params *UpdateIdentityProviderSettingsInput, optFns ...func(*Options)) (*UpdateIdentityProviderSettingsOutput, error) 
}
