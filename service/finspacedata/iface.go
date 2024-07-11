
package finspacedata

import (
    "github.com/aws/aws-sdk-go-v2/service/finspacedata"
)

// IFinspacedata defines the interface for finspacedata
type IFinspacedata interface {
 Options() Options 
 AssociateUserToPermissionGroup(ctx context.Context, params *AssociateUserToPermissionGroupInput, optFns ...func(*Options)) (*AssociateUserToPermissionGroupOutput, error) 
 CreateChangeset(ctx context.Context, params *CreateChangesetInput, optFns ...func(*Options)) (*CreateChangesetOutput, error) 
 CreateDataView(ctx context.Context, params *CreateDataViewInput, optFns ...func(*Options)) (*CreateDataViewOutput, error) 
 CreateDataset(ctx context.Context, params *CreateDatasetInput, optFns ...func(*Options)) (*CreateDatasetOutput, error) 
 CreatePermissionGroup(ctx context.Context, params *CreatePermissionGroupInput, optFns ...func(*Options)) (*CreatePermissionGroupOutput, error) 
 CreateUser(ctx context.Context, params *CreateUserInput, optFns ...func(*Options)) (*CreateUserOutput, error) 
 DeleteDataset(ctx context.Context, params *DeleteDatasetInput, optFns ...func(*Options)) (*DeleteDatasetOutput, error) 
 DeletePermissionGroup(ctx context.Context, params *DeletePermissionGroupInput, optFns ...func(*Options)) (*DeletePermissionGroupOutput, error) 
 DisableUser(ctx context.Context, params *DisableUserInput, optFns ...func(*Options)) (*DisableUserOutput, error) 
 DisassociateUserFromPermissionGroup(ctx context.Context, params *DisassociateUserFromPermissionGroupInput, optFns ...func(*Options)) (*DisassociateUserFromPermissionGroupOutput, error) 
 EnableUser(ctx context.Context, params *EnableUserInput, optFns ...func(*Options)) (*EnableUserOutput, error) 
 GetChangeset(ctx context.Context, params *GetChangesetInput, optFns ...func(*Options)) (*GetChangesetOutput, error) 
 GetDataView(ctx context.Context, params *GetDataViewInput, optFns ...func(*Options)) (*GetDataViewOutput, error) 
 GetDataset(ctx context.Context, params *GetDatasetInput, optFns ...func(*Options)) (*GetDatasetOutput, error) 
 GetExternalDataViewAccessDetails(ctx context.Context, params *GetExternalDataViewAccessDetailsInput, optFns ...func(*Options)) (*GetExternalDataViewAccessDetailsOutput, error) 
 GetPermissionGroup(ctx context.Context, params *GetPermissionGroupInput, optFns ...func(*Options)) (*GetPermissionGroupOutput, error) 
 GetProgrammaticAccessCredentials(ctx context.Context, params *GetProgrammaticAccessCredentialsInput, optFns ...func(*Options)) (*GetProgrammaticAccessCredentialsOutput, error) 
 GetUser(ctx context.Context, params *GetUserInput, optFns ...func(*Options)) (*GetUserOutput, error) 
 GetWorkingLocation(ctx context.Context, params *GetWorkingLocationInput, optFns ...func(*Options)) (*GetWorkingLocationOutput, error) 
 ListChangesets(ctx context.Context, params *ListChangesetsInput, optFns ...func(*Options)) (*ListChangesetsOutput, error) 
 ListDataViews(ctx context.Context, params *ListDataViewsInput, optFns ...func(*Options)) (*ListDataViewsOutput, error) 
 ListDatasets(ctx context.Context, params *ListDatasetsInput, optFns ...func(*Options)) (*ListDatasetsOutput, error) 
 ListPermissionGroups(ctx context.Context, params *ListPermissionGroupsInput, optFns ...func(*Options)) (*ListPermissionGroupsOutput, error) 
 ListPermissionGroupsByUser(ctx context.Context, params *ListPermissionGroupsByUserInput, optFns ...func(*Options)) (*ListPermissionGroupsByUserOutput, error) 
 ListUsers(ctx context.Context, params *ListUsersInput, optFns ...func(*Options)) (*ListUsersOutput, error) 
 ListUsersByPermissionGroup(ctx context.Context, params *ListUsersByPermissionGroupInput, optFns ...func(*Options)) (*ListUsersByPermissionGroupOutput, error) 
 ResetUserPassword(ctx context.Context, params *ResetUserPasswordInput, optFns ...func(*Options)) (*ResetUserPasswordOutput, error) 
 UpdateChangeset(ctx context.Context, params *UpdateChangesetInput, optFns ...func(*Options)) (*UpdateChangesetOutput, error) 
 UpdateDataset(ctx context.Context, params *UpdateDatasetInput, optFns ...func(*Options)) (*UpdateDatasetOutput, error) 
 UpdatePermissionGroup(ctx context.Context, params *UpdatePermissionGroupInput, optFns ...func(*Options)) (*UpdatePermissionGroupOutput, error) 
 UpdateUser(ctx context.Context, params *UpdateUserInput, optFns ...func(*Options)) (*UpdateUserOutput, error) 
}
