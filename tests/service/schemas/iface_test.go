package schemas_test

// tests for the schemas service interface for this ../../../service/schemas/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/schemas"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/schemas/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/schemas/schemas_iface"
	"github.com/stretchr/testify/assert"
)

func TestSchemasServiceCanBeMocked(t *testing.T) {
	var iface schemas_iface.IClient
	iface = &schemas.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := schemas.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateDiscoverer", func(t *testing.T) {
        input := &schemas.CreateDiscovererInput{}
        output := &schemas.CreateDiscovererOutput{}

        mockClient.On("CreateDiscoverer", ctx, input).Return(output, nil)

        result, err := mockClient.CreateDiscoverer(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateRegistry", func(t *testing.T) {
        input := &schemas.CreateRegistryInput{}
        output := &schemas.CreateRegistryOutput{}

        mockClient.On("CreateRegistry", ctx, input).Return(output, nil)

        result, err := mockClient.CreateRegistry(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateSchema", func(t *testing.T) {
        input := &schemas.CreateSchemaInput{}
        output := &schemas.CreateSchemaOutput{}

        mockClient.On("CreateSchema", ctx, input).Return(output, nil)

        result, err := mockClient.CreateSchema(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteDiscoverer", func(t *testing.T) {
        input := &schemas.DeleteDiscovererInput{}
        output := &schemas.DeleteDiscovererOutput{}

        mockClient.On("DeleteDiscoverer", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteDiscoverer(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteRegistry", func(t *testing.T) {
        input := &schemas.DeleteRegistryInput{}
        output := &schemas.DeleteRegistryOutput{}

        mockClient.On("DeleteRegistry", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteRegistry(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteResourcePolicy", func(t *testing.T) {
        input := &schemas.DeleteResourcePolicyInput{}
        output := &schemas.DeleteResourcePolicyOutput{}

        mockClient.On("DeleteResourcePolicy", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteResourcePolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteSchema", func(t *testing.T) {
        input := &schemas.DeleteSchemaInput{}
        output := &schemas.DeleteSchemaOutput{}

        mockClient.On("DeleteSchema", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteSchema(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteSchemaVersion", func(t *testing.T) {
        input := &schemas.DeleteSchemaVersionInput{}
        output := &schemas.DeleteSchemaVersionOutput{}

        mockClient.On("DeleteSchemaVersion", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteSchemaVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeCodeBinding", func(t *testing.T) {
        input := &schemas.DescribeCodeBindingInput{}
        output := &schemas.DescribeCodeBindingOutput{}

        mockClient.On("DescribeCodeBinding", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeCodeBinding(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeDiscoverer", func(t *testing.T) {
        input := &schemas.DescribeDiscovererInput{}
        output := &schemas.DescribeDiscovererOutput{}

        mockClient.On("DescribeDiscoverer", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeDiscoverer(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeRegistry", func(t *testing.T) {
        input := &schemas.DescribeRegistryInput{}
        output := &schemas.DescribeRegistryOutput{}

        mockClient.On("DescribeRegistry", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeRegistry(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeSchema", func(t *testing.T) {
        input := &schemas.DescribeSchemaInput{}
        output := &schemas.DescribeSchemaOutput{}

        mockClient.On("DescribeSchema", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeSchema(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestExportSchema", func(t *testing.T) {
        input := &schemas.ExportSchemaInput{}
        output := &schemas.ExportSchemaOutput{}

        mockClient.On("ExportSchema", ctx, input).Return(output, nil)

        result, err := mockClient.ExportSchema(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetCodeBindingSource", func(t *testing.T) {
        input := &schemas.GetCodeBindingSourceInput{}
        output := &schemas.GetCodeBindingSourceOutput{}

        mockClient.On("GetCodeBindingSource", ctx, input).Return(output, nil)

        result, err := mockClient.GetCodeBindingSource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDiscoveredSchema", func(t *testing.T) {
        input := &schemas.GetDiscoveredSchemaInput{}
        output := &schemas.GetDiscoveredSchemaOutput{}

        mockClient.On("GetDiscoveredSchema", ctx, input).Return(output, nil)

        result, err := mockClient.GetDiscoveredSchema(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetResourcePolicy", func(t *testing.T) {
        input := &schemas.GetResourcePolicyInput{}
        output := &schemas.GetResourcePolicyOutput{}

        mockClient.On("GetResourcePolicy", ctx, input).Return(output, nil)

        result, err := mockClient.GetResourcePolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDiscoverers", func(t *testing.T) {
        input := &schemas.ListDiscoverersInput{}
        output := &schemas.ListDiscoverersOutput{}

        mockClient.On("ListDiscoverers", ctx, input).Return(output, nil)

        result, err := mockClient.ListDiscoverers(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListRegistries", func(t *testing.T) {
        input := &schemas.ListRegistriesInput{}
        output := &schemas.ListRegistriesOutput{}

        mockClient.On("ListRegistries", ctx, input).Return(output, nil)

        result, err := mockClient.ListRegistries(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListSchemaVersions", func(t *testing.T) {
        input := &schemas.ListSchemaVersionsInput{}
        output := &schemas.ListSchemaVersionsOutput{}

        mockClient.On("ListSchemaVersions", ctx, input).Return(output, nil)

        result, err := mockClient.ListSchemaVersions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListSchemas", func(t *testing.T) {
        input := &schemas.ListSchemasInput{}
        output := &schemas.ListSchemasOutput{}

        mockClient.On("ListSchemas", ctx, input).Return(output, nil)

        result, err := mockClient.ListSchemas(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &schemas.ListTagsForResourceInput{}
        output := &schemas.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutCodeBinding", func(t *testing.T) {
        input := &schemas.PutCodeBindingInput{}
        output := &schemas.PutCodeBindingOutput{}

        mockClient.On("PutCodeBinding", ctx, input).Return(output, nil)

        result, err := mockClient.PutCodeBinding(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutResourcePolicy", func(t *testing.T) {
        input := &schemas.PutResourcePolicyInput{}
        output := &schemas.PutResourcePolicyOutput{}

        mockClient.On("PutResourcePolicy", ctx, input).Return(output, nil)

        result, err := mockClient.PutResourcePolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSearchSchemas", func(t *testing.T) {
        input := &schemas.SearchSchemasInput{}
        output := &schemas.SearchSchemasOutput{}

        mockClient.On("SearchSchemas", ctx, input).Return(output, nil)

        result, err := mockClient.SearchSchemas(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartDiscoverer", func(t *testing.T) {
        input := &schemas.StartDiscovererInput{}
        output := &schemas.StartDiscovererOutput{}

        mockClient.On("StartDiscoverer", ctx, input).Return(output, nil)

        result, err := mockClient.StartDiscoverer(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopDiscoverer", func(t *testing.T) {
        input := &schemas.StopDiscovererInput{}
        output := &schemas.StopDiscovererOutput{}

        mockClient.On("StopDiscoverer", ctx, input).Return(output, nil)

        result, err := mockClient.StopDiscoverer(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &schemas.TagResourceInput{}
        output := &schemas.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &schemas.UntagResourceInput{}
        output := &schemas.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateDiscoverer", func(t *testing.T) {
        input := &schemas.UpdateDiscovererInput{}
        output := &schemas.UpdateDiscovererOutput{}

        mockClient.On("UpdateDiscoverer", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateDiscoverer(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateRegistry", func(t *testing.T) {
        input := &schemas.UpdateRegistryInput{}
        output := &schemas.UpdateRegistryOutput{}

        mockClient.On("UpdateRegistry", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateRegistry(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateSchema", func(t *testing.T) {
        input := &schemas.UpdateSchemaInput{}
        output := &schemas.UpdateSchemaOutput{}

        mockClient.On("UpdateSchema", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateSchema(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
