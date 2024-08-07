
package identitystore_iface

// This file was generated by scripts/gen_ifaces.sh do not edit it manually
import (
    "context"
    . "github.com/aws/aws-sdk-go-v2/service/identitystore"
)

// IClient defines the interface for identitystore
type IClient interface {
 Options() Options 
 CreateGroup(ctx context.Context, params *CreateGroupInput, optFns ...func(*Options)) (*CreateGroupOutput, error) 
 CreateGroupMembership(ctx context.Context, params *CreateGroupMembershipInput, optFns ...func(*Options)) (*CreateGroupMembershipOutput, error) 
 CreateUser(ctx context.Context, params *CreateUserInput, optFns ...func(*Options)) (*CreateUserOutput, error) 
 DeleteGroup(ctx context.Context, params *DeleteGroupInput, optFns ...func(*Options)) (*DeleteGroupOutput, error) 
 DeleteGroupMembership(ctx context.Context, params *DeleteGroupMembershipInput, optFns ...func(*Options)) (*DeleteGroupMembershipOutput, error) 
 DeleteUser(ctx context.Context, params *DeleteUserInput, optFns ...func(*Options)) (*DeleteUserOutput, error) 
 DescribeGroup(ctx context.Context, params *DescribeGroupInput, optFns ...func(*Options)) (*DescribeGroupOutput, error) 
 DescribeGroupMembership(ctx context.Context, params *DescribeGroupMembershipInput, optFns ...func(*Options)) (*DescribeGroupMembershipOutput, error) 
 DescribeUser(ctx context.Context, params *DescribeUserInput, optFns ...func(*Options)) (*DescribeUserOutput, error) 
 GetGroupId(ctx context.Context, params *GetGroupIdInput, optFns ...func(*Options)) (*GetGroupIdOutput, error) 
 GetGroupMembershipId(ctx context.Context, params *GetGroupMembershipIdInput, optFns ...func(*Options)) (*GetGroupMembershipIdOutput, error) 
 GetUserId(ctx context.Context, params *GetUserIdInput, optFns ...func(*Options)) (*GetUserIdOutput, error) 
 IsMemberInGroups(ctx context.Context, params *IsMemberInGroupsInput, optFns ...func(*Options)) (*IsMemberInGroupsOutput, error) 
 ListGroupMemberships(ctx context.Context, params *ListGroupMembershipsInput, optFns ...func(*Options)) (*ListGroupMembershipsOutput, error) 
 ListGroupMembershipsForMember(ctx context.Context, params *ListGroupMembershipsForMemberInput, optFns ...func(*Options)) (*ListGroupMembershipsForMemberOutput, error) 
 ListGroups(ctx context.Context, params *ListGroupsInput, optFns ...func(*Options)) (*ListGroupsOutput, error) 
 ListUsers(ctx context.Context, params *ListUsersInput, optFns ...func(*Options)) (*ListUsersOutput, error) 
 UpdateGroup(ctx context.Context, params *UpdateGroupInput, optFns ...func(*Options)) (*UpdateGroupOutput, error) 
 UpdateUser(ctx context.Context, params *UpdateUserInput, optFns ...func(*Options)) (*UpdateUserOutput, error) 
}
