
package apigatewayv2_iface

// This file was generated by scripts/gen_ifaces.sh do not edit it manually
import (
    "context"
    . "github.com/aws/aws-sdk-go-v2/service/apigatewayv2"
)

// IClient defines the interface for apigatewayv2
type IClient interface {
 Options() Options 
 CreateApi(ctx context.Context, params *CreateApiInput, optFns ...func(*Options)) (*CreateApiOutput, error) 
 CreateApiMapping(ctx context.Context, params *CreateApiMappingInput, optFns ...func(*Options)) (*CreateApiMappingOutput, error) 
 CreateAuthorizer(ctx context.Context, params *CreateAuthorizerInput, optFns ...func(*Options)) (*CreateAuthorizerOutput, error) 
 CreateDeployment(ctx context.Context, params *CreateDeploymentInput, optFns ...func(*Options)) (*CreateDeploymentOutput, error) 
 CreateDomainName(ctx context.Context, params *CreateDomainNameInput, optFns ...func(*Options)) (*CreateDomainNameOutput, error) 
 CreateIntegration(ctx context.Context, params *CreateIntegrationInput, optFns ...func(*Options)) (*CreateIntegrationOutput, error) 
 CreateIntegrationResponse(ctx context.Context, params *CreateIntegrationResponseInput, optFns ...func(*Options)) (*CreateIntegrationResponseOutput, error) 
 CreateModel(ctx context.Context, params *CreateModelInput, optFns ...func(*Options)) (*CreateModelOutput, error) 
 CreateRoute(ctx context.Context, params *CreateRouteInput, optFns ...func(*Options)) (*CreateRouteOutput, error) 
 CreateRouteResponse(ctx context.Context, params *CreateRouteResponseInput, optFns ...func(*Options)) (*CreateRouteResponseOutput, error) 
 CreateStage(ctx context.Context, params *CreateStageInput, optFns ...func(*Options)) (*CreateStageOutput, error) 
 CreateVpcLink(ctx context.Context, params *CreateVpcLinkInput, optFns ...func(*Options)) (*CreateVpcLinkOutput, error) 
 DeleteAccessLogSettings(ctx context.Context, params *DeleteAccessLogSettingsInput, optFns ...func(*Options)) (*DeleteAccessLogSettingsOutput, error) 
 DeleteApi(ctx context.Context, params *DeleteApiInput, optFns ...func(*Options)) (*DeleteApiOutput, error) 
 DeleteApiMapping(ctx context.Context, params *DeleteApiMappingInput, optFns ...func(*Options)) (*DeleteApiMappingOutput, error) 
 DeleteAuthorizer(ctx context.Context, params *DeleteAuthorizerInput, optFns ...func(*Options)) (*DeleteAuthorizerOutput, error) 
 DeleteCorsConfiguration(ctx context.Context, params *DeleteCorsConfigurationInput, optFns ...func(*Options)) (*DeleteCorsConfigurationOutput, error) 
 DeleteDeployment(ctx context.Context, params *DeleteDeploymentInput, optFns ...func(*Options)) (*DeleteDeploymentOutput, error) 
 DeleteDomainName(ctx context.Context, params *DeleteDomainNameInput, optFns ...func(*Options)) (*DeleteDomainNameOutput, error) 
 DeleteIntegration(ctx context.Context, params *DeleteIntegrationInput, optFns ...func(*Options)) (*DeleteIntegrationOutput, error) 
 DeleteIntegrationResponse(ctx context.Context, params *DeleteIntegrationResponseInput, optFns ...func(*Options)) (*DeleteIntegrationResponseOutput, error) 
 DeleteModel(ctx context.Context, params *DeleteModelInput, optFns ...func(*Options)) (*DeleteModelOutput, error) 
 DeleteRoute(ctx context.Context, params *DeleteRouteInput, optFns ...func(*Options)) (*DeleteRouteOutput, error) 
 DeleteRouteRequestParameter(ctx context.Context, params *DeleteRouteRequestParameterInput, optFns ...func(*Options)) (*DeleteRouteRequestParameterOutput, error) 
 DeleteRouteResponse(ctx context.Context, params *DeleteRouteResponseInput, optFns ...func(*Options)) (*DeleteRouteResponseOutput, error) 
 DeleteRouteSettings(ctx context.Context, params *DeleteRouteSettingsInput, optFns ...func(*Options)) (*DeleteRouteSettingsOutput, error) 
 DeleteStage(ctx context.Context, params *DeleteStageInput, optFns ...func(*Options)) (*DeleteStageOutput, error) 
 DeleteVpcLink(ctx context.Context, params *DeleteVpcLinkInput, optFns ...func(*Options)) (*DeleteVpcLinkOutput, error) 
 ExportApi(ctx context.Context, params *ExportApiInput, optFns ...func(*Options)) (*ExportApiOutput, error) 
 GetApi(ctx context.Context, params *GetApiInput, optFns ...func(*Options)) (*GetApiOutput, error) 
 GetApiMapping(ctx context.Context, params *GetApiMappingInput, optFns ...func(*Options)) (*GetApiMappingOutput, error) 
 GetApiMappings(ctx context.Context, params *GetApiMappingsInput, optFns ...func(*Options)) (*GetApiMappingsOutput, error) 
 GetApis(ctx context.Context, params *GetApisInput, optFns ...func(*Options)) (*GetApisOutput, error) 
 GetAuthorizer(ctx context.Context, params *GetAuthorizerInput, optFns ...func(*Options)) (*GetAuthorizerOutput, error) 
 GetAuthorizers(ctx context.Context, params *GetAuthorizersInput, optFns ...func(*Options)) (*GetAuthorizersOutput, error) 
 GetDeployment(ctx context.Context, params *GetDeploymentInput, optFns ...func(*Options)) (*GetDeploymentOutput, error) 
 GetDeployments(ctx context.Context, params *GetDeploymentsInput, optFns ...func(*Options)) (*GetDeploymentsOutput, error) 
 GetDomainName(ctx context.Context, params *GetDomainNameInput, optFns ...func(*Options)) (*GetDomainNameOutput, error) 
 GetDomainNames(ctx context.Context, params *GetDomainNamesInput, optFns ...func(*Options)) (*GetDomainNamesOutput, error) 
 GetIntegration(ctx context.Context, params *GetIntegrationInput, optFns ...func(*Options)) (*GetIntegrationOutput, error) 
 GetIntegrationResponse(ctx context.Context, params *GetIntegrationResponseInput, optFns ...func(*Options)) (*GetIntegrationResponseOutput, error) 
 GetIntegrationResponses(ctx context.Context, params *GetIntegrationResponsesInput, optFns ...func(*Options)) (*GetIntegrationResponsesOutput, error) 
 GetIntegrations(ctx context.Context, params *GetIntegrationsInput, optFns ...func(*Options)) (*GetIntegrationsOutput, error) 
 GetModel(ctx context.Context, params *GetModelInput, optFns ...func(*Options)) (*GetModelOutput, error) 
 GetModelTemplate(ctx context.Context, params *GetModelTemplateInput, optFns ...func(*Options)) (*GetModelTemplateOutput, error) 
 GetModels(ctx context.Context, params *GetModelsInput, optFns ...func(*Options)) (*GetModelsOutput, error) 
 GetRoute(ctx context.Context, params *GetRouteInput, optFns ...func(*Options)) (*GetRouteOutput, error) 
 GetRouteResponse(ctx context.Context, params *GetRouteResponseInput, optFns ...func(*Options)) (*GetRouteResponseOutput, error) 
 GetRouteResponses(ctx context.Context, params *GetRouteResponsesInput, optFns ...func(*Options)) (*GetRouteResponsesOutput, error) 
 GetRoutes(ctx context.Context, params *GetRoutesInput, optFns ...func(*Options)) (*GetRoutesOutput, error) 
 GetStage(ctx context.Context, params *GetStageInput, optFns ...func(*Options)) (*GetStageOutput, error) 
 GetStages(ctx context.Context, params *GetStagesInput, optFns ...func(*Options)) (*GetStagesOutput, error) 
 GetTags(ctx context.Context, params *GetTagsInput, optFns ...func(*Options)) (*GetTagsOutput, error) 
 GetVpcLink(ctx context.Context, params *GetVpcLinkInput, optFns ...func(*Options)) (*GetVpcLinkOutput, error) 
 GetVpcLinks(ctx context.Context, params *GetVpcLinksInput, optFns ...func(*Options)) (*GetVpcLinksOutput, error) 
 ImportApi(ctx context.Context, params *ImportApiInput, optFns ...func(*Options)) (*ImportApiOutput, error) 
 ReimportApi(ctx context.Context, params *ReimportApiInput, optFns ...func(*Options)) (*ReimportApiOutput, error) 
 ResetAuthorizersCache(ctx context.Context, params *ResetAuthorizersCacheInput, optFns ...func(*Options)) (*ResetAuthorizersCacheOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateApi(ctx context.Context, params *UpdateApiInput, optFns ...func(*Options)) (*UpdateApiOutput, error) 
 UpdateApiMapping(ctx context.Context, params *UpdateApiMappingInput, optFns ...func(*Options)) (*UpdateApiMappingOutput, error) 
 UpdateAuthorizer(ctx context.Context, params *UpdateAuthorizerInput, optFns ...func(*Options)) (*UpdateAuthorizerOutput, error) 
 UpdateDeployment(ctx context.Context, params *UpdateDeploymentInput, optFns ...func(*Options)) (*UpdateDeploymentOutput, error) 
 UpdateDomainName(ctx context.Context, params *UpdateDomainNameInput, optFns ...func(*Options)) (*UpdateDomainNameOutput, error) 
 UpdateIntegration(ctx context.Context, params *UpdateIntegrationInput, optFns ...func(*Options)) (*UpdateIntegrationOutput, error) 
 UpdateIntegrationResponse(ctx context.Context, params *UpdateIntegrationResponseInput, optFns ...func(*Options)) (*UpdateIntegrationResponseOutput, error) 
 UpdateModel(ctx context.Context, params *UpdateModelInput, optFns ...func(*Options)) (*UpdateModelOutput, error) 
 UpdateRoute(ctx context.Context, params *UpdateRouteInput, optFns ...func(*Options)) (*UpdateRouteOutput, error) 
 UpdateRouteResponse(ctx context.Context, params *UpdateRouteResponseInput, optFns ...func(*Options)) (*UpdateRouteResponseOutput, error) 
 UpdateStage(ctx context.Context, params *UpdateStageInput, optFns ...func(*Options)) (*UpdateStageOutput, error) 
 UpdateVpcLink(ctx context.Context, params *UpdateVpcLinkInput, optFns ...func(*Options)) (*UpdateVpcLinkOutput, error) 
}
