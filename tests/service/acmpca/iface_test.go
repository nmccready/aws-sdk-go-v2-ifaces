package acmpca_test

// tests for the acmpca service interface for this ../../../service/acmpca/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/acmpca"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/acmpca/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/acmpca/acmpca_iface"
	"github.com/stretchr/testify/assert"
)

func TestAcmpcaServiceCanBeMocked(t *testing.T) {
	var iface acmpca_iface.IClient
	iface = &acmpca.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := acmpca.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateCertificateAuthority", func(t *testing.T) {
        input := &acmpca.CreateCertificateAuthorityInput{}
        output := &acmpca.CreateCertificateAuthorityOutput{}

        mockClient.On("CreateCertificateAuthority", ctx, input).Return(output, nil)

        result, err := mockClient.CreateCertificateAuthority(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateCertificateAuthorityAuditReport", func(t *testing.T) {
        input := &acmpca.CreateCertificateAuthorityAuditReportInput{}
        output := &acmpca.CreateCertificateAuthorityAuditReportOutput{}

        mockClient.On("CreateCertificateAuthorityAuditReport", ctx, input).Return(output, nil)

        result, err := mockClient.CreateCertificateAuthorityAuditReport(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreatePermission", func(t *testing.T) {
        input := &acmpca.CreatePermissionInput{}
        output := &acmpca.CreatePermissionOutput{}

        mockClient.On("CreatePermission", ctx, input).Return(output, nil)

        result, err := mockClient.CreatePermission(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteCertificateAuthority", func(t *testing.T) {
        input := &acmpca.DeleteCertificateAuthorityInput{}
        output := &acmpca.DeleteCertificateAuthorityOutput{}

        mockClient.On("DeleteCertificateAuthority", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteCertificateAuthority(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeletePermission", func(t *testing.T) {
        input := &acmpca.DeletePermissionInput{}
        output := &acmpca.DeletePermissionOutput{}

        mockClient.On("DeletePermission", ctx, input).Return(output, nil)

        result, err := mockClient.DeletePermission(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeletePolicy", func(t *testing.T) {
        input := &acmpca.DeletePolicyInput{}
        output := &acmpca.DeletePolicyOutput{}

        mockClient.On("DeletePolicy", ctx, input).Return(output, nil)

        result, err := mockClient.DeletePolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeCertificateAuthority", func(t *testing.T) {
        input := &acmpca.DescribeCertificateAuthorityInput{}
        output := &acmpca.DescribeCertificateAuthorityOutput{}

        mockClient.On("DescribeCertificateAuthority", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeCertificateAuthority(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeCertificateAuthorityAuditReport", func(t *testing.T) {
        input := &acmpca.DescribeCertificateAuthorityAuditReportInput{}
        output := &acmpca.DescribeCertificateAuthorityAuditReportOutput{}

        mockClient.On("DescribeCertificateAuthorityAuditReport", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeCertificateAuthorityAuditReport(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetCertificate", func(t *testing.T) {
        input := &acmpca.GetCertificateInput{}
        output := &acmpca.GetCertificateOutput{}

        mockClient.On("GetCertificate", ctx, input).Return(output, nil)

        result, err := mockClient.GetCertificate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetCertificateAuthorityCertificate", func(t *testing.T) {
        input := &acmpca.GetCertificateAuthorityCertificateInput{}
        output := &acmpca.GetCertificateAuthorityCertificateOutput{}

        mockClient.On("GetCertificateAuthorityCertificate", ctx, input).Return(output, nil)

        result, err := mockClient.GetCertificateAuthorityCertificate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetCertificateAuthorityCsr", func(t *testing.T) {
        input := &acmpca.GetCertificateAuthorityCsrInput{}
        output := &acmpca.GetCertificateAuthorityCsrOutput{}

        mockClient.On("GetCertificateAuthorityCsr", ctx, input).Return(output, nil)

        result, err := mockClient.GetCertificateAuthorityCsr(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetPolicy", func(t *testing.T) {
        input := &acmpca.GetPolicyInput{}
        output := &acmpca.GetPolicyOutput{}

        mockClient.On("GetPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.GetPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestImportCertificateAuthorityCertificate", func(t *testing.T) {
        input := &acmpca.ImportCertificateAuthorityCertificateInput{}
        output := &acmpca.ImportCertificateAuthorityCertificateOutput{}

        mockClient.On("ImportCertificateAuthorityCertificate", ctx, input).Return(output, nil)

        result, err := mockClient.ImportCertificateAuthorityCertificate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestIssueCertificate", func(t *testing.T) {
        input := &acmpca.IssueCertificateInput{}
        output := &acmpca.IssueCertificateOutput{}

        mockClient.On("IssueCertificate", ctx, input).Return(output, nil)

        result, err := mockClient.IssueCertificate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListCertificateAuthorities", func(t *testing.T) {
        input := &acmpca.ListCertificateAuthoritiesInput{}
        output := &acmpca.ListCertificateAuthoritiesOutput{}

        mockClient.On("ListCertificateAuthorities", ctx, input).Return(output, nil)

        result, err := mockClient.ListCertificateAuthorities(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListPermissions", func(t *testing.T) {
        input := &acmpca.ListPermissionsInput{}
        output := &acmpca.ListPermissionsOutput{}

        mockClient.On("ListPermissions", ctx, input).Return(output, nil)

        result, err := mockClient.ListPermissions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTags", func(t *testing.T) {
        input := &acmpca.ListTagsInput{}
        output := &acmpca.ListTagsOutput{}

        mockClient.On("ListTags", ctx, input).Return(output, nil)

        result, err := mockClient.ListTags(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutPolicy", func(t *testing.T) {
        input := &acmpca.PutPolicyInput{}
        output := &acmpca.PutPolicyOutput{}

        mockClient.On("PutPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.PutPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRestoreCertificateAuthority", func(t *testing.T) {
        input := &acmpca.RestoreCertificateAuthorityInput{}
        output := &acmpca.RestoreCertificateAuthorityOutput{}

        mockClient.On("RestoreCertificateAuthority", ctx, input).Return(output, nil)

        result, err := mockClient.RestoreCertificateAuthority(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRevokeCertificate", func(t *testing.T) {
        input := &acmpca.RevokeCertificateInput{}
        output := &acmpca.RevokeCertificateOutput{}

        mockClient.On("RevokeCertificate", ctx, input).Return(output, nil)

        result, err := mockClient.RevokeCertificate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagCertificateAuthority", func(t *testing.T) {
        input := &acmpca.TagCertificateAuthorityInput{}
        output := &acmpca.TagCertificateAuthorityOutput{}

        mockClient.On("TagCertificateAuthority", ctx, input).Return(output, nil)

        result, err := mockClient.TagCertificateAuthority(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagCertificateAuthority", func(t *testing.T) {
        input := &acmpca.UntagCertificateAuthorityInput{}
        output := &acmpca.UntagCertificateAuthorityOutput{}

        mockClient.On("UntagCertificateAuthority", ctx, input).Return(output, nil)

        result, err := mockClient.UntagCertificateAuthority(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateCertificateAuthority", func(t *testing.T) {
        input := &acmpca.UpdateCertificateAuthorityInput{}
        output := &acmpca.UpdateCertificateAuthorityOutput{}

        mockClient.On("UpdateCertificateAuthority", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateCertificateAuthority(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
