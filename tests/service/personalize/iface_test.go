package personalize_test

// tests for the personalize service interface for this ../../../service/personalize/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/personalize"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/personalize/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/personalize/personalize_iface"
	"github.com/stretchr/testify/assert"
)

func TestPersonalizeServiceCanBeMocked(t *testing.T) {
	var iface personalize_iface.IClient
	iface = &personalize.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := personalize.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateBatchInferenceJob", func(t *testing.T) {
        input := &personalize.CreateBatchInferenceJobInput{}
        output := &personalize.CreateBatchInferenceJobOutput{}

        mockClient.On("CreateBatchInferenceJob", ctx, input).Return(output, nil)

        result, err := mockClient.CreateBatchInferenceJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateBatchSegmentJob", func(t *testing.T) {
        input := &personalize.CreateBatchSegmentJobInput{}
        output := &personalize.CreateBatchSegmentJobOutput{}

        mockClient.On("CreateBatchSegmentJob", ctx, input).Return(output, nil)

        result, err := mockClient.CreateBatchSegmentJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateCampaign", func(t *testing.T) {
        input := &personalize.CreateCampaignInput{}
        output := &personalize.CreateCampaignOutput{}

        mockClient.On("CreateCampaign", ctx, input).Return(output, nil)

        result, err := mockClient.CreateCampaign(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateDataDeletionJob", func(t *testing.T) {
        input := &personalize.CreateDataDeletionJobInput{}
        output := &personalize.CreateDataDeletionJobOutput{}

        mockClient.On("CreateDataDeletionJob", ctx, input).Return(output, nil)

        result, err := mockClient.CreateDataDeletionJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateDataset", func(t *testing.T) {
        input := &personalize.CreateDatasetInput{}
        output := &personalize.CreateDatasetOutput{}

        mockClient.On("CreateDataset", ctx, input).Return(output, nil)

        result, err := mockClient.CreateDataset(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateDatasetExportJob", func(t *testing.T) {
        input := &personalize.CreateDatasetExportJobInput{}
        output := &personalize.CreateDatasetExportJobOutput{}

        mockClient.On("CreateDatasetExportJob", ctx, input).Return(output, nil)

        result, err := mockClient.CreateDatasetExportJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateDatasetGroup", func(t *testing.T) {
        input := &personalize.CreateDatasetGroupInput{}
        output := &personalize.CreateDatasetGroupOutput{}

        mockClient.On("CreateDatasetGroup", ctx, input).Return(output, nil)

        result, err := mockClient.CreateDatasetGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateDatasetImportJob", func(t *testing.T) {
        input := &personalize.CreateDatasetImportJobInput{}
        output := &personalize.CreateDatasetImportJobOutput{}

        mockClient.On("CreateDatasetImportJob", ctx, input).Return(output, nil)

        result, err := mockClient.CreateDatasetImportJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateEventTracker", func(t *testing.T) {
        input := &personalize.CreateEventTrackerInput{}
        output := &personalize.CreateEventTrackerOutput{}

        mockClient.On("CreateEventTracker", ctx, input).Return(output, nil)

        result, err := mockClient.CreateEventTracker(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateFilter", func(t *testing.T) {
        input := &personalize.CreateFilterInput{}
        output := &personalize.CreateFilterOutput{}

        mockClient.On("CreateFilter", ctx, input).Return(output, nil)

        result, err := mockClient.CreateFilter(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateMetricAttribution", func(t *testing.T) {
        input := &personalize.CreateMetricAttributionInput{}
        output := &personalize.CreateMetricAttributionOutput{}

        mockClient.On("CreateMetricAttribution", ctx, input).Return(output, nil)

        result, err := mockClient.CreateMetricAttribution(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateRecommender", func(t *testing.T) {
        input := &personalize.CreateRecommenderInput{}
        output := &personalize.CreateRecommenderOutput{}

        mockClient.On("CreateRecommender", ctx, input).Return(output, nil)

        result, err := mockClient.CreateRecommender(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateSchema", func(t *testing.T) {
        input := &personalize.CreateSchemaInput{}
        output := &personalize.CreateSchemaOutput{}

        mockClient.On("CreateSchema", ctx, input).Return(output, nil)

        result, err := mockClient.CreateSchema(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateSolution", func(t *testing.T) {
        input := &personalize.CreateSolutionInput{}
        output := &personalize.CreateSolutionOutput{}

        mockClient.On("CreateSolution", ctx, input).Return(output, nil)

        result, err := mockClient.CreateSolution(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateSolutionVersion", func(t *testing.T) {
        input := &personalize.CreateSolutionVersionInput{}
        output := &personalize.CreateSolutionVersionOutput{}

        mockClient.On("CreateSolutionVersion", ctx, input).Return(output, nil)

        result, err := mockClient.CreateSolutionVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteCampaign", func(t *testing.T) {
        input := &personalize.DeleteCampaignInput{}
        output := &personalize.DeleteCampaignOutput{}

        mockClient.On("DeleteCampaign", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteCampaign(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteDataset", func(t *testing.T) {
        input := &personalize.DeleteDatasetInput{}
        output := &personalize.DeleteDatasetOutput{}

        mockClient.On("DeleteDataset", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteDataset(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteDatasetGroup", func(t *testing.T) {
        input := &personalize.DeleteDatasetGroupInput{}
        output := &personalize.DeleteDatasetGroupOutput{}

        mockClient.On("DeleteDatasetGroup", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteDatasetGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteEventTracker", func(t *testing.T) {
        input := &personalize.DeleteEventTrackerInput{}
        output := &personalize.DeleteEventTrackerOutput{}

        mockClient.On("DeleteEventTracker", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteEventTracker(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteFilter", func(t *testing.T) {
        input := &personalize.DeleteFilterInput{}
        output := &personalize.DeleteFilterOutput{}

        mockClient.On("DeleteFilter", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteFilter(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteMetricAttribution", func(t *testing.T) {
        input := &personalize.DeleteMetricAttributionInput{}
        output := &personalize.DeleteMetricAttributionOutput{}

        mockClient.On("DeleteMetricAttribution", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteMetricAttribution(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteRecommender", func(t *testing.T) {
        input := &personalize.DeleteRecommenderInput{}
        output := &personalize.DeleteRecommenderOutput{}

        mockClient.On("DeleteRecommender", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteRecommender(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteSchema", func(t *testing.T) {
        input := &personalize.DeleteSchemaInput{}
        output := &personalize.DeleteSchemaOutput{}

        mockClient.On("DeleteSchema", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteSchema(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteSolution", func(t *testing.T) {
        input := &personalize.DeleteSolutionInput{}
        output := &personalize.DeleteSolutionOutput{}

        mockClient.On("DeleteSolution", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteSolution(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeAlgorithm", func(t *testing.T) {
        input := &personalize.DescribeAlgorithmInput{}
        output := &personalize.DescribeAlgorithmOutput{}

        mockClient.On("DescribeAlgorithm", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeAlgorithm(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeBatchInferenceJob", func(t *testing.T) {
        input := &personalize.DescribeBatchInferenceJobInput{}
        output := &personalize.DescribeBatchInferenceJobOutput{}

        mockClient.On("DescribeBatchInferenceJob", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeBatchInferenceJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeBatchSegmentJob", func(t *testing.T) {
        input := &personalize.DescribeBatchSegmentJobInput{}
        output := &personalize.DescribeBatchSegmentJobOutput{}

        mockClient.On("DescribeBatchSegmentJob", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeBatchSegmentJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeCampaign", func(t *testing.T) {
        input := &personalize.DescribeCampaignInput{}
        output := &personalize.DescribeCampaignOutput{}

        mockClient.On("DescribeCampaign", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeCampaign(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeDataDeletionJob", func(t *testing.T) {
        input := &personalize.DescribeDataDeletionJobInput{}
        output := &personalize.DescribeDataDeletionJobOutput{}

        mockClient.On("DescribeDataDeletionJob", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeDataDeletionJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeDataset", func(t *testing.T) {
        input := &personalize.DescribeDatasetInput{}
        output := &personalize.DescribeDatasetOutput{}

        mockClient.On("DescribeDataset", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeDataset(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeDatasetExportJob", func(t *testing.T) {
        input := &personalize.DescribeDatasetExportJobInput{}
        output := &personalize.DescribeDatasetExportJobOutput{}

        mockClient.On("DescribeDatasetExportJob", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeDatasetExportJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeDatasetGroup", func(t *testing.T) {
        input := &personalize.DescribeDatasetGroupInput{}
        output := &personalize.DescribeDatasetGroupOutput{}

        mockClient.On("DescribeDatasetGroup", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeDatasetGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeDatasetImportJob", func(t *testing.T) {
        input := &personalize.DescribeDatasetImportJobInput{}
        output := &personalize.DescribeDatasetImportJobOutput{}

        mockClient.On("DescribeDatasetImportJob", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeDatasetImportJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeEventTracker", func(t *testing.T) {
        input := &personalize.DescribeEventTrackerInput{}
        output := &personalize.DescribeEventTrackerOutput{}

        mockClient.On("DescribeEventTracker", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeEventTracker(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeFeatureTransformation", func(t *testing.T) {
        input := &personalize.DescribeFeatureTransformationInput{}
        output := &personalize.DescribeFeatureTransformationOutput{}

        mockClient.On("DescribeFeatureTransformation", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeFeatureTransformation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeFilter", func(t *testing.T) {
        input := &personalize.DescribeFilterInput{}
        output := &personalize.DescribeFilterOutput{}

        mockClient.On("DescribeFilter", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeFilter(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeMetricAttribution", func(t *testing.T) {
        input := &personalize.DescribeMetricAttributionInput{}
        output := &personalize.DescribeMetricAttributionOutput{}

        mockClient.On("DescribeMetricAttribution", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeMetricAttribution(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeRecipe", func(t *testing.T) {
        input := &personalize.DescribeRecipeInput{}
        output := &personalize.DescribeRecipeOutput{}

        mockClient.On("DescribeRecipe", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeRecipe(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeRecommender", func(t *testing.T) {
        input := &personalize.DescribeRecommenderInput{}
        output := &personalize.DescribeRecommenderOutput{}

        mockClient.On("DescribeRecommender", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeRecommender(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeSchema", func(t *testing.T) {
        input := &personalize.DescribeSchemaInput{}
        output := &personalize.DescribeSchemaOutput{}

        mockClient.On("DescribeSchema", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeSchema(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeSolution", func(t *testing.T) {
        input := &personalize.DescribeSolutionInput{}
        output := &personalize.DescribeSolutionOutput{}

        mockClient.On("DescribeSolution", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeSolution(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeSolutionVersion", func(t *testing.T) {
        input := &personalize.DescribeSolutionVersionInput{}
        output := &personalize.DescribeSolutionVersionOutput{}

        mockClient.On("DescribeSolutionVersion", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeSolutionVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSolutionMetrics", func(t *testing.T) {
        input := &personalize.GetSolutionMetricsInput{}
        output := &personalize.GetSolutionMetricsOutput{}

        mockClient.On("GetSolutionMetrics", ctx, input).Return(output, nil)

        result, err := mockClient.GetSolutionMetrics(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListBatchInferenceJobs", func(t *testing.T) {
        input := &personalize.ListBatchInferenceJobsInput{}
        output := &personalize.ListBatchInferenceJobsOutput{}

        mockClient.On("ListBatchInferenceJobs", ctx, input).Return(output, nil)

        result, err := mockClient.ListBatchInferenceJobs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListBatchSegmentJobs", func(t *testing.T) {
        input := &personalize.ListBatchSegmentJobsInput{}
        output := &personalize.ListBatchSegmentJobsOutput{}

        mockClient.On("ListBatchSegmentJobs", ctx, input).Return(output, nil)

        result, err := mockClient.ListBatchSegmentJobs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListCampaigns", func(t *testing.T) {
        input := &personalize.ListCampaignsInput{}
        output := &personalize.ListCampaignsOutput{}

        mockClient.On("ListCampaigns", ctx, input).Return(output, nil)

        result, err := mockClient.ListCampaigns(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDataDeletionJobs", func(t *testing.T) {
        input := &personalize.ListDataDeletionJobsInput{}
        output := &personalize.ListDataDeletionJobsOutput{}

        mockClient.On("ListDataDeletionJobs", ctx, input).Return(output, nil)

        result, err := mockClient.ListDataDeletionJobs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDatasetExportJobs", func(t *testing.T) {
        input := &personalize.ListDatasetExportJobsInput{}
        output := &personalize.ListDatasetExportJobsOutput{}

        mockClient.On("ListDatasetExportJobs", ctx, input).Return(output, nil)

        result, err := mockClient.ListDatasetExportJobs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDatasetGroups", func(t *testing.T) {
        input := &personalize.ListDatasetGroupsInput{}
        output := &personalize.ListDatasetGroupsOutput{}

        mockClient.On("ListDatasetGroups", ctx, input).Return(output, nil)

        result, err := mockClient.ListDatasetGroups(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDatasetImportJobs", func(t *testing.T) {
        input := &personalize.ListDatasetImportJobsInput{}
        output := &personalize.ListDatasetImportJobsOutput{}

        mockClient.On("ListDatasetImportJobs", ctx, input).Return(output, nil)

        result, err := mockClient.ListDatasetImportJobs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDatasets", func(t *testing.T) {
        input := &personalize.ListDatasetsInput{}
        output := &personalize.ListDatasetsOutput{}

        mockClient.On("ListDatasets", ctx, input).Return(output, nil)

        result, err := mockClient.ListDatasets(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListEventTrackers", func(t *testing.T) {
        input := &personalize.ListEventTrackersInput{}
        output := &personalize.ListEventTrackersOutput{}

        mockClient.On("ListEventTrackers", ctx, input).Return(output, nil)

        result, err := mockClient.ListEventTrackers(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListFilters", func(t *testing.T) {
        input := &personalize.ListFiltersInput{}
        output := &personalize.ListFiltersOutput{}

        mockClient.On("ListFilters", ctx, input).Return(output, nil)

        result, err := mockClient.ListFilters(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListMetricAttributionMetrics", func(t *testing.T) {
        input := &personalize.ListMetricAttributionMetricsInput{}
        output := &personalize.ListMetricAttributionMetricsOutput{}

        mockClient.On("ListMetricAttributionMetrics", ctx, input).Return(output, nil)

        result, err := mockClient.ListMetricAttributionMetrics(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListMetricAttributions", func(t *testing.T) {
        input := &personalize.ListMetricAttributionsInput{}
        output := &personalize.ListMetricAttributionsOutput{}

        mockClient.On("ListMetricAttributions", ctx, input).Return(output, nil)

        result, err := mockClient.ListMetricAttributions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListRecipes", func(t *testing.T) {
        input := &personalize.ListRecipesInput{}
        output := &personalize.ListRecipesOutput{}

        mockClient.On("ListRecipes", ctx, input).Return(output, nil)

        result, err := mockClient.ListRecipes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListRecommenders", func(t *testing.T) {
        input := &personalize.ListRecommendersInput{}
        output := &personalize.ListRecommendersOutput{}

        mockClient.On("ListRecommenders", ctx, input).Return(output, nil)

        result, err := mockClient.ListRecommenders(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListSchemas", func(t *testing.T) {
        input := &personalize.ListSchemasInput{}
        output := &personalize.ListSchemasOutput{}

        mockClient.On("ListSchemas", ctx, input).Return(output, nil)

        result, err := mockClient.ListSchemas(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListSolutionVersions", func(t *testing.T) {
        input := &personalize.ListSolutionVersionsInput{}
        output := &personalize.ListSolutionVersionsOutput{}

        mockClient.On("ListSolutionVersions", ctx, input).Return(output, nil)

        result, err := mockClient.ListSolutionVersions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListSolutions", func(t *testing.T) {
        input := &personalize.ListSolutionsInput{}
        output := &personalize.ListSolutionsOutput{}

        mockClient.On("ListSolutions", ctx, input).Return(output, nil)

        result, err := mockClient.ListSolutions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &personalize.ListTagsForResourceInput{}
        output := &personalize.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartRecommender", func(t *testing.T) {
        input := &personalize.StartRecommenderInput{}
        output := &personalize.StartRecommenderOutput{}

        mockClient.On("StartRecommender", ctx, input).Return(output, nil)

        result, err := mockClient.StartRecommender(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopRecommender", func(t *testing.T) {
        input := &personalize.StopRecommenderInput{}
        output := &personalize.StopRecommenderOutput{}

        mockClient.On("StopRecommender", ctx, input).Return(output, nil)

        result, err := mockClient.StopRecommender(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopSolutionVersionCreation", func(t *testing.T) {
        input := &personalize.StopSolutionVersionCreationInput{}
        output := &personalize.StopSolutionVersionCreationOutput{}

        mockClient.On("StopSolutionVersionCreation", ctx, input).Return(output, nil)

        result, err := mockClient.StopSolutionVersionCreation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &personalize.TagResourceInput{}
        output := &personalize.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &personalize.UntagResourceInput{}
        output := &personalize.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateCampaign", func(t *testing.T) {
        input := &personalize.UpdateCampaignInput{}
        output := &personalize.UpdateCampaignOutput{}

        mockClient.On("UpdateCampaign", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateCampaign(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateDataset", func(t *testing.T) {
        input := &personalize.UpdateDatasetInput{}
        output := &personalize.UpdateDatasetOutput{}

        mockClient.On("UpdateDataset", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateDataset(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateMetricAttribution", func(t *testing.T) {
        input := &personalize.UpdateMetricAttributionInput{}
        output := &personalize.UpdateMetricAttributionOutput{}

        mockClient.On("UpdateMetricAttribution", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateMetricAttribution(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateRecommender", func(t *testing.T) {
        input := &personalize.UpdateRecommenderInput{}
        output := &personalize.UpdateRecommenderOutput{}

        mockClient.On("UpdateRecommender", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateRecommender(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
