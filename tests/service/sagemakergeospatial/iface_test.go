package sagemakergeospatial_test

// tests for the sagemakergeospatial service interface for this ../../../service/sagemakergeospatial/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/sagemakergeospatial"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/sagemakergeospatial/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/sagemakergeospatial/sagemakergeospatial_iface"
	"github.com/stretchr/testify/assert"
)

func TestSagemakergeospatialServiceCanBeMocked(t *testing.T) {
	var iface sagemakergeospatial_iface.IClient
	iface = &sagemakergeospatial.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := sagemakergeospatial.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteEarthObservationJob", func(t *testing.T) {
        input := &sagemakergeospatial.DeleteEarthObservationJobInput{}
        output := &sagemakergeospatial.DeleteEarthObservationJobOutput{}

        mockClient.On("DeleteEarthObservationJob", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteEarthObservationJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteVectorEnrichmentJob", func(t *testing.T) {
        input := &sagemakergeospatial.DeleteVectorEnrichmentJobInput{}
        output := &sagemakergeospatial.DeleteVectorEnrichmentJobOutput{}

        mockClient.On("DeleteVectorEnrichmentJob", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteVectorEnrichmentJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestExportEarthObservationJob", func(t *testing.T) {
        input := &sagemakergeospatial.ExportEarthObservationJobInput{}
        output := &sagemakergeospatial.ExportEarthObservationJobOutput{}

        mockClient.On("ExportEarthObservationJob", ctx, input).Return(output, nil)

        result, err := mockClient.ExportEarthObservationJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestExportVectorEnrichmentJob", func(t *testing.T) {
        input := &sagemakergeospatial.ExportVectorEnrichmentJobInput{}
        output := &sagemakergeospatial.ExportVectorEnrichmentJobOutput{}

        mockClient.On("ExportVectorEnrichmentJob", ctx, input).Return(output, nil)

        result, err := mockClient.ExportVectorEnrichmentJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetEarthObservationJob", func(t *testing.T) {
        input := &sagemakergeospatial.GetEarthObservationJobInput{}
        output := &sagemakergeospatial.GetEarthObservationJobOutput{}

        mockClient.On("GetEarthObservationJob", ctx, input).Return(output, nil)

        result, err := mockClient.GetEarthObservationJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetRasterDataCollection", func(t *testing.T) {
        input := &sagemakergeospatial.GetRasterDataCollectionInput{}
        output := &sagemakergeospatial.GetRasterDataCollectionOutput{}

        mockClient.On("GetRasterDataCollection", ctx, input).Return(output, nil)

        result, err := mockClient.GetRasterDataCollection(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetTile", func(t *testing.T) {
        input := &sagemakergeospatial.GetTileInput{}
        output := &sagemakergeospatial.GetTileOutput{}

        mockClient.On("GetTile", ctx, input).Return(output, nil)

        result, err := mockClient.GetTile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetVectorEnrichmentJob", func(t *testing.T) {
        input := &sagemakergeospatial.GetVectorEnrichmentJobInput{}
        output := &sagemakergeospatial.GetVectorEnrichmentJobOutput{}

        mockClient.On("GetVectorEnrichmentJob", ctx, input).Return(output, nil)

        result, err := mockClient.GetVectorEnrichmentJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListEarthObservationJobs", func(t *testing.T) {
        input := &sagemakergeospatial.ListEarthObservationJobsInput{}
        output := &sagemakergeospatial.ListEarthObservationJobsOutput{}

        mockClient.On("ListEarthObservationJobs", ctx, input).Return(output, nil)

        result, err := mockClient.ListEarthObservationJobs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListRasterDataCollections", func(t *testing.T) {
        input := &sagemakergeospatial.ListRasterDataCollectionsInput{}
        output := &sagemakergeospatial.ListRasterDataCollectionsOutput{}

        mockClient.On("ListRasterDataCollections", ctx, input).Return(output, nil)

        result, err := mockClient.ListRasterDataCollections(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &sagemakergeospatial.ListTagsForResourceInput{}
        output := &sagemakergeospatial.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListVectorEnrichmentJobs", func(t *testing.T) {
        input := &sagemakergeospatial.ListVectorEnrichmentJobsInput{}
        output := &sagemakergeospatial.ListVectorEnrichmentJobsOutput{}

        mockClient.On("ListVectorEnrichmentJobs", ctx, input).Return(output, nil)

        result, err := mockClient.ListVectorEnrichmentJobs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSearchRasterDataCollection", func(t *testing.T) {
        input := &sagemakergeospatial.SearchRasterDataCollectionInput{}
        output := &sagemakergeospatial.SearchRasterDataCollectionOutput{}

        mockClient.On("SearchRasterDataCollection", ctx, input).Return(output, nil)

        result, err := mockClient.SearchRasterDataCollection(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartEarthObservationJob", func(t *testing.T) {
        input := &sagemakergeospatial.StartEarthObservationJobInput{}
        output := &sagemakergeospatial.StartEarthObservationJobOutput{}

        mockClient.On("StartEarthObservationJob", ctx, input).Return(output, nil)

        result, err := mockClient.StartEarthObservationJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartVectorEnrichmentJob", func(t *testing.T) {
        input := &sagemakergeospatial.StartVectorEnrichmentJobInput{}
        output := &sagemakergeospatial.StartVectorEnrichmentJobOutput{}

        mockClient.On("StartVectorEnrichmentJob", ctx, input).Return(output, nil)

        result, err := mockClient.StartVectorEnrichmentJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopEarthObservationJob", func(t *testing.T) {
        input := &sagemakergeospatial.StopEarthObservationJobInput{}
        output := &sagemakergeospatial.StopEarthObservationJobOutput{}

        mockClient.On("StopEarthObservationJob", ctx, input).Return(output, nil)

        result, err := mockClient.StopEarthObservationJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopVectorEnrichmentJob", func(t *testing.T) {
        input := &sagemakergeospatial.StopVectorEnrichmentJobInput{}
        output := &sagemakergeospatial.StopVectorEnrichmentJobOutput{}

        mockClient.On("StopVectorEnrichmentJob", ctx, input).Return(output, nil)

        result, err := mockClient.StopVectorEnrichmentJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &sagemakergeospatial.TagResourceInput{}
        output := &sagemakergeospatial.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &sagemakergeospatial.UntagResourceInput{}
        output := &sagemakergeospatial.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
