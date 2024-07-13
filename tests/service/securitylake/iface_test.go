package securitylake_test

// tests for the securitylake service interface for this ../../../service/securitylake/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/securitylake"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/securitylake/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/securitylake/securitylake_iface"
	"github.com/stretchr/testify/assert"
)

func TestSecuritylakeServiceCanBeMocked(t *testing.T) {
	var iface securitylake_iface.IClient
	iface = &securitylake.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := securitylake.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateAwsLogSource", func(t *testing.T) {
        input := &securitylake.CreateAwsLogSourceInput{}
        output := &securitylake.CreateAwsLogSourceOutput{}

        mockClient.On("CreateAwsLogSource", ctx, input).Return(output, nil)

        result, err := mockClient.CreateAwsLogSource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateCustomLogSource", func(t *testing.T) {
        input := &securitylake.CreateCustomLogSourceInput{}
        output := &securitylake.CreateCustomLogSourceOutput{}

        mockClient.On("CreateCustomLogSource", ctx, input).Return(output, nil)

        result, err := mockClient.CreateCustomLogSource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateDataLake", func(t *testing.T) {
        input := &securitylake.CreateDataLakeInput{}
        output := &securitylake.CreateDataLakeOutput{}

        mockClient.On("CreateDataLake", ctx, input).Return(output, nil)

        result, err := mockClient.CreateDataLake(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateDataLakeExceptionSubscription", func(t *testing.T) {
        input := &securitylake.CreateDataLakeExceptionSubscriptionInput{}
        output := &securitylake.CreateDataLakeExceptionSubscriptionOutput{}

        mockClient.On("CreateDataLakeExceptionSubscription", ctx, input).Return(output, nil)

        result, err := mockClient.CreateDataLakeExceptionSubscription(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateDataLakeOrganizationConfiguration", func(t *testing.T) {
        input := &securitylake.CreateDataLakeOrganizationConfigurationInput{}
        output := &securitylake.CreateDataLakeOrganizationConfigurationOutput{}

        mockClient.On("CreateDataLakeOrganizationConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.CreateDataLakeOrganizationConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateSubscriber", func(t *testing.T) {
        input := &securitylake.CreateSubscriberInput{}
        output := &securitylake.CreateSubscriberOutput{}

        mockClient.On("CreateSubscriber", ctx, input).Return(output, nil)

        result, err := mockClient.CreateSubscriber(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateSubscriberNotification", func(t *testing.T) {
        input := &securitylake.CreateSubscriberNotificationInput{}
        output := &securitylake.CreateSubscriberNotificationOutput{}

        mockClient.On("CreateSubscriberNotification", ctx, input).Return(output, nil)

        result, err := mockClient.CreateSubscriberNotification(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteAwsLogSource", func(t *testing.T) {
        input := &securitylake.DeleteAwsLogSourceInput{}
        output := &securitylake.DeleteAwsLogSourceOutput{}

        mockClient.On("DeleteAwsLogSource", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteAwsLogSource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteCustomLogSource", func(t *testing.T) {
        input := &securitylake.DeleteCustomLogSourceInput{}
        output := &securitylake.DeleteCustomLogSourceOutput{}

        mockClient.On("DeleteCustomLogSource", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteCustomLogSource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteDataLake", func(t *testing.T) {
        input := &securitylake.DeleteDataLakeInput{}
        output := &securitylake.DeleteDataLakeOutput{}

        mockClient.On("DeleteDataLake", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteDataLake(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteDataLakeExceptionSubscription", func(t *testing.T) {
        input := &securitylake.DeleteDataLakeExceptionSubscriptionInput{}
        output := &securitylake.DeleteDataLakeExceptionSubscriptionOutput{}

        mockClient.On("DeleteDataLakeExceptionSubscription", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteDataLakeExceptionSubscription(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteDataLakeOrganizationConfiguration", func(t *testing.T) {
        input := &securitylake.DeleteDataLakeOrganizationConfigurationInput{}
        output := &securitylake.DeleteDataLakeOrganizationConfigurationOutput{}

        mockClient.On("DeleteDataLakeOrganizationConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteDataLakeOrganizationConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteSubscriber", func(t *testing.T) {
        input := &securitylake.DeleteSubscriberInput{}
        output := &securitylake.DeleteSubscriberOutput{}

        mockClient.On("DeleteSubscriber", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteSubscriber(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteSubscriberNotification", func(t *testing.T) {
        input := &securitylake.DeleteSubscriberNotificationInput{}
        output := &securitylake.DeleteSubscriberNotificationOutput{}

        mockClient.On("DeleteSubscriberNotification", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteSubscriberNotification(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeregisterDataLakeDelegatedAdministrator", func(t *testing.T) {
        input := &securitylake.DeregisterDataLakeDelegatedAdministratorInput{}
        output := &securitylake.DeregisterDataLakeDelegatedAdministratorOutput{}

        mockClient.On("DeregisterDataLakeDelegatedAdministrator", ctx, input).Return(output, nil)

        result, err := mockClient.DeregisterDataLakeDelegatedAdministrator(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDataLakeExceptionSubscription", func(t *testing.T) {
        input := &securitylake.GetDataLakeExceptionSubscriptionInput{}
        output := &securitylake.GetDataLakeExceptionSubscriptionOutput{}

        mockClient.On("GetDataLakeExceptionSubscription", ctx, input).Return(output, nil)

        result, err := mockClient.GetDataLakeExceptionSubscription(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDataLakeOrganizationConfiguration", func(t *testing.T) {
        input := &securitylake.GetDataLakeOrganizationConfigurationInput{}
        output := &securitylake.GetDataLakeOrganizationConfigurationOutput{}

        mockClient.On("GetDataLakeOrganizationConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.GetDataLakeOrganizationConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDataLakeSources", func(t *testing.T) {
        input := &securitylake.GetDataLakeSourcesInput{}
        output := &securitylake.GetDataLakeSourcesOutput{}

        mockClient.On("GetDataLakeSources", ctx, input).Return(output, nil)

        result, err := mockClient.GetDataLakeSources(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSubscriber", func(t *testing.T) {
        input := &securitylake.GetSubscriberInput{}
        output := &securitylake.GetSubscriberOutput{}

        mockClient.On("GetSubscriber", ctx, input).Return(output, nil)

        result, err := mockClient.GetSubscriber(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDataLakeExceptions", func(t *testing.T) {
        input := &securitylake.ListDataLakeExceptionsInput{}
        output := &securitylake.ListDataLakeExceptionsOutput{}

        mockClient.On("ListDataLakeExceptions", ctx, input).Return(output, nil)

        result, err := mockClient.ListDataLakeExceptions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDataLakes", func(t *testing.T) {
        input := &securitylake.ListDataLakesInput{}
        output := &securitylake.ListDataLakesOutput{}

        mockClient.On("ListDataLakes", ctx, input).Return(output, nil)

        result, err := mockClient.ListDataLakes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListLogSources", func(t *testing.T) {
        input := &securitylake.ListLogSourcesInput{}
        output := &securitylake.ListLogSourcesOutput{}

        mockClient.On("ListLogSources", ctx, input).Return(output, nil)

        result, err := mockClient.ListLogSources(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListSubscribers", func(t *testing.T) {
        input := &securitylake.ListSubscribersInput{}
        output := &securitylake.ListSubscribersOutput{}

        mockClient.On("ListSubscribers", ctx, input).Return(output, nil)

        result, err := mockClient.ListSubscribers(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &securitylake.ListTagsForResourceInput{}
        output := &securitylake.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRegisterDataLakeDelegatedAdministrator", func(t *testing.T) {
        input := &securitylake.RegisterDataLakeDelegatedAdministratorInput{}
        output := &securitylake.RegisterDataLakeDelegatedAdministratorOutput{}

        mockClient.On("RegisterDataLakeDelegatedAdministrator", ctx, input).Return(output, nil)

        result, err := mockClient.RegisterDataLakeDelegatedAdministrator(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &securitylake.TagResourceInput{}
        output := &securitylake.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &securitylake.UntagResourceInput{}
        output := &securitylake.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateDataLake", func(t *testing.T) {
        input := &securitylake.UpdateDataLakeInput{}
        output := &securitylake.UpdateDataLakeOutput{}

        mockClient.On("UpdateDataLake", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateDataLake(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateDataLakeExceptionSubscription", func(t *testing.T) {
        input := &securitylake.UpdateDataLakeExceptionSubscriptionInput{}
        output := &securitylake.UpdateDataLakeExceptionSubscriptionOutput{}

        mockClient.On("UpdateDataLakeExceptionSubscription", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateDataLakeExceptionSubscription(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateSubscriber", func(t *testing.T) {
        input := &securitylake.UpdateSubscriberInput{}
        output := &securitylake.UpdateSubscriberOutput{}

        mockClient.On("UpdateSubscriber", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateSubscriber(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateSubscriberNotification", func(t *testing.T) {
        input := &securitylake.UpdateSubscriberNotificationInput{}
        output := &securitylake.UpdateSubscriberNotificationOutput{}

        mockClient.On("UpdateSubscriberNotification", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateSubscriberNotification(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
