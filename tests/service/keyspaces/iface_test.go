package keyspaces_test

// tests for the keyspaces service interface for this ../../../service/keyspaces/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/keyspaces"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/keyspaces/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/keyspaces/keyspaces_iface"
	"github.com/stretchr/testify/assert"
)

func TestKeyspacesServiceCanBeMocked(t *testing.T) {
	var iface keyspaces_iface.IClient
	iface = &keyspaces.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := keyspaces.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateKeyspace", func(t *testing.T) {
        input := &keyspaces.CreateKeyspaceInput{}
        output := &keyspaces.CreateKeyspaceOutput{}

        mockClient.On("CreateKeyspace", ctx, input).Return(output, nil)

        result, err := mockClient.CreateKeyspace(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateTable", func(t *testing.T) {
        input := &keyspaces.CreateTableInput{}
        output := &keyspaces.CreateTableOutput{}

        mockClient.On("CreateTable", ctx, input).Return(output, nil)

        result, err := mockClient.CreateTable(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteKeyspace", func(t *testing.T) {
        input := &keyspaces.DeleteKeyspaceInput{}
        output := &keyspaces.DeleteKeyspaceOutput{}

        mockClient.On("DeleteKeyspace", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteKeyspace(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteTable", func(t *testing.T) {
        input := &keyspaces.DeleteTableInput{}
        output := &keyspaces.DeleteTableOutput{}

        mockClient.On("DeleteTable", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteTable(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetKeyspace", func(t *testing.T) {
        input := &keyspaces.GetKeyspaceInput{}
        output := &keyspaces.GetKeyspaceOutput{}

        mockClient.On("GetKeyspace", ctx, input).Return(output, nil)

        result, err := mockClient.GetKeyspace(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetTable", func(t *testing.T) {
        input := &keyspaces.GetTableInput{}
        output := &keyspaces.GetTableOutput{}

        mockClient.On("GetTable", ctx, input).Return(output, nil)

        result, err := mockClient.GetTable(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetTableAutoScalingSettings", func(t *testing.T) {
        input := &keyspaces.GetTableAutoScalingSettingsInput{}
        output := &keyspaces.GetTableAutoScalingSettingsOutput{}

        mockClient.On("GetTableAutoScalingSettings", ctx, input).Return(output, nil)

        result, err := mockClient.GetTableAutoScalingSettings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListKeyspaces", func(t *testing.T) {
        input := &keyspaces.ListKeyspacesInput{}
        output := &keyspaces.ListKeyspacesOutput{}

        mockClient.On("ListKeyspaces", ctx, input).Return(output, nil)

        result, err := mockClient.ListKeyspaces(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTables", func(t *testing.T) {
        input := &keyspaces.ListTablesInput{}
        output := &keyspaces.ListTablesOutput{}

        mockClient.On("ListTables", ctx, input).Return(output, nil)

        result, err := mockClient.ListTables(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &keyspaces.ListTagsForResourceInput{}
        output := &keyspaces.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRestoreTable", func(t *testing.T) {
        input := &keyspaces.RestoreTableInput{}
        output := &keyspaces.RestoreTableOutput{}

        mockClient.On("RestoreTable", ctx, input).Return(output, nil)

        result, err := mockClient.RestoreTable(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &keyspaces.TagResourceInput{}
        output := &keyspaces.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &keyspaces.UntagResourceInput{}
        output := &keyspaces.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateTable", func(t *testing.T) {
        input := &keyspaces.UpdateTableInput{}
        output := &keyspaces.UpdateTableOutput{}

        mockClient.On("UpdateTable", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateTable(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
