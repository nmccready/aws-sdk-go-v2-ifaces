
package servicediscovery

import (
    "github.com/aws/aws-sdk-go-v2/service/servicediscovery"
)

// IClient defines the interface for servicediscovery
type IClient interface {
 Options() Options 
 CreateHttpNamespace(ctx context.Context, params *CreateHttpNamespaceInput, optFns ...func(*Options)) (*CreateHttpNamespaceOutput, error) 
 CreatePrivateDnsNamespace(ctx context.Context, params *CreatePrivateDnsNamespaceInput, optFns ...func(*Options)) (*CreatePrivateDnsNamespaceOutput, error) 
 CreatePublicDnsNamespace(ctx context.Context, params *CreatePublicDnsNamespaceInput, optFns ...func(*Options)) (*CreatePublicDnsNamespaceOutput, error) 
 CreateService(ctx context.Context, params *CreateServiceInput, optFns ...func(*Options)) (*CreateServiceOutput, error) 
 DeleteNamespace(ctx context.Context, params *DeleteNamespaceInput, optFns ...func(*Options)) (*DeleteNamespaceOutput, error) 
 DeleteService(ctx context.Context, params *DeleteServiceInput, optFns ...func(*Options)) (*DeleteServiceOutput, error) 
 DeregisterInstance(ctx context.Context, params *DeregisterInstanceInput, optFns ...func(*Options)) (*DeregisterInstanceOutput, error) 
 DiscoverInstances(ctx context.Context, params *DiscoverInstancesInput, optFns ...func(*Options)) (*DiscoverInstancesOutput, error) 
 DiscoverInstancesRevision(ctx context.Context, params *DiscoverInstancesRevisionInput, optFns ...func(*Options)) (*DiscoverInstancesRevisionOutput, error) 
 GetInstance(ctx context.Context, params *GetInstanceInput, optFns ...func(*Options)) (*GetInstanceOutput, error) 
 GetInstancesHealthStatus(ctx context.Context, params *GetInstancesHealthStatusInput, optFns ...func(*Options)) (*GetInstancesHealthStatusOutput, error) 
 GetNamespace(ctx context.Context, params *GetNamespaceInput, optFns ...func(*Options)) (*GetNamespaceOutput, error) 
 GetOperation(ctx context.Context, params *GetOperationInput, optFns ...func(*Options)) (*GetOperationOutput, error) 
 GetService(ctx context.Context, params *GetServiceInput, optFns ...func(*Options)) (*GetServiceOutput, error) 
 ListInstances(ctx context.Context, params *ListInstancesInput, optFns ...func(*Options)) (*ListInstancesOutput, error) 
 ListNamespaces(ctx context.Context, params *ListNamespacesInput, optFns ...func(*Options)) (*ListNamespacesOutput, error) 
 ListOperations(ctx context.Context, params *ListOperationsInput, optFns ...func(*Options)) (*ListOperationsOutput, error) 
 ListServices(ctx context.Context, params *ListServicesInput, optFns ...func(*Options)) (*ListServicesOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 RegisterInstance(ctx context.Context, params *RegisterInstanceInput, optFns ...func(*Options)) (*RegisterInstanceOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateHttpNamespace(ctx context.Context, params *UpdateHttpNamespaceInput, optFns ...func(*Options)) (*UpdateHttpNamespaceOutput, error) 
 UpdateInstanceCustomHealthStatus(ctx context.Context, params *UpdateInstanceCustomHealthStatusInput, optFns ...func(*Options)) (*UpdateInstanceCustomHealthStatusOutput, error) 
 UpdatePrivateDnsNamespace(ctx context.Context, params *UpdatePrivateDnsNamespaceInput, optFns ...func(*Options)) (*UpdatePrivateDnsNamespaceOutput, error) 
 UpdatePublicDnsNamespace(ctx context.Context, params *UpdatePublicDnsNamespaceInput, optFns ...func(*Options)) (*UpdatePublicDnsNamespaceOutput, error) 
 UpdateService(ctx context.Context, params *UpdateServiceInput, optFns ...func(*Options)) (*UpdateServiceOutput, error) 
}
