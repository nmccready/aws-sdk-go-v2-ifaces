package rekognition_test

// tests for the rekognition service interface for this ../../../service/rekognition/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/rekognition"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/rekognition/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/rekognition/rekognition_iface"
	"github.com/stretchr/testify/assert"
)

func TestRekognitionServiceCanBeMocked(t *testing.T) {
	var iface rekognition_iface.IClient
	iface = &rekognition.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := rekognition.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateFaces", func(t *testing.T) {
        input := &rekognition.AssociateFacesInput{}
        output := &rekognition.AssociateFacesOutput{}

        mockClient.On("AssociateFaces", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateFaces(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCompareFaces", func(t *testing.T) {
        input := &rekognition.CompareFacesInput{}
        output := &rekognition.CompareFacesOutput{}

        mockClient.On("CompareFaces", ctx, input).Return(output, nil)

        result, err := mockClient.CompareFaces(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCopyProjectVersion", func(t *testing.T) {
        input := &rekognition.CopyProjectVersionInput{}
        output := &rekognition.CopyProjectVersionOutput{}

        mockClient.On("CopyProjectVersion", ctx, input).Return(output, nil)

        result, err := mockClient.CopyProjectVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateCollection", func(t *testing.T) {
        input := &rekognition.CreateCollectionInput{}
        output := &rekognition.CreateCollectionOutput{}

        mockClient.On("CreateCollection", ctx, input).Return(output, nil)

        result, err := mockClient.CreateCollection(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateDataset", func(t *testing.T) {
        input := &rekognition.CreateDatasetInput{}
        output := &rekognition.CreateDatasetOutput{}

        mockClient.On("CreateDataset", ctx, input).Return(output, nil)

        result, err := mockClient.CreateDataset(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateFaceLivenessSession", func(t *testing.T) {
        input := &rekognition.CreateFaceLivenessSessionInput{}
        output := &rekognition.CreateFaceLivenessSessionOutput{}

        mockClient.On("CreateFaceLivenessSession", ctx, input).Return(output, nil)

        result, err := mockClient.CreateFaceLivenessSession(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateProject", func(t *testing.T) {
        input := &rekognition.CreateProjectInput{}
        output := &rekognition.CreateProjectOutput{}

        mockClient.On("CreateProject", ctx, input).Return(output, nil)

        result, err := mockClient.CreateProject(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateProjectVersion", func(t *testing.T) {
        input := &rekognition.CreateProjectVersionInput{}
        output := &rekognition.CreateProjectVersionOutput{}

        mockClient.On("CreateProjectVersion", ctx, input).Return(output, nil)

        result, err := mockClient.CreateProjectVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateStreamProcessor", func(t *testing.T) {
        input := &rekognition.CreateStreamProcessorInput{}
        output := &rekognition.CreateStreamProcessorOutput{}

        mockClient.On("CreateStreamProcessor", ctx, input).Return(output, nil)

        result, err := mockClient.CreateStreamProcessor(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateUser", func(t *testing.T) {
        input := &rekognition.CreateUserInput{}
        output := &rekognition.CreateUserOutput{}

        mockClient.On("CreateUser", ctx, input).Return(output, nil)

        result, err := mockClient.CreateUser(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteCollection", func(t *testing.T) {
        input := &rekognition.DeleteCollectionInput{}
        output := &rekognition.DeleteCollectionOutput{}

        mockClient.On("DeleteCollection", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteCollection(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteDataset", func(t *testing.T) {
        input := &rekognition.DeleteDatasetInput{}
        output := &rekognition.DeleteDatasetOutput{}

        mockClient.On("DeleteDataset", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteDataset(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteFaces", func(t *testing.T) {
        input := &rekognition.DeleteFacesInput{}
        output := &rekognition.DeleteFacesOutput{}

        mockClient.On("DeleteFaces", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteFaces(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteProject", func(t *testing.T) {
        input := &rekognition.DeleteProjectInput{}
        output := &rekognition.DeleteProjectOutput{}

        mockClient.On("DeleteProject", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteProject(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteProjectPolicy", func(t *testing.T) {
        input := &rekognition.DeleteProjectPolicyInput{}
        output := &rekognition.DeleteProjectPolicyOutput{}

        mockClient.On("DeleteProjectPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteProjectPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteProjectVersion", func(t *testing.T) {
        input := &rekognition.DeleteProjectVersionInput{}
        output := &rekognition.DeleteProjectVersionOutput{}

        mockClient.On("DeleteProjectVersion", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteProjectVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteStreamProcessor", func(t *testing.T) {
        input := &rekognition.DeleteStreamProcessorInput{}
        output := &rekognition.DeleteStreamProcessorOutput{}

        mockClient.On("DeleteStreamProcessor", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteStreamProcessor(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteUser", func(t *testing.T) {
        input := &rekognition.DeleteUserInput{}
        output := &rekognition.DeleteUserOutput{}

        mockClient.On("DeleteUser", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteUser(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeCollection", func(t *testing.T) {
        input := &rekognition.DescribeCollectionInput{}
        output := &rekognition.DescribeCollectionOutput{}

        mockClient.On("DescribeCollection", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeCollection(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeDataset", func(t *testing.T) {
        input := &rekognition.DescribeDatasetInput{}
        output := &rekognition.DescribeDatasetOutput{}

        mockClient.On("DescribeDataset", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeDataset(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeProjectVersions", func(t *testing.T) {
        input := &rekognition.DescribeProjectVersionsInput{}
        output := &rekognition.DescribeProjectVersionsOutput{}

        mockClient.On("DescribeProjectVersions", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeProjectVersions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeProjects", func(t *testing.T) {
        input := &rekognition.DescribeProjectsInput{}
        output := &rekognition.DescribeProjectsOutput{}

        mockClient.On("DescribeProjects", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeProjects(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeStreamProcessor", func(t *testing.T) {
        input := &rekognition.DescribeStreamProcessorInput{}
        output := &rekognition.DescribeStreamProcessorOutput{}

        mockClient.On("DescribeStreamProcessor", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeStreamProcessor(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDetectCustomLabels", func(t *testing.T) {
        input := &rekognition.DetectCustomLabelsInput{}
        output := &rekognition.DetectCustomLabelsOutput{}

        mockClient.On("DetectCustomLabels", ctx, input).Return(output, nil)

        result, err := mockClient.DetectCustomLabels(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDetectFaces", func(t *testing.T) {
        input := &rekognition.DetectFacesInput{}
        output := &rekognition.DetectFacesOutput{}

        mockClient.On("DetectFaces", ctx, input).Return(output, nil)

        result, err := mockClient.DetectFaces(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDetectLabels", func(t *testing.T) {
        input := &rekognition.DetectLabelsInput{}
        output := &rekognition.DetectLabelsOutput{}

        mockClient.On("DetectLabels", ctx, input).Return(output, nil)

        result, err := mockClient.DetectLabels(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDetectModerationLabels", func(t *testing.T) {
        input := &rekognition.DetectModerationLabelsInput{}
        output := &rekognition.DetectModerationLabelsOutput{}

        mockClient.On("DetectModerationLabels", ctx, input).Return(output, nil)

        result, err := mockClient.DetectModerationLabels(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDetectProtectiveEquipment", func(t *testing.T) {
        input := &rekognition.DetectProtectiveEquipmentInput{}
        output := &rekognition.DetectProtectiveEquipmentOutput{}

        mockClient.On("DetectProtectiveEquipment", ctx, input).Return(output, nil)

        result, err := mockClient.DetectProtectiveEquipment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDetectText", func(t *testing.T) {
        input := &rekognition.DetectTextInput{}
        output := &rekognition.DetectTextOutput{}

        mockClient.On("DetectText", ctx, input).Return(output, nil)

        result, err := mockClient.DetectText(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateFaces", func(t *testing.T) {
        input := &rekognition.DisassociateFacesInput{}
        output := &rekognition.DisassociateFacesOutput{}

        mockClient.On("DisassociateFaces", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateFaces(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDistributeDatasetEntries", func(t *testing.T) {
        input := &rekognition.DistributeDatasetEntriesInput{}
        output := &rekognition.DistributeDatasetEntriesOutput{}

        mockClient.On("DistributeDatasetEntries", ctx, input).Return(output, nil)

        result, err := mockClient.DistributeDatasetEntries(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetCelebrityInfo", func(t *testing.T) {
        input := &rekognition.GetCelebrityInfoInput{}
        output := &rekognition.GetCelebrityInfoOutput{}

        mockClient.On("GetCelebrityInfo", ctx, input).Return(output, nil)

        result, err := mockClient.GetCelebrityInfo(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetCelebrityRecognition", func(t *testing.T) {
        input := &rekognition.GetCelebrityRecognitionInput{}
        output := &rekognition.GetCelebrityRecognitionOutput{}

        mockClient.On("GetCelebrityRecognition", ctx, input).Return(output, nil)

        result, err := mockClient.GetCelebrityRecognition(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetContentModeration", func(t *testing.T) {
        input := &rekognition.GetContentModerationInput{}
        output := &rekognition.GetContentModerationOutput{}

        mockClient.On("GetContentModeration", ctx, input).Return(output, nil)

        result, err := mockClient.GetContentModeration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetFaceDetection", func(t *testing.T) {
        input := &rekognition.GetFaceDetectionInput{}
        output := &rekognition.GetFaceDetectionOutput{}

        mockClient.On("GetFaceDetection", ctx, input).Return(output, nil)

        result, err := mockClient.GetFaceDetection(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetFaceLivenessSessionResults", func(t *testing.T) {
        input := &rekognition.GetFaceLivenessSessionResultsInput{}
        output := &rekognition.GetFaceLivenessSessionResultsOutput{}

        mockClient.On("GetFaceLivenessSessionResults", ctx, input).Return(output, nil)

        result, err := mockClient.GetFaceLivenessSessionResults(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetFaceSearch", func(t *testing.T) {
        input := &rekognition.GetFaceSearchInput{}
        output := &rekognition.GetFaceSearchOutput{}

        mockClient.On("GetFaceSearch", ctx, input).Return(output, nil)

        result, err := mockClient.GetFaceSearch(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetLabelDetection", func(t *testing.T) {
        input := &rekognition.GetLabelDetectionInput{}
        output := &rekognition.GetLabelDetectionOutput{}

        mockClient.On("GetLabelDetection", ctx, input).Return(output, nil)

        result, err := mockClient.GetLabelDetection(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetMediaAnalysisJob", func(t *testing.T) {
        input := &rekognition.GetMediaAnalysisJobInput{}
        output := &rekognition.GetMediaAnalysisJobOutput{}

        mockClient.On("GetMediaAnalysisJob", ctx, input).Return(output, nil)

        result, err := mockClient.GetMediaAnalysisJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetPersonTracking", func(t *testing.T) {
        input := &rekognition.GetPersonTrackingInput{}
        output := &rekognition.GetPersonTrackingOutput{}

        mockClient.On("GetPersonTracking", ctx, input).Return(output, nil)

        result, err := mockClient.GetPersonTracking(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSegmentDetection", func(t *testing.T) {
        input := &rekognition.GetSegmentDetectionInput{}
        output := &rekognition.GetSegmentDetectionOutput{}

        mockClient.On("GetSegmentDetection", ctx, input).Return(output, nil)

        result, err := mockClient.GetSegmentDetection(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetTextDetection", func(t *testing.T) {
        input := &rekognition.GetTextDetectionInput{}
        output := &rekognition.GetTextDetectionOutput{}

        mockClient.On("GetTextDetection", ctx, input).Return(output, nil)

        result, err := mockClient.GetTextDetection(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestIndexFaces", func(t *testing.T) {
        input := &rekognition.IndexFacesInput{}
        output := &rekognition.IndexFacesOutput{}

        mockClient.On("IndexFaces", ctx, input).Return(output, nil)

        result, err := mockClient.IndexFaces(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListCollections", func(t *testing.T) {
        input := &rekognition.ListCollectionsInput{}
        output := &rekognition.ListCollectionsOutput{}

        mockClient.On("ListCollections", ctx, input).Return(output, nil)

        result, err := mockClient.ListCollections(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDatasetEntries", func(t *testing.T) {
        input := &rekognition.ListDatasetEntriesInput{}
        output := &rekognition.ListDatasetEntriesOutput{}

        mockClient.On("ListDatasetEntries", ctx, input).Return(output, nil)

        result, err := mockClient.ListDatasetEntries(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDatasetLabels", func(t *testing.T) {
        input := &rekognition.ListDatasetLabelsInput{}
        output := &rekognition.ListDatasetLabelsOutput{}

        mockClient.On("ListDatasetLabels", ctx, input).Return(output, nil)

        result, err := mockClient.ListDatasetLabels(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListFaces", func(t *testing.T) {
        input := &rekognition.ListFacesInput{}
        output := &rekognition.ListFacesOutput{}

        mockClient.On("ListFaces", ctx, input).Return(output, nil)

        result, err := mockClient.ListFaces(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListMediaAnalysisJobs", func(t *testing.T) {
        input := &rekognition.ListMediaAnalysisJobsInput{}
        output := &rekognition.ListMediaAnalysisJobsOutput{}

        mockClient.On("ListMediaAnalysisJobs", ctx, input).Return(output, nil)

        result, err := mockClient.ListMediaAnalysisJobs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListProjectPolicies", func(t *testing.T) {
        input := &rekognition.ListProjectPoliciesInput{}
        output := &rekognition.ListProjectPoliciesOutput{}

        mockClient.On("ListProjectPolicies", ctx, input).Return(output, nil)

        result, err := mockClient.ListProjectPolicies(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListStreamProcessors", func(t *testing.T) {
        input := &rekognition.ListStreamProcessorsInput{}
        output := &rekognition.ListStreamProcessorsOutput{}

        mockClient.On("ListStreamProcessors", ctx, input).Return(output, nil)

        result, err := mockClient.ListStreamProcessors(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &rekognition.ListTagsForResourceInput{}
        output := &rekognition.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListUsers", func(t *testing.T) {
        input := &rekognition.ListUsersInput{}
        output := &rekognition.ListUsersOutput{}

        mockClient.On("ListUsers", ctx, input).Return(output, nil)

        result, err := mockClient.ListUsers(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutProjectPolicy", func(t *testing.T) {
        input := &rekognition.PutProjectPolicyInput{}
        output := &rekognition.PutProjectPolicyOutput{}

        mockClient.On("PutProjectPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.PutProjectPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRecognizeCelebrities", func(t *testing.T) {
        input := &rekognition.RecognizeCelebritiesInput{}
        output := &rekognition.RecognizeCelebritiesOutput{}

        mockClient.On("RecognizeCelebrities", ctx, input).Return(output, nil)

        result, err := mockClient.RecognizeCelebrities(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSearchFaces", func(t *testing.T) {
        input := &rekognition.SearchFacesInput{}
        output := &rekognition.SearchFacesOutput{}

        mockClient.On("SearchFaces", ctx, input).Return(output, nil)

        result, err := mockClient.SearchFaces(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSearchFacesByImage", func(t *testing.T) {
        input := &rekognition.SearchFacesByImageInput{}
        output := &rekognition.SearchFacesByImageOutput{}

        mockClient.On("SearchFacesByImage", ctx, input).Return(output, nil)

        result, err := mockClient.SearchFacesByImage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSearchUsers", func(t *testing.T) {
        input := &rekognition.SearchUsersInput{}
        output := &rekognition.SearchUsersOutput{}

        mockClient.On("SearchUsers", ctx, input).Return(output, nil)

        result, err := mockClient.SearchUsers(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSearchUsersByImage", func(t *testing.T) {
        input := &rekognition.SearchUsersByImageInput{}
        output := &rekognition.SearchUsersByImageOutput{}

        mockClient.On("SearchUsersByImage", ctx, input).Return(output, nil)

        result, err := mockClient.SearchUsersByImage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartCelebrityRecognition", func(t *testing.T) {
        input := &rekognition.StartCelebrityRecognitionInput{}
        output := &rekognition.StartCelebrityRecognitionOutput{}

        mockClient.On("StartCelebrityRecognition", ctx, input).Return(output, nil)

        result, err := mockClient.StartCelebrityRecognition(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartContentModeration", func(t *testing.T) {
        input := &rekognition.StartContentModerationInput{}
        output := &rekognition.StartContentModerationOutput{}

        mockClient.On("StartContentModeration", ctx, input).Return(output, nil)

        result, err := mockClient.StartContentModeration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartFaceDetection", func(t *testing.T) {
        input := &rekognition.StartFaceDetectionInput{}
        output := &rekognition.StartFaceDetectionOutput{}

        mockClient.On("StartFaceDetection", ctx, input).Return(output, nil)

        result, err := mockClient.StartFaceDetection(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartFaceSearch", func(t *testing.T) {
        input := &rekognition.StartFaceSearchInput{}
        output := &rekognition.StartFaceSearchOutput{}

        mockClient.On("StartFaceSearch", ctx, input).Return(output, nil)

        result, err := mockClient.StartFaceSearch(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartLabelDetection", func(t *testing.T) {
        input := &rekognition.StartLabelDetectionInput{}
        output := &rekognition.StartLabelDetectionOutput{}

        mockClient.On("StartLabelDetection", ctx, input).Return(output, nil)

        result, err := mockClient.StartLabelDetection(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartMediaAnalysisJob", func(t *testing.T) {
        input := &rekognition.StartMediaAnalysisJobInput{}
        output := &rekognition.StartMediaAnalysisJobOutput{}

        mockClient.On("StartMediaAnalysisJob", ctx, input).Return(output, nil)

        result, err := mockClient.StartMediaAnalysisJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartPersonTracking", func(t *testing.T) {
        input := &rekognition.StartPersonTrackingInput{}
        output := &rekognition.StartPersonTrackingOutput{}

        mockClient.On("StartPersonTracking", ctx, input).Return(output, nil)

        result, err := mockClient.StartPersonTracking(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartProjectVersion", func(t *testing.T) {
        input := &rekognition.StartProjectVersionInput{}
        output := &rekognition.StartProjectVersionOutput{}

        mockClient.On("StartProjectVersion", ctx, input).Return(output, nil)

        result, err := mockClient.StartProjectVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartSegmentDetection", func(t *testing.T) {
        input := &rekognition.StartSegmentDetectionInput{}
        output := &rekognition.StartSegmentDetectionOutput{}

        mockClient.On("StartSegmentDetection", ctx, input).Return(output, nil)

        result, err := mockClient.StartSegmentDetection(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartStreamProcessor", func(t *testing.T) {
        input := &rekognition.StartStreamProcessorInput{}
        output := &rekognition.StartStreamProcessorOutput{}

        mockClient.On("StartStreamProcessor", ctx, input).Return(output, nil)

        result, err := mockClient.StartStreamProcessor(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartTextDetection", func(t *testing.T) {
        input := &rekognition.StartTextDetectionInput{}
        output := &rekognition.StartTextDetectionOutput{}

        mockClient.On("StartTextDetection", ctx, input).Return(output, nil)

        result, err := mockClient.StartTextDetection(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopProjectVersion", func(t *testing.T) {
        input := &rekognition.StopProjectVersionInput{}
        output := &rekognition.StopProjectVersionOutput{}

        mockClient.On("StopProjectVersion", ctx, input).Return(output, nil)

        result, err := mockClient.StopProjectVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopStreamProcessor", func(t *testing.T) {
        input := &rekognition.StopStreamProcessorInput{}
        output := &rekognition.StopStreamProcessorOutput{}

        mockClient.On("StopStreamProcessor", ctx, input).Return(output, nil)

        result, err := mockClient.StopStreamProcessor(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &rekognition.TagResourceInput{}
        output := &rekognition.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &rekognition.UntagResourceInput{}
        output := &rekognition.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateDatasetEntries", func(t *testing.T) {
        input := &rekognition.UpdateDatasetEntriesInput{}
        output := &rekognition.UpdateDatasetEntriesOutput{}

        mockClient.On("UpdateDatasetEntries", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateDatasetEntries(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateStreamProcessor", func(t *testing.T) {
        input := &rekognition.UpdateStreamProcessorInput{}
        output := &rekognition.UpdateStreamProcessorOutput{}

        mockClient.On("UpdateStreamProcessor", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateStreamProcessor(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
