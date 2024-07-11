
package servicequotas

import (
    "github.com/aws/aws-sdk-go-v2/service/servicequotas"
)

// IClient defines the interface for servicequotas
type IClient interface {
 Options() Options 
 AssociateServiceQuotaTemplate(ctx context.Context, params *AssociateServiceQuotaTemplateInput, optFns ...func(*Options)) (*AssociateServiceQuotaTemplateOutput, error) 
 DeleteServiceQuotaIncreaseRequestFromTemplate(ctx context.Context, params *DeleteServiceQuotaIncreaseRequestFromTemplateInput, optFns ...func(*Options)) (*DeleteServiceQuotaIncreaseRequestFromTemplateOutput, error) 
 DisassociateServiceQuotaTemplate(ctx context.Context, params *DisassociateServiceQuotaTemplateInput, optFns ...func(*Options)) (*DisassociateServiceQuotaTemplateOutput, error) 
 GetAWSDefaultServiceQuota(ctx context.Context, params *GetAWSDefaultServiceQuotaInput, optFns ...func(*Options)) (*GetAWSDefaultServiceQuotaOutput, error) 
 GetAssociationForServiceQuotaTemplate(ctx context.Context, params *GetAssociationForServiceQuotaTemplateInput, optFns ...func(*Options)) (*GetAssociationForServiceQuotaTemplateOutput, error) 
 GetRequestedServiceQuotaChange(ctx context.Context, params *GetRequestedServiceQuotaChangeInput, optFns ...func(*Options)) (*GetRequestedServiceQuotaChangeOutput, error) 
 GetServiceQuota(ctx context.Context, params *GetServiceQuotaInput, optFns ...func(*Options)) (*GetServiceQuotaOutput, error) 
 GetServiceQuotaIncreaseRequestFromTemplate(ctx context.Context, params *GetServiceQuotaIncreaseRequestFromTemplateInput, optFns ...func(*Options)) (*GetServiceQuotaIncreaseRequestFromTemplateOutput, error) 
 ListAWSDefaultServiceQuotas(ctx context.Context, params *ListAWSDefaultServiceQuotasInput, optFns ...func(*Options)) (*ListAWSDefaultServiceQuotasOutput, error) 
 ListRequestedServiceQuotaChangeHistory(ctx context.Context, params *ListRequestedServiceQuotaChangeHistoryInput, optFns ...func(*Options)) (*ListRequestedServiceQuotaChangeHistoryOutput, error) 
 ListRequestedServiceQuotaChangeHistoryByQuota(ctx context.Context, params *ListRequestedServiceQuotaChangeHistoryByQuotaInput, optFns ...func(*Options)) (*ListRequestedServiceQuotaChangeHistoryByQuotaOutput, error) 
 ListServiceQuotaIncreaseRequestsInTemplate(ctx context.Context, params *ListServiceQuotaIncreaseRequestsInTemplateInput, optFns ...func(*Options)) (*ListServiceQuotaIncreaseRequestsInTemplateOutput, error) 
 ListServiceQuotas(ctx context.Context, params *ListServiceQuotasInput, optFns ...func(*Options)) (*ListServiceQuotasOutput, error) 
 ListServices(ctx context.Context, params *ListServicesInput, optFns ...func(*Options)) (*ListServicesOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 PutServiceQuotaIncreaseRequestIntoTemplate(ctx context.Context, params *PutServiceQuotaIncreaseRequestIntoTemplateInput, optFns ...func(*Options)) (*PutServiceQuotaIncreaseRequestIntoTemplateOutput, error) 
 RequestServiceQuotaIncrease(ctx context.Context, params *RequestServiceQuotaIncreaseInput, optFns ...func(*Options)) (*RequestServiceQuotaIncreaseOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
}
