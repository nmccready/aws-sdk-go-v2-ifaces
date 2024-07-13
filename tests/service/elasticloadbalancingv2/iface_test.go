package elasticloadbalancingv2_test

// tests for the elasticloadbalancingv2 service interface for this ../../../service/elasticloadbalancingv2/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/elasticloadbalancingv2/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/elasticloadbalancingv2/elasticloadbalancingv2_iface"
	"github.com/stretchr/testify/assert"
)

func TestElasticloadbalancingv2ServiceCanBeMocked(t *testing.T) {
	var iface elasticloadbalancingv2_iface.IClient
	iface = &elasticloadbalancingv2.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := elasticloadbalancingv2.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAddListenerCertificates", func(t *testing.T) {
        input := &elasticloadbalancingv2.AddListenerCertificatesInput{}
        output := &elasticloadbalancingv2.AddListenerCertificatesOutput{}

        mockClient.On("AddListenerCertificates", ctx, input).Return(output, nil)

        result, err := mockClient.AddListenerCertificates(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAddTags", func(t *testing.T) {
        input := &elasticloadbalancingv2.AddTagsInput{}
        output := &elasticloadbalancingv2.AddTagsOutput{}

        mockClient.On("AddTags", ctx, input).Return(output, nil)

        result, err := mockClient.AddTags(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAddTrustStoreRevocations", func(t *testing.T) {
        input := &elasticloadbalancingv2.AddTrustStoreRevocationsInput{}
        output := &elasticloadbalancingv2.AddTrustStoreRevocationsOutput{}

        mockClient.On("AddTrustStoreRevocations", ctx, input).Return(output, nil)

        result, err := mockClient.AddTrustStoreRevocations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateListener", func(t *testing.T) {
        input := &elasticloadbalancingv2.CreateListenerInput{}
        output := &elasticloadbalancingv2.CreateListenerOutput{}

        mockClient.On("CreateListener", ctx, input).Return(output, nil)

        result, err := mockClient.CreateListener(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateLoadBalancer", func(t *testing.T) {
        input := &elasticloadbalancingv2.CreateLoadBalancerInput{}
        output := &elasticloadbalancingv2.CreateLoadBalancerOutput{}

        mockClient.On("CreateLoadBalancer", ctx, input).Return(output, nil)

        result, err := mockClient.CreateLoadBalancer(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateRule", func(t *testing.T) {
        input := &elasticloadbalancingv2.CreateRuleInput{}
        output := &elasticloadbalancingv2.CreateRuleOutput{}

        mockClient.On("CreateRule", ctx, input).Return(output, nil)

        result, err := mockClient.CreateRule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateTargetGroup", func(t *testing.T) {
        input := &elasticloadbalancingv2.CreateTargetGroupInput{}
        output := &elasticloadbalancingv2.CreateTargetGroupOutput{}

        mockClient.On("CreateTargetGroup", ctx, input).Return(output, nil)

        result, err := mockClient.CreateTargetGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateTrustStore", func(t *testing.T) {
        input := &elasticloadbalancingv2.CreateTrustStoreInput{}
        output := &elasticloadbalancingv2.CreateTrustStoreOutput{}

        mockClient.On("CreateTrustStore", ctx, input).Return(output, nil)

        result, err := mockClient.CreateTrustStore(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteListener", func(t *testing.T) {
        input := &elasticloadbalancingv2.DeleteListenerInput{}
        output := &elasticloadbalancingv2.DeleteListenerOutput{}

        mockClient.On("DeleteListener", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteListener(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteLoadBalancer", func(t *testing.T) {
        input := &elasticloadbalancingv2.DeleteLoadBalancerInput{}
        output := &elasticloadbalancingv2.DeleteLoadBalancerOutput{}

        mockClient.On("DeleteLoadBalancer", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteLoadBalancer(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteRule", func(t *testing.T) {
        input := &elasticloadbalancingv2.DeleteRuleInput{}
        output := &elasticloadbalancingv2.DeleteRuleOutput{}

        mockClient.On("DeleteRule", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteRule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteTargetGroup", func(t *testing.T) {
        input := &elasticloadbalancingv2.DeleteTargetGroupInput{}
        output := &elasticloadbalancingv2.DeleteTargetGroupOutput{}

        mockClient.On("DeleteTargetGroup", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteTargetGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteTrustStore", func(t *testing.T) {
        input := &elasticloadbalancingv2.DeleteTrustStoreInput{}
        output := &elasticloadbalancingv2.DeleteTrustStoreOutput{}

        mockClient.On("DeleteTrustStore", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteTrustStore(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeregisterTargets", func(t *testing.T) {
        input := &elasticloadbalancingv2.DeregisterTargetsInput{}
        output := &elasticloadbalancingv2.DeregisterTargetsOutput{}

        mockClient.On("DeregisterTargets", ctx, input).Return(output, nil)

        result, err := mockClient.DeregisterTargets(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeAccountLimits", func(t *testing.T) {
        input := &elasticloadbalancingv2.DescribeAccountLimitsInput{}
        output := &elasticloadbalancingv2.DescribeAccountLimitsOutput{}

        mockClient.On("DescribeAccountLimits", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeAccountLimits(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeListenerCertificates", func(t *testing.T) {
        input := &elasticloadbalancingv2.DescribeListenerCertificatesInput{}
        output := &elasticloadbalancingv2.DescribeListenerCertificatesOutput{}

        mockClient.On("DescribeListenerCertificates", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeListenerCertificates(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeListeners", func(t *testing.T) {
        input := &elasticloadbalancingv2.DescribeListenersInput{}
        output := &elasticloadbalancingv2.DescribeListenersOutput{}

        mockClient.On("DescribeListeners", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeListeners(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeLoadBalancerAttributes", func(t *testing.T) {
        input := &elasticloadbalancingv2.DescribeLoadBalancerAttributesInput{}
        output := &elasticloadbalancingv2.DescribeLoadBalancerAttributesOutput{}

        mockClient.On("DescribeLoadBalancerAttributes", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeLoadBalancerAttributes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeLoadBalancers", func(t *testing.T) {
        input := &elasticloadbalancingv2.DescribeLoadBalancersInput{}
        output := &elasticloadbalancingv2.DescribeLoadBalancersOutput{}

        mockClient.On("DescribeLoadBalancers", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeLoadBalancers(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeRules", func(t *testing.T) {
        input := &elasticloadbalancingv2.DescribeRulesInput{}
        output := &elasticloadbalancingv2.DescribeRulesOutput{}

        mockClient.On("DescribeRules", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeRules(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeSSLPolicies", func(t *testing.T) {
        input := &elasticloadbalancingv2.DescribeSSLPoliciesInput{}
        output := &elasticloadbalancingv2.DescribeSSLPoliciesOutput{}

        mockClient.On("DescribeSSLPolicies", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeSSLPolicies(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeTags", func(t *testing.T) {
        input := &elasticloadbalancingv2.DescribeTagsInput{}
        output := &elasticloadbalancingv2.DescribeTagsOutput{}

        mockClient.On("DescribeTags", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeTags(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeTargetGroupAttributes", func(t *testing.T) {
        input := &elasticloadbalancingv2.DescribeTargetGroupAttributesInput{}
        output := &elasticloadbalancingv2.DescribeTargetGroupAttributesOutput{}

        mockClient.On("DescribeTargetGroupAttributes", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeTargetGroupAttributes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeTargetGroups", func(t *testing.T) {
        input := &elasticloadbalancingv2.DescribeTargetGroupsInput{}
        output := &elasticloadbalancingv2.DescribeTargetGroupsOutput{}

        mockClient.On("DescribeTargetGroups", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeTargetGroups(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeTargetHealth", func(t *testing.T) {
        input := &elasticloadbalancingv2.DescribeTargetHealthInput{}
        output := &elasticloadbalancingv2.DescribeTargetHealthOutput{}

        mockClient.On("DescribeTargetHealth", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeTargetHealth(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeTrustStoreAssociations", func(t *testing.T) {
        input := &elasticloadbalancingv2.DescribeTrustStoreAssociationsInput{}
        output := &elasticloadbalancingv2.DescribeTrustStoreAssociationsOutput{}

        mockClient.On("DescribeTrustStoreAssociations", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeTrustStoreAssociations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeTrustStoreRevocations", func(t *testing.T) {
        input := &elasticloadbalancingv2.DescribeTrustStoreRevocationsInput{}
        output := &elasticloadbalancingv2.DescribeTrustStoreRevocationsOutput{}

        mockClient.On("DescribeTrustStoreRevocations", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeTrustStoreRevocations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeTrustStores", func(t *testing.T) {
        input := &elasticloadbalancingv2.DescribeTrustStoresInput{}
        output := &elasticloadbalancingv2.DescribeTrustStoresOutput{}

        mockClient.On("DescribeTrustStores", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeTrustStores(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetTrustStoreCaCertificatesBundle", func(t *testing.T) {
        input := &elasticloadbalancingv2.GetTrustStoreCaCertificatesBundleInput{}
        output := &elasticloadbalancingv2.GetTrustStoreCaCertificatesBundleOutput{}

        mockClient.On("GetTrustStoreCaCertificatesBundle", ctx, input).Return(output, nil)

        result, err := mockClient.GetTrustStoreCaCertificatesBundle(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetTrustStoreRevocationContent", func(t *testing.T) {
        input := &elasticloadbalancingv2.GetTrustStoreRevocationContentInput{}
        output := &elasticloadbalancingv2.GetTrustStoreRevocationContentOutput{}

        mockClient.On("GetTrustStoreRevocationContent", ctx, input).Return(output, nil)

        result, err := mockClient.GetTrustStoreRevocationContent(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyListener", func(t *testing.T) {
        input := &elasticloadbalancingv2.ModifyListenerInput{}
        output := &elasticloadbalancingv2.ModifyListenerOutput{}

        mockClient.On("ModifyListener", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyListener(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyLoadBalancerAttributes", func(t *testing.T) {
        input := &elasticloadbalancingv2.ModifyLoadBalancerAttributesInput{}
        output := &elasticloadbalancingv2.ModifyLoadBalancerAttributesOutput{}

        mockClient.On("ModifyLoadBalancerAttributes", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyLoadBalancerAttributes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyRule", func(t *testing.T) {
        input := &elasticloadbalancingv2.ModifyRuleInput{}
        output := &elasticloadbalancingv2.ModifyRuleOutput{}

        mockClient.On("ModifyRule", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyRule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyTargetGroup", func(t *testing.T) {
        input := &elasticloadbalancingv2.ModifyTargetGroupInput{}
        output := &elasticloadbalancingv2.ModifyTargetGroupOutput{}

        mockClient.On("ModifyTargetGroup", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyTargetGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyTargetGroupAttributes", func(t *testing.T) {
        input := &elasticloadbalancingv2.ModifyTargetGroupAttributesInput{}
        output := &elasticloadbalancingv2.ModifyTargetGroupAttributesOutput{}

        mockClient.On("ModifyTargetGroupAttributes", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyTargetGroupAttributes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyTrustStore", func(t *testing.T) {
        input := &elasticloadbalancingv2.ModifyTrustStoreInput{}
        output := &elasticloadbalancingv2.ModifyTrustStoreOutput{}

        mockClient.On("ModifyTrustStore", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyTrustStore(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRegisterTargets", func(t *testing.T) {
        input := &elasticloadbalancingv2.RegisterTargetsInput{}
        output := &elasticloadbalancingv2.RegisterTargetsOutput{}

        mockClient.On("RegisterTargets", ctx, input).Return(output, nil)

        result, err := mockClient.RegisterTargets(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRemoveListenerCertificates", func(t *testing.T) {
        input := &elasticloadbalancingv2.RemoveListenerCertificatesInput{}
        output := &elasticloadbalancingv2.RemoveListenerCertificatesOutput{}

        mockClient.On("RemoveListenerCertificates", ctx, input).Return(output, nil)

        result, err := mockClient.RemoveListenerCertificates(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRemoveTags", func(t *testing.T) {
        input := &elasticloadbalancingv2.RemoveTagsInput{}
        output := &elasticloadbalancingv2.RemoveTagsOutput{}

        mockClient.On("RemoveTags", ctx, input).Return(output, nil)

        result, err := mockClient.RemoveTags(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRemoveTrustStoreRevocations", func(t *testing.T) {
        input := &elasticloadbalancingv2.RemoveTrustStoreRevocationsInput{}
        output := &elasticloadbalancingv2.RemoveTrustStoreRevocationsOutput{}

        mockClient.On("RemoveTrustStoreRevocations", ctx, input).Return(output, nil)

        result, err := mockClient.RemoveTrustStoreRevocations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSetIpAddressType", func(t *testing.T) {
        input := &elasticloadbalancingv2.SetIpAddressTypeInput{}
        output := &elasticloadbalancingv2.SetIpAddressTypeOutput{}

        mockClient.On("SetIpAddressType", ctx, input).Return(output, nil)

        result, err := mockClient.SetIpAddressType(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSetRulePriorities", func(t *testing.T) {
        input := &elasticloadbalancingv2.SetRulePrioritiesInput{}
        output := &elasticloadbalancingv2.SetRulePrioritiesOutput{}

        mockClient.On("SetRulePriorities", ctx, input).Return(output, nil)

        result, err := mockClient.SetRulePriorities(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSetSecurityGroups", func(t *testing.T) {
        input := &elasticloadbalancingv2.SetSecurityGroupsInput{}
        output := &elasticloadbalancingv2.SetSecurityGroupsOutput{}

        mockClient.On("SetSecurityGroups", ctx, input).Return(output, nil)

        result, err := mockClient.SetSecurityGroups(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSetSubnets", func(t *testing.T) {
        input := &elasticloadbalancingv2.SetSubnetsInput{}
        output := &elasticloadbalancingv2.SetSubnetsOutput{}

        mockClient.On("SetSubnets", ctx, input).Return(output, nil)

        result, err := mockClient.SetSubnets(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
