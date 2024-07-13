package acm_test

// tests for the acm service interface for this ../../../service/acm/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/acm"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/acm/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/acm/acm_iface"
	"github.com/stretchr/testify/assert"
)

func TestAcmServiceCanBeMocked(t *testing.T) {
	var iface acm_iface.IClient
	iface = &acm.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := acm.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAddTagsToCertificate", func(t *testing.T) {
        input := &acm.AddTagsToCertificateInput{}
        output := &acm.AddTagsToCertificateOutput{}

        mockClient.On("AddTagsToCertificate", ctx, input).Return(output, nil)

        result, err := mockClient.AddTagsToCertificate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteCertificate", func(t *testing.T) {
        input := &acm.DeleteCertificateInput{}
        output := &acm.DeleteCertificateOutput{}

        mockClient.On("DeleteCertificate", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteCertificate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeCertificate", func(t *testing.T) {
        input := &acm.DescribeCertificateInput{}
        output := &acm.DescribeCertificateOutput{}

        mockClient.On("DescribeCertificate", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeCertificate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestExportCertificate", func(t *testing.T) {
        input := &acm.ExportCertificateInput{}
        output := &acm.ExportCertificateOutput{}

        mockClient.On("ExportCertificate", ctx, input).Return(output, nil)

        result, err := mockClient.ExportCertificate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetAccountConfiguration", func(t *testing.T) {
        input := &acm.GetAccountConfigurationInput{}
        output := &acm.GetAccountConfigurationOutput{}

        mockClient.On("GetAccountConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.GetAccountConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetCertificate", func(t *testing.T) {
        input := &acm.GetCertificateInput{}
        output := &acm.GetCertificateOutput{}

        mockClient.On("GetCertificate", ctx, input).Return(output, nil)

        result, err := mockClient.GetCertificate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestImportCertificate", func(t *testing.T) {
        input := &acm.ImportCertificateInput{}
        output := &acm.ImportCertificateOutput{}

        mockClient.On("ImportCertificate", ctx, input).Return(output, nil)

        result, err := mockClient.ImportCertificate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListCertificates", func(t *testing.T) {
        input := &acm.ListCertificatesInput{}
        output := &acm.ListCertificatesOutput{}

        mockClient.On("ListCertificates", ctx, input).Return(output, nil)

        result, err := mockClient.ListCertificates(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForCertificate", func(t *testing.T) {
        input := &acm.ListTagsForCertificateInput{}
        output := &acm.ListTagsForCertificateOutput{}

        mockClient.On("ListTagsForCertificate", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForCertificate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutAccountConfiguration", func(t *testing.T) {
        input := &acm.PutAccountConfigurationInput{}
        output := &acm.PutAccountConfigurationOutput{}

        mockClient.On("PutAccountConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.PutAccountConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRemoveTagsFromCertificate", func(t *testing.T) {
        input := &acm.RemoveTagsFromCertificateInput{}
        output := &acm.RemoveTagsFromCertificateOutput{}

        mockClient.On("RemoveTagsFromCertificate", ctx, input).Return(output, nil)

        result, err := mockClient.RemoveTagsFromCertificate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRenewCertificate", func(t *testing.T) {
        input := &acm.RenewCertificateInput{}
        output := &acm.RenewCertificateOutput{}

        mockClient.On("RenewCertificate", ctx, input).Return(output, nil)

        result, err := mockClient.RenewCertificate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRequestCertificate", func(t *testing.T) {
        input := &acm.RequestCertificateInput{}
        output := &acm.RequestCertificateOutput{}

        mockClient.On("RequestCertificate", ctx, input).Return(output, nil)

        result, err := mockClient.RequestCertificate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestResendValidationEmail", func(t *testing.T) {
        input := &acm.ResendValidationEmailInput{}
        output := &acm.ResendValidationEmailOutput{}

        mockClient.On("ResendValidationEmail", ctx, input).Return(output, nil)

        result, err := mockClient.ResendValidationEmail(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateCertificateOptions", func(t *testing.T) {
        input := &acm.UpdateCertificateOptionsInput{}
        output := &acm.UpdateCertificateOptionsOutput{}

        mockClient.On("UpdateCertificateOptions", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateCertificateOptions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
