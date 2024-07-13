package appsync_test

// tests for the appsync service interface for this ../../../service/appsync/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/appsync"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/appsync/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/appsync/appsync_iface"
	"github.com/stretchr/testify/assert"
)

func TestAppsyncServiceCanBeMocked(t *testing.T) {
	var iface appsync_iface.IClient
	iface = &appsync.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := appsync.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateApi", func(t *testing.T) {
        input := &appsync.AssociateApiInput{}
        output := &appsync.AssociateApiOutput{}

        mockClient.On("AssociateApi", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateApi(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateMergedGraphqlApi", func(t *testing.T) {
        input := &appsync.AssociateMergedGraphqlApiInput{}
        output := &appsync.AssociateMergedGraphqlApiOutput{}

        mockClient.On("AssociateMergedGraphqlApi", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateMergedGraphqlApi(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateSourceGraphqlApi", func(t *testing.T) {
        input := &appsync.AssociateSourceGraphqlApiInput{}
        output := &appsync.AssociateSourceGraphqlApiOutput{}

        mockClient.On("AssociateSourceGraphqlApi", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateSourceGraphqlApi(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateApiCache", func(t *testing.T) {
        input := &appsync.CreateApiCacheInput{}
        output := &appsync.CreateApiCacheOutput{}

        mockClient.On("CreateApiCache", ctx, input).Return(output, nil)

        result, err := mockClient.CreateApiCache(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateApiKey", func(t *testing.T) {
        input := &appsync.CreateApiKeyInput{}
        output := &appsync.CreateApiKeyOutput{}

        mockClient.On("CreateApiKey", ctx, input).Return(output, nil)

        result, err := mockClient.CreateApiKey(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateDataSource", func(t *testing.T) {
        input := &appsync.CreateDataSourceInput{}
        output := &appsync.CreateDataSourceOutput{}

        mockClient.On("CreateDataSource", ctx, input).Return(output, nil)

        result, err := mockClient.CreateDataSource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateDomainName", func(t *testing.T) {
        input := &appsync.CreateDomainNameInput{}
        output := &appsync.CreateDomainNameOutput{}

        mockClient.On("CreateDomainName", ctx, input).Return(output, nil)

        result, err := mockClient.CreateDomainName(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateFunction", func(t *testing.T) {
        input := &appsync.CreateFunctionInput{}
        output := &appsync.CreateFunctionOutput{}

        mockClient.On("CreateFunction", ctx, input).Return(output, nil)

        result, err := mockClient.CreateFunction(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateGraphqlApi", func(t *testing.T) {
        input := &appsync.CreateGraphqlApiInput{}
        output := &appsync.CreateGraphqlApiOutput{}

        mockClient.On("CreateGraphqlApi", ctx, input).Return(output, nil)

        result, err := mockClient.CreateGraphqlApi(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateResolver", func(t *testing.T) {
        input := &appsync.CreateResolverInput{}
        output := &appsync.CreateResolverOutput{}

        mockClient.On("CreateResolver", ctx, input).Return(output, nil)

        result, err := mockClient.CreateResolver(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateType", func(t *testing.T) {
        input := &appsync.CreateTypeInput{}
        output := &appsync.CreateTypeOutput{}

        mockClient.On("CreateType", ctx, input).Return(output, nil)

        result, err := mockClient.CreateType(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteApiCache", func(t *testing.T) {
        input := &appsync.DeleteApiCacheInput{}
        output := &appsync.DeleteApiCacheOutput{}

        mockClient.On("DeleteApiCache", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteApiCache(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteApiKey", func(t *testing.T) {
        input := &appsync.DeleteApiKeyInput{}
        output := &appsync.DeleteApiKeyOutput{}

        mockClient.On("DeleteApiKey", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteApiKey(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteDataSource", func(t *testing.T) {
        input := &appsync.DeleteDataSourceInput{}
        output := &appsync.DeleteDataSourceOutput{}

        mockClient.On("DeleteDataSource", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteDataSource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteDomainName", func(t *testing.T) {
        input := &appsync.DeleteDomainNameInput{}
        output := &appsync.DeleteDomainNameOutput{}

        mockClient.On("DeleteDomainName", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteDomainName(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteFunction", func(t *testing.T) {
        input := &appsync.DeleteFunctionInput{}
        output := &appsync.DeleteFunctionOutput{}

        mockClient.On("DeleteFunction", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteFunction(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteGraphqlApi", func(t *testing.T) {
        input := &appsync.DeleteGraphqlApiInput{}
        output := &appsync.DeleteGraphqlApiOutput{}

        mockClient.On("DeleteGraphqlApi", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteGraphqlApi(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteResolver", func(t *testing.T) {
        input := &appsync.DeleteResolverInput{}
        output := &appsync.DeleteResolverOutput{}

        mockClient.On("DeleteResolver", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteResolver(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteType", func(t *testing.T) {
        input := &appsync.DeleteTypeInput{}
        output := &appsync.DeleteTypeOutput{}

        mockClient.On("DeleteType", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteType(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateApi", func(t *testing.T) {
        input := &appsync.DisassociateApiInput{}
        output := &appsync.DisassociateApiOutput{}

        mockClient.On("DisassociateApi", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateApi(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateMergedGraphqlApi", func(t *testing.T) {
        input := &appsync.DisassociateMergedGraphqlApiInput{}
        output := &appsync.DisassociateMergedGraphqlApiOutput{}

        mockClient.On("DisassociateMergedGraphqlApi", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateMergedGraphqlApi(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateSourceGraphqlApi", func(t *testing.T) {
        input := &appsync.DisassociateSourceGraphqlApiInput{}
        output := &appsync.DisassociateSourceGraphqlApiOutput{}

        mockClient.On("DisassociateSourceGraphqlApi", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateSourceGraphqlApi(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestEvaluateCode", func(t *testing.T) {
        input := &appsync.EvaluateCodeInput{}
        output := &appsync.EvaluateCodeOutput{}

        mockClient.On("EvaluateCode", ctx, input).Return(output, nil)

        result, err := mockClient.EvaluateCode(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestEvaluateMappingTemplate", func(t *testing.T) {
        input := &appsync.EvaluateMappingTemplateInput{}
        output := &appsync.EvaluateMappingTemplateOutput{}

        mockClient.On("EvaluateMappingTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.EvaluateMappingTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestFlushApiCache", func(t *testing.T) {
        input := &appsync.FlushApiCacheInput{}
        output := &appsync.FlushApiCacheOutput{}

        mockClient.On("FlushApiCache", ctx, input).Return(output, nil)

        result, err := mockClient.FlushApiCache(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetApiAssociation", func(t *testing.T) {
        input := &appsync.GetApiAssociationInput{}
        output := &appsync.GetApiAssociationOutput{}

        mockClient.On("GetApiAssociation", ctx, input).Return(output, nil)

        result, err := mockClient.GetApiAssociation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetApiCache", func(t *testing.T) {
        input := &appsync.GetApiCacheInput{}
        output := &appsync.GetApiCacheOutput{}

        mockClient.On("GetApiCache", ctx, input).Return(output, nil)

        result, err := mockClient.GetApiCache(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDataSource", func(t *testing.T) {
        input := &appsync.GetDataSourceInput{}
        output := &appsync.GetDataSourceOutput{}

        mockClient.On("GetDataSource", ctx, input).Return(output, nil)

        result, err := mockClient.GetDataSource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDataSourceIntrospection", func(t *testing.T) {
        input := &appsync.GetDataSourceIntrospectionInput{}
        output := &appsync.GetDataSourceIntrospectionOutput{}

        mockClient.On("GetDataSourceIntrospection", ctx, input).Return(output, nil)

        result, err := mockClient.GetDataSourceIntrospection(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDomainName", func(t *testing.T) {
        input := &appsync.GetDomainNameInput{}
        output := &appsync.GetDomainNameOutput{}

        mockClient.On("GetDomainName", ctx, input).Return(output, nil)

        result, err := mockClient.GetDomainName(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetFunction", func(t *testing.T) {
        input := &appsync.GetFunctionInput{}
        output := &appsync.GetFunctionOutput{}

        mockClient.On("GetFunction", ctx, input).Return(output, nil)

        result, err := mockClient.GetFunction(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetGraphqlApi", func(t *testing.T) {
        input := &appsync.GetGraphqlApiInput{}
        output := &appsync.GetGraphqlApiOutput{}

        mockClient.On("GetGraphqlApi", ctx, input).Return(output, nil)

        result, err := mockClient.GetGraphqlApi(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetGraphqlApiEnvironmentVariables", func(t *testing.T) {
        input := &appsync.GetGraphqlApiEnvironmentVariablesInput{}
        output := &appsync.GetGraphqlApiEnvironmentVariablesOutput{}

        mockClient.On("GetGraphqlApiEnvironmentVariables", ctx, input).Return(output, nil)

        result, err := mockClient.GetGraphqlApiEnvironmentVariables(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetIntrospectionSchema", func(t *testing.T) {
        input := &appsync.GetIntrospectionSchemaInput{}
        output := &appsync.GetIntrospectionSchemaOutput{}

        mockClient.On("GetIntrospectionSchema", ctx, input).Return(output, nil)

        result, err := mockClient.GetIntrospectionSchema(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetResolver", func(t *testing.T) {
        input := &appsync.GetResolverInput{}
        output := &appsync.GetResolverOutput{}

        mockClient.On("GetResolver", ctx, input).Return(output, nil)

        result, err := mockClient.GetResolver(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSchemaCreationStatus", func(t *testing.T) {
        input := &appsync.GetSchemaCreationStatusInput{}
        output := &appsync.GetSchemaCreationStatusOutput{}

        mockClient.On("GetSchemaCreationStatus", ctx, input).Return(output, nil)

        result, err := mockClient.GetSchemaCreationStatus(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSourceApiAssociation", func(t *testing.T) {
        input := &appsync.GetSourceApiAssociationInput{}
        output := &appsync.GetSourceApiAssociationOutput{}

        mockClient.On("GetSourceApiAssociation", ctx, input).Return(output, nil)

        result, err := mockClient.GetSourceApiAssociation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetType", func(t *testing.T) {
        input := &appsync.GetTypeInput{}
        output := &appsync.GetTypeOutput{}

        mockClient.On("GetType", ctx, input).Return(output, nil)

        result, err := mockClient.GetType(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListApiKeys", func(t *testing.T) {
        input := &appsync.ListApiKeysInput{}
        output := &appsync.ListApiKeysOutput{}

        mockClient.On("ListApiKeys", ctx, input).Return(output, nil)

        result, err := mockClient.ListApiKeys(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDataSources", func(t *testing.T) {
        input := &appsync.ListDataSourcesInput{}
        output := &appsync.ListDataSourcesOutput{}

        mockClient.On("ListDataSources", ctx, input).Return(output, nil)

        result, err := mockClient.ListDataSources(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDomainNames", func(t *testing.T) {
        input := &appsync.ListDomainNamesInput{}
        output := &appsync.ListDomainNamesOutput{}

        mockClient.On("ListDomainNames", ctx, input).Return(output, nil)

        result, err := mockClient.ListDomainNames(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListFunctions", func(t *testing.T) {
        input := &appsync.ListFunctionsInput{}
        output := &appsync.ListFunctionsOutput{}

        mockClient.On("ListFunctions", ctx, input).Return(output, nil)

        result, err := mockClient.ListFunctions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListGraphqlApis", func(t *testing.T) {
        input := &appsync.ListGraphqlApisInput{}
        output := &appsync.ListGraphqlApisOutput{}

        mockClient.On("ListGraphqlApis", ctx, input).Return(output, nil)

        result, err := mockClient.ListGraphqlApis(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListResolvers", func(t *testing.T) {
        input := &appsync.ListResolversInput{}
        output := &appsync.ListResolversOutput{}

        mockClient.On("ListResolvers", ctx, input).Return(output, nil)

        result, err := mockClient.ListResolvers(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListResolversByFunction", func(t *testing.T) {
        input := &appsync.ListResolversByFunctionInput{}
        output := &appsync.ListResolversByFunctionOutput{}

        mockClient.On("ListResolversByFunction", ctx, input).Return(output, nil)

        result, err := mockClient.ListResolversByFunction(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListSourceApiAssociations", func(t *testing.T) {
        input := &appsync.ListSourceApiAssociationsInput{}
        output := &appsync.ListSourceApiAssociationsOutput{}

        mockClient.On("ListSourceApiAssociations", ctx, input).Return(output, nil)

        result, err := mockClient.ListSourceApiAssociations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &appsync.ListTagsForResourceInput{}
        output := &appsync.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTypes", func(t *testing.T) {
        input := &appsync.ListTypesInput{}
        output := &appsync.ListTypesOutput{}

        mockClient.On("ListTypes", ctx, input).Return(output, nil)

        result, err := mockClient.ListTypes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTypesByAssociation", func(t *testing.T) {
        input := &appsync.ListTypesByAssociationInput{}
        output := &appsync.ListTypesByAssociationOutput{}

        mockClient.On("ListTypesByAssociation", ctx, input).Return(output, nil)

        result, err := mockClient.ListTypesByAssociation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutGraphqlApiEnvironmentVariables", func(t *testing.T) {
        input := &appsync.PutGraphqlApiEnvironmentVariablesInput{}
        output := &appsync.PutGraphqlApiEnvironmentVariablesOutput{}

        mockClient.On("PutGraphqlApiEnvironmentVariables", ctx, input).Return(output, nil)

        result, err := mockClient.PutGraphqlApiEnvironmentVariables(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartDataSourceIntrospection", func(t *testing.T) {
        input := &appsync.StartDataSourceIntrospectionInput{}
        output := &appsync.StartDataSourceIntrospectionOutput{}

        mockClient.On("StartDataSourceIntrospection", ctx, input).Return(output, nil)

        result, err := mockClient.StartDataSourceIntrospection(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartSchemaCreation", func(t *testing.T) {
        input := &appsync.StartSchemaCreationInput{}
        output := &appsync.StartSchemaCreationOutput{}

        mockClient.On("StartSchemaCreation", ctx, input).Return(output, nil)

        result, err := mockClient.StartSchemaCreation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartSchemaMerge", func(t *testing.T) {
        input := &appsync.StartSchemaMergeInput{}
        output := &appsync.StartSchemaMergeOutput{}

        mockClient.On("StartSchemaMerge", ctx, input).Return(output, nil)

        result, err := mockClient.StartSchemaMerge(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &appsync.TagResourceInput{}
        output := &appsync.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &appsync.UntagResourceInput{}
        output := &appsync.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateApiCache", func(t *testing.T) {
        input := &appsync.UpdateApiCacheInput{}
        output := &appsync.UpdateApiCacheOutput{}

        mockClient.On("UpdateApiCache", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateApiCache(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateApiKey", func(t *testing.T) {
        input := &appsync.UpdateApiKeyInput{}
        output := &appsync.UpdateApiKeyOutput{}

        mockClient.On("UpdateApiKey", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateApiKey(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateDataSource", func(t *testing.T) {
        input := &appsync.UpdateDataSourceInput{}
        output := &appsync.UpdateDataSourceOutput{}

        mockClient.On("UpdateDataSource", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateDataSource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateDomainName", func(t *testing.T) {
        input := &appsync.UpdateDomainNameInput{}
        output := &appsync.UpdateDomainNameOutput{}

        mockClient.On("UpdateDomainName", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateDomainName(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateFunction", func(t *testing.T) {
        input := &appsync.UpdateFunctionInput{}
        output := &appsync.UpdateFunctionOutput{}

        mockClient.On("UpdateFunction", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateFunction(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateGraphqlApi", func(t *testing.T) {
        input := &appsync.UpdateGraphqlApiInput{}
        output := &appsync.UpdateGraphqlApiOutput{}

        mockClient.On("UpdateGraphqlApi", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateGraphqlApi(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateResolver", func(t *testing.T) {
        input := &appsync.UpdateResolverInput{}
        output := &appsync.UpdateResolverOutput{}

        mockClient.On("UpdateResolver", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateResolver(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateSourceApiAssociation", func(t *testing.T) {
        input := &appsync.UpdateSourceApiAssociationInput{}
        output := &appsync.UpdateSourceApiAssociationOutput{}

        mockClient.On("UpdateSourceApiAssociation", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateSourceApiAssociation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateType", func(t *testing.T) {
        input := &appsync.UpdateTypeInput{}
        output := &appsync.UpdateTypeOutput{}

        mockClient.On("UpdateType", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateType(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
