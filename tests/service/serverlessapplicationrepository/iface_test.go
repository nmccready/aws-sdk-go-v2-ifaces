// Code generated by ./scripts/gen_tests.sh DO NOT EDIT.
package serverlessapplicationrepository_test

// tests for the serverlessapplicationrepository service interface for 
// this ../../../service/serverlessapplicationrepository/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/serverlessapplicationrepository"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/serverlessapplicationrepository/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/serverlessapplicationrepository/serverlessapplicationrepository_iface"
	"github.com/stretchr/testify/assert"
)

func TestServerlessapplicationrepositoryServiceCanBeMocked(t *testing.T) {
	var iface serverlessapplicationrepository_iface.IClient
	iface = &serverlessapplicationrepository.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := serverlessapplicationrepository.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateApplication", func(t *testing.T) {
        input := &serverlessapplicationrepository.CreateApplicationInput{}
        output := &serverlessapplicationrepository.CreateApplicationOutput{}

        mockClient.On("CreateApplication", ctx, input).Return(output, nil)

        result, err := mockClient.CreateApplication(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateApplicationVersion", func(t *testing.T) {
        input := &serverlessapplicationrepository.CreateApplicationVersionInput{}
        output := &serverlessapplicationrepository.CreateApplicationVersionOutput{}

        mockClient.On("CreateApplicationVersion", ctx, input).Return(output, nil)

        result, err := mockClient.CreateApplicationVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateCloudFormationChangeSet", func(t *testing.T) {
        input := &serverlessapplicationrepository.CreateCloudFormationChangeSetInput{}
        output := &serverlessapplicationrepository.CreateCloudFormationChangeSetOutput{}

        mockClient.On("CreateCloudFormationChangeSet", ctx, input).Return(output, nil)

        result, err := mockClient.CreateCloudFormationChangeSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateCloudFormationTemplate", func(t *testing.T) {
        input := &serverlessapplicationrepository.CreateCloudFormationTemplateInput{}
        output := &serverlessapplicationrepository.CreateCloudFormationTemplateOutput{}

        mockClient.On("CreateCloudFormationTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.CreateCloudFormationTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteApplication", func(t *testing.T) {
        input := &serverlessapplicationrepository.DeleteApplicationInput{}
        output := &serverlessapplicationrepository.DeleteApplicationOutput{}

        mockClient.On("DeleteApplication", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteApplication(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetApplication", func(t *testing.T) {
        input := &serverlessapplicationrepository.GetApplicationInput{}
        output := &serverlessapplicationrepository.GetApplicationOutput{}

        mockClient.On("GetApplication", ctx, input).Return(output, nil)

        result, err := mockClient.GetApplication(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetApplicationPolicy", func(t *testing.T) {
        input := &serverlessapplicationrepository.GetApplicationPolicyInput{}
        output := &serverlessapplicationrepository.GetApplicationPolicyOutput{}

        mockClient.On("GetApplicationPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.GetApplicationPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetCloudFormationTemplate", func(t *testing.T) {
        input := &serverlessapplicationrepository.GetCloudFormationTemplateInput{}
        output := &serverlessapplicationrepository.GetCloudFormationTemplateOutput{}

        mockClient.On("GetCloudFormationTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.GetCloudFormationTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListApplicationDependencies", func(t *testing.T) {
        input := &serverlessapplicationrepository.ListApplicationDependenciesInput{}
        output := &serverlessapplicationrepository.ListApplicationDependenciesOutput{}

        mockClient.On("ListApplicationDependencies", ctx, input).Return(output, nil)

        result, err := mockClient.ListApplicationDependencies(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListApplicationVersions", func(t *testing.T) {
        input := &serverlessapplicationrepository.ListApplicationVersionsInput{}
        output := &serverlessapplicationrepository.ListApplicationVersionsOutput{}

        mockClient.On("ListApplicationVersions", ctx, input).Return(output, nil)

        result, err := mockClient.ListApplicationVersions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListApplications", func(t *testing.T) {
        input := &serverlessapplicationrepository.ListApplicationsInput{}
        output := &serverlessapplicationrepository.ListApplicationsOutput{}

        mockClient.On("ListApplications", ctx, input).Return(output, nil)

        result, err := mockClient.ListApplications(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutApplicationPolicy", func(t *testing.T) {
        input := &serverlessapplicationrepository.PutApplicationPolicyInput{}
        output := &serverlessapplicationrepository.PutApplicationPolicyOutput{}

        mockClient.On("PutApplicationPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.PutApplicationPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUnshareApplication", func(t *testing.T) {
        input := &serverlessapplicationrepository.UnshareApplicationInput{}
        output := &serverlessapplicationrepository.UnshareApplicationOutput{}

        mockClient.On("UnshareApplication", ctx, input).Return(output, nil)

        result, err := mockClient.UnshareApplication(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateApplication", func(t *testing.T) {
        input := &serverlessapplicationrepository.UpdateApplicationInput{}
        output := &serverlessapplicationrepository.UpdateApplicationOutput{}

        mockClient.On("UpdateApplication", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateApplication(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
