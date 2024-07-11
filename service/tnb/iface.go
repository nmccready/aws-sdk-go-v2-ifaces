
package tnb

// This file was generated by scripts/gen_ifaces.sh do not edit it manually
import (
    "context"
    . "github.com/aws/aws-sdk-go-v2/service/tnb"
)

// IClient defines the interface for tnb
type IClient interface {
 Options() Options 
 CancelSolNetworkOperation(ctx context.Context, params *CancelSolNetworkOperationInput, optFns ...func(*Options)) (*CancelSolNetworkOperationOutput, error) 
 CreateSolFunctionPackage(ctx context.Context, params *CreateSolFunctionPackageInput, optFns ...func(*Options)) (*CreateSolFunctionPackageOutput, error) 
 CreateSolNetworkInstance(ctx context.Context, params *CreateSolNetworkInstanceInput, optFns ...func(*Options)) (*CreateSolNetworkInstanceOutput, error) 
 CreateSolNetworkPackage(ctx context.Context, params *CreateSolNetworkPackageInput, optFns ...func(*Options)) (*CreateSolNetworkPackageOutput, error) 
 DeleteSolFunctionPackage(ctx context.Context, params *DeleteSolFunctionPackageInput, optFns ...func(*Options)) (*DeleteSolFunctionPackageOutput, error) 
 DeleteSolNetworkInstance(ctx context.Context, params *DeleteSolNetworkInstanceInput, optFns ...func(*Options)) (*DeleteSolNetworkInstanceOutput, error) 
 DeleteSolNetworkPackage(ctx context.Context, params *DeleteSolNetworkPackageInput, optFns ...func(*Options)) (*DeleteSolNetworkPackageOutput, error) 
 GetSolFunctionInstance(ctx context.Context, params *GetSolFunctionInstanceInput, optFns ...func(*Options)) (*GetSolFunctionInstanceOutput, error) 
 GetSolFunctionPackage(ctx context.Context, params *GetSolFunctionPackageInput, optFns ...func(*Options)) (*GetSolFunctionPackageOutput, error) 
 GetSolFunctionPackageContent(ctx context.Context, params *GetSolFunctionPackageContentInput, optFns ...func(*Options)) (*GetSolFunctionPackageContentOutput, error) 
 GetSolFunctionPackageDescriptor(ctx context.Context, params *GetSolFunctionPackageDescriptorInput, optFns ...func(*Options)) (*GetSolFunctionPackageDescriptorOutput, error) 
 GetSolNetworkInstance(ctx context.Context, params *GetSolNetworkInstanceInput, optFns ...func(*Options)) (*GetSolNetworkInstanceOutput, error) 
 GetSolNetworkOperation(ctx context.Context, params *GetSolNetworkOperationInput, optFns ...func(*Options)) (*GetSolNetworkOperationOutput, error) 
 GetSolNetworkPackage(ctx context.Context, params *GetSolNetworkPackageInput, optFns ...func(*Options)) (*GetSolNetworkPackageOutput, error) 
 GetSolNetworkPackageContent(ctx context.Context, params *GetSolNetworkPackageContentInput, optFns ...func(*Options)) (*GetSolNetworkPackageContentOutput, error) 
 GetSolNetworkPackageDescriptor(ctx context.Context, params *GetSolNetworkPackageDescriptorInput, optFns ...func(*Options)) (*GetSolNetworkPackageDescriptorOutput, error) 
 InstantiateSolNetworkInstance(ctx context.Context, params *InstantiateSolNetworkInstanceInput, optFns ...func(*Options)) (*InstantiateSolNetworkInstanceOutput, error) 
 ListSolFunctionInstances(ctx context.Context, params *ListSolFunctionInstancesInput, optFns ...func(*Options)) (*ListSolFunctionInstancesOutput, error) 
 ListSolFunctionPackages(ctx context.Context, params *ListSolFunctionPackagesInput, optFns ...func(*Options)) (*ListSolFunctionPackagesOutput, error) 
 ListSolNetworkInstances(ctx context.Context, params *ListSolNetworkInstancesInput, optFns ...func(*Options)) (*ListSolNetworkInstancesOutput, error) 
 ListSolNetworkOperations(ctx context.Context, params *ListSolNetworkOperationsInput, optFns ...func(*Options)) (*ListSolNetworkOperationsOutput, error) 
 ListSolNetworkPackages(ctx context.Context, params *ListSolNetworkPackagesInput, optFns ...func(*Options)) (*ListSolNetworkPackagesOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 PutSolFunctionPackageContent(ctx context.Context, params *PutSolFunctionPackageContentInput, optFns ...func(*Options)) (*PutSolFunctionPackageContentOutput, error) 
 PutSolNetworkPackageContent(ctx context.Context, params *PutSolNetworkPackageContentInput, optFns ...func(*Options)) (*PutSolNetworkPackageContentOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 TerminateSolNetworkInstance(ctx context.Context, params *TerminateSolNetworkInstanceInput, optFns ...func(*Options)) (*TerminateSolNetworkInstanceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateSolFunctionPackage(ctx context.Context, params *UpdateSolFunctionPackageInput, optFns ...func(*Options)) (*UpdateSolFunctionPackageOutput, error) 
 UpdateSolNetworkInstance(ctx context.Context, params *UpdateSolNetworkInstanceInput, optFns ...func(*Options)) (*UpdateSolNetworkInstanceOutput, error) 
 UpdateSolNetworkPackage(ctx context.Context, params *UpdateSolNetworkPackageInput, optFns ...func(*Options)) (*UpdateSolNetworkPackageOutput, error) 
 ValidateSolFunctionPackageContent(ctx context.Context, params *ValidateSolFunctionPackageContentInput, optFns ...func(*Options)) (*ValidateSolFunctionPackageContentOutput, error) 
 ValidateSolNetworkPackageContent(ctx context.Context, params *ValidateSolNetworkPackageContentInput, optFns ...func(*Options)) (*ValidateSolNetworkPackageContentOutput, error) 
}
