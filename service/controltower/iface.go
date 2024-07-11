
package controltower

import (
    "github.com/aws/aws-sdk-go-v2/service/controltower"
)

// IControltower defines the interface for controltower
type IControltower interface {
 Options() Options 
 CreateLandingZone(ctx context.Context, params *CreateLandingZoneInput, optFns ...func(*Options)) (*CreateLandingZoneOutput, error) 
 DeleteLandingZone(ctx context.Context, params *DeleteLandingZoneInput, optFns ...func(*Options)) (*DeleteLandingZoneOutput, error) 
 DisableBaseline(ctx context.Context, params *DisableBaselineInput, optFns ...func(*Options)) (*DisableBaselineOutput, error) 
 DisableControl(ctx context.Context, params *DisableControlInput, optFns ...func(*Options)) (*DisableControlOutput, error) 
 EnableBaseline(ctx context.Context, params *EnableBaselineInput, optFns ...func(*Options)) (*EnableBaselineOutput, error) 
 EnableControl(ctx context.Context, params *EnableControlInput, optFns ...func(*Options)) (*EnableControlOutput, error) 
 GetBaseline(ctx context.Context, params *GetBaselineInput, optFns ...func(*Options)) (*GetBaselineOutput, error) 
 GetBaselineOperation(ctx context.Context, params *GetBaselineOperationInput, optFns ...func(*Options)) (*GetBaselineOperationOutput, error) 
 GetControlOperation(ctx context.Context, params *GetControlOperationInput, optFns ...func(*Options)) (*GetControlOperationOutput, error) 
 GetEnabledBaseline(ctx context.Context, params *GetEnabledBaselineInput, optFns ...func(*Options)) (*GetEnabledBaselineOutput, error) 
 GetEnabledControl(ctx context.Context, params *GetEnabledControlInput, optFns ...func(*Options)) (*GetEnabledControlOutput, error) 
 GetLandingZone(ctx context.Context, params *GetLandingZoneInput, optFns ...func(*Options)) (*GetLandingZoneOutput, error) 
 GetLandingZoneOperation(ctx context.Context, params *GetLandingZoneOperationInput, optFns ...func(*Options)) (*GetLandingZoneOperationOutput, error) 
 ListBaselines(ctx context.Context, params *ListBaselinesInput, optFns ...func(*Options)) (*ListBaselinesOutput, error) 
 ListControlOperations(ctx context.Context, params *ListControlOperationsInput, optFns ...func(*Options)) (*ListControlOperationsOutput, error) 
 ListEnabledBaselines(ctx context.Context, params *ListEnabledBaselinesInput, optFns ...func(*Options)) (*ListEnabledBaselinesOutput, error) 
 ListEnabledControls(ctx context.Context, params *ListEnabledControlsInput, optFns ...func(*Options)) (*ListEnabledControlsOutput, error) 
 ListLandingZoneOperations(ctx context.Context, params *ListLandingZoneOperationsInput, optFns ...func(*Options)) (*ListLandingZoneOperationsOutput, error) 
 ListLandingZones(ctx context.Context, params *ListLandingZonesInput, optFns ...func(*Options)) (*ListLandingZonesOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 ResetEnabledBaseline(ctx context.Context, params *ResetEnabledBaselineInput, optFns ...func(*Options)) (*ResetEnabledBaselineOutput, error) 
 ResetLandingZone(ctx context.Context, params *ResetLandingZoneInput, optFns ...func(*Options)) (*ResetLandingZoneOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateEnabledBaseline(ctx context.Context, params *UpdateEnabledBaselineInput, optFns ...func(*Options)) (*UpdateEnabledBaselineOutput, error) 
 UpdateEnabledControl(ctx context.Context, params *UpdateEnabledControlInput, optFns ...func(*Options)) (*UpdateEnabledControlOutput, error) 
 UpdateLandingZone(ctx context.Context, params *UpdateLandingZoneInput, optFns ...func(*Options)) (*UpdateLandingZoneOutput, error) 
}
