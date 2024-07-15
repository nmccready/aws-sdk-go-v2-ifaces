package paymentcryptography_test

// tests for the paymentcryptography service interface for this ../../../service/paymentcryptography/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/paymentcryptography"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/paymentcryptography/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/paymentcryptography/paymentcryptography_iface"
	"github.com/stretchr/testify/assert"
)

func TestPaymentcryptographyServiceCanBeMocked(t *testing.T) {
	var iface paymentcryptography_iface.IClient
	iface = &paymentcryptography.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := paymentcryptography.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateAlias", func(t *testing.T) {
        input := &paymentcryptography.CreateAliasInput{}
        output := &paymentcryptography.CreateAliasOutput{}

        mockClient.On("CreateAlias", ctx, input).Return(output, nil)

        result, err := mockClient.CreateAlias(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateKey", func(t *testing.T) {
        input := &paymentcryptography.CreateKeyInput{}
        output := &paymentcryptography.CreateKeyOutput{}

        mockClient.On("CreateKey", ctx, input).Return(output, nil)

        result, err := mockClient.CreateKey(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteAlias", func(t *testing.T) {
        input := &paymentcryptography.DeleteAliasInput{}
        output := &paymentcryptography.DeleteAliasOutput{}

        mockClient.On("DeleteAlias", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteAlias(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteKey", func(t *testing.T) {
        input := &paymentcryptography.DeleteKeyInput{}
        output := &paymentcryptography.DeleteKeyOutput{}

        mockClient.On("DeleteKey", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteKey(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestExportKey", func(t *testing.T) {
        input := &paymentcryptography.ExportKeyInput{}
        output := &paymentcryptography.ExportKeyOutput{}

        mockClient.On("ExportKey", ctx, input).Return(output, nil)

        result, err := mockClient.ExportKey(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetAlias", func(t *testing.T) {
        input := &paymentcryptography.GetAliasInput{}
        output := &paymentcryptography.GetAliasOutput{}

        mockClient.On("GetAlias", ctx, input).Return(output, nil)

        result, err := mockClient.GetAlias(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetKey", func(t *testing.T) {
        input := &paymentcryptography.GetKeyInput{}
        output := &paymentcryptography.GetKeyOutput{}

        mockClient.On("GetKey", ctx, input).Return(output, nil)

        result, err := mockClient.GetKey(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetParametersForExport", func(t *testing.T) {
        input := &paymentcryptography.GetParametersForExportInput{}
        output := &paymentcryptography.GetParametersForExportOutput{}

        mockClient.On("GetParametersForExport", ctx, input).Return(output, nil)

        result, err := mockClient.GetParametersForExport(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetParametersForImport", func(t *testing.T) {
        input := &paymentcryptography.GetParametersForImportInput{}
        output := &paymentcryptography.GetParametersForImportOutput{}

        mockClient.On("GetParametersForImport", ctx, input).Return(output, nil)

        result, err := mockClient.GetParametersForImport(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetPublicKeyCertificate", func(t *testing.T) {
        input := &paymentcryptography.GetPublicKeyCertificateInput{}
        output := &paymentcryptography.GetPublicKeyCertificateOutput{}

        mockClient.On("GetPublicKeyCertificate", ctx, input).Return(output, nil)

        result, err := mockClient.GetPublicKeyCertificate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestImportKey", func(t *testing.T) {
        input := &paymentcryptography.ImportKeyInput{}
        output := &paymentcryptography.ImportKeyOutput{}

        mockClient.On("ImportKey", ctx, input).Return(output, nil)

        result, err := mockClient.ImportKey(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAliases", func(t *testing.T) {
        input := &paymentcryptography.ListAliasesInput{}
        output := &paymentcryptography.ListAliasesOutput{}

        mockClient.On("ListAliases", ctx, input).Return(output, nil)

        result, err := mockClient.ListAliases(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListKeys", func(t *testing.T) {
        input := &paymentcryptography.ListKeysInput{}
        output := &paymentcryptography.ListKeysOutput{}

        mockClient.On("ListKeys", ctx, input).Return(output, nil)

        result, err := mockClient.ListKeys(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &paymentcryptography.ListTagsForResourceInput{}
        output := &paymentcryptography.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRestoreKey", func(t *testing.T) {
        input := &paymentcryptography.RestoreKeyInput{}
        output := &paymentcryptography.RestoreKeyOutput{}

        mockClient.On("RestoreKey", ctx, input).Return(output, nil)

        result, err := mockClient.RestoreKey(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartKeyUsage", func(t *testing.T) {
        input := &paymentcryptography.StartKeyUsageInput{}
        output := &paymentcryptography.StartKeyUsageOutput{}

        mockClient.On("StartKeyUsage", ctx, input).Return(output, nil)

        result, err := mockClient.StartKeyUsage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopKeyUsage", func(t *testing.T) {
        input := &paymentcryptography.StopKeyUsageInput{}
        output := &paymentcryptography.StopKeyUsageOutput{}

        mockClient.On("StopKeyUsage", ctx, input).Return(output, nil)

        result, err := mockClient.StopKeyUsage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &paymentcryptography.TagResourceInput{}
        output := &paymentcryptography.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &paymentcryptography.UntagResourceInput{}
        output := &paymentcryptography.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateAlias", func(t *testing.T) {
        input := &paymentcryptography.UpdateAliasInput{}
        output := &paymentcryptography.UpdateAliasOutput{}

        mockClient.On("UpdateAlias", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateAlias(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
