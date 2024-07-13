
package greengrass_iface

// This file was generated by scripts/gen_ifaces.sh do not edit it manually
import (
    "context"
    . "github.com/aws/aws-sdk-go-v2/service/greengrass"
)

// IClient defines the interface for greengrass
type IClient interface {
 Options() Options 
 AssociateRoleToGroup(ctx context.Context, params *AssociateRoleToGroupInput, optFns ...func(*Options)) (*AssociateRoleToGroupOutput, error) 
 AssociateServiceRoleToAccount(ctx context.Context, params *AssociateServiceRoleToAccountInput, optFns ...func(*Options)) (*AssociateServiceRoleToAccountOutput, error) 
 CreateConnectorDefinition(ctx context.Context, params *CreateConnectorDefinitionInput, optFns ...func(*Options)) (*CreateConnectorDefinitionOutput, error) 
 CreateConnectorDefinitionVersion(ctx context.Context, params *CreateConnectorDefinitionVersionInput, optFns ...func(*Options)) (*CreateConnectorDefinitionVersionOutput, error) 
 CreateCoreDefinition(ctx context.Context, params *CreateCoreDefinitionInput, optFns ...func(*Options)) (*CreateCoreDefinitionOutput, error) 
 CreateCoreDefinitionVersion(ctx context.Context, params *CreateCoreDefinitionVersionInput, optFns ...func(*Options)) (*CreateCoreDefinitionVersionOutput, error) 
 CreateDeployment(ctx context.Context, params *CreateDeploymentInput, optFns ...func(*Options)) (*CreateDeploymentOutput, error) 
 CreateDeviceDefinition(ctx context.Context, params *CreateDeviceDefinitionInput, optFns ...func(*Options)) (*CreateDeviceDefinitionOutput, error) 
 CreateDeviceDefinitionVersion(ctx context.Context, params *CreateDeviceDefinitionVersionInput, optFns ...func(*Options)) (*CreateDeviceDefinitionVersionOutput, error) 
 CreateFunctionDefinition(ctx context.Context, params *CreateFunctionDefinitionInput, optFns ...func(*Options)) (*CreateFunctionDefinitionOutput, error) 
 CreateFunctionDefinitionVersion(ctx context.Context, params *CreateFunctionDefinitionVersionInput, optFns ...func(*Options)) (*CreateFunctionDefinitionVersionOutput, error) 
 CreateGroup(ctx context.Context, params *CreateGroupInput, optFns ...func(*Options)) (*CreateGroupOutput, error) 
 CreateGroupCertificateAuthority(ctx context.Context, params *CreateGroupCertificateAuthorityInput, optFns ...func(*Options)) (*CreateGroupCertificateAuthorityOutput, error) 
 CreateGroupVersion(ctx context.Context, params *CreateGroupVersionInput, optFns ...func(*Options)) (*CreateGroupVersionOutput, error) 
 CreateLoggerDefinition(ctx context.Context, params *CreateLoggerDefinitionInput, optFns ...func(*Options)) (*CreateLoggerDefinitionOutput, error) 
 CreateLoggerDefinitionVersion(ctx context.Context, params *CreateLoggerDefinitionVersionInput, optFns ...func(*Options)) (*CreateLoggerDefinitionVersionOutput, error) 
 CreateResourceDefinition(ctx context.Context, params *CreateResourceDefinitionInput, optFns ...func(*Options)) (*CreateResourceDefinitionOutput, error) 
 CreateResourceDefinitionVersion(ctx context.Context, params *CreateResourceDefinitionVersionInput, optFns ...func(*Options)) (*CreateResourceDefinitionVersionOutput, error) 
 CreateSoftwareUpdateJob(ctx context.Context, params *CreateSoftwareUpdateJobInput, optFns ...func(*Options)) (*CreateSoftwareUpdateJobOutput, error) 
 CreateSubscriptionDefinition(ctx context.Context, params *CreateSubscriptionDefinitionInput, optFns ...func(*Options)) (*CreateSubscriptionDefinitionOutput, error) 
 CreateSubscriptionDefinitionVersion(ctx context.Context, params *CreateSubscriptionDefinitionVersionInput, optFns ...func(*Options)) (*CreateSubscriptionDefinitionVersionOutput, error) 
 DeleteConnectorDefinition(ctx context.Context, params *DeleteConnectorDefinitionInput, optFns ...func(*Options)) (*DeleteConnectorDefinitionOutput, error) 
 DeleteCoreDefinition(ctx context.Context, params *DeleteCoreDefinitionInput, optFns ...func(*Options)) (*DeleteCoreDefinitionOutput, error) 
 DeleteDeviceDefinition(ctx context.Context, params *DeleteDeviceDefinitionInput, optFns ...func(*Options)) (*DeleteDeviceDefinitionOutput, error) 
 DeleteFunctionDefinition(ctx context.Context, params *DeleteFunctionDefinitionInput, optFns ...func(*Options)) (*DeleteFunctionDefinitionOutput, error) 
 DeleteGroup(ctx context.Context, params *DeleteGroupInput, optFns ...func(*Options)) (*DeleteGroupOutput, error) 
 DeleteLoggerDefinition(ctx context.Context, params *DeleteLoggerDefinitionInput, optFns ...func(*Options)) (*DeleteLoggerDefinitionOutput, error) 
 DeleteResourceDefinition(ctx context.Context, params *DeleteResourceDefinitionInput, optFns ...func(*Options)) (*DeleteResourceDefinitionOutput, error) 
 DeleteSubscriptionDefinition(ctx context.Context, params *DeleteSubscriptionDefinitionInput, optFns ...func(*Options)) (*DeleteSubscriptionDefinitionOutput, error) 
 DisassociateRoleFromGroup(ctx context.Context, params *DisassociateRoleFromGroupInput, optFns ...func(*Options)) (*DisassociateRoleFromGroupOutput, error) 
 DisassociateServiceRoleFromAccount(ctx context.Context, params *DisassociateServiceRoleFromAccountInput, optFns ...func(*Options)) (*DisassociateServiceRoleFromAccountOutput, error) 
 GetAssociatedRole(ctx context.Context, params *GetAssociatedRoleInput, optFns ...func(*Options)) (*GetAssociatedRoleOutput, error) 
 GetBulkDeploymentStatus(ctx context.Context, params *GetBulkDeploymentStatusInput, optFns ...func(*Options)) (*GetBulkDeploymentStatusOutput, error) 
 GetConnectivityInfo(ctx context.Context, params *GetConnectivityInfoInput, optFns ...func(*Options)) (*GetConnectivityInfoOutput, error) 
 GetConnectorDefinition(ctx context.Context, params *GetConnectorDefinitionInput, optFns ...func(*Options)) (*GetConnectorDefinitionOutput, error) 
 GetConnectorDefinitionVersion(ctx context.Context, params *GetConnectorDefinitionVersionInput, optFns ...func(*Options)) (*GetConnectorDefinitionVersionOutput, error) 
 GetCoreDefinition(ctx context.Context, params *GetCoreDefinitionInput, optFns ...func(*Options)) (*GetCoreDefinitionOutput, error) 
 GetCoreDefinitionVersion(ctx context.Context, params *GetCoreDefinitionVersionInput, optFns ...func(*Options)) (*GetCoreDefinitionVersionOutput, error) 
 GetDeploymentStatus(ctx context.Context, params *GetDeploymentStatusInput, optFns ...func(*Options)) (*GetDeploymentStatusOutput, error) 
 GetDeviceDefinition(ctx context.Context, params *GetDeviceDefinitionInput, optFns ...func(*Options)) (*GetDeviceDefinitionOutput, error) 
 GetDeviceDefinitionVersion(ctx context.Context, params *GetDeviceDefinitionVersionInput, optFns ...func(*Options)) (*GetDeviceDefinitionVersionOutput, error) 
 GetFunctionDefinition(ctx context.Context, params *GetFunctionDefinitionInput, optFns ...func(*Options)) (*GetFunctionDefinitionOutput, error) 
 GetFunctionDefinitionVersion(ctx context.Context, params *GetFunctionDefinitionVersionInput, optFns ...func(*Options)) (*GetFunctionDefinitionVersionOutput, error) 
 GetGroup(ctx context.Context, params *GetGroupInput, optFns ...func(*Options)) (*GetGroupOutput, error) 
 GetGroupCertificateAuthority(ctx context.Context, params *GetGroupCertificateAuthorityInput, optFns ...func(*Options)) (*GetGroupCertificateAuthorityOutput, error) 
 GetGroupCertificateConfiguration(ctx context.Context, params *GetGroupCertificateConfigurationInput, optFns ...func(*Options)) (*GetGroupCertificateConfigurationOutput, error) 
 GetGroupVersion(ctx context.Context, params *GetGroupVersionInput, optFns ...func(*Options)) (*GetGroupVersionOutput, error) 
 GetLoggerDefinition(ctx context.Context, params *GetLoggerDefinitionInput, optFns ...func(*Options)) (*GetLoggerDefinitionOutput, error) 
 GetLoggerDefinitionVersion(ctx context.Context, params *GetLoggerDefinitionVersionInput, optFns ...func(*Options)) (*GetLoggerDefinitionVersionOutput, error) 
 GetResourceDefinition(ctx context.Context, params *GetResourceDefinitionInput, optFns ...func(*Options)) (*GetResourceDefinitionOutput, error) 
 GetResourceDefinitionVersion(ctx context.Context, params *GetResourceDefinitionVersionInput, optFns ...func(*Options)) (*GetResourceDefinitionVersionOutput, error) 
 GetServiceRoleForAccount(ctx context.Context, params *GetServiceRoleForAccountInput, optFns ...func(*Options)) (*GetServiceRoleForAccountOutput, error) 
 GetSubscriptionDefinition(ctx context.Context, params *GetSubscriptionDefinitionInput, optFns ...func(*Options)) (*GetSubscriptionDefinitionOutput, error) 
 GetSubscriptionDefinitionVersion(ctx context.Context, params *GetSubscriptionDefinitionVersionInput, optFns ...func(*Options)) (*GetSubscriptionDefinitionVersionOutput, error) 
 GetThingRuntimeConfiguration(ctx context.Context, params *GetThingRuntimeConfigurationInput, optFns ...func(*Options)) (*GetThingRuntimeConfigurationOutput, error) 
 ListBulkDeploymentDetailedReports(ctx context.Context, params *ListBulkDeploymentDetailedReportsInput, optFns ...func(*Options)) (*ListBulkDeploymentDetailedReportsOutput, error) 
 ListBulkDeployments(ctx context.Context, params *ListBulkDeploymentsInput, optFns ...func(*Options)) (*ListBulkDeploymentsOutput, error) 
 ListConnectorDefinitionVersions(ctx context.Context, params *ListConnectorDefinitionVersionsInput, optFns ...func(*Options)) (*ListConnectorDefinitionVersionsOutput, error) 
 ListConnectorDefinitions(ctx context.Context, params *ListConnectorDefinitionsInput, optFns ...func(*Options)) (*ListConnectorDefinitionsOutput, error) 
 ListCoreDefinitionVersions(ctx context.Context, params *ListCoreDefinitionVersionsInput, optFns ...func(*Options)) (*ListCoreDefinitionVersionsOutput, error) 
 ListCoreDefinitions(ctx context.Context, params *ListCoreDefinitionsInput, optFns ...func(*Options)) (*ListCoreDefinitionsOutput, error) 
 ListDeployments(ctx context.Context, params *ListDeploymentsInput, optFns ...func(*Options)) (*ListDeploymentsOutput, error) 
 ListDeviceDefinitionVersions(ctx context.Context, params *ListDeviceDefinitionVersionsInput, optFns ...func(*Options)) (*ListDeviceDefinitionVersionsOutput, error) 
 ListDeviceDefinitions(ctx context.Context, params *ListDeviceDefinitionsInput, optFns ...func(*Options)) (*ListDeviceDefinitionsOutput, error) 
 ListFunctionDefinitionVersions(ctx context.Context, params *ListFunctionDefinitionVersionsInput, optFns ...func(*Options)) (*ListFunctionDefinitionVersionsOutput, error) 
 ListFunctionDefinitions(ctx context.Context, params *ListFunctionDefinitionsInput, optFns ...func(*Options)) (*ListFunctionDefinitionsOutput, error) 
 ListGroupCertificateAuthorities(ctx context.Context, params *ListGroupCertificateAuthoritiesInput, optFns ...func(*Options)) (*ListGroupCertificateAuthoritiesOutput, error) 
 ListGroupVersions(ctx context.Context, params *ListGroupVersionsInput, optFns ...func(*Options)) (*ListGroupVersionsOutput, error) 
 ListGroups(ctx context.Context, params *ListGroupsInput, optFns ...func(*Options)) (*ListGroupsOutput, error) 
 ListLoggerDefinitionVersions(ctx context.Context, params *ListLoggerDefinitionVersionsInput, optFns ...func(*Options)) (*ListLoggerDefinitionVersionsOutput, error) 
 ListLoggerDefinitions(ctx context.Context, params *ListLoggerDefinitionsInput, optFns ...func(*Options)) (*ListLoggerDefinitionsOutput, error) 
 ListResourceDefinitionVersions(ctx context.Context, params *ListResourceDefinitionVersionsInput, optFns ...func(*Options)) (*ListResourceDefinitionVersionsOutput, error) 
 ListResourceDefinitions(ctx context.Context, params *ListResourceDefinitionsInput, optFns ...func(*Options)) (*ListResourceDefinitionsOutput, error) 
 ListSubscriptionDefinitionVersions(ctx context.Context, params *ListSubscriptionDefinitionVersionsInput, optFns ...func(*Options)) (*ListSubscriptionDefinitionVersionsOutput, error) 
 ListSubscriptionDefinitions(ctx context.Context, params *ListSubscriptionDefinitionsInput, optFns ...func(*Options)) (*ListSubscriptionDefinitionsOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 ResetDeployments(ctx context.Context, params *ResetDeploymentsInput, optFns ...func(*Options)) (*ResetDeploymentsOutput, error) 
 StartBulkDeployment(ctx context.Context, params *StartBulkDeploymentInput, optFns ...func(*Options)) (*StartBulkDeploymentOutput, error) 
 StopBulkDeployment(ctx context.Context, params *StopBulkDeploymentInput, optFns ...func(*Options)) (*StopBulkDeploymentOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateConnectivityInfo(ctx context.Context, params *UpdateConnectivityInfoInput, optFns ...func(*Options)) (*UpdateConnectivityInfoOutput, error) 
 UpdateConnectorDefinition(ctx context.Context, params *UpdateConnectorDefinitionInput, optFns ...func(*Options)) (*UpdateConnectorDefinitionOutput, error) 
 UpdateCoreDefinition(ctx context.Context, params *UpdateCoreDefinitionInput, optFns ...func(*Options)) (*UpdateCoreDefinitionOutput, error) 
 UpdateDeviceDefinition(ctx context.Context, params *UpdateDeviceDefinitionInput, optFns ...func(*Options)) (*UpdateDeviceDefinitionOutput, error) 
 UpdateFunctionDefinition(ctx context.Context, params *UpdateFunctionDefinitionInput, optFns ...func(*Options)) (*UpdateFunctionDefinitionOutput, error) 
 UpdateGroup(ctx context.Context, params *UpdateGroupInput, optFns ...func(*Options)) (*UpdateGroupOutput, error) 
 UpdateGroupCertificateConfiguration(ctx context.Context, params *UpdateGroupCertificateConfigurationInput, optFns ...func(*Options)) (*UpdateGroupCertificateConfigurationOutput, error) 
 UpdateLoggerDefinition(ctx context.Context, params *UpdateLoggerDefinitionInput, optFns ...func(*Options)) (*UpdateLoggerDefinitionOutput, error) 
 UpdateResourceDefinition(ctx context.Context, params *UpdateResourceDefinitionInput, optFns ...func(*Options)) (*UpdateResourceDefinitionOutput, error) 
 UpdateSubscriptionDefinition(ctx context.Context, params *UpdateSubscriptionDefinitionInput, optFns ...func(*Options)) (*UpdateSubscriptionDefinitionOutput, error) 
 UpdateThingRuntimeConfiguration(ctx context.Context, params *UpdateThingRuntimeConfigurationInput, optFns ...func(*Options)) (*UpdateThingRuntimeConfigurationOutput, error) 
}