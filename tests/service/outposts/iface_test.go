package outposts_test

// tests for the outposts service interface for this ../../../service/outposts/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/outposts"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/outposts/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/outposts/outposts_iface"
	"github.com/stretchr/testify/assert"
)

func TestOutpostsServiceCanBeMocked(t *testing.T) {
	var iface outposts_iface.IClient
	iface = &outposts.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := outposts.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCancelCapacityTask", func(t *testing.T) {
        input := &outposts.CancelCapacityTaskInput{}
        output := &outposts.CancelCapacityTaskOutput{}

        mockClient.On("CancelCapacityTask", ctx, input).Return(output, nil)

        result, err := mockClient.CancelCapacityTask(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCancelOrder", func(t *testing.T) {
        input := &outposts.CancelOrderInput{}
        output := &outposts.CancelOrderOutput{}

        mockClient.On("CancelOrder", ctx, input).Return(output, nil)

        result, err := mockClient.CancelOrder(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateOrder", func(t *testing.T) {
        input := &outposts.CreateOrderInput{}
        output := &outposts.CreateOrderOutput{}

        mockClient.On("CreateOrder", ctx, input).Return(output, nil)

        result, err := mockClient.CreateOrder(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateOutpost", func(t *testing.T) {
        input := &outposts.CreateOutpostInput{}
        output := &outposts.CreateOutpostOutput{}

        mockClient.On("CreateOutpost", ctx, input).Return(output, nil)

        result, err := mockClient.CreateOutpost(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateSite", func(t *testing.T) {
        input := &outposts.CreateSiteInput{}
        output := &outposts.CreateSiteOutput{}

        mockClient.On("CreateSite", ctx, input).Return(output, nil)

        result, err := mockClient.CreateSite(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteOutpost", func(t *testing.T) {
        input := &outposts.DeleteOutpostInput{}
        output := &outposts.DeleteOutpostOutput{}

        mockClient.On("DeleteOutpost", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteOutpost(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteSite", func(t *testing.T) {
        input := &outposts.DeleteSiteInput{}
        output := &outposts.DeleteSiteOutput{}

        mockClient.On("DeleteSite", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteSite(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetCapacityTask", func(t *testing.T) {
        input := &outposts.GetCapacityTaskInput{}
        output := &outposts.GetCapacityTaskOutput{}

        mockClient.On("GetCapacityTask", ctx, input).Return(output, nil)

        result, err := mockClient.GetCapacityTask(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetCatalogItem", func(t *testing.T) {
        input := &outposts.GetCatalogItemInput{}
        output := &outposts.GetCatalogItemOutput{}

        mockClient.On("GetCatalogItem", ctx, input).Return(output, nil)

        result, err := mockClient.GetCatalogItem(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetConnection", func(t *testing.T) {
        input := &outposts.GetConnectionInput{}
        output := &outposts.GetConnectionOutput{}

        mockClient.On("GetConnection", ctx, input).Return(output, nil)

        result, err := mockClient.GetConnection(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetOrder", func(t *testing.T) {
        input := &outposts.GetOrderInput{}
        output := &outposts.GetOrderOutput{}

        mockClient.On("GetOrder", ctx, input).Return(output, nil)

        result, err := mockClient.GetOrder(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetOutpost", func(t *testing.T) {
        input := &outposts.GetOutpostInput{}
        output := &outposts.GetOutpostOutput{}

        mockClient.On("GetOutpost", ctx, input).Return(output, nil)

        result, err := mockClient.GetOutpost(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetOutpostInstanceTypes", func(t *testing.T) {
        input := &outposts.GetOutpostInstanceTypesInput{}
        output := &outposts.GetOutpostInstanceTypesOutput{}

        mockClient.On("GetOutpostInstanceTypes", ctx, input).Return(output, nil)

        result, err := mockClient.GetOutpostInstanceTypes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetOutpostSupportedInstanceTypes", func(t *testing.T) {
        input := &outposts.GetOutpostSupportedInstanceTypesInput{}
        output := &outposts.GetOutpostSupportedInstanceTypesOutput{}

        mockClient.On("GetOutpostSupportedInstanceTypes", ctx, input).Return(output, nil)

        result, err := mockClient.GetOutpostSupportedInstanceTypes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSite", func(t *testing.T) {
        input := &outposts.GetSiteInput{}
        output := &outposts.GetSiteOutput{}

        mockClient.On("GetSite", ctx, input).Return(output, nil)

        result, err := mockClient.GetSite(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSiteAddress", func(t *testing.T) {
        input := &outposts.GetSiteAddressInput{}
        output := &outposts.GetSiteAddressOutput{}

        mockClient.On("GetSiteAddress", ctx, input).Return(output, nil)

        result, err := mockClient.GetSiteAddress(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAssets", func(t *testing.T) {
        input := &outposts.ListAssetsInput{}
        output := &outposts.ListAssetsOutput{}

        mockClient.On("ListAssets", ctx, input).Return(output, nil)

        result, err := mockClient.ListAssets(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListCapacityTasks", func(t *testing.T) {
        input := &outposts.ListCapacityTasksInput{}
        output := &outposts.ListCapacityTasksOutput{}

        mockClient.On("ListCapacityTasks", ctx, input).Return(output, nil)

        result, err := mockClient.ListCapacityTasks(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListCatalogItems", func(t *testing.T) {
        input := &outposts.ListCatalogItemsInput{}
        output := &outposts.ListCatalogItemsOutput{}

        mockClient.On("ListCatalogItems", ctx, input).Return(output, nil)

        result, err := mockClient.ListCatalogItems(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListOrders", func(t *testing.T) {
        input := &outposts.ListOrdersInput{}
        output := &outposts.ListOrdersOutput{}

        mockClient.On("ListOrders", ctx, input).Return(output, nil)

        result, err := mockClient.ListOrders(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListOutposts", func(t *testing.T) {
        input := &outposts.ListOutpostsInput{}
        output := &outposts.ListOutpostsOutput{}

        mockClient.On("ListOutposts", ctx, input).Return(output, nil)

        result, err := mockClient.ListOutposts(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListSites", func(t *testing.T) {
        input := &outposts.ListSitesInput{}
        output := &outposts.ListSitesOutput{}

        mockClient.On("ListSites", ctx, input).Return(output, nil)

        result, err := mockClient.ListSites(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &outposts.ListTagsForResourceInput{}
        output := &outposts.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartCapacityTask", func(t *testing.T) {
        input := &outposts.StartCapacityTaskInput{}
        output := &outposts.StartCapacityTaskOutput{}

        mockClient.On("StartCapacityTask", ctx, input).Return(output, nil)

        result, err := mockClient.StartCapacityTask(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartConnection", func(t *testing.T) {
        input := &outposts.StartConnectionInput{}
        output := &outposts.StartConnectionOutput{}

        mockClient.On("StartConnection", ctx, input).Return(output, nil)

        result, err := mockClient.StartConnection(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &outposts.TagResourceInput{}
        output := &outposts.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &outposts.UntagResourceInput{}
        output := &outposts.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateOutpost", func(t *testing.T) {
        input := &outposts.UpdateOutpostInput{}
        output := &outposts.UpdateOutpostOutput{}

        mockClient.On("UpdateOutpost", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateOutpost(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateSite", func(t *testing.T) {
        input := &outposts.UpdateSiteInput{}
        output := &outposts.UpdateSiteOutput{}

        mockClient.On("UpdateSite", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateSite(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateSiteAddress", func(t *testing.T) {
        input := &outposts.UpdateSiteAddressInput{}
        output := &outposts.UpdateSiteAddressOutput{}

        mockClient.On("UpdateSiteAddress", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateSiteAddress(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateSiteRackPhysicalProperties", func(t *testing.T) {
        input := &outposts.UpdateSiteRackPhysicalPropertiesInput{}
        output := &outposts.UpdateSiteRackPhysicalPropertiesOutput{}

        mockClient.On("UpdateSiteRackPhysicalProperties", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateSiteRackPhysicalProperties(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
