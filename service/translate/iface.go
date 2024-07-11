
package translate

import (
    "github.com/aws/aws-sdk-go-v2/service/translate"
)

// ITranslate defines the interface for translate
type ITranslate interface {
 Options() Options 
 CreateParallelData(ctx context.Context, params *CreateParallelDataInput, optFns ...func(*Options)) (*CreateParallelDataOutput, error) 
 DeleteParallelData(ctx context.Context, params *DeleteParallelDataInput, optFns ...func(*Options)) (*DeleteParallelDataOutput, error) 
 DeleteTerminology(ctx context.Context, params *DeleteTerminologyInput, optFns ...func(*Options)) (*DeleteTerminologyOutput, error) 
 DescribeTextTranslationJob(ctx context.Context, params *DescribeTextTranslationJobInput, optFns ...func(*Options)) (*DescribeTextTranslationJobOutput, error) 
 GetParallelData(ctx context.Context, params *GetParallelDataInput, optFns ...func(*Options)) (*GetParallelDataOutput, error) 
 GetTerminology(ctx context.Context, params *GetTerminologyInput, optFns ...func(*Options)) (*GetTerminologyOutput, error) 
 ImportTerminology(ctx context.Context, params *ImportTerminologyInput, optFns ...func(*Options)) (*ImportTerminologyOutput, error) 
 ListLanguages(ctx context.Context, params *ListLanguagesInput, optFns ...func(*Options)) (*ListLanguagesOutput, error) 
 ListParallelData(ctx context.Context, params *ListParallelDataInput, optFns ...func(*Options)) (*ListParallelDataOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 ListTerminologies(ctx context.Context, params *ListTerminologiesInput, optFns ...func(*Options)) (*ListTerminologiesOutput, error) 
 ListTextTranslationJobs(ctx context.Context, params *ListTextTranslationJobsInput, optFns ...func(*Options)) (*ListTextTranslationJobsOutput, error) 
 StartTextTranslationJob(ctx context.Context, params *StartTextTranslationJobInput, optFns ...func(*Options)) (*StartTextTranslationJobOutput, error) 
 StopTextTranslationJob(ctx context.Context, params *StopTextTranslationJobInput, optFns ...func(*Options)) (*StopTextTranslationJobOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 TranslateDocument(ctx context.Context, params *TranslateDocumentInput, optFns ...func(*Options)) (*TranslateDocumentOutput, error) 
 TranslateText(ctx context.Context, params *TranslateTextInput, optFns ...func(*Options)) (*TranslateTextOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateParallelData(ctx context.Context, params *UpdateParallelDataInput, optFns ...func(*Options)) (*UpdateParallelDataOutput, error) 
}
