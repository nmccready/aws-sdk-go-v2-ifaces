// Code generated by ./scripts/gen_tests.sh DO NOT EDIT.
package mediaconvert_test

// tests for the mediaconvert service interface for 
// this ../../../service/mediaconvert/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/mediaconvert"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/mediaconvert/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/mediaconvert/mediaconvert_iface"
	"github.com/stretchr/testify/assert"
)

func TestMediaconvertServiceCanBeMocked(t *testing.T) {
	var iface mediaconvert_iface.IClient
	iface = &mediaconvert.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := mediaconvert.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateCertificate", func(t *testing.T) {
        input := &mediaconvert.AssociateCertificateInput{}
        output := &mediaconvert.AssociateCertificateOutput{}

        mockClient.On("AssociateCertificate", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateCertificate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCancelJob", func(t *testing.T) {
        input := &mediaconvert.CancelJobInput{}
        output := &mediaconvert.CancelJobOutput{}

        mockClient.On("CancelJob", ctx, input).Return(output, nil)

        result, err := mockClient.CancelJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateJob", func(t *testing.T) {
        input := &mediaconvert.CreateJobInput{}
        output := &mediaconvert.CreateJobOutput{}

        mockClient.On("CreateJob", ctx, input).Return(output, nil)

        result, err := mockClient.CreateJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateJobTemplate", func(t *testing.T) {
        input := &mediaconvert.CreateJobTemplateInput{}
        output := &mediaconvert.CreateJobTemplateOutput{}

        mockClient.On("CreateJobTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.CreateJobTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreatePreset", func(t *testing.T) {
        input := &mediaconvert.CreatePresetInput{}
        output := &mediaconvert.CreatePresetOutput{}

        mockClient.On("CreatePreset", ctx, input).Return(output, nil)

        result, err := mockClient.CreatePreset(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateQueue", func(t *testing.T) {
        input := &mediaconvert.CreateQueueInput{}
        output := &mediaconvert.CreateQueueOutput{}

        mockClient.On("CreateQueue", ctx, input).Return(output, nil)

        result, err := mockClient.CreateQueue(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteJobTemplate", func(t *testing.T) {
        input := &mediaconvert.DeleteJobTemplateInput{}
        output := &mediaconvert.DeleteJobTemplateOutput{}

        mockClient.On("DeleteJobTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteJobTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeletePolicy", func(t *testing.T) {
        input := &mediaconvert.DeletePolicyInput{}
        output := &mediaconvert.DeletePolicyOutput{}

        mockClient.On("DeletePolicy", ctx, input).Return(output, nil)

        result, err := mockClient.DeletePolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeletePreset", func(t *testing.T) {
        input := &mediaconvert.DeletePresetInput{}
        output := &mediaconvert.DeletePresetOutput{}

        mockClient.On("DeletePreset", ctx, input).Return(output, nil)

        result, err := mockClient.DeletePreset(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteQueue", func(t *testing.T) {
        input := &mediaconvert.DeleteQueueInput{}
        output := &mediaconvert.DeleteQueueOutput{}

        mockClient.On("DeleteQueue", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteQueue(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeEndpoints", func(t *testing.T) {
        input := &mediaconvert.DescribeEndpointsInput{}
        output := &mediaconvert.DescribeEndpointsOutput{}

        mockClient.On("DescribeEndpoints", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeEndpoints(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateCertificate", func(t *testing.T) {
        input := &mediaconvert.DisassociateCertificateInput{}
        output := &mediaconvert.DisassociateCertificateOutput{}

        mockClient.On("DisassociateCertificate", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateCertificate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetJob", func(t *testing.T) {
        input := &mediaconvert.GetJobInput{}
        output := &mediaconvert.GetJobOutput{}

        mockClient.On("GetJob", ctx, input).Return(output, nil)

        result, err := mockClient.GetJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetJobTemplate", func(t *testing.T) {
        input := &mediaconvert.GetJobTemplateInput{}
        output := &mediaconvert.GetJobTemplateOutput{}

        mockClient.On("GetJobTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.GetJobTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetPolicy", func(t *testing.T) {
        input := &mediaconvert.GetPolicyInput{}
        output := &mediaconvert.GetPolicyOutput{}

        mockClient.On("GetPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.GetPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetPreset", func(t *testing.T) {
        input := &mediaconvert.GetPresetInput{}
        output := &mediaconvert.GetPresetOutput{}

        mockClient.On("GetPreset", ctx, input).Return(output, nil)

        result, err := mockClient.GetPreset(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetQueue", func(t *testing.T) {
        input := &mediaconvert.GetQueueInput{}
        output := &mediaconvert.GetQueueOutput{}

        mockClient.On("GetQueue", ctx, input).Return(output, nil)

        result, err := mockClient.GetQueue(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListJobTemplates", func(t *testing.T) {
        input := &mediaconvert.ListJobTemplatesInput{}
        output := &mediaconvert.ListJobTemplatesOutput{}

        mockClient.On("ListJobTemplates", ctx, input).Return(output, nil)

        result, err := mockClient.ListJobTemplates(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListJobs", func(t *testing.T) {
        input := &mediaconvert.ListJobsInput{}
        output := &mediaconvert.ListJobsOutput{}

        mockClient.On("ListJobs", ctx, input).Return(output, nil)

        result, err := mockClient.ListJobs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListPresets", func(t *testing.T) {
        input := &mediaconvert.ListPresetsInput{}
        output := &mediaconvert.ListPresetsOutput{}

        mockClient.On("ListPresets", ctx, input).Return(output, nil)

        result, err := mockClient.ListPresets(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListQueues", func(t *testing.T) {
        input := &mediaconvert.ListQueuesInput{}
        output := &mediaconvert.ListQueuesOutput{}

        mockClient.On("ListQueues", ctx, input).Return(output, nil)

        result, err := mockClient.ListQueues(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &mediaconvert.ListTagsForResourceInput{}
        output := &mediaconvert.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListVersions", func(t *testing.T) {
        input := &mediaconvert.ListVersionsInput{}
        output := &mediaconvert.ListVersionsOutput{}

        mockClient.On("ListVersions", ctx, input).Return(output, nil)

        result, err := mockClient.ListVersions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutPolicy", func(t *testing.T) {
        input := &mediaconvert.PutPolicyInput{}
        output := &mediaconvert.PutPolicyOutput{}

        mockClient.On("PutPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.PutPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSearchJobs", func(t *testing.T) {
        input := &mediaconvert.SearchJobsInput{}
        output := &mediaconvert.SearchJobsOutput{}

        mockClient.On("SearchJobs", ctx, input).Return(output, nil)

        result, err := mockClient.SearchJobs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &mediaconvert.TagResourceInput{}
        output := &mediaconvert.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &mediaconvert.UntagResourceInput{}
        output := &mediaconvert.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateJobTemplate", func(t *testing.T) {
        input := &mediaconvert.UpdateJobTemplateInput{}
        output := &mediaconvert.UpdateJobTemplateOutput{}

        mockClient.On("UpdateJobTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateJobTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdatePreset", func(t *testing.T) {
        input := &mediaconvert.UpdatePresetInput{}
        output := &mediaconvert.UpdatePresetOutput{}

        mockClient.On("UpdatePreset", ctx, input).Return(output, nil)

        result, err := mockClient.UpdatePreset(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateQueue", func(t *testing.T) {
        input := &mediaconvert.UpdateQueueInput{}
        output := &mediaconvert.UpdateQueueOutput{}

        mockClient.On("UpdateQueue", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateQueue(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
