
package sagemakergeospatial

import (
    "github.com/aws/aws-sdk-go-v2/service/sagemakergeospatial"
)

// IClient defines the interface for sagemakergeospatial
type IClient interface {
 Options() Options 
 DeleteEarthObservationJob(ctx context.Context, params *DeleteEarthObservationJobInput, optFns ...func(*Options)) (*DeleteEarthObservationJobOutput, error) 
 DeleteVectorEnrichmentJob(ctx context.Context, params *DeleteVectorEnrichmentJobInput, optFns ...func(*Options)) (*DeleteVectorEnrichmentJobOutput, error) 
 ExportEarthObservationJob(ctx context.Context, params *ExportEarthObservationJobInput, optFns ...func(*Options)) (*ExportEarthObservationJobOutput, error) 
 ExportVectorEnrichmentJob(ctx context.Context, params *ExportVectorEnrichmentJobInput, optFns ...func(*Options)) (*ExportVectorEnrichmentJobOutput, error) 
 GetEarthObservationJob(ctx context.Context, params *GetEarthObservationJobInput, optFns ...func(*Options)) (*GetEarthObservationJobOutput, error) 
 GetRasterDataCollection(ctx context.Context, params *GetRasterDataCollectionInput, optFns ...func(*Options)) (*GetRasterDataCollectionOutput, error) 
 GetTile(ctx context.Context, params *GetTileInput, optFns ...func(*Options)) (*GetTileOutput, error) 
 GetVectorEnrichmentJob(ctx context.Context, params *GetVectorEnrichmentJobInput, optFns ...func(*Options)) (*GetVectorEnrichmentJobOutput, error) 
 ListEarthObservationJobs(ctx context.Context, params *ListEarthObservationJobsInput, optFns ...func(*Options)) (*ListEarthObservationJobsOutput, error) 
 ListRasterDataCollections(ctx context.Context, params *ListRasterDataCollectionsInput, optFns ...func(*Options)) (*ListRasterDataCollectionsOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 ListVectorEnrichmentJobs(ctx context.Context, params *ListVectorEnrichmentJobsInput, optFns ...func(*Options)) (*ListVectorEnrichmentJobsOutput, error) 
 SearchRasterDataCollection(ctx context.Context, params *SearchRasterDataCollectionInput, optFns ...func(*Options)) (*SearchRasterDataCollectionOutput, error) 
 StartEarthObservationJob(ctx context.Context, params *StartEarthObservationJobInput, optFns ...func(*Options)) (*StartEarthObservationJobOutput, error) 
 StartVectorEnrichmentJob(ctx context.Context, params *StartVectorEnrichmentJobInput, optFns ...func(*Options)) (*StartVectorEnrichmentJobOutput, error) 
 StopEarthObservationJob(ctx context.Context, params *StopEarthObservationJobInput, optFns ...func(*Options)) (*StopEarthObservationJobOutput, error) 
 StopVectorEnrichmentJob(ctx context.Context, params *StopVectorEnrichmentJobInput, optFns ...func(*Options)) (*StopVectorEnrichmentJobOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
}
