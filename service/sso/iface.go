
package sso

import (
    "github.com/aws/aws-sdk-go-v2/service/sso"
)

// ISso defines the interface for sso
type ISso interface {
 Options() Options 
 GetRoleCredentials(ctx context.Context, params *GetRoleCredentialsInput, optFns ...func(*Options)) (*GetRoleCredentialsOutput, error) 
 ListAccountRoles(ctx context.Context, params *ListAccountRolesInput, optFns ...func(*Options)) (*ListAccountRolesOutput, error) 
 ListAccounts(ctx context.Context, params *ListAccountsInput, optFns ...func(*Options)) (*ListAccountsOutput, error) 
 Logout(ctx context.Context, params *LogoutInput, optFns ...func(*Options)) (*LogoutOutput, error) 
}
