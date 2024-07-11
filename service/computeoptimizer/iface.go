
package computeoptimizer

import (
    "github.com/aws/aws-sdk-go-v2/service/computeoptimizer"
)

// IComputeoptimizer defines the interface for computeoptimizer
type IComputeoptimizer interface {
 Options() Options 
 DeleteRecommendationPreferences(ctx context.Context, params *DeleteRecommendationPreferencesInput, optFns ...func(*Options)) (*DeleteRecommendationPreferencesOutput, error) 
 DescribeRecommendationExportJobs(ctx context.Context, params *DescribeRecommendationExportJobsInput, optFns ...func(*Options)) (*DescribeRecommendationExportJobsOutput, error) 
 ExportAutoScalingGroupRecommendations(ctx context.Context, params *ExportAutoScalingGroupRecommendationsInput, optFns ...func(*Options)) (*ExportAutoScalingGroupRecommendationsOutput, error) 
 ExportEBSVolumeRecommendations(ctx context.Context, params *ExportEBSVolumeRecommendationsInput, optFns ...func(*Options)) (*ExportEBSVolumeRecommendationsOutput, error) 
 ExportEC2InstanceRecommendations(ctx context.Context, params *ExportEC2InstanceRecommendationsInput, optFns ...func(*Options)) (*ExportEC2InstanceRecommendationsOutput, error) 
 ExportECSServiceRecommendations(ctx context.Context, params *ExportECSServiceRecommendationsInput, optFns ...func(*Options)) (*ExportECSServiceRecommendationsOutput, error) 
 ExportLambdaFunctionRecommendations(ctx context.Context, params *ExportLambdaFunctionRecommendationsInput, optFns ...func(*Options)) (*ExportLambdaFunctionRecommendationsOutput, error) 
 ExportLicenseRecommendations(ctx context.Context, params *ExportLicenseRecommendationsInput, optFns ...func(*Options)) (*ExportLicenseRecommendationsOutput, error) 
 ExportRDSDatabaseRecommendations(ctx context.Context, params *ExportRDSDatabaseRecommendationsInput, optFns ...func(*Options)) (*ExportRDSDatabaseRecommendationsOutput, error) 
 GetAutoScalingGroupRecommendations(ctx context.Context, params *GetAutoScalingGroupRecommendationsInput, optFns ...func(*Options)) (*GetAutoScalingGroupRecommendationsOutput, error) 
 GetEBSVolumeRecommendations(ctx context.Context, params *GetEBSVolumeRecommendationsInput, optFns ...func(*Options)) (*GetEBSVolumeRecommendationsOutput, error) 
 GetEC2InstanceRecommendations(ctx context.Context, params *GetEC2InstanceRecommendationsInput, optFns ...func(*Options)) (*GetEC2InstanceRecommendationsOutput, error) 
 GetEC2RecommendationProjectedMetrics(ctx context.Context, params *GetEC2RecommendationProjectedMetricsInput, optFns ...func(*Options)) (*GetEC2RecommendationProjectedMetricsOutput, error) 
 GetECSServiceRecommendationProjectedMetrics(ctx context.Context, params *GetECSServiceRecommendationProjectedMetricsInput, optFns ...func(*Options)) (*GetECSServiceRecommendationProjectedMetricsOutput, error) 
 GetECSServiceRecommendations(ctx context.Context, params *GetECSServiceRecommendationsInput, optFns ...func(*Options)) (*GetECSServiceRecommendationsOutput, error) 
 GetEffectiveRecommendationPreferences(ctx context.Context, params *GetEffectiveRecommendationPreferencesInput, optFns ...func(*Options)) (*GetEffectiveRecommendationPreferencesOutput, error) 
 GetEnrollmentStatus(ctx context.Context, params *GetEnrollmentStatusInput, optFns ...func(*Options)) (*GetEnrollmentStatusOutput, error) 
 GetEnrollmentStatusesForOrganization(ctx context.Context, params *GetEnrollmentStatusesForOrganizationInput, optFns ...func(*Options)) (*GetEnrollmentStatusesForOrganizationOutput, error) 
 GetLambdaFunctionRecommendations(ctx context.Context, params *GetLambdaFunctionRecommendationsInput, optFns ...func(*Options)) (*GetLambdaFunctionRecommendationsOutput, error) 
 GetLicenseRecommendations(ctx context.Context, params *GetLicenseRecommendationsInput, optFns ...func(*Options)) (*GetLicenseRecommendationsOutput, error) 
 GetRDSDatabaseRecommendationProjectedMetrics(ctx context.Context, params *GetRDSDatabaseRecommendationProjectedMetricsInput, optFns ...func(*Options)) (*GetRDSDatabaseRecommendationProjectedMetricsOutput, error) 
 GetRDSDatabaseRecommendations(ctx context.Context, params *GetRDSDatabaseRecommendationsInput, optFns ...func(*Options)) (*GetRDSDatabaseRecommendationsOutput, error) 
 GetRecommendationPreferences(ctx context.Context, params *GetRecommendationPreferencesInput, optFns ...func(*Options)) (*GetRecommendationPreferencesOutput, error) 
 GetRecommendationSummaries(ctx context.Context, params *GetRecommendationSummariesInput, optFns ...func(*Options)) (*GetRecommendationSummariesOutput, error) 
 PutRecommendationPreferences(ctx context.Context, params *PutRecommendationPreferencesInput, optFns ...func(*Options)) (*PutRecommendationPreferencesOutput, error) 
 UpdateEnrollmentStatus(ctx context.Context, params *UpdateEnrollmentStatusInput, optFns ...func(*Options)) (*UpdateEnrollmentStatusOutput, error) 
}
