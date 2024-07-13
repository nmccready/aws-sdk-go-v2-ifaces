package cloudtraildata_test

// tests for the cloudtraildata service interface for this ../../../service/cloudtraildata/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/cloudtraildata"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/cloudtraildata/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/cloudtraildata/cloudtraildata_iface"
	"github.com/stretchr/testify/assert"
)

func TestCloudtraildataServiceCanBeMocked(t *testing.T) {
	var iface cloudtraildata_iface.IClient
	iface = &cloudtraildata.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := cloudtraildata.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutAuditEvents", func(t *testing.T) {
        input := &cloudtraildata.PutAuditEventsInput{}
        output := &cloudtraildata.PutAuditEventsOutput{}

        mockClient.On("PutAuditEvents", ctx, input).Return(output, nil)

        result, err := mockClient.PutAuditEvents(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
