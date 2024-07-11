
package amplifyuibuilder

import (
    "github.com/aws/aws-sdk-go-v2/service/amplifyuibuilder"
)

// IAmplifyuibuilder defines the interface for amplifyuibuilder
type IAmplifyuibuilder interface {
 Options() Options 
 CreateComponent(ctx context.Context, params *CreateComponentInput, optFns ...func(*Options)) (*CreateComponentOutput, error) 
 CreateForm(ctx context.Context, params *CreateFormInput, optFns ...func(*Options)) (*CreateFormOutput, error) 
 CreateTheme(ctx context.Context, params *CreateThemeInput, optFns ...func(*Options)) (*CreateThemeOutput, error) 
 DeleteComponent(ctx context.Context, params *DeleteComponentInput, optFns ...func(*Options)) (*DeleteComponentOutput, error) 
 DeleteForm(ctx context.Context, params *DeleteFormInput, optFns ...func(*Options)) (*DeleteFormOutput, error) 
 DeleteTheme(ctx context.Context, params *DeleteThemeInput, optFns ...func(*Options)) (*DeleteThemeOutput, error) 
 ExchangeCodeForToken(ctx context.Context, params *ExchangeCodeForTokenInput, optFns ...func(*Options)) (*ExchangeCodeForTokenOutput, error) 
 ExportComponents(ctx context.Context, params *ExportComponentsInput, optFns ...func(*Options)) (*ExportComponentsOutput, error) 
 ExportForms(ctx context.Context, params *ExportFormsInput, optFns ...func(*Options)) (*ExportFormsOutput, error) 
 ExportThemes(ctx context.Context, params *ExportThemesInput, optFns ...func(*Options)) (*ExportThemesOutput, error) 
 GetCodegenJob(ctx context.Context, params *GetCodegenJobInput, optFns ...func(*Options)) (*GetCodegenJobOutput, error) 
 GetComponent(ctx context.Context, params *GetComponentInput, optFns ...func(*Options)) (*GetComponentOutput, error) 
 GetForm(ctx context.Context, params *GetFormInput, optFns ...func(*Options)) (*GetFormOutput, error) 
 GetMetadata(ctx context.Context, params *GetMetadataInput, optFns ...func(*Options)) (*GetMetadataOutput, error) 
 GetTheme(ctx context.Context, params *GetThemeInput, optFns ...func(*Options)) (*GetThemeOutput, error) 
 ListCodegenJobs(ctx context.Context, params *ListCodegenJobsInput, optFns ...func(*Options)) (*ListCodegenJobsOutput, error) 
 ListComponents(ctx context.Context, params *ListComponentsInput, optFns ...func(*Options)) (*ListComponentsOutput, error) 
 ListForms(ctx context.Context, params *ListFormsInput, optFns ...func(*Options)) (*ListFormsOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 ListThemes(ctx context.Context, params *ListThemesInput, optFns ...func(*Options)) (*ListThemesOutput, error) 
 PutMetadataFlag(ctx context.Context, params *PutMetadataFlagInput, optFns ...func(*Options)) (*PutMetadataFlagOutput, error) 
 RefreshToken(ctx context.Context, params *RefreshTokenInput, optFns ...func(*Options)) (*RefreshTokenOutput, error) 
 StartCodegenJob(ctx context.Context, params *StartCodegenJobInput, optFns ...func(*Options)) (*StartCodegenJobOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateComponent(ctx context.Context, params *UpdateComponentInput, optFns ...func(*Options)) (*UpdateComponentOutput, error) 
 UpdateForm(ctx context.Context, params *UpdateFormInput, optFns ...func(*Options)) (*UpdateFormOutput, error) 
 UpdateTheme(ctx context.Context, params *UpdateThemeInput, optFns ...func(*Options)) (*UpdateThemeOutput, error) 
}
