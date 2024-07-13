package medicalimaging_test

// tests for the medicalimaging service interface for this ../../../service/medicalimaging/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/medicalimaging"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/medicalimaging/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/medicalimaging/medicalimaging_iface"
	"github.com/stretchr/testify/assert"
)

func TestMedicalimagingServiceCanBeMocked(t *testing.T) {
	var iface medicalimaging_iface.IClient
	iface = &medicalimaging.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := medicalimaging.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCopyImageSet", func(t *testing.T) {
        input := &medicalimaging.CopyImageSetInput{}
        output := &medicalimaging.CopyImageSetOutput{}

        mockClient.On("CopyImageSet", ctx, input).Return(output, nil)

        result, err := mockClient.CopyImageSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateDatastore", func(t *testing.T) {
        input := &medicalimaging.CreateDatastoreInput{}
        output := &medicalimaging.CreateDatastoreOutput{}

        mockClient.On("CreateDatastore", ctx, input).Return(output, nil)

        result, err := mockClient.CreateDatastore(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteDatastore", func(t *testing.T) {
        input := &medicalimaging.DeleteDatastoreInput{}
        output := &medicalimaging.DeleteDatastoreOutput{}

        mockClient.On("DeleteDatastore", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteDatastore(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteImageSet", func(t *testing.T) {
        input := &medicalimaging.DeleteImageSetInput{}
        output := &medicalimaging.DeleteImageSetOutput{}

        mockClient.On("DeleteImageSet", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteImageSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDICOMImportJob", func(t *testing.T) {
        input := &medicalimaging.GetDICOMImportJobInput{}
        output := &medicalimaging.GetDICOMImportJobOutput{}

        mockClient.On("GetDICOMImportJob", ctx, input).Return(output, nil)

        result, err := mockClient.GetDICOMImportJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDatastore", func(t *testing.T) {
        input := &medicalimaging.GetDatastoreInput{}
        output := &medicalimaging.GetDatastoreOutput{}

        mockClient.On("GetDatastore", ctx, input).Return(output, nil)

        result, err := mockClient.GetDatastore(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetImageFrame", func(t *testing.T) {
        input := &medicalimaging.GetImageFrameInput{}
        output := &medicalimaging.GetImageFrameOutput{}

        mockClient.On("GetImageFrame", ctx, input).Return(output, nil)

        result, err := mockClient.GetImageFrame(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetImageSet", func(t *testing.T) {
        input := &medicalimaging.GetImageSetInput{}
        output := &medicalimaging.GetImageSetOutput{}

        mockClient.On("GetImageSet", ctx, input).Return(output, nil)

        result, err := mockClient.GetImageSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetImageSetMetadata", func(t *testing.T) {
        input := &medicalimaging.GetImageSetMetadataInput{}
        output := &medicalimaging.GetImageSetMetadataOutput{}

        mockClient.On("GetImageSetMetadata", ctx, input).Return(output, nil)

        result, err := mockClient.GetImageSetMetadata(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDICOMImportJobs", func(t *testing.T) {
        input := &medicalimaging.ListDICOMImportJobsInput{}
        output := &medicalimaging.ListDICOMImportJobsOutput{}

        mockClient.On("ListDICOMImportJobs", ctx, input).Return(output, nil)

        result, err := mockClient.ListDICOMImportJobs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDatastores", func(t *testing.T) {
        input := &medicalimaging.ListDatastoresInput{}
        output := &medicalimaging.ListDatastoresOutput{}

        mockClient.On("ListDatastores", ctx, input).Return(output, nil)

        result, err := mockClient.ListDatastores(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListImageSetVersions", func(t *testing.T) {
        input := &medicalimaging.ListImageSetVersionsInput{}
        output := &medicalimaging.ListImageSetVersionsOutput{}

        mockClient.On("ListImageSetVersions", ctx, input).Return(output, nil)

        result, err := mockClient.ListImageSetVersions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &medicalimaging.ListTagsForResourceInput{}
        output := &medicalimaging.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSearchImageSets", func(t *testing.T) {
        input := &medicalimaging.SearchImageSetsInput{}
        output := &medicalimaging.SearchImageSetsOutput{}

        mockClient.On("SearchImageSets", ctx, input).Return(output, nil)

        result, err := mockClient.SearchImageSets(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartDICOMImportJob", func(t *testing.T) {
        input := &medicalimaging.StartDICOMImportJobInput{}
        output := &medicalimaging.StartDICOMImportJobOutput{}

        mockClient.On("StartDICOMImportJob", ctx, input).Return(output, nil)

        result, err := mockClient.StartDICOMImportJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &medicalimaging.TagResourceInput{}
        output := &medicalimaging.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &medicalimaging.UntagResourceInput{}
        output := &medicalimaging.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateImageSetMetadata", func(t *testing.T) {
        input := &medicalimaging.UpdateImageSetMetadataInput{}
        output := &medicalimaging.UpdateImageSetMetadataOutput{}

        mockClient.On("UpdateImageSetMetadata", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateImageSetMetadata(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
