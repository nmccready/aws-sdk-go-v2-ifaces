package apptest_test

// tests for the apptest service interface for this ../../../service/apptest/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/apptest"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/apptest/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/apptest/apptest_iface"
	"github.com/stretchr/testify/assert"
)

func TestApptestServiceCanBeMocked(t *testing.T) {
	var iface apptest_iface.IClient
	iface = &apptest.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := apptest.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateTestCase", func(t *testing.T) {
        input := &apptest.CreateTestCaseInput{}
        output := &apptest.CreateTestCaseOutput{}

        mockClient.On("CreateTestCase", ctx, input).Return(output, nil)

        result, err := mockClient.CreateTestCase(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateTestConfiguration", func(t *testing.T) {
        input := &apptest.CreateTestConfigurationInput{}
        output := &apptest.CreateTestConfigurationOutput{}

        mockClient.On("CreateTestConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.CreateTestConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateTestSuite", func(t *testing.T) {
        input := &apptest.CreateTestSuiteInput{}
        output := &apptest.CreateTestSuiteOutput{}

        mockClient.On("CreateTestSuite", ctx, input).Return(output, nil)

        result, err := mockClient.CreateTestSuite(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteTestCase", func(t *testing.T) {
        input := &apptest.DeleteTestCaseInput{}
        output := &apptest.DeleteTestCaseOutput{}

        mockClient.On("DeleteTestCase", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteTestCase(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteTestConfiguration", func(t *testing.T) {
        input := &apptest.DeleteTestConfigurationInput{}
        output := &apptest.DeleteTestConfigurationOutput{}

        mockClient.On("DeleteTestConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteTestConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteTestRun", func(t *testing.T) {
        input := &apptest.DeleteTestRunInput{}
        output := &apptest.DeleteTestRunOutput{}

        mockClient.On("DeleteTestRun", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteTestRun(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteTestSuite", func(t *testing.T) {
        input := &apptest.DeleteTestSuiteInput{}
        output := &apptest.DeleteTestSuiteOutput{}

        mockClient.On("DeleteTestSuite", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteTestSuite(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetTestCase", func(t *testing.T) {
        input := &apptest.GetTestCaseInput{}
        output := &apptest.GetTestCaseOutput{}

        mockClient.On("GetTestCase", ctx, input).Return(output, nil)

        result, err := mockClient.GetTestCase(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetTestConfiguration", func(t *testing.T) {
        input := &apptest.GetTestConfigurationInput{}
        output := &apptest.GetTestConfigurationOutput{}

        mockClient.On("GetTestConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.GetTestConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetTestRunStep", func(t *testing.T) {
        input := &apptest.GetTestRunStepInput{}
        output := &apptest.GetTestRunStepOutput{}

        mockClient.On("GetTestRunStep", ctx, input).Return(output, nil)

        result, err := mockClient.GetTestRunStep(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetTestSuite", func(t *testing.T) {
        input := &apptest.GetTestSuiteInput{}
        output := &apptest.GetTestSuiteOutput{}

        mockClient.On("GetTestSuite", ctx, input).Return(output, nil)

        result, err := mockClient.GetTestSuite(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &apptest.ListTagsForResourceInput{}
        output := &apptest.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTestCases", func(t *testing.T) {
        input := &apptest.ListTestCasesInput{}
        output := &apptest.ListTestCasesOutput{}

        mockClient.On("ListTestCases", ctx, input).Return(output, nil)

        result, err := mockClient.ListTestCases(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTestConfigurations", func(t *testing.T) {
        input := &apptest.ListTestConfigurationsInput{}
        output := &apptest.ListTestConfigurationsOutput{}

        mockClient.On("ListTestConfigurations", ctx, input).Return(output, nil)

        result, err := mockClient.ListTestConfigurations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTestRunSteps", func(t *testing.T) {
        input := &apptest.ListTestRunStepsInput{}
        output := &apptest.ListTestRunStepsOutput{}

        mockClient.On("ListTestRunSteps", ctx, input).Return(output, nil)

        result, err := mockClient.ListTestRunSteps(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTestRunTestCases", func(t *testing.T) {
        input := &apptest.ListTestRunTestCasesInput{}
        output := &apptest.ListTestRunTestCasesOutput{}

        mockClient.On("ListTestRunTestCases", ctx, input).Return(output, nil)

        result, err := mockClient.ListTestRunTestCases(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTestRuns", func(t *testing.T) {
        input := &apptest.ListTestRunsInput{}
        output := &apptest.ListTestRunsOutput{}

        mockClient.On("ListTestRuns", ctx, input).Return(output, nil)

        result, err := mockClient.ListTestRuns(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTestSuites", func(t *testing.T) {
        input := &apptest.ListTestSuitesInput{}
        output := &apptest.ListTestSuitesOutput{}

        mockClient.On("ListTestSuites", ctx, input).Return(output, nil)

        result, err := mockClient.ListTestSuites(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartTestRun", func(t *testing.T) {
        input := &apptest.StartTestRunInput{}
        output := &apptest.StartTestRunOutput{}

        mockClient.On("StartTestRun", ctx, input).Return(output, nil)

        result, err := mockClient.StartTestRun(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &apptest.TagResourceInput{}
        output := &apptest.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &apptest.UntagResourceInput{}
        output := &apptest.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateTestCase", func(t *testing.T) {
        input := &apptest.UpdateTestCaseInput{}
        output := &apptest.UpdateTestCaseOutput{}

        mockClient.On("UpdateTestCase", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateTestCase(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateTestConfiguration", func(t *testing.T) {
        input := &apptest.UpdateTestConfigurationInput{}
        output := &apptest.UpdateTestConfigurationOutput{}

        mockClient.On("UpdateTestConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateTestConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateTestSuite", func(t *testing.T) {
        input := &apptest.UpdateTestSuiteInput{}
        output := &apptest.UpdateTestSuiteOutput{}

        mockClient.On("UpdateTestSuite", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateTestSuite(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
