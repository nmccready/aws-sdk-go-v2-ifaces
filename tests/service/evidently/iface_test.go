package evidently_test

// tests for the evidently service interface for this ../../../service/evidently/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/evidently"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/evidently/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/evidently/evidently_iface"
	"github.com/stretchr/testify/assert"
)

func TestEvidentlyServiceCanBeMocked(t *testing.T) {
	var iface evidently_iface.IClient
	iface = &evidently.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := evidently.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchEvaluateFeature", func(t *testing.T) {
        input := &evidently.BatchEvaluateFeatureInput{}
        output := &evidently.BatchEvaluateFeatureOutput{}

        mockClient.On("BatchEvaluateFeature", ctx, input).Return(output, nil)

        result, err := mockClient.BatchEvaluateFeature(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateExperiment", func(t *testing.T) {
        input := &evidently.CreateExperimentInput{}
        output := &evidently.CreateExperimentOutput{}

        mockClient.On("CreateExperiment", ctx, input).Return(output, nil)

        result, err := mockClient.CreateExperiment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateFeature", func(t *testing.T) {
        input := &evidently.CreateFeatureInput{}
        output := &evidently.CreateFeatureOutput{}

        mockClient.On("CreateFeature", ctx, input).Return(output, nil)

        result, err := mockClient.CreateFeature(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateLaunch", func(t *testing.T) {
        input := &evidently.CreateLaunchInput{}
        output := &evidently.CreateLaunchOutput{}

        mockClient.On("CreateLaunch", ctx, input).Return(output, nil)

        result, err := mockClient.CreateLaunch(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateProject", func(t *testing.T) {
        input := &evidently.CreateProjectInput{}
        output := &evidently.CreateProjectOutput{}

        mockClient.On("CreateProject", ctx, input).Return(output, nil)

        result, err := mockClient.CreateProject(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateSegment", func(t *testing.T) {
        input := &evidently.CreateSegmentInput{}
        output := &evidently.CreateSegmentOutput{}

        mockClient.On("CreateSegment", ctx, input).Return(output, nil)

        result, err := mockClient.CreateSegment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteExperiment", func(t *testing.T) {
        input := &evidently.DeleteExperimentInput{}
        output := &evidently.DeleteExperimentOutput{}

        mockClient.On("DeleteExperiment", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteExperiment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteFeature", func(t *testing.T) {
        input := &evidently.DeleteFeatureInput{}
        output := &evidently.DeleteFeatureOutput{}

        mockClient.On("DeleteFeature", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteFeature(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteLaunch", func(t *testing.T) {
        input := &evidently.DeleteLaunchInput{}
        output := &evidently.DeleteLaunchOutput{}

        mockClient.On("DeleteLaunch", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteLaunch(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteProject", func(t *testing.T) {
        input := &evidently.DeleteProjectInput{}
        output := &evidently.DeleteProjectOutput{}

        mockClient.On("DeleteProject", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteProject(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteSegment", func(t *testing.T) {
        input := &evidently.DeleteSegmentInput{}
        output := &evidently.DeleteSegmentOutput{}

        mockClient.On("DeleteSegment", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteSegment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestEvaluateFeature", func(t *testing.T) {
        input := &evidently.EvaluateFeatureInput{}
        output := &evidently.EvaluateFeatureOutput{}

        mockClient.On("EvaluateFeature", ctx, input).Return(output, nil)

        result, err := mockClient.EvaluateFeature(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetExperiment", func(t *testing.T) {
        input := &evidently.GetExperimentInput{}
        output := &evidently.GetExperimentOutput{}

        mockClient.On("GetExperiment", ctx, input).Return(output, nil)

        result, err := mockClient.GetExperiment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetExperimentResults", func(t *testing.T) {
        input := &evidently.GetExperimentResultsInput{}
        output := &evidently.GetExperimentResultsOutput{}

        mockClient.On("GetExperimentResults", ctx, input).Return(output, nil)

        result, err := mockClient.GetExperimentResults(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetFeature", func(t *testing.T) {
        input := &evidently.GetFeatureInput{}
        output := &evidently.GetFeatureOutput{}

        mockClient.On("GetFeature", ctx, input).Return(output, nil)

        result, err := mockClient.GetFeature(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetLaunch", func(t *testing.T) {
        input := &evidently.GetLaunchInput{}
        output := &evidently.GetLaunchOutput{}

        mockClient.On("GetLaunch", ctx, input).Return(output, nil)

        result, err := mockClient.GetLaunch(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetProject", func(t *testing.T) {
        input := &evidently.GetProjectInput{}
        output := &evidently.GetProjectOutput{}

        mockClient.On("GetProject", ctx, input).Return(output, nil)

        result, err := mockClient.GetProject(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSegment", func(t *testing.T) {
        input := &evidently.GetSegmentInput{}
        output := &evidently.GetSegmentOutput{}

        mockClient.On("GetSegment", ctx, input).Return(output, nil)

        result, err := mockClient.GetSegment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListExperiments", func(t *testing.T) {
        input := &evidently.ListExperimentsInput{}
        output := &evidently.ListExperimentsOutput{}

        mockClient.On("ListExperiments", ctx, input).Return(output, nil)

        result, err := mockClient.ListExperiments(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListFeatures", func(t *testing.T) {
        input := &evidently.ListFeaturesInput{}
        output := &evidently.ListFeaturesOutput{}

        mockClient.On("ListFeatures", ctx, input).Return(output, nil)

        result, err := mockClient.ListFeatures(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListLaunches", func(t *testing.T) {
        input := &evidently.ListLaunchesInput{}
        output := &evidently.ListLaunchesOutput{}

        mockClient.On("ListLaunches", ctx, input).Return(output, nil)

        result, err := mockClient.ListLaunches(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListProjects", func(t *testing.T) {
        input := &evidently.ListProjectsInput{}
        output := &evidently.ListProjectsOutput{}

        mockClient.On("ListProjects", ctx, input).Return(output, nil)

        result, err := mockClient.ListProjects(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListSegmentReferences", func(t *testing.T) {
        input := &evidently.ListSegmentReferencesInput{}
        output := &evidently.ListSegmentReferencesOutput{}

        mockClient.On("ListSegmentReferences", ctx, input).Return(output, nil)

        result, err := mockClient.ListSegmentReferences(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListSegments", func(t *testing.T) {
        input := &evidently.ListSegmentsInput{}
        output := &evidently.ListSegmentsOutput{}

        mockClient.On("ListSegments", ctx, input).Return(output, nil)

        result, err := mockClient.ListSegments(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &evidently.ListTagsForResourceInput{}
        output := &evidently.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutProjectEvents", func(t *testing.T) {
        input := &evidently.PutProjectEventsInput{}
        output := &evidently.PutProjectEventsOutput{}

        mockClient.On("PutProjectEvents", ctx, input).Return(output, nil)

        result, err := mockClient.PutProjectEvents(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartExperiment", func(t *testing.T) {
        input := &evidently.StartExperimentInput{}
        output := &evidently.StartExperimentOutput{}

        mockClient.On("StartExperiment", ctx, input).Return(output, nil)

        result, err := mockClient.StartExperiment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartLaunch", func(t *testing.T) {
        input := &evidently.StartLaunchInput{}
        output := &evidently.StartLaunchOutput{}

        mockClient.On("StartLaunch", ctx, input).Return(output, nil)

        result, err := mockClient.StartLaunch(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopExperiment", func(t *testing.T) {
        input := &evidently.StopExperimentInput{}
        output := &evidently.StopExperimentOutput{}

        mockClient.On("StopExperiment", ctx, input).Return(output, nil)

        result, err := mockClient.StopExperiment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopLaunch", func(t *testing.T) {
        input := &evidently.StopLaunchInput{}
        output := &evidently.StopLaunchOutput{}

        mockClient.On("StopLaunch", ctx, input).Return(output, nil)

        result, err := mockClient.StopLaunch(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &evidently.TagResourceInput{}
        output := &evidently.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTestSegmentPattern", func(t *testing.T) {
        input := &evidently.TestSegmentPatternInput{}
        output := &evidently.TestSegmentPatternOutput{}

        mockClient.On("TestSegmentPattern", ctx, input).Return(output, nil)

        result, err := mockClient.TestSegmentPattern(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &evidently.UntagResourceInput{}
        output := &evidently.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateExperiment", func(t *testing.T) {
        input := &evidently.UpdateExperimentInput{}
        output := &evidently.UpdateExperimentOutput{}

        mockClient.On("UpdateExperiment", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateExperiment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateFeature", func(t *testing.T) {
        input := &evidently.UpdateFeatureInput{}
        output := &evidently.UpdateFeatureOutput{}

        mockClient.On("UpdateFeature", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateFeature(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateLaunch", func(t *testing.T) {
        input := &evidently.UpdateLaunchInput{}
        output := &evidently.UpdateLaunchOutput{}

        mockClient.On("UpdateLaunch", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateLaunch(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateProject", func(t *testing.T) {
        input := &evidently.UpdateProjectInput{}
        output := &evidently.UpdateProjectOutput{}

        mockClient.On("UpdateProject", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateProject(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateProjectDataDelivery", func(t *testing.T) {
        input := &evidently.UpdateProjectDataDeliveryInput{}
        output := &evidently.UpdateProjectDataDeliveryOutput{}

        mockClient.On("UpdateProjectDataDelivery", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateProjectDataDelivery(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
