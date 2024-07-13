
package licensemanagerlinuxsubscriptions_iface

// This file was generated by scripts/gen_ifaces.sh do not edit it manually
import (
    "context"
    . "github.com/aws/aws-sdk-go-v2/service/licensemanagerlinuxsubscriptions"
)

// IClient defines the interface for licensemanagerlinuxsubscriptions
type IClient interface {
 Options() Options 
 DeregisterSubscriptionProvider(ctx context.Context, params *DeregisterSubscriptionProviderInput, optFns ...func(*Options)) (*DeregisterSubscriptionProviderOutput, error) 
 GetRegisteredSubscriptionProvider(ctx context.Context, params *GetRegisteredSubscriptionProviderInput, optFns ...func(*Options)) (*GetRegisteredSubscriptionProviderOutput, error) 
 GetServiceSettings(ctx context.Context, params *GetServiceSettingsInput, optFns ...func(*Options)) (*GetServiceSettingsOutput, error) 
 ListLinuxSubscriptionInstances(ctx context.Context, params *ListLinuxSubscriptionInstancesInput, optFns ...func(*Options)) (*ListLinuxSubscriptionInstancesOutput, error) 
 ListLinuxSubscriptions(ctx context.Context, params *ListLinuxSubscriptionsInput, optFns ...func(*Options)) (*ListLinuxSubscriptionsOutput, error) 
 ListRegisteredSubscriptionProviders(ctx context.Context, params *ListRegisteredSubscriptionProvidersInput, optFns ...func(*Options)) (*ListRegisteredSubscriptionProvidersOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 RegisterSubscriptionProvider(ctx context.Context, params *RegisterSubscriptionProviderInput, optFns ...func(*Options)) (*RegisterSubscriptionProviderOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateServiceSettings(ctx context.Context, params *UpdateServiceSettingsInput, optFns ...func(*Options)) (*UpdateServiceSettingsOutput, error) 
}