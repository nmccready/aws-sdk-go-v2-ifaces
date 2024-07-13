package comprehend_test

// tests for the comprehend service interface for this ../../../service/comprehend/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/comprehend"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/comprehend/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/comprehend/comprehend_iface"
	"github.com/stretchr/testify/assert"
)

func TestComprehendServiceCanBeMocked(t *testing.T) {
	var iface comprehend_iface.IClient
	iface = &comprehend.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := comprehend.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchDetectDominantLanguage", func(t *testing.T) {
        input := &comprehend.BatchDetectDominantLanguageInput{}
        output := &comprehend.BatchDetectDominantLanguageOutput{}

        mockClient.On("BatchDetectDominantLanguage", ctx, input).Return(output, nil)

        result, err := mockClient.BatchDetectDominantLanguage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchDetectEntities", func(t *testing.T) {
        input := &comprehend.BatchDetectEntitiesInput{}
        output := &comprehend.BatchDetectEntitiesOutput{}

        mockClient.On("BatchDetectEntities", ctx, input).Return(output, nil)

        result, err := mockClient.BatchDetectEntities(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchDetectKeyPhrases", func(t *testing.T) {
        input := &comprehend.BatchDetectKeyPhrasesInput{}
        output := &comprehend.BatchDetectKeyPhrasesOutput{}

        mockClient.On("BatchDetectKeyPhrases", ctx, input).Return(output, nil)

        result, err := mockClient.BatchDetectKeyPhrases(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchDetectSentiment", func(t *testing.T) {
        input := &comprehend.BatchDetectSentimentInput{}
        output := &comprehend.BatchDetectSentimentOutput{}

        mockClient.On("BatchDetectSentiment", ctx, input).Return(output, nil)

        result, err := mockClient.BatchDetectSentiment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchDetectSyntax", func(t *testing.T) {
        input := &comprehend.BatchDetectSyntaxInput{}
        output := &comprehend.BatchDetectSyntaxOutput{}

        mockClient.On("BatchDetectSyntax", ctx, input).Return(output, nil)

        result, err := mockClient.BatchDetectSyntax(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchDetectTargetedSentiment", func(t *testing.T) {
        input := &comprehend.BatchDetectTargetedSentimentInput{}
        output := &comprehend.BatchDetectTargetedSentimentOutput{}

        mockClient.On("BatchDetectTargetedSentiment", ctx, input).Return(output, nil)

        result, err := mockClient.BatchDetectTargetedSentiment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestClassifyDocument", func(t *testing.T) {
        input := &comprehend.ClassifyDocumentInput{}
        output := &comprehend.ClassifyDocumentOutput{}

        mockClient.On("ClassifyDocument", ctx, input).Return(output, nil)

        result, err := mockClient.ClassifyDocument(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestContainsPiiEntities", func(t *testing.T) {
        input := &comprehend.ContainsPiiEntitiesInput{}
        output := &comprehend.ContainsPiiEntitiesOutput{}

        mockClient.On("ContainsPiiEntities", ctx, input).Return(output, nil)

        result, err := mockClient.ContainsPiiEntities(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateDataset", func(t *testing.T) {
        input := &comprehend.CreateDatasetInput{}
        output := &comprehend.CreateDatasetOutput{}

        mockClient.On("CreateDataset", ctx, input).Return(output, nil)

        result, err := mockClient.CreateDataset(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateDocumentClassifier", func(t *testing.T) {
        input := &comprehend.CreateDocumentClassifierInput{}
        output := &comprehend.CreateDocumentClassifierOutput{}

        mockClient.On("CreateDocumentClassifier", ctx, input).Return(output, nil)

        result, err := mockClient.CreateDocumentClassifier(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateEndpoint", func(t *testing.T) {
        input := &comprehend.CreateEndpointInput{}
        output := &comprehend.CreateEndpointOutput{}

        mockClient.On("CreateEndpoint", ctx, input).Return(output, nil)

        result, err := mockClient.CreateEndpoint(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateEntityRecognizer", func(t *testing.T) {
        input := &comprehend.CreateEntityRecognizerInput{}
        output := &comprehend.CreateEntityRecognizerOutput{}

        mockClient.On("CreateEntityRecognizer", ctx, input).Return(output, nil)

        result, err := mockClient.CreateEntityRecognizer(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateFlywheel", func(t *testing.T) {
        input := &comprehend.CreateFlywheelInput{}
        output := &comprehend.CreateFlywheelOutput{}

        mockClient.On("CreateFlywheel", ctx, input).Return(output, nil)

        result, err := mockClient.CreateFlywheel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteDocumentClassifier", func(t *testing.T) {
        input := &comprehend.DeleteDocumentClassifierInput{}
        output := &comprehend.DeleteDocumentClassifierOutput{}

        mockClient.On("DeleteDocumentClassifier", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteDocumentClassifier(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteEndpoint", func(t *testing.T) {
        input := &comprehend.DeleteEndpointInput{}
        output := &comprehend.DeleteEndpointOutput{}

        mockClient.On("DeleteEndpoint", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteEndpoint(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteEntityRecognizer", func(t *testing.T) {
        input := &comprehend.DeleteEntityRecognizerInput{}
        output := &comprehend.DeleteEntityRecognizerOutput{}

        mockClient.On("DeleteEntityRecognizer", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteEntityRecognizer(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteFlywheel", func(t *testing.T) {
        input := &comprehend.DeleteFlywheelInput{}
        output := &comprehend.DeleteFlywheelOutput{}

        mockClient.On("DeleteFlywheel", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteFlywheel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteResourcePolicy", func(t *testing.T) {
        input := &comprehend.DeleteResourcePolicyInput{}
        output := &comprehend.DeleteResourcePolicyOutput{}

        mockClient.On("DeleteResourcePolicy", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteResourcePolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeDataset", func(t *testing.T) {
        input := &comprehend.DescribeDatasetInput{}
        output := &comprehend.DescribeDatasetOutput{}

        mockClient.On("DescribeDataset", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeDataset(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeDocumentClassificationJob", func(t *testing.T) {
        input := &comprehend.DescribeDocumentClassificationJobInput{}
        output := &comprehend.DescribeDocumentClassificationJobOutput{}

        mockClient.On("DescribeDocumentClassificationJob", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeDocumentClassificationJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeDocumentClassifier", func(t *testing.T) {
        input := &comprehend.DescribeDocumentClassifierInput{}
        output := &comprehend.DescribeDocumentClassifierOutput{}

        mockClient.On("DescribeDocumentClassifier", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeDocumentClassifier(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeDominantLanguageDetectionJob", func(t *testing.T) {
        input := &comprehend.DescribeDominantLanguageDetectionJobInput{}
        output := &comprehend.DescribeDominantLanguageDetectionJobOutput{}

        mockClient.On("DescribeDominantLanguageDetectionJob", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeDominantLanguageDetectionJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeEndpoint", func(t *testing.T) {
        input := &comprehend.DescribeEndpointInput{}
        output := &comprehend.DescribeEndpointOutput{}

        mockClient.On("DescribeEndpoint", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeEndpoint(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeEntitiesDetectionJob", func(t *testing.T) {
        input := &comprehend.DescribeEntitiesDetectionJobInput{}
        output := &comprehend.DescribeEntitiesDetectionJobOutput{}

        mockClient.On("DescribeEntitiesDetectionJob", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeEntitiesDetectionJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeEntityRecognizer", func(t *testing.T) {
        input := &comprehend.DescribeEntityRecognizerInput{}
        output := &comprehend.DescribeEntityRecognizerOutput{}

        mockClient.On("DescribeEntityRecognizer", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeEntityRecognizer(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeEventsDetectionJob", func(t *testing.T) {
        input := &comprehend.DescribeEventsDetectionJobInput{}
        output := &comprehend.DescribeEventsDetectionJobOutput{}

        mockClient.On("DescribeEventsDetectionJob", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeEventsDetectionJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeFlywheel", func(t *testing.T) {
        input := &comprehend.DescribeFlywheelInput{}
        output := &comprehend.DescribeFlywheelOutput{}

        mockClient.On("DescribeFlywheel", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeFlywheel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeFlywheelIteration", func(t *testing.T) {
        input := &comprehend.DescribeFlywheelIterationInput{}
        output := &comprehend.DescribeFlywheelIterationOutput{}

        mockClient.On("DescribeFlywheelIteration", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeFlywheelIteration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeKeyPhrasesDetectionJob", func(t *testing.T) {
        input := &comprehend.DescribeKeyPhrasesDetectionJobInput{}
        output := &comprehend.DescribeKeyPhrasesDetectionJobOutput{}

        mockClient.On("DescribeKeyPhrasesDetectionJob", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeKeyPhrasesDetectionJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribePiiEntitiesDetectionJob", func(t *testing.T) {
        input := &comprehend.DescribePiiEntitiesDetectionJobInput{}
        output := &comprehend.DescribePiiEntitiesDetectionJobOutput{}

        mockClient.On("DescribePiiEntitiesDetectionJob", ctx, input).Return(output, nil)

        result, err := mockClient.DescribePiiEntitiesDetectionJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeResourcePolicy", func(t *testing.T) {
        input := &comprehend.DescribeResourcePolicyInput{}
        output := &comprehend.DescribeResourcePolicyOutput{}

        mockClient.On("DescribeResourcePolicy", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeResourcePolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeSentimentDetectionJob", func(t *testing.T) {
        input := &comprehend.DescribeSentimentDetectionJobInput{}
        output := &comprehend.DescribeSentimentDetectionJobOutput{}

        mockClient.On("DescribeSentimentDetectionJob", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeSentimentDetectionJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeTargetedSentimentDetectionJob", func(t *testing.T) {
        input := &comprehend.DescribeTargetedSentimentDetectionJobInput{}
        output := &comprehend.DescribeTargetedSentimentDetectionJobOutput{}

        mockClient.On("DescribeTargetedSentimentDetectionJob", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeTargetedSentimentDetectionJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeTopicsDetectionJob", func(t *testing.T) {
        input := &comprehend.DescribeTopicsDetectionJobInput{}
        output := &comprehend.DescribeTopicsDetectionJobOutput{}

        mockClient.On("DescribeTopicsDetectionJob", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeTopicsDetectionJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDetectDominantLanguage", func(t *testing.T) {
        input := &comprehend.DetectDominantLanguageInput{}
        output := &comprehend.DetectDominantLanguageOutput{}

        mockClient.On("DetectDominantLanguage", ctx, input).Return(output, nil)

        result, err := mockClient.DetectDominantLanguage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDetectEntities", func(t *testing.T) {
        input := &comprehend.DetectEntitiesInput{}
        output := &comprehend.DetectEntitiesOutput{}

        mockClient.On("DetectEntities", ctx, input).Return(output, nil)

        result, err := mockClient.DetectEntities(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDetectKeyPhrases", func(t *testing.T) {
        input := &comprehend.DetectKeyPhrasesInput{}
        output := &comprehend.DetectKeyPhrasesOutput{}

        mockClient.On("DetectKeyPhrases", ctx, input).Return(output, nil)

        result, err := mockClient.DetectKeyPhrases(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDetectPiiEntities", func(t *testing.T) {
        input := &comprehend.DetectPiiEntitiesInput{}
        output := &comprehend.DetectPiiEntitiesOutput{}

        mockClient.On("DetectPiiEntities", ctx, input).Return(output, nil)

        result, err := mockClient.DetectPiiEntities(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDetectSentiment", func(t *testing.T) {
        input := &comprehend.DetectSentimentInput{}
        output := &comprehend.DetectSentimentOutput{}

        mockClient.On("DetectSentiment", ctx, input).Return(output, nil)

        result, err := mockClient.DetectSentiment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDetectSyntax", func(t *testing.T) {
        input := &comprehend.DetectSyntaxInput{}
        output := &comprehend.DetectSyntaxOutput{}

        mockClient.On("DetectSyntax", ctx, input).Return(output, nil)

        result, err := mockClient.DetectSyntax(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDetectTargetedSentiment", func(t *testing.T) {
        input := &comprehend.DetectTargetedSentimentInput{}
        output := &comprehend.DetectTargetedSentimentOutput{}

        mockClient.On("DetectTargetedSentiment", ctx, input).Return(output, nil)

        result, err := mockClient.DetectTargetedSentiment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDetectToxicContent", func(t *testing.T) {
        input := &comprehend.DetectToxicContentInput{}
        output := &comprehend.DetectToxicContentOutput{}

        mockClient.On("DetectToxicContent", ctx, input).Return(output, nil)

        result, err := mockClient.DetectToxicContent(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestImportModel", func(t *testing.T) {
        input := &comprehend.ImportModelInput{}
        output := &comprehend.ImportModelOutput{}

        mockClient.On("ImportModel", ctx, input).Return(output, nil)

        result, err := mockClient.ImportModel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDatasets", func(t *testing.T) {
        input := &comprehend.ListDatasetsInput{}
        output := &comprehend.ListDatasetsOutput{}

        mockClient.On("ListDatasets", ctx, input).Return(output, nil)

        result, err := mockClient.ListDatasets(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDocumentClassificationJobs", func(t *testing.T) {
        input := &comprehend.ListDocumentClassificationJobsInput{}
        output := &comprehend.ListDocumentClassificationJobsOutput{}

        mockClient.On("ListDocumentClassificationJobs", ctx, input).Return(output, nil)

        result, err := mockClient.ListDocumentClassificationJobs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDocumentClassifierSummaries", func(t *testing.T) {
        input := &comprehend.ListDocumentClassifierSummariesInput{}
        output := &comprehend.ListDocumentClassifierSummariesOutput{}

        mockClient.On("ListDocumentClassifierSummaries", ctx, input).Return(output, nil)

        result, err := mockClient.ListDocumentClassifierSummaries(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDocumentClassifiers", func(t *testing.T) {
        input := &comprehend.ListDocumentClassifiersInput{}
        output := &comprehend.ListDocumentClassifiersOutput{}

        mockClient.On("ListDocumentClassifiers", ctx, input).Return(output, nil)

        result, err := mockClient.ListDocumentClassifiers(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDominantLanguageDetectionJobs", func(t *testing.T) {
        input := &comprehend.ListDominantLanguageDetectionJobsInput{}
        output := &comprehend.ListDominantLanguageDetectionJobsOutput{}

        mockClient.On("ListDominantLanguageDetectionJobs", ctx, input).Return(output, nil)

        result, err := mockClient.ListDominantLanguageDetectionJobs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListEndpoints", func(t *testing.T) {
        input := &comprehend.ListEndpointsInput{}
        output := &comprehend.ListEndpointsOutput{}

        mockClient.On("ListEndpoints", ctx, input).Return(output, nil)

        result, err := mockClient.ListEndpoints(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListEntitiesDetectionJobs", func(t *testing.T) {
        input := &comprehend.ListEntitiesDetectionJobsInput{}
        output := &comprehend.ListEntitiesDetectionJobsOutput{}

        mockClient.On("ListEntitiesDetectionJobs", ctx, input).Return(output, nil)

        result, err := mockClient.ListEntitiesDetectionJobs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListEntityRecognizerSummaries", func(t *testing.T) {
        input := &comprehend.ListEntityRecognizerSummariesInput{}
        output := &comprehend.ListEntityRecognizerSummariesOutput{}

        mockClient.On("ListEntityRecognizerSummaries", ctx, input).Return(output, nil)

        result, err := mockClient.ListEntityRecognizerSummaries(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListEntityRecognizers", func(t *testing.T) {
        input := &comprehend.ListEntityRecognizersInput{}
        output := &comprehend.ListEntityRecognizersOutput{}

        mockClient.On("ListEntityRecognizers", ctx, input).Return(output, nil)

        result, err := mockClient.ListEntityRecognizers(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListEventsDetectionJobs", func(t *testing.T) {
        input := &comprehend.ListEventsDetectionJobsInput{}
        output := &comprehend.ListEventsDetectionJobsOutput{}

        mockClient.On("ListEventsDetectionJobs", ctx, input).Return(output, nil)

        result, err := mockClient.ListEventsDetectionJobs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListFlywheelIterationHistory", func(t *testing.T) {
        input := &comprehend.ListFlywheelIterationHistoryInput{}
        output := &comprehend.ListFlywheelIterationHistoryOutput{}

        mockClient.On("ListFlywheelIterationHistory", ctx, input).Return(output, nil)

        result, err := mockClient.ListFlywheelIterationHistory(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListFlywheels", func(t *testing.T) {
        input := &comprehend.ListFlywheelsInput{}
        output := &comprehend.ListFlywheelsOutput{}

        mockClient.On("ListFlywheels", ctx, input).Return(output, nil)

        result, err := mockClient.ListFlywheels(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListKeyPhrasesDetectionJobs", func(t *testing.T) {
        input := &comprehend.ListKeyPhrasesDetectionJobsInput{}
        output := &comprehend.ListKeyPhrasesDetectionJobsOutput{}

        mockClient.On("ListKeyPhrasesDetectionJobs", ctx, input).Return(output, nil)

        result, err := mockClient.ListKeyPhrasesDetectionJobs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListPiiEntitiesDetectionJobs", func(t *testing.T) {
        input := &comprehend.ListPiiEntitiesDetectionJobsInput{}
        output := &comprehend.ListPiiEntitiesDetectionJobsOutput{}

        mockClient.On("ListPiiEntitiesDetectionJobs", ctx, input).Return(output, nil)

        result, err := mockClient.ListPiiEntitiesDetectionJobs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListSentimentDetectionJobs", func(t *testing.T) {
        input := &comprehend.ListSentimentDetectionJobsInput{}
        output := &comprehend.ListSentimentDetectionJobsOutput{}

        mockClient.On("ListSentimentDetectionJobs", ctx, input).Return(output, nil)

        result, err := mockClient.ListSentimentDetectionJobs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &comprehend.ListTagsForResourceInput{}
        output := &comprehend.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTargetedSentimentDetectionJobs", func(t *testing.T) {
        input := &comprehend.ListTargetedSentimentDetectionJobsInput{}
        output := &comprehend.ListTargetedSentimentDetectionJobsOutput{}

        mockClient.On("ListTargetedSentimentDetectionJobs", ctx, input).Return(output, nil)

        result, err := mockClient.ListTargetedSentimentDetectionJobs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTopicsDetectionJobs", func(t *testing.T) {
        input := &comprehend.ListTopicsDetectionJobsInput{}
        output := &comprehend.ListTopicsDetectionJobsOutput{}

        mockClient.On("ListTopicsDetectionJobs", ctx, input).Return(output, nil)

        result, err := mockClient.ListTopicsDetectionJobs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutResourcePolicy", func(t *testing.T) {
        input := &comprehend.PutResourcePolicyInput{}
        output := &comprehend.PutResourcePolicyOutput{}

        mockClient.On("PutResourcePolicy", ctx, input).Return(output, nil)

        result, err := mockClient.PutResourcePolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartDocumentClassificationJob", func(t *testing.T) {
        input := &comprehend.StartDocumentClassificationJobInput{}
        output := &comprehend.StartDocumentClassificationJobOutput{}

        mockClient.On("StartDocumentClassificationJob", ctx, input).Return(output, nil)

        result, err := mockClient.StartDocumentClassificationJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartDominantLanguageDetectionJob", func(t *testing.T) {
        input := &comprehend.StartDominantLanguageDetectionJobInput{}
        output := &comprehend.StartDominantLanguageDetectionJobOutput{}

        mockClient.On("StartDominantLanguageDetectionJob", ctx, input).Return(output, nil)

        result, err := mockClient.StartDominantLanguageDetectionJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartEntitiesDetectionJob", func(t *testing.T) {
        input := &comprehend.StartEntitiesDetectionJobInput{}
        output := &comprehend.StartEntitiesDetectionJobOutput{}

        mockClient.On("StartEntitiesDetectionJob", ctx, input).Return(output, nil)

        result, err := mockClient.StartEntitiesDetectionJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartEventsDetectionJob", func(t *testing.T) {
        input := &comprehend.StartEventsDetectionJobInput{}
        output := &comprehend.StartEventsDetectionJobOutput{}

        mockClient.On("StartEventsDetectionJob", ctx, input).Return(output, nil)

        result, err := mockClient.StartEventsDetectionJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartFlywheelIteration", func(t *testing.T) {
        input := &comprehend.StartFlywheelIterationInput{}
        output := &comprehend.StartFlywheelIterationOutput{}

        mockClient.On("StartFlywheelIteration", ctx, input).Return(output, nil)

        result, err := mockClient.StartFlywheelIteration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartKeyPhrasesDetectionJob", func(t *testing.T) {
        input := &comprehend.StartKeyPhrasesDetectionJobInput{}
        output := &comprehend.StartKeyPhrasesDetectionJobOutput{}

        mockClient.On("StartKeyPhrasesDetectionJob", ctx, input).Return(output, nil)

        result, err := mockClient.StartKeyPhrasesDetectionJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartPiiEntitiesDetectionJob", func(t *testing.T) {
        input := &comprehend.StartPiiEntitiesDetectionJobInput{}
        output := &comprehend.StartPiiEntitiesDetectionJobOutput{}

        mockClient.On("StartPiiEntitiesDetectionJob", ctx, input).Return(output, nil)

        result, err := mockClient.StartPiiEntitiesDetectionJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartSentimentDetectionJob", func(t *testing.T) {
        input := &comprehend.StartSentimentDetectionJobInput{}
        output := &comprehend.StartSentimentDetectionJobOutput{}

        mockClient.On("StartSentimentDetectionJob", ctx, input).Return(output, nil)

        result, err := mockClient.StartSentimentDetectionJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartTargetedSentimentDetectionJob", func(t *testing.T) {
        input := &comprehend.StartTargetedSentimentDetectionJobInput{}
        output := &comprehend.StartTargetedSentimentDetectionJobOutput{}

        mockClient.On("StartTargetedSentimentDetectionJob", ctx, input).Return(output, nil)

        result, err := mockClient.StartTargetedSentimentDetectionJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartTopicsDetectionJob", func(t *testing.T) {
        input := &comprehend.StartTopicsDetectionJobInput{}
        output := &comprehend.StartTopicsDetectionJobOutput{}

        mockClient.On("StartTopicsDetectionJob", ctx, input).Return(output, nil)

        result, err := mockClient.StartTopicsDetectionJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopDominantLanguageDetectionJob", func(t *testing.T) {
        input := &comprehend.StopDominantLanguageDetectionJobInput{}
        output := &comprehend.StopDominantLanguageDetectionJobOutput{}

        mockClient.On("StopDominantLanguageDetectionJob", ctx, input).Return(output, nil)

        result, err := mockClient.StopDominantLanguageDetectionJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopEntitiesDetectionJob", func(t *testing.T) {
        input := &comprehend.StopEntitiesDetectionJobInput{}
        output := &comprehend.StopEntitiesDetectionJobOutput{}

        mockClient.On("StopEntitiesDetectionJob", ctx, input).Return(output, nil)

        result, err := mockClient.StopEntitiesDetectionJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopEventsDetectionJob", func(t *testing.T) {
        input := &comprehend.StopEventsDetectionJobInput{}
        output := &comprehend.StopEventsDetectionJobOutput{}

        mockClient.On("StopEventsDetectionJob", ctx, input).Return(output, nil)

        result, err := mockClient.StopEventsDetectionJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopKeyPhrasesDetectionJob", func(t *testing.T) {
        input := &comprehend.StopKeyPhrasesDetectionJobInput{}
        output := &comprehend.StopKeyPhrasesDetectionJobOutput{}

        mockClient.On("StopKeyPhrasesDetectionJob", ctx, input).Return(output, nil)

        result, err := mockClient.StopKeyPhrasesDetectionJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopPiiEntitiesDetectionJob", func(t *testing.T) {
        input := &comprehend.StopPiiEntitiesDetectionJobInput{}
        output := &comprehend.StopPiiEntitiesDetectionJobOutput{}

        mockClient.On("StopPiiEntitiesDetectionJob", ctx, input).Return(output, nil)

        result, err := mockClient.StopPiiEntitiesDetectionJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopSentimentDetectionJob", func(t *testing.T) {
        input := &comprehend.StopSentimentDetectionJobInput{}
        output := &comprehend.StopSentimentDetectionJobOutput{}

        mockClient.On("StopSentimentDetectionJob", ctx, input).Return(output, nil)

        result, err := mockClient.StopSentimentDetectionJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopTargetedSentimentDetectionJob", func(t *testing.T) {
        input := &comprehend.StopTargetedSentimentDetectionJobInput{}
        output := &comprehend.StopTargetedSentimentDetectionJobOutput{}

        mockClient.On("StopTargetedSentimentDetectionJob", ctx, input).Return(output, nil)

        result, err := mockClient.StopTargetedSentimentDetectionJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopTrainingDocumentClassifier", func(t *testing.T) {
        input := &comprehend.StopTrainingDocumentClassifierInput{}
        output := &comprehend.StopTrainingDocumentClassifierOutput{}

        mockClient.On("StopTrainingDocumentClassifier", ctx, input).Return(output, nil)

        result, err := mockClient.StopTrainingDocumentClassifier(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopTrainingEntityRecognizer", func(t *testing.T) {
        input := &comprehend.StopTrainingEntityRecognizerInput{}
        output := &comprehend.StopTrainingEntityRecognizerOutput{}

        mockClient.On("StopTrainingEntityRecognizer", ctx, input).Return(output, nil)

        result, err := mockClient.StopTrainingEntityRecognizer(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &comprehend.TagResourceInput{}
        output := &comprehend.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &comprehend.UntagResourceInput{}
        output := &comprehend.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateEndpoint", func(t *testing.T) {
        input := &comprehend.UpdateEndpointInput{}
        output := &comprehend.UpdateEndpointOutput{}

        mockClient.On("UpdateEndpoint", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateEndpoint(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateFlywheel", func(t *testing.T) {
        input := &comprehend.UpdateFlywheelInput{}
        output := &comprehend.UpdateFlywheelOutput{}

        mockClient.On("UpdateFlywheel", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateFlywheel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
