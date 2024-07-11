
package pcaconnectorad

import (
    "github.com/aws/aws-sdk-go-v2/service/pcaconnectorad"
)

// IClient defines the interface for pcaconnectorad
type IClient interface {
 Options() Options 
 CreateConnector(ctx context.Context, params *CreateConnectorInput, optFns ...func(*Options)) (*CreateConnectorOutput, error) 
 CreateDirectoryRegistration(ctx context.Context, params *CreateDirectoryRegistrationInput, optFns ...func(*Options)) (*CreateDirectoryRegistrationOutput, error) 
 CreateServicePrincipalName(ctx context.Context, params *CreateServicePrincipalNameInput, optFns ...func(*Options)) (*CreateServicePrincipalNameOutput, error) 
 CreateTemplate(ctx context.Context, params *CreateTemplateInput, optFns ...func(*Options)) (*CreateTemplateOutput, error) 
 CreateTemplateGroupAccessControlEntry(ctx context.Context, params *CreateTemplateGroupAccessControlEntryInput, optFns ...func(*Options)) (*CreateTemplateGroupAccessControlEntryOutput, error) 
 DeleteConnector(ctx context.Context, params *DeleteConnectorInput, optFns ...func(*Options)) (*DeleteConnectorOutput, error) 
 DeleteDirectoryRegistration(ctx context.Context, params *DeleteDirectoryRegistrationInput, optFns ...func(*Options)) (*DeleteDirectoryRegistrationOutput, error) 
 DeleteServicePrincipalName(ctx context.Context, params *DeleteServicePrincipalNameInput, optFns ...func(*Options)) (*DeleteServicePrincipalNameOutput, error) 
 DeleteTemplate(ctx context.Context, params *DeleteTemplateInput, optFns ...func(*Options)) (*DeleteTemplateOutput, error) 
 DeleteTemplateGroupAccessControlEntry(ctx context.Context, params *DeleteTemplateGroupAccessControlEntryInput, optFns ...func(*Options)) (*DeleteTemplateGroupAccessControlEntryOutput, error) 
 GetConnector(ctx context.Context, params *GetConnectorInput, optFns ...func(*Options)) (*GetConnectorOutput, error) 
 GetDirectoryRegistration(ctx context.Context, params *GetDirectoryRegistrationInput, optFns ...func(*Options)) (*GetDirectoryRegistrationOutput, error) 
 GetServicePrincipalName(ctx context.Context, params *GetServicePrincipalNameInput, optFns ...func(*Options)) (*GetServicePrincipalNameOutput, error) 
 GetTemplate(ctx context.Context, params *GetTemplateInput, optFns ...func(*Options)) (*GetTemplateOutput, error) 
 GetTemplateGroupAccessControlEntry(ctx context.Context, params *GetTemplateGroupAccessControlEntryInput, optFns ...func(*Options)) (*GetTemplateGroupAccessControlEntryOutput, error) 
 ListConnectors(ctx context.Context, params *ListConnectorsInput, optFns ...func(*Options)) (*ListConnectorsOutput, error) 
 ListDirectoryRegistrations(ctx context.Context, params *ListDirectoryRegistrationsInput, optFns ...func(*Options)) (*ListDirectoryRegistrationsOutput, error) 
 ListServicePrincipalNames(ctx context.Context, params *ListServicePrincipalNamesInput, optFns ...func(*Options)) (*ListServicePrincipalNamesOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 ListTemplateGroupAccessControlEntries(ctx context.Context, params *ListTemplateGroupAccessControlEntriesInput, optFns ...func(*Options)) (*ListTemplateGroupAccessControlEntriesOutput, error) 
 ListTemplates(ctx context.Context, params *ListTemplatesInput, optFns ...func(*Options)) (*ListTemplatesOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateTemplate(ctx context.Context, params *UpdateTemplateInput, optFns ...func(*Options)) (*UpdateTemplateOutput, error) 
 UpdateTemplateGroupAccessControlEntry(ctx context.Context, params *UpdateTemplateGroupAccessControlEntryInput, optFns ...func(*Options)) (*UpdateTemplateGroupAccessControlEntryOutput, error) 
}
