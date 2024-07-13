package timestreaminfluxdb_test

// tests for the timestreaminfluxdb service interface for this ../../../service/timestreaminfluxdb/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/timestreaminfluxdb"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/timestreaminfluxdb/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/timestreaminfluxdb/timestreaminfluxdb_iface"
	"github.com/stretchr/testify/assert"
)

func TestTimestreaminfluxdbServiceCanBeMocked(t *testing.T) {
	var iface timestreaminfluxdb_iface.IClient
	iface = &timestreaminfluxdb.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := timestreaminfluxdb.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateDbInstance", func(t *testing.T) {
        input := &timestreaminfluxdb.CreateDbInstanceInput{}
        output := &timestreaminfluxdb.CreateDbInstanceOutput{}

        mockClient.On("CreateDbInstance", ctx, input).Return(output, nil)

        result, err := mockClient.CreateDbInstance(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateDbParameterGroup", func(t *testing.T) {
        input := &timestreaminfluxdb.CreateDbParameterGroupInput{}
        output := &timestreaminfluxdb.CreateDbParameterGroupOutput{}

        mockClient.On("CreateDbParameterGroup", ctx, input).Return(output, nil)

        result, err := mockClient.CreateDbParameterGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteDbInstance", func(t *testing.T) {
        input := &timestreaminfluxdb.DeleteDbInstanceInput{}
        output := &timestreaminfluxdb.DeleteDbInstanceOutput{}

        mockClient.On("DeleteDbInstance", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteDbInstance(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDbInstance", func(t *testing.T) {
        input := &timestreaminfluxdb.GetDbInstanceInput{}
        output := &timestreaminfluxdb.GetDbInstanceOutput{}

        mockClient.On("GetDbInstance", ctx, input).Return(output, nil)

        result, err := mockClient.GetDbInstance(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDbParameterGroup", func(t *testing.T) {
        input := &timestreaminfluxdb.GetDbParameterGroupInput{}
        output := &timestreaminfluxdb.GetDbParameterGroupOutput{}

        mockClient.On("GetDbParameterGroup", ctx, input).Return(output, nil)

        result, err := mockClient.GetDbParameterGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDbInstances", func(t *testing.T) {
        input := &timestreaminfluxdb.ListDbInstancesInput{}
        output := &timestreaminfluxdb.ListDbInstancesOutput{}

        mockClient.On("ListDbInstances", ctx, input).Return(output, nil)

        result, err := mockClient.ListDbInstances(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDbParameterGroups", func(t *testing.T) {
        input := &timestreaminfluxdb.ListDbParameterGroupsInput{}
        output := &timestreaminfluxdb.ListDbParameterGroupsOutput{}

        mockClient.On("ListDbParameterGroups", ctx, input).Return(output, nil)

        result, err := mockClient.ListDbParameterGroups(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &timestreaminfluxdb.ListTagsForResourceInput{}
        output := &timestreaminfluxdb.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &timestreaminfluxdb.TagResourceInput{}
        output := &timestreaminfluxdb.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &timestreaminfluxdb.UntagResourceInput{}
        output := &timestreaminfluxdb.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateDbInstance", func(t *testing.T) {
        input := &timestreaminfluxdb.UpdateDbInstanceInput{}
        output := &timestreaminfluxdb.UpdateDbInstanceOutput{}

        mockClient.On("UpdateDbInstance", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateDbInstance(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
