
package iotevents

import (
    "github.com/aws/aws-sdk-go-v2/service/iotevents"
)

// IClient defines the interface for iotevents
type IClient interface {
 Options() Options 
 CreateAlarmModel(ctx context.Context, params *CreateAlarmModelInput, optFns ...func(*Options)) (*CreateAlarmModelOutput, error) 
 CreateDetectorModel(ctx context.Context, params *CreateDetectorModelInput, optFns ...func(*Options)) (*CreateDetectorModelOutput, error) 
 CreateInput(ctx context.Context, params *CreateInputInput, optFns ...func(*Options)) (*CreateInputOutput, error) 
 DeleteAlarmModel(ctx context.Context, params *DeleteAlarmModelInput, optFns ...func(*Options)) (*DeleteAlarmModelOutput, error) 
 DeleteDetectorModel(ctx context.Context, params *DeleteDetectorModelInput, optFns ...func(*Options)) (*DeleteDetectorModelOutput, error) 
 DeleteInput(ctx context.Context, params *DeleteInputInput, optFns ...func(*Options)) (*DeleteInputOutput, error) 
 DescribeAlarmModel(ctx context.Context, params *DescribeAlarmModelInput, optFns ...func(*Options)) (*DescribeAlarmModelOutput, error) 
 DescribeDetectorModel(ctx context.Context, params *DescribeDetectorModelInput, optFns ...func(*Options)) (*DescribeDetectorModelOutput, error) 
 DescribeDetectorModelAnalysis(ctx context.Context, params *DescribeDetectorModelAnalysisInput, optFns ...func(*Options)) (*DescribeDetectorModelAnalysisOutput, error) 
 DescribeInput(ctx context.Context, params *DescribeInputInput, optFns ...func(*Options)) (*DescribeInputOutput, error) 
 DescribeLoggingOptions(ctx context.Context, params *DescribeLoggingOptionsInput, optFns ...func(*Options)) (*DescribeLoggingOptionsOutput, error) 
 GetDetectorModelAnalysisResults(ctx context.Context, params *GetDetectorModelAnalysisResultsInput, optFns ...func(*Options)) (*GetDetectorModelAnalysisResultsOutput, error) 
 ListAlarmModelVersions(ctx context.Context, params *ListAlarmModelVersionsInput, optFns ...func(*Options)) (*ListAlarmModelVersionsOutput, error) 
 ListAlarmModels(ctx context.Context, params *ListAlarmModelsInput, optFns ...func(*Options)) (*ListAlarmModelsOutput, error) 
 ListDetectorModelVersions(ctx context.Context, params *ListDetectorModelVersionsInput, optFns ...func(*Options)) (*ListDetectorModelVersionsOutput, error) 
 ListDetectorModels(ctx context.Context, params *ListDetectorModelsInput, optFns ...func(*Options)) (*ListDetectorModelsOutput, error) 
 ListInputRoutings(ctx context.Context, params *ListInputRoutingsInput, optFns ...func(*Options)) (*ListInputRoutingsOutput, error) 
 ListInputs(ctx context.Context, params *ListInputsInput, optFns ...func(*Options)) (*ListInputsOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 PutLoggingOptions(ctx context.Context, params *PutLoggingOptionsInput, optFns ...func(*Options)) (*PutLoggingOptionsOutput, error) 
 StartDetectorModelAnalysis(ctx context.Context, params *StartDetectorModelAnalysisInput, optFns ...func(*Options)) (*StartDetectorModelAnalysisOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateAlarmModel(ctx context.Context, params *UpdateAlarmModelInput, optFns ...func(*Options)) (*UpdateAlarmModelOutput, error) 
 UpdateDetectorModel(ctx context.Context, params *UpdateDetectorModelInput, optFns ...func(*Options)) (*UpdateDetectorModelOutput, error) 
 UpdateInput(ctx context.Context, params *UpdateInputInput, optFns ...func(*Options)) (*UpdateInputOutput, error) 
}
