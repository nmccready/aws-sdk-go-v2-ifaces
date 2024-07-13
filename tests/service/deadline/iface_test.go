package deadline_test

// tests for the deadline service interface for this ../../../service/deadline/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/deadline"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/deadline/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/deadline/deadline_iface"
	"github.com/stretchr/testify/assert"
)

func TestDeadlineServiceCanBeMocked(t *testing.T) {
	var iface deadline_iface.IClient
	iface = &deadline.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := deadline.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateMemberToFarm", func(t *testing.T) {
        input := &deadline.AssociateMemberToFarmInput{}
        output := &deadline.AssociateMemberToFarmOutput{}

        mockClient.On("AssociateMemberToFarm", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateMemberToFarm(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateMemberToFleet", func(t *testing.T) {
        input := &deadline.AssociateMemberToFleetInput{}
        output := &deadline.AssociateMemberToFleetOutput{}

        mockClient.On("AssociateMemberToFleet", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateMemberToFleet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateMemberToJob", func(t *testing.T) {
        input := &deadline.AssociateMemberToJobInput{}
        output := &deadline.AssociateMemberToJobOutput{}

        mockClient.On("AssociateMemberToJob", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateMemberToJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateMemberToQueue", func(t *testing.T) {
        input := &deadline.AssociateMemberToQueueInput{}
        output := &deadline.AssociateMemberToQueueOutput{}

        mockClient.On("AssociateMemberToQueue", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateMemberToQueue(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssumeFleetRoleForRead", func(t *testing.T) {
        input := &deadline.AssumeFleetRoleForReadInput{}
        output := &deadline.AssumeFleetRoleForReadOutput{}

        mockClient.On("AssumeFleetRoleForRead", ctx, input).Return(output, nil)

        result, err := mockClient.AssumeFleetRoleForRead(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssumeFleetRoleForWorker", func(t *testing.T) {
        input := &deadline.AssumeFleetRoleForWorkerInput{}
        output := &deadline.AssumeFleetRoleForWorkerOutput{}

        mockClient.On("AssumeFleetRoleForWorker", ctx, input).Return(output, nil)

        result, err := mockClient.AssumeFleetRoleForWorker(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssumeQueueRoleForRead", func(t *testing.T) {
        input := &deadline.AssumeQueueRoleForReadInput{}
        output := &deadline.AssumeQueueRoleForReadOutput{}

        mockClient.On("AssumeQueueRoleForRead", ctx, input).Return(output, nil)

        result, err := mockClient.AssumeQueueRoleForRead(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssumeQueueRoleForUser", func(t *testing.T) {
        input := &deadline.AssumeQueueRoleForUserInput{}
        output := &deadline.AssumeQueueRoleForUserOutput{}

        mockClient.On("AssumeQueueRoleForUser", ctx, input).Return(output, nil)

        result, err := mockClient.AssumeQueueRoleForUser(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssumeQueueRoleForWorker", func(t *testing.T) {
        input := &deadline.AssumeQueueRoleForWorkerInput{}
        output := &deadline.AssumeQueueRoleForWorkerOutput{}

        mockClient.On("AssumeQueueRoleForWorker", ctx, input).Return(output, nil)

        result, err := mockClient.AssumeQueueRoleForWorker(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestBatchGetJobEntity", func(t *testing.T) {
        input := &deadline.BatchGetJobEntityInput{}
        output := &deadline.BatchGetJobEntityOutput{}

        mockClient.On("BatchGetJobEntity", ctx, input).Return(output, nil)

        result, err := mockClient.BatchGetJobEntity(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCopyJobTemplate", func(t *testing.T) {
        input := &deadline.CopyJobTemplateInput{}
        output := &deadline.CopyJobTemplateOutput{}

        mockClient.On("CopyJobTemplate", ctx, input).Return(output, nil)

        result, err := mockClient.CopyJobTemplate(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateBudget", func(t *testing.T) {
        input := &deadline.CreateBudgetInput{}
        output := &deadline.CreateBudgetOutput{}

        mockClient.On("CreateBudget", ctx, input).Return(output, nil)

        result, err := mockClient.CreateBudget(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateFarm", func(t *testing.T) {
        input := &deadline.CreateFarmInput{}
        output := &deadline.CreateFarmOutput{}

        mockClient.On("CreateFarm", ctx, input).Return(output, nil)

        result, err := mockClient.CreateFarm(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateFleet", func(t *testing.T) {
        input := &deadline.CreateFleetInput{}
        output := &deadline.CreateFleetOutput{}

        mockClient.On("CreateFleet", ctx, input).Return(output, nil)

        result, err := mockClient.CreateFleet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateJob", func(t *testing.T) {
        input := &deadline.CreateJobInput{}
        output := &deadline.CreateJobOutput{}

        mockClient.On("CreateJob", ctx, input).Return(output, nil)

        result, err := mockClient.CreateJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateLicenseEndpoint", func(t *testing.T) {
        input := &deadline.CreateLicenseEndpointInput{}
        output := &deadline.CreateLicenseEndpointOutput{}

        mockClient.On("CreateLicenseEndpoint", ctx, input).Return(output, nil)

        result, err := mockClient.CreateLicenseEndpoint(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateMonitor", func(t *testing.T) {
        input := &deadline.CreateMonitorInput{}
        output := &deadline.CreateMonitorOutput{}

        mockClient.On("CreateMonitor", ctx, input).Return(output, nil)

        result, err := mockClient.CreateMonitor(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateQueue", func(t *testing.T) {
        input := &deadline.CreateQueueInput{}
        output := &deadline.CreateQueueOutput{}

        mockClient.On("CreateQueue", ctx, input).Return(output, nil)

        result, err := mockClient.CreateQueue(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateQueueEnvironment", func(t *testing.T) {
        input := &deadline.CreateQueueEnvironmentInput{}
        output := &deadline.CreateQueueEnvironmentOutput{}

        mockClient.On("CreateQueueEnvironment", ctx, input).Return(output, nil)

        result, err := mockClient.CreateQueueEnvironment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateQueueFleetAssociation", func(t *testing.T) {
        input := &deadline.CreateQueueFleetAssociationInput{}
        output := &deadline.CreateQueueFleetAssociationOutput{}

        mockClient.On("CreateQueueFleetAssociation", ctx, input).Return(output, nil)

        result, err := mockClient.CreateQueueFleetAssociation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateStorageProfile", func(t *testing.T) {
        input := &deadline.CreateStorageProfileInput{}
        output := &deadline.CreateStorageProfileOutput{}

        mockClient.On("CreateStorageProfile", ctx, input).Return(output, nil)

        result, err := mockClient.CreateStorageProfile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateWorker", func(t *testing.T) {
        input := &deadline.CreateWorkerInput{}
        output := &deadline.CreateWorkerOutput{}

        mockClient.On("CreateWorker", ctx, input).Return(output, nil)

        result, err := mockClient.CreateWorker(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteBudget", func(t *testing.T) {
        input := &deadline.DeleteBudgetInput{}
        output := &deadline.DeleteBudgetOutput{}

        mockClient.On("DeleteBudget", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteBudget(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteFarm", func(t *testing.T) {
        input := &deadline.DeleteFarmInput{}
        output := &deadline.DeleteFarmOutput{}

        mockClient.On("DeleteFarm", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteFarm(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteFleet", func(t *testing.T) {
        input := &deadline.DeleteFleetInput{}
        output := &deadline.DeleteFleetOutput{}

        mockClient.On("DeleteFleet", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteFleet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteLicenseEndpoint", func(t *testing.T) {
        input := &deadline.DeleteLicenseEndpointInput{}
        output := &deadline.DeleteLicenseEndpointOutput{}

        mockClient.On("DeleteLicenseEndpoint", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteLicenseEndpoint(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteMeteredProduct", func(t *testing.T) {
        input := &deadline.DeleteMeteredProductInput{}
        output := &deadline.DeleteMeteredProductOutput{}

        mockClient.On("DeleteMeteredProduct", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteMeteredProduct(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteMonitor", func(t *testing.T) {
        input := &deadline.DeleteMonitorInput{}
        output := &deadline.DeleteMonitorOutput{}

        mockClient.On("DeleteMonitor", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteMonitor(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteQueue", func(t *testing.T) {
        input := &deadline.DeleteQueueInput{}
        output := &deadline.DeleteQueueOutput{}

        mockClient.On("DeleteQueue", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteQueue(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteQueueEnvironment", func(t *testing.T) {
        input := &deadline.DeleteQueueEnvironmentInput{}
        output := &deadline.DeleteQueueEnvironmentOutput{}

        mockClient.On("DeleteQueueEnvironment", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteQueueEnvironment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteQueueFleetAssociation", func(t *testing.T) {
        input := &deadline.DeleteQueueFleetAssociationInput{}
        output := &deadline.DeleteQueueFleetAssociationOutput{}

        mockClient.On("DeleteQueueFleetAssociation", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteQueueFleetAssociation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteStorageProfile", func(t *testing.T) {
        input := &deadline.DeleteStorageProfileInput{}
        output := &deadline.DeleteStorageProfileOutput{}

        mockClient.On("DeleteStorageProfile", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteStorageProfile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteWorker", func(t *testing.T) {
        input := &deadline.DeleteWorkerInput{}
        output := &deadline.DeleteWorkerOutput{}

        mockClient.On("DeleteWorker", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteWorker(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateMemberFromFarm", func(t *testing.T) {
        input := &deadline.DisassociateMemberFromFarmInput{}
        output := &deadline.DisassociateMemberFromFarmOutput{}

        mockClient.On("DisassociateMemberFromFarm", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateMemberFromFarm(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateMemberFromFleet", func(t *testing.T) {
        input := &deadline.DisassociateMemberFromFleetInput{}
        output := &deadline.DisassociateMemberFromFleetOutput{}

        mockClient.On("DisassociateMemberFromFleet", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateMemberFromFleet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateMemberFromJob", func(t *testing.T) {
        input := &deadline.DisassociateMemberFromJobInput{}
        output := &deadline.DisassociateMemberFromJobOutput{}

        mockClient.On("DisassociateMemberFromJob", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateMemberFromJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateMemberFromQueue", func(t *testing.T) {
        input := &deadline.DisassociateMemberFromQueueInput{}
        output := &deadline.DisassociateMemberFromQueueOutput{}

        mockClient.On("DisassociateMemberFromQueue", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateMemberFromQueue(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetBudget", func(t *testing.T) {
        input := &deadline.GetBudgetInput{}
        output := &deadline.GetBudgetOutput{}

        mockClient.On("GetBudget", ctx, input).Return(output, nil)

        result, err := mockClient.GetBudget(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetFarm", func(t *testing.T) {
        input := &deadline.GetFarmInput{}
        output := &deadline.GetFarmOutput{}

        mockClient.On("GetFarm", ctx, input).Return(output, nil)

        result, err := mockClient.GetFarm(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetFleet", func(t *testing.T) {
        input := &deadline.GetFleetInput{}
        output := &deadline.GetFleetOutput{}

        mockClient.On("GetFleet", ctx, input).Return(output, nil)

        result, err := mockClient.GetFleet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetJob", func(t *testing.T) {
        input := &deadline.GetJobInput{}
        output := &deadline.GetJobOutput{}

        mockClient.On("GetJob", ctx, input).Return(output, nil)

        result, err := mockClient.GetJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetLicenseEndpoint", func(t *testing.T) {
        input := &deadline.GetLicenseEndpointInput{}
        output := &deadline.GetLicenseEndpointOutput{}

        mockClient.On("GetLicenseEndpoint", ctx, input).Return(output, nil)

        result, err := mockClient.GetLicenseEndpoint(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetMonitor", func(t *testing.T) {
        input := &deadline.GetMonitorInput{}
        output := &deadline.GetMonitorOutput{}

        mockClient.On("GetMonitor", ctx, input).Return(output, nil)

        result, err := mockClient.GetMonitor(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetQueue", func(t *testing.T) {
        input := &deadline.GetQueueInput{}
        output := &deadline.GetQueueOutput{}

        mockClient.On("GetQueue", ctx, input).Return(output, nil)

        result, err := mockClient.GetQueue(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetQueueEnvironment", func(t *testing.T) {
        input := &deadline.GetQueueEnvironmentInput{}
        output := &deadline.GetQueueEnvironmentOutput{}

        mockClient.On("GetQueueEnvironment", ctx, input).Return(output, nil)

        result, err := mockClient.GetQueueEnvironment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetQueueFleetAssociation", func(t *testing.T) {
        input := &deadline.GetQueueFleetAssociationInput{}
        output := &deadline.GetQueueFleetAssociationOutput{}

        mockClient.On("GetQueueFleetAssociation", ctx, input).Return(output, nil)

        result, err := mockClient.GetQueueFleetAssociation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSession", func(t *testing.T) {
        input := &deadline.GetSessionInput{}
        output := &deadline.GetSessionOutput{}

        mockClient.On("GetSession", ctx, input).Return(output, nil)

        result, err := mockClient.GetSession(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSessionAction", func(t *testing.T) {
        input := &deadline.GetSessionActionInput{}
        output := &deadline.GetSessionActionOutput{}

        mockClient.On("GetSessionAction", ctx, input).Return(output, nil)

        result, err := mockClient.GetSessionAction(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSessionsStatisticsAggregation", func(t *testing.T) {
        input := &deadline.GetSessionsStatisticsAggregationInput{}
        output := &deadline.GetSessionsStatisticsAggregationOutput{}

        mockClient.On("GetSessionsStatisticsAggregation", ctx, input).Return(output, nil)

        result, err := mockClient.GetSessionsStatisticsAggregation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetStep", func(t *testing.T) {
        input := &deadline.GetStepInput{}
        output := &deadline.GetStepOutput{}

        mockClient.On("GetStep", ctx, input).Return(output, nil)

        result, err := mockClient.GetStep(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetStorageProfile", func(t *testing.T) {
        input := &deadline.GetStorageProfileInput{}
        output := &deadline.GetStorageProfileOutput{}

        mockClient.On("GetStorageProfile", ctx, input).Return(output, nil)

        result, err := mockClient.GetStorageProfile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetStorageProfileForQueue", func(t *testing.T) {
        input := &deadline.GetStorageProfileForQueueInput{}
        output := &deadline.GetStorageProfileForQueueOutput{}

        mockClient.On("GetStorageProfileForQueue", ctx, input).Return(output, nil)

        result, err := mockClient.GetStorageProfileForQueue(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetTask", func(t *testing.T) {
        input := &deadline.GetTaskInput{}
        output := &deadline.GetTaskOutput{}

        mockClient.On("GetTask", ctx, input).Return(output, nil)

        result, err := mockClient.GetTask(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetWorker", func(t *testing.T) {
        input := &deadline.GetWorkerInput{}
        output := &deadline.GetWorkerOutput{}

        mockClient.On("GetWorker", ctx, input).Return(output, nil)

        result, err := mockClient.GetWorker(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListAvailableMeteredProducts", func(t *testing.T) {
        input := &deadline.ListAvailableMeteredProductsInput{}
        output := &deadline.ListAvailableMeteredProductsOutput{}

        mockClient.On("ListAvailableMeteredProducts", ctx, input).Return(output, nil)

        result, err := mockClient.ListAvailableMeteredProducts(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListBudgets", func(t *testing.T) {
        input := &deadline.ListBudgetsInput{}
        output := &deadline.ListBudgetsOutput{}

        mockClient.On("ListBudgets", ctx, input).Return(output, nil)

        result, err := mockClient.ListBudgets(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListFarmMembers", func(t *testing.T) {
        input := &deadline.ListFarmMembersInput{}
        output := &deadline.ListFarmMembersOutput{}

        mockClient.On("ListFarmMembers", ctx, input).Return(output, nil)

        result, err := mockClient.ListFarmMembers(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListFarms", func(t *testing.T) {
        input := &deadline.ListFarmsInput{}
        output := &deadline.ListFarmsOutput{}

        mockClient.On("ListFarms", ctx, input).Return(output, nil)

        result, err := mockClient.ListFarms(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListFleetMembers", func(t *testing.T) {
        input := &deadline.ListFleetMembersInput{}
        output := &deadline.ListFleetMembersOutput{}

        mockClient.On("ListFleetMembers", ctx, input).Return(output, nil)

        result, err := mockClient.ListFleetMembers(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListFleets", func(t *testing.T) {
        input := &deadline.ListFleetsInput{}
        output := &deadline.ListFleetsOutput{}

        mockClient.On("ListFleets", ctx, input).Return(output, nil)

        result, err := mockClient.ListFleets(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListJobMembers", func(t *testing.T) {
        input := &deadline.ListJobMembersInput{}
        output := &deadline.ListJobMembersOutput{}

        mockClient.On("ListJobMembers", ctx, input).Return(output, nil)

        result, err := mockClient.ListJobMembers(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListJobs", func(t *testing.T) {
        input := &deadline.ListJobsInput{}
        output := &deadline.ListJobsOutput{}

        mockClient.On("ListJobs", ctx, input).Return(output, nil)

        result, err := mockClient.ListJobs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListLicenseEndpoints", func(t *testing.T) {
        input := &deadline.ListLicenseEndpointsInput{}
        output := &deadline.ListLicenseEndpointsOutput{}

        mockClient.On("ListLicenseEndpoints", ctx, input).Return(output, nil)

        result, err := mockClient.ListLicenseEndpoints(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListMeteredProducts", func(t *testing.T) {
        input := &deadline.ListMeteredProductsInput{}
        output := &deadline.ListMeteredProductsOutput{}

        mockClient.On("ListMeteredProducts", ctx, input).Return(output, nil)

        result, err := mockClient.ListMeteredProducts(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListMonitors", func(t *testing.T) {
        input := &deadline.ListMonitorsInput{}
        output := &deadline.ListMonitorsOutput{}

        mockClient.On("ListMonitors", ctx, input).Return(output, nil)

        result, err := mockClient.ListMonitors(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListQueueEnvironments", func(t *testing.T) {
        input := &deadline.ListQueueEnvironmentsInput{}
        output := &deadline.ListQueueEnvironmentsOutput{}

        mockClient.On("ListQueueEnvironments", ctx, input).Return(output, nil)

        result, err := mockClient.ListQueueEnvironments(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListQueueFleetAssociations", func(t *testing.T) {
        input := &deadline.ListQueueFleetAssociationsInput{}
        output := &deadline.ListQueueFleetAssociationsOutput{}

        mockClient.On("ListQueueFleetAssociations", ctx, input).Return(output, nil)

        result, err := mockClient.ListQueueFleetAssociations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListQueueMembers", func(t *testing.T) {
        input := &deadline.ListQueueMembersInput{}
        output := &deadline.ListQueueMembersOutput{}

        mockClient.On("ListQueueMembers", ctx, input).Return(output, nil)

        result, err := mockClient.ListQueueMembers(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListQueues", func(t *testing.T) {
        input := &deadline.ListQueuesInput{}
        output := &deadline.ListQueuesOutput{}

        mockClient.On("ListQueues", ctx, input).Return(output, nil)

        result, err := mockClient.ListQueues(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListSessionActions", func(t *testing.T) {
        input := &deadline.ListSessionActionsInput{}
        output := &deadline.ListSessionActionsOutput{}

        mockClient.On("ListSessionActions", ctx, input).Return(output, nil)

        result, err := mockClient.ListSessionActions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListSessions", func(t *testing.T) {
        input := &deadline.ListSessionsInput{}
        output := &deadline.ListSessionsOutput{}

        mockClient.On("ListSessions", ctx, input).Return(output, nil)

        result, err := mockClient.ListSessions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListSessionsForWorker", func(t *testing.T) {
        input := &deadline.ListSessionsForWorkerInput{}
        output := &deadline.ListSessionsForWorkerOutput{}

        mockClient.On("ListSessionsForWorker", ctx, input).Return(output, nil)

        result, err := mockClient.ListSessionsForWorker(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListStepConsumers", func(t *testing.T) {
        input := &deadline.ListStepConsumersInput{}
        output := &deadline.ListStepConsumersOutput{}

        mockClient.On("ListStepConsumers", ctx, input).Return(output, nil)

        result, err := mockClient.ListStepConsumers(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListStepDependencies", func(t *testing.T) {
        input := &deadline.ListStepDependenciesInput{}
        output := &deadline.ListStepDependenciesOutput{}

        mockClient.On("ListStepDependencies", ctx, input).Return(output, nil)

        result, err := mockClient.ListStepDependencies(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListSteps", func(t *testing.T) {
        input := &deadline.ListStepsInput{}
        output := &deadline.ListStepsOutput{}

        mockClient.On("ListSteps", ctx, input).Return(output, nil)

        result, err := mockClient.ListSteps(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListStorageProfiles", func(t *testing.T) {
        input := &deadline.ListStorageProfilesInput{}
        output := &deadline.ListStorageProfilesOutput{}

        mockClient.On("ListStorageProfiles", ctx, input).Return(output, nil)

        result, err := mockClient.ListStorageProfiles(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListStorageProfilesForQueue", func(t *testing.T) {
        input := &deadline.ListStorageProfilesForQueueInput{}
        output := &deadline.ListStorageProfilesForQueueOutput{}

        mockClient.On("ListStorageProfilesForQueue", ctx, input).Return(output, nil)

        result, err := mockClient.ListStorageProfilesForQueue(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &deadline.ListTagsForResourceInput{}
        output := &deadline.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTasks", func(t *testing.T) {
        input := &deadline.ListTasksInput{}
        output := &deadline.ListTasksOutput{}

        mockClient.On("ListTasks", ctx, input).Return(output, nil)

        result, err := mockClient.ListTasks(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListWorkers", func(t *testing.T) {
        input := &deadline.ListWorkersInput{}
        output := &deadline.ListWorkersOutput{}

        mockClient.On("ListWorkers", ctx, input).Return(output, nil)

        result, err := mockClient.ListWorkers(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutMeteredProduct", func(t *testing.T) {
        input := &deadline.PutMeteredProductInput{}
        output := &deadline.PutMeteredProductOutput{}

        mockClient.On("PutMeteredProduct", ctx, input).Return(output, nil)

        result, err := mockClient.PutMeteredProduct(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSearchJobs", func(t *testing.T) {
        input := &deadline.SearchJobsInput{}
        output := &deadline.SearchJobsOutput{}

        mockClient.On("SearchJobs", ctx, input).Return(output, nil)

        result, err := mockClient.SearchJobs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSearchSteps", func(t *testing.T) {
        input := &deadline.SearchStepsInput{}
        output := &deadline.SearchStepsOutput{}

        mockClient.On("SearchSteps", ctx, input).Return(output, nil)

        result, err := mockClient.SearchSteps(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSearchTasks", func(t *testing.T) {
        input := &deadline.SearchTasksInput{}
        output := &deadline.SearchTasksOutput{}

        mockClient.On("SearchTasks", ctx, input).Return(output, nil)

        result, err := mockClient.SearchTasks(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestSearchWorkers", func(t *testing.T) {
        input := &deadline.SearchWorkersInput{}
        output := &deadline.SearchWorkersOutput{}

        mockClient.On("SearchWorkers", ctx, input).Return(output, nil)

        result, err := mockClient.SearchWorkers(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartSessionsStatisticsAggregation", func(t *testing.T) {
        input := &deadline.StartSessionsStatisticsAggregationInput{}
        output := &deadline.StartSessionsStatisticsAggregationOutput{}

        mockClient.On("StartSessionsStatisticsAggregation", ctx, input).Return(output, nil)

        result, err := mockClient.StartSessionsStatisticsAggregation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &deadline.TagResourceInput{}
        output := &deadline.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &deadline.UntagResourceInput{}
        output := &deadline.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateBudget", func(t *testing.T) {
        input := &deadline.UpdateBudgetInput{}
        output := &deadline.UpdateBudgetOutput{}

        mockClient.On("UpdateBudget", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateBudget(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateFarm", func(t *testing.T) {
        input := &deadline.UpdateFarmInput{}
        output := &deadline.UpdateFarmOutput{}

        mockClient.On("UpdateFarm", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateFarm(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateFleet", func(t *testing.T) {
        input := &deadline.UpdateFleetInput{}
        output := &deadline.UpdateFleetOutput{}

        mockClient.On("UpdateFleet", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateFleet(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateJob", func(t *testing.T) {
        input := &deadline.UpdateJobInput{}
        output := &deadline.UpdateJobOutput{}

        mockClient.On("UpdateJob", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateMonitor", func(t *testing.T) {
        input := &deadline.UpdateMonitorInput{}
        output := &deadline.UpdateMonitorOutput{}

        mockClient.On("UpdateMonitor", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateMonitor(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateQueue", func(t *testing.T) {
        input := &deadline.UpdateQueueInput{}
        output := &deadline.UpdateQueueOutput{}

        mockClient.On("UpdateQueue", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateQueue(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateQueueEnvironment", func(t *testing.T) {
        input := &deadline.UpdateQueueEnvironmentInput{}
        output := &deadline.UpdateQueueEnvironmentOutput{}

        mockClient.On("UpdateQueueEnvironment", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateQueueEnvironment(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateQueueFleetAssociation", func(t *testing.T) {
        input := &deadline.UpdateQueueFleetAssociationInput{}
        output := &deadline.UpdateQueueFleetAssociationOutput{}

        mockClient.On("UpdateQueueFleetAssociation", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateQueueFleetAssociation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateSession", func(t *testing.T) {
        input := &deadline.UpdateSessionInput{}
        output := &deadline.UpdateSessionOutput{}

        mockClient.On("UpdateSession", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateSession(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateStep", func(t *testing.T) {
        input := &deadline.UpdateStepInput{}
        output := &deadline.UpdateStepOutput{}

        mockClient.On("UpdateStep", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateStep(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateStorageProfile", func(t *testing.T) {
        input := &deadline.UpdateStorageProfileInput{}
        output := &deadline.UpdateStorageProfileOutput{}

        mockClient.On("UpdateStorageProfile", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateStorageProfile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateTask", func(t *testing.T) {
        input := &deadline.UpdateTaskInput{}
        output := &deadline.UpdateTaskOutput{}

        mockClient.On("UpdateTask", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateTask(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateWorker", func(t *testing.T) {
        input := &deadline.UpdateWorkerInput{}
        output := &deadline.UpdateWorkerOutput{}

        mockClient.On("UpdateWorker", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateWorker(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateWorkerSchedule", func(t *testing.T) {
        input := &deadline.UpdateWorkerScheduleInput{}
        output := &deadline.UpdateWorkerScheduleOutput{}

        mockClient.On("UpdateWorkerSchedule", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateWorkerSchedule(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
