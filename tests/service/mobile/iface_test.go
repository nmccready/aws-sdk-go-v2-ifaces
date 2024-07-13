package mobile_test

// tests for the mobile service interface for this ../../../service/mobile/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/mobile"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/mobile/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/mobile/mobile_iface"
	"github.com/stretchr/testify/assert"
)

func TestMobileServiceCanBeMocked(t *testing.T) {
	var iface mobile_iface.IClient
	iface = &mobile.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := mobile.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateProject", func(t *testing.T) {
        input := &mobile.CreateProjectInput{}
        output := &mobile.CreateProjectOutput{}

        mockClient.On("CreateProject", ctx, input).Return(output, nil)

        result, err := mockClient.CreateProject(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteProject", func(t *testing.T) {
        input := &mobile.DeleteProjectInput{}
        output := &mobile.DeleteProjectOutput{}

        mockClient.On("DeleteProject", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteProject(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeBundle", func(t *testing.T) {
        input := &mobile.DescribeBundleInput{}
        output := &mobile.DescribeBundleOutput{}

        mockClient.On("DescribeBundle", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeBundle(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeProject", func(t *testing.T) {
        input := &mobile.DescribeProjectInput{}
        output := &mobile.DescribeProjectOutput{}

        mockClient.On("DescribeProject", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeProject(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestExportBundle", func(t *testing.T) {
        input := &mobile.ExportBundleInput{}
        output := &mobile.ExportBundleOutput{}

        mockClient.On("ExportBundle", ctx, input).Return(output, nil)

        result, err := mockClient.ExportBundle(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestExportProject", func(t *testing.T) {
        input := &mobile.ExportProjectInput{}
        output := &mobile.ExportProjectOutput{}

        mockClient.On("ExportProject", ctx, input).Return(output, nil)

        result, err := mockClient.ExportProject(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListBundles", func(t *testing.T) {
        input := &mobile.ListBundlesInput{}
        output := &mobile.ListBundlesOutput{}

        mockClient.On("ListBundles", ctx, input).Return(output, nil)

        result, err := mockClient.ListBundles(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListProjects", func(t *testing.T) {
        input := &mobile.ListProjectsInput{}
        output := &mobile.ListProjectsOutput{}

        mockClient.On("ListProjects", ctx, input).Return(output, nil)

        result, err := mockClient.ListProjects(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateProject", func(t *testing.T) {
        input := &mobile.UpdateProjectInput{}
        output := &mobile.UpdateProjectOutput{}

        mockClient.On("UpdateProject", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateProject(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
