
package frauddetector_iface

// This file was generated by scripts/gen_ifaces.sh do not edit it manually
import (
    "context"
    . "github.com/aws/aws-sdk-go-v2/service/frauddetector"
)

// IClient defines the interface for frauddetector
type IClient interface {
 Options() Options 
 BatchCreateVariable(ctx context.Context, params *BatchCreateVariableInput, optFns ...func(*Options)) (*BatchCreateVariableOutput, error) 
 BatchGetVariable(ctx context.Context, params *BatchGetVariableInput, optFns ...func(*Options)) (*BatchGetVariableOutput, error) 
 CancelBatchImportJob(ctx context.Context, params *CancelBatchImportJobInput, optFns ...func(*Options)) (*CancelBatchImportJobOutput, error) 
 CancelBatchPredictionJob(ctx context.Context, params *CancelBatchPredictionJobInput, optFns ...func(*Options)) (*CancelBatchPredictionJobOutput, error) 
 CreateBatchImportJob(ctx context.Context, params *CreateBatchImportJobInput, optFns ...func(*Options)) (*CreateBatchImportJobOutput, error) 
 CreateBatchPredictionJob(ctx context.Context, params *CreateBatchPredictionJobInput, optFns ...func(*Options)) (*CreateBatchPredictionJobOutput, error) 
 CreateDetectorVersion(ctx context.Context, params *CreateDetectorVersionInput, optFns ...func(*Options)) (*CreateDetectorVersionOutput, error) 
 CreateList(ctx context.Context, params *CreateListInput, optFns ...func(*Options)) (*CreateListOutput, error) 
 CreateModel(ctx context.Context, params *CreateModelInput, optFns ...func(*Options)) (*CreateModelOutput, error) 
 CreateModelVersion(ctx context.Context, params *CreateModelVersionInput, optFns ...func(*Options)) (*CreateModelVersionOutput, error) 
 CreateRule(ctx context.Context, params *CreateRuleInput, optFns ...func(*Options)) (*CreateRuleOutput, error) 
 CreateVariable(ctx context.Context, params *CreateVariableInput, optFns ...func(*Options)) (*CreateVariableOutput, error) 
 DeleteBatchImportJob(ctx context.Context, params *DeleteBatchImportJobInput, optFns ...func(*Options)) (*DeleteBatchImportJobOutput, error) 
 DeleteBatchPredictionJob(ctx context.Context, params *DeleteBatchPredictionJobInput, optFns ...func(*Options)) (*DeleteBatchPredictionJobOutput, error) 
 DeleteDetector(ctx context.Context, params *DeleteDetectorInput, optFns ...func(*Options)) (*DeleteDetectorOutput, error) 
 DeleteDetectorVersion(ctx context.Context, params *DeleteDetectorVersionInput, optFns ...func(*Options)) (*DeleteDetectorVersionOutput, error) 
 DeleteEntityType(ctx context.Context, params *DeleteEntityTypeInput, optFns ...func(*Options)) (*DeleteEntityTypeOutput, error) 
 DeleteEvent(ctx context.Context, params *DeleteEventInput, optFns ...func(*Options)) (*DeleteEventOutput, error) 
 DeleteEventType(ctx context.Context, params *DeleteEventTypeInput, optFns ...func(*Options)) (*DeleteEventTypeOutput, error) 
 DeleteEventsByEventType(ctx context.Context, params *DeleteEventsByEventTypeInput, optFns ...func(*Options)) (*DeleteEventsByEventTypeOutput, error) 
 DeleteExternalModel(ctx context.Context, params *DeleteExternalModelInput, optFns ...func(*Options)) (*DeleteExternalModelOutput, error) 
 DeleteLabel(ctx context.Context, params *DeleteLabelInput, optFns ...func(*Options)) (*DeleteLabelOutput, error) 
 DeleteList(ctx context.Context, params *DeleteListInput, optFns ...func(*Options)) (*DeleteListOutput, error) 
 DeleteModel(ctx context.Context, params *DeleteModelInput, optFns ...func(*Options)) (*DeleteModelOutput, error) 
 DeleteModelVersion(ctx context.Context, params *DeleteModelVersionInput, optFns ...func(*Options)) (*DeleteModelVersionOutput, error) 
 DeleteOutcome(ctx context.Context, params *DeleteOutcomeInput, optFns ...func(*Options)) (*DeleteOutcomeOutput, error) 
 DeleteRule(ctx context.Context, params *DeleteRuleInput, optFns ...func(*Options)) (*DeleteRuleOutput, error) 
 DeleteVariable(ctx context.Context, params *DeleteVariableInput, optFns ...func(*Options)) (*DeleteVariableOutput, error) 
 DescribeDetector(ctx context.Context, params *DescribeDetectorInput, optFns ...func(*Options)) (*DescribeDetectorOutput, error) 
 DescribeModelVersions(ctx context.Context, params *DescribeModelVersionsInput, optFns ...func(*Options)) (*DescribeModelVersionsOutput, error) 
 GetBatchImportJobs(ctx context.Context, params *GetBatchImportJobsInput, optFns ...func(*Options)) (*GetBatchImportJobsOutput, error) 
 GetBatchPredictionJobs(ctx context.Context, params *GetBatchPredictionJobsInput, optFns ...func(*Options)) (*GetBatchPredictionJobsOutput, error) 
 GetDeleteEventsByEventTypeStatus(ctx context.Context, params *GetDeleteEventsByEventTypeStatusInput, optFns ...func(*Options)) (*GetDeleteEventsByEventTypeStatusOutput, error) 
 GetDetectorVersion(ctx context.Context, params *GetDetectorVersionInput, optFns ...func(*Options)) (*GetDetectorVersionOutput, error) 
 GetDetectors(ctx context.Context, params *GetDetectorsInput, optFns ...func(*Options)) (*GetDetectorsOutput, error) 
 GetEntityTypes(ctx context.Context, params *GetEntityTypesInput, optFns ...func(*Options)) (*GetEntityTypesOutput, error) 
 GetEvent(ctx context.Context, params *GetEventInput, optFns ...func(*Options)) (*GetEventOutput, error) 
 GetEventPrediction(ctx context.Context, params *GetEventPredictionInput, optFns ...func(*Options)) (*GetEventPredictionOutput, error) 
 GetEventPredictionMetadata(ctx context.Context, params *GetEventPredictionMetadataInput, optFns ...func(*Options)) (*GetEventPredictionMetadataOutput, error) 
 GetEventTypes(ctx context.Context, params *GetEventTypesInput, optFns ...func(*Options)) (*GetEventTypesOutput, error) 
 GetExternalModels(ctx context.Context, params *GetExternalModelsInput, optFns ...func(*Options)) (*GetExternalModelsOutput, error) 
 GetKMSEncryptionKey(ctx context.Context, params *GetKMSEncryptionKeyInput, optFns ...func(*Options)) (*GetKMSEncryptionKeyOutput, error) 
 GetLabels(ctx context.Context, params *GetLabelsInput, optFns ...func(*Options)) (*GetLabelsOutput, error) 
 GetListElements(ctx context.Context, params *GetListElementsInput, optFns ...func(*Options)) (*GetListElementsOutput, error) 
 GetListsMetadata(ctx context.Context, params *GetListsMetadataInput, optFns ...func(*Options)) (*GetListsMetadataOutput, error) 
 GetModelVersion(ctx context.Context, params *GetModelVersionInput, optFns ...func(*Options)) (*GetModelVersionOutput, error) 
 GetModels(ctx context.Context, params *GetModelsInput, optFns ...func(*Options)) (*GetModelsOutput, error) 
 GetOutcomes(ctx context.Context, params *GetOutcomesInput, optFns ...func(*Options)) (*GetOutcomesOutput, error) 
 GetRules(ctx context.Context, params *GetRulesInput, optFns ...func(*Options)) (*GetRulesOutput, error) 
 GetVariables(ctx context.Context, params *GetVariablesInput, optFns ...func(*Options)) (*GetVariablesOutput, error) 
 ListEventPredictions(ctx context.Context, params *ListEventPredictionsInput, optFns ...func(*Options)) (*ListEventPredictionsOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 PutDetector(ctx context.Context, params *PutDetectorInput, optFns ...func(*Options)) (*PutDetectorOutput, error) 
 PutEntityType(ctx context.Context, params *PutEntityTypeInput, optFns ...func(*Options)) (*PutEntityTypeOutput, error) 
 PutEventType(ctx context.Context, params *PutEventTypeInput, optFns ...func(*Options)) (*PutEventTypeOutput, error) 
 PutExternalModel(ctx context.Context, params *PutExternalModelInput, optFns ...func(*Options)) (*PutExternalModelOutput, error) 
 PutKMSEncryptionKey(ctx context.Context, params *PutKMSEncryptionKeyInput, optFns ...func(*Options)) (*PutKMSEncryptionKeyOutput, error) 
 PutLabel(ctx context.Context, params *PutLabelInput, optFns ...func(*Options)) (*PutLabelOutput, error) 
 PutOutcome(ctx context.Context, params *PutOutcomeInput, optFns ...func(*Options)) (*PutOutcomeOutput, error) 
 SendEvent(ctx context.Context, params *SendEventInput, optFns ...func(*Options)) (*SendEventOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateDetectorVersion(ctx context.Context, params *UpdateDetectorVersionInput, optFns ...func(*Options)) (*UpdateDetectorVersionOutput, error) 
 UpdateDetectorVersionMetadata(ctx context.Context, params *UpdateDetectorVersionMetadataInput, optFns ...func(*Options)) (*UpdateDetectorVersionMetadataOutput, error) 
 UpdateDetectorVersionStatus(ctx context.Context, params *UpdateDetectorVersionStatusInput, optFns ...func(*Options)) (*UpdateDetectorVersionStatusOutput, error) 
 UpdateEventLabel(ctx context.Context, params *UpdateEventLabelInput, optFns ...func(*Options)) (*UpdateEventLabelOutput, error) 
 UpdateList(ctx context.Context, params *UpdateListInput, optFns ...func(*Options)) (*UpdateListOutput, error) 
 UpdateModel(ctx context.Context, params *UpdateModelInput, optFns ...func(*Options)) (*UpdateModelOutput, error) 
 UpdateModelVersion(ctx context.Context, params *UpdateModelVersionInput, optFns ...func(*Options)) (*UpdateModelVersionOutput, error) 
 UpdateModelVersionStatus(ctx context.Context, params *UpdateModelVersionStatusInput, optFns ...func(*Options)) (*UpdateModelVersionStatusOutput, error) 
 UpdateRuleMetadata(ctx context.Context, params *UpdateRuleMetadataInput, optFns ...func(*Options)) (*UpdateRuleMetadataOutput, error) 
 UpdateRuleVersion(ctx context.Context, params *UpdateRuleVersionInput, optFns ...func(*Options)) (*UpdateRuleVersionOutput, error) 
 UpdateVariable(ctx context.Context, params *UpdateVariableInput, optFns ...func(*Options)) (*UpdateVariableOutput, error) 
}
