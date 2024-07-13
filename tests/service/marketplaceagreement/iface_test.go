package marketplaceagreement_test

// tests for the marketplaceagreement service interface for this ../../../service/marketplaceagreement/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/marketplaceagreement"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/marketplaceagreement/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/marketplaceagreement/marketplaceagreement_iface"
	"github.com/stretchr/testify/assert"
)

func TestMarketplaceagreementServiceCanBeMocked(t *testing.T) {
	var iface marketplaceagreement_iface.IClient
	iface = &marketplaceagreement.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := marketplaceagreement.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeAgreement", func(t *testing.T) {
        input := &marketplaceagreement.DescribeAgreementInput{}
        output := &marketplaceagreement.DescribeAgreementOutput{}

        mockClient.On("DescribeAgreement", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeAgreement(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetAgreementTerms", func(t *testing.T) {
        input := &marketplaceagreement.GetAgreementTermsInput{}
        output := &marketplaceagreement.GetAgreementTermsOutput{}

        mockClient.On("GetAgreementTerms", ctx, input).Return(output, nil)

        result, err := mockClient.GetAgreementTerms(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSearchAgreements", func(t *testing.T) {
        input := &marketplaceagreement.SearchAgreementsInput{}
        output := &marketplaceagreement.SearchAgreementsOutput{}

        mockClient.On("SearchAgreements", ctx, input).Return(output, nil)

        result, err := mockClient.SearchAgreements(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
