
package iotjobsdataplane

import (
    "github.com/aws/aws-sdk-go-v2/service/iotjobsdataplane"
)

// IIotjobsdataplane defines the interface for iotjobsdataplane
type IIotjobsdataplane interface {
 Options() Options 
 DescribeJobExecution(ctx context.Context, params *DescribeJobExecutionInput, optFns ...func(*Options)) (*DescribeJobExecutionOutput, error) 
 GetPendingJobExecutions(ctx context.Context, params *GetPendingJobExecutionsInput, optFns ...func(*Options)) (*GetPendingJobExecutionsOutput, error) 
 StartNextPendingJobExecution(ctx context.Context, params *StartNextPendingJobExecutionInput, optFns ...func(*Options)) (*StartNextPendingJobExecutionOutput, error) 
 UpdateJobExecution(ctx context.Context, params *UpdateJobExecutionInput, optFns ...func(*Options)) (*UpdateJobExecutionOutput, error) 
}
