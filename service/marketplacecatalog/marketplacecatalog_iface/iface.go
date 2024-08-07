
package marketplacecatalog_iface

// This file was generated by scripts/gen_ifaces.sh do not edit it manually
import (
    "context"
    . "github.com/aws/aws-sdk-go-v2/service/marketplacecatalog"
)

// IClient defines the interface for marketplacecatalog
type IClient interface {
 Options() Options 
 BatchDescribeEntities(ctx context.Context, params *BatchDescribeEntitiesInput, optFns ...func(*Options)) (*BatchDescribeEntitiesOutput, error) 
 CancelChangeSet(ctx context.Context, params *CancelChangeSetInput, optFns ...func(*Options)) (*CancelChangeSetOutput, error) 
 DeleteResourcePolicy(ctx context.Context, params *DeleteResourcePolicyInput, optFns ...func(*Options)) (*DeleteResourcePolicyOutput, error) 
 DescribeChangeSet(ctx context.Context, params *DescribeChangeSetInput, optFns ...func(*Options)) (*DescribeChangeSetOutput, error) 
 DescribeEntity(ctx context.Context, params *DescribeEntityInput, optFns ...func(*Options)) (*DescribeEntityOutput, error) 
 GetResourcePolicy(ctx context.Context, params *GetResourcePolicyInput, optFns ...func(*Options)) (*GetResourcePolicyOutput, error) 
 ListChangeSets(ctx context.Context, params *ListChangeSetsInput, optFns ...func(*Options)) (*ListChangeSetsOutput, error) 
 ListEntities(ctx context.Context, params *ListEntitiesInput, optFns ...func(*Options)) (*ListEntitiesOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 PutResourcePolicy(ctx context.Context, params *PutResourcePolicyInput, optFns ...func(*Options)) (*PutResourcePolicyOutput, error) 
 StartChangeSet(ctx context.Context, params *StartChangeSetInput, optFns ...func(*Options)) (*StartChangeSetOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
}
