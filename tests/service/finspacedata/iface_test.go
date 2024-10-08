// Code generated by ./scripts/gen_tests.sh DO NOT EDIT.
package finspacedata_test

// tests for the finspacedata service interface for 
// this ../../../service/finspacedata/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/finspacedata"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/finspacedata/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/finspacedata/finspacedata_iface"
	"github.com/stretchr/testify/assert"
)

func TestFinspacedataServiceCanBeMocked(t *testing.T) {
	var iface finspacedata_iface.IClient
	iface = &finspacedata.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := finspacedata.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateUserToPermissionGroup", func(t *testing.T) {
        input := &finspacedata.AssociateUserToPermissionGroupInput{}
        output := &finspacedata.AssociateUserToPermissionGroupOutput{}

        mockClient.On("AssociateUserToPermissionGroup", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateUserToPermissionGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateChangeset", func(t *testing.T) {
        input := &finspacedata.CreateChangesetInput{}
        output := &finspacedata.CreateChangesetOutput{}

        mockClient.On("CreateChangeset", ctx, input).Return(output, nil)

        result, err := mockClient.CreateChangeset(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateDataView", func(t *testing.T) {
        input := &finspacedata.CreateDataViewInput{}
        output := &finspacedata.CreateDataViewOutput{}

        mockClient.On("CreateDataView", ctx, input).Return(output, nil)

        result, err := mockClient.CreateDataView(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateDataset", func(t *testing.T) {
        input := &finspacedata.CreateDatasetInput{}
        output := &finspacedata.CreateDatasetOutput{}

        mockClient.On("CreateDataset", ctx, input).Return(output, nil)

        result, err := mockClient.CreateDataset(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreatePermissionGroup", func(t *testing.T) {
        input := &finspacedata.CreatePermissionGroupInput{}
        output := &finspacedata.CreatePermissionGroupOutput{}

        mockClient.On("CreatePermissionGroup", ctx, input).Return(output, nil)

        result, err := mockClient.CreatePermissionGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateUser", func(t *testing.T) {
        input := &finspacedata.CreateUserInput{}
        output := &finspacedata.CreateUserOutput{}

        mockClient.On("CreateUser", ctx, input).Return(output, nil)

        result, err := mockClient.CreateUser(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteDataset", func(t *testing.T) {
        input := &finspacedata.DeleteDatasetInput{}
        output := &finspacedata.DeleteDatasetOutput{}

        mockClient.On("DeleteDataset", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteDataset(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeletePermissionGroup", func(t *testing.T) {
        input := &finspacedata.DeletePermissionGroupInput{}
        output := &finspacedata.DeletePermissionGroupOutput{}

        mockClient.On("DeletePermissionGroup", ctx, input).Return(output, nil)

        result, err := mockClient.DeletePermissionGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisableUser", func(t *testing.T) {
        input := &finspacedata.DisableUserInput{}
        output := &finspacedata.DisableUserOutput{}

        mockClient.On("DisableUser", ctx, input).Return(output, nil)

        result, err := mockClient.DisableUser(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateUserFromPermissionGroup", func(t *testing.T) {
        input := &finspacedata.DisassociateUserFromPermissionGroupInput{}
        output := &finspacedata.DisassociateUserFromPermissionGroupOutput{}

        mockClient.On("DisassociateUserFromPermissionGroup", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateUserFromPermissionGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestEnableUser", func(t *testing.T) {
        input := &finspacedata.EnableUserInput{}
        output := &finspacedata.EnableUserOutput{}

        mockClient.On("EnableUser", ctx, input).Return(output, nil)

        result, err := mockClient.EnableUser(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetChangeset", func(t *testing.T) {
        input := &finspacedata.GetChangesetInput{}
        output := &finspacedata.GetChangesetOutput{}

        mockClient.On("GetChangeset", ctx, input).Return(output, nil)

        result, err := mockClient.GetChangeset(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDataView", func(t *testing.T) {
        input := &finspacedata.GetDataViewInput{}
        output := &finspacedata.GetDataViewOutput{}

        mockClient.On("GetDataView", ctx, input).Return(output, nil)

        result, err := mockClient.GetDataView(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDataset", func(t *testing.T) {
        input := &finspacedata.GetDatasetInput{}
        output := &finspacedata.GetDatasetOutput{}

        mockClient.On("GetDataset", ctx, input).Return(output, nil)

        result, err := mockClient.GetDataset(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetExternalDataViewAccessDetails", func(t *testing.T) {
        input := &finspacedata.GetExternalDataViewAccessDetailsInput{}
        output := &finspacedata.GetExternalDataViewAccessDetailsOutput{}

        mockClient.On("GetExternalDataViewAccessDetails", ctx, input).Return(output, nil)

        result, err := mockClient.GetExternalDataViewAccessDetails(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetPermissionGroup", func(t *testing.T) {
        input := &finspacedata.GetPermissionGroupInput{}
        output := &finspacedata.GetPermissionGroupOutput{}

        mockClient.On("GetPermissionGroup", ctx, input).Return(output, nil)

        result, err := mockClient.GetPermissionGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetProgrammaticAccessCredentials", func(t *testing.T) {
        input := &finspacedata.GetProgrammaticAccessCredentialsInput{}
        output := &finspacedata.GetProgrammaticAccessCredentialsOutput{}

        mockClient.On("GetProgrammaticAccessCredentials", ctx, input).Return(output, nil)

        result, err := mockClient.GetProgrammaticAccessCredentials(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetUser", func(t *testing.T) {
        input := &finspacedata.GetUserInput{}
        output := &finspacedata.GetUserOutput{}

        mockClient.On("GetUser", ctx, input).Return(output, nil)

        result, err := mockClient.GetUser(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetWorkingLocation", func(t *testing.T) {
        input := &finspacedata.GetWorkingLocationInput{}
        output := &finspacedata.GetWorkingLocationOutput{}

        mockClient.On("GetWorkingLocation", ctx, input).Return(output, nil)

        result, err := mockClient.GetWorkingLocation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListChangesets", func(t *testing.T) {
        input := &finspacedata.ListChangesetsInput{}
        output := &finspacedata.ListChangesetsOutput{}

        mockClient.On("ListChangesets", ctx, input).Return(output, nil)

        result, err := mockClient.ListChangesets(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDataViews", func(t *testing.T) {
        input := &finspacedata.ListDataViewsInput{}
        output := &finspacedata.ListDataViewsOutput{}

        mockClient.On("ListDataViews", ctx, input).Return(output, nil)

        result, err := mockClient.ListDataViews(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDatasets", func(t *testing.T) {
        input := &finspacedata.ListDatasetsInput{}
        output := &finspacedata.ListDatasetsOutput{}

        mockClient.On("ListDatasets", ctx, input).Return(output, nil)

        result, err := mockClient.ListDatasets(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListPermissionGroups", func(t *testing.T) {
        input := &finspacedata.ListPermissionGroupsInput{}
        output := &finspacedata.ListPermissionGroupsOutput{}

        mockClient.On("ListPermissionGroups", ctx, input).Return(output, nil)

        result, err := mockClient.ListPermissionGroups(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListPermissionGroupsByUser", func(t *testing.T) {
        input := &finspacedata.ListPermissionGroupsByUserInput{}
        output := &finspacedata.ListPermissionGroupsByUserOutput{}

        mockClient.On("ListPermissionGroupsByUser", ctx, input).Return(output, nil)

        result, err := mockClient.ListPermissionGroupsByUser(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListUsers", func(t *testing.T) {
        input := &finspacedata.ListUsersInput{}
        output := &finspacedata.ListUsersOutput{}

        mockClient.On("ListUsers", ctx, input).Return(output, nil)

        result, err := mockClient.ListUsers(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListUsersByPermissionGroup", func(t *testing.T) {
        input := &finspacedata.ListUsersByPermissionGroupInput{}
        output := &finspacedata.ListUsersByPermissionGroupOutput{}

        mockClient.On("ListUsersByPermissionGroup", ctx, input).Return(output, nil)

        result, err := mockClient.ListUsersByPermissionGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestResetUserPassword", func(t *testing.T) {
        input := &finspacedata.ResetUserPasswordInput{}
        output := &finspacedata.ResetUserPasswordOutput{}

        mockClient.On("ResetUserPassword", ctx, input).Return(output, nil)

        result, err := mockClient.ResetUserPassword(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateChangeset", func(t *testing.T) {
        input := &finspacedata.UpdateChangesetInput{}
        output := &finspacedata.UpdateChangesetOutput{}

        mockClient.On("UpdateChangeset", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateChangeset(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateDataset", func(t *testing.T) {
        input := &finspacedata.UpdateDatasetInput{}
        output := &finspacedata.UpdateDatasetOutput{}

        mockClient.On("UpdateDataset", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateDataset(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdatePermissionGroup", func(t *testing.T) {
        input := &finspacedata.UpdatePermissionGroupInput{}
        output := &finspacedata.UpdatePermissionGroupOutput{}

        mockClient.On("UpdatePermissionGroup", ctx, input).Return(output, nil)

        result, err := mockClient.UpdatePermissionGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateUser", func(t *testing.T) {
        input := &finspacedata.UpdateUserInput{}
        output := &finspacedata.UpdateUserOutput{}

        mockClient.On("UpdateUser", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateUser(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
