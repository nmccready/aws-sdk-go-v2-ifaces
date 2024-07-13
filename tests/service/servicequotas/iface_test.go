package servicequotas_test

// tests for the servicequotas service interface for this ../../../service/servicequotas/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/servicequotas"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/servicequotas/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/servicequotas/servicequotas_iface"
	"github.com/stretchr/testify/assert"
)

func TestServicequotasServiceCanBeMocked(t *testing.T) {
	var iface servicequotas_iface.IClient
	iface = &servicequotas.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := servicequotas.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateServiceQuotaTemplate", func(t *testing.T) {
        input := &servicequotas.AssociateServiceQuotaTemplateInput{}
        output := &servicequotas.AssociateServiceQuotaTemplateOutput{}

        mockClient.On("AssociateServiceQuotaTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateServiceQuotaTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteServiceQuotaIncreaseRequestFromTemplate", func(t *testing.T) {
        input := &servicequotas.DeleteServiceQuotaIncreaseRequestFromTemplateInput{}
        output := &servicequotas.DeleteServiceQuotaIncreaseRequestFromTemplateOutput{}

        mockClient.On("DeleteServiceQuotaIncreaseRequestFromTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteServiceQuotaIncreaseRequestFromTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateServiceQuotaTemplate", func(t *testing.T) {
        input := &servicequotas.DisassociateServiceQuotaTemplateInput{}
        output := &servicequotas.DisassociateServiceQuotaTemplateOutput{}

        mockClient.On("DisassociateServiceQuotaTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateServiceQuotaTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetAWSDefaultServiceQuota", func(t *testing.T) {
        input := &servicequotas.GetAWSDefaultServiceQuotaInput{}
        output := &servicequotas.GetAWSDefaultServiceQuotaOutput{}

        mockClient.On("GetAWSDefaultServiceQuota", ctx, input).Return(output, nil)

        result, err := mockClient.GetAWSDefaultServiceQuota(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetAssociationForServiceQuotaTemplate", func(t *testing.T) {
        input := &servicequotas.GetAssociationForServiceQuotaTemplateInput{}
        output := &servicequotas.GetAssociationForServiceQuotaTemplateOutput{}

        mockClient.On("GetAssociationForServiceQuotaTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.GetAssociationForServiceQuotaTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetRequestedServiceQuotaChange", func(t *testing.T) {
        input := &servicequotas.GetRequestedServiceQuotaChangeInput{}
        output := &servicequotas.GetRequestedServiceQuotaChangeOutput{}

        mockClient.On("GetRequestedServiceQuotaChange", ctx, input).Return(output, nil)

        result, err := mockClient.GetRequestedServiceQuotaChange(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetServiceQuota", func(t *testing.T) {
        input := &servicequotas.GetServiceQuotaInput{}
        output := &servicequotas.GetServiceQuotaOutput{}

        mockClient.On("GetServiceQuota", ctx, input).Return(output, nil)

        result, err := mockClient.GetServiceQuota(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetServiceQuotaIncreaseRequestFromTemplate", func(t *testing.T) {
        input := &servicequotas.GetServiceQuotaIncreaseRequestFromTemplateInput{}
        output := &servicequotas.GetServiceQuotaIncreaseRequestFromTemplateOutput{}

        mockClient.On("GetServiceQuotaIncreaseRequestFromTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.GetServiceQuotaIncreaseRequestFromTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAWSDefaultServiceQuotas", func(t *testing.T) {
        input := &servicequotas.ListAWSDefaultServiceQuotasInput{}
        output := &servicequotas.ListAWSDefaultServiceQuotasOutput{}

        mockClient.On("ListAWSDefaultServiceQuotas", ctx, input).Return(output, nil)

        result, err := mockClient.ListAWSDefaultServiceQuotas(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListRequestedServiceQuotaChangeHistory", func(t *testing.T) {
        input := &servicequotas.ListRequestedServiceQuotaChangeHistoryInput{}
        output := &servicequotas.ListRequestedServiceQuotaChangeHistoryOutput{}

        mockClient.On("ListRequestedServiceQuotaChangeHistory", ctx, input).Return(output, nil)

        result, err := mockClient.ListRequestedServiceQuotaChangeHistory(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListRequestedServiceQuotaChangeHistoryByQuota", func(t *testing.T) {
        input := &servicequotas.ListRequestedServiceQuotaChangeHistoryByQuotaInput{}
        output := &servicequotas.ListRequestedServiceQuotaChangeHistoryByQuotaOutput{}

        mockClient.On("ListRequestedServiceQuotaChangeHistoryByQuota", ctx, input).Return(output, nil)

        result, err := mockClient.ListRequestedServiceQuotaChangeHistoryByQuota(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListServiceQuotaIncreaseRequestsInTemplate", func(t *testing.T) {
        input := &servicequotas.ListServiceQuotaIncreaseRequestsInTemplateInput{}
        output := &servicequotas.ListServiceQuotaIncreaseRequestsInTemplateOutput{}

        mockClient.On("ListServiceQuotaIncreaseRequestsInTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.ListServiceQuotaIncreaseRequestsInTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListServiceQuotas", func(t *testing.T) {
        input := &servicequotas.ListServiceQuotasInput{}
        output := &servicequotas.ListServiceQuotasOutput{}

        mockClient.On("ListServiceQuotas", ctx, input).Return(output, nil)

        result, err := mockClient.ListServiceQuotas(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListServices", func(t *testing.T) {
        input := &servicequotas.ListServicesInput{}
        output := &servicequotas.ListServicesOutput{}

        mockClient.On("ListServices", ctx, input).Return(output, nil)

        result, err := mockClient.ListServices(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &servicequotas.ListTagsForResourceInput{}
        output := &servicequotas.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutServiceQuotaIncreaseRequestIntoTemplate", func(t *testing.T) {
        input := &servicequotas.PutServiceQuotaIncreaseRequestIntoTemplateInput{}
        output := &servicequotas.PutServiceQuotaIncreaseRequestIntoTemplateOutput{}

        mockClient.On("PutServiceQuotaIncreaseRequestIntoTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.PutServiceQuotaIncreaseRequestIntoTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRequestServiceQuotaIncrease", func(t *testing.T) {
        input := &servicequotas.RequestServiceQuotaIncreaseInput{}
        output := &servicequotas.RequestServiceQuotaIncreaseOutput{}

        mockClient.On("RequestServiceQuotaIncrease", ctx, input).Return(output, nil)

        result, err := mockClient.RequestServiceQuotaIncrease(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &servicequotas.TagResourceInput{}
        output := &servicequotas.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &servicequotas.UntagResourceInput{}
        output := &servicequotas.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
