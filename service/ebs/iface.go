
package ebs

import (
    "github.com/aws/aws-sdk-go-v2/service/ebs"
)

// IClient defines the interface for ebs
type IClient interface {
 Options() Options 
 CompleteSnapshot(ctx context.Context, params *CompleteSnapshotInput, optFns ...func(*Options)) (*CompleteSnapshotOutput, error) 
 GetSnapshotBlock(ctx context.Context, params *GetSnapshotBlockInput, optFns ...func(*Options)) (*GetSnapshotBlockOutput, error) 
 ListChangedBlocks(ctx context.Context, params *ListChangedBlocksInput, optFns ...func(*Options)) (*ListChangedBlocksOutput, error) 
 ListSnapshotBlocks(ctx context.Context, params *ListSnapshotBlocksInput, optFns ...func(*Options)) (*ListSnapshotBlocksOutput, error) 
 PutSnapshotBlock(ctx context.Context, params *PutSnapshotBlockInput, optFns ...func(*Options)) (*PutSnapshotBlockOutput, error) 
 StartSnapshot(ctx context.Context, params *StartSnapshotInput, optFns ...func(*Options)) (*StartSnapshotOutput, error) 
}
