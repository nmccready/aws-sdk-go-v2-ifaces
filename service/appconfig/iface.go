
package appconfig

import (
    "github.com/aws/aws-sdk-go-v2/service/appconfig"
)

// IClient defines the interface for appconfig
type IClient interface {
 Options() Options 
 CreateApplication(ctx context.Context, params *CreateApplicationInput, optFns ...func(*Options)) (*CreateApplicationOutput, error) 
 CreateConfigurationProfile(ctx context.Context, params *CreateConfigurationProfileInput, optFns ...func(*Options)) (*CreateConfigurationProfileOutput, error) 
 CreateDeploymentStrategy(ctx context.Context, params *CreateDeploymentStrategyInput, optFns ...func(*Options)) (*CreateDeploymentStrategyOutput, error) 
 CreateEnvironment(ctx context.Context, params *CreateEnvironmentInput, optFns ...func(*Options)) (*CreateEnvironmentOutput, error) 
 CreateExtension(ctx context.Context, params *CreateExtensionInput, optFns ...func(*Options)) (*CreateExtensionOutput, error) 
 CreateExtensionAssociation(ctx context.Context, params *CreateExtensionAssociationInput, optFns ...func(*Options)) (*CreateExtensionAssociationOutput, error) 
 CreateHostedConfigurationVersion(ctx context.Context, params *CreateHostedConfigurationVersionInput, optFns ...func(*Options)) (*CreateHostedConfigurationVersionOutput, error) 
 DeleteApplication(ctx context.Context, params *DeleteApplicationInput, optFns ...func(*Options)) (*DeleteApplicationOutput, error) 
 DeleteConfigurationProfile(ctx context.Context, params *DeleteConfigurationProfileInput, optFns ...func(*Options)) (*DeleteConfigurationProfileOutput, error) 
 DeleteDeploymentStrategy(ctx context.Context, params *DeleteDeploymentStrategyInput, optFns ...func(*Options)) (*DeleteDeploymentStrategyOutput, error) 
 DeleteEnvironment(ctx context.Context, params *DeleteEnvironmentInput, optFns ...func(*Options)) (*DeleteEnvironmentOutput, error) 
 DeleteExtension(ctx context.Context, params *DeleteExtensionInput, optFns ...func(*Options)) (*DeleteExtensionOutput, error) 
 DeleteExtensionAssociation(ctx context.Context, params *DeleteExtensionAssociationInput, optFns ...func(*Options)) (*DeleteExtensionAssociationOutput, error) 
 DeleteHostedConfigurationVersion(ctx context.Context, params *DeleteHostedConfigurationVersionInput, optFns ...func(*Options)) (*DeleteHostedConfigurationVersionOutput, error) 
 GetApplication(ctx context.Context, params *GetApplicationInput, optFns ...func(*Options)) (*GetApplicationOutput, error) 
 GetConfiguration(ctx context.Context, params *GetConfigurationInput, optFns ...func(*Options)) (*GetConfigurationOutput, error) 
 GetConfigurationProfile(ctx context.Context, params *GetConfigurationProfileInput, optFns ...func(*Options)) (*GetConfigurationProfileOutput, error) 
 GetDeployment(ctx context.Context, params *GetDeploymentInput, optFns ...func(*Options)) (*GetDeploymentOutput, error) 
 GetDeploymentStrategy(ctx context.Context, params *GetDeploymentStrategyInput, optFns ...func(*Options)) (*GetDeploymentStrategyOutput, error) 
 GetEnvironment(ctx context.Context, params *GetEnvironmentInput, optFns ...func(*Options)) (*GetEnvironmentOutput, error) 
 GetExtension(ctx context.Context, params *GetExtensionInput, optFns ...func(*Options)) (*GetExtensionOutput, error) 
 GetExtensionAssociation(ctx context.Context, params *GetExtensionAssociationInput, optFns ...func(*Options)) (*GetExtensionAssociationOutput, error) 
 GetHostedConfigurationVersion(ctx context.Context, params *GetHostedConfigurationVersionInput, optFns ...func(*Options)) (*GetHostedConfigurationVersionOutput, error) 
 ListApplications(ctx context.Context, params *ListApplicationsInput, optFns ...func(*Options)) (*ListApplicationsOutput, error) 
 ListConfigurationProfiles(ctx context.Context, params *ListConfigurationProfilesInput, optFns ...func(*Options)) (*ListConfigurationProfilesOutput, error) 
 ListDeploymentStrategies(ctx context.Context, params *ListDeploymentStrategiesInput, optFns ...func(*Options)) (*ListDeploymentStrategiesOutput, error) 
 ListDeployments(ctx context.Context, params *ListDeploymentsInput, optFns ...func(*Options)) (*ListDeploymentsOutput, error) 
 ListEnvironments(ctx context.Context, params *ListEnvironmentsInput, optFns ...func(*Options)) (*ListEnvironmentsOutput, error) 
 ListExtensionAssociations(ctx context.Context, params *ListExtensionAssociationsInput, optFns ...func(*Options)) (*ListExtensionAssociationsOutput, error) 
 ListExtensions(ctx context.Context, params *ListExtensionsInput, optFns ...func(*Options)) (*ListExtensionsOutput, error) 
 ListHostedConfigurationVersions(ctx context.Context, params *ListHostedConfigurationVersionsInput, optFns ...func(*Options)) (*ListHostedConfigurationVersionsOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 StartDeployment(ctx context.Context, params *StartDeploymentInput, optFns ...func(*Options)) (*StartDeploymentOutput, error) 
 StopDeployment(ctx context.Context, params *StopDeploymentInput, optFns ...func(*Options)) (*StopDeploymentOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateApplication(ctx context.Context, params *UpdateApplicationInput, optFns ...func(*Options)) (*UpdateApplicationOutput, error) 
 UpdateConfigurationProfile(ctx context.Context, params *UpdateConfigurationProfileInput, optFns ...func(*Options)) (*UpdateConfigurationProfileOutput, error) 
 UpdateDeploymentStrategy(ctx context.Context, params *UpdateDeploymentStrategyInput, optFns ...func(*Options)) (*UpdateDeploymentStrategyOutput, error) 
 UpdateEnvironment(ctx context.Context, params *UpdateEnvironmentInput, optFns ...func(*Options)) (*UpdateEnvironmentOutput, error) 
 UpdateExtension(ctx context.Context, params *UpdateExtensionInput, optFns ...func(*Options)) (*UpdateExtensionOutput, error) 
 UpdateExtensionAssociation(ctx context.Context, params *UpdateExtensionAssociationInput, optFns ...func(*Options)) (*UpdateExtensionAssociationOutput, error) 
 ValidateConfiguration(ctx context.Context, params *ValidateConfigurationInput, optFns ...func(*Options)) (*ValidateConfigurationOutput, error) 
}
