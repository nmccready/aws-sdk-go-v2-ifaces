package elasticloadbalancing_test

// tests for the elasticloadbalancing service interface for this ../../../service/elasticloadbalancing/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/elasticloadbalancing/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/elasticloadbalancing/elasticloadbalancing_iface"
	"github.com/stretchr/testify/assert"
)

func TestElasticloadbalancingServiceCanBeMocked(t *testing.T) {
	var iface elasticloadbalancing_iface.IClient
	iface = &elasticloadbalancing.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := elasticloadbalancing.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAddTags", func(t *testing.T) {
        input := &elasticloadbalancing.AddTagsInput{}
        output := &elasticloadbalancing.AddTagsOutput{}

        mockClient.On("AddTags", ctx, input).Return(output, nil)

        result, err := mockClient.AddTags(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestApplySecurityGroupsToLoadBalancer", func(t *testing.T) {
        input := &elasticloadbalancing.ApplySecurityGroupsToLoadBalancerInput{}
        output := &elasticloadbalancing.ApplySecurityGroupsToLoadBalancerOutput{}

        mockClient.On("ApplySecurityGroupsToLoadBalancer", ctx, input).Return(output, nil)

        result, err := mockClient.ApplySecurityGroupsToLoadBalancer(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAttachLoadBalancerToSubnets", func(t *testing.T) {
        input := &elasticloadbalancing.AttachLoadBalancerToSubnetsInput{}
        output := &elasticloadbalancing.AttachLoadBalancerToSubnetsOutput{}

        mockClient.On("AttachLoadBalancerToSubnets", ctx, input).Return(output, nil)

        result, err := mockClient.AttachLoadBalancerToSubnets(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestConfigureHealthCheck", func(t *testing.T) {
        input := &elasticloadbalancing.ConfigureHealthCheckInput{}
        output := &elasticloadbalancing.ConfigureHealthCheckOutput{}

        mockClient.On("ConfigureHealthCheck", ctx, input).Return(output, nil)

        result, err := mockClient.ConfigureHealthCheck(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateAppCookieStickinessPolicy", func(t *testing.T) {
        input := &elasticloadbalancing.CreateAppCookieStickinessPolicyInput{}
        output := &elasticloadbalancing.CreateAppCookieStickinessPolicyOutput{}

        mockClient.On("CreateAppCookieStickinessPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.CreateAppCookieStickinessPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateLBCookieStickinessPolicy", func(t *testing.T) {
        input := &elasticloadbalancing.CreateLBCookieStickinessPolicyInput{}
        output := &elasticloadbalancing.CreateLBCookieStickinessPolicyOutput{}

        mockClient.On("CreateLBCookieStickinessPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.CreateLBCookieStickinessPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateLoadBalancer", func(t *testing.T) {
        input := &elasticloadbalancing.CreateLoadBalancerInput{}
        output := &elasticloadbalancing.CreateLoadBalancerOutput{}

        mockClient.On("CreateLoadBalancer", ctx, input).Return(output, nil)

        result, err := mockClient.CreateLoadBalancer(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateLoadBalancerListeners", func(t *testing.T) {
        input := &elasticloadbalancing.CreateLoadBalancerListenersInput{}
        output := &elasticloadbalancing.CreateLoadBalancerListenersOutput{}

        mockClient.On("CreateLoadBalancerListeners", ctx, input).Return(output, nil)

        result, err := mockClient.CreateLoadBalancerListeners(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateLoadBalancerPolicy", func(t *testing.T) {
        input := &elasticloadbalancing.CreateLoadBalancerPolicyInput{}
        output := &elasticloadbalancing.CreateLoadBalancerPolicyOutput{}

        mockClient.On("CreateLoadBalancerPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.CreateLoadBalancerPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteLoadBalancer", func(t *testing.T) {
        input := &elasticloadbalancing.DeleteLoadBalancerInput{}
        output := &elasticloadbalancing.DeleteLoadBalancerOutput{}

        mockClient.On("DeleteLoadBalancer", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteLoadBalancer(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteLoadBalancerListeners", func(t *testing.T) {
        input := &elasticloadbalancing.DeleteLoadBalancerListenersInput{}
        output := &elasticloadbalancing.DeleteLoadBalancerListenersOutput{}

        mockClient.On("DeleteLoadBalancerListeners", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteLoadBalancerListeners(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteLoadBalancerPolicy", func(t *testing.T) {
        input := &elasticloadbalancing.DeleteLoadBalancerPolicyInput{}
        output := &elasticloadbalancing.DeleteLoadBalancerPolicyOutput{}

        mockClient.On("DeleteLoadBalancerPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteLoadBalancerPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeregisterInstancesFromLoadBalancer", func(t *testing.T) {
        input := &elasticloadbalancing.DeregisterInstancesFromLoadBalancerInput{}
        output := &elasticloadbalancing.DeregisterInstancesFromLoadBalancerOutput{}

        mockClient.On("DeregisterInstancesFromLoadBalancer", ctx, input).Return(output, nil)

        result, err := mockClient.DeregisterInstancesFromLoadBalancer(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeAccountLimits", func(t *testing.T) {
        input := &elasticloadbalancing.DescribeAccountLimitsInput{}
        output := &elasticloadbalancing.DescribeAccountLimitsOutput{}

        mockClient.On("DescribeAccountLimits", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeAccountLimits(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeInstanceHealth", func(t *testing.T) {
        input := &elasticloadbalancing.DescribeInstanceHealthInput{}
        output := &elasticloadbalancing.DescribeInstanceHealthOutput{}

        mockClient.On("DescribeInstanceHealth", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeInstanceHealth(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeLoadBalancerAttributes", func(t *testing.T) {
        input := &elasticloadbalancing.DescribeLoadBalancerAttributesInput{}
        output := &elasticloadbalancing.DescribeLoadBalancerAttributesOutput{}

        mockClient.On("DescribeLoadBalancerAttributes", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeLoadBalancerAttributes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeLoadBalancerPolicies", func(t *testing.T) {
        input := &elasticloadbalancing.DescribeLoadBalancerPoliciesInput{}
        output := &elasticloadbalancing.DescribeLoadBalancerPoliciesOutput{}

        mockClient.On("DescribeLoadBalancerPolicies", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeLoadBalancerPolicies(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeLoadBalancerPolicyTypes", func(t *testing.T) {
        input := &elasticloadbalancing.DescribeLoadBalancerPolicyTypesInput{}
        output := &elasticloadbalancing.DescribeLoadBalancerPolicyTypesOutput{}

        mockClient.On("DescribeLoadBalancerPolicyTypes", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeLoadBalancerPolicyTypes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeLoadBalancers", func(t *testing.T) {
        input := &elasticloadbalancing.DescribeLoadBalancersInput{}
        output := &elasticloadbalancing.DescribeLoadBalancersOutput{}

        mockClient.On("DescribeLoadBalancers", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeLoadBalancers(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeTags", func(t *testing.T) {
        input := &elasticloadbalancing.DescribeTagsInput{}
        output := &elasticloadbalancing.DescribeTagsOutput{}

        mockClient.On("DescribeTags", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeTags(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDetachLoadBalancerFromSubnets", func(t *testing.T) {
        input := &elasticloadbalancing.DetachLoadBalancerFromSubnetsInput{}
        output := &elasticloadbalancing.DetachLoadBalancerFromSubnetsOutput{}

        mockClient.On("DetachLoadBalancerFromSubnets", ctx, input).Return(output, nil)

        result, err := mockClient.DetachLoadBalancerFromSubnets(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisableAvailabilityZonesForLoadBalancer", func(t *testing.T) {
        input := &elasticloadbalancing.DisableAvailabilityZonesForLoadBalancerInput{}
        output := &elasticloadbalancing.DisableAvailabilityZonesForLoadBalancerOutput{}

        mockClient.On("DisableAvailabilityZonesForLoadBalancer", ctx, input).Return(output, nil)

        result, err := mockClient.DisableAvailabilityZonesForLoadBalancer(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestEnableAvailabilityZonesForLoadBalancer", func(t *testing.T) {
        input := &elasticloadbalancing.EnableAvailabilityZonesForLoadBalancerInput{}
        output := &elasticloadbalancing.EnableAvailabilityZonesForLoadBalancerOutput{}

        mockClient.On("EnableAvailabilityZonesForLoadBalancer", ctx, input).Return(output, nil)

        result, err := mockClient.EnableAvailabilityZonesForLoadBalancer(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyLoadBalancerAttributes", func(t *testing.T) {
        input := &elasticloadbalancing.ModifyLoadBalancerAttributesInput{}
        output := &elasticloadbalancing.ModifyLoadBalancerAttributesOutput{}

        mockClient.On("ModifyLoadBalancerAttributes", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyLoadBalancerAttributes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRegisterInstancesWithLoadBalancer", func(t *testing.T) {
        input := &elasticloadbalancing.RegisterInstancesWithLoadBalancerInput{}
        output := &elasticloadbalancing.RegisterInstancesWithLoadBalancerOutput{}

        mockClient.On("RegisterInstancesWithLoadBalancer", ctx, input).Return(output, nil)

        result, err := mockClient.RegisterInstancesWithLoadBalancer(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRemoveTags", func(t *testing.T) {
        input := &elasticloadbalancing.RemoveTagsInput{}
        output := &elasticloadbalancing.RemoveTagsOutput{}

        mockClient.On("RemoveTags", ctx, input).Return(output, nil)

        result, err := mockClient.RemoveTags(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSetLoadBalancerListenerSSLCertificate", func(t *testing.T) {
        input := &elasticloadbalancing.SetLoadBalancerListenerSSLCertificateInput{}
        output := &elasticloadbalancing.SetLoadBalancerListenerSSLCertificateOutput{}

        mockClient.On("SetLoadBalancerListenerSSLCertificate", ctx, input).Return(output, nil)

        result, err := mockClient.SetLoadBalancerListenerSSLCertificate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSetLoadBalancerPoliciesForBackendServer", func(t *testing.T) {
        input := &elasticloadbalancing.SetLoadBalancerPoliciesForBackendServerInput{}
        output := &elasticloadbalancing.SetLoadBalancerPoliciesForBackendServerOutput{}

        mockClient.On("SetLoadBalancerPoliciesForBackendServer", ctx, input).Return(output, nil)

        result, err := mockClient.SetLoadBalancerPoliciesForBackendServer(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSetLoadBalancerPoliciesOfListener", func(t *testing.T) {
        input := &elasticloadbalancing.SetLoadBalancerPoliciesOfListenerInput{}
        output := &elasticloadbalancing.SetLoadBalancerPoliciesOfListenerOutput{}

        mockClient.On("SetLoadBalancerPoliciesOfListener", ctx, input).Return(output, nil)

        result, err := mockClient.SetLoadBalancerPoliciesOfListener(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
