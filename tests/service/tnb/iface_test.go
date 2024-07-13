package tnb_test

// tests for the tnb service interface for this ../../../service/tnb/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/tnb"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/tnb/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/tnb/tnb_iface"
	"github.com/stretchr/testify/assert"
)

func TestTnbServiceCanBeMocked(t *testing.T) {
	var iface tnb_iface.IClient
	iface = &tnb.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := tnb.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCancelSolNetworkOperation", func(t *testing.T) {
        input := &tnb.CancelSolNetworkOperationInput{}
        output := &tnb.CancelSolNetworkOperationOutput{}

        mockClient.On("CancelSolNetworkOperation", ctx, input).Return(output, nil)

        result, err := mockClient.CancelSolNetworkOperation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateSolFunctionPackage", func(t *testing.T) {
        input := &tnb.CreateSolFunctionPackageInput{}
        output := &tnb.CreateSolFunctionPackageOutput{}

        mockClient.On("CreateSolFunctionPackage", ctx, input).Return(output, nil)

        result, err := mockClient.CreateSolFunctionPackage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateSolNetworkInstance", func(t *testing.T) {
        input := &tnb.CreateSolNetworkInstanceInput{}
        output := &tnb.CreateSolNetworkInstanceOutput{}

        mockClient.On("CreateSolNetworkInstance", ctx, input).Return(output, nil)

        result, err := mockClient.CreateSolNetworkInstance(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateSolNetworkPackage", func(t *testing.T) {
        input := &tnb.CreateSolNetworkPackageInput{}
        output := &tnb.CreateSolNetworkPackageOutput{}

        mockClient.On("CreateSolNetworkPackage", ctx, input).Return(output, nil)

        result, err := mockClient.CreateSolNetworkPackage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteSolFunctionPackage", func(t *testing.T) {
        input := &tnb.DeleteSolFunctionPackageInput{}
        output := &tnb.DeleteSolFunctionPackageOutput{}

        mockClient.On("DeleteSolFunctionPackage", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteSolFunctionPackage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteSolNetworkInstance", func(t *testing.T) {
        input := &tnb.DeleteSolNetworkInstanceInput{}
        output := &tnb.DeleteSolNetworkInstanceOutput{}

        mockClient.On("DeleteSolNetworkInstance", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteSolNetworkInstance(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteSolNetworkPackage", func(t *testing.T) {
        input := &tnb.DeleteSolNetworkPackageInput{}
        output := &tnb.DeleteSolNetworkPackageOutput{}

        mockClient.On("DeleteSolNetworkPackage", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteSolNetworkPackage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSolFunctionInstance", func(t *testing.T) {
        input := &tnb.GetSolFunctionInstanceInput{}
        output := &tnb.GetSolFunctionInstanceOutput{}

        mockClient.On("GetSolFunctionInstance", ctx, input).Return(output, nil)

        result, err := mockClient.GetSolFunctionInstance(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSolFunctionPackage", func(t *testing.T) {
        input := &tnb.GetSolFunctionPackageInput{}
        output := &tnb.GetSolFunctionPackageOutput{}

        mockClient.On("GetSolFunctionPackage", ctx, input).Return(output, nil)

        result, err := mockClient.GetSolFunctionPackage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSolFunctionPackageContent", func(t *testing.T) {
        input := &tnb.GetSolFunctionPackageContentInput{}
        output := &tnb.GetSolFunctionPackageContentOutput{}

        mockClient.On("GetSolFunctionPackageContent", ctx, input).Return(output, nil)

        result, err := mockClient.GetSolFunctionPackageContent(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSolFunctionPackageDescriptor", func(t *testing.T) {
        input := &tnb.GetSolFunctionPackageDescriptorInput{}
        output := &tnb.GetSolFunctionPackageDescriptorOutput{}

        mockClient.On("GetSolFunctionPackageDescriptor", ctx, input).Return(output, nil)

        result, err := mockClient.GetSolFunctionPackageDescriptor(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSolNetworkInstance", func(t *testing.T) {
        input := &tnb.GetSolNetworkInstanceInput{}
        output := &tnb.GetSolNetworkInstanceOutput{}

        mockClient.On("GetSolNetworkInstance", ctx, input).Return(output, nil)

        result, err := mockClient.GetSolNetworkInstance(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSolNetworkOperation", func(t *testing.T) {
        input := &tnb.GetSolNetworkOperationInput{}
        output := &tnb.GetSolNetworkOperationOutput{}

        mockClient.On("GetSolNetworkOperation", ctx, input).Return(output, nil)

        result, err := mockClient.GetSolNetworkOperation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSolNetworkPackage", func(t *testing.T) {
        input := &tnb.GetSolNetworkPackageInput{}
        output := &tnb.GetSolNetworkPackageOutput{}

        mockClient.On("GetSolNetworkPackage", ctx, input).Return(output, nil)

        result, err := mockClient.GetSolNetworkPackage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSolNetworkPackageContent", func(t *testing.T) {
        input := &tnb.GetSolNetworkPackageContentInput{}
        output := &tnb.GetSolNetworkPackageContentOutput{}

        mockClient.On("GetSolNetworkPackageContent", ctx, input).Return(output, nil)

        result, err := mockClient.GetSolNetworkPackageContent(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSolNetworkPackageDescriptor", func(t *testing.T) {
        input := &tnb.GetSolNetworkPackageDescriptorInput{}
        output := &tnb.GetSolNetworkPackageDescriptorOutput{}

        mockClient.On("GetSolNetworkPackageDescriptor", ctx, input).Return(output, nil)

        result, err := mockClient.GetSolNetworkPackageDescriptor(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestInstantiateSolNetworkInstance", func(t *testing.T) {
        input := &tnb.InstantiateSolNetworkInstanceInput{}
        output := &tnb.InstantiateSolNetworkInstanceOutput{}

        mockClient.On("InstantiateSolNetworkInstance", ctx, input).Return(output, nil)

        result, err := mockClient.InstantiateSolNetworkInstance(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListSolFunctionInstances", func(t *testing.T) {
        input := &tnb.ListSolFunctionInstancesInput{}
        output := &tnb.ListSolFunctionInstancesOutput{}

        mockClient.On("ListSolFunctionInstances", ctx, input).Return(output, nil)

        result, err := mockClient.ListSolFunctionInstances(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListSolFunctionPackages", func(t *testing.T) {
        input := &tnb.ListSolFunctionPackagesInput{}
        output := &tnb.ListSolFunctionPackagesOutput{}

        mockClient.On("ListSolFunctionPackages", ctx, input).Return(output, nil)

        result, err := mockClient.ListSolFunctionPackages(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListSolNetworkInstances", func(t *testing.T) {
        input := &tnb.ListSolNetworkInstancesInput{}
        output := &tnb.ListSolNetworkInstancesOutput{}

        mockClient.On("ListSolNetworkInstances", ctx, input).Return(output, nil)

        result, err := mockClient.ListSolNetworkInstances(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListSolNetworkOperations", func(t *testing.T) {
        input := &tnb.ListSolNetworkOperationsInput{}
        output := &tnb.ListSolNetworkOperationsOutput{}

        mockClient.On("ListSolNetworkOperations", ctx, input).Return(output, nil)

        result, err := mockClient.ListSolNetworkOperations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListSolNetworkPackages", func(t *testing.T) {
        input := &tnb.ListSolNetworkPackagesInput{}
        output := &tnb.ListSolNetworkPackagesOutput{}

        mockClient.On("ListSolNetworkPackages", ctx, input).Return(output, nil)

        result, err := mockClient.ListSolNetworkPackages(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &tnb.ListTagsForResourceInput{}
        output := &tnb.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutSolFunctionPackageContent", func(t *testing.T) {
        input := &tnb.PutSolFunctionPackageContentInput{}
        output := &tnb.PutSolFunctionPackageContentOutput{}

        mockClient.On("PutSolFunctionPackageContent", ctx, input).Return(output, nil)

        result, err := mockClient.PutSolFunctionPackageContent(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutSolNetworkPackageContent", func(t *testing.T) {
        input := &tnb.PutSolNetworkPackageContentInput{}
        output := &tnb.PutSolNetworkPackageContentOutput{}

        mockClient.On("PutSolNetworkPackageContent", ctx, input).Return(output, nil)

        result, err := mockClient.PutSolNetworkPackageContent(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &tnb.TagResourceInput{}
        output := &tnb.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTerminateSolNetworkInstance", func(t *testing.T) {
        input := &tnb.TerminateSolNetworkInstanceInput{}
        output := &tnb.TerminateSolNetworkInstanceOutput{}

        mockClient.On("TerminateSolNetworkInstance", ctx, input).Return(output, nil)

        result, err := mockClient.TerminateSolNetworkInstance(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &tnb.UntagResourceInput{}
        output := &tnb.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateSolFunctionPackage", func(t *testing.T) {
        input := &tnb.UpdateSolFunctionPackageInput{}
        output := &tnb.UpdateSolFunctionPackageOutput{}

        mockClient.On("UpdateSolFunctionPackage", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateSolFunctionPackage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateSolNetworkInstance", func(t *testing.T) {
        input := &tnb.UpdateSolNetworkInstanceInput{}
        output := &tnb.UpdateSolNetworkInstanceOutput{}

        mockClient.On("UpdateSolNetworkInstance", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateSolNetworkInstance(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateSolNetworkPackage", func(t *testing.T) {
        input := &tnb.UpdateSolNetworkPackageInput{}
        output := &tnb.UpdateSolNetworkPackageOutput{}

        mockClient.On("UpdateSolNetworkPackage", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateSolNetworkPackage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestValidateSolFunctionPackageContent", func(t *testing.T) {
        input := &tnb.ValidateSolFunctionPackageContentInput{}
        output := &tnb.ValidateSolFunctionPackageContentOutput{}

        mockClient.On("ValidateSolFunctionPackageContent", ctx, input).Return(output, nil)

        result, err := mockClient.ValidateSolFunctionPackageContent(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestValidateSolNetworkPackageContent", func(t *testing.T) {
        input := &tnb.ValidateSolNetworkPackageContentInput{}
        output := &tnb.ValidateSolNetworkPackageContentOutput{}

        mockClient.On("ValidateSolNetworkPackageContent", ctx, input).Return(output, nil)

        result, err := mockClient.ValidateSolNetworkPackageContent(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
