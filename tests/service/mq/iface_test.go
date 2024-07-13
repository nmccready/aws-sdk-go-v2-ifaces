package mq_test

// tests for the mq service interface for this ../../../service/mq/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/mq"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/mq/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/mq/mq_iface"
	"github.com/stretchr/testify/assert"
)

func TestMqServiceCanBeMocked(t *testing.T) {
	var iface mq_iface.IClient
	iface = &mq.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := mq.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateBroker", func(t *testing.T) {
        input := &mq.CreateBrokerInput{}
        output := &mq.CreateBrokerOutput{}

        mockClient.On("CreateBroker", ctx, input).Return(output, nil)

        result, err := mockClient.CreateBroker(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateConfiguration", func(t *testing.T) {
        input := &mq.CreateConfigurationInput{}
        output := &mq.CreateConfigurationOutput{}

        mockClient.On("CreateConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.CreateConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateTags", func(t *testing.T) {
        input := &mq.CreateTagsInput{}
        output := &mq.CreateTagsOutput{}

        mockClient.On("CreateTags", ctx, input).Return(output, nil)

        result, err := mockClient.CreateTags(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateUser", func(t *testing.T) {
        input := &mq.CreateUserInput{}
        output := &mq.CreateUserOutput{}

        mockClient.On("CreateUser", ctx, input).Return(output, nil)

        result, err := mockClient.CreateUser(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteBroker", func(t *testing.T) {
        input := &mq.DeleteBrokerInput{}
        output := &mq.DeleteBrokerOutput{}

        mockClient.On("DeleteBroker", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteBroker(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteTags", func(t *testing.T) {
        input := &mq.DeleteTagsInput{}
        output := &mq.DeleteTagsOutput{}

        mockClient.On("DeleteTags", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteTags(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteUser", func(t *testing.T) {
        input := &mq.DeleteUserInput{}
        output := &mq.DeleteUserOutput{}

        mockClient.On("DeleteUser", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteUser(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeBroker", func(t *testing.T) {
        input := &mq.DescribeBrokerInput{}
        output := &mq.DescribeBrokerOutput{}

        mockClient.On("DescribeBroker", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeBroker(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeBrokerEngineTypes", func(t *testing.T) {
        input := &mq.DescribeBrokerEngineTypesInput{}
        output := &mq.DescribeBrokerEngineTypesOutput{}

        mockClient.On("DescribeBrokerEngineTypes", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeBrokerEngineTypes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeBrokerInstanceOptions", func(t *testing.T) {
        input := &mq.DescribeBrokerInstanceOptionsInput{}
        output := &mq.DescribeBrokerInstanceOptionsOutput{}

        mockClient.On("DescribeBrokerInstanceOptions", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeBrokerInstanceOptions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeConfiguration", func(t *testing.T) {
        input := &mq.DescribeConfigurationInput{}
        output := &mq.DescribeConfigurationOutput{}

        mockClient.On("DescribeConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeConfigurationRevision", func(t *testing.T) {
        input := &mq.DescribeConfigurationRevisionInput{}
        output := &mq.DescribeConfigurationRevisionOutput{}

        mockClient.On("DescribeConfigurationRevision", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeConfigurationRevision(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeUser", func(t *testing.T) {
        input := &mq.DescribeUserInput{}
        output := &mq.DescribeUserOutput{}

        mockClient.On("DescribeUser", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeUser(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListBrokers", func(t *testing.T) {
        input := &mq.ListBrokersInput{}
        output := &mq.ListBrokersOutput{}

        mockClient.On("ListBrokers", ctx, input).Return(output, nil)

        result, err := mockClient.ListBrokers(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListConfigurationRevisions", func(t *testing.T) {
        input := &mq.ListConfigurationRevisionsInput{}
        output := &mq.ListConfigurationRevisionsOutput{}

        mockClient.On("ListConfigurationRevisions", ctx, input).Return(output, nil)

        result, err := mockClient.ListConfigurationRevisions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListConfigurations", func(t *testing.T) {
        input := &mq.ListConfigurationsInput{}
        output := &mq.ListConfigurationsOutput{}

        mockClient.On("ListConfigurations", ctx, input).Return(output, nil)

        result, err := mockClient.ListConfigurations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTags", func(t *testing.T) {
        input := &mq.ListTagsInput{}
        output := &mq.ListTagsOutput{}

        mockClient.On("ListTags", ctx, input).Return(output, nil)

        result, err := mockClient.ListTags(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListUsers", func(t *testing.T) {
        input := &mq.ListUsersInput{}
        output := &mq.ListUsersOutput{}

        mockClient.On("ListUsers", ctx, input).Return(output, nil)

        result, err := mockClient.ListUsers(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPromote", func(t *testing.T) {
        input := &mq.PromoteInput{}
        output := &mq.PromoteOutput{}

        mockClient.On("Promote", ctx, input).Return(output, nil)

        result, err := mockClient.Promote(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRebootBroker", func(t *testing.T) {
        input := &mq.RebootBrokerInput{}
        output := &mq.RebootBrokerOutput{}

        mockClient.On("RebootBroker", ctx, input).Return(output, nil)

        result, err := mockClient.RebootBroker(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateBroker", func(t *testing.T) {
        input := &mq.UpdateBrokerInput{}
        output := &mq.UpdateBrokerOutput{}

        mockClient.On("UpdateBroker", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateBroker(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateConfiguration", func(t *testing.T) {
        input := &mq.UpdateConfigurationInput{}
        output := &mq.UpdateConfigurationOutput{}

        mockClient.On("UpdateConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateUser", func(t *testing.T) {
        input := &mq.UpdateUserInput{}
        output := &mq.UpdateUserOutput{}

        mockClient.On("UpdateUser", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateUser(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
