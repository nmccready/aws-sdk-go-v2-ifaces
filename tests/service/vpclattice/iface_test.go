// Code generated by ./scripts/gen_tests.sh DO NOT EDIT.
package vpclattice_test

// tests for the vpclattice service interface for 
// this ../../../service/vpclattice/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/vpclattice"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/vpclattice/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/vpclattice/vpclattice_iface"
	"github.com/stretchr/testify/assert"
)

func TestVpclatticeServiceCanBeMocked(t *testing.T) {
	var iface vpclattice_iface.IClient
	iface = &vpclattice.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := vpclattice.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchUpdateRule", func(t *testing.T) {
        input := &vpclattice.BatchUpdateRuleInput{}
        output := &vpclattice.BatchUpdateRuleOutput{}

        mockClient.On("BatchUpdateRule", ctx, input).Return(output, nil)

        result, err := mockClient.BatchUpdateRule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateAccessLogSubscription", func(t *testing.T) {
        input := &vpclattice.CreateAccessLogSubscriptionInput{}
        output := &vpclattice.CreateAccessLogSubscriptionOutput{}

        mockClient.On("CreateAccessLogSubscription", ctx, input).Return(output, nil)

        result, err := mockClient.CreateAccessLogSubscription(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateListener", func(t *testing.T) {
        input := &vpclattice.CreateListenerInput{}
        output := &vpclattice.CreateListenerOutput{}

        mockClient.On("CreateListener", ctx, input).Return(output, nil)

        result, err := mockClient.CreateListener(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateResourceConfiguration", func(t *testing.T) {
        input := &vpclattice.CreateResourceConfigurationInput{}
        output := &vpclattice.CreateResourceConfigurationOutput{}

        mockClient.On("CreateResourceConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.CreateResourceConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateResourceGateway", func(t *testing.T) {
        input := &vpclattice.CreateResourceGatewayInput{}
        output := &vpclattice.CreateResourceGatewayOutput{}

        mockClient.On("CreateResourceGateway", ctx, input).Return(output, nil)

        result, err := mockClient.CreateResourceGateway(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateRule", func(t *testing.T) {
        input := &vpclattice.CreateRuleInput{}
        output := &vpclattice.CreateRuleOutput{}

        mockClient.On("CreateRule", ctx, input).Return(output, nil)

        result, err := mockClient.CreateRule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateService", func(t *testing.T) {
        input := &vpclattice.CreateServiceInput{}
        output := &vpclattice.CreateServiceOutput{}

        mockClient.On("CreateService", ctx, input).Return(output, nil)

        result, err := mockClient.CreateService(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateServiceNetwork", func(t *testing.T) {
        input := &vpclattice.CreateServiceNetworkInput{}
        output := &vpclattice.CreateServiceNetworkOutput{}

        mockClient.On("CreateServiceNetwork", ctx, input).Return(output, nil)

        result, err := mockClient.CreateServiceNetwork(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateServiceNetworkResourceAssociation", func(t *testing.T) {
        input := &vpclattice.CreateServiceNetworkResourceAssociationInput{}
        output := &vpclattice.CreateServiceNetworkResourceAssociationOutput{}

        mockClient.On("CreateServiceNetworkResourceAssociation", ctx, input).Return(output, nil)

        result, err := mockClient.CreateServiceNetworkResourceAssociation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateServiceNetworkServiceAssociation", func(t *testing.T) {
        input := &vpclattice.CreateServiceNetworkServiceAssociationInput{}
        output := &vpclattice.CreateServiceNetworkServiceAssociationOutput{}

        mockClient.On("CreateServiceNetworkServiceAssociation", ctx, input).Return(output, nil)

        result, err := mockClient.CreateServiceNetworkServiceAssociation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateServiceNetworkVpcAssociation", func(t *testing.T) {
        input := &vpclattice.CreateServiceNetworkVpcAssociationInput{}
        output := &vpclattice.CreateServiceNetworkVpcAssociationOutput{}

        mockClient.On("CreateServiceNetworkVpcAssociation", ctx, input).Return(output, nil)

        result, err := mockClient.CreateServiceNetworkVpcAssociation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateTargetGroup", func(t *testing.T) {
        input := &vpclattice.CreateTargetGroupInput{}
        output := &vpclattice.CreateTargetGroupOutput{}

        mockClient.On("CreateTargetGroup", ctx, input).Return(output, nil)

        result, err := mockClient.CreateTargetGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteAccessLogSubscription", func(t *testing.T) {
        input := &vpclattice.DeleteAccessLogSubscriptionInput{}
        output := &vpclattice.DeleteAccessLogSubscriptionOutput{}

        mockClient.On("DeleteAccessLogSubscription", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteAccessLogSubscription(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteAuthPolicy", func(t *testing.T) {
        input := &vpclattice.DeleteAuthPolicyInput{}
        output := &vpclattice.DeleteAuthPolicyOutput{}

        mockClient.On("DeleteAuthPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteAuthPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteListener", func(t *testing.T) {
        input := &vpclattice.DeleteListenerInput{}
        output := &vpclattice.DeleteListenerOutput{}

        mockClient.On("DeleteListener", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteListener(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteResourceConfiguration", func(t *testing.T) {
        input := &vpclattice.DeleteResourceConfigurationInput{}
        output := &vpclattice.DeleteResourceConfigurationOutput{}

        mockClient.On("DeleteResourceConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteResourceConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteResourceEndpointAssociation", func(t *testing.T) {
        input := &vpclattice.DeleteResourceEndpointAssociationInput{}
        output := &vpclattice.DeleteResourceEndpointAssociationOutput{}

        mockClient.On("DeleteResourceEndpointAssociation", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteResourceEndpointAssociation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteResourceGateway", func(t *testing.T) {
        input := &vpclattice.DeleteResourceGatewayInput{}
        output := &vpclattice.DeleteResourceGatewayOutput{}

        mockClient.On("DeleteResourceGateway", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteResourceGateway(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteResourcePolicy", func(t *testing.T) {
        input := &vpclattice.DeleteResourcePolicyInput{}
        output := &vpclattice.DeleteResourcePolicyOutput{}

        mockClient.On("DeleteResourcePolicy", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteResourcePolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteRule", func(t *testing.T) {
        input := &vpclattice.DeleteRuleInput{}
        output := &vpclattice.DeleteRuleOutput{}

        mockClient.On("DeleteRule", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteRule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteService", func(t *testing.T) {
        input := &vpclattice.DeleteServiceInput{}
        output := &vpclattice.DeleteServiceOutput{}

        mockClient.On("DeleteService", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteService(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteServiceNetwork", func(t *testing.T) {
        input := &vpclattice.DeleteServiceNetworkInput{}
        output := &vpclattice.DeleteServiceNetworkOutput{}

        mockClient.On("DeleteServiceNetwork", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteServiceNetwork(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteServiceNetworkResourceAssociation", func(t *testing.T) {
        input := &vpclattice.DeleteServiceNetworkResourceAssociationInput{}
        output := &vpclattice.DeleteServiceNetworkResourceAssociationOutput{}

        mockClient.On("DeleteServiceNetworkResourceAssociation", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteServiceNetworkResourceAssociation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteServiceNetworkServiceAssociation", func(t *testing.T) {
        input := &vpclattice.DeleteServiceNetworkServiceAssociationInput{}
        output := &vpclattice.DeleteServiceNetworkServiceAssociationOutput{}

        mockClient.On("DeleteServiceNetworkServiceAssociation", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteServiceNetworkServiceAssociation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteServiceNetworkVpcAssociation", func(t *testing.T) {
        input := &vpclattice.DeleteServiceNetworkVpcAssociationInput{}
        output := &vpclattice.DeleteServiceNetworkVpcAssociationOutput{}

        mockClient.On("DeleteServiceNetworkVpcAssociation", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteServiceNetworkVpcAssociation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteTargetGroup", func(t *testing.T) {
        input := &vpclattice.DeleteTargetGroupInput{}
        output := &vpclattice.DeleteTargetGroupOutput{}

        mockClient.On("DeleteTargetGroup", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteTargetGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeregisterTargets", func(t *testing.T) {
        input := &vpclattice.DeregisterTargetsInput{}
        output := &vpclattice.DeregisterTargetsOutput{}

        mockClient.On("DeregisterTargets", ctx, input).Return(output, nil)

        result, err := mockClient.DeregisterTargets(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetAccessLogSubscription", func(t *testing.T) {
        input := &vpclattice.GetAccessLogSubscriptionInput{}
        output := &vpclattice.GetAccessLogSubscriptionOutput{}

        mockClient.On("GetAccessLogSubscription", ctx, input).Return(output, nil)

        result, err := mockClient.GetAccessLogSubscription(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetAuthPolicy", func(t *testing.T) {
        input := &vpclattice.GetAuthPolicyInput{}
        output := &vpclattice.GetAuthPolicyOutput{}

        mockClient.On("GetAuthPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.GetAuthPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetListener", func(t *testing.T) {
        input := &vpclattice.GetListenerInput{}
        output := &vpclattice.GetListenerOutput{}

        mockClient.On("GetListener", ctx, input).Return(output, nil)

        result, err := mockClient.GetListener(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetResourceConfiguration", func(t *testing.T) {
        input := &vpclattice.GetResourceConfigurationInput{}
        output := &vpclattice.GetResourceConfigurationOutput{}

        mockClient.On("GetResourceConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.GetResourceConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetResourceGateway", func(t *testing.T) {
        input := &vpclattice.GetResourceGatewayInput{}
        output := &vpclattice.GetResourceGatewayOutput{}

        mockClient.On("GetResourceGateway", ctx, input).Return(output, nil)

        result, err := mockClient.GetResourceGateway(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetResourcePolicy", func(t *testing.T) {
        input := &vpclattice.GetResourcePolicyInput{}
        output := &vpclattice.GetResourcePolicyOutput{}

        mockClient.On("GetResourcePolicy", ctx, input).Return(output, nil)

        result, err := mockClient.GetResourcePolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetRule", func(t *testing.T) {
        input := &vpclattice.GetRuleInput{}
        output := &vpclattice.GetRuleOutput{}

        mockClient.On("GetRule", ctx, input).Return(output, nil)

        result, err := mockClient.GetRule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetService", func(t *testing.T) {
        input := &vpclattice.GetServiceInput{}
        output := &vpclattice.GetServiceOutput{}

        mockClient.On("GetService", ctx, input).Return(output, nil)

        result, err := mockClient.GetService(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetServiceNetwork", func(t *testing.T) {
        input := &vpclattice.GetServiceNetworkInput{}
        output := &vpclattice.GetServiceNetworkOutput{}

        mockClient.On("GetServiceNetwork", ctx, input).Return(output, nil)

        result, err := mockClient.GetServiceNetwork(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetServiceNetworkResourceAssociation", func(t *testing.T) {
        input := &vpclattice.GetServiceNetworkResourceAssociationInput{}
        output := &vpclattice.GetServiceNetworkResourceAssociationOutput{}

        mockClient.On("GetServiceNetworkResourceAssociation", ctx, input).Return(output, nil)

        result, err := mockClient.GetServiceNetworkResourceAssociation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetServiceNetworkServiceAssociation", func(t *testing.T) {
        input := &vpclattice.GetServiceNetworkServiceAssociationInput{}
        output := &vpclattice.GetServiceNetworkServiceAssociationOutput{}

        mockClient.On("GetServiceNetworkServiceAssociation", ctx, input).Return(output, nil)

        result, err := mockClient.GetServiceNetworkServiceAssociation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetServiceNetworkVpcAssociation", func(t *testing.T) {
        input := &vpclattice.GetServiceNetworkVpcAssociationInput{}
        output := &vpclattice.GetServiceNetworkVpcAssociationOutput{}

        mockClient.On("GetServiceNetworkVpcAssociation", ctx, input).Return(output, nil)

        result, err := mockClient.GetServiceNetworkVpcAssociation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetTargetGroup", func(t *testing.T) {
        input := &vpclattice.GetTargetGroupInput{}
        output := &vpclattice.GetTargetGroupOutput{}

        mockClient.On("GetTargetGroup", ctx, input).Return(output, nil)

        result, err := mockClient.GetTargetGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAccessLogSubscriptions", func(t *testing.T) {
        input := &vpclattice.ListAccessLogSubscriptionsInput{}
        output := &vpclattice.ListAccessLogSubscriptionsOutput{}

        mockClient.On("ListAccessLogSubscriptions", ctx, input).Return(output, nil)

        result, err := mockClient.ListAccessLogSubscriptions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListListeners", func(t *testing.T) {
        input := &vpclattice.ListListenersInput{}
        output := &vpclattice.ListListenersOutput{}

        mockClient.On("ListListeners", ctx, input).Return(output, nil)

        result, err := mockClient.ListListeners(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListResourceConfigurations", func(t *testing.T) {
        input := &vpclattice.ListResourceConfigurationsInput{}
        output := &vpclattice.ListResourceConfigurationsOutput{}

        mockClient.On("ListResourceConfigurations", ctx, input).Return(output, nil)

        result, err := mockClient.ListResourceConfigurations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListResourceEndpointAssociations", func(t *testing.T) {
        input := &vpclattice.ListResourceEndpointAssociationsInput{}
        output := &vpclattice.ListResourceEndpointAssociationsOutput{}

        mockClient.On("ListResourceEndpointAssociations", ctx, input).Return(output, nil)

        result, err := mockClient.ListResourceEndpointAssociations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListResourceGateways", func(t *testing.T) {
        input := &vpclattice.ListResourceGatewaysInput{}
        output := &vpclattice.ListResourceGatewaysOutput{}

        mockClient.On("ListResourceGateways", ctx, input).Return(output, nil)

        result, err := mockClient.ListResourceGateways(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListRules", func(t *testing.T) {
        input := &vpclattice.ListRulesInput{}
        output := &vpclattice.ListRulesOutput{}

        mockClient.On("ListRules", ctx, input).Return(output, nil)

        result, err := mockClient.ListRules(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListServiceNetworkResourceAssociations", func(t *testing.T) {
        input := &vpclattice.ListServiceNetworkResourceAssociationsInput{}
        output := &vpclattice.ListServiceNetworkResourceAssociationsOutput{}

        mockClient.On("ListServiceNetworkResourceAssociations", ctx, input).Return(output, nil)

        result, err := mockClient.ListServiceNetworkResourceAssociations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListServiceNetworkServiceAssociations", func(t *testing.T) {
        input := &vpclattice.ListServiceNetworkServiceAssociationsInput{}
        output := &vpclattice.ListServiceNetworkServiceAssociationsOutput{}

        mockClient.On("ListServiceNetworkServiceAssociations", ctx, input).Return(output, nil)

        result, err := mockClient.ListServiceNetworkServiceAssociations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListServiceNetworkVpcAssociations", func(t *testing.T) {
        input := &vpclattice.ListServiceNetworkVpcAssociationsInput{}
        output := &vpclattice.ListServiceNetworkVpcAssociationsOutput{}

        mockClient.On("ListServiceNetworkVpcAssociations", ctx, input).Return(output, nil)

        result, err := mockClient.ListServiceNetworkVpcAssociations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListServiceNetworkVpcEndpointAssociations", func(t *testing.T) {
        input := &vpclattice.ListServiceNetworkVpcEndpointAssociationsInput{}
        output := &vpclattice.ListServiceNetworkVpcEndpointAssociationsOutput{}

        mockClient.On("ListServiceNetworkVpcEndpointAssociations", ctx, input).Return(output, nil)

        result, err := mockClient.ListServiceNetworkVpcEndpointAssociations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListServiceNetworks", func(t *testing.T) {
        input := &vpclattice.ListServiceNetworksInput{}
        output := &vpclattice.ListServiceNetworksOutput{}

        mockClient.On("ListServiceNetworks", ctx, input).Return(output, nil)

        result, err := mockClient.ListServiceNetworks(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListServices", func(t *testing.T) {
        input := &vpclattice.ListServicesInput{}
        output := &vpclattice.ListServicesOutput{}

        mockClient.On("ListServices", ctx, input).Return(output, nil)

        result, err := mockClient.ListServices(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &vpclattice.ListTagsForResourceInput{}
        output := &vpclattice.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTargetGroups", func(t *testing.T) {
        input := &vpclattice.ListTargetGroupsInput{}
        output := &vpclattice.ListTargetGroupsOutput{}

        mockClient.On("ListTargetGroups", ctx, input).Return(output, nil)

        result, err := mockClient.ListTargetGroups(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTargets", func(t *testing.T) {
        input := &vpclattice.ListTargetsInput{}
        output := &vpclattice.ListTargetsOutput{}

        mockClient.On("ListTargets", ctx, input).Return(output, nil)

        result, err := mockClient.ListTargets(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutAuthPolicy", func(t *testing.T) {
        input := &vpclattice.PutAuthPolicyInput{}
        output := &vpclattice.PutAuthPolicyOutput{}

        mockClient.On("PutAuthPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.PutAuthPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutResourcePolicy", func(t *testing.T) {
        input := &vpclattice.PutResourcePolicyInput{}
        output := &vpclattice.PutResourcePolicyOutput{}

        mockClient.On("PutResourcePolicy", ctx, input).Return(output, nil)

        result, err := mockClient.PutResourcePolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRegisterTargets", func(t *testing.T) {
        input := &vpclattice.RegisterTargetsInput{}
        output := &vpclattice.RegisterTargetsOutput{}

        mockClient.On("RegisterTargets", ctx, input).Return(output, nil)

        result, err := mockClient.RegisterTargets(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &vpclattice.TagResourceInput{}
        output := &vpclattice.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &vpclattice.UntagResourceInput{}
        output := &vpclattice.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateAccessLogSubscription", func(t *testing.T) {
        input := &vpclattice.UpdateAccessLogSubscriptionInput{}
        output := &vpclattice.UpdateAccessLogSubscriptionOutput{}

        mockClient.On("UpdateAccessLogSubscription", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateAccessLogSubscription(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateListener", func(t *testing.T) {
        input := &vpclattice.UpdateListenerInput{}
        output := &vpclattice.UpdateListenerOutput{}

        mockClient.On("UpdateListener", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateListener(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateResourceConfiguration", func(t *testing.T) {
        input := &vpclattice.UpdateResourceConfigurationInput{}
        output := &vpclattice.UpdateResourceConfigurationOutput{}

        mockClient.On("UpdateResourceConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateResourceConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateResourceGateway", func(t *testing.T) {
        input := &vpclattice.UpdateResourceGatewayInput{}
        output := &vpclattice.UpdateResourceGatewayOutput{}

        mockClient.On("UpdateResourceGateway", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateResourceGateway(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateRule", func(t *testing.T) {
        input := &vpclattice.UpdateRuleInput{}
        output := &vpclattice.UpdateRuleOutput{}

        mockClient.On("UpdateRule", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateRule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateService", func(t *testing.T) {
        input := &vpclattice.UpdateServiceInput{}
        output := &vpclattice.UpdateServiceOutput{}

        mockClient.On("UpdateService", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateService(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateServiceNetwork", func(t *testing.T) {
        input := &vpclattice.UpdateServiceNetworkInput{}
        output := &vpclattice.UpdateServiceNetworkOutput{}

        mockClient.On("UpdateServiceNetwork", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateServiceNetwork(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateServiceNetworkVpcAssociation", func(t *testing.T) {
        input := &vpclattice.UpdateServiceNetworkVpcAssociationInput{}
        output := &vpclattice.UpdateServiceNetworkVpcAssociationOutput{}

        mockClient.On("UpdateServiceNetworkVpcAssociation", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateServiceNetworkVpcAssociation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateTargetGroup", func(t *testing.T) {
        input := &vpclattice.UpdateTargetGroupInput{}
        output := &vpclattice.UpdateTargetGroupOutput{}

        mockClient.On("UpdateTargetGroup", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateTargetGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
