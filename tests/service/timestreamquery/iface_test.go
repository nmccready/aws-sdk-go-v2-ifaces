package timestreamquery_test

// tests for the timestreamquery service interface for this ../../../service/timestreamquery/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/timestreamquery"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/timestreamquery/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/timestreamquery/timestreamquery_iface"
	"github.com/stretchr/testify/assert"
)

func TestTimestreamqueryServiceCanBeMocked(t *testing.T) {
	var iface timestreamquery_iface.IClient
	iface = &timestreamquery.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := timestreamquery.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCancelQuery", func(t *testing.T) {
        input := &timestreamquery.CancelQueryInput{}
        output := &timestreamquery.CancelQueryOutput{}

        mockClient.On("CancelQuery", ctx, input).Return(output, nil)

        result, err := mockClient.CancelQuery(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateScheduledQuery", func(t *testing.T) {
        input := &timestreamquery.CreateScheduledQueryInput{}
        output := &timestreamquery.CreateScheduledQueryOutput{}

        mockClient.On("CreateScheduledQuery", ctx, input).Return(output, nil)

        result, err := mockClient.CreateScheduledQuery(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteScheduledQuery", func(t *testing.T) {
        input := &timestreamquery.DeleteScheduledQueryInput{}
        output := &timestreamquery.DeleteScheduledQueryOutput{}

        mockClient.On("DeleteScheduledQuery", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteScheduledQuery(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeAccountSettings", func(t *testing.T) {
        input := &timestreamquery.DescribeAccountSettingsInput{}
        output := &timestreamquery.DescribeAccountSettingsOutput{}

        mockClient.On("DescribeAccountSettings", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeAccountSettings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeEndpoints", func(t *testing.T) {
        input := &timestreamquery.DescribeEndpointsInput{}
        output := &timestreamquery.DescribeEndpointsOutput{}

        mockClient.On("DescribeEndpoints", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeEndpoints(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeScheduledQuery", func(t *testing.T) {
        input := &timestreamquery.DescribeScheduledQueryInput{}
        output := &timestreamquery.DescribeScheduledQueryOutput{}

        mockClient.On("DescribeScheduledQuery", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeScheduledQuery(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestExecuteScheduledQuery", func(t *testing.T) {
        input := &timestreamquery.ExecuteScheduledQueryInput{}
        output := &timestreamquery.ExecuteScheduledQueryOutput{}

        mockClient.On("ExecuteScheduledQuery", ctx, input).Return(output, nil)

        result, err := mockClient.ExecuteScheduledQuery(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListScheduledQueries", func(t *testing.T) {
        input := &timestreamquery.ListScheduledQueriesInput{}
        output := &timestreamquery.ListScheduledQueriesOutput{}

        mockClient.On("ListScheduledQueries", ctx, input).Return(output, nil)

        result, err := mockClient.ListScheduledQueries(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &timestreamquery.ListTagsForResourceInput{}
        output := &timestreamquery.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPrepareQuery", func(t *testing.T) {
        input := &timestreamquery.PrepareQueryInput{}
        output := &timestreamquery.PrepareQueryOutput{}

        mockClient.On("PrepareQuery", ctx, input).Return(output, nil)

        result, err := mockClient.PrepareQuery(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestQuery", func(t *testing.T) {
        input := &timestreamquery.QueryInput{}
        output := &timestreamquery.QueryOutput{}

        mockClient.On("Query", ctx, input).Return(output, nil)

        result, err := mockClient.Query(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &timestreamquery.TagResourceInput{}
        output := &timestreamquery.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &timestreamquery.UntagResourceInput{}
        output := &timestreamquery.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateAccountSettings", func(t *testing.T) {
        input := &timestreamquery.UpdateAccountSettingsInput{}
        output := &timestreamquery.UpdateAccountSettingsOutput{}

        mockClient.On("UpdateAccountSettings", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateAccountSettings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateScheduledQuery", func(t *testing.T) {
        input := &timestreamquery.UpdateScheduledQueryInput{}
        output := &timestreamquery.UpdateScheduledQueryOutput{}

        mockClient.On("UpdateScheduledQuery", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateScheduledQuery(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
