package secretsmanager_test

// tests for the secretsmanager service interface for this ../../../service/secretsmanager/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/secretsmanager/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/secretsmanager/secretsmanager_iface"
	"github.com/stretchr/testify/assert"
)

func TestSecretsmanagerServiceCanBeMocked(t *testing.T) {
	var iface secretsmanager_iface.IClient
	iface = &secretsmanager.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := secretsmanager.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchGetSecretValue", func(t *testing.T) {
        input := &secretsmanager.BatchGetSecretValueInput{}
        output := &secretsmanager.BatchGetSecretValueOutput{}

        mockClient.On("BatchGetSecretValue", ctx, input).Return(output, nil)

        result, err := mockClient.BatchGetSecretValue(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCancelRotateSecret", func(t *testing.T) {
        input := &secretsmanager.CancelRotateSecretInput{}
        output := &secretsmanager.CancelRotateSecretOutput{}

        mockClient.On("CancelRotateSecret", ctx, input).Return(output, nil)

        result, err := mockClient.CancelRotateSecret(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateSecret", func(t *testing.T) {
        input := &secretsmanager.CreateSecretInput{}
        output := &secretsmanager.CreateSecretOutput{}

        mockClient.On("CreateSecret", ctx, input).Return(output, nil)

        result, err := mockClient.CreateSecret(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteResourcePolicy", func(t *testing.T) {
        input := &secretsmanager.DeleteResourcePolicyInput{}
        output := &secretsmanager.DeleteResourcePolicyOutput{}

        mockClient.On("DeleteResourcePolicy", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteResourcePolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteSecret", func(t *testing.T) {
        input := &secretsmanager.DeleteSecretInput{}
        output := &secretsmanager.DeleteSecretOutput{}

        mockClient.On("DeleteSecret", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteSecret(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeSecret", func(t *testing.T) {
        input := &secretsmanager.DescribeSecretInput{}
        output := &secretsmanager.DescribeSecretOutput{}

        mockClient.On("DescribeSecret", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeSecret(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetRandomPassword", func(t *testing.T) {
        input := &secretsmanager.GetRandomPasswordInput{}
        output := &secretsmanager.GetRandomPasswordOutput{}

        mockClient.On("GetRandomPassword", ctx, input).Return(output, nil)

        result, err := mockClient.GetRandomPassword(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetResourcePolicy", func(t *testing.T) {
        input := &secretsmanager.GetResourcePolicyInput{}
        output := &secretsmanager.GetResourcePolicyOutput{}

        mockClient.On("GetResourcePolicy", ctx, input).Return(output, nil)

        result, err := mockClient.GetResourcePolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSecretValue", func(t *testing.T) {
        input := &secretsmanager.GetSecretValueInput{}
        output := &secretsmanager.GetSecretValueOutput{}

        mockClient.On("GetSecretValue", ctx, input).Return(output, nil)

        result, err := mockClient.GetSecretValue(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListSecretVersionIds", func(t *testing.T) {
        input := &secretsmanager.ListSecretVersionIdsInput{}
        output := &secretsmanager.ListSecretVersionIdsOutput{}

        mockClient.On("ListSecretVersionIds", ctx, input).Return(output, nil)

        result, err := mockClient.ListSecretVersionIds(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListSecrets", func(t *testing.T) {
        input := &secretsmanager.ListSecretsInput{}
        output := &secretsmanager.ListSecretsOutput{}

        mockClient.On("ListSecrets", ctx, input).Return(output, nil)

        result, err := mockClient.ListSecrets(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutResourcePolicy", func(t *testing.T) {
        input := &secretsmanager.PutResourcePolicyInput{}
        output := &secretsmanager.PutResourcePolicyOutput{}

        mockClient.On("PutResourcePolicy", ctx, input).Return(output, nil)

        result, err := mockClient.PutResourcePolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutSecretValue", func(t *testing.T) {
        input := &secretsmanager.PutSecretValueInput{}
        output := &secretsmanager.PutSecretValueOutput{}

        mockClient.On("PutSecretValue", ctx, input).Return(output, nil)

        result, err := mockClient.PutSecretValue(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRemoveRegionsFromReplication", func(t *testing.T) {
        input := &secretsmanager.RemoveRegionsFromReplicationInput{}
        output := &secretsmanager.RemoveRegionsFromReplicationOutput{}

        mockClient.On("RemoveRegionsFromReplication", ctx, input).Return(output, nil)

        result, err := mockClient.RemoveRegionsFromReplication(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestReplicateSecretToRegions", func(t *testing.T) {
        input := &secretsmanager.ReplicateSecretToRegionsInput{}
        output := &secretsmanager.ReplicateSecretToRegionsOutput{}

        mockClient.On("ReplicateSecretToRegions", ctx, input).Return(output, nil)

        result, err := mockClient.ReplicateSecretToRegions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRestoreSecret", func(t *testing.T) {
        input := &secretsmanager.RestoreSecretInput{}
        output := &secretsmanager.RestoreSecretOutput{}

        mockClient.On("RestoreSecret", ctx, input).Return(output, nil)

        result, err := mockClient.RestoreSecret(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRotateSecret", func(t *testing.T) {
        input := &secretsmanager.RotateSecretInput{}
        output := &secretsmanager.RotateSecretOutput{}

        mockClient.On("RotateSecret", ctx, input).Return(output, nil)

        result, err := mockClient.RotateSecret(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopReplicationToReplica", func(t *testing.T) {
        input := &secretsmanager.StopReplicationToReplicaInput{}
        output := &secretsmanager.StopReplicationToReplicaOutput{}

        mockClient.On("StopReplicationToReplica", ctx, input).Return(output, nil)

        result, err := mockClient.StopReplicationToReplica(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &secretsmanager.TagResourceInput{}
        output := &secretsmanager.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &secretsmanager.UntagResourceInput{}
        output := &secretsmanager.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateSecret", func(t *testing.T) {
        input := &secretsmanager.UpdateSecretInput{}
        output := &secretsmanager.UpdateSecretOutput{}

        mockClient.On("UpdateSecret", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateSecret(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateSecretVersionStage", func(t *testing.T) {
        input := &secretsmanager.UpdateSecretVersionStageInput{}
        output := &secretsmanager.UpdateSecretVersionStageOutput{}

        mockClient.On("UpdateSecretVersionStage", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateSecretVersionStage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestValidateResourcePolicy", func(t *testing.T) {
        input := &secretsmanager.ValidateResourcePolicyInput{}
        output := &secretsmanager.ValidateResourcePolicyOutput{}

        mockClient.On("ValidateResourcePolicy", ctx, input).Return(output, nil)

        result, err := mockClient.ValidateResourcePolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
