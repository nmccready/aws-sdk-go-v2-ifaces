package opsworkscm_test

// tests for the opsworkscm service interface for this ../../../service/opsworkscm/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/opsworkscm"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/opsworkscm/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/opsworkscm/opsworkscm_iface"
	"github.com/stretchr/testify/assert"
)

func TestOpsworkscmServiceCanBeMocked(t *testing.T) {
	var iface opsworkscm_iface.IClient
	iface = &opsworkscm.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := opsworkscm.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateNode", func(t *testing.T) {
        input := &opsworkscm.AssociateNodeInput{}
        output := &opsworkscm.AssociateNodeOutput{}

        mockClient.On("AssociateNode", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateNode(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateBackup", func(t *testing.T) {
        input := &opsworkscm.CreateBackupInput{}
        output := &opsworkscm.CreateBackupOutput{}

        mockClient.On("CreateBackup", ctx, input).Return(output, nil)

        result, err := mockClient.CreateBackup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateServer", func(t *testing.T) {
        input := &opsworkscm.CreateServerInput{}
        output := &opsworkscm.CreateServerOutput{}

        mockClient.On("CreateServer", ctx, input).Return(output, nil)

        result, err := mockClient.CreateServer(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteBackup", func(t *testing.T) {
        input := &opsworkscm.DeleteBackupInput{}
        output := &opsworkscm.DeleteBackupOutput{}

        mockClient.On("DeleteBackup", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteBackup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteServer", func(t *testing.T) {
        input := &opsworkscm.DeleteServerInput{}
        output := &opsworkscm.DeleteServerOutput{}

        mockClient.On("DeleteServer", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteServer(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeAccountAttributes", func(t *testing.T) {
        input := &opsworkscm.DescribeAccountAttributesInput{}
        output := &opsworkscm.DescribeAccountAttributesOutput{}

        mockClient.On("DescribeAccountAttributes", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeAccountAttributes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeBackups", func(t *testing.T) {
        input := &opsworkscm.DescribeBackupsInput{}
        output := &opsworkscm.DescribeBackupsOutput{}

        mockClient.On("DescribeBackups", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeBackups(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeEvents", func(t *testing.T) {
        input := &opsworkscm.DescribeEventsInput{}
        output := &opsworkscm.DescribeEventsOutput{}

        mockClient.On("DescribeEvents", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeEvents(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeNodeAssociationStatus", func(t *testing.T) {
        input := &opsworkscm.DescribeNodeAssociationStatusInput{}
        output := &opsworkscm.DescribeNodeAssociationStatusOutput{}

        mockClient.On("DescribeNodeAssociationStatus", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeNodeAssociationStatus(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeServers", func(t *testing.T) {
        input := &opsworkscm.DescribeServersInput{}
        output := &opsworkscm.DescribeServersOutput{}

        mockClient.On("DescribeServers", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeServers(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateNode", func(t *testing.T) {
        input := &opsworkscm.DisassociateNodeInput{}
        output := &opsworkscm.DisassociateNodeOutput{}

        mockClient.On("DisassociateNode", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateNode(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestExportServerEngineAttribute", func(t *testing.T) {
        input := &opsworkscm.ExportServerEngineAttributeInput{}
        output := &opsworkscm.ExportServerEngineAttributeOutput{}

        mockClient.On("ExportServerEngineAttribute", ctx, input).Return(output, nil)

        result, err := mockClient.ExportServerEngineAttribute(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &opsworkscm.ListTagsForResourceInput{}
        output := &opsworkscm.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRestoreServer", func(t *testing.T) {
        input := &opsworkscm.RestoreServerInput{}
        output := &opsworkscm.RestoreServerOutput{}

        mockClient.On("RestoreServer", ctx, input).Return(output, nil)

        result, err := mockClient.RestoreServer(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartMaintenance", func(t *testing.T) {
        input := &opsworkscm.StartMaintenanceInput{}
        output := &opsworkscm.StartMaintenanceOutput{}

        mockClient.On("StartMaintenance", ctx, input).Return(output, nil)

        result, err := mockClient.StartMaintenance(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &opsworkscm.TagResourceInput{}
        output := &opsworkscm.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &opsworkscm.UntagResourceInput{}
        output := &opsworkscm.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateServer", func(t *testing.T) {
        input := &opsworkscm.UpdateServerInput{}
        output := &opsworkscm.UpdateServerOutput{}

        mockClient.On("UpdateServer", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateServer(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateServerEngineAttributes", func(t *testing.T) {
        input := &opsworkscm.UpdateServerEngineAttributesInput{}
        output := &opsworkscm.UpdateServerEngineAttributesOutput{}

        mockClient.On("UpdateServerEngineAttributes", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateServerEngineAttributes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
