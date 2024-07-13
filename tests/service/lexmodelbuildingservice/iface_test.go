package lexmodelbuildingservice_test

// tests for the lexmodelbuildingservice service interface for this ../../../service/lexmodelbuildingservice/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/lexmodelbuildingservice"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/lexmodelbuildingservice/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/lexmodelbuildingservice/lexmodelbuildingservice_iface"
	"github.com/stretchr/testify/assert"
)

func TestLexmodelbuildingserviceServiceCanBeMocked(t *testing.T) {
	var iface lexmodelbuildingservice_iface.IClient
	iface = &lexmodelbuildingservice.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := lexmodelbuildingservice.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateBotVersion", func(t *testing.T) {
        input := &lexmodelbuildingservice.CreateBotVersionInput{}
        output := &lexmodelbuildingservice.CreateBotVersionOutput{}

        mockClient.On("CreateBotVersion", ctx, input).Return(output, nil)

        result, err := mockClient.CreateBotVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateIntentVersion", func(t *testing.T) {
        input := &lexmodelbuildingservice.CreateIntentVersionInput{}
        output := &lexmodelbuildingservice.CreateIntentVersionOutput{}

        mockClient.On("CreateIntentVersion", ctx, input).Return(output, nil)

        result, err := mockClient.CreateIntentVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateSlotTypeVersion", func(t *testing.T) {
        input := &lexmodelbuildingservice.CreateSlotTypeVersionInput{}
        output := &lexmodelbuildingservice.CreateSlotTypeVersionOutput{}

        mockClient.On("CreateSlotTypeVersion", ctx, input).Return(output, nil)

        result, err := mockClient.CreateSlotTypeVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteBot", func(t *testing.T) {
        input := &lexmodelbuildingservice.DeleteBotInput{}
        output := &lexmodelbuildingservice.DeleteBotOutput{}

        mockClient.On("DeleteBot", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteBot(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteBotAlias", func(t *testing.T) {
        input := &lexmodelbuildingservice.DeleteBotAliasInput{}
        output := &lexmodelbuildingservice.DeleteBotAliasOutput{}

        mockClient.On("DeleteBotAlias", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteBotAlias(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteBotChannelAssociation", func(t *testing.T) {
        input := &lexmodelbuildingservice.DeleteBotChannelAssociationInput{}
        output := &lexmodelbuildingservice.DeleteBotChannelAssociationOutput{}

        mockClient.On("DeleteBotChannelAssociation", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteBotChannelAssociation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteBotVersion", func(t *testing.T) {
        input := &lexmodelbuildingservice.DeleteBotVersionInput{}
        output := &lexmodelbuildingservice.DeleteBotVersionOutput{}

        mockClient.On("DeleteBotVersion", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteBotVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteIntent", func(t *testing.T) {
        input := &lexmodelbuildingservice.DeleteIntentInput{}
        output := &lexmodelbuildingservice.DeleteIntentOutput{}

        mockClient.On("DeleteIntent", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteIntent(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteIntentVersion", func(t *testing.T) {
        input := &lexmodelbuildingservice.DeleteIntentVersionInput{}
        output := &lexmodelbuildingservice.DeleteIntentVersionOutput{}

        mockClient.On("DeleteIntentVersion", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteIntentVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteSlotType", func(t *testing.T) {
        input := &lexmodelbuildingservice.DeleteSlotTypeInput{}
        output := &lexmodelbuildingservice.DeleteSlotTypeOutput{}

        mockClient.On("DeleteSlotType", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteSlotType(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteSlotTypeVersion", func(t *testing.T) {
        input := &lexmodelbuildingservice.DeleteSlotTypeVersionInput{}
        output := &lexmodelbuildingservice.DeleteSlotTypeVersionOutput{}

        mockClient.On("DeleteSlotTypeVersion", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteSlotTypeVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteUtterances", func(t *testing.T) {
        input := &lexmodelbuildingservice.DeleteUtterancesInput{}
        output := &lexmodelbuildingservice.DeleteUtterancesOutput{}

        mockClient.On("DeleteUtterances", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteUtterances(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetBot", func(t *testing.T) {
        input := &lexmodelbuildingservice.GetBotInput{}
        output := &lexmodelbuildingservice.GetBotOutput{}

        mockClient.On("GetBot", ctx, input).Return(output, nil)

        result, err := mockClient.GetBot(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetBotAlias", func(t *testing.T) {
        input := &lexmodelbuildingservice.GetBotAliasInput{}
        output := &lexmodelbuildingservice.GetBotAliasOutput{}

        mockClient.On("GetBotAlias", ctx, input).Return(output, nil)

        result, err := mockClient.GetBotAlias(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetBotAliases", func(t *testing.T) {
        input := &lexmodelbuildingservice.GetBotAliasesInput{}
        output := &lexmodelbuildingservice.GetBotAliasesOutput{}

        mockClient.On("GetBotAliases", ctx, input).Return(output, nil)

        result, err := mockClient.GetBotAliases(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetBotChannelAssociation", func(t *testing.T) {
        input := &lexmodelbuildingservice.GetBotChannelAssociationInput{}
        output := &lexmodelbuildingservice.GetBotChannelAssociationOutput{}

        mockClient.On("GetBotChannelAssociation", ctx, input).Return(output, nil)

        result, err := mockClient.GetBotChannelAssociation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetBotChannelAssociations", func(t *testing.T) {
        input := &lexmodelbuildingservice.GetBotChannelAssociationsInput{}
        output := &lexmodelbuildingservice.GetBotChannelAssociationsOutput{}

        mockClient.On("GetBotChannelAssociations", ctx, input).Return(output, nil)

        result, err := mockClient.GetBotChannelAssociations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetBotVersions", func(t *testing.T) {
        input := &lexmodelbuildingservice.GetBotVersionsInput{}
        output := &lexmodelbuildingservice.GetBotVersionsOutput{}

        mockClient.On("GetBotVersions", ctx, input).Return(output, nil)

        result, err := mockClient.GetBotVersions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetBots", func(t *testing.T) {
        input := &lexmodelbuildingservice.GetBotsInput{}
        output := &lexmodelbuildingservice.GetBotsOutput{}

        mockClient.On("GetBots", ctx, input).Return(output, nil)

        result, err := mockClient.GetBots(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetBuiltinIntent", func(t *testing.T) {
        input := &lexmodelbuildingservice.GetBuiltinIntentInput{}
        output := &lexmodelbuildingservice.GetBuiltinIntentOutput{}

        mockClient.On("GetBuiltinIntent", ctx, input).Return(output, nil)

        result, err := mockClient.GetBuiltinIntent(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetBuiltinIntents", func(t *testing.T) {
        input := &lexmodelbuildingservice.GetBuiltinIntentsInput{}
        output := &lexmodelbuildingservice.GetBuiltinIntentsOutput{}

        mockClient.On("GetBuiltinIntents", ctx, input).Return(output, nil)

        result, err := mockClient.GetBuiltinIntents(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetBuiltinSlotTypes", func(t *testing.T) {
        input := &lexmodelbuildingservice.GetBuiltinSlotTypesInput{}
        output := &lexmodelbuildingservice.GetBuiltinSlotTypesOutput{}

        mockClient.On("GetBuiltinSlotTypes", ctx, input).Return(output, nil)

        result, err := mockClient.GetBuiltinSlotTypes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetExport", func(t *testing.T) {
        input := &lexmodelbuildingservice.GetExportInput{}
        output := &lexmodelbuildingservice.GetExportOutput{}

        mockClient.On("GetExport", ctx, input).Return(output, nil)

        result, err := mockClient.GetExport(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetImport", func(t *testing.T) {
        input := &lexmodelbuildingservice.GetImportInput{}
        output := &lexmodelbuildingservice.GetImportOutput{}

        mockClient.On("GetImport", ctx, input).Return(output, nil)

        result, err := mockClient.GetImport(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetIntent", func(t *testing.T) {
        input := &lexmodelbuildingservice.GetIntentInput{}
        output := &lexmodelbuildingservice.GetIntentOutput{}

        mockClient.On("GetIntent", ctx, input).Return(output, nil)

        result, err := mockClient.GetIntent(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetIntentVersions", func(t *testing.T) {
        input := &lexmodelbuildingservice.GetIntentVersionsInput{}
        output := &lexmodelbuildingservice.GetIntentVersionsOutput{}

        mockClient.On("GetIntentVersions", ctx, input).Return(output, nil)

        result, err := mockClient.GetIntentVersions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetIntents", func(t *testing.T) {
        input := &lexmodelbuildingservice.GetIntentsInput{}
        output := &lexmodelbuildingservice.GetIntentsOutput{}

        mockClient.On("GetIntents", ctx, input).Return(output, nil)

        result, err := mockClient.GetIntents(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetMigration", func(t *testing.T) {
        input := &lexmodelbuildingservice.GetMigrationInput{}
        output := &lexmodelbuildingservice.GetMigrationOutput{}

        mockClient.On("GetMigration", ctx, input).Return(output, nil)

        result, err := mockClient.GetMigration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetMigrations", func(t *testing.T) {
        input := &lexmodelbuildingservice.GetMigrationsInput{}
        output := &lexmodelbuildingservice.GetMigrationsOutput{}

        mockClient.On("GetMigrations", ctx, input).Return(output, nil)

        result, err := mockClient.GetMigrations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSlotType", func(t *testing.T) {
        input := &lexmodelbuildingservice.GetSlotTypeInput{}
        output := &lexmodelbuildingservice.GetSlotTypeOutput{}

        mockClient.On("GetSlotType", ctx, input).Return(output, nil)

        result, err := mockClient.GetSlotType(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSlotTypeVersions", func(t *testing.T) {
        input := &lexmodelbuildingservice.GetSlotTypeVersionsInput{}
        output := &lexmodelbuildingservice.GetSlotTypeVersionsOutput{}

        mockClient.On("GetSlotTypeVersions", ctx, input).Return(output, nil)

        result, err := mockClient.GetSlotTypeVersions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSlotTypes", func(t *testing.T) {
        input := &lexmodelbuildingservice.GetSlotTypesInput{}
        output := &lexmodelbuildingservice.GetSlotTypesOutput{}

        mockClient.On("GetSlotTypes", ctx, input).Return(output, nil)

        result, err := mockClient.GetSlotTypes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetUtterancesView", func(t *testing.T) {
        input := &lexmodelbuildingservice.GetUtterancesViewInput{}
        output := &lexmodelbuildingservice.GetUtterancesViewOutput{}

        mockClient.On("GetUtterancesView", ctx, input).Return(output, nil)

        result, err := mockClient.GetUtterancesView(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &lexmodelbuildingservice.ListTagsForResourceInput{}
        output := &lexmodelbuildingservice.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutBot", func(t *testing.T) {
        input := &lexmodelbuildingservice.PutBotInput{}
        output := &lexmodelbuildingservice.PutBotOutput{}

        mockClient.On("PutBot", ctx, input).Return(output, nil)

        result, err := mockClient.PutBot(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutBotAlias", func(t *testing.T) {
        input := &lexmodelbuildingservice.PutBotAliasInput{}
        output := &lexmodelbuildingservice.PutBotAliasOutput{}

        mockClient.On("PutBotAlias", ctx, input).Return(output, nil)

        result, err := mockClient.PutBotAlias(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutIntent", func(t *testing.T) {
        input := &lexmodelbuildingservice.PutIntentInput{}
        output := &lexmodelbuildingservice.PutIntentOutput{}

        mockClient.On("PutIntent", ctx, input).Return(output, nil)

        result, err := mockClient.PutIntent(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutSlotType", func(t *testing.T) {
        input := &lexmodelbuildingservice.PutSlotTypeInput{}
        output := &lexmodelbuildingservice.PutSlotTypeOutput{}

        mockClient.On("PutSlotType", ctx, input).Return(output, nil)

        result, err := mockClient.PutSlotType(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartImport", func(t *testing.T) {
        input := &lexmodelbuildingservice.StartImportInput{}
        output := &lexmodelbuildingservice.StartImportOutput{}

        mockClient.On("StartImport", ctx, input).Return(output, nil)

        result, err := mockClient.StartImport(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartMigration", func(t *testing.T) {
        input := &lexmodelbuildingservice.StartMigrationInput{}
        output := &lexmodelbuildingservice.StartMigrationOutput{}

        mockClient.On("StartMigration", ctx, input).Return(output, nil)

        result, err := mockClient.StartMigration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &lexmodelbuildingservice.TagResourceInput{}
        output := &lexmodelbuildingservice.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &lexmodelbuildingservice.UntagResourceInput{}
        output := &lexmodelbuildingservice.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
