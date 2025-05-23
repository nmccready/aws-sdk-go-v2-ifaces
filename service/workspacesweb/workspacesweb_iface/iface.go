
package workspacesweb_iface

// This file was generated by scripts/gen_ifaces.sh do not edit it manually
import (
    "context"
    . "github.com/aws/aws-sdk-go-v2/service/workspacesweb"
)

// IClient defines the interface for workspacesweb
type IClient interface {
 Options() Options 
 AssociateBrowserSettings(ctx context.Context, params *AssociateBrowserSettingsInput, optFns ...func(*Options)) (*AssociateBrowserSettingsOutput, error) 
 AssociateDataProtectionSettings(ctx context.Context, params *AssociateDataProtectionSettingsInput, optFns ...func(*Options)) (*AssociateDataProtectionSettingsOutput, error) 
 AssociateIpAccessSettings(ctx context.Context, params *AssociateIpAccessSettingsInput, optFns ...func(*Options)) (*AssociateIpAccessSettingsOutput, error) 
 AssociateNetworkSettings(ctx context.Context, params *AssociateNetworkSettingsInput, optFns ...func(*Options)) (*AssociateNetworkSettingsOutput, error) 
 AssociateTrustStore(ctx context.Context, params *AssociateTrustStoreInput, optFns ...func(*Options)) (*AssociateTrustStoreOutput, error) 
 AssociateUserAccessLoggingSettings(ctx context.Context, params *AssociateUserAccessLoggingSettingsInput, optFns ...func(*Options)) (*AssociateUserAccessLoggingSettingsOutput, error) 
 AssociateUserSettings(ctx context.Context, params *AssociateUserSettingsInput, optFns ...func(*Options)) (*AssociateUserSettingsOutput, error) 
 CreateBrowserSettings(ctx context.Context, params *CreateBrowserSettingsInput, optFns ...func(*Options)) (*CreateBrowserSettingsOutput, error) 
 CreateDataProtectionSettings(ctx context.Context, params *CreateDataProtectionSettingsInput, optFns ...func(*Options)) (*CreateDataProtectionSettingsOutput, error) 
 CreateIdentityProvider(ctx context.Context, params *CreateIdentityProviderInput, optFns ...func(*Options)) (*CreateIdentityProviderOutput, error) 
 CreateIpAccessSettings(ctx context.Context, params *CreateIpAccessSettingsInput, optFns ...func(*Options)) (*CreateIpAccessSettingsOutput, error) 
 CreateNetworkSettings(ctx context.Context, params *CreateNetworkSettingsInput, optFns ...func(*Options)) (*CreateNetworkSettingsOutput, error) 
 CreatePortal(ctx context.Context, params *CreatePortalInput, optFns ...func(*Options)) (*CreatePortalOutput, error) 
 CreateTrustStore(ctx context.Context, params *CreateTrustStoreInput, optFns ...func(*Options)) (*CreateTrustStoreOutput, error) 
 CreateUserAccessLoggingSettings(ctx context.Context, params *CreateUserAccessLoggingSettingsInput, optFns ...func(*Options)) (*CreateUserAccessLoggingSettingsOutput, error) 
 CreateUserSettings(ctx context.Context, params *CreateUserSettingsInput, optFns ...func(*Options)) (*CreateUserSettingsOutput, error) 
 DeleteBrowserSettings(ctx context.Context, params *DeleteBrowserSettingsInput, optFns ...func(*Options)) (*DeleteBrowserSettingsOutput, error) 
 DeleteDataProtectionSettings(ctx context.Context, params *DeleteDataProtectionSettingsInput, optFns ...func(*Options)) (*DeleteDataProtectionSettingsOutput, error) 
 DeleteIdentityProvider(ctx context.Context, params *DeleteIdentityProviderInput, optFns ...func(*Options)) (*DeleteIdentityProviderOutput, error) 
 DeleteIpAccessSettings(ctx context.Context, params *DeleteIpAccessSettingsInput, optFns ...func(*Options)) (*DeleteIpAccessSettingsOutput, error) 
 DeleteNetworkSettings(ctx context.Context, params *DeleteNetworkSettingsInput, optFns ...func(*Options)) (*DeleteNetworkSettingsOutput, error) 
 DeletePortal(ctx context.Context, params *DeletePortalInput, optFns ...func(*Options)) (*DeletePortalOutput, error) 
 DeleteTrustStore(ctx context.Context, params *DeleteTrustStoreInput, optFns ...func(*Options)) (*DeleteTrustStoreOutput, error) 
 DeleteUserAccessLoggingSettings(ctx context.Context, params *DeleteUserAccessLoggingSettingsInput, optFns ...func(*Options)) (*DeleteUserAccessLoggingSettingsOutput, error) 
 DeleteUserSettings(ctx context.Context, params *DeleteUserSettingsInput, optFns ...func(*Options)) (*DeleteUserSettingsOutput, error) 
 DisassociateBrowserSettings(ctx context.Context, params *DisassociateBrowserSettingsInput, optFns ...func(*Options)) (*DisassociateBrowserSettingsOutput, error) 
 DisassociateDataProtectionSettings(ctx context.Context, params *DisassociateDataProtectionSettingsInput, optFns ...func(*Options)) (*DisassociateDataProtectionSettingsOutput, error) 
 DisassociateIpAccessSettings(ctx context.Context, params *DisassociateIpAccessSettingsInput, optFns ...func(*Options)) (*DisassociateIpAccessSettingsOutput, error) 
 DisassociateNetworkSettings(ctx context.Context, params *DisassociateNetworkSettingsInput, optFns ...func(*Options)) (*DisassociateNetworkSettingsOutput, error) 
 DisassociateTrustStore(ctx context.Context, params *DisassociateTrustStoreInput, optFns ...func(*Options)) (*DisassociateTrustStoreOutput, error) 
 DisassociateUserAccessLoggingSettings(ctx context.Context, params *DisassociateUserAccessLoggingSettingsInput, optFns ...func(*Options)) (*DisassociateUserAccessLoggingSettingsOutput, error) 
 DisassociateUserSettings(ctx context.Context, params *DisassociateUserSettingsInput, optFns ...func(*Options)) (*DisassociateUserSettingsOutput, error) 
 ExpireSession(ctx context.Context, params *ExpireSessionInput, optFns ...func(*Options)) (*ExpireSessionOutput, error) 
 GetBrowserSettings(ctx context.Context, params *GetBrowserSettingsInput, optFns ...func(*Options)) (*GetBrowserSettingsOutput, error) 
 GetDataProtectionSettings(ctx context.Context, params *GetDataProtectionSettingsInput, optFns ...func(*Options)) (*GetDataProtectionSettingsOutput, error) 
 GetIdentityProvider(ctx context.Context, params *GetIdentityProviderInput, optFns ...func(*Options)) (*GetIdentityProviderOutput, error) 
 GetIpAccessSettings(ctx context.Context, params *GetIpAccessSettingsInput, optFns ...func(*Options)) (*GetIpAccessSettingsOutput, error) 
 GetNetworkSettings(ctx context.Context, params *GetNetworkSettingsInput, optFns ...func(*Options)) (*GetNetworkSettingsOutput, error) 
 GetPortal(ctx context.Context, params *GetPortalInput, optFns ...func(*Options)) (*GetPortalOutput, error) 
 GetPortalServiceProviderMetadata(ctx context.Context, params *GetPortalServiceProviderMetadataInput, optFns ...func(*Options)) (*GetPortalServiceProviderMetadataOutput, error) 
 GetSession(ctx context.Context, params *GetSessionInput, optFns ...func(*Options)) (*GetSessionOutput, error) 
 GetTrustStore(ctx context.Context, params *GetTrustStoreInput, optFns ...func(*Options)) (*GetTrustStoreOutput, error) 
 GetTrustStoreCertificate(ctx context.Context, params *GetTrustStoreCertificateInput, optFns ...func(*Options)) (*GetTrustStoreCertificateOutput, error) 
 GetUserAccessLoggingSettings(ctx context.Context, params *GetUserAccessLoggingSettingsInput, optFns ...func(*Options)) (*GetUserAccessLoggingSettingsOutput, error) 
 GetUserSettings(ctx context.Context, params *GetUserSettingsInput, optFns ...func(*Options)) (*GetUserSettingsOutput, error) 
 ListBrowserSettings(ctx context.Context, params *ListBrowserSettingsInput, optFns ...func(*Options)) (*ListBrowserSettingsOutput, error) 
 ListDataProtectionSettings(ctx context.Context, params *ListDataProtectionSettingsInput, optFns ...func(*Options)) (*ListDataProtectionSettingsOutput, error) 
 ListIdentityProviders(ctx context.Context, params *ListIdentityProvidersInput, optFns ...func(*Options)) (*ListIdentityProvidersOutput, error) 
 ListIpAccessSettings(ctx context.Context, params *ListIpAccessSettingsInput, optFns ...func(*Options)) (*ListIpAccessSettingsOutput, error) 
 ListNetworkSettings(ctx context.Context, params *ListNetworkSettingsInput, optFns ...func(*Options)) (*ListNetworkSettingsOutput, error) 
 ListPortals(ctx context.Context, params *ListPortalsInput, optFns ...func(*Options)) (*ListPortalsOutput, error) 
 ListSessions(ctx context.Context, params *ListSessionsInput, optFns ...func(*Options)) (*ListSessionsOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 ListTrustStoreCertificates(ctx context.Context, params *ListTrustStoreCertificatesInput, optFns ...func(*Options)) (*ListTrustStoreCertificatesOutput, error) 
 ListTrustStores(ctx context.Context, params *ListTrustStoresInput, optFns ...func(*Options)) (*ListTrustStoresOutput, error) 
 ListUserAccessLoggingSettings(ctx context.Context, params *ListUserAccessLoggingSettingsInput, optFns ...func(*Options)) (*ListUserAccessLoggingSettingsOutput, error) 
 ListUserSettings(ctx context.Context, params *ListUserSettingsInput, optFns ...func(*Options)) (*ListUserSettingsOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateBrowserSettings(ctx context.Context, params *UpdateBrowserSettingsInput, optFns ...func(*Options)) (*UpdateBrowserSettingsOutput, error) 
 UpdateDataProtectionSettings(ctx context.Context, params *UpdateDataProtectionSettingsInput, optFns ...func(*Options)) (*UpdateDataProtectionSettingsOutput, error) 
 UpdateIdentityProvider(ctx context.Context, params *UpdateIdentityProviderInput, optFns ...func(*Options)) (*UpdateIdentityProviderOutput, error) 
 UpdateIpAccessSettings(ctx context.Context, params *UpdateIpAccessSettingsInput, optFns ...func(*Options)) (*UpdateIpAccessSettingsOutput, error) 
 UpdateNetworkSettings(ctx context.Context, params *UpdateNetworkSettingsInput, optFns ...func(*Options)) (*UpdateNetworkSettingsOutput, error) 
 UpdatePortal(ctx context.Context, params *UpdatePortalInput, optFns ...func(*Options)) (*UpdatePortalOutput, error) 
 UpdateTrustStore(ctx context.Context, params *UpdateTrustStoreInput, optFns ...func(*Options)) (*UpdateTrustStoreOutput, error) 
 UpdateUserAccessLoggingSettings(ctx context.Context, params *UpdateUserAccessLoggingSettingsInput, optFns ...func(*Options)) (*UpdateUserAccessLoggingSettingsOutput, error) 
 UpdateUserSettings(ctx context.Context, params *UpdateUserSettingsInput, optFns ...func(*Options)) (*UpdateUserSettingsOutput, error) 
}
