
package comprehendmedical

import (
    "github.com/aws/aws-sdk-go-v2/service/comprehendmedical"
)

// IComprehendmedical defines the interface for comprehendmedical
type IComprehendmedical interface {
 Options() Options 
 DescribeEntitiesDetectionV2Job(ctx context.Context, params *DescribeEntitiesDetectionV2JobInput, optFns ...func(*Options)) (*DescribeEntitiesDetectionV2JobOutput, error) 
 DescribeICD10CMInferenceJob(ctx context.Context, params *DescribeICD10CMInferenceJobInput, optFns ...func(*Options)) (*DescribeICD10CMInferenceJobOutput, error) 
 DescribePHIDetectionJob(ctx context.Context, params *DescribePHIDetectionJobInput, optFns ...func(*Options)) (*DescribePHIDetectionJobOutput, error) 
 DescribeRxNormInferenceJob(ctx context.Context, params *DescribeRxNormInferenceJobInput, optFns ...func(*Options)) (*DescribeRxNormInferenceJobOutput, error) 
 DescribeSNOMEDCTInferenceJob(ctx context.Context, params *DescribeSNOMEDCTInferenceJobInput, optFns ...func(*Options)) (*DescribeSNOMEDCTInferenceJobOutput, error) 
 DetectEntities(ctx context.Context, params *DetectEntitiesInput, optFns ...func(*Options)) (*DetectEntitiesOutput, error) 
 DetectEntitiesV2(ctx context.Context, params *DetectEntitiesV2Input, optFns ...func(*Options)) (*DetectEntitiesV2Output, error) 
 DetectPHI(ctx context.Context, params *DetectPHIInput, optFns ...func(*Options)) (*DetectPHIOutput, error) 
 InferICD10CM(ctx context.Context, params *InferICD10CMInput, optFns ...func(*Options)) (*InferICD10CMOutput, error) 
 InferRxNorm(ctx context.Context, params *InferRxNormInput, optFns ...func(*Options)) (*InferRxNormOutput, error) 
 InferSNOMEDCT(ctx context.Context, params *InferSNOMEDCTInput, optFns ...func(*Options)) (*InferSNOMEDCTOutput, error) 
 ListEntitiesDetectionV2Jobs(ctx context.Context, params *ListEntitiesDetectionV2JobsInput, optFns ...func(*Options)) (*ListEntitiesDetectionV2JobsOutput, error) 
 ListICD10CMInferenceJobs(ctx context.Context, params *ListICD10CMInferenceJobsInput, optFns ...func(*Options)) (*ListICD10CMInferenceJobsOutput, error) 
 ListPHIDetectionJobs(ctx context.Context, params *ListPHIDetectionJobsInput, optFns ...func(*Options)) (*ListPHIDetectionJobsOutput, error) 
 ListRxNormInferenceJobs(ctx context.Context, params *ListRxNormInferenceJobsInput, optFns ...func(*Options)) (*ListRxNormInferenceJobsOutput, error) 
 ListSNOMEDCTInferenceJobs(ctx context.Context, params *ListSNOMEDCTInferenceJobsInput, optFns ...func(*Options)) (*ListSNOMEDCTInferenceJobsOutput, error) 
 StartEntitiesDetectionV2Job(ctx context.Context, params *StartEntitiesDetectionV2JobInput, optFns ...func(*Options)) (*StartEntitiesDetectionV2JobOutput, error) 
 StartICD10CMInferenceJob(ctx context.Context, params *StartICD10CMInferenceJobInput, optFns ...func(*Options)) (*StartICD10CMInferenceJobOutput, error) 
 StartPHIDetectionJob(ctx context.Context, params *StartPHIDetectionJobInput, optFns ...func(*Options)) (*StartPHIDetectionJobOutput, error) 
 StartRxNormInferenceJob(ctx context.Context, params *StartRxNormInferenceJobInput, optFns ...func(*Options)) (*StartRxNormInferenceJobOutput, error) 
 StartSNOMEDCTInferenceJob(ctx context.Context, params *StartSNOMEDCTInferenceJobInput, optFns ...func(*Options)) (*StartSNOMEDCTInferenceJobOutput, error) 
 StopEntitiesDetectionV2Job(ctx context.Context, params *StopEntitiesDetectionV2JobInput, optFns ...func(*Options)) (*StopEntitiesDetectionV2JobOutput, error) 
 StopICD10CMInferenceJob(ctx context.Context, params *StopICD10CMInferenceJobInput, optFns ...func(*Options)) (*StopICD10CMInferenceJobOutput, error) 
 StopPHIDetectionJob(ctx context.Context, params *StopPHIDetectionJobInput, optFns ...func(*Options)) (*StopPHIDetectionJobOutput, error) 
 StopRxNormInferenceJob(ctx context.Context, params *StopRxNormInferenceJobInput, optFns ...func(*Options)) (*StopRxNormInferenceJobOutput, error) 
 StopSNOMEDCTInferenceJob(ctx context.Context, params *StopSNOMEDCTInferenceJobInput, optFns ...func(*Options)) (*StopSNOMEDCTInferenceJobOutput, error) 
}
