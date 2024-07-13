package kms_test

// tests for the kms service interface for this ../../../service/kms/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/kms"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/kms/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/kms/kms_iface"
	"github.com/stretchr/testify/assert"
)

func TestKmsServiceCanBeMocked(t *testing.T) {
	var iface kms_iface.IClient
	iface = &kms.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := kms.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCancelKeyDeletion", func(t *testing.T) {
        input := &kms.CancelKeyDeletionInput{}
        output := &kms.CancelKeyDeletionOutput{}

        mockClient.On("CancelKeyDeletion", ctx, input).Return(output, nil)

        result, err := mockClient.CancelKeyDeletion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestConnectCustomKeyStore", func(t *testing.T) {
        input := &kms.ConnectCustomKeyStoreInput{}
        output := &kms.ConnectCustomKeyStoreOutput{}

        mockClient.On("ConnectCustomKeyStore", ctx, input).Return(output, nil)

        result, err := mockClient.ConnectCustomKeyStore(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateAlias", func(t *testing.T) {
        input := &kms.CreateAliasInput{}
        output := &kms.CreateAliasOutput{}

        mockClient.On("CreateAlias", ctx, input).Return(output, nil)

        result, err := mockClient.CreateAlias(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateCustomKeyStore", func(t *testing.T) {
        input := &kms.CreateCustomKeyStoreInput{}
        output := &kms.CreateCustomKeyStoreOutput{}

        mockClient.On("CreateCustomKeyStore", ctx, input).Return(output, nil)

        result, err := mockClient.CreateCustomKeyStore(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateGrant", func(t *testing.T) {
        input := &kms.CreateGrantInput{}
        output := &kms.CreateGrantOutput{}

        mockClient.On("CreateGrant", ctx, input).Return(output, nil)

        result, err := mockClient.CreateGrant(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateKey", func(t *testing.T) {
        input := &kms.CreateKeyInput{}
        output := &kms.CreateKeyOutput{}

        mockClient.On("CreateKey", ctx, input).Return(output, nil)

        result, err := mockClient.CreateKey(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDecrypt", func(t *testing.T) {
        input := &kms.DecryptInput{}
        output := &kms.DecryptOutput{}

        mockClient.On("Decrypt", ctx, input).Return(output, nil)

        result, err := mockClient.Decrypt(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteAlias", func(t *testing.T) {
        input := &kms.DeleteAliasInput{}
        output := &kms.DeleteAliasOutput{}

        mockClient.On("DeleteAlias", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteAlias(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteCustomKeyStore", func(t *testing.T) {
        input := &kms.DeleteCustomKeyStoreInput{}
        output := &kms.DeleteCustomKeyStoreOutput{}

        mockClient.On("DeleteCustomKeyStore", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteCustomKeyStore(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteImportedKeyMaterial", func(t *testing.T) {
        input := &kms.DeleteImportedKeyMaterialInput{}
        output := &kms.DeleteImportedKeyMaterialOutput{}

        mockClient.On("DeleteImportedKeyMaterial", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteImportedKeyMaterial(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeriveSharedSecret", func(t *testing.T) {
        input := &kms.DeriveSharedSecretInput{}
        output := &kms.DeriveSharedSecretOutput{}

        mockClient.On("DeriveSharedSecret", ctx, input).Return(output, nil)

        result, err := mockClient.DeriveSharedSecret(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeCustomKeyStores", func(t *testing.T) {
        input := &kms.DescribeCustomKeyStoresInput{}
        output := &kms.DescribeCustomKeyStoresOutput{}

        mockClient.On("DescribeCustomKeyStores", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeCustomKeyStores(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeKey", func(t *testing.T) {
        input := &kms.DescribeKeyInput{}
        output := &kms.DescribeKeyOutput{}

        mockClient.On("DescribeKey", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeKey(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisableKey", func(t *testing.T) {
        input := &kms.DisableKeyInput{}
        output := &kms.DisableKeyOutput{}

        mockClient.On("DisableKey", ctx, input).Return(output, nil)

        result, err := mockClient.DisableKey(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisableKeyRotation", func(t *testing.T) {
        input := &kms.DisableKeyRotationInput{}
        output := &kms.DisableKeyRotationOutput{}

        mockClient.On("DisableKeyRotation", ctx, input).Return(output, nil)

        result, err := mockClient.DisableKeyRotation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisconnectCustomKeyStore", func(t *testing.T) {
        input := &kms.DisconnectCustomKeyStoreInput{}
        output := &kms.DisconnectCustomKeyStoreOutput{}

        mockClient.On("DisconnectCustomKeyStore", ctx, input).Return(output, nil)

        result, err := mockClient.DisconnectCustomKeyStore(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestEnableKey", func(t *testing.T) {
        input := &kms.EnableKeyInput{}
        output := &kms.EnableKeyOutput{}

        mockClient.On("EnableKey", ctx, input).Return(output, nil)

        result, err := mockClient.EnableKey(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestEnableKeyRotation", func(t *testing.T) {
        input := &kms.EnableKeyRotationInput{}
        output := &kms.EnableKeyRotationOutput{}

        mockClient.On("EnableKeyRotation", ctx, input).Return(output, nil)

        result, err := mockClient.EnableKeyRotation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestEncrypt", func(t *testing.T) {
        input := &kms.EncryptInput{}
        output := &kms.EncryptOutput{}

        mockClient.On("Encrypt", ctx, input).Return(output, nil)

        result, err := mockClient.Encrypt(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGenerateDataKey", func(t *testing.T) {
        input := &kms.GenerateDataKeyInput{}
        output := &kms.GenerateDataKeyOutput{}

        mockClient.On("GenerateDataKey", ctx, input).Return(output, nil)

        result, err := mockClient.GenerateDataKey(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGenerateDataKeyPair", func(t *testing.T) {
        input := &kms.GenerateDataKeyPairInput{}
        output := &kms.GenerateDataKeyPairOutput{}

        mockClient.On("GenerateDataKeyPair", ctx, input).Return(output, nil)

        result, err := mockClient.GenerateDataKeyPair(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGenerateDataKeyPairWithoutPlaintext", func(t *testing.T) {
        input := &kms.GenerateDataKeyPairWithoutPlaintextInput{}
        output := &kms.GenerateDataKeyPairWithoutPlaintextOutput{}

        mockClient.On("GenerateDataKeyPairWithoutPlaintext", ctx, input).Return(output, nil)

        result, err := mockClient.GenerateDataKeyPairWithoutPlaintext(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGenerateDataKeyWithoutPlaintext", func(t *testing.T) {
        input := &kms.GenerateDataKeyWithoutPlaintextInput{}
        output := &kms.GenerateDataKeyWithoutPlaintextOutput{}

        mockClient.On("GenerateDataKeyWithoutPlaintext", ctx, input).Return(output, nil)

        result, err := mockClient.GenerateDataKeyWithoutPlaintext(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGenerateMac", func(t *testing.T) {
        input := &kms.GenerateMacInput{}
        output := &kms.GenerateMacOutput{}

        mockClient.On("GenerateMac", ctx, input).Return(output, nil)

        result, err := mockClient.GenerateMac(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGenerateRandom", func(t *testing.T) {
        input := &kms.GenerateRandomInput{}
        output := &kms.GenerateRandomOutput{}

        mockClient.On("GenerateRandom", ctx, input).Return(output, nil)

        result, err := mockClient.GenerateRandom(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetKeyPolicy", func(t *testing.T) {
        input := &kms.GetKeyPolicyInput{}
        output := &kms.GetKeyPolicyOutput{}

        mockClient.On("GetKeyPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.GetKeyPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetKeyRotationStatus", func(t *testing.T) {
        input := &kms.GetKeyRotationStatusInput{}
        output := &kms.GetKeyRotationStatusOutput{}

        mockClient.On("GetKeyRotationStatus", ctx, input).Return(output, nil)

        result, err := mockClient.GetKeyRotationStatus(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetParametersForImport", func(t *testing.T) {
        input := &kms.GetParametersForImportInput{}
        output := &kms.GetParametersForImportOutput{}

        mockClient.On("GetParametersForImport", ctx, input).Return(output, nil)

        result, err := mockClient.GetParametersForImport(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetPublicKey", func(t *testing.T) {
        input := &kms.GetPublicKeyInput{}
        output := &kms.GetPublicKeyOutput{}

        mockClient.On("GetPublicKey", ctx, input).Return(output, nil)

        result, err := mockClient.GetPublicKey(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestImportKeyMaterial", func(t *testing.T) {
        input := &kms.ImportKeyMaterialInput{}
        output := &kms.ImportKeyMaterialOutput{}

        mockClient.On("ImportKeyMaterial", ctx, input).Return(output, nil)

        result, err := mockClient.ImportKeyMaterial(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAliases", func(t *testing.T) {
        input := &kms.ListAliasesInput{}
        output := &kms.ListAliasesOutput{}

        mockClient.On("ListAliases", ctx, input).Return(output, nil)

        result, err := mockClient.ListAliases(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListGrants", func(t *testing.T) {
        input := &kms.ListGrantsInput{}
        output := &kms.ListGrantsOutput{}

        mockClient.On("ListGrants", ctx, input).Return(output, nil)

        result, err := mockClient.ListGrants(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListKeyPolicies", func(t *testing.T) {
        input := &kms.ListKeyPoliciesInput{}
        output := &kms.ListKeyPoliciesOutput{}

        mockClient.On("ListKeyPolicies", ctx, input).Return(output, nil)

        result, err := mockClient.ListKeyPolicies(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListKeyRotations", func(t *testing.T) {
        input := &kms.ListKeyRotationsInput{}
        output := &kms.ListKeyRotationsOutput{}

        mockClient.On("ListKeyRotations", ctx, input).Return(output, nil)

        result, err := mockClient.ListKeyRotations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListKeys", func(t *testing.T) {
        input := &kms.ListKeysInput{}
        output := &kms.ListKeysOutput{}

        mockClient.On("ListKeys", ctx, input).Return(output, nil)

        result, err := mockClient.ListKeys(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListResourceTags", func(t *testing.T) {
        input := &kms.ListResourceTagsInput{}
        output := &kms.ListResourceTagsOutput{}

        mockClient.On("ListResourceTags", ctx, input).Return(output, nil)

        result, err := mockClient.ListResourceTags(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListRetirableGrants", func(t *testing.T) {
        input := &kms.ListRetirableGrantsInput{}
        output := &kms.ListRetirableGrantsOutput{}

        mockClient.On("ListRetirableGrants", ctx, input).Return(output, nil)

        result, err := mockClient.ListRetirableGrants(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutKeyPolicy", func(t *testing.T) {
        input := &kms.PutKeyPolicyInput{}
        output := &kms.PutKeyPolicyOutput{}

        mockClient.On("PutKeyPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.PutKeyPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestReEncrypt", func(t *testing.T) {
        input := &kms.ReEncryptInput{}
        output := &kms.ReEncryptOutput{}

        mockClient.On("ReEncrypt", ctx, input).Return(output, nil)

        result, err := mockClient.ReEncrypt(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestReplicateKey", func(t *testing.T) {
        input := &kms.ReplicateKeyInput{}
        output := &kms.ReplicateKeyOutput{}

        mockClient.On("ReplicateKey", ctx, input).Return(output, nil)

        result, err := mockClient.ReplicateKey(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRetireGrant", func(t *testing.T) {
        input := &kms.RetireGrantInput{}
        output := &kms.RetireGrantOutput{}

        mockClient.On("RetireGrant", ctx, input).Return(output, nil)

        result, err := mockClient.RetireGrant(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRevokeGrant", func(t *testing.T) {
        input := &kms.RevokeGrantInput{}
        output := &kms.RevokeGrantOutput{}

        mockClient.On("RevokeGrant", ctx, input).Return(output, nil)

        result, err := mockClient.RevokeGrant(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRotateKeyOnDemand", func(t *testing.T) {
        input := &kms.RotateKeyOnDemandInput{}
        output := &kms.RotateKeyOnDemandOutput{}

        mockClient.On("RotateKeyOnDemand", ctx, input).Return(output, nil)

        result, err := mockClient.RotateKeyOnDemand(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestScheduleKeyDeletion", func(t *testing.T) {
        input := &kms.ScheduleKeyDeletionInput{}
        output := &kms.ScheduleKeyDeletionOutput{}

        mockClient.On("ScheduleKeyDeletion", ctx, input).Return(output, nil)

        result, err := mockClient.ScheduleKeyDeletion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSign", func(t *testing.T) {
        input := &kms.SignInput{}
        output := &kms.SignOutput{}

        mockClient.On("Sign", ctx, input).Return(output, nil)

        result, err := mockClient.Sign(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &kms.TagResourceInput{}
        output := &kms.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &kms.UntagResourceInput{}
        output := &kms.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateAlias", func(t *testing.T) {
        input := &kms.UpdateAliasInput{}
        output := &kms.UpdateAliasOutput{}

        mockClient.On("UpdateAlias", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateAlias(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateCustomKeyStore", func(t *testing.T) {
        input := &kms.UpdateCustomKeyStoreInput{}
        output := &kms.UpdateCustomKeyStoreOutput{}

        mockClient.On("UpdateCustomKeyStore", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateCustomKeyStore(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateKeyDescription", func(t *testing.T) {
        input := &kms.UpdateKeyDescriptionInput{}
        output := &kms.UpdateKeyDescriptionOutput{}

        mockClient.On("UpdateKeyDescription", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateKeyDescription(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdatePrimaryRegion", func(t *testing.T) {
        input := &kms.UpdatePrimaryRegionInput{}
        output := &kms.UpdatePrimaryRegionOutput{}

        mockClient.On("UpdatePrimaryRegion", ctx, input).Return(output, nil)

        result, err := mockClient.UpdatePrimaryRegion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestVerify", func(t *testing.T) {
        input := &kms.VerifyInput{}
        output := &kms.VerifyOutput{}

        mockClient.On("Verify", ctx, input).Return(output, nil)

        result, err := mockClient.Verify(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestVerifyMac", func(t *testing.T) {
        input := &kms.VerifyMacInput{}
        output := &kms.VerifyMacOutput{}

        mockClient.On("VerifyMac", ctx, input).Return(output, nil)

        result, err := mockClient.VerifyMac(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
