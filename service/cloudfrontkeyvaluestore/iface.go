
package cloudfrontkeyvaluestore

import (
    "github.com/aws/aws-sdk-go-v2/service/cloudfrontkeyvaluestore"
)

// ICloudfrontkeyvaluestore defines the interface for cloudfrontkeyvaluestore
type ICloudfrontkeyvaluestore interface {
 Options() Options 
 DeleteKey(ctx context.Context, params *DeleteKeyInput, optFns ...func(*Options)) (*DeleteKeyOutput, error) 
 DescribeKeyValueStore(ctx context.Context, params *DescribeKeyValueStoreInput, optFns ...func(*Options)) (*DescribeKeyValueStoreOutput, error) 
 GetKey(ctx context.Context, params *GetKeyInput, optFns ...func(*Options)) (*GetKeyOutput, error) 
 ListKeys(ctx context.Context, params *ListKeysInput, optFns ...func(*Options)) (*ListKeysOutput, error) 
 PutKey(ctx context.Context, params *PutKeyInput, optFns ...func(*Options)) (*PutKeyOutput, error) 
 UpdateKeys(ctx context.Context, params *UpdateKeysInput, optFns ...func(*Options)) (*UpdateKeysOutput, error) 
}
