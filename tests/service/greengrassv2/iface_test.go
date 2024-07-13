package greengrassv2_test

// tests for the greengrassv2 service interface for this ../../../service/greengrassv2/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/greengrassv2"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/greengrassv2/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/greengrassv2/greengrassv2_iface"
	"github.com/stretchr/testify/assert"
)

func TestGreengrassv2ServiceCanBeMocked(t *testing.T) {
	var iface greengrassv2_iface.IClient
	iface = &greengrassv2.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := greengrassv2.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateServiceRoleToAccount", func(t *testing.T) {
        input := &greengrassv2.AssociateServiceRoleToAccountInput{}
        output := &greengrassv2.AssociateServiceRoleToAccountOutput{}

        mockClient.On("AssociateServiceRoleToAccount", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateServiceRoleToAccount(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchAssociateClientDeviceWithCoreDevice", func(t *testing.T) {
        input := &greengrassv2.BatchAssociateClientDeviceWithCoreDeviceInput{}
        output := &greengrassv2.BatchAssociateClientDeviceWithCoreDeviceOutput{}

        mockClient.On("BatchAssociateClientDeviceWithCoreDevice", ctx, input).Return(output, nil)

        result, err := mockClient.BatchAssociateClientDeviceWithCoreDevice(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchDisassociateClientDeviceFromCoreDevice", func(t *testing.T) {
        input := &greengrassv2.BatchDisassociateClientDeviceFromCoreDeviceInput{}
        output := &greengrassv2.BatchDisassociateClientDeviceFromCoreDeviceOutput{}

        mockClient.On("BatchDisassociateClientDeviceFromCoreDevice", ctx, input).Return(output, nil)

        result, err := mockClient.BatchDisassociateClientDeviceFromCoreDevice(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCancelDeployment", func(t *testing.T) {
        input := &greengrassv2.CancelDeploymentInput{}
        output := &greengrassv2.CancelDeploymentOutput{}

        mockClient.On("CancelDeployment", ctx, input).Return(output, nil)

        result, err := mockClient.CancelDeployment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateComponentVersion", func(t *testing.T) {
        input := &greengrassv2.CreateComponentVersionInput{}
        output := &greengrassv2.CreateComponentVersionOutput{}

        mockClient.On("CreateComponentVersion", ctx, input).Return(output, nil)

        result, err := mockClient.CreateComponentVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateDeployment", func(t *testing.T) {
        input := &greengrassv2.CreateDeploymentInput{}
        output := &greengrassv2.CreateDeploymentOutput{}

        mockClient.On("CreateDeployment", ctx, input).Return(output, nil)

        result, err := mockClient.CreateDeployment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteComponent", func(t *testing.T) {
        input := &greengrassv2.DeleteComponentInput{}
        output := &greengrassv2.DeleteComponentOutput{}

        mockClient.On("DeleteComponent", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteComponent(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteCoreDevice", func(t *testing.T) {
        input := &greengrassv2.DeleteCoreDeviceInput{}
        output := &greengrassv2.DeleteCoreDeviceOutput{}

        mockClient.On("DeleteCoreDevice", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteCoreDevice(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteDeployment", func(t *testing.T) {
        input := &greengrassv2.DeleteDeploymentInput{}
        output := &greengrassv2.DeleteDeploymentOutput{}

        mockClient.On("DeleteDeployment", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteDeployment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeComponent", func(t *testing.T) {
        input := &greengrassv2.DescribeComponentInput{}
        output := &greengrassv2.DescribeComponentOutput{}

        mockClient.On("DescribeComponent", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeComponent(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateServiceRoleFromAccount", func(t *testing.T) {
        input := &greengrassv2.DisassociateServiceRoleFromAccountInput{}
        output := &greengrassv2.DisassociateServiceRoleFromAccountOutput{}

        mockClient.On("DisassociateServiceRoleFromAccount", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateServiceRoleFromAccount(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetComponent", func(t *testing.T) {
        input := &greengrassv2.GetComponentInput{}
        output := &greengrassv2.GetComponentOutput{}

        mockClient.On("GetComponent", ctx, input).Return(output, nil)

        result, err := mockClient.GetComponent(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetComponentVersionArtifact", func(t *testing.T) {
        input := &greengrassv2.GetComponentVersionArtifactInput{}
        output := &greengrassv2.GetComponentVersionArtifactOutput{}

        mockClient.On("GetComponentVersionArtifact", ctx, input).Return(output, nil)

        result, err := mockClient.GetComponentVersionArtifact(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetConnectivityInfo", func(t *testing.T) {
        input := &greengrassv2.GetConnectivityInfoInput{}
        output := &greengrassv2.GetConnectivityInfoOutput{}

        mockClient.On("GetConnectivityInfo", ctx, input).Return(output, nil)

        result, err := mockClient.GetConnectivityInfo(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetCoreDevice", func(t *testing.T) {
        input := &greengrassv2.GetCoreDeviceInput{}
        output := &greengrassv2.GetCoreDeviceOutput{}

        mockClient.On("GetCoreDevice", ctx, input).Return(output, nil)

        result, err := mockClient.GetCoreDevice(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDeployment", func(t *testing.T) {
        input := &greengrassv2.GetDeploymentInput{}
        output := &greengrassv2.GetDeploymentOutput{}

        mockClient.On("GetDeployment", ctx, input).Return(output, nil)

        result, err := mockClient.GetDeployment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetServiceRoleForAccount", func(t *testing.T) {
        input := &greengrassv2.GetServiceRoleForAccountInput{}
        output := &greengrassv2.GetServiceRoleForAccountOutput{}

        mockClient.On("GetServiceRoleForAccount", ctx, input).Return(output, nil)

        result, err := mockClient.GetServiceRoleForAccount(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListClientDevicesAssociatedWithCoreDevice", func(t *testing.T) {
        input := &greengrassv2.ListClientDevicesAssociatedWithCoreDeviceInput{}
        output := &greengrassv2.ListClientDevicesAssociatedWithCoreDeviceOutput{}

        mockClient.On("ListClientDevicesAssociatedWithCoreDevice", ctx, input).Return(output, nil)

        result, err := mockClient.ListClientDevicesAssociatedWithCoreDevice(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListComponentVersions", func(t *testing.T) {
        input := &greengrassv2.ListComponentVersionsInput{}
        output := &greengrassv2.ListComponentVersionsOutput{}

        mockClient.On("ListComponentVersions", ctx, input).Return(output, nil)

        result, err := mockClient.ListComponentVersions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListComponents", func(t *testing.T) {
        input := &greengrassv2.ListComponentsInput{}
        output := &greengrassv2.ListComponentsOutput{}

        mockClient.On("ListComponents", ctx, input).Return(output, nil)

        result, err := mockClient.ListComponents(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListCoreDevices", func(t *testing.T) {
        input := &greengrassv2.ListCoreDevicesInput{}
        output := &greengrassv2.ListCoreDevicesOutput{}

        mockClient.On("ListCoreDevices", ctx, input).Return(output, nil)

        result, err := mockClient.ListCoreDevices(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDeployments", func(t *testing.T) {
        input := &greengrassv2.ListDeploymentsInput{}
        output := &greengrassv2.ListDeploymentsOutput{}

        mockClient.On("ListDeployments", ctx, input).Return(output, nil)

        result, err := mockClient.ListDeployments(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListEffectiveDeployments", func(t *testing.T) {
        input := &greengrassv2.ListEffectiveDeploymentsInput{}
        output := &greengrassv2.ListEffectiveDeploymentsOutput{}

        mockClient.On("ListEffectiveDeployments", ctx, input).Return(output, nil)

        result, err := mockClient.ListEffectiveDeployments(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListInstalledComponents", func(t *testing.T) {
        input := &greengrassv2.ListInstalledComponentsInput{}
        output := &greengrassv2.ListInstalledComponentsOutput{}

        mockClient.On("ListInstalledComponents", ctx, input).Return(output, nil)

        result, err := mockClient.ListInstalledComponents(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &greengrassv2.ListTagsForResourceInput{}
        output := &greengrassv2.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestResolveComponentCandidates", func(t *testing.T) {
        input := &greengrassv2.ResolveComponentCandidatesInput{}
        output := &greengrassv2.ResolveComponentCandidatesOutput{}

        mockClient.On("ResolveComponentCandidates", ctx, input).Return(output, nil)

        result, err := mockClient.ResolveComponentCandidates(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &greengrassv2.TagResourceInput{}
        output := &greengrassv2.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &greengrassv2.UntagResourceInput{}
        output := &greengrassv2.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateConnectivityInfo", func(t *testing.T) {
        input := &greengrassv2.UpdateConnectivityInfoInput{}
        output := &greengrassv2.UpdateConnectivityInfoOutput{}

        mockClient.On("UpdateConnectivityInfo", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateConnectivityInfo(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
