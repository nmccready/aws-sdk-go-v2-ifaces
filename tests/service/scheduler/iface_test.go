package scheduler_test

// tests for the scheduler service interface for this ../../../service/scheduler/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/scheduler"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/scheduler/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/scheduler/scheduler_iface"
	"github.com/stretchr/testify/assert"
)

func TestSchedulerServiceCanBeMocked(t *testing.T) {
	var iface scheduler_iface.IClient
	iface = &scheduler.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := scheduler.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateSchedule", func(t *testing.T) {
        input := &scheduler.CreateScheduleInput{}
        output := &scheduler.CreateScheduleOutput{}

        mockClient.On("CreateSchedule", ctx, input).Return(output, nil)

        result, err := mockClient.CreateSchedule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateScheduleGroup", func(t *testing.T) {
        input := &scheduler.CreateScheduleGroupInput{}
        output := &scheduler.CreateScheduleGroupOutput{}

        mockClient.On("CreateScheduleGroup", ctx, input).Return(output, nil)

        result, err := mockClient.CreateScheduleGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteSchedule", func(t *testing.T) {
        input := &scheduler.DeleteScheduleInput{}
        output := &scheduler.DeleteScheduleOutput{}

        mockClient.On("DeleteSchedule", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteSchedule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteScheduleGroup", func(t *testing.T) {
        input := &scheduler.DeleteScheduleGroupInput{}
        output := &scheduler.DeleteScheduleGroupOutput{}

        mockClient.On("DeleteScheduleGroup", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteScheduleGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSchedule", func(t *testing.T) {
        input := &scheduler.GetScheduleInput{}
        output := &scheduler.GetScheduleOutput{}

        mockClient.On("GetSchedule", ctx, input).Return(output, nil)

        result, err := mockClient.GetSchedule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetScheduleGroup", func(t *testing.T) {
        input := &scheduler.GetScheduleGroupInput{}
        output := &scheduler.GetScheduleGroupOutput{}

        mockClient.On("GetScheduleGroup", ctx, input).Return(output, nil)

        result, err := mockClient.GetScheduleGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListScheduleGroups", func(t *testing.T) {
        input := &scheduler.ListScheduleGroupsInput{}
        output := &scheduler.ListScheduleGroupsOutput{}

        mockClient.On("ListScheduleGroups", ctx, input).Return(output, nil)

        result, err := mockClient.ListScheduleGroups(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListSchedules", func(t *testing.T) {
        input := &scheduler.ListSchedulesInput{}
        output := &scheduler.ListSchedulesOutput{}

        mockClient.On("ListSchedules", ctx, input).Return(output, nil)

        result, err := mockClient.ListSchedules(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &scheduler.ListTagsForResourceInput{}
        output := &scheduler.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &scheduler.TagResourceInput{}
        output := &scheduler.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &scheduler.UntagResourceInput{}
        output := &scheduler.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateSchedule", func(t *testing.T) {
        input := &scheduler.UpdateScheduleInput{}
        output := &scheduler.UpdateScheduleOutput{}

        mockClient.On("UpdateSchedule", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateSchedule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
