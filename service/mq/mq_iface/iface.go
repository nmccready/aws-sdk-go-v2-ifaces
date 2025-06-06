
package mq_iface

// This file was generated by scripts/gen_ifaces.sh do not edit it manually
import (
    "context"
    . "github.com/aws/aws-sdk-go-v2/service/mq"
)

// IClient defines the interface for mq
type IClient interface {
 Options() Options 
 CreateBroker(ctx context.Context, params *CreateBrokerInput, optFns ...func(*Options)) (*CreateBrokerOutput, error) 
 CreateConfiguration(ctx context.Context, params *CreateConfigurationInput, optFns ...func(*Options)) (*CreateConfigurationOutput, error) 
 CreateTags(ctx context.Context, params *CreateTagsInput, optFns ...func(*Options)) (*CreateTagsOutput, error) 
 CreateUser(ctx context.Context, params *CreateUserInput, optFns ...func(*Options)) (*CreateUserOutput, error) 
 DeleteBroker(ctx context.Context, params *DeleteBrokerInput, optFns ...func(*Options)) (*DeleteBrokerOutput, error) 
 DeleteConfiguration(ctx context.Context, params *DeleteConfigurationInput, optFns ...func(*Options)) (*DeleteConfigurationOutput, error) 
 DeleteTags(ctx context.Context, params *DeleteTagsInput, optFns ...func(*Options)) (*DeleteTagsOutput, error) 
 DeleteUser(ctx context.Context, params *DeleteUserInput, optFns ...func(*Options)) (*DeleteUserOutput, error) 
 DescribeBroker(ctx context.Context, params *DescribeBrokerInput, optFns ...func(*Options)) (*DescribeBrokerOutput, error) 
 DescribeBrokerEngineTypes(ctx context.Context, params *DescribeBrokerEngineTypesInput, optFns ...func(*Options)) (*DescribeBrokerEngineTypesOutput, error) 
 DescribeBrokerInstanceOptions(ctx context.Context, params *DescribeBrokerInstanceOptionsInput, optFns ...func(*Options)) (*DescribeBrokerInstanceOptionsOutput, error) 
 DescribeConfiguration(ctx context.Context, params *DescribeConfigurationInput, optFns ...func(*Options)) (*DescribeConfigurationOutput, error) 
 DescribeConfigurationRevision(ctx context.Context, params *DescribeConfigurationRevisionInput, optFns ...func(*Options)) (*DescribeConfigurationRevisionOutput, error) 
 DescribeUser(ctx context.Context, params *DescribeUserInput, optFns ...func(*Options)) (*DescribeUserOutput, error) 
 ListBrokers(ctx context.Context, params *ListBrokersInput, optFns ...func(*Options)) (*ListBrokersOutput, error) 
 ListConfigurationRevisions(ctx context.Context, params *ListConfigurationRevisionsInput, optFns ...func(*Options)) (*ListConfigurationRevisionsOutput, error) 
 ListConfigurations(ctx context.Context, params *ListConfigurationsInput, optFns ...func(*Options)) (*ListConfigurationsOutput, error) 
 ListTags(ctx context.Context, params *ListTagsInput, optFns ...func(*Options)) (*ListTagsOutput, error) 
 ListUsers(ctx context.Context, params *ListUsersInput, optFns ...func(*Options)) (*ListUsersOutput, error) 
 Promote(ctx context.Context, params *PromoteInput, optFns ...func(*Options)) (*PromoteOutput, error) 
 RebootBroker(ctx context.Context, params *RebootBrokerInput, optFns ...func(*Options)) (*RebootBrokerOutput, error) 
 UpdateBroker(ctx context.Context, params *UpdateBrokerInput, optFns ...func(*Options)) (*UpdateBrokerOutput, error) 
 UpdateConfiguration(ctx context.Context, params *UpdateConfigurationInput, optFns ...func(*Options)) (*UpdateConfigurationOutput, error) 
 UpdateUser(ctx context.Context, params *UpdateUserInput, optFns ...func(*Options)) (*UpdateUserOutput, error) 
}
