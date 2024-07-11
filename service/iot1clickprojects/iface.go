
package iot1clickprojects

import (
    "github.com/aws/aws-sdk-go-v2/service/iot1clickprojects"
)

// IIot1clickprojects defines the interface for iot1clickprojects
type IIot1clickprojects interface {
 Options() Options 
 AssociateDeviceWithPlacement(ctx context.Context, params *AssociateDeviceWithPlacementInput, optFns ...func(*Options)) (*AssociateDeviceWithPlacementOutput, error) 
 CreatePlacement(ctx context.Context, params *CreatePlacementInput, optFns ...func(*Options)) (*CreatePlacementOutput, error) 
 CreateProject(ctx context.Context, params *CreateProjectInput, optFns ...func(*Options)) (*CreateProjectOutput, error) 
 DeletePlacement(ctx context.Context, params *DeletePlacementInput, optFns ...func(*Options)) (*DeletePlacementOutput, error) 
 DeleteProject(ctx context.Context, params *DeleteProjectInput, optFns ...func(*Options)) (*DeleteProjectOutput, error) 
 DescribePlacement(ctx context.Context, params *DescribePlacementInput, optFns ...func(*Options)) (*DescribePlacementOutput, error) 
 DescribeProject(ctx context.Context, params *DescribeProjectInput, optFns ...func(*Options)) (*DescribeProjectOutput, error) 
 DisassociateDeviceFromPlacement(ctx context.Context, params *DisassociateDeviceFromPlacementInput, optFns ...func(*Options)) (*DisassociateDeviceFromPlacementOutput, error) 
 GetDevicesInPlacement(ctx context.Context, params *GetDevicesInPlacementInput, optFns ...func(*Options)) (*GetDevicesInPlacementOutput, error) 
 ListPlacements(ctx context.Context, params *ListPlacementsInput, optFns ...func(*Options)) (*ListPlacementsOutput, error) 
 ListProjects(ctx context.Context, params *ListProjectsInput, optFns ...func(*Options)) (*ListProjectsOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdatePlacement(ctx context.Context, params *UpdatePlacementInput, optFns ...func(*Options)) (*UpdatePlacementOutput, error) 
 UpdateProject(ctx context.Context, params *UpdateProjectInput, optFns ...func(*Options)) (*UpdateProjectOutput, error) 
}
