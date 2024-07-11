
package lexmodelbuildingservice

import (
    "github.com/aws/aws-sdk-go-v2/service/lexmodelbuildingservice"
)

// ILexmodelbuildingservice defines the interface for lexmodelbuildingservice
type ILexmodelbuildingservice interface {
 Options() Options 
 CreateBotVersion(ctx context.Context, params *CreateBotVersionInput, optFns ...func(*Options)) (*CreateBotVersionOutput, error) 
 CreateIntentVersion(ctx context.Context, params *CreateIntentVersionInput, optFns ...func(*Options)) (*CreateIntentVersionOutput, error) 
 CreateSlotTypeVersion(ctx context.Context, params *CreateSlotTypeVersionInput, optFns ...func(*Options)) (*CreateSlotTypeVersionOutput, error) 
 DeleteBot(ctx context.Context, params *DeleteBotInput, optFns ...func(*Options)) (*DeleteBotOutput, error) 
 DeleteBotAlias(ctx context.Context, params *DeleteBotAliasInput, optFns ...func(*Options)) (*DeleteBotAliasOutput, error) 
 DeleteBotChannelAssociation(ctx context.Context, params *DeleteBotChannelAssociationInput, optFns ...func(*Options)) (*DeleteBotChannelAssociationOutput, error) 
 DeleteBotVersion(ctx context.Context, params *DeleteBotVersionInput, optFns ...func(*Options)) (*DeleteBotVersionOutput, error) 
 DeleteIntent(ctx context.Context, params *DeleteIntentInput, optFns ...func(*Options)) (*DeleteIntentOutput, error) 
 DeleteIntentVersion(ctx context.Context, params *DeleteIntentVersionInput, optFns ...func(*Options)) (*DeleteIntentVersionOutput, error) 
 DeleteSlotType(ctx context.Context, params *DeleteSlotTypeInput, optFns ...func(*Options)) (*DeleteSlotTypeOutput, error) 
 DeleteSlotTypeVersion(ctx context.Context, params *DeleteSlotTypeVersionInput, optFns ...func(*Options)) (*DeleteSlotTypeVersionOutput, error) 
 DeleteUtterances(ctx context.Context, params *DeleteUtterancesInput, optFns ...func(*Options)) (*DeleteUtterancesOutput, error) 
 GetBot(ctx context.Context, params *GetBotInput, optFns ...func(*Options)) (*GetBotOutput, error) 
 GetBotAlias(ctx context.Context, params *GetBotAliasInput, optFns ...func(*Options)) (*GetBotAliasOutput, error) 
 GetBotAliases(ctx context.Context, params *GetBotAliasesInput, optFns ...func(*Options)) (*GetBotAliasesOutput, error) 
 GetBotChannelAssociation(ctx context.Context, params *GetBotChannelAssociationInput, optFns ...func(*Options)) (*GetBotChannelAssociationOutput, error) 
 GetBotChannelAssociations(ctx context.Context, params *GetBotChannelAssociationsInput, optFns ...func(*Options)) (*GetBotChannelAssociationsOutput, error) 
 GetBotVersions(ctx context.Context, params *GetBotVersionsInput, optFns ...func(*Options)) (*GetBotVersionsOutput, error) 
 GetBots(ctx context.Context, params *GetBotsInput, optFns ...func(*Options)) (*GetBotsOutput, error) 
 GetBuiltinIntent(ctx context.Context, params *GetBuiltinIntentInput, optFns ...func(*Options)) (*GetBuiltinIntentOutput, error) 
 GetBuiltinIntents(ctx context.Context, params *GetBuiltinIntentsInput, optFns ...func(*Options)) (*GetBuiltinIntentsOutput, error) 
 GetBuiltinSlotTypes(ctx context.Context, params *GetBuiltinSlotTypesInput, optFns ...func(*Options)) (*GetBuiltinSlotTypesOutput, error) 
 GetExport(ctx context.Context, params *GetExportInput, optFns ...func(*Options)) (*GetExportOutput, error) 
 GetImport(ctx context.Context, params *GetImportInput, optFns ...func(*Options)) (*GetImportOutput, error) 
 GetIntent(ctx context.Context, params *GetIntentInput, optFns ...func(*Options)) (*GetIntentOutput, error) 
 GetIntentVersions(ctx context.Context, params *GetIntentVersionsInput, optFns ...func(*Options)) (*GetIntentVersionsOutput, error) 
 GetIntents(ctx context.Context, params *GetIntentsInput, optFns ...func(*Options)) (*GetIntentsOutput, error) 
 GetMigration(ctx context.Context, params *GetMigrationInput, optFns ...func(*Options)) (*GetMigrationOutput, error) 
 GetMigrations(ctx context.Context, params *GetMigrationsInput, optFns ...func(*Options)) (*GetMigrationsOutput, error) 
 GetSlotType(ctx context.Context, params *GetSlotTypeInput, optFns ...func(*Options)) (*GetSlotTypeOutput, error) 
 GetSlotTypeVersions(ctx context.Context, params *GetSlotTypeVersionsInput, optFns ...func(*Options)) (*GetSlotTypeVersionsOutput, error) 
 GetSlotTypes(ctx context.Context, params *GetSlotTypesInput, optFns ...func(*Options)) (*GetSlotTypesOutput, error) 
 GetUtterancesView(ctx context.Context, params *GetUtterancesViewInput, optFns ...func(*Options)) (*GetUtterancesViewOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 PutBot(ctx context.Context, params *PutBotInput, optFns ...func(*Options)) (*PutBotOutput, error) 
 PutBotAlias(ctx context.Context, params *PutBotAliasInput, optFns ...func(*Options)) (*PutBotAliasOutput, error) 
 PutIntent(ctx context.Context, params *PutIntentInput, optFns ...func(*Options)) (*PutIntentOutput, error) 
 PutSlotType(ctx context.Context, params *PutSlotTypeInput, optFns ...func(*Options)) (*PutSlotTypeOutput, error) 
 StartImport(ctx context.Context, params *StartImportInput, optFns ...func(*Options)) (*StartImportOutput, error) 
 StartMigration(ctx context.Context, params *StartMigrationInput, optFns ...func(*Options)) (*StartMigrationOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
}
