package cloudsearch_test

// tests for the cloudsearch service interface for this ../../../service/cloudsearch/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/cloudsearch"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/cloudsearch/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/cloudsearch/cloudsearch_iface"
	"github.com/stretchr/testify/assert"
)

func TestCloudsearchServiceCanBeMocked(t *testing.T) {
	var iface cloudsearch_iface.IClient
	iface = &cloudsearch.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := cloudsearch.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBuildSuggesters", func(t *testing.T) {
        input := &cloudsearch.BuildSuggestersInput{}
        output := &cloudsearch.BuildSuggestersOutput{}

        mockClient.On("BuildSuggesters", ctx, input).Return(output, nil)

        result, err := mockClient.BuildSuggesters(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateDomain", func(t *testing.T) {
        input := &cloudsearch.CreateDomainInput{}
        output := &cloudsearch.CreateDomainOutput{}

        mockClient.On("CreateDomain", ctx, input).Return(output, nil)

        result, err := mockClient.CreateDomain(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDefineAnalysisScheme", func(t *testing.T) {
        input := &cloudsearch.DefineAnalysisSchemeInput{}
        output := &cloudsearch.DefineAnalysisSchemeOutput{}

        mockClient.On("DefineAnalysisScheme", ctx, input).Return(output, nil)

        result, err := mockClient.DefineAnalysisScheme(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDefineExpression", func(t *testing.T) {
        input := &cloudsearch.DefineExpressionInput{}
        output := &cloudsearch.DefineExpressionOutput{}

        mockClient.On("DefineExpression", ctx, input).Return(output, nil)

        result, err := mockClient.DefineExpression(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDefineIndexField", func(t *testing.T) {
        input := &cloudsearch.DefineIndexFieldInput{}
        output := &cloudsearch.DefineIndexFieldOutput{}

        mockClient.On("DefineIndexField", ctx, input).Return(output, nil)

        result, err := mockClient.DefineIndexField(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDefineSuggester", func(t *testing.T) {
        input := &cloudsearch.DefineSuggesterInput{}
        output := &cloudsearch.DefineSuggesterOutput{}

        mockClient.On("DefineSuggester", ctx, input).Return(output, nil)

        result, err := mockClient.DefineSuggester(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteAnalysisScheme", func(t *testing.T) {
        input := &cloudsearch.DeleteAnalysisSchemeInput{}
        output := &cloudsearch.DeleteAnalysisSchemeOutput{}

        mockClient.On("DeleteAnalysisScheme", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteAnalysisScheme(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteDomain", func(t *testing.T) {
        input := &cloudsearch.DeleteDomainInput{}
        output := &cloudsearch.DeleteDomainOutput{}

        mockClient.On("DeleteDomain", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteDomain(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteExpression", func(t *testing.T) {
        input := &cloudsearch.DeleteExpressionInput{}
        output := &cloudsearch.DeleteExpressionOutput{}

        mockClient.On("DeleteExpression", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteExpression(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteIndexField", func(t *testing.T) {
        input := &cloudsearch.DeleteIndexFieldInput{}
        output := &cloudsearch.DeleteIndexFieldOutput{}

        mockClient.On("DeleteIndexField", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteIndexField(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteSuggester", func(t *testing.T) {
        input := &cloudsearch.DeleteSuggesterInput{}
        output := &cloudsearch.DeleteSuggesterOutput{}

        mockClient.On("DeleteSuggester", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteSuggester(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeAnalysisSchemes", func(t *testing.T) {
        input := &cloudsearch.DescribeAnalysisSchemesInput{}
        output := &cloudsearch.DescribeAnalysisSchemesOutput{}

        mockClient.On("DescribeAnalysisSchemes", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeAnalysisSchemes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeAvailabilityOptions", func(t *testing.T) {
        input := &cloudsearch.DescribeAvailabilityOptionsInput{}
        output := &cloudsearch.DescribeAvailabilityOptionsOutput{}

        mockClient.On("DescribeAvailabilityOptions", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeAvailabilityOptions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeDomainEndpointOptions", func(t *testing.T) {
        input := &cloudsearch.DescribeDomainEndpointOptionsInput{}
        output := &cloudsearch.DescribeDomainEndpointOptionsOutput{}

        mockClient.On("DescribeDomainEndpointOptions", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeDomainEndpointOptions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeDomains", func(t *testing.T) {
        input := &cloudsearch.DescribeDomainsInput{}
        output := &cloudsearch.DescribeDomainsOutput{}

        mockClient.On("DescribeDomains", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeDomains(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeExpressions", func(t *testing.T) {
        input := &cloudsearch.DescribeExpressionsInput{}
        output := &cloudsearch.DescribeExpressionsOutput{}

        mockClient.On("DescribeExpressions", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeExpressions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeIndexFields", func(t *testing.T) {
        input := &cloudsearch.DescribeIndexFieldsInput{}
        output := &cloudsearch.DescribeIndexFieldsOutput{}

        mockClient.On("DescribeIndexFields", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeIndexFields(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeScalingParameters", func(t *testing.T) {
        input := &cloudsearch.DescribeScalingParametersInput{}
        output := &cloudsearch.DescribeScalingParametersOutput{}

        mockClient.On("DescribeScalingParameters", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeScalingParameters(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeServiceAccessPolicies", func(t *testing.T) {
        input := &cloudsearch.DescribeServiceAccessPoliciesInput{}
        output := &cloudsearch.DescribeServiceAccessPoliciesOutput{}

        mockClient.On("DescribeServiceAccessPolicies", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeServiceAccessPolicies(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeSuggesters", func(t *testing.T) {
        input := &cloudsearch.DescribeSuggestersInput{}
        output := &cloudsearch.DescribeSuggestersOutput{}

        mockClient.On("DescribeSuggesters", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeSuggesters(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestIndexDocuments", func(t *testing.T) {
        input := &cloudsearch.IndexDocumentsInput{}
        output := &cloudsearch.IndexDocumentsOutput{}

        mockClient.On("IndexDocuments", ctx, input).Return(output, nil)

        result, err := mockClient.IndexDocuments(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDomainNames", func(t *testing.T) {
        input := &cloudsearch.ListDomainNamesInput{}
        output := &cloudsearch.ListDomainNamesOutput{}

        mockClient.On("ListDomainNames", ctx, input).Return(output, nil)

        result, err := mockClient.ListDomainNames(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateAvailabilityOptions", func(t *testing.T) {
        input := &cloudsearch.UpdateAvailabilityOptionsInput{}
        output := &cloudsearch.UpdateAvailabilityOptionsOutput{}

        mockClient.On("UpdateAvailabilityOptions", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateAvailabilityOptions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateDomainEndpointOptions", func(t *testing.T) {
        input := &cloudsearch.UpdateDomainEndpointOptionsInput{}
        output := &cloudsearch.UpdateDomainEndpointOptionsOutput{}

        mockClient.On("UpdateDomainEndpointOptions", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateDomainEndpointOptions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateScalingParameters", func(t *testing.T) {
        input := &cloudsearch.UpdateScalingParametersInput{}
        output := &cloudsearch.UpdateScalingParametersOutput{}

        mockClient.On("UpdateScalingParameters", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateScalingParameters(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateServiceAccessPolicies", func(t *testing.T) {
        input := &cloudsearch.UpdateServiceAccessPoliciesInput{}
        output := &cloudsearch.UpdateServiceAccessPoliciesOutput{}

        mockClient.On("UpdateServiceAccessPolicies", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateServiceAccessPolicies(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
