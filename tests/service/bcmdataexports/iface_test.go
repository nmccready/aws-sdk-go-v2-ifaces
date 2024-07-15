// Code generated by ./scripts/gen_tests.sh DO NOT EDIT.
package bcmdataexports_test

// tests for the bcmdataexports service interface for 
// this ../../../service/bcmdataexports/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/bcmdataexports"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/bcmdataexports/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/bcmdataexports/bcmdataexports_iface"
	"github.com/stretchr/testify/assert"
)

func TestBcmdataexportsServiceCanBeMocked(t *testing.T) {
	var iface bcmdataexports_iface.IClient
	iface = &bcmdataexports.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := bcmdataexports.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateExport", func(t *testing.T) {
        input := &bcmdataexports.CreateExportInput{}
        output := &bcmdataexports.CreateExportOutput{}

        mockClient.On("CreateExport", ctx, input).Return(output, nil)

        result, err := mockClient.CreateExport(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteExport", func(t *testing.T) {
        input := &bcmdataexports.DeleteExportInput{}
        output := &bcmdataexports.DeleteExportOutput{}

        mockClient.On("DeleteExport", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteExport(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetExecution", func(t *testing.T) {
        input := &bcmdataexports.GetExecutionInput{}
        output := &bcmdataexports.GetExecutionOutput{}

        mockClient.On("GetExecution", ctx, input).Return(output, nil)

        result, err := mockClient.GetExecution(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetExport", func(t *testing.T) {
        input := &bcmdataexports.GetExportInput{}
        output := &bcmdataexports.GetExportOutput{}

        mockClient.On("GetExport", ctx, input).Return(output, nil)

        result, err := mockClient.GetExport(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetTable", func(t *testing.T) {
        input := &bcmdataexports.GetTableInput{}
        output := &bcmdataexports.GetTableOutput{}

        mockClient.On("GetTable", ctx, input).Return(output, nil)

        result, err := mockClient.GetTable(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListExecutions", func(t *testing.T) {
        input := &bcmdataexports.ListExecutionsInput{}
        output := &bcmdataexports.ListExecutionsOutput{}

        mockClient.On("ListExecutions", ctx, input).Return(output, nil)

        result, err := mockClient.ListExecutions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListExports", func(t *testing.T) {
        input := &bcmdataexports.ListExportsInput{}
        output := &bcmdataexports.ListExportsOutput{}

        mockClient.On("ListExports", ctx, input).Return(output, nil)

        result, err := mockClient.ListExports(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTables", func(t *testing.T) {
        input := &bcmdataexports.ListTablesInput{}
        output := &bcmdataexports.ListTablesOutput{}

        mockClient.On("ListTables", ctx, input).Return(output, nil)

        result, err := mockClient.ListTables(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &bcmdataexports.ListTagsForResourceInput{}
        output := &bcmdataexports.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &bcmdataexports.TagResourceInput{}
        output := &bcmdataexports.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &bcmdataexports.UntagResourceInput{}
        output := &bcmdataexports.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateExport", func(t *testing.T) {
        input := &bcmdataexports.UpdateExportInput{}
        output := &bcmdataexports.UpdateExportOutput{}

        mockClient.On("UpdateExport", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateExport(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
