package connectcontactlens_test

// tests for the connectcontactlens service interface for this ../../../service/connectcontactlens/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/connectcontactlens"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/connectcontactlens/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/connectcontactlens/connectcontactlens_iface"
	"github.com/stretchr/testify/assert"
)

func TestConnectcontactlensServiceCanBeMocked(t *testing.T) {
	var iface connectcontactlens_iface.IClient
	iface = &connectcontactlens.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := connectcontactlens.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListRealtimeContactAnalysisSegments", func(t *testing.T) {
        input := &connectcontactlens.ListRealtimeContactAnalysisSegmentsInput{}
        output := &connectcontactlens.ListRealtimeContactAnalysisSegmentsOutput{}

        mockClient.On("ListRealtimeContactAnalysisSegments", ctx, input).Return(output, nil)

        result, err := mockClient.ListRealtimeContactAnalysisSegments(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
