package route53recoverycontrolconfig_test

// tests for the route53recoverycontrolconfig service interface for this ../../../service/route53recoverycontrolconfig/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/route53recoverycontrolconfig"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/route53recoverycontrolconfig/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/route53recoverycontrolconfig/route53recoverycontrolconfig_iface"
	"github.com/stretchr/testify/assert"
)

func TestRoute53recoverycontrolconfigServiceCanBeMocked(t *testing.T) {
	var iface route53recoverycontrolconfig_iface.IClient
	iface = &route53recoverycontrolconfig.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := route53recoverycontrolconfig.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateCluster", func(t *testing.T) {
        input := &route53recoverycontrolconfig.CreateClusterInput{}
        output := &route53recoverycontrolconfig.CreateClusterOutput{}

        mockClient.On("CreateCluster", ctx, input).Return(output, nil)

        result, err := mockClient.CreateCluster(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateControlPanel", func(t *testing.T) {
        input := &route53recoverycontrolconfig.CreateControlPanelInput{}
        output := &route53recoverycontrolconfig.CreateControlPanelOutput{}

        mockClient.On("CreateControlPanel", ctx, input).Return(output, nil)

        result, err := mockClient.CreateControlPanel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateRoutingControl", func(t *testing.T) {
        input := &route53recoverycontrolconfig.CreateRoutingControlInput{}
        output := &route53recoverycontrolconfig.CreateRoutingControlOutput{}

        mockClient.On("CreateRoutingControl", ctx, input).Return(output, nil)

        result, err := mockClient.CreateRoutingControl(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateSafetyRule", func(t *testing.T) {
        input := &route53recoverycontrolconfig.CreateSafetyRuleInput{}
        output := &route53recoverycontrolconfig.CreateSafetyRuleOutput{}

        mockClient.On("CreateSafetyRule", ctx, input).Return(output, nil)

        result, err := mockClient.CreateSafetyRule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteCluster", func(t *testing.T) {
        input := &route53recoverycontrolconfig.DeleteClusterInput{}
        output := &route53recoverycontrolconfig.DeleteClusterOutput{}

        mockClient.On("DeleteCluster", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteCluster(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteControlPanel", func(t *testing.T) {
        input := &route53recoverycontrolconfig.DeleteControlPanelInput{}
        output := &route53recoverycontrolconfig.DeleteControlPanelOutput{}

        mockClient.On("DeleteControlPanel", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteControlPanel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteRoutingControl", func(t *testing.T) {
        input := &route53recoverycontrolconfig.DeleteRoutingControlInput{}
        output := &route53recoverycontrolconfig.DeleteRoutingControlOutput{}

        mockClient.On("DeleteRoutingControl", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteRoutingControl(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteSafetyRule", func(t *testing.T) {
        input := &route53recoverycontrolconfig.DeleteSafetyRuleInput{}
        output := &route53recoverycontrolconfig.DeleteSafetyRuleOutput{}

        mockClient.On("DeleteSafetyRule", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteSafetyRule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeCluster", func(t *testing.T) {
        input := &route53recoverycontrolconfig.DescribeClusterInput{}
        output := &route53recoverycontrolconfig.DescribeClusterOutput{}

        mockClient.On("DescribeCluster", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeCluster(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeControlPanel", func(t *testing.T) {
        input := &route53recoverycontrolconfig.DescribeControlPanelInput{}
        output := &route53recoverycontrolconfig.DescribeControlPanelOutput{}

        mockClient.On("DescribeControlPanel", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeControlPanel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeRoutingControl", func(t *testing.T) {
        input := &route53recoverycontrolconfig.DescribeRoutingControlInput{}
        output := &route53recoverycontrolconfig.DescribeRoutingControlOutput{}

        mockClient.On("DescribeRoutingControl", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeRoutingControl(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeSafetyRule", func(t *testing.T) {
        input := &route53recoverycontrolconfig.DescribeSafetyRuleInput{}
        output := &route53recoverycontrolconfig.DescribeSafetyRuleOutput{}

        mockClient.On("DescribeSafetyRule", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeSafetyRule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetResourcePolicy", func(t *testing.T) {
        input := &route53recoverycontrolconfig.GetResourcePolicyInput{}
        output := &route53recoverycontrolconfig.GetResourcePolicyOutput{}

        mockClient.On("GetResourcePolicy", ctx, input).Return(output, nil)

        result, err := mockClient.GetResourcePolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAssociatedRoute53HealthChecks", func(t *testing.T) {
        input := &route53recoverycontrolconfig.ListAssociatedRoute53HealthChecksInput{}
        output := &route53recoverycontrolconfig.ListAssociatedRoute53HealthChecksOutput{}

        mockClient.On("ListAssociatedRoute53HealthChecks", ctx, input).Return(output, nil)

        result, err := mockClient.ListAssociatedRoute53HealthChecks(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListClusters", func(t *testing.T) {
        input := &route53recoverycontrolconfig.ListClustersInput{}
        output := &route53recoverycontrolconfig.ListClustersOutput{}

        mockClient.On("ListClusters", ctx, input).Return(output, nil)

        result, err := mockClient.ListClusters(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListControlPanels", func(t *testing.T) {
        input := &route53recoverycontrolconfig.ListControlPanelsInput{}
        output := &route53recoverycontrolconfig.ListControlPanelsOutput{}

        mockClient.On("ListControlPanels", ctx, input).Return(output, nil)

        result, err := mockClient.ListControlPanels(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListRoutingControls", func(t *testing.T) {
        input := &route53recoverycontrolconfig.ListRoutingControlsInput{}
        output := &route53recoverycontrolconfig.ListRoutingControlsOutput{}

        mockClient.On("ListRoutingControls", ctx, input).Return(output, nil)

        result, err := mockClient.ListRoutingControls(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListSafetyRules", func(t *testing.T) {
        input := &route53recoverycontrolconfig.ListSafetyRulesInput{}
        output := &route53recoverycontrolconfig.ListSafetyRulesOutput{}

        mockClient.On("ListSafetyRules", ctx, input).Return(output, nil)

        result, err := mockClient.ListSafetyRules(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &route53recoverycontrolconfig.ListTagsForResourceInput{}
        output := &route53recoverycontrolconfig.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &route53recoverycontrolconfig.TagResourceInput{}
        output := &route53recoverycontrolconfig.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &route53recoverycontrolconfig.UntagResourceInput{}
        output := &route53recoverycontrolconfig.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateControlPanel", func(t *testing.T) {
        input := &route53recoverycontrolconfig.UpdateControlPanelInput{}
        output := &route53recoverycontrolconfig.UpdateControlPanelOutput{}

        mockClient.On("UpdateControlPanel", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateControlPanel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateRoutingControl", func(t *testing.T) {
        input := &route53recoverycontrolconfig.UpdateRoutingControlInput{}
        output := &route53recoverycontrolconfig.UpdateRoutingControlOutput{}

        mockClient.On("UpdateRoutingControl", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateRoutingControl(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateSafetyRule", func(t *testing.T) {
        input := &route53recoverycontrolconfig.UpdateSafetyRuleInput{}
        output := &route53recoverycontrolconfig.UpdateSafetyRuleOutput{}

        mockClient.On("UpdateSafetyRule", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateSafetyRule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
