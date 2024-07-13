package route53resolver_test

// tests for the route53resolver service interface for this ../../../service/route53resolver/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/route53resolver"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/route53resolver/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/route53resolver/route53resolver_iface"
	"github.com/stretchr/testify/assert"
)

func TestRoute53resolverServiceCanBeMocked(t *testing.T) {
	var iface route53resolver_iface.IClient
	iface = &route53resolver.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := route53resolver.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateFirewallRuleGroup", func(t *testing.T) {
        input := &route53resolver.AssociateFirewallRuleGroupInput{}
        output := &route53resolver.AssociateFirewallRuleGroupOutput{}

        mockClient.On("AssociateFirewallRuleGroup", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateFirewallRuleGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateResolverEndpointIpAddress", func(t *testing.T) {
        input := &route53resolver.AssociateResolverEndpointIpAddressInput{}
        output := &route53resolver.AssociateResolverEndpointIpAddressOutput{}

        mockClient.On("AssociateResolverEndpointIpAddress", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateResolverEndpointIpAddress(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateResolverQueryLogConfig", func(t *testing.T) {
        input := &route53resolver.AssociateResolverQueryLogConfigInput{}
        output := &route53resolver.AssociateResolverQueryLogConfigOutput{}

        mockClient.On("AssociateResolverQueryLogConfig", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateResolverQueryLogConfig(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateResolverRule", func(t *testing.T) {
        input := &route53resolver.AssociateResolverRuleInput{}
        output := &route53resolver.AssociateResolverRuleOutput{}

        mockClient.On("AssociateResolverRule", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateResolverRule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateFirewallDomainList", func(t *testing.T) {
        input := &route53resolver.CreateFirewallDomainListInput{}
        output := &route53resolver.CreateFirewallDomainListOutput{}

        mockClient.On("CreateFirewallDomainList", ctx, input).Return(output, nil)

        result, err := mockClient.CreateFirewallDomainList(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateFirewallRule", func(t *testing.T) {
        input := &route53resolver.CreateFirewallRuleInput{}
        output := &route53resolver.CreateFirewallRuleOutput{}

        mockClient.On("CreateFirewallRule", ctx, input).Return(output, nil)

        result, err := mockClient.CreateFirewallRule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateFirewallRuleGroup", func(t *testing.T) {
        input := &route53resolver.CreateFirewallRuleGroupInput{}
        output := &route53resolver.CreateFirewallRuleGroupOutput{}

        mockClient.On("CreateFirewallRuleGroup", ctx, input).Return(output, nil)

        result, err := mockClient.CreateFirewallRuleGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateOutpostResolver", func(t *testing.T) {
        input := &route53resolver.CreateOutpostResolverInput{}
        output := &route53resolver.CreateOutpostResolverOutput{}

        mockClient.On("CreateOutpostResolver", ctx, input).Return(output, nil)

        result, err := mockClient.CreateOutpostResolver(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateResolverEndpoint", func(t *testing.T) {
        input := &route53resolver.CreateResolverEndpointInput{}
        output := &route53resolver.CreateResolverEndpointOutput{}

        mockClient.On("CreateResolverEndpoint", ctx, input).Return(output, nil)

        result, err := mockClient.CreateResolverEndpoint(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateResolverQueryLogConfig", func(t *testing.T) {
        input := &route53resolver.CreateResolverQueryLogConfigInput{}
        output := &route53resolver.CreateResolverQueryLogConfigOutput{}

        mockClient.On("CreateResolverQueryLogConfig", ctx, input).Return(output, nil)

        result, err := mockClient.CreateResolverQueryLogConfig(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateResolverRule", func(t *testing.T) {
        input := &route53resolver.CreateResolverRuleInput{}
        output := &route53resolver.CreateResolverRuleOutput{}

        mockClient.On("CreateResolverRule", ctx, input).Return(output, nil)

        result, err := mockClient.CreateResolverRule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteFirewallDomainList", func(t *testing.T) {
        input := &route53resolver.DeleteFirewallDomainListInput{}
        output := &route53resolver.DeleteFirewallDomainListOutput{}

        mockClient.On("DeleteFirewallDomainList", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteFirewallDomainList(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteFirewallRule", func(t *testing.T) {
        input := &route53resolver.DeleteFirewallRuleInput{}
        output := &route53resolver.DeleteFirewallRuleOutput{}

        mockClient.On("DeleteFirewallRule", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteFirewallRule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteFirewallRuleGroup", func(t *testing.T) {
        input := &route53resolver.DeleteFirewallRuleGroupInput{}
        output := &route53resolver.DeleteFirewallRuleGroupOutput{}

        mockClient.On("DeleteFirewallRuleGroup", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteFirewallRuleGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteOutpostResolver", func(t *testing.T) {
        input := &route53resolver.DeleteOutpostResolverInput{}
        output := &route53resolver.DeleteOutpostResolverOutput{}

        mockClient.On("DeleteOutpostResolver", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteOutpostResolver(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteResolverEndpoint", func(t *testing.T) {
        input := &route53resolver.DeleteResolverEndpointInput{}
        output := &route53resolver.DeleteResolverEndpointOutput{}

        mockClient.On("DeleteResolverEndpoint", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteResolverEndpoint(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteResolverQueryLogConfig", func(t *testing.T) {
        input := &route53resolver.DeleteResolverQueryLogConfigInput{}
        output := &route53resolver.DeleteResolverQueryLogConfigOutput{}

        mockClient.On("DeleteResolverQueryLogConfig", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteResolverQueryLogConfig(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteResolverRule", func(t *testing.T) {
        input := &route53resolver.DeleteResolverRuleInput{}
        output := &route53resolver.DeleteResolverRuleOutput{}

        mockClient.On("DeleteResolverRule", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteResolverRule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateFirewallRuleGroup", func(t *testing.T) {
        input := &route53resolver.DisassociateFirewallRuleGroupInput{}
        output := &route53resolver.DisassociateFirewallRuleGroupOutput{}

        mockClient.On("DisassociateFirewallRuleGroup", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateFirewallRuleGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateResolverEndpointIpAddress", func(t *testing.T) {
        input := &route53resolver.DisassociateResolverEndpointIpAddressInput{}
        output := &route53resolver.DisassociateResolverEndpointIpAddressOutput{}

        mockClient.On("DisassociateResolverEndpointIpAddress", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateResolverEndpointIpAddress(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateResolverQueryLogConfig", func(t *testing.T) {
        input := &route53resolver.DisassociateResolverQueryLogConfigInput{}
        output := &route53resolver.DisassociateResolverQueryLogConfigOutput{}

        mockClient.On("DisassociateResolverQueryLogConfig", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateResolverQueryLogConfig(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateResolverRule", func(t *testing.T) {
        input := &route53resolver.DisassociateResolverRuleInput{}
        output := &route53resolver.DisassociateResolverRuleOutput{}

        mockClient.On("DisassociateResolverRule", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateResolverRule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetFirewallConfig", func(t *testing.T) {
        input := &route53resolver.GetFirewallConfigInput{}
        output := &route53resolver.GetFirewallConfigOutput{}

        mockClient.On("GetFirewallConfig", ctx, input).Return(output, nil)

        result, err := mockClient.GetFirewallConfig(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetFirewallDomainList", func(t *testing.T) {
        input := &route53resolver.GetFirewallDomainListInput{}
        output := &route53resolver.GetFirewallDomainListOutput{}

        mockClient.On("GetFirewallDomainList", ctx, input).Return(output, nil)

        result, err := mockClient.GetFirewallDomainList(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetFirewallRuleGroup", func(t *testing.T) {
        input := &route53resolver.GetFirewallRuleGroupInput{}
        output := &route53resolver.GetFirewallRuleGroupOutput{}

        mockClient.On("GetFirewallRuleGroup", ctx, input).Return(output, nil)

        result, err := mockClient.GetFirewallRuleGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetFirewallRuleGroupAssociation", func(t *testing.T) {
        input := &route53resolver.GetFirewallRuleGroupAssociationInput{}
        output := &route53resolver.GetFirewallRuleGroupAssociationOutput{}

        mockClient.On("GetFirewallRuleGroupAssociation", ctx, input).Return(output, nil)

        result, err := mockClient.GetFirewallRuleGroupAssociation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetFirewallRuleGroupPolicy", func(t *testing.T) {
        input := &route53resolver.GetFirewallRuleGroupPolicyInput{}
        output := &route53resolver.GetFirewallRuleGroupPolicyOutput{}

        mockClient.On("GetFirewallRuleGroupPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.GetFirewallRuleGroupPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetOutpostResolver", func(t *testing.T) {
        input := &route53resolver.GetOutpostResolverInput{}
        output := &route53resolver.GetOutpostResolverOutput{}

        mockClient.On("GetOutpostResolver", ctx, input).Return(output, nil)

        result, err := mockClient.GetOutpostResolver(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetResolverConfig", func(t *testing.T) {
        input := &route53resolver.GetResolverConfigInput{}
        output := &route53resolver.GetResolverConfigOutput{}

        mockClient.On("GetResolverConfig", ctx, input).Return(output, nil)

        result, err := mockClient.GetResolverConfig(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetResolverDnssecConfig", func(t *testing.T) {
        input := &route53resolver.GetResolverDnssecConfigInput{}
        output := &route53resolver.GetResolverDnssecConfigOutput{}

        mockClient.On("GetResolverDnssecConfig", ctx, input).Return(output, nil)

        result, err := mockClient.GetResolverDnssecConfig(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetResolverEndpoint", func(t *testing.T) {
        input := &route53resolver.GetResolverEndpointInput{}
        output := &route53resolver.GetResolverEndpointOutput{}

        mockClient.On("GetResolverEndpoint", ctx, input).Return(output, nil)

        result, err := mockClient.GetResolverEndpoint(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetResolverQueryLogConfig", func(t *testing.T) {
        input := &route53resolver.GetResolverQueryLogConfigInput{}
        output := &route53resolver.GetResolverQueryLogConfigOutput{}

        mockClient.On("GetResolverQueryLogConfig", ctx, input).Return(output, nil)

        result, err := mockClient.GetResolverQueryLogConfig(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetResolverQueryLogConfigAssociation", func(t *testing.T) {
        input := &route53resolver.GetResolverQueryLogConfigAssociationInput{}
        output := &route53resolver.GetResolverQueryLogConfigAssociationOutput{}

        mockClient.On("GetResolverQueryLogConfigAssociation", ctx, input).Return(output, nil)

        result, err := mockClient.GetResolverQueryLogConfigAssociation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetResolverQueryLogConfigPolicy", func(t *testing.T) {
        input := &route53resolver.GetResolverQueryLogConfigPolicyInput{}
        output := &route53resolver.GetResolverQueryLogConfigPolicyOutput{}

        mockClient.On("GetResolverQueryLogConfigPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.GetResolverQueryLogConfigPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetResolverRule", func(t *testing.T) {
        input := &route53resolver.GetResolverRuleInput{}
        output := &route53resolver.GetResolverRuleOutput{}

        mockClient.On("GetResolverRule", ctx, input).Return(output, nil)

        result, err := mockClient.GetResolverRule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetResolverRuleAssociation", func(t *testing.T) {
        input := &route53resolver.GetResolverRuleAssociationInput{}
        output := &route53resolver.GetResolverRuleAssociationOutput{}

        mockClient.On("GetResolverRuleAssociation", ctx, input).Return(output, nil)

        result, err := mockClient.GetResolverRuleAssociation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetResolverRulePolicy", func(t *testing.T) {
        input := &route53resolver.GetResolverRulePolicyInput{}
        output := &route53resolver.GetResolverRulePolicyOutput{}

        mockClient.On("GetResolverRulePolicy", ctx, input).Return(output, nil)

        result, err := mockClient.GetResolverRulePolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestImportFirewallDomains", func(t *testing.T) {
        input := &route53resolver.ImportFirewallDomainsInput{}
        output := &route53resolver.ImportFirewallDomainsOutput{}

        mockClient.On("ImportFirewallDomains", ctx, input).Return(output, nil)

        result, err := mockClient.ImportFirewallDomains(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListFirewallConfigs", func(t *testing.T) {
        input := &route53resolver.ListFirewallConfigsInput{}
        output := &route53resolver.ListFirewallConfigsOutput{}

        mockClient.On("ListFirewallConfigs", ctx, input).Return(output, nil)

        result, err := mockClient.ListFirewallConfigs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListFirewallDomainLists", func(t *testing.T) {
        input := &route53resolver.ListFirewallDomainListsInput{}
        output := &route53resolver.ListFirewallDomainListsOutput{}

        mockClient.On("ListFirewallDomainLists", ctx, input).Return(output, nil)

        result, err := mockClient.ListFirewallDomainLists(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListFirewallDomains", func(t *testing.T) {
        input := &route53resolver.ListFirewallDomainsInput{}
        output := &route53resolver.ListFirewallDomainsOutput{}

        mockClient.On("ListFirewallDomains", ctx, input).Return(output, nil)

        result, err := mockClient.ListFirewallDomains(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListFirewallRuleGroupAssociations", func(t *testing.T) {
        input := &route53resolver.ListFirewallRuleGroupAssociationsInput{}
        output := &route53resolver.ListFirewallRuleGroupAssociationsOutput{}

        mockClient.On("ListFirewallRuleGroupAssociations", ctx, input).Return(output, nil)

        result, err := mockClient.ListFirewallRuleGroupAssociations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListFirewallRuleGroups", func(t *testing.T) {
        input := &route53resolver.ListFirewallRuleGroupsInput{}
        output := &route53resolver.ListFirewallRuleGroupsOutput{}

        mockClient.On("ListFirewallRuleGroups", ctx, input).Return(output, nil)

        result, err := mockClient.ListFirewallRuleGroups(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListFirewallRules", func(t *testing.T) {
        input := &route53resolver.ListFirewallRulesInput{}
        output := &route53resolver.ListFirewallRulesOutput{}

        mockClient.On("ListFirewallRules", ctx, input).Return(output, nil)

        result, err := mockClient.ListFirewallRules(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListOutpostResolvers", func(t *testing.T) {
        input := &route53resolver.ListOutpostResolversInput{}
        output := &route53resolver.ListOutpostResolversOutput{}

        mockClient.On("ListOutpostResolvers", ctx, input).Return(output, nil)

        result, err := mockClient.ListOutpostResolvers(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListResolverConfigs", func(t *testing.T) {
        input := &route53resolver.ListResolverConfigsInput{}
        output := &route53resolver.ListResolverConfigsOutput{}

        mockClient.On("ListResolverConfigs", ctx, input).Return(output, nil)

        result, err := mockClient.ListResolverConfigs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListResolverDnssecConfigs", func(t *testing.T) {
        input := &route53resolver.ListResolverDnssecConfigsInput{}
        output := &route53resolver.ListResolverDnssecConfigsOutput{}

        mockClient.On("ListResolverDnssecConfigs", ctx, input).Return(output, nil)

        result, err := mockClient.ListResolverDnssecConfigs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListResolverEndpointIpAddresses", func(t *testing.T) {
        input := &route53resolver.ListResolverEndpointIpAddressesInput{}
        output := &route53resolver.ListResolverEndpointIpAddressesOutput{}

        mockClient.On("ListResolverEndpointIpAddresses", ctx, input).Return(output, nil)

        result, err := mockClient.ListResolverEndpointIpAddresses(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListResolverEndpoints", func(t *testing.T) {
        input := &route53resolver.ListResolverEndpointsInput{}
        output := &route53resolver.ListResolverEndpointsOutput{}

        mockClient.On("ListResolverEndpoints", ctx, input).Return(output, nil)

        result, err := mockClient.ListResolverEndpoints(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListResolverQueryLogConfigAssociations", func(t *testing.T) {
        input := &route53resolver.ListResolverQueryLogConfigAssociationsInput{}
        output := &route53resolver.ListResolverQueryLogConfigAssociationsOutput{}

        mockClient.On("ListResolverQueryLogConfigAssociations", ctx, input).Return(output, nil)

        result, err := mockClient.ListResolverQueryLogConfigAssociations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListResolverQueryLogConfigs", func(t *testing.T) {
        input := &route53resolver.ListResolverQueryLogConfigsInput{}
        output := &route53resolver.ListResolverQueryLogConfigsOutput{}

        mockClient.On("ListResolverQueryLogConfigs", ctx, input).Return(output, nil)

        result, err := mockClient.ListResolverQueryLogConfigs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListResolverRuleAssociations", func(t *testing.T) {
        input := &route53resolver.ListResolverRuleAssociationsInput{}
        output := &route53resolver.ListResolverRuleAssociationsOutput{}

        mockClient.On("ListResolverRuleAssociations", ctx, input).Return(output, nil)

        result, err := mockClient.ListResolverRuleAssociations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListResolverRules", func(t *testing.T) {
        input := &route53resolver.ListResolverRulesInput{}
        output := &route53resolver.ListResolverRulesOutput{}

        mockClient.On("ListResolverRules", ctx, input).Return(output, nil)

        result, err := mockClient.ListResolverRules(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &route53resolver.ListTagsForResourceInput{}
        output := &route53resolver.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutFirewallRuleGroupPolicy", func(t *testing.T) {
        input := &route53resolver.PutFirewallRuleGroupPolicyInput{}
        output := &route53resolver.PutFirewallRuleGroupPolicyOutput{}

        mockClient.On("PutFirewallRuleGroupPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.PutFirewallRuleGroupPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutResolverQueryLogConfigPolicy", func(t *testing.T) {
        input := &route53resolver.PutResolverQueryLogConfigPolicyInput{}
        output := &route53resolver.PutResolverQueryLogConfigPolicyOutput{}

        mockClient.On("PutResolverQueryLogConfigPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.PutResolverQueryLogConfigPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutResolverRulePolicy", func(t *testing.T) {
        input := &route53resolver.PutResolverRulePolicyInput{}
        output := &route53resolver.PutResolverRulePolicyOutput{}

        mockClient.On("PutResolverRulePolicy", ctx, input).Return(output, nil)

        result, err := mockClient.PutResolverRulePolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &route53resolver.TagResourceInput{}
        output := &route53resolver.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &route53resolver.UntagResourceInput{}
        output := &route53resolver.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateFirewallConfig", func(t *testing.T) {
        input := &route53resolver.UpdateFirewallConfigInput{}
        output := &route53resolver.UpdateFirewallConfigOutput{}

        mockClient.On("UpdateFirewallConfig", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateFirewallConfig(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateFirewallDomains", func(t *testing.T) {
        input := &route53resolver.UpdateFirewallDomainsInput{}
        output := &route53resolver.UpdateFirewallDomainsOutput{}

        mockClient.On("UpdateFirewallDomains", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateFirewallDomains(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateFirewallRule", func(t *testing.T) {
        input := &route53resolver.UpdateFirewallRuleInput{}
        output := &route53resolver.UpdateFirewallRuleOutput{}

        mockClient.On("UpdateFirewallRule", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateFirewallRule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateFirewallRuleGroupAssociation", func(t *testing.T) {
        input := &route53resolver.UpdateFirewallRuleGroupAssociationInput{}
        output := &route53resolver.UpdateFirewallRuleGroupAssociationOutput{}

        mockClient.On("UpdateFirewallRuleGroupAssociation", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateFirewallRuleGroupAssociation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateOutpostResolver", func(t *testing.T) {
        input := &route53resolver.UpdateOutpostResolverInput{}
        output := &route53resolver.UpdateOutpostResolverOutput{}

        mockClient.On("UpdateOutpostResolver", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateOutpostResolver(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateResolverConfig", func(t *testing.T) {
        input := &route53resolver.UpdateResolverConfigInput{}
        output := &route53resolver.UpdateResolverConfigOutput{}

        mockClient.On("UpdateResolverConfig", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateResolverConfig(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateResolverDnssecConfig", func(t *testing.T) {
        input := &route53resolver.UpdateResolverDnssecConfigInput{}
        output := &route53resolver.UpdateResolverDnssecConfigOutput{}

        mockClient.On("UpdateResolverDnssecConfig", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateResolverDnssecConfig(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateResolverEndpoint", func(t *testing.T) {
        input := &route53resolver.UpdateResolverEndpointInput{}
        output := &route53resolver.UpdateResolverEndpointOutput{}

        mockClient.On("UpdateResolverEndpoint", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateResolverEndpoint(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateResolverRule", func(t *testing.T) {
        input := &route53resolver.UpdateResolverRuleInput{}
        output := &route53resolver.UpdateResolverRuleOutput{}

        mockClient.On("UpdateResolverRule", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateResolverRule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
