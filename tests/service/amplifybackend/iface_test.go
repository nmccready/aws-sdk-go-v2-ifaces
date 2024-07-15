// Code generated by ./scripts/gen_tests.sh DO NOT EDIT.
package amplifybackend_test

// tests for the amplifybackend service interface for 
// this ../../../service/amplifybackend/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/amplifybackend"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/amplifybackend/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/amplifybackend/amplifybackend_iface"
	"github.com/stretchr/testify/assert"
)

func TestAmplifybackendServiceCanBeMocked(t *testing.T) {
	var iface amplifybackend_iface.IClient
	iface = &amplifybackend.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := amplifybackend.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCloneBackend", func(t *testing.T) {
        input := &amplifybackend.CloneBackendInput{}
        output := &amplifybackend.CloneBackendOutput{}

        mockClient.On("CloneBackend", ctx, input).Return(output, nil)

        result, err := mockClient.CloneBackend(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateBackend", func(t *testing.T) {
        input := &amplifybackend.CreateBackendInput{}
        output := &amplifybackend.CreateBackendOutput{}

        mockClient.On("CreateBackend", ctx, input).Return(output, nil)

        result, err := mockClient.CreateBackend(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateBackendAPI", func(t *testing.T) {
        input := &amplifybackend.CreateBackendAPIInput{}
        output := &amplifybackend.CreateBackendAPIOutput{}

        mockClient.On("CreateBackendAPI", ctx, input).Return(output, nil)

        result, err := mockClient.CreateBackendAPI(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateBackendAuth", func(t *testing.T) {
        input := &amplifybackend.CreateBackendAuthInput{}
        output := &amplifybackend.CreateBackendAuthOutput{}

        mockClient.On("CreateBackendAuth", ctx, input).Return(output, nil)

        result, err := mockClient.CreateBackendAuth(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateBackendConfig", func(t *testing.T) {
        input := &amplifybackend.CreateBackendConfigInput{}
        output := &amplifybackend.CreateBackendConfigOutput{}

        mockClient.On("CreateBackendConfig", ctx, input).Return(output, nil)

        result, err := mockClient.CreateBackendConfig(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateBackendStorage", func(t *testing.T) {
        input := &amplifybackend.CreateBackendStorageInput{}
        output := &amplifybackend.CreateBackendStorageOutput{}

        mockClient.On("CreateBackendStorage", ctx, input).Return(output, nil)

        result, err := mockClient.CreateBackendStorage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateToken", func(t *testing.T) {
        input := &amplifybackend.CreateTokenInput{}
        output := &amplifybackend.CreateTokenOutput{}

        mockClient.On("CreateToken", ctx, input).Return(output, nil)

        result, err := mockClient.CreateToken(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteBackend", func(t *testing.T) {
        input := &amplifybackend.DeleteBackendInput{}
        output := &amplifybackend.DeleteBackendOutput{}

        mockClient.On("DeleteBackend", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteBackend(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteBackendAPI", func(t *testing.T) {
        input := &amplifybackend.DeleteBackendAPIInput{}
        output := &amplifybackend.DeleteBackendAPIOutput{}

        mockClient.On("DeleteBackendAPI", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteBackendAPI(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteBackendAuth", func(t *testing.T) {
        input := &amplifybackend.DeleteBackendAuthInput{}
        output := &amplifybackend.DeleteBackendAuthOutput{}

        mockClient.On("DeleteBackendAuth", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteBackendAuth(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteBackendStorage", func(t *testing.T) {
        input := &amplifybackend.DeleteBackendStorageInput{}
        output := &amplifybackend.DeleteBackendStorageOutput{}

        mockClient.On("DeleteBackendStorage", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteBackendStorage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteToken", func(t *testing.T) {
        input := &amplifybackend.DeleteTokenInput{}
        output := &amplifybackend.DeleteTokenOutput{}

        mockClient.On("DeleteToken", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteToken(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGenerateBackendAPIModels", func(t *testing.T) {
        input := &amplifybackend.GenerateBackendAPIModelsInput{}
        output := &amplifybackend.GenerateBackendAPIModelsOutput{}

        mockClient.On("GenerateBackendAPIModels", ctx, input).Return(output, nil)

        result, err := mockClient.GenerateBackendAPIModels(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetBackend", func(t *testing.T) {
        input := &amplifybackend.GetBackendInput{}
        output := &amplifybackend.GetBackendOutput{}

        mockClient.On("GetBackend", ctx, input).Return(output, nil)

        result, err := mockClient.GetBackend(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetBackendAPI", func(t *testing.T) {
        input := &amplifybackend.GetBackendAPIInput{}
        output := &amplifybackend.GetBackendAPIOutput{}

        mockClient.On("GetBackendAPI", ctx, input).Return(output, nil)

        result, err := mockClient.GetBackendAPI(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetBackendAPIModels", func(t *testing.T) {
        input := &amplifybackend.GetBackendAPIModelsInput{}
        output := &amplifybackend.GetBackendAPIModelsOutput{}

        mockClient.On("GetBackendAPIModels", ctx, input).Return(output, nil)

        result, err := mockClient.GetBackendAPIModels(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetBackendAuth", func(t *testing.T) {
        input := &amplifybackend.GetBackendAuthInput{}
        output := &amplifybackend.GetBackendAuthOutput{}

        mockClient.On("GetBackendAuth", ctx, input).Return(output, nil)

        result, err := mockClient.GetBackendAuth(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetBackendJob", func(t *testing.T) {
        input := &amplifybackend.GetBackendJobInput{}
        output := &amplifybackend.GetBackendJobOutput{}

        mockClient.On("GetBackendJob", ctx, input).Return(output, nil)

        result, err := mockClient.GetBackendJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetBackendStorage", func(t *testing.T) {
        input := &amplifybackend.GetBackendStorageInput{}
        output := &amplifybackend.GetBackendStorageOutput{}

        mockClient.On("GetBackendStorage", ctx, input).Return(output, nil)

        result, err := mockClient.GetBackendStorage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetToken", func(t *testing.T) {
        input := &amplifybackend.GetTokenInput{}
        output := &amplifybackend.GetTokenOutput{}

        mockClient.On("GetToken", ctx, input).Return(output, nil)

        result, err := mockClient.GetToken(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestImportBackendAuth", func(t *testing.T) {
        input := &amplifybackend.ImportBackendAuthInput{}
        output := &amplifybackend.ImportBackendAuthOutput{}

        mockClient.On("ImportBackendAuth", ctx, input).Return(output, nil)

        result, err := mockClient.ImportBackendAuth(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestImportBackendStorage", func(t *testing.T) {
        input := &amplifybackend.ImportBackendStorageInput{}
        output := &amplifybackend.ImportBackendStorageOutput{}

        mockClient.On("ImportBackendStorage", ctx, input).Return(output, nil)

        result, err := mockClient.ImportBackendStorage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListBackendJobs", func(t *testing.T) {
        input := &amplifybackend.ListBackendJobsInput{}
        output := &amplifybackend.ListBackendJobsOutput{}

        mockClient.On("ListBackendJobs", ctx, input).Return(output, nil)

        result, err := mockClient.ListBackendJobs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListS3Buckets", func(t *testing.T) {
        input := &amplifybackend.ListS3BucketsInput{}
        output := &amplifybackend.ListS3BucketsOutput{}

        mockClient.On("ListS3Buckets", ctx, input).Return(output, nil)

        result, err := mockClient.ListS3Buckets(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRemoveAllBackends", func(t *testing.T) {
        input := &amplifybackend.RemoveAllBackendsInput{}
        output := &amplifybackend.RemoveAllBackendsOutput{}

        mockClient.On("RemoveAllBackends", ctx, input).Return(output, nil)

        result, err := mockClient.RemoveAllBackends(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRemoveBackendConfig", func(t *testing.T) {
        input := &amplifybackend.RemoveBackendConfigInput{}
        output := &amplifybackend.RemoveBackendConfigOutput{}

        mockClient.On("RemoveBackendConfig", ctx, input).Return(output, nil)

        result, err := mockClient.RemoveBackendConfig(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateBackendAPI", func(t *testing.T) {
        input := &amplifybackend.UpdateBackendAPIInput{}
        output := &amplifybackend.UpdateBackendAPIOutput{}

        mockClient.On("UpdateBackendAPI", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateBackendAPI(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateBackendAuth", func(t *testing.T) {
        input := &amplifybackend.UpdateBackendAuthInput{}
        output := &amplifybackend.UpdateBackendAuthOutput{}

        mockClient.On("UpdateBackendAuth", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateBackendAuth(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateBackendConfig", func(t *testing.T) {
        input := &amplifybackend.UpdateBackendConfigInput{}
        output := &amplifybackend.UpdateBackendConfigOutput{}

        mockClient.On("UpdateBackendConfig", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateBackendConfig(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateBackendJob", func(t *testing.T) {
        input := &amplifybackend.UpdateBackendJobInput{}
        output := &amplifybackend.UpdateBackendJobOutput{}

        mockClient.On("UpdateBackendJob", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateBackendJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateBackendStorage", func(t *testing.T) {
        input := &amplifybackend.UpdateBackendStorageInput{}
        output := &amplifybackend.UpdateBackendStorageOutput{}

        mockClient.On("UpdateBackendStorage", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateBackendStorage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}