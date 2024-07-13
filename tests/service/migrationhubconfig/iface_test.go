package migrationhubconfig_test

// tests for the migrationhubconfig service interface for this ../../../service/migrationhubconfig/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/migrationhubconfig"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/migrationhubconfig/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/migrationhubconfig/migrationhubconfig_iface"
	"github.com/stretchr/testify/assert"
)

func TestMigrationhubconfigServiceCanBeMocked(t *testing.T) {
	var iface migrationhubconfig_iface.IClient
	iface = &migrationhubconfig.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := migrationhubconfig.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateHomeRegionControl", func(t *testing.T) {
        input := &migrationhubconfig.CreateHomeRegionControlInput{}
        output := &migrationhubconfig.CreateHomeRegionControlOutput{}

        mockClient.On("CreateHomeRegionControl", ctx, input).Return(output, nil)

        result, err := mockClient.CreateHomeRegionControl(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteHomeRegionControl", func(t *testing.T) {
        input := &migrationhubconfig.DeleteHomeRegionControlInput{}
        output := &migrationhubconfig.DeleteHomeRegionControlOutput{}

        mockClient.On("DeleteHomeRegionControl", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteHomeRegionControl(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeHomeRegionControls", func(t *testing.T) {
        input := &migrationhubconfig.DescribeHomeRegionControlsInput{}
        output := &migrationhubconfig.DescribeHomeRegionControlsOutput{}

        mockClient.On("DescribeHomeRegionControls", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeHomeRegionControls(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetHomeRegion", func(t *testing.T) {
        input := &migrationhubconfig.GetHomeRegionInput{}
        output := &migrationhubconfig.GetHomeRegionOutput{}

        mockClient.On("GetHomeRegion", ctx, input).Return(output, nil)

        result, err := mockClient.GetHomeRegion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
