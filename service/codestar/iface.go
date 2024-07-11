
package codestar

import (
    "github.com/aws/aws-sdk-go-v2/service/codestar"
)

// ICodestar defines the interface for codestar
type ICodestar interface {
 Options() Options 
 AssociateTeamMember(ctx context.Context, params *AssociateTeamMemberInput, optFns ...func(*Options)) (*AssociateTeamMemberOutput, error) 
 CreateProject(ctx context.Context, params *CreateProjectInput, optFns ...func(*Options)) (*CreateProjectOutput, error) 
 CreateUserProfile(ctx context.Context, params *CreateUserProfileInput, optFns ...func(*Options)) (*CreateUserProfileOutput, error) 
 DeleteProject(ctx context.Context, params *DeleteProjectInput, optFns ...func(*Options)) (*DeleteProjectOutput, error) 
 DeleteUserProfile(ctx context.Context, params *DeleteUserProfileInput, optFns ...func(*Options)) (*DeleteUserProfileOutput, error) 
 DescribeProject(ctx context.Context, params *DescribeProjectInput, optFns ...func(*Options)) (*DescribeProjectOutput, error) 
 DescribeUserProfile(ctx context.Context, params *DescribeUserProfileInput, optFns ...func(*Options)) (*DescribeUserProfileOutput, error) 
 DisassociateTeamMember(ctx context.Context, params *DisassociateTeamMemberInput, optFns ...func(*Options)) (*DisassociateTeamMemberOutput, error) 
 ListProjects(ctx context.Context, params *ListProjectsInput, optFns ...func(*Options)) (*ListProjectsOutput, error) 
 ListResources(ctx context.Context, params *ListResourcesInput, optFns ...func(*Options)) (*ListResourcesOutput, error) 
 ListTagsForProject(ctx context.Context, params *ListTagsForProjectInput, optFns ...func(*Options)) (*ListTagsForProjectOutput, error) 
 ListTeamMembers(ctx context.Context, params *ListTeamMembersInput, optFns ...func(*Options)) (*ListTeamMembersOutput, error) 
 ListUserProfiles(ctx context.Context, params *ListUserProfilesInput, optFns ...func(*Options)) (*ListUserProfilesOutput, error) 
 TagProject(ctx context.Context, params *TagProjectInput, optFns ...func(*Options)) (*TagProjectOutput, error) 
 UntagProject(ctx context.Context, params *UntagProjectInput, optFns ...func(*Options)) (*UntagProjectOutput, error) 
 UpdateProject(ctx context.Context, params *UpdateProjectInput, optFns ...func(*Options)) (*UpdateProjectOutput, error) 
 UpdateTeamMember(ctx context.Context, params *UpdateTeamMemberInput, optFns ...func(*Options)) (*UpdateTeamMemberOutput, error) 
 UpdateUserProfile(ctx context.Context, params *UpdateUserProfileInput, optFns ...func(*Options)) (*UpdateUserProfileOutput, error) 
}
