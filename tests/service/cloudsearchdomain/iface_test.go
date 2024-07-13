package cloudsearchdomain_test

// tests for the cloudsearchdomain service interface for this ../../../service/cloudsearchdomain/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/cloudsearchdomain"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/cloudsearchdomain/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/cloudsearchdomain/cloudsearchdomain_iface"
	"github.com/stretchr/testify/assert"
)

func TestCloudsearchdomainServiceCanBeMocked(t *testing.T) {
	var iface cloudsearchdomain_iface.IClient
	iface = &cloudsearchdomain.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := cloudsearchdomain.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSearch", func(t *testing.T) {
        input := &cloudsearchdomain.SearchInput{}
        output := &cloudsearchdomain.SearchOutput{}

        mockClient.On("Search", ctx, input).Return(output, nil)

        result, err := mockClient.Search(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSuggest", func(t *testing.T) {
        input := &cloudsearchdomain.SuggestInput{}
        output := &cloudsearchdomain.SuggestOutput{}

        mockClient.On("Suggest", ctx, input).Return(output, nil)

        result, err := mockClient.Suggest(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUploadDocuments", func(t *testing.T) {
        input := &cloudsearchdomain.UploadDocumentsInput{}
        output := &cloudsearchdomain.UploadDocumentsOutput{}

        mockClient.On("UploadDocuments", ctx, input).Return(output, nil)

        result, err := mockClient.UploadDocuments(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
