
package rekognition

import (
    "github.com/aws/aws-sdk-go-v2/service/rekognition"
)

// IRekognition defines the interface for rekognition
type IRekognition interface {
 Options() Options 
 AssociateFaces(ctx context.Context, params *AssociateFacesInput, optFns ...func(*Options)) (*AssociateFacesOutput, error) 
 CompareFaces(ctx context.Context, params *CompareFacesInput, optFns ...func(*Options)) (*CompareFacesOutput, error) 
 CopyProjectVersion(ctx context.Context, params *CopyProjectVersionInput, optFns ...func(*Options)) (*CopyProjectVersionOutput, error) 
 CreateCollection(ctx context.Context, params *CreateCollectionInput, optFns ...func(*Options)) (*CreateCollectionOutput, error) 
 CreateDataset(ctx context.Context, params *CreateDatasetInput, optFns ...func(*Options)) (*CreateDatasetOutput, error) 
 CreateFaceLivenessSession(ctx context.Context, params *CreateFaceLivenessSessionInput, optFns ...func(*Options)) (*CreateFaceLivenessSessionOutput, error) 
 CreateProject(ctx context.Context, params *CreateProjectInput, optFns ...func(*Options)) (*CreateProjectOutput, error) 
 CreateProjectVersion(ctx context.Context, params *CreateProjectVersionInput, optFns ...func(*Options)) (*CreateProjectVersionOutput, error) 
 CreateStreamProcessor(ctx context.Context, params *CreateStreamProcessorInput, optFns ...func(*Options)) (*CreateStreamProcessorOutput, error) 
 CreateUser(ctx context.Context, params *CreateUserInput, optFns ...func(*Options)) (*CreateUserOutput, error) 
 DeleteCollection(ctx context.Context, params *DeleteCollectionInput, optFns ...func(*Options)) (*DeleteCollectionOutput, error) 
 DeleteDataset(ctx context.Context, params *DeleteDatasetInput, optFns ...func(*Options)) (*DeleteDatasetOutput, error) 
 DeleteFaces(ctx context.Context, params *DeleteFacesInput, optFns ...func(*Options)) (*DeleteFacesOutput, error) 
 DeleteProject(ctx context.Context, params *DeleteProjectInput, optFns ...func(*Options)) (*DeleteProjectOutput, error) 
 DeleteProjectPolicy(ctx context.Context, params *DeleteProjectPolicyInput, optFns ...func(*Options)) (*DeleteProjectPolicyOutput, error) 
 DeleteProjectVersion(ctx context.Context, params *DeleteProjectVersionInput, optFns ...func(*Options)) (*DeleteProjectVersionOutput, error) 
 DeleteStreamProcessor(ctx context.Context, params *DeleteStreamProcessorInput, optFns ...func(*Options)) (*DeleteStreamProcessorOutput, error) 
 DeleteUser(ctx context.Context, params *DeleteUserInput, optFns ...func(*Options)) (*DeleteUserOutput, error) 
 DescribeCollection(ctx context.Context, params *DescribeCollectionInput, optFns ...func(*Options)) (*DescribeCollectionOutput, error) 
 DescribeDataset(ctx context.Context, params *DescribeDatasetInput, optFns ...func(*Options)) (*DescribeDatasetOutput, error) 
 DescribeProjectVersions(ctx context.Context, params *DescribeProjectVersionsInput, optFns ...func(*Options)) (*DescribeProjectVersionsOutput, error) 
 DescribeProjects(ctx context.Context, params *DescribeProjectsInput, optFns ...func(*Options)) (*DescribeProjectsOutput, error) 
 DescribeStreamProcessor(ctx context.Context, params *DescribeStreamProcessorInput, optFns ...func(*Options)) (*DescribeStreamProcessorOutput, error) 
 DetectCustomLabels(ctx context.Context, params *DetectCustomLabelsInput, optFns ...func(*Options)) (*DetectCustomLabelsOutput, error) 
 DetectFaces(ctx context.Context, params *DetectFacesInput, optFns ...func(*Options)) (*DetectFacesOutput, error) 
 DetectLabels(ctx context.Context, params *DetectLabelsInput, optFns ...func(*Options)) (*DetectLabelsOutput, error) 
 DetectModerationLabels(ctx context.Context, params *DetectModerationLabelsInput, optFns ...func(*Options)) (*DetectModerationLabelsOutput, error) 
 DetectProtectiveEquipment(ctx context.Context, params *DetectProtectiveEquipmentInput, optFns ...func(*Options)) (*DetectProtectiveEquipmentOutput, error) 
 DetectText(ctx context.Context, params *DetectTextInput, optFns ...func(*Options)) (*DetectTextOutput, error) 
 DisassociateFaces(ctx context.Context, params *DisassociateFacesInput, optFns ...func(*Options)) (*DisassociateFacesOutput, error) 
 DistributeDatasetEntries(ctx context.Context, params *DistributeDatasetEntriesInput, optFns ...func(*Options)) (*DistributeDatasetEntriesOutput, error) 
 GetCelebrityInfo(ctx context.Context, params *GetCelebrityInfoInput, optFns ...func(*Options)) (*GetCelebrityInfoOutput, error) 
 GetCelebrityRecognition(ctx context.Context, params *GetCelebrityRecognitionInput, optFns ...func(*Options)) (*GetCelebrityRecognitionOutput, error) 
 GetContentModeration(ctx context.Context, params *GetContentModerationInput, optFns ...func(*Options)) (*GetContentModerationOutput, error) 
 GetFaceDetection(ctx context.Context, params *GetFaceDetectionInput, optFns ...func(*Options)) (*GetFaceDetectionOutput, error) 
 GetFaceLivenessSessionResults(ctx context.Context, params *GetFaceLivenessSessionResultsInput, optFns ...func(*Options)) (*GetFaceLivenessSessionResultsOutput, error) 
 GetFaceSearch(ctx context.Context, params *GetFaceSearchInput, optFns ...func(*Options)) (*GetFaceSearchOutput, error) 
 GetLabelDetection(ctx context.Context, params *GetLabelDetectionInput, optFns ...func(*Options)) (*GetLabelDetectionOutput, error) 
 GetMediaAnalysisJob(ctx context.Context, params *GetMediaAnalysisJobInput, optFns ...func(*Options)) (*GetMediaAnalysisJobOutput, error) 
 GetPersonTracking(ctx context.Context, params *GetPersonTrackingInput, optFns ...func(*Options)) (*GetPersonTrackingOutput, error) 
 GetSegmentDetection(ctx context.Context, params *GetSegmentDetectionInput, optFns ...func(*Options)) (*GetSegmentDetectionOutput, error) 
 GetTextDetection(ctx context.Context, params *GetTextDetectionInput, optFns ...func(*Options)) (*GetTextDetectionOutput, error) 
 IndexFaces(ctx context.Context, params *IndexFacesInput, optFns ...func(*Options)) (*IndexFacesOutput, error) 
 ListCollections(ctx context.Context, params *ListCollectionsInput, optFns ...func(*Options)) (*ListCollectionsOutput, error) 
 ListDatasetEntries(ctx context.Context, params *ListDatasetEntriesInput, optFns ...func(*Options)) (*ListDatasetEntriesOutput, error) 
 ListDatasetLabels(ctx context.Context, params *ListDatasetLabelsInput, optFns ...func(*Options)) (*ListDatasetLabelsOutput, error) 
 ListFaces(ctx context.Context, params *ListFacesInput, optFns ...func(*Options)) (*ListFacesOutput, error) 
 ListMediaAnalysisJobs(ctx context.Context, params *ListMediaAnalysisJobsInput, optFns ...func(*Options)) (*ListMediaAnalysisJobsOutput, error) 
 ListProjectPolicies(ctx context.Context, params *ListProjectPoliciesInput, optFns ...func(*Options)) (*ListProjectPoliciesOutput, error) 
 ListStreamProcessors(ctx context.Context, params *ListStreamProcessorsInput, optFns ...func(*Options)) (*ListStreamProcessorsOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 ListUsers(ctx context.Context, params *ListUsersInput, optFns ...func(*Options)) (*ListUsersOutput, error) 
 PutProjectPolicy(ctx context.Context, params *PutProjectPolicyInput, optFns ...func(*Options)) (*PutProjectPolicyOutput, error) 
 RecognizeCelebrities(ctx context.Context, params *RecognizeCelebritiesInput, optFns ...func(*Options)) (*RecognizeCelebritiesOutput, error) 
 SearchFaces(ctx context.Context, params *SearchFacesInput, optFns ...func(*Options)) (*SearchFacesOutput, error) 
 SearchFacesByImage(ctx context.Context, params *SearchFacesByImageInput, optFns ...func(*Options)) (*SearchFacesByImageOutput, error) 
 SearchUsers(ctx context.Context, params *SearchUsersInput, optFns ...func(*Options)) (*SearchUsersOutput, error) 
 SearchUsersByImage(ctx context.Context, params *SearchUsersByImageInput, optFns ...func(*Options)) (*SearchUsersByImageOutput, error) 
 StartCelebrityRecognition(ctx context.Context, params *StartCelebrityRecognitionInput, optFns ...func(*Options)) (*StartCelebrityRecognitionOutput, error) 
 StartContentModeration(ctx context.Context, params *StartContentModerationInput, optFns ...func(*Options)) (*StartContentModerationOutput, error) 
 StartFaceDetection(ctx context.Context, params *StartFaceDetectionInput, optFns ...func(*Options)) (*StartFaceDetectionOutput, error) 
 StartFaceSearch(ctx context.Context, params *StartFaceSearchInput, optFns ...func(*Options)) (*StartFaceSearchOutput, error) 
 StartLabelDetection(ctx context.Context, params *StartLabelDetectionInput, optFns ...func(*Options)) (*StartLabelDetectionOutput, error) 
 StartMediaAnalysisJob(ctx context.Context, params *StartMediaAnalysisJobInput, optFns ...func(*Options)) (*StartMediaAnalysisJobOutput, error) 
 StartPersonTracking(ctx context.Context, params *StartPersonTrackingInput, optFns ...func(*Options)) (*StartPersonTrackingOutput, error) 
 StartProjectVersion(ctx context.Context, params *StartProjectVersionInput, optFns ...func(*Options)) (*StartProjectVersionOutput, error) 
 StartSegmentDetection(ctx context.Context, params *StartSegmentDetectionInput, optFns ...func(*Options)) (*StartSegmentDetectionOutput, error) 
 StartStreamProcessor(ctx context.Context, params *StartStreamProcessorInput, optFns ...func(*Options)) (*StartStreamProcessorOutput, error) 
 StartTextDetection(ctx context.Context, params *StartTextDetectionInput, optFns ...func(*Options)) (*StartTextDetectionOutput, error) 
 StopProjectVersion(ctx context.Context, params *StopProjectVersionInput, optFns ...func(*Options)) (*StopProjectVersionOutput, error) 
 StopStreamProcessor(ctx context.Context, params *StopStreamProcessorInput, optFns ...func(*Options)) (*StopStreamProcessorOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateDatasetEntries(ctx context.Context, params *UpdateDatasetEntriesInput, optFns ...func(*Options)) (*UpdateDatasetEntriesOutput, error) 
 UpdateStreamProcessor(ctx context.Context, params *UpdateStreamProcessorInput, optFns ...func(*Options)) (*UpdateStreamProcessorOutput, error) 
}
