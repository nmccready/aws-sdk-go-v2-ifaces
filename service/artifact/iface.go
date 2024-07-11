
package artifact

// This file was generated by scripts/gen_ifaces.sh do not edit it manually
import (
    "github.com/aws/aws-sdk-go-v2/service/artifact"
)

// IClient defines the interface for artifact
type IClient interface {
 Options() Options 
 GetAccountSettings(ctx context.Context, params *GetAccountSettingsInput, optFns ...func(*Options)) (*GetAccountSettingsOutput, error) 
 GetReport(ctx context.Context, params *GetReportInput, optFns ...func(*Options)) (*GetReportOutput, error) 
 GetReportMetadata(ctx context.Context, params *GetReportMetadataInput, optFns ...func(*Options)) (*GetReportMetadataOutput, error) 
 GetTermForReport(ctx context.Context, params *GetTermForReportInput, optFns ...func(*Options)) (*GetTermForReportOutput, error) 
 ListReports(ctx context.Context, params *ListReportsInput, optFns ...func(*Options)) (*ListReportsOutput, error) 
 PutAccountSettings(ctx context.Context, params *PutAccountSettingsInput, optFns ...func(*Options)) (*PutAccountSettingsOutput, error) 
}
