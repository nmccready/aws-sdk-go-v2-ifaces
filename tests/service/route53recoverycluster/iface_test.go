package route53recoverycluster_test

// tests for the route53recoverycluster service interface for this ../../../service/route53recoverycluster/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/route53recoverycluster"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/route53recoverycluster/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/route53recoverycluster/route53recoverycluster_iface"
	"github.com/stretchr/testify/assert"
)

func TestRoute53recoveryclusterServiceCanBeMocked(t *testing.T) {
	var iface route53recoverycluster_iface.IClient
	iface = &route53recoverycluster.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := route53recoverycluster.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetRoutingControlState", func(t *testing.T) {
        input := &route53recoverycluster.GetRoutingControlStateInput{}
        output := &route53recoverycluster.GetRoutingControlStateOutput{}

        mockClient.On("GetRoutingControlState", ctx, input).Return(output, nil)

        result, err := mockClient.GetRoutingControlState(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListRoutingControls", func(t *testing.T) {
        input := &route53recoverycluster.ListRoutingControlsInput{}
        output := &route53recoverycluster.ListRoutingControlsOutput{}

        mockClient.On("ListRoutingControls", ctx, input).Return(output, nil)

        result, err := mockClient.ListRoutingControls(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateRoutingControlState", func(t *testing.T) {
        input := &route53recoverycluster.UpdateRoutingControlStateInput{}
        output := &route53recoverycluster.UpdateRoutingControlStateOutput{}

        mockClient.On("UpdateRoutingControlState", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateRoutingControlState(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateRoutingControlStates", func(t *testing.T) {
        input := &route53recoverycluster.UpdateRoutingControlStatesInput{}
        output := &route53recoverycluster.UpdateRoutingControlStatesOutput{}

        mockClient.On("UpdateRoutingControlStates", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateRoutingControlStates(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
