package lexruntimeservice_test

// tests for the lexruntimeservice service interface for this ../../../service/lexruntimeservice/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/lexruntimeservice"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/lexruntimeservice/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/lexruntimeservice/lexruntimeservice_iface"
	"github.com/stretchr/testify/assert"
)

func TestLexruntimeserviceServiceCanBeMocked(t *testing.T) {
	var iface lexruntimeservice_iface.IClient
	iface = &lexruntimeservice.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := lexruntimeservice.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteSession", func(t *testing.T) {
        input := &lexruntimeservice.DeleteSessionInput{}
        output := &lexruntimeservice.DeleteSessionOutput{}

        mockClient.On("DeleteSession", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteSession(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSession", func(t *testing.T) {
        input := &lexruntimeservice.GetSessionInput{}
        output := &lexruntimeservice.GetSessionOutput{}

        mockClient.On("GetSession", ctx, input).Return(output, nil)

        result, err := mockClient.GetSession(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPostContent", func(t *testing.T) {
        input := &lexruntimeservice.PostContentInput{}
        output := &lexruntimeservice.PostContentOutput{}

        mockClient.On("PostContent", ctx, input).Return(output, nil)

        result, err := mockClient.PostContent(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPostText", func(t *testing.T) {
        input := &lexruntimeservice.PostTextInput{}
        output := &lexruntimeservice.PostTextOutput{}

        mockClient.On("PostText", ctx, input).Return(output, nil)

        result, err := mockClient.PostText(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutSession", func(t *testing.T) {
        input := &lexruntimeservice.PutSessionInput{}
        output := &lexruntimeservice.PutSessionOutput{}

        mockClient.On("PutSession", ctx, input).Return(output, nil)

        result, err := mockClient.PutSession(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
