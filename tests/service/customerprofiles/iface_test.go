package customerprofiles_test

// tests for the customerprofiles service interface for this ../../../service/customerprofiles/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/customerprofiles"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/customerprofiles/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/customerprofiles/customerprofiles_iface"
	"github.com/stretchr/testify/assert"
)

func TestCustomerprofilesServiceCanBeMocked(t *testing.T) {
	var iface customerprofiles_iface.IClient
	iface = &customerprofiles.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := customerprofiles.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAddProfileKey", func(t *testing.T) {
        input := &customerprofiles.AddProfileKeyInput{}
        output := &customerprofiles.AddProfileKeyOutput{}

        mockClient.On("AddProfileKey", ctx, input).Return(output, nil)

        result, err := mockClient.AddProfileKey(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateCalculatedAttributeDefinition", func(t *testing.T) {
        input := &customerprofiles.CreateCalculatedAttributeDefinitionInput{}
        output := &customerprofiles.CreateCalculatedAttributeDefinitionOutput{}

        mockClient.On("CreateCalculatedAttributeDefinition", ctx, input).Return(output, nil)

        result, err := mockClient.CreateCalculatedAttributeDefinition(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateDomain", func(t *testing.T) {
        input := &customerprofiles.CreateDomainInput{}
        output := &customerprofiles.CreateDomainOutput{}

        mockClient.On("CreateDomain", ctx, input).Return(output, nil)

        result, err := mockClient.CreateDomain(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateEventStream", func(t *testing.T) {
        input := &customerprofiles.CreateEventStreamInput{}
        output := &customerprofiles.CreateEventStreamOutput{}

        mockClient.On("CreateEventStream", ctx, input).Return(output, nil)

        result, err := mockClient.CreateEventStream(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateIntegrationWorkflow", func(t *testing.T) {
        input := &customerprofiles.CreateIntegrationWorkflowInput{}
        output := &customerprofiles.CreateIntegrationWorkflowOutput{}

        mockClient.On("CreateIntegrationWorkflow", ctx, input).Return(output, nil)

        result, err := mockClient.CreateIntegrationWorkflow(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateProfile", func(t *testing.T) {
        input := &customerprofiles.CreateProfileInput{}
        output := &customerprofiles.CreateProfileOutput{}

        mockClient.On("CreateProfile", ctx, input).Return(output, nil)

        result, err := mockClient.CreateProfile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteCalculatedAttributeDefinition", func(t *testing.T) {
        input := &customerprofiles.DeleteCalculatedAttributeDefinitionInput{}
        output := &customerprofiles.DeleteCalculatedAttributeDefinitionOutput{}

        mockClient.On("DeleteCalculatedAttributeDefinition", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteCalculatedAttributeDefinition(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteDomain", func(t *testing.T) {
        input := &customerprofiles.DeleteDomainInput{}
        output := &customerprofiles.DeleteDomainOutput{}

        mockClient.On("DeleteDomain", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteDomain(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteEventStream", func(t *testing.T) {
        input := &customerprofiles.DeleteEventStreamInput{}
        output := &customerprofiles.DeleteEventStreamOutput{}

        mockClient.On("DeleteEventStream", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteEventStream(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteIntegration", func(t *testing.T) {
        input := &customerprofiles.DeleteIntegrationInput{}
        output := &customerprofiles.DeleteIntegrationOutput{}

        mockClient.On("DeleteIntegration", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteIntegration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteProfile", func(t *testing.T) {
        input := &customerprofiles.DeleteProfileInput{}
        output := &customerprofiles.DeleteProfileOutput{}

        mockClient.On("DeleteProfile", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteProfile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteProfileKey", func(t *testing.T) {
        input := &customerprofiles.DeleteProfileKeyInput{}
        output := &customerprofiles.DeleteProfileKeyOutput{}

        mockClient.On("DeleteProfileKey", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteProfileKey(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteProfileObject", func(t *testing.T) {
        input := &customerprofiles.DeleteProfileObjectInput{}
        output := &customerprofiles.DeleteProfileObjectOutput{}

        mockClient.On("DeleteProfileObject", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteProfileObject(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteProfileObjectType", func(t *testing.T) {
        input := &customerprofiles.DeleteProfileObjectTypeInput{}
        output := &customerprofiles.DeleteProfileObjectTypeOutput{}

        mockClient.On("DeleteProfileObjectType", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteProfileObjectType(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteWorkflow", func(t *testing.T) {
        input := &customerprofiles.DeleteWorkflowInput{}
        output := &customerprofiles.DeleteWorkflowOutput{}

        mockClient.On("DeleteWorkflow", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteWorkflow(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDetectProfileObjectType", func(t *testing.T) {
        input := &customerprofiles.DetectProfileObjectTypeInput{}
        output := &customerprofiles.DetectProfileObjectTypeOutput{}

        mockClient.On("DetectProfileObjectType", ctx, input).Return(output, nil)

        result, err := mockClient.DetectProfileObjectType(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetAutoMergingPreview", func(t *testing.T) {
        input := &customerprofiles.GetAutoMergingPreviewInput{}
        output := &customerprofiles.GetAutoMergingPreviewOutput{}

        mockClient.On("GetAutoMergingPreview", ctx, input).Return(output, nil)

        result, err := mockClient.GetAutoMergingPreview(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetCalculatedAttributeDefinition", func(t *testing.T) {
        input := &customerprofiles.GetCalculatedAttributeDefinitionInput{}
        output := &customerprofiles.GetCalculatedAttributeDefinitionOutput{}

        mockClient.On("GetCalculatedAttributeDefinition", ctx, input).Return(output, nil)

        result, err := mockClient.GetCalculatedAttributeDefinition(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetCalculatedAttributeForProfile", func(t *testing.T) {
        input := &customerprofiles.GetCalculatedAttributeForProfileInput{}
        output := &customerprofiles.GetCalculatedAttributeForProfileOutput{}

        mockClient.On("GetCalculatedAttributeForProfile", ctx, input).Return(output, nil)

        result, err := mockClient.GetCalculatedAttributeForProfile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDomain", func(t *testing.T) {
        input := &customerprofiles.GetDomainInput{}
        output := &customerprofiles.GetDomainOutput{}

        mockClient.On("GetDomain", ctx, input).Return(output, nil)

        result, err := mockClient.GetDomain(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetEventStream", func(t *testing.T) {
        input := &customerprofiles.GetEventStreamInput{}
        output := &customerprofiles.GetEventStreamOutput{}

        mockClient.On("GetEventStream", ctx, input).Return(output, nil)

        result, err := mockClient.GetEventStream(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetIdentityResolutionJob", func(t *testing.T) {
        input := &customerprofiles.GetIdentityResolutionJobInput{}
        output := &customerprofiles.GetIdentityResolutionJobOutput{}

        mockClient.On("GetIdentityResolutionJob", ctx, input).Return(output, nil)

        result, err := mockClient.GetIdentityResolutionJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetIntegration", func(t *testing.T) {
        input := &customerprofiles.GetIntegrationInput{}
        output := &customerprofiles.GetIntegrationOutput{}

        mockClient.On("GetIntegration", ctx, input).Return(output, nil)

        result, err := mockClient.GetIntegration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetMatches", func(t *testing.T) {
        input := &customerprofiles.GetMatchesInput{}
        output := &customerprofiles.GetMatchesOutput{}

        mockClient.On("GetMatches", ctx, input).Return(output, nil)

        result, err := mockClient.GetMatches(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetProfileObjectType", func(t *testing.T) {
        input := &customerprofiles.GetProfileObjectTypeInput{}
        output := &customerprofiles.GetProfileObjectTypeOutput{}

        mockClient.On("GetProfileObjectType", ctx, input).Return(output, nil)

        result, err := mockClient.GetProfileObjectType(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetProfileObjectTypeTemplate", func(t *testing.T) {
        input := &customerprofiles.GetProfileObjectTypeTemplateInput{}
        output := &customerprofiles.GetProfileObjectTypeTemplateOutput{}

        mockClient.On("GetProfileObjectTypeTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.GetProfileObjectTypeTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSimilarProfiles", func(t *testing.T) {
        input := &customerprofiles.GetSimilarProfilesInput{}
        output := &customerprofiles.GetSimilarProfilesOutput{}

        mockClient.On("GetSimilarProfiles", ctx, input).Return(output, nil)

        result, err := mockClient.GetSimilarProfiles(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetWorkflow", func(t *testing.T) {
        input := &customerprofiles.GetWorkflowInput{}
        output := &customerprofiles.GetWorkflowOutput{}

        mockClient.On("GetWorkflow", ctx, input).Return(output, nil)

        result, err := mockClient.GetWorkflow(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetWorkflowSteps", func(t *testing.T) {
        input := &customerprofiles.GetWorkflowStepsInput{}
        output := &customerprofiles.GetWorkflowStepsOutput{}

        mockClient.On("GetWorkflowSteps", ctx, input).Return(output, nil)

        result, err := mockClient.GetWorkflowSteps(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAccountIntegrations", func(t *testing.T) {
        input := &customerprofiles.ListAccountIntegrationsInput{}
        output := &customerprofiles.ListAccountIntegrationsOutput{}

        mockClient.On("ListAccountIntegrations", ctx, input).Return(output, nil)

        result, err := mockClient.ListAccountIntegrations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListCalculatedAttributeDefinitions", func(t *testing.T) {
        input := &customerprofiles.ListCalculatedAttributeDefinitionsInput{}
        output := &customerprofiles.ListCalculatedAttributeDefinitionsOutput{}

        mockClient.On("ListCalculatedAttributeDefinitions", ctx, input).Return(output, nil)

        result, err := mockClient.ListCalculatedAttributeDefinitions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListCalculatedAttributesForProfile", func(t *testing.T) {
        input := &customerprofiles.ListCalculatedAttributesForProfileInput{}
        output := &customerprofiles.ListCalculatedAttributesForProfileOutput{}

        mockClient.On("ListCalculatedAttributesForProfile", ctx, input).Return(output, nil)

        result, err := mockClient.ListCalculatedAttributesForProfile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDomains", func(t *testing.T) {
        input := &customerprofiles.ListDomainsInput{}
        output := &customerprofiles.ListDomainsOutput{}

        mockClient.On("ListDomains", ctx, input).Return(output, nil)

        result, err := mockClient.ListDomains(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListEventStreams", func(t *testing.T) {
        input := &customerprofiles.ListEventStreamsInput{}
        output := &customerprofiles.ListEventStreamsOutput{}

        mockClient.On("ListEventStreams", ctx, input).Return(output, nil)

        result, err := mockClient.ListEventStreams(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListIdentityResolutionJobs", func(t *testing.T) {
        input := &customerprofiles.ListIdentityResolutionJobsInput{}
        output := &customerprofiles.ListIdentityResolutionJobsOutput{}

        mockClient.On("ListIdentityResolutionJobs", ctx, input).Return(output, nil)

        result, err := mockClient.ListIdentityResolutionJobs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListIntegrations", func(t *testing.T) {
        input := &customerprofiles.ListIntegrationsInput{}
        output := &customerprofiles.ListIntegrationsOutput{}

        mockClient.On("ListIntegrations", ctx, input).Return(output, nil)

        result, err := mockClient.ListIntegrations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListProfileObjectTypeTemplates", func(t *testing.T) {
        input := &customerprofiles.ListProfileObjectTypeTemplatesInput{}
        output := &customerprofiles.ListProfileObjectTypeTemplatesOutput{}

        mockClient.On("ListProfileObjectTypeTemplates", ctx, input).Return(output, nil)

        result, err := mockClient.ListProfileObjectTypeTemplates(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListProfileObjectTypes", func(t *testing.T) {
        input := &customerprofiles.ListProfileObjectTypesInput{}
        output := &customerprofiles.ListProfileObjectTypesOutput{}

        mockClient.On("ListProfileObjectTypes", ctx, input).Return(output, nil)

        result, err := mockClient.ListProfileObjectTypes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListProfileObjects", func(t *testing.T) {
        input := &customerprofiles.ListProfileObjectsInput{}
        output := &customerprofiles.ListProfileObjectsOutput{}

        mockClient.On("ListProfileObjects", ctx, input).Return(output, nil)

        result, err := mockClient.ListProfileObjects(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListRuleBasedMatches", func(t *testing.T) {
        input := &customerprofiles.ListRuleBasedMatchesInput{}
        output := &customerprofiles.ListRuleBasedMatchesOutput{}

        mockClient.On("ListRuleBasedMatches", ctx, input).Return(output, nil)

        result, err := mockClient.ListRuleBasedMatches(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &customerprofiles.ListTagsForResourceInput{}
        output := &customerprofiles.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListWorkflows", func(t *testing.T) {
        input := &customerprofiles.ListWorkflowsInput{}
        output := &customerprofiles.ListWorkflowsOutput{}

        mockClient.On("ListWorkflows", ctx, input).Return(output, nil)

        result, err := mockClient.ListWorkflows(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestMergeProfiles", func(t *testing.T) {
        input := &customerprofiles.MergeProfilesInput{}
        output := &customerprofiles.MergeProfilesOutput{}

        mockClient.On("MergeProfiles", ctx, input).Return(output, nil)

        result, err := mockClient.MergeProfiles(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutIntegration", func(t *testing.T) {
        input := &customerprofiles.PutIntegrationInput{}
        output := &customerprofiles.PutIntegrationOutput{}

        mockClient.On("PutIntegration", ctx, input).Return(output, nil)

        result, err := mockClient.PutIntegration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutProfileObject", func(t *testing.T) {
        input := &customerprofiles.PutProfileObjectInput{}
        output := &customerprofiles.PutProfileObjectOutput{}

        mockClient.On("PutProfileObject", ctx, input).Return(output, nil)

        result, err := mockClient.PutProfileObject(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutProfileObjectType", func(t *testing.T) {
        input := &customerprofiles.PutProfileObjectTypeInput{}
        output := &customerprofiles.PutProfileObjectTypeOutput{}

        mockClient.On("PutProfileObjectType", ctx, input).Return(output, nil)

        result, err := mockClient.PutProfileObjectType(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSearchProfiles", func(t *testing.T) {
        input := &customerprofiles.SearchProfilesInput{}
        output := &customerprofiles.SearchProfilesOutput{}

        mockClient.On("SearchProfiles", ctx, input).Return(output, nil)

        result, err := mockClient.SearchProfiles(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &customerprofiles.TagResourceInput{}
        output := &customerprofiles.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &customerprofiles.UntagResourceInput{}
        output := &customerprofiles.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateCalculatedAttributeDefinition", func(t *testing.T) {
        input := &customerprofiles.UpdateCalculatedAttributeDefinitionInput{}
        output := &customerprofiles.UpdateCalculatedAttributeDefinitionOutput{}

        mockClient.On("UpdateCalculatedAttributeDefinition", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateCalculatedAttributeDefinition(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateDomain", func(t *testing.T) {
        input := &customerprofiles.UpdateDomainInput{}
        output := &customerprofiles.UpdateDomainOutput{}

        mockClient.On("UpdateDomain", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateDomain(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateProfile", func(t *testing.T) {
        input := &customerprofiles.UpdateProfileInput{}
        output := &customerprofiles.UpdateProfileOutput{}

        mockClient.On("UpdateProfile", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateProfile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
