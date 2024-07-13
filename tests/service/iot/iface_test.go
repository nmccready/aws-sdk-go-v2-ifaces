package iot_test

// tests for the iot service interface for this ../../../service/iot/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/iot"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/iot/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/iot/iot_iface"
	"github.com/stretchr/testify/assert"
)

func TestIotServiceCanBeMocked(t *testing.T) {
	var iface iot_iface.IClient
	iface = &iot.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := iot.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAcceptCertificateTransfer", func(t *testing.T) {
        input := &iot.AcceptCertificateTransferInput{}
        output := &iot.AcceptCertificateTransferOutput{}

        mockClient.On("AcceptCertificateTransfer", ctx, input).Return(output, nil)

        result, err := mockClient.AcceptCertificateTransfer(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAddThingToBillingGroup", func(t *testing.T) {
        input := &iot.AddThingToBillingGroupInput{}
        output := &iot.AddThingToBillingGroupOutput{}

        mockClient.On("AddThingToBillingGroup", ctx, input).Return(output, nil)

        result, err := mockClient.AddThingToBillingGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAddThingToThingGroup", func(t *testing.T) {
        input := &iot.AddThingToThingGroupInput{}
        output := &iot.AddThingToThingGroupOutput{}

        mockClient.On("AddThingToThingGroup", ctx, input).Return(output, nil)

        result, err := mockClient.AddThingToThingGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateTargetsWithJob", func(t *testing.T) {
        input := &iot.AssociateTargetsWithJobInput{}
        output := &iot.AssociateTargetsWithJobOutput{}

        mockClient.On("AssociateTargetsWithJob", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateTargetsWithJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAttachPolicy", func(t *testing.T) {
        input := &iot.AttachPolicyInput{}
        output := &iot.AttachPolicyOutput{}

        mockClient.On("AttachPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.AttachPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAttachPrincipalPolicy", func(t *testing.T) {
        input := &iot.AttachPrincipalPolicyInput{}
        output := &iot.AttachPrincipalPolicyOutput{}

        mockClient.On("AttachPrincipalPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.AttachPrincipalPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAttachSecurityProfile", func(t *testing.T) {
        input := &iot.AttachSecurityProfileInput{}
        output := &iot.AttachSecurityProfileOutput{}

        mockClient.On("AttachSecurityProfile", ctx, input).Return(output, nil)

        result, err := mockClient.AttachSecurityProfile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAttachThingPrincipal", func(t *testing.T) {
        input := &iot.AttachThingPrincipalInput{}
        output := &iot.AttachThingPrincipalOutput{}

        mockClient.On("AttachThingPrincipal", ctx, input).Return(output, nil)

        result, err := mockClient.AttachThingPrincipal(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCancelAuditMitigationActionsTask", func(t *testing.T) {
        input := &iot.CancelAuditMitigationActionsTaskInput{}
        output := &iot.CancelAuditMitigationActionsTaskOutput{}

        mockClient.On("CancelAuditMitigationActionsTask", ctx, input).Return(output, nil)

        result, err := mockClient.CancelAuditMitigationActionsTask(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCancelAuditTask", func(t *testing.T) {
        input := &iot.CancelAuditTaskInput{}
        output := &iot.CancelAuditTaskOutput{}

        mockClient.On("CancelAuditTask", ctx, input).Return(output, nil)

        result, err := mockClient.CancelAuditTask(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCancelCertificateTransfer", func(t *testing.T) {
        input := &iot.CancelCertificateTransferInput{}
        output := &iot.CancelCertificateTransferOutput{}

        mockClient.On("CancelCertificateTransfer", ctx, input).Return(output, nil)

        result, err := mockClient.CancelCertificateTransfer(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCancelDetectMitigationActionsTask", func(t *testing.T) {
        input := &iot.CancelDetectMitigationActionsTaskInput{}
        output := &iot.CancelDetectMitigationActionsTaskOutput{}

        mockClient.On("CancelDetectMitigationActionsTask", ctx, input).Return(output, nil)

        result, err := mockClient.CancelDetectMitigationActionsTask(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCancelJob", func(t *testing.T) {
        input := &iot.CancelJobInput{}
        output := &iot.CancelJobOutput{}

        mockClient.On("CancelJob", ctx, input).Return(output, nil)

        result, err := mockClient.CancelJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCancelJobExecution", func(t *testing.T) {
        input := &iot.CancelJobExecutionInput{}
        output := &iot.CancelJobExecutionOutput{}

        mockClient.On("CancelJobExecution", ctx, input).Return(output, nil)

        result, err := mockClient.CancelJobExecution(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestClearDefaultAuthorizer", func(t *testing.T) {
        input := &iot.ClearDefaultAuthorizerInput{}
        output := &iot.ClearDefaultAuthorizerOutput{}

        mockClient.On("ClearDefaultAuthorizer", ctx, input).Return(output, nil)

        result, err := mockClient.ClearDefaultAuthorizer(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestConfirmTopicRuleDestination", func(t *testing.T) {
        input := &iot.ConfirmTopicRuleDestinationInput{}
        output := &iot.ConfirmTopicRuleDestinationOutput{}

        mockClient.On("ConfirmTopicRuleDestination", ctx, input).Return(output, nil)

        result, err := mockClient.ConfirmTopicRuleDestination(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateAuditSuppression", func(t *testing.T) {
        input := &iot.CreateAuditSuppressionInput{}
        output := &iot.CreateAuditSuppressionOutput{}

        mockClient.On("CreateAuditSuppression", ctx, input).Return(output, nil)

        result, err := mockClient.CreateAuditSuppression(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateAuthorizer", func(t *testing.T) {
        input := &iot.CreateAuthorizerInput{}
        output := &iot.CreateAuthorizerOutput{}

        mockClient.On("CreateAuthorizer", ctx, input).Return(output, nil)

        result, err := mockClient.CreateAuthorizer(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateBillingGroup", func(t *testing.T) {
        input := &iot.CreateBillingGroupInput{}
        output := &iot.CreateBillingGroupOutput{}

        mockClient.On("CreateBillingGroup", ctx, input).Return(output, nil)

        result, err := mockClient.CreateBillingGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateCertificateFromCsr", func(t *testing.T) {
        input := &iot.CreateCertificateFromCsrInput{}
        output := &iot.CreateCertificateFromCsrOutput{}

        mockClient.On("CreateCertificateFromCsr", ctx, input).Return(output, nil)

        result, err := mockClient.CreateCertificateFromCsr(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateCertificateProvider", func(t *testing.T) {
        input := &iot.CreateCertificateProviderInput{}
        output := &iot.CreateCertificateProviderOutput{}

        mockClient.On("CreateCertificateProvider", ctx, input).Return(output, nil)

        result, err := mockClient.CreateCertificateProvider(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateCustomMetric", func(t *testing.T) {
        input := &iot.CreateCustomMetricInput{}
        output := &iot.CreateCustomMetricOutput{}

        mockClient.On("CreateCustomMetric", ctx, input).Return(output, nil)

        result, err := mockClient.CreateCustomMetric(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateDimension", func(t *testing.T) {
        input := &iot.CreateDimensionInput{}
        output := &iot.CreateDimensionOutput{}

        mockClient.On("CreateDimension", ctx, input).Return(output, nil)

        result, err := mockClient.CreateDimension(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateDomainConfiguration", func(t *testing.T) {
        input := &iot.CreateDomainConfigurationInput{}
        output := &iot.CreateDomainConfigurationOutput{}

        mockClient.On("CreateDomainConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.CreateDomainConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateDynamicThingGroup", func(t *testing.T) {
        input := &iot.CreateDynamicThingGroupInput{}
        output := &iot.CreateDynamicThingGroupOutput{}

        mockClient.On("CreateDynamicThingGroup", ctx, input).Return(output, nil)

        result, err := mockClient.CreateDynamicThingGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateFleetMetric", func(t *testing.T) {
        input := &iot.CreateFleetMetricInput{}
        output := &iot.CreateFleetMetricOutput{}

        mockClient.On("CreateFleetMetric", ctx, input).Return(output, nil)

        result, err := mockClient.CreateFleetMetric(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateJob", func(t *testing.T) {
        input := &iot.CreateJobInput{}
        output := &iot.CreateJobOutput{}

        mockClient.On("CreateJob", ctx, input).Return(output, nil)

        result, err := mockClient.CreateJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateJobTemplate", func(t *testing.T) {
        input := &iot.CreateJobTemplateInput{}
        output := &iot.CreateJobTemplateOutput{}

        mockClient.On("CreateJobTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.CreateJobTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateKeysAndCertificate", func(t *testing.T) {
        input := &iot.CreateKeysAndCertificateInput{}
        output := &iot.CreateKeysAndCertificateOutput{}

        mockClient.On("CreateKeysAndCertificate", ctx, input).Return(output, nil)

        result, err := mockClient.CreateKeysAndCertificate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateMitigationAction", func(t *testing.T) {
        input := &iot.CreateMitigationActionInput{}
        output := &iot.CreateMitigationActionOutput{}

        mockClient.On("CreateMitigationAction", ctx, input).Return(output, nil)

        result, err := mockClient.CreateMitigationAction(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateOTAUpdate", func(t *testing.T) {
        input := &iot.CreateOTAUpdateInput{}
        output := &iot.CreateOTAUpdateOutput{}

        mockClient.On("CreateOTAUpdate", ctx, input).Return(output, nil)

        result, err := mockClient.CreateOTAUpdate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreatePackage", func(t *testing.T) {
        input := &iot.CreatePackageInput{}
        output := &iot.CreatePackageOutput{}

        mockClient.On("CreatePackage", ctx, input).Return(output, nil)

        result, err := mockClient.CreatePackage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreatePackageVersion", func(t *testing.T) {
        input := &iot.CreatePackageVersionInput{}
        output := &iot.CreatePackageVersionOutput{}

        mockClient.On("CreatePackageVersion", ctx, input).Return(output, nil)

        result, err := mockClient.CreatePackageVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreatePolicy", func(t *testing.T) {
        input := &iot.CreatePolicyInput{}
        output := &iot.CreatePolicyOutput{}

        mockClient.On("CreatePolicy", ctx, input).Return(output, nil)

        result, err := mockClient.CreatePolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreatePolicyVersion", func(t *testing.T) {
        input := &iot.CreatePolicyVersionInput{}
        output := &iot.CreatePolicyVersionOutput{}

        mockClient.On("CreatePolicyVersion", ctx, input).Return(output, nil)

        result, err := mockClient.CreatePolicyVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateProvisioningClaim", func(t *testing.T) {
        input := &iot.CreateProvisioningClaimInput{}
        output := &iot.CreateProvisioningClaimOutput{}

        mockClient.On("CreateProvisioningClaim", ctx, input).Return(output, nil)

        result, err := mockClient.CreateProvisioningClaim(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateProvisioningTemplate", func(t *testing.T) {
        input := &iot.CreateProvisioningTemplateInput{}
        output := &iot.CreateProvisioningTemplateOutput{}

        mockClient.On("CreateProvisioningTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.CreateProvisioningTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateProvisioningTemplateVersion", func(t *testing.T) {
        input := &iot.CreateProvisioningTemplateVersionInput{}
        output := &iot.CreateProvisioningTemplateVersionOutput{}

        mockClient.On("CreateProvisioningTemplateVersion", ctx, input).Return(output, nil)

        result, err := mockClient.CreateProvisioningTemplateVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateRoleAlias", func(t *testing.T) {
        input := &iot.CreateRoleAliasInput{}
        output := &iot.CreateRoleAliasOutput{}

        mockClient.On("CreateRoleAlias", ctx, input).Return(output, nil)

        result, err := mockClient.CreateRoleAlias(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateScheduledAudit", func(t *testing.T) {
        input := &iot.CreateScheduledAuditInput{}
        output := &iot.CreateScheduledAuditOutput{}

        mockClient.On("CreateScheduledAudit", ctx, input).Return(output, nil)

        result, err := mockClient.CreateScheduledAudit(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateSecurityProfile", func(t *testing.T) {
        input := &iot.CreateSecurityProfileInput{}
        output := &iot.CreateSecurityProfileOutput{}

        mockClient.On("CreateSecurityProfile", ctx, input).Return(output, nil)

        result, err := mockClient.CreateSecurityProfile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateStream", func(t *testing.T) {
        input := &iot.CreateStreamInput{}
        output := &iot.CreateStreamOutput{}

        mockClient.On("CreateStream", ctx, input).Return(output, nil)

        result, err := mockClient.CreateStream(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateThing", func(t *testing.T) {
        input := &iot.CreateThingInput{}
        output := &iot.CreateThingOutput{}

        mockClient.On("CreateThing", ctx, input).Return(output, nil)

        result, err := mockClient.CreateThing(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateThingGroup", func(t *testing.T) {
        input := &iot.CreateThingGroupInput{}
        output := &iot.CreateThingGroupOutput{}

        mockClient.On("CreateThingGroup", ctx, input).Return(output, nil)

        result, err := mockClient.CreateThingGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateThingType", func(t *testing.T) {
        input := &iot.CreateThingTypeInput{}
        output := &iot.CreateThingTypeOutput{}

        mockClient.On("CreateThingType", ctx, input).Return(output, nil)

        result, err := mockClient.CreateThingType(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateTopicRule", func(t *testing.T) {
        input := &iot.CreateTopicRuleInput{}
        output := &iot.CreateTopicRuleOutput{}

        mockClient.On("CreateTopicRule", ctx, input).Return(output, nil)

        result, err := mockClient.CreateTopicRule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateTopicRuleDestination", func(t *testing.T) {
        input := &iot.CreateTopicRuleDestinationInput{}
        output := &iot.CreateTopicRuleDestinationOutput{}

        mockClient.On("CreateTopicRuleDestination", ctx, input).Return(output, nil)

        result, err := mockClient.CreateTopicRuleDestination(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteAccountAuditConfiguration", func(t *testing.T) {
        input := &iot.DeleteAccountAuditConfigurationInput{}
        output := &iot.DeleteAccountAuditConfigurationOutput{}

        mockClient.On("DeleteAccountAuditConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteAccountAuditConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteAuditSuppression", func(t *testing.T) {
        input := &iot.DeleteAuditSuppressionInput{}
        output := &iot.DeleteAuditSuppressionOutput{}

        mockClient.On("DeleteAuditSuppression", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteAuditSuppression(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteAuthorizer", func(t *testing.T) {
        input := &iot.DeleteAuthorizerInput{}
        output := &iot.DeleteAuthorizerOutput{}

        mockClient.On("DeleteAuthorizer", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteAuthorizer(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteBillingGroup", func(t *testing.T) {
        input := &iot.DeleteBillingGroupInput{}
        output := &iot.DeleteBillingGroupOutput{}

        mockClient.On("DeleteBillingGroup", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteBillingGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteCACertificate", func(t *testing.T) {
        input := &iot.DeleteCACertificateInput{}
        output := &iot.DeleteCACertificateOutput{}

        mockClient.On("DeleteCACertificate", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteCACertificate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteCertificate", func(t *testing.T) {
        input := &iot.DeleteCertificateInput{}
        output := &iot.DeleteCertificateOutput{}

        mockClient.On("DeleteCertificate", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteCertificate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteCertificateProvider", func(t *testing.T) {
        input := &iot.DeleteCertificateProviderInput{}
        output := &iot.DeleteCertificateProviderOutput{}

        mockClient.On("DeleteCertificateProvider", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteCertificateProvider(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteCustomMetric", func(t *testing.T) {
        input := &iot.DeleteCustomMetricInput{}
        output := &iot.DeleteCustomMetricOutput{}

        mockClient.On("DeleteCustomMetric", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteCustomMetric(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteDimension", func(t *testing.T) {
        input := &iot.DeleteDimensionInput{}
        output := &iot.DeleteDimensionOutput{}

        mockClient.On("DeleteDimension", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteDimension(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteDomainConfiguration", func(t *testing.T) {
        input := &iot.DeleteDomainConfigurationInput{}
        output := &iot.DeleteDomainConfigurationOutput{}

        mockClient.On("DeleteDomainConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteDomainConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteDynamicThingGroup", func(t *testing.T) {
        input := &iot.DeleteDynamicThingGroupInput{}
        output := &iot.DeleteDynamicThingGroupOutput{}

        mockClient.On("DeleteDynamicThingGroup", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteDynamicThingGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteFleetMetric", func(t *testing.T) {
        input := &iot.DeleteFleetMetricInput{}
        output := &iot.DeleteFleetMetricOutput{}

        mockClient.On("DeleteFleetMetric", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteFleetMetric(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteJob", func(t *testing.T) {
        input := &iot.DeleteJobInput{}
        output := &iot.DeleteJobOutput{}

        mockClient.On("DeleteJob", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteJobExecution", func(t *testing.T) {
        input := &iot.DeleteJobExecutionInput{}
        output := &iot.DeleteJobExecutionOutput{}

        mockClient.On("DeleteJobExecution", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteJobExecution(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteJobTemplate", func(t *testing.T) {
        input := &iot.DeleteJobTemplateInput{}
        output := &iot.DeleteJobTemplateOutput{}

        mockClient.On("DeleteJobTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteJobTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteMitigationAction", func(t *testing.T) {
        input := &iot.DeleteMitigationActionInput{}
        output := &iot.DeleteMitigationActionOutput{}

        mockClient.On("DeleteMitigationAction", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteMitigationAction(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteOTAUpdate", func(t *testing.T) {
        input := &iot.DeleteOTAUpdateInput{}
        output := &iot.DeleteOTAUpdateOutput{}

        mockClient.On("DeleteOTAUpdate", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteOTAUpdate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeletePackage", func(t *testing.T) {
        input := &iot.DeletePackageInput{}
        output := &iot.DeletePackageOutput{}

        mockClient.On("DeletePackage", ctx, input).Return(output, nil)

        result, err := mockClient.DeletePackage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeletePackageVersion", func(t *testing.T) {
        input := &iot.DeletePackageVersionInput{}
        output := &iot.DeletePackageVersionOutput{}

        mockClient.On("DeletePackageVersion", ctx, input).Return(output, nil)

        result, err := mockClient.DeletePackageVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeletePolicy", func(t *testing.T) {
        input := &iot.DeletePolicyInput{}
        output := &iot.DeletePolicyOutput{}

        mockClient.On("DeletePolicy", ctx, input).Return(output, nil)

        result, err := mockClient.DeletePolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeletePolicyVersion", func(t *testing.T) {
        input := &iot.DeletePolicyVersionInput{}
        output := &iot.DeletePolicyVersionOutput{}

        mockClient.On("DeletePolicyVersion", ctx, input).Return(output, nil)

        result, err := mockClient.DeletePolicyVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteProvisioningTemplate", func(t *testing.T) {
        input := &iot.DeleteProvisioningTemplateInput{}
        output := &iot.DeleteProvisioningTemplateOutput{}

        mockClient.On("DeleteProvisioningTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteProvisioningTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteProvisioningTemplateVersion", func(t *testing.T) {
        input := &iot.DeleteProvisioningTemplateVersionInput{}
        output := &iot.DeleteProvisioningTemplateVersionOutput{}

        mockClient.On("DeleteProvisioningTemplateVersion", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteProvisioningTemplateVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteRegistrationCode", func(t *testing.T) {
        input := &iot.DeleteRegistrationCodeInput{}
        output := &iot.DeleteRegistrationCodeOutput{}

        mockClient.On("DeleteRegistrationCode", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteRegistrationCode(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteRoleAlias", func(t *testing.T) {
        input := &iot.DeleteRoleAliasInput{}
        output := &iot.DeleteRoleAliasOutput{}

        mockClient.On("DeleteRoleAlias", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteRoleAlias(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteScheduledAudit", func(t *testing.T) {
        input := &iot.DeleteScheduledAuditInput{}
        output := &iot.DeleteScheduledAuditOutput{}

        mockClient.On("DeleteScheduledAudit", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteScheduledAudit(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteSecurityProfile", func(t *testing.T) {
        input := &iot.DeleteSecurityProfileInput{}
        output := &iot.DeleteSecurityProfileOutput{}

        mockClient.On("DeleteSecurityProfile", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteSecurityProfile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteStream", func(t *testing.T) {
        input := &iot.DeleteStreamInput{}
        output := &iot.DeleteStreamOutput{}

        mockClient.On("DeleteStream", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteStream(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteThing", func(t *testing.T) {
        input := &iot.DeleteThingInput{}
        output := &iot.DeleteThingOutput{}

        mockClient.On("DeleteThing", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteThing(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteThingGroup", func(t *testing.T) {
        input := &iot.DeleteThingGroupInput{}
        output := &iot.DeleteThingGroupOutput{}

        mockClient.On("DeleteThingGroup", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteThingGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteThingType", func(t *testing.T) {
        input := &iot.DeleteThingTypeInput{}
        output := &iot.DeleteThingTypeOutput{}

        mockClient.On("DeleteThingType", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteThingType(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteTopicRule", func(t *testing.T) {
        input := &iot.DeleteTopicRuleInput{}
        output := &iot.DeleteTopicRuleOutput{}

        mockClient.On("DeleteTopicRule", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteTopicRule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteTopicRuleDestination", func(t *testing.T) {
        input := &iot.DeleteTopicRuleDestinationInput{}
        output := &iot.DeleteTopicRuleDestinationOutput{}

        mockClient.On("DeleteTopicRuleDestination", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteTopicRuleDestination(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteV2LoggingLevel", func(t *testing.T) {
        input := &iot.DeleteV2LoggingLevelInput{}
        output := &iot.DeleteV2LoggingLevelOutput{}

        mockClient.On("DeleteV2LoggingLevel", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteV2LoggingLevel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeprecateThingType", func(t *testing.T) {
        input := &iot.DeprecateThingTypeInput{}
        output := &iot.DeprecateThingTypeOutput{}

        mockClient.On("DeprecateThingType", ctx, input).Return(output, nil)

        result, err := mockClient.DeprecateThingType(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeAccountAuditConfiguration", func(t *testing.T) {
        input := &iot.DescribeAccountAuditConfigurationInput{}
        output := &iot.DescribeAccountAuditConfigurationOutput{}

        mockClient.On("DescribeAccountAuditConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeAccountAuditConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeAuditFinding", func(t *testing.T) {
        input := &iot.DescribeAuditFindingInput{}
        output := &iot.DescribeAuditFindingOutput{}

        mockClient.On("DescribeAuditFinding", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeAuditFinding(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeAuditMitigationActionsTask", func(t *testing.T) {
        input := &iot.DescribeAuditMitigationActionsTaskInput{}
        output := &iot.DescribeAuditMitigationActionsTaskOutput{}

        mockClient.On("DescribeAuditMitigationActionsTask", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeAuditMitigationActionsTask(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeAuditSuppression", func(t *testing.T) {
        input := &iot.DescribeAuditSuppressionInput{}
        output := &iot.DescribeAuditSuppressionOutput{}

        mockClient.On("DescribeAuditSuppression", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeAuditSuppression(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeAuditTask", func(t *testing.T) {
        input := &iot.DescribeAuditTaskInput{}
        output := &iot.DescribeAuditTaskOutput{}

        mockClient.On("DescribeAuditTask", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeAuditTask(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeAuthorizer", func(t *testing.T) {
        input := &iot.DescribeAuthorizerInput{}
        output := &iot.DescribeAuthorizerOutput{}

        mockClient.On("DescribeAuthorizer", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeAuthorizer(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeBillingGroup", func(t *testing.T) {
        input := &iot.DescribeBillingGroupInput{}
        output := &iot.DescribeBillingGroupOutput{}

        mockClient.On("DescribeBillingGroup", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeBillingGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeCACertificate", func(t *testing.T) {
        input := &iot.DescribeCACertificateInput{}
        output := &iot.DescribeCACertificateOutput{}

        mockClient.On("DescribeCACertificate", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeCACertificate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeCertificate", func(t *testing.T) {
        input := &iot.DescribeCertificateInput{}
        output := &iot.DescribeCertificateOutput{}

        mockClient.On("DescribeCertificate", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeCertificate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeCertificateProvider", func(t *testing.T) {
        input := &iot.DescribeCertificateProviderInput{}
        output := &iot.DescribeCertificateProviderOutput{}

        mockClient.On("DescribeCertificateProvider", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeCertificateProvider(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeCustomMetric", func(t *testing.T) {
        input := &iot.DescribeCustomMetricInput{}
        output := &iot.DescribeCustomMetricOutput{}

        mockClient.On("DescribeCustomMetric", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeCustomMetric(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeDefaultAuthorizer", func(t *testing.T) {
        input := &iot.DescribeDefaultAuthorizerInput{}
        output := &iot.DescribeDefaultAuthorizerOutput{}

        mockClient.On("DescribeDefaultAuthorizer", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeDefaultAuthorizer(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeDetectMitigationActionsTask", func(t *testing.T) {
        input := &iot.DescribeDetectMitigationActionsTaskInput{}
        output := &iot.DescribeDetectMitigationActionsTaskOutput{}

        mockClient.On("DescribeDetectMitigationActionsTask", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeDetectMitigationActionsTask(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeDimension", func(t *testing.T) {
        input := &iot.DescribeDimensionInput{}
        output := &iot.DescribeDimensionOutput{}

        mockClient.On("DescribeDimension", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeDimension(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeDomainConfiguration", func(t *testing.T) {
        input := &iot.DescribeDomainConfigurationInput{}
        output := &iot.DescribeDomainConfigurationOutput{}

        mockClient.On("DescribeDomainConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeDomainConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeEndpoint", func(t *testing.T) {
        input := &iot.DescribeEndpointInput{}
        output := &iot.DescribeEndpointOutput{}

        mockClient.On("DescribeEndpoint", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeEndpoint(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeEventConfigurations", func(t *testing.T) {
        input := &iot.DescribeEventConfigurationsInput{}
        output := &iot.DescribeEventConfigurationsOutput{}

        mockClient.On("DescribeEventConfigurations", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeEventConfigurations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeFleetMetric", func(t *testing.T) {
        input := &iot.DescribeFleetMetricInput{}
        output := &iot.DescribeFleetMetricOutput{}

        mockClient.On("DescribeFleetMetric", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeFleetMetric(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeIndex", func(t *testing.T) {
        input := &iot.DescribeIndexInput{}
        output := &iot.DescribeIndexOutput{}

        mockClient.On("DescribeIndex", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeIndex(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeJob", func(t *testing.T) {
        input := &iot.DescribeJobInput{}
        output := &iot.DescribeJobOutput{}

        mockClient.On("DescribeJob", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeJobExecution", func(t *testing.T) {
        input := &iot.DescribeJobExecutionInput{}
        output := &iot.DescribeJobExecutionOutput{}

        mockClient.On("DescribeJobExecution", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeJobExecution(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeJobTemplate", func(t *testing.T) {
        input := &iot.DescribeJobTemplateInput{}
        output := &iot.DescribeJobTemplateOutput{}

        mockClient.On("DescribeJobTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeJobTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeManagedJobTemplate", func(t *testing.T) {
        input := &iot.DescribeManagedJobTemplateInput{}
        output := &iot.DescribeManagedJobTemplateOutput{}

        mockClient.On("DescribeManagedJobTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeManagedJobTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeMitigationAction", func(t *testing.T) {
        input := &iot.DescribeMitigationActionInput{}
        output := &iot.DescribeMitigationActionOutput{}

        mockClient.On("DescribeMitigationAction", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeMitigationAction(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeProvisioningTemplate", func(t *testing.T) {
        input := &iot.DescribeProvisioningTemplateInput{}
        output := &iot.DescribeProvisioningTemplateOutput{}

        mockClient.On("DescribeProvisioningTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeProvisioningTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeProvisioningTemplateVersion", func(t *testing.T) {
        input := &iot.DescribeProvisioningTemplateVersionInput{}
        output := &iot.DescribeProvisioningTemplateVersionOutput{}

        mockClient.On("DescribeProvisioningTemplateVersion", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeProvisioningTemplateVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeRoleAlias", func(t *testing.T) {
        input := &iot.DescribeRoleAliasInput{}
        output := &iot.DescribeRoleAliasOutput{}

        mockClient.On("DescribeRoleAlias", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeRoleAlias(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeScheduledAudit", func(t *testing.T) {
        input := &iot.DescribeScheduledAuditInput{}
        output := &iot.DescribeScheduledAuditOutput{}

        mockClient.On("DescribeScheduledAudit", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeScheduledAudit(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeSecurityProfile", func(t *testing.T) {
        input := &iot.DescribeSecurityProfileInput{}
        output := &iot.DescribeSecurityProfileOutput{}

        mockClient.On("DescribeSecurityProfile", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeSecurityProfile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeStream", func(t *testing.T) {
        input := &iot.DescribeStreamInput{}
        output := &iot.DescribeStreamOutput{}

        mockClient.On("DescribeStream", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeStream(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeThing", func(t *testing.T) {
        input := &iot.DescribeThingInput{}
        output := &iot.DescribeThingOutput{}

        mockClient.On("DescribeThing", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeThing(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeThingGroup", func(t *testing.T) {
        input := &iot.DescribeThingGroupInput{}
        output := &iot.DescribeThingGroupOutput{}

        mockClient.On("DescribeThingGroup", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeThingGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeThingRegistrationTask", func(t *testing.T) {
        input := &iot.DescribeThingRegistrationTaskInput{}
        output := &iot.DescribeThingRegistrationTaskOutput{}

        mockClient.On("DescribeThingRegistrationTask", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeThingRegistrationTask(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeThingType", func(t *testing.T) {
        input := &iot.DescribeThingTypeInput{}
        output := &iot.DescribeThingTypeOutput{}

        mockClient.On("DescribeThingType", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeThingType(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDetachPolicy", func(t *testing.T) {
        input := &iot.DetachPolicyInput{}
        output := &iot.DetachPolicyOutput{}

        mockClient.On("DetachPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.DetachPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDetachPrincipalPolicy", func(t *testing.T) {
        input := &iot.DetachPrincipalPolicyInput{}
        output := &iot.DetachPrincipalPolicyOutput{}

        mockClient.On("DetachPrincipalPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.DetachPrincipalPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDetachSecurityProfile", func(t *testing.T) {
        input := &iot.DetachSecurityProfileInput{}
        output := &iot.DetachSecurityProfileOutput{}

        mockClient.On("DetachSecurityProfile", ctx, input).Return(output, nil)

        result, err := mockClient.DetachSecurityProfile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDetachThingPrincipal", func(t *testing.T) {
        input := &iot.DetachThingPrincipalInput{}
        output := &iot.DetachThingPrincipalOutput{}

        mockClient.On("DetachThingPrincipal", ctx, input).Return(output, nil)

        result, err := mockClient.DetachThingPrincipal(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisableTopicRule", func(t *testing.T) {
        input := &iot.DisableTopicRuleInput{}
        output := &iot.DisableTopicRuleOutput{}

        mockClient.On("DisableTopicRule", ctx, input).Return(output, nil)

        result, err := mockClient.DisableTopicRule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestEnableTopicRule", func(t *testing.T) {
        input := &iot.EnableTopicRuleInput{}
        output := &iot.EnableTopicRuleOutput{}

        mockClient.On("EnableTopicRule", ctx, input).Return(output, nil)

        result, err := mockClient.EnableTopicRule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetBehaviorModelTrainingSummaries", func(t *testing.T) {
        input := &iot.GetBehaviorModelTrainingSummariesInput{}
        output := &iot.GetBehaviorModelTrainingSummariesOutput{}

        mockClient.On("GetBehaviorModelTrainingSummaries", ctx, input).Return(output, nil)

        result, err := mockClient.GetBehaviorModelTrainingSummaries(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetBucketsAggregation", func(t *testing.T) {
        input := &iot.GetBucketsAggregationInput{}
        output := &iot.GetBucketsAggregationOutput{}

        mockClient.On("GetBucketsAggregation", ctx, input).Return(output, nil)

        result, err := mockClient.GetBucketsAggregation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetCardinality", func(t *testing.T) {
        input := &iot.GetCardinalityInput{}
        output := &iot.GetCardinalityOutput{}

        mockClient.On("GetCardinality", ctx, input).Return(output, nil)

        result, err := mockClient.GetCardinality(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetEffectivePolicies", func(t *testing.T) {
        input := &iot.GetEffectivePoliciesInput{}
        output := &iot.GetEffectivePoliciesOutput{}

        mockClient.On("GetEffectivePolicies", ctx, input).Return(output, nil)

        result, err := mockClient.GetEffectivePolicies(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetIndexingConfiguration", func(t *testing.T) {
        input := &iot.GetIndexingConfigurationInput{}
        output := &iot.GetIndexingConfigurationOutput{}

        mockClient.On("GetIndexingConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.GetIndexingConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetJobDocument", func(t *testing.T) {
        input := &iot.GetJobDocumentInput{}
        output := &iot.GetJobDocumentOutput{}

        mockClient.On("GetJobDocument", ctx, input).Return(output, nil)

        result, err := mockClient.GetJobDocument(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetLoggingOptions", func(t *testing.T) {
        input := &iot.GetLoggingOptionsInput{}
        output := &iot.GetLoggingOptionsOutput{}

        mockClient.On("GetLoggingOptions", ctx, input).Return(output, nil)

        result, err := mockClient.GetLoggingOptions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetOTAUpdate", func(t *testing.T) {
        input := &iot.GetOTAUpdateInput{}
        output := &iot.GetOTAUpdateOutput{}

        mockClient.On("GetOTAUpdate", ctx, input).Return(output, nil)

        result, err := mockClient.GetOTAUpdate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetPackage", func(t *testing.T) {
        input := &iot.GetPackageInput{}
        output := &iot.GetPackageOutput{}

        mockClient.On("GetPackage", ctx, input).Return(output, nil)

        result, err := mockClient.GetPackage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetPackageConfiguration", func(t *testing.T) {
        input := &iot.GetPackageConfigurationInput{}
        output := &iot.GetPackageConfigurationOutput{}

        mockClient.On("GetPackageConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.GetPackageConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetPackageVersion", func(t *testing.T) {
        input := &iot.GetPackageVersionInput{}
        output := &iot.GetPackageVersionOutput{}

        mockClient.On("GetPackageVersion", ctx, input).Return(output, nil)

        result, err := mockClient.GetPackageVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetPercentiles", func(t *testing.T) {
        input := &iot.GetPercentilesInput{}
        output := &iot.GetPercentilesOutput{}

        mockClient.On("GetPercentiles", ctx, input).Return(output, nil)

        result, err := mockClient.GetPercentiles(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetPolicy", func(t *testing.T) {
        input := &iot.GetPolicyInput{}
        output := &iot.GetPolicyOutput{}

        mockClient.On("GetPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.GetPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetPolicyVersion", func(t *testing.T) {
        input := &iot.GetPolicyVersionInput{}
        output := &iot.GetPolicyVersionOutput{}

        mockClient.On("GetPolicyVersion", ctx, input).Return(output, nil)

        result, err := mockClient.GetPolicyVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetRegistrationCode", func(t *testing.T) {
        input := &iot.GetRegistrationCodeInput{}
        output := &iot.GetRegistrationCodeOutput{}

        mockClient.On("GetRegistrationCode", ctx, input).Return(output, nil)

        result, err := mockClient.GetRegistrationCode(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetStatistics", func(t *testing.T) {
        input := &iot.GetStatisticsInput{}
        output := &iot.GetStatisticsOutput{}

        mockClient.On("GetStatistics", ctx, input).Return(output, nil)

        result, err := mockClient.GetStatistics(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetTopicRule", func(t *testing.T) {
        input := &iot.GetTopicRuleInput{}
        output := &iot.GetTopicRuleOutput{}

        mockClient.On("GetTopicRule", ctx, input).Return(output, nil)

        result, err := mockClient.GetTopicRule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetTopicRuleDestination", func(t *testing.T) {
        input := &iot.GetTopicRuleDestinationInput{}
        output := &iot.GetTopicRuleDestinationOutput{}

        mockClient.On("GetTopicRuleDestination", ctx, input).Return(output, nil)

        result, err := mockClient.GetTopicRuleDestination(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetV2LoggingOptions", func(t *testing.T) {
        input := &iot.GetV2LoggingOptionsInput{}
        output := &iot.GetV2LoggingOptionsOutput{}

        mockClient.On("GetV2LoggingOptions", ctx, input).Return(output, nil)

        result, err := mockClient.GetV2LoggingOptions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListActiveViolations", func(t *testing.T) {
        input := &iot.ListActiveViolationsInput{}
        output := &iot.ListActiveViolationsOutput{}

        mockClient.On("ListActiveViolations", ctx, input).Return(output, nil)

        result, err := mockClient.ListActiveViolations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAttachedPolicies", func(t *testing.T) {
        input := &iot.ListAttachedPoliciesInput{}
        output := &iot.ListAttachedPoliciesOutput{}

        mockClient.On("ListAttachedPolicies", ctx, input).Return(output, nil)

        result, err := mockClient.ListAttachedPolicies(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAuditFindings", func(t *testing.T) {
        input := &iot.ListAuditFindingsInput{}
        output := &iot.ListAuditFindingsOutput{}

        mockClient.On("ListAuditFindings", ctx, input).Return(output, nil)

        result, err := mockClient.ListAuditFindings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAuditMitigationActionsExecutions", func(t *testing.T) {
        input := &iot.ListAuditMitigationActionsExecutionsInput{}
        output := &iot.ListAuditMitigationActionsExecutionsOutput{}

        mockClient.On("ListAuditMitigationActionsExecutions", ctx, input).Return(output, nil)

        result, err := mockClient.ListAuditMitigationActionsExecutions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAuditMitigationActionsTasks", func(t *testing.T) {
        input := &iot.ListAuditMitigationActionsTasksInput{}
        output := &iot.ListAuditMitigationActionsTasksOutput{}

        mockClient.On("ListAuditMitigationActionsTasks", ctx, input).Return(output, nil)

        result, err := mockClient.ListAuditMitigationActionsTasks(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAuditSuppressions", func(t *testing.T) {
        input := &iot.ListAuditSuppressionsInput{}
        output := &iot.ListAuditSuppressionsOutput{}

        mockClient.On("ListAuditSuppressions", ctx, input).Return(output, nil)

        result, err := mockClient.ListAuditSuppressions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAuditTasks", func(t *testing.T) {
        input := &iot.ListAuditTasksInput{}
        output := &iot.ListAuditTasksOutput{}

        mockClient.On("ListAuditTasks", ctx, input).Return(output, nil)

        result, err := mockClient.ListAuditTasks(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAuthorizers", func(t *testing.T) {
        input := &iot.ListAuthorizersInput{}
        output := &iot.ListAuthorizersOutput{}

        mockClient.On("ListAuthorizers", ctx, input).Return(output, nil)

        result, err := mockClient.ListAuthorizers(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListBillingGroups", func(t *testing.T) {
        input := &iot.ListBillingGroupsInput{}
        output := &iot.ListBillingGroupsOutput{}

        mockClient.On("ListBillingGroups", ctx, input).Return(output, nil)

        result, err := mockClient.ListBillingGroups(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListCACertificates", func(t *testing.T) {
        input := &iot.ListCACertificatesInput{}
        output := &iot.ListCACertificatesOutput{}

        mockClient.On("ListCACertificates", ctx, input).Return(output, nil)

        result, err := mockClient.ListCACertificates(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListCertificateProviders", func(t *testing.T) {
        input := &iot.ListCertificateProvidersInput{}
        output := &iot.ListCertificateProvidersOutput{}

        mockClient.On("ListCertificateProviders", ctx, input).Return(output, nil)

        result, err := mockClient.ListCertificateProviders(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListCertificates", func(t *testing.T) {
        input := &iot.ListCertificatesInput{}
        output := &iot.ListCertificatesOutput{}

        mockClient.On("ListCertificates", ctx, input).Return(output, nil)

        result, err := mockClient.ListCertificates(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListCertificatesByCA", func(t *testing.T) {
        input := &iot.ListCertificatesByCAInput{}
        output := &iot.ListCertificatesByCAOutput{}

        mockClient.On("ListCertificatesByCA", ctx, input).Return(output, nil)

        result, err := mockClient.ListCertificatesByCA(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListCustomMetrics", func(t *testing.T) {
        input := &iot.ListCustomMetricsInput{}
        output := &iot.ListCustomMetricsOutput{}

        mockClient.On("ListCustomMetrics", ctx, input).Return(output, nil)

        result, err := mockClient.ListCustomMetrics(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDetectMitigationActionsExecutions", func(t *testing.T) {
        input := &iot.ListDetectMitigationActionsExecutionsInput{}
        output := &iot.ListDetectMitigationActionsExecutionsOutput{}

        mockClient.On("ListDetectMitigationActionsExecutions", ctx, input).Return(output, nil)

        result, err := mockClient.ListDetectMitigationActionsExecutions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDetectMitigationActionsTasks", func(t *testing.T) {
        input := &iot.ListDetectMitigationActionsTasksInput{}
        output := &iot.ListDetectMitigationActionsTasksOutput{}

        mockClient.On("ListDetectMitigationActionsTasks", ctx, input).Return(output, nil)

        result, err := mockClient.ListDetectMitigationActionsTasks(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDimensions", func(t *testing.T) {
        input := &iot.ListDimensionsInput{}
        output := &iot.ListDimensionsOutput{}

        mockClient.On("ListDimensions", ctx, input).Return(output, nil)

        result, err := mockClient.ListDimensions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDomainConfigurations", func(t *testing.T) {
        input := &iot.ListDomainConfigurationsInput{}
        output := &iot.ListDomainConfigurationsOutput{}

        mockClient.On("ListDomainConfigurations", ctx, input).Return(output, nil)

        result, err := mockClient.ListDomainConfigurations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListFleetMetrics", func(t *testing.T) {
        input := &iot.ListFleetMetricsInput{}
        output := &iot.ListFleetMetricsOutput{}

        mockClient.On("ListFleetMetrics", ctx, input).Return(output, nil)

        result, err := mockClient.ListFleetMetrics(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListIndices", func(t *testing.T) {
        input := &iot.ListIndicesInput{}
        output := &iot.ListIndicesOutput{}

        mockClient.On("ListIndices", ctx, input).Return(output, nil)

        result, err := mockClient.ListIndices(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListJobExecutionsForJob", func(t *testing.T) {
        input := &iot.ListJobExecutionsForJobInput{}
        output := &iot.ListJobExecutionsForJobOutput{}

        mockClient.On("ListJobExecutionsForJob", ctx, input).Return(output, nil)

        result, err := mockClient.ListJobExecutionsForJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListJobExecutionsForThing", func(t *testing.T) {
        input := &iot.ListJobExecutionsForThingInput{}
        output := &iot.ListJobExecutionsForThingOutput{}

        mockClient.On("ListJobExecutionsForThing", ctx, input).Return(output, nil)

        result, err := mockClient.ListJobExecutionsForThing(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListJobTemplates", func(t *testing.T) {
        input := &iot.ListJobTemplatesInput{}
        output := &iot.ListJobTemplatesOutput{}

        mockClient.On("ListJobTemplates", ctx, input).Return(output, nil)

        result, err := mockClient.ListJobTemplates(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListJobs", func(t *testing.T) {
        input := &iot.ListJobsInput{}
        output := &iot.ListJobsOutput{}

        mockClient.On("ListJobs", ctx, input).Return(output, nil)

        result, err := mockClient.ListJobs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListManagedJobTemplates", func(t *testing.T) {
        input := &iot.ListManagedJobTemplatesInput{}
        output := &iot.ListManagedJobTemplatesOutput{}

        mockClient.On("ListManagedJobTemplates", ctx, input).Return(output, nil)

        result, err := mockClient.ListManagedJobTemplates(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListMetricValues", func(t *testing.T) {
        input := &iot.ListMetricValuesInput{}
        output := &iot.ListMetricValuesOutput{}

        mockClient.On("ListMetricValues", ctx, input).Return(output, nil)

        result, err := mockClient.ListMetricValues(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListMitigationActions", func(t *testing.T) {
        input := &iot.ListMitigationActionsInput{}
        output := &iot.ListMitigationActionsOutput{}

        mockClient.On("ListMitigationActions", ctx, input).Return(output, nil)

        result, err := mockClient.ListMitigationActions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListOTAUpdates", func(t *testing.T) {
        input := &iot.ListOTAUpdatesInput{}
        output := &iot.ListOTAUpdatesOutput{}

        mockClient.On("ListOTAUpdates", ctx, input).Return(output, nil)

        result, err := mockClient.ListOTAUpdates(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListOutgoingCertificates", func(t *testing.T) {
        input := &iot.ListOutgoingCertificatesInput{}
        output := &iot.ListOutgoingCertificatesOutput{}

        mockClient.On("ListOutgoingCertificates", ctx, input).Return(output, nil)

        result, err := mockClient.ListOutgoingCertificates(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListPackageVersions", func(t *testing.T) {
        input := &iot.ListPackageVersionsInput{}
        output := &iot.ListPackageVersionsOutput{}

        mockClient.On("ListPackageVersions", ctx, input).Return(output, nil)

        result, err := mockClient.ListPackageVersions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListPackages", func(t *testing.T) {
        input := &iot.ListPackagesInput{}
        output := &iot.ListPackagesOutput{}

        mockClient.On("ListPackages", ctx, input).Return(output, nil)

        result, err := mockClient.ListPackages(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListPolicies", func(t *testing.T) {
        input := &iot.ListPoliciesInput{}
        output := &iot.ListPoliciesOutput{}

        mockClient.On("ListPolicies", ctx, input).Return(output, nil)

        result, err := mockClient.ListPolicies(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListPolicyPrincipals", func(t *testing.T) {
        input := &iot.ListPolicyPrincipalsInput{}
        output := &iot.ListPolicyPrincipalsOutput{}

        mockClient.On("ListPolicyPrincipals", ctx, input).Return(output, nil)

        result, err := mockClient.ListPolicyPrincipals(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListPolicyVersions", func(t *testing.T) {
        input := &iot.ListPolicyVersionsInput{}
        output := &iot.ListPolicyVersionsOutput{}

        mockClient.On("ListPolicyVersions", ctx, input).Return(output, nil)

        result, err := mockClient.ListPolicyVersions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListPrincipalPolicies", func(t *testing.T) {
        input := &iot.ListPrincipalPoliciesInput{}
        output := &iot.ListPrincipalPoliciesOutput{}

        mockClient.On("ListPrincipalPolicies", ctx, input).Return(output, nil)

        result, err := mockClient.ListPrincipalPolicies(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListPrincipalThings", func(t *testing.T) {
        input := &iot.ListPrincipalThingsInput{}
        output := &iot.ListPrincipalThingsOutput{}

        mockClient.On("ListPrincipalThings", ctx, input).Return(output, nil)

        result, err := mockClient.ListPrincipalThings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListProvisioningTemplateVersions", func(t *testing.T) {
        input := &iot.ListProvisioningTemplateVersionsInput{}
        output := &iot.ListProvisioningTemplateVersionsOutput{}

        mockClient.On("ListProvisioningTemplateVersions", ctx, input).Return(output, nil)

        result, err := mockClient.ListProvisioningTemplateVersions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListProvisioningTemplates", func(t *testing.T) {
        input := &iot.ListProvisioningTemplatesInput{}
        output := &iot.ListProvisioningTemplatesOutput{}

        mockClient.On("ListProvisioningTemplates", ctx, input).Return(output, nil)

        result, err := mockClient.ListProvisioningTemplates(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListRelatedResourcesForAuditFinding", func(t *testing.T) {
        input := &iot.ListRelatedResourcesForAuditFindingInput{}
        output := &iot.ListRelatedResourcesForAuditFindingOutput{}

        mockClient.On("ListRelatedResourcesForAuditFinding", ctx, input).Return(output, nil)

        result, err := mockClient.ListRelatedResourcesForAuditFinding(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListRoleAliases", func(t *testing.T) {
        input := &iot.ListRoleAliasesInput{}
        output := &iot.ListRoleAliasesOutput{}

        mockClient.On("ListRoleAliases", ctx, input).Return(output, nil)

        result, err := mockClient.ListRoleAliases(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListScheduledAudits", func(t *testing.T) {
        input := &iot.ListScheduledAuditsInput{}
        output := &iot.ListScheduledAuditsOutput{}

        mockClient.On("ListScheduledAudits", ctx, input).Return(output, nil)

        result, err := mockClient.ListScheduledAudits(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListSecurityProfiles", func(t *testing.T) {
        input := &iot.ListSecurityProfilesInput{}
        output := &iot.ListSecurityProfilesOutput{}

        mockClient.On("ListSecurityProfiles", ctx, input).Return(output, nil)

        result, err := mockClient.ListSecurityProfiles(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListSecurityProfilesForTarget", func(t *testing.T) {
        input := &iot.ListSecurityProfilesForTargetInput{}
        output := &iot.ListSecurityProfilesForTargetOutput{}

        mockClient.On("ListSecurityProfilesForTarget", ctx, input).Return(output, nil)

        result, err := mockClient.ListSecurityProfilesForTarget(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListStreams", func(t *testing.T) {
        input := &iot.ListStreamsInput{}
        output := &iot.ListStreamsOutput{}

        mockClient.On("ListStreams", ctx, input).Return(output, nil)

        result, err := mockClient.ListStreams(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &iot.ListTagsForResourceInput{}
        output := &iot.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTargetsForPolicy", func(t *testing.T) {
        input := &iot.ListTargetsForPolicyInput{}
        output := &iot.ListTargetsForPolicyOutput{}

        mockClient.On("ListTargetsForPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.ListTargetsForPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTargetsForSecurityProfile", func(t *testing.T) {
        input := &iot.ListTargetsForSecurityProfileInput{}
        output := &iot.ListTargetsForSecurityProfileOutput{}

        mockClient.On("ListTargetsForSecurityProfile", ctx, input).Return(output, nil)

        result, err := mockClient.ListTargetsForSecurityProfile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListThingGroups", func(t *testing.T) {
        input := &iot.ListThingGroupsInput{}
        output := &iot.ListThingGroupsOutput{}

        mockClient.On("ListThingGroups", ctx, input).Return(output, nil)

        result, err := mockClient.ListThingGroups(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListThingGroupsForThing", func(t *testing.T) {
        input := &iot.ListThingGroupsForThingInput{}
        output := &iot.ListThingGroupsForThingOutput{}

        mockClient.On("ListThingGroupsForThing", ctx, input).Return(output, nil)

        result, err := mockClient.ListThingGroupsForThing(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListThingPrincipals", func(t *testing.T) {
        input := &iot.ListThingPrincipalsInput{}
        output := &iot.ListThingPrincipalsOutput{}

        mockClient.On("ListThingPrincipals", ctx, input).Return(output, nil)

        result, err := mockClient.ListThingPrincipals(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListThingRegistrationTaskReports", func(t *testing.T) {
        input := &iot.ListThingRegistrationTaskReportsInput{}
        output := &iot.ListThingRegistrationTaskReportsOutput{}

        mockClient.On("ListThingRegistrationTaskReports", ctx, input).Return(output, nil)

        result, err := mockClient.ListThingRegistrationTaskReports(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListThingRegistrationTasks", func(t *testing.T) {
        input := &iot.ListThingRegistrationTasksInput{}
        output := &iot.ListThingRegistrationTasksOutput{}

        mockClient.On("ListThingRegistrationTasks", ctx, input).Return(output, nil)

        result, err := mockClient.ListThingRegistrationTasks(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListThingTypes", func(t *testing.T) {
        input := &iot.ListThingTypesInput{}
        output := &iot.ListThingTypesOutput{}

        mockClient.On("ListThingTypes", ctx, input).Return(output, nil)

        result, err := mockClient.ListThingTypes(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListThings", func(t *testing.T) {
        input := &iot.ListThingsInput{}
        output := &iot.ListThingsOutput{}

        mockClient.On("ListThings", ctx, input).Return(output, nil)

        result, err := mockClient.ListThings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListThingsInBillingGroup", func(t *testing.T) {
        input := &iot.ListThingsInBillingGroupInput{}
        output := &iot.ListThingsInBillingGroupOutput{}

        mockClient.On("ListThingsInBillingGroup", ctx, input).Return(output, nil)

        result, err := mockClient.ListThingsInBillingGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListThingsInThingGroup", func(t *testing.T) {
        input := &iot.ListThingsInThingGroupInput{}
        output := &iot.ListThingsInThingGroupOutput{}

        mockClient.On("ListThingsInThingGroup", ctx, input).Return(output, nil)

        result, err := mockClient.ListThingsInThingGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTopicRuleDestinations", func(t *testing.T) {
        input := &iot.ListTopicRuleDestinationsInput{}
        output := &iot.ListTopicRuleDestinationsOutput{}

        mockClient.On("ListTopicRuleDestinations", ctx, input).Return(output, nil)

        result, err := mockClient.ListTopicRuleDestinations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTopicRules", func(t *testing.T) {
        input := &iot.ListTopicRulesInput{}
        output := &iot.ListTopicRulesOutput{}

        mockClient.On("ListTopicRules", ctx, input).Return(output, nil)

        result, err := mockClient.ListTopicRules(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListV2LoggingLevels", func(t *testing.T) {
        input := &iot.ListV2LoggingLevelsInput{}
        output := &iot.ListV2LoggingLevelsOutput{}

        mockClient.On("ListV2LoggingLevels", ctx, input).Return(output, nil)

        result, err := mockClient.ListV2LoggingLevels(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListViolationEvents", func(t *testing.T) {
        input := &iot.ListViolationEventsInput{}
        output := &iot.ListViolationEventsOutput{}

        mockClient.On("ListViolationEvents", ctx, input).Return(output, nil)

        result, err := mockClient.ListViolationEvents(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutVerificationStateOnViolation", func(t *testing.T) {
        input := &iot.PutVerificationStateOnViolationInput{}
        output := &iot.PutVerificationStateOnViolationOutput{}

        mockClient.On("PutVerificationStateOnViolation", ctx, input).Return(output, nil)

        result, err := mockClient.PutVerificationStateOnViolation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRegisterCACertificate", func(t *testing.T) {
        input := &iot.RegisterCACertificateInput{}
        output := &iot.RegisterCACertificateOutput{}

        mockClient.On("RegisterCACertificate", ctx, input).Return(output, nil)

        result, err := mockClient.RegisterCACertificate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRegisterCertificate", func(t *testing.T) {
        input := &iot.RegisterCertificateInput{}
        output := &iot.RegisterCertificateOutput{}

        mockClient.On("RegisterCertificate", ctx, input).Return(output, nil)

        result, err := mockClient.RegisterCertificate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRegisterCertificateWithoutCA", func(t *testing.T) {
        input := &iot.RegisterCertificateWithoutCAInput{}
        output := &iot.RegisterCertificateWithoutCAOutput{}

        mockClient.On("RegisterCertificateWithoutCA", ctx, input).Return(output, nil)

        result, err := mockClient.RegisterCertificateWithoutCA(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRegisterThing", func(t *testing.T) {
        input := &iot.RegisterThingInput{}
        output := &iot.RegisterThingOutput{}

        mockClient.On("RegisterThing", ctx, input).Return(output, nil)

        result, err := mockClient.RegisterThing(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRejectCertificateTransfer", func(t *testing.T) {
        input := &iot.RejectCertificateTransferInput{}
        output := &iot.RejectCertificateTransferOutput{}

        mockClient.On("RejectCertificateTransfer", ctx, input).Return(output, nil)

        result, err := mockClient.RejectCertificateTransfer(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRemoveThingFromBillingGroup", func(t *testing.T) {
        input := &iot.RemoveThingFromBillingGroupInput{}
        output := &iot.RemoveThingFromBillingGroupOutput{}

        mockClient.On("RemoveThingFromBillingGroup", ctx, input).Return(output, nil)

        result, err := mockClient.RemoveThingFromBillingGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRemoveThingFromThingGroup", func(t *testing.T) {
        input := &iot.RemoveThingFromThingGroupInput{}
        output := &iot.RemoveThingFromThingGroupOutput{}

        mockClient.On("RemoveThingFromThingGroup", ctx, input).Return(output, nil)

        result, err := mockClient.RemoveThingFromThingGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestReplaceTopicRule", func(t *testing.T) {
        input := &iot.ReplaceTopicRuleInput{}
        output := &iot.ReplaceTopicRuleOutput{}

        mockClient.On("ReplaceTopicRule", ctx, input).Return(output, nil)

        result, err := mockClient.ReplaceTopicRule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSearchIndex", func(t *testing.T) {
        input := &iot.SearchIndexInput{}
        output := &iot.SearchIndexOutput{}

        mockClient.On("SearchIndex", ctx, input).Return(output, nil)

        result, err := mockClient.SearchIndex(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSetDefaultAuthorizer", func(t *testing.T) {
        input := &iot.SetDefaultAuthorizerInput{}
        output := &iot.SetDefaultAuthorizerOutput{}

        mockClient.On("SetDefaultAuthorizer", ctx, input).Return(output, nil)

        result, err := mockClient.SetDefaultAuthorizer(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSetDefaultPolicyVersion", func(t *testing.T) {
        input := &iot.SetDefaultPolicyVersionInput{}
        output := &iot.SetDefaultPolicyVersionOutput{}

        mockClient.On("SetDefaultPolicyVersion", ctx, input).Return(output, nil)

        result, err := mockClient.SetDefaultPolicyVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSetLoggingOptions", func(t *testing.T) {
        input := &iot.SetLoggingOptionsInput{}
        output := &iot.SetLoggingOptionsOutput{}

        mockClient.On("SetLoggingOptions", ctx, input).Return(output, nil)

        result, err := mockClient.SetLoggingOptions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSetV2LoggingLevel", func(t *testing.T) {
        input := &iot.SetV2LoggingLevelInput{}
        output := &iot.SetV2LoggingLevelOutput{}

        mockClient.On("SetV2LoggingLevel", ctx, input).Return(output, nil)

        result, err := mockClient.SetV2LoggingLevel(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSetV2LoggingOptions", func(t *testing.T) {
        input := &iot.SetV2LoggingOptionsInput{}
        output := &iot.SetV2LoggingOptionsOutput{}

        mockClient.On("SetV2LoggingOptions", ctx, input).Return(output, nil)

        result, err := mockClient.SetV2LoggingOptions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartAuditMitigationActionsTask", func(t *testing.T) {
        input := &iot.StartAuditMitigationActionsTaskInput{}
        output := &iot.StartAuditMitigationActionsTaskOutput{}

        mockClient.On("StartAuditMitigationActionsTask", ctx, input).Return(output, nil)

        result, err := mockClient.StartAuditMitigationActionsTask(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartDetectMitigationActionsTask", func(t *testing.T) {
        input := &iot.StartDetectMitigationActionsTaskInput{}
        output := &iot.StartDetectMitigationActionsTaskOutput{}

        mockClient.On("StartDetectMitigationActionsTask", ctx, input).Return(output, nil)

        result, err := mockClient.StartDetectMitigationActionsTask(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartOnDemandAuditTask", func(t *testing.T) {
        input := &iot.StartOnDemandAuditTaskInput{}
        output := &iot.StartOnDemandAuditTaskOutput{}

        mockClient.On("StartOnDemandAuditTask", ctx, input).Return(output, nil)

        result, err := mockClient.StartOnDemandAuditTask(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartThingRegistrationTask", func(t *testing.T) {
        input := &iot.StartThingRegistrationTaskInput{}
        output := &iot.StartThingRegistrationTaskOutput{}

        mockClient.On("StartThingRegistrationTask", ctx, input).Return(output, nil)

        result, err := mockClient.StartThingRegistrationTask(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopThingRegistrationTask", func(t *testing.T) {
        input := &iot.StopThingRegistrationTaskInput{}
        output := &iot.StopThingRegistrationTaskOutput{}

        mockClient.On("StopThingRegistrationTask", ctx, input).Return(output, nil)

        result, err := mockClient.StopThingRegistrationTask(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &iot.TagResourceInput{}
        output := &iot.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTestAuthorization", func(t *testing.T) {
        input := &iot.TestAuthorizationInput{}
        output := &iot.TestAuthorizationOutput{}

        mockClient.On("TestAuthorization", ctx, input).Return(output, nil)

        result, err := mockClient.TestAuthorization(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTestInvokeAuthorizer", func(t *testing.T) {
        input := &iot.TestInvokeAuthorizerInput{}
        output := &iot.TestInvokeAuthorizerOutput{}

        mockClient.On("TestInvokeAuthorizer", ctx, input).Return(output, nil)

        result, err := mockClient.TestInvokeAuthorizer(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTransferCertificate", func(t *testing.T) {
        input := &iot.TransferCertificateInput{}
        output := &iot.TransferCertificateOutput{}

        mockClient.On("TransferCertificate", ctx, input).Return(output, nil)

        result, err := mockClient.TransferCertificate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &iot.UntagResourceInput{}
        output := &iot.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateAccountAuditConfiguration", func(t *testing.T) {
        input := &iot.UpdateAccountAuditConfigurationInput{}
        output := &iot.UpdateAccountAuditConfigurationOutput{}

        mockClient.On("UpdateAccountAuditConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateAccountAuditConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateAuditSuppression", func(t *testing.T) {
        input := &iot.UpdateAuditSuppressionInput{}
        output := &iot.UpdateAuditSuppressionOutput{}

        mockClient.On("UpdateAuditSuppression", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateAuditSuppression(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateAuthorizer", func(t *testing.T) {
        input := &iot.UpdateAuthorizerInput{}
        output := &iot.UpdateAuthorizerOutput{}

        mockClient.On("UpdateAuthorizer", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateAuthorizer(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateBillingGroup", func(t *testing.T) {
        input := &iot.UpdateBillingGroupInput{}
        output := &iot.UpdateBillingGroupOutput{}

        mockClient.On("UpdateBillingGroup", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateBillingGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateCACertificate", func(t *testing.T) {
        input := &iot.UpdateCACertificateInput{}
        output := &iot.UpdateCACertificateOutput{}

        mockClient.On("UpdateCACertificate", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateCACertificate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateCertificate", func(t *testing.T) {
        input := &iot.UpdateCertificateInput{}
        output := &iot.UpdateCertificateOutput{}

        mockClient.On("UpdateCertificate", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateCertificate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateCertificateProvider", func(t *testing.T) {
        input := &iot.UpdateCertificateProviderInput{}
        output := &iot.UpdateCertificateProviderOutput{}

        mockClient.On("UpdateCertificateProvider", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateCertificateProvider(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateCustomMetric", func(t *testing.T) {
        input := &iot.UpdateCustomMetricInput{}
        output := &iot.UpdateCustomMetricOutput{}

        mockClient.On("UpdateCustomMetric", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateCustomMetric(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateDimension", func(t *testing.T) {
        input := &iot.UpdateDimensionInput{}
        output := &iot.UpdateDimensionOutput{}

        mockClient.On("UpdateDimension", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateDimension(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateDomainConfiguration", func(t *testing.T) {
        input := &iot.UpdateDomainConfigurationInput{}
        output := &iot.UpdateDomainConfigurationOutput{}

        mockClient.On("UpdateDomainConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateDomainConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateDynamicThingGroup", func(t *testing.T) {
        input := &iot.UpdateDynamicThingGroupInput{}
        output := &iot.UpdateDynamicThingGroupOutput{}

        mockClient.On("UpdateDynamicThingGroup", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateDynamicThingGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateEventConfigurations", func(t *testing.T) {
        input := &iot.UpdateEventConfigurationsInput{}
        output := &iot.UpdateEventConfigurationsOutput{}

        mockClient.On("UpdateEventConfigurations", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateEventConfigurations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateFleetMetric", func(t *testing.T) {
        input := &iot.UpdateFleetMetricInput{}
        output := &iot.UpdateFleetMetricOutput{}

        mockClient.On("UpdateFleetMetric", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateFleetMetric(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateIndexingConfiguration", func(t *testing.T) {
        input := &iot.UpdateIndexingConfigurationInput{}
        output := &iot.UpdateIndexingConfigurationOutput{}

        mockClient.On("UpdateIndexingConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateIndexingConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateJob", func(t *testing.T) {
        input := &iot.UpdateJobInput{}
        output := &iot.UpdateJobOutput{}

        mockClient.On("UpdateJob", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateMitigationAction", func(t *testing.T) {
        input := &iot.UpdateMitigationActionInput{}
        output := &iot.UpdateMitigationActionOutput{}

        mockClient.On("UpdateMitigationAction", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateMitigationAction(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdatePackage", func(t *testing.T) {
        input := &iot.UpdatePackageInput{}
        output := &iot.UpdatePackageOutput{}

        mockClient.On("UpdatePackage", ctx, input).Return(output, nil)

        result, err := mockClient.UpdatePackage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdatePackageConfiguration", func(t *testing.T) {
        input := &iot.UpdatePackageConfigurationInput{}
        output := &iot.UpdatePackageConfigurationOutput{}

        mockClient.On("UpdatePackageConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.UpdatePackageConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdatePackageVersion", func(t *testing.T) {
        input := &iot.UpdatePackageVersionInput{}
        output := &iot.UpdatePackageVersionOutput{}

        mockClient.On("UpdatePackageVersion", ctx, input).Return(output, nil)

        result, err := mockClient.UpdatePackageVersion(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateProvisioningTemplate", func(t *testing.T) {
        input := &iot.UpdateProvisioningTemplateInput{}
        output := &iot.UpdateProvisioningTemplateOutput{}

        mockClient.On("UpdateProvisioningTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateProvisioningTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateRoleAlias", func(t *testing.T) {
        input := &iot.UpdateRoleAliasInput{}
        output := &iot.UpdateRoleAliasOutput{}

        mockClient.On("UpdateRoleAlias", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateRoleAlias(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateScheduledAudit", func(t *testing.T) {
        input := &iot.UpdateScheduledAuditInput{}
        output := &iot.UpdateScheduledAuditOutput{}

        mockClient.On("UpdateScheduledAudit", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateScheduledAudit(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateSecurityProfile", func(t *testing.T) {
        input := &iot.UpdateSecurityProfileInput{}
        output := &iot.UpdateSecurityProfileOutput{}

        mockClient.On("UpdateSecurityProfile", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateSecurityProfile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateStream", func(t *testing.T) {
        input := &iot.UpdateStreamInput{}
        output := &iot.UpdateStreamOutput{}

        mockClient.On("UpdateStream", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateStream(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateThing", func(t *testing.T) {
        input := &iot.UpdateThingInput{}
        output := &iot.UpdateThingOutput{}

        mockClient.On("UpdateThing", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateThing(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateThingGroup", func(t *testing.T) {
        input := &iot.UpdateThingGroupInput{}
        output := &iot.UpdateThingGroupOutput{}

        mockClient.On("UpdateThingGroup", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateThingGroup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateThingGroupsForThing", func(t *testing.T) {
        input := &iot.UpdateThingGroupsForThingInput{}
        output := &iot.UpdateThingGroupsForThingOutput{}

        mockClient.On("UpdateThingGroupsForThing", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateThingGroupsForThing(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateTopicRuleDestination", func(t *testing.T) {
        input := &iot.UpdateTopicRuleDestinationInput{}
        output := &iot.UpdateTopicRuleDestinationOutput{}

        mockClient.On("UpdateTopicRuleDestination", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateTopicRuleDestination(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestValidateSecurityProfileBehaviors", func(t *testing.T) {
        input := &iot.ValidateSecurityProfileBehaviorsInput{}
        output := &iot.ValidateSecurityProfileBehaviorsOutput{}

        mockClient.On("ValidateSecurityProfileBehaviors", ctx, input).Return(output, nil)

        result, err := mockClient.ValidateSecurityProfileBehaviors(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
