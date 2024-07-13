package healthlake_test

// tests for the healthlake service interface for this ../../../service/healthlake/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/healthlake"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/healthlake/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/healthlake/healthlake_iface"
	"github.com/stretchr/testify/assert"
)

func TestHealthlakeServiceCanBeMocked(t *testing.T) {
	var iface healthlake_iface.IClient
	iface = &healthlake.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := healthlake.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateFHIRDatastore", func(t *testing.T) {
        input := &healthlake.CreateFHIRDatastoreInput{}
        output := &healthlake.CreateFHIRDatastoreOutput{}

        mockClient.On("CreateFHIRDatastore", ctx, input).Return(output, nil)

        result, err := mockClient.CreateFHIRDatastore(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteFHIRDatastore", func(t *testing.T) {
        input := &healthlake.DeleteFHIRDatastoreInput{}
        output := &healthlake.DeleteFHIRDatastoreOutput{}

        mockClient.On("DeleteFHIRDatastore", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteFHIRDatastore(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeFHIRDatastore", func(t *testing.T) {
        input := &healthlake.DescribeFHIRDatastoreInput{}
        output := &healthlake.DescribeFHIRDatastoreOutput{}

        mockClient.On("DescribeFHIRDatastore", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeFHIRDatastore(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeFHIRExportJob", func(t *testing.T) {
        input := &healthlake.DescribeFHIRExportJobInput{}
        output := &healthlake.DescribeFHIRExportJobOutput{}

        mockClient.On("DescribeFHIRExportJob", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeFHIRExportJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeFHIRImportJob", func(t *testing.T) {
        input := &healthlake.DescribeFHIRImportJobInput{}
        output := &healthlake.DescribeFHIRImportJobOutput{}

        mockClient.On("DescribeFHIRImportJob", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeFHIRImportJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListFHIRDatastores", func(t *testing.T) {
        input := &healthlake.ListFHIRDatastoresInput{}
        output := &healthlake.ListFHIRDatastoresOutput{}

        mockClient.On("ListFHIRDatastores", ctx, input).Return(output, nil)

        result, err := mockClient.ListFHIRDatastores(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListFHIRExportJobs", func(t *testing.T) {
        input := &healthlake.ListFHIRExportJobsInput{}
        output := &healthlake.ListFHIRExportJobsOutput{}

        mockClient.On("ListFHIRExportJobs", ctx, input).Return(output, nil)

        result, err := mockClient.ListFHIRExportJobs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListFHIRImportJobs", func(t *testing.T) {
        input := &healthlake.ListFHIRImportJobsInput{}
        output := &healthlake.ListFHIRImportJobsOutput{}

        mockClient.On("ListFHIRImportJobs", ctx, input).Return(output, nil)

        result, err := mockClient.ListFHIRImportJobs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &healthlake.ListTagsForResourceInput{}
        output := &healthlake.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartFHIRExportJob", func(t *testing.T) {
        input := &healthlake.StartFHIRExportJobInput{}
        output := &healthlake.StartFHIRExportJobOutput{}

        mockClient.On("StartFHIRExportJob", ctx, input).Return(output, nil)

        result, err := mockClient.StartFHIRExportJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartFHIRImportJob", func(t *testing.T) {
        input := &healthlake.StartFHIRImportJobInput{}
        output := &healthlake.StartFHIRImportJobOutput{}

        mockClient.On("StartFHIRImportJob", ctx, input).Return(output, nil)

        result, err := mockClient.StartFHIRImportJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &healthlake.TagResourceInput{}
        output := &healthlake.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &healthlake.UntagResourceInput{}
        output := &healthlake.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
