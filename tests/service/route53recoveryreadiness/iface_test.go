package route53recoveryreadiness_test

// tests for the route53recoveryreadiness service interface for this ../../../service/route53recoveryreadiness/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/route53recoveryreadiness"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/route53recoveryreadiness/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/route53recoveryreadiness/route53recoveryreadiness_iface"
	"github.com/stretchr/testify/assert"
)

func TestRoute53recoveryreadinessServiceCanBeMocked(t *testing.T) {
	var iface route53recoveryreadiness_iface.IClient
	iface = &route53recoveryreadiness.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := route53recoveryreadiness.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateCell", func(t *testing.T) {
        input := &route53recoveryreadiness.CreateCellInput{}
        output := &route53recoveryreadiness.CreateCellOutput{}

        mockClient.On("CreateCell", ctx, input).Return(output, nil)

        result, err := mockClient.CreateCell(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateCrossAccountAuthorization", func(t *testing.T) {
        input := &route53recoveryreadiness.CreateCrossAccountAuthorizationInput{}
        output := &route53recoveryreadiness.CreateCrossAccountAuthorizationOutput{}

        mockClient.On("CreateCrossAccountAuthorization", ctx, input).Return(output, nil)

        result, err := mockClient.CreateCrossAccountAuthorization(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateReadinessCheck", func(t *testing.T) {
        input := &route53recoveryreadiness.CreateReadinessCheckInput{}
        output := &route53recoveryreadiness.CreateReadinessCheckOutput{}

        mockClient.On("CreateReadinessCheck", ctx, input).Return(output, nil)

        result, err := mockClient.CreateReadinessCheck(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateRecoveryGroup", func(t *testing.T) {
        input := &route53recoveryreadiness.CreateRecoveryGroupInput{}
        output := &route53recoveryreadiness.CreateRecoveryGroupOutput{}

        mockClient.On("CreateRecoveryGroup", ctx, input).Return(output, nil)

        result, err := mockClient.CreateRecoveryGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateResourceSet", func(t *testing.T) {
        input := &route53recoveryreadiness.CreateResourceSetInput{}
        output := &route53recoveryreadiness.CreateResourceSetOutput{}

        mockClient.On("CreateResourceSet", ctx, input).Return(output, nil)

        result, err := mockClient.CreateResourceSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteCell", func(t *testing.T) {
        input := &route53recoveryreadiness.DeleteCellInput{}
        output := &route53recoveryreadiness.DeleteCellOutput{}

        mockClient.On("DeleteCell", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteCell(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteCrossAccountAuthorization", func(t *testing.T) {
        input := &route53recoveryreadiness.DeleteCrossAccountAuthorizationInput{}
        output := &route53recoveryreadiness.DeleteCrossAccountAuthorizationOutput{}

        mockClient.On("DeleteCrossAccountAuthorization", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteCrossAccountAuthorization(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteReadinessCheck", func(t *testing.T) {
        input := &route53recoveryreadiness.DeleteReadinessCheckInput{}
        output := &route53recoveryreadiness.DeleteReadinessCheckOutput{}

        mockClient.On("DeleteReadinessCheck", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteReadinessCheck(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteRecoveryGroup", func(t *testing.T) {
        input := &route53recoveryreadiness.DeleteRecoveryGroupInput{}
        output := &route53recoveryreadiness.DeleteRecoveryGroupOutput{}

        mockClient.On("DeleteRecoveryGroup", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteRecoveryGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteResourceSet", func(t *testing.T) {
        input := &route53recoveryreadiness.DeleteResourceSetInput{}
        output := &route53recoveryreadiness.DeleteResourceSetOutput{}

        mockClient.On("DeleteResourceSet", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteResourceSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetArchitectureRecommendations", func(t *testing.T) {
        input := &route53recoveryreadiness.GetArchitectureRecommendationsInput{}
        output := &route53recoveryreadiness.GetArchitectureRecommendationsOutput{}

        mockClient.On("GetArchitectureRecommendations", ctx, input).Return(output, nil)

        result, err := mockClient.GetArchitectureRecommendations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetCell", func(t *testing.T) {
        input := &route53recoveryreadiness.GetCellInput{}
        output := &route53recoveryreadiness.GetCellOutput{}

        mockClient.On("GetCell", ctx, input).Return(output, nil)

        result, err := mockClient.GetCell(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetCellReadinessSummary", func(t *testing.T) {
        input := &route53recoveryreadiness.GetCellReadinessSummaryInput{}
        output := &route53recoveryreadiness.GetCellReadinessSummaryOutput{}

        mockClient.On("GetCellReadinessSummary", ctx, input).Return(output, nil)

        result, err := mockClient.GetCellReadinessSummary(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetReadinessCheck", func(t *testing.T) {
        input := &route53recoveryreadiness.GetReadinessCheckInput{}
        output := &route53recoveryreadiness.GetReadinessCheckOutput{}

        mockClient.On("GetReadinessCheck", ctx, input).Return(output, nil)

        result, err := mockClient.GetReadinessCheck(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetReadinessCheckResourceStatus", func(t *testing.T) {
        input := &route53recoveryreadiness.GetReadinessCheckResourceStatusInput{}
        output := &route53recoveryreadiness.GetReadinessCheckResourceStatusOutput{}

        mockClient.On("GetReadinessCheckResourceStatus", ctx, input).Return(output, nil)

        result, err := mockClient.GetReadinessCheckResourceStatus(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetReadinessCheckStatus", func(t *testing.T) {
        input := &route53recoveryreadiness.GetReadinessCheckStatusInput{}
        output := &route53recoveryreadiness.GetReadinessCheckStatusOutput{}

        mockClient.On("GetReadinessCheckStatus", ctx, input).Return(output, nil)

        result, err := mockClient.GetReadinessCheckStatus(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetRecoveryGroup", func(t *testing.T) {
        input := &route53recoveryreadiness.GetRecoveryGroupInput{}
        output := &route53recoveryreadiness.GetRecoveryGroupOutput{}

        mockClient.On("GetRecoveryGroup", ctx, input).Return(output, nil)

        result, err := mockClient.GetRecoveryGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetRecoveryGroupReadinessSummary", func(t *testing.T) {
        input := &route53recoveryreadiness.GetRecoveryGroupReadinessSummaryInput{}
        output := &route53recoveryreadiness.GetRecoveryGroupReadinessSummaryOutput{}

        mockClient.On("GetRecoveryGroupReadinessSummary", ctx, input).Return(output, nil)

        result, err := mockClient.GetRecoveryGroupReadinessSummary(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetResourceSet", func(t *testing.T) {
        input := &route53recoveryreadiness.GetResourceSetInput{}
        output := &route53recoveryreadiness.GetResourceSetOutput{}

        mockClient.On("GetResourceSet", ctx, input).Return(output, nil)

        result, err := mockClient.GetResourceSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListCells", func(t *testing.T) {
        input := &route53recoveryreadiness.ListCellsInput{}
        output := &route53recoveryreadiness.ListCellsOutput{}

        mockClient.On("ListCells", ctx, input).Return(output, nil)

        result, err := mockClient.ListCells(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListCrossAccountAuthorizations", func(t *testing.T) {
        input := &route53recoveryreadiness.ListCrossAccountAuthorizationsInput{}
        output := &route53recoveryreadiness.ListCrossAccountAuthorizationsOutput{}

        mockClient.On("ListCrossAccountAuthorizations", ctx, input).Return(output, nil)

        result, err := mockClient.ListCrossAccountAuthorizations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListReadinessChecks", func(t *testing.T) {
        input := &route53recoveryreadiness.ListReadinessChecksInput{}
        output := &route53recoveryreadiness.ListReadinessChecksOutput{}

        mockClient.On("ListReadinessChecks", ctx, input).Return(output, nil)

        result, err := mockClient.ListReadinessChecks(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListRecoveryGroups", func(t *testing.T) {
        input := &route53recoveryreadiness.ListRecoveryGroupsInput{}
        output := &route53recoveryreadiness.ListRecoveryGroupsOutput{}

        mockClient.On("ListRecoveryGroups", ctx, input).Return(output, nil)

        result, err := mockClient.ListRecoveryGroups(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListResourceSets", func(t *testing.T) {
        input := &route53recoveryreadiness.ListResourceSetsInput{}
        output := &route53recoveryreadiness.ListResourceSetsOutput{}

        mockClient.On("ListResourceSets", ctx, input).Return(output, nil)

        result, err := mockClient.ListResourceSets(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListRules", func(t *testing.T) {
        input := &route53recoveryreadiness.ListRulesInput{}
        output := &route53recoveryreadiness.ListRulesOutput{}

        mockClient.On("ListRules", ctx, input).Return(output, nil)

        result, err := mockClient.ListRules(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResources", func(t *testing.T) {
        input := &route53recoveryreadiness.ListTagsForResourcesInput{}
        output := &route53recoveryreadiness.ListTagsForResourcesOutput{}

        mockClient.On("ListTagsForResources", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResources(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &route53recoveryreadiness.TagResourceInput{}
        output := &route53recoveryreadiness.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &route53recoveryreadiness.UntagResourceInput{}
        output := &route53recoveryreadiness.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateCell", func(t *testing.T) {
        input := &route53recoveryreadiness.UpdateCellInput{}
        output := &route53recoveryreadiness.UpdateCellOutput{}

        mockClient.On("UpdateCell", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateCell(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateReadinessCheck", func(t *testing.T) {
        input := &route53recoveryreadiness.UpdateReadinessCheckInput{}
        output := &route53recoveryreadiness.UpdateReadinessCheckOutput{}

        mockClient.On("UpdateReadinessCheck", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateReadinessCheck(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateRecoveryGroup", func(t *testing.T) {
        input := &route53recoveryreadiness.UpdateRecoveryGroupInput{}
        output := &route53recoveryreadiness.UpdateRecoveryGroupOutput{}

        mockClient.On("UpdateRecoveryGroup", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateRecoveryGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateResourceSet", func(t *testing.T) {
        input := &route53recoveryreadiness.UpdateResourceSetInput{}
        output := &route53recoveryreadiness.UpdateResourceSetOutput{}

        mockClient.On("UpdateResourceSet", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateResourceSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
