package iotevents_test

// tests for the iotevents service interface for this ../../../service/iotevents/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/iotevents"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/iotevents/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/iotevents/iotevents_iface"
	"github.com/stretchr/testify/assert"
)

func TestIoteventsServiceCanBeMocked(t *testing.T) {
	var iface iotevents_iface.IClient
	iface = &iotevents.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := iotevents.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateAlarmModel", func(t *testing.T) {
        input := &iotevents.CreateAlarmModelInput{}
        output := &iotevents.CreateAlarmModelOutput{}

        mockClient.On("CreateAlarmModel", ctx, input).Return(output, nil)

        result, err := mockClient.CreateAlarmModel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateDetectorModel", func(t *testing.T) {
        input := &iotevents.CreateDetectorModelInput{}
        output := &iotevents.CreateDetectorModelOutput{}

        mockClient.On("CreateDetectorModel", ctx, input).Return(output, nil)

        result, err := mockClient.CreateDetectorModel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateInput", func(t *testing.T) {
        input := &iotevents.CreateInputInput{}
        output := &iotevents.CreateInputOutput{}

        mockClient.On("CreateInput", ctx, input).Return(output, nil)

        result, err := mockClient.CreateInput(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteAlarmModel", func(t *testing.T) {
        input := &iotevents.DeleteAlarmModelInput{}
        output := &iotevents.DeleteAlarmModelOutput{}

        mockClient.On("DeleteAlarmModel", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteAlarmModel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteDetectorModel", func(t *testing.T) {
        input := &iotevents.DeleteDetectorModelInput{}
        output := &iotevents.DeleteDetectorModelOutput{}

        mockClient.On("DeleteDetectorModel", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteDetectorModel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteInput", func(t *testing.T) {
        input := &iotevents.DeleteInputInput{}
        output := &iotevents.DeleteInputOutput{}

        mockClient.On("DeleteInput", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteInput(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeAlarmModel", func(t *testing.T) {
        input := &iotevents.DescribeAlarmModelInput{}
        output := &iotevents.DescribeAlarmModelOutput{}

        mockClient.On("DescribeAlarmModel", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeAlarmModel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeDetectorModel", func(t *testing.T) {
        input := &iotevents.DescribeDetectorModelInput{}
        output := &iotevents.DescribeDetectorModelOutput{}

        mockClient.On("DescribeDetectorModel", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeDetectorModel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeDetectorModelAnalysis", func(t *testing.T) {
        input := &iotevents.DescribeDetectorModelAnalysisInput{}
        output := &iotevents.DescribeDetectorModelAnalysisOutput{}

        mockClient.On("DescribeDetectorModelAnalysis", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeDetectorModelAnalysis(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeInput", func(t *testing.T) {
        input := &iotevents.DescribeInputInput{}
        output := &iotevents.DescribeInputOutput{}

        mockClient.On("DescribeInput", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeInput(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeLoggingOptions", func(t *testing.T) {
        input := &iotevents.DescribeLoggingOptionsInput{}
        output := &iotevents.DescribeLoggingOptionsOutput{}

        mockClient.On("DescribeLoggingOptions", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeLoggingOptions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDetectorModelAnalysisResults", func(t *testing.T) {
        input := &iotevents.GetDetectorModelAnalysisResultsInput{}
        output := &iotevents.GetDetectorModelAnalysisResultsOutput{}

        mockClient.On("GetDetectorModelAnalysisResults", ctx, input).Return(output, nil)

        result, err := mockClient.GetDetectorModelAnalysisResults(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAlarmModelVersions", func(t *testing.T) {
        input := &iotevents.ListAlarmModelVersionsInput{}
        output := &iotevents.ListAlarmModelVersionsOutput{}

        mockClient.On("ListAlarmModelVersions", ctx, input).Return(output, nil)

        result, err := mockClient.ListAlarmModelVersions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAlarmModels", func(t *testing.T) {
        input := &iotevents.ListAlarmModelsInput{}
        output := &iotevents.ListAlarmModelsOutput{}

        mockClient.On("ListAlarmModels", ctx, input).Return(output, nil)

        result, err := mockClient.ListAlarmModels(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDetectorModelVersions", func(t *testing.T) {
        input := &iotevents.ListDetectorModelVersionsInput{}
        output := &iotevents.ListDetectorModelVersionsOutput{}

        mockClient.On("ListDetectorModelVersions", ctx, input).Return(output, nil)

        result, err := mockClient.ListDetectorModelVersions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDetectorModels", func(t *testing.T) {
        input := &iotevents.ListDetectorModelsInput{}
        output := &iotevents.ListDetectorModelsOutput{}

        mockClient.On("ListDetectorModels", ctx, input).Return(output, nil)

        result, err := mockClient.ListDetectorModels(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListInputRoutings", func(t *testing.T) {
        input := &iotevents.ListInputRoutingsInput{}
        output := &iotevents.ListInputRoutingsOutput{}

        mockClient.On("ListInputRoutings", ctx, input).Return(output, nil)

        result, err := mockClient.ListInputRoutings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListInputs", func(t *testing.T) {
        input := &iotevents.ListInputsInput{}
        output := &iotevents.ListInputsOutput{}

        mockClient.On("ListInputs", ctx, input).Return(output, nil)

        result, err := mockClient.ListInputs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &iotevents.ListTagsForResourceInput{}
        output := &iotevents.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutLoggingOptions", func(t *testing.T) {
        input := &iotevents.PutLoggingOptionsInput{}
        output := &iotevents.PutLoggingOptionsOutput{}

        mockClient.On("PutLoggingOptions", ctx, input).Return(output, nil)

        result, err := mockClient.PutLoggingOptions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartDetectorModelAnalysis", func(t *testing.T) {
        input := &iotevents.StartDetectorModelAnalysisInput{}
        output := &iotevents.StartDetectorModelAnalysisOutput{}

        mockClient.On("StartDetectorModelAnalysis", ctx, input).Return(output, nil)

        result, err := mockClient.StartDetectorModelAnalysis(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &iotevents.TagResourceInput{}
        output := &iotevents.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &iotevents.UntagResourceInput{}
        output := &iotevents.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateAlarmModel", func(t *testing.T) {
        input := &iotevents.UpdateAlarmModelInput{}
        output := &iotevents.UpdateAlarmModelOutput{}

        mockClient.On("UpdateAlarmModel", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateAlarmModel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateDetectorModel", func(t *testing.T) {
        input := &iotevents.UpdateDetectorModelInput{}
        output := &iotevents.UpdateDetectorModelOutput{}

        mockClient.On("UpdateDetectorModel", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateDetectorModel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateInput", func(t *testing.T) {
        input := &iotevents.UpdateInputInput{}
        output := &iotevents.UpdateInputOutput{}

        mockClient.On("UpdateInput", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateInput(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
