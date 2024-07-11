
package apigatewaymanagementapi

import (
    "github.com/aws/aws-sdk-go-v2/service/apigatewaymanagementapi"
)

// IClient defines the interface for apigatewaymanagementapi
type IClient interface {
 Options() Options 
 DeleteConnection(ctx context.Context, params *DeleteConnectionInput, optFns ...func(*Options)) (*DeleteConnectionOutput, error) 
 GetConnection(ctx context.Context, params *GetConnectionInput, optFns ...func(*Options)) (*GetConnectionOutput, error) 
 PostToConnection(ctx context.Context, params *PostToConnectionInput, optFns ...func(*Options)) (*PostToConnectionOutput, error) 
}
