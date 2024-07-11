
package textract

import (
    "github.com/aws/aws-sdk-go-v2/service/textract"
)

// ITextract defines the interface for textract
type ITextract interface {
 Options() Options 
 AnalyzeDocument(ctx context.Context, params *AnalyzeDocumentInput, optFns ...func(*Options)) (*AnalyzeDocumentOutput, error) 
 AnalyzeExpense(ctx context.Context, params *AnalyzeExpenseInput, optFns ...func(*Options)) (*AnalyzeExpenseOutput, error) 
 AnalyzeID(ctx context.Context, params *AnalyzeIDInput, optFns ...func(*Options)) (*AnalyzeIDOutput, error) 
 CreateAdapter(ctx context.Context, params *CreateAdapterInput, optFns ...func(*Options)) (*CreateAdapterOutput, error) 
 CreateAdapterVersion(ctx context.Context, params *CreateAdapterVersionInput, optFns ...func(*Options)) (*CreateAdapterVersionOutput, error) 
 DeleteAdapter(ctx context.Context, params *DeleteAdapterInput, optFns ...func(*Options)) (*DeleteAdapterOutput, error) 
 DeleteAdapterVersion(ctx context.Context, params *DeleteAdapterVersionInput, optFns ...func(*Options)) (*DeleteAdapterVersionOutput, error) 
 DetectDocumentText(ctx context.Context, params *DetectDocumentTextInput, optFns ...func(*Options)) (*DetectDocumentTextOutput, error) 
 GetAdapter(ctx context.Context, params *GetAdapterInput, optFns ...func(*Options)) (*GetAdapterOutput, error) 
 GetAdapterVersion(ctx context.Context, params *GetAdapterVersionInput, optFns ...func(*Options)) (*GetAdapterVersionOutput, error) 
 GetDocumentAnalysis(ctx context.Context, params *GetDocumentAnalysisInput, optFns ...func(*Options)) (*GetDocumentAnalysisOutput, error) 
 GetDocumentTextDetection(ctx context.Context, params *GetDocumentTextDetectionInput, optFns ...func(*Options)) (*GetDocumentTextDetectionOutput, error) 
 GetExpenseAnalysis(ctx context.Context, params *GetExpenseAnalysisInput, optFns ...func(*Options)) (*GetExpenseAnalysisOutput, error) 
 GetLendingAnalysis(ctx context.Context, params *GetLendingAnalysisInput, optFns ...func(*Options)) (*GetLendingAnalysisOutput, error) 
 GetLendingAnalysisSummary(ctx context.Context, params *GetLendingAnalysisSummaryInput, optFns ...func(*Options)) (*GetLendingAnalysisSummaryOutput, error) 
 ListAdapterVersions(ctx context.Context, params *ListAdapterVersionsInput, optFns ...func(*Options)) (*ListAdapterVersionsOutput, error) 
 ListAdapters(ctx context.Context, params *ListAdaptersInput, optFns ...func(*Options)) (*ListAdaptersOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 StartDocumentAnalysis(ctx context.Context, params *StartDocumentAnalysisInput, optFns ...func(*Options)) (*StartDocumentAnalysisOutput, error) 
 StartDocumentTextDetection(ctx context.Context, params *StartDocumentTextDetectionInput, optFns ...func(*Options)) (*StartDocumentTextDetectionOutput, error) 
 StartExpenseAnalysis(ctx context.Context, params *StartExpenseAnalysisInput, optFns ...func(*Options)) (*StartExpenseAnalysisOutput, error) 
 StartLendingAnalysis(ctx context.Context, params *StartLendingAnalysisInput, optFns ...func(*Options)) (*StartLendingAnalysisOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateAdapter(ctx context.Context, params *UpdateAdapterInput, optFns ...func(*Options)) (*UpdateAdapterOutput, error) 
}
