
package comprehend_iface

// This file was generated by scripts/gen_ifaces.sh do not edit it manually
import (
    "context"
    . "github.com/aws/aws-sdk-go-v2/service/comprehend"
)

// IClient defines the interface for comprehend
type IClient interface {
 Options() Options 
 BatchDetectDominantLanguage(ctx context.Context, params *BatchDetectDominantLanguageInput, optFns ...func(*Options)) (*BatchDetectDominantLanguageOutput, error) 
 BatchDetectEntities(ctx context.Context, params *BatchDetectEntitiesInput, optFns ...func(*Options)) (*BatchDetectEntitiesOutput, error) 
 BatchDetectKeyPhrases(ctx context.Context, params *BatchDetectKeyPhrasesInput, optFns ...func(*Options)) (*BatchDetectKeyPhrasesOutput, error) 
 BatchDetectSentiment(ctx context.Context, params *BatchDetectSentimentInput, optFns ...func(*Options)) (*BatchDetectSentimentOutput, error) 
 BatchDetectSyntax(ctx context.Context, params *BatchDetectSyntaxInput, optFns ...func(*Options)) (*BatchDetectSyntaxOutput, error) 
 BatchDetectTargetedSentiment(ctx context.Context, params *BatchDetectTargetedSentimentInput, optFns ...func(*Options)) (*BatchDetectTargetedSentimentOutput, error) 
 ClassifyDocument(ctx context.Context, params *ClassifyDocumentInput, optFns ...func(*Options)) (*ClassifyDocumentOutput, error) 
 ContainsPiiEntities(ctx context.Context, params *ContainsPiiEntitiesInput, optFns ...func(*Options)) (*ContainsPiiEntitiesOutput, error) 
 CreateDataset(ctx context.Context, params *CreateDatasetInput, optFns ...func(*Options)) (*CreateDatasetOutput, error) 
 CreateDocumentClassifier(ctx context.Context, params *CreateDocumentClassifierInput, optFns ...func(*Options)) (*CreateDocumentClassifierOutput, error) 
 CreateEndpoint(ctx context.Context, params *CreateEndpointInput, optFns ...func(*Options)) (*CreateEndpointOutput, error) 
 CreateEntityRecognizer(ctx context.Context, params *CreateEntityRecognizerInput, optFns ...func(*Options)) (*CreateEntityRecognizerOutput, error) 
 CreateFlywheel(ctx context.Context, params *CreateFlywheelInput, optFns ...func(*Options)) (*CreateFlywheelOutput, error) 
 DeleteDocumentClassifier(ctx context.Context, params *DeleteDocumentClassifierInput, optFns ...func(*Options)) (*DeleteDocumentClassifierOutput, error) 
 DeleteEndpoint(ctx context.Context, params *DeleteEndpointInput, optFns ...func(*Options)) (*DeleteEndpointOutput, error) 
 DeleteEntityRecognizer(ctx context.Context, params *DeleteEntityRecognizerInput, optFns ...func(*Options)) (*DeleteEntityRecognizerOutput, error) 
 DeleteFlywheel(ctx context.Context, params *DeleteFlywheelInput, optFns ...func(*Options)) (*DeleteFlywheelOutput, error) 
 DeleteResourcePolicy(ctx context.Context, params *DeleteResourcePolicyInput, optFns ...func(*Options)) (*DeleteResourcePolicyOutput, error) 
 DescribeDataset(ctx context.Context, params *DescribeDatasetInput, optFns ...func(*Options)) (*DescribeDatasetOutput, error) 
 DescribeDocumentClassificationJob(ctx context.Context, params *DescribeDocumentClassificationJobInput, optFns ...func(*Options)) (*DescribeDocumentClassificationJobOutput, error) 
 DescribeDocumentClassifier(ctx context.Context, params *DescribeDocumentClassifierInput, optFns ...func(*Options)) (*DescribeDocumentClassifierOutput, error) 
 DescribeDominantLanguageDetectionJob(ctx context.Context, params *DescribeDominantLanguageDetectionJobInput, optFns ...func(*Options)) (*DescribeDominantLanguageDetectionJobOutput, error) 
 DescribeEndpoint(ctx context.Context, params *DescribeEndpointInput, optFns ...func(*Options)) (*DescribeEndpointOutput, error) 
 DescribeEntitiesDetectionJob(ctx context.Context, params *DescribeEntitiesDetectionJobInput, optFns ...func(*Options)) (*DescribeEntitiesDetectionJobOutput, error) 
 DescribeEntityRecognizer(ctx context.Context, params *DescribeEntityRecognizerInput, optFns ...func(*Options)) (*DescribeEntityRecognizerOutput, error) 
 DescribeEventsDetectionJob(ctx context.Context, params *DescribeEventsDetectionJobInput, optFns ...func(*Options)) (*DescribeEventsDetectionJobOutput, error) 
 DescribeFlywheel(ctx context.Context, params *DescribeFlywheelInput, optFns ...func(*Options)) (*DescribeFlywheelOutput, error) 
 DescribeFlywheelIteration(ctx context.Context, params *DescribeFlywheelIterationInput, optFns ...func(*Options)) (*DescribeFlywheelIterationOutput, error) 
 DescribeKeyPhrasesDetectionJob(ctx context.Context, params *DescribeKeyPhrasesDetectionJobInput, optFns ...func(*Options)) (*DescribeKeyPhrasesDetectionJobOutput, error) 
 DescribePiiEntitiesDetectionJob(ctx context.Context, params *DescribePiiEntitiesDetectionJobInput, optFns ...func(*Options)) (*DescribePiiEntitiesDetectionJobOutput, error) 
 DescribeResourcePolicy(ctx context.Context, params *DescribeResourcePolicyInput, optFns ...func(*Options)) (*DescribeResourcePolicyOutput, error) 
 DescribeSentimentDetectionJob(ctx context.Context, params *DescribeSentimentDetectionJobInput, optFns ...func(*Options)) (*DescribeSentimentDetectionJobOutput, error) 
 DescribeTargetedSentimentDetectionJob(ctx context.Context, params *DescribeTargetedSentimentDetectionJobInput, optFns ...func(*Options)) (*DescribeTargetedSentimentDetectionJobOutput, error) 
 DescribeTopicsDetectionJob(ctx context.Context, params *DescribeTopicsDetectionJobInput, optFns ...func(*Options)) (*DescribeTopicsDetectionJobOutput, error) 
 DetectDominantLanguage(ctx context.Context, params *DetectDominantLanguageInput, optFns ...func(*Options)) (*DetectDominantLanguageOutput, error) 
 DetectEntities(ctx context.Context, params *DetectEntitiesInput, optFns ...func(*Options)) (*DetectEntitiesOutput, error) 
 DetectKeyPhrases(ctx context.Context, params *DetectKeyPhrasesInput, optFns ...func(*Options)) (*DetectKeyPhrasesOutput, error) 
 DetectPiiEntities(ctx context.Context, params *DetectPiiEntitiesInput, optFns ...func(*Options)) (*DetectPiiEntitiesOutput, error) 
 DetectSentiment(ctx context.Context, params *DetectSentimentInput, optFns ...func(*Options)) (*DetectSentimentOutput, error) 
 DetectSyntax(ctx context.Context, params *DetectSyntaxInput, optFns ...func(*Options)) (*DetectSyntaxOutput, error) 
 DetectTargetedSentiment(ctx context.Context, params *DetectTargetedSentimentInput, optFns ...func(*Options)) (*DetectTargetedSentimentOutput, error) 
 DetectToxicContent(ctx context.Context, params *DetectToxicContentInput, optFns ...func(*Options)) (*DetectToxicContentOutput, error) 
 ImportModel(ctx context.Context, params *ImportModelInput, optFns ...func(*Options)) (*ImportModelOutput, error) 
 ListDatasets(ctx context.Context, params *ListDatasetsInput, optFns ...func(*Options)) (*ListDatasetsOutput, error) 
 ListDocumentClassificationJobs(ctx context.Context, params *ListDocumentClassificationJobsInput, optFns ...func(*Options)) (*ListDocumentClassificationJobsOutput, error) 
 ListDocumentClassifierSummaries(ctx context.Context, params *ListDocumentClassifierSummariesInput, optFns ...func(*Options)) (*ListDocumentClassifierSummariesOutput, error) 
 ListDocumentClassifiers(ctx context.Context, params *ListDocumentClassifiersInput, optFns ...func(*Options)) (*ListDocumentClassifiersOutput, error) 
 ListDominantLanguageDetectionJobs(ctx context.Context, params *ListDominantLanguageDetectionJobsInput, optFns ...func(*Options)) (*ListDominantLanguageDetectionJobsOutput, error) 
 ListEndpoints(ctx context.Context, params *ListEndpointsInput, optFns ...func(*Options)) (*ListEndpointsOutput, error) 
 ListEntitiesDetectionJobs(ctx context.Context, params *ListEntitiesDetectionJobsInput, optFns ...func(*Options)) (*ListEntitiesDetectionJobsOutput, error) 
 ListEntityRecognizerSummaries(ctx context.Context, params *ListEntityRecognizerSummariesInput, optFns ...func(*Options)) (*ListEntityRecognizerSummariesOutput, error) 
 ListEntityRecognizers(ctx context.Context, params *ListEntityRecognizersInput, optFns ...func(*Options)) (*ListEntityRecognizersOutput, error) 
 ListEventsDetectionJobs(ctx context.Context, params *ListEventsDetectionJobsInput, optFns ...func(*Options)) (*ListEventsDetectionJobsOutput, error) 
 ListFlywheelIterationHistory(ctx context.Context, params *ListFlywheelIterationHistoryInput, optFns ...func(*Options)) (*ListFlywheelIterationHistoryOutput, error) 
 ListFlywheels(ctx context.Context, params *ListFlywheelsInput, optFns ...func(*Options)) (*ListFlywheelsOutput, error) 
 ListKeyPhrasesDetectionJobs(ctx context.Context, params *ListKeyPhrasesDetectionJobsInput, optFns ...func(*Options)) (*ListKeyPhrasesDetectionJobsOutput, error) 
 ListPiiEntitiesDetectionJobs(ctx context.Context, params *ListPiiEntitiesDetectionJobsInput, optFns ...func(*Options)) (*ListPiiEntitiesDetectionJobsOutput, error) 
 ListSentimentDetectionJobs(ctx context.Context, params *ListSentimentDetectionJobsInput, optFns ...func(*Options)) (*ListSentimentDetectionJobsOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 ListTargetedSentimentDetectionJobs(ctx context.Context, params *ListTargetedSentimentDetectionJobsInput, optFns ...func(*Options)) (*ListTargetedSentimentDetectionJobsOutput, error) 
 ListTopicsDetectionJobs(ctx context.Context, params *ListTopicsDetectionJobsInput, optFns ...func(*Options)) (*ListTopicsDetectionJobsOutput, error) 
 PutResourcePolicy(ctx context.Context, params *PutResourcePolicyInput, optFns ...func(*Options)) (*PutResourcePolicyOutput, error) 
 StartDocumentClassificationJob(ctx context.Context, params *StartDocumentClassificationJobInput, optFns ...func(*Options)) (*StartDocumentClassificationJobOutput, error) 
 StartDominantLanguageDetectionJob(ctx context.Context, params *StartDominantLanguageDetectionJobInput, optFns ...func(*Options)) (*StartDominantLanguageDetectionJobOutput, error) 
 StartEntitiesDetectionJob(ctx context.Context, params *StartEntitiesDetectionJobInput, optFns ...func(*Options)) (*StartEntitiesDetectionJobOutput, error) 
 StartEventsDetectionJob(ctx context.Context, params *StartEventsDetectionJobInput, optFns ...func(*Options)) (*StartEventsDetectionJobOutput, error) 
 StartFlywheelIteration(ctx context.Context, params *StartFlywheelIterationInput, optFns ...func(*Options)) (*StartFlywheelIterationOutput, error) 
 StartKeyPhrasesDetectionJob(ctx context.Context, params *StartKeyPhrasesDetectionJobInput, optFns ...func(*Options)) (*StartKeyPhrasesDetectionJobOutput, error) 
 StartPiiEntitiesDetectionJob(ctx context.Context, params *StartPiiEntitiesDetectionJobInput, optFns ...func(*Options)) (*StartPiiEntitiesDetectionJobOutput, error) 
 StartSentimentDetectionJob(ctx context.Context, params *StartSentimentDetectionJobInput, optFns ...func(*Options)) (*StartSentimentDetectionJobOutput, error) 
 StartTargetedSentimentDetectionJob(ctx context.Context, params *StartTargetedSentimentDetectionJobInput, optFns ...func(*Options)) (*StartTargetedSentimentDetectionJobOutput, error) 
 StartTopicsDetectionJob(ctx context.Context, params *StartTopicsDetectionJobInput, optFns ...func(*Options)) (*StartTopicsDetectionJobOutput, error) 
 StopDominantLanguageDetectionJob(ctx context.Context, params *StopDominantLanguageDetectionJobInput, optFns ...func(*Options)) (*StopDominantLanguageDetectionJobOutput, error) 
 StopEntitiesDetectionJob(ctx context.Context, params *StopEntitiesDetectionJobInput, optFns ...func(*Options)) (*StopEntitiesDetectionJobOutput, error) 
 StopEventsDetectionJob(ctx context.Context, params *StopEventsDetectionJobInput, optFns ...func(*Options)) (*StopEventsDetectionJobOutput, error) 
 StopKeyPhrasesDetectionJob(ctx context.Context, params *StopKeyPhrasesDetectionJobInput, optFns ...func(*Options)) (*StopKeyPhrasesDetectionJobOutput, error) 
 StopPiiEntitiesDetectionJob(ctx context.Context, params *StopPiiEntitiesDetectionJobInput, optFns ...func(*Options)) (*StopPiiEntitiesDetectionJobOutput, error) 
 StopSentimentDetectionJob(ctx context.Context, params *StopSentimentDetectionJobInput, optFns ...func(*Options)) (*StopSentimentDetectionJobOutput, error) 
 StopTargetedSentimentDetectionJob(ctx context.Context, params *StopTargetedSentimentDetectionJobInput, optFns ...func(*Options)) (*StopTargetedSentimentDetectionJobOutput, error) 
 StopTrainingDocumentClassifier(ctx context.Context, params *StopTrainingDocumentClassifierInput, optFns ...func(*Options)) (*StopTrainingDocumentClassifierOutput, error) 
 StopTrainingEntityRecognizer(ctx context.Context, params *StopTrainingEntityRecognizerInput, optFns ...func(*Options)) (*StopTrainingEntityRecognizerOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateEndpoint(ctx context.Context, params *UpdateEndpointInput, optFns ...func(*Options)) (*UpdateEndpointOutput, error) 
 UpdateFlywheel(ctx context.Context, params *UpdateFlywheelInput, optFns ...func(*Options)) (*UpdateFlywheelOutput, error) 
}
