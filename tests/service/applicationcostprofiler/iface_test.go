package applicationcostprofiler_test

// tests for the applicationcostprofiler service interface for this ../../../service/applicationcostprofiler/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/applicationcostprofiler"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/applicationcostprofiler/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/applicationcostprofiler/applicationcostprofiler_iface"
	"github.com/stretchr/testify/assert"
)

func TestApplicationcostprofilerServiceCanBeMocked(t *testing.T) {
	var iface applicationcostprofiler_iface.IClient
	iface = &applicationcostprofiler.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := applicationcostprofiler.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteReportDefinition", func(t *testing.T) {
        input := &applicationcostprofiler.DeleteReportDefinitionInput{}
        output := &applicationcostprofiler.DeleteReportDefinitionOutput{}

        mockClient.On("DeleteReportDefinition", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteReportDefinition(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetReportDefinition", func(t *testing.T) {
        input := &applicationcostprofiler.GetReportDefinitionInput{}
        output := &applicationcostprofiler.GetReportDefinitionOutput{}

        mockClient.On("GetReportDefinition", ctx, input).Return(output, nil)

        result, err := mockClient.GetReportDefinition(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestImportApplicationUsage", func(t *testing.T) {
        input := &applicationcostprofiler.ImportApplicationUsageInput{}
        output := &applicationcostprofiler.ImportApplicationUsageOutput{}

        mockClient.On("ImportApplicationUsage", ctx, input).Return(output, nil)

        result, err := mockClient.ImportApplicationUsage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListReportDefinitions", func(t *testing.T) {
        input := &applicationcostprofiler.ListReportDefinitionsInput{}
        output := &applicationcostprofiler.ListReportDefinitionsOutput{}

        mockClient.On("ListReportDefinitions", ctx, input).Return(output, nil)

        result, err := mockClient.ListReportDefinitions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutReportDefinition", func(t *testing.T) {
        input := &applicationcostprofiler.PutReportDefinitionInput{}
        output := &applicationcostprofiler.PutReportDefinitionOutput{}

        mockClient.On("PutReportDefinition", ctx, input).Return(output, nil)

        result, err := mockClient.PutReportDefinition(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateReportDefinition", func(t *testing.T) {
        input := &applicationcostprofiler.UpdateReportDefinitionInput{}
        output := &applicationcostprofiler.UpdateReportDefinitionOutput{}

        mockClient.On("UpdateReportDefinition", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateReportDefinition(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
