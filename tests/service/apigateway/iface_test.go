package apigateway_test

// tests for the apigateway service interface for this ../../../service/apigateway/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/apigateway"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/apigateway/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/apigateway/apigateway_iface"
	"github.com/stretchr/testify/assert"
)

func TestApigatewayServiceCanBeMocked(t *testing.T) {
	var iface apigateway_iface.IClient
	iface = &apigateway.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := apigateway.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateApiKey", func(t *testing.T) {
        input := &apigateway.CreateApiKeyInput{}
        output := &apigateway.CreateApiKeyOutput{}

        mockClient.On("CreateApiKey", ctx, input).Return(output, nil)

        result, err := mockClient.CreateApiKey(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateAuthorizer", func(t *testing.T) {
        input := &apigateway.CreateAuthorizerInput{}
        output := &apigateway.CreateAuthorizerOutput{}

        mockClient.On("CreateAuthorizer", ctx, input).Return(output, nil)

        result, err := mockClient.CreateAuthorizer(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateBasePathMapping", func(t *testing.T) {
        input := &apigateway.CreateBasePathMappingInput{}
        output := &apigateway.CreateBasePathMappingOutput{}

        mockClient.On("CreateBasePathMapping", ctx, input).Return(output, nil)

        result, err := mockClient.CreateBasePathMapping(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateDeployment", func(t *testing.T) {
        input := &apigateway.CreateDeploymentInput{}
        output := &apigateway.CreateDeploymentOutput{}

        mockClient.On("CreateDeployment", ctx, input).Return(output, nil)

        result, err := mockClient.CreateDeployment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateDocumentationPart", func(t *testing.T) {
        input := &apigateway.CreateDocumentationPartInput{}
        output := &apigateway.CreateDocumentationPartOutput{}

        mockClient.On("CreateDocumentationPart", ctx, input).Return(output, nil)

        result, err := mockClient.CreateDocumentationPart(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateDocumentationVersion", func(t *testing.T) {
        input := &apigateway.CreateDocumentationVersionInput{}
        output := &apigateway.CreateDocumentationVersionOutput{}

        mockClient.On("CreateDocumentationVersion", ctx, input).Return(output, nil)

        result, err := mockClient.CreateDocumentationVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateDomainName", func(t *testing.T) {
        input := &apigateway.CreateDomainNameInput{}
        output := &apigateway.CreateDomainNameOutput{}

        mockClient.On("CreateDomainName", ctx, input).Return(output, nil)

        result, err := mockClient.CreateDomainName(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateModel", func(t *testing.T) {
        input := &apigateway.CreateModelInput{}
        output := &apigateway.CreateModelOutput{}

        mockClient.On("CreateModel", ctx, input).Return(output, nil)

        result, err := mockClient.CreateModel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateRequestValidator", func(t *testing.T) {
        input := &apigateway.CreateRequestValidatorInput{}
        output := &apigateway.CreateRequestValidatorOutput{}

        mockClient.On("CreateRequestValidator", ctx, input).Return(output, nil)

        result, err := mockClient.CreateRequestValidator(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateResource", func(t *testing.T) {
        input := &apigateway.CreateResourceInput{}
        output := &apigateway.CreateResourceOutput{}

        mockClient.On("CreateResource", ctx, input).Return(output, nil)

        result, err := mockClient.CreateResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateRestApi", func(t *testing.T) {
        input := &apigateway.CreateRestApiInput{}
        output := &apigateway.CreateRestApiOutput{}

        mockClient.On("CreateRestApi", ctx, input).Return(output, nil)

        result, err := mockClient.CreateRestApi(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateStage", func(t *testing.T) {
        input := &apigateway.CreateStageInput{}
        output := &apigateway.CreateStageOutput{}

        mockClient.On("CreateStage", ctx, input).Return(output, nil)

        result, err := mockClient.CreateStage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateUsagePlan", func(t *testing.T) {
        input := &apigateway.CreateUsagePlanInput{}
        output := &apigateway.CreateUsagePlanOutput{}

        mockClient.On("CreateUsagePlan", ctx, input).Return(output, nil)

        result, err := mockClient.CreateUsagePlan(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateUsagePlanKey", func(t *testing.T) {
        input := &apigateway.CreateUsagePlanKeyInput{}
        output := &apigateway.CreateUsagePlanKeyOutput{}

        mockClient.On("CreateUsagePlanKey", ctx, input).Return(output, nil)

        result, err := mockClient.CreateUsagePlanKey(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateVpcLink", func(t *testing.T) {
        input := &apigateway.CreateVpcLinkInput{}
        output := &apigateway.CreateVpcLinkOutput{}

        mockClient.On("CreateVpcLink", ctx, input).Return(output, nil)

        result, err := mockClient.CreateVpcLink(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteApiKey", func(t *testing.T) {
        input := &apigateway.DeleteApiKeyInput{}
        output := &apigateway.DeleteApiKeyOutput{}

        mockClient.On("DeleteApiKey", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteApiKey(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteAuthorizer", func(t *testing.T) {
        input := &apigateway.DeleteAuthorizerInput{}
        output := &apigateway.DeleteAuthorizerOutput{}

        mockClient.On("DeleteAuthorizer", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteAuthorizer(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteBasePathMapping", func(t *testing.T) {
        input := &apigateway.DeleteBasePathMappingInput{}
        output := &apigateway.DeleteBasePathMappingOutput{}

        mockClient.On("DeleteBasePathMapping", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteBasePathMapping(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteClientCertificate", func(t *testing.T) {
        input := &apigateway.DeleteClientCertificateInput{}
        output := &apigateway.DeleteClientCertificateOutput{}

        mockClient.On("DeleteClientCertificate", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteClientCertificate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteDeployment", func(t *testing.T) {
        input := &apigateway.DeleteDeploymentInput{}
        output := &apigateway.DeleteDeploymentOutput{}

        mockClient.On("DeleteDeployment", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteDeployment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteDocumentationPart", func(t *testing.T) {
        input := &apigateway.DeleteDocumentationPartInput{}
        output := &apigateway.DeleteDocumentationPartOutput{}

        mockClient.On("DeleteDocumentationPart", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteDocumentationPart(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteDocumentationVersion", func(t *testing.T) {
        input := &apigateway.DeleteDocumentationVersionInput{}
        output := &apigateway.DeleteDocumentationVersionOutput{}

        mockClient.On("DeleteDocumentationVersion", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteDocumentationVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteDomainName", func(t *testing.T) {
        input := &apigateway.DeleteDomainNameInput{}
        output := &apigateway.DeleteDomainNameOutput{}

        mockClient.On("DeleteDomainName", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteDomainName(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteGatewayResponse", func(t *testing.T) {
        input := &apigateway.DeleteGatewayResponseInput{}
        output := &apigateway.DeleteGatewayResponseOutput{}

        mockClient.On("DeleteGatewayResponse", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteGatewayResponse(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteIntegration", func(t *testing.T) {
        input := &apigateway.DeleteIntegrationInput{}
        output := &apigateway.DeleteIntegrationOutput{}

        mockClient.On("DeleteIntegration", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteIntegration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteIntegrationResponse", func(t *testing.T) {
        input := &apigateway.DeleteIntegrationResponseInput{}
        output := &apigateway.DeleteIntegrationResponseOutput{}

        mockClient.On("DeleteIntegrationResponse", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteIntegrationResponse(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteMethod", func(t *testing.T) {
        input := &apigateway.DeleteMethodInput{}
        output := &apigateway.DeleteMethodOutput{}

        mockClient.On("DeleteMethod", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteMethod(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteMethodResponse", func(t *testing.T) {
        input := &apigateway.DeleteMethodResponseInput{}
        output := &apigateway.DeleteMethodResponseOutput{}

        mockClient.On("DeleteMethodResponse", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteMethodResponse(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteModel", func(t *testing.T) {
        input := &apigateway.DeleteModelInput{}
        output := &apigateway.DeleteModelOutput{}

        mockClient.On("DeleteModel", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteModel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteRequestValidator", func(t *testing.T) {
        input := &apigateway.DeleteRequestValidatorInput{}
        output := &apigateway.DeleteRequestValidatorOutput{}

        mockClient.On("DeleteRequestValidator", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteRequestValidator(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteResource", func(t *testing.T) {
        input := &apigateway.DeleteResourceInput{}
        output := &apigateway.DeleteResourceOutput{}

        mockClient.On("DeleteResource", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteRestApi", func(t *testing.T) {
        input := &apigateway.DeleteRestApiInput{}
        output := &apigateway.DeleteRestApiOutput{}

        mockClient.On("DeleteRestApi", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteRestApi(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteStage", func(t *testing.T) {
        input := &apigateway.DeleteStageInput{}
        output := &apigateway.DeleteStageOutput{}

        mockClient.On("DeleteStage", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteStage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteUsagePlan", func(t *testing.T) {
        input := &apigateway.DeleteUsagePlanInput{}
        output := &apigateway.DeleteUsagePlanOutput{}

        mockClient.On("DeleteUsagePlan", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteUsagePlan(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteUsagePlanKey", func(t *testing.T) {
        input := &apigateway.DeleteUsagePlanKeyInput{}
        output := &apigateway.DeleteUsagePlanKeyOutput{}

        mockClient.On("DeleteUsagePlanKey", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteUsagePlanKey(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteVpcLink", func(t *testing.T) {
        input := &apigateway.DeleteVpcLinkInput{}
        output := &apigateway.DeleteVpcLinkOutput{}

        mockClient.On("DeleteVpcLink", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteVpcLink(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestFlushStageAuthorizersCache", func(t *testing.T) {
        input := &apigateway.FlushStageAuthorizersCacheInput{}
        output := &apigateway.FlushStageAuthorizersCacheOutput{}

        mockClient.On("FlushStageAuthorizersCache", ctx, input).Return(output, nil)

        result, err := mockClient.FlushStageAuthorizersCache(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestFlushStageCache", func(t *testing.T) {
        input := &apigateway.FlushStageCacheInput{}
        output := &apigateway.FlushStageCacheOutput{}

        mockClient.On("FlushStageCache", ctx, input).Return(output, nil)

        result, err := mockClient.FlushStageCache(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGenerateClientCertificate", func(t *testing.T) {
        input := &apigateway.GenerateClientCertificateInput{}
        output := &apigateway.GenerateClientCertificateOutput{}

        mockClient.On("GenerateClientCertificate", ctx, input).Return(output, nil)

        result, err := mockClient.GenerateClientCertificate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetAccount", func(t *testing.T) {
        input := &apigateway.GetAccountInput{}
        output := &apigateway.GetAccountOutput{}

        mockClient.On("GetAccount", ctx, input).Return(output, nil)

        result, err := mockClient.GetAccount(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetApiKey", func(t *testing.T) {
        input := &apigateway.GetApiKeyInput{}
        output := &apigateway.GetApiKeyOutput{}

        mockClient.On("GetApiKey", ctx, input).Return(output, nil)

        result, err := mockClient.GetApiKey(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetApiKeys", func(t *testing.T) {
        input := &apigateway.GetApiKeysInput{}
        output := &apigateway.GetApiKeysOutput{}

        mockClient.On("GetApiKeys", ctx, input).Return(output, nil)

        result, err := mockClient.GetApiKeys(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetAuthorizer", func(t *testing.T) {
        input := &apigateway.GetAuthorizerInput{}
        output := &apigateway.GetAuthorizerOutput{}

        mockClient.On("GetAuthorizer", ctx, input).Return(output, nil)

        result, err := mockClient.GetAuthorizer(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetAuthorizers", func(t *testing.T) {
        input := &apigateway.GetAuthorizersInput{}
        output := &apigateway.GetAuthorizersOutput{}

        mockClient.On("GetAuthorizers", ctx, input).Return(output, nil)

        result, err := mockClient.GetAuthorizers(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetBasePathMapping", func(t *testing.T) {
        input := &apigateway.GetBasePathMappingInput{}
        output := &apigateway.GetBasePathMappingOutput{}

        mockClient.On("GetBasePathMapping", ctx, input).Return(output, nil)

        result, err := mockClient.GetBasePathMapping(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetBasePathMappings", func(t *testing.T) {
        input := &apigateway.GetBasePathMappingsInput{}
        output := &apigateway.GetBasePathMappingsOutput{}

        mockClient.On("GetBasePathMappings", ctx, input).Return(output, nil)

        result, err := mockClient.GetBasePathMappings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetClientCertificate", func(t *testing.T) {
        input := &apigateway.GetClientCertificateInput{}
        output := &apigateway.GetClientCertificateOutput{}

        mockClient.On("GetClientCertificate", ctx, input).Return(output, nil)

        result, err := mockClient.GetClientCertificate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetClientCertificates", func(t *testing.T) {
        input := &apigateway.GetClientCertificatesInput{}
        output := &apigateway.GetClientCertificatesOutput{}

        mockClient.On("GetClientCertificates", ctx, input).Return(output, nil)

        result, err := mockClient.GetClientCertificates(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDeployment", func(t *testing.T) {
        input := &apigateway.GetDeploymentInput{}
        output := &apigateway.GetDeploymentOutput{}

        mockClient.On("GetDeployment", ctx, input).Return(output, nil)

        result, err := mockClient.GetDeployment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDeployments", func(t *testing.T) {
        input := &apigateway.GetDeploymentsInput{}
        output := &apigateway.GetDeploymentsOutput{}

        mockClient.On("GetDeployments", ctx, input).Return(output, nil)

        result, err := mockClient.GetDeployments(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDocumentationPart", func(t *testing.T) {
        input := &apigateway.GetDocumentationPartInput{}
        output := &apigateway.GetDocumentationPartOutput{}

        mockClient.On("GetDocumentationPart", ctx, input).Return(output, nil)

        result, err := mockClient.GetDocumentationPart(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDocumentationParts", func(t *testing.T) {
        input := &apigateway.GetDocumentationPartsInput{}
        output := &apigateway.GetDocumentationPartsOutput{}

        mockClient.On("GetDocumentationParts", ctx, input).Return(output, nil)

        result, err := mockClient.GetDocumentationParts(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDocumentationVersion", func(t *testing.T) {
        input := &apigateway.GetDocumentationVersionInput{}
        output := &apigateway.GetDocumentationVersionOutput{}

        mockClient.On("GetDocumentationVersion", ctx, input).Return(output, nil)

        result, err := mockClient.GetDocumentationVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDocumentationVersions", func(t *testing.T) {
        input := &apigateway.GetDocumentationVersionsInput{}
        output := &apigateway.GetDocumentationVersionsOutput{}

        mockClient.On("GetDocumentationVersions", ctx, input).Return(output, nil)

        result, err := mockClient.GetDocumentationVersions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDomainName", func(t *testing.T) {
        input := &apigateway.GetDomainNameInput{}
        output := &apigateway.GetDomainNameOutput{}

        mockClient.On("GetDomainName", ctx, input).Return(output, nil)

        result, err := mockClient.GetDomainName(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDomainNames", func(t *testing.T) {
        input := &apigateway.GetDomainNamesInput{}
        output := &apigateway.GetDomainNamesOutput{}

        mockClient.On("GetDomainNames", ctx, input).Return(output, nil)

        result, err := mockClient.GetDomainNames(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetExport", func(t *testing.T) {
        input := &apigateway.GetExportInput{}
        output := &apigateway.GetExportOutput{}

        mockClient.On("GetExport", ctx, input).Return(output, nil)

        result, err := mockClient.GetExport(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetGatewayResponse", func(t *testing.T) {
        input := &apigateway.GetGatewayResponseInput{}
        output := &apigateway.GetGatewayResponseOutput{}

        mockClient.On("GetGatewayResponse", ctx, input).Return(output, nil)

        result, err := mockClient.GetGatewayResponse(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetGatewayResponses", func(t *testing.T) {
        input := &apigateway.GetGatewayResponsesInput{}
        output := &apigateway.GetGatewayResponsesOutput{}

        mockClient.On("GetGatewayResponses", ctx, input).Return(output, nil)

        result, err := mockClient.GetGatewayResponses(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetIntegration", func(t *testing.T) {
        input := &apigateway.GetIntegrationInput{}
        output := &apigateway.GetIntegrationOutput{}

        mockClient.On("GetIntegration", ctx, input).Return(output, nil)

        result, err := mockClient.GetIntegration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetIntegrationResponse", func(t *testing.T) {
        input := &apigateway.GetIntegrationResponseInput{}
        output := &apigateway.GetIntegrationResponseOutput{}

        mockClient.On("GetIntegrationResponse", ctx, input).Return(output, nil)

        result, err := mockClient.GetIntegrationResponse(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetMethod", func(t *testing.T) {
        input := &apigateway.GetMethodInput{}
        output := &apigateway.GetMethodOutput{}

        mockClient.On("GetMethod", ctx, input).Return(output, nil)

        result, err := mockClient.GetMethod(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetMethodResponse", func(t *testing.T) {
        input := &apigateway.GetMethodResponseInput{}
        output := &apigateway.GetMethodResponseOutput{}

        mockClient.On("GetMethodResponse", ctx, input).Return(output, nil)

        result, err := mockClient.GetMethodResponse(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetModel", func(t *testing.T) {
        input := &apigateway.GetModelInput{}
        output := &apigateway.GetModelOutput{}

        mockClient.On("GetModel", ctx, input).Return(output, nil)

        result, err := mockClient.GetModel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetModelTemplate", func(t *testing.T) {
        input := &apigateway.GetModelTemplateInput{}
        output := &apigateway.GetModelTemplateOutput{}

        mockClient.On("GetModelTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.GetModelTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetModels", func(t *testing.T) {
        input := &apigateway.GetModelsInput{}
        output := &apigateway.GetModelsOutput{}

        mockClient.On("GetModels", ctx, input).Return(output, nil)

        result, err := mockClient.GetModels(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetRequestValidator", func(t *testing.T) {
        input := &apigateway.GetRequestValidatorInput{}
        output := &apigateway.GetRequestValidatorOutput{}

        mockClient.On("GetRequestValidator", ctx, input).Return(output, nil)

        result, err := mockClient.GetRequestValidator(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetRequestValidators", func(t *testing.T) {
        input := &apigateway.GetRequestValidatorsInput{}
        output := &apigateway.GetRequestValidatorsOutput{}

        mockClient.On("GetRequestValidators", ctx, input).Return(output, nil)

        result, err := mockClient.GetRequestValidators(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetResource", func(t *testing.T) {
        input := &apigateway.GetResourceInput{}
        output := &apigateway.GetResourceOutput{}

        mockClient.On("GetResource", ctx, input).Return(output, nil)

        result, err := mockClient.GetResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetResources", func(t *testing.T) {
        input := &apigateway.GetResourcesInput{}
        output := &apigateway.GetResourcesOutput{}

        mockClient.On("GetResources", ctx, input).Return(output, nil)

        result, err := mockClient.GetResources(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetRestApi", func(t *testing.T) {
        input := &apigateway.GetRestApiInput{}
        output := &apigateway.GetRestApiOutput{}

        mockClient.On("GetRestApi", ctx, input).Return(output, nil)

        result, err := mockClient.GetRestApi(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetRestApis", func(t *testing.T) {
        input := &apigateway.GetRestApisInput{}
        output := &apigateway.GetRestApisOutput{}

        mockClient.On("GetRestApis", ctx, input).Return(output, nil)

        result, err := mockClient.GetRestApis(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSdk", func(t *testing.T) {
        input := &apigateway.GetSdkInput{}
        output := &apigateway.GetSdkOutput{}

        mockClient.On("GetSdk", ctx, input).Return(output, nil)

        result, err := mockClient.GetSdk(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSdkType", func(t *testing.T) {
        input := &apigateway.GetSdkTypeInput{}
        output := &apigateway.GetSdkTypeOutput{}

        mockClient.On("GetSdkType", ctx, input).Return(output, nil)

        result, err := mockClient.GetSdkType(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSdkTypes", func(t *testing.T) {
        input := &apigateway.GetSdkTypesInput{}
        output := &apigateway.GetSdkTypesOutput{}

        mockClient.On("GetSdkTypes", ctx, input).Return(output, nil)

        result, err := mockClient.GetSdkTypes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetStage", func(t *testing.T) {
        input := &apigateway.GetStageInput{}
        output := &apigateway.GetStageOutput{}

        mockClient.On("GetStage", ctx, input).Return(output, nil)

        result, err := mockClient.GetStage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetStages", func(t *testing.T) {
        input := &apigateway.GetStagesInput{}
        output := &apigateway.GetStagesOutput{}

        mockClient.On("GetStages", ctx, input).Return(output, nil)

        result, err := mockClient.GetStages(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetTags", func(t *testing.T) {
        input := &apigateway.GetTagsInput{}
        output := &apigateway.GetTagsOutput{}

        mockClient.On("GetTags", ctx, input).Return(output, nil)

        result, err := mockClient.GetTags(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetUsage", func(t *testing.T) {
        input := &apigateway.GetUsageInput{}
        output := &apigateway.GetUsageOutput{}

        mockClient.On("GetUsage", ctx, input).Return(output, nil)

        result, err := mockClient.GetUsage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetUsagePlan", func(t *testing.T) {
        input := &apigateway.GetUsagePlanInput{}
        output := &apigateway.GetUsagePlanOutput{}

        mockClient.On("GetUsagePlan", ctx, input).Return(output, nil)

        result, err := mockClient.GetUsagePlan(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetUsagePlanKey", func(t *testing.T) {
        input := &apigateway.GetUsagePlanKeyInput{}
        output := &apigateway.GetUsagePlanKeyOutput{}

        mockClient.On("GetUsagePlanKey", ctx, input).Return(output, nil)

        result, err := mockClient.GetUsagePlanKey(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetUsagePlanKeys", func(t *testing.T) {
        input := &apigateway.GetUsagePlanKeysInput{}
        output := &apigateway.GetUsagePlanKeysOutput{}

        mockClient.On("GetUsagePlanKeys", ctx, input).Return(output, nil)

        result, err := mockClient.GetUsagePlanKeys(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetUsagePlans", func(t *testing.T) {
        input := &apigateway.GetUsagePlansInput{}
        output := &apigateway.GetUsagePlansOutput{}

        mockClient.On("GetUsagePlans", ctx, input).Return(output, nil)

        result, err := mockClient.GetUsagePlans(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetVpcLink", func(t *testing.T) {
        input := &apigateway.GetVpcLinkInput{}
        output := &apigateway.GetVpcLinkOutput{}

        mockClient.On("GetVpcLink", ctx, input).Return(output, nil)

        result, err := mockClient.GetVpcLink(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetVpcLinks", func(t *testing.T) {
        input := &apigateway.GetVpcLinksInput{}
        output := &apigateway.GetVpcLinksOutput{}

        mockClient.On("GetVpcLinks", ctx, input).Return(output, nil)

        result, err := mockClient.GetVpcLinks(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestImportApiKeys", func(t *testing.T) {
        input := &apigateway.ImportApiKeysInput{}
        output := &apigateway.ImportApiKeysOutput{}

        mockClient.On("ImportApiKeys", ctx, input).Return(output, nil)

        result, err := mockClient.ImportApiKeys(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestImportDocumentationParts", func(t *testing.T) {
        input := &apigateway.ImportDocumentationPartsInput{}
        output := &apigateway.ImportDocumentationPartsOutput{}

        mockClient.On("ImportDocumentationParts", ctx, input).Return(output, nil)

        result, err := mockClient.ImportDocumentationParts(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestImportRestApi", func(t *testing.T) {
        input := &apigateway.ImportRestApiInput{}
        output := &apigateway.ImportRestApiOutput{}

        mockClient.On("ImportRestApi", ctx, input).Return(output, nil)

        result, err := mockClient.ImportRestApi(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutGatewayResponse", func(t *testing.T) {
        input := &apigateway.PutGatewayResponseInput{}
        output := &apigateway.PutGatewayResponseOutput{}

        mockClient.On("PutGatewayResponse", ctx, input).Return(output, nil)

        result, err := mockClient.PutGatewayResponse(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutIntegration", func(t *testing.T) {
        input := &apigateway.PutIntegrationInput{}
        output := &apigateway.PutIntegrationOutput{}

        mockClient.On("PutIntegration", ctx, input).Return(output, nil)

        result, err := mockClient.PutIntegration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutIntegrationResponse", func(t *testing.T) {
        input := &apigateway.PutIntegrationResponseInput{}
        output := &apigateway.PutIntegrationResponseOutput{}

        mockClient.On("PutIntegrationResponse", ctx, input).Return(output, nil)

        result, err := mockClient.PutIntegrationResponse(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutMethod", func(t *testing.T) {
        input := &apigateway.PutMethodInput{}
        output := &apigateway.PutMethodOutput{}

        mockClient.On("PutMethod", ctx, input).Return(output, nil)

        result, err := mockClient.PutMethod(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutMethodResponse", func(t *testing.T) {
        input := &apigateway.PutMethodResponseInput{}
        output := &apigateway.PutMethodResponseOutput{}

        mockClient.On("PutMethodResponse", ctx, input).Return(output, nil)

        result, err := mockClient.PutMethodResponse(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutRestApi", func(t *testing.T) {
        input := &apigateway.PutRestApiInput{}
        output := &apigateway.PutRestApiOutput{}

        mockClient.On("PutRestApi", ctx, input).Return(output, nil)

        result, err := mockClient.PutRestApi(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &apigateway.TagResourceInput{}
        output := &apigateway.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTestInvokeAuthorizer", func(t *testing.T) {
        input := &apigateway.TestInvokeAuthorizerInput{}
        output := &apigateway.TestInvokeAuthorizerOutput{}

        mockClient.On("TestInvokeAuthorizer", ctx, input).Return(output, nil)

        result, err := mockClient.TestInvokeAuthorizer(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTestInvokeMethod", func(t *testing.T) {
        input := &apigateway.TestInvokeMethodInput{}
        output := &apigateway.TestInvokeMethodOutput{}

        mockClient.On("TestInvokeMethod", ctx, input).Return(output, nil)

        result, err := mockClient.TestInvokeMethod(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &apigateway.UntagResourceInput{}
        output := &apigateway.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateAccount", func(t *testing.T) {
        input := &apigateway.UpdateAccountInput{}
        output := &apigateway.UpdateAccountOutput{}

        mockClient.On("UpdateAccount", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateAccount(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateApiKey", func(t *testing.T) {
        input := &apigateway.UpdateApiKeyInput{}
        output := &apigateway.UpdateApiKeyOutput{}

        mockClient.On("UpdateApiKey", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateApiKey(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateAuthorizer", func(t *testing.T) {
        input := &apigateway.UpdateAuthorizerInput{}
        output := &apigateway.UpdateAuthorizerOutput{}

        mockClient.On("UpdateAuthorizer", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateAuthorizer(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateBasePathMapping", func(t *testing.T) {
        input := &apigateway.UpdateBasePathMappingInput{}
        output := &apigateway.UpdateBasePathMappingOutput{}

        mockClient.On("UpdateBasePathMapping", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateBasePathMapping(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateClientCertificate", func(t *testing.T) {
        input := &apigateway.UpdateClientCertificateInput{}
        output := &apigateway.UpdateClientCertificateOutput{}

        mockClient.On("UpdateClientCertificate", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateClientCertificate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateDeployment", func(t *testing.T) {
        input := &apigateway.UpdateDeploymentInput{}
        output := &apigateway.UpdateDeploymentOutput{}

        mockClient.On("UpdateDeployment", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateDeployment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateDocumentationPart", func(t *testing.T) {
        input := &apigateway.UpdateDocumentationPartInput{}
        output := &apigateway.UpdateDocumentationPartOutput{}

        mockClient.On("UpdateDocumentationPart", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateDocumentationPart(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateDocumentationVersion", func(t *testing.T) {
        input := &apigateway.UpdateDocumentationVersionInput{}
        output := &apigateway.UpdateDocumentationVersionOutput{}

        mockClient.On("UpdateDocumentationVersion", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateDocumentationVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateDomainName", func(t *testing.T) {
        input := &apigateway.UpdateDomainNameInput{}
        output := &apigateway.UpdateDomainNameOutput{}

        mockClient.On("UpdateDomainName", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateDomainName(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateGatewayResponse", func(t *testing.T) {
        input := &apigateway.UpdateGatewayResponseInput{}
        output := &apigateway.UpdateGatewayResponseOutput{}

        mockClient.On("UpdateGatewayResponse", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateGatewayResponse(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateIntegration", func(t *testing.T) {
        input := &apigateway.UpdateIntegrationInput{}
        output := &apigateway.UpdateIntegrationOutput{}

        mockClient.On("UpdateIntegration", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateIntegration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateIntegrationResponse", func(t *testing.T) {
        input := &apigateway.UpdateIntegrationResponseInput{}
        output := &apigateway.UpdateIntegrationResponseOutput{}

        mockClient.On("UpdateIntegrationResponse", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateIntegrationResponse(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateMethod", func(t *testing.T) {
        input := &apigateway.UpdateMethodInput{}
        output := &apigateway.UpdateMethodOutput{}

        mockClient.On("UpdateMethod", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateMethod(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateMethodResponse", func(t *testing.T) {
        input := &apigateway.UpdateMethodResponseInput{}
        output := &apigateway.UpdateMethodResponseOutput{}

        mockClient.On("UpdateMethodResponse", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateMethodResponse(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateModel", func(t *testing.T) {
        input := &apigateway.UpdateModelInput{}
        output := &apigateway.UpdateModelOutput{}

        mockClient.On("UpdateModel", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateModel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateRequestValidator", func(t *testing.T) {
        input := &apigateway.UpdateRequestValidatorInput{}
        output := &apigateway.UpdateRequestValidatorOutput{}

        mockClient.On("UpdateRequestValidator", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateRequestValidator(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateResource", func(t *testing.T) {
        input := &apigateway.UpdateResourceInput{}
        output := &apigateway.UpdateResourceOutput{}

        mockClient.On("UpdateResource", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateRestApi", func(t *testing.T) {
        input := &apigateway.UpdateRestApiInput{}
        output := &apigateway.UpdateRestApiOutput{}

        mockClient.On("UpdateRestApi", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateRestApi(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateStage", func(t *testing.T) {
        input := &apigateway.UpdateStageInput{}
        output := &apigateway.UpdateStageOutput{}

        mockClient.On("UpdateStage", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateStage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateUsage", func(t *testing.T) {
        input := &apigateway.UpdateUsageInput{}
        output := &apigateway.UpdateUsageOutput{}

        mockClient.On("UpdateUsage", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateUsage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateUsagePlan", func(t *testing.T) {
        input := &apigateway.UpdateUsagePlanInput{}
        output := &apigateway.UpdateUsagePlanOutput{}

        mockClient.On("UpdateUsagePlan", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateUsagePlan(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateVpcLink", func(t *testing.T) {
        input := &apigateway.UpdateVpcLinkInput{}
        output := &apigateway.UpdateVpcLinkOutput{}

        mockClient.On("UpdateVpcLink", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateVpcLink(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
