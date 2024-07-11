
package transcribe

import (
    "github.com/aws/aws-sdk-go-v2/service/transcribe"
)

// ITranscribe defines the interface for transcribe
type ITranscribe interface {
 Options() Options 
 CreateCallAnalyticsCategory(ctx context.Context, params *CreateCallAnalyticsCategoryInput, optFns ...func(*Options)) (*CreateCallAnalyticsCategoryOutput, error) 
 CreateLanguageModel(ctx context.Context, params *CreateLanguageModelInput, optFns ...func(*Options)) (*CreateLanguageModelOutput, error) 
 CreateMedicalVocabulary(ctx context.Context, params *CreateMedicalVocabularyInput, optFns ...func(*Options)) (*CreateMedicalVocabularyOutput, error) 
 CreateVocabulary(ctx context.Context, params *CreateVocabularyInput, optFns ...func(*Options)) (*CreateVocabularyOutput, error) 
 CreateVocabularyFilter(ctx context.Context, params *CreateVocabularyFilterInput, optFns ...func(*Options)) (*CreateVocabularyFilterOutput, error) 
 DeleteCallAnalyticsCategory(ctx context.Context, params *DeleteCallAnalyticsCategoryInput, optFns ...func(*Options)) (*DeleteCallAnalyticsCategoryOutput, error) 
 DeleteCallAnalyticsJob(ctx context.Context, params *DeleteCallAnalyticsJobInput, optFns ...func(*Options)) (*DeleteCallAnalyticsJobOutput, error) 
 DeleteLanguageModel(ctx context.Context, params *DeleteLanguageModelInput, optFns ...func(*Options)) (*DeleteLanguageModelOutput, error) 
 DeleteMedicalScribeJob(ctx context.Context, params *DeleteMedicalScribeJobInput, optFns ...func(*Options)) (*DeleteMedicalScribeJobOutput, error) 
 DeleteMedicalTranscriptionJob(ctx context.Context, params *DeleteMedicalTranscriptionJobInput, optFns ...func(*Options)) (*DeleteMedicalTranscriptionJobOutput, error) 
 DeleteMedicalVocabulary(ctx context.Context, params *DeleteMedicalVocabularyInput, optFns ...func(*Options)) (*DeleteMedicalVocabularyOutput, error) 
 DeleteTranscriptionJob(ctx context.Context, params *DeleteTranscriptionJobInput, optFns ...func(*Options)) (*DeleteTranscriptionJobOutput, error) 
 DeleteVocabulary(ctx context.Context, params *DeleteVocabularyInput, optFns ...func(*Options)) (*DeleteVocabularyOutput, error) 
 DeleteVocabularyFilter(ctx context.Context, params *DeleteVocabularyFilterInput, optFns ...func(*Options)) (*DeleteVocabularyFilterOutput, error) 
 DescribeLanguageModel(ctx context.Context, params *DescribeLanguageModelInput, optFns ...func(*Options)) (*DescribeLanguageModelOutput, error) 
 GetCallAnalyticsCategory(ctx context.Context, params *GetCallAnalyticsCategoryInput, optFns ...func(*Options)) (*GetCallAnalyticsCategoryOutput, error) 
 GetCallAnalyticsJob(ctx context.Context, params *GetCallAnalyticsJobInput, optFns ...func(*Options)) (*GetCallAnalyticsJobOutput, error) 
 GetMedicalScribeJob(ctx context.Context, params *GetMedicalScribeJobInput, optFns ...func(*Options)) (*GetMedicalScribeJobOutput, error) 
 GetMedicalTranscriptionJob(ctx context.Context, params *GetMedicalTranscriptionJobInput, optFns ...func(*Options)) (*GetMedicalTranscriptionJobOutput, error) 
 GetMedicalVocabulary(ctx context.Context, params *GetMedicalVocabularyInput, optFns ...func(*Options)) (*GetMedicalVocabularyOutput, error) 
 GetTranscriptionJob(ctx context.Context, params *GetTranscriptionJobInput, optFns ...func(*Options)) (*GetTranscriptionJobOutput, error) 
 GetVocabulary(ctx context.Context, params *GetVocabularyInput, optFns ...func(*Options)) (*GetVocabularyOutput, error) 
 GetVocabularyFilter(ctx context.Context, params *GetVocabularyFilterInput, optFns ...func(*Options)) (*GetVocabularyFilterOutput, error) 
 ListCallAnalyticsCategories(ctx context.Context, params *ListCallAnalyticsCategoriesInput, optFns ...func(*Options)) (*ListCallAnalyticsCategoriesOutput, error) 
 ListCallAnalyticsJobs(ctx context.Context, params *ListCallAnalyticsJobsInput, optFns ...func(*Options)) (*ListCallAnalyticsJobsOutput, error) 
 ListLanguageModels(ctx context.Context, params *ListLanguageModelsInput, optFns ...func(*Options)) (*ListLanguageModelsOutput, error) 
 ListMedicalScribeJobs(ctx context.Context, params *ListMedicalScribeJobsInput, optFns ...func(*Options)) (*ListMedicalScribeJobsOutput, error) 
 ListMedicalTranscriptionJobs(ctx context.Context, params *ListMedicalTranscriptionJobsInput, optFns ...func(*Options)) (*ListMedicalTranscriptionJobsOutput, error) 
 ListMedicalVocabularies(ctx context.Context, params *ListMedicalVocabulariesInput, optFns ...func(*Options)) (*ListMedicalVocabulariesOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 ListTranscriptionJobs(ctx context.Context, params *ListTranscriptionJobsInput, optFns ...func(*Options)) (*ListTranscriptionJobsOutput, error) 
 ListVocabularies(ctx context.Context, params *ListVocabulariesInput, optFns ...func(*Options)) (*ListVocabulariesOutput, error) 
 ListVocabularyFilters(ctx context.Context, params *ListVocabularyFiltersInput, optFns ...func(*Options)) (*ListVocabularyFiltersOutput, error) 
 StartCallAnalyticsJob(ctx context.Context, params *StartCallAnalyticsJobInput, optFns ...func(*Options)) (*StartCallAnalyticsJobOutput, error) 
 StartMedicalScribeJob(ctx context.Context, params *StartMedicalScribeJobInput, optFns ...func(*Options)) (*StartMedicalScribeJobOutput, error) 
 StartMedicalTranscriptionJob(ctx context.Context, params *StartMedicalTranscriptionJobInput, optFns ...func(*Options)) (*StartMedicalTranscriptionJobOutput, error) 
 StartTranscriptionJob(ctx context.Context, params *StartTranscriptionJobInput, optFns ...func(*Options)) (*StartTranscriptionJobOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateCallAnalyticsCategory(ctx context.Context, params *UpdateCallAnalyticsCategoryInput, optFns ...func(*Options)) (*UpdateCallAnalyticsCategoryOutput, error) 
 UpdateMedicalVocabulary(ctx context.Context, params *UpdateMedicalVocabularyInput, optFns ...func(*Options)) (*UpdateMedicalVocabularyOutput, error) 
 UpdateVocabulary(ctx context.Context, params *UpdateVocabularyInput, optFns ...func(*Options)) (*UpdateVocabularyOutput, error) 
 UpdateVocabularyFilter(ctx context.Context, params *UpdateVocabularyFilterInput, optFns ...func(*Options)) (*UpdateVocabularyFilterOutput, error) 
}
