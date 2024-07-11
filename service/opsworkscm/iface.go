
package opsworkscm

import (
    "github.com/aws/aws-sdk-go-v2/service/opsworkscm"
)

// IOpsworkscm defines the interface for opsworkscm
type IOpsworkscm interface {
 Options() Options 
 AssociateNode(ctx context.Context, params *AssociateNodeInput, optFns ...func(*Options)) (*AssociateNodeOutput, error) 
 CreateBackup(ctx context.Context, params *CreateBackupInput, optFns ...func(*Options)) (*CreateBackupOutput, error) 
 CreateServer(ctx context.Context, params *CreateServerInput, optFns ...func(*Options)) (*CreateServerOutput, error) 
 DeleteBackup(ctx context.Context, params *DeleteBackupInput, optFns ...func(*Options)) (*DeleteBackupOutput, error) 
 DeleteServer(ctx context.Context, params *DeleteServerInput, optFns ...func(*Options)) (*DeleteServerOutput, error) 
 DescribeAccountAttributes(ctx context.Context, params *DescribeAccountAttributesInput, optFns ...func(*Options)) (*DescribeAccountAttributesOutput, error) 
 DescribeBackups(ctx context.Context, params *DescribeBackupsInput, optFns ...func(*Options)) (*DescribeBackupsOutput, error) 
 DescribeEvents(ctx context.Context, params *DescribeEventsInput, optFns ...func(*Options)) (*DescribeEventsOutput, error) 
 DescribeNodeAssociationStatus(ctx context.Context, params *DescribeNodeAssociationStatusInput, optFns ...func(*Options)) (*DescribeNodeAssociationStatusOutput, error) 
 DescribeServers(ctx context.Context, params *DescribeServersInput, optFns ...func(*Options)) (*DescribeServersOutput, error) 
 DisassociateNode(ctx context.Context, params *DisassociateNodeInput, optFns ...func(*Options)) (*DisassociateNodeOutput, error) 
 ExportServerEngineAttribute(ctx context.Context, params *ExportServerEngineAttributeInput, optFns ...func(*Options)) (*ExportServerEngineAttributeOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 RestoreServer(ctx context.Context, params *RestoreServerInput, optFns ...func(*Options)) (*RestoreServerOutput, error) 
 StartMaintenance(ctx context.Context, params *StartMaintenanceInput, optFns ...func(*Options)) (*StartMaintenanceOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateServer(ctx context.Context, params *UpdateServerInput, optFns ...func(*Options)) (*UpdateServerOutput, error) 
 UpdateServerEngineAttributes(ctx context.Context, params *UpdateServerEngineAttributesInput, optFns ...func(*Options)) (*UpdateServerEngineAttributesOutput, error) 
}
