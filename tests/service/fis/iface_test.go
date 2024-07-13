package fis_test

// tests for the fis service interface for this ../../../service/fis/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/fis"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/fis/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/fis/fis_iface"
	"github.com/stretchr/testify/assert"
)

func TestFisServiceCanBeMocked(t *testing.T) {
	var iface fis_iface.IClient
	iface = &fis.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := fis.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateExperimentTemplate", func(t *testing.T) {
        input := &fis.CreateExperimentTemplateInput{}
        output := &fis.CreateExperimentTemplateOutput{}

        mockClient.On("CreateExperimentTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.CreateExperimentTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateTargetAccountConfiguration", func(t *testing.T) {
        input := &fis.CreateTargetAccountConfigurationInput{}
        output := &fis.CreateTargetAccountConfigurationOutput{}

        mockClient.On("CreateTargetAccountConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.CreateTargetAccountConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteExperimentTemplate", func(t *testing.T) {
        input := &fis.DeleteExperimentTemplateInput{}
        output := &fis.DeleteExperimentTemplateOutput{}

        mockClient.On("DeleteExperimentTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteExperimentTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteTargetAccountConfiguration", func(t *testing.T) {
        input := &fis.DeleteTargetAccountConfigurationInput{}
        output := &fis.DeleteTargetAccountConfigurationOutput{}

        mockClient.On("DeleteTargetAccountConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteTargetAccountConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetAction", func(t *testing.T) {
        input := &fis.GetActionInput{}
        output := &fis.GetActionOutput{}

        mockClient.On("GetAction", ctx, input).Return(output, nil)

        result, err := mockClient.GetAction(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetExperiment", func(t *testing.T) {
        input := &fis.GetExperimentInput{}
        output := &fis.GetExperimentOutput{}

        mockClient.On("GetExperiment", ctx, input).Return(output, nil)

        result, err := mockClient.GetExperiment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetExperimentTargetAccountConfiguration", func(t *testing.T) {
        input := &fis.GetExperimentTargetAccountConfigurationInput{}
        output := &fis.GetExperimentTargetAccountConfigurationOutput{}

        mockClient.On("GetExperimentTargetAccountConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.GetExperimentTargetAccountConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetExperimentTemplate", func(t *testing.T) {
        input := &fis.GetExperimentTemplateInput{}
        output := &fis.GetExperimentTemplateOutput{}

        mockClient.On("GetExperimentTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.GetExperimentTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetTargetAccountConfiguration", func(t *testing.T) {
        input := &fis.GetTargetAccountConfigurationInput{}
        output := &fis.GetTargetAccountConfigurationOutput{}

        mockClient.On("GetTargetAccountConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.GetTargetAccountConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetTargetResourceType", func(t *testing.T) {
        input := &fis.GetTargetResourceTypeInput{}
        output := &fis.GetTargetResourceTypeOutput{}

        mockClient.On("GetTargetResourceType", ctx, input).Return(output, nil)

        result, err := mockClient.GetTargetResourceType(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListActions", func(t *testing.T) {
        input := &fis.ListActionsInput{}
        output := &fis.ListActionsOutput{}

        mockClient.On("ListActions", ctx, input).Return(output, nil)

        result, err := mockClient.ListActions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListExperimentResolvedTargets", func(t *testing.T) {
        input := &fis.ListExperimentResolvedTargetsInput{}
        output := &fis.ListExperimentResolvedTargetsOutput{}

        mockClient.On("ListExperimentResolvedTargets", ctx, input).Return(output, nil)

        result, err := mockClient.ListExperimentResolvedTargets(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListExperimentTargetAccountConfigurations", func(t *testing.T) {
        input := &fis.ListExperimentTargetAccountConfigurationsInput{}
        output := &fis.ListExperimentTargetAccountConfigurationsOutput{}

        mockClient.On("ListExperimentTargetAccountConfigurations", ctx, input).Return(output, nil)

        result, err := mockClient.ListExperimentTargetAccountConfigurations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListExperimentTemplates", func(t *testing.T) {
        input := &fis.ListExperimentTemplatesInput{}
        output := &fis.ListExperimentTemplatesOutput{}

        mockClient.On("ListExperimentTemplates", ctx, input).Return(output, nil)

        result, err := mockClient.ListExperimentTemplates(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListExperiments", func(t *testing.T) {
        input := &fis.ListExperimentsInput{}
        output := &fis.ListExperimentsOutput{}

        mockClient.On("ListExperiments", ctx, input).Return(output, nil)

        result, err := mockClient.ListExperiments(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &fis.ListTagsForResourceInput{}
        output := &fis.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTargetAccountConfigurations", func(t *testing.T) {
        input := &fis.ListTargetAccountConfigurationsInput{}
        output := &fis.ListTargetAccountConfigurationsOutput{}

        mockClient.On("ListTargetAccountConfigurations", ctx, input).Return(output, nil)

        result, err := mockClient.ListTargetAccountConfigurations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTargetResourceTypes", func(t *testing.T) {
        input := &fis.ListTargetResourceTypesInput{}
        output := &fis.ListTargetResourceTypesOutput{}

        mockClient.On("ListTargetResourceTypes", ctx, input).Return(output, nil)

        result, err := mockClient.ListTargetResourceTypes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartExperiment", func(t *testing.T) {
        input := &fis.StartExperimentInput{}
        output := &fis.StartExperimentOutput{}

        mockClient.On("StartExperiment", ctx, input).Return(output, nil)

        result, err := mockClient.StartExperiment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopExperiment", func(t *testing.T) {
        input := &fis.StopExperimentInput{}
        output := &fis.StopExperimentOutput{}

        mockClient.On("StopExperiment", ctx, input).Return(output, nil)

        result, err := mockClient.StopExperiment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &fis.TagResourceInput{}
        output := &fis.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &fis.UntagResourceInput{}
        output := &fis.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateExperimentTemplate", func(t *testing.T) {
        input := &fis.UpdateExperimentTemplateInput{}
        output := &fis.UpdateExperimentTemplateOutput{}

        mockClient.On("UpdateExperimentTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateExperimentTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateTargetAccountConfiguration", func(t *testing.T) {
        input := &fis.UpdateTargetAccountConfigurationInput{}
        output := &fis.UpdateTargetAccountConfigurationOutput{}

        mockClient.On("UpdateTargetAccountConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateTargetAccountConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
