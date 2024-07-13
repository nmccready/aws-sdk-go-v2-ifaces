
package location_iface

// This file was generated by scripts/gen_ifaces.sh do not edit it manually
import (
    "context"
    . "github.com/aws/aws-sdk-go-v2/service/location"
)

// IClient defines the interface for location
type IClient interface {
 Options() Options 
 AssociateTrackerConsumer(ctx context.Context, params *AssociateTrackerConsumerInput, optFns ...func(*Options)) (*AssociateTrackerConsumerOutput, error) 
 BatchDeleteDevicePositionHistory(ctx context.Context, params *BatchDeleteDevicePositionHistoryInput, optFns ...func(*Options)) (*BatchDeleteDevicePositionHistoryOutput, error) 
 BatchDeleteGeofence(ctx context.Context, params *BatchDeleteGeofenceInput, optFns ...func(*Options)) (*BatchDeleteGeofenceOutput, error) 
 BatchEvaluateGeofences(ctx context.Context, params *BatchEvaluateGeofencesInput, optFns ...func(*Options)) (*BatchEvaluateGeofencesOutput, error) 
 BatchGetDevicePosition(ctx context.Context, params *BatchGetDevicePositionInput, optFns ...func(*Options)) (*BatchGetDevicePositionOutput, error) 
 BatchPutGeofence(ctx context.Context, params *BatchPutGeofenceInput, optFns ...func(*Options)) (*BatchPutGeofenceOutput, error) 
 BatchUpdateDevicePosition(ctx context.Context, params *BatchUpdateDevicePositionInput, optFns ...func(*Options)) (*BatchUpdateDevicePositionOutput, error) 
 CalculateRoute(ctx context.Context, params *CalculateRouteInput, optFns ...func(*Options)) (*CalculateRouteOutput, error) 
 CalculateRouteMatrix(ctx context.Context, params *CalculateRouteMatrixInput, optFns ...func(*Options)) (*CalculateRouteMatrixOutput, error) 
 CreateGeofenceCollection(ctx context.Context, params *CreateGeofenceCollectionInput, optFns ...func(*Options)) (*CreateGeofenceCollectionOutput, error) 
 CreateKey(ctx context.Context, params *CreateKeyInput, optFns ...func(*Options)) (*CreateKeyOutput, error) 
 CreateMap(ctx context.Context, params *CreateMapInput, optFns ...func(*Options)) (*CreateMapOutput, error) 
 CreatePlaceIndex(ctx context.Context, params *CreatePlaceIndexInput, optFns ...func(*Options)) (*CreatePlaceIndexOutput, error) 
 CreateRouteCalculator(ctx context.Context, params *CreateRouteCalculatorInput, optFns ...func(*Options)) (*CreateRouteCalculatorOutput, error) 
 CreateTracker(ctx context.Context, params *CreateTrackerInput, optFns ...func(*Options)) (*CreateTrackerOutput, error) 
 DeleteGeofenceCollection(ctx context.Context, params *DeleteGeofenceCollectionInput, optFns ...func(*Options)) (*DeleteGeofenceCollectionOutput, error) 
 DeleteKey(ctx context.Context, params *DeleteKeyInput, optFns ...func(*Options)) (*DeleteKeyOutput, error) 
 DeleteMap(ctx context.Context, params *DeleteMapInput, optFns ...func(*Options)) (*DeleteMapOutput, error) 
 DeletePlaceIndex(ctx context.Context, params *DeletePlaceIndexInput, optFns ...func(*Options)) (*DeletePlaceIndexOutput, error) 
 DeleteRouteCalculator(ctx context.Context, params *DeleteRouteCalculatorInput, optFns ...func(*Options)) (*DeleteRouteCalculatorOutput, error) 
 DeleteTracker(ctx context.Context, params *DeleteTrackerInput, optFns ...func(*Options)) (*DeleteTrackerOutput, error) 
 DescribeGeofenceCollection(ctx context.Context, params *DescribeGeofenceCollectionInput, optFns ...func(*Options)) (*DescribeGeofenceCollectionOutput, error) 
 DescribeKey(ctx context.Context, params *DescribeKeyInput, optFns ...func(*Options)) (*DescribeKeyOutput, error) 
 DescribeMap(ctx context.Context, params *DescribeMapInput, optFns ...func(*Options)) (*DescribeMapOutput, error) 
 DescribePlaceIndex(ctx context.Context, params *DescribePlaceIndexInput, optFns ...func(*Options)) (*DescribePlaceIndexOutput, error) 
 DescribeRouteCalculator(ctx context.Context, params *DescribeRouteCalculatorInput, optFns ...func(*Options)) (*DescribeRouteCalculatorOutput, error) 
 DescribeTracker(ctx context.Context, params *DescribeTrackerInput, optFns ...func(*Options)) (*DescribeTrackerOutput, error) 
 DisassociateTrackerConsumer(ctx context.Context, params *DisassociateTrackerConsumerInput, optFns ...func(*Options)) (*DisassociateTrackerConsumerOutput, error) 
 ForecastGeofenceEvents(ctx context.Context, params *ForecastGeofenceEventsInput, optFns ...func(*Options)) (*ForecastGeofenceEventsOutput, error) 
 GetDevicePosition(ctx context.Context, params *GetDevicePositionInput, optFns ...func(*Options)) (*GetDevicePositionOutput, error) 
 GetDevicePositionHistory(ctx context.Context, params *GetDevicePositionHistoryInput, optFns ...func(*Options)) (*GetDevicePositionHistoryOutput, error) 
 GetGeofence(ctx context.Context, params *GetGeofenceInput, optFns ...func(*Options)) (*GetGeofenceOutput, error) 
 GetMapGlyphs(ctx context.Context, params *GetMapGlyphsInput, optFns ...func(*Options)) (*GetMapGlyphsOutput, error) 
 GetMapSprites(ctx context.Context, params *GetMapSpritesInput, optFns ...func(*Options)) (*GetMapSpritesOutput, error) 
 GetMapStyleDescriptor(ctx context.Context, params *GetMapStyleDescriptorInput, optFns ...func(*Options)) (*GetMapStyleDescriptorOutput, error) 
 GetMapTile(ctx context.Context, params *GetMapTileInput, optFns ...func(*Options)) (*GetMapTileOutput, error) 
 GetPlace(ctx context.Context, params *GetPlaceInput, optFns ...func(*Options)) (*GetPlaceOutput, error) 
 ListDevicePositions(ctx context.Context, params *ListDevicePositionsInput, optFns ...func(*Options)) (*ListDevicePositionsOutput, error) 
 ListGeofenceCollections(ctx context.Context, params *ListGeofenceCollectionsInput, optFns ...func(*Options)) (*ListGeofenceCollectionsOutput, error) 
 ListGeofences(ctx context.Context, params *ListGeofencesInput, optFns ...func(*Options)) (*ListGeofencesOutput, error) 
 ListKeys(ctx context.Context, params *ListKeysInput, optFns ...func(*Options)) (*ListKeysOutput, error) 
 ListMaps(ctx context.Context, params *ListMapsInput, optFns ...func(*Options)) (*ListMapsOutput, error) 
 ListPlaceIndexes(ctx context.Context, params *ListPlaceIndexesInput, optFns ...func(*Options)) (*ListPlaceIndexesOutput, error) 
 ListRouteCalculators(ctx context.Context, params *ListRouteCalculatorsInput, optFns ...func(*Options)) (*ListRouteCalculatorsOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 ListTrackerConsumers(ctx context.Context, params *ListTrackerConsumersInput, optFns ...func(*Options)) (*ListTrackerConsumersOutput, error) 
 ListTrackers(ctx context.Context, params *ListTrackersInput, optFns ...func(*Options)) (*ListTrackersOutput, error) 
 PutGeofence(ctx context.Context, params *PutGeofenceInput, optFns ...func(*Options)) (*PutGeofenceOutput, error) 
 SearchPlaceIndexForPosition(ctx context.Context, params *SearchPlaceIndexForPositionInput, optFns ...func(*Options)) (*SearchPlaceIndexForPositionOutput, error) 
 SearchPlaceIndexForSuggestions(ctx context.Context, params *SearchPlaceIndexForSuggestionsInput, optFns ...func(*Options)) (*SearchPlaceIndexForSuggestionsOutput, error) 
 SearchPlaceIndexForText(ctx context.Context, params *SearchPlaceIndexForTextInput, optFns ...func(*Options)) (*SearchPlaceIndexForTextOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateGeofenceCollection(ctx context.Context, params *UpdateGeofenceCollectionInput, optFns ...func(*Options)) (*UpdateGeofenceCollectionOutput, error) 
 UpdateKey(ctx context.Context, params *UpdateKeyInput, optFns ...func(*Options)) (*UpdateKeyOutput, error) 
 UpdateMap(ctx context.Context, params *UpdateMapInput, optFns ...func(*Options)) (*UpdateMapOutput, error) 
 UpdatePlaceIndex(ctx context.Context, params *UpdatePlaceIndexInput, optFns ...func(*Options)) (*UpdatePlaceIndexOutput, error) 
 UpdateRouteCalculator(ctx context.Context, params *UpdateRouteCalculatorInput, optFns ...func(*Options)) (*UpdateRouteCalculatorOutput, error) 
 UpdateTracker(ctx context.Context, params *UpdateTrackerInput, optFns ...func(*Options)) (*UpdateTrackerOutput, error) 
 VerifyDevicePosition(ctx context.Context, params *VerifyDevicePositionInput, optFns ...func(*Options)) (*VerifyDevicePositionOutput, error) 
}