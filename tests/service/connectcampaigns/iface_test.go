// Code generated by ./scripts/gen_tests.sh DO NOT EDIT.
package connectcampaigns_test

// tests for the connectcampaigns service interface for 
// this ../../../service/connectcampaigns/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/connectcampaigns"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/connectcampaigns/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/connectcampaigns/connectcampaigns_iface"
	"github.com/stretchr/testify/assert"
)

func TestConnectcampaignsServiceCanBeMocked(t *testing.T) {
	var iface connectcampaigns_iface.IClient
	iface = &connectcampaigns.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := connectcampaigns.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateCampaign", func(t *testing.T) {
        input := &connectcampaigns.CreateCampaignInput{}
        output := &connectcampaigns.CreateCampaignOutput{}

        mockClient.On("CreateCampaign", ctx, input).Return(output, nil)

        result, err := mockClient.CreateCampaign(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteCampaign", func(t *testing.T) {
        input := &connectcampaigns.DeleteCampaignInput{}
        output := &connectcampaigns.DeleteCampaignOutput{}

        mockClient.On("DeleteCampaign", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteCampaign(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteConnectInstanceConfig", func(t *testing.T) {
        input := &connectcampaigns.DeleteConnectInstanceConfigInput{}
        output := &connectcampaigns.DeleteConnectInstanceConfigOutput{}

        mockClient.On("DeleteConnectInstanceConfig", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteConnectInstanceConfig(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteInstanceOnboardingJob", func(t *testing.T) {
        input := &connectcampaigns.DeleteInstanceOnboardingJobInput{}
        output := &connectcampaigns.DeleteInstanceOnboardingJobOutput{}

        mockClient.On("DeleteInstanceOnboardingJob", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteInstanceOnboardingJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeCampaign", func(t *testing.T) {
        input := &connectcampaigns.DescribeCampaignInput{}
        output := &connectcampaigns.DescribeCampaignOutput{}

        mockClient.On("DescribeCampaign", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeCampaign(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetCampaignState", func(t *testing.T) {
        input := &connectcampaigns.GetCampaignStateInput{}
        output := &connectcampaigns.GetCampaignStateOutput{}

        mockClient.On("GetCampaignState", ctx, input).Return(output, nil)

        result, err := mockClient.GetCampaignState(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetCampaignStateBatch", func(t *testing.T) {
        input := &connectcampaigns.GetCampaignStateBatchInput{}
        output := &connectcampaigns.GetCampaignStateBatchOutput{}

        mockClient.On("GetCampaignStateBatch", ctx, input).Return(output, nil)

        result, err := mockClient.GetCampaignStateBatch(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetConnectInstanceConfig", func(t *testing.T) {
        input := &connectcampaigns.GetConnectInstanceConfigInput{}
        output := &connectcampaigns.GetConnectInstanceConfigOutput{}

        mockClient.On("GetConnectInstanceConfig", ctx, input).Return(output, nil)

        result, err := mockClient.GetConnectInstanceConfig(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetInstanceOnboardingJobStatus", func(t *testing.T) {
        input := &connectcampaigns.GetInstanceOnboardingJobStatusInput{}
        output := &connectcampaigns.GetInstanceOnboardingJobStatusOutput{}

        mockClient.On("GetInstanceOnboardingJobStatus", ctx, input).Return(output, nil)

        result, err := mockClient.GetInstanceOnboardingJobStatus(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListCampaigns", func(t *testing.T) {
        input := &connectcampaigns.ListCampaignsInput{}
        output := &connectcampaigns.ListCampaignsOutput{}

        mockClient.On("ListCampaigns", ctx, input).Return(output, nil)

        result, err := mockClient.ListCampaigns(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &connectcampaigns.ListTagsForResourceInput{}
        output := &connectcampaigns.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPauseCampaign", func(t *testing.T) {
        input := &connectcampaigns.PauseCampaignInput{}
        output := &connectcampaigns.PauseCampaignOutput{}

        mockClient.On("PauseCampaign", ctx, input).Return(output, nil)

        result, err := mockClient.PauseCampaign(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutDialRequestBatch", func(t *testing.T) {
        input := &connectcampaigns.PutDialRequestBatchInput{}
        output := &connectcampaigns.PutDialRequestBatchOutput{}

        mockClient.On("PutDialRequestBatch", ctx, input).Return(output, nil)

        result, err := mockClient.PutDialRequestBatch(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestResumeCampaign", func(t *testing.T) {
        input := &connectcampaigns.ResumeCampaignInput{}
        output := &connectcampaigns.ResumeCampaignOutput{}

        mockClient.On("ResumeCampaign", ctx, input).Return(output, nil)

        result, err := mockClient.ResumeCampaign(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartCampaign", func(t *testing.T) {
        input := &connectcampaigns.StartCampaignInput{}
        output := &connectcampaigns.StartCampaignOutput{}

        mockClient.On("StartCampaign", ctx, input).Return(output, nil)

        result, err := mockClient.StartCampaign(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartInstanceOnboardingJob", func(t *testing.T) {
        input := &connectcampaigns.StartInstanceOnboardingJobInput{}
        output := &connectcampaigns.StartInstanceOnboardingJobOutput{}

        mockClient.On("StartInstanceOnboardingJob", ctx, input).Return(output, nil)

        result, err := mockClient.StartInstanceOnboardingJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopCampaign", func(t *testing.T) {
        input := &connectcampaigns.StopCampaignInput{}
        output := &connectcampaigns.StopCampaignOutput{}

        mockClient.On("StopCampaign", ctx, input).Return(output, nil)

        result, err := mockClient.StopCampaign(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &connectcampaigns.TagResourceInput{}
        output := &connectcampaigns.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &connectcampaigns.UntagResourceInput{}
        output := &connectcampaigns.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateCampaignDialerConfig", func(t *testing.T) {
        input := &connectcampaigns.UpdateCampaignDialerConfigInput{}
        output := &connectcampaigns.UpdateCampaignDialerConfigOutput{}

        mockClient.On("UpdateCampaignDialerConfig", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateCampaignDialerConfig(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateCampaignName", func(t *testing.T) {
        input := &connectcampaigns.UpdateCampaignNameInput{}
        output := &connectcampaigns.UpdateCampaignNameOutput{}

        mockClient.On("UpdateCampaignName", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateCampaignName(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateCampaignOutboundCallConfig", func(t *testing.T) {
        input := &connectcampaigns.UpdateCampaignOutboundCallConfigInput{}
        output := &connectcampaigns.UpdateCampaignOutboundCallConfigOutput{}

        mockClient.On("UpdateCampaignOutboundCallConfig", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateCampaignOutboundCallConfig(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
