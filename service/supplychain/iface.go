
package supplychain

import (
    "github.com/aws/aws-sdk-go-v2/service/supplychain"
)

// IClient defines the interface for supplychain
type IClient interface {
 Options() Options 
 CreateBillOfMaterialsImportJob(ctx context.Context, params *CreateBillOfMaterialsImportJobInput, optFns ...func(*Options)) (*CreateBillOfMaterialsImportJobOutput, error) 
 GetBillOfMaterialsImportJob(ctx context.Context, params *GetBillOfMaterialsImportJobInput, optFns ...func(*Options)) (*GetBillOfMaterialsImportJobOutput, error) 
 SendDataIntegrationEvent(ctx context.Context, params *SendDataIntegrationEventInput, optFns ...func(*Options)) (*SendDataIntegrationEventOutput, error) 
}
