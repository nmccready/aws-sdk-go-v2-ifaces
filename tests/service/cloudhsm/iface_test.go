package cloudhsm_test

// tests for the cloudhsm service interface for this ../../../service/cloudhsm/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/cloudhsm"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/cloudhsm/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/cloudhsm/cloudhsm_iface"
	"github.com/stretchr/testify/assert"
)

func TestCloudhsmServiceCanBeMocked(t *testing.T) {
	var iface cloudhsm_iface.IClient
	iface = &cloudhsm.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := cloudhsm.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAddTagsToResource", func(t *testing.T) {
        input := &cloudhsm.AddTagsToResourceInput{}
        output := &cloudhsm.AddTagsToResourceOutput{}

        mockClient.On("AddTagsToResource", ctx, input).Return(output, nil)

        result, err := mockClient.AddTagsToResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateHapg", func(t *testing.T) {
        input := &cloudhsm.CreateHapgInput{}
        output := &cloudhsm.CreateHapgOutput{}

        mockClient.On("CreateHapg", ctx, input).Return(output, nil)

        result, err := mockClient.CreateHapg(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateHsm", func(t *testing.T) {
        input := &cloudhsm.CreateHsmInput{}
        output := &cloudhsm.CreateHsmOutput{}

        mockClient.On("CreateHsm", ctx, input).Return(output, nil)

        result, err := mockClient.CreateHsm(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateLunaClient", func(t *testing.T) {
        input := &cloudhsm.CreateLunaClientInput{}
        output := &cloudhsm.CreateLunaClientOutput{}

        mockClient.On("CreateLunaClient", ctx, input).Return(output, nil)

        result, err := mockClient.CreateLunaClient(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteHapg", func(t *testing.T) {
        input := &cloudhsm.DeleteHapgInput{}
        output := &cloudhsm.DeleteHapgOutput{}

        mockClient.On("DeleteHapg", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteHapg(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteHsm", func(t *testing.T) {
        input := &cloudhsm.DeleteHsmInput{}
        output := &cloudhsm.DeleteHsmOutput{}

        mockClient.On("DeleteHsm", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteHsm(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteLunaClient", func(t *testing.T) {
        input := &cloudhsm.DeleteLunaClientInput{}
        output := &cloudhsm.DeleteLunaClientOutput{}

        mockClient.On("DeleteLunaClient", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteLunaClient(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeHapg", func(t *testing.T) {
        input := &cloudhsm.DescribeHapgInput{}
        output := &cloudhsm.DescribeHapgOutput{}

        mockClient.On("DescribeHapg", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeHapg(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeHsm", func(t *testing.T) {
        input := &cloudhsm.DescribeHsmInput{}
        output := &cloudhsm.DescribeHsmOutput{}

        mockClient.On("DescribeHsm", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeHsm(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeLunaClient", func(t *testing.T) {
        input := &cloudhsm.DescribeLunaClientInput{}
        output := &cloudhsm.DescribeLunaClientOutput{}

        mockClient.On("DescribeLunaClient", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeLunaClient(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetConfig", func(t *testing.T) {
        input := &cloudhsm.GetConfigInput{}
        output := &cloudhsm.GetConfigOutput{}

        mockClient.On("GetConfig", ctx, input).Return(output, nil)

        result, err := mockClient.GetConfig(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAvailableZones", func(t *testing.T) {
        input := &cloudhsm.ListAvailableZonesInput{}
        output := &cloudhsm.ListAvailableZonesOutput{}

        mockClient.On("ListAvailableZones", ctx, input).Return(output, nil)

        result, err := mockClient.ListAvailableZones(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListHapgs", func(t *testing.T) {
        input := &cloudhsm.ListHapgsInput{}
        output := &cloudhsm.ListHapgsOutput{}

        mockClient.On("ListHapgs", ctx, input).Return(output, nil)

        result, err := mockClient.ListHapgs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListHsms", func(t *testing.T) {
        input := &cloudhsm.ListHsmsInput{}
        output := &cloudhsm.ListHsmsOutput{}

        mockClient.On("ListHsms", ctx, input).Return(output, nil)

        result, err := mockClient.ListHsms(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListLunaClients", func(t *testing.T) {
        input := &cloudhsm.ListLunaClientsInput{}
        output := &cloudhsm.ListLunaClientsOutput{}

        mockClient.On("ListLunaClients", ctx, input).Return(output, nil)

        result, err := mockClient.ListLunaClients(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &cloudhsm.ListTagsForResourceInput{}
        output := &cloudhsm.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyHapg", func(t *testing.T) {
        input := &cloudhsm.ModifyHapgInput{}
        output := &cloudhsm.ModifyHapgOutput{}

        mockClient.On("ModifyHapg", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyHapg(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyHsm", func(t *testing.T) {
        input := &cloudhsm.ModifyHsmInput{}
        output := &cloudhsm.ModifyHsmOutput{}

        mockClient.On("ModifyHsm", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyHsm(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyLunaClient", func(t *testing.T) {
        input := &cloudhsm.ModifyLunaClientInput{}
        output := &cloudhsm.ModifyLunaClientOutput{}

        mockClient.On("ModifyLunaClient", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyLunaClient(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRemoveTagsFromResource", func(t *testing.T) {
        input := &cloudhsm.RemoveTagsFromResourceInput{}
        output := &cloudhsm.RemoveTagsFromResourceOutput{}

        mockClient.On("RemoveTagsFromResource", ctx, input).Return(output, nil)

        result, err := mockClient.RemoveTagsFromResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
