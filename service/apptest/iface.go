
package apptest

import (
    "github.com/aws/aws-sdk-go-v2/service/apptest"
)

// IClient defines the interface for apptest
type IClient interface {
 Options() Options 
 CreateTestCase(ctx context.Context, params *CreateTestCaseInput, optFns ...func(*Options)) (*CreateTestCaseOutput, error) 
 CreateTestConfiguration(ctx context.Context, params *CreateTestConfigurationInput, optFns ...func(*Options)) (*CreateTestConfigurationOutput, error) 
 CreateTestSuite(ctx context.Context, params *CreateTestSuiteInput, optFns ...func(*Options)) (*CreateTestSuiteOutput, error) 
 DeleteTestCase(ctx context.Context, params *DeleteTestCaseInput, optFns ...func(*Options)) (*DeleteTestCaseOutput, error) 
 DeleteTestConfiguration(ctx context.Context, params *DeleteTestConfigurationInput, optFns ...func(*Options)) (*DeleteTestConfigurationOutput, error) 
 DeleteTestRun(ctx context.Context, params *DeleteTestRunInput, optFns ...func(*Options)) (*DeleteTestRunOutput, error) 
 DeleteTestSuite(ctx context.Context, params *DeleteTestSuiteInput, optFns ...func(*Options)) (*DeleteTestSuiteOutput, error) 
 GetTestCase(ctx context.Context, params *GetTestCaseInput, optFns ...func(*Options)) (*GetTestCaseOutput, error) 
 GetTestConfiguration(ctx context.Context, params *GetTestConfigurationInput, optFns ...func(*Options)) (*GetTestConfigurationOutput, error) 
 GetTestRunStep(ctx context.Context, params *GetTestRunStepInput, optFns ...func(*Options)) (*GetTestRunStepOutput, error) 
 GetTestSuite(ctx context.Context, params *GetTestSuiteInput, optFns ...func(*Options)) (*GetTestSuiteOutput, error) 
 ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) 
 ListTestCases(ctx context.Context, params *ListTestCasesInput, optFns ...func(*Options)) (*ListTestCasesOutput, error) 
 ListTestConfigurations(ctx context.Context, params *ListTestConfigurationsInput, optFns ...func(*Options)) (*ListTestConfigurationsOutput, error) 
 ListTestRunSteps(ctx context.Context, params *ListTestRunStepsInput, optFns ...func(*Options)) (*ListTestRunStepsOutput, error) 
 ListTestRunTestCases(ctx context.Context, params *ListTestRunTestCasesInput, optFns ...func(*Options)) (*ListTestRunTestCasesOutput, error) 
 ListTestRuns(ctx context.Context, params *ListTestRunsInput, optFns ...func(*Options)) (*ListTestRunsOutput, error) 
 ListTestSuites(ctx context.Context, params *ListTestSuitesInput, optFns ...func(*Options)) (*ListTestSuitesOutput, error) 
 StartTestRun(ctx context.Context, params *StartTestRunInput, optFns ...func(*Options)) (*StartTestRunOutput, error) 
 TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) 
 UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) 
 UpdateTestCase(ctx context.Context, params *UpdateTestCaseInput, optFns ...func(*Options)) (*UpdateTestCaseOutput, error) 
 UpdateTestConfiguration(ctx context.Context, params *UpdateTestConfigurationInput, optFns ...func(*Options)) (*UpdateTestConfigurationOutput, error) 
 UpdateTestSuite(ctx context.Context, params *UpdateTestSuiteInput, optFns ...func(*Options)) (*UpdateTestSuiteOutput, error) 
}
