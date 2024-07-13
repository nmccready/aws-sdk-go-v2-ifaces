package computeoptimizer_test

// tests for the computeoptimizer service interface for this ../../../service/computeoptimizer/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/computeoptimizer"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/computeoptimizer/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/computeoptimizer/computeoptimizer_iface"
	"github.com/stretchr/testify/assert"
)

func TestComputeoptimizerServiceCanBeMocked(t *testing.T) {
	var iface computeoptimizer_iface.IClient
	iface = &computeoptimizer.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := computeoptimizer.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteRecommendationPreferences", func(t *testing.T) {
        input := &computeoptimizer.DeleteRecommendationPreferencesInput{}
        output := &computeoptimizer.DeleteRecommendationPreferencesOutput{}

        mockClient.On("DeleteRecommendationPreferences", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteRecommendationPreferences(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeRecommendationExportJobs", func(t *testing.T) {
        input := &computeoptimizer.DescribeRecommendationExportJobsInput{}
        output := &computeoptimizer.DescribeRecommendationExportJobsOutput{}

        mockClient.On("DescribeRecommendationExportJobs", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeRecommendationExportJobs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestExportAutoScalingGroupRecommendations", func(t *testing.T) {
        input := &computeoptimizer.ExportAutoScalingGroupRecommendationsInput{}
        output := &computeoptimizer.ExportAutoScalingGroupRecommendationsOutput{}

        mockClient.On("ExportAutoScalingGroupRecommendations", ctx, input).Return(output, nil)

        result, err := mockClient.ExportAutoScalingGroupRecommendations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestExportEBSVolumeRecommendations", func(t *testing.T) {
        input := &computeoptimizer.ExportEBSVolumeRecommendationsInput{}
        output := &computeoptimizer.ExportEBSVolumeRecommendationsOutput{}

        mockClient.On("ExportEBSVolumeRecommendations", ctx, input).Return(output, nil)

        result, err := mockClient.ExportEBSVolumeRecommendations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestExportEC2InstanceRecommendations", func(t *testing.T) {
        input := &computeoptimizer.ExportEC2InstanceRecommendationsInput{}
        output := &computeoptimizer.ExportEC2InstanceRecommendationsOutput{}

        mockClient.On("ExportEC2InstanceRecommendations", ctx, input).Return(output, nil)

        result, err := mockClient.ExportEC2InstanceRecommendations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestExportECSServiceRecommendations", func(t *testing.T) {
        input := &computeoptimizer.ExportECSServiceRecommendationsInput{}
        output := &computeoptimizer.ExportECSServiceRecommendationsOutput{}

        mockClient.On("ExportECSServiceRecommendations", ctx, input).Return(output, nil)

        result, err := mockClient.ExportECSServiceRecommendations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestExportLambdaFunctionRecommendations", func(t *testing.T) {
        input := &computeoptimizer.ExportLambdaFunctionRecommendationsInput{}
        output := &computeoptimizer.ExportLambdaFunctionRecommendationsOutput{}

        mockClient.On("ExportLambdaFunctionRecommendations", ctx, input).Return(output, nil)

        result, err := mockClient.ExportLambdaFunctionRecommendations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestExportLicenseRecommendations", func(t *testing.T) {
        input := &computeoptimizer.ExportLicenseRecommendationsInput{}
        output := &computeoptimizer.ExportLicenseRecommendationsOutput{}

        mockClient.On("ExportLicenseRecommendations", ctx, input).Return(output, nil)

        result, err := mockClient.ExportLicenseRecommendations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestExportRDSDatabaseRecommendations", func(t *testing.T) {
        input := &computeoptimizer.ExportRDSDatabaseRecommendationsInput{}
        output := &computeoptimizer.ExportRDSDatabaseRecommendationsOutput{}

        mockClient.On("ExportRDSDatabaseRecommendations", ctx, input).Return(output, nil)

        result, err := mockClient.ExportRDSDatabaseRecommendations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetAutoScalingGroupRecommendations", func(t *testing.T) {
        input := &computeoptimizer.GetAutoScalingGroupRecommendationsInput{}
        output := &computeoptimizer.GetAutoScalingGroupRecommendationsOutput{}

        mockClient.On("GetAutoScalingGroupRecommendations", ctx, input).Return(output, nil)

        result, err := mockClient.GetAutoScalingGroupRecommendations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetEBSVolumeRecommendations", func(t *testing.T) {
        input := &computeoptimizer.GetEBSVolumeRecommendationsInput{}
        output := &computeoptimizer.GetEBSVolumeRecommendationsOutput{}

        mockClient.On("GetEBSVolumeRecommendations", ctx, input).Return(output, nil)

        result, err := mockClient.GetEBSVolumeRecommendations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetEC2InstanceRecommendations", func(t *testing.T) {
        input := &computeoptimizer.GetEC2InstanceRecommendationsInput{}
        output := &computeoptimizer.GetEC2InstanceRecommendationsOutput{}

        mockClient.On("GetEC2InstanceRecommendations", ctx, input).Return(output, nil)

        result, err := mockClient.GetEC2InstanceRecommendations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetEC2RecommendationProjectedMetrics", func(t *testing.T) {
        input := &computeoptimizer.GetEC2RecommendationProjectedMetricsInput{}
        output := &computeoptimizer.GetEC2RecommendationProjectedMetricsOutput{}

        mockClient.On("GetEC2RecommendationProjectedMetrics", ctx, input).Return(output, nil)

        result, err := mockClient.GetEC2RecommendationProjectedMetrics(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetECSServiceRecommendationProjectedMetrics", func(t *testing.T) {
        input := &computeoptimizer.GetECSServiceRecommendationProjectedMetricsInput{}
        output := &computeoptimizer.GetECSServiceRecommendationProjectedMetricsOutput{}

        mockClient.On("GetECSServiceRecommendationProjectedMetrics", ctx, input).Return(output, nil)

        result, err := mockClient.GetECSServiceRecommendationProjectedMetrics(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetECSServiceRecommendations", func(t *testing.T) {
        input := &computeoptimizer.GetECSServiceRecommendationsInput{}
        output := &computeoptimizer.GetECSServiceRecommendationsOutput{}

        mockClient.On("GetECSServiceRecommendations", ctx, input).Return(output, nil)

        result, err := mockClient.GetECSServiceRecommendations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetEffectiveRecommendationPreferences", func(t *testing.T) {
        input := &computeoptimizer.GetEffectiveRecommendationPreferencesInput{}
        output := &computeoptimizer.GetEffectiveRecommendationPreferencesOutput{}

        mockClient.On("GetEffectiveRecommendationPreferences", ctx, input).Return(output, nil)

        result, err := mockClient.GetEffectiveRecommendationPreferences(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetEnrollmentStatus", func(t *testing.T) {
        input := &computeoptimizer.GetEnrollmentStatusInput{}
        output := &computeoptimizer.GetEnrollmentStatusOutput{}

        mockClient.On("GetEnrollmentStatus", ctx, input).Return(output, nil)

        result, err := mockClient.GetEnrollmentStatus(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetEnrollmentStatusesForOrganization", func(t *testing.T) {
        input := &computeoptimizer.GetEnrollmentStatusesForOrganizationInput{}
        output := &computeoptimizer.GetEnrollmentStatusesForOrganizationOutput{}

        mockClient.On("GetEnrollmentStatusesForOrganization", ctx, input).Return(output, nil)

        result, err := mockClient.GetEnrollmentStatusesForOrganization(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetLambdaFunctionRecommendations", func(t *testing.T) {
        input := &computeoptimizer.GetLambdaFunctionRecommendationsInput{}
        output := &computeoptimizer.GetLambdaFunctionRecommendationsOutput{}

        mockClient.On("GetLambdaFunctionRecommendations", ctx, input).Return(output, nil)

        result, err := mockClient.GetLambdaFunctionRecommendations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetLicenseRecommendations", func(t *testing.T) {
        input := &computeoptimizer.GetLicenseRecommendationsInput{}
        output := &computeoptimizer.GetLicenseRecommendationsOutput{}

        mockClient.On("GetLicenseRecommendations", ctx, input).Return(output, nil)

        result, err := mockClient.GetLicenseRecommendations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetRDSDatabaseRecommendationProjectedMetrics", func(t *testing.T) {
        input := &computeoptimizer.GetRDSDatabaseRecommendationProjectedMetricsInput{}
        output := &computeoptimizer.GetRDSDatabaseRecommendationProjectedMetricsOutput{}

        mockClient.On("GetRDSDatabaseRecommendationProjectedMetrics", ctx, input).Return(output, nil)

        result, err := mockClient.GetRDSDatabaseRecommendationProjectedMetrics(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetRDSDatabaseRecommendations", func(t *testing.T) {
        input := &computeoptimizer.GetRDSDatabaseRecommendationsInput{}
        output := &computeoptimizer.GetRDSDatabaseRecommendationsOutput{}

        mockClient.On("GetRDSDatabaseRecommendations", ctx, input).Return(output, nil)

        result, err := mockClient.GetRDSDatabaseRecommendations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetRecommendationPreferences", func(t *testing.T) {
        input := &computeoptimizer.GetRecommendationPreferencesInput{}
        output := &computeoptimizer.GetRecommendationPreferencesOutput{}

        mockClient.On("GetRecommendationPreferences", ctx, input).Return(output, nil)

        result, err := mockClient.GetRecommendationPreferences(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetRecommendationSummaries", func(t *testing.T) {
        input := &computeoptimizer.GetRecommendationSummariesInput{}
        output := &computeoptimizer.GetRecommendationSummariesOutput{}

        mockClient.On("GetRecommendationSummaries", ctx, input).Return(output, nil)

        result, err := mockClient.GetRecommendationSummaries(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutRecommendationPreferences", func(t *testing.T) {
        input := &computeoptimizer.PutRecommendationPreferencesInput{}
        output := &computeoptimizer.PutRecommendationPreferencesOutput{}

        mockClient.On("PutRecommendationPreferences", ctx, input).Return(output, nil)

        result, err := mockClient.PutRecommendationPreferences(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateEnrollmentStatus", func(t *testing.T) {
        input := &computeoptimizer.UpdateEnrollmentStatusInput{}
        output := &computeoptimizer.UpdateEnrollmentStatusOutput{}

        mockClient.On("UpdateEnrollmentStatus", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateEnrollmentStatus(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
