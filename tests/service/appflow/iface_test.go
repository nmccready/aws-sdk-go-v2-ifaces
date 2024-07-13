package appflow_test

// tests for the appflow service interface for this ../../../service/appflow/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/appflow"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/appflow/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/appflow/appflow_iface"
	"github.com/stretchr/testify/assert"
)

func TestAppflowServiceCanBeMocked(t *testing.T) {
	var iface appflow_iface.IClient
	iface = &appflow.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := appflow.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCancelFlowExecutions", func(t *testing.T) {
        input := &appflow.CancelFlowExecutionsInput{}
        output := &appflow.CancelFlowExecutionsOutput{}

        mockClient.On("CancelFlowExecutions", ctx, input).Return(output, nil)

        result, err := mockClient.CancelFlowExecutions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateConnectorProfile", func(t *testing.T) {
        input := &appflow.CreateConnectorProfileInput{}
        output := &appflow.CreateConnectorProfileOutput{}

        mockClient.On("CreateConnectorProfile", ctx, input).Return(output, nil)

        result, err := mockClient.CreateConnectorProfile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateFlow", func(t *testing.T) {
        input := &appflow.CreateFlowInput{}
        output := &appflow.CreateFlowOutput{}

        mockClient.On("CreateFlow", ctx, input).Return(output, nil)

        result, err := mockClient.CreateFlow(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteConnectorProfile", func(t *testing.T) {
        input := &appflow.DeleteConnectorProfileInput{}
        output := &appflow.DeleteConnectorProfileOutput{}

        mockClient.On("DeleteConnectorProfile", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteConnectorProfile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteFlow", func(t *testing.T) {
        input := &appflow.DeleteFlowInput{}
        output := &appflow.DeleteFlowOutput{}

        mockClient.On("DeleteFlow", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteFlow(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeConnector", func(t *testing.T) {
        input := &appflow.DescribeConnectorInput{}
        output := &appflow.DescribeConnectorOutput{}

        mockClient.On("DescribeConnector", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeConnector(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeConnectorEntity", func(t *testing.T) {
        input := &appflow.DescribeConnectorEntityInput{}
        output := &appflow.DescribeConnectorEntityOutput{}

        mockClient.On("DescribeConnectorEntity", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeConnectorEntity(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeConnectorProfiles", func(t *testing.T) {
        input := &appflow.DescribeConnectorProfilesInput{}
        output := &appflow.DescribeConnectorProfilesOutput{}

        mockClient.On("DescribeConnectorProfiles", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeConnectorProfiles(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeConnectors", func(t *testing.T) {
        input := &appflow.DescribeConnectorsInput{}
        output := &appflow.DescribeConnectorsOutput{}

        mockClient.On("DescribeConnectors", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeConnectors(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeFlow", func(t *testing.T) {
        input := &appflow.DescribeFlowInput{}
        output := &appflow.DescribeFlowOutput{}

        mockClient.On("DescribeFlow", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeFlow(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeFlowExecutionRecords", func(t *testing.T) {
        input := &appflow.DescribeFlowExecutionRecordsInput{}
        output := &appflow.DescribeFlowExecutionRecordsOutput{}

        mockClient.On("DescribeFlowExecutionRecords", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeFlowExecutionRecords(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListConnectorEntities", func(t *testing.T) {
        input := &appflow.ListConnectorEntitiesInput{}
        output := &appflow.ListConnectorEntitiesOutput{}

        mockClient.On("ListConnectorEntities", ctx, input).Return(output, nil)

        result, err := mockClient.ListConnectorEntities(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListConnectors", func(t *testing.T) {
        input := &appflow.ListConnectorsInput{}
        output := &appflow.ListConnectorsOutput{}

        mockClient.On("ListConnectors", ctx, input).Return(output, nil)

        result, err := mockClient.ListConnectors(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListFlows", func(t *testing.T) {
        input := &appflow.ListFlowsInput{}
        output := &appflow.ListFlowsOutput{}

        mockClient.On("ListFlows", ctx, input).Return(output, nil)

        result, err := mockClient.ListFlows(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &appflow.ListTagsForResourceInput{}
        output := &appflow.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRegisterConnector", func(t *testing.T) {
        input := &appflow.RegisterConnectorInput{}
        output := &appflow.RegisterConnectorOutput{}

        mockClient.On("RegisterConnector", ctx, input).Return(output, nil)

        result, err := mockClient.RegisterConnector(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestResetConnectorMetadataCache", func(t *testing.T) {
        input := &appflow.ResetConnectorMetadataCacheInput{}
        output := &appflow.ResetConnectorMetadataCacheOutput{}

        mockClient.On("ResetConnectorMetadataCache", ctx, input).Return(output, nil)

        result, err := mockClient.ResetConnectorMetadataCache(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartFlow", func(t *testing.T) {
        input := &appflow.StartFlowInput{}
        output := &appflow.StartFlowOutput{}

        mockClient.On("StartFlow", ctx, input).Return(output, nil)

        result, err := mockClient.StartFlow(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopFlow", func(t *testing.T) {
        input := &appflow.StopFlowInput{}
        output := &appflow.StopFlowOutput{}

        mockClient.On("StopFlow", ctx, input).Return(output, nil)

        result, err := mockClient.StopFlow(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &appflow.TagResourceInput{}
        output := &appflow.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUnregisterConnector", func(t *testing.T) {
        input := &appflow.UnregisterConnectorInput{}
        output := &appflow.UnregisterConnectorOutput{}

        mockClient.On("UnregisterConnector", ctx, input).Return(output, nil)

        result, err := mockClient.UnregisterConnector(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &appflow.UntagResourceInput{}
        output := &appflow.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateConnectorProfile", func(t *testing.T) {
        input := &appflow.UpdateConnectorProfileInput{}
        output := &appflow.UpdateConnectorProfileOutput{}

        mockClient.On("UpdateConnectorProfile", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateConnectorProfile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateConnectorRegistration", func(t *testing.T) {
        input := &appflow.UpdateConnectorRegistrationInput{}
        output := &appflow.UpdateConnectorRegistrationOutput{}

        mockClient.On("UpdateConnectorRegistration", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateConnectorRegistration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateFlow", func(t *testing.T) {
        input := &appflow.UpdateFlowInput{}
        output := &appflow.UpdateFlowOutput{}

        mockClient.On("UpdateFlow", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateFlow(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
