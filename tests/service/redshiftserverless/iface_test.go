// Code generated by ./scripts/gen_tests.sh DO NOT EDIT.
package redshiftserverless_test

// tests for the redshiftserverless service interface for 
// this ../../../service/redshiftserverless/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/redshiftserverless"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/redshiftserverless/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/redshiftserverless/redshiftserverless_iface"
	"github.com/stretchr/testify/assert"
)

func TestRedshiftserverlessServiceCanBeMocked(t *testing.T) {
	var iface redshiftserverless_iface.IClient
	iface = &redshiftserverless.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := redshiftserverless.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestConvertRecoveryPointToSnapshot", func(t *testing.T) {
        input := &redshiftserverless.ConvertRecoveryPointToSnapshotInput{}
        output := &redshiftserverless.ConvertRecoveryPointToSnapshotOutput{}

        mockClient.On("ConvertRecoveryPointToSnapshot", ctx, input).Return(output, nil)

        result, err := mockClient.ConvertRecoveryPointToSnapshot(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateCustomDomainAssociation", func(t *testing.T) {
        input := &redshiftserverless.CreateCustomDomainAssociationInput{}
        output := &redshiftserverless.CreateCustomDomainAssociationOutput{}

        mockClient.On("CreateCustomDomainAssociation", ctx, input).Return(output, nil)

        result, err := mockClient.CreateCustomDomainAssociation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateEndpointAccess", func(t *testing.T) {
        input := &redshiftserverless.CreateEndpointAccessInput{}
        output := &redshiftserverless.CreateEndpointAccessOutput{}

        mockClient.On("CreateEndpointAccess", ctx, input).Return(output, nil)

        result, err := mockClient.CreateEndpointAccess(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateNamespace", func(t *testing.T) {
        input := &redshiftserverless.CreateNamespaceInput{}
        output := &redshiftserverless.CreateNamespaceOutput{}

        mockClient.On("CreateNamespace", ctx, input).Return(output, nil)

        result, err := mockClient.CreateNamespace(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateReservation", func(t *testing.T) {
        input := &redshiftserverless.CreateReservationInput{}
        output := &redshiftserverless.CreateReservationOutput{}

        mockClient.On("CreateReservation", ctx, input).Return(output, nil)

        result, err := mockClient.CreateReservation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateScheduledAction", func(t *testing.T) {
        input := &redshiftserverless.CreateScheduledActionInput{}
        output := &redshiftserverless.CreateScheduledActionOutput{}

        mockClient.On("CreateScheduledAction", ctx, input).Return(output, nil)

        result, err := mockClient.CreateScheduledAction(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateSnapshot", func(t *testing.T) {
        input := &redshiftserverless.CreateSnapshotInput{}
        output := &redshiftserverless.CreateSnapshotOutput{}

        mockClient.On("CreateSnapshot", ctx, input).Return(output, nil)

        result, err := mockClient.CreateSnapshot(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateSnapshotCopyConfiguration", func(t *testing.T) {
        input := &redshiftserverless.CreateSnapshotCopyConfigurationInput{}
        output := &redshiftserverless.CreateSnapshotCopyConfigurationOutput{}

        mockClient.On("CreateSnapshotCopyConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.CreateSnapshotCopyConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateUsageLimit", func(t *testing.T) {
        input := &redshiftserverless.CreateUsageLimitInput{}
        output := &redshiftserverless.CreateUsageLimitOutput{}

        mockClient.On("CreateUsageLimit", ctx, input).Return(output, nil)

        result, err := mockClient.CreateUsageLimit(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateWorkgroup", func(t *testing.T) {
        input := &redshiftserverless.CreateWorkgroupInput{}
        output := &redshiftserverless.CreateWorkgroupOutput{}

        mockClient.On("CreateWorkgroup", ctx, input).Return(output, nil)

        result, err := mockClient.CreateWorkgroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteCustomDomainAssociation", func(t *testing.T) {
        input := &redshiftserverless.DeleteCustomDomainAssociationInput{}
        output := &redshiftserverless.DeleteCustomDomainAssociationOutput{}

        mockClient.On("DeleteCustomDomainAssociation", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteCustomDomainAssociation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteEndpointAccess", func(t *testing.T) {
        input := &redshiftserverless.DeleteEndpointAccessInput{}
        output := &redshiftserverless.DeleteEndpointAccessOutput{}

        mockClient.On("DeleteEndpointAccess", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteEndpointAccess(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteNamespace", func(t *testing.T) {
        input := &redshiftserverless.DeleteNamespaceInput{}
        output := &redshiftserverless.DeleteNamespaceOutput{}

        mockClient.On("DeleteNamespace", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteNamespace(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteResourcePolicy", func(t *testing.T) {
        input := &redshiftserverless.DeleteResourcePolicyInput{}
        output := &redshiftserverless.DeleteResourcePolicyOutput{}

        mockClient.On("DeleteResourcePolicy", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteResourcePolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteScheduledAction", func(t *testing.T) {
        input := &redshiftserverless.DeleteScheduledActionInput{}
        output := &redshiftserverless.DeleteScheduledActionOutput{}

        mockClient.On("DeleteScheduledAction", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteScheduledAction(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteSnapshot", func(t *testing.T) {
        input := &redshiftserverless.DeleteSnapshotInput{}
        output := &redshiftserverless.DeleteSnapshotOutput{}

        mockClient.On("DeleteSnapshot", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteSnapshot(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteSnapshotCopyConfiguration", func(t *testing.T) {
        input := &redshiftserverless.DeleteSnapshotCopyConfigurationInput{}
        output := &redshiftserverless.DeleteSnapshotCopyConfigurationOutput{}

        mockClient.On("DeleteSnapshotCopyConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteSnapshotCopyConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteUsageLimit", func(t *testing.T) {
        input := &redshiftserverless.DeleteUsageLimitInput{}
        output := &redshiftserverless.DeleteUsageLimitOutput{}

        mockClient.On("DeleteUsageLimit", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteUsageLimit(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteWorkgroup", func(t *testing.T) {
        input := &redshiftserverless.DeleteWorkgroupInput{}
        output := &redshiftserverless.DeleteWorkgroupOutput{}

        mockClient.On("DeleteWorkgroup", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteWorkgroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetCredentials", func(t *testing.T) {
        input := &redshiftserverless.GetCredentialsInput{}
        output := &redshiftserverless.GetCredentialsOutput{}

        mockClient.On("GetCredentials", ctx, input).Return(output, nil)

        result, err := mockClient.GetCredentials(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetCustomDomainAssociation", func(t *testing.T) {
        input := &redshiftserverless.GetCustomDomainAssociationInput{}
        output := &redshiftserverless.GetCustomDomainAssociationOutput{}

        mockClient.On("GetCustomDomainAssociation", ctx, input).Return(output, nil)

        result, err := mockClient.GetCustomDomainAssociation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetEndpointAccess", func(t *testing.T) {
        input := &redshiftserverless.GetEndpointAccessInput{}
        output := &redshiftserverless.GetEndpointAccessOutput{}

        mockClient.On("GetEndpointAccess", ctx, input).Return(output, nil)

        result, err := mockClient.GetEndpointAccess(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetNamespace", func(t *testing.T) {
        input := &redshiftserverless.GetNamespaceInput{}
        output := &redshiftserverless.GetNamespaceOutput{}

        mockClient.On("GetNamespace", ctx, input).Return(output, nil)

        result, err := mockClient.GetNamespace(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetRecoveryPoint", func(t *testing.T) {
        input := &redshiftserverless.GetRecoveryPointInput{}
        output := &redshiftserverless.GetRecoveryPointOutput{}

        mockClient.On("GetRecoveryPoint", ctx, input).Return(output, nil)

        result, err := mockClient.GetRecoveryPoint(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetReservation", func(t *testing.T) {
        input := &redshiftserverless.GetReservationInput{}
        output := &redshiftserverless.GetReservationOutput{}

        mockClient.On("GetReservation", ctx, input).Return(output, nil)

        result, err := mockClient.GetReservation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetReservationOffering", func(t *testing.T) {
        input := &redshiftserverless.GetReservationOfferingInput{}
        output := &redshiftserverless.GetReservationOfferingOutput{}

        mockClient.On("GetReservationOffering", ctx, input).Return(output, nil)

        result, err := mockClient.GetReservationOffering(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetResourcePolicy", func(t *testing.T) {
        input := &redshiftserverless.GetResourcePolicyInput{}
        output := &redshiftserverless.GetResourcePolicyOutput{}

        mockClient.On("GetResourcePolicy", ctx, input).Return(output, nil)

        result, err := mockClient.GetResourcePolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetScheduledAction", func(t *testing.T) {
        input := &redshiftserverless.GetScheduledActionInput{}
        output := &redshiftserverless.GetScheduledActionOutput{}

        mockClient.On("GetScheduledAction", ctx, input).Return(output, nil)

        result, err := mockClient.GetScheduledAction(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSnapshot", func(t *testing.T) {
        input := &redshiftserverless.GetSnapshotInput{}
        output := &redshiftserverless.GetSnapshotOutput{}

        mockClient.On("GetSnapshot", ctx, input).Return(output, nil)

        result, err := mockClient.GetSnapshot(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetTableRestoreStatus", func(t *testing.T) {
        input := &redshiftserverless.GetTableRestoreStatusInput{}
        output := &redshiftserverless.GetTableRestoreStatusOutput{}

        mockClient.On("GetTableRestoreStatus", ctx, input).Return(output, nil)

        result, err := mockClient.GetTableRestoreStatus(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetTrack", func(t *testing.T) {
        input := &redshiftserverless.GetTrackInput{}
        output := &redshiftserverless.GetTrackOutput{}

        mockClient.On("GetTrack", ctx, input).Return(output, nil)

        result, err := mockClient.GetTrack(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetUsageLimit", func(t *testing.T) {
        input := &redshiftserverless.GetUsageLimitInput{}
        output := &redshiftserverless.GetUsageLimitOutput{}

        mockClient.On("GetUsageLimit", ctx, input).Return(output, nil)

        result, err := mockClient.GetUsageLimit(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetWorkgroup", func(t *testing.T) {
        input := &redshiftserverless.GetWorkgroupInput{}
        output := &redshiftserverless.GetWorkgroupOutput{}

        mockClient.On("GetWorkgroup", ctx, input).Return(output, nil)

        result, err := mockClient.GetWorkgroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListCustomDomainAssociations", func(t *testing.T) {
        input := &redshiftserverless.ListCustomDomainAssociationsInput{}
        output := &redshiftserverless.ListCustomDomainAssociationsOutput{}

        mockClient.On("ListCustomDomainAssociations", ctx, input).Return(output, nil)

        result, err := mockClient.ListCustomDomainAssociations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListEndpointAccess", func(t *testing.T) {
        input := &redshiftserverless.ListEndpointAccessInput{}
        output := &redshiftserverless.ListEndpointAccessOutput{}

        mockClient.On("ListEndpointAccess", ctx, input).Return(output, nil)

        result, err := mockClient.ListEndpointAccess(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListManagedWorkgroups", func(t *testing.T) {
        input := &redshiftserverless.ListManagedWorkgroupsInput{}
        output := &redshiftserverless.ListManagedWorkgroupsOutput{}

        mockClient.On("ListManagedWorkgroups", ctx, input).Return(output, nil)

        result, err := mockClient.ListManagedWorkgroups(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListNamespaces", func(t *testing.T) {
        input := &redshiftserverless.ListNamespacesInput{}
        output := &redshiftserverless.ListNamespacesOutput{}

        mockClient.On("ListNamespaces", ctx, input).Return(output, nil)

        result, err := mockClient.ListNamespaces(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListRecoveryPoints", func(t *testing.T) {
        input := &redshiftserverless.ListRecoveryPointsInput{}
        output := &redshiftserverless.ListRecoveryPointsOutput{}

        mockClient.On("ListRecoveryPoints", ctx, input).Return(output, nil)

        result, err := mockClient.ListRecoveryPoints(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListReservationOfferings", func(t *testing.T) {
        input := &redshiftserverless.ListReservationOfferingsInput{}
        output := &redshiftserverless.ListReservationOfferingsOutput{}

        mockClient.On("ListReservationOfferings", ctx, input).Return(output, nil)

        result, err := mockClient.ListReservationOfferings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListReservations", func(t *testing.T) {
        input := &redshiftserverless.ListReservationsInput{}
        output := &redshiftserverless.ListReservationsOutput{}

        mockClient.On("ListReservations", ctx, input).Return(output, nil)

        result, err := mockClient.ListReservations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListScheduledActions", func(t *testing.T) {
        input := &redshiftserverless.ListScheduledActionsInput{}
        output := &redshiftserverless.ListScheduledActionsOutput{}

        mockClient.On("ListScheduledActions", ctx, input).Return(output, nil)

        result, err := mockClient.ListScheduledActions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListSnapshotCopyConfigurations", func(t *testing.T) {
        input := &redshiftserverless.ListSnapshotCopyConfigurationsInput{}
        output := &redshiftserverless.ListSnapshotCopyConfigurationsOutput{}

        mockClient.On("ListSnapshotCopyConfigurations", ctx, input).Return(output, nil)

        result, err := mockClient.ListSnapshotCopyConfigurations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListSnapshots", func(t *testing.T) {
        input := &redshiftserverless.ListSnapshotsInput{}
        output := &redshiftserverless.ListSnapshotsOutput{}

        mockClient.On("ListSnapshots", ctx, input).Return(output, nil)

        result, err := mockClient.ListSnapshots(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTableRestoreStatus", func(t *testing.T) {
        input := &redshiftserverless.ListTableRestoreStatusInput{}
        output := &redshiftserverless.ListTableRestoreStatusOutput{}

        mockClient.On("ListTableRestoreStatus", ctx, input).Return(output, nil)

        result, err := mockClient.ListTableRestoreStatus(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &redshiftserverless.ListTagsForResourceInput{}
        output := &redshiftserverless.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTracks", func(t *testing.T) {
        input := &redshiftserverless.ListTracksInput{}
        output := &redshiftserverless.ListTracksOutput{}

        mockClient.On("ListTracks", ctx, input).Return(output, nil)

        result, err := mockClient.ListTracks(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListUsageLimits", func(t *testing.T) {
        input := &redshiftserverless.ListUsageLimitsInput{}
        output := &redshiftserverless.ListUsageLimitsOutput{}

        mockClient.On("ListUsageLimits", ctx, input).Return(output, nil)

        result, err := mockClient.ListUsageLimits(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListWorkgroups", func(t *testing.T) {
        input := &redshiftserverless.ListWorkgroupsInput{}
        output := &redshiftserverless.ListWorkgroupsOutput{}

        mockClient.On("ListWorkgroups", ctx, input).Return(output, nil)

        result, err := mockClient.ListWorkgroups(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutResourcePolicy", func(t *testing.T) {
        input := &redshiftserverless.PutResourcePolicyInput{}
        output := &redshiftserverless.PutResourcePolicyOutput{}

        mockClient.On("PutResourcePolicy", ctx, input).Return(output, nil)

        result, err := mockClient.PutResourcePolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRestoreFromRecoveryPoint", func(t *testing.T) {
        input := &redshiftserverless.RestoreFromRecoveryPointInput{}
        output := &redshiftserverless.RestoreFromRecoveryPointOutput{}

        mockClient.On("RestoreFromRecoveryPoint", ctx, input).Return(output, nil)

        result, err := mockClient.RestoreFromRecoveryPoint(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRestoreFromSnapshot", func(t *testing.T) {
        input := &redshiftserverless.RestoreFromSnapshotInput{}
        output := &redshiftserverless.RestoreFromSnapshotOutput{}

        mockClient.On("RestoreFromSnapshot", ctx, input).Return(output, nil)

        result, err := mockClient.RestoreFromSnapshot(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRestoreTableFromRecoveryPoint", func(t *testing.T) {
        input := &redshiftserverless.RestoreTableFromRecoveryPointInput{}
        output := &redshiftserverless.RestoreTableFromRecoveryPointOutput{}

        mockClient.On("RestoreTableFromRecoveryPoint", ctx, input).Return(output, nil)

        result, err := mockClient.RestoreTableFromRecoveryPoint(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRestoreTableFromSnapshot", func(t *testing.T) {
        input := &redshiftserverless.RestoreTableFromSnapshotInput{}
        output := &redshiftserverless.RestoreTableFromSnapshotOutput{}

        mockClient.On("RestoreTableFromSnapshot", ctx, input).Return(output, nil)

        result, err := mockClient.RestoreTableFromSnapshot(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &redshiftserverless.TagResourceInput{}
        output := &redshiftserverless.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &redshiftserverless.UntagResourceInput{}
        output := &redshiftserverless.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateCustomDomainAssociation", func(t *testing.T) {
        input := &redshiftserverless.UpdateCustomDomainAssociationInput{}
        output := &redshiftserverless.UpdateCustomDomainAssociationOutput{}

        mockClient.On("UpdateCustomDomainAssociation", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateCustomDomainAssociation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateEndpointAccess", func(t *testing.T) {
        input := &redshiftserverless.UpdateEndpointAccessInput{}
        output := &redshiftserverless.UpdateEndpointAccessOutput{}

        mockClient.On("UpdateEndpointAccess", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateEndpointAccess(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateNamespace", func(t *testing.T) {
        input := &redshiftserverless.UpdateNamespaceInput{}
        output := &redshiftserverless.UpdateNamespaceOutput{}

        mockClient.On("UpdateNamespace", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateNamespace(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateScheduledAction", func(t *testing.T) {
        input := &redshiftserverless.UpdateScheduledActionInput{}
        output := &redshiftserverless.UpdateScheduledActionOutput{}

        mockClient.On("UpdateScheduledAction", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateScheduledAction(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateSnapshot", func(t *testing.T) {
        input := &redshiftserverless.UpdateSnapshotInput{}
        output := &redshiftserverless.UpdateSnapshotOutput{}

        mockClient.On("UpdateSnapshot", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateSnapshot(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateSnapshotCopyConfiguration", func(t *testing.T) {
        input := &redshiftserverless.UpdateSnapshotCopyConfigurationInput{}
        output := &redshiftserverless.UpdateSnapshotCopyConfigurationOutput{}

        mockClient.On("UpdateSnapshotCopyConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateSnapshotCopyConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateUsageLimit", func(t *testing.T) {
        input := &redshiftserverless.UpdateUsageLimitInput{}
        output := &redshiftserverless.UpdateUsageLimitOutput{}

        mockClient.On("UpdateUsageLimit", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateUsageLimit(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateWorkgroup", func(t *testing.T) {
        input := &redshiftserverless.UpdateWorkgroupInput{}
        output := &redshiftserverless.UpdateWorkgroupOutput{}

        mockClient.On("UpdateWorkgroup", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateWorkgroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
