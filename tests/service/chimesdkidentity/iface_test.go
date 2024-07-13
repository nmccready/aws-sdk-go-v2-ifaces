package chimesdkidentity_test

// tests for the chimesdkidentity service interface for this ../../../service/chimesdkidentity/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/chimesdkidentity"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/chimesdkidentity/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/chimesdkidentity/chimesdkidentity_iface"
	"github.com/stretchr/testify/assert"
)

func TestChimesdkidentityServiceCanBeMocked(t *testing.T) {
	var iface chimesdkidentity_iface.IClient
	iface = &chimesdkidentity.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := chimesdkidentity.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateAppInstance", func(t *testing.T) {
        input := &chimesdkidentity.CreateAppInstanceInput{}
        output := &chimesdkidentity.CreateAppInstanceOutput{}

        mockClient.On("CreateAppInstance", ctx, input).Return(output, nil)

        result, err := mockClient.CreateAppInstance(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateAppInstanceAdmin", func(t *testing.T) {
        input := &chimesdkidentity.CreateAppInstanceAdminInput{}
        output := &chimesdkidentity.CreateAppInstanceAdminOutput{}

        mockClient.On("CreateAppInstanceAdmin", ctx, input).Return(output, nil)

        result, err := mockClient.CreateAppInstanceAdmin(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateAppInstanceBot", func(t *testing.T) {
        input := &chimesdkidentity.CreateAppInstanceBotInput{}
        output := &chimesdkidentity.CreateAppInstanceBotOutput{}

        mockClient.On("CreateAppInstanceBot", ctx, input).Return(output, nil)

        result, err := mockClient.CreateAppInstanceBot(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateAppInstanceUser", func(t *testing.T) {
        input := &chimesdkidentity.CreateAppInstanceUserInput{}
        output := &chimesdkidentity.CreateAppInstanceUserOutput{}

        mockClient.On("CreateAppInstanceUser", ctx, input).Return(output, nil)

        result, err := mockClient.CreateAppInstanceUser(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteAppInstance", func(t *testing.T) {
        input := &chimesdkidentity.DeleteAppInstanceInput{}
        output := &chimesdkidentity.DeleteAppInstanceOutput{}

        mockClient.On("DeleteAppInstance", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteAppInstance(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteAppInstanceAdmin", func(t *testing.T) {
        input := &chimesdkidentity.DeleteAppInstanceAdminInput{}
        output := &chimesdkidentity.DeleteAppInstanceAdminOutput{}

        mockClient.On("DeleteAppInstanceAdmin", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteAppInstanceAdmin(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteAppInstanceBot", func(t *testing.T) {
        input := &chimesdkidentity.DeleteAppInstanceBotInput{}
        output := &chimesdkidentity.DeleteAppInstanceBotOutput{}

        mockClient.On("DeleteAppInstanceBot", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteAppInstanceBot(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteAppInstanceUser", func(t *testing.T) {
        input := &chimesdkidentity.DeleteAppInstanceUserInput{}
        output := &chimesdkidentity.DeleteAppInstanceUserOutput{}

        mockClient.On("DeleteAppInstanceUser", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteAppInstanceUser(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeregisterAppInstanceUserEndpoint", func(t *testing.T) {
        input := &chimesdkidentity.DeregisterAppInstanceUserEndpointInput{}
        output := &chimesdkidentity.DeregisterAppInstanceUserEndpointOutput{}

        mockClient.On("DeregisterAppInstanceUserEndpoint", ctx, input).Return(output, nil)

        result, err := mockClient.DeregisterAppInstanceUserEndpoint(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeAppInstance", func(t *testing.T) {
        input := &chimesdkidentity.DescribeAppInstanceInput{}
        output := &chimesdkidentity.DescribeAppInstanceOutput{}

        mockClient.On("DescribeAppInstance", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeAppInstance(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeAppInstanceAdmin", func(t *testing.T) {
        input := &chimesdkidentity.DescribeAppInstanceAdminInput{}
        output := &chimesdkidentity.DescribeAppInstanceAdminOutput{}

        mockClient.On("DescribeAppInstanceAdmin", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeAppInstanceAdmin(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeAppInstanceBot", func(t *testing.T) {
        input := &chimesdkidentity.DescribeAppInstanceBotInput{}
        output := &chimesdkidentity.DescribeAppInstanceBotOutput{}

        mockClient.On("DescribeAppInstanceBot", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeAppInstanceBot(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeAppInstanceUser", func(t *testing.T) {
        input := &chimesdkidentity.DescribeAppInstanceUserInput{}
        output := &chimesdkidentity.DescribeAppInstanceUserOutput{}

        mockClient.On("DescribeAppInstanceUser", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeAppInstanceUser(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeAppInstanceUserEndpoint", func(t *testing.T) {
        input := &chimesdkidentity.DescribeAppInstanceUserEndpointInput{}
        output := &chimesdkidentity.DescribeAppInstanceUserEndpointOutput{}

        mockClient.On("DescribeAppInstanceUserEndpoint", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeAppInstanceUserEndpoint(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetAppInstanceRetentionSettings", func(t *testing.T) {
        input := &chimesdkidentity.GetAppInstanceRetentionSettingsInput{}
        output := &chimesdkidentity.GetAppInstanceRetentionSettingsOutput{}

        mockClient.On("GetAppInstanceRetentionSettings", ctx, input).Return(output, nil)

        result, err := mockClient.GetAppInstanceRetentionSettings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAppInstanceAdmins", func(t *testing.T) {
        input := &chimesdkidentity.ListAppInstanceAdminsInput{}
        output := &chimesdkidentity.ListAppInstanceAdminsOutput{}

        mockClient.On("ListAppInstanceAdmins", ctx, input).Return(output, nil)

        result, err := mockClient.ListAppInstanceAdmins(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAppInstanceBots", func(t *testing.T) {
        input := &chimesdkidentity.ListAppInstanceBotsInput{}
        output := &chimesdkidentity.ListAppInstanceBotsOutput{}

        mockClient.On("ListAppInstanceBots", ctx, input).Return(output, nil)

        result, err := mockClient.ListAppInstanceBots(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAppInstanceUserEndpoints", func(t *testing.T) {
        input := &chimesdkidentity.ListAppInstanceUserEndpointsInput{}
        output := &chimesdkidentity.ListAppInstanceUserEndpointsOutput{}

        mockClient.On("ListAppInstanceUserEndpoints", ctx, input).Return(output, nil)

        result, err := mockClient.ListAppInstanceUserEndpoints(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAppInstanceUsers", func(t *testing.T) {
        input := &chimesdkidentity.ListAppInstanceUsersInput{}
        output := &chimesdkidentity.ListAppInstanceUsersOutput{}

        mockClient.On("ListAppInstanceUsers", ctx, input).Return(output, nil)

        result, err := mockClient.ListAppInstanceUsers(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAppInstances", func(t *testing.T) {
        input := &chimesdkidentity.ListAppInstancesInput{}
        output := &chimesdkidentity.ListAppInstancesOutput{}

        mockClient.On("ListAppInstances", ctx, input).Return(output, nil)

        result, err := mockClient.ListAppInstances(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &chimesdkidentity.ListTagsForResourceInput{}
        output := &chimesdkidentity.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutAppInstanceRetentionSettings", func(t *testing.T) {
        input := &chimesdkidentity.PutAppInstanceRetentionSettingsInput{}
        output := &chimesdkidentity.PutAppInstanceRetentionSettingsOutput{}

        mockClient.On("PutAppInstanceRetentionSettings", ctx, input).Return(output, nil)

        result, err := mockClient.PutAppInstanceRetentionSettings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutAppInstanceUserExpirationSettings", func(t *testing.T) {
        input := &chimesdkidentity.PutAppInstanceUserExpirationSettingsInput{}
        output := &chimesdkidentity.PutAppInstanceUserExpirationSettingsOutput{}

        mockClient.On("PutAppInstanceUserExpirationSettings", ctx, input).Return(output, nil)

        result, err := mockClient.PutAppInstanceUserExpirationSettings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRegisterAppInstanceUserEndpoint", func(t *testing.T) {
        input := &chimesdkidentity.RegisterAppInstanceUserEndpointInput{}
        output := &chimesdkidentity.RegisterAppInstanceUserEndpointOutput{}

        mockClient.On("RegisterAppInstanceUserEndpoint", ctx, input).Return(output, nil)

        result, err := mockClient.RegisterAppInstanceUserEndpoint(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &chimesdkidentity.TagResourceInput{}
        output := &chimesdkidentity.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &chimesdkidentity.UntagResourceInput{}
        output := &chimesdkidentity.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateAppInstance", func(t *testing.T) {
        input := &chimesdkidentity.UpdateAppInstanceInput{}
        output := &chimesdkidentity.UpdateAppInstanceOutput{}

        mockClient.On("UpdateAppInstance", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateAppInstance(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateAppInstanceBot", func(t *testing.T) {
        input := &chimesdkidentity.UpdateAppInstanceBotInput{}
        output := &chimesdkidentity.UpdateAppInstanceBotOutput{}

        mockClient.On("UpdateAppInstanceBot", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateAppInstanceBot(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateAppInstanceUser", func(t *testing.T) {
        input := &chimesdkidentity.UpdateAppInstanceUserInput{}
        output := &chimesdkidentity.UpdateAppInstanceUserOutput{}

        mockClient.On("UpdateAppInstanceUser", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateAppInstanceUser(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateAppInstanceUserEndpoint", func(t *testing.T) {
        input := &chimesdkidentity.UpdateAppInstanceUserEndpointInput{}
        output := &chimesdkidentity.UpdateAppInstanceUserEndpointOutput{}

        mockClient.On("UpdateAppInstanceUserEndpoint", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateAppInstanceUserEndpoint(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
