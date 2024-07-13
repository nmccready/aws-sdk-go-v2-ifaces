package migrationhubrefactorspaces_test

// tests for the migrationhubrefactorspaces service interface for this ../../../service/migrationhubrefactorspaces/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/migrationhubrefactorspaces"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/migrationhubrefactorspaces/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/migrationhubrefactorspaces/migrationhubrefactorspaces_iface"
	"github.com/stretchr/testify/assert"
)

func TestMigrationhubrefactorspacesServiceCanBeMocked(t *testing.T) {
	var iface migrationhubrefactorspaces_iface.IClient
	iface = &migrationhubrefactorspaces.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := migrationhubrefactorspaces.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateApplication", func(t *testing.T) {
        input := &migrationhubrefactorspaces.CreateApplicationInput{}
        output := &migrationhubrefactorspaces.CreateApplicationOutput{}

        mockClient.On("CreateApplication", ctx, input).Return(output, nil)

        result, err := mockClient.CreateApplication(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateEnvironment", func(t *testing.T) {
        input := &migrationhubrefactorspaces.CreateEnvironmentInput{}
        output := &migrationhubrefactorspaces.CreateEnvironmentOutput{}

        mockClient.On("CreateEnvironment", ctx, input).Return(output, nil)

        result, err := mockClient.CreateEnvironment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateRoute", func(t *testing.T) {
        input := &migrationhubrefactorspaces.CreateRouteInput{}
        output := &migrationhubrefactorspaces.CreateRouteOutput{}

        mockClient.On("CreateRoute", ctx, input).Return(output, nil)

        result, err := mockClient.CreateRoute(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateService", func(t *testing.T) {
        input := &migrationhubrefactorspaces.CreateServiceInput{}
        output := &migrationhubrefactorspaces.CreateServiceOutput{}

        mockClient.On("CreateService", ctx, input).Return(output, nil)

        result, err := mockClient.CreateService(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteApplication", func(t *testing.T) {
        input := &migrationhubrefactorspaces.DeleteApplicationInput{}
        output := &migrationhubrefactorspaces.DeleteApplicationOutput{}

        mockClient.On("DeleteApplication", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteApplication(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteEnvironment", func(t *testing.T) {
        input := &migrationhubrefactorspaces.DeleteEnvironmentInput{}
        output := &migrationhubrefactorspaces.DeleteEnvironmentOutput{}

        mockClient.On("DeleteEnvironment", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteEnvironment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteResourcePolicy", func(t *testing.T) {
        input := &migrationhubrefactorspaces.DeleteResourcePolicyInput{}
        output := &migrationhubrefactorspaces.DeleteResourcePolicyOutput{}

        mockClient.On("DeleteResourcePolicy", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteResourcePolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteRoute", func(t *testing.T) {
        input := &migrationhubrefactorspaces.DeleteRouteInput{}
        output := &migrationhubrefactorspaces.DeleteRouteOutput{}

        mockClient.On("DeleteRoute", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteRoute(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteService", func(t *testing.T) {
        input := &migrationhubrefactorspaces.DeleteServiceInput{}
        output := &migrationhubrefactorspaces.DeleteServiceOutput{}

        mockClient.On("DeleteService", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteService(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetApplication", func(t *testing.T) {
        input := &migrationhubrefactorspaces.GetApplicationInput{}
        output := &migrationhubrefactorspaces.GetApplicationOutput{}

        mockClient.On("GetApplication", ctx, input).Return(output, nil)

        result, err := mockClient.GetApplication(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetEnvironment", func(t *testing.T) {
        input := &migrationhubrefactorspaces.GetEnvironmentInput{}
        output := &migrationhubrefactorspaces.GetEnvironmentOutput{}

        mockClient.On("GetEnvironment", ctx, input).Return(output, nil)

        result, err := mockClient.GetEnvironment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetResourcePolicy", func(t *testing.T) {
        input := &migrationhubrefactorspaces.GetResourcePolicyInput{}
        output := &migrationhubrefactorspaces.GetResourcePolicyOutput{}

        mockClient.On("GetResourcePolicy", ctx, input).Return(output, nil)

        result, err := mockClient.GetResourcePolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetRoute", func(t *testing.T) {
        input := &migrationhubrefactorspaces.GetRouteInput{}
        output := &migrationhubrefactorspaces.GetRouteOutput{}

        mockClient.On("GetRoute", ctx, input).Return(output, nil)

        result, err := mockClient.GetRoute(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetService", func(t *testing.T) {
        input := &migrationhubrefactorspaces.GetServiceInput{}
        output := &migrationhubrefactorspaces.GetServiceOutput{}

        mockClient.On("GetService", ctx, input).Return(output, nil)

        result, err := mockClient.GetService(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListApplications", func(t *testing.T) {
        input := &migrationhubrefactorspaces.ListApplicationsInput{}
        output := &migrationhubrefactorspaces.ListApplicationsOutput{}

        mockClient.On("ListApplications", ctx, input).Return(output, nil)

        result, err := mockClient.ListApplications(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListEnvironmentVpcs", func(t *testing.T) {
        input := &migrationhubrefactorspaces.ListEnvironmentVpcsInput{}
        output := &migrationhubrefactorspaces.ListEnvironmentVpcsOutput{}

        mockClient.On("ListEnvironmentVpcs", ctx, input).Return(output, nil)

        result, err := mockClient.ListEnvironmentVpcs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListEnvironments", func(t *testing.T) {
        input := &migrationhubrefactorspaces.ListEnvironmentsInput{}
        output := &migrationhubrefactorspaces.ListEnvironmentsOutput{}

        mockClient.On("ListEnvironments", ctx, input).Return(output, nil)

        result, err := mockClient.ListEnvironments(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListRoutes", func(t *testing.T) {
        input := &migrationhubrefactorspaces.ListRoutesInput{}
        output := &migrationhubrefactorspaces.ListRoutesOutput{}

        mockClient.On("ListRoutes", ctx, input).Return(output, nil)

        result, err := mockClient.ListRoutes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListServices", func(t *testing.T) {
        input := &migrationhubrefactorspaces.ListServicesInput{}
        output := &migrationhubrefactorspaces.ListServicesOutput{}

        mockClient.On("ListServices", ctx, input).Return(output, nil)

        result, err := mockClient.ListServices(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &migrationhubrefactorspaces.ListTagsForResourceInput{}
        output := &migrationhubrefactorspaces.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutResourcePolicy", func(t *testing.T) {
        input := &migrationhubrefactorspaces.PutResourcePolicyInput{}
        output := &migrationhubrefactorspaces.PutResourcePolicyOutput{}

        mockClient.On("PutResourcePolicy", ctx, input).Return(output, nil)

        result, err := mockClient.PutResourcePolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &migrationhubrefactorspaces.TagResourceInput{}
        output := &migrationhubrefactorspaces.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &migrationhubrefactorspaces.UntagResourceInput{}
        output := &migrationhubrefactorspaces.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateRoute", func(t *testing.T) {
        input := &migrationhubrefactorspaces.UpdateRouteInput{}
        output := &migrationhubrefactorspaces.UpdateRouteOutput{}

        mockClient.On("UpdateRoute", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateRoute(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
