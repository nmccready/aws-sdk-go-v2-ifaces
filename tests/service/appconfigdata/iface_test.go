// Code generated by ./scripts/gen_tests.sh DO NOT EDIT.
package appconfigdata_test

// tests for the appconfigdata service interface for 
// this ../../../service/appconfigdata/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/appconfigdata"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/appconfigdata/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/appconfigdata/appconfigdata_iface"
	"github.com/stretchr/testify/assert"
)

func TestAppconfigdataServiceCanBeMocked(t *testing.T) {
	var iface appconfigdata_iface.IClient
	iface = &appconfigdata.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := appconfigdata.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetLatestConfiguration", func(t *testing.T) {
        input := &appconfigdata.GetLatestConfigurationInput{}
        output := &appconfigdata.GetLatestConfigurationOutput{}

        mockClient.On("GetLatestConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.GetLatestConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartConfigurationSession", func(t *testing.T) {
        input := &appconfigdata.StartConfigurationSessionInput{}
        output := &appconfigdata.StartConfigurationSessionOutput{}

        mockClient.On("StartConfigurationSession", ctx, input).Return(output, nil)

        result, err := mockClient.StartConfigurationSession(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
