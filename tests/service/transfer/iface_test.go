package transfer_test

// tests for the transfer service interface for this ../../../service/transfer/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/transfer"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/transfer/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/transfer/transfer_iface"
	"github.com/stretchr/testify/assert"
)

func TestTransferServiceCanBeMocked(t *testing.T) {
	var iface transfer_iface.IClient
	iface = &transfer.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := transfer.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateAccess", func(t *testing.T) {
        input := &transfer.CreateAccessInput{}
        output := &transfer.CreateAccessOutput{}

        mockClient.On("CreateAccess", ctx, input).Return(output, nil)

        result, err := mockClient.CreateAccess(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateAgreement", func(t *testing.T) {
        input := &transfer.CreateAgreementInput{}
        output := &transfer.CreateAgreementOutput{}

        mockClient.On("CreateAgreement", ctx, input).Return(output, nil)

        result, err := mockClient.CreateAgreement(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateConnector", func(t *testing.T) {
        input := &transfer.CreateConnectorInput{}
        output := &transfer.CreateConnectorOutput{}

        mockClient.On("CreateConnector", ctx, input).Return(output, nil)

        result, err := mockClient.CreateConnector(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateProfile", func(t *testing.T) {
        input := &transfer.CreateProfileInput{}
        output := &transfer.CreateProfileOutput{}

        mockClient.On("CreateProfile", ctx, input).Return(output, nil)

        result, err := mockClient.CreateProfile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateServer", func(t *testing.T) {
        input := &transfer.CreateServerInput{}
        output := &transfer.CreateServerOutput{}

        mockClient.On("CreateServer", ctx, input).Return(output, nil)

        result, err := mockClient.CreateServer(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateUser", func(t *testing.T) {
        input := &transfer.CreateUserInput{}
        output := &transfer.CreateUserOutput{}

        mockClient.On("CreateUser", ctx, input).Return(output, nil)

        result, err := mockClient.CreateUser(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateWorkflow", func(t *testing.T) {
        input := &transfer.CreateWorkflowInput{}
        output := &transfer.CreateWorkflowOutput{}

        mockClient.On("CreateWorkflow", ctx, input).Return(output, nil)

        result, err := mockClient.CreateWorkflow(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteAccess", func(t *testing.T) {
        input := &transfer.DeleteAccessInput{}
        output := &transfer.DeleteAccessOutput{}

        mockClient.On("DeleteAccess", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteAccess(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteAgreement", func(t *testing.T) {
        input := &transfer.DeleteAgreementInput{}
        output := &transfer.DeleteAgreementOutput{}

        mockClient.On("DeleteAgreement", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteAgreement(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteCertificate", func(t *testing.T) {
        input := &transfer.DeleteCertificateInput{}
        output := &transfer.DeleteCertificateOutput{}

        mockClient.On("DeleteCertificate", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteCertificate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteConnector", func(t *testing.T) {
        input := &transfer.DeleteConnectorInput{}
        output := &transfer.DeleteConnectorOutput{}

        mockClient.On("DeleteConnector", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteConnector(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteHostKey", func(t *testing.T) {
        input := &transfer.DeleteHostKeyInput{}
        output := &transfer.DeleteHostKeyOutput{}

        mockClient.On("DeleteHostKey", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteHostKey(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteProfile", func(t *testing.T) {
        input := &transfer.DeleteProfileInput{}
        output := &transfer.DeleteProfileOutput{}

        mockClient.On("DeleteProfile", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteProfile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteServer", func(t *testing.T) {
        input := &transfer.DeleteServerInput{}
        output := &transfer.DeleteServerOutput{}

        mockClient.On("DeleteServer", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteServer(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteSshPublicKey", func(t *testing.T) {
        input := &transfer.DeleteSshPublicKeyInput{}
        output := &transfer.DeleteSshPublicKeyOutput{}

        mockClient.On("DeleteSshPublicKey", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteSshPublicKey(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteUser", func(t *testing.T) {
        input := &transfer.DeleteUserInput{}
        output := &transfer.DeleteUserOutput{}

        mockClient.On("DeleteUser", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteUser(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteWorkflow", func(t *testing.T) {
        input := &transfer.DeleteWorkflowInput{}
        output := &transfer.DeleteWorkflowOutput{}

        mockClient.On("DeleteWorkflow", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteWorkflow(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeAccess", func(t *testing.T) {
        input := &transfer.DescribeAccessInput{}
        output := &transfer.DescribeAccessOutput{}

        mockClient.On("DescribeAccess", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeAccess(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeAgreement", func(t *testing.T) {
        input := &transfer.DescribeAgreementInput{}
        output := &transfer.DescribeAgreementOutput{}

        mockClient.On("DescribeAgreement", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeAgreement(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeCertificate", func(t *testing.T) {
        input := &transfer.DescribeCertificateInput{}
        output := &transfer.DescribeCertificateOutput{}

        mockClient.On("DescribeCertificate", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeCertificate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeConnector", func(t *testing.T) {
        input := &transfer.DescribeConnectorInput{}
        output := &transfer.DescribeConnectorOutput{}

        mockClient.On("DescribeConnector", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeConnector(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeExecution", func(t *testing.T) {
        input := &transfer.DescribeExecutionInput{}
        output := &transfer.DescribeExecutionOutput{}

        mockClient.On("DescribeExecution", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeExecution(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeHostKey", func(t *testing.T) {
        input := &transfer.DescribeHostKeyInput{}
        output := &transfer.DescribeHostKeyOutput{}

        mockClient.On("DescribeHostKey", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeHostKey(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeProfile", func(t *testing.T) {
        input := &transfer.DescribeProfileInput{}
        output := &transfer.DescribeProfileOutput{}

        mockClient.On("DescribeProfile", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeProfile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeSecurityPolicy", func(t *testing.T) {
        input := &transfer.DescribeSecurityPolicyInput{}
        output := &transfer.DescribeSecurityPolicyOutput{}

        mockClient.On("DescribeSecurityPolicy", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeSecurityPolicy(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeServer", func(t *testing.T) {
        input := &transfer.DescribeServerInput{}
        output := &transfer.DescribeServerOutput{}

        mockClient.On("DescribeServer", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeServer(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeUser", func(t *testing.T) {
        input := &transfer.DescribeUserInput{}
        output := &transfer.DescribeUserOutput{}

        mockClient.On("DescribeUser", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeUser(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDescribeWorkflow", func(t *testing.T) {
        input := &transfer.DescribeWorkflowInput{}
        output := &transfer.DescribeWorkflowOutput{}

        mockClient.On("DescribeWorkflow", ctx, input).Return(output, nil)

        result, err := mockClient.DescribeWorkflow(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestImportCertificate", func(t *testing.T) {
        input := &transfer.ImportCertificateInput{}
        output := &transfer.ImportCertificateOutput{}

        mockClient.On("ImportCertificate", ctx, input).Return(output, nil)

        result, err := mockClient.ImportCertificate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestImportHostKey", func(t *testing.T) {
        input := &transfer.ImportHostKeyInput{}
        output := &transfer.ImportHostKeyOutput{}

        mockClient.On("ImportHostKey", ctx, input).Return(output, nil)

        result, err := mockClient.ImportHostKey(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestImportSshPublicKey", func(t *testing.T) {
        input := &transfer.ImportSshPublicKeyInput{}
        output := &transfer.ImportSshPublicKeyOutput{}

        mockClient.On("ImportSshPublicKey", ctx, input).Return(output, nil)

        result, err := mockClient.ImportSshPublicKey(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAccesses", func(t *testing.T) {
        input := &transfer.ListAccessesInput{}
        output := &transfer.ListAccessesOutput{}

        mockClient.On("ListAccesses", ctx, input).Return(output, nil)

        result, err := mockClient.ListAccesses(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAgreements", func(t *testing.T) {
        input := &transfer.ListAgreementsInput{}
        output := &transfer.ListAgreementsOutput{}

        mockClient.On("ListAgreements", ctx, input).Return(output, nil)

        result, err := mockClient.ListAgreements(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListCertificates", func(t *testing.T) {
        input := &transfer.ListCertificatesInput{}
        output := &transfer.ListCertificatesOutput{}

        mockClient.On("ListCertificates", ctx, input).Return(output, nil)

        result, err := mockClient.ListCertificates(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListConnectors", func(t *testing.T) {
        input := &transfer.ListConnectorsInput{}
        output := &transfer.ListConnectorsOutput{}

        mockClient.On("ListConnectors", ctx, input).Return(output, nil)

        result, err := mockClient.ListConnectors(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListExecutions", func(t *testing.T) {
        input := &transfer.ListExecutionsInput{}
        output := &transfer.ListExecutionsOutput{}

        mockClient.On("ListExecutions", ctx, input).Return(output, nil)

        result, err := mockClient.ListExecutions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListHostKeys", func(t *testing.T) {
        input := &transfer.ListHostKeysInput{}
        output := &transfer.ListHostKeysOutput{}

        mockClient.On("ListHostKeys", ctx, input).Return(output, nil)

        result, err := mockClient.ListHostKeys(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListProfiles", func(t *testing.T) {
        input := &transfer.ListProfilesInput{}
        output := &transfer.ListProfilesOutput{}

        mockClient.On("ListProfiles", ctx, input).Return(output, nil)

        result, err := mockClient.ListProfiles(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListSecurityPolicies", func(t *testing.T) {
        input := &transfer.ListSecurityPoliciesInput{}
        output := &transfer.ListSecurityPoliciesOutput{}

        mockClient.On("ListSecurityPolicies", ctx, input).Return(output, nil)

        result, err := mockClient.ListSecurityPolicies(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListServers", func(t *testing.T) {
        input := &transfer.ListServersInput{}
        output := &transfer.ListServersOutput{}

        mockClient.On("ListServers", ctx, input).Return(output, nil)

        result, err := mockClient.ListServers(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &transfer.ListTagsForResourceInput{}
        output := &transfer.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListUsers", func(t *testing.T) {
        input := &transfer.ListUsersInput{}
        output := &transfer.ListUsersOutput{}

        mockClient.On("ListUsers", ctx, input).Return(output, nil)

        result, err := mockClient.ListUsers(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListWorkflows", func(t *testing.T) {
        input := &transfer.ListWorkflowsInput{}
        output := &transfer.ListWorkflowsOutput{}

        mockClient.On("ListWorkflows", ctx, input).Return(output, nil)

        result, err := mockClient.ListWorkflows(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSendWorkflowStepState", func(t *testing.T) {
        input := &transfer.SendWorkflowStepStateInput{}
        output := &transfer.SendWorkflowStepStateOutput{}

        mockClient.On("SendWorkflowStepState", ctx, input).Return(output, nil)

        result, err := mockClient.SendWorkflowStepState(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartDirectoryListing", func(t *testing.T) {
        input := &transfer.StartDirectoryListingInput{}
        output := &transfer.StartDirectoryListingOutput{}

        mockClient.On("StartDirectoryListing", ctx, input).Return(output, nil)

        result, err := mockClient.StartDirectoryListing(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartFileTransfer", func(t *testing.T) {
        input := &transfer.StartFileTransferInput{}
        output := &transfer.StartFileTransferOutput{}

        mockClient.On("StartFileTransfer", ctx, input).Return(output, nil)

        result, err := mockClient.StartFileTransfer(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartServer", func(t *testing.T) {
        input := &transfer.StartServerInput{}
        output := &transfer.StartServerOutput{}

        mockClient.On("StartServer", ctx, input).Return(output, nil)

        result, err := mockClient.StartServer(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopServer", func(t *testing.T) {
        input := &transfer.StopServerInput{}
        output := &transfer.StopServerOutput{}

        mockClient.On("StopServer", ctx, input).Return(output, nil)

        result, err := mockClient.StopServer(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &transfer.TagResourceInput{}
        output := &transfer.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTestConnection", func(t *testing.T) {
        input := &transfer.TestConnectionInput{}
        output := &transfer.TestConnectionOutput{}

        mockClient.On("TestConnection", ctx, input).Return(output, nil)

        result, err := mockClient.TestConnection(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTestIdentityProvider", func(t *testing.T) {
        input := &transfer.TestIdentityProviderInput{}
        output := &transfer.TestIdentityProviderOutput{}

        mockClient.On("TestIdentityProvider", ctx, input).Return(output, nil)

        result, err := mockClient.TestIdentityProvider(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &transfer.UntagResourceInput{}
        output := &transfer.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateAccess", func(t *testing.T) {
        input := &transfer.UpdateAccessInput{}
        output := &transfer.UpdateAccessOutput{}

        mockClient.On("UpdateAccess", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateAccess(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateAgreement", func(t *testing.T) {
        input := &transfer.UpdateAgreementInput{}
        output := &transfer.UpdateAgreementOutput{}

        mockClient.On("UpdateAgreement", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateAgreement(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateCertificate", func(t *testing.T) {
        input := &transfer.UpdateCertificateInput{}
        output := &transfer.UpdateCertificateOutput{}

        mockClient.On("UpdateCertificate", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateCertificate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateConnector", func(t *testing.T) {
        input := &transfer.UpdateConnectorInput{}
        output := &transfer.UpdateConnectorOutput{}

        mockClient.On("UpdateConnector", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateConnector(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateHostKey", func(t *testing.T) {
        input := &transfer.UpdateHostKeyInput{}
        output := &transfer.UpdateHostKeyOutput{}

        mockClient.On("UpdateHostKey", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateHostKey(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateProfile", func(t *testing.T) {
        input := &transfer.UpdateProfileInput{}
        output := &transfer.UpdateProfileOutput{}

        mockClient.On("UpdateProfile", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateProfile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateServer", func(t *testing.T) {
        input := &transfer.UpdateServerInput{}
        output := &transfer.UpdateServerOutput{}

        mockClient.On("UpdateServer", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateServer(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateUser", func(t *testing.T) {
        input := &transfer.UpdateUserInput{}
        output := &transfer.UpdateUserOutput{}

        mockClient.On("UpdateUser", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateUser(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
