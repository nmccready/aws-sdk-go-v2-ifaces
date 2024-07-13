package cloudfrontkeyvaluestore_test

// tests for the cloudfrontkeyvaluestore service interface for this ../../../service/cloudfrontkeyvaluestore/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/cloudfrontkeyvaluestore"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/cloudfrontkeyvaluestore/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/cloudfrontkeyvaluestore/cloudfrontkeyvaluestore_iface"
	"github.com/stretchr/testify/assert"
)

func TestCloudfrontkeyvaluestoreServiceCanBeMocked(t *testing.T) {
	var iface cloudfrontkeyvaluestore_iface.IClient
	iface = &cloudfrontkeyvaluestore.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := cloudfrontkeyvaluestore.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteKey", func(t *testing.T) {
        input := &cloudfrontkeyvaluestore.DeleteKeyInput{}
        output := &cloudfrontkeyvaluestore.DeleteKeyOutput{}

        mockClient.On("DeleteKey", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteKey(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeKeyValueStore", func(t *testing.T) {
        input := &cloudfrontkeyvaluestore.DescribeKeyValueStoreInput{}
        output := &cloudfrontkeyvaluestore.DescribeKeyValueStoreOutput{}

        mockClient.On("DescribeKeyValueStore", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeKeyValueStore(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetKey", func(t *testing.T) {
        input := &cloudfrontkeyvaluestore.GetKeyInput{}
        output := &cloudfrontkeyvaluestore.GetKeyOutput{}

        mockClient.On("GetKey", ctx, input).Return(output, nil)

        result, err := mockClient.GetKey(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListKeys", func(t *testing.T) {
        input := &cloudfrontkeyvaluestore.ListKeysInput{}
        output := &cloudfrontkeyvaluestore.ListKeysOutput{}

        mockClient.On("ListKeys", ctx, input).Return(output, nil)

        result, err := mockClient.ListKeys(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutKey", func(t *testing.T) {
        input := &cloudfrontkeyvaluestore.PutKeyInput{}
        output := &cloudfrontkeyvaluestore.PutKeyOutput{}

        mockClient.On("PutKey", ctx, input).Return(output, nil)

        result, err := mockClient.PutKey(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateKeys", func(t *testing.T) {
        input := &cloudfrontkeyvaluestore.UpdateKeysInput{}
        output := &cloudfrontkeyvaluestore.UpdateKeysOutput{}

        mockClient.On("UpdateKeys", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateKeys(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
