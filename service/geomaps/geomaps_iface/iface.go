
package geomaps_iface

// This file was generated by scripts/gen_ifaces.sh do not edit it manually
import (
    "context"
    . "github.com/aws/aws-sdk-go-v2/service/geomaps"
)

// IClient defines the interface for geomaps
type IClient interface {
 Options() Options 
 GetGlyphs(ctx context.Context, params *GetGlyphsInput, optFns ...func(*Options)) (*GetGlyphsOutput, error) 
 GetSprites(ctx context.Context, params *GetSpritesInput, optFns ...func(*Options)) (*GetSpritesOutput, error) 
 GetStaticMap(ctx context.Context, params *GetStaticMapInput, optFns ...func(*Options)) (*GetStaticMapOutput, error) 
 GetStyleDescriptor(ctx context.Context, params *GetStyleDescriptorInput, optFns ...func(*Options)) (*GetStyleDescriptorOutput, error) 
 GetTile(ctx context.Context, params *GetTileInput, optFns ...func(*Options)) (*GetTileOutput, error) 
}