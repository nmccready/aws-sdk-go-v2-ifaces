// Code generated by ./scripts/gen_tests.sh DO NOT EDIT.
package workmailmessageflow_test

// tests for the workmailmessageflow service interface for 
// this ../../../service/workmailmessageflow/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/workmailmessageflow"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/workmailmessageflow/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/workmailmessageflow/workmailmessageflow_iface"
	"github.com/stretchr/testify/assert"
)

func TestWorkmailmessageflowServiceCanBeMocked(t *testing.T) {
	var iface workmailmessageflow_iface.IClient
	iface = &workmailmessageflow.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := workmailmessageflow.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetRawMessageContent", func(t *testing.T) {
        input := &workmailmessageflow.GetRawMessageContentInput{}
        output := &workmailmessageflow.GetRawMessageContentOutput{}

        mockClient.On("GetRawMessageContent", ctx, input).Return(output, nil)

        result, err := mockClient.GetRawMessageContent(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutRawMessageContent", func(t *testing.T) {
        input := &workmailmessageflow.PutRawMessageContentInput{}
        output := &workmailmessageflow.PutRawMessageContentOutput{}

        mockClient.On("PutRawMessageContent", ctx, input).Return(output, nil)

        result, err := mockClient.PutRawMessageContent(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
