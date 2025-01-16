// Code generated by ./scripts/gen_tests.sh DO NOT EDIT.
package backupsearch_test

// tests for the backupsearch service interface for 
// this ../../../service/backupsearch/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/backupsearch"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/backupsearch/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/backupsearch/backupsearch_iface"
	"github.com/stretchr/testify/assert"
)

func TestBackupsearchServiceCanBeMocked(t *testing.T) {
	var iface backupsearch_iface.IClient
	iface = &backupsearch.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := backupsearch.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSearchJob", func(t *testing.T) {
        input := &backupsearch.GetSearchJobInput{}
        output := &backupsearch.GetSearchJobOutput{}

        mockClient.On("GetSearchJob", ctx, input).Return(output, nil)

        result, err := mockClient.GetSearchJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSearchResultExportJob", func(t *testing.T) {
        input := &backupsearch.GetSearchResultExportJobInput{}
        output := &backupsearch.GetSearchResultExportJobOutput{}

        mockClient.On("GetSearchResultExportJob", ctx, input).Return(output, nil)

        result, err := mockClient.GetSearchResultExportJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListSearchJobBackups", func(t *testing.T) {
        input := &backupsearch.ListSearchJobBackupsInput{}
        output := &backupsearch.ListSearchJobBackupsOutput{}

        mockClient.On("ListSearchJobBackups", ctx, input).Return(output, nil)

        result, err := mockClient.ListSearchJobBackups(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListSearchJobResults", func(t *testing.T) {
        input := &backupsearch.ListSearchJobResultsInput{}
        output := &backupsearch.ListSearchJobResultsOutput{}

        mockClient.On("ListSearchJobResults", ctx, input).Return(output, nil)

        result, err := mockClient.ListSearchJobResults(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListSearchJobs", func(t *testing.T) {
        input := &backupsearch.ListSearchJobsInput{}
        output := &backupsearch.ListSearchJobsOutput{}

        mockClient.On("ListSearchJobs", ctx, input).Return(output, nil)

        result, err := mockClient.ListSearchJobs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListSearchResultExportJobs", func(t *testing.T) {
        input := &backupsearch.ListSearchResultExportJobsInput{}
        output := &backupsearch.ListSearchResultExportJobsOutput{}

        mockClient.On("ListSearchResultExportJobs", ctx, input).Return(output, nil)

        result, err := mockClient.ListSearchResultExportJobs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &backupsearch.ListTagsForResourceInput{}
        output := &backupsearch.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartSearchJob", func(t *testing.T) {
        input := &backupsearch.StartSearchJobInput{}
        output := &backupsearch.StartSearchJobOutput{}

        mockClient.On("StartSearchJob", ctx, input).Return(output, nil)

        result, err := mockClient.StartSearchJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartSearchResultExportJob", func(t *testing.T) {
        input := &backupsearch.StartSearchResultExportJobInput{}
        output := &backupsearch.StartSearchResultExportJobOutput{}

        mockClient.On("StartSearchResultExportJob", ctx, input).Return(output, nil)

        result, err := mockClient.StartSearchResultExportJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopSearchJob", func(t *testing.T) {
        input := &backupsearch.StopSearchJobInput{}
        output := &backupsearch.StopSearchJobOutput{}

        mockClient.On("StopSearchJob", ctx, input).Return(output, nil)

        result, err := mockClient.StopSearchJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &backupsearch.TagResourceInput{}
        output := &backupsearch.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &backupsearch.UntagResourceInput{}
        output := &backupsearch.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}