
package elasticbeanstalk

import (
    "github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk"
)

// IElasticbeanstalk defines the interface for elasticbeanstalk
type IElasticbeanstalk interface {
 Options() Options 
 AbortEnvironmentUpdate(ctx context.Context, params *AbortEnvironmentUpdateInput, optFns ...func(*Options)) (*AbortEnvironmentUpdateOutput, error) 
 ApplyEnvironmentManagedAction(ctx context.Context, params *ApplyEnvironmentManagedActionInput, optFns ...func(*Options)) (*ApplyEnvironmentManagedActionOutput, error) 
 AssociateEnvironmentOperationsRole(ctx context.Context, params *AssociateEnvironmentOperationsRoleInput, optFns ...func(*Options)) (*AssociateEnvironmentOperationsRoleOutput, error) 
 CheckDNSAvailability(ctx context.Context, params *CheckDNSAvailabilityInput, optFns ...func(*Options)) (*CheckDNSAvailabilityOutput, error) 
 ComposeEnvironments(ctx context.Context, params *ComposeEnvironmentsInput, optFns ...func(*Options)) (*ComposeEnvironmentsOutput, error) 
 CreateApplication(ctx context.Context, params *CreateApplicationInput, optFns ...func(*Options)) (*CreateApplicationOutput, error) 
 CreateApplicationVersion(ctx context.Context, params *CreateApplicationVersionInput, optFns ...func(*Options)) (*CreateApplicationVersionOutput, error) 
 CreateConfigurationTemplate(ctx context.Context, params *CreateConfigurationTemplateInput, optFns ...func(*Options)) (*CreateConfigurationTemplateOutput, error) 
 CreateEnvironment(ctx context.Context, params *CreateEnvironmentInput, optFns ...func(*Options)) (*CreateEnvironmentOutput, error) 
 CreatePlatformVersion(ctx context.Context, params *CreatePlatformVersionInput, optFns ...func(*Options)) (*CreatePlatformVersionOutput, error) 
 CreateStorageLocation(ctx context.Context, params *CreateStorageLocationInput, optFns ...func(*Options)) (*CreateStorageLocationOutput, error) 
 DeleteApplication(ctx context.Context, params *DeleteApplicationInput, optFns ...func(*Options)) (*DeleteApplicationOutput, error) 
 DeleteApplicationVersion(ctx context.Context, params *DeleteApplicationVersionInput, optFns ...func(*Options)) (*DeleteApplicationVersionOutput, error) 
 DeleteConfigurationTemplate(ctx context.Context, params *DeleteConfigurationTemplateInput, optFns ...func(*Options)) (*DeleteConfigurationTemplateOutput, error) 
 DeleteEnvironmentConfiguration(ctx context.Context, params *DeleteEnvironmentConfigurationInput, optFns ...func(*Options)) (*DeleteEnvironmentConfigurationOutput, error) 
 DeletePlatformVersion(ctx context.Context, params *DeletePlatformVersionInput, optFns ...func(*Options)) (*DeletePlatformVersionOutput, error) 
 DescribeAccountAttributes(ctx context.Context, params *DescribeAccountAttributesInput, optFns ...func(*Options)) (*DescribeAccountAttributesOutput, error) 
 DescribeApplicationVersions(ctx context.Context, params *DescribeApplicationVersionsInput, optFns ...func(*Options)) (*DescribeApplicationVersionsOutput, error) 
 DescribeApplications(ctx context.Context, params *DescribeApplicationsInput, optFns ...func(*Options)) (*DescribeApplicationsOutput, error) 
 DescribeConfigurationOptions(ctx context.Context, params *DescribeConfigurationOptionsInput, optFns ...func(*Options)) (*DescribeConfigurationOptionsOutput, error) 
 DescribeConfigurationSettings(ctx context.Context, params *DescribeConfigurationSettingsInput, optFns ...func(*Options)) (*DescribeConfigurationSettingsOutput, error) 
 DescribeEnvironmentHealth(ctx context.Context, params *DescribeEnvironmentHealthInput, optFns ...func(*Options)) (*DescribeEnvironmentHealthOutput, error) 
 DescribeEnvironmentManagedActionHistory(ctx context.Context, params *DescribeEnvironmentManagedActionHistoryInput, optFns ...func(*Options)) (*DescribeEnvironmentManagedActionHistoryOutput, error) 
 DescribeEnvironmentManagedActions(ctx context.Context, params *DescribeEnvironmentManagedActionsInput, optFns ...func(*Options)) (*DescribeEnvironmentManagedActionsOutput, error) 
 DescribeEnvironmentResources(ctx context.Context, params *DescribeEnvironmentResourcesInput, optFns ...func(*Options)) (*DescribeEnvironmentResourcesOutput, error) 
 DescribeEnvironments(ctx context.Context, params *DescribeEnvironmentsInput, optFns ...func(*Options)) (*DescribeEnvironmentsOutput, error) 
 DescribeEvents(ctx context.Context, params *DescribeEventsInput, optFns ...func(*Options)) (*DescribeEventsOutput, error) 
 DescribeInstancesHealth(ctx context.Context, params *DescribeInstancesHealthInput, optFns ...func(*Options)) (*DescribeInstancesHealthOutput, error) 
 DescribePlatformVersion(ctx context.Context, params *DescribePlatformVersionInput, optFns ...func(*Options)) (*DescribePlatformVersionOutput, error) 
 DisassociateEnvironmentOperationsRole(ctx context.Context, params *DisassociateEnvironmentOperationsRoleInput, optFns ...func(*Options)) (*DisassociateEnvironmentOperationsRoleOutput, error) 
 ListAvailableSolutionStacks(ctx context.Context, params *ListAvailableSolutionStacksInput, optFns ...func(*Options)) (*ListAvailableSolutionStacksOutput, error) 
 ListPlatformBranches(ctx context.Context, params *ListPlatformBranchesInput, optFns ...func(*Options)) (*ListPlatformBranchesOutput, error) 
 ListPlatformVersions(ctx context.Context, params *ListPlatformVersionsInput, optFns ...func(*Options)) (*ListPlatformVersionsOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 RebuildEnvironment(ctx context.Context, params *RebuildEnvironmentInput, optFns ...func(*Options)) (*RebuildEnvironmentOutput, error) 
 RequestEnvironmentInfo(ctx context.Context, params *RequestEnvironmentInfoInput, optFns ...func(*Options)) (*RequestEnvironmentInfoOutput, error) 
 RestartAppServer(ctx context.Context, params *RestartAppServerInput, optFns ...func(*Options)) (*RestartAppServerOutput, error) 
 RetrieveEnvironmentInfo(ctx context.Context, params *RetrieveEnvironmentInfoInput, optFns ...func(*Options)) (*RetrieveEnvironmentInfoOutput, error) 
 SwapEnvironmentCNAMEs(ctx context.Context, params *SwapEnvironmentCNAMEsInput, optFns ...func(*Options)) (*SwapEnvironmentCNAMEsOutput, error) 
 TerminateEnvironment(ctx context.Context, params *TerminateEnvironmentInput, optFns ...func(*Options)) (*TerminateEnvironmentOutput, error) 
 UpdateApplication(ctx context.Context, params *UpdateApplicationInput, optFns ...func(*Options)) (*UpdateApplicationOutput, error) 
 UpdateApplicationResourceLifecycle(ctx context.Context, params *UpdateApplicationResourceLifecycleInput, optFns ...func(*Options)) (*UpdateApplicationResourceLifecycleOutput, error) 
 UpdateApplicationVersion(ctx context.Context, params *UpdateApplicationVersionInput, optFns ...func(*Options)) (*UpdateApplicationVersionOutput, error) 
 UpdateConfigurationTemplate(ctx context.Context, params *UpdateConfigurationTemplateInput, optFns ...func(*Options)) (*UpdateConfigurationTemplateOutput, error) 
 UpdateEnvironment(ctx context.Context, params *UpdateEnvironmentInput, optFns ...func(*Options)) (*UpdateEnvironmentOutput, error) 
 UpdateTagsForResource(ctx context.Context, params *UpdateTagsForResourceInput, optFns ...func(*Options)) (*UpdateTagsForResourceOutput, error) 
 ValidateConfigurationSettings(ctx context.Context, params *ValidateConfigurationSettingsInput, optFns ...func(*Options)) (*ValidateConfigurationSettingsOutput, error) 
}
