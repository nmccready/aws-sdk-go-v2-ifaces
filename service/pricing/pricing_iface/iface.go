
package pricing_iface

// This file was generated by scripts/gen_ifaces.sh do not edit it manually
import (
    "context"
    . "github.com/aws/aws-sdk-go-v2/service/pricing"
)

// IClient defines the interface for pricing
type IClient interface {
 Options() Options 
 DescribeServices(ctx context.Context, params *DescribeServicesInput, optFns ...func(*Options)) (*DescribeServicesOutput, error) 
 GetAttributeValues(ctx context.Context, params *GetAttributeValuesInput, optFns ...func(*Options)) (*GetAttributeValuesOutput, error) 
 GetPriceListFileUrl(ctx context.Context, params *GetPriceListFileUrlInput, optFns ...func(*Options)) (*GetPriceListFileUrlOutput, error) 
 GetProducts(ctx context.Context, params *GetProductsInput, optFns ...func(*Options)) (*GetProductsOutput, error) 
 ListPriceLists(ctx context.Context, params *ListPriceListsInput, optFns ...func(*Options)) (*ListPriceListsOutput, error) 
}