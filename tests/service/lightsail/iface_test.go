package lightsail_test

// tests for the lightsail service interface for this ../../../service/lightsail/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/lightsail"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/lightsail/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/lightsail/lightsail_iface"
	"github.com/stretchr/testify/assert"
)

func TestLightsailServiceCanBeMocked(t *testing.T) {
	var iface lightsail_iface.IClient
	iface = &lightsail.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := lightsail.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAllocateStaticIp", func(t *testing.T) {
        input := &lightsail.AllocateStaticIpInput{}
        output := &lightsail.AllocateStaticIpOutput{}

        mockClient.On("AllocateStaticIp", ctx, input).Return(output, nil)

        result, err := mockClient.AllocateStaticIp(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAttachCertificateToDistribution", func(t *testing.T) {
        input := &lightsail.AttachCertificateToDistributionInput{}
        output := &lightsail.AttachCertificateToDistributionOutput{}

        mockClient.On("AttachCertificateToDistribution", ctx, input).Return(output, nil)

        result, err := mockClient.AttachCertificateToDistribution(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAttachDisk", func(t *testing.T) {
        input := &lightsail.AttachDiskInput{}
        output := &lightsail.AttachDiskOutput{}

        mockClient.On("AttachDisk", ctx, input).Return(output, nil)

        result, err := mockClient.AttachDisk(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAttachInstancesToLoadBalancer", func(t *testing.T) {
        input := &lightsail.AttachInstancesToLoadBalancerInput{}
        output := &lightsail.AttachInstancesToLoadBalancerOutput{}

        mockClient.On("AttachInstancesToLoadBalancer", ctx, input).Return(output, nil)

        result, err := mockClient.AttachInstancesToLoadBalancer(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAttachLoadBalancerTlsCertificate", func(t *testing.T) {
        input := &lightsail.AttachLoadBalancerTlsCertificateInput{}
        output := &lightsail.AttachLoadBalancerTlsCertificateOutput{}

        mockClient.On("AttachLoadBalancerTlsCertificate", ctx, input).Return(output, nil)

        result, err := mockClient.AttachLoadBalancerTlsCertificate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAttachStaticIp", func(t *testing.T) {
        input := &lightsail.AttachStaticIpInput{}
        output := &lightsail.AttachStaticIpOutput{}

        mockClient.On("AttachStaticIp", ctx, input).Return(output, nil)

        result, err := mockClient.AttachStaticIp(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCloseInstancePublicPorts", func(t *testing.T) {
        input := &lightsail.CloseInstancePublicPortsInput{}
        output := &lightsail.CloseInstancePublicPortsOutput{}

        mockClient.On("CloseInstancePublicPorts", ctx, input).Return(output, nil)

        result, err := mockClient.CloseInstancePublicPorts(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCopySnapshot", func(t *testing.T) {
        input := &lightsail.CopySnapshotInput{}
        output := &lightsail.CopySnapshotOutput{}

        mockClient.On("CopySnapshot", ctx, input).Return(output, nil)

        result, err := mockClient.CopySnapshot(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateBucket", func(t *testing.T) {
        input := &lightsail.CreateBucketInput{}
        output := &lightsail.CreateBucketOutput{}

        mockClient.On("CreateBucket", ctx, input).Return(output, nil)

        result, err := mockClient.CreateBucket(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateBucketAccessKey", func(t *testing.T) {
        input := &lightsail.CreateBucketAccessKeyInput{}
        output := &lightsail.CreateBucketAccessKeyOutput{}

        mockClient.On("CreateBucketAccessKey", ctx, input).Return(output, nil)

        result, err := mockClient.CreateBucketAccessKey(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateCertificate", func(t *testing.T) {
        input := &lightsail.CreateCertificateInput{}
        output := &lightsail.CreateCertificateOutput{}

        mockClient.On("CreateCertificate", ctx, input).Return(output, nil)

        result, err := mockClient.CreateCertificate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateCloudFormationStack", func(t *testing.T) {
        input := &lightsail.CreateCloudFormationStackInput{}
        output := &lightsail.CreateCloudFormationStackOutput{}

        mockClient.On("CreateCloudFormationStack", ctx, input).Return(output, nil)

        result, err := mockClient.CreateCloudFormationStack(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateContactMethod", func(t *testing.T) {
        input := &lightsail.CreateContactMethodInput{}
        output := &lightsail.CreateContactMethodOutput{}

        mockClient.On("CreateContactMethod", ctx, input).Return(output, nil)

        result, err := mockClient.CreateContactMethod(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateContainerService", func(t *testing.T) {
        input := &lightsail.CreateContainerServiceInput{}
        output := &lightsail.CreateContainerServiceOutput{}

        mockClient.On("CreateContainerService", ctx, input).Return(output, nil)

        result, err := mockClient.CreateContainerService(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateContainerServiceDeployment", func(t *testing.T) {
        input := &lightsail.CreateContainerServiceDeploymentInput{}
        output := &lightsail.CreateContainerServiceDeploymentOutput{}

        mockClient.On("CreateContainerServiceDeployment", ctx, input).Return(output, nil)

        result, err := mockClient.CreateContainerServiceDeployment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateContainerServiceRegistryLogin", func(t *testing.T) {
        input := &lightsail.CreateContainerServiceRegistryLoginInput{}
        output := &lightsail.CreateContainerServiceRegistryLoginOutput{}

        mockClient.On("CreateContainerServiceRegistryLogin", ctx, input).Return(output, nil)

        result, err := mockClient.CreateContainerServiceRegistryLogin(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateDisk", func(t *testing.T) {
        input := &lightsail.CreateDiskInput{}
        output := &lightsail.CreateDiskOutput{}

        mockClient.On("CreateDisk", ctx, input).Return(output, nil)

        result, err := mockClient.CreateDisk(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateDiskFromSnapshot", func(t *testing.T) {
        input := &lightsail.CreateDiskFromSnapshotInput{}
        output := &lightsail.CreateDiskFromSnapshotOutput{}

        mockClient.On("CreateDiskFromSnapshot", ctx, input).Return(output, nil)

        result, err := mockClient.CreateDiskFromSnapshot(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateDiskSnapshot", func(t *testing.T) {
        input := &lightsail.CreateDiskSnapshotInput{}
        output := &lightsail.CreateDiskSnapshotOutput{}

        mockClient.On("CreateDiskSnapshot", ctx, input).Return(output, nil)

        result, err := mockClient.CreateDiskSnapshot(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateDistribution", func(t *testing.T) {
        input := &lightsail.CreateDistributionInput{}
        output := &lightsail.CreateDistributionOutput{}

        mockClient.On("CreateDistribution", ctx, input).Return(output, nil)

        result, err := mockClient.CreateDistribution(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateDomain", func(t *testing.T) {
        input := &lightsail.CreateDomainInput{}
        output := &lightsail.CreateDomainOutput{}

        mockClient.On("CreateDomain", ctx, input).Return(output, nil)

        result, err := mockClient.CreateDomain(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateDomainEntry", func(t *testing.T) {
        input := &lightsail.CreateDomainEntryInput{}
        output := &lightsail.CreateDomainEntryOutput{}

        mockClient.On("CreateDomainEntry", ctx, input).Return(output, nil)

        result, err := mockClient.CreateDomainEntry(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateGUISessionAccessDetails", func(t *testing.T) {
        input := &lightsail.CreateGUISessionAccessDetailsInput{}
        output := &lightsail.CreateGUISessionAccessDetailsOutput{}

        mockClient.On("CreateGUISessionAccessDetails", ctx, input).Return(output, nil)

        result, err := mockClient.CreateGUISessionAccessDetails(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateInstanceSnapshot", func(t *testing.T) {
        input := &lightsail.CreateInstanceSnapshotInput{}
        output := &lightsail.CreateInstanceSnapshotOutput{}

        mockClient.On("CreateInstanceSnapshot", ctx, input).Return(output, nil)

        result, err := mockClient.CreateInstanceSnapshot(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateInstances", func(t *testing.T) {
        input := &lightsail.CreateInstancesInput{}
        output := &lightsail.CreateInstancesOutput{}

        mockClient.On("CreateInstances", ctx, input).Return(output, nil)

        result, err := mockClient.CreateInstances(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateInstancesFromSnapshot", func(t *testing.T) {
        input := &lightsail.CreateInstancesFromSnapshotInput{}
        output := &lightsail.CreateInstancesFromSnapshotOutput{}

        mockClient.On("CreateInstancesFromSnapshot", ctx, input).Return(output, nil)

        result, err := mockClient.CreateInstancesFromSnapshot(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateKeyPair", func(t *testing.T) {
        input := &lightsail.CreateKeyPairInput{}
        output := &lightsail.CreateKeyPairOutput{}

        mockClient.On("CreateKeyPair", ctx, input).Return(output, nil)

        result, err := mockClient.CreateKeyPair(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateLoadBalancer", func(t *testing.T) {
        input := &lightsail.CreateLoadBalancerInput{}
        output := &lightsail.CreateLoadBalancerOutput{}

        mockClient.On("CreateLoadBalancer", ctx, input).Return(output, nil)

        result, err := mockClient.CreateLoadBalancer(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateLoadBalancerTlsCertificate", func(t *testing.T) {
        input := &lightsail.CreateLoadBalancerTlsCertificateInput{}
        output := &lightsail.CreateLoadBalancerTlsCertificateOutput{}

        mockClient.On("CreateLoadBalancerTlsCertificate", ctx, input).Return(output, nil)

        result, err := mockClient.CreateLoadBalancerTlsCertificate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateRelationalDatabase", func(t *testing.T) {
        input := &lightsail.CreateRelationalDatabaseInput{}
        output := &lightsail.CreateRelationalDatabaseOutput{}

        mockClient.On("CreateRelationalDatabase", ctx, input).Return(output, nil)

        result, err := mockClient.CreateRelationalDatabase(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateRelationalDatabaseFromSnapshot", func(t *testing.T) {
        input := &lightsail.CreateRelationalDatabaseFromSnapshotInput{}
        output := &lightsail.CreateRelationalDatabaseFromSnapshotOutput{}

        mockClient.On("CreateRelationalDatabaseFromSnapshot", ctx, input).Return(output, nil)

        result, err := mockClient.CreateRelationalDatabaseFromSnapshot(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateRelationalDatabaseSnapshot", func(t *testing.T) {
        input := &lightsail.CreateRelationalDatabaseSnapshotInput{}
        output := &lightsail.CreateRelationalDatabaseSnapshotOutput{}

        mockClient.On("CreateRelationalDatabaseSnapshot", ctx, input).Return(output, nil)

        result, err := mockClient.CreateRelationalDatabaseSnapshot(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteAlarm", func(t *testing.T) {
        input := &lightsail.DeleteAlarmInput{}
        output := &lightsail.DeleteAlarmOutput{}

        mockClient.On("DeleteAlarm", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteAlarm(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteAutoSnapshot", func(t *testing.T) {
        input := &lightsail.DeleteAutoSnapshotInput{}
        output := &lightsail.DeleteAutoSnapshotOutput{}

        mockClient.On("DeleteAutoSnapshot", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteAutoSnapshot(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteBucket", func(t *testing.T) {
        input := &lightsail.DeleteBucketInput{}
        output := &lightsail.DeleteBucketOutput{}

        mockClient.On("DeleteBucket", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteBucket(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteBucketAccessKey", func(t *testing.T) {
        input := &lightsail.DeleteBucketAccessKeyInput{}
        output := &lightsail.DeleteBucketAccessKeyOutput{}

        mockClient.On("DeleteBucketAccessKey", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteBucketAccessKey(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteCertificate", func(t *testing.T) {
        input := &lightsail.DeleteCertificateInput{}
        output := &lightsail.DeleteCertificateOutput{}

        mockClient.On("DeleteCertificate", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteCertificate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteContactMethod", func(t *testing.T) {
        input := &lightsail.DeleteContactMethodInput{}
        output := &lightsail.DeleteContactMethodOutput{}

        mockClient.On("DeleteContactMethod", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteContactMethod(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteContainerImage", func(t *testing.T) {
        input := &lightsail.DeleteContainerImageInput{}
        output := &lightsail.DeleteContainerImageOutput{}

        mockClient.On("DeleteContainerImage", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteContainerImage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteContainerService", func(t *testing.T) {
        input := &lightsail.DeleteContainerServiceInput{}
        output := &lightsail.DeleteContainerServiceOutput{}

        mockClient.On("DeleteContainerService", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteContainerService(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteDisk", func(t *testing.T) {
        input := &lightsail.DeleteDiskInput{}
        output := &lightsail.DeleteDiskOutput{}

        mockClient.On("DeleteDisk", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteDisk(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteDiskSnapshot", func(t *testing.T) {
        input := &lightsail.DeleteDiskSnapshotInput{}
        output := &lightsail.DeleteDiskSnapshotOutput{}

        mockClient.On("DeleteDiskSnapshot", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteDiskSnapshot(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteDistribution", func(t *testing.T) {
        input := &lightsail.DeleteDistributionInput{}
        output := &lightsail.DeleteDistributionOutput{}

        mockClient.On("DeleteDistribution", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteDistribution(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteDomain", func(t *testing.T) {
        input := &lightsail.DeleteDomainInput{}
        output := &lightsail.DeleteDomainOutput{}

        mockClient.On("DeleteDomain", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteDomain(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteDomainEntry", func(t *testing.T) {
        input := &lightsail.DeleteDomainEntryInput{}
        output := &lightsail.DeleteDomainEntryOutput{}

        mockClient.On("DeleteDomainEntry", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteDomainEntry(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteInstance", func(t *testing.T) {
        input := &lightsail.DeleteInstanceInput{}
        output := &lightsail.DeleteInstanceOutput{}

        mockClient.On("DeleteInstance", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteInstance(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteInstanceSnapshot", func(t *testing.T) {
        input := &lightsail.DeleteInstanceSnapshotInput{}
        output := &lightsail.DeleteInstanceSnapshotOutput{}

        mockClient.On("DeleteInstanceSnapshot", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteInstanceSnapshot(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteKeyPair", func(t *testing.T) {
        input := &lightsail.DeleteKeyPairInput{}
        output := &lightsail.DeleteKeyPairOutput{}

        mockClient.On("DeleteKeyPair", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteKeyPair(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteKnownHostKeys", func(t *testing.T) {
        input := &lightsail.DeleteKnownHostKeysInput{}
        output := &lightsail.DeleteKnownHostKeysOutput{}

        mockClient.On("DeleteKnownHostKeys", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteKnownHostKeys(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteLoadBalancer", func(t *testing.T) {
        input := &lightsail.DeleteLoadBalancerInput{}
        output := &lightsail.DeleteLoadBalancerOutput{}

        mockClient.On("DeleteLoadBalancer", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteLoadBalancer(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteLoadBalancerTlsCertificate", func(t *testing.T) {
        input := &lightsail.DeleteLoadBalancerTlsCertificateInput{}
        output := &lightsail.DeleteLoadBalancerTlsCertificateOutput{}

        mockClient.On("DeleteLoadBalancerTlsCertificate", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteLoadBalancerTlsCertificate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteRelationalDatabase", func(t *testing.T) {
        input := &lightsail.DeleteRelationalDatabaseInput{}
        output := &lightsail.DeleteRelationalDatabaseOutput{}

        mockClient.On("DeleteRelationalDatabase", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteRelationalDatabase(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteRelationalDatabaseSnapshot", func(t *testing.T) {
        input := &lightsail.DeleteRelationalDatabaseSnapshotInput{}
        output := &lightsail.DeleteRelationalDatabaseSnapshotOutput{}

        mockClient.On("DeleteRelationalDatabaseSnapshot", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteRelationalDatabaseSnapshot(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDetachCertificateFromDistribution", func(t *testing.T) {
        input := &lightsail.DetachCertificateFromDistributionInput{}
        output := &lightsail.DetachCertificateFromDistributionOutput{}

        mockClient.On("DetachCertificateFromDistribution", ctx, input).Return(output, nil)

        result, err := mockClient.DetachCertificateFromDistribution(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDetachDisk", func(t *testing.T) {
        input := &lightsail.DetachDiskInput{}
        output := &lightsail.DetachDiskOutput{}

        mockClient.On("DetachDisk", ctx, input).Return(output, nil)

        result, err := mockClient.DetachDisk(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDetachInstancesFromLoadBalancer", func(t *testing.T) {
        input := &lightsail.DetachInstancesFromLoadBalancerInput{}
        output := &lightsail.DetachInstancesFromLoadBalancerOutput{}

        mockClient.On("DetachInstancesFromLoadBalancer", ctx, input).Return(output, nil)

        result, err := mockClient.DetachInstancesFromLoadBalancer(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDetachStaticIp", func(t *testing.T) {
        input := &lightsail.DetachStaticIpInput{}
        output := &lightsail.DetachStaticIpOutput{}

        mockClient.On("DetachStaticIp", ctx, input).Return(output, nil)

        result, err := mockClient.DetachStaticIp(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisableAddOn", func(t *testing.T) {
        input := &lightsail.DisableAddOnInput{}
        output := &lightsail.DisableAddOnOutput{}

        mockClient.On("DisableAddOn", ctx, input).Return(output, nil)

        result, err := mockClient.DisableAddOn(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDownloadDefaultKeyPair", func(t *testing.T) {
        input := &lightsail.DownloadDefaultKeyPairInput{}
        output := &lightsail.DownloadDefaultKeyPairOutput{}

        mockClient.On("DownloadDefaultKeyPair", ctx, input).Return(output, nil)

        result, err := mockClient.DownloadDefaultKeyPair(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestEnableAddOn", func(t *testing.T) {
        input := &lightsail.EnableAddOnInput{}
        output := &lightsail.EnableAddOnOutput{}

        mockClient.On("EnableAddOn", ctx, input).Return(output, nil)

        result, err := mockClient.EnableAddOn(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestExportSnapshot", func(t *testing.T) {
        input := &lightsail.ExportSnapshotInput{}
        output := &lightsail.ExportSnapshotOutput{}

        mockClient.On("ExportSnapshot", ctx, input).Return(output, nil)

        result, err := mockClient.ExportSnapshot(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetActiveNames", func(t *testing.T) {
        input := &lightsail.GetActiveNamesInput{}
        output := &lightsail.GetActiveNamesOutput{}

        mockClient.On("GetActiveNames", ctx, input).Return(output, nil)

        result, err := mockClient.GetActiveNames(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetAlarms", func(t *testing.T) {
        input := &lightsail.GetAlarmsInput{}
        output := &lightsail.GetAlarmsOutput{}

        mockClient.On("GetAlarms", ctx, input).Return(output, nil)

        result, err := mockClient.GetAlarms(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetAutoSnapshots", func(t *testing.T) {
        input := &lightsail.GetAutoSnapshotsInput{}
        output := &lightsail.GetAutoSnapshotsOutput{}

        mockClient.On("GetAutoSnapshots", ctx, input).Return(output, nil)

        result, err := mockClient.GetAutoSnapshots(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetBlueprints", func(t *testing.T) {
        input := &lightsail.GetBlueprintsInput{}
        output := &lightsail.GetBlueprintsOutput{}

        mockClient.On("GetBlueprints", ctx, input).Return(output, nil)

        result, err := mockClient.GetBlueprints(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetBucketAccessKeys", func(t *testing.T) {
        input := &lightsail.GetBucketAccessKeysInput{}
        output := &lightsail.GetBucketAccessKeysOutput{}

        mockClient.On("GetBucketAccessKeys", ctx, input).Return(output, nil)

        result, err := mockClient.GetBucketAccessKeys(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetBucketBundles", func(t *testing.T) {
        input := &lightsail.GetBucketBundlesInput{}
        output := &lightsail.GetBucketBundlesOutput{}

        mockClient.On("GetBucketBundles", ctx, input).Return(output, nil)

        result, err := mockClient.GetBucketBundles(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetBucketMetricData", func(t *testing.T) {
        input := &lightsail.GetBucketMetricDataInput{}
        output := &lightsail.GetBucketMetricDataOutput{}

        mockClient.On("GetBucketMetricData", ctx, input).Return(output, nil)

        result, err := mockClient.GetBucketMetricData(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetBuckets", func(t *testing.T) {
        input := &lightsail.GetBucketsInput{}
        output := &lightsail.GetBucketsOutput{}

        mockClient.On("GetBuckets", ctx, input).Return(output, nil)

        result, err := mockClient.GetBuckets(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetBundles", func(t *testing.T) {
        input := &lightsail.GetBundlesInput{}
        output := &lightsail.GetBundlesOutput{}

        mockClient.On("GetBundles", ctx, input).Return(output, nil)

        result, err := mockClient.GetBundles(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetCertificates", func(t *testing.T) {
        input := &lightsail.GetCertificatesInput{}
        output := &lightsail.GetCertificatesOutput{}

        mockClient.On("GetCertificates", ctx, input).Return(output, nil)

        result, err := mockClient.GetCertificates(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetCloudFormationStackRecords", func(t *testing.T) {
        input := &lightsail.GetCloudFormationStackRecordsInput{}
        output := &lightsail.GetCloudFormationStackRecordsOutput{}

        mockClient.On("GetCloudFormationStackRecords", ctx, input).Return(output, nil)

        result, err := mockClient.GetCloudFormationStackRecords(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetContactMethods", func(t *testing.T) {
        input := &lightsail.GetContactMethodsInput{}
        output := &lightsail.GetContactMethodsOutput{}

        mockClient.On("GetContactMethods", ctx, input).Return(output, nil)

        result, err := mockClient.GetContactMethods(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetContainerAPIMetadata", func(t *testing.T) {
        input := &lightsail.GetContainerAPIMetadataInput{}
        output := &lightsail.GetContainerAPIMetadataOutput{}

        mockClient.On("GetContainerAPIMetadata", ctx, input).Return(output, nil)

        result, err := mockClient.GetContainerAPIMetadata(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetContainerImages", func(t *testing.T) {
        input := &lightsail.GetContainerImagesInput{}
        output := &lightsail.GetContainerImagesOutput{}

        mockClient.On("GetContainerImages", ctx, input).Return(output, nil)

        result, err := mockClient.GetContainerImages(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetContainerLog", func(t *testing.T) {
        input := &lightsail.GetContainerLogInput{}
        output := &lightsail.GetContainerLogOutput{}

        mockClient.On("GetContainerLog", ctx, input).Return(output, nil)

        result, err := mockClient.GetContainerLog(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetContainerServiceDeployments", func(t *testing.T) {
        input := &lightsail.GetContainerServiceDeploymentsInput{}
        output := &lightsail.GetContainerServiceDeploymentsOutput{}

        mockClient.On("GetContainerServiceDeployments", ctx, input).Return(output, nil)

        result, err := mockClient.GetContainerServiceDeployments(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetContainerServiceMetricData", func(t *testing.T) {
        input := &lightsail.GetContainerServiceMetricDataInput{}
        output := &lightsail.GetContainerServiceMetricDataOutput{}

        mockClient.On("GetContainerServiceMetricData", ctx, input).Return(output, nil)

        result, err := mockClient.GetContainerServiceMetricData(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetContainerServicePowers", func(t *testing.T) {
        input := &lightsail.GetContainerServicePowersInput{}
        output := &lightsail.GetContainerServicePowersOutput{}

        mockClient.On("GetContainerServicePowers", ctx, input).Return(output, nil)

        result, err := mockClient.GetContainerServicePowers(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetContainerServices", func(t *testing.T) {
        input := &lightsail.GetContainerServicesInput{}
        output := &lightsail.GetContainerServicesOutput{}

        mockClient.On("GetContainerServices", ctx, input).Return(output, nil)

        result, err := mockClient.GetContainerServices(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetCostEstimate", func(t *testing.T) {
        input := &lightsail.GetCostEstimateInput{}
        output := &lightsail.GetCostEstimateOutput{}

        mockClient.On("GetCostEstimate", ctx, input).Return(output, nil)

        result, err := mockClient.GetCostEstimate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDisk", func(t *testing.T) {
        input := &lightsail.GetDiskInput{}
        output := &lightsail.GetDiskOutput{}

        mockClient.On("GetDisk", ctx, input).Return(output, nil)

        result, err := mockClient.GetDisk(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDiskSnapshot", func(t *testing.T) {
        input := &lightsail.GetDiskSnapshotInput{}
        output := &lightsail.GetDiskSnapshotOutput{}

        mockClient.On("GetDiskSnapshot", ctx, input).Return(output, nil)

        result, err := mockClient.GetDiskSnapshot(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDiskSnapshots", func(t *testing.T) {
        input := &lightsail.GetDiskSnapshotsInput{}
        output := &lightsail.GetDiskSnapshotsOutput{}

        mockClient.On("GetDiskSnapshots", ctx, input).Return(output, nil)

        result, err := mockClient.GetDiskSnapshots(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDisks", func(t *testing.T) {
        input := &lightsail.GetDisksInput{}
        output := &lightsail.GetDisksOutput{}

        mockClient.On("GetDisks", ctx, input).Return(output, nil)

        result, err := mockClient.GetDisks(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDistributionBundles", func(t *testing.T) {
        input := &lightsail.GetDistributionBundlesInput{}
        output := &lightsail.GetDistributionBundlesOutput{}

        mockClient.On("GetDistributionBundles", ctx, input).Return(output, nil)

        result, err := mockClient.GetDistributionBundles(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDistributionLatestCacheReset", func(t *testing.T) {
        input := &lightsail.GetDistributionLatestCacheResetInput{}
        output := &lightsail.GetDistributionLatestCacheResetOutput{}

        mockClient.On("GetDistributionLatestCacheReset", ctx, input).Return(output, nil)

        result, err := mockClient.GetDistributionLatestCacheReset(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDistributionMetricData", func(t *testing.T) {
        input := &lightsail.GetDistributionMetricDataInput{}
        output := &lightsail.GetDistributionMetricDataOutput{}

        mockClient.On("GetDistributionMetricData", ctx, input).Return(output, nil)

        result, err := mockClient.GetDistributionMetricData(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDistributions", func(t *testing.T) {
        input := &lightsail.GetDistributionsInput{}
        output := &lightsail.GetDistributionsOutput{}

        mockClient.On("GetDistributions", ctx, input).Return(output, nil)

        result, err := mockClient.GetDistributions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDomain", func(t *testing.T) {
        input := &lightsail.GetDomainInput{}
        output := &lightsail.GetDomainOutput{}

        mockClient.On("GetDomain", ctx, input).Return(output, nil)

        result, err := mockClient.GetDomain(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDomains", func(t *testing.T) {
        input := &lightsail.GetDomainsInput{}
        output := &lightsail.GetDomainsOutput{}

        mockClient.On("GetDomains", ctx, input).Return(output, nil)

        result, err := mockClient.GetDomains(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetExportSnapshotRecords", func(t *testing.T) {
        input := &lightsail.GetExportSnapshotRecordsInput{}
        output := &lightsail.GetExportSnapshotRecordsOutput{}

        mockClient.On("GetExportSnapshotRecords", ctx, input).Return(output, nil)

        result, err := mockClient.GetExportSnapshotRecords(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetInstance", func(t *testing.T) {
        input := &lightsail.GetInstanceInput{}
        output := &lightsail.GetInstanceOutput{}

        mockClient.On("GetInstance", ctx, input).Return(output, nil)

        result, err := mockClient.GetInstance(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetInstanceAccessDetails", func(t *testing.T) {
        input := &lightsail.GetInstanceAccessDetailsInput{}
        output := &lightsail.GetInstanceAccessDetailsOutput{}

        mockClient.On("GetInstanceAccessDetails", ctx, input).Return(output, nil)

        result, err := mockClient.GetInstanceAccessDetails(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetInstanceMetricData", func(t *testing.T) {
        input := &lightsail.GetInstanceMetricDataInput{}
        output := &lightsail.GetInstanceMetricDataOutput{}

        mockClient.On("GetInstanceMetricData", ctx, input).Return(output, nil)

        result, err := mockClient.GetInstanceMetricData(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetInstancePortStates", func(t *testing.T) {
        input := &lightsail.GetInstancePortStatesInput{}
        output := &lightsail.GetInstancePortStatesOutput{}

        mockClient.On("GetInstancePortStates", ctx, input).Return(output, nil)

        result, err := mockClient.GetInstancePortStates(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetInstanceSnapshot", func(t *testing.T) {
        input := &lightsail.GetInstanceSnapshotInput{}
        output := &lightsail.GetInstanceSnapshotOutput{}

        mockClient.On("GetInstanceSnapshot", ctx, input).Return(output, nil)

        result, err := mockClient.GetInstanceSnapshot(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetInstanceSnapshots", func(t *testing.T) {
        input := &lightsail.GetInstanceSnapshotsInput{}
        output := &lightsail.GetInstanceSnapshotsOutput{}

        mockClient.On("GetInstanceSnapshots", ctx, input).Return(output, nil)

        result, err := mockClient.GetInstanceSnapshots(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetInstanceState", func(t *testing.T) {
        input := &lightsail.GetInstanceStateInput{}
        output := &lightsail.GetInstanceStateOutput{}

        mockClient.On("GetInstanceState", ctx, input).Return(output, nil)

        result, err := mockClient.GetInstanceState(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetInstances", func(t *testing.T) {
        input := &lightsail.GetInstancesInput{}
        output := &lightsail.GetInstancesOutput{}

        mockClient.On("GetInstances", ctx, input).Return(output, nil)

        result, err := mockClient.GetInstances(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetKeyPair", func(t *testing.T) {
        input := &lightsail.GetKeyPairInput{}
        output := &lightsail.GetKeyPairOutput{}

        mockClient.On("GetKeyPair", ctx, input).Return(output, nil)

        result, err := mockClient.GetKeyPair(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetKeyPairs", func(t *testing.T) {
        input := &lightsail.GetKeyPairsInput{}
        output := &lightsail.GetKeyPairsOutput{}

        mockClient.On("GetKeyPairs", ctx, input).Return(output, nil)

        result, err := mockClient.GetKeyPairs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetLoadBalancer", func(t *testing.T) {
        input := &lightsail.GetLoadBalancerInput{}
        output := &lightsail.GetLoadBalancerOutput{}

        mockClient.On("GetLoadBalancer", ctx, input).Return(output, nil)

        result, err := mockClient.GetLoadBalancer(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetLoadBalancerMetricData", func(t *testing.T) {
        input := &lightsail.GetLoadBalancerMetricDataInput{}
        output := &lightsail.GetLoadBalancerMetricDataOutput{}

        mockClient.On("GetLoadBalancerMetricData", ctx, input).Return(output, nil)

        result, err := mockClient.GetLoadBalancerMetricData(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetLoadBalancerTlsCertificates", func(t *testing.T) {
        input := &lightsail.GetLoadBalancerTlsCertificatesInput{}
        output := &lightsail.GetLoadBalancerTlsCertificatesOutput{}

        mockClient.On("GetLoadBalancerTlsCertificates", ctx, input).Return(output, nil)

        result, err := mockClient.GetLoadBalancerTlsCertificates(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetLoadBalancerTlsPolicies", func(t *testing.T) {
        input := &lightsail.GetLoadBalancerTlsPoliciesInput{}
        output := &lightsail.GetLoadBalancerTlsPoliciesOutput{}

        mockClient.On("GetLoadBalancerTlsPolicies", ctx, input).Return(output, nil)

        result, err := mockClient.GetLoadBalancerTlsPolicies(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetLoadBalancers", func(t *testing.T) {
        input := &lightsail.GetLoadBalancersInput{}
        output := &lightsail.GetLoadBalancersOutput{}

        mockClient.On("GetLoadBalancers", ctx, input).Return(output, nil)

        result, err := mockClient.GetLoadBalancers(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetOperation", func(t *testing.T) {
        input := &lightsail.GetOperationInput{}
        output := &lightsail.GetOperationOutput{}

        mockClient.On("GetOperation", ctx, input).Return(output, nil)

        result, err := mockClient.GetOperation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetOperations", func(t *testing.T) {
        input := &lightsail.GetOperationsInput{}
        output := &lightsail.GetOperationsOutput{}

        mockClient.On("GetOperations", ctx, input).Return(output, nil)

        result, err := mockClient.GetOperations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetOperationsForResource", func(t *testing.T) {
        input := &lightsail.GetOperationsForResourceInput{}
        output := &lightsail.GetOperationsForResourceOutput{}

        mockClient.On("GetOperationsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.GetOperationsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetRegions", func(t *testing.T) {
        input := &lightsail.GetRegionsInput{}
        output := &lightsail.GetRegionsOutput{}

        mockClient.On("GetRegions", ctx, input).Return(output, nil)

        result, err := mockClient.GetRegions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetRelationalDatabase", func(t *testing.T) {
        input := &lightsail.GetRelationalDatabaseInput{}
        output := &lightsail.GetRelationalDatabaseOutput{}

        mockClient.On("GetRelationalDatabase", ctx, input).Return(output, nil)

        result, err := mockClient.GetRelationalDatabase(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetRelationalDatabaseBlueprints", func(t *testing.T) {
        input := &lightsail.GetRelationalDatabaseBlueprintsInput{}
        output := &lightsail.GetRelationalDatabaseBlueprintsOutput{}

        mockClient.On("GetRelationalDatabaseBlueprints", ctx, input).Return(output, nil)

        result, err := mockClient.GetRelationalDatabaseBlueprints(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetRelationalDatabaseBundles", func(t *testing.T) {
        input := &lightsail.GetRelationalDatabaseBundlesInput{}
        output := &lightsail.GetRelationalDatabaseBundlesOutput{}

        mockClient.On("GetRelationalDatabaseBundles", ctx, input).Return(output, nil)

        result, err := mockClient.GetRelationalDatabaseBundles(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetRelationalDatabaseEvents", func(t *testing.T) {
        input := &lightsail.GetRelationalDatabaseEventsInput{}
        output := &lightsail.GetRelationalDatabaseEventsOutput{}

        mockClient.On("GetRelationalDatabaseEvents", ctx, input).Return(output, nil)

        result, err := mockClient.GetRelationalDatabaseEvents(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetRelationalDatabaseLogEvents", func(t *testing.T) {
        input := &lightsail.GetRelationalDatabaseLogEventsInput{}
        output := &lightsail.GetRelationalDatabaseLogEventsOutput{}

        mockClient.On("GetRelationalDatabaseLogEvents", ctx, input).Return(output, nil)

        result, err := mockClient.GetRelationalDatabaseLogEvents(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetRelationalDatabaseLogStreams", func(t *testing.T) {
        input := &lightsail.GetRelationalDatabaseLogStreamsInput{}
        output := &lightsail.GetRelationalDatabaseLogStreamsOutput{}

        mockClient.On("GetRelationalDatabaseLogStreams", ctx, input).Return(output, nil)

        result, err := mockClient.GetRelationalDatabaseLogStreams(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetRelationalDatabaseMasterUserPassword", func(t *testing.T) {
        input := &lightsail.GetRelationalDatabaseMasterUserPasswordInput{}
        output := &lightsail.GetRelationalDatabaseMasterUserPasswordOutput{}

        mockClient.On("GetRelationalDatabaseMasterUserPassword", ctx, input).Return(output, nil)

        result, err := mockClient.GetRelationalDatabaseMasterUserPassword(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetRelationalDatabaseMetricData", func(t *testing.T) {
        input := &lightsail.GetRelationalDatabaseMetricDataInput{}
        output := &lightsail.GetRelationalDatabaseMetricDataOutput{}

        mockClient.On("GetRelationalDatabaseMetricData", ctx, input).Return(output, nil)

        result, err := mockClient.GetRelationalDatabaseMetricData(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetRelationalDatabaseParameters", func(t *testing.T) {
        input := &lightsail.GetRelationalDatabaseParametersInput{}
        output := &lightsail.GetRelationalDatabaseParametersOutput{}

        mockClient.On("GetRelationalDatabaseParameters", ctx, input).Return(output, nil)

        result, err := mockClient.GetRelationalDatabaseParameters(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetRelationalDatabaseSnapshot", func(t *testing.T) {
        input := &lightsail.GetRelationalDatabaseSnapshotInput{}
        output := &lightsail.GetRelationalDatabaseSnapshotOutput{}

        mockClient.On("GetRelationalDatabaseSnapshot", ctx, input).Return(output, nil)

        result, err := mockClient.GetRelationalDatabaseSnapshot(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetRelationalDatabaseSnapshots", func(t *testing.T) {
        input := &lightsail.GetRelationalDatabaseSnapshotsInput{}
        output := &lightsail.GetRelationalDatabaseSnapshotsOutput{}

        mockClient.On("GetRelationalDatabaseSnapshots", ctx, input).Return(output, nil)

        result, err := mockClient.GetRelationalDatabaseSnapshots(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetRelationalDatabases", func(t *testing.T) {
        input := &lightsail.GetRelationalDatabasesInput{}
        output := &lightsail.GetRelationalDatabasesOutput{}

        mockClient.On("GetRelationalDatabases", ctx, input).Return(output, nil)

        result, err := mockClient.GetRelationalDatabases(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSetupHistory", func(t *testing.T) {
        input := &lightsail.GetSetupHistoryInput{}
        output := &lightsail.GetSetupHistoryOutput{}

        mockClient.On("GetSetupHistory", ctx, input).Return(output, nil)

        result, err := mockClient.GetSetupHistory(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetStaticIp", func(t *testing.T) {
        input := &lightsail.GetStaticIpInput{}
        output := &lightsail.GetStaticIpOutput{}

        mockClient.On("GetStaticIp", ctx, input).Return(output, nil)

        result, err := mockClient.GetStaticIp(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetStaticIps", func(t *testing.T) {
        input := &lightsail.GetStaticIpsInput{}
        output := &lightsail.GetStaticIpsOutput{}

        mockClient.On("GetStaticIps", ctx, input).Return(output, nil)

        result, err := mockClient.GetStaticIps(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestImportKeyPair", func(t *testing.T) {
        input := &lightsail.ImportKeyPairInput{}
        output := &lightsail.ImportKeyPairOutput{}

        mockClient.On("ImportKeyPair", ctx, input).Return(output, nil)

        result, err := mockClient.ImportKeyPair(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestIsVpcPeered", func(t *testing.T) {
        input := &lightsail.IsVpcPeeredInput{}
        output := &lightsail.IsVpcPeeredOutput{}

        mockClient.On("IsVpcPeered", ctx, input).Return(output, nil)

        result, err := mockClient.IsVpcPeered(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestOpenInstancePublicPorts", func(t *testing.T) {
        input := &lightsail.OpenInstancePublicPortsInput{}
        output := &lightsail.OpenInstancePublicPortsOutput{}

        mockClient.On("OpenInstancePublicPorts", ctx, input).Return(output, nil)

        result, err := mockClient.OpenInstancePublicPorts(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPeerVpc", func(t *testing.T) {
        input := &lightsail.PeerVpcInput{}
        output := &lightsail.PeerVpcOutput{}

        mockClient.On("PeerVpc", ctx, input).Return(output, nil)

        result, err := mockClient.PeerVpc(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutAlarm", func(t *testing.T) {
        input := &lightsail.PutAlarmInput{}
        output := &lightsail.PutAlarmOutput{}

        mockClient.On("PutAlarm", ctx, input).Return(output, nil)

        result, err := mockClient.PutAlarm(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutInstancePublicPorts", func(t *testing.T) {
        input := &lightsail.PutInstancePublicPortsInput{}
        output := &lightsail.PutInstancePublicPortsOutput{}

        mockClient.On("PutInstancePublicPorts", ctx, input).Return(output, nil)

        result, err := mockClient.PutInstancePublicPorts(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRebootInstance", func(t *testing.T) {
        input := &lightsail.RebootInstanceInput{}
        output := &lightsail.RebootInstanceOutput{}

        mockClient.On("RebootInstance", ctx, input).Return(output, nil)

        result, err := mockClient.RebootInstance(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRebootRelationalDatabase", func(t *testing.T) {
        input := &lightsail.RebootRelationalDatabaseInput{}
        output := &lightsail.RebootRelationalDatabaseOutput{}

        mockClient.On("RebootRelationalDatabase", ctx, input).Return(output, nil)

        result, err := mockClient.RebootRelationalDatabase(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRegisterContainerImage", func(t *testing.T) {
        input := &lightsail.RegisterContainerImageInput{}
        output := &lightsail.RegisterContainerImageOutput{}

        mockClient.On("RegisterContainerImage", ctx, input).Return(output, nil)

        result, err := mockClient.RegisterContainerImage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestReleaseStaticIp", func(t *testing.T) {
        input := &lightsail.ReleaseStaticIpInput{}
        output := &lightsail.ReleaseStaticIpOutput{}

        mockClient.On("ReleaseStaticIp", ctx, input).Return(output, nil)

        result, err := mockClient.ReleaseStaticIp(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestResetDistributionCache", func(t *testing.T) {
        input := &lightsail.ResetDistributionCacheInput{}
        output := &lightsail.ResetDistributionCacheOutput{}

        mockClient.On("ResetDistributionCache", ctx, input).Return(output, nil)

        result, err := mockClient.ResetDistributionCache(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSendContactMethodVerification", func(t *testing.T) {
        input := &lightsail.SendContactMethodVerificationInput{}
        output := &lightsail.SendContactMethodVerificationOutput{}

        mockClient.On("SendContactMethodVerification", ctx, input).Return(output, nil)

        result, err := mockClient.SendContactMethodVerification(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSetIpAddressType", func(t *testing.T) {
        input := &lightsail.SetIpAddressTypeInput{}
        output := &lightsail.SetIpAddressTypeOutput{}

        mockClient.On("SetIpAddressType", ctx, input).Return(output, nil)

        result, err := mockClient.SetIpAddressType(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSetResourceAccessForBucket", func(t *testing.T) {
        input := &lightsail.SetResourceAccessForBucketInput{}
        output := &lightsail.SetResourceAccessForBucketOutput{}

        mockClient.On("SetResourceAccessForBucket", ctx, input).Return(output, nil)

        result, err := mockClient.SetResourceAccessForBucket(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSetupInstanceHttps", func(t *testing.T) {
        input := &lightsail.SetupInstanceHttpsInput{}
        output := &lightsail.SetupInstanceHttpsOutput{}

        mockClient.On("SetupInstanceHttps", ctx, input).Return(output, nil)

        result, err := mockClient.SetupInstanceHttps(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartGUISession", func(t *testing.T) {
        input := &lightsail.StartGUISessionInput{}
        output := &lightsail.StartGUISessionOutput{}

        mockClient.On("StartGUISession", ctx, input).Return(output, nil)

        result, err := mockClient.StartGUISession(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartInstance", func(t *testing.T) {
        input := &lightsail.StartInstanceInput{}
        output := &lightsail.StartInstanceOutput{}

        mockClient.On("StartInstance", ctx, input).Return(output, nil)

        result, err := mockClient.StartInstance(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartRelationalDatabase", func(t *testing.T) {
        input := &lightsail.StartRelationalDatabaseInput{}
        output := &lightsail.StartRelationalDatabaseOutput{}

        mockClient.On("StartRelationalDatabase", ctx, input).Return(output, nil)

        result, err := mockClient.StartRelationalDatabase(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopGUISession", func(t *testing.T) {
        input := &lightsail.StopGUISessionInput{}
        output := &lightsail.StopGUISessionOutput{}

        mockClient.On("StopGUISession", ctx, input).Return(output, nil)

        result, err := mockClient.StopGUISession(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopInstance", func(t *testing.T) {
        input := &lightsail.StopInstanceInput{}
        output := &lightsail.StopInstanceOutput{}

        mockClient.On("StopInstance", ctx, input).Return(output, nil)

        result, err := mockClient.StopInstance(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopRelationalDatabase", func(t *testing.T) {
        input := &lightsail.StopRelationalDatabaseInput{}
        output := &lightsail.StopRelationalDatabaseOutput{}

        mockClient.On("StopRelationalDatabase", ctx, input).Return(output, nil)

        result, err := mockClient.StopRelationalDatabase(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &lightsail.TagResourceInput{}
        output := &lightsail.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTestAlarm", func(t *testing.T) {
        input := &lightsail.TestAlarmInput{}
        output := &lightsail.TestAlarmOutput{}

        mockClient.On("TestAlarm", ctx, input).Return(output, nil)

        result, err := mockClient.TestAlarm(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUnpeerVpc", func(t *testing.T) {
        input := &lightsail.UnpeerVpcInput{}
        output := &lightsail.UnpeerVpcOutput{}

        mockClient.On("UnpeerVpc", ctx, input).Return(output, nil)

        result, err := mockClient.UnpeerVpc(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &lightsail.UntagResourceInput{}
        output := &lightsail.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateBucket", func(t *testing.T) {
        input := &lightsail.UpdateBucketInput{}
        output := &lightsail.UpdateBucketOutput{}

        mockClient.On("UpdateBucket", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateBucket(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateBucketBundle", func(t *testing.T) {
        input := &lightsail.UpdateBucketBundleInput{}
        output := &lightsail.UpdateBucketBundleOutput{}

        mockClient.On("UpdateBucketBundle", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateBucketBundle(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateContainerService", func(t *testing.T) {
        input := &lightsail.UpdateContainerServiceInput{}
        output := &lightsail.UpdateContainerServiceOutput{}

        mockClient.On("UpdateContainerService", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateContainerService(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateDistribution", func(t *testing.T) {
        input := &lightsail.UpdateDistributionInput{}
        output := &lightsail.UpdateDistributionOutput{}

        mockClient.On("UpdateDistribution", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateDistribution(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateDistributionBundle", func(t *testing.T) {
        input := &lightsail.UpdateDistributionBundleInput{}
        output := &lightsail.UpdateDistributionBundleOutput{}

        mockClient.On("UpdateDistributionBundle", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateDistributionBundle(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateDomainEntry", func(t *testing.T) {
        input := &lightsail.UpdateDomainEntryInput{}
        output := &lightsail.UpdateDomainEntryOutput{}

        mockClient.On("UpdateDomainEntry", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateDomainEntry(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateInstanceMetadataOptions", func(t *testing.T) {
        input := &lightsail.UpdateInstanceMetadataOptionsInput{}
        output := &lightsail.UpdateInstanceMetadataOptionsOutput{}

        mockClient.On("UpdateInstanceMetadataOptions", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateInstanceMetadataOptions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateLoadBalancerAttribute", func(t *testing.T) {
        input := &lightsail.UpdateLoadBalancerAttributeInput{}
        output := &lightsail.UpdateLoadBalancerAttributeOutput{}

        mockClient.On("UpdateLoadBalancerAttribute", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateLoadBalancerAttribute(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateRelationalDatabase", func(t *testing.T) {
        input := &lightsail.UpdateRelationalDatabaseInput{}
        output := &lightsail.UpdateRelationalDatabaseOutput{}

        mockClient.On("UpdateRelationalDatabase", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateRelationalDatabase(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateRelationalDatabaseParameters", func(t *testing.T) {
        input := &lightsail.UpdateRelationalDatabaseParametersInput{}
        output := &lightsail.UpdateRelationalDatabaseParametersOutput{}

        mockClient.On("UpdateRelationalDatabaseParameters", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateRelationalDatabaseParameters(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
