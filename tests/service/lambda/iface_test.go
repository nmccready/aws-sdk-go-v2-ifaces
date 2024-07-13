package lambda_test

// tests for the lambda service interface for this ../../../service/lambda/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/lambda/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/lambda/lambda_iface"
	"github.com/stretchr/testify/assert"
)

func TestLambdaServiceCanBeMocked(t *testing.T) {
	var iface lambda_iface.IClient
	iface = &lambda.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := lambda.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAddLayerVersionPermission", func(t *testing.T) {
        input := &lambda.AddLayerVersionPermissionInput{}
        output := &lambda.AddLayerVersionPermissionOutput{}

        mockClient.On("AddLayerVersionPermission", ctx, input).Return(output, nil)

        result, err := mockClient.AddLayerVersionPermission(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAddPermission", func(t *testing.T) {
        input := &lambda.AddPermissionInput{}
        output := &lambda.AddPermissionOutput{}

        mockClient.On("AddPermission", ctx, input).Return(output, nil)

        result, err := mockClient.AddPermission(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateAlias", func(t *testing.T) {
        input := &lambda.CreateAliasInput{}
        output := &lambda.CreateAliasOutput{}

        mockClient.On("CreateAlias", ctx, input).Return(output, nil)

        result, err := mockClient.CreateAlias(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateCodeSigningConfig", func(t *testing.T) {
        input := &lambda.CreateCodeSigningConfigInput{}
        output := &lambda.CreateCodeSigningConfigOutput{}

        mockClient.On("CreateCodeSigningConfig", ctx, input).Return(output, nil)

        result, err := mockClient.CreateCodeSigningConfig(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateEventSourceMapping", func(t *testing.T) {
        input := &lambda.CreateEventSourceMappingInput{}
        output := &lambda.CreateEventSourceMappingOutput{}

        mockClient.On("CreateEventSourceMapping", ctx, input).Return(output, nil)

        result, err := mockClient.CreateEventSourceMapping(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateFunction", func(t *testing.T) {
        input := &lambda.CreateFunctionInput{}
        output := &lambda.CreateFunctionOutput{}

        mockClient.On("CreateFunction", ctx, input).Return(output, nil)

        result, err := mockClient.CreateFunction(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateFunctionUrlConfig", func(t *testing.T) {
        input := &lambda.CreateFunctionUrlConfigInput{}
        output := &lambda.CreateFunctionUrlConfigOutput{}

        mockClient.On("CreateFunctionUrlConfig", ctx, input).Return(output, nil)

        result, err := mockClient.CreateFunctionUrlConfig(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteAlias", func(t *testing.T) {
        input := &lambda.DeleteAliasInput{}
        output := &lambda.DeleteAliasOutput{}

        mockClient.On("DeleteAlias", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteAlias(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteCodeSigningConfig", func(t *testing.T) {
        input := &lambda.DeleteCodeSigningConfigInput{}
        output := &lambda.DeleteCodeSigningConfigOutput{}

        mockClient.On("DeleteCodeSigningConfig", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteCodeSigningConfig(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteEventSourceMapping", func(t *testing.T) {
        input := &lambda.DeleteEventSourceMappingInput{}
        output := &lambda.DeleteEventSourceMappingOutput{}

        mockClient.On("DeleteEventSourceMapping", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteEventSourceMapping(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteFunction", func(t *testing.T) {
        input := &lambda.DeleteFunctionInput{}
        output := &lambda.DeleteFunctionOutput{}

        mockClient.On("DeleteFunction", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteFunction(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteFunctionCodeSigningConfig", func(t *testing.T) {
        input := &lambda.DeleteFunctionCodeSigningConfigInput{}
        output := &lambda.DeleteFunctionCodeSigningConfigOutput{}

        mockClient.On("DeleteFunctionCodeSigningConfig", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteFunctionCodeSigningConfig(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteFunctionConcurrency", func(t *testing.T) {
        input := &lambda.DeleteFunctionConcurrencyInput{}
        output := &lambda.DeleteFunctionConcurrencyOutput{}

        mockClient.On("DeleteFunctionConcurrency", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteFunctionConcurrency(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteFunctionEventInvokeConfig", func(t *testing.T) {
        input := &lambda.DeleteFunctionEventInvokeConfigInput{}
        output := &lambda.DeleteFunctionEventInvokeConfigOutput{}

        mockClient.On("DeleteFunctionEventInvokeConfig", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteFunctionEventInvokeConfig(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteFunctionUrlConfig", func(t *testing.T) {
        input := &lambda.DeleteFunctionUrlConfigInput{}
        output := &lambda.DeleteFunctionUrlConfigOutput{}

        mockClient.On("DeleteFunctionUrlConfig", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteFunctionUrlConfig(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteLayerVersion", func(t *testing.T) {
        input := &lambda.DeleteLayerVersionInput{}
        output := &lambda.DeleteLayerVersionOutput{}

        mockClient.On("DeleteLayerVersion", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteLayerVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteProvisionedConcurrencyConfig", func(t *testing.T) {
        input := &lambda.DeleteProvisionedConcurrencyConfigInput{}
        output := &lambda.DeleteProvisionedConcurrencyConfigOutput{}

        mockClient.On("DeleteProvisionedConcurrencyConfig", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteProvisionedConcurrencyConfig(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetAccountSettings", func(t *testing.T) {
        input := &lambda.GetAccountSettingsInput{}
        output := &lambda.GetAccountSettingsOutput{}

        mockClient.On("GetAccountSettings", ctx, input).Return(output, nil)

        result, err := mockClient.GetAccountSettings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetAlias", func(t *testing.T) {
        input := &lambda.GetAliasInput{}
        output := &lambda.GetAliasOutput{}

        mockClient.On("GetAlias", ctx, input).Return(output, nil)

        result, err := mockClient.GetAlias(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetCodeSigningConfig", func(t *testing.T) {
        input := &lambda.GetCodeSigningConfigInput{}
        output := &lambda.GetCodeSigningConfigOutput{}

        mockClient.On("GetCodeSigningConfig", ctx, input).Return(output, nil)

        result, err := mockClient.GetCodeSigningConfig(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetEventSourceMapping", func(t *testing.T) {
        input := &lambda.GetEventSourceMappingInput{}
        output := &lambda.GetEventSourceMappingOutput{}

        mockClient.On("GetEventSourceMapping", ctx, input).Return(output, nil)

        result, err := mockClient.GetEventSourceMapping(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetFunction", func(t *testing.T) {
        input := &lambda.GetFunctionInput{}
        output := &lambda.GetFunctionOutput{}

        mockClient.On("GetFunction", ctx, input).Return(output, nil)

        result, err := mockClient.GetFunction(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetFunctionCodeSigningConfig", func(t *testing.T) {
        input := &lambda.GetFunctionCodeSigningConfigInput{}
        output := &lambda.GetFunctionCodeSigningConfigOutput{}

        mockClient.On("GetFunctionCodeSigningConfig", ctx, input).Return(output, nil)

        result, err := mockClient.GetFunctionCodeSigningConfig(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetFunctionConcurrency", func(t *testing.T) {
        input := &lambda.GetFunctionConcurrencyInput{}
        output := &lambda.GetFunctionConcurrencyOutput{}

        mockClient.On("GetFunctionConcurrency", ctx, input).Return(output, nil)

        result, err := mockClient.GetFunctionConcurrency(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetFunctionConfiguration", func(t *testing.T) {
        input := &lambda.GetFunctionConfigurationInput{}
        output := &lambda.GetFunctionConfigurationOutput{}

        mockClient.On("GetFunctionConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.GetFunctionConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetFunctionEventInvokeConfig", func(t *testing.T) {
        input := &lambda.GetFunctionEventInvokeConfigInput{}
        output := &lambda.GetFunctionEventInvokeConfigOutput{}

        mockClient.On("GetFunctionEventInvokeConfig", ctx, input).Return(output, nil)

        result, err := mockClient.GetFunctionEventInvokeConfig(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetFunctionUrlConfig", func(t *testing.T) {
        input := &lambda.GetFunctionUrlConfigInput{}
        output := &lambda.GetFunctionUrlConfigOutput{}

        mockClient.On("GetFunctionUrlConfig", ctx, input).Return(output, nil)

        result, err := mockClient.GetFunctionUrlConfig(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetLayerVersion", func(t *testing.T) {
        input := &lambda.GetLayerVersionInput{}
        output := &lambda.GetLayerVersionOutput{}

        mockClient.On("GetLayerVersion", ctx, input).Return(output, nil)

        result, err := mockClient.GetLayerVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetLayerVersionByArn", func(t *testing.T) {
        input := &lambda.GetLayerVersionByArnInput{}
        output := &lambda.GetLayerVersionByArnOutput{}

        mockClient.On("GetLayerVersionByArn", ctx, input).Return(output, nil)

        result, err := mockClient.GetLayerVersionByArn(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetLayerVersionPolicy", func(t *testing.T) {
        input := &lambda.GetLayerVersionPolicyInput{}
        output := &lambda.GetLayerVersionPolicyOutput{}

        mockClient.On("GetLayerVersionPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.GetLayerVersionPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetPolicy", func(t *testing.T) {
        input := &lambda.GetPolicyInput{}
        output := &lambda.GetPolicyOutput{}

        mockClient.On("GetPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.GetPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetProvisionedConcurrencyConfig", func(t *testing.T) {
        input := &lambda.GetProvisionedConcurrencyConfigInput{}
        output := &lambda.GetProvisionedConcurrencyConfigOutput{}

        mockClient.On("GetProvisionedConcurrencyConfig", ctx, input).Return(output, nil)

        result, err := mockClient.GetProvisionedConcurrencyConfig(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetRuntimeManagementConfig", func(t *testing.T) {
        input := &lambda.GetRuntimeManagementConfigInput{}
        output := &lambda.GetRuntimeManagementConfigOutput{}

        mockClient.On("GetRuntimeManagementConfig", ctx, input).Return(output, nil)

        result, err := mockClient.GetRuntimeManagementConfig(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestInvoke", func(t *testing.T) {
        input := &lambda.InvokeInput{}
        output := &lambda.InvokeOutput{}

        mockClient.On("Invoke", ctx, input).Return(output, nil)

        result, err := mockClient.Invoke(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestInvokeAsync", func(t *testing.T) {
        input := &lambda.InvokeAsyncInput{}
        output := &lambda.InvokeAsyncOutput{}

        mockClient.On("InvokeAsync", ctx, input).Return(output, nil)

        result, err := mockClient.InvokeAsync(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestInvokeWithResponseStream", func(t *testing.T) {
        input := &lambda.InvokeWithResponseStreamInput{}
        output := &lambda.InvokeWithResponseStreamOutput{}

        mockClient.On("InvokeWithResponseStream", ctx, input).Return(output, nil)

        result, err := mockClient.InvokeWithResponseStream(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAliases", func(t *testing.T) {
        input := &lambda.ListAliasesInput{}
        output := &lambda.ListAliasesOutput{}

        mockClient.On("ListAliases", ctx, input).Return(output, nil)

        result, err := mockClient.ListAliases(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListCodeSigningConfigs", func(t *testing.T) {
        input := &lambda.ListCodeSigningConfigsInput{}
        output := &lambda.ListCodeSigningConfigsOutput{}

        mockClient.On("ListCodeSigningConfigs", ctx, input).Return(output, nil)

        result, err := mockClient.ListCodeSigningConfigs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListEventSourceMappings", func(t *testing.T) {
        input := &lambda.ListEventSourceMappingsInput{}
        output := &lambda.ListEventSourceMappingsOutput{}

        mockClient.On("ListEventSourceMappings", ctx, input).Return(output, nil)

        result, err := mockClient.ListEventSourceMappings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListFunctionEventInvokeConfigs", func(t *testing.T) {
        input := &lambda.ListFunctionEventInvokeConfigsInput{}
        output := &lambda.ListFunctionEventInvokeConfigsOutput{}

        mockClient.On("ListFunctionEventInvokeConfigs", ctx, input).Return(output, nil)

        result, err := mockClient.ListFunctionEventInvokeConfigs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListFunctionUrlConfigs", func(t *testing.T) {
        input := &lambda.ListFunctionUrlConfigsInput{}
        output := &lambda.ListFunctionUrlConfigsOutput{}

        mockClient.On("ListFunctionUrlConfigs", ctx, input).Return(output, nil)

        result, err := mockClient.ListFunctionUrlConfigs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListFunctions", func(t *testing.T) {
        input := &lambda.ListFunctionsInput{}
        output := &lambda.ListFunctionsOutput{}

        mockClient.On("ListFunctions", ctx, input).Return(output, nil)

        result, err := mockClient.ListFunctions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListFunctionsByCodeSigningConfig", func(t *testing.T) {
        input := &lambda.ListFunctionsByCodeSigningConfigInput{}
        output := &lambda.ListFunctionsByCodeSigningConfigOutput{}

        mockClient.On("ListFunctionsByCodeSigningConfig", ctx, input).Return(output, nil)

        result, err := mockClient.ListFunctionsByCodeSigningConfig(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListLayerVersions", func(t *testing.T) {
        input := &lambda.ListLayerVersionsInput{}
        output := &lambda.ListLayerVersionsOutput{}

        mockClient.On("ListLayerVersions", ctx, input).Return(output, nil)

        result, err := mockClient.ListLayerVersions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListLayers", func(t *testing.T) {
        input := &lambda.ListLayersInput{}
        output := &lambda.ListLayersOutput{}

        mockClient.On("ListLayers", ctx, input).Return(output, nil)

        result, err := mockClient.ListLayers(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListProvisionedConcurrencyConfigs", func(t *testing.T) {
        input := &lambda.ListProvisionedConcurrencyConfigsInput{}
        output := &lambda.ListProvisionedConcurrencyConfigsOutput{}

        mockClient.On("ListProvisionedConcurrencyConfigs", ctx, input).Return(output, nil)

        result, err := mockClient.ListProvisionedConcurrencyConfigs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTags", func(t *testing.T) {
        input := &lambda.ListTagsInput{}
        output := &lambda.ListTagsOutput{}

        mockClient.On("ListTags", ctx, input).Return(output, nil)

        result, err := mockClient.ListTags(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListVersionsByFunction", func(t *testing.T) {
        input := &lambda.ListVersionsByFunctionInput{}
        output := &lambda.ListVersionsByFunctionOutput{}

        mockClient.On("ListVersionsByFunction", ctx, input).Return(output, nil)

        result, err := mockClient.ListVersionsByFunction(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPublishLayerVersion", func(t *testing.T) {
        input := &lambda.PublishLayerVersionInput{}
        output := &lambda.PublishLayerVersionOutput{}

        mockClient.On("PublishLayerVersion", ctx, input).Return(output, nil)

        result, err := mockClient.PublishLayerVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPublishVersion", func(t *testing.T) {
        input := &lambda.PublishVersionInput{}
        output := &lambda.PublishVersionOutput{}

        mockClient.On("PublishVersion", ctx, input).Return(output, nil)

        result, err := mockClient.PublishVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutFunctionCodeSigningConfig", func(t *testing.T) {
        input := &lambda.PutFunctionCodeSigningConfigInput{}
        output := &lambda.PutFunctionCodeSigningConfigOutput{}

        mockClient.On("PutFunctionCodeSigningConfig", ctx, input).Return(output, nil)

        result, err := mockClient.PutFunctionCodeSigningConfig(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutFunctionConcurrency", func(t *testing.T) {
        input := &lambda.PutFunctionConcurrencyInput{}
        output := &lambda.PutFunctionConcurrencyOutput{}

        mockClient.On("PutFunctionConcurrency", ctx, input).Return(output, nil)

        result, err := mockClient.PutFunctionConcurrency(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutFunctionEventInvokeConfig", func(t *testing.T) {
        input := &lambda.PutFunctionEventInvokeConfigInput{}
        output := &lambda.PutFunctionEventInvokeConfigOutput{}

        mockClient.On("PutFunctionEventInvokeConfig", ctx, input).Return(output, nil)

        result, err := mockClient.PutFunctionEventInvokeConfig(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutProvisionedConcurrencyConfig", func(t *testing.T) {
        input := &lambda.PutProvisionedConcurrencyConfigInput{}
        output := &lambda.PutProvisionedConcurrencyConfigOutput{}

        mockClient.On("PutProvisionedConcurrencyConfig", ctx, input).Return(output, nil)

        result, err := mockClient.PutProvisionedConcurrencyConfig(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutRuntimeManagementConfig", func(t *testing.T) {
        input := &lambda.PutRuntimeManagementConfigInput{}
        output := &lambda.PutRuntimeManagementConfigOutput{}

        mockClient.On("PutRuntimeManagementConfig", ctx, input).Return(output, nil)

        result, err := mockClient.PutRuntimeManagementConfig(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRemoveLayerVersionPermission", func(t *testing.T) {
        input := &lambda.RemoveLayerVersionPermissionInput{}
        output := &lambda.RemoveLayerVersionPermissionOutput{}

        mockClient.On("RemoveLayerVersionPermission", ctx, input).Return(output, nil)

        result, err := mockClient.RemoveLayerVersionPermission(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRemovePermission", func(t *testing.T) {
        input := &lambda.RemovePermissionInput{}
        output := &lambda.RemovePermissionOutput{}

        mockClient.On("RemovePermission", ctx, input).Return(output, nil)

        result, err := mockClient.RemovePermission(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &lambda.TagResourceInput{}
        output := &lambda.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &lambda.UntagResourceInput{}
        output := &lambda.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateAlias", func(t *testing.T) {
        input := &lambda.UpdateAliasInput{}
        output := &lambda.UpdateAliasOutput{}

        mockClient.On("UpdateAlias", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateAlias(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateCodeSigningConfig", func(t *testing.T) {
        input := &lambda.UpdateCodeSigningConfigInput{}
        output := &lambda.UpdateCodeSigningConfigOutput{}

        mockClient.On("UpdateCodeSigningConfig", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateCodeSigningConfig(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateEventSourceMapping", func(t *testing.T) {
        input := &lambda.UpdateEventSourceMappingInput{}
        output := &lambda.UpdateEventSourceMappingOutput{}

        mockClient.On("UpdateEventSourceMapping", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateEventSourceMapping(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateFunctionCode", func(t *testing.T) {
        input := &lambda.UpdateFunctionCodeInput{}
        output := &lambda.UpdateFunctionCodeOutput{}

        mockClient.On("UpdateFunctionCode", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateFunctionCode(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateFunctionConfiguration", func(t *testing.T) {
        input := &lambda.UpdateFunctionConfigurationInput{}
        output := &lambda.UpdateFunctionConfigurationOutput{}

        mockClient.On("UpdateFunctionConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateFunctionConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateFunctionEventInvokeConfig", func(t *testing.T) {
        input := &lambda.UpdateFunctionEventInvokeConfigInput{}
        output := &lambda.UpdateFunctionEventInvokeConfigOutput{}

        mockClient.On("UpdateFunctionEventInvokeConfig", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateFunctionEventInvokeConfig(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateFunctionUrlConfig", func(t *testing.T) {
        input := &lambda.UpdateFunctionUrlConfigInput{}
        output := &lambda.UpdateFunctionUrlConfigOutput{}

        mockClient.On("UpdateFunctionUrlConfig", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateFunctionUrlConfig(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
