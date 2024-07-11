
package greengrassv2

import (
    "github.com/aws/aws-sdk-go-v2/service/greengrassv2"
)

// IGreengrassv2 defines the interface for greengrassv2
type IGreengrassv2 interface {
 Options() Options 
 AssociateServiceRoleToAccount(ctx context.Context, params *AssociateServiceRoleToAccountInput, optFns ...func(*Options)) (*AssociateServiceRoleToAccountOutput, error) 
 BatchAssociateClientDeviceWithCoreDevice(ctx context.Context, params *BatchAssociateClientDeviceWithCoreDeviceInput, optFns ...func(*Options)) (*BatchAssociateClientDeviceWithCoreDeviceOutput, error) 
 BatchDisassociateClientDeviceFromCoreDevice(ctx context.Context, params *BatchDisassociateClientDeviceFromCoreDeviceInput, optFns ...func(*Options)) (*BatchDisassociateClientDeviceFromCoreDeviceOutput, error) 
 CancelDeployment(ctx context.Context, params *CancelDeploymentInput, optFns ...func(*Options)) (*CancelDeploymentOutput, error) 
 CreateComponentVersion(ctx context.Context, params *CreateComponentVersionInput, optFns ...func(*Options)) (*CreateComponentVersionOutput, error) 
 CreateDeployment(ctx context.Context, params *CreateDeploymentInput, optFns ...func(*Options)) (*CreateDeploymentOutput, error) 
 DeleteComponent(ctx context.Context, params *DeleteComponentInput, optFns ...func(*Options)) (*DeleteComponentOutput, error) 
 DeleteCoreDevice(ctx context.Context, params *DeleteCoreDeviceInput, optFns ...func(*Options)) (*DeleteCoreDeviceOutput, error) 
 DeleteDeployment(ctx context.Context, params *DeleteDeploymentInput, optFns ...func(*Options)) (*DeleteDeploymentOutput, error) 
 DescribeComponent(ctx context.Context, params *DescribeComponentInput, optFns ...func(*Options)) (*DescribeComponentOutput, error) 
 DisassociateServiceRoleFromAccount(ctx context.Context, params *DisassociateServiceRoleFromAccountInput, optFns ...func(*Options)) (*DisassociateServiceRoleFromAccountOutput, error) 
 GetComponent(ctx context.Context, params *GetComponentInput, optFns ...func(*Options)) (*GetComponentOutput, error) 
 GetComponentVersionArtifact(ctx context.Context, params *GetComponentVersionArtifactInput, optFns ...func(*Options)) (*GetComponentVersionArtifactOutput, error) 
 GetConnectivityInfo(ctx context.Context, params *GetConnectivityInfoInput, optFns ...func(*Options)) (*GetConnectivityInfoOutput, error) 
 GetCoreDevice(ctx context.Context, params *GetCoreDeviceInput, optFns ...func(*Options)) (*GetCoreDeviceOutput, error) 
 GetDeployment(ctx context.Context, params *GetDeploymentInput, optFns ...func(*Options)) (*GetDeploymentOutput, error) 
 GetServiceRoleForAccount(ctx context.Context, params *GetServiceRoleForAccountInput, optFns ...func(*Options)) (*GetServiceRoleForAccountOutput, error) 
 ListClientDevicesAssociatedWithCoreDevice(ctx context.Context, params *ListClientDevicesAssociatedWithCoreDeviceInput, optFns ...func(*Options)) (*ListClientDevicesAssociatedWithCoreDeviceOutput, error) 
 ListComponentVersions(ctx context.Context, params *ListComponentVersionsInput, optFns ...func(*Options)) (*ListComponentVersionsOutput, error) 
 ListComponents(ctx context.Context, params *ListComponentsInput, optFns ...func(*Options)) (*ListComponentsOutput, error) 
 ListCoreDevices(ctx context.Context, params *ListCoreDevicesInput, optFns ...func(*Options)) (*ListCoreDevicesOutput, error) 
 ListDeployments(ctx context.Context, params *ListDeploymentsInput, optFns ...func(*Options)) (*ListDeploymentsOutput, error) 
 ListEffectiveDeployments(ctx context.Context, params *ListEffectiveDeploymentsInput, optFns ...func(*Options)) (*ListEffectiveDeploymentsOutput, error) 
 ListInstalledComponents(ctx context.Context, params *ListInstalledComponentsInput, optFns ...func(*Options)) (*ListInstalledComponentsOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 ResolveComponentCandidates(ctx context.Context, params *ResolveComponentCandidatesInput, optFns ...func(*Options)) (*ResolveComponentCandidatesOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateConnectivityInfo(ctx context.Context, params *UpdateConnectivityInfoInput, optFns ...func(*Options)) (*UpdateConnectivityInfoOutput, error) 
}
