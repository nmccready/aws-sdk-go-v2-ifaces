// Code generated by ./scripts/gen_tests.sh DO NOT EDIT.
package iotfleetwise_test

// tests for the iotfleetwise service interface for 
// this ../../../service/iotfleetwise/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/iotfleetwise"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/iotfleetwise/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/iotfleetwise/iotfleetwise_iface"
	"github.com/stretchr/testify/assert"
)

func TestIotfleetwiseServiceCanBeMocked(t *testing.T) {
	var iface iotfleetwise_iface.IClient
	iface = &iotfleetwise.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := iotfleetwise.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateVehicleFleet", func(t *testing.T) {
        input := &iotfleetwise.AssociateVehicleFleetInput{}
        output := &iotfleetwise.AssociateVehicleFleetOutput{}

        mockClient.On("AssociateVehicleFleet", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateVehicleFleet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchCreateVehicle", func(t *testing.T) {
        input := &iotfleetwise.BatchCreateVehicleInput{}
        output := &iotfleetwise.BatchCreateVehicleOutput{}

        mockClient.On("BatchCreateVehicle", ctx, input).Return(output, nil)

        result, err := mockClient.BatchCreateVehicle(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchUpdateVehicle", func(t *testing.T) {
        input := &iotfleetwise.BatchUpdateVehicleInput{}
        output := &iotfleetwise.BatchUpdateVehicleOutput{}

        mockClient.On("BatchUpdateVehicle", ctx, input).Return(output, nil)

        result, err := mockClient.BatchUpdateVehicle(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateCampaign", func(t *testing.T) {
        input := &iotfleetwise.CreateCampaignInput{}
        output := &iotfleetwise.CreateCampaignOutput{}

        mockClient.On("CreateCampaign", ctx, input).Return(output, nil)

        result, err := mockClient.CreateCampaign(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateDecoderManifest", func(t *testing.T) {
        input := &iotfleetwise.CreateDecoderManifestInput{}
        output := &iotfleetwise.CreateDecoderManifestOutput{}

        mockClient.On("CreateDecoderManifest", ctx, input).Return(output, nil)

        result, err := mockClient.CreateDecoderManifest(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateFleet", func(t *testing.T) {
        input := &iotfleetwise.CreateFleetInput{}
        output := &iotfleetwise.CreateFleetOutput{}

        mockClient.On("CreateFleet", ctx, input).Return(output, nil)

        result, err := mockClient.CreateFleet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateModelManifest", func(t *testing.T) {
        input := &iotfleetwise.CreateModelManifestInput{}
        output := &iotfleetwise.CreateModelManifestOutput{}

        mockClient.On("CreateModelManifest", ctx, input).Return(output, nil)

        result, err := mockClient.CreateModelManifest(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateSignalCatalog", func(t *testing.T) {
        input := &iotfleetwise.CreateSignalCatalogInput{}
        output := &iotfleetwise.CreateSignalCatalogOutput{}

        mockClient.On("CreateSignalCatalog", ctx, input).Return(output, nil)

        result, err := mockClient.CreateSignalCatalog(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateVehicle", func(t *testing.T) {
        input := &iotfleetwise.CreateVehicleInput{}
        output := &iotfleetwise.CreateVehicleOutput{}

        mockClient.On("CreateVehicle", ctx, input).Return(output, nil)

        result, err := mockClient.CreateVehicle(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteCampaign", func(t *testing.T) {
        input := &iotfleetwise.DeleteCampaignInput{}
        output := &iotfleetwise.DeleteCampaignOutput{}

        mockClient.On("DeleteCampaign", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteCampaign(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteDecoderManifest", func(t *testing.T) {
        input := &iotfleetwise.DeleteDecoderManifestInput{}
        output := &iotfleetwise.DeleteDecoderManifestOutput{}

        mockClient.On("DeleteDecoderManifest", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteDecoderManifest(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteFleet", func(t *testing.T) {
        input := &iotfleetwise.DeleteFleetInput{}
        output := &iotfleetwise.DeleteFleetOutput{}

        mockClient.On("DeleteFleet", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteFleet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteModelManifest", func(t *testing.T) {
        input := &iotfleetwise.DeleteModelManifestInput{}
        output := &iotfleetwise.DeleteModelManifestOutput{}

        mockClient.On("DeleteModelManifest", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteModelManifest(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteSignalCatalog", func(t *testing.T) {
        input := &iotfleetwise.DeleteSignalCatalogInput{}
        output := &iotfleetwise.DeleteSignalCatalogOutput{}

        mockClient.On("DeleteSignalCatalog", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteSignalCatalog(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteVehicle", func(t *testing.T) {
        input := &iotfleetwise.DeleteVehicleInput{}
        output := &iotfleetwise.DeleteVehicleOutput{}

        mockClient.On("DeleteVehicle", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteVehicle(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateVehicleFleet", func(t *testing.T) {
        input := &iotfleetwise.DisassociateVehicleFleetInput{}
        output := &iotfleetwise.DisassociateVehicleFleetOutput{}

        mockClient.On("DisassociateVehicleFleet", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateVehicleFleet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetCampaign", func(t *testing.T) {
        input := &iotfleetwise.GetCampaignInput{}
        output := &iotfleetwise.GetCampaignOutput{}

        mockClient.On("GetCampaign", ctx, input).Return(output, nil)

        result, err := mockClient.GetCampaign(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDecoderManifest", func(t *testing.T) {
        input := &iotfleetwise.GetDecoderManifestInput{}
        output := &iotfleetwise.GetDecoderManifestOutput{}

        mockClient.On("GetDecoderManifest", ctx, input).Return(output, nil)

        result, err := mockClient.GetDecoderManifest(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetEncryptionConfiguration", func(t *testing.T) {
        input := &iotfleetwise.GetEncryptionConfigurationInput{}
        output := &iotfleetwise.GetEncryptionConfigurationOutput{}

        mockClient.On("GetEncryptionConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.GetEncryptionConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetFleet", func(t *testing.T) {
        input := &iotfleetwise.GetFleetInput{}
        output := &iotfleetwise.GetFleetOutput{}

        mockClient.On("GetFleet", ctx, input).Return(output, nil)

        result, err := mockClient.GetFleet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetLoggingOptions", func(t *testing.T) {
        input := &iotfleetwise.GetLoggingOptionsInput{}
        output := &iotfleetwise.GetLoggingOptionsOutput{}

        mockClient.On("GetLoggingOptions", ctx, input).Return(output, nil)

        result, err := mockClient.GetLoggingOptions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetModelManifest", func(t *testing.T) {
        input := &iotfleetwise.GetModelManifestInput{}
        output := &iotfleetwise.GetModelManifestOutput{}

        mockClient.On("GetModelManifest", ctx, input).Return(output, nil)

        result, err := mockClient.GetModelManifest(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetRegisterAccountStatus", func(t *testing.T) {
        input := &iotfleetwise.GetRegisterAccountStatusInput{}
        output := &iotfleetwise.GetRegisterAccountStatusOutput{}

        mockClient.On("GetRegisterAccountStatus", ctx, input).Return(output, nil)

        result, err := mockClient.GetRegisterAccountStatus(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSignalCatalog", func(t *testing.T) {
        input := &iotfleetwise.GetSignalCatalogInput{}
        output := &iotfleetwise.GetSignalCatalogOutput{}

        mockClient.On("GetSignalCatalog", ctx, input).Return(output, nil)

        result, err := mockClient.GetSignalCatalog(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetVehicle", func(t *testing.T) {
        input := &iotfleetwise.GetVehicleInput{}
        output := &iotfleetwise.GetVehicleOutput{}

        mockClient.On("GetVehicle", ctx, input).Return(output, nil)

        result, err := mockClient.GetVehicle(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetVehicleStatus", func(t *testing.T) {
        input := &iotfleetwise.GetVehicleStatusInput{}
        output := &iotfleetwise.GetVehicleStatusOutput{}

        mockClient.On("GetVehicleStatus", ctx, input).Return(output, nil)

        result, err := mockClient.GetVehicleStatus(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestImportDecoderManifest", func(t *testing.T) {
        input := &iotfleetwise.ImportDecoderManifestInput{}
        output := &iotfleetwise.ImportDecoderManifestOutput{}

        mockClient.On("ImportDecoderManifest", ctx, input).Return(output, nil)

        result, err := mockClient.ImportDecoderManifest(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestImportSignalCatalog", func(t *testing.T) {
        input := &iotfleetwise.ImportSignalCatalogInput{}
        output := &iotfleetwise.ImportSignalCatalogOutput{}

        mockClient.On("ImportSignalCatalog", ctx, input).Return(output, nil)

        result, err := mockClient.ImportSignalCatalog(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListCampaigns", func(t *testing.T) {
        input := &iotfleetwise.ListCampaignsInput{}
        output := &iotfleetwise.ListCampaignsOutput{}

        mockClient.On("ListCampaigns", ctx, input).Return(output, nil)

        result, err := mockClient.ListCampaigns(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDecoderManifestNetworkInterfaces", func(t *testing.T) {
        input := &iotfleetwise.ListDecoderManifestNetworkInterfacesInput{}
        output := &iotfleetwise.ListDecoderManifestNetworkInterfacesOutput{}

        mockClient.On("ListDecoderManifestNetworkInterfaces", ctx, input).Return(output, nil)

        result, err := mockClient.ListDecoderManifestNetworkInterfaces(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDecoderManifestSignals", func(t *testing.T) {
        input := &iotfleetwise.ListDecoderManifestSignalsInput{}
        output := &iotfleetwise.ListDecoderManifestSignalsOutput{}

        mockClient.On("ListDecoderManifestSignals", ctx, input).Return(output, nil)

        result, err := mockClient.ListDecoderManifestSignals(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDecoderManifests", func(t *testing.T) {
        input := &iotfleetwise.ListDecoderManifestsInput{}
        output := &iotfleetwise.ListDecoderManifestsOutput{}

        mockClient.On("ListDecoderManifests", ctx, input).Return(output, nil)

        result, err := mockClient.ListDecoderManifests(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListFleets", func(t *testing.T) {
        input := &iotfleetwise.ListFleetsInput{}
        output := &iotfleetwise.ListFleetsOutput{}

        mockClient.On("ListFleets", ctx, input).Return(output, nil)

        result, err := mockClient.ListFleets(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListFleetsForVehicle", func(t *testing.T) {
        input := &iotfleetwise.ListFleetsForVehicleInput{}
        output := &iotfleetwise.ListFleetsForVehicleOutput{}

        mockClient.On("ListFleetsForVehicle", ctx, input).Return(output, nil)

        result, err := mockClient.ListFleetsForVehicle(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListModelManifestNodes", func(t *testing.T) {
        input := &iotfleetwise.ListModelManifestNodesInput{}
        output := &iotfleetwise.ListModelManifestNodesOutput{}

        mockClient.On("ListModelManifestNodes", ctx, input).Return(output, nil)

        result, err := mockClient.ListModelManifestNodes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListModelManifests", func(t *testing.T) {
        input := &iotfleetwise.ListModelManifestsInput{}
        output := &iotfleetwise.ListModelManifestsOutput{}

        mockClient.On("ListModelManifests", ctx, input).Return(output, nil)

        result, err := mockClient.ListModelManifests(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListSignalCatalogNodes", func(t *testing.T) {
        input := &iotfleetwise.ListSignalCatalogNodesInput{}
        output := &iotfleetwise.ListSignalCatalogNodesOutput{}

        mockClient.On("ListSignalCatalogNodes", ctx, input).Return(output, nil)

        result, err := mockClient.ListSignalCatalogNodes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListSignalCatalogs", func(t *testing.T) {
        input := &iotfleetwise.ListSignalCatalogsInput{}
        output := &iotfleetwise.ListSignalCatalogsOutput{}

        mockClient.On("ListSignalCatalogs", ctx, input).Return(output, nil)

        result, err := mockClient.ListSignalCatalogs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &iotfleetwise.ListTagsForResourceInput{}
        output := &iotfleetwise.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListVehicles", func(t *testing.T) {
        input := &iotfleetwise.ListVehiclesInput{}
        output := &iotfleetwise.ListVehiclesOutput{}

        mockClient.On("ListVehicles", ctx, input).Return(output, nil)

        result, err := mockClient.ListVehicles(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListVehiclesInFleet", func(t *testing.T) {
        input := &iotfleetwise.ListVehiclesInFleetInput{}
        output := &iotfleetwise.ListVehiclesInFleetOutput{}

        mockClient.On("ListVehiclesInFleet", ctx, input).Return(output, nil)

        result, err := mockClient.ListVehiclesInFleet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutEncryptionConfiguration", func(t *testing.T) {
        input := &iotfleetwise.PutEncryptionConfigurationInput{}
        output := &iotfleetwise.PutEncryptionConfigurationOutput{}

        mockClient.On("PutEncryptionConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.PutEncryptionConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutLoggingOptions", func(t *testing.T) {
        input := &iotfleetwise.PutLoggingOptionsInput{}
        output := &iotfleetwise.PutLoggingOptionsOutput{}

        mockClient.On("PutLoggingOptions", ctx, input).Return(output, nil)

        result, err := mockClient.PutLoggingOptions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRegisterAccount", func(t *testing.T) {
        input := &iotfleetwise.RegisterAccountInput{}
        output := &iotfleetwise.RegisterAccountOutput{}

        mockClient.On("RegisterAccount", ctx, input).Return(output, nil)

        result, err := mockClient.RegisterAccount(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &iotfleetwise.TagResourceInput{}
        output := &iotfleetwise.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &iotfleetwise.UntagResourceInput{}
        output := &iotfleetwise.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateCampaign", func(t *testing.T) {
        input := &iotfleetwise.UpdateCampaignInput{}
        output := &iotfleetwise.UpdateCampaignOutput{}

        mockClient.On("UpdateCampaign", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateCampaign(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateDecoderManifest", func(t *testing.T) {
        input := &iotfleetwise.UpdateDecoderManifestInput{}
        output := &iotfleetwise.UpdateDecoderManifestOutput{}

        mockClient.On("UpdateDecoderManifest", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateDecoderManifest(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateFleet", func(t *testing.T) {
        input := &iotfleetwise.UpdateFleetInput{}
        output := &iotfleetwise.UpdateFleetOutput{}

        mockClient.On("UpdateFleet", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateFleet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateModelManifest", func(t *testing.T) {
        input := &iotfleetwise.UpdateModelManifestInput{}
        output := &iotfleetwise.UpdateModelManifestOutput{}

        mockClient.On("UpdateModelManifest", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateModelManifest(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateSignalCatalog", func(t *testing.T) {
        input := &iotfleetwise.UpdateSignalCatalogInput{}
        output := &iotfleetwise.UpdateSignalCatalogOutput{}

        mockClient.On("UpdateSignalCatalog", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateSignalCatalog(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateVehicle", func(t *testing.T) {
        input := &iotfleetwise.UpdateVehicleInput{}
        output := &iotfleetwise.UpdateVehicleOutput{}

        mockClient.On("UpdateVehicle", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateVehicle(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
