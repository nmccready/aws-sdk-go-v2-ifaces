// Code generated by ./scripts/gen_tests.sh DO NOT EDIT.
package kafkaconnect_test

// tests for the kafkaconnect service interface for 
// this ../../../service/kafkaconnect/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/kafkaconnect"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/kafkaconnect/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/kafkaconnect/kafkaconnect_iface"
	"github.com/stretchr/testify/assert"
)

func TestKafkaconnectServiceCanBeMocked(t *testing.T) {
	var iface kafkaconnect_iface.IClient
	iface = &kafkaconnect.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := kafkaconnect.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateConnector", func(t *testing.T) {
        input := &kafkaconnect.CreateConnectorInput{}
        output := &kafkaconnect.CreateConnectorOutput{}

        mockClient.On("CreateConnector", ctx, input).Return(output, nil)

        result, err := mockClient.CreateConnector(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateCustomPlugin", func(t *testing.T) {
        input := &kafkaconnect.CreateCustomPluginInput{}
        output := &kafkaconnect.CreateCustomPluginOutput{}

        mockClient.On("CreateCustomPlugin", ctx, input).Return(output, nil)

        result, err := mockClient.CreateCustomPlugin(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateWorkerConfiguration", func(t *testing.T) {
        input := &kafkaconnect.CreateWorkerConfigurationInput{}
        output := &kafkaconnect.CreateWorkerConfigurationOutput{}

        mockClient.On("CreateWorkerConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.CreateWorkerConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteConnector", func(t *testing.T) {
        input := &kafkaconnect.DeleteConnectorInput{}
        output := &kafkaconnect.DeleteConnectorOutput{}

        mockClient.On("DeleteConnector", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteConnector(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteCustomPlugin", func(t *testing.T) {
        input := &kafkaconnect.DeleteCustomPluginInput{}
        output := &kafkaconnect.DeleteCustomPluginOutput{}

        mockClient.On("DeleteCustomPlugin", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteCustomPlugin(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteWorkerConfiguration", func(t *testing.T) {
        input := &kafkaconnect.DeleteWorkerConfigurationInput{}
        output := &kafkaconnect.DeleteWorkerConfigurationOutput{}

        mockClient.On("DeleteWorkerConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteWorkerConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeConnector", func(t *testing.T) {
        input := &kafkaconnect.DescribeConnectorInput{}
        output := &kafkaconnect.DescribeConnectorOutput{}

        mockClient.On("DescribeConnector", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeConnector(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeCustomPlugin", func(t *testing.T) {
        input := &kafkaconnect.DescribeCustomPluginInput{}
        output := &kafkaconnect.DescribeCustomPluginOutput{}

        mockClient.On("DescribeCustomPlugin", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeCustomPlugin(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeWorkerConfiguration", func(t *testing.T) {
        input := &kafkaconnect.DescribeWorkerConfigurationInput{}
        output := &kafkaconnect.DescribeWorkerConfigurationOutput{}

        mockClient.On("DescribeWorkerConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeWorkerConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListConnectors", func(t *testing.T) {
        input := &kafkaconnect.ListConnectorsInput{}
        output := &kafkaconnect.ListConnectorsOutput{}

        mockClient.On("ListConnectors", ctx, input).Return(output, nil)

        result, err := mockClient.ListConnectors(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListCustomPlugins", func(t *testing.T) {
        input := &kafkaconnect.ListCustomPluginsInput{}
        output := &kafkaconnect.ListCustomPluginsOutput{}

        mockClient.On("ListCustomPlugins", ctx, input).Return(output, nil)

        result, err := mockClient.ListCustomPlugins(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &kafkaconnect.ListTagsForResourceInput{}
        output := &kafkaconnect.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListWorkerConfigurations", func(t *testing.T) {
        input := &kafkaconnect.ListWorkerConfigurationsInput{}
        output := &kafkaconnect.ListWorkerConfigurationsOutput{}

        mockClient.On("ListWorkerConfigurations", ctx, input).Return(output, nil)

        result, err := mockClient.ListWorkerConfigurations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &kafkaconnect.TagResourceInput{}
        output := &kafkaconnect.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &kafkaconnect.UntagResourceInput{}
        output := &kafkaconnect.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateConnector", func(t *testing.T) {
        input := &kafkaconnect.UpdateConnectorInput{}
        output := &kafkaconnect.UpdateConnectorOutput{}

        mockClient.On("UpdateConnector", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateConnector(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}