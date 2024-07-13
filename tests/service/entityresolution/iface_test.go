package entityresolution_test

// tests for the entityresolution service interface for this ../../../service/entityresolution/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/entityresolution"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/entityresolution/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/entityresolution/entityresolution_iface"
	"github.com/stretchr/testify/assert"
)

func TestEntityresolutionServiceCanBeMocked(t *testing.T) {
	var iface entityresolution_iface.IClient
	iface = &entityresolution.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := entityresolution.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAddPolicyStatement", func(t *testing.T) {
        input := &entityresolution.AddPolicyStatementInput{}
        output := &entityresolution.AddPolicyStatementOutput{}

        mockClient.On("AddPolicyStatement", ctx, input).Return(output, nil)

        result, err := mockClient.AddPolicyStatement(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchDeleteUniqueId", func(t *testing.T) {
        input := &entityresolution.BatchDeleteUniqueIdInput{}
        output := &entityresolution.BatchDeleteUniqueIdOutput{}

        mockClient.On("BatchDeleteUniqueId", ctx, input).Return(output, nil)

        result, err := mockClient.BatchDeleteUniqueId(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateIdMappingWorkflow", func(t *testing.T) {
        input := &entityresolution.CreateIdMappingWorkflowInput{}
        output := &entityresolution.CreateIdMappingWorkflowOutput{}

        mockClient.On("CreateIdMappingWorkflow", ctx, input).Return(output, nil)

        result, err := mockClient.CreateIdMappingWorkflow(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateIdNamespace", func(t *testing.T) {
        input := &entityresolution.CreateIdNamespaceInput{}
        output := &entityresolution.CreateIdNamespaceOutput{}

        mockClient.On("CreateIdNamespace", ctx, input).Return(output, nil)

        result, err := mockClient.CreateIdNamespace(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateMatchingWorkflow", func(t *testing.T) {
        input := &entityresolution.CreateMatchingWorkflowInput{}
        output := &entityresolution.CreateMatchingWorkflowOutput{}

        mockClient.On("CreateMatchingWorkflow", ctx, input).Return(output, nil)

        result, err := mockClient.CreateMatchingWorkflow(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateSchemaMapping", func(t *testing.T) {
        input := &entityresolution.CreateSchemaMappingInput{}
        output := &entityresolution.CreateSchemaMappingOutput{}

        mockClient.On("CreateSchemaMapping", ctx, input).Return(output, nil)

        result, err := mockClient.CreateSchemaMapping(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteIdMappingWorkflow", func(t *testing.T) {
        input := &entityresolution.DeleteIdMappingWorkflowInput{}
        output := &entityresolution.DeleteIdMappingWorkflowOutput{}

        mockClient.On("DeleteIdMappingWorkflow", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteIdMappingWorkflow(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteIdNamespace", func(t *testing.T) {
        input := &entityresolution.DeleteIdNamespaceInput{}
        output := &entityresolution.DeleteIdNamespaceOutput{}

        mockClient.On("DeleteIdNamespace", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteIdNamespace(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteMatchingWorkflow", func(t *testing.T) {
        input := &entityresolution.DeleteMatchingWorkflowInput{}
        output := &entityresolution.DeleteMatchingWorkflowOutput{}

        mockClient.On("DeleteMatchingWorkflow", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteMatchingWorkflow(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeletePolicyStatement", func(t *testing.T) {
        input := &entityresolution.DeletePolicyStatementInput{}
        output := &entityresolution.DeletePolicyStatementOutput{}

        mockClient.On("DeletePolicyStatement", ctx, input).Return(output, nil)

        result, err := mockClient.DeletePolicyStatement(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteSchemaMapping", func(t *testing.T) {
        input := &entityresolution.DeleteSchemaMappingInput{}
        output := &entityresolution.DeleteSchemaMappingOutput{}

        mockClient.On("DeleteSchemaMapping", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteSchemaMapping(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetIdMappingJob", func(t *testing.T) {
        input := &entityresolution.GetIdMappingJobInput{}
        output := &entityresolution.GetIdMappingJobOutput{}

        mockClient.On("GetIdMappingJob", ctx, input).Return(output, nil)

        result, err := mockClient.GetIdMappingJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetIdMappingWorkflow", func(t *testing.T) {
        input := &entityresolution.GetIdMappingWorkflowInput{}
        output := &entityresolution.GetIdMappingWorkflowOutput{}

        mockClient.On("GetIdMappingWorkflow", ctx, input).Return(output, nil)

        result, err := mockClient.GetIdMappingWorkflow(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetIdNamespace", func(t *testing.T) {
        input := &entityresolution.GetIdNamespaceInput{}
        output := &entityresolution.GetIdNamespaceOutput{}

        mockClient.On("GetIdNamespace", ctx, input).Return(output, nil)

        result, err := mockClient.GetIdNamespace(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetMatchId", func(t *testing.T) {
        input := &entityresolution.GetMatchIdInput{}
        output := &entityresolution.GetMatchIdOutput{}

        mockClient.On("GetMatchId", ctx, input).Return(output, nil)

        result, err := mockClient.GetMatchId(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetMatchingJob", func(t *testing.T) {
        input := &entityresolution.GetMatchingJobInput{}
        output := &entityresolution.GetMatchingJobOutput{}

        mockClient.On("GetMatchingJob", ctx, input).Return(output, nil)

        result, err := mockClient.GetMatchingJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetMatchingWorkflow", func(t *testing.T) {
        input := &entityresolution.GetMatchingWorkflowInput{}
        output := &entityresolution.GetMatchingWorkflowOutput{}

        mockClient.On("GetMatchingWorkflow", ctx, input).Return(output, nil)

        result, err := mockClient.GetMatchingWorkflow(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetPolicy", func(t *testing.T) {
        input := &entityresolution.GetPolicyInput{}
        output := &entityresolution.GetPolicyOutput{}

        mockClient.On("GetPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.GetPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetProviderService", func(t *testing.T) {
        input := &entityresolution.GetProviderServiceInput{}
        output := &entityresolution.GetProviderServiceOutput{}

        mockClient.On("GetProviderService", ctx, input).Return(output, nil)

        result, err := mockClient.GetProviderService(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSchemaMapping", func(t *testing.T) {
        input := &entityresolution.GetSchemaMappingInput{}
        output := &entityresolution.GetSchemaMappingOutput{}

        mockClient.On("GetSchemaMapping", ctx, input).Return(output, nil)

        result, err := mockClient.GetSchemaMapping(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListIdMappingJobs", func(t *testing.T) {
        input := &entityresolution.ListIdMappingJobsInput{}
        output := &entityresolution.ListIdMappingJobsOutput{}

        mockClient.On("ListIdMappingJobs", ctx, input).Return(output, nil)

        result, err := mockClient.ListIdMappingJobs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListIdMappingWorkflows", func(t *testing.T) {
        input := &entityresolution.ListIdMappingWorkflowsInput{}
        output := &entityresolution.ListIdMappingWorkflowsOutput{}

        mockClient.On("ListIdMappingWorkflows", ctx, input).Return(output, nil)

        result, err := mockClient.ListIdMappingWorkflows(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListIdNamespaces", func(t *testing.T) {
        input := &entityresolution.ListIdNamespacesInput{}
        output := &entityresolution.ListIdNamespacesOutput{}

        mockClient.On("ListIdNamespaces", ctx, input).Return(output, nil)

        result, err := mockClient.ListIdNamespaces(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListMatchingJobs", func(t *testing.T) {
        input := &entityresolution.ListMatchingJobsInput{}
        output := &entityresolution.ListMatchingJobsOutput{}

        mockClient.On("ListMatchingJobs", ctx, input).Return(output, nil)

        result, err := mockClient.ListMatchingJobs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListMatchingWorkflows", func(t *testing.T) {
        input := &entityresolution.ListMatchingWorkflowsInput{}
        output := &entityresolution.ListMatchingWorkflowsOutput{}

        mockClient.On("ListMatchingWorkflows", ctx, input).Return(output, nil)

        result, err := mockClient.ListMatchingWorkflows(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListProviderServices", func(t *testing.T) {
        input := &entityresolution.ListProviderServicesInput{}
        output := &entityresolution.ListProviderServicesOutput{}

        mockClient.On("ListProviderServices", ctx, input).Return(output, nil)

        result, err := mockClient.ListProviderServices(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListSchemaMappings", func(t *testing.T) {
        input := &entityresolution.ListSchemaMappingsInput{}
        output := &entityresolution.ListSchemaMappingsOutput{}

        mockClient.On("ListSchemaMappings", ctx, input).Return(output, nil)

        result, err := mockClient.ListSchemaMappings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &entityresolution.ListTagsForResourceInput{}
        output := &entityresolution.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutPolicy", func(t *testing.T) {
        input := &entityresolution.PutPolicyInput{}
        output := &entityresolution.PutPolicyOutput{}

        mockClient.On("PutPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.PutPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartIdMappingJob", func(t *testing.T) {
        input := &entityresolution.StartIdMappingJobInput{}
        output := &entityresolution.StartIdMappingJobOutput{}

        mockClient.On("StartIdMappingJob", ctx, input).Return(output, nil)

        result, err := mockClient.StartIdMappingJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartMatchingJob", func(t *testing.T) {
        input := &entityresolution.StartMatchingJobInput{}
        output := &entityresolution.StartMatchingJobOutput{}

        mockClient.On("StartMatchingJob", ctx, input).Return(output, nil)

        result, err := mockClient.StartMatchingJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &entityresolution.TagResourceInput{}
        output := &entityresolution.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &entityresolution.UntagResourceInput{}
        output := &entityresolution.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateIdMappingWorkflow", func(t *testing.T) {
        input := &entityresolution.UpdateIdMappingWorkflowInput{}
        output := &entityresolution.UpdateIdMappingWorkflowOutput{}

        mockClient.On("UpdateIdMappingWorkflow", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateIdMappingWorkflow(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateIdNamespace", func(t *testing.T) {
        input := &entityresolution.UpdateIdNamespaceInput{}
        output := &entityresolution.UpdateIdNamespaceOutput{}

        mockClient.On("UpdateIdNamespace", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateIdNamespace(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateMatchingWorkflow", func(t *testing.T) {
        input := &entityresolution.UpdateMatchingWorkflowInput{}
        output := &entityresolution.UpdateMatchingWorkflowOutput{}

        mockClient.On("UpdateMatchingWorkflow", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateMatchingWorkflow(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateSchemaMapping", func(t *testing.T) {
        input := &entityresolution.UpdateSchemaMappingInput{}
        output := &entityresolution.UpdateSchemaMappingOutput{}

        mockClient.On("UpdateSchemaMapping", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateSchemaMapping(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
