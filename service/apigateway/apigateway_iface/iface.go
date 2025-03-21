
package apigateway_iface

// This file was generated by scripts/gen_ifaces.sh do not edit it manually
import (
    "context"
    . "github.com/aws/aws-sdk-go-v2/service/apigateway"
)

// IClient defines the interface for apigateway
type IClient interface {
 Options() Options 
 CreateApiKey(ctx context.Context, params *CreateApiKeyInput, optFns ...func(*Options)) (*CreateApiKeyOutput, error) 
 CreateAuthorizer(ctx context.Context, params *CreateAuthorizerInput, optFns ...func(*Options)) (*CreateAuthorizerOutput, error) 
 CreateBasePathMapping(ctx context.Context, params *CreateBasePathMappingInput, optFns ...func(*Options)) (*CreateBasePathMappingOutput, error) 
 CreateDeployment(ctx context.Context, params *CreateDeploymentInput, optFns ...func(*Options)) (*CreateDeploymentOutput, error) 
 CreateDocumentationPart(ctx context.Context, params *CreateDocumentationPartInput, optFns ...func(*Options)) (*CreateDocumentationPartOutput, error) 
 CreateDocumentationVersion(ctx context.Context, params *CreateDocumentationVersionInput, optFns ...func(*Options)) (*CreateDocumentationVersionOutput, error) 
 CreateDomainName(ctx context.Context, params *CreateDomainNameInput, optFns ...func(*Options)) (*CreateDomainNameOutput, error) 
 CreateDomainNameAccessAssociation(ctx context.Context, params *CreateDomainNameAccessAssociationInput, optFns ...func(*Options)) (*CreateDomainNameAccessAssociationOutput, error) 
 CreateModel(ctx context.Context, params *CreateModelInput, optFns ...func(*Options)) (*CreateModelOutput, error) 
 CreateRequestValidator(ctx context.Context, params *CreateRequestValidatorInput, optFns ...func(*Options)) (*CreateRequestValidatorOutput, error) 
 CreateResource(ctx context.Context, params *CreateResourceInput, optFns ...func(*Options)) (*CreateResourceOutput, error) 
 CreateRestApi(ctx context.Context, params *CreateRestApiInput, optFns ...func(*Options)) (*CreateRestApiOutput, error) 
 CreateStage(ctx context.Context, params *CreateStageInput, optFns ...func(*Options)) (*CreateStageOutput, error) 
 CreateUsagePlan(ctx context.Context, params *CreateUsagePlanInput, optFns ...func(*Options)) (*CreateUsagePlanOutput, error) 
 CreateUsagePlanKey(ctx context.Context, params *CreateUsagePlanKeyInput, optFns ...func(*Options)) (*CreateUsagePlanKeyOutput, error) 
 CreateVpcLink(ctx context.Context, params *CreateVpcLinkInput, optFns ...func(*Options)) (*CreateVpcLinkOutput, error) 
 DeleteApiKey(ctx context.Context, params *DeleteApiKeyInput, optFns ...func(*Options)) (*DeleteApiKeyOutput, error) 
 DeleteAuthorizer(ctx context.Context, params *DeleteAuthorizerInput, optFns ...func(*Options)) (*DeleteAuthorizerOutput, error) 
 DeleteBasePathMapping(ctx context.Context, params *DeleteBasePathMappingInput, optFns ...func(*Options)) (*DeleteBasePathMappingOutput, error) 
 DeleteClientCertificate(ctx context.Context, params *DeleteClientCertificateInput, optFns ...func(*Options)) (*DeleteClientCertificateOutput, error) 
 DeleteDeployment(ctx context.Context, params *DeleteDeploymentInput, optFns ...func(*Options)) (*DeleteDeploymentOutput, error) 
 DeleteDocumentationPart(ctx context.Context, params *DeleteDocumentationPartInput, optFns ...func(*Options)) (*DeleteDocumentationPartOutput, error) 
 DeleteDocumentationVersion(ctx context.Context, params *DeleteDocumentationVersionInput, optFns ...func(*Options)) (*DeleteDocumentationVersionOutput, error) 
 DeleteDomainName(ctx context.Context, params *DeleteDomainNameInput, optFns ...func(*Options)) (*DeleteDomainNameOutput, error) 
 DeleteDomainNameAccessAssociation(ctx context.Context, params *DeleteDomainNameAccessAssociationInput, optFns ...func(*Options)) (*DeleteDomainNameAccessAssociationOutput, error) 
 DeleteGatewayResponse(ctx context.Context, params *DeleteGatewayResponseInput, optFns ...func(*Options)) (*DeleteGatewayResponseOutput, error) 
 DeleteIntegration(ctx context.Context, params *DeleteIntegrationInput, optFns ...func(*Options)) (*DeleteIntegrationOutput, error) 
 DeleteIntegrationResponse(ctx context.Context, params *DeleteIntegrationResponseInput, optFns ...func(*Options)) (*DeleteIntegrationResponseOutput, error) 
 DeleteMethod(ctx context.Context, params *DeleteMethodInput, optFns ...func(*Options)) (*DeleteMethodOutput, error) 
 DeleteMethodResponse(ctx context.Context, params *DeleteMethodResponseInput, optFns ...func(*Options)) (*DeleteMethodResponseOutput, error) 
 DeleteModel(ctx context.Context, params *DeleteModelInput, optFns ...func(*Options)) (*DeleteModelOutput, error) 
 DeleteRequestValidator(ctx context.Context, params *DeleteRequestValidatorInput, optFns ...func(*Options)) (*DeleteRequestValidatorOutput, error) 
 DeleteResource(ctx context.Context, params *DeleteResourceInput, optFns ...func(*Options)) (*DeleteResourceOutput, error) 
 DeleteRestApi(ctx context.Context, params *DeleteRestApiInput, optFns ...func(*Options)) (*DeleteRestApiOutput, error) 
 DeleteStage(ctx context.Context, params *DeleteStageInput, optFns ...func(*Options)) (*DeleteStageOutput, error) 
 DeleteUsagePlan(ctx context.Context, params *DeleteUsagePlanInput, optFns ...func(*Options)) (*DeleteUsagePlanOutput, error) 
 DeleteUsagePlanKey(ctx context.Context, params *DeleteUsagePlanKeyInput, optFns ...func(*Options)) (*DeleteUsagePlanKeyOutput, error) 
 DeleteVpcLink(ctx context.Context, params *DeleteVpcLinkInput, optFns ...func(*Options)) (*DeleteVpcLinkOutput, error) 
 FlushStageAuthorizersCache(ctx context.Context, params *FlushStageAuthorizersCacheInput, optFns ...func(*Options)) (*FlushStageAuthorizersCacheOutput, error) 
 FlushStageCache(ctx context.Context, params *FlushStageCacheInput, optFns ...func(*Options)) (*FlushStageCacheOutput, error) 
 GenerateClientCertificate(ctx context.Context, params *GenerateClientCertificateInput, optFns ...func(*Options)) (*GenerateClientCertificateOutput, error) 
 GetAccount(ctx context.Context, params *GetAccountInput, optFns ...func(*Options)) (*GetAccountOutput, error) 
 GetApiKey(ctx context.Context, params *GetApiKeyInput, optFns ...func(*Options)) (*GetApiKeyOutput, error) 
 GetApiKeys(ctx context.Context, params *GetApiKeysInput, optFns ...func(*Options)) (*GetApiKeysOutput, error) 
 GetAuthorizer(ctx context.Context, params *GetAuthorizerInput, optFns ...func(*Options)) (*GetAuthorizerOutput, error) 
 GetAuthorizers(ctx context.Context, params *GetAuthorizersInput, optFns ...func(*Options)) (*GetAuthorizersOutput, error) 
 GetBasePathMapping(ctx context.Context, params *GetBasePathMappingInput, optFns ...func(*Options)) (*GetBasePathMappingOutput, error) 
 GetBasePathMappings(ctx context.Context, params *GetBasePathMappingsInput, optFns ...func(*Options)) (*GetBasePathMappingsOutput, error) 
 GetClientCertificate(ctx context.Context, params *GetClientCertificateInput, optFns ...func(*Options)) (*GetClientCertificateOutput, error) 
 GetClientCertificates(ctx context.Context, params *GetClientCertificatesInput, optFns ...func(*Options)) (*GetClientCertificatesOutput, error) 
 GetDeployment(ctx context.Context, params *GetDeploymentInput, optFns ...func(*Options)) (*GetDeploymentOutput, error) 
 GetDeployments(ctx context.Context, params *GetDeploymentsInput, optFns ...func(*Options)) (*GetDeploymentsOutput, error) 
 GetDocumentationPart(ctx context.Context, params *GetDocumentationPartInput, optFns ...func(*Options)) (*GetDocumentationPartOutput, error) 
 GetDocumentationParts(ctx context.Context, params *GetDocumentationPartsInput, optFns ...func(*Options)) (*GetDocumentationPartsOutput, error) 
 GetDocumentationVersion(ctx context.Context, params *GetDocumentationVersionInput, optFns ...func(*Options)) (*GetDocumentationVersionOutput, error) 
 GetDocumentationVersions(ctx context.Context, params *GetDocumentationVersionsInput, optFns ...func(*Options)) (*GetDocumentationVersionsOutput, error) 
 GetDomainName(ctx context.Context, params *GetDomainNameInput, optFns ...func(*Options)) (*GetDomainNameOutput, error) 
 GetDomainNameAccessAssociations(ctx context.Context, params *GetDomainNameAccessAssociationsInput, optFns ...func(*Options)) (*GetDomainNameAccessAssociationsOutput, error) 
 GetDomainNames(ctx context.Context, params *GetDomainNamesInput, optFns ...func(*Options)) (*GetDomainNamesOutput, error) 
 GetExport(ctx context.Context, params *GetExportInput, optFns ...func(*Options)) (*GetExportOutput, error) 
 GetGatewayResponse(ctx context.Context, params *GetGatewayResponseInput, optFns ...func(*Options)) (*GetGatewayResponseOutput, error) 
 GetGatewayResponses(ctx context.Context, params *GetGatewayResponsesInput, optFns ...func(*Options)) (*GetGatewayResponsesOutput, error) 
 GetIntegration(ctx context.Context, params *GetIntegrationInput, optFns ...func(*Options)) (*GetIntegrationOutput, error) 
 GetIntegrationResponse(ctx context.Context, params *GetIntegrationResponseInput, optFns ...func(*Options)) (*GetIntegrationResponseOutput, error) 
 GetMethod(ctx context.Context, params *GetMethodInput, optFns ...func(*Options)) (*GetMethodOutput, error) 
 GetMethodResponse(ctx context.Context, params *GetMethodResponseInput, optFns ...func(*Options)) (*GetMethodResponseOutput, error) 
 GetModel(ctx context.Context, params *GetModelInput, optFns ...func(*Options)) (*GetModelOutput, error) 
 GetModelTemplate(ctx context.Context, params *GetModelTemplateInput, optFns ...func(*Options)) (*GetModelTemplateOutput, error) 
 GetModels(ctx context.Context, params *GetModelsInput, optFns ...func(*Options)) (*GetModelsOutput, error) 
 GetRequestValidator(ctx context.Context, params *GetRequestValidatorInput, optFns ...func(*Options)) (*GetRequestValidatorOutput, error) 
 GetRequestValidators(ctx context.Context, params *GetRequestValidatorsInput, optFns ...func(*Options)) (*GetRequestValidatorsOutput, error) 
 GetResource(ctx context.Context, params *GetResourceInput, optFns ...func(*Options)) (*GetResourceOutput, error) 
 GetResources(ctx context.Context, params *GetResourcesInput, optFns ...func(*Options)) (*GetResourcesOutput, error) 
 GetRestApi(ctx context.Context, params *GetRestApiInput, optFns ...func(*Options)) (*GetRestApiOutput, error) 
 GetRestApis(ctx context.Context, params *GetRestApisInput, optFns ...func(*Options)) (*GetRestApisOutput, error) 
 GetSdk(ctx context.Context, params *GetSdkInput, optFns ...func(*Options)) (*GetSdkOutput, error) 
 GetSdkType(ctx context.Context, params *GetSdkTypeInput, optFns ...func(*Options)) (*GetSdkTypeOutput, error) 
 GetSdkTypes(ctx context.Context, params *GetSdkTypesInput, optFns ...func(*Options)) (*GetSdkTypesOutput, error) 
 GetStage(ctx context.Context, params *GetStageInput, optFns ...func(*Options)) (*GetStageOutput, error) 
 GetStages(ctx context.Context, params *GetStagesInput, optFns ...func(*Options)) (*GetStagesOutput, error) 
 GetTags(ctx context.Context, params *GetTagsInput, optFns ...func(*Options)) (*GetTagsOutput, error) 
 GetUsage(ctx context.Context, params *GetUsageInput, optFns ...func(*Options)) (*GetUsageOutput, error) 
 GetUsagePlan(ctx context.Context, params *GetUsagePlanInput, optFns ...func(*Options)) (*GetUsagePlanOutput, error) 
 GetUsagePlanKey(ctx context.Context, params *GetUsagePlanKeyInput, optFns ...func(*Options)) (*GetUsagePlanKeyOutput, error) 
 GetUsagePlanKeys(ctx context.Context, params *GetUsagePlanKeysInput, optFns ...func(*Options)) (*GetUsagePlanKeysOutput, error) 
 GetUsagePlans(ctx context.Context, params *GetUsagePlansInput, optFns ...func(*Options)) (*GetUsagePlansOutput, error) 
 GetVpcLink(ctx context.Context, params *GetVpcLinkInput, optFns ...func(*Options)) (*GetVpcLinkOutput, error) 
 GetVpcLinks(ctx context.Context, params *GetVpcLinksInput, optFns ...func(*Options)) (*GetVpcLinksOutput, error) 
 ImportApiKeys(ctx context.Context, params *ImportApiKeysInput, optFns ...func(*Options)) (*ImportApiKeysOutput, error) 
 ImportDocumentationParts(ctx context.Context, params *ImportDocumentationPartsInput, optFns ...func(*Options)) (*ImportDocumentationPartsOutput, error) 
 ImportRestApi(ctx context.Context, params *ImportRestApiInput, optFns ...func(*Options)) (*ImportRestApiOutput, error) 
 PutGatewayResponse(ctx context.Context, params *PutGatewayResponseInput, optFns ...func(*Options)) (*PutGatewayResponseOutput, error) 
 PutIntegration(ctx context.Context, params *PutIntegrationInput, optFns ...func(*Options)) (*PutIntegrationOutput, error) 
 PutIntegrationResponse(ctx context.Context, params *PutIntegrationResponseInput, optFns ...func(*Options)) (*PutIntegrationResponseOutput, error) 
 PutMethod(ctx context.Context, params *PutMethodInput, optFns ...func(*Options)) (*PutMethodOutput, error) 
 PutMethodResponse(ctx context.Context, params *PutMethodResponseInput, optFns ...func(*Options)) (*PutMethodResponseOutput, error) 
 PutRestApi(ctx context.Context, params *PutRestApiInput, optFns ...func(*Options)) (*PutRestApiOutput, error) 
 RejectDomainNameAccessAssociation(ctx context.Context, params *RejectDomainNameAccessAssociationInput, optFns ...func(*Options)) (*RejectDomainNameAccessAssociationOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 TestInvokeAuthorizer(ctx context.Context, params *TestInvokeAuthorizerInput, optFns ...func(*Options)) (*TestInvokeAuthorizerOutput, error) 
 TestInvokeMethod(ctx context.Context, params *TestInvokeMethodInput, optFns ...func(*Options)) (*TestInvokeMethodOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateAccount(ctx context.Context, params *UpdateAccountInput, optFns ...func(*Options)) (*UpdateAccountOutput, error) 
 UpdateApiKey(ctx context.Context, params *UpdateApiKeyInput, optFns ...func(*Options)) (*UpdateApiKeyOutput, error) 
 UpdateAuthorizer(ctx context.Context, params *UpdateAuthorizerInput, optFns ...func(*Options)) (*UpdateAuthorizerOutput, error) 
 UpdateBasePathMapping(ctx context.Context, params *UpdateBasePathMappingInput, optFns ...func(*Options)) (*UpdateBasePathMappingOutput, error) 
 UpdateClientCertificate(ctx context.Context, params *UpdateClientCertificateInput, optFns ...func(*Options)) (*UpdateClientCertificateOutput, error) 
 UpdateDeployment(ctx context.Context, params *UpdateDeploymentInput, optFns ...func(*Options)) (*UpdateDeploymentOutput, error) 
 UpdateDocumentationPart(ctx context.Context, params *UpdateDocumentationPartInput, optFns ...func(*Options)) (*UpdateDocumentationPartOutput, error) 
 UpdateDocumentationVersion(ctx context.Context, params *UpdateDocumentationVersionInput, optFns ...func(*Options)) (*UpdateDocumentationVersionOutput, error) 
 UpdateDomainName(ctx context.Context, params *UpdateDomainNameInput, optFns ...func(*Options)) (*UpdateDomainNameOutput, error) 
 UpdateGatewayResponse(ctx context.Context, params *UpdateGatewayResponseInput, optFns ...func(*Options)) (*UpdateGatewayResponseOutput, error) 
 UpdateIntegration(ctx context.Context, params *UpdateIntegrationInput, optFns ...func(*Options)) (*UpdateIntegrationOutput, error) 
 UpdateIntegrationResponse(ctx context.Context, params *UpdateIntegrationResponseInput, optFns ...func(*Options)) (*UpdateIntegrationResponseOutput, error) 
 UpdateMethod(ctx context.Context, params *UpdateMethodInput, optFns ...func(*Options)) (*UpdateMethodOutput, error) 
 UpdateMethodResponse(ctx context.Context, params *UpdateMethodResponseInput, optFns ...func(*Options)) (*UpdateMethodResponseOutput, error) 
 UpdateModel(ctx context.Context, params *UpdateModelInput, optFns ...func(*Options)) (*UpdateModelOutput, error) 
 UpdateRequestValidator(ctx context.Context, params *UpdateRequestValidatorInput, optFns ...func(*Options)) (*UpdateRequestValidatorOutput, error) 
 UpdateResource(ctx context.Context, params *UpdateResourceInput, optFns ...func(*Options)) (*UpdateResourceOutput, error) 
 UpdateRestApi(ctx context.Context, params *UpdateRestApiInput, optFns ...func(*Options)) (*UpdateRestApiOutput, error) 
 UpdateStage(ctx context.Context, params *UpdateStageInput, optFns ...func(*Options)) (*UpdateStageOutput, error) 
 UpdateUsage(ctx context.Context, params *UpdateUsageInput, optFns ...func(*Options)) (*UpdateUsageOutput, error) 
 UpdateUsagePlan(ctx context.Context, params *UpdateUsagePlanInput, optFns ...func(*Options)) (*UpdateUsagePlanOutput, error) 
 UpdateVpcLink(ctx context.Context, params *UpdateVpcLinkInput, optFns ...func(*Options)) (*UpdateVpcLinkOutput, error) 
}
