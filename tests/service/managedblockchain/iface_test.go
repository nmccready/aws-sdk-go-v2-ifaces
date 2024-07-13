package managedblockchain_test

// tests for the managedblockchain service interface for this ../../../service/managedblockchain/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/managedblockchain"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/managedblockchain/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/managedblockchain/managedblockchain_iface"
	"github.com/stretchr/testify/assert"
)

func TestManagedblockchainServiceCanBeMocked(t *testing.T) {
	var iface managedblockchain_iface.IClient
	iface = &managedblockchain.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := managedblockchain.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateAccessor", func(t *testing.T) {
        input := &managedblockchain.CreateAccessorInput{}
        output := &managedblockchain.CreateAccessorOutput{}

        mockClient.On("CreateAccessor", ctx, input).Return(output, nil)

        result, err := mockClient.CreateAccessor(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateMember", func(t *testing.T) {
        input := &managedblockchain.CreateMemberInput{}
        output := &managedblockchain.CreateMemberOutput{}

        mockClient.On("CreateMember", ctx, input).Return(output, nil)

        result, err := mockClient.CreateMember(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateNetwork", func(t *testing.T) {
        input := &managedblockchain.CreateNetworkInput{}
        output := &managedblockchain.CreateNetworkOutput{}

        mockClient.On("CreateNetwork", ctx, input).Return(output, nil)

        result, err := mockClient.CreateNetwork(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateNode", func(t *testing.T) {
        input := &managedblockchain.CreateNodeInput{}
        output := &managedblockchain.CreateNodeOutput{}

        mockClient.On("CreateNode", ctx, input).Return(output, nil)

        result, err := mockClient.CreateNode(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateProposal", func(t *testing.T) {
        input := &managedblockchain.CreateProposalInput{}
        output := &managedblockchain.CreateProposalOutput{}

        mockClient.On("CreateProposal", ctx, input).Return(output, nil)

        result, err := mockClient.CreateProposal(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteAccessor", func(t *testing.T) {
        input := &managedblockchain.DeleteAccessorInput{}
        output := &managedblockchain.DeleteAccessorOutput{}

        mockClient.On("DeleteAccessor", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteAccessor(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteMember", func(t *testing.T) {
        input := &managedblockchain.DeleteMemberInput{}
        output := &managedblockchain.DeleteMemberOutput{}

        mockClient.On("DeleteMember", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteMember(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteNode", func(t *testing.T) {
        input := &managedblockchain.DeleteNodeInput{}
        output := &managedblockchain.DeleteNodeOutput{}

        mockClient.On("DeleteNode", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteNode(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetAccessor", func(t *testing.T) {
        input := &managedblockchain.GetAccessorInput{}
        output := &managedblockchain.GetAccessorOutput{}

        mockClient.On("GetAccessor", ctx, input).Return(output, nil)

        result, err := mockClient.GetAccessor(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetMember", func(t *testing.T) {
        input := &managedblockchain.GetMemberInput{}
        output := &managedblockchain.GetMemberOutput{}

        mockClient.On("GetMember", ctx, input).Return(output, nil)

        result, err := mockClient.GetMember(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetNetwork", func(t *testing.T) {
        input := &managedblockchain.GetNetworkInput{}
        output := &managedblockchain.GetNetworkOutput{}

        mockClient.On("GetNetwork", ctx, input).Return(output, nil)

        result, err := mockClient.GetNetwork(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetNode", func(t *testing.T) {
        input := &managedblockchain.GetNodeInput{}
        output := &managedblockchain.GetNodeOutput{}

        mockClient.On("GetNode", ctx, input).Return(output, nil)

        result, err := mockClient.GetNode(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetProposal", func(t *testing.T) {
        input := &managedblockchain.GetProposalInput{}
        output := &managedblockchain.GetProposalOutput{}

        mockClient.On("GetProposal", ctx, input).Return(output, nil)

        result, err := mockClient.GetProposal(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAccessors", func(t *testing.T) {
        input := &managedblockchain.ListAccessorsInput{}
        output := &managedblockchain.ListAccessorsOutput{}

        mockClient.On("ListAccessors", ctx, input).Return(output, nil)

        result, err := mockClient.ListAccessors(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListInvitations", func(t *testing.T) {
        input := &managedblockchain.ListInvitationsInput{}
        output := &managedblockchain.ListInvitationsOutput{}

        mockClient.On("ListInvitations", ctx, input).Return(output, nil)

        result, err := mockClient.ListInvitations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListMembers", func(t *testing.T) {
        input := &managedblockchain.ListMembersInput{}
        output := &managedblockchain.ListMembersOutput{}

        mockClient.On("ListMembers", ctx, input).Return(output, nil)

        result, err := mockClient.ListMembers(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListNetworks", func(t *testing.T) {
        input := &managedblockchain.ListNetworksInput{}
        output := &managedblockchain.ListNetworksOutput{}

        mockClient.On("ListNetworks", ctx, input).Return(output, nil)

        result, err := mockClient.ListNetworks(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListNodes", func(t *testing.T) {
        input := &managedblockchain.ListNodesInput{}
        output := &managedblockchain.ListNodesOutput{}

        mockClient.On("ListNodes", ctx, input).Return(output, nil)

        result, err := mockClient.ListNodes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListProposalVotes", func(t *testing.T) {
        input := &managedblockchain.ListProposalVotesInput{}
        output := &managedblockchain.ListProposalVotesOutput{}

        mockClient.On("ListProposalVotes", ctx, input).Return(output, nil)

        result, err := mockClient.ListProposalVotes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListProposals", func(t *testing.T) {
        input := &managedblockchain.ListProposalsInput{}
        output := &managedblockchain.ListProposalsOutput{}

        mockClient.On("ListProposals", ctx, input).Return(output, nil)

        result, err := mockClient.ListProposals(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &managedblockchain.ListTagsForResourceInput{}
        output := &managedblockchain.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRejectInvitation", func(t *testing.T) {
        input := &managedblockchain.RejectInvitationInput{}
        output := &managedblockchain.RejectInvitationOutput{}

        mockClient.On("RejectInvitation", ctx, input).Return(output, nil)

        result, err := mockClient.RejectInvitation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &managedblockchain.TagResourceInput{}
        output := &managedblockchain.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &managedblockchain.UntagResourceInput{}
        output := &managedblockchain.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateMember", func(t *testing.T) {
        input := &managedblockchain.UpdateMemberInput{}
        output := &managedblockchain.UpdateMemberOutput{}

        mockClient.On("UpdateMember", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateMember(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateNode", func(t *testing.T) {
        input := &managedblockchain.UpdateNodeInput{}
        output := &managedblockchain.UpdateNodeOutput{}

        mockClient.On("UpdateNode", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateNode(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestVoteOnProposal", func(t *testing.T) {
        input := &managedblockchain.VoteOnProposalInput{}
        output := &managedblockchain.VoteOnProposalOutput{}

        mockClient.On("VoteOnProposal", ctx, input).Return(output, nil)

        result, err := mockClient.VoteOnProposal(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
