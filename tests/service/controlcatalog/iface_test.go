package controlcatalog_test

// tests for the controlcatalog service interface for this ../../../service/controlcatalog/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/controlcatalog"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/controlcatalog/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/controlcatalog/controlcatalog_iface"
	"github.com/stretchr/testify/assert"
)

func TestControlcatalogServiceCanBeMocked(t *testing.T) {
	var iface controlcatalog_iface.IClient
	iface = &controlcatalog.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := controlcatalog.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListCommonControls", func(t *testing.T) {
        input := &controlcatalog.ListCommonControlsInput{}
        output := &controlcatalog.ListCommonControlsOutput{}

        mockClient.On("ListCommonControls", ctx, input).Return(output, nil)

        result, err := mockClient.ListCommonControls(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDomains", func(t *testing.T) {
        input := &controlcatalog.ListDomainsInput{}
        output := &controlcatalog.ListDomainsOutput{}

        mockClient.On("ListDomains", ctx, input).Return(output, nil)

        result, err := mockClient.ListDomains(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListObjectives", func(t *testing.T) {
        input := &controlcatalog.ListObjectivesInput{}
        output := &controlcatalog.ListObjectivesOutput{}

        mockClient.On("ListObjectives", ctx, input).Return(output, nil)

        result, err := mockClient.ListObjectives(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
