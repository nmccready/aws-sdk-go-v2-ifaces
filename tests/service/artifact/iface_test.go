package artifact_test

// tests for the artifact service interface for this ../../../service/artifact/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/artifact"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/artifact/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/artifact/artifact_iface"
	"github.com/stretchr/testify/assert"
)

func TestArtifactServiceCanBeMocked(t *testing.T) {
	var iface artifact_iface.IClient
	iface = &artifact.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := artifact.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetAccountSettings", func(t *testing.T) {
        input := &artifact.GetAccountSettingsInput{}
        output := &artifact.GetAccountSettingsOutput{}

        mockClient.On("GetAccountSettings", ctx, input).Return(output, nil)

        result, err := mockClient.GetAccountSettings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetReport", func(t *testing.T) {
        input := &artifact.GetReportInput{}
        output := &artifact.GetReportOutput{}

        mockClient.On("GetReport", ctx, input).Return(output, nil)

        result, err := mockClient.GetReport(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetReportMetadata", func(t *testing.T) {
        input := &artifact.GetReportMetadataInput{}
        output := &artifact.GetReportMetadataOutput{}

        mockClient.On("GetReportMetadata", ctx, input).Return(output, nil)

        result, err := mockClient.GetReportMetadata(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetTermForReport", func(t *testing.T) {
        input := &artifact.GetTermForReportInput{}
        output := &artifact.GetTermForReportOutput{}

        mockClient.On("GetTermForReport", ctx, input).Return(output, nil)

        result, err := mockClient.GetTermForReport(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListReports", func(t *testing.T) {
        input := &artifact.ListReportsInput{}
        output := &artifact.ListReportsOutput{}

        mockClient.On("ListReports", ctx, input).Return(output, nil)

        result, err := mockClient.ListReports(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutAccountSettings", func(t *testing.T) {
        input := &artifact.PutAccountSettingsInput{}
        output := &artifact.PutAccountSettingsOutput{}

        mockClient.On("PutAccountSettings", ctx, input).Return(output, nil)

        result, err := mockClient.PutAccountSettings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
