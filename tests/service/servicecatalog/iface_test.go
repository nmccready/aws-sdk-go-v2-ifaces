package servicecatalog_test

// tests for the servicecatalog service interface for this ../../../service/servicecatalog/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/servicecatalog"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/servicecatalog/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/servicecatalog/servicecatalog_iface"
	"github.com/stretchr/testify/assert"
)

func TestServicecatalogServiceCanBeMocked(t *testing.T) {
	var iface servicecatalog_iface.IClient
	iface = &servicecatalog.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := servicecatalog.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAcceptPortfolioShare", func(t *testing.T) {
        input := &servicecatalog.AcceptPortfolioShareInput{}
        output := &servicecatalog.AcceptPortfolioShareOutput{}

        mockClient.On("AcceptPortfolioShare", ctx, input).Return(output, nil)

        result, err := mockClient.AcceptPortfolioShare(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateBudgetWithResource", func(t *testing.T) {
        input := &servicecatalog.AssociateBudgetWithResourceInput{}
        output := &servicecatalog.AssociateBudgetWithResourceOutput{}

        mockClient.On("AssociateBudgetWithResource", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateBudgetWithResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociatePrincipalWithPortfolio", func(t *testing.T) {
        input := &servicecatalog.AssociatePrincipalWithPortfolioInput{}
        output := &servicecatalog.AssociatePrincipalWithPortfolioOutput{}

        mockClient.On("AssociatePrincipalWithPortfolio", ctx, input).Return(output, nil)

        result, err := mockClient.AssociatePrincipalWithPortfolio(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateProductWithPortfolio", func(t *testing.T) {
        input := &servicecatalog.AssociateProductWithPortfolioInput{}
        output := &servicecatalog.AssociateProductWithPortfolioOutput{}

        mockClient.On("AssociateProductWithPortfolio", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateProductWithPortfolio(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateServiceActionWithProvisioningArtifact", func(t *testing.T) {
        input := &servicecatalog.AssociateServiceActionWithProvisioningArtifactInput{}
        output := &servicecatalog.AssociateServiceActionWithProvisioningArtifactOutput{}

        mockClient.On("AssociateServiceActionWithProvisioningArtifact", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateServiceActionWithProvisioningArtifact(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateTagOptionWithResource", func(t *testing.T) {
        input := &servicecatalog.AssociateTagOptionWithResourceInput{}
        output := &servicecatalog.AssociateTagOptionWithResourceOutput{}

        mockClient.On("AssociateTagOptionWithResource", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateTagOptionWithResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchAssociateServiceActionWithProvisioningArtifact", func(t *testing.T) {
        input := &servicecatalog.BatchAssociateServiceActionWithProvisioningArtifactInput{}
        output := &servicecatalog.BatchAssociateServiceActionWithProvisioningArtifactOutput{}

        mockClient.On("BatchAssociateServiceActionWithProvisioningArtifact", ctx, input).Return(output, nil)

        result, err := mockClient.BatchAssociateServiceActionWithProvisioningArtifact(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchDisassociateServiceActionFromProvisioningArtifact", func(t *testing.T) {
        input := &servicecatalog.BatchDisassociateServiceActionFromProvisioningArtifactInput{}
        output := &servicecatalog.BatchDisassociateServiceActionFromProvisioningArtifactOutput{}

        mockClient.On("BatchDisassociateServiceActionFromProvisioningArtifact", ctx, input).Return(output, nil)

        result, err := mockClient.BatchDisassociateServiceActionFromProvisioningArtifact(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCopyProduct", func(t *testing.T) {
        input := &servicecatalog.CopyProductInput{}
        output := &servicecatalog.CopyProductOutput{}

        mockClient.On("CopyProduct", ctx, input).Return(output, nil)

        result, err := mockClient.CopyProduct(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateConstraint", func(t *testing.T) {
        input := &servicecatalog.CreateConstraintInput{}
        output := &servicecatalog.CreateConstraintOutput{}

        mockClient.On("CreateConstraint", ctx, input).Return(output, nil)

        result, err := mockClient.CreateConstraint(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreatePortfolio", func(t *testing.T) {
        input := &servicecatalog.CreatePortfolioInput{}
        output := &servicecatalog.CreatePortfolioOutput{}

        mockClient.On("CreatePortfolio", ctx, input).Return(output, nil)

        result, err := mockClient.CreatePortfolio(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreatePortfolioShare", func(t *testing.T) {
        input := &servicecatalog.CreatePortfolioShareInput{}
        output := &servicecatalog.CreatePortfolioShareOutput{}

        mockClient.On("CreatePortfolioShare", ctx, input).Return(output, nil)

        result, err := mockClient.CreatePortfolioShare(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateProduct", func(t *testing.T) {
        input := &servicecatalog.CreateProductInput{}
        output := &servicecatalog.CreateProductOutput{}

        mockClient.On("CreateProduct", ctx, input).Return(output, nil)

        result, err := mockClient.CreateProduct(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateProvisionedProductPlan", func(t *testing.T) {
        input := &servicecatalog.CreateProvisionedProductPlanInput{}
        output := &servicecatalog.CreateProvisionedProductPlanOutput{}

        mockClient.On("CreateProvisionedProductPlan", ctx, input).Return(output, nil)

        result, err := mockClient.CreateProvisionedProductPlan(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateProvisioningArtifact", func(t *testing.T) {
        input := &servicecatalog.CreateProvisioningArtifactInput{}
        output := &servicecatalog.CreateProvisioningArtifactOutput{}

        mockClient.On("CreateProvisioningArtifact", ctx, input).Return(output, nil)

        result, err := mockClient.CreateProvisioningArtifact(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateServiceAction", func(t *testing.T) {
        input := &servicecatalog.CreateServiceActionInput{}
        output := &servicecatalog.CreateServiceActionOutput{}

        mockClient.On("CreateServiceAction", ctx, input).Return(output, nil)

        result, err := mockClient.CreateServiceAction(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateTagOption", func(t *testing.T) {
        input := &servicecatalog.CreateTagOptionInput{}
        output := &servicecatalog.CreateTagOptionOutput{}

        mockClient.On("CreateTagOption", ctx, input).Return(output, nil)

        result, err := mockClient.CreateTagOption(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteConstraint", func(t *testing.T) {
        input := &servicecatalog.DeleteConstraintInput{}
        output := &servicecatalog.DeleteConstraintOutput{}

        mockClient.On("DeleteConstraint", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteConstraint(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeletePortfolio", func(t *testing.T) {
        input := &servicecatalog.DeletePortfolioInput{}
        output := &servicecatalog.DeletePortfolioOutput{}

        mockClient.On("DeletePortfolio", ctx, input).Return(output, nil)

        result, err := mockClient.DeletePortfolio(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeletePortfolioShare", func(t *testing.T) {
        input := &servicecatalog.DeletePortfolioShareInput{}
        output := &servicecatalog.DeletePortfolioShareOutput{}

        mockClient.On("DeletePortfolioShare", ctx, input).Return(output, nil)

        result, err := mockClient.DeletePortfolioShare(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteProduct", func(t *testing.T) {
        input := &servicecatalog.DeleteProductInput{}
        output := &servicecatalog.DeleteProductOutput{}

        mockClient.On("DeleteProduct", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteProduct(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteProvisionedProductPlan", func(t *testing.T) {
        input := &servicecatalog.DeleteProvisionedProductPlanInput{}
        output := &servicecatalog.DeleteProvisionedProductPlanOutput{}

        mockClient.On("DeleteProvisionedProductPlan", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteProvisionedProductPlan(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteProvisioningArtifact", func(t *testing.T) {
        input := &servicecatalog.DeleteProvisioningArtifactInput{}
        output := &servicecatalog.DeleteProvisioningArtifactOutput{}

        mockClient.On("DeleteProvisioningArtifact", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteProvisioningArtifact(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteServiceAction", func(t *testing.T) {
        input := &servicecatalog.DeleteServiceActionInput{}
        output := &servicecatalog.DeleteServiceActionOutput{}

        mockClient.On("DeleteServiceAction", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteServiceAction(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteTagOption", func(t *testing.T) {
        input := &servicecatalog.DeleteTagOptionInput{}
        output := &servicecatalog.DeleteTagOptionOutput{}

        mockClient.On("DeleteTagOption", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteTagOption(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeConstraint", func(t *testing.T) {
        input := &servicecatalog.DescribeConstraintInput{}
        output := &servicecatalog.DescribeConstraintOutput{}

        mockClient.On("DescribeConstraint", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeConstraint(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeCopyProductStatus", func(t *testing.T) {
        input := &servicecatalog.DescribeCopyProductStatusInput{}
        output := &servicecatalog.DescribeCopyProductStatusOutput{}

        mockClient.On("DescribeCopyProductStatus", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeCopyProductStatus(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribePortfolio", func(t *testing.T) {
        input := &servicecatalog.DescribePortfolioInput{}
        output := &servicecatalog.DescribePortfolioOutput{}

        mockClient.On("DescribePortfolio", ctx, input).Return(output, nil)

        result, err := mockClient.DescribePortfolio(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribePortfolioShareStatus", func(t *testing.T) {
        input := &servicecatalog.DescribePortfolioShareStatusInput{}
        output := &servicecatalog.DescribePortfolioShareStatusOutput{}

        mockClient.On("DescribePortfolioShareStatus", ctx, input).Return(output, nil)

        result, err := mockClient.DescribePortfolioShareStatus(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribePortfolioShares", func(t *testing.T) {
        input := &servicecatalog.DescribePortfolioSharesInput{}
        output := &servicecatalog.DescribePortfolioSharesOutput{}

        mockClient.On("DescribePortfolioShares", ctx, input).Return(output, nil)

        result, err := mockClient.DescribePortfolioShares(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeProduct", func(t *testing.T) {
        input := &servicecatalog.DescribeProductInput{}
        output := &servicecatalog.DescribeProductOutput{}

        mockClient.On("DescribeProduct", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeProduct(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeProductAsAdmin", func(t *testing.T) {
        input := &servicecatalog.DescribeProductAsAdminInput{}
        output := &servicecatalog.DescribeProductAsAdminOutput{}

        mockClient.On("DescribeProductAsAdmin", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeProductAsAdmin(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeProductView", func(t *testing.T) {
        input := &servicecatalog.DescribeProductViewInput{}
        output := &servicecatalog.DescribeProductViewOutput{}

        mockClient.On("DescribeProductView", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeProductView(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeProvisionedProduct", func(t *testing.T) {
        input := &servicecatalog.DescribeProvisionedProductInput{}
        output := &servicecatalog.DescribeProvisionedProductOutput{}

        mockClient.On("DescribeProvisionedProduct", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeProvisionedProduct(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeProvisionedProductPlan", func(t *testing.T) {
        input := &servicecatalog.DescribeProvisionedProductPlanInput{}
        output := &servicecatalog.DescribeProvisionedProductPlanOutput{}

        mockClient.On("DescribeProvisionedProductPlan", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeProvisionedProductPlan(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeProvisioningArtifact", func(t *testing.T) {
        input := &servicecatalog.DescribeProvisioningArtifactInput{}
        output := &servicecatalog.DescribeProvisioningArtifactOutput{}

        mockClient.On("DescribeProvisioningArtifact", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeProvisioningArtifact(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeProvisioningParameters", func(t *testing.T) {
        input := &servicecatalog.DescribeProvisioningParametersInput{}
        output := &servicecatalog.DescribeProvisioningParametersOutput{}

        mockClient.On("DescribeProvisioningParameters", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeProvisioningParameters(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeRecord", func(t *testing.T) {
        input := &servicecatalog.DescribeRecordInput{}
        output := &servicecatalog.DescribeRecordOutput{}

        mockClient.On("DescribeRecord", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeRecord(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeServiceAction", func(t *testing.T) {
        input := &servicecatalog.DescribeServiceActionInput{}
        output := &servicecatalog.DescribeServiceActionOutput{}

        mockClient.On("DescribeServiceAction", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeServiceAction(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeServiceActionExecutionParameters", func(t *testing.T) {
        input := &servicecatalog.DescribeServiceActionExecutionParametersInput{}
        output := &servicecatalog.DescribeServiceActionExecutionParametersOutput{}

        mockClient.On("DescribeServiceActionExecutionParameters", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeServiceActionExecutionParameters(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeTagOption", func(t *testing.T) {
        input := &servicecatalog.DescribeTagOptionInput{}
        output := &servicecatalog.DescribeTagOptionOutput{}

        mockClient.On("DescribeTagOption", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeTagOption(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisableAWSOrganizationsAccess", func(t *testing.T) {
        input := &servicecatalog.DisableAWSOrganizationsAccessInput{}
        output := &servicecatalog.DisableAWSOrganizationsAccessOutput{}

        mockClient.On("DisableAWSOrganizationsAccess", ctx, input).Return(output, nil)

        result, err := mockClient.DisableAWSOrganizationsAccess(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateBudgetFromResource", func(t *testing.T) {
        input := &servicecatalog.DisassociateBudgetFromResourceInput{}
        output := &servicecatalog.DisassociateBudgetFromResourceOutput{}

        mockClient.On("DisassociateBudgetFromResource", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateBudgetFromResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociatePrincipalFromPortfolio", func(t *testing.T) {
        input := &servicecatalog.DisassociatePrincipalFromPortfolioInput{}
        output := &servicecatalog.DisassociatePrincipalFromPortfolioOutput{}

        mockClient.On("DisassociatePrincipalFromPortfolio", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociatePrincipalFromPortfolio(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateProductFromPortfolio", func(t *testing.T) {
        input := &servicecatalog.DisassociateProductFromPortfolioInput{}
        output := &servicecatalog.DisassociateProductFromPortfolioOutput{}

        mockClient.On("DisassociateProductFromPortfolio", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateProductFromPortfolio(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateServiceActionFromProvisioningArtifact", func(t *testing.T) {
        input := &servicecatalog.DisassociateServiceActionFromProvisioningArtifactInput{}
        output := &servicecatalog.DisassociateServiceActionFromProvisioningArtifactOutput{}

        mockClient.On("DisassociateServiceActionFromProvisioningArtifact", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateServiceActionFromProvisioningArtifact(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateTagOptionFromResource", func(t *testing.T) {
        input := &servicecatalog.DisassociateTagOptionFromResourceInput{}
        output := &servicecatalog.DisassociateTagOptionFromResourceOutput{}

        mockClient.On("DisassociateTagOptionFromResource", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateTagOptionFromResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestEnableAWSOrganizationsAccess", func(t *testing.T) {
        input := &servicecatalog.EnableAWSOrganizationsAccessInput{}
        output := &servicecatalog.EnableAWSOrganizationsAccessOutput{}

        mockClient.On("EnableAWSOrganizationsAccess", ctx, input).Return(output, nil)

        result, err := mockClient.EnableAWSOrganizationsAccess(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestExecuteProvisionedProductPlan", func(t *testing.T) {
        input := &servicecatalog.ExecuteProvisionedProductPlanInput{}
        output := &servicecatalog.ExecuteProvisionedProductPlanOutput{}

        mockClient.On("ExecuteProvisionedProductPlan", ctx, input).Return(output, nil)

        result, err := mockClient.ExecuteProvisionedProductPlan(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestExecuteProvisionedProductServiceAction", func(t *testing.T) {
        input := &servicecatalog.ExecuteProvisionedProductServiceActionInput{}
        output := &servicecatalog.ExecuteProvisionedProductServiceActionOutput{}

        mockClient.On("ExecuteProvisionedProductServiceAction", ctx, input).Return(output, nil)

        result, err := mockClient.ExecuteProvisionedProductServiceAction(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetAWSOrganizationsAccessStatus", func(t *testing.T) {
        input := &servicecatalog.GetAWSOrganizationsAccessStatusInput{}
        output := &servicecatalog.GetAWSOrganizationsAccessStatusOutput{}

        mockClient.On("GetAWSOrganizationsAccessStatus", ctx, input).Return(output, nil)

        result, err := mockClient.GetAWSOrganizationsAccessStatus(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetProvisionedProductOutputs", func(t *testing.T) {
        input := &servicecatalog.GetProvisionedProductOutputsInput{}
        output := &servicecatalog.GetProvisionedProductOutputsOutput{}

        mockClient.On("GetProvisionedProductOutputs", ctx, input).Return(output, nil)

        result, err := mockClient.GetProvisionedProductOutputs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestImportAsProvisionedProduct", func(t *testing.T) {
        input := &servicecatalog.ImportAsProvisionedProductInput{}
        output := &servicecatalog.ImportAsProvisionedProductOutput{}

        mockClient.On("ImportAsProvisionedProduct", ctx, input).Return(output, nil)

        result, err := mockClient.ImportAsProvisionedProduct(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAcceptedPortfolioShares", func(t *testing.T) {
        input := &servicecatalog.ListAcceptedPortfolioSharesInput{}
        output := &servicecatalog.ListAcceptedPortfolioSharesOutput{}

        mockClient.On("ListAcceptedPortfolioShares", ctx, input).Return(output, nil)

        result, err := mockClient.ListAcceptedPortfolioShares(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListBudgetsForResource", func(t *testing.T) {
        input := &servicecatalog.ListBudgetsForResourceInput{}
        output := &servicecatalog.ListBudgetsForResourceOutput{}

        mockClient.On("ListBudgetsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListBudgetsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListConstraintsForPortfolio", func(t *testing.T) {
        input := &servicecatalog.ListConstraintsForPortfolioInput{}
        output := &servicecatalog.ListConstraintsForPortfolioOutput{}

        mockClient.On("ListConstraintsForPortfolio", ctx, input).Return(output, nil)

        result, err := mockClient.ListConstraintsForPortfolio(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListLaunchPaths", func(t *testing.T) {
        input := &servicecatalog.ListLaunchPathsInput{}
        output := &servicecatalog.ListLaunchPathsOutput{}

        mockClient.On("ListLaunchPaths", ctx, input).Return(output, nil)

        result, err := mockClient.ListLaunchPaths(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListOrganizationPortfolioAccess", func(t *testing.T) {
        input := &servicecatalog.ListOrganizationPortfolioAccessInput{}
        output := &servicecatalog.ListOrganizationPortfolioAccessOutput{}

        mockClient.On("ListOrganizationPortfolioAccess", ctx, input).Return(output, nil)

        result, err := mockClient.ListOrganizationPortfolioAccess(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListPortfolioAccess", func(t *testing.T) {
        input := &servicecatalog.ListPortfolioAccessInput{}
        output := &servicecatalog.ListPortfolioAccessOutput{}

        mockClient.On("ListPortfolioAccess", ctx, input).Return(output, nil)

        result, err := mockClient.ListPortfolioAccess(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListPortfolios", func(t *testing.T) {
        input := &servicecatalog.ListPortfoliosInput{}
        output := &servicecatalog.ListPortfoliosOutput{}

        mockClient.On("ListPortfolios", ctx, input).Return(output, nil)

        result, err := mockClient.ListPortfolios(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListPortfoliosForProduct", func(t *testing.T) {
        input := &servicecatalog.ListPortfoliosForProductInput{}
        output := &servicecatalog.ListPortfoliosForProductOutput{}

        mockClient.On("ListPortfoliosForProduct", ctx, input).Return(output, nil)

        result, err := mockClient.ListPortfoliosForProduct(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListPrincipalsForPortfolio", func(t *testing.T) {
        input := &servicecatalog.ListPrincipalsForPortfolioInput{}
        output := &servicecatalog.ListPrincipalsForPortfolioOutput{}

        mockClient.On("ListPrincipalsForPortfolio", ctx, input).Return(output, nil)

        result, err := mockClient.ListPrincipalsForPortfolio(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListProvisionedProductPlans", func(t *testing.T) {
        input := &servicecatalog.ListProvisionedProductPlansInput{}
        output := &servicecatalog.ListProvisionedProductPlansOutput{}

        mockClient.On("ListProvisionedProductPlans", ctx, input).Return(output, nil)

        result, err := mockClient.ListProvisionedProductPlans(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListProvisioningArtifacts", func(t *testing.T) {
        input := &servicecatalog.ListProvisioningArtifactsInput{}
        output := &servicecatalog.ListProvisioningArtifactsOutput{}

        mockClient.On("ListProvisioningArtifacts", ctx, input).Return(output, nil)

        result, err := mockClient.ListProvisioningArtifacts(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListProvisioningArtifactsForServiceAction", func(t *testing.T) {
        input := &servicecatalog.ListProvisioningArtifactsForServiceActionInput{}
        output := &servicecatalog.ListProvisioningArtifactsForServiceActionOutput{}

        mockClient.On("ListProvisioningArtifactsForServiceAction", ctx, input).Return(output, nil)

        result, err := mockClient.ListProvisioningArtifactsForServiceAction(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListRecordHistory", func(t *testing.T) {
        input := &servicecatalog.ListRecordHistoryInput{}
        output := &servicecatalog.ListRecordHistoryOutput{}

        mockClient.On("ListRecordHistory", ctx, input).Return(output, nil)

        result, err := mockClient.ListRecordHistory(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListResourcesForTagOption", func(t *testing.T) {
        input := &servicecatalog.ListResourcesForTagOptionInput{}
        output := &servicecatalog.ListResourcesForTagOptionOutput{}

        mockClient.On("ListResourcesForTagOption", ctx, input).Return(output, nil)

        result, err := mockClient.ListResourcesForTagOption(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListServiceActions", func(t *testing.T) {
        input := &servicecatalog.ListServiceActionsInput{}
        output := &servicecatalog.ListServiceActionsOutput{}

        mockClient.On("ListServiceActions", ctx, input).Return(output, nil)

        result, err := mockClient.ListServiceActions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListServiceActionsForProvisioningArtifact", func(t *testing.T) {
        input := &servicecatalog.ListServiceActionsForProvisioningArtifactInput{}
        output := &servicecatalog.ListServiceActionsForProvisioningArtifactOutput{}

        mockClient.On("ListServiceActionsForProvisioningArtifact", ctx, input).Return(output, nil)

        result, err := mockClient.ListServiceActionsForProvisioningArtifact(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListStackInstancesForProvisionedProduct", func(t *testing.T) {
        input := &servicecatalog.ListStackInstancesForProvisionedProductInput{}
        output := &servicecatalog.ListStackInstancesForProvisionedProductOutput{}

        mockClient.On("ListStackInstancesForProvisionedProduct", ctx, input).Return(output, nil)

        result, err := mockClient.ListStackInstancesForProvisionedProduct(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagOptions", func(t *testing.T) {
        input := &servicecatalog.ListTagOptionsInput{}
        output := &servicecatalog.ListTagOptionsOutput{}

        mockClient.On("ListTagOptions", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagOptions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestNotifyProvisionProductEngineWorkflowResult", func(t *testing.T) {
        input := &servicecatalog.NotifyProvisionProductEngineWorkflowResultInput{}
        output := &servicecatalog.NotifyProvisionProductEngineWorkflowResultOutput{}

        mockClient.On("NotifyProvisionProductEngineWorkflowResult", ctx, input).Return(output, nil)

        result, err := mockClient.NotifyProvisionProductEngineWorkflowResult(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestNotifyTerminateProvisionedProductEngineWorkflowResult", func(t *testing.T) {
        input := &servicecatalog.NotifyTerminateProvisionedProductEngineWorkflowResultInput{}
        output := &servicecatalog.NotifyTerminateProvisionedProductEngineWorkflowResultOutput{}

        mockClient.On("NotifyTerminateProvisionedProductEngineWorkflowResult", ctx, input).Return(output, nil)

        result, err := mockClient.NotifyTerminateProvisionedProductEngineWorkflowResult(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestNotifyUpdateProvisionedProductEngineWorkflowResult", func(t *testing.T) {
        input := &servicecatalog.NotifyUpdateProvisionedProductEngineWorkflowResultInput{}
        output := &servicecatalog.NotifyUpdateProvisionedProductEngineWorkflowResultOutput{}

        mockClient.On("NotifyUpdateProvisionedProductEngineWorkflowResult", ctx, input).Return(output, nil)

        result, err := mockClient.NotifyUpdateProvisionedProductEngineWorkflowResult(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestProvisionProduct", func(t *testing.T) {
        input := &servicecatalog.ProvisionProductInput{}
        output := &servicecatalog.ProvisionProductOutput{}

        mockClient.On("ProvisionProduct", ctx, input).Return(output, nil)

        result, err := mockClient.ProvisionProduct(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRejectPortfolioShare", func(t *testing.T) {
        input := &servicecatalog.RejectPortfolioShareInput{}
        output := &servicecatalog.RejectPortfolioShareOutput{}

        mockClient.On("RejectPortfolioShare", ctx, input).Return(output, nil)

        result, err := mockClient.RejectPortfolioShare(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestScanProvisionedProducts", func(t *testing.T) {
        input := &servicecatalog.ScanProvisionedProductsInput{}
        output := &servicecatalog.ScanProvisionedProductsOutput{}

        mockClient.On("ScanProvisionedProducts", ctx, input).Return(output, nil)

        result, err := mockClient.ScanProvisionedProducts(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSearchProducts", func(t *testing.T) {
        input := &servicecatalog.SearchProductsInput{}
        output := &servicecatalog.SearchProductsOutput{}

        mockClient.On("SearchProducts", ctx, input).Return(output, nil)

        result, err := mockClient.SearchProducts(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSearchProductsAsAdmin", func(t *testing.T) {
        input := &servicecatalog.SearchProductsAsAdminInput{}
        output := &servicecatalog.SearchProductsAsAdminOutput{}

        mockClient.On("SearchProductsAsAdmin", ctx, input).Return(output, nil)

        result, err := mockClient.SearchProductsAsAdmin(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSearchProvisionedProducts", func(t *testing.T) {
        input := &servicecatalog.SearchProvisionedProductsInput{}
        output := &servicecatalog.SearchProvisionedProductsOutput{}

        mockClient.On("SearchProvisionedProducts", ctx, input).Return(output, nil)

        result, err := mockClient.SearchProvisionedProducts(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTerminateProvisionedProduct", func(t *testing.T) {
        input := &servicecatalog.TerminateProvisionedProductInput{}
        output := &servicecatalog.TerminateProvisionedProductOutput{}

        mockClient.On("TerminateProvisionedProduct", ctx, input).Return(output, nil)

        result, err := mockClient.TerminateProvisionedProduct(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateConstraint", func(t *testing.T) {
        input := &servicecatalog.UpdateConstraintInput{}
        output := &servicecatalog.UpdateConstraintOutput{}

        mockClient.On("UpdateConstraint", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateConstraint(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdatePortfolio", func(t *testing.T) {
        input := &servicecatalog.UpdatePortfolioInput{}
        output := &servicecatalog.UpdatePortfolioOutput{}

        mockClient.On("UpdatePortfolio", ctx, input).Return(output, nil)

        result, err := mockClient.UpdatePortfolio(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdatePortfolioShare", func(t *testing.T) {
        input := &servicecatalog.UpdatePortfolioShareInput{}
        output := &servicecatalog.UpdatePortfolioShareOutput{}

        mockClient.On("UpdatePortfolioShare", ctx, input).Return(output, nil)

        result, err := mockClient.UpdatePortfolioShare(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateProduct", func(t *testing.T) {
        input := &servicecatalog.UpdateProductInput{}
        output := &servicecatalog.UpdateProductOutput{}

        mockClient.On("UpdateProduct", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateProduct(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateProvisionedProduct", func(t *testing.T) {
        input := &servicecatalog.UpdateProvisionedProductInput{}
        output := &servicecatalog.UpdateProvisionedProductOutput{}

        mockClient.On("UpdateProvisionedProduct", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateProvisionedProduct(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateProvisionedProductProperties", func(t *testing.T) {
        input := &servicecatalog.UpdateProvisionedProductPropertiesInput{}
        output := &servicecatalog.UpdateProvisionedProductPropertiesOutput{}

        mockClient.On("UpdateProvisionedProductProperties", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateProvisionedProductProperties(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateProvisioningArtifact", func(t *testing.T) {
        input := &servicecatalog.UpdateProvisioningArtifactInput{}
        output := &servicecatalog.UpdateProvisioningArtifactOutput{}

        mockClient.On("UpdateProvisioningArtifact", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateProvisioningArtifact(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateServiceAction", func(t *testing.T) {
        input := &servicecatalog.UpdateServiceActionInput{}
        output := &servicecatalog.UpdateServiceActionOutput{}

        mockClient.On("UpdateServiceAction", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateServiceAction(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateTagOption", func(t *testing.T) {
        input := &servicecatalog.UpdateTagOptionInput{}
        output := &servicecatalog.UpdateTagOptionOutput{}

        mockClient.On("UpdateTagOption", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateTagOption(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
