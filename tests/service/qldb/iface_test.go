// Code generated by ./scripts/gen_tests.sh DO NOT EDIT.
package qldb_test

// tests for the qldb service interface for 
// this ../../../service/qldb/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/qldb"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/qldb/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/qldb/qldb_iface"
	"github.com/stretchr/testify/assert"
)

func TestQldbServiceCanBeMocked(t *testing.T) {
	var iface qldb_iface.IClient
	iface = &qldb.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := qldb.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCancelJournalKinesisStream", func(t *testing.T) {
        input := &qldb.CancelJournalKinesisStreamInput{}
        output := &qldb.CancelJournalKinesisStreamOutput{}

        mockClient.On("CancelJournalKinesisStream", ctx, input).Return(output, nil)

        result, err := mockClient.CancelJournalKinesisStream(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateLedger", func(t *testing.T) {
        input := &qldb.CreateLedgerInput{}
        output := &qldb.CreateLedgerOutput{}

        mockClient.On("CreateLedger", ctx, input).Return(output, nil)

        result, err := mockClient.CreateLedger(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteLedger", func(t *testing.T) {
        input := &qldb.DeleteLedgerInput{}
        output := &qldb.DeleteLedgerOutput{}

        mockClient.On("DeleteLedger", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteLedger(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeJournalKinesisStream", func(t *testing.T) {
        input := &qldb.DescribeJournalKinesisStreamInput{}
        output := &qldb.DescribeJournalKinesisStreamOutput{}

        mockClient.On("DescribeJournalKinesisStream", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeJournalKinesisStream(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeJournalS3Export", func(t *testing.T) {
        input := &qldb.DescribeJournalS3ExportInput{}
        output := &qldb.DescribeJournalS3ExportOutput{}

        mockClient.On("DescribeJournalS3Export", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeJournalS3Export(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeLedger", func(t *testing.T) {
        input := &qldb.DescribeLedgerInput{}
        output := &qldb.DescribeLedgerOutput{}

        mockClient.On("DescribeLedger", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeLedger(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestExportJournalToS3", func(t *testing.T) {
        input := &qldb.ExportJournalToS3Input{}
        output := &qldb.ExportJournalToS3Output{}

        mockClient.On("ExportJournalToS3", ctx, input).Return(output, nil)

        result, err := mockClient.ExportJournalToS3(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetBlock", func(t *testing.T) {
        input := &qldb.GetBlockInput{}
        output := &qldb.GetBlockOutput{}

        mockClient.On("GetBlock", ctx, input).Return(output, nil)

        result, err := mockClient.GetBlock(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDigest", func(t *testing.T) {
        input := &qldb.GetDigestInput{}
        output := &qldb.GetDigestOutput{}

        mockClient.On("GetDigest", ctx, input).Return(output, nil)

        result, err := mockClient.GetDigest(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetRevision", func(t *testing.T) {
        input := &qldb.GetRevisionInput{}
        output := &qldb.GetRevisionOutput{}

        mockClient.On("GetRevision", ctx, input).Return(output, nil)

        result, err := mockClient.GetRevision(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListJournalKinesisStreamsForLedger", func(t *testing.T) {
        input := &qldb.ListJournalKinesisStreamsForLedgerInput{}
        output := &qldb.ListJournalKinesisStreamsForLedgerOutput{}

        mockClient.On("ListJournalKinesisStreamsForLedger", ctx, input).Return(output, nil)

        result, err := mockClient.ListJournalKinesisStreamsForLedger(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListJournalS3Exports", func(t *testing.T) {
        input := &qldb.ListJournalS3ExportsInput{}
        output := &qldb.ListJournalS3ExportsOutput{}

        mockClient.On("ListJournalS3Exports", ctx, input).Return(output, nil)

        result, err := mockClient.ListJournalS3Exports(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListJournalS3ExportsForLedger", func(t *testing.T) {
        input := &qldb.ListJournalS3ExportsForLedgerInput{}
        output := &qldb.ListJournalS3ExportsForLedgerOutput{}

        mockClient.On("ListJournalS3ExportsForLedger", ctx, input).Return(output, nil)

        result, err := mockClient.ListJournalS3ExportsForLedger(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListLedgers", func(t *testing.T) {
        input := &qldb.ListLedgersInput{}
        output := &qldb.ListLedgersOutput{}

        mockClient.On("ListLedgers", ctx, input).Return(output, nil)

        result, err := mockClient.ListLedgers(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &qldb.ListTagsForResourceInput{}
        output := &qldb.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStreamJournalToKinesis", func(t *testing.T) {
        input := &qldb.StreamJournalToKinesisInput{}
        output := &qldb.StreamJournalToKinesisOutput{}

        mockClient.On("StreamJournalToKinesis", ctx, input).Return(output, nil)

        result, err := mockClient.StreamJournalToKinesis(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &qldb.TagResourceInput{}
        output := &qldb.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &qldb.UntagResourceInput{}
        output := &qldb.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateLedger", func(t *testing.T) {
        input := &qldb.UpdateLedgerInput{}
        output := &qldb.UpdateLedgerOutput{}

        mockClient.On("UpdateLedger", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateLedger(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateLedgerPermissionsMode", func(t *testing.T) {
        input := &qldb.UpdateLedgerPermissionsModeInput{}
        output := &qldb.UpdateLedgerPermissionsModeOutput{}

        mockClient.On("UpdateLedgerPermissionsMode", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateLedgerPermissionsMode(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
