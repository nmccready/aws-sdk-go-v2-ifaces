
package appintegrations

import (
    "github.com/aws/aws-sdk-go-v2/service/appintegrations"
)

// IAppintegrations defines the interface for appintegrations
type IAppintegrations interface {
 Options() Options 
 CreateApplication(ctx context.Context, params *CreateApplicationInput, optFns ...func(*Options)) (*CreateApplicationOutput, error) 
 CreateDataIntegration(ctx context.Context, params *CreateDataIntegrationInput, optFns ...func(*Options)) (*CreateDataIntegrationOutput, error) 
 CreateEventIntegration(ctx context.Context, params *CreateEventIntegrationInput, optFns ...func(*Options)) (*CreateEventIntegrationOutput, error) 
 DeleteApplication(ctx context.Context, params *DeleteApplicationInput, optFns ...func(*Options)) (*DeleteApplicationOutput, error) 
 DeleteDataIntegration(ctx context.Context, params *DeleteDataIntegrationInput, optFns ...func(*Options)) (*DeleteDataIntegrationOutput, error) 
 DeleteEventIntegration(ctx context.Context, params *DeleteEventIntegrationInput, optFns ...func(*Options)) (*DeleteEventIntegrationOutput, error) 
 GetApplication(ctx context.Context, params *GetApplicationInput, optFns ...func(*Options)) (*GetApplicationOutput, error) 
 GetDataIntegration(ctx context.Context, params *GetDataIntegrationInput, optFns ...func(*Options)) (*GetDataIntegrationOutput, error) 
 GetEventIntegration(ctx context.Context, params *GetEventIntegrationInput, optFns ...func(*Options)) (*GetEventIntegrationOutput, error) 
 ListApplicationAssociations(ctx context.Context, params *ListApplicationAssociationsInput, optFns ...func(*Options)) (*ListApplicationAssociationsOutput, error) 
 ListApplications(ctx context.Context, params *ListApplicationsInput, optFns ...func(*Options)) (*ListApplicationsOutput, error) 
 ListDataIntegrationAssociations(ctx context.Context, params *ListDataIntegrationAssociationsInput, optFns ...func(*Options)) (*ListDataIntegrationAssociationsOutput, error) 
 ListDataIntegrations(ctx context.Context, params *ListDataIntegrationsInput, optFns ...func(*Options)) (*ListDataIntegrationsOutput, error) 
 ListEventIntegrationAssociations(ctx context.Context, params *ListEventIntegrationAssociationsInput, optFns ...func(*Options)) (*ListEventIntegrationAssociationsOutput, error) 
 ListEventIntegrations(ctx context.Context, params *ListEventIntegrationsInput, optFns ...func(*Options)) (*ListEventIntegrationsOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateApplication(ctx context.Context, params *UpdateApplicationInput, optFns ...func(*Options)) (*UpdateApplicationOutput, error) 
 UpdateDataIntegration(ctx context.Context, params *UpdateDataIntegrationInput, optFns ...func(*Options)) (*UpdateDataIntegrationOutput, error) 
 UpdateEventIntegration(ctx context.Context, params *UpdateEventIntegrationInput, optFns ...func(*Options)) (*UpdateEventIntegrationOutput, error) 
}
