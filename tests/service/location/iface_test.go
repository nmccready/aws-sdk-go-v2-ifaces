package location_test

// tests for the location service interface for this ../../../service/location/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/location"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/location/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/location/location_iface"
	"github.com/stretchr/testify/assert"
)

func TestLocationServiceCanBeMocked(t *testing.T) {
	var iface location_iface.IClient
	iface = &location.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := location.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateTrackerConsumer", func(t *testing.T) {
        input := &location.AssociateTrackerConsumerInput{}
        output := &location.AssociateTrackerConsumerOutput{}

        mockClient.On("AssociateTrackerConsumer", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateTrackerConsumer(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchDeleteDevicePositionHistory", func(t *testing.T) {
        input := &location.BatchDeleteDevicePositionHistoryInput{}
        output := &location.BatchDeleteDevicePositionHistoryOutput{}

        mockClient.On("BatchDeleteDevicePositionHistory", ctx, input).Return(output, nil)

        result, err := mockClient.BatchDeleteDevicePositionHistory(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchDeleteGeofence", func(t *testing.T) {
        input := &location.BatchDeleteGeofenceInput{}
        output := &location.BatchDeleteGeofenceOutput{}

        mockClient.On("BatchDeleteGeofence", ctx, input).Return(output, nil)

        result, err := mockClient.BatchDeleteGeofence(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchEvaluateGeofences", func(t *testing.T) {
        input := &location.BatchEvaluateGeofencesInput{}
        output := &location.BatchEvaluateGeofencesOutput{}

        mockClient.On("BatchEvaluateGeofences", ctx, input).Return(output, nil)

        result, err := mockClient.BatchEvaluateGeofences(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchGetDevicePosition", func(t *testing.T) {
        input := &location.BatchGetDevicePositionInput{}
        output := &location.BatchGetDevicePositionOutput{}

        mockClient.On("BatchGetDevicePosition", ctx, input).Return(output, nil)

        result, err := mockClient.BatchGetDevicePosition(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchPutGeofence", func(t *testing.T) {
        input := &location.BatchPutGeofenceInput{}
        output := &location.BatchPutGeofenceOutput{}

        mockClient.On("BatchPutGeofence", ctx, input).Return(output, nil)

        result, err := mockClient.BatchPutGeofence(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchUpdateDevicePosition", func(t *testing.T) {
        input := &location.BatchUpdateDevicePositionInput{}
        output := &location.BatchUpdateDevicePositionOutput{}

        mockClient.On("BatchUpdateDevicePosition", ctx, input).Return(output, nil)

        result, err := mockClient.BatchUpdateDevicePosition(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCalculateRoute", func(t *testing.T) {
        input := &location.CalculateRouteInput{}
        output := &location.CalculateRouteOutput{}

        mockClient.On("CalculateRoute", ctx, input).Return(output, nil)

        result, err := mockClient.CalculateRoute(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCalculateRouteMatrix", func(t *testing.T) {
        input := &location.CalculateRouteMatrixInput{}
        output := &location.CalculateRouteMatrixOutput{}

        mockClient.On("CalculateRouteMatrix", ctx, input).Return(output, nil)

        result, err := mockClient.CalculateRouteMatrix(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateGeofenceCollection", func(t *testing.T) {
        input := &location.CreateGeofenceCollectionInput{}
        output := &location.CreateGeofenceCollectionOutput{}

        mockClient.On("CreateGeofenceCollection", ctx, input).Return(output, nil)

        result, err := mockClient.CreateGeofenceCollection(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateKey", func(t *testing.T) {
        input := &location.CreateKeyInput{}
        output := &location.CreateKeyOutput{}

        mockClient.On("CreateKey", ctx, input).Return(output, nil)

        result, err := mockClient.CreateKey(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateMap", func(t *testing.T) {
        input := &location.CreateMapInput{}
        output := &location.CreateMapOutput{}

        mockClient.On("CreateMap", ctx, input).Return(output, nil)

        result, err := mockClient.CreateMap(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreatePlaceIndex", func(t *testing.T) {
        input := &location.CreatePlaceIndexInput{}
        output := &location.CreatePlaceIndexOutput{}

        mockClient.On("CreatePlaceIndex", ctx, input).Return(output, nil)

        result, err := mockClient.CreatePlaceIndex(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateRouteCalculator", func(t *testing.T) {
        input := &location.CreateRouteCalculatorInput{}
        output := &location.CreateRouteCalculatorOutput{}

        mockClient.On("CreateRouteCalculator", ctx, input).Return(output, nil)

        result, err := mockClient.CreateRouteCalculator(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateTracker", func(t *testing.T) {
        input := &location.CreateTrackerInput{}
        output := &location.CreateTrackerOutput{}

        mockClient.On("CreateTracker", ctx, input).Return(output, nil)

        result, err := mockClient.CreateTracker(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteGeofenceCollection", func(t *testing.T) {
        input := &location.DeleteGeofenceCollectionInput{}
        output := &location.DeleteGeofenceCollectionOutput{}

        mockClient.On("DeleteGeofenceCollection", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteGeofenceCollection(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteKey", func(t *testing.T) {
        input := &location.DeleteKeyInput{}
        output := &location.DeleteKeyOutput{}

        mockClient.On("DeleteKey", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteKey(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteMap", func(t *testing.T) {
        input := &location.DeleteMapInput{}
        output := &location.DeleteMapOutput{}

        mockClient.On("DeleteMap", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteMap(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeletePlaceIndex", func(t *testing.T) {
        input := &location.DeletePlaceIndexInput{}
        output := &location.DeletePlaceIndexOutput{}

        mockClient.On("DeletePlaceIndex", ctx, input).Return(output, nil)

        result, err := mockClient.DeletePlaceIndex(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteRouteCalculator", func(t *testing.T) {
        input := &location.DeleteRouteCalculatorInput{}
        output := &location.DeleteRouteCalculatorOutput{}

        mockClient.On("DeleteRouteCalculator", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteRouteCalculator(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteTracker", func(t *testing.T) {
        input := &location.DeleteTrackerInput{}
        output := &location.DeleteTrackerOutput{}

        mockClient.On("DeleteTracker", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteTracker(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeGeofenceCollection", func(t *testing.T) {
        input := &location.DescribeGeofenceCollectionInput{}
        output := &location.DescribeGeofenceCollectionOutput{}

        mockClient.On("DescribeGeofenceCollection", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeGeofenceCollection(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeKey", func(t *testing.T) {
        input := &location.DescribeKeyInput{}
        output := &location.DescribeKeyOutput{}

        mockClient.On("DescribeKey", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeKey(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeMap", func(t *testing.T) {
        input := &location.DescribeMapInput{}
        output := &location.DescribeMapOutput{}

        mockClient.On("DescribeMap", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeMap(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribePlaceIndex", func(t *testing.T) {
        input := &location.DescribePlaceIndexInput{}
        output := &location.DescribePlaceIndexOutput{}

        mockClient.On("DescribePlaceIndex", ctx, input).Return(output, nil)

        result, err := mockClient.DescribePlaceIndex(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeRouteCalculator", func(t *testing.T) {
        input := &location.DescribeRouteCalculatorInput{}
        output := &location.DescribeRouteCalculatorOutput{}

        mockClient.On("DescribeRouteCalculator", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeRouteCalculator(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeTracker", func(t *testing.T) {
        input := &location.DescribeTrackerInput{}
        output := &location.DescribeTrackerOutput{}

        mockClient.On("DescribeTracker", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeTracker(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateTrackerConsumer", func(t *testing.T) {
        input := &location.DisassociateTrackerConsumerInput{}
        output := &location.DisassociateTrackerConsumerOutput{}

        mockClient.On("DisassociateTrackerConsumer", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateTrackerConsumer(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestForecastGeofenceEvents", func(t *testing.T) {
        input := &location.ForecastGeofenceEventsInput{}
        output := &location.ForecastGeofenceEventsOutput{}

        mockClient.On("ForecastGeofenceEvents", ctx, input).Return(output, nil)

        result, err := mockClient.ForecastGeofenceEvents(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDevicePosition", func(t *testing.T) {
        input := &location.GetDevicePositionInput{}
        output := &location.GetDevicePositionOutput{}

        mockClient.On("GetDevicePosition", ctx, input).Return(output, nil)

        result, err := mockClient.GetDevicePosition(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDevicePositionHistory", func(t *testing.T) {
        input := &location.GetDevicePositionHistoryInput{}
        output := &location.GetDevicePositionHistoryOutput{}

        mockClient.On("GetDevicePositionHistory", ctx, input).Return(output, nil)

        result, err := mockClient.GetDevicePositionHistory(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetGeofence", func(t *testing.T) {
        input := &location.GetGeofenceInput{}
        output := &location.GetGeofenceOutput{}

        mockClient.On("GetGeofence", ctx, input).Return(output, nil)

        result, err := mockClient.GetGeofence(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetMapGlyphs", func(t *testing.T) {
        input := &location.GetMapGlyphsInput{}
        output := &location.GetMapGlyphsOutput{}

        mockClient.On("GetMapGlyphs", ctx, input).Return(output, nil)

        result, err := mockClient.GetMapGlyphs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetMapSprites", func(t *testing.T) {
        input := &location.GetMapSpritesInput{}
        output := &location.GetMapSpritesOutput{}

        mockClient.On("GetMapSprites", ctx, input).Return(output, nil)

        result, err := mockClient.GetMapSprites(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetMapStyleDescriptor", func(t *testing.T) {
        input := &location.GetMapStyleDescriptorInput{}
        output := &location.GetMapStyleDescriptorOutput{}

        mockClient.On("GetMapStyleDescriptor", ctx, input).Return(output, nil)

        result, err := mockClient.GetMapStyleDescriptor(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetMapTile", func(t *testing.T) {
        input := &location.GetMapTileInput{}
        output := &location.GetMapTileOutput{}

        mockClient.On("GetMapTile", ctx, input).Return(output, nil)

        result, err := mockClient.GetMapTile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetPlace", func(t *testing.T) {
        input := &location.GetPlaceInput{}
        output := &location.GetPlaceOutput{}

        mockClient.On("GetPlace", ctx, input).Return(output, nil)

        result, err := mockClient.GetPlace(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDevicePositions", func(t *testing.T) {
        input := &location.ListDevicePositionsInput{}
        output := &location.ListDevicePositionsOutput{}

        mockClient.On("ListDevicePositions", ctx, input).Return(output, nil)

        result, err := mockClient.ListDevicePositions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListGeofenceCollections", func(t *testing.T) {
        input := &location.ListGeofenceCollectionsInput{}
        output := &location.ListGeofenceCollectionsOutput{}

        mockClient.On("ListGeofenceCollections", ctx, input).Return(output, nil)

        result, err := mockClient.ListGeofenceCollections(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListGeofences", func(t *testing.T) {
        input := &location.ListGeofencesInput{}
        output := &location.ListGeofencesOutput{}

        mockClient.On("ListGeofences", ctx, input).Return(output, nil)

        result, err := mockClient.ListGeofences(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListKeys", func(t *testing.T) {
        input := &location.ListKeysInput{}
        output := &location.ListKeysOutput{}

        mockClient.On("ListKeys", ctx, input).Return(output, nil)

        result, err := mockClient.ListKeys(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListMaps", func(t *testing.T) {
        input := &location.ListMapsInput{}
        output := &location.ListMapsOutput{}

        mockClient.On("ListMaps", ctx, input).Return(output, nil)

        result, err := mockClient.ListMaps(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListPlaceIndexes", func(t *testing.T) {
        input := &location.ListPlaceIndexesInput{}
        output := &location.ListPlaceIndexesOutput{}

        mockClient.On("ListPlaceIndexes", ctx, input).Return(output, nil)

        result, err := mockClient.ListPlaceIndexes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListRouteCalculators", func(t *testing.T) {
        input := &location.ListRouteCalculatorsInput{}
        output := &location.ListRouteCalculatorsOutput{}

        mockClient.On("ListRouteCalculators", ctx, input).Return(output, nil)

        result, err := mockClient.ListRouteCalculators(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &location.ListTagsForResourceInput{}
        output := &location.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTrackerConsumers", func(t *testing.T) {
        input := &location.ListTrackerConsumersInput{}
        output := &location.ListTrackerConsumersOutput{}

        mockClient.On("ListTrackerConsumers", ctx, input).Return(output, nil)

        result, err := mockClient.ListTrackerConsumers(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTrackers", func(t *testing.T) {
        input := &location.ListTrackersInput{}
        output := &location.ListTrackersOutput{}

        mockClient.On("ListTrackers", ctx, input).Return(output, nil)

        result, err := mockClient.ListTrackers(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutGeofence", func(t *testing.T) {
        input := &location.PutGeofenceInput{}
        output := &location.PutGeofenceOutput{}

        mockClient.On("PutGeofence", ctx, input).Return(output, nil)

        result, err := mockClient.PutGeofence(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSearchPlaceIndexForPosition", func(t *testing.T) {
        input := &location.SearchPlaceIndexForPositionInput{}
        output := &location.SearchPlaceIndexForPositionOutput{}

        mockClient.On("SearchPlaceIndexForPosition", ctx, input).Return(output, nil)

        result, err := mockClient.SearchPlaceIndexForPosition(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSearchPlaceIndexForSuggestions", func(t *testing.T) {
        input := &location.SearchPlaceIndexForSuggestionsInput{}
        output := &location.SearchPlaceIndexForSuggestionsOutput{}

        mockClient.On("SearchPlaceIndexForSuggestions", ctx, input).Return(output, nil)

        result, err := mockClient.SearchPlaceIndexForSuggestions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSearchPlaceIndexForText", func(t *testing.T) {
        input := &location.SearchPlaceIndexForTextInput{}
        output := &location.SearchPlaceIndexForTextOutput{}

        mockClient.On("SearchPlaceIndexForText", ctx, input).Return(output, nil)

        result, err := mockClient.SearchPlaceIndexForText(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &location.TagResourceInput{}
        output := &location.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &location.UntagResourceInput{}
        output := &location.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateGeofenceCollection", func(t *testing.T) {
        input := &location.UpdateGeofenceCollectionInput{}
        output := &location.UpdateGeofenceCollectionOutput{}

        mockClient.On("UpdateGeofenceCollection", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateGeofenceCollection(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateKey", func(t *testing.T) {
        input := &location.UpdateKeyInput{}
        output := &location.UpdateKeyOutput{}

        mockClient.On("UpdateKey", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateKey(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateMap", func(t *testing.T) {
        input := &location.UpdateMapInput{}
        output := &location.UpdateMapOutput{}

        mockClient.On("UpdateMap", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateMap(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdatePlaceIndex", func(t *testing.T) {
        input := &location.UpdatePlaceIndexInput{}
        output := &location.UpdatePlaceIndexOutput{}

        mockClient.On("UpdatePlaceIndex", ctx, input).Return(output, nil)

        result, err := mockClient.UpdatePlaceIndex(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateRouteCalculator", func(t *testing.T) {
        input := &location.UpdateRouteCalculatorInput{}
        output := &location.UpdateRouteCalculatorOutput{}

        mockClient.On("UpdateRouteCalculator", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateRouteCalculator(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateTracker", func(t *testing.T) {
        input := &location.UpdateTrackerInput{}
        output := &location.UpdateTrackerOutput{}

        mockClient.On("UpdateTracker", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateTracker(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestVerifyDevicePosition", func(t *testing.T) {
        input := &location.VerifyDevicePositionInput{}
        output := &location.VerifyDevicePositionOutput{}

        mockClient.On("VerifyDevicePosition", ctx, input).Return(output, nil)

        result, err := mockClient.VerifyDevicePosition(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
