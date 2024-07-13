package costandusagereportservice_test

// tests for the costandusagereportservice service interface for this ../../../service/costandusagereportservice/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/costandusagereportservice"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/costandusagereportservice/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/costandusagereportservice/costandusagereportservice_iface"
	"github.com/stretchr/testify/assert"
)

func TestCostandusagereportserviceServiceCanBeMocked(t *testing.T) {
	var iface costandusagereportservice_iface.IClient
	iface = &costandusagereportservice.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := costandusagereportservice.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteReportDefinition", func(t *testing.T) {
        input := &costandusagereportservice.DeleteReportDefinitionInput{}
        output := &costandusagereportservice.DeleteReportDefinitionOutput{}

        mockClient.On("DeleteReportDefinition", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteReportDefinition(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeReportDefinitions", func(t *testing.T) {
        input := &costandusagereportservice.DescribeReportDefinitionsInput{}
        output := &costandusagereportservice.DescribeReportDefinitionsOutput{}

        mockClient.On("DescribeReportDefinitions", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeReportDefinitions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &costandusagereportservice.ListTagsForResourceInput{}
        output := &costandusagereportservice.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestModifyReportDefinition", func(t *testing.T) {
        input := &costandusagereportservice.ModifyReportDefinitionInput{}
        output := &costandusagereportservice.ModifyReportDefinitionOutput{}

        mockClient.On("ModifyReportDefinition", ctx, input).Return(output, nil)

        result, err := mockClient.ModifyReportDefinition(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutReportDefinition", func(t *testing.T) {
        input := &costandusagereportservice.PutReportDefinitionInput{}
        output := &costandusagereportservice.PutReportDefinitionOutput{}

        mockClient.On("PutReportDefinition", ctx, input).Return(output, nil)

        result, err := mockClient.PutReportDefinition(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &costandusagereportservice.TagResourceInput{}
        output := &costandusagereportservice.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &costandusagereportservice.UntagResourceInput{}
        output := &costandusagereportservice.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
