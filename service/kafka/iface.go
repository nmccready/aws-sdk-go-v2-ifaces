
package kafka

import (
    "github.com/aws/aws-sdk-go-v2/service/kafka"
)

// IKafka defines the interface for kafka
type IKafka interface {
 Options() Options 
 BatchAssociateScramSecret(ctx context.Context, params *BatchAssociateScramSecretInput, optFns ...func(*Options)) (*BatchAssociateScramSecretOutput, error) 
 BatchDisassociateScramSecret(ctx context.Context, params *BatchDisassociateScramSecretInput, optFns ...func(*Options)) (*BatchDisassociateScramSecretOutput, error) 
 CreateCluster(ctx context.Context, params *CreateClusterInput, optFns ...func(*Options)) (*CreateClusterOutput, error) 
 CreateClusterV2(ctx context.Context, params *CreateClusterV2Input, optFns ...func(*Options)) (*CreateClusterV2Output, error) 
 CreateConfiguration(ctx context.Context, params *CreateConfigurationInput, optFns ...func(*Options)) (*CreateConfigurationOutput, error) 
 CreateReplicator(ctx context.Context, params *CreateReplicatorInput, optFns ...func(*Options)) (*CreateReplicatorOutput, error) 
 CreateVpcConnection(ctx context.Context, params *CreateVpcConnectionInput, optFns ...func(*Options)) (*CreateVpcConnectionOutput, error) 
 DeleteCluster(ctx context.Context, params *DeleteClusterInput, optFns ...func(*Options)) (*DeleteClusterOutput, error) 
 DeleteClusterPolicy(ctx context.Context, params *DeleteClusterPolicyInput, optFns ...func(*Options)) (*DeleteClusterPolicyOutput, error) 
 DeleteConfiguration(ctx context.Context, params *DeleteConfigurationInput, optFns ...func(*Options)) (*DeleteConfigurationOutput, error) 
 DeleteReplicator(ctx context.Context, params *DeleteReplicatorInput, optFns ...func(*Options)) (*DeleteReplicatorOutput, error) 
 DeleteVpcConnection(ctx context.Context, params *DeleteVpcConnectionInput, optFns ...func(*Options)) (*DeleteVpcConnectionOutput, error) 
 DescribeCluster(ctx context.Context, params *DescribeClusterInput, optFns ...func(*Options)) (*DescribeClusterOutput, error) 
 DescribeClusterOperation(ctx context.Context, params *DescribeClusterOperationInput, optFns ...func(*Options)) (*DescribeClusterOperationOutput, error) 
 DescribeClusterOperationV2(ctx context.Context, params *DescribeClusterOperationV2Input, optFns ...func(*Options)) (*DescribeClusterOperationV2Output, error) 
 DescribeClusterV2(ctx context.Context, params *DescribeClusterV2Input, optFns ...func(*Options)) (*DescribeClusterV2Output, error) 
 DescribeConfiguration(ctx context.Context, params *DescribeConfigurationInput, optFns ...func(*Options)) (*DescribeConfigurationOutput, error) 
 DescribeConfigurationRevision(ctx context.Context, params *DescribeConfigurationRevisionInput, optFns ...func(*Options)) (*DescribeConfigurationRevisionOutput, error) 
 DescribeReplicator(ctx context.Context, params *DescribeReplicatorInput, optFns ...func(*Options)) (*DescribeReplicatorOutput, error) 
 DescribeVpcConnection(ctx context.Context, params *DescribeVpcConnectionInput, optFns ...func(*Options)) (*DescribeVpcConnectionOutput, error) 
 GetBootstrapBrokers(ctx context.Context, params *GetBootstrapBrokersInput, optFns ...func(*Options)) (*GetBootstrapBrokersOutput, error) 
 GetClusterPolicy(ctx context.Context, params *GetClusterPolicyInput, optFns ...func(*Options)) (*GetClusterPolicyOutput, error) 
 GetCompatibleKafkaVersions(ctx context.Context, params *GetCompatibleKafkaVersionsInput, optFns ...func(*Options)) (*GetCompatibleKafkaVersionsOutput, error) 
 ListClientVpcConnections(ctx context.Context, params *ListClientVpcConnectionsInput, optFns ...func(*Options)) (*ListClientVpcConnectionsOutput, error) 
 ListClusterOperations(ctx context.Context, params *ListClusterOperationsInput, optFns ...func(*Options)) (*ListClusterOperationsOutput, error) 
 ListClusterOperationsV2(ctx context.Context, params *ListClusterOperationsV2Input, optFns ...func(*Options)) (*ListClusterOperationsV2Output, error) 
 ListClusters(ctx context.Context, params *ListClustersInput, optFns ...func(*Options)) (*ListClustersOutput, error) 
 ListClustersV2(ctx context.Context, params *ListClustersV2Input, optFns ...func(*Options)) (*ListClustersV2Output, error) 
 ListConfigurationRevisions(ctx context.Context, params *ListConfigurationRevisionsInput, optFns ...func(*Options)) (*ListConfigurationRevisionsOutput, error) 
 ListConfigurations(ctx context.Context, params *ListConfigurationsInput, optFns ...func(*Options)) (*ListConfigurationsOutput, error) 
 ListKafkaVersions(ctx context.Context, params *ListKafkaVersionsInput, optFns ...func(*Options)) (*ListKafkaVersionsOutput, error) 
 ListNodes(ctx context.Context, params *ListNodesInput, optFns ...func(*Options)) (*ListNodesOutput, error) 
 ListReplicators(ctx context.Context, params *ListReplicatorsInput, optFns ...func(*Options)) (*ListReplicatorsOutput, error) 
 ListScramSecrets(ctx context.Context, params *ListScramSecretsInput, optFns ...func(*Options)) (*ListScramSecretsOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 ListVpcConnections(ctx context.Context, params *ListVpcConnectionsInput, optFns ...func(*Options)) (*ListVpcConnectionsOutput, error) 
 PutClusterPolicy(ctx context.Context, params *PutClusterPolicyInput, optFns ...func(*Options)) (*PutClusterPolicyOutput, error) 
 RebootBroker(ctx context.Context, params *RebootBrokerInput, optFns ...func(*Options)) (*RebootBrokerOutput, error) 
 RejectClientVpcConnection(ctx context.Context, params *RejectClientVpcConnectionInput, optFns ...func(*Options)) (*RejectClientVpcConnectionOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateBrokerCount(ctx context.Context, params *UpdateBrokerCountInput, optFns ...func(*Options)) (*UpdateBrokerCountOutput, error) 
 UpdateBrokerStorage(ctx context.Context, params *UpdateBrokerStorageInput, optFns ...func(*Options)) (*UpdateBrokerStorageOutput, error) 
 UpdateBrokerType(ctx context.Context, params *UpdateBrokerTypeInput, optFns ...func(*Options)) (*UpdateBrokerTypeOutput, error) 
 UpdateClusterConfiguration(ctx context.Context, params *UpdateClusterConfigurationInput, optFns ...func(*Options)) (*UpdateClusterConfigurationOutput, error) 
 UpdateClusterKafkaVersion(ctx context.Context, params *UpdateClusterKafkaVersionInput, optFns ...func(*Options)) (*UpdateClusterKafkaVersionOutput, error) 
 UpdateConfiguration(ctx context.Context, params *UpdateConfigurationInput, optFns ...func(*Options)) (*UpdateConfigurationOutput, error) 
 UpdateConnectivity(ctx context.Context, params *UpdateConnectivityInput, optFns ...func(*Options)) (*UpdateConnectivityOutput, error) 
 UpdateMonitoring(ctx context.Context, params *UpdateMonitoringInput, optFns ...func(*Options)) (*UpdateMonitoringOutput, error) 
 UpdateReplicationInfo(ctx context.Context, params *UpdateReplicationInfoInput, optFns ...func(*Options)) (*UpdateReplicationInfoOutput, error) 
 UpdateSecurity(ctx context.Context, params *UpdateSecurityInput, optFns ...func(*Options)) (*UpdateSecurityOutput, error) 
 UpdateStorage(ctx context.Context, params *UpdateStorageInput, optFns ...func(*Options)) (*UpdateStorageOutput, error) 
}
