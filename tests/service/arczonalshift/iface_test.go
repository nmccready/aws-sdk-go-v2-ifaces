package arczonalshift_test

// tests for the arczonalshift service interface for this ../../../service/arczonalshift/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/arczonalshift"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/arczonalshift/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/arczonalshift/arczonalshift_iface"
	"github.com/stretchr/testify/assert"
)

func TestArczonalshiftServiceCanBeMocked(t *testing.T) {
	var iface arczonalshift_iface.IClient
	iface = &arczonalshift.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := arczonalshift.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCancelZonalShift", func(t *testing.T) {
        input := &arczonalshift.CancelZonalShiftInput{}
        output := &arczonalshift.CancelZonalShiftOutput{}

        mockClient.On("CancelZonalShift", ctx, input).Return(output, nil)

        result, err := mockClient.CancelZonalShift(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreatePracticeRunConfiguration", func(t *testing.T) {
        input := &arczonalshift.CreatePracticeRunConfigurationInput{}
        output := &arczonalshift.CreatePracticeRunConfigurationOutput{}

        mockClient.On("CreatePracticeRunConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.CreatePracticeRunConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeletePracticeRunConfiguration", func(t *testing.T) {
        input := &arczonalshift.DeletePracticeRunConfigurationInput{}
        output := &arczonalshift.DeletePracticeRunConfigurationOutput{}

        mockClient.On("DeletePracticeRunConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.DeletePracticeRunConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetManagedResource", func(t *testing.T) {
        input := &arczonalshift.GetManagedResourceInput{}
        output := &arczonalshift.GetManagedResourceOutput{}

        mockClient.On("GetManagedResource", ctx, input).Return(output, nil)

        result, err := mockClient.GetManagedResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAutoshifts", func(t *testing.T) {
        input := &arczonalshift.ListAutoshiftsInput{}
        output := &arczonalshift.ListAutoshiftsOutput{}

        mockClient.On("ListAutoshifts", ctx, input).Return(output, nil)

        result, err := mockClient.ListAutoshifts(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListManagedResources", func(t *testing.T) {
        input := &arczonalshift.ListManagedResourcesInput{}
        output := &arczonalshift.ListManagedResourcesOutput{}

        mockClient.On("ListManagedResources", ctx, input).Return(output, nil)

        result, err := mockClient.ListManagedResources(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListZonalShifts", func(t *testing.T) {
        input := &arczonalshift.ListZonalShiftsInput{}
        output := &arczonalshift.ListZonalShiftsOutput{}

        mockClient.On("ListZonalShifts", ctx, input).Return(output, nil)

        result, err := mockClient.ListZonalShifts(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartZonalShift", func(t *testing.T) {
        input := &arczonalshift.StartZonalShiftInput{}
        output := &arczonalshift.StartZonalShiftOutput{}

        mockClient.On("StartZonalShift", ctx, input).Return(output, nil)

        result, err := mockClient.StartZonalShift(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdatePracticeRunConfiguration", func(t *testing.T) {
        input := &arczonalshift.UpdatePracticeRunConfigurationInput{}
        output := &arczonalshift.UpdatePracticeRunConfigurationOutput{}

        mockClient.On("UpdatePracticeRunConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.UpdatePracticeRunConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateZonalAutoshiftConfiguration", func(t *testing.T) {
        input := &arczonalshift.UpdateZonalAutoshiftConfigurationInput{}
        output := &arczonalshift.UpdateZonalAutoshiftConfigurationOutput{}

        mockClient.On("UpdateZonalAutoshiftConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateZonalAutoshiftConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateZonalShift", func(t *testing.T) {
        input := &arczonalshift.UpdateZonalShiftInput{}
        output := &arczonalshift.UpdateZonalShiftOutput{}

        mockClient.On("UpdateZonalShift", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateZonalShift(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
