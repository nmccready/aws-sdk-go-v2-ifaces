
package migrationhubconfig_iface

// This file was generated by scripts/gen_ifaces.sh do not edit it manually
import (
    "context"
    . "github.com/aws/aws-sdk-go-v2/service/migrationhubconfig"
)

// IClient defines the interface for migrationhubconfig
type IClient interface {
 Options() Options 
 CreateHomeRegionControl(ctx context.Context, params *CreateHomeRegionControlInput, optFns ...func(*Options)) (*CreateHomeRegionControlOutput, error) 
 DeleteHomeRegionControl(ctx context.Context, params *DeleteHomeRegionControlInput, optFns ...func(*Options)) (*DeleteHomeRegionControlOutput, error) 
 DescribeHomeRegionControls(ctx context.Context, params *DescribeHomeRegionControlsInput, optFns ...func(*Options)) (*DescribeHomeRegionControlsOutput, error) 
 GetHomeRegion(ctx context.Context, params *GetHomeRegionInput, optFns ...func(*Options)) (*GetHomeRegionOutput, error) 
}