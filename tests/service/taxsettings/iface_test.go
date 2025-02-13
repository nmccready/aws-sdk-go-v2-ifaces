// Code generated by ./scripts/gen_tests.sh DO NOT EDIT.
package taxsettings_test

// tests for the taxsettings service interface for 
// this ../../../service/taxsettings/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/taxsettings"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/taxsettings/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/taxsettings/taxsettings_iface"
	"github.com/stretchr/testify/assert"
)

func TestTaxsettingsServiceCanBeMocked(t *testing.T) {
	var iface taxsettings_iface.IClient
	iface = &taxsettings.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := taxsettings.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchDeleteTaxRegistration", func(t *testing.T) {
        input := &taxsettings.BatchDeleteTaxRegistrationInput{}
        output := &taxsettings.BatchDeleteTaxRegistrationOutput{}

        mockClient.On("BatchDeleteTaxRegistration", ctx, input).Return(output, nil)

        result, err := mockClient.BatchDeleteTaxRegistration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchGetTaxExemptions", func(t *testing.T) {
        input := &taxsettings.BatchGetTaxExemptionsInput{}
        output := &taxsettings.BatchGetTaxExemptionsOutput{}

        mockClient.On("BatchGetTaxExemptions", ctx, input).Return(output, nil)

        result, err := mockClient.BatchGetTaxExemptions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchPutTaxRegistration", func(t *testing.T) {
        input := &taxsettings.BatchPutTaxRegistrationInput{}
        output := &taxsettings.BatchPutTaxRegistrationOutput{}

        mockClient.On("BatchPutTaxRegistration", ctx, input).Return(output, nil)

        result, err := mockClient.BatchPutTaxRegistration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteSupplementalTaxRegistration", func(t *testing.T) {
        input := &taxsettings.DeleteSupplementalTaxRegistrationInput{}
        output := &taxsettings.DeleteSupplementalTaxRegistrationOutput{}

        mockClient.On("DeleteSupplementalTaxRegistration", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteSupplementalTaxRegistration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteTaxRegistration", func(t *testing.T) {
        input := &taxsettings.DeleteTaxRegistrationInput{}
        output := &taxsettings.DeleteTaxRegistrationOutput{}

        mockClient.On("DeleteTaxRegistration", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteTaxRegistration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetTaxExemptionTypes", func(t *testing.T) {
        input := &taxsettings.GetTaxExemptionTypesInput{}
        output := &taxsettings.GetTaxExemptionTypesOutput{}

        mockClient.On("GetTaxExemptionTypes", ctx, input).Return(output, nil)

        result, err := mockClient.GetTaxExemptionTypes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetTaxInheritance", func(t *testing.T) {
        input := &taxsettings.GetTaxInheritanceInput{}
        output := &taxsettings.GetTaxInheritanceOutput{}

        mockClient.On("GetTaxInheritance", ctx, input).Return(output, nil)

        result, err := mockClient.GetTaxInheritance(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetTaxRegistration", func(t *testing.T) {
        input := &taxsettings.GetTaxRegistrationInput{}
        output := &taxsettings.GetTaxRegistrationOutput{}

        mockClient.On("GetTaxRegistration", ctx, input).Return(output, nil)

        result, err := mockClient.GetTaxRegistration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetTaxRegistrationDocument", func(t *testing.T) {
        input := &taxsettings.GetTaxRegistrationDocumentInput{}
        output := &taxsettings.GetTaxRegistrationDocumentOutput{}

        mockClient.On("GetTaxRegistrationDocument", ctx, input).Return(output, nil)

        result, err := mockClient.GetTaxRegistrationDocument(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListSupplementalTaxRegistrations", func(t *testing.T) {
        input := &taxsettings.ListSupplementalTaxRegistrationsInput{}
        output := &taxsettings.ListSupplementalTaxRegistrationsOutput{}

        mockClient.On("ListSupplementalTaxRegistrations", ctx, input).Return(output, nil)

        result, err := mockClient.ListSupplementalTaxRegistrations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTaxExemptions", func(t *testing.T) {
        input := &taxsettings.ListTaxExemptionsInput{}
        output := &taxsettings.ListTaxExemptionsOutput{}

        mockClient.On("ListTaxExemptions", ctx, input).Return(output, nil)

        result, err := mockClient.ListTaxExemptions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTaxRegistrations", func(t *testing.T) {
        input := &taxsettings.ListTaxRegistrationsInput{}
        output := &taxsettings.ListTaxRegistrationsOutput{}

        mockClient.On("ListTaxRegistrations", ctx, input).Return(output, nil)

        result, err := mockClient.ListTaxRegistrations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutSupplementalTaxRegistration", func(t *testing.T) {
        input := &taxsettings.PutSupplementalTaxRegistrationInput{}
        output := &taxsettings.PutSupplementalTaxRegistrationOutput{}

        mockClient.On("PutSupplementalTaxRegistration", ctx, input).Return(output, nil)

        result, err := mockClient.PutSupplementalTaxRegistration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutTaxExemption", func(t *testing.T) {
        input := &taxsettings.PutTaxExemptionInput{}
        output := &taxsettings.PutTaxExemptionOutput{}

        mockClient.On("PutTaxExemption", ctx, input).Return(output, nil)

        result, err := mockClient.PutTaxExemption(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutTaxInheritance", func(t *testing.T) {
        input := &taxsettings.PutTaxInheritanceInput{}
        output := &taxsettings.PutTaxInheritanceOutput{}

        mockClient.On("PutTaxInheritance", ctx, input).Return(output, nil)

        result, err := mockClient.PutTaxInheritance(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutTaxRegistration", func(t *testing.T) {
        input := &taxsettings.PutTaxRegistrationInput{}
        output := &taxsettings.PutTaxRegistrationOutput{}

        mockClient.On("PutTaxRegistration", ctx, input).Return(output, nil)

        result, err := mockClient.PutTaxRegistration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
