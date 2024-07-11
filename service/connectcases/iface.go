
package connectcases

import (
    "github.com/aws/aws-sdk-go-v2/service/connectcases"
)

// IConnectcases defines the interface for connectcases
type IConnectcases interface {
 Options() Options 
 BatchGetField(ctx context.Context, params *BatchGetFieldInput, optFns ...func(*Options)) (*BatchGetFieldOutput, error) 
 BatchPutFieldOptions(ctx context.Context, params *BatchPutFieldOptionsInput, optFns ...func(*Options)) (*BatchPutFieldOptionsOutput, error) 
 CreateCase(ctx context.Context, params *CreateCaseInput, optFns ...func(*Options)) (*CreateCaseOutput, error) 
 CreateDomain(ctx context.Context, params *CreateDomainInput, optFns ...func(*Options)) (*CreateDomainOutput, error) 
 CreateField(ctx context.Context, params *CreateFieldInput, optFns ...func(*Options)) (*CreateFieldOutput, error) 
 CreateLayout(ctx context.Context, params *CreateLayoutInput, optFns ...func(*Options)) (*CreateLayoutOutput, error) 
 CreateRelatedItem(ctx context.Context, params *CreateRelatedItemInput, optFns ...func(*Options)) (*CreateRelatedItemOutput, error) 
 CreateTemplate(ctx context.Context, params *CreateTemplateInput, optFns ...func(*Options)) (*CreateTemplateOutput, error) 
 DeleteDomain(ctx context.Context, params *DeleteDomainInput, optFns ...func(*Options)) (*DeleteDomainOutput, error) 
 DeleteField(ctx context.Context, params *DeleteFieldInput, optFns ...func(*Options)) (*DeleteFieldOutput, error) 
 DeleteLayout(ctx context.Context, params *DeleteLayoutInput, optFns ...func(*Options)) (*DeleteLayoutOutput, error) 
 DeleteTemplate(ctx context.Context, params *DeleteTemplateInput, optFns ...func(*Options)) (*DeleteTemplateOutput, error) 
 GetCase(ctx context.Context, params *GetCaseInput, optFns ...func(*Options)) (*GetCaseOutput, error) 
 GetCaseAuditEvents(ctx context.Context, params *GetCaseAuditEventsInput, optFns ...func(*Options)) (*GetCaseAuditEventsOutput, error) 
 GetCaseEventConfiguration(ctx context.Context, params *GetCaseEventConfigurationInput, optFns ...func(*Options)) (*GetCaseEventConfigurationOutput, error) 
 GetDomain(ctx context.Context, params *GetDomainInput, optFns ...func(*Options)) (*GetDomainOutput, error) 
 GetLayout(ctx context.Context, params *GetLayoutInput, optFns ...func(*Options)) (*GetLayoutOutput, error) 
 GetTemplate(ctx context.Context, params *GetTemplateInput, optFns ...func(*Options)) (*GetTemplateOutput, error) 
 ListCasesForContact(ctx context.Context, params *ListCasesForContactInput, optFns ...func(*Options)) (*ListCasesForContactOutput, error) 
 ListDomains(ctx context.Context, params *ListDomainsInput, optFns ...func(*Options)) (*ListDomainsOutput, error) 
 ListFieldOptions(ctx context.Context, params *ListFieldOptionsInput, optFns ...func(*Options)) (*ListFieldOptionsOutput, error) 
 ListFields(ctx context.Context, params *ListFieldsInput, optFns ...func(*Options)) (*ListFieldsOutput, error) 
 ListLayouts(ctx context.Context, params *ListLayoutsInput, optFns ...func(*Options)) (*ListLayoutsOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 ListTemplates(ctx context.Context, params *ListTemplatesInput, optFns ...func(*Options)) (*ListTemplatesOutput, error) 
 PutCaseEventConfiguration(ctx context.Context, params *PutCaseEventConfigurationInput, optFns ...func(*Options)) (*PutCaseEventConfigurationOutput, error) 
 SearchCases(ctx context.Context, params *SearchCasesInput, optFns ...func(*Options)) (*SearchCasesOutput, error) 
 SearchRelatedItems(ctx context.Context, params *SearchRelatedItemsInput, optFns ...func(*Options)) (*SearchRelatedItemsOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateCase(ctx context.Context, params *UpdateCaseInput, optFns ...func(*Options)) (*UpdateCaseOutput, error) 
 UpdateField(ctx context.Context, params *UpdateFieldInput, optFns ...func(*Options)) (*UpdateFieldOutput, error) 
 UpdateLayout(ctx context.Context, params *UpdateLayoutInput, optFns ...func(*Options)) (*UpdateLayoutOutput, error) 
 UpdateTemplate(ctx context.Context, params *UpdateTemplateInput, optFns ...func(*Options)) (*UpdateTemplateOutput, error) 
}
