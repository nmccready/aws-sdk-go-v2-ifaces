
package arczonalshift

import (
    "github.com/aws/aws-sdk-go-v2/service/arczonalshift"
)

// IClient defines the interface for arczonalshift
type IClient interface {
 Options() Options 
 CancelZonalShift(ctx context.Context, params *CancelZonalShiftInput, optFns ...func(*Options)) (*CancelZonalShiftOutput, error) 
 CreatePracticeRunConfiguration(ctx context.Context, params *CreatePracticeRunConfigurationInput, optFns ...func(*Options)) (*CreatePracticeRunConfigurationOutput, error) 
 DeletePracticeRunConfiguration(ctx context.Context, params *DeletePracticeRunConfigurationInput, optFns ...func(*Options)) (*DeletePracticeRunConfigurationOutput, error) 
 GetManagedResource(ctx context.Context, params *GetManagedResourceInput, optFns ...func(*Options)) (*GetManagedResourceOutput, error) 
 ListAutoshifts(ctx context.Context, params *ListAutoshiftsInput, optFns ...func(*Options)) (*ListAutoshiftsOutput, error) 
 ListManagedResources(ctx context.Context, params *ListManagedResourcesInput, optFns ...func(*Options)) (*ListManagedResourcesOutput, error) 
 ListZonalShifts(ctx context.Context, params *ListZonalShiftsInput, optFns ...func(*Options)) (*ListZonalShiftsOutput, error) 
 StartZonalShift(ctx context.Context, params *StartZonalShiftInput, optFns ...func(*Options)) (*StartZonalShiftOutput, error) 
 UpdatePracticeRunConfiguration(ctx context.Context, params *UpdatePracticeRunConfigurationInput, optFns ...func(*Options)) (*UpdatePracticeRunConfigurationOutput, error) 
 UpdateZonalAutoshiftConfiguration(ctx context.Context, params *UpdateZonalAutoshiftConfigurationInput, optFns ...func(*Options)) (*UpdateZonalAutoshiftConfigurationOutput, error) 
 UpdateZonalShift(ctx context.Context, params *UpdateZonalShiftInput, optFns ...func(*Options)) (*UpdateZonalShiftOutput, error) 
}
