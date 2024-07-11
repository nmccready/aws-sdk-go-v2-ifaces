
package apprunner

import (
    "github.com/aws/aws-sdk-go-v2/service/apprunner"
)

// IClient defines the interface for apprunner
type IClient interface {
 Options() Options 
 AssociateCustomDomain(ctx context.Context, params *AssociateCustomDomainInput, optFns ...func(*Options)) (*AssociateCustomDomainOutput, error) 
 CreateAutoScalingConfiguration(ctx context.Context, params *CreateAutoScalingConfigurationInput, optFns ...func(*Options)) (*CreateAutoScalingConfigurationOutput, error) 
 CreateConnection(ctx context.Context, params *CreateConnectionInput, optFns ...func(*Options)) (*CreateConnectionOutput, error) 
 CreateObservabilityConfiguration(ctx context.Context, params *CreateObservabilityConfigurationInput, optFns ...func(*Options)) (*CreateObservabilityConfigurationOutput, error) 
 CreateService(ctx context.Context, params *CreateServiceInput, optFns ...func(*Options)) (*CreateServiceOutput, error) 
 CreateVpcConnector(ctx context.Context, params *CreateVpcConnectorInput, optFns ...func(*Options)) (*CreateVpcConnectorOutput, error) 
 CreateVpcIngressConnection(ctx context.Context, params *CreateVpcIngressConnectionInput, optFns ...func(*Options)) (*CreateVpcIngressConnectionOutput, error) 
 DeleteAutoScalingConfiguration(ctx context.Context, params *DeleteAutoScalingConfigurationInput, optFns ...func(*Options)) (*DeleteAutoScalingConfigurationOutput, error) 
 DeleteConnection(ctx context.Context, params *DeleteConnectionInput, optFns ...func(*Options)) (*DeleteConnectionOutput, error) 
 DeleteObservabilityConfiguration(ctx context.Context, params *DeleteObservabilityConfigurationInput, optFns ...func(*Options)) (*DeleteObservabilityConfigurationOutput, error) 
 DeleteService(ctx context.Context, params *DeleteServiceInput, optFns ...func(*Options)) (*DeleteServiceOutput, error) 
 DeleteVpcConnector(ctx context.Context, params *DeleteVpcConnectorInput, optFns ...func(*Options)) (*DeleteVpcConnectorOutput, error) 
 DeleteVpcIngressConnection(ctx context.Context, params *DeleteVpcIngressConnectionInput, optFns ...func(*Options)) (*DeleteVpcIngressConnectionOutput, error) 
 DescribeAutoScalingConfiguration(ctx context.Context, params *DescribeAutoScalingConfigurationInput, optFns ...func(*Options)) (*DescribeAutoScalingConfigurationOutput, error) 
 DescribeCustomDomains(ctx context.Context, params *DescribeCustomDomainsInput, optFns ...func(*Options)) (*DescribeCustomDomainsOutput, error) 
 DescribeObservabilityConfiguration(ctx context.Context, params *DescribeObservabilityConfigurationInput, optFns ...func(*Options)) (*DescribeObservabilityConfigurationOutput, error) 
 DescribeService(ctx context.Context, params *DescribeServiceInput, optFns ...func(*Options)) (*DescribeServiceOutput, error) 
 DescribeVpcConnector(ctx context.Context, params *DescribeVpcConnectorInput, optFns ...func(*Options)) (*DescribeVpcConnectorOutput, error) 
 DescribeVpcIngressConnection(ctx context.Context, params *DescribeVpcIngressConnectionInput, optFns ...func(*Options)) (*DescribeVpcIngressConnectionOutput, error) 
 DisassociateCustomDomain(ctx context.Context, params *DisassociateCustomDomainInput, optFns ...func(*Options)) (*DisassociateCustomDomainOutput, error) 
 ListAutoScalingConfigurations(ctx context.Context, params *ListAutoScalingConfigurationsInput, optFns ...func(*Options)) (*ListAutoScalingConfigurationsOutput, error) 
 ListConnections(ctx context.Context, params *ListConnectionsInput, optFns ...func(*Options)) (*ListConnectionsOutput, error) 
 ListObservabilityConfigurations(ctx context.Context, params *ListObservabilityConfigurationsInput, optFns ...func(*Options)) (*ListObservabilityConfigurationsOutput, error) 
 ListOperations(ctx context.Context, params *ListOperationsInput, optFns ...func(*Options)) (*ListOperationsOutput, error) 
 ListServices(ctx context.Context, params *ListServicesInput, optFns ...func(*Options)) (*ListServicesOutput, error) 
 ListServicesForAutoScalingConfiguration(ctx context.Context, params *ListServicesForAutoScalingConfigurationInput, optFns ...func(*Options)) (*ListServicesForAutoScalingConfigurationOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 ListVpcConnectors(ctx context.Context, params *ListVpcConnectorsInput, optFns ...func(*Options)) (*ListVpcConnectorsOutput, error) 
 ListVpcIngressConnections(ctx context.Context, params *ListVpcIngressConnectionsInput, optFns ...func(*Options)) (*ListVpcIngressConnectionsOutput, error) 
 PauseService(ctx context.Context, params *PauseServiceInput, optFns ...func(*Options)) (*PauseServiceOutput, error) 
 ResumeService(ctx context.Context, params *ResumeServiceInput, optFns ...func(*Options)) (*ResumeServiceOutput, error) 
 StartDeployment(ctx context.Context, params *StartDeploymentInput, optFns ...func(*Options)) (*StartDeploymentOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateDefaultAutoScalingConfiguration(ctx context.Context, params *UpdateDefaultAutoScalingConfigurationInput, optFns ...func(*Options)) (*UpdateDefaultAutoScalingConfigurationOutput, error) 
 UpdateService(ctx context.Context, params *UpdateServiceInput, optFns ...func(*Options)) (*UpdateServiceOutput, error) 
 UpdateVpcIngressConnection(ctx context.Context, params *UpdateVpcIngressConnectionInput, optFns ...func(*Options)) (*UpdateVpcIngressConnectionOutput, error) 
}
