
package elasticinference

import (
    "github.com/aws/aws-sdk-go-v2/service/elasticinference"
)

// IElasticinference defines the interface for elasticinference
type IElasticinference interface {
 Options() Options 
 DescribeAcceleratorOfferings(ctx context.Context, params *DescribeAcceleratorOfferingsInput, optFns ...func(*Options)) (*DescribeAcceleratorOfferingsOutput, error) 
 DescribeAcceleratorTypes(ctx context.Context, params *DescribeAcceleratorTypesInput, optFns ...func(*Options)) (*DescribeAcceleratorTypesOutput, error) 
 DescribeAccelerators(ctx context.Context, params *DescribeAcceleratorsInput, optFns ...func(*Options)) (*DescribeAcceleratorsOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
}
