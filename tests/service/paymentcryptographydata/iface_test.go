package paymentcryptographydata_test

// tests for the paymentcryptographydata service interface for this ../../../service/paymentcryptographydata/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/paymentcryptographydata"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/paymentcryptographydata/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/paymentcryptographydata/paymentcryptographydata_iface"
	"github.com/stretchr/testify/assert"
)

func TestPaymentcryptographydataServiceCanBeMocked(t *testing.T) {
	var iface paymentcryptographydata_iface.IClient
	iface = &paymentcryptographydata.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := paymentcryptographydata.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDecryptData", func(t *testing.T) {
        input := &paymentcryptographydata.DecryptDataInput{}
        output := &paymentcryptographydata.DecryptDataOutput{}

        mockClient.On("DecryptData", ctx, input).Return(output, nil)

        result, err := mockClient.DecryptData(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestEncryptData", func(t *testing.T) {
        input := &paymentcryptographydata.EncryptDataInput{}
        output := &paymentcryptographydata.EncryptDataOutput{}

        mockClient.On("EncryptData", ctx, input).Return(output, nil)

        result, err := mockClient.EncryptData(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGenerateCardValidationData", func(t *testing.T) {
        input := &paymentcryptographydata.GenerateCardValidationDataInput{}
        output := &paymentcryptographydata.GenerateCardValidationDataOutput{}

        mockClient.On("GenerateCardValidationData", ctx, input).Return(output, nil)

        result, err := mockClient.GenerateCardValidationData(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGenerateMac", func(t *testing.T) {
        input := &paymentcryptographydata.GenerateMacInput{}
        output := &paymentcryptographydata.GenerateMacOutput{}

        mockClient.On("GenerateMac", ctx, input).Return(output, nil)

        result, err := mockClient.GenerateMac(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGeneratePinData", func(t *testing.T) {
        input := &paymentcryptographydata.GeneratePinDataInput{}
        output := &paymentcryptographydata.GeneratePinDataOutput{}

        mockClient.On("GeneratePinData", ctx, input).Return(output, nil)

        result, err := mockClient.GeneratePinData(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestReEncryptData", func(t *testing.T) {
        input := &paymentcryptographydata.ReEncryptDataInput{}
        output := &paymentcryptographydata.ReEncryptDataOutput{}

        mockClient.On("ReEncryptData", ctx, input).Return(output, nil)

        result, err := mockClient.ReEncryptData(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTranslatePinData", func(t *testing.T) {
        input := &paymentcryptographydata.TranslatePinDataInput{}
        output := &paymentcryptographydata.TranslatePinDataOutput{}

        mockClient.On("TranslatePinData", ctx, input).Return(output, nil)

        result, err := mockClient.TranslatePinData(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestVerifyAuthRequestCryptogram", func(t *testing.T) {
        input := &paymentcryptographydata.VerifyAuthRequestCryptogramInput{}
        output := &paymentcryptographydata.VerifyAuthRequestCryptogramOutput{}

        mockClient.On("VerifyAuthRequestCryptogram", ctx, input).Return(output, nil)

        result, err := mockClient.VerifyAuthRequestCryptogram(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestVerifyCardValidationData", func(t *testing.T) {
        input := &paymentcryptographydata.VerifyCardValidationDataInput{}
        output := &paymentcryptographydata.VerifyCardValidationDataOutput{}

        mockClient.On("VerifyCardValidationData", ctx, input).Return(output, nil)

        result, err := mockClient.VerifyCardValidationData(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestVerifyMac", func(t *testing.T) {
        input := &paymentcryptographydata.VerifyMacInput{}
        output := &paymentcryptographydata.VerifyMacOutput{}

        mockClient.On("VerifyMac", ctx, input).Return(output, nil)

        result, err := mockClient.VerifyMac(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestVerifyPinData", func(t *testing.T) {
        input := &paymentcryptographydata.VerifyPinDataInput{}
        output := &paymentcryptographydata.VerifyPinDataOutput{}

        mockClient.On("VerifyPinData", ctx, input).Return(output, nil)

        result, err := mockClient.VerifyPinData(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
