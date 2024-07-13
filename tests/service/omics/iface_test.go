package omics_test

// tests for the omics service interface for this ../../../service/omics/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/omics"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/omics/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/omics/omics_iface"
	"github.com/stretchr/testify/assert"
)

func TestOmicsServiceCanBeMocked(t *testing.T) {
	var iface omics_iface.IClient
	iface = &omics.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := omics.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAbortMultipartReadSetUpload", func(t *testing.T) {
        input := &omics.AbortMultipartReadSetUploadInput{}
        output := &omics.AbortMultipartReadSetUploadOutput{}

        mockClient.On("AbortMultipartReadSetUpload", ctx, input).Return(output, nil)

        result, err := mockClient.AbortMultipartReadSetUpload(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAcceptShare", func(t *testing.T) {
        input := &omics.AcceptShareInput{}
        output := &omics.AcceptShareOutput{}

        mockClient.On("AcceptShare", ctx, input).Return(output, nil)

        result, err := mockClient.AcceptShare(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchDeleteReadSet", func(t *testing.T) {
        input := &omics.BatchDeleteReadSetInput{}
        output := &omics.BatchDeleteReadSetOutput{}

        mockClient.On("BatchDeleteReadSet", ctx, input).Return(output, nil)

        result, err := mockClient.BatchDeleteReadSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCancelAnnotationImportJob", func(t *testing.T) {
        input := &omics.CancelAnnotationImportJobInput{}
        output := &omics.CancelAnnotationImportJobOutput{}

        mockClient.On("CancelAnnotationImportJob", ctx, input).Return(output, nil)

        result, err := mockClient.CancelAnnotationImportJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCancelRun", func(t *testing.T) {
        input := &omics.CancelRunInput{}
        output := &omics.CancelRunOutput{}

        mockClient.On("CancelRun", ctx, input).Return(output, nil)

        result, err := mockClient.CancelRun(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCancelVariantImportJob", func(t *testing.T) {
        input := &omics.CancelVariantImportJobInput{}
        output := &omics.CancelVariantImportJobOutput{}

        mockClient.On("CancelVariantImportJob", ctx, input).Return(output, nil)

        result, err := mockClient.CancelVariantImportJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCompleteMultipartReadSetUpload", func(t *testing.T) {
        input := &omics.CompleteMultipartReadSetUploadInput{}
        output := &omics.CompleteMultipartReadSetUploadOutput{}

        mockClient.On("CompleteMultipartReadSetUpload", ctx, input).Return(output, nil)

        result, err := mockClient.CompleteMultipartReadSetUpload(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateAnnotationStore", func(t *testing.T) {
        input := &omics.CreateAnnotationStoreInput{}
        output := &omics.CreateAnnotationStoreOutput{}

        mockClient.On("CreateAnnotationStore", ctx, input).Return(output, nil)

        result, err := mockClient.CreateAnnotationStore(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateAnnotationStoreVersion", func(t *testing.T) {
        input := &omics.CreateAnnotationStoreVersionInput{}
        output := &omics.CreateAnnotationStoreVersionOutput{}

        mockClient.On("CreateAnnotationStoreVersion", ctx, input).Return(output, nil)

        result, err := mockClient.CreateAnnotationStoreVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateMultipartReadSetUpload", func(t *testing.T) {
        input := &omics.CreateMultipartReadSetUploadInput{}
        output := &omics.CreateMultipartReadSetUploadOutput{}

        mockClient.On("CreateMultipartReadSetUpload", ctx, input).Return(output, nil)

        result, err := mockClient.CreateMultipartReadSetUpload(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateReferenceStore", func(t *testing.T) {
        input := &omics.CreateReferenceStoreInput{}
        output := &omics.CreateReferenceStoreOutput{}

        mockClient.On("CreateReferenceStore", ctx, input).Return(output, nil)

        result, err := mockClient.CreateReferenceStore(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateRunGroup", func(t *testing.T) {
        input := &omics.CreateRunGroupInput{}
        output := &omics.CreateRunGroupOutput{}

        mockClient.On("CreateRunGroup", ctx, input).Return(output, nil)

        result, err := mockClient.CreateRunGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateSequenceStore", func(t *testing.T) {
        input := &omics.CreateSequenceStoreInput{}
        output := &omics.CreateSequenceStoreOutput{}

        mockClient.On("CreateSequenceStore", ctx, input).Return(output, nil)

        result, err := mockClient.CreateSequenceStore(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateShare", func(t *testing.T) {
        input := &omics.CreateShareInput{}
        output := &omics.CreateShareOutput{}

        mockClient.On("CreateShare", ctx, input).Return(output, nil)

        result, err := mockClient.CreateShare(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateVariantStore", func(t *testing.T) {
        input := &omics.CreateVariantStoreInput{}
        output := &omics.CreateVariantStoreOutput{}

        mockClient.On("CreateVariantStore", ctx, input).Return(output, nil)

        result, err := mockClient.CreateVariantStore(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateWorkflow", func(t *testing.T) {
        input := &omics.CreateWorkflowInput{}
        output := &omics.CreateWorkflowOutput{}

        mockClient.On("CreateWorkflow", ctx, input).Return(output, nil)

        result, err := mockClient.CreateWorkflow(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteAnnotationStore", func(t *testing.T) {
        input := &omics.DeleteAnnotationStoreInput{}
        output := &omics.DeleteAnnotationStoreOutput{}

        mockClient.On("DeleteAnnotationStore", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteAnnotationStore(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteAnnotationStoreVersions", func(t *testing.T) {
        input := &omics.DeleteAnnotationStoreVersionsInput{}
        output := &omics.DeleteAnnotationStoreVersionsOutput{}

        mockClient.On("DeleteAnnotationStoreVersions", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteAnnotationStoreVersions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteReference", func(t *testing.T) {
        input := &omics.DeleteReferenceInput{}
        output := &omics.DeleteReferenceOutput{}

        mockClient.On("DeleteReference", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteReference(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteReferenceStore", func(t *testing.T) {
        input := &omics.DeleteReferenceStoreInput{}
        output := &omics.DeleteReferenceStoreOutput{}

        mockClient.On("DeleteReferenceStore", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteReferenceStore(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteRun", func(t *testing.T) {
        input := &omics.DeleteRunInput{}
        output := &omics.DeleteRunOutput{}

        mockClient.On("DeleteRun", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteRun(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteRunGroup", func(t *testing.T) {
        input := &omics.DeleteRunGroupInput{}
        output := &omics.DeleteRunGroupOutput{}

        mockClient.On("DeleteRunGroup", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteRunGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteSequenceStore", func(t *testing.T) {
        input := &omics.DeleteSequenceStoreInput{}
        output := &omics.DeleteSequenceStoreOutput{}

        mockClient.On("DeleteSequenceStore", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteSequenceStore(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteShare", func(t *testing.T) {
        input := &omics.DeleteShareInput{}
        output := &omics.DeleteShareOutput{}

        mockClient.On("DeleteShare", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteShare(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteVariantStore", func(t *testing.T) {
        input := &omics.DeleteVariantStoreInput{}
        output := &omics.DeleteVariantStoreOutput{}

        mockClient.On("DeleteVariantStore", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteVariantStore(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteWorkflow", func(t *testing.T) {
        input := &omics.DeleteWorkflowInput{}
        output := &omics.DeleteWorkflowOutput{}

        mockClient.On("DeleteWorkflow", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteWorkflow(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetAnnotationImportJob", func(t *testing.T) {
        input := &omics.GetAnnotationImportJobInput{}
        output := &omics.GetAnnotationImportJobOutput{}

        mockClient.On("GetAnnotationImportJob", ctx, input).Return(output, nil)

        result, err := mockClient.GetAnnotationImportJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetAnnotationStore", func(t *testing.T) {
        input := &omics.GetAnnotationStoreInput{}
        output := &omics.GetAnnotationStoreOutput{}

        mockClient.On("GetAnnotationStore", ctx, input).Return(output, nil)

        result, err := mockClient.GetAnnotationStore(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetAnnotationStoreVersion", func(t *testing.T) {
        input := &omics.GetAnnotationStoreVersionInput{}
        output := &omics.GetAnnotationStoreVersionOutput{}

        mockClient.On("GetAnnotationStoreVersion", ctx, input).Return(output, nil)

        result, err := mockClient.GetAnnotationStoreVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetReadSet", func(t *testing.T) {
        input := &omics.GetReadSetInput{}
        output := &omics.GetReadSetOutput{}

        mockClient.On("GetReadSet", ctx, input).Return(output, nil)

        result, err := mockClient.GetReadSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetReadSetActivationJob", func(t *testing.T) {
        input := &omics.GetReadSetActivationJobInput{}
        output := &omics.GetReadSetActivationJobOutput{}

        mockClient.On("GetReadSetActivationJob", ctx, input).Return(output, nil)

        result, err := mockClient.GetReadSetActivationJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetReadSetExportJob", func(t *testing.T) {
        input := &omics.GetReadSetExportJobInput{}
        output := &omics.GetReadSetExportJobOutput{}

        mockClient.On("GetReadSetExportJob", ctx, input).Return(output, nil)

        result, err := mockClient.GetReadSetExportJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetReadSetImportJob", func(t *testing.T) {
        input := &omics.GetReadSetImportJobInput{}
        output := &omics.GetReadSetImportJobOutput{}

        mockClient.On("GetReadSetImportJob", ctx, input).Return(output, nil)

        result, err := mockClient.GetReadSetImportJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetReadSetMetadata", func(t *testing.T) {
        input := &omics.GetReadSetMetadataInput{}
        output := &omics.GetReadSetMetadataOutput{}

        mockClient.On("GetReadSetMetadata", ctx, input).Return(output, nil)

        result, err := mockClient.GetReadSetMetadata(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetReference", func(t *testing.T) {
        input := &omics.GetReferenceInput{}
        output := &omics.GetReferenceOutput{}

        mockClient.On("GetReference", ctx, input).Return(output, nil)

        result, err := mockClient.GetReference(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetReferenceImportJob", func(t *testing.T) {
        input := &omics.GetReferenceImportJobInput{}
        output := &omics.GetReferenceImportJobOutput{}

        mockClient.On("GetReferenceImportJob", ctx, input).Return(output, nil)

        result, err := mockClient.GetReferenceImportJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetReferenceMetadata", func(t *testing.T) {
        input := &omics.GetReferenceMetadataInput{}
        output := &omics.GetReferenceMetadataOutput{}

        mockClient.On("GetReferenceMetadata", ctx, input).Return(output, nil)

        result, err := mockClient.GetReferenceMetadata(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetReferenceStore", func(t *testing.T) {
        input := &omics.GetReferenceStoreInput{}
        output := &omics.GetReferenceStoreOutput{}

        mockClient.On("GetReferenceStore", ctx, input).Return(output, nil)

        result, err := mockClient.GetReferenceStore(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetRun", func(t *testing.T) {
        input := &omics.GetRunInput{}
        output := &omics.GetRunOutput{}

        mockClient.On("GetRun", ctx, input).Return(output, nil)

        result, err := mockClient.GetRun(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetRunGroup", func(t *testing.T) {
        input := &omics.GetRunGroupInput{}
        output := &omics.GetRunGroupOutput{}

        mockClient.On("GetRunGroup", ctx, input).Return(output, nil)

        result, err := mockClient.GetRunGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetRunTask", func(t *testing.T) {
        input := &omics.GetRunTaskInput{}
        output := &omics.GetRunTaskOutput{}

        mockClient.On("GetRunTask", ctx, input).Return(output, nil)

        result, err := mockClient.GetRunTask(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSequenceStore", func(t *testing.T) {
        input := &omics.GetSequenceStoreInput{}
        output := &omics.GetSequenceStoreOutput{}

        mockClient.On("GetSequenceStore", ctx, input).Return(output, nil)

        result, err := mockClient.GetSequenceStore(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetShare", func(t *testing.T) {
        input := &omics.GetShareInput{}
        output := &omics.GetShareOutput{}

        mockClient.On("GetShare", ctx, input).Return(output, nil)

        result, err := mockClient.GetShare(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetVariantImportJob", func(t *testing.T) {
        input := &omics.GetVariantImportJobInput{}
        output := &omics.GetVariantImportJobOutput{}

        mockClient.On("GetVariantImportJob", ctx, input).Return(output, nil)

        result, err := mockClient.GetVariantImportJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetVariantStore", func(t *testing.T) {
        input := &omics.GetVariantStoreInput{}
        output := &omics.GetVariantStoreOutput{}

        mockClient.On("GetVariantStore", ctx, input).Return(output, nil)

        result, err := mockClient.GetVariantStore(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetWorkflow", func(t *testing.T) {
        input := &omics.GetWorkflowInput{}
        output := &omics.GetWorkflowOutput{}

        mockClient.On("GetWorkflow", ctx, input).Return(output, nil)

        result, err := mockClient.GetWorkflow(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAnnotationImportJobs", func(t *testing.T) {
        input := &omics.ListAnnotationImportJobsInput{}
        output := &omics.ListAnnotationImportJobsOutput{}

        mockClient.On("ListAnnotationImportJobs", ctx, input).Return(output, nil)

        result, err := mockClient.ListAnnotationImportJobs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAnnotationStoreVersions", func(t *testing.T) {
        input := &omics.ListAnnotationStoreVersionsInput{}
        output := &omics.ListAnnotationStoreVersionsOutput{}

        mockClient.On("ListAnnotationStoreVersions", ctx, input).Return(output, nil)

        result, err := mockClient.ListAnnotationStoreVersions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAnnotationStores", func(t *testing.T) {
        input := &omics.ListAnnotationStoresInput{}
        output := &omics.ListAnnotationStoresOutput{}

        mockClient.On("ListAnnotationStores", ctx, input).Return(output, nil)

        result, err := mockClient.ListAnnotationStores(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListMultipartReadSetUploads", func(t *testing.T) {
        input := &omics.ListMultipartReadSetUploadsInput{}
        output := &omics.ListMultipartReadSetUploadsOutput{}

        mockClient.On("ListMultipartReadSetUploads", ctx, input).Return(output, nil)

        result, err := mockClient.ListMultipartReadSetUploads(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListReadSetActivationJobs", func(t *testing.T) {
        input := &omics.ListReadSetActivationJobsInput{}
        output := &omics.ListReadSetActivationJobsOutput{}

        mockClient.On("ListReadSetActivationJobs", ctx, input).Return(output, nil)

        result, err := mockClient.ListReadSetActivationJobs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListReadSetExportJobs", func(t *testing.T) {
        input := &omics.ListReadSetExportJobsInput{}
        output := &omics.ListReadSetExportJobsOutput{}

        mockClient.On("ListReadSetExportJobs", ctx, input).Return(output, nil)

        result, err := mockClient.ListReadSetExportJobs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListReadSetImportJobs", func(t *testing.T) {
        input := &omics.ListReadSetImportJobsInput{}
        output := &omics.ListReadSetImportJobsOutput{}

        mockClient.On("ListReadSetImportJobs", ctx, input).Return(output, nil)

        result, err := mockClient.ListReadSetImportJobs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListReadSetUploadParts", func(t *testing.T) {
        input := &omics.ListReadSetUploadPartsInput{}
        output := &omics.ListReadSetUploadPartsOutput{}

        mockClient.On("ListReadSetUploadParts", ctx, input).Return(output, nil)

        result, err := mockClient.ListReadSetUploadParts(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListReadSets", func(t *testing.T) {
        input := &omics.ListReadSetsInput{}
        output := &omics.ListReadSetsOutput{}

        mockClient.On("ListReadSets", ctx, input).Return(output, nil)

        result, err := mockClient.ListReadSets(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListReferenceImportJobs", func(t *testing.T) {
        input := &omics.ListReferenceImportJobsInput{}
        output := &omics.ListReferenceImportJobsOutput{}

        mockClient.On("ListReferenceImportJobs", ctx, input).Return(output, nil)

        result, err := mockClient.ListReferenceImportJobs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListReferenceStores", func(t *testing.T) {
        input := &omics.ListReferenceStoresInput{}
        output := &omics.ListReferenceStoresOutput{}

        mockClient.On("ListReferenceStores", ctx, input).Return(output, nil)

        result, err := mockClient.ListReferenceStores(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListReferences", func(t *testing.T) {
        input := &omics.ListReferencesInput{}
        output := &omics.ListReferencesOutput{}

        mockClient.On("ListReferences", ctx, input).Return(output, nil)

        result, err := mockClient.ListReferences(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListRunGroups", func(t *testing.T) {
        input := &omics.ListRunGroupsInput{}
        output := &omics.ListRunGroupsOutput{}

        mockClient.On("ListRunGroups", ctx, input).Return(output, nil)

        result, err := mockClient.ListRunGroups(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListRunTasks", func(t *testing.T) {
        input := &omics.ListRunTasksInput{}
        output := &omics.ListRunTasksOutput{}

        mockClient.On("ListRunTasks", ctx, input).Return(output, nil)

        result, err := mockClient.ListRunTasks(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListRuns", func(t *testing.T) {
        input := &omics.ListRunsInput{}
        output := &omics.ListRunsOutput{}

        mockClient.On("ListRuns", ctx, input).Return(output, nil)

        result, err := mockClient.ListRuns(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListSequenceStores", func(t *testing.T) {
        input := &omics.ListSequenceStoresInput{}
        output := &omics.ListSequenceStoresOutput{}

        mockClient.On("ListSequenceStores", ctx, input).Return(output, nil)

        result, err := mockClient.ListSequenceStores(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListShares", func(t *testing.T) {
        input := &omics.ListSharesInput{}
        output := &omics.ListSharesOutput{}

        mockClient.On("ListShares", ctx, input).Return(output, nil)

        result, err := mockClient.ListShares(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &omics.ListTagsForResourceInput{}
        output := &omics.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListVariantImportJobs", func(t *testing.T) {
        input := &omics.ListVariantImportJobsInput{}
        output := &omics.ListVariantImportJobsOutput{}

        mockClient.On("ListVariantImportJobs", ctx, input).Return(output, nil)

        result, err := mockClient.ListVariantImportJobs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListVariantStores", func(t *testing.T) {
        input := &omics.ListVariantStoresInput{}
        output := &omics.ListVariantStoresOutput{}

        mockClient.On("ListVariantStores", ctx, input).Return(output, nil)

        result, err := mockClient.ListVariantStores(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListWorkflows", func(t *testing.T) {
        input := &omics.ListWorkflowsInput{}
        output := &omics.ListWorkflowsOutput{}

        mockClient.On("ListWorkflows", ctx, input).Return(output, nil)

        result, err := mockClient.ListWorkflows(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartAnnotationImportJob", func(t *testing.T) {
        input := &omics.StartAnnotationImportJobInput{}
        output := &omics.StartAnnotationImportJobOutput{}

        mockClient.On("StartAnnotationImportJob", ctx, input).Return(output, nil)

        result, err := mockClient.StartAnnotationImportJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartReadSetActivationJob", func(t *testing.T) {
        input := &omics.StartReadSetActivationJobInput{}
        output := &omics.StartReadSetActivationJobOutput{}

        mockClient.On("StartReadSetActivationJob", ctx, input).Return(output, nil)

        result, err := mockClient.StartReadSetActivationJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartReadSetExportJob", func(t *testing.T) {
        input := &omics.StartReadSetExportJobInput{}
        output := &omics.StartReadSetExportJobOutput{}

        mockClient.On("StartReadSetExportJob", ctx, input).Return(output, nil)

        result, err := mockClient.StartReadSetExportJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartReadSetImportJob", func(t *testing.T) {
        input := &omics.StartReadSetImportJobInput{}
        output := &omics.StartReadSetImportJobOutput{}

        mockClient.On("StartReadSetImportJob", ctx, input).Return(output, nil)

        result, err := mockClient.StartReadSetImportJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartReferenceImportJob", func(t *testing.T) {
        input := &omics.StartReferenceImportJobInput{}
        output := &omics.StartReferenceImportJobOutput{}

        mockClient.On("StartReferenceImportJob", ctx, input).Return(output, nil)

        result, err := mockClient.StartReferenceImportJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartRun", func(t *testing.T) {
        input := &omics.StartRunInput{}
        output := &omics.StartRunOutput{}

        mockClient.On("StartRun", ctx, input).Return(output, nil)

        result, err := mockClient.StartRun(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartVariantImportJob", func(t *testing.T) {
        input := &omics.StartVariantImportJobInput{}
        output := &omics.StartVariantImportJobOutput{}

        mockClient.On("StartVariantImportJob", ctx, input).Return(output, nil)

        result, err := mockClient.StartVariantImportJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &omics.TagResourceInput{}
        output := &omics.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &omics.UntagResourceInput{}
        output := &omics.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateAnnotationStore", func(t *testing.T) {
        input := &omics.UpdateAnnotationStoreInput{}
        output := &omics.UpdateAnnotationStoreOutput{}

        mockClient.On("UpdateAnnotationStore", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateAnnotationStore(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateAnnotationStoreVersion", func(t *testing.T) {
        input := &omics.UpdateAnnotationStoreVersionInput{}
        output := &omics.UpdateAnnotationStoreVersionOutput{}

        mockClient.On("UpdateAnnotationStoreVersion", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateAnnotationStoreVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateRunGroup", func(t *testing.T) {
        input := &omics.UpdateRunGroupInput{}
        output := &omics.UpdateRunGroupOutput{}

        mockClient.On("UpdateRunGroup", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateRunGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateVariantStore", func(t *testing.T) {
        input := &omics.UpdateVariantStoreInput{}
        output := &omics.UpdateVariantStoreOutput{}

        mockClient.On("UpdateVariantStore", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateVariantStore(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateWorkflow", func(t *testing.T) {
        input := &omics.UpdateWorkflowInput{}
        output := &omics.UpdateWorkflowOutput{}

        mockClient.On("UpdateWorkflow", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateWorkflow(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUploadReadSetPart", func(t *testing.T) {
        input := &omics.UploadReadSetPartInput{}
        output := &omics.UploadReadSetPartOutput{}

        mockClient.On("UploadReadSetPart", ctx, input).Return(output, nil)

        result, err := mockClient.UploadReadSetPart(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
