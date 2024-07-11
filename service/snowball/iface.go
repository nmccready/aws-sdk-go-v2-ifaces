
package snowball

import (
    "github.com/aws/aws-sdk-go-v2/service/snowball"
)

// IClient defines the interface for snowball
type IClient interface {
 Options() Options 
 CancelCluster(ctx context.Context, params *CancelClusterInput, optFns ...func(*Options)) (*CancelClusterOutput, error) 
 CancelJob(ctx context.Context, params *CancelJobInput, optFns ...func(*Options)) (*CancelJobOutput, error) 
 CreateAddress(ctx context.Context, params *CreateAddressInput, optFns ...func(*Options)) (*CreateAddressOutput, error) 
 CreateCluster(ctx context.Context, params *CreateClusterInput, optFns ...func(*Options)) (*CreateClusterOutput, error) 
 CreateJob(ctx context.Context, params *CreateJobInput, optFns ...func(*Options)) (*CreateJobOutput, error) 
 CreateLongTermPricing(ctx context.Context, params *CreateLongTermPricingInput, optFns ...func(*Options)) (*CreateLongTermPricingOutput, error) 
 CreateReturnShippingLabel(ctx context.Context, params *CreateReturnShippingLabelInput, optFns ...func(*Options)) (*CreateReturnShippingLabelOutput, error) 
 DescribeAddress(ctx context.Context, params *DescribeAddressInput, optFns ...func(*Options)) (*DescribeAddressOutput, error) 
 DescribeAddresses(ctx context.Context, params *DescribeAddressesInput, optFns ...func(*Options)) (*DescribeAddressesOutput, error) 
 DescribeCluster(ctx context.Context, params *DescribeClusterInput, optFns ...func(*Options)) (*DescribeClusterOutput, error) 
 DescribeJob(ctx context.Context, params *DescribeJobInput, optFns ...func(*Options)) (*DescribeJobOutput, error) 
 DescribeReturnShippingLabel(ctx context.Context, params *DescribeReturnShippingLabelInput, optFns ...func(*Options)) (*DescribeReturnShippingLabelOutput, error) 
 GetJobManifest(ctx context.Context, params *GetJobManifestInput, optFns ...func(*Options)) (*GetJobManifestOutput, error) 
 GetJobUnlockCode(ctx context.Context, params *GetJobUnlockCodeInput, optFns ...func(*Options)) (*GetJobUnlockCodeOutput, error) 
 GetSnowballUsage(ctx context.Context, params *GetSnowballUsageInput, optFns ...func(*Options)) (*GetSnowballUsageOutput, error) 
 GetSoftwareUpdates(ctx context.Context, params *GetSoftwareUpdatesInput, optFns ...func(*Options)) (*GetSoftwareUpdatesOutput, error) 
 ListClusterJobs(ctx context.Context, params *ListClusterJobsInput, optFns ...func(*Options)) (*ListClusterJobsOutput, error) 
 ListClusters(ctx context.Context, params *ListClustersInput, optFns ...func(*Options)) (*ListClustersOutput, error) 
 ListCompatibleImages(ctx context.Context, params *ListCompatibleImagesInput, optFns ...func(*Options)) (*ListCompatibleImagesOutput, error) 
 ListJobs(ctx context.Context, params *ListJobsInput, optFns ...func(*Options)) (*ListJobsOutput, error) 
 ListLongTermPricing(ctx context.Context, params *ListLongTermPricingInput, optFns ...func(*Options)) (*ListLongTermPricingOutput, error) 
 ListPickupLocations(ctx context.Context, params *ListPickupLocationsInput, optFns ...func(*Options)) (*ListPickupLocationsOutput, error) 
 ListServiceVersions(ctx context.Context, params *ListServiceVersionsInput, optFns ...func(*Options)) (*ListServiceVersionsOutput, error) 
 UpdateCluster(ctx context.Context, params *UpdateClusterInput, optFns ...func(*Options)) (*UpdateClusterOutput, error) 
 UpdateJob(ctx context.Context, params *UpdateJobInput, optFns ...func(*Options)) (*UpdateJobOutput, error) 
 UpdateJobShipmentState(ctx context.Context, params *UpdateJobShipmentStateInput, optFns ...func(*Options)) (*UpdateJobShipmentStateOutput, error) 
 UpdateLongTermPricing(ctx context.Context, params *UpdateLongTermPricingInput, optFns ...func(*Options)) (*UpdateLongTermPricingOutput, error) 
}
