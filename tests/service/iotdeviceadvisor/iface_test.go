package iotdeviceadvisor_test

// tests for the iotdeviceadvisor service interface for this ../../../service/iotdeviceadvisor/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/iotdeviceadvisor"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/iotdeviceadvisor/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/iotdeviceadvisor/iotdeviceadvisor_iface"
	"github.com/stretchr/testify/assert"
)

func TestIotdeviceadvisorServiceCanBeMocked(t *testing.T) {
	var iface iotdeviceadvisor_iface.IClient
	iface = &iotdeviceadvisor.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := iotdeviceadvisor.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateSuiteDefinition", func(t *testing.T) {
        input := &iotdeviceadvisor.CreateSuiteDefinitionInput{}
        output := &iotdeviceadvisor.CreateSuiteDefinitionOutput{}

        mockClient.On("CreateSuiteDefinition", ctx, input).Return(output, nil)

        result, err := mockClient.CreateSuiteDefinition(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteSuiteDefinition", func(t *testing.T) {
        input := &iotdeviceadvisor.DeleteSuiteDefinitionInput{}
        output := &iotdeviceadvisor.DeleteSuiteDefinitionOutput{}

        mockClient.On("DeleteSuiteDefinition", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteSuiteDefinition(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetEndpoint", func(t *testing.T) {
        input := &iotdeviceadvisor.GetEndpointInput{}
        output := &iotdeviceadvisor.GetEndpointOutput{}

        mockClient.On("GetEndpoint", ctx, input).Return(output, nil)

        result, err := mockClient.GetEndpoint(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSuiteDefinition", func(t *testing.T) {
        input := &iotdeviceadvisor.GetSuiteDefinitionInput{}
        output := &iotdeviceadvisor.GetSuiteDefinitionOutput{}

        mockClient.On("GetSuiteDefinition", ctx, input).Return(output, nil)

        result, err := mockClient.GetSuiteDefinition(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSuiteRun", func(t *testing.T) {
        input := &iotdeviceadvisor.GetSuiteRunInput{}
        output := &iotdeviceadvisor.GetSuiteRunOutput{}

        mockClient.On("GetSuiteRun", ctx, input).Return(output, nil)

        result, err := mockClient.GetSuiteRun(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSuiteRunReport", func(t *testing.T) {
        input := &iotdeviceadvisor.GetSuiteRunReportInput{}
        output := &iotdeviceadvisor.GetSuiteRunReportOutput{}

        mockClient.On("GetSuiteRunReport", ctx, input).Return(output, nil)

        result, err := mockClient.GetSuiteRunReport(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListSuiteDefinitions", func(t *testing.T) {
        input := &iotdeviceadvisor.ListSuiteDefinitionsInput{}
        output := &iotdeviceadvisor.ListSuiteDefinitionsOutput{}

        mockClient.On("ListSuiteDefinitions", ctx, input).Return(output, nil)

        result, err := mockClient.ListSuiteDefinitions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListSuiteRuns", func(t *testing.T) {
        input := &iotdeviceadvisor.ListSuiteRunsInput{}
        output := &iotdeviceadvisor.ListSuiteRunsOutput{}

        mockClient.On("ListSuiteRuns", ctx, input).Return(output, nil)

        result, err := mockClient.ListSuiteRuns(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &iotdeviceadvisor.ListTagsForResourceInput{}
        output := &iotdeviceadvisor.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartSuiteRun", func(t *testing.T) {
        input := &iotdeviceadvisor.StartSuiteRunInput{}
        output := &iotdeviceadvisor.StartSuiteRunOutput{}

        mockClient.On("StartSuiteRun", ctx, input).Return(output, nil)

        result, err := mockClient.StartSuiteRun(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopSuiteRun", func(t *testing.T) {
        input := &iotdeviceadvisor.StopSuiteRunInput{}
        output := &iotdeviceadvisor.StopSuiteRunOutput{}

        mockClient.On("StopSuiteRun", ctx, input).Return(output, nil)

        result, err := mockClient.StopSuiteRun(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &iotdeviceadvisor.TagResourceInput{}
        output := &iotdeviceadvisor.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &iotdeviceadvisor.UntagResourceInput{}
        output := &iotdeviceadvisor.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateSuiteDefinition", func(t *testing.T) {
        input := &iotdeviceadvisor.UpdateSuiteDefinitionInput{}
        output := &iotdeviceadvisor.UpdateSuiteDefinitionOutput{}

        mockClient.On("UpdateSuiteDefinition", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateSuiteDefinition(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
