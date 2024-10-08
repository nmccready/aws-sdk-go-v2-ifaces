
package managedblockchain_iface

// This file was generated by scripts/gen_ifaces.sh do not edit it manually
import (
    "context"
    . "github.com/aws/aws-sdk-go-v2/service/managedblockchain"
)

// IClient defines the interface for managedblockchain
type IClient interface {
 Options() Options 
 CreateAccessor(ctx context.Context, params *CreateAccessorInput, optFns ...func(*Options)) (*CreateAccessorOutput, error) 
 CreateMember(ctx context.Context, params *CreateMemberInput, optFns ...func(*Options)) (*CreateMemberOutput, error) 
 CreateNetwork(ctx context.Context, params *CreateNetworkInput, optFns ...func(*Options)) (*CreateNetworkOutput, error) 
 CreateNode(ctx context.Context, params *CreateNodeInput, optFns ...func(*Options)) (*CreateNodeOutput, error) 
 CreateProposal(ctx context.Context, params *CreateProposalInput, optFns ...func(*Options)) (*CreateProposalOutput, error) 
 DeleteAccessor(ctx context.Context, params *DeleteAccessorInput, optFns ...func(*Options)) (*DeleteAccessorOutput, error) 
 DeleteMember(ctx context.Context, params *DeleteMemberInput, optFns ...func(*Options)) (*DeleteMemberOutput, error) 
 DeleteNode(ctx context.Context, params *DeleteNodeInput, optFns ...func(*Options)) (*DeleteNodeOutput, error) 
 GetAccessor(ctx context.Context, params *GetAccessorInput, optFns ...func(*Options)) (*GetAccessorOutput, error) 
 GetMember(ctx context.Context, params *GetMemberInput, optFns ...func(*Options)) (*GetMemberOutput, error) 
 GetNetwork(ctx context.Context, params *GetNetworkInput, optFns ...func(*Options)) (*GetNetworkOutput, error) 
 GetNode(ctx context.Context, params *GetNodeInput, optFns ...func(*Options)) (*GetNodeOutput, error) 
 GetProposal(ctx context.Context, params *GetProposalInput, optFns ...func(*Options)) (*GetProposalOutput, error) 
 ListAccessors(ctx context.Context, params *ListAccessorsInput, optFns ...func(*Options)) (*ListAccessorsOutput, error) 
 ListInvitations(ctx context.Context, params *ListInvitationsInput, optFns ...func(*Options)) (*ListInvitationsOutput, error) 
 ListMembers(ctx context.Context, params *ListMembersInput, optFns ...func(*Options)) (*ListMembersOutput, error) 
 ListNetworks(ctx context.Context, params *ListNetworksInput, optFns ...func(*Options)) (*ListNetworksOutput, error) 
 ListNodes(ctx context.Context, params *ListNodesInput, optFns ...func(*Options)) (*ListNodesOutput, error) 
 ListProposalVotes(ctx context.Context, params *ListProposalVotesInput, optFns ...func(*Options)) (*ListProposalVotesOutput, error) 
 ListProposals(ctx context.Context, params *ListProposalsInput, optFns ...func(*Options)) (*ListProposalsOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 RejectInvitation(ctx context.Context, params *RejectInvitationInput, optFns ...func(*Options)) (*RejectInvitationOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateMember(ctx context.Context, params *UpdateMemberInput, optFns ...func(*Options)) (*UpdateMemberOutput, error) 
 UpdateNode(ctx context.Context, params *UpdateNodeInput, optFns ...func(*Options)) (*UpdateNodeOutput, error) 
 VoteOnProposal(ctx context.Context, params *VoteOnProposalInput, optFns ...func(*Options)) (*VoteOnProposalOutput, error) 
}
