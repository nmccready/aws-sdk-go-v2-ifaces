package redshift_test

// tests for the redshift service interface for this ../../../service/redshift/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/redshift"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/redshift/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/redshift/redshift_iface"
	"github.com/stretchr/testify/assert"
)

func TestRedshiftServiceCanBeMocked(t *testing.T) {
	var iface redshift_iface.IClient
	iface = &redshift.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := redshift.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAcceptReservedNodeExchange", func(t *testing.T) {
        input := &redshift.AcceptReservedNodeExchangeInput{}
        output := &redshift.AcceptReservedNodeExchangeOutput{}

        mockClient.On("AcceptReservedNodeExchange", ctx, input).Return(output, nil)

        result, err := mockClient.AcceptReservedNodeExchange(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAddPartner", func(t *testing.T) {
        input := &redshift.AddPartnerInput{}
        output := &redshift.AddPartnerOutput{}

        mockClient.On("AddPartner", ctx, input).Return(output, nil)

        result, err := mockClient.AddPartner(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateDataShareConsumer", func(t *testing.T) {
        input := &redshift.AssociateDataShareConsumerInput{}
        output := &redshift.AssociateDataShareConsumerOutput{}

        mockClient.On("AssociateDataShareConsumer", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateDataShareConsumer(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAuthorizeClusterSecurityGroupIngress", func(t *testing.T) {
        input := &redshift.AuthorizeClusterSecurityGroupIngressInput{}
        output := &redshift.AuthorizeClusterSecurityGroupIngressOutput{}

        mockClient.On("AuthorizeClusterSecurityGroupIngress", ctx, input).Return(output, nil)

        result, err := mockClient.AuthorizeClusterSecurityGroupIngress(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAuthorizeDataShare", func(t *testing.T) {
        input := &redshift.AuthorizeDataShareInput{}
        output := &redshift.AuthorizeDataShareOutput{}

        mockClient.On("AuthorizeDataShare", ctx, input).Return(output, nil)

        result, err := mockClient.AuthorizeDataShare(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAuthorizeEndpointAccess", func(t *testing.T) {
        input := &redshift.AuthorizeEndpointAccessInput{}
        output := &redshift.AuthorizeEndpointAccessOutput{}

        mockClient.On("AuthorizeEndpointAccess", ctx, input).Return(output, nil)

        result, err := mockClient.AuthorizeEndpointAccess(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAuthorizeSnapshotAccess", func(t *testing.T) {
        input := &redshift.AuthorizeSnapshotAccessInput{}
        output := &redshift.AuthorizeSnapshotAccessOutput{}

        mockClient.On("AuthorizeSnapshotAccess", ctx, input).Return(output, nil)

        result, err := mockClient.AuthorizeSnapshotAccess(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchDeleteClusterSnapshots", func(t *testing.T) {
        input := &redshift.BatchDeleteClusterSnapshotsInput{}
        output := &redshift.BatchDeleteClusterSnapshotsOutput{}

        mockClient.On("BatchDeleteClusterSnapshots", ctx, input).Return(output, nil)

        result, err := mockClient.BatchDeleteClusterSnapshots(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchModifyClusterSnapshots", func(t *testing.T) {
        input := &redshift.BatchModifyClusterSnapshotsInput{}
        output := &redshift.BatchModifyClusterSnapshotsOutput{}

        mockClient.On("BatchModifyClusterSnapshots", ctx, input).Return(output, nil)

        result, err := mockClient.BatchModifyClusterSnapshots(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCancelResize", func(t *testing.T) {
        input := &redshift.CancelResizeInput{}
        output := &redshift.CancelResizeOutput{}

        mockClient.On("CancelResize", ctx, input).Return(output, nil)

        result, err := mockClient.CancelResize(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCopyClusterSnapshot", func(t *testing.T) {
        input := &redshift.CopyClusterSnapshotInput{}
        output := &redshift.CopyClusterSnapshotOutput{}

        mockClient.On("CopyClusterSnapshot", ctx, input).Return(output, nil)

        result, err := mockClient.CopyClusterSnapshot(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateAuthenticationProfile", func(t *testing.T) {
        input := &redshift.CreateAuthenticationProfileInput{}
        output := &redshift.CreateAuthenticationProfileOutput{}

        mockClient.On("CreateAuthenticationProfile", ctx, input).Return(output, nil)

        result, err := mockClient.CreateAuthenticationProfile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateCluster", func(t *testing.T) {
        input := &redshift.CreateClusterInput{}
        output := &redshift.CreateClusterOutput{}

        mockClient.On("CreateCluster", ctx, input).Return(output, nil)

        result, err := mockClient.CreateCluster(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateClusterParameterGroup", func(t *testing.T) {
        input := &redshift.CreateClusterParameterGroupInput{}
        output := &redshift.CreateClusterParameterGroupOutput{}

        mockClient.On("CreateClusterParameterGroup", ctx, input).Return(output, nil)

        result, err := mockClient.CreateClusterParameterGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateClusterSecurityGroup", func(t *testing.T) {
        input := &redshift.CreateClusterSecurityGroupInput{}
        output := &redshift.CreateClusterSecurityGroupOutput{}

        mockClient.On("CreateClusterSecurityGroup", ctx, input).Return(output, nil)

        result, err := mockClient.CreateClusterSecurityGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateClusterSnapshot", func(t *testing.T) {
        input := &redshift.CreateClusterSnapshotInput{}
        output := &redshift.CreateClusterSnapshotOutput{}

        mockClient.On("CreateClusterSnapshot", ctx, input).Return(output, nil)

        result, err := mockClient.CreateClusterSnapshot(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateClusterSubnetGroup", func(t *testing.T) {
        input := &redshift.CreateClusterSubnetGroupInput{}
        output := &redshift.CreateClusterSubnetGroupOutput{}

        mockClient.On("CreateClusterSubnetGroup", ctx, input).Return(output, nil)

        result, err := mockClient.CreateClusterSubnetGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateCustomDomainAssociation", func(t *testing.T) {
        input := &redshift.CreateCustomDomainAssociationInput{}
        output := &redshift.CreateCustomDomainAssociationOutput{}

        mockClient.On("CreateCustomDomainAssociation", ctx, input).Return(output, nil)

        result, err := mockClient.CreateCustomDomainAssociation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateEndpointAccess", func(t *testing.T) {
        input := &redshift.CreateEndpointAccessInput{}
        output := &redshift.CreateEndpointAccessOutput{}

        mockClient.On("CreateEndpointAccess", ctx, input).Return(output, nil)

        result, err := mockClient.CreateEndpointAccess(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateEventSubscription", func(t *testing.T) {
        input := &redshift.CreateEventSubscriptionInput{}
        output := &redshift.CreateEventSubscriptionOutput{}

        mockClient.On("CreateEventSubscription", ctx, input).Return(output, nil)

        result, err := mockClient.CreateEventSubscription(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateHsmClientCertificate", func(t *testing.T) {
        input := &redshift.CreateHsmClientCertificateInput{}
        output := &redshift.CreateHsmClientCertificateOutput{}

        mockClient.On("CreateHsmClientCertificate", ctx, input).Return(output, nil)

        result, err := mockClient.CreateHsmClientCertificate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateHsmConfiguration", func(t *testing.T) {
        input := &redshift.CreateHsmConfigurationInput{}
        output := &redshift.CreateHsmConfigurationOutput{}

        mockClient.On("CreateHsmConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.CreateHsmConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateRedshiftIdcApplication", func(t *testing.T) {
        input := &redshift.CreateRedshiftIdcApplicationInput{}
        output := &redshift.CreateRedshiftIdcApplicationOutput{}

        mockClient.On("CreateRedshiftIdcApplication", ctx, input).Return(output, nil)

        result, err := mockClient.CreateRedshiftIdcApplication(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateScheduledAction", func(t *testing.T) {
        input := &redshift.CreateScheduledActionInput{}
        output := &redshift.CreateScheduledActionOutput{}

        mockClient.On("CreateScheduledAction", ctx, input).Return(output, nil)

        result, err := mockClient.CreateScheduledAction(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateSnapshotCopyGrant", func(t *testing.T) {
        input := &redshift.CreateSnapshotCopyGrantInput{}
        output := &redshift.CreateSnapshotCopyGrantOutput{}

        mockClient.On("CreateSnapshotCopyGrant", ctx, input).Return(output, nil)

        result, err := mockClient.CreateSnapshotCopyGrant(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateSnapshotSchedule", func(t *testing.T) {
        input := &redshift.CreateSnapshotScheduleInput{}
        output := &redshift.CreateSnapshotScheduleOutput{}

        mockClient.On("CreateSnapshotSchedule", ctx, input).Return(output, nil)

        result, err := mockClient.CreateSnapshotSchedule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateTags", func(t *testing.T) {
        input := &redshift.CreateTagsInput{}
        output := &redshift.CreateTagsOutput{}

        mockClient.On("CreateTags", ctx, input).Return(output, nil)

        result, err := mockClient.CreateTags(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateUsageLimit", func(t *testing.T) {
        input := &redshift.CreateUsageLimitInput{}
        output := &redshift.CreateUsageLimitOutput{}

        mockClient.On("CreateUsageLimit", ctx, input).Return(output, nil)

        result, err := mockClient.CreateUsageLimit(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeauthorizeDataShare", func(t *testing.T) {
        input := &redshift.DeauthorizeDataShareInput{}
        output := &redshift.DeauthorizeDataShareOutput{}

        mockClient.On("DeauthorizeDataShare", ctx, input).Return(output, nil)

        result, err := mockClient.DeauthorizeDataShare(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteAuthenticationProfile", func(t *testing.T) {
        input := &redshift.DeleteAuthenticationProfileInput{}
        output := &redshift.DeleteAuthenticationProfileOutput{}

        mockClient.On("DeleteAuthenticationProfile", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteAuthenticationProfile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteCluster", func(t *testing.T) {
        input := &redshift.DeleteClusterInput{}
        output := &redshift.DeleteClusterOutput{}

        mockClient.On("DeleteCluster", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteCluster(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteClusterParameterGroup", func(t *testing.T) {
        input := &redshift.DeleteClusterParameterGroupInput{}
        output := &redshift.DeleteClusterParameterGroupOutput{}

        mockClient.On("DeleteClusterParameterGroup", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteClusterParameterGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteClusterSecurityGroup", func(t *testing.T) {
        input := &redshift.DeleteClusterSecurityGroupInput{}
        output := &redshift.DeleteClusterSecurityGroupOutput{}

        mockClient.On("DeleteClusterSecurityGroup", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteClusterSecurityGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteClusterSnapshot", func(t *testing.T) {
        input := &redshift.DeleteClusterSnapshotInput{}
        output := &redshift.DeleteClusterSnapshotOutput{}

        mockClient.On("DeleteClusterSnapshot", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteClusterSnapshot(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteClusterSubnetGroup", func(t *testing.T) {
        input := &redshift.DeleteClusterSubnetGroupInput{}
        output := &redshift.DeleteClusterSubnetGroupOutput{}

        mockClient.On("DeleteClusterSubnetGroup", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteClusterSubnetGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteCustomDomainAssociation", func(t *testing.T) {
        input := &redshift.DeleteCustomDomainAssociationInput{}
        output := &redshift.DeleteCustomDomainAssociationOutput{}

        mockClient.On("DeleteCustomDomainAssociation", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteCustomDomainAssociation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteEndpointAccess", func(t *testing.T) {
        input := &redshift.DeleteEndpointAccessInput{}
        output := &redshift.DeleteEndpointAccessOutput{}

        mockClient.On("DeleteEndpointAccess", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteEndpointAccess(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteEventSubscription", func(t *testing.T) {
        input := &redshift.DeleteEventSubscriptionInput{}
        output := &redshift.DeleteEventSubscriptionOutput{}

        mockClient.On("DeleteEventSubscription", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteEventSubscription(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteHsmClientCertificate", func(t *testing.T) {
        input := &redshift.DeleteHsmClientCertificateInput{}
        output := &redshift.DeleteHsmClientCertificateOutput{}

        mockClient.On("DeleteHsmClientCertificate", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteHsmClientCertificate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteHsmConfiguration", func(t *testing.T) {
        input := &redshift.DeleteHsmConfigurationInput{}
        output := &redshift.DeleteHsmConfigurationOutput{}

        mockClient.On("DeleteHsmConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteHsmConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeletePartner", func(t *testing.T) {
        input := &redshift.DeletePartnerInput{}
        output := &redshift.DeletePartnerOutput{}

        mockClient.On("DeletePartner", ctx, input).Return(output, nil)

        result, err := mockClient.DeletePartner(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteRedshiftIdcApplication", func(t *testing.T) {
        input := &redshift.DeleteRedshiftIdcApplicationInput{}
        output := &redshift.DeleteRedshiftIdcApplicationOutput{}

        mockClient.On("DeleteRedshiftIdcApplication", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteRedshiftIdcApplication(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteResourcePolicy", func(t *testing.T) {
        input := &redshift.DeleteResourcePolicyInput{}
        output := &redshift.DeleteResourcePolicyOutput{}

        mockClient.On("DeleteResourcePolicy", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteResourcePolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteScheduledAction", func(t *testing.T) {
        input := &redshift.DeleteScheduledActionInput{}
        output := &redshift.DeleteScheduledActionOutput{}

        mockClient.On("DeleteScheduledAction", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteScheduledAction(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteSnapshotCopyGrant", func(t *testing.T) {
        input := &redshift.DeleteSnapshotCopyGrantInput{}
        output := &redshift.DeleteSnapshotCopyGrantOutput{}

        mockClient.On("DeleteSnapshotCopyGrant", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteSnapshotCopyGrant(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteSnapshotSchedule", func(t *testing.T) {
        input := &redshift.DeleteSnapshotScheduleInput{}
        output := &redshift.DeleteSnapshotScheduleOutput{}

        mockClient.On("DeleteSnapshotSchedule", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteSnapshotSchedule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteTags", func(t *testing.T) {
        input := &redshift.DeleteTagsInput{}
        output := &redshift.DeleteTagsOutput{}

        mockClient.On("DeleteTags", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteTags(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteUsageLimit", func(t *testing.T) {
        input := &redshift.DeleteUsageLimitInput{}
        output := &redshift.DeleteUsageLimitOutput{}

        mockClient.On("DeleteUsageLimit", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteUsageLimit(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeAccountAttributes", func(t *testing.T) {
        input := &redshift.DescribeAccountAttributesInput{}
        output := &redshift.DescribeAccountAttributesOutput{}

        mockClient.On("DescribeAccountAttributes", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeAccountAttributes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeAuthenticationProfiles", func(t *testing.T) {
        input := &redshift.DescribeAuthenticationProfilesInput{}
        output := &redshift.DescribeAuthenticationProfilesOutput{}

        mockClient.On("DescribeAuthenticationProfiles", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeAuthenticationProfiles(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeClusterDbRevisions", func(t *testing.T) {
        input := &redshift.DescribeClusterDbRevisionsInput{}
        output := &redshift.DescribeClusterDbRevisionsOutput{}

        mockClient.On("DescribeClusterDbRevisions", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeClusterDbRevisions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeClusterParameterGroups", func(t *testing.T) {
        input := &redshift.DescribeClusterParameterGroupsInput{}
        output := &redshift.DescribeClusterParameterGroupsOutput{}

        mockClient.On("DescribeClusterParameterGroups", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeClusterParameterGroups(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeClusterParameters", func(t *testing.T) {
        input := &redshift.DescribeClusterParametersInput{}
        output := &redshift.DescribeClusterParametersOutput{}

        mockClient.On("DescribeClusterParameters", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeClusterParameters(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeClusterSecurityGroups", func(t *testing.T) {
        input := &redshift.DescribeClusterSecurityGroupsInput{}
        output := &redshift.DescribeClusterSecurityGroupsOutput{}

        mockClient.On("DescribeClusterSecurityGroups", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeClusterSecurityGroups(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeClusterSnapshots", func(t *testing.T) {
        input := &redshift.DescribeClusterSnapshotsInput{}
        output := &redshift.DescribeClusterSnapshotsOutput{}

        mockClient.On("DescribeClusterSnapshots", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeClusterSnapshots(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeClusterSubnetGroups", func(t *testing.T) {
        input := &redshift.DescribeClusterSubnetGroupsInput{}
        output := &redshift.DescribeClusterSubnetGroupsOutput{}

        mockClient.On("DescribeClusterSubnetGroups", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeClusterSubnetGroups(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeClusterTracks", func(t *testing.T) {
        input := &redshift.DescribeClusterTracksInput{}
        output := &redshift.DescribeClusterTracksOutput{}

        mockClient.On("DescribeClusterTracks", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeClusterTracks(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeClusterVersions", func(t *testing.T) {
        input := &redshift.DescribeClusterVersionsInput{}
        output := &redshift.DescribeClusterVersionsOutput{}

        mockClient.On("DescribeClusterVersions", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeClusterVersions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeClusters", func(t *testing.T) {
        input := &redshift.DescribeClustersInput{}
        output := &redshift.DescribeClustersOutput{}

        mockClient.On("DescribeClusters", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeClusters(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeCustomDomainAssociations", func(t *testing.T) {
        input := &redshift.DescribeCustomDomainAssociationsInput{}
        output := &redshift.DescribeCustomDomainAssociationsOutput{}

        mockClient.On("DescribeCustomDomainAssociations", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeCustomDomainAssociations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeDataShares", func(t *testing.T) {
        input := &redshift.DescribeDataSharesInput{}
        output := &redshift.DescribeDataSharesOutput{}

        mockClient.On("DescribeDataShares", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeDataShares(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeDataSharesForConsumer", func(t *testing.T) {
        input := &redshift.DescribeDataSharesForConsumerInput{}
        output := &redshift.DescribeDataSharesForConsumerOutput{}

        mockClient.On("DescribeDataSharesForConsumer", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeDataSharesForConsumer(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeDataSharesForProducer", func(t *testing.T) {
        input := &redshift.DescribeDataSharesForProducerInput{}
        output := &redshift.DescribeDataSharesForProducerOutput{}

        mockClient.On("DescribeDataSharesForProducer", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeDataSharesForProducer(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeDefaultClusterParameters", func(t *testing.T) {
        input := &redshift.DescribeDefaultClusterParametersInput{}
        output := &redshift.DescribeDefaultClusterParametersOutput{}

        mockClient.On("DescribeDefaultClusterParameters", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeDefaultClusterParameters(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeEndpointAccess", func(t *testing.T) {
        input := &redshift.DescribeEndpointAccessInput{}
        output := &redshift.DescribeEndpointAccessOutput{}

        mockClient.On("DescribeEndpointAccess", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeEndpointAccess(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeEndpointAuthorization", func(t *testing.T) {
        input := &redshift.DescribeEndpointAuthorizationInput{}
        output := &redshift.DescribeEndpointAuthorizationOutput{}

        mockClient.On("DescribeEndpointAuthorization", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeEndpointAuthorization(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeEventCategories", func(t *testing.T) {
        input := &redshift.DescribeEventCategoriesInput{}
        output := &redshift.DescribeEventCategoriesOutput{}

        mockClient.On("DescribeEventCategories", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeEventCategories(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeEventSubscriptions", func(t *testing.T) {
        input := &redshift.DescribeEventSubscriptionsInput{}
        output := &redshift.DescribeEventSubscriptionsOutput{}

        mockClient.On("DescribeEventSubscriptions", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeEventSubscriptions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeEvents", func(t *testing.T) {
        input := &redshift.DescribeEventsInput{}
        output := &redshift.DescribeEventsOutput{}

        mockClient.On("DescribeEvents", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeEvents(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeHsmClientCertificates", func(t *testing.T) {
        input := &redshift.DescribeHsmClientCertificatesInput{}
        output := &redshift.DescribeHsmClientCertificatesOutput{}

        mockClient.On("DescribeHsmClientCertificates", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeHsmClientCertificates(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeHsmConfigurations", func(t *testing.T) {
        input := &redshift.DescribeHsmConfigurationsInput{}
        output := &redshift.DescribeHsmConfigurationsOutput{}

        mockClient.On("DescribeHsmConfigurations", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeHsmConfigurations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeInboundIntegrations", func(t *testing.T) {
        input := &redshift.DescribeInboundIntegrationsInput{}
        output := &redshift.DescribeInboundIntegrationsOutput{}

        mockClient.On("DescribeInboundIntegrations", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeInboundIntegrations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeLoggingStatus", func(t *testing.T) {
        input := &redshift.DescribeLoggingStatusInput{}
        output := &redshift.DescribeLoggingStatusOutput{}

        mockClient.On("DescribeLoggingStatus", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeLoggingStatus(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeNodeConfigurationOptions", func(t *testing.T) {
        input := &redshift.DescribeNodeConfigurationOptionsInput{}
        output := &redshift.DescribeNodeConfigurationOptionsOutput{}

        mockClient.On("DescribeNodeConfigurationOptions", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeNodeConfigurationOptions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeOrderableClusterOptions", func(t *testing.T) {
        input := &redshift.DescribeOrderableClusterOptionsInput{}
        output := &redshift.DescribeOrderableClusterOptionsOutput{}

        mockClient.On("DescribeOrderableClusterOptions", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeOrderableClusterOptions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribePartners", func(t *testing.T) {
        input := &redshift.DescribePartnersInput{}
        output := &redshift.DescribePartnersOutput{}

        mockClient.On("DescribePartners", ctx, input).Return(output, nil)

        result, err := mockClient.DescribePartners(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeRedshiftIdcApplications", func(t *testing.T) {
        input := &redshift.DescribeRedshiftIdcApplicationsInput{}
        output := &redshift.DescribeRedshiftIdcApplicationsOutput{}

        mockClient.On("DescribeRedshiftIdcApplications", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeRedshiftIdcApplications(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeReservedNodeExchangeStatus", func(t *testing.T) {
        input := &redshift.DescribeReservedNodeExchangeStatusInput{}
        output := &redshift.DescribeReservedNodeExchangeStatusOutput{}

        mockClient.On("DescribeReservedNodeExchangeStatus", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeReservedNodeExchangeStatus(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeReservedNodeOfferings", func(t *testing.T) {
        input := &redshift.DescribeReservedNodeOfferingsInput{}
        output := &redshift.DescribeReservedNodeOfferingsOutput{}

        mockClient.On("DescribeReservedNodeOfferings", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeReservedNodeOfferings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeReservedNodes", func(t *testing.T) {
        input := &redshift.DescribeReservedNodesInput{}
        output := &redshift.DescribeReservedNodesOutput{}

        mockClient.On("DescribeReservedNodes", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeReservedNodes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeResize", func(t *testing.T) {
        input := &redshift.DescribeResizeInput{}
        output := &redshift.DescribeResizeOutput{}

        mockClient.On("DescribeResize", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeResize(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeScheduledActions", func(t *testing.T) {
        input := &redshift.DescribeScheduledActionsInput{}
        output := &redshift.DescribeScheduledActionsOutput{}

        mockClient.On("DescribeScheduledActions", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeScheduledActions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeSnapshotCopyGrants", func(t *testing.T) {
        input := &redshift.DescribeSnapshotCopyGrantsInput{}
        output := &redshift.DescribeSnapshotCopyGrantsOutput{}

        mockClient.On("DescribeSnapshotCopyGrants", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeSnapshotCopyGrants(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeSnapshotSchedules", func(t *testing.T) {
        input := &redshift.DescribeSnapshotSchedulesInput{}
        output := &redshift.DescribeSnapshotSchedulesOutput{}

        mockClient.On("DescribeSnapshotSchedules", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeSnapshotSchedules(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeStorage", func(t *testing.T) {
        input := &redshift.DescribeStorageInput{}
        output := &redshift.DescribeStorageOutput{}

        mockClient.On("DescribeStorage", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeStorage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeTableRestoreStatus", func(t *testing.T) {
        input := &redshift.DescribeTableRestoreStatusInput{}
        output := &redshift.DescribeTableRestoreStatusOutput{}

        mockClient.On("DescribeTableRestoreStatus", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeTableRestoreStatus(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeTags", func(t *testing.T) {
        input := &redshift.DescribeTagsInput{}
        output := &redshift.DescribeTagsOutput{}

        mockClient.On("DescribeTags", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeTags(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeUsageLimits", func(t *testing.T) {
        input := &redshift.DescribeUsageLimitsInput{}
        output := &redshift.DescribeUsageLimitsOutput{}

        mockClient.On("DescribeUsageLimits", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeUsageLimits(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisableLogging", func(t *testing.T) {
        input := &redshift.DisableLoggingInput{}
        output := &redshift.DisableLoggingOutput{}

        mockClient.On("DisableLogging", ctx, input).Return(output, nil)

        result, err := mockClient.DisableLogging(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisableSnapshotCopy", func(t *testing.T) {
        input := &redshift.DisableSnapshotCopyInput{}
        output := &redshift.DisableSnapshotCopyOutput{}

        mockClient.On("DisableSnapshotCopy", ctx, input).Return(output, nil)

        result, err := mockClient.DisableSnapshotCopy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateDataShareConsumer", func(t *testing.T) {
        input := &redshift.DisassociateDataShareConsumerInput{}
        output := &redshift.DisassociateDataShareConsumerOutput{}

        mockClient.On("DisassociateDataShareConsumer", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateDataShareConsumer(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestEnableLogging", func(t *testing.T) {
        input := &redshift.EnableLoggingInput{}
        output := &redshift.EnableLoggingOutput{}

        mockClient.On("EnableLogging", ctx, input).Return(output, nil)

        result, err := mockClient.EnableLogging(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestEnableSnapshotCopy", func(t *testing.T) {
        input := &redshift.EnableSnapshotCopyInput{}
        output := &redshift.EnableSnapshotCopyOutput{}

        mockClient.On("EnableSnapshotCopy", ctx, input).Return(output, nil)

        result, err := mockClient.EnableSnapshotCopy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestFailoverPrimaryCompute", func(t *testing.T) {
        input := &redshift.FailoverPrimaryComputeInput{}
        output := &redshift.FailoverPrimaryComputeOutput{}

        mockClient.On("FailoverPrimaryCompute", ctx, input).Return(output, nil)

        result, err := mockClient.FailoverPrimaryCompute(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetClusterCredentials", func(t *testing.T) {
        input := &redshift.GetClusterCredentialsInput{}
        output := &redshift.GetClusterCredentialsOutput{}

        mockClient.On("GetClusterCredentials", ctx, input).Return(output, nil)

        result, err := mockClient.GetClusterCredentials(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetClusterCredentialsWithIAM", func(t *testing.T) {
        input := &redshift.GetClusterCredentialsWithIAMInput{}
        output := &redshift.GetClusterCredentialsWithIAMOutput{}

        mockClient.On("GetClusterCredentialsWithIAM", ctx, input).Return(output, nil)

        result, err := mockClient.GetClusterCredentialsWithIAM(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetReservedNodeExchangeConfigurationOptions", func(t *testing.T) {
        input := &redshift.GetReservedNodeExchangeConfigurationOptionsInput{}
        output := &redshift.GetReservedNodeExchangeConfigurationOptionsOutput{}

        mockClient.On("GetReservedNodeExchangeConfigurationOptions", ctx, input).Return(output, nil)

        result, err := mockClient.GetReservedNodeExchangeConfigurationOptions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetReservedNodeExchangeOfferings", func(t *testing.T) {
        input := &redshift.GetReservedNodeExchangeOfferingsInput{}
        output := &redshift.GetReservedNodeExchangeOfferingsOutput{}

        mockClient.On("GetReservedNodeExchangeOfferings", ctx, input).Return(output, nil)

        result, err := mockClient.GetReservedNodeExchangeOfferings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetResourcePolicy", func(t *testing.T) {
        input := &redshift.GetResourcePolicyInput{}
        output := &redshift.GetResourcePolicyOutput{}

        mockClient.On("GetResourcePolicy", ctx, input).Return(output, nil)

        result, err := mockClient.GetResourcePolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListRecommendations", func(t *testing.T) {
        input := &redshift.ListRecommendationsInput{}
        output := &redshift.ListRecommendationsOutput{}

        mockClient.On("ListRecommendations", ctx, input).Return(output, nil)

        result, err := mockClient.ListRecommendations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyAquaConfiguration", func(t *testing.T) {
        input := &redshift.ModifyAquaConfigurationInput{}
        output := &redshift.ModifyAquaConfigurationOutput{}

        mockClient.On("ModifyAquaConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyAquaConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyAuthenticationProfile", func(t *testing.T) {
        input := &redshift.ModifyAuthenticationProfileInput{}
        output := &redshift.ModifyAuthenticationProfileOutput{}

        mockClient.On("ModifyAuthenticationProfile", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyAuthenticationProfile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyCluster", func(t *testing.T) {
        input := &redshift.ModifyClusterInput{}
        output := &redshift.ModifyClusterOutput{}

        mockClient.On("ModifyCluster", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyCluster(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyClusterDbRevision", func(t *testing.T) {
        input := &redshift.ModifyClusterDbRevisionInput{}
        output := &redshift.ModifyClusterDbRevisionOutput{}

        mockClient.On("ModifyClusterDbRevision", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyClusterDbRevision(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyClusterIamRoles", func(t *testing.T) {
        input := &redshift.ModifyClusterIamRolesInput{}
        output := &redshift.ModifyClusterIamRolesOutput{}

        mockClient.On("ModifyClusterIamRoles", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyClusterIamRoles(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyClusterMaintenance", func(t *testing.T) {
        input := &redshift.ModifyClusterMaintenanceInput{}
        output := &redshift.ModifyClusterMaintenanceOutput{}

        mockClient.On("ModifyClusterMaintenance", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyClusterMaintenance(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyClusterParameterGroup", func(t *testing.T) {
        input := &redshift.ModifyClusterParameterGroupInput{}
        output := &redshift.ModifyClusterParameterGroupOutput{}

        mockClient.On("ModifyClusterParameterGroup", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyClusterParameterGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyClusterSnapshot", func(t *testing.T) {
        input := &redshift.ModifyClusterSnapshotInput{}
        output := &redshift.ModifyClusterSnapshotOutput{}

        mockClient.On("ModifyClusterSnapshot", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyClusterSnapshot(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyClusterSnapshotSchedule", func(t *testing.T) {
        input := &redshift.ModifyClusterSnapshotScheduleInput{}
        output := &redshift.ModifyClusterSnapshotScheduleOutput{}

        mockClient.On("ModifyClusterSnapshotSchedule", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyClusterSnapshotSchedule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyClusterSubnetGroup", func(t *testing.T) {
        input := &redshift.ModifyClusterSubnetGroupInput{}
        output := &redshift.ModifyClusterSubnetGroupOutput{}

        mockClient.On("ModifyClusterSubnetGroup", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyClusterSubnetGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyCustomDomainAssociation", func(t *testing.T) {
        input := &redshift.ModifyCustomDomainAssociationInput{}
        output := &redshift.ModifyCustomDomainAssociationOutput{}

        mockClient.On("ModifyCustomDomainAssociation", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyCustomDomainAssociation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyEndpointAccess", func(t *testing.T) {
        input := &redshift.ModifyEndpointAccessInput{}
        output := &redshift.ModifyEndpointAccessOutput{}

        mockClient.On("ModifyEndpointAccess", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyEndpointAccess(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyEventSubscription", func(t *testing.T) {
        input := &redshift.ModifyEventSubscriptionInput{}
        output := &redshift.ModifyEventSubscriptionOutput{}

        mockClient.On("ModifyEventSubscription", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyEventSubscription(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyRedshiftIdcApplication", func(t *testing.T) {
        input := &redshift.ModifyRedshiftIdcApplicationInput{}
        output := &redshift.ModifyRedshiftIdcApplicationOutput{}

        mockClient.On("ModifyRedshiftIdcApplication", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyRedshiftIdcApplication(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyScheduledAction", func(t *testing.T) {
        input := &redshift.ModifyScheduledActionInput{}
        output := &redshift.ModifyScheduledActionOutput{}

        mockClient.On("ModifyScheduledAction", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyScheduledAction(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifySnapshotCopyRetentionPeriod", func(t *testing.T) {
        input := &redshift.ModifySnapshotCopyRetentionPeriodInput{}
        output := &redshift.ModifySnapshotCopyRetentionPeriodOutput{}

        mockClient.On("ModifySnapshotCopyRetentionPeriod", ctx, input).Return(output, nil)

        result, err := mockClient.ModifySnapshotCopyRetentionPeriod(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifySnapshotSchedule", func(t *testing.T) {
        input := &redshift.ModifySnapshotScheduleInput{}
        output := &redshift.ModifySnapshotScheduleOutput{}

        mockClient.On("ModifySnapshotSchedule", ctx, input).Return(output, nil)

        result, err := mockClient.ModifySnapshotSchedule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyUsageLimit", func(t *testing.T) {
        input := &redshift.ModifyUsageLimitInput{}
        output := &redshift.ModifyUsageLimitOutput{}

        mockClient.On("ModifyUsageLimit", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyUsageLimit(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPauseCluster", func(t *testing.T) {
        input := &redshift.PauseClusterInput{}
        output := &redshift.PauseClusterOutput{}

        mockClient.On("PauseCluster", ctx, input).Return(output, nil)

        result, err := mockClient.PauseCluster(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPurchaseReservedNodeOffering", func(t *testing.T) {
        input := &redshift.PurchaseReservedNodeOfferingInput{}
        output := &redshift.PurchaseReservedNodeOfferingOutput{}

        mockClient.On("PurchaseReservedNodeOffering", ctx, input).Return(output, nil)

        result, err := mockClient.PurchaseReservedNodeOffering(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutResourcePolicy", func(t *testing.T) {
        input := &redshift.PutResourcePolicyInput{}
        output := &redshift.PutResourcePolicyOutput{}

        mockClient.On("PutResourcePolicy", ctx, input).Return(output, nil)

        result, err := mockClient.PutResourcePolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRebootCluster", func(t *testing.T) {
        input := &redshift.RebootClusterInput{}
        output := &redshift.RebootClusterOutput{}

        mockClient.On("RebootCluster", ctx, input).Return(output, nil)

        result, err := mockClient.RebootCluster(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRejectDataShare", func(t *testing.T) {
        input := &redshift.RejectDataShareInput{}
        output := &redshift.RejectDataShareOutput{}

        mockClient.On("RejectDataShare", ctx, input).Return(output, nil)

        result, err := mockClient.RejectDataShare(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestResetClusterParameterGroup", func(t *testing.T) {
        input := &redshift.ResetClusterParameterGroupInput{}
        output := &redshift.ResetClusterParameterGroupOutput{}

        mockClient.On("ResetClusterParameterGroup", ctx, input).Return(output, nil)

        result, err := mockClient.ResetClusterParameterGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestResizeCluster", func(t *testing.T) {
        input := &redshift.ResizeClusterInput{}
        output := &redshift.ResizeClusterOutput{}

        mockClient.On("ResizeCluster", ctx, input).Return(output, nil)

        result, err := mockClient.ResizeCluster(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRestoreFromClusterSnapshot", func(t *testing.T) {
        input := &redshift.RestoreFromClusterSnapshotInput{}
        output := &redshift.RestoreFromClusterSnapshotOutput{}

        mockClient.On("RestoreFromClusterSnapshot", ctx, input).Return(output, nil)

        result, err := mockClient.RestoreFromClusterSnapshot(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRestoreTableFromClusterSnapshot", func(t *testing.T) {
        input := &redshift.RestoreTableFromClusterSnapshotInput{}
        output := &redshift.RestoreTableFromClusterSnapshotOutput{}

        mockClient.On("RestoreTableFromClusterSnapshot", ctx, input).Return(output, nil)

        result, err := mockClient.RestoreTableFromClusterSnapshot(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestResumeCluster", func(t *testing.T) {
        input := &redshift.ResumeClusterInput{}
        output := &redshift.ResumeClusterOutput{}

        mockClient.On("ResumeCluster", ctx, input).Return(output, nil)

        result, err := mockClient.ResumeCluster(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRevokeClusterSecurityGroupIngress", func(t *testing.T) {
        input := &redshift.RevokeClusterSecurityGroupIngressInput{}
        output := &redshift.RevokeClusterSecurityGroupIngressOutput{}

        mockClient.On("RevokeClusterSecurityGroupIngress", ctx, input).Return(output, nil)

        result, err := mockClient.RevokeClusterSecurityGroupIngress(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRevokeEndpointAccess", func(t *testing.T) {
        input := &redshift.RevokeEndpointAccessInput{}
        output := &redshift.RevokeEndpointAccessOutput{}

        mockClient.On("RevokeEndpointAccess", ctx, input).Return(output, nil)

        result, err := mockClient.RevokeEndpointAccess(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRevokeSnapshotAccess", func(t *testing.T) {
        input := &redshift.RevokeSnapshotAccessInput{}
        output := &redshift.RevokeSnapshotAccessOutput{}

        mockClient.On("RevokeSnapshotAccess", ctx, input).Return(output, nil)

        result, err := mockClient.RevokeSnapshotAccess(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRotateEncryptionKey", func(t *testing.T) {
        input := &redshift.RotateEncryptionKeyInput{}
        output := &redshift.RotateEncryptionKeyOutput{}

        mockClient.On("RotateEncryptionKey", ctx, input).Return(output, nil)

        result, err := mockClient.RotateEncryptionKey(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdatePartnerStatus", func(t *testing.T) {
        input := &redshift.UpdatePartnerStatusInput{}
        output := &redshift.UpdatePartnerStatusOutput{}

        mockClient.On("UpdatePartnerStatus", ctx, input).Return(output, nil)

        result, err := mockClient.UpdatePartnerStatus(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
