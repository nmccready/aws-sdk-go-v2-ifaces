package iot1clickprojects_test

// tests for the iot1clickprojects service interface for this ../../../service/iot1clickprojects/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/iot1clickprojects"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/iot1clickprojects/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/iot1clickprojects/iot1clickprojects_iface"
	"github.com/stretchr/testify/assert"
)

func TestIot1clickprojectsServiceCanBeMocked(t *testing.T) {
	var iface iot1clickprojects_iface.IClient
	iface = &iot1clickprojects.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := iot1clickprojects.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateDeviceWithPlacement", func(t *testing.T) {
        input := &iot1clickprojects.AssociateDeviceWithPlacementInput{}
        output := &iot1clickprojects.AssociateDeviceWithPlacementOutput{}

        mockClient.On("AssociateDeviceWithPlacement", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateDeviceWithPlacement(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreatePlacement", func(t *testing.T) {
        input := &iot1clickprojects.CreatePlacementInput{}
        output := &iot1clickprojects.CreatePlacementOutput{}

        mockClient.On("CreatePlacement", ctx, input).Return(output, nil)

        result, err := mockClient.CreatePlacement(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateProject", func(t *testing.T) {
        input := &iot1clickprojects.CreateProjectInput{}
        output := &iot1clickprojects.CreateProjectOutput{}

        mockClient.On("CreateProject", ctx, input).Return(output, nil)

        result, err := mockClient.CreateProject(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeletePlacement", func(t *testing.T) {
        input := &iot1clickprojects.DeletePlacementInput{}
        output := &iot1clickprojects.DeletePlacementOutput{}

        mockClient.On("DeletePlacement", ctx, input).Return(output, nil)

        result, err := mockClient.DeletePlacement(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteProject", func(t *testing.T) {
        input := &iot1clickprojects.DeleteProjectInput{}
        output := &iot1clickprojects.DeleteProjectOutput{}

        mockClient.On("DeleteProject", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteProject(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribePlacement", func(t *testing.T) {
        input := &iot1clickprojects.DescribePlacementInput{}
        output := &iot1clickprojects.DescribePlacementOutput{}

        mockClient.On("DescribePlacement", ctx, input).Return(output, nil)

        result, err := mockClient.DescribePlacement(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeProject", func(t *testing.T) {
        input := &iot1clickprojects.DescribeProjectInput{}
        output := &iot1clickprojects.DescribeProjectOutput{}

        mockClient.On("DescribeProject", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeProject(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateDeviceFromPlacement", func(t *testing.T) {
        input := &iot1clickprojects.DisassociateDeviceFromPlacementInput{}
        output := &iot1clickprojects.DisassociateDeviceFromPlacementOutput{}

        mockClient.On("DisassociateDeviceFromPlacement", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateDeviceFromPlacement(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDevicesInPlacement", func(t *testing.T) {
        input := &iot1clickprojects.GetDevicesInPlacementInput{}
        output := &iot1clickprojects.GetDevicesInPlacementOutput{}

        mockClient.On("GetDevicesInPlacement", ctx, input).Return(output, nil)

        result, err := mockClient.GetDevicesInPlacement(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListPlacements", func(t *testing.T) {
        input := &iot1clickprojects.ListPlacementsInput{}
        output := &iot1clickprojects.ListPlacementsOutput{}

        mockClient.On("ListPlacements", ctx, input).Return(output, nil)

        result, err := mockClient.ListPlacements(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListProjects", func(t *testing.T) {
        input := &iot1clickprojects.ListProjectsInput{}
        output := &iot1clickprojects.ListProjectsOutput{}

        mockClient.On("ListProjects", ctx, input).Return(output, nil)

        result, err := mockClient.ListProjects(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &iot1clickprojects.ListTagsForResourceInput{}
        output := &iot1clickprojects.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &iot1clickprojects.TagResourceInput{}
        output := &iot1clickprojects.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &iot1clickprojects.UntagResourceInput{}
        output := &iot1clickprojects.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdatePlacement", func(t *testing.T) {
        input := &iot1clickprojects.UpdatePlacementInput{}
        output := &iot1clickprojects.UpdatePlacementOutput{}

        mockClient.On("UpdatePlacement", ctx, input).Return(output, nil)

        result, err := mockClient.UpdatePlacement(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateProject", func(t *testing.T) {
        input := &iot1clickprojects.UpdateProjectInput{}
        output := &iot1clickprojects.UpdateProjectOutput{}

        mockClient.On("UpdateProject", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateProject(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
