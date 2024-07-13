package kafka_test

// tests for the kafka service interface for this ../../../service/kafka/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/kafka"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/kafka/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/kafka/kafka_iface"
	"github.com/stretchr/testify/assert"
)

func TestKafkaServiceCanBeMocked(t *testing.T) {
	var iface kafka_iface.IClient
	iface = &kafka.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := kafka.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchAssociateScramSecret", func(t *testing.T) {
        input := &kafka.BatchAssociateScramSecretInput{}
        output := &kafka.BatchAssociateScramSecretOutput{}

        mockClient.On("BatchAssociateScramSecret", ctx, input).Return(output, nil)

        result, err := mockClient.BatchAssociateScramSecret(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchDisassociateScramSecret", func(t *testing.T) {
        input := &kafka.BatchDisassociateScramSecretInput{}
        output := &kafka.BatchDisassociateScramSecretOutput{}

        mockClient.On("BatchDisassociateScramSecret", ctx, input).Return(output, nil)

        result, err := mockClient.BatchDisassociateScramSecret(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateCluster", func(t *testing.T) {
        input := &kafka.CreateClusterInput{}
        output := &kafka.CreateClusterOutput{}

        mockClient.On("CreateCluster", ctx, input).Return(output, nil)

        result, err := mockClient.CreateCluster(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateClusterV2", func(t *testing.T) {
        input := &kafka.CreateClusterV2Input{}
        output := &kafka.CreateClusterV2Output{}

        mockClient.On("CreateClusterV2", ctx, input).Return(output, nil)

        result, err := mockClient.CreateClusterV2(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateConfiguration", func(t *testing.T) {
        input := &kafka.CreateConfigurationInput{}
        output := &kafka.CreateConfigurationOutput{}

        mockClient.On("CreateConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.CreateConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateReplicator", func(t *testing.T) {
        input := &kafka.CreateReplicatorInput{}
        output := &kafka.CreateReplicatorOutput{}

        mockClient.On("CreateReplicator", ctx, input).Return(output, nil)

        result, err := mockClient.CreateReplicator(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateVpcConnection", func(t *testing.T) {
        input := &kafka.CreateVpcConnectionInput{}
        output := &kafka.CreateVpcConnectionOutput{}

        mockClient.On("CreateVpcConnection", ctx, input).Return(output, nil)

        result, err := mockClient.CreateVpcConnection(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteCluster", func(t *testing.T) {
        input := &kafka.DeleteClusterInput{}
        output := &kafka.DeleteClusterOutput{}

        mockClient.On("DeleteCluster", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteCluster(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteClusterPolicy", func(t *testing.T) {
        input := &kafka.DeleteClusterPolicyInput{}
        output := &kafka.DeleteClusterPolicyOutput{}

        mockClient.On("DeleteClusterPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteClusterPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteConfiguration", func(t *testing.T) {
        input := &kafka.DeleteConfigurationInput{}
        output := &kafka.DeleteConfigurationOutput{}

        mockClient.On("DeleteConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteReplicator", func(t *testing.T) {
        input := &kafka.DeleteReplicatorInput{}
        output := &kafka.DeleteReplicatorOutput{}

        mockClient.On("DeleteReplicator", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteReplicator(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteVpcConnection", func(t *testing.T) {
        input := &kafka.DeleteVpcConnectionInput{}
        output := &kafka.DeleteVpcConnectionOutput{}

        mockClient.On("DeleteVpcConnection", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteVpcConnection(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeCluster", func(t *testing.T) {
        input := &kafka.DescribeClusterInput{}
        output := &kafka.DescribeClusterOutput{}

        mockClient.On("DescribeCluster", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeCluster(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeClusterOperation", func(t *testing.T) {
        input := &kafka.DescribeClusterOperationInput{}
        output := &kafka.DescribeClusterOperationOutput{}

        mockClient.On("DescribeClusterOperation", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeClusterOperation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeClusterOperationV2", func(t *testing.T) {
        input := &kafka.DescribeClusterOperationV2Input{}
        output := &kafka.DescribeClusterOperationV2Output{}

        mockClient.On("DescribeClusterOperationV2", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeClusterOperationV2(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeClusterV2", func(t *testing.T) {
        input := &kafka.DescribeClusterV2Input{}
        output := &kafka.DescribeClusterV2Output{}

        mockClient.On("DescribeClusterV2", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeClusterV2(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeConfiguration", func(t *testing.T) {
        input := &kafka.DescribeConfigurationInput{}
        output := &kafka.DescribeConfigurationOutput{}

        mockClient.On("DescribeConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeConfigurationRevision", func(t *testing.T) {
        input := &kafka.DescribeConfigurationRevisionInput{}
        output := &kafka.DescribeConfigurationRevisionOutput{}

        mockClient.On("DescribeConfigurationRevision", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeConfigurationRevision(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeReplicator", func(t *testing.T) {
        input := &kafka.DescribeReplicatorInput{}
        output := &kafka.DescribeReplicatorOutput{}

        mockClient.On("DescribeReplicator", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeReplicator(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeVpcConnection", func(t *testing.T) {
        input := &kafka.DescribeVpcConnectionInput{}
        output := &kafka.DescribeVpcConnectionOutput{}

        mockClient.On("DescribeVpcConnection", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeVpcConnection(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetBootstrapBrokers", func(t *testing.T) {
        input := &kafka.GetBootstrapBrokersInput{}
        output := &kafka.GetBootstrapBrokersOutput{}

        mockClient.On("GetBootstrapBrokers", ctx, input).Return(output, nil)

        result, err := mockClient.GetBootstrapBrokers(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetClusterPolicy", func(t *testing.T) {
        input := &kafka.GetClusterPolicyInput{}
        output := &kafka.GetClusterPolicyOutput{}

        mockClient.On("GetClusterPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.GetClusterPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetCompatibleKafkaVersions", func(t *testing.T) {
        input := &kafka.GetCompatibleKafkaVersionsInput{}
        output := &kafka.GetCompatibleKafkaVersionsOutput{}

        mockClient.On("GetCompatibleKafkaVersions", ctx, input).Return(output, nil)

        result, err := mockClient.GetCompatibleKafkaVersions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListClientVpcConnections", func(t *testing.T) {
        input := &kafka.ListClientVpcConnectionsInput{}
        output := &kafka.ListClientVpcConnectionsOutput{}

        mockClient.On("ListClientVpcConnections", ctx, input).Return(output, nil)

        result, err := mockClient.ListClientVpcConnections(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListClusterOperations", func(t *testing.T) {
        input := &kafka.ListClusterOperationsInput{}
        output := &kafka.ListClusterOperationsOutput{}

        mockClient.On("ListClusterOperations", ctx, input).Return(output, nil)

        result, err := mockClient.ListClusterOperations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListClusterOperationsV2", func(t *testing.T) {
        input := &kafka.ListClusterOperationsV2Input{}
        output := &kafka.ListClusterOperationsV2Output{}

        mockClient.On("ListClusterOperationsV2", ctx, input).Return(output, nil)

        result, err := mockClient.ListClusterOperationsV2(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListClusters", func(t *testing.T) {
        input := &kafka.ListClustersInput{}
        output := &kafka.ListClustersOutput{}

        mockClient.On("ListClusters", ctx, input).Return(output, nil)

        result, err := mockClient.ListClusters(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListClustersV2", func(t *testing.T) {
        input := &kafka.ListClustersV2Input{}
        output := &kafka.ListClustersV2Output{}

        mockClient.On("ListClustersV2", ctx, input).Return(output, nil)

        result, err := mockClient.ListClustersV2(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListConfigurationRevisions", func(t *testing.T) {
        input := &kafka.ListConfigurationRevisionsInput{}
        output := &kafka.ListConfigurationRevisionsOutput{}

        mockClient.On("ListConfigurationRevisions", ctx, input).Return(output, nil)

        result, err := mockClient.ListConfigurationRevisions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListConfigurations", func(t *testing.T) {
        input := &kafka.ListConfigurationsInput{}
        output := &kafka.ListConfigurationsOutput{}

        mockClient.On("ListConfigurations", ctx, input).Return(output, nil)

        result, err := mockClient.ListConfigurations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListKafkaVersions", func(t *testing.T) {
        input := &kafka.ListKafkaVersionsInput{}
        output := &kafka.ListKafkaVersionsOutput{}

        mockClient.On("ListKafkaVersions", ctx, input).Return(output, nil)

        result, err := mockClient.ListKafkaVersions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListNodes", func(t *testing.T) {
        input := &kafka.ListNodesInput{}
        output := &kafka.ListNodesOutput{}

        mockClient.On("ListNodes", ctx, input).Return(output, nil)

        result, err := mockClient.ListNodes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListReplicators", func(t *testing.T) {
        input := &kafka.ListReplicatorsInput{}
        output := &kafka.ListReplicatorsOutput{}

        mockClient.On("ListReplicators", ctx, input).Return(output, nil)

        result, err := mockClient.ListReplicators(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListScramSecrets", func(t *testing.T) {
        input := &kafka.ListScramSecretsInput{}
        output := &kafka.ListScramSecretsOutput{}

        mockClient.On("ListScramSecrets", ctx, input).Return(output, nil)

        result, err := mockClient.ListScramSecrets(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &kafka.ListTagsForResourceInput{}
        output := &kafka.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListVpcConnections", func(t *testing.T) {
        input := &kafka.ListVpcConnectionsInput{}
        output := &kafka.ListVpcConnectionsOutput{}

        mockClient.On("ListVpcConnections", ctx, input).Return(output, nil)

        result, err := mockClient.ListVpcConnections(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutClusterPolicy", func(t *testing.T) {
        input := &kafka.PutClusterPolicyInput{}
        output := &kafka.PutClusterPolicyOutput{}

        mockClient.On("PutClusterPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.PutClusterPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRebootBroker", func(t *testing.T) {
        input := &kafka.RebootBrokerInput{}
        output := &kafka.RebootBrokerOutput{}

        mockClient.On("RebootBroker", ctx, input).Return(output, nil)

        result, err := mockClient.RebootBroker(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRejectClientVpcConnection", func(t *testing.T) {
        input := &kafka.RejectClientVpcConnectionInput{}
        output := &kafka.RejectClientVpcConnectionOutput{}

        mockClient.On("RejectClientVpcConnection", ctx, input).Return(output, nil)

        result, err := mockClient.RejectClientVpcConnection(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &kafka.TagResourceInput{}
        output := &kafka.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &kafka.UntagResourceInput{}
        output := &kafka.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateBrokerCount", func(t *testing.T) {
        input := &kafka.UpdateBrokerCountInput{}
        output := &kafka.UpdateBrokerCountOutput{}

        mockClient.On("UpdateBrokerCount", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateBrokerCount(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateBrokerStorage", func(t *testing.T) {
        input := &kafka.UpdateBrokerStorageInput{}
        output := &kafka.UpdateBrokerStorageOutput{}

        mockClient.On("UpdateBrokerStorage", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateBrokerStorage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateBrokerType", func(t *testing.T) {
        input := &kafka.UpdateBrokerTypeInput{}
        output := &kafka.UpdateBrokerTypeOutput{}

        mockClient.On("UpdateBrokerType", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateBrokerType(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateClusterConfiguration", func(t *testing.T) {
        input := &kafka.UpdateClusterConfigurationInput{}
        output := &kafka.UpdateClusterConfigurationOutput{}

        mockClient.On("UpdateClusterConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateClusterConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateClusterKafkaVersion", func(t *testing.T) {
        input := &kafka.UpdateClusterKafkaVersionInput{}
        output := &kafka.UpdateClusterKafkaVersionOutput{}

        mockClient.On("UpdateClusterKafkaVersion", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateClusterKafkaVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateConfiguration", func(t *testing.T) {
        input := &kafka.UpdateConfigurationInput{}
        output := &kafka.UpdateConfigurationOutput{}

        mockClient.On("UpdateConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateConnectivity", func(t *testing.T) {
        input := &kafka.UpdateConnectivityInput{}
        output := &kafka.UpdateConnectivityOutput{}

        mockClient.On("UpdateConnectivity", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateConnectivity(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateMonitoring", func(t *testing.T) {
        input := &kafka.UpdateMonitoringInput{}
        output := &kafka.UpdateMonitoringOutput{}

        mockClient.On("UpdateMonitoring", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateMonitoring(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateReplicationInfo", func(t *testing.T) {
        input := &kafka.UpdateReplicationInfoInput{}
        output := &kafka.UpdateReplicationInfoOutput{}

        mockClient.On("UpdateReplicationInfo", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateReplicationInfo(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateSecurity", func(t *testing.T) {
        input := &kafka.UpdateSecurityInput{}
        output := &kafka.UpdateSecurityOutput{}

        mockClient.On("UpdateSecurity", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateSecurity(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateStorage", func(t *testing.T) {
        input := &kafka.UpdateStorageInput{}
        output := &kafka.UpdateStorageOutput{}

        mockClient.On("UpdateStorage", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateStorage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
