package apigatewayv2_test

// tests for the apigatewayv2 service interface for this ../../../service/apigatewayv2/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/apigatewayv2/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/apigatewayv2/apigatewayv2_iface"
	"github.com/stretchr/testify/assert"
)

func TestApigatewayv2ServiceCanBeMocked(t *testing.T) {
	var iface apigatewayv2_iface.IClient
	iface = &apigatewayv2.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := apigatewayv2.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateApi", func(t *testing.T) {
        input := &apigatewayv2.CreateApiInput{}
        output := &apigatewayv2.CreateApiOutput{}

        mockClient.On("CreateApi", ctx, input).Return(output, nil)

        result, err := mockClient.CreateApi(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateApiMapping", func(t *testing.T) {
        input := &apigatewayv2.CreateApiMappingInput{}
        output := &apigatewayv2.CreateApiMappingOutput{}

        mockClient.On("CreateApiMapping", ctx, input).Return(output, nil)

        result, err := mockClient.CreateApiMapping(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateAuthorizer", func(t *testing.T) {
        input := &apigatewayv2.CreateAuthorizerInput{}
        output := &apigatewayv2.CreateAuthorizerOutput{}

        mockClient.On("CreateAuthorizer", ctx, input).Return(output, nil)

        result, err := mockClient.CreateAuthorizer(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateDeployment", func(t *testing.T) {
        input := &apigatewayv2.CreateDeploymentInput{}
        output := &apigatewayv2.CreateDeploymentOutput{}

        mockClient.On("CreateDeployment", ctx, input).Return(output, nil)

        result, err := mockClient.CreateDeployment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateDomainName", func(t *testing.T) {
        input := &apigatewayv2.CreateDomainNameInput{}
        output := &apigatewayv2.CreateDomainNameOutput{}

        mockClient.On("CreateDomainName", ctx, input).Return(output, nil)

        result, err := mockClient.CreateDomainName(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateIntegration", func(t *testing.T) {
        input := &apigatewayv2.CreateIntegrationInput{}
        output := &apigatewayv2.CreateIntegrationOutput{}

        mockClient.On("CreateIntegration", ctx, input).Return(output, nil)

        result, err := mockClient.CreateIntegration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateIntegrationResponse", func(t *testing.T) {
        input := &apigatewayv2.CreateIntegrationResponseInput{}
        output := &apigatewayv2.CreateIntegrationResponseOutput{}

        mockClient.On("CreateIntegrationResponse", ctx, input).Return(output, nil)

        result, err := mockClient.CreateIntegrationResponse(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateModel", func(t *testing.T) {
        input := &apigatewayv2.CreateModelInput{}
        output := &apigatewayv2.CreateModelOutput{}

        mockClient.On("CreateModel", ctx, input).Return(output, nil)

        result, err := mockClient.CreateModel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateRoute", func(t *testing.T) {
        input := &apigatewayv2.CreateRouteInput{}
        output := &apigatewayv2.CreateRouteOutput{}

        mockClient.On("CreateRoute", ctx, input).Return(output, nil)

        result, err := mockClient.CreateRoute(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateRouteResponse", func(t *testing.T) {
        input := &apigatewayv2.CreateRouteResponseInput{}
        output := &apigatewayv2.CreateRouteResponseOutput{}

        mockClient.On("CreateRouteResponse", ctx, input).Return(output, nil)

        result, err := mockClient.CreateRouteResponse(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateStage", func(t *testing.T) {
        input := &apigatewayv2.CreateStageInput{}
        output := &apigatewayv2.CreateStageOutput{}

        mockClient.On("CreateStage", ctx, input).Return(output, nil)

        result, err := mockClient.CreateStage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateVpcLink", func(t *testing.T) {
        input := &apigatewayv2.CreateVpcLinkInput{}
        output := &apigatewayv2.CreateVpcLinkOutput{}

        mockClient.On("CreateVpcLink", ctx, input).Return(output, nil)

        result, err := mockClient.CreateVpcLink(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteAccessLogSettings", func(t *testing.T) {
        input := &apigatewayv2.DeleteAccessLogSettingsInput{}
        output := &apigatewayv2.DeleteAccessLogSettingsOutput{}

        mockClient.On("DeleteAccessLogSettings", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteAccessLogSettings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteApi", func(t *testing.T) {
        input := &apigatewayv2.DeleteApiInput{}
        output := &apigatewayv2.DeleteApiOutput{}

        mockClient.On("DeleteApi", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteApi(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteApiMapping", func(t *testing.T) {
        input := &apigatewayv2.DeleteApiMappingInput{}
        output := &apigatewayv2.DeleteApiMappingOutput{}

        mockClient.On("DeleteApiMapping", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteApiMapping(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteAuthorizer", func(t *testing.T) {
        input := &apigatewayv2.DeleteAuthorizerInput{}
        output := &apigatewayv2.DeleteAuthorizerOutput{}

        mockClient.On("DeleteAuthorizer", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteAuthorizer(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteCorsConfiguration", func(t *testing.T) {
        input := &apigatewayv2.DeleteCorsConfigurationInput{}
        output := &apigatewayv2.DeleteCorsConfigurationOutput{}

        mockClient.On("DeleteCorsConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteCorsConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteDeployment", func(t *testing.T) {
        input := &apigatewayv2.DeleteDeploymentInput{}
        output := &apigatewayv2.DeleteDeploymentOutput{}

        mockClient.On("DeleteDeployment", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteDeployment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteDomainName", func(t *testing.T) {
        input := &apigatewayv2.DeleteDomainNameInput{}
        output := &apigatewayv2.DeleteDomainNameOutput{}

        mockClient.On("DeleteDomainName", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteDomainName(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteIntegration", func(t *testing.T) {
        input := &apigatewayv2.DeleteIntegrationInput{}
        output := &apigatewayv2.DeleteIntegrationOutput{}

        mockClient.On("DeleteIntegration", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteIntegration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteIntegrationResponse", func(t *testing.T) {
        input := &apigatewayv2.DeleteIntegrationResponseInput{}
        output := &apigatewayv2.DeleteIntegrationResponseOutput{}

        mockClient.On("DeleteIntegrationResponse", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteIntegrationResponse(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteModel", func(t *testing.T) {
        input := &apigatewayv2.DeleteModelInput{}
        output := &apigatewayv2.DeleteModelOutput{}

        mockClient.On("DeleteModel", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteModel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteRoute", func(t *testing.T) {
        input := &apigatewayv2.DeleteRouteInput{}
        output := &apigatewayv2.DeleteRouteOutput{}

        mockClient.On("DeleteRoute", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteRoute(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteRouteRequestParameter", func(t *testing.T) {
        input := &apigatewayv2.DeleteRouteRequestParameterInput{}
        output := &apigatewayv2.DeleteRouteRequestParameterOutput{}

        mockClient.On("DeleteRouteRequestParameter", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteRouteRequestParameter(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteRouteResponse", func(t *testing.T) {
        input := &apigatewayv2.DeleteRouteResponseInput{}
        output := &apigatewayv2.DeleteRouteResponseOutput{}

        mockClient.On("DeleteRouteResponse", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteRouteResponse(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteRouteSettings", func(t *testing.T) {
        input := &apigatewayv2.DeleteRouteSettingsInput{}
        output := &apigatewayv2.DeleteRouteSettingsOutput{}

        mockClient.On("DeleteRouteSettings", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteRouteSettings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteStage", func(t *testing.T) {
        input := &apigatewayv2.DeleteStageInput{}
        output := &apigatewayv2.DeleteStageOutput{}

        mockClient.On("DeleteStage", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteStage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteVpcLink", func(t *testing.T) {
        input := &apigatewayv2.DeleteVpcLinkInput{}
        output := &apigatewayv2.DeleteVpcLinkOutput{}

        mockClient.On("DeleteVpcLink", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteVpcLink(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestExportApi", func(t *testing.T) {
        input := &apigatewayv2.ExportApiInput{}
        output := &apigatewayv2.ExportApiOutput{}

        mockClient.On("ExportApi", ctx, input).Return(output, nil)

        result, err := mockClient.ExportApi(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetApi", func(t *testing.T) {
        input := &apigatewayv2.GetApiInput{}
        output := &apigatewayv2.GetApiOutput{}

        mockClient.On("GetApi", ctx, input).Return(output, nil)

        result, err := mockClient.GetApi(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetApiMapping", func(t *testing.T) {
        input := &apigatewayv2.GetApiMappingInput{}
        output := &apigatewayv2.GetApiMappingOutput{}

        mockClient.On("GetApiMapping", ctx, input).Return(output, nil)

        result, err := mockClient.GetApiMapping(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetApiMappings", func(t *testing.T) {
        input := &apigatewayv2.GetApiMappingsInput{}
        output := &apigatewayv2.GetApiMappingsOutput{}

        mockClient.On("GetApiMappings", ctx, input).Return(output, nil)

        result, err := mockClient.GetApiMappings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetApis", func(t *testing.T) {
        input := &apigatewayv2.GetApisInput{}
        output := &apigatewayv2.GetApisOutput{}

        mockClient.On("GetApis", ctx, input).Return(output, nil)

        result, err := mockClient.GetApis(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetAuthorizer", func(t *testing.T) {
        input := &apigatewayv2.GetAuthorizerInput{}
        output := &apigatewayv2.GetAuthorizerOutput{}

        mockClient.On("GetAuthorizer", ctx, input).Return(output, nil)

        result, err := mockClient.GetAuthorizer(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetAuthorizers", func(t *testing.T) {
        input := &apigatewayv2.GetAuthorizersInput{}
        output := &apigatewayv2.GetAuthorizersOutput{}

        mockClient.On("GetAuthorizers", ctx, input).Return(output, nil)

        result, err := mockClient.GetAuthorizers(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDeployment", func(t *testing.T) {
        input := &apigatewayv2.GetDeploymentInput{}
        output := &apigatewayv2.GetDeploymentOutput{}

        mockClient.On("GetDeployment", ctx, input).Return(output, nil)

        result, err := mockClient.GetDeployment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDeployments", func(t *testing.T) {
        input := &apigatewayv2.GetDeploymentsInput{}
        output := &apigatewayv2.GetDeploymentsOutput{}

        mockClient.On("GetDeployments", ctx, input).Return(output, nil)

        result, err := mockClient.GetDeployments(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDomainName", func(t *testing.T) {
        input := &apigatewayv2.GetDomainNameInput{}
        output := &apigatewayv2.GetDomainNameOutput{}

        mockClient.On("GetDomainName", ctx, input).Return(output, nil)

        result, err := mockClient.GetDomainName(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDomainNames", func(t *testing.T) {
        input := &apigatewayv2.GetDomainNamesInput{}
        output := &apigatewayv2.GetDomainNamesOutput{}

        mockClient.On("GetDomainNames", ctx, input).Return(output, nil)

        result, err := mockClient.GetDomainNames(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetIntegration", func(t *testing.T) {
        input := &apigatewayv2.GetIntegrationInput{}
        output := &apigatewayv2.GetIntegrationOutput{}

        mockClient.On("GetIntegration", ctx, input).Return(output, nil)

        result, err := mockClient.GetIntegration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetIntegrationResponse", func(t *testing.T) {
        input := &apigatewayv2.GetIntegrationResponseInput{}
        output := &apigatewayv2.GetIntegrationResponseOutput{}

        mockClient.On("GetIntegrationResponse", ctx, input).Return(output, nil)

        result, err := mockClient.GetIntegrationResponse(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetIntegrationResponses", func(t *testing.T) {
        input := &apigatewayv2.GetIntegrationResponsesInput{}
        output := &apigatewayv2.GetIntegrationResponsesOutput{}

        mockClient.On("GetIntegrationResponses", ctx, input).Return(output, nil)

        result, err := mockClient.GetIntegrationResponses(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetIntegrations", func(t *testing.T) {
        input := &apigatewayv2.GetIntegrationsInput{}
        output := &apigatewayv2.GetIntegrationsOutput{}

        mockClient.On("GetIntegrations", ctx, input).Return(output, nil)

        result, err := mockClient.GetIntegrations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetModel", func(t *testing.T) {
        input := &apigatewayv2.GetModelInput{}
        output := &apigatewayv2.GetModelOutput{}

        mockClient.On("GetModel", ctx, input).Return(output, nil)

        result, err := mockClient.GetModel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetModelTemplate", func(t *testing.T) {
        input := &apigatewayv2.GetModelTemplateInput{}
        output := &apigatewayv2.GetModelTemplateOutput{}

        mockClient.On("GetModelTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.GetModelTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetModels", func(t *testing.T) {
        input := &apigatewayv2.GetModelsInput{}
        output := &apigatewayv2.GetModelsOutput{}

        mockClient.On("GetModels", ctx, input).Return(output, nil)

        result, err := mockClient.GetModels(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetRoute", func(t *testing.T) {
        input := &apigatewayv2.GetRouteInput{}
        output := &apigatewayv2.GetRouteOutput{}

        mockClient.On("GetRoute", ctx, input).Return(output, nil)

        result, err := mockClient.GetRoute(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetRouteResponse", func(t *testing.T) {
        input := &apigatewayv2.GetRouteResponseInput{}
        output := &apigatewayv2.GetRouteResponseOutput{}

        mockClient.On("GetRouteResponse", ctx, input).Return(output, nil)

        result, err := mockClient.GetRouteResponse(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetRouteResponses", func(t *testing.T) {
        input := &apigatewayv2.GetRouteResponsesInput{}
        output := &apigatewayv2.GetRouteResponsesOutput{}

        mockClient.On("GetRouteResponses", ctx, input).Return(output, nil)

        result, err := mockClient.GetRouteResponses(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetRoutes", func(t *testing.T) {
        input := &apigatewayv2.GetRoutesInput{}
        output := &apigatewayv2.GetRoutesOutput{}

        mockClient.On("GetRoutes", ctx, input).Return(output, nil)

        result, err := mockClient.GetRoutes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetStage", func(t *testing.T) {
        input := &apigatewayv2.GetStageInput{}
        output := &apigatewayv2.GetStageOutput{}

        mockClient.On("GetStage", ctx, input).Return(output, nil)

        result, err := mockClient.GetStage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetStages", func(t *testing.T) {
        input := &apigatewayv2.GetStagesInput{}
        output := &apigatewayv2.GetStagesOutput{}

        mockClient.On("GetStages", ctx, input).Return(output, nil)

        result, err := mockClient.GetStages(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetTags", func(t *testing.T) {
        input := &apigatewayv2.GetTagsInput{}
        output := &apigatewayv2.GetTagsOutput{}

        mockClient.On("GetTags", ctx, input).Return(output, nil)

        result, err := mockClient.GetTags(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetVpcLink", func(t *testing.T) {
        input := &apigatewayv2.GetVpcLinkInput{}
        output := &apigatewayv2.GetVpcLinkOutput{}

        mockClient.On("GetVpcLink", ctx, input).Return(output, nil)

        result, err := mockClient.GetVpcLink(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetVpcLinks", func(t *testing.T) {
        input := &apigatewayv2.GetVpcLinksInput{}
        output := &apigatewayv2.GetVpcLinksOutput{}

        mockClient.On("GetVpcLinks", ctx, input).Return(output, nil)

        result, err := mockClient.GetVpcLinks(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestImportApi", func(t *testing.T) {
        input := &apigatewayv2.ImportApiInput{}
        output := &apigatewayv2.ImportApiOutput{}

        mockClient.On("ImportApi", ctx, input).Return(output, nil)

        result, err := mockClient.ImportApi(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestReimportApi", func(t *testing.T) {
        input := &apigatewayv2.ReimportApiInput{}
        output := &apigatewayv2.ReimportApiOutput{}

        mockClient.On("ReimportApi", ctx, input).Return(output, nil)

        result, err := mockClient.ReimportApi(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestResetAuthorizersCache", func(t *testing.T) {
        input := &apigatewayv2.ResetAuthorizersCacheInput{}
        output := &apigatewayv2.ResetAuthorizersCacheOutput{}

        mockClient.On("ResetAuthorizersCache", ctx, input).Return(output, nil)

        result, err := mockClient.ResetAuthorizersCache(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &apigatewayv2.TagResourceInput{}
        output := &apigatewayv2.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &apigatewayv2.UntagResourceInput{}
        output := &apigatewayv2.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateApi", func(t *testing.T) {
        input := &apigatewayv2.UpdateApiInput{}
        output := &apigatewayv2.UpdateApiOutput{}

        mockClient.On("UpdateApi", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateApi(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateApiMapping", func(t *testing.T) {
        input := &apigatewayv2.UpdateApiMappingInput{}
        output := &apigatewayv2.UpdateApiMappingOutput{}

        mockClient.On("UpdateApiMapping", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateApiMapping(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateAuthorizer", func(t *testing.T) {
        input := &apigatewayv2.UpdateAuthorizerInput{}
        output := &apigatewayv2.UpdateAuthorizerOutput{}

        mockClient.On("UpdateAuthorizer", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateAuthorizer(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateDeployment", func(t *testing.T) {
        input := &apigatewayv2.UpdateDeploymentInput{}
        output := &apigatewayv2.UpdateDeploymentOutput{}

        mockClient.On("UpdateDeployment", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateDeployment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateDomainName", func(t *testing.T) {
        input := &apigatewayv2.UpdateDomainNameInput{}
        output := &apigatewayv2.UpdateDomainNameOutput{}

        mockClient.On("UpdateDomainName", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateDomainName(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateIntegration", func(t *testing.T) {
        input := &apigatewayv2.UpdateIntegrationInput{}
        output := &apigatewayv2.UpdateIntegrationOutput{}

        mockClient.On("UpdateIntegration", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateIntegration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateIntegrationResponse", func(t *testing.T) {
        input := &apigatewayv2.UpdateIntegrationResponseInput{}
        output := &apigatewayv2.UpdateIntegrationResponseOutput{}

        mockClient.On("UpdateIntegrationResponse", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateIntegrationResponse(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateModel", func(t *testing.T) {
        input := &apigatewayv2.UpdateModelInput{}
        output := &apigatewayv2.UpdateModelOutput{}

        mockClient.On("UpdateModel", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateModel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateRoute", func(t *testing.T) {
        input := &apigatewayv2.UpdateRouteInput{}
        output := &apigatewayv2.UpdateRouteOutput{}

        mockClient.On("UpdateRoute", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateRoute(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateRouteResponse", func(t *testing.T) {
        input := &apigatewayv2.UpdateRouteResponseInput{}
        output := &apigatewayv2.UpdateRouteResponseOutput{}

        mockClient.On("UpdateRouteResponse", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateRouteResponse(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateStage", func(t *testing.T) {
        input := &apigatewayv2.UpdateStageInput{}
        output := &apigatewayv2.UpdateStageOutput{}

        mockClient.On("UpdateStage", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateStage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateVpcLink", func(t *testing.T) {
        input := &apigatewayv2.UpdateVpcLinkInput{}
        output := &apigatewayv2.UpdateVpcLinkOutput{}

        mockClient.On("UpdateVpcLink", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateVpcLink(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
