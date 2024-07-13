package detective_test

// tests for the detective service interface for this ../../../service/detective/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/detective"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/detective/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/detective/detective_iface"
	"github.com/stretchr/testify/assert"
)

func TestDetectiveServiceCanBeMocked(t *testing.T) {
	var iface detective_iface.IClient
	iface = &detective.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := detective.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAcceptInvitation", func(t *testing.T) {
        input := &detective.AcceptInvitationInput{}
        output := &detective.AcceptInvitationOutput{}

        mockClient.On("AcceptInvitation", ctx, input).Return(output, nil)

        result, err := mockClient.AcceptInvitation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchGetGraphMemberDatasources", func(t *testing.T) {
        input := &detective.BatchGetGraphMemberDatasourcesInput{}
        output := &detective.BatchGetGraphMemberDatasourcesOutput{}

        mockClient.On("BatchGetGraphMemberDatasources", ctx, input).Return(output, nil)

        result, err := mockClient.BatchGetGraphMemberDatasources(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchGetMembershipDatasources", func(t *testing.T) {
        input := &detective.BatchGetMembershipDatasourcesInput{}
        output := &detective.BatchGetMembershipDatasourcesOutput{}

        mockClient.On("BatchGetMembershipDatasources", ctx, input).Return(output, nil)

        result, err := mockClient.BatchGetMembershipDatasources(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateGraph", func(t *testing.T) {
        input := &detective.CreateGraphInput{}
        output := &detective.CreateGraphOutput{}

        mockClient.On("CreateGraph", ctx, input).Return(output, nil)

        result, err := mockClient.CreateGraph(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateMembers", func(t *testing.T) {
        input := &detective.CreateMembersInput{}
        output := &detective.CreateMembersOutput{}

        mockClient.On("CreateMembers", ctx, input).Return(output, nil)

        result, err := mockClient.CreateMembers(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteGraph", func(t *testing.T) {
        input := &detective.DeleteGraphInput{}
        output := &detective.DeleteGraphOutput{}

        mockClient.On("DeleteGraph", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteGraph(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteMembers", func(t *testing.T) {
        input := &detective.DeleteMembersInput{}
        output := &detective.DeleteMembersOutput{}

        mockClient.On("DeleteMembers", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteMembers(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeOrganizationConfiguration", func(t *testing.T) {
        input := &detective.DescribeOrganizationConfigurationInput{}
        output := &detective.DescribeOrganizationConfigurationOutput{}

        mockClient.On("DescribeOrganizationConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeOrganizationConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisableOrganizationAdminAccount", func(t *testing.T) {
        input := &detective.DisableOrganizationAdminAccountInput{}
        output := &detective.DisableOrganizationAdminAccountOutput{}

        mockClient.On("DisableOrganizationAdminAccount", ctx, input).Return(output, nil)

        result, err := mockClient.DisableOrganizationAdminAccount(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateMembership", func(t *testing.T) {
        input := &detective.DisassociateMembershipInput{}
        output := &detective.DisassociateMembershipOutput{}

        mockClient.On("DisassociateMembership", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateMembership(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestEnableOrganizationAdminAccount", func(t *testing.T) {
        input := &detective.EnableOrganizationAdminAccountInput{}
        output := &detective.EnableOrganizationAdminAccountOutput{}

        mockClient.On("EnableOrganizationAdminAccount", ctx, input).Return(output, nil)

        result, err := mockClient.EnableOrganizationAdminAccount(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetInvestigation", func(t *testing.T) {
        input := &detective.GetInvestigationInput{}
        output := &detective.GetInvestigationOutput{}

        mockClient.On("GetInvestigation", ctx, input).Return(output, nil)

        result, err := mockClient.GetInvestigation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetMembers", func(t *testing.T) {
        input := &detective.GetMembersInput{}
        output := &detective.GetMembersOutput{}

        mockClient.On("GetMembers", ctx, input).Return(output, nil)

        result, err := mockClient.GetMembers(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDatasourcePackages", func(t *testing.T) {
        input := &detective.ListDatasourcePackagesInput{}
        output := &detective.ListDatasourcePackagesOutput{}

        mockClient.On("ListDatasourcePackages", ctx, input).Return(output, nil)

        result, err := mockClient.ListDatasourcePackages(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListGraphs", func(t *testing.T) {
        input := &detective.ListGraphsInput{}
        output := &detective.ListGraphsOutput{}

        mockClient.On("ListGraphs", ctx, input).Return(output, nil)

        result, err := mockClient.ListGraphs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListIndicators", func(t *testing.T) {
        input := &detective.ListIndicatorsInput{}
        output := &detective.ListIndicatorsOutput{}

        mockClient.On("ListIndicators", ctx, input).Return(output, nil)

        result, err := mockClient.ListIndicators(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListInvestigations", func(t *testing.T) {
        input := &detective.ListInvestigationsInput{}
        output := &detective.ListInvestigationsOutput{}

        mockClient.On("ListInvestigations", ctx, input).Return(output, nil)

        result, err := mockClient.ListInvestigations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListInvitations", func(t *testing.T) {
        input := &detective.ListInvitationsInput{}
        output := &detective.ListInvitationsOutput{}

        mockClient.On("ListInvitations", ctx, input).Return(output, nil)

        result, err := mockClient.ListInvitations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListMembers", func(t *testing.T) {
        input := &detective.ListMembersInput{}
        output := &detective.ListMembersOutput{}

        mockClient.On("ListMembers", ctx, input).Return(output, nil)

        result, err := mockClient.ListMembers(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListOrganizationAdminAccounts", func(t *testing.T) {
        input := &detective.ListOrganizationAdminAccountsInput{}
        output := &detective.ListOrganizationAdminAccountsOutput{}

        mockClient.On("ListOrganizationAdminAccounts", ctx, input).Return(output, nil)

        result, err := mockClient.ListOrganizationAdminAccounts(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &detective.ListTagsForResourceInput{}
        output := &detective.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRejectInvitation", func(t *testing.T) {
        input := &detective.RejectInvitationInput{}
        output := &detective.RejectInvitationOutput{}

        mockClient.On("RejectInvitation", ctx, input).Return(output, nil)

        result, err := mockClient.RejectInvitation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartInvestigation", func(t *testing.T) {
        input := &detective.StartInvestigationInput{}
        output := &detective.StartInvestigationOutput{}

        mockClient.On("StartInvestigation", ctx, input).Return(output, nil)

        result, err := mockClient.StartInvestigation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartMonitoringMember", func(t *testing.T) {
        input := &detective.StartMonitoringMemberInput{}
        output := &detective.StartMonitoringMemberOutput{}

        mockClient.On("StartMonitoringMember", ctx, input).Return(output, nil)

        result, err := mockClient.StartMonitoringMember(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &detective.TagResourceInput{}
        output := &detective.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &detective.UntagResourceInput{}
        output := &detective.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateDatasourcePackages", func(t *testing.T) {
        input := &detective.UpdateDatasourcePackagesInput{}
        output := &detective.UpdateDatasourcePackagesOutput{}

        mockClient.On("UpdateDatasourcePackages", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateDatasourcePackages(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateInvestigationState", func(t *testing.T) {
        input := &detective.UpdateInvestigationStateInput{}
        output := &detective.UpdateInvestigationStateOutput{}

        mockClient.On("UpdateInvestigationState", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateInvestigationState(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateOrganizationConfiguration", func(t *testing.T) {
        input := &detective.UpdateOrganizationConfigurationInput{}
        output := &detective.UpdateOrganizationConfigurationOutput{}

        mockClient.On("UpdateOrganizationConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateOrganizationConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
