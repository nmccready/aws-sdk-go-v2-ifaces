package kendra_test

// tests for the kendra service interface for this ../../../service/kendra/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/kendra"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/kendra/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/kendra/kendra_iface"
	"github.com/stretchr/testify/assert"
)

func TestKendraServiceCanBeMocked(t *testing.T) {
	var iface kendra_iface.IClient
	iface = &kendra.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := kendra.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateEntitiesToExperience", func(t *testing.T) {
        input := &kendra.AssociateEntitiesToExperienceInput{}
        output := &kendra.AssociateEntitiesToExperienceOutput{}

        mockClient.On("AssociateEntitiesToExperience", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateEntitiesToExperience(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociatePersonasToEntities", func(t *testing.T) {
        input := &kendra.AssociatePersonasToEntitiesInput{}
        output := &kendra.AssociatePersonasToEntitiesOutput{}

        mockClient.On("AssociatePersonasToEntities", ctx, input).Return(output, nil)

        result, err := mockClient.AssociatePersonasToEntities(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchDeleteDocument", func(t *testing.T) {
        input := &kendra.BatchDeleteDocumentInput{}
        output := &kendra.BatchDeleteDocumentOutput{}

        mockClient.On("BatchDeleteDocument", ctx, input).Return(output, nil)

        result, err := mockClient.BatchDeleteDocument(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchDeleteFeaturedResultsSet", func(t *testing.T) {
        input := &kendra.BatchDeleteFeaturedResultsSetInput{}
        output := &kendra.BatchDeleteFeaturedResultsSetOutput{}

        mockClient.On("BatchDeleteFeaturedResultsSet", ctx, input).Return(output, nil)

        result, err := mockClient.BatchDeleteFeaturedResultsSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchGetDocumentStatus", func(t *testing.T) {
        input := &kendra.BatchGetDocumentStatusInput{}
        output := &kendra.BatchGetDocumentStatusOutput{}

        mockClient.On("BatchGetDocumentStatus", ctx, input).Return(output, nil)

        result, err := mockClient.BatchGetDocumentStatus(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchPutDocument", func(t *testing.T) {
        input := &kendra.BatchPutDocumentInput{}
        output := &kendra.BatchPutDocumentOutput{}

        mockClient.On("BatchPutDocument", ctx, input).Return(output, nil)

        result, err := mockClient.BatchPutDocument(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestClearQuerySuggestions", func(t *testing.T) {
        input := &kendra.ClearQuerySuggestionsInput{}
        output := &kendra.ClearQuerySuggestionsOutput{}

        mockClient.On("ClearQuerySuggestions", ctx, input).Return(output, nil)

        result, err := mockClient.ClearQuerySuggestions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateAccessControlConfiguration", func(t *testing.T) {
        input := &kendra.CreateAccessControlConfigurationInput{}
        output := &kendra.CreateAccessControlConfigurationOutput{}

        mockClient.On("CreateAccessControlConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.CreateAccessControlConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateDataSource", func(t *testing.T) {
        input := &kendra.CreateDataSourceInput{}
        output := &kendra.CreateDataSourceOutput{}

        mockClient.On("CreateDataSource", ctx, input).Return(output, nil)

        result, err := mockClient.CreateDataSource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateExperience", func(t *testing.T) {
        input := &kendra.CreateExperienceInput{}
        output := &kendra.CreateExperienceOutput{}

        mockClient.On("CreateExperience", ctx, input).Return(output, nil)

        result, err := mockClient.CreateExperience(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateFaq", func(t *testing.T) {
        input := &kendra.CreateFaqInput{}
        output := &kendra.CreateFaqOutput{}

        mockClient.On("CreateFaq", ctx, input).Return(output, nil)

        result, err := mockClient.CreateFaq(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateFeaturedResultsSet", func(t *testing.T) {
        input := &kendra.CreateFeaturedResultsSetInput{}
        output := &kendra.CreateFeaturedResultsSetOutput{}

        mockClient.On("CreateFeaturedResultsSet", ctx, input).Return(output, nil)

        result, err := mockClient.CreateFeaturedResultsSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateIndex", func(t *testing.T) {
        input := &kendra.CreateIndexInput{}
        output := &kendra.CreateIndexOutput{}

        mockClient.On("CreateIndex", ctx, input).Return(output, nil)

        result, err := mockClient.CreateIndex(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateQuerySuggestionsBlockList", func(t *testing.T) {
        input := &kendra.CreateQuerySuggestionsBlockListInput{}
        output := &kendra.CreateQuerySuggestionsBlockListOutput{}

        mockClient.On("CreateQuerySuggestionsBlockList", ctx, input).Return(output, nil)

        result, err := mockClient.CreateQuerySuggestionsBlockList(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateThesaurus", func(t *testing.T) {
        input := &kendra.CreateThesaurusInput{}
        output := &kendra.CreateThesaurusOutput{}

        mockClient.On("CreateThesaurus", ctx, input).Return(output, nil)

        result, err := mockClient.CreateThesaurus(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteAccessControlConfiguration", func(t *testing.T) {
        input := &kendra.DeleteAccessControlConfigurationInput{}
        output := &kendra.DeleteAccessControlConfigurationOutput{}

        mockClient.On("DeleteAccessControlConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteAccessControlConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteDataSource", func(t *testing.T) {
        input := &kendra.DeleteDataSourceInput{}
        output := &kendra.DeleteDataSourceOutput{}

        mockClient.On("DeleteDataSource", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteDataSource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteExperience", func(t *testing.T) {
        input := &kendra.DeleteExperienceInput{}
        output := &kendra.DeleteExperienceOutput{}

        mockClient.On("DeleteExperience", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteExperience(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteFaq", func(t *testing.T) {
        input := &kendra.DeleteFaqInput{}
        output := &kendra.DeleteFaqOutput{}

        mockClient.On("DeleteFaq", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteFaq(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteIndex", func(t *testing.T) {
        input := &kendra.DeleteIndexInput{}
        output := &kendra.DeleteIndexOutput{}

        mockClient.On("DeleteIndex", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteIndex(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeletePrincipalMapping", func(t *testing.T) {
        input := &kendra.DeletePrincipalMappingInput{}
        output := &kendra.DeletePrincipalMappingOutput{}

        mockClient.On("DeletePrincipalMapping", ctx, input).Return(output, nil)

        result, err := mockClient.DeletePrincipalMapping(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteQuerySuggestionsBlockList", func(t *testing.T) {
        input := &kendra.DeleteQuerySuggestionsBlockListInput{}
        output := &kendra.DeleteQuerySuggestionsBlockListOutput{}

        mockClient.On("DeleteQuerySuggestionsBlockList", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteQuerySuggestionsBlockList(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteThesaurus", func(t *testing.T) {
        input := &kendra.DeleteThesaurusInput{}
        output := &kendra.DeleteThesaurusOutput{}

        mockClient.On("DeleteThesaurus", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteThesaurus(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeAccessControlConfiguration", func(t *testing.T) {
        input := &kendra.DescribeAccessControlConfigurationInput{}
        output := &kendra.DescribeAccessControlConfigurationOutput{}

        mockClient.On("DescribeAccessControlConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeAccessControlConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeDataSource", func(t *testing.T) {
        input := &kendra.DescribeDataSourceInput{}
        output := &kendra.DescribeDataSourceOutput{}

        mockClient.On("DescribeDataSource", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeDataSource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeExperience", func(t *testing.T) {
        input := &kendra.DescribeExperienceInput{}
        output := &kendra.DescribeExperienceOutput{}

        mockClient.On("DescribeExperience", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeExperience(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeFaq", func(t *testing.T) {
        input := &kendra.DescribeFaqInput{}
        output := &kendra.DescribeFaqOutput{}

        mockClient.On("DescribeFaq", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeFaq(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeFeaturedResultsSet", func(t *testing.T) {
        input := &kendra.DescribeFeaturedResultsSetInput{}
        output := &kendra.DescribeFeaturedResultsSetOutput{}

        mockClient.On("DescribeFeaturedResultsSet", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeFeaturedResultsSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeIndex", func(t *testing.T) {
        input := &kendra.DescribeIndexInput{}
        output := &kendra.DescribeIndexOutput{}

        mockClient.On("DescribeIndex", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeIndex(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribePrincipalMapping", func(t *testing.T) {
        input := &kendra.DescribePrincipalMappingInput{}
        output := &kendra.DescribePrincipalMappingOutput{}

        mockClient.On("DescribePrincipalMapping", ctx, input).Return(output, nil)

        result, err := mockClient.DescribePrincipalMapping(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeQuerySuggestionsBlockList", func(t *testing.T) {
        input := &kendra.DescribeQuerySuggestionsBlockListInput{}
        output := &kendra.DescribeQuerySuggestionsBlockListOutput{}

        mockClient.On("DescribeQuerySuggestionsBlockList", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeQuerySuggestionsBlockList(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeQuerySuggestionsConfig", func(t *testing.T) {
        input := &kendra.DescribeQuerySuggestionsConfigInput{}
        output := &kendra.DescribeQuerySuggestionsConfigOutput{}

        mockClient.On("DescribeQuerySuggestionsConfig", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeQuerySuggestionsConfig(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeThesaurus", func(t *testing.T) {
        input := &kendra.DescribeThesaurusInput{}
        output := &kendra.DescribeThesaurusOutput{}

        mockClient.On("DescribeThesaurus", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeThesaurus(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateEntitiesFromExperience", func(t *testing.T) {
        input := &kendra.DisassociateEntitiesFromExperienceInput{}
        output := &kendra.DisassociateEntitiesFromExperienceOutput{}

        mockClient.On("DisassociateEntitiesFromExperience", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateEntitiesFromExperience(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociatePersonasFromEntities", func(t *testing.T) {
        input := &kendra.DisassociatePersonasFromEntitiesInput{}
        output := &kendra.DisassociatePersonasFromEntitiesOutput{}

        mockClient.On("DisassociatePersonasFromEntities", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociatePersonasFromEntities(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetQuerySuggestions", func(t *testing.T) {
        input := &kendra.GetQuerySuggestionsInput{}
        output := &kendra.GetQuerySuggestionsOutput{}

        mockClient.On("GetQuerySuggestions", ctx, input).Return(output, nil)

        result, err := mockClient.GetQuerySuggestions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSnapshots", func(t *testing.T) {
        input := &kendra.GetSnapshotsInput{}
        output := &kendra.GetSnapshotsOutput{}

        mockClient.On("GetSnapshots", ctx, input).Return(output, nil)

        result, err := mockClient.GetSnapshots(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAccessControlConfigurations", func(t *testing.T) {
        input := &kendra.ListAccessControlConfigurationsInput{}
        output := &kendra.ListAccessControlConfigurationsOutput{}

        mockClient.On("ListAccessControlConfigurations", ctx, input).Return(output, nil)

        result, err := mockClient.ListAccessControlConfigurations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDataSourceSyncJobs", func(t *testing.T) {
        input := &kendra.ListDataSourceSyncJobsInput{}
        output := &kendra.ListDataSourceSyncJobsOutput{}

        mockClient.On("ListDataSourceSyncJobs", ctx, input).Return(output, nil)

        result, err := mockClient.ListDataSourceSyncJobs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDataSources", func(t *testing.T) {
        input := &kendra.ListDataSourcesInput{}
        output := &kendra.ListDataSourcesOutput{}

        mockClient.On("ListDataSources", ctx, input).Return(output, nil)

        result, err := mockClient.ListDataSources(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListEntityPersonas", func(t *testing.T) {
        input := &kendra.ListEntityPersonasInput{}
        output := &kendra.ListEntityPersonasOutput{}

        mockClient.On("ListEntityPersonas", ctx, input).Return(output, nil)

        result, err := mockClient.ListEntityPersonas(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListExperienceEntities", func(t *testing.T) {
        input := &kendra.ListExperienceEntitiesInput{}
        output := &kendra.ListExperienceEntitiesOutput{}

        mockClient.On("ListExperienceEntities", ctx, input).Return(output, nil)

        result, err := mockClient.ListExperienceEntities(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListExperiences", func(t *testing.T) {
        input := &kendra.ListExperiencesInput{}
        output := &kendra.ListExperiencesOutput{}

        mockClient.On("ListExperiences", ctx, input).Return(output, nil)

        result, err := mockClient.ListExperiences(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListFaqs", func(t *testing.T) {
        input := &kendra.ListFaqsInput{}
        output := &kendra.ListFaqsOutput{}

        mockClient.On("ListFaqs", ctx, input).Return(output, nil)

        result, err := mockClient.ListFaqs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListFeaturedResultsSets", func(t *testing.T) {
        input := &kendra.ListFeaturedResultsSetsInput{}
        output := &kendra.ListFeaturedResultsSetsOutput{}

        mockClient.On("ListFeaturedResultsSets", ctx, input).Return(output, nil)

        result, err := mockClient.ListFeaturedResultsSets(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListGroupsOlderThanOrderingId", func(t *testing.T) {
        input := &kendra.ListGroupsOlderThanOrderingIdInput{}
        output := &kendra.ListGroupsOlderThanOrderingIdOutput{}

        mockClient.On("ListGroupsOlderThanOrderingId", ctx, input).Return(output, nil)

        result, err := mockClient.ListGroupsOlderThanOrderingId(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListIndices", func(t *testing.T) {
        input := &kendra.ListIndicesInput{}
        output := &kendra.ListIndicesOutput{}

        mockClient.On("ListIndices", ctx, input).Return(output, nil)

        result, err := mockClient.ListIndices(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListQuerySuggestionsBlockLists", func(t *testing.T) {
        input := &kendra.ListQuerySuggestionsBlockListsInput{}
        output := &kendra.ListQuerySuggestionsBlockListsOutput{}

        mockClient.On("ListQuerySuggestionsBlockLists", ctx, input).Return(output, nil)

        result, err := mockClient.ListQuerySuggestionsBlockLists(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &kendra.ListTagsForResourceInput{}
        output := &kendra.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListThesauri", func(t *testing.T) {
        input := &kendra.ListThesauriInput{}
        output := &kendra.ListThesauriOutput{}

        mockClient.On("ListThesauri", ctx, input).Return(output, nil)

        result, err := mockClient.ListThesauri(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutPrincipalMapping", func(t *testing.T) {
        input := &kendra.PutPrincipalMappingInput{}
        output := &kendra.PutPrincipalMappingOutput{}

        mockClient.On("PutPrincipalMapping", ctx, input).Return(output, nil)

        result, err := mockClient.PutPrincipalMapping(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestQuery", func(t *testing.T) {
        input := &kendra.QueryInput{}
        output := &kendra.QueryOutput{}

        mockClient.On("Query", ctx, input).Return(output, nil)

        result, err := mockClient.Query(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRetrieve", func(t *testing.T) {
        input := &kendra.RetrieveInput{}
        output := &kendra.RetrieveOutput{}

        mockClient.On("Retrieve", ctx, input).Return(output, nil)

        result, err := mockClient.Retrieve(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartDataSourceSyncJob", func(t *testing.T) {
        input := &kendra.StartDataSourceSyncJobInput{}
        output := &kendra.StartDataSourceSyncJobOutput{}

        mockClient.On("StartDataSourceSyncJob", ctx, input).Return(output, nil)

        result, err := mockClient.StartDataSourceSyncJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopDataSourceSyncJob", func(t *testing.T) {
        input := &kendra.StopDataSourceSyncJobInput{}
        output := &kendra.StopDataSourceSyncJobOutput{}

        mockClient.On("StopDataSourceSyncJob", ctx, input).Return(output, nil)

        result, err := mockClient.StopDataSourceSyncJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSubmitFeedback", func(t *testing.T) {
        input := &kendra.SubmitFeedbackInput{}
        output := &kendra.SubmitFeedbackOutput{}

        mockClient.On("SubmitFeedback", ctx, input).Return(output, nil)

        result, err := mockClient.SubmitFeedback(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &kendra.TagResourceInput{}
        output := &kendra.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &kendra.UntagResourceInput{}
        output := &kendra.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateAccessControlConfiguration", func(t *testing.T) {
        input := &kendra.UpdateAccessControlConfigurationInput{}
        output := &kendra.UpdateAccessControlConfigurationOutput{}

        mockClient.On("UpdateAccessControlConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateAccessControlConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateDataSource", func(t *testing.T) {
        input := &kendra.UpdateDataSourceInput{}
        output := &kendra.UpdateDataSourceOutput{}

        mockClient.On("UpdateDataSource", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateDataSource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateExperience", func(t *testing.T) {
        input := &kendra.UpdateExperienceInput{}
        output := &kendra.UpdateExperienceOutput{}

        mockClient.On("UpdateExperience", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateExperience(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateFeaturedResultsSet", func(t *testing.T) {
        input := &kendra.UpdateFeaturedResultsSetInput{}
        output := &kendra.UpdateFeaturedResultsSetOutput{}

        mockClient.On("UpdateFeaturedResultsSet", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateFeaturedResultsSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateIndex", func(t *testing.T) {
        input := &kendra.UpdateIndexInput{}
        output := &kendra.UpdateIndexOutput{}

        mockClient.On("UpdateIndex", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateIndex(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateQuerySuggestionsBlockList", func(t *testing.T) {
        input := &kendra.UpdateQuerySuggestionsBlockListInput{}
        output := &kendra.UpdateQuerySuggestionsBlockListOutput{}

        mockClient.On("UpdateQuerySuggestionsBlockList", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateQuerySuggestionsBlockList(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateQuerySuggestionsConfig", func(t *testing.T) {
        input := &kendra.UpdateQuerySuggestionsConfigInput{}
        output := &kendra.UpdateQuerySuggestionsConfigOutput{}

        mockClient.On("UpdateQuerySuggestionsConfig", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateQuerySuggestionsConfig(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateThesaurus", func(t *testing.T) {
        input := &kendra.UpdateThesaurusInput{}
        output := &kendra.UpdateThesaurusOutput{}

        mockClient.On("UpdateThesaurus", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateThesaurus(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
