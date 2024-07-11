
package serverlessapplicationrepository

import (
    "github.com/aws/aws-sdk-go-v2/service/serverlessapplicationrepository"
)

// IClient defines the interface for serverlessapplicationrepository
type IClient interface {
 Options() Options 
 CreateApplication(ctx context.Context, params *CreateApplicationInput, optFns ...func(*Options)) (*CreateApplicationOutput, error) 
 CreateApplicationVersion(ctx context.Context, params *CreateApplicationVersionInput, optFns ...func(*Options)) (*CreateApplicationVersionOutput, error) 
 CreateCloudFormationChangeSet(ctx context.Context, params *CreateCloudFormationChangeSetInput, optFns ...func(*Options)) (*CreateCloudFormationChangeSetOutput, error) 
 CreateCloudFormationTemplate(ctx context.Context, params *CreateCloudFormationTemplateInput, optFns ...func(*Options)) (*CreateCloudFormationTemplateOutput, error) 
 DeleteApplication(ctx context.Context, params *DeleteApplicationInput, optFns ...func(*Options)) (*DeleteApplicationOutput, error) 
 GetApplication(ctx context.Context, params *GetApplicationInput, optFns ...func(*Options)) (*GetApplicationOutput, error) 
 GetApplicationPolicy(ctx context.Context, params *GetApplicationPolicyInput, optFns ...func(*Options)) (*GetApplicationPolicyOutput, error) 
 GetCloudFormationTemplate(ctx context.Context, params *GetCloudFormationTemplateInput, optFns ...func(*Options)) (*GetCloudFormationTemplateOutput, error) 
 ListApplicationDependencies(ctx context.Context, params *ListApplicationDependenciesInput, optFns ...func(*Options)) (*ListApplicationDependenciesOutput, error) 
 ListApplicationVersions(ctx context.Context, params *ListApplicationVersionsInput, optFns ...func(*Options)) (*ListApplicationVersionsOutput, error) 
 ListApplications(ctx context.Context, params *ListApplicationsInput, optFns ...func(*Options)) (*ListApplicationsOutput, error) 
 PutApplicationPolicy(ctx context.Context, params *PutApplicationPolicyInput, optFns ...func(*Options)) (*PutApplicationPolicyOutput, error) 
 UnshareApplication(ctx context.Context, params *UnshareApplicationInput, optFns ...func(*Options)) (*UnshareApplicationOutput, error) 
 UpdateApplication(ctx context.Context, params *UpdateApplicationInput, optFns ...func(*Options)) (*UpdateApplicationOutput, error) 
}
