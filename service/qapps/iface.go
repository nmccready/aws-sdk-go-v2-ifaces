
package qapps

import (
    "github.com/aws/aws-sdk-go-v2/service/qapps"
)

// IClient defines the interface for qapps
type IClient interface {
 Options() Options 
 AssociateLibraryItemReview(ctx context.Context, params *AssociateLibraryItemReviewInput, optFns ...func(*Options)) (*AssociateLibraryItemReviewOutput, error) 
 AssociateQAppWithUser(ctx context.Context, params *AssociateQAppWithUserInput, optFns ...func(*Options)) (*AssociateQAppWithUserOutput, error) 
 CreateLibraryItem(ctx context.Context, params *CreateLibraryItemInput, optFns ...func(*Options)) (*CreateLibraryItemOutput, error) 
 CreateQApp(ctx context.Context, params *CreateQAppInput, optFns ...func(*Options)) (*CreateQAppOutput, error) 
 DeleteLibraryItem(ctx context.Context, params *DeleteLibraryItemInput, optFns ...func(*Options)) (*DeleteLibraryItemOutput, error) 
 DeleteQApp(ctx context.Context, params *DeleteQAppInput, optFns ...func(*Options)) (*DeleteQAppOutput, error) 
 DisassociateLibraryItemReview(ctx context.Context, params *DisassociateLibraryItemReviewInput, optFns ...func(*Options)) (*DisassociateLibraryItemReviewOutput, error) 
 DisassociateQAppFromUser(ctx context.Context, params *DisassociateQAppFromUserInput, optFns ...func(*Options)) (*DisassociateQAppFromUserOutput, error) 
 GetLibraryItem(ctx context.Context, params *GetLibraryItemInput, optFns ...func(*Options)) (*GetLibraryItemOutput, error) 
 GetQApp(ctx context.Context, params *GetQAppInput, optFns ...func(*Options)) (*GetQAppOutput, error) 
 GetQAppSession(ctx context.Context, params *GetQAppSessionInput, optFns ...func(*Options)) (*GetQAppSessionOutput, error) 
 ImportDocument(ctx context.Context, params *ImportDocumentInput, optFns ...func(*Options)) (*ImportDocumentOutput, error) 
 ListLibraryItems(ctx context.Context, params *ListLibraryItemsInput, optFns ...func(*Options)) (*ListLibraryItemsOutput, error) 
 ListQApps(ctx context.Context, params *ListQAppsInput, optFns ...func(*Options)) (*ListQAppsOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 PredictQApp(ctx context.Context, params *PredictQAppInput, optFns ...func(*Options)) (*PredictQAppOutput, error) 
 StartQAppSession(ctx context.Context, params *StartQAppSessionInput, optFns ...func(*Options)) (*StartQAppSessionOutput, error) 
 StopQAppSession(ctx context.Context, params *StopQAppSessionInput, optFns ...func(*Options)) (*StopQAppSessionOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateLibraryItem(ctx context.Context, params *UpdateLibraryItemInput, optFns ...func(*Options)) (*UpdateLibraryItemOutput, error) 
 UpdateQApp(ctx context.Context, params *UpdateQAppInput, optFns ...func(*Options)) (*UpdateQAppOutput, error) 
 UpdateQAppSession(ctx context.Context, params *UpdateQAppSessionInput, optFns ...func(*Options)) (*UpdateQAppSessionOutput, error) 
}
