package controltower_test

// tests for the controltower service interface for this ../../../service/controltower/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/controltower"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/controltower/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/controltower/controltower_iface"
	"github.com/stretchr/testify/assert"
)

func TestControltowerServiceCanBeMocked(t *testing.T) {
	var iface controltower_iface.IClient
	iface = &controltower.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := controltower.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateLandingZone", func(t *testing.T) {
        input := &controltower.CreateLandingZoneInput{}
        output := &controltower.CreateLandingZoneOutput{}

        mockClient.On("CreateLandingZone", ctx, input).Return(output, nil)

        result, err := mockClient.CreateLandingZone(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteLandingZone", func(t *testing.T) {
        input := &controltower.DeleteLandingZoneInput{}
        output := &controltower.DeleteLandingZoneOutput{}

        mockClient.On("DeleteLandingZone", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteLandingZone(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisableBaseline", func(t *testing.T) {
        input := &controltower.DisableBaselineInput{}
        output := &controltower.DisableBaselineOutput{}

        mockClient.On("DisableBaseline", ctx, input).Return(output, nil)

        result, err := mockClient.DisableBaseline(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisableControl", func(t *testing.T) {
        input := &controltower.DisableControlInput{}
        output := &controltower.DisableControlOutput{}

        mockClient.On("DisableControl", ctx, input).Return(output, nil)

        result, err := mockClient.DisableControl(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestEnableBaseline", func(t *testing.T) {
        input := &controltower.EnableBaselineInput{}
        output := &controltower.EnableBaselineOutput{}

        mockClient.On("EnableBaseline", ctx, input).Return(output, nil)

        result, err := mockClient.EnableBaseline(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestEnableControl", func(t *testing.T) {
        input := &controltower.EnableControlInput{}
        output := &controltower.EnableControlOutput{}

        mockClient.On("EnableControl", ctx, input).Return(output, nil)

        result, err := mockClient.EnableControl(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetBaseline", func(t *testing.T) {
        input := &controltower.GetBaselineInput{}
        output := &controltower.GetBaselineOutput{}

        mockClient.On("GetBaseline", ctx, input).Return(output, nil)

        result, err := mockClient.GetBaseline(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetBaselineOperation", func(t *testing.T) {
        input := &controltower.GetBaselineOperationInput{}
        output := &controltower.GetBaselineOperationOutput{}

        mockClient.On("GetBaselineOperation", ctx, input).Return(output, nil)

        result, err := mockClient.GetBaselineOperation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetControlOperation", func(t *testing.T) {
        input := &controltower.GetControlOperationInput{}
        output := &controltower.GetControlOperationOutput{}

        mockClient.On("GetControlOperation", ctx, input).Return(output, nil)

        result, err := mockClient.GetControlOperation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetEnabledBaseline", func(t *testing.T) {
        input := &controltower.GetEnabledBaselineInput{}
        output := &controltower.GetEnabledBaselineOutput{}

        mockClient.On("GetEnabledBaseline", ctx, input).Return(output, nil)

        result, err := mockClient.GetEnabledBaseline(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetEnabledControl", func(t *testing.T) {
        input := &controltower.GetEnabledControlInput{}
        output := &controltower.GetEnabledControlOutput{}

        mockClient.On("GetEnabledControl", ctx, input).Return(output, nil)

        result, err := mockClient.GetEnabledControl(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetLandingZone", func(t *testing.T) {
        input := &controltower.GetLandingZoneInput{}
        output := &controltower.GetLandingZoneOutput{}

        mockClient.On("GetLandingZone", ctx, input).Return(output, nil)

        result, err := mockClient.GetLandingZone(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetLandingZoneOperation", func(t *testing.T) {
        input := &controltower.GetLandingZoneOperationInput{}
        output := &controltower.GetLandingZoneOperationOutput{}

        mockClient.On("GetLandingZoneOperation", ctx, input).Return(output, nil)

        result, err := mockClient.GetLandingZoneOperation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListBaselines", func(t *testing.T) {
        input := &controltower.ListBaselinesInput{}
        output := &controltower.ListBaselinesOutput{}

        mockClient.On("ListBaselines", ctx, input).Return(output, nil)

        result, err := mockClient.ListBaselines(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListControlOperations", func(t *testing.T) {
        input := &controltower.ListControlOperationsInput{}
        output := &controltower.ListControlOperationsOutput{}

        mockClient.On("ListControlOperations", ctx, input).Return(output, nil)

        result, err := mockClient.ListControlOperations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListEnabledBaselines", func(t *testing.T) {
        input := &controltower.ListEnabledBaselinesInput{}
        output := &controltower.ListEnabledBaselinesOutput{}

        mockClient.On("ListEnabledBaselines", ctx, input).Return(output, nil)

        result, err := mockClient.ListEnabledBaselines(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListEnabledControls", func(t *testing.T) {
        input := &controltower.ListEnabledControlsInput{}
        output := &controltower.ListEnabledControlsOutput{}

        mockClient.On("ListEnabledControls", ctx, input).Return(output, nil)

        result, err := mockClient.ListEnabledControls(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListLandingZoneOperations", func(t *testing.T) {
        input := &controltower.ListLandingZoneOperationsInput{}
        output := &controltower.ListLandingZoneOperationsOutput{}

        mockClient.On("ListLandingZoneOperations", ctx, input).Return(output, nil)

        result, err := mockClient.ListLandingZoneOperations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListLandingZones", func(t *testing.T) {
        input := &controltower.ListLandingZonesInput{}
        output := &controltower.ListLandingZonesOutput{}

        mockClient.On("ListLandingZones", ctx, input).Return(output, nil)

        result, err := mockClient.ListLandingZones(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &controltower.ListTagsForResourceInput{}
        output := &controltower.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestResetEnabledBaseline", func(t *testing.T) {
        input := &controltower.ResetEnabledBaselineInput{}
        output := &controltower.ResetEnabledBaselineOutput{}

        mockClient.On("ResetEnabledBaseline", ctx, input).Return(output, nil)

        result, err := mockClient.ResetEnabledBaseline(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestResetLandingZone", func(t *testing.T) {
        input := &controltower.ResetLandingZoneInput{}
        output := &controltower.ResetLandingZoneOutput{}

        mockClient.On("ResetLandingZone", ctx, input).Return(output, nil)

        result, err := mockClient.ResetLandingZone(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &controltower.TagResourceInput{}
        output := &controltower.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &controltower.UntagResourceInput{}
        output := &controltower.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateEnabledBaseline", func(t *testing.T) {
        input := &controltower.UpdateEnabledBaselineInput{}
        output := &controltower.UpdateEnabledBaselineOutput{}

        mockClient.On("UpdateEnabledBaseline", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateEnabledBaseline(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateEnabledControl", func(t *testing.T) {
        input := &controltower.UpdateEnabledControlInput{}
        output := &controltower.UpdateEnabledControlOutput{}

        mockClient.On("UpdateEnabledControl", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateEnabledControl(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateLandingZone", func(t *testing.T) {
        input := &controltower.UpdateLandingZoneInput{}
        output := &controltower.UpdateLandingZoneOutput{}

        mockClient.On("UpdateLandingZone", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateLandingZone(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
