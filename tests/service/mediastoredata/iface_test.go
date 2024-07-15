// Code generated by ./scripts/gen_tests.sh DO NOT EDIT.
package mediastoredata_test

// tests for the mediastoredata service interface for 
// this ../../../service/mediastoredata/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/mediastoredata"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/mediastoredata/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/mediastoredata/mediastoredata_iface"
	"github.com/stretchr/testify/assert"
)

func TestMediastoredataServiceCanBeMocked(t *testing.T) {
	var iface mediastoredata_iface.IClient
	iface = &mediastoredata.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := mediastoredata.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteObject", func(t *testing.T) {
        input := &mediastoredata.DeleteObjectInput{}
        output := &mediastoredata.DeleteObjectOutput{}

        mockClient.On("DeleteObject", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteObject(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeObject", func(t *testing.T) {
        input := &mediastoredata.DescribeObjectInput{}
        output := &mediastoredata.DescribeObjectOutput{}

        mockClient.On("DescribeObject", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeObject(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetObject", func(t *testing.T) {
        input := &mediastoredata.GetObjectInput{}
        output := &mediastoredata.GetObjectOutput{}

        mockClient.On("GetObject", ctx, input).Return(output, nil)

        result, err := mockClient.GetObject(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListItems", func(t *testing.T) {
        input := &mediastoredata.ListItemsInput{}
        output := &mediastoredata.ListItemsOutput{}

        mockClient.On("ListItems", ctx, input).Return(output, nil)

        result, err := mockClient.ListItems(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutObject", func(t *testing.T) {
        input := &mediastoredata.PutObjectInput{}
        output := &mediastoredata.PutObjectOutput{}

        mockClient.On("PutObject", ctx, input).Return(output, nil)

        result, err := mockClient.PutObject(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
