
package mobile_iface

// This file was generated by scripts/gen_ifaces.sh do not edit it manually
import (
    "context"
    . "github.com/aws/aws-sdk-go-v2/service/mobile"
)

// IClient defines the interface for mobile
type IClient interface {
 Options() Options 
 CreateProject(ctx context.Context, params *CreateProjectInput, optFns ...func(*Options)) (*CreateProjectOutput, error) 
 DeleteProject(ctx context.Context, params *DeleteProjectInput, optFns ...func(*Options)) (*DeleteProjectOutput, error) 
 DescribeBundle(ctx context.Context, params *DescribeBundleInput, optFns ...func(*Options)) (*DescribeBundleOutput, error) 
 DescribeProject(ctx context.Context, params *DescribeProjectInput, optFns ...func(*Options)) (*DescribeProjectOutput, error) 
 ExportBundle(ctx context.Context, params *ExportBundleInput, optFns ...func(*Options)) (*ExportBundleOutput, error) 
 ExportProject(ctx context.Context, params *ExportProjectInput, optFns ...func(*Options)) (*ExportProjectOutput, error) 
 ListBundles(ctx context.Context, params *ListBundlesInput, optFns ...func(*Options)) (*ListBundlesOutput, error) 
 ListProjects(ctx context.Context, params *ListProjectsInput, optFns ...func(*Options)) (*ListProjectsOutput, error) 
 UpdateProject(ctx context.Context, params *UpdateProjectInput, optFns ...func(*Options)) (*UpdateProjectOutput, error) 
}
