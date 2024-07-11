
package qldb

import (
    "github.com/aws/aws-sdk-go-v2/service/qldb"
)

// IQldb defines the interface for qldb
type IQldb interface {
 Options() Options 
 CancelJournalKinesisStream(ctx context.Context, params *CancelJournalKinesisStreamInput, optFns ...func(*Options)) (*CancelJournalKinesisStreamOutput, error) 
 CreateLedger(ctx context.Context, params *CreateLedgerInput, optFns ...func(*Options)) (*CreateLedgerOutput, error) 
 DeleteLedger(ctx context.Context, params *DeleteLedgerInput, optFns ...func(*Options)) (*DeleteLedgerOutput, error) 
 DescribeJournalKinesisStream(ctx context.Context, params *DescribeJournalKinesisStreamInput, optFns ...func(*Options)) (*DescribeJournalKinesisStreamOutput, error) 
 DescribeJournalS3Export(ctx context.Context, params *DescribeJournalS3ExportInput, optFns ...func(*Options)) (*DescribeJournalS3ExportOutput, error) 
 DescribeLedger(ctx context.Context, params *DescribeLedgerInput, optFns ...func(*Options)) (*DescribeLedgerOutput, error) 
 ExportJournalToS3(ctx context.Context, params *ExportJournalToS3Input, optFns ...func(*Options)) (*ExportJournalToS3Output, error) 
 GetBlock(ctx context.Context, params *GetBlockInput, optFns ...func(*Options)) (*GetBlockOutput, error) 
 GetDigest(ctx context.Context, params *GetDigestInput, optFns ...func(*Options)) (*GetDigestOutput, error) 
 GetRevision(ctx context.Context, params *GetRevisionInput, optFns ...func(*Options)) (*GetRevisionOutput, error) 
 ListJournalKinesisStreamsForLedger(ctx context.Context, params *ListJournalKinesisStreamsForLedgerInput, optFns ...func(*Options)) (*ListJournalKinesisStreamsForLedgerOutput, error) 
 ListJournalS3Exports(ctx context.Context, params *ListJournalS3ExportsInput, optFns ...func(*Options)) (*ListJournalS3ExportsOutput, error) 
 ListJournalS3ExportsForLedger(ctx context.Context, params *ListJournalS3ExportsForLedgerInput, optFns ...func(*Options)) (*ListJournalS3ExportsForLedgerOutput, error) 
 ListLedgers(ctx context.Context, params *ListLedgersInput, optFns ...func(*Options)) (*ListLedgersOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 StreamJournalToKinesis(ctx context.Context, params *StreamJournalToKinesisInput, optFns ...func(*Options)) (*StreamJournalToKinesisOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateLedger(ctx context.Context, params *UpdateLedgerInput, optFns ...func(*Options)) (*UpdateLedgerOutput, error) 
 UpdateLedgerPermissionsMode(ctx context.Context, params *UpdateLedgerPermissionsModeInput, optFns ...func(*Options)) (*UpdateLedgerPermissionsModeOutput, error) 
}
