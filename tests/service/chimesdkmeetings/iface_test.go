package chimesdkmeetings_test

// tests for the chimesdkmeetings service interface for this ../../../service/chimesdkmeetings/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/chimesdkmeetings"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/chimesdkmeetings/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/chimesdkmeetings/chimesdkmeetings_iface"
	"github.com/stretchr/testify/assert"
)

func TestChimesdkmeetingsServiceCanBeMocked(t *testing.T) {
	var iface chimesdkmeetings_iface.IClient
	iface = &chimesdkmeetings.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := chimesdkmeetings.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchCreateAttendee", func(t *testing.T) {
        input := &chimesdkmeetings.BatchCreateAttendeeInput{}
        output := &chimesdkmeetings.BatchCreateAttendeeOutput{}

        mockClient.On("BatchCreateAttendee", ctx, input).Return(output, nil)

        result, err := mockClient.BatchCreateAttendee(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchUpdateAttendeeCapabilitiesExcept", func(t *testing.T) {
        input := &chimesdkmeetings.BatchUpdateAttendeeCapabilitiesExceptInput{}
        output := &chimesdkmeetings.BatchUpdateAttendeeCapabilitiesExceptOutput{}

        mockClient.On("BatchUpdateAttendeeCapabilitiesExcept", ctx, input).Return(output, nil)

        result, err := mockClient.BatchUpdateAttendeeCapabilitiesExcept(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateAttendee", func(t *testing.T) {
        input := &chimesdkmeetings.CreateAttendeeInput{}
        output := &chimesdkmeetings.CreateAttendeeOutput{}

        mockClient.On("CreateAttendee", ctx, input).Return(output, nil)

        result, err := mockClient.CreateAttendee(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateMeeting", func(t *testing.T) {
        input := &chimesdkmeetings.CreateMeetingInput{}
        output := &chimesdkmeetings.CreateMeetingOutput{}

        mockClient.On("CreateMeeting", ctx, input).Return(output, nil)

        result, err := mockClient.CreateMeeting(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateMeetingWithAttendees", func(t *testing.T) {
        input := &chimesdkmeetings.CreateMeetingWithAttendeesInput{}
        output := &chimesdkmeetings.CreateMeetingWithAttendeesOutput{}

        mockClient.On("CreateMeetingWithAttendees", ctx, input).Return(output, nil)

        result, err := mockClient.CreateMeetingWithAttendees(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteAttendee", func(t *testing.T) {
        input := &chimesdkmeetings.DeleteAttendeeInput{}
        output := &chimesdkmeetings.DeleteAttendeeOutput{}

        mockClient.On("DeleteAttendee", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteAttendee(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteMeeting", func(t *testing.T) {
        input := &chimesdkmeetings.DeleteMeetingInput{}
        output := &chimesdkmeetings.DeleteMeetingOutput{}

        mockClient.On("DeleteMeeting", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteMeeting(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetAttendee", func(t *testing.T) {
        input := &chimesdkmeetings.GetAttendeeInput{}
        output := &chimesdkmeetings.GetAttendeeOutput{}

        mockClient.On("GetAttendee", ctx, input).Return(output, nil)

        result, err := mockClient.GetAttendee(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetMeeting", func(t *testing.T) {
        input := &chimesdkmeetings.GetMeetingInput{}
        output := &chimesdkmeetings.GetMeetingOutput{}

        mockClient.On("GetMeeting", ctx, input).Return(output, nil)

        result, err := mockClient.GetMeeting(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAttendees", func(t *testing.T) {
        input := &chimesdkmeetings.ListAttendeesInput{}
        output := &chimesdkmeetings.ListAttendeesOutput{}

        mockClient.On("ListAttendees", ctx, input).Return(output, nil)

        result, err := mockClient.ListAttendees(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &chimesdkmeetings.ListTagsForResourceInput{}
        output := &chimesdkmeetings.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartMeetingTranscription", func(t *testing.T) {
        input := &chimesdkmeetings.StartMeetingTranscriptionInput{}
        output := &chimesdkmeetings.StartMeetingTranscriptionOutput{}

        mockClient.On("StartMeetingTranscription", ctx, input).Return(output, nil)

        result, err := mockClient.StartMeetingTranscription(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopMeetingTranscription", func(t *testing.T) {
        input := &chimesdkmeetings.StopMeetingTranscriptionInput{}
        output := &chimesdkmeetings.StopMeetingTranscriptionOutput{}

        mockClient.On("StopMeetingTranscription", ctx, input).Return(output, nil)

        result, err := mockClient.StopMeetingTranscription(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &chimesdkmeetings.TagResourceInput{}
        output := &chimesdkmeetings.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &chimesdkmeetings.UntagResourceInput{}
        output := &chimesdkmeetings.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateAttendeeCapabilities", func(t *testing.T) {
        input := &chimesdkmeetings.UpdateAttendeeCapabilitiesInput{}
        output := &chimesdkmeetings.UpdateAttendeeCapabilitiesOutput{}

        mockClient.On("UpdateAttendeeCapabilities", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateAttendeeCapabilities(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
