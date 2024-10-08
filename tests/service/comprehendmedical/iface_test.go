// Code generated by ./scripts/gen_tests.sh DO NOT EDIT.
package comprehendmedical_test

// tests for the comprehendmedical service interface for 
// this ../../../service/comprehendmedical/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/comprehendmedical"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/comprehendmedical/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/comprehendmedical/comprehendmedical_iface"
	"github.com/stretchr/testify/assert"
)

func TestComprehendmedicalServiceCanBeMocked(t *testing.T) {
	var iface comprehendmedical_iface.IClient
	iface = &comprehendmedical.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := comprehendmedical.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeEntitiesDetectionV2Job", func(t *testing.T) {
        input := &comprehendmedical.DescribeEntitiesDetectionV2JobInput{}
        output := &comprehendmedical.DescribeEntitiesDetectionV2JobOutput{}

        mockClient.On("DescribeEntitiesDetectionV2Job", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeEntitiesDetectionV2Job(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeICD10CMInferenceJob", func(t *testing.T) {
        input := &comprehendmedical.DescribeICD10CMInferenceJobInput{}
        output := &comprehendmedical.DescribeICD10CMInferenceJobOutput{}

        mockClient.On("DescribeICD10CMInferenceJob", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeICD10CMInferenceJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribePHIDetectionJob", func(t *testing.T) {
        input := &comprehendmedical.DescribePHIDetectionJobInput{}
        output := &comprehendmedical.DescribePHIDetectionJobOutput{}

        mockClient.On("DescribePHIDetectionJob", ctx, input).Return(output, nil)

        result, err := mockClient.DescribePHIDetectionJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeRxNormInferenceJob", func(t *testing.T) {
        input := &comprehendmedical.DescribeRxNormInferenceJobInput{}
        output := &comprehendmedical.DescribeRxNormInferenceJobOutput{}

        mockClient.On("DescribeRxNormInferenceJob", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeRxNormInferenceJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeSNOMEDCTInferenceJob", func(t *testing.T) {
        input := &comprehendmedical.DescribeSNOMEDCTInferenceJobInput{}
        output := &comprehendmedical.DescribeSNOMEDCTInferenceJobOutput{}

        mockClient.On("DescribeSNOMEDCTInferenceJob", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeSNOMEDCTInferenceJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDetectEntities", func(t *testing.T) {
        input := &comprehendmedical.DetectEntitiesInput{}
        output := &comprehendmedical.DetectEntitiesOutput{}

        mockClient.On("DetectEntities", ctx, input).Return(output, nil)

        result, err := mockClient.DetectEntities(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDetectEntitiesV2", func(t *testing.T) {
        input := &comprehendmedical.DetectEntitiesV2Input{}
        output := &comprehendmedical.DetectEntitiesV2Output{}

        mockClient.On("DetectEntitiesV2", ctx, input).Return(output, nil)

        result, err := mockClient.DetectEntitiesV2(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDetectPHI", func(t *testing.T) {
        input := &comprehendmedical.DetectPHIInput{}
        output := &comprehendmedical.DetectPHIOutput{}

        mockClient.On("DetectPHI", ctx, input).Return(output, nil)

        result, err := mockClient.DetectPHI(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestInferICD10CM", func(t *testing.T) {
        input := &comprehendmedical.InferICD10CMInput{}
        output := &comprehendmedical.InferICD10CMOutput{}

        mockClient.On("InferICD10CM", ctx, input).Return(output, nil)

        result, err := mockClient.InferICD10CM(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestInferRxNorm", func(t *testing.T) {
        input := &comprehendmedical.InferRxNormInput{}
        output := &comprehendmedical.InferRxNormOutput{}

        mockClient.On("InferRxNorm", ctx, input).Return(output, nil)

        result, err := mockClient.InferRxNorm(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestInferSNOMEDCT", func(t *testing.T) {
        input := &comprehendmedical.InferSNOMEDCTInput{}
        output := &comprehendmedical.InferSNOMEDCTOutput{}

        mockClient.On("InferSNOMEDCT", ctx, input).Return(output, nil)

        result, err := mockClient.InferSNOMEDCT(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListEntitiesDetectionV2Jobs", func(t *testing.T) {
        input := &comprehendmedical.ListEntitiesDetectionV2JobsInput{}
        output := &comprehendmedical.ListEntitiesDetectionV2JobsOutput{}

        mockClient.On("ListEntitiesDetectionV2Jobs", ctx, input).Return(output, nil)

        result, err := mockClient.ListEntitiesDetectionV2Jobs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListICD10CMInferenceJobs", func(t *testing.T) {
        input := &comprehendmedical.ListICD10CMInferenceJobsInput{}
        output := &comprehendmedical.ListICD10CMInferenceJobsOutput{}

        mockClient.On("ListICD10CMInferenceJobs", ctx, input).Return(output, nil)

        result, err := mockClient.ListICD10CMInferenceJobs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListPHIDetectionJobs", func(t *testing.T) {
        input := &comprehendmedical.ListPHIDetectionJobsInput{}
        output := &comprehendmedical.ListPHIDetectionJobsOutput{}

        mockClient.On("ListPHIDetectionJobs", ctx, input).Return(output, nil)

        result, err := mockClient.ListPHIDetectionJobs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListRxNormInferenceJobs", func(t *testing.T) {
        input := &comprehendmedical.ListRxNormInferenceJobsInput{}
        output := &comprehendmedical.ListRxNormInferenceJobsOutput{}

        mockClient.On("ListRxNormInferenceJobs", ctx, input).Return(output, nil)

        result, err := mockClient.ListRxNormInferenceJobs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListSNOMEDCTInferenceJobs", func(t *testing.T) {
        input := &comprehendmedical.ListSNOMEDCTInferenceJobsInput{}
        output := &comprehendmedical.ListSNOMEDCTInferenceJobsOutput{}

        mockClient.On("ListSNOMEDCTInferenceJobs", ctx, input).Return(output, nil)

        result, err := mockClient.ListSNOMEDCTInferenceJobs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartEntitiesDetectionV2Job", func(t *testing.T) {
        input := &comprehendmedical.StartEntitiesDetectionV2JobInput{}
        output := &comprehendmedical.StartEntitiesDetectionV2JobOutput{}

        mockClient.On("StartEntitiesDetectionV2Job", ctx, input).Return(output, nil)

        result, err := mockClient.StartEntitiesDetectionV2Job(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartICD10CMInferenceJob", func(t *testing.T) {
        input := &comprehendmedical.StartICD10CMInferenceJobInput{}
        output := &comprehendmedical.StartICD10CMInferenceJobOutput{}

        mockClient.On("StartICD10CMInferenceJob", ctx, input).Return(output, nil)

        result, err := mockClient.StartICD10CMInferenceJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartPHIDetectionJob", func(t *testing.T) {
        input := &comprehendmedical.StartPHIDetectionJobInput{}
        output := &comprehendmedical.StartPHIDetectionJobOutput{}

        mockClient.On("StartPHIDetectionJob", ctx, input).Return(output, nil)

        result, err := mockClient.StartPHIDetectionJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartRxNormInferenceJob", func(t *testing.T) {
        input := &comprehendmedical.StartRxNormInferenceJobInput{}
        output := &comprehendmedical.StartRxNormInferenceJobOutput{}

        mockClient.On("StartRxNormInferenceJob", ctx, input).Return(output, nil)

        result, err := mockClient.StartRxNormInferenceJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartSNOMEDCTInferenceJob", func(t *testing.T) {
        input := &comprehendmedical.StartSNOMEDCTInferenceJobInput{}
        output := &comprehendmedical.StartSNOMEDCTInferenceJobOutput{}

        mockClient.On("StartSNOMEDCTInferenceJob", ctx, input).Return(output, nil)

        result, err := mockClient.StartSNOMEDCTInferenceJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopEntitiesDetectionV2Job", func(t *testing.T) {
        input := &comprehendmedical.StopEntitiesDetectionV2JobInput{}
        output := &comprehendmedical.StopEntitiesDetectionV2JobOutput{}

        mockClient.On("StopEntitiesDetectionV2Job", ctx, input).Return(output, nil)

        result, err := mockClient.StopEntitiesDetectionV2Job(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopICD10CMInferenceJob", func(t *testing.T) {
        input := &comprehendmedical.StopICD10CMInferenceJobInput{}
        output := &comprehendmedical.StopICD10CMInferenceJobOutput{}

        mockClient.On("StopICD10CMInferenceJob", ctx, input).Return(output, nil)

        result, err := mockClient.StopICD10CMInferenceJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopPHIDetectionJob", func(t *testing.T) {
        input := &comprehendmedical.StopPHIDetectionJobInput{}
        output := &comprehendmedical.StopPHIDetectionJobOutput{}

        mockClient.On("StopPHIDetectionJob", ctx, input).Return(output, nil)

        result, err := mockClient.StopPHIDetectionJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopRxNormInferenceJob", func(t *testing.T) {
        input := &comprehendmedical.StopRxNormInferenceJobInput{}
        output := &comprehendmedical.StopRxNormInferenceJobOutput{}

        mockClient.On("StopRxNormInferenceJob", ctx, input).Return(output, nil)

        result, err := mockClient.StopRxNormInferenceJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopSNOMEDCTInferenceJob", func(t *testing.T) {
        input := &comprehendmedical.StopSNOMEDCTInferenceJobInput{}
        output := &comprehendmedical.StopSNOMEDCTInferenceJobOutput{}

        mockClient.On("StopSNOMEDCTInferenceJob", ctx, input).Return(output, nil)

        result, err := mockClient.StopSNOMEDCTInferenceJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
