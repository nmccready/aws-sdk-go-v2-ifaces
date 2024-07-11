
package evidently

import (
    "github.com/aws/aws-sdk-go-v2/service/evidently"
)

// IClient defines the interface for evidently
type IClient interface {
 Options() Options 
 BatchEvaluateFeature(ctx context.Context, params *BatchEvaluateFeatureInput, optFns ...func(*Options)) (*BatchEvaluateFeatureOutput, error) 
 CreateExperiment(ctx context.Context, params *CreateExperimentInput, optFns ...func(*Options)) (*CreateExperimentOutput, error) 
 CreateFeature(ctx context.Context, params *CreateFeatureInput, optFns ...func(*Options)) (*CreateFeatureOutput, error) 
 CreateLaunch(ctx context.Context, params *CreateLaunchInput, optFns ...func(*Options)) (*CreateLaunchOutput, error) 
 CreateProject(ctx context.Context, params *CreateProjectInput, optFns ...func(*Options)) (*CreateProjectOutput, error) 
 CreateSegment(ctx context.Context, params *CreateSegmentInput, optFns ...func(*Options)) (*CreateSegmentOutput, error) 
 DeleteExperiment(ctx context.Context, params *DeleteExperimentInput, optFns ...func(*Options)) (*DeleteExperimentOutput, error) 
 DeleteFeature(ctx context.Context, params *DeleteFeatureInput, optFns ...func(*Options)) (*DeleteFeatureOutput, error) 
 DeleteLaunch(ctx context.Context, params *DeleteLaunchInput, optFns ...func(*Options)) (*DeleteLaunchOutput, error) 
 DeleteProject(ctx context.Context, params *DeleteProjectInput, optFns ...func(*Options)) (*DeleteProjectOutput, error) 
 DeleteSegment(ctx context.Context, params *DeleteSegmentInput, optFns ...func(*Options)) (*DeleteSegmentOutput, error) 
 EvaluateFeature(ctx context.Context, params *EvaluateFeatureInput, optFns ...func(*Options)) (*EvaluateFeatureOutput, error) 
 GetExperiment(ctx context.Context, params *GetExperimentInput, optFns ...func(*Options)) (*GetExperimentOutput, error) 
 GetExperimentResults(ctx context.Context, params *GetExperimentResultsInput, optFns ...func(*Options)) (*GetExperimentResultsOutput, error) 
 GetFeature(ctx context.Context, params *GetFeatureInput, optFns ...func(*Options)) (*GetFeatureOutput, error) 
 GetLaunch(ctx context.Context, params *GetLaunchInput, optFns ...func(*Options)) (*GetLaunchOutput, error) 
 GetProject(ctx context.Context, params *GetProjectInput, optFns ...func(*Options)) (*GetProjectOutput, error) 
 GetSegment(ctx context.Context, params *GetSegmentInput, optFns ...func(*Options)) (*GetSegmentOutput, error) 
 ListExperiments(ctx context.Context, params *ListExperimentsInput, optFns ...func(*Options)) (*ListExperimentsOutput, error) 
 ListFeatures(ctx context.Context, params *ListFeaturesInput, optFns ...func(*Options)) (*ListFeaturesOutput, error) 
 ListLaunches(ctx context.Context, params *ListLaunchesInput, optFns ...func(*Options)) (*ListLaunchesOutput, error) 
 ListProjects(ctx context.Context, params *ListProjectsInput, optFns ...func(*Options)) (*ListProjectsOutput, error) 
 ListSegmentReferences(ctx context.Context, params *ListSegmentReferencesInput, optFns ...func(*Options)) (*ListSegmentReferencesOutput, error) 
 ListSegments(ctx context.Context, params *ListSegmentsInput, optFns ...func(*Options)) (*ListSegmentsOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 PutProjectEvents(ctx context.Context, params *PutProjectEventsInput, optFns ...func(*Options)) (*PutProjectEventsOutput, error) 
 StartExperiment(ctx context.Context, params *StartExperimentInput, optFns ...func(*Options)) (*StartExperimentOutput, error) 
 StartLaunch(ctx context.Context, params *StartLaunchInput, optFns ...func(*Options)) (*StartLaunchOutput, error) 
 StopExperiment(ctx context.Context, params *StopExperimentInput, optFns ...func(*Options)) (*StopExperimentOutput, error) 
 StopLaunch(ctx context.Context, params *StopLaunchInput, optFns ...func(*Options)) (*StopLaunchOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 TestSegmentPattern(ctx context.Context, params *TestSegmentPatternInput, optFns ...func(*Options)) (*TestSegmentPatternOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateExperiment(ctx context.Context, params *UpdateExperimentInput, optFns ...func(*Options)) (*UpdateExperimentOutput, error) 
 UpdateFeature(ctx context.Context, params *UpdateFeatureInput, optFns ...func(*Options)) (*UpdateFeatureOutput, error) 
 UpdateLaunch(ctx context.Context, params *UpdateLaunchInput, optFns ...func(*Options)) (*UpdateLaunchOutput, error) 
 UpdateProject(ctx context.Context, params *UpdateProjectInput, optFns ...func(*Options)) (*UpdateProjectOutput, error) 
 UpdateProjectDataDelivery(ctx context.Context, params *UpdateProjectDataDeliveryInput, optFns ...func(*Options)) (*UpdateProjectDataDeliveryOutput, error) 
}
