// Code generated by ./scripts/gen_tests.sh DO NOT EDIT.
package frauddetector_test

// tests for the frauddetector service interface for 
// this ../../../service/frauddetector/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/frauddetector"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/frauddetector/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/frauddetector/frauddetector_iface"
	"github.com/stretchr/testify/assert"
)

func TestFrauddetectorServiceCanBeMocked(t *testing.T) {
	var iface frauddetector_iface.IClient
	iface = &frauddetector.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := frauddetector.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchCreateVariable", func(t *testing.T) {
        input := &frauddetector.BatchCreateVariableInput{}
        output := &frauddetector.BatchCreateVariableOutput{}

        mockClient.On("BatchCreateVariable", ctx, input).Return(output, nil)

        result, err := mockClient.BatchCreateVariable(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchGetVariable", func(t *testing.T) {
        input := &frauddetector.BatchGetVariableInput{}
        output := &frauddetector.BatchGetVariableOutput{}

        mockClient.On("BatchGetVariable", ctx, input).Return(output, nil)

        result, err := mockClient.BatchGetVariable(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCancelBatchImportJob", func(t *testing.T) {
        input := &frauddetector.CancelBatchImportJobInput{}
        output := &frauddetector.CancelBatchImportJobOutput{}

        mockClient.On("CancelBatchImportJob", ctx, input).Return(output, nil)

        result, err := mockClient.CancelBatchImportJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCancelBatchPredictionJob", func(t *testing.T) {
        input := &frauddetector.CancelBatchPredictionJobInput{}
        output := &frauddetector.CancelBatchPredictionJobOutput{}

        mockClient.On("CancelBatchPredictionJob", ctx, input).Return(output, nil)

        result, err := mockClient.CancelBatchPredictionJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateBatchImportJob", func(t *testing.T) {
        input := &frauddetector.CreateBatchImportJobInput{}
        output := &frauddetector.CreateBatchImportJobOutput{}

        mockClient.On("CreateBatchImportJob", ctx, input).Return(output, nil)

        result, err := mockClient.CreateBatchImportJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateBatchPredictionJob", func(t *testing.T) {
        input := &frauddetector.CreateBatchPredictionJobInput{}
        output := &frauddetector.CreateBatchPredictionJobOutput{}

        mockClient.On("CreateBatchPredictionJob", ctx, input).Return(output, nil)

        result, err := mockClient.CreateBatchPredictionJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateDetectorVersion", func(t *testing.T) {
        input := &frauddetector.CreateDetectorVersionInput{}
        output := &frauddetector.CreateDetectorVersionOutput{}

        mockClient.On("CreateDetectorVersion", ctx, input).Return(output, nil)

        result, err := mockClient.CreateDetectorVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateList", func(t *testing.T) {
        input := &frauddetector.CreateListInput{}
        output := &frauddetector.CreateListOutput{}

        mockClient.On("CreateList", ctx, input).Return(output, nil)

        result, err := mockClient.CreateList(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateModel", func(t *testing.T) {
        input := &frauddetector.CreateModelInput{}
        output := &frauddetector.CreateModelOutput{}

        mockClient.On("CreateModel", ctx, input).Return(output, nil)

        result, err := mockClient.CreateModel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateModelVersion", func(t *testing.T) {
        input := &frauddetector.CreateModelVersionInput{}
        output := &frauddetector.CreateModelVersionOutput{}

        mockClient.On("CreateModelVersion", ctx, input).Return(output, nil)

        result, err := mockClient.CreateModelVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateRule", func(t *testing.T) {
        input := &frauddetector.CreateRuleInput{}
        output := &frauddetector.CreateRuleOutput{}

        mockClient.On("CreateRule", ctx, input).Return(output, nil)

        result, err := mockClient.CreateRule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateVariable", func(t *testing.T) {
        input := &frauddetector.CreateVariableInput{}
        output := &frauddetector.CreateVariableOutput{}

        mockClient.On("CreateVariable", ctx, input).Return(output, nil)

        result, err := mockClient.CreateVariable(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteBatchImportJob", func(t *testing.T) {
        input := &frauddetector.DeleteBatchImportJobInput{}
        output := &frauddetector.DeleteBatchImportJobOutput{}

        mockClient.On("DeleteBatchImportJob", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteBatchImportJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteBatchPredictionJob", func(t *testing.T) {
        input := &frauddetector.DeleteBatchPredictionJobInput{}
        output := &frauddetector.DeleteBatchPredictionJobOutput{}

        mockClient.On("DeleteBatchPredictionJob", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteBatchPredictionJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteDetector", func(t *testing.T) {
        input := &frauddetector.DeleteDetectorInput{}
        output := &frauddetector.DeleteDetectorOutput{}

        mockClient.On("DeleteDetector", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteDetector(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteDetectorVersion", func(t *testing.T) {
        input := &frauddetector.DeleteDetectorVersionInput{}
        output := &frauddetector.DeleteDetectorVersionOutput{}

        mockClient.On("DeleteDetectorVersion", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteDetectorVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteEntityType", func(t *testing.T) {
        input := &frauddetector.DeleteEntityTypeInput{}
        output := &frauddetector.DeleteEntityTypeOutput{}

        mockClient.On("DeleteEntityType", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteEntityType(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteEvent", func(t *testing.T) {
        input := &frauddetector.DeleteEventInput{}
        output := &frauddetector.DeleteEventOutput{}

        mockClient.On("DeleteEvent", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteEvent(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteEventType", func(t *testing.T) {
        input := &frauddetector.DeleteEventTypeInput{}
        output := &frauddetector.DeleteEventTypeOutput{}

        mockClient.On("DeleteEventType", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteEventType(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteEventsByEventType", func(t *testing.T) {
        input := &frauddetector.DeleteEventsByEventTypeInput{}
        output := &frauddetector.DeleteEventsByEventTypeOutput{}

        mockClient.On("DeleteEventsByEventType", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteEventsByEventType(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteExternalModel", func(t *testing.T) {
        input := &frauddetector.DeleteExternalModelInput{}
        output := &frauddetector.DeleteExternalModelOutput{}

        mockClient.On("DeleteExternalModel", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteExternalModel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteLabel", func(t *testing.T) {
        input := &frauddetector.DeleteLabelInput{}
        output := &frauddetector.DeleteLabelOutput{}

        mockClient.On("DeleteLabel", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteLabel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteList", func(t *testing.T) {
        input := &frauddetector.DeleteListInput{}
        output := &frauddetector.DeleteListOutput{}

        mockClient.On("DeleteList", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteList(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteModel", func(t *testing.T) {
        input := &frauddetector.DeleteModelInput{}
        output := &frauddetector.DeleteModelOutput{}

        mockClient.On("DeleteModel", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteModel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteModelVersion", func(t *testing.T) {
        input := &frauddetector.DeleteModelVersionInput{}
        output := &frauddetector.DeleteModelVersionOutput{}

        mockClient.On("DeleteModelVersion", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteModelVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteOutcome", func(t *testing.T) {
        input := &frauddetector.DeleteOutcomeInput{}
        output := &frauddetector.DeleteOutcomeOutput{}

        mockClient.On("DeleteOutcome", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteOutcome(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteRule", func(t *testing.T) {
        input := &frauddetector.DeleteRuleInput{}
        output := &frauddetector.DeleteRuleOutput{}

        mockClient.On("DeleteRule", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteRule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteVariable", func(t *testing.T) {
        input := &frauddetector.DeleteVariableInput{}
        output := &frauddetector.DeleteVariableOutput{}

        mockClient.On("DeleteVariable", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteVariable(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeDetector", func(t *testing.T) {
        input := &frauddetector.DescribeDetectorInput{}
        output := &frauddetector.DescribeDetectorOutput{}

        mockClient.On("DescribeDetector", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeDetector(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeModelVersions", func(t *testing.T) {
        input := &frauddetector.DescribeModelVersionsInput{}
        output := &frauddetector.DescribeModelVersionsOutput{}

        mockClient.On("DescribeModelVersions", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeModelVersions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetBatchImportJobs", func(t *testing.T) {
        input := &frauddetector.GetBatchImportJobsInput{}
        output := &frauddetector.GetBatchImportJobsOutput{}

        mockClient.On("GetBatchImportJobs", ctx, input).Return(output, nil)

        result, err := mockClient.GetBatchImportJobs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetBatchPredictionJobs", func(t *testing.T) {
        input := &frauddetector.GetBatchPredictionJobsInput{}
        output := &frauddetector.GetBatchPredictionJobsOutput{}

        mockClient.On("GetBatchPredictionJobs", ctx, input).Return(output, nil)

        result, err := mockClient.GetBatchPredictionJobs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDeleteEventsByEventTypeStatus", func(t *testing.T) {
        input := &frauddetector.GetDeleteEventsByEventTypeStatusInput{}
        output := &frauddetector.GetDeleteEventsByEventTypeStatusOutput{}

        mockClient.On("GetDeleteEventsByEventTypeStatus", ctx, input).Return(output, nil)

        result, err := mockClient.GetDeleteEventsByEventTypeStatus(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDetectorVersion", func(t *testing.T) {
        input := &frauddetector.GetDetectorVersionInput{}
        output := &frauddetector.GetDetectorVersionOutput{}

        mockClient.On("GetDetectorVersion", ctx, input).Return(output, nil)

        result, err := mockClient.GetDetectorVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDetectors", func(t *testing.T) {
        input := &frauddetector.GetDetectorsInput{}
        output := &frauddetector.GetDetectorsOutput{}

        mockClient.On("GetDetectors", ctx, input).Return(output, nil)

        result, err := mockClient.GetDetectors(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetEntityTypes", func(t *testing.T) {
        input := &frauddetector.GetEntityTypesInput{}
        output := &frauddetector.GetEntityTypesOutput{}

        mockClient.On("GetEntityTypes", ctx, input).Return(output, nil)

        result, err := mockClient.GetEntityTypes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetEvent", func(t *testing.T) {
        input := &frauddetector.GetEventInput{}
        output := &frauddetector.GetEventOutput{}

        mockClient.On("GetEvent", ctx, input).Return(output, nil)

        result, err := mockClient.GetEvent(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetEventPrediction", func(t *testing.T) {
        input := &frauddetector.GetEventPredictionInput{}
        output := &frauddetector.GetEventPredictionOutput{}

        mockClient.On("GetEventPrediction", ctx, input).Return(output, nil)

        result, err := mockClient.GetEventPrediction(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetEventPredictionMetadata", func(t *testing.T) {
        input := &frauddetector.GetEventPredictionMetadataInput{}
        output := &frauddetector.GetEventPredictionMetadataOutput{}

        mockClient.On("GetEventPredictionMetadata", ctx, input).Return(output, nil)

        result, err := mockClient.GetEventPredictionMetadata(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetEventTypes", func(t *testing.T) {
        input := &frauddetector.GetEventTypesInput{}
        output := &frauddetector.GetEventTypesOutput{}

        mockClient.On("GetEventTypes", ctx, input).Return(output, nil)

        result, err := mockClient.GetEventTypes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetExternalModels", func(t *testing.T) {
        input := &frauddetector.GetExternalModelsInput{}
        output := &frauddetector.GetExternalModelsOutput{}

        mockClient.On("GetExternalModels", ctx, input).Return(output, nil)

        result, err := mockClient.GetExternalModels(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetKMSEncryptionKey", func(t *testing.T) {
        input := &frauddetector.GetKMSEncryptionKeyInput{}
        output := &frauddetector.GetKMSEncryptionKeyOutput{}

        mockClient.On("GetKMSEncryptionKey", ctx, input).Return(output, nil)

        result, err := mockClient.GetKMSEncryptionKey(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetLabels", func(t *testing.T) {
        input := &frauddetector.GetLabelsInput{}
        output := &frauddetector.GetLabelsOutput{}

        mockClient.On("GetLabels", ctx, input).Return(output, nil)

        result, err := mockClient.GetLabels(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetListElements", func(t *testing.T) {
        input := &frauddetector.GetListElementsInput{}
        output := &frauddetector.GetListElementsOutput{}

        mockClient.On("GetListElements", ctx, input).Return(output, nil)

        result, err := mockClient.GetListElements(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetListsMetadata", func(t *testing.T) {
        input := &frauddetector.GetListsMetadataInput{}
        output := &frauddetector.GetListsMetadataOutput{}

        mockClient.On("GetListsMetadata", ctx, input).Return(output, nil)

        result, err := mockClient.GetListsMetadata(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetModelVersion", func(t *testing.T) {
        input := &frauddetector.GetModelVersionInput{}
        output := &frauddetector.GetModelVersionOutput{}

        mockClient.On("GetModelVersion", ctx, input).Return(output, nil)

        result, err := mockClient.GetModelVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetModels", func(t *testing.T) {
        input := &frauddetector.GetModelsInput{}
        output := &frauddetector.GetModelsOutput{}

        mockClient.On("GetModels", ctx, input).Return(output, nil)

        result, err := mockClient.GetModels(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetOutcomes", func(t *testing.T) {
        input := &frauddetector.GetOutcomesInput{}
        output := &frauddetector.GetOutcomesOutput{}

        mockClient.On("GetOutcomes", ctx, input).Return(output, nil)

        result, err := mockClient.GetOutcomes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetRules", func(t *testing.T) {
        input := &frauddetector.GetRulesInput{}
        output := &frauddetector.GetRulesOutput{}

        mockClient.On("GetRules", ctx, input).Return(output, nil)

        result, err := mockClient.GetRules(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetVariables", func(t *testing.T) {
        input := &frauddetector.GetVariablesInput{}
        output := &frauddetector.GetVariablesOutput{}

        mockClient.On("GetVariables", ctx, input).Return(output, nil)

        result, err := mockClient.GetVariables(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListEventPredictions", func(t *testing.T) {
        input := &frauddetector.ListEventPredictionsInput{}
        output := &frauddetector.ListEventPredictionsOutput{}

        mockClient.On("ListEventPredictions", ctx, input).Return(output, nil)

        result, err := mockClient.ListEventPredictions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &frauddetector.ListTagsForResourceInput{}
        output := &frauddetector.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutDetector", func(t *testing.T) {
        input := &frauddetector.PutDetectorInput{}
        output := &frauddetector.PutDetectorOutput{}

        mockClient.On("PutDetector", ctx, input).Return(output, nil)

        result, err := mockClient.PutDetector(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutEntityType", func(t *testing.T) {
        input := &frauddetector.PutEntityTypeInput{}
        output := &frauddetector.PutEntityTypeOutput{}

        mockClient.On("PutEntityType", ctx, input).Return(output, nil)

        result, err := mockClient.PutEntityType(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutEventType", func(t *testing.T) {
        input := &frauddetector.PutEventTypeInput{}
        output := &frauddetector.PutEventTypeOutput{}

        mockClient.On("PutEventType", ctx, input).Return(output, nil)

        result, err := mockClient.PutEventType(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutExternalModel", func(t *testing.T) {
        input := &frauddetector.PutExternalModelInput{}
        output := &frauddetector.PutExternalModelOutput{}

        mockClient.On("PutExternalModel", ctx, input).Return(output, nil)

        result, err := mockClient.PutExternalModel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutKMSEncryptionKey", func(t *testing.T) {
        input := &frauddetector.PutKMSEncryptionKeyInput{}
        output := &frauddetector.PutKMSEncryptionKeyOutput{}

        mockClient.On("PutKMSEncryptionKey", ctx, input).Return(output, nil)

        result, err := mockClient.PutKMSEncryptionKey(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutLabel", func(t *testing.T) {
        input := &frauddetector.PutLabelInput{}
        output := &frauddetector.PutLabelOutput{}

        mockClient.On("PutLabel", ctx, input).Return(output, nil)

        result, err := mockClient.PutLabel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutOutcome", func(t *testing.T) {
        input := &frauddetector.PutOutcomeInput{}
        output := &frauddetector.PutOutcomeOutput{}

        mockClient.On("PutOutcome", ctx, input).Return(output, nil)

        result, err := mockClient.PutOutcome(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSendEvent", func(t *testing.T) {
        input := &frauddetector.SendEventInput{}
        output := &frauddetector.SendEventOutput{}

        mockClient.On("SendEvent", ctx, input).Return(output, nil)

        result, err := mockClient.SendEvent(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &frauddetector.TagResourceInput{}
        output := &frauddetector.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &frauddetector.UntagResourceInput{}
        output := &frauddetector.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateDetectorVersion", func(t *testing.T) {
        input := &frauddetector.UpdateDetectorVersionInput{}
        output := &frauddetector.UpdateDetectorVersionOutput{}

        mockClient.On("UpdateDetectorVersion", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateDetectorVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateDetectorVersionMetadata", func(t *testing.T) {
        input := &frauddetector.UpdateDetectorVersionMetadataInput{}
        output := &frauddetector.UpdateDetectorVersionMetadataOutput{}

        mockClient.On("UpdateDetectorVersionMetadata", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateDetectorVersionMetadata(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateDetectorVersionStatus", func(t *testing.T) {
        input := &frauddetector.UpdateDetectorVersionStatusInput{}
        output := &frauddetector.UpdateDetectorVersionStatusOutput{}

        mockClient.On("UpdateDetectorVersionStatus", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateDetectorVersionStatus(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateEventLabel", func(t *testing.T) {
        input := &frauddetector.UpdateEventLabelInput{}
        output := &frauddetector.UpdateEventLabelOutput{}

        mockClient.On("UpdateEventLabel", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateEventLabel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateList", func(t *testing.T) {
        input := &frauddetector.UpdateListInput{}
        output := &frauddetector.UpdateListOutput{}

        mockClient.On("UpdateList", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateList(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateModel", func(t *testing.T) {
        input := &frauddetector.UpdateModelInput{}
        output := &frauddetector.UpdateModelOutput{}

        mockClient.On("UpdateModel", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateModel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateModelVersion", func(t *testing.T) {
        input := &frauddetector.UpdateModelVersionInput{}
        output := &frauddetector.UpdateModelVersionOutput{}

        mockClient.On("UpdateModelVersion", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateModelVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateModelVersionStatus", func(t *testing.T) {
        input := &frauddetector.UpdateModelVersionStatusInput{}
        output := &frauddetector.UpdateModelVersionStatusOutput{}

        mockClient.On("UpdateModelVersionStatus", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateModelVersionStatus(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateRuleMetadata", func(t *testing.T) {
        input := &frauddetector.UpdateRuleMetadataInput{}
        output := &frauddetector.UpdateRuleMetadataOutput{}

        mockClient.On("UpdateRuleMetadata", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateRuleMetadata(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateRuleVersion", func(t *testing.T) {
        input := &frauddetector.UpdateRuleVersionInput{}
        output := &frauddetector.UpdateRuleVersionOutput{}

        mockClient.On("UpdateRuleVersion", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateRuleVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateVariable", func(t *testing.T) {
        input := &frauddetector.UpdateVariableInput{}
        output := &frauddetector.UpdateVariableOutput{}

        mockClient.On("UpdateVariable", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateVariable(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
