package datazone_test

// tests for the datazone service interface for this ../../../service/datazone/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/datazone"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/datazone/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/datazone/datazone_iface"
	"github.com/stretchr/testify/assert"
)

func TestDatazoneServiceCanBeMocked(t *testing.T) {
	var iface datazone_iface.IClient
	iface = &datazone.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := datazone.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAcceptPredictions", func(t *testing.T) {
        input := &datazone.AcceptPredictionsInput{}
        output := &datazone.AcceptPredictionsOutput{}

        mockClient.On("AcceptPredictions", ctx, input).Return(output, nil)

        result, err := mockClient.AcceptPredictions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAcceptSubscriptionRequest", func(t *testing.T) {
        input := &datazone.AcceptSubscriptionRequestInput{}
        output := &datazone.AcceptSubscriptionRequestOutput{}

        mockClient.On("AcceptSubscriptionRequest", ctx, input).Return(output, nil)

        result, err := mockClient.AcceptSubscriptionRequest(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateEnvironmentRole", func(t *testing.T) {
        input := &datazone.AssociateEnvironmentRoleInput{}
        output := &datazone.AssociateEnvironmentRoleOutput{}

        mockClient.On("AssociateEnvironmentRole", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateEnvironmentRole(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCancelMetadataGenerationRun", func(t *testing.T) {
        input := &datazone.CancelMetadataGenerationRunInput{}
        output := &datazone.CancelMetadataGenerationRunOutput{}

        mockClient.On("CancelMetadataGenerationRun", ctx, input).Return(output, nil)

        result, err := mockClient.CancelMetadataGenerationRun(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCancelSubscription", func(t *testing.T) {
        input := &datazone.CancelSubscriptionInput{}
        output := &datazone.CancelSubscriptionOutput{}

        mockClient.On("CancelSubscription", ctx, input).Return(output, nil)

        result, err := mockClient.CancelSubscription(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateAsset", func(t *testing.T) {
        input := &datazone.CreateAssetInput{}
        output := &datazone.CreateAssetOutput{}

        mockClient.On("CreateAsset", ctx, input).Return(output, nil)

        result, err := mockClient.CreateAsset(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateAssetRevision", func(t *testing.T) {
        input := &datazone.CreateAssetRevisionInput{}
        output := &datazone.CreateAssetRevisionOutput{}

        mockClient.On("CreateAssetRevision", ctx, input).Return(output, nil)

        result, err := mockClient.CreateAssetRevision(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateAssetType", func(t *testing.T) {
        input := &datazone.CreateAssetTypeInput{}
        output := &datazone.CreateAssetTypeOutput{}

        mockClient.On("CreateAssetType", ctx, input).Return(output, nil)

        result, err := mockClient.CreateAssetType(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateDataSource", func(t *testing.T) {
        input := &datazone.CreateDataSourceInput{}
        output := &datazone.CreateDataSourceOutput{}

        mockClient.On("CreateDataSource", ctx, input).Return(output, nil)

        result, err := mockClient.CreateDataSource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateDomain", func(t *testing.T) {
        input := &datazone.CreateDomainInput{}
        output := &datazone.CreateDomainOutput{}

        mockClient.On("CreateDomain", ctx, input).Return(output, nil)

        result, err := mockClient.CreateDomain(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateEnvironment", func(t *testing.T) {
        input := &datazone.CreateEnvironmentInput{}
        output := &datazone.CreateEnvironmentOutput{}

        mockClient.On("CreateEnvironment", ctx, input).Return(output, nil)

        result, err := mockClient.CreateEnvironment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateEnvironmentAction", func(t *testing.T) {
        input := &datazone.CreateEnvironmentActionInput{}
        output := &datazone.CreateEnvironmentActionOutput{}

        mockClient.On("CreateEnvironmentAction", ctx, input).Return(output, nil)

        result, err := mockClient.CreateEnvironmentAction(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateEnvironmentProfile", func(t *testing.T) {
        input := &datazone.CreateEnvironmentProfileInput{}
        output := &datazone.CreateEnvironmentProfileOutput{}

        mockClient.On("CreateEnvironmentProfile", ctx, input).Return(output, nil)

        result, err := mockClient.CreateEnvironmentProfile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateFormType", func(t *testing.T) {
        input := &datazone.CreateFormTypeInput{}
        output := &datazone.CreateFormTypeOutput{}

        mockClient.On("CreateFormType", ctx, input).Return(output, nil)

        result, err := mockClient.CreateFormType(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateGlossary", func(t *testing.T) {
        input := &datazone.CreateGlossaryInput{}
        output := &datazone.CreateGlossaryOutput{}

        mockClient.On("CreateGlossary", ctx, input).Return(output, nil)

        result, err := mockClient.CreateGlossary(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateGlossaryTerm", func(t *testing.T) {
        input := &datazone.CreateGlossaryTermInput{}
        output := &datazone.CreateGlossaryTermOutput{}

        mockClient.On("CreateGlossaryTerm", ctx, input).Return(output, nil)

        result, err := mockClient.CreateGlossaryTerm(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateGroupProfile", func(t *testing.T) {
        input := &datazone.CreateGroupProfileInput{}
        output := &datazone.CreateGroupProfileOutput{}

        mockClient.On("CreateGroupProfile", ctx, input).Return(output, nil)

        result, err := mockClient.CreateGroupProfile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateListingChangeSet", func(t *testing.T) {
        input := &datazone.CreateListingChangeSetInput{}
        output := &datazone.CreateListingChangeSetOutput{}

        mockClient.On("CreateListingChangeSet", ctx, input).Return(output, nil)

        result, err := mockClient.CreateListingChangeSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateProject", func(t *testing.T) {
        input := &datazone.CreateProjectInput{}
        output := &datazone.CreateProjectOutput{}

        mockClient.On("CreateProject", ctx, input).Return(output, nil)

        result, err := mockClient.CreateProject(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateProjectMembership", func(t *testing.T) {
        input := &datazone.CreateProjectMembershipInput{}
        output := &datazone.CreateProjectMembershipOutput{}

        mockClient.On("CreateProjectMembership", ctx, input).Return(output, nil)

        result, err := mockClient.CreateProjectMembership(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateSubscriptionGrant", func(t *testing.T) {
        input := &datazone.CreateSubscriptionGrantInput{}
        output := &datazone.CreateSubscriptionGrantOutput{}

        mockClient.On("CreateSubscriptionGrant", ctx, input).Return(output, nil)

        result, err := mockClient.CreateSubscriptionGrant(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateSubscriptionRequest", func(t *testing.T) {
        input := &datazone.CreateSubscriptionRequestInput{}
        output := &datazone.CreateSubscriptionRequestOutput{}

        mockClient.On("CreateSubscriptionRequest", ctx, input).Return(output, nil)

        result, err := mockClient.CreateSubscriptionRequest(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateSubscriptionTarget", func(t *testing.T) {
        input := &datazone.CreateSubscriptionTargetInput{}
        output := &datazone.CreateSubscriptionTargetOutput{}

        mockClient.On("CreateSubscriptionTarget", ctx, input).Return(output, nil)

        result, err := mockClient.CreateSubscriptionTarget(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateUserProfile", func(t *testing.T) {
        input := &datazone.CreateUserProfileInput{}
        output := &datazone.CreateUserProfileOutput{}

        mockClient.On("CreateUserProfile", ctx, input).Return(output, nil)

        result, err := mockClient.CreateUserProfile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteAsset", func(t *testing.T) {
        input := &datazone.DeleteAssetInput{}
        output := &datazone.DeleteAssetOutput{}

        mockClient.On("DeleteAsset", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteAsset(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteAssetType", func(t *testing.T) {
        input := &datazone.DeleteAssetTypeInput{}
        output := &datazone.DeleteAssetTypeOutput{}

        mockClient.On("DeleteAssetType", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteAssetType(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteDataSource", func(t *testing.T) {
        input := &datazone.DeleteDataSourceInput{}
        output := &datazone.DeleteDataSourceOutput{}

        mockClient.On("DeleteDataSource", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteDataSource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteDomain", func(t *testing.T) {
        input := &datazone.DeleteDomainInput{}
        output := &datazone.DeleteDomainOutput{}

        mockClient.On("DeleteDomain", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteDomain(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteEnvironment", func(t *testing.T) {
        input := &datazone.DeleteEnvironmentInput{}
        output := &datazone.DeleteEnvironmentOutput{}

        mockClient.On("DeleteEnvironment", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteEnvironment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteEnvironmentAction", func(t *testing.T) {
        input := &datazone.DeleteEnvironmentActionInput{}
        output := &datazone.DeleteEnvironmentActionOutput{}

        mockClient.On("DeleteEnvironmentAction", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteEnvironmentAction(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteEnvironmentBlueprintConfiguration", func(t *testing.T) {
        input := &datazone.DeleteEnvironmentBlueprintConfigurationInput{}
        output := &datazone.DeleteEnvironmentBlueprintConfigurationOutput{}

        mockClient.On("DeleteEnvironmentBlueprintConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteEnvironmentBlueprintConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteEnvironmentProfile", func(t *testing.T) {
        input := &datazone.DeleteEnvironmentProfileInput{}
        output := &datazone.DeleteEnvironmentProfileOutput{}

        mockClient.On("DeleteEnvironmentProfile", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteEnvironmentProfile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteFormType", func(t *testing.T) {
        input := &datazone.DeleteFormTypeInput{}
        output := &datazone.DeleteFormTypeOutput{}

        mockClient.On("DeleteFormType", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteFormType(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteGlossary", func(t *testing.T) {
        input := &datazone.DeleteGlossaryInput{}
        output := &datazone.DeleteGlossaryOutput{}

        mockClient.On("DeleteGlossary", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteGlossary(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteGlossaryTerm", func(t *testing.T) {
        input := &datazone.DeleteGlossaryTermInput{}
        output := &datazone.DeleteGlossaryTermOutput{}

        mockClient.On("DeleteGlossaryTerm", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteGlossaryTerm(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteListing", func(t *testing.T) {
        input := &datazone.DeleteListingInput{}
        output := &datazone.DeleteListingOutput{}

        mockClient.On("DeleteListing", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteListing(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteProject", func(t *testing.T) {
        input := &datazone.DeleteProjectInput{}
        output := &datazone.DeleteProjectOutput{}

        mockClient.On("DeleteProject", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteProject(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteProjectMembership", func(t *testing.T) {
        input := &datazone.DeleteProjectMembershipInput{}
        output := &datazone.DeleteProjectMembershipOutput{}

        mockClient.On("DeleteProjectMembership", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteProjectMembership(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteSubscriptionGrant", func(t *testing.T) {
        input := &datazone.DeleteSubscriptionGrantInput{}
        output := &datazone.DeleteSubscriptionGrantOutput{}

        mockClient.On("DeleteSubscriptionGrant", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteSubscriptionGrant(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteSubscriptionRequest", func(t *testing.T) {
        input := &datazone.DeleteSubscriptionRequestInput{}
        output := &datazone.DeleteSubscriptionRequestOutput{}

        mockClient.On("DeleteSubscriptionRequest", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteSubscriptionRequest(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteSubscriptionTarget", func(t *testing.T) {
        input := &datazone.DeleteSubscriptionTargetInput{}
        output := &datazone.DeleteSubscriptionTargetOutput{}

        mockClient.On("DeleteSubscriptionTarget", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteSubscriptionTarget(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteTimeSeriesDataPoints", func(t *testing.T) {
        input := &datazone.DeleteTimeSeriesDataPointsInput{}
        output := &datazone.DeleteTimeSeriesDataPointsOutput{}

        mockClient.On("DeleteTimeSeriesDataPoints", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteTimeSeriesDataPoints(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateEnvironmentRole", func(t *testing.T) {
        input := &datazone.DisassociateEnvironmentRoleInput{}
        output := &datazone.DisassociateEnvironmentRoleOutput{}

        mockClient.On("DisassociateEnvironmentRole", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateEnvironmentRole(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetAsset", func(t *testing.T) {
        input := &datazone.GetAssetInput{}
        output := &datazone.GetAssetOutput{}

        mockClient.On("GetAsset", ctx, input).Return(output, nil)

        result, err := mockClient.GetAsset(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetAssetType", func(t *testing.T) {
        input := &datazone.GetAssetTypeInput{}
        output := &datazone.GetAssetTypeOutput{}

        mockClient.On("GetAssetType", ctx, input).Return(output, nil)

        result, err := mockClient.GetAssetType(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDataSource", func(t *testing.T) {
        input := &datazone.GetDataSourceInput{}
        output := &datazone.GetDataSourceOutput{}

        mockClient.On("GetDataSource", ctx, input).Return(output, nil)

        result, err := mockClient.GetDataSource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDataSourceRun", func(t *testing.T) {
        input := &datazone.GetDataSourceRunInput{}
        output := &datazone.GetDataSourceRunOutput{}

        mockClient.On("GetDataSourceRun", ctx, input).Return(output, nil)

        result, err := mockClient.GetDataSourceRun(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDomain", func(t *testing.T) {
        input := &datazone.GetDomainInput{}
        output := &datazone.GetDomainOutput{}

        mockClient.On("GetDomain", ctx, input).Return(output, nil)

        result, err := mockClient.GetDomain(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetEnvironment", func(t *testing.T) {
        input := &datazone.GetEnvironmentInput{}
        output := &datazone.GetEnvironmentOutput{}

        mockClient.On("GetEnvironment", ctx, input).Return(output, nil)

        result, err := mockClient.GetEnvironment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetEnvironmentAction", func(t *testing.T) {
        input := &datazone.GetEnvironmentActionInput{}
        output := &datazone.GetEnvironmentActionOutput{}

        mockClient.On("GetEnvironmentAction", ctx, input).Return(output, nil)

        result, err := mockClient.GetEnvironmentAction(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetEnvironmentBlueprint", func(t *testing.T) {
        input := &datazone.GetEnvironmentBlueprintInput{}
        output := &datazone.GetEnvironmentBlueprintOutput{}

        mockClient.On("GetEnvironmentBlueprint", ctx, input).Return(output, nil)

        result, err := mockClient.GetEnvironmentBlueprint(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetEnvironmentBlueprintConfiguration", func(t *testing.T) {
        input := &datazone.GetEnvironmentBlueprintConfigurationInput{}
        output := &datazone.GetEnvironmentBlueprintConfigurationOutput{}

        mockClient.On("GetEnvironmentBlueprintConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.GetEnvironmentBlueprintConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetEnvironmentProfile", func(t *testing.T) {
        input := &datazone.GetEnvironmentProfileInput{}
        output := &datazone.GetEnvironmentProfileOutput{}

        mockClient.On("GetEnvironmentProfile", ctx, input).Return(output, nil)

        result, err := mockClient.GetEnvironmentProfile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetFormType", func(t *testing.T) {
        input := &datazone.GetFormTypeInput{}
        output := &datazone.GetFormTypeOutput{}

        mockClient.On("GetFormType", ctx, input).Return(output, nil)

        result, err := mockClient.GetFormType(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetGlossary", func(t *testing.T) {
        input := &datazone.GetGlossaryInput{}
        output := &datazone.GetGlossaryOutput{}

        mockClient.On("GetGlossary", ctx, input).Return(output, nil)

        result, err := mockClient.GetGlossary(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetGlossaryTerm", func(t *testing.T) {
        input := &datazone.GetGlossaryTermInput{}
        output := &datazone.GetGlossaryTermOutput{}

        mockClient.On("GetGlossaryTerm", ctx, input).Return(output, nil)

        result, err := mockClient.GetGlossaryTerm(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetGroupProfile", func(t *testing.T) {
        input := &datazone.GetGroupProfileInput{}
        output := &datazone.GetGroupProfileOutput{}

        mockClient.On("GetGroupProfile", ctx, input).Return(output, nil)

        result, err := mockClient.GetGroupProfile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetIamPortalLoginUrl", func(t *testing.T) {
        input := &datazone.GetIamPortalLoginUrlInput{}
        output := &datazone.GetIamPortalLoginUrlOutput{}

        mockClient.On("GetIamPortalLoginUrl", ctx, input).Return(output, nil)

        result, err := mockClient.GetIamPortalLoginUrl(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetLineageNode", func(t *testing.T) {
        input := &datazone.GetLineageNodeInput{}
        output := &datazone.GetLineageNodeOutput{}

        mockClient.On("GetLineageNode", ctx, input).Return(output, nil)

        result, err := mockClient.GetLineageNode(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetListing", func(t *testing.T) {
        input := &datazone.GetListingInput{}
        output := &datazone.GetListingOutput{}

        mockClient.On("GetListing", ctx, input).Return(output, nil)

        result, err := mockClient.GetListing(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetMetadataGenerationRun", func(t *testing.T) {
        input := &datazone.GetMetadataGenerationRunInput{}
        output := &datazone.GetMetadataGenerationRunOutput{}

        mockClient.On("GetMetadataGenerationRun", ctx, input).Return(output, nil)

        result, err := mockClient.GetMetadataGenerationRun(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetProject", func(t *testing.T) {
        input := &datazone.GetProjectInput{}
        output := &datazone.GetProjectOutput{}

        mockClient.On("GetProject", ctx, input).Return(output, nil)

        result, err := mockClient.GetProject(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSubscription", func(t *testing.T) {
        input := &datazone.GetSubscriptionInput{}
        output := &datazone.GetSubscriptionOutput{}

        mockClient.On("GetSubscription", ctx, input).Return(output, nil)

        result, err := mockClient.GetSubscription(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSubscriptionGrant", func(t *testing.T) {
        input := &datazone.GetSubscriptionGrantInput{}
        output := &datazone.GetSubscriptionGrantOutput{}

        mockClient.On("GetSubscriptionGrant", ctx, input).Return(output, nil)

        result, err := mockClient.GetSubscriptionGrant(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSubscriptionRequestDetails", func(t *testing.T) {
        input := &datazone.GetSubscriptionRequestDetailsInput{}
        output := &datazone.GetSubscriptionRequestDetailsOutput{}

        mockClient.On("GetSubscriptionRequestDetails", ctx, input).Return(output, nil)

        result, err := mockClient.GetSubscriptionRequestDetails(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSubscriptionTarget", func(t *testing.T) {
        input := &datazone.GetSubscriptionTargetInput{}
        output := &datazone.GetSubscriptionTargetOutput{}

        mockClient.On("GetSubscriptionTarget", ctx, input).Return(output, nil)

        result, err := mockClient.GetSubscriptionTarget(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetTimeSeriesDataPoint", func(t *testing.T) {
        input := &datazone.GetTimeSeriesDataPointInput{}
        output := &datazone.GetTimeSeriesDataPointOutput{}

        mockClient.On("GetTimeSeriesDataPoint", ctx, input).Return(output, nil)

        result, err := mockClient.GetTimeSeriesDataPoint(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetUserProfile", func(t *testing.T) {
        input := &datazone.GetUserProfileInput{}
        output := &datazone.GetUserProfileOutput{}

        mockClient.On("GetUserProfile", ctx, input).Return(output, nil)

        result, err := mockClient.GetUserProfile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAssetRevisions", func(t *testing.T) {
        input := &datazone.ListAssetRevisionsInput{}
        output := &datazone.ListAssetRevisionsOutput{}

        mockClient.On("ListAssetRevisions", ctx, input).Return(output, nil)

        result, err := mockClient.ListAssetRevisions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDataSourceRunActivities", func(t *testing.T) {
        input := &datazone.ListDataSourceRunActivitiesInput{}
        output := &datazone.ListDataSourceRunActivitiesOutput{}

        mockClient.On("ListDataSourceRunActivities", ctx, input).Return(output, nil)

        result, err := mockClient.ListDataSourceRunActivities(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDataSourceRuns", func(t *testing.T) {
        input := &datazone.ListDataSourceRunsInput{}
        output := &datazone.ListDataSourceRunsOutput{}

        mockClient.On("ListDataSourceRuns", ctx, input).Return(output, nil)

        result, err := mockClient.ListDataSourceRuns(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDataSources", func(t *testing.T) {
        input := &datazone.ListDataSourcesInput{}
        output := &datazone.ListDataSourcesOutput{}

        mockClient.On("ListDataSources", ctx, input).Return(output, nil)

        result, err := mockClient.ListDataSources(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDomains", func(t *testing.T) {
        input := &datazone.ListDomainsInput{}
        output := &datazone.ListDomainsOutput{}

        mockClient.On("ListDomains", ctx, input).Return(output, nil)

        result, err := mockClient.ListDomains(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListEnvironmentActions", func(t *testing.T) {
        input := &datazone.ListEnvironmentActionsInput{}
        output := &datazone.ListEnvironmentActionsOutput{}

        mockClient.On("ListEnvironmentActions", ctx, input).Return(output, nil)

        result, err := mockClient.ListEnvironmentActions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListEnvironmentBlueprintConfigurations", func(t *testing.T) {
        input := &datazone.ListEnvironmentBlueprintConfigurationsInput{}
        output := &datazone.ListEnvironmentBlueprintConfigurationsOutput{}

        mockClient.On("ListEnvironmentBlueprintConfigurations", ctx, input).Return(output, nil)

        result, err := mockClient.ListEnvironmentBlueprintConfigurations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListEnvironmentBlueprints", func(t *testing.T) {
        input := &datazone.ListEnvironmentBlueprintsInput{}
        output := &datazone.ListEnvironmentBlueprintsOutput{}

        mockClient.On("ListEnvironmentBlueprints", ctx, input).Return(output, nil)

        result, err := mockClient.ListEnvironmentBlueprints(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListEnvironmentProfiles", func(t *testing.T) {
        input := &datazone.ListEnvironmentProfilesInput{}
        output := &datazone.ListEnvironmentProfilesOutput{}

        mockClient.On("ListEnvironmentProfiles", ctx, input).Return(output, nil)

        result, err := mockClient.ListEnvironmentProfiles(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListEnvironments", func(t *testing.T) {
        input := &datazone.ListEnvironmentsInput{}
        output := &datazone.ListEnvironmentsOutput{}

        mockClient.On("ListEnvironments", ctx, input).Return(output, nil)

        result, err := mockClient.ListEnvironments(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListLineageNodeHistory", func(t *testing.T) {
        input := &datazone.ListLineageNodeHistoryInput{}
        output := &datazone.ListLineageNodeHistoryOutput{}

        mockClient.On("ListLineageNodeHistory", ctx, input).Return(output, nil)

        result, err := mockClient.ListLineageNodeHistory(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListMetadataGenerationRuns", func(t *testing.T) {
        input := &datazone.ListMetadataGenerationRunsInput{}
        output := &datazone.ListMetadataGenerationRunsOutput{}

        mockClient.On("ListMetadataGenerationRuns", ctx, input).Return(output, nil)

        result, err := mockClient.ListMetadataGenerationRuns(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListNotifications", func(t *testing.T) {
        input := &datazone.ListNotificationsInput{}
        output := &datazone.ListNotificationsOutput{}

        mockClient.On("ListNotifications", ctx, input).Return(output, nil)

        result, err := mockClient.ListNotifications(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListProjectMemberships", func(t *testing.T) {
        input := &datazone.ListProjectMembershipsInput{}
        output := &datazone.ListProjectMembershipsOutput{}

        mockClient.On("ListProjectMemberships", ctx, input).Return(output, nil)

        result, err := mockClient.ListProjectMemberships(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListProjects", func(t *testing.T) {
        input := &datazone.ListProjectsInput{}
        output := &datazone.ListProjectsOutput{}

        mockClient.On("ListProjects", ctx, input).Return(output, nil)

        result, err := mockClient.ListProjects(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListSubscriptionGrants", func(t *testing.T) {
        input := &datazone.ListSubscriptionGrantsInput{}
        output := &datazone.ListSubscriptionGrantsOutput{}

        mockClient.On("ListSubscriptionGrants", ctx, input).Return(output, nil)

        result, err := mockClient.ListSubscriptionGrants(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListSubscriptionRequests", func(t *testing.T) {
        input := &datazone.ListSubscriptionRequestsInput{}
        output := &datazone.ListSubscriptionRequestsOutput{}

        mockClient.On("ListSubscriptionRequests", ctx, input).Return(output, nil)

        result, err := mockClient.ListSubscriptionRequests(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListSubscriptionTargets", func(t *testing.T) {
        input := &datazone.ListSubscriptionTargetsInput{}
        output := &datazone.ListSubscriptionTargetsOutput{}

        mockClient.On("ListSubscriptionTargets", ctx, input).Return(output, nil)

        result, err := mockClient.ListSubscriptionTargets(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListSubscriptions", func(t *testing.T) {
        input := &datazone.ListSubscriptionsInput{}
        output := &datazone.ListSubscriptionsOutput{}

        mockClient.On("ListSubscriptions", ctx, input).Return(output, nil)

        result, err := mockClient.ListSubscriptions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &datazone.ListTagsForResourceInput{}
        output := &datazone.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTimeSeriesDataPoints", func(t *testing.T) {
        input := &datazone.ListTimeSeriesDataPointsInput{}
        output := &datazone.ListTimeSeriesDataPointsOutput{}

        mockClient.On("ListTimeSeriesDataPoints", ctx, input).Return(output, nil)

        result, err := mockClient.ListTimeSeriesDataPoints(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPostLineageEvent", func(t *testing.T) {
        input := &datazone.PostLineageEventInput{}
        output := &datazone.PostLineageEventOutput{}

        mockClient.On("PostLineageEvent", ctx, input).Return(output, nil)

        result, err := mockClient.PostLineageEvent(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPostTimeSeriesDataPoints", func(t *testing.T) {
        input := &datazone.PostTimeSeriesDataPointsInput{}
        output := &datazone.PostTimeSeriesDataPointsOutput{}

        mockClient.On("PostTimeSeriesDataPoints", ctx, input).Return(output, nil)

        result, err := mockClient.PostTimeSeriesDataPoints(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutEnvironmentBlueprintConfiguration", func(t *testing.T) {
        input := &datazone.PutEnvironmentBlueprintConfigurationInput{}
        output := &datazone.PutEnvironmentBlueprintConfigurationOutput{}

        mockClient.On("PutEnvironmentBlueprintConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.PutEnvironmentBlueprintConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRejectPredictions", func(t *testing.T) {
        input := &datazone.RejectPredictionsInput{}
        output := &datazone.RejectPredictionsOutput{}

        mockClient.On("RejectPredictions", ctx, input).Return(output, nil)

        result, err := mockClient.RejectPredictions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRejectSubscriptionRequest", func(t *testing.T) {
        input := &datazone.RejectSubscriptionRequestInput{}
        output := &datazone.RejectSubscriptionRequestOutput{}

        mockClient.On("RejectSubscriptionRequest", ctx, input).Return(output, nil)

        result, err := mockClient.RejectSubscriptionRequest(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRevokeSubscription", func(t *testing.T) {
        input := &datazone.RevokeSubscriptionInput{}
        output := &datazone.RevokeSubscriptionOutput{}

        mockClient.On("RevokeSubscription", ctx, input).Return(output, nil)

        result, err := mockClient.RevokeSubscription(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSearch", func(t *testing.T) {
        input := &datazone.SearchInput{}
        output := &datazone.SearchOutput{}

        mockClient.On("Search", ctx, input).Return(output, nil)

        result, err := mockClient.Search(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSearchGroupProfiles", func(t *testing.T) {
        input := &datazone.SearchGroupProfilesInput{}
        output := &datazone.SearchGroupProfilesOutput{}

        mockClient.On("SearchGroupProfiles", ctx, input).Return(output, nil)

        result, err := mockClient.SearchGroupProfiles(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSearchListings", func(t *testing.T) {
        input := &datazone.SearchListingsInput{}
        output := &datazone.SearchListingsOutput{}

        mockClient.On("SearchListings", ctx, input).Return(output, nil)

        result, err := mockClient.SearchListings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSearchTypes", func(t *testing.T) {
        input := &datazone.SearchTypesInput{}
        output := &datazone.SearchTypesOutput{}

        mockClient.On("SearchTypes", ctx, input).Return(output, nil)

        result, err := mockClient.SearchTypes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSearchUserProfiles", func(t *testing.T) {
        input := &datazone.SearchUserProfilesInput{}
        output := &datazone.SearchUserProfilesOutput{}

        mockClient.On("SearchUserProfiles", ctx, input).Return(output, nil)

        result, err := mockClient.SearchUserProfiles(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartDataSourceRun", func(t *testing.T) {
        input := &datazone.StartDataSourceRunInput{}
        output := &datazone.StartDataSourceRunOutput{}

        mockClient.On("StartDataSourceRun", ctx, input).Return(output, nil)

        result, err := mockClient.StartDataSourceRun(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartMetadataGenerationRun", func(t *testing.T) {
        input := &datazone.StartMetadataGenerationRunInput{}
        output := &datazone.StartMetadataGenerationRunOutput{}

        mockClient.On("StartMetadataGenerationRun", ctx, input).Return(output, nil)

        result, err := mockClient.StartMetadataGenerationRun(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &datazone.TagResourceInput{}
        output := &datazone.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &datazone.UntagResourceInput{}
        output := &datazone.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateDataSource", func(t *testing.T) {
        input := &datazone.UpdateDataSourceInput{}
        output := &datazone.UpdateDataSourceOutput{}

        mockClient.On("UpdateDataSource", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateDataSource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateDomain", func(t *testing.T) {
        input := &datazone.UpdateDomainInput{}
        output := &datazone.UpdateDomainOutput{}

        mockClient.On("UpdateDomain", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateDomain(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateEnvironment", func(t *testing.T) {
        input := &datazone.UpdateEnvironmentInput{}
        output := &datazone.UpdateEnvironmentOutput{}

        mockClient.On("UpdateEnvironment", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateEnvironment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateEnvironmentAction", func(t *testing.T) {
        input := &datazone.UpdateEnvironmentActionInput{}
        output := &datazone.UpdateEnvironmentActionOutput{}

        mockClient.On("UpdateEnvironmentAction", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateEnvironmentAction(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateEnvironmentProfile", func(t *testing.T) {
        input := &datazone.UpdateEnvironmentProfileInput{}
        output := &datazone.UpdateEnvironmentProfileOutput{}

        mockClient.On("UpdateEnvironmentProfile", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateEnvironmentProfile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateGlossary", func(t *testing.T) {
        input := &datazone.UpdateGlossaryInput{}
        output := &datazone.UpdateGlossaryOutput{}

        mockClient.On("UpdateGlossary", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateGlossary(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateGlossaryTerm", func(t *testing.T) {
        input := &datazone.UpdateGlossaryTermInput{}
        output := &datazone.UpdateGlossaryTermOutput{}

        mockClient.On("UpdateGlossaryTerm", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateGlossaryTerm(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateGroupProfile", func(t *testing.T) {
        input := &datazone.UpdateGroupProfileInput{}
        output := &datazone.UpdateGroupProfileOutput{}

        mockClient.On("UpdateGroupProfile", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateGroupProfile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateProject", func(t *testing.T) {
        input := &datazone.UpdateProjectInput{}
        output := &datazone.UpdateProjectOutput{}

        mockClient.On("UpdateProject", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateProject(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateSubscriptionGrantStatus", func(t *testing.T) {
        input := &datazone.UpdateSubscriptionGrantStatusInput{}
        output := &datazone.UpdateSubscriptionGrantStatusOutput{}

        mockClient.On("UpdateSubscriptionGrantStatus", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateSubscriptionGrantStatus(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateSubscriptionRequest", func(t *testing.T) {
        input := &datazone.UpdateSubscriptionRequestInput{}
        output := &datazone.UpdateSubscriptionRequestOutput{}

        mockClient.On("UpdateSubscriptionRequest", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateSubscriptionRequest(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateSubscriptionTarget", func(t *testing.T) {
        input := &datazone.UpdateSubscriptionTargetInput{}
        output := &datazone.UpdateSubscriptionTargetOutput{}

        mockClient.On("UpdateSubscriptionTarget", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateSubscriptionTarget(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateUserProfile", func(t *testing.T) {
        input := &datazone.UpdateUserProfileInput{}
        output := &datazone.UpdateUserProfileOutput{}

        mockClient.On("UpdateUserProfile", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateUserProfile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
