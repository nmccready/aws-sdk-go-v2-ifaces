
package sso_iface

// This file was generated by scripts/gen_ifaces.sh do not edit it manually
import (
    "context"
    . "github.com/aws/aws-sdk-go-v2/service/sso"
)

// IClient defines the interface for sso
type IClient interface {
 Options() Options 
 GetRoleCredentials(ctx context.Context, params *GetRoleCredentialsInput, optFns ...func(*Options)) (*GetRoleCredentialsOutput, error) 
 ListAccountRoles(ctx context.Context, params *ListAccountRolesInput, optFns ...func(*Options)) (*ListAccountRolesOutput, error) 
 ListAccounts(ctx context.Context, params *ListAccountsInput, optFns ...func(*Options)) (*ListAccountsOutput, error) 
 Logout(ctx context.Context, params *LogoutInput, optFns ...func(*Options)) (*LogoutOutput, error) 
}
