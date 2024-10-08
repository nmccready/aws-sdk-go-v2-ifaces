
package bedrockagentruntime_iface

// This file was generated by scripts/gen_ifaces.sh do not edit it manually
import (
    "context"
    . "github.com/aws/aws-sdk-go-v2/service/bedrockagentruntime"
)

// IClient defines the interface for bedrockagentruntime
type IClient interface {
 Options() Options 
 DeleteAgentMemory(ctx context.Context, params *DeleteAgentMemoryInput, optFns ...func(*Options)) (*DeleteAgentMemoryOutput, error) 
 GetAgentMemory(ctx context.Context, params *GetAgentMemoryInput, optFns ...func(*Options)) (*GetAgentMemoryOutput, error) 
 InvokeAgent(ctx context.Context, params *InvokeAgentInput, optFns ...func(*Options)) (*InvokeAgentOutput, error) 
 InvokeFlow(ctx context.Context, params *InvokeFlowInput, optFns ...func(*Options)) (*InvokeFlowOutput, error) 
 Retrieve(ctx context.Context, params *RetrieveInput, optFns ...func(*Options)) (*RetrieveOutput, error) 
 RetrieveAndGenerate(ctx context.Context, params *RetrieveAndGenerateInput, optFns ...func(*Options)) (*RetrieveAndGenerateOutput, error) 
}
