package support_test

// tests for the support service interface for this ../../../service/support/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/support"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/support/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/support/support_iface"
	"github.com/stretchr/testify/assert"
)

func TestSupportServiceCanBeMocked(t *testing.T) {
	var iface support_iface.IClient
	iface = &support.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := support.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAddAttachmentsToSet", func(t *testing.T) {
        input := &support.AddAttachmentsToSetInput{}
        output := &support.AddAttachmentsToSetOutput{}

        mockClient.On("AddAttachmentsToSet", ctx, input).Return(output, nil)

        result, err := mockClient.AddAttachmentsToSet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAddCommunicationToCase", func(t *testing.T) {
        input := &support.AddCommunicationToCaseInput{}
        output := &support.AddCommunicationToCaseOutput{}

        mockClient.On("AddCommunicationToCase", ctx, input).Return(output, nil)

        result, err := mockClient.AddCommunicationToCase(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateCase", func(t *testing.T) {
        input := &support.CreateCaseInput{}
        output := &support.CreateCaseOutput{}

        mockClient.On("CreateCase", ctx, input).Return(output, nil)

        result, err := mockClient.CreateCase(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeAttachment", func(t *testing.T) {
        input := &support.DescribeAttachmentInput{}
        output := &support.DescribeAttachmentOutput{}

        mockClient.On("DescribeAttachment", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeAttachment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeCases", func(t *testing.T) {
        input := &support.DescribeCasesInput{}
        output := &support.DescribeCasesOutput{}

        mockClient.On("DescribeCases", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeCases(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeCommunications", func(t *testing.T) {
        input := &support.DescribeCommunicationsInput{}
        output := &support.DescribeCommunicationsOutput{}

        mockClient.On("DescribeCommunications", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeCommunications(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeCreateCaseOptions", func(t *testing.T) {
        input := &support.DescribeCreateCaseOptionsInput{}
        output := &support.DescribeCreateCaseOptionsOutput{}

        mockClient.On("DescribeCreateCaseOptions", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeCreateCaseOptions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeServices", func(t *testing.T) {
        input := &support.DescribeServicesInput{}
        output := &support.DescribeServicesOutput{}

        mockClient.On("DescribeServices", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeServices(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeSeverityLevels", func(t *testing.T) {
        input := &support.DescribeSeverityLevelsInput{}
        output := &support.DescribeSeverityLevelsOutput{}

        mockClient.On("DescribeSeverityLevels", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeSeverityLevels(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeSupportedLanguages", func(t *testing.T) {
        input := &support.DescribeSupportedLanguagesInput{}
        output := &support.DescribeSupportedLanguagesOutput{}

        mockClient.On("DescribeSupportedLanguages", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeSupportedLanguages(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeTrustedAdvisorCheckRefreshStatuses", func(t *testing.T) {
        input := &support.DescribeTrustedAdvisorCheckRefreshStatusesInput{}
        output := &support.DescribeTrustedAdvisorCheckRefreshStatusesOutput{}

        mockClient.On("DescribeTrustedAdvisorCheckRefreshStatuses", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeTrustedAdvisorCheckRefreshStatuses(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeTrustedAdvisorCheckResult", func(t *testing.T) {
        input := &support.DescribeTrustedAdvisorCheckResultInput{}
        output := &support.DescribeTrustedAdvisorCheckResultOutput{}

        mockClient.On("DescribeTrustedAdvisorCheckResult", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeTrustedAdvisorCheckResult(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeTrustedAdvisorCheckSummaries", func(t *testing.T) {
        input := &support.DescribeTrustedAdvisorCheckSummariesInput{}
        output := &support.DescribeTrustedAdvisorCheckSummariesOutput{}

        mockClient.On("DescribeTrustedAdvisorCheckSummaries", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeTrustedAdvisorCheckSummaries(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeTrustedAdvisorChecks", func(t *testing.T) {
        input := &support.DescribeTrustedAdvisorChecksInput{}
        output := &support.DescribeTrustedAdvisorChecksOutput{}

        mockClient.On("DescribeTrustedAdvisorChecks", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeTrustedAdvisorChecks(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRefreshTrustedAdvisorCheck", func(t *testing.T) {
        input := &support.RefreshTrustedAdvisorCheckInput{}
        output := &support.RefreshTrustedAdvisorCheckOutput{}

        mockClient.On("RefreshTrustedAdvisorCheck", ctx, input).Return(output, nil)

        result, err := mockClient.RefreshTrustedAdvisorCheck(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestResolveCase", func(t *testing.T) {
        input := &support.ResolveCaseInput{}
        output := &support.ResolveCaseOutput{}

        mockClient.On("ResolveCase", ctx, input).Return(output, nil)

        result, err := mockClient.ResolveCase(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
