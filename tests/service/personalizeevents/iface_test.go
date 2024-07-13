package personalizeevents_test

// tests for the personalizeevents service interface for this ../../../service/personalizeevents/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/personalizeevents"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/personalizeevents/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/personalizeevents/personalizeevents_iface"
	"github.com/stretchr/testify/assert"
)

func TestPersonalizeeventsServiceCanBeMocked(t *testing.T) {
	var iface personalizeevents_iface.IClient
	iface = &personalizeevents.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := personalizeevents.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutActionInteractions", func(t *testing.T) {
        input := &personalizeevents.PutActionInteractionsInput{}
        output := &personalizeevents.PutActionInteractionsOutput{}

        mockClient.On("PutActionInteractions", ctx, input).Return(output, nil)

        result, err := mockClient.PutActionInteractions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutActions", func(t *testing.T) {
        input := &personalizeevents.PutActionsInput{}
        output := &personalizeevents.PutActionsOutput{}

        mockClient.On("PutActions", ctx, input).Return(output, nil)

        result, err := mockClient.PutActions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutEvents", func(t *testing.T) {
        input := &personalizeevents.PutEventsInput{}
        output := &personalizeevents.PutEventsOutput{}

        mockClient.On("PutEvents", ctx, input).Return(output, nil)

        result, err := mockClient.PutEvents(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutItems", func(t *testing.T) {
        input := &personalizeevents.PutItemsInput{}
        output := &personalizeevents.PutItemsOutput{}

        mockClient.On("PutItems", ctx, input).Return(output, nil)

        result, err := mockClient.PutItems(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutUsers", func(t *testing.T) {
        input := &personalizeevents.PutUsersInput{}
        output := &personalizeevents.PutUsersOutput{}

        mockClient.On("PutUsers", ctx, input).Return(output, nil)

        result, err := mockClient.PutUsers(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
