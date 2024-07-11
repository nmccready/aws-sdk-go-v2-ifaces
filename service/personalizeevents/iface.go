
package personalizeevents

import (
    "github.com/aws/aws-sdk-go-v2/service/personalizeevents"
)

// IPersonalizeevents defines the interface for personalizeevents
type IPersonalizeevents interface {
 Options() Options 
 PutActionInteractions(ctx context.Context, params *PutActionInteractionsInput, optFns ...func(*Options)) (*PutActionInteractionsOutput, error) 
 PutActions(ctx context.Context, params *PutActionsInput, optFns ...func(*Options)) (*PutActionsOutput, error) 
 PutEvents(ctx context.Context, params *PutEventsInput, optFns ...func(*Options)) (*PutEventsOutput, error) 
 PutItems(ctx context.Context, params *PutItemsInput, optFns ...func(*Options)) (*PutItemsOutput, error) 
 PutUsers(ctx context.Context, params *PutUsersInput, optFns ...func(*Options)) (*PutUsersOutput, error) 
}
