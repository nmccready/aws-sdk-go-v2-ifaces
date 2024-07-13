package devicefarm_test

// tests for the devicefarm service interface for this ../../../service/devicefarm/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/devicefarm"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/devicefarm/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/devicefarm/devicefarm_iface"
	"github.com/stretchr/testify/assert"
)

func TestDevicefarmServiceCanBeMocked(t *testing.T) {
	var iface devicefarm_iface.IClient
	iface = &devicefarm.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := devicefarm.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateDevicePool", func(t *testing.T) {
        input := &devicefarm.CreateDevicePoolInput{}
        output := &devicefarm.CreateDevicePoolOutput{}

        mockClient.On("CreateDevicePool", ctx, input).Return(output, nil)

        result, err := mockClient.CreateDevicePool(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateInstanceProfile", func(t *testing.T) {
        input := &devicefarm.CreateInstanceProfileInput{}
        output := &devicefarm.CreateInstanceProfileOutput{}

        mockClient.On("CreateInstanceProfile", ctx, input).Return(output, nil)

        result, err := mockClient.CreateInstanceProfile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateNetworkProfile", func(t *testing.T) {
        input := &devicefarm.CreateNetworkProfileInput{}
        output := &devicefarm.CreateNetworkProfileOutput{}

        mockClient.On("CreateNetworkProfile", ctx, input).Return(output, nil)

        result, err := mockClient.CreateNetworkProfile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateProject", func(t *testing.T) {
        input := &devicefarm.CreateProjectInput{}
        output := &devicefarm.CreateProjectOutput{}

        mockClient.On("CreateProject", ctx, input).Return(output, nil)

        result, err := mockClient.CreateProject(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateRemoteAccessSession", func(t *testing.T) {
        input := &devicefarm.CreateRemoteAccessSessionInput{}
        output := &devicefarm.CreateRemoteAccessSessionOutput{}

        mockClient.On("CreateRemoteAccessSession", ctx, input).Return(output, nil)

        result, err := mockClient.CreateRemoteAccessSession(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateTestGridProject", func(t *testing.T) {
        input := &devicefarm.CreateTestGridProjectInput{}
        output := &devicefarm.CreateTestGridProjectOutput{}

        mockClient.On("CreateTestGridProject", ctx, input).Return(output, nil)

        result, err := mockClient.CreateTestGridProject(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateTestGridUrl", func(t *testing.T) {
        input := &devicefarm.CreateTestGridUrlInput{}
        output := &devicefarm.CreateTestGridUrlOutput{}

        mockClient.On("CreateTestGridUrl", ctx, input).Return(output, nil)

        result, err := mockClient.CreateTestGridUrl(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateUpload", func(t *testing.T) {
        input := &devicefarm.CreateUploadInput{}
        output := &devicefarm.CreateUploadOutput{}

        mockClient.On("CreateUpload", ctx, input).Return(output, nil)

        result, err := mockClient.CreateUpload(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateVPCEConfiguration", func(t *testing.T) {
        input := &devicefarm.CreateVPCEConfigurationInput{}
        output := &devicefarm.CreateVPCEConfigurationOutput{}

        mockClient.On("CreateVPCEConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.CreateVPCEConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteDevicePool", func(t *testing.T) {
        input := &devicefarm.DeleteDevicePoolInput{}
        output := &devicefarm.DeleteDevicePoolOutput{}

        mockClient.On("DeleteDevicePool", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteDevicePool(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteInstanceProfile", func(t *testing.T) {
        input := &devicefarm.DeleteInstanceProfileInput{}
        output := &devicefarm.DeleteInstanceProfileOutput{}

        mockClient.On("DeleteInstanceProfile", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteInstanceProfile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteNetworkProfile", func(t *testing.T) {
        input := &devicefarm.DeleteNetworkProfileInput{}
        output := &devicefarm.DeleteNetworkProfileOutput{}

        mockClient.On("DeleteNetworkProfile", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteNetworkProfile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteProject", func(t *testing.T) {
        input := &devicefarm.DeleteProjectInput{}
        output := &devicefarm.DeleteProjectOutput{}

        mockClient.On("DeleteProject", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteProject(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteRemoteAccessSession", func(t *testing.T) {
        input := &devicefarm.DeleteRemoteAccessSessionInput{}
        output := &devicefarm.DeleteRemoteAccessSessionOutput{}

        mockClient.On("DeleteRemoteAccessSession", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteRemoteAccessSession(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteRun", func(t *testing.T) {
        input := &devicefarm.DeleteRunInput{}
        output := &devicefarm.DeleteRunOutput{}

        mockClient.On("DeleteRun", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteRun(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteTestGridProject", func(t *testing.T) {
        input := &devicefarm.DeleteTestGridProjectInput{}
        output := &devicefarm.DeleteTestGridProjectOutput{}

        mockClient.On("DeleteTestGridProject", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteTestGridProject(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteUpload", func(t *testing.T) {
        input := &devicefarm.DeleteUploadInput{}
        output := &devicefarm.DeleteUploadOutput{}

        mockClient.On("DeleteUpload", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteUpload(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteVPCEConfiguration", func(t *testing.T) {
        input := &devicefarm.DeleteVPCEConfigurationInput{}
        output := &devicefarm.DeleteVPCEConfigurationOutput{}

        mockClient.On("DeleteVPCEConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteVPCEConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetAccountSettings", func(t *testing.T) {
        input := &devicefarm.GetAccountSettingsInput{}
        output := &devicefarm.GetAccountSettingsOutput{}

        mockClient.On("GetAccountSettings", ctx, input).Return(output, nil)

        result, err := mockClient.GetAccountSettings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDevice", func(t *testing.T) {
        input := &devicefarm.GetDeviceInput{}
        output := &devicefarm.GetDeviceOutput{}

        mockClient.On("GetDevice", ctx, input).Return(output, nil)

        result, err := mockClient.GetDevice(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDeviceInstance", func(t *testing.T) {
        input := &devicefarm.GetDeviceInstanceInput{}
        output := &devicefarm.GetDeviceInstanceOutput{}

        mockClient.On("GetDeviceInstance", ctx, input).Return(output, nil)

        result, err := mockClient.GetDeviceInstance(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDevicePool", func(t *testing.T) {
        input := &devicefarm.GetDevicePoolInput{}
        output := &devicefarm.GetDevicePoolOutput{}

        mockClient.On("GetDevicePool", ctx, input).Return(output, nil)

        result, err := mockClient.GetDevicePool(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetDevicePoolCompatibility", func(t *testing.T) {
        input := &devicefarm.GetDevicePoolCompatibilityInput{}
        output := &devicefarm.GetDevicePoolCompatibilityOutput{}

        mockClient.On("GetDevicePoolCompatibility", ctx, input).Return(output, nil)

        result, err := mockClient.GetDevicePoolCompatibility(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetInstanceProfile", func(t *testing.T) {
        input := &devicefarm.GetInstanceProfileInput{}
        output := &devicefarm.GetInstanceProfileOutput{}

        mockClient.On("GetInstanceProfile", ctx, input).Return(output, nil)

        result, err := mockClient.GetInstanceProfile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetJob", func(t *testing.T) {
        input := &devicefarm.GetJobInput{}
        output := &devicefarm.GetJobOutput{}

        mockClient.On("GetJob", ctx, input).Return(output, nil)

        result, err := mockClient.GetJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetNetworkProfile", func(t *testing.T) {
        input := &devicefarm.GetNetworkProfileInput{}
        output := &devicefarm.GetNetworkProfileOutput{}

        mockClient.On("GetNetworkProfile", ctx, input).Return(output, nil)

        result, err := mockClient.GetNetworkProfile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetOfferingStatus", func(t *testing.T) {
        input := &devicefarm.GetOfferingStatusInput{}
        output := &devicefarm.GetOfferingStatusOutput{}

        mockClient.On("GetOfferingStatus", ctx, input).Return(output, nil)

        result, err := mockClient.GetOfferingStatus(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetProject", func(t *testing.T) {
        input := &devicefarm.GetProjectInput{}
        output := &devicefarm.GetProjectOutput{}

        mockClient.On("GetProject", ctx, input).Return(output, nil)

        result, err := mockClient.GetProject(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetRemoteAccessSession", func(t *testing.T) {
        input := &devicefarm.GetRemoteAccessSessionInput{}
        output := &devicefarm.GetRemoteAccessSessionOutput{}

        mockClient.On("GetRemoteAccessSession", ctx, input).Return(output, nil)

        result, err := mockClient.GetRemoteAccessSession(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetRun", func(t *testing.T) {
        input := &devicefarm.GetRunInput{}
        output := &devicefarm.GetRunOutput{}

        mockClient.On("GetRun", ctx, input).Return(output, nil)

        result, err := mockClient.GetRun(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSuite", func(t *testing.T) {
        input := &devicefarm.GetSuiteInput{}
        output := &devicefarm.GetSuiteOutput{}

        mockClient.On("GetSuite", ctx, input).Return(output, nil)

        result, err := mockClient.GetSuite(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetTest", func(t *testing.T) {
        input := &devicefarm.GetTestInput{}
        output := &devicefarm.GetTestOutput{}

        mockClient.On("GetTest", ctx, input).Return(output, nil)

        result, err := mockClient.GetTest(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetTestGridProject", func(t *testing.T) {
        input := &devicefarm.GetTestGridProjectInput{}
        output := &devicefarm.GetTestGridProjectOutput{}

        mockClient.On("GetTestGridProject", ctx, input).Return(output, nil)

        result, err := mockClient.GetTestGridProject(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetTestGridSession", func(t *testing.T) {
        input := &devicefarm.GetTestGridSessionInput{}
        output := &devicefarm.GetTestGridSessionOutput{}

        mockClient.On("GetTestGridSession", ctx, input).Return(output, nil)

        result, err := mockClient.GetTestGridSession(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetUpload", func(t *testing.T) {
        input := &devicefarm.GetUploadInput{}
        output := &devicefarm.GetUploadOutput{}

        mockClient.On("GetUpload", ctx, input).Return(output, nil)

        result, err := mockClient.GetUpload(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetVPCEConfiguration", func(t *testing.T) {
        input := &devicefarm.GetVPCEConfigurationInput{}
        output := &devicefarm.GetVPCEConfigurationOutput{}

        mockClient.On("GetVPCEConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.GetVPCEConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestInstallToRemoteAccessSession", func(t *testing.T) {
        input := &devicefarm.InstallToRemoteAccessSessionInput{}
        output := &devicefarm.InstallToRemoteAccessSessionOutput{}

        mockClient.On("InstallToRemoteAccessSession", ctx, input).Return(output, nil)

        result, err := mockClient.InstallToRemoteAccessSession(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListArtifacts", func(t *testing.T) {
        input := &devicefarm.ListArtifactsInput{}
        output := &devicefarm.ListArtifactsOutput{}

        mockClient.On("ListArtifacts", ctx, input).Return(output, nil)

        result, err := mockClient.ListArtifacts(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDeviceInstances", func(t *testing.T) {
        input := &devicefarm.ListDeviceInstancesInput{}
        output := &devicefarm.ListDeviceInstancesOutput{}

        mockClient.On("ListDeviceInstances", ctx, input).Return(output, nil)

        result, err := mockClient.ListDeviceInstances(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDevicePools", func(t *testing.T) {
        input := &devicefarm.ListDevicePoolsInput{}
        output := &devicefarm.ListDevicePoolsOutput{}

        mockClient.On("ListDevicePools", ctx, input).Return(output, nil)

        result, err := mockClient.ListDevicePools(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListDevices", func(t *testing.T) {
        input := &devicefarm.ListDevicesInput{}
        output := &devicefarm.ListDevicesOutput{}

        mockClient.On("ListDevices", ctx, input).Return(output, nil)

        result, err := mockClient.ListDevices(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListInstanceProfiles", func(t *testing.T) {
        input := &devicefarm.ListInstanceProfilesInput{}
        output := &devicefarm.ListInstanceProfilesOutput{}

        mockClient.On("ListInstanceProfiles", ctx, input).Return(output, nil)

        result, err := mockClient.ListInstanceProfiles(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListJobs", func(t *testing.T) {
        input := &devicefarm.ListJobsInput{}
        output := &devicefarm.ListJobsOutput{}

        mockClient.On("ListJobs", ctx, input).Return(output, nil)

        result, err := mockClient.ListJobs(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListNetworkProfiles", func(t *testing.T) {
        input := &devicefarm.ListNetworkProfilesInput{}
        output := &devicefarm.ListNetworkProfilesOutput{}

        mockClient.On("ListNetworkProfiles", ctx, input).Return(output, nil)

        result, err := mockClient.ListNetworkProfiles(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListOfferingPromotions", func(t *testing.T) {
        input := &devicefarm.ListOfferingPromotionsInput{}
        output := &devicefarm.ListOfferingPromotionsOutput{}

        mockClient.On("ListOfferingPromotions", ctx, input).Return(output, nil)

        result, err := mockClient.ListOfferingPromotions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListOfferingTransactions", func(t *testing.T) {
        input := &devicefarm.ListOfferingTransactionsInput{}
        output := &devicefarm.ListOfferingTransactionsOutput{}

        mockClient.On("ListOfferingTransactions", ctx, input).Return(output, nil)

        result, err := mockClient.ListOfferingTransactions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListOfferings", func(t *testing.T) {
        input := &devicefarm.ListOfferingsInput{}
        output := &devicefarm.ListOfferingsOutput{}

        mockClient.On("ListOfferings", ctx, input).Return(output, nil)

        result, err := mockClient.ListOfferings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListProjects", func(t *testing.T) {
        input := &devicefarm.ListProjectsInput{}
        output := &devicefarm.ListProjectsOutput{}

        mockClient.On("ListProjects", ctx, input).Return(output, nil)

        result, err := mockClient.ListProjects(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListRemoteAccessSessions", func(t *testing.T) {
        input := &devicefarm.ListRemoteAccessSessionsInput{}
        output := &devicefarm.ListRemoteAccessSessionsOutput{}

        mockClient.On("ListRemoteAccessSessions", ctx, input).Return(output, nil)

        result, err := mockClient.ListRemoteAccessSessions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListRuns", func(t *testing.T) {
        input := &devicefarm.ListRunsInput{}
        output := &devicefarm.ListRunsOutput{}

        mockClient.On("ListRuns", ctx, input).Return(output, nil)

        result, err := mockClient.ListRuns(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListSamples", func(t *testing.T) {
        input := &devicefarm.ListSamplesInput{}
        output := &devicefarm.ListSamplesOutput{}

        mockClient.On("ListSamples", ctx, input).Return(output, nil)

        result, err := mockClient.ListSamples(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListSuites", func(t *testing.T) {
        input := &devicefarm.ListSuitesInput{}
        output := &devicefarm.ListSuitesOutput{}

        mockClient.On("ListSuites", ctx, input).Return(output, nil)

        result, err := mockClient.ListSuites(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &devicefarm.ListTagsForResourceInput{}
        output := &devicefarm.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTestGridProjects", func(t *testing.T) {
        input := &devicefarm.ListTestGridProjectsInput{}
        output := &devicefarm.ListTestGridProjectsOutput{}

        mockClient.On("ListTestGridProjects", ctx, input).Return(output, nil)

        result, err := mockClient.ListTestGridProjects(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTestGridSessionActions", func(t *testing.T) {
        input := &devicefarm.ListTestGridSessionActionsInput{}
        output := &devicefarm.ListTestGridSessionActionsOutput{}

        mockClient.On("ListTestGridSessionActions", ctx, input).Return(output, nil)

        result, err := mockClient.ListTestGridSessionActions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTestGridSessionArtifacts", func(t *testing.T) {
        input := &devicefarm.ListTestGridSessionArtifactsInput{}
        output := &devicefarm.ListTestGridSessionArtifactsOutput{}

        mockClient.On("ListTestGridSessionArtifacts", ctx, input).Return(output, nil)

        result, err := mockClient.ListTestGridSessionArtifacts(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTestGridSessions", func(t *testing.T) {
        input := &devicefarm.ListTestGridSessionsInput{}
        output := &devicefarm.ListTestGridSessionsOutput{}

        mockClient.On("ListTestGridSessions", ctx, input).Return(output, nil)

        result, err := mockClient.ListTestGridSessions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTests", func(t *testing.T) {
        input := &devicefarm.ListTestsInput{}
        output := &devicefarm.ListTestsOutput{}

        mockClient.On("ListTests", ctx, input).Return(output, nil)

        result, err := mockClient.ListTests(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListUniqueProblems", func(t *testing.T) {
        input := &devicefarm.ListUniqueProblemsInput{}
        output := &devicefarm.ListUniqueProblemsOutput{}

        mockClient.On("ListUniqueProblems", ctx, input).Return(output, nil)

        result, err := mockClient.ListUniqueProblems(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListUploads", func(t *testing.T) {
        input := &devicefarm.ListUploadsInput{}
        output := &devicefarm.ListUploadsOutput{}

        mockClient.On("ListUploads", ctx, input).Return(output, nil)

        result, err := mockClient.ListUploads(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListVPCEConfigurations", func(t *testing.T) {
        input := &devicefarm.ListVPCEConfigurationsInput{}
        output := &devicefarm.ListVPCEConfigurationsOutput{}

        mockClient.On("ListVPCEConfigurations", ctx, input).Return(output, nil)

        result, err := mockClient.ListVPCEConfigurations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPurchaseOffering", func(t *testing.T) {
        input := &devicefarm.PurchaseOfferingInput{}
        output := &devicefarm.PurchaseOfferingOutput{}

        mockClient.On("PurchaseOffering", ctx, input).Return(output, nil)

        result, err := mockClient.PurchaseOffering(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestRenewOffering", func(t *testing.T) {
        input := &devicefarm.RenewOfferingInput{}
        output := &devicefarm.RenewOfferingOutput{}

        mockClient.On("RenewOffering", ctx, input).Return(output, nil)

        result, err := mockClient.RenewOffering(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestScheduleRun", func(t *testing.T) {
        input := &devicefarm.ScheduleRunInput{}
        output := &devicefarm.ScheduleRunOutput{}

        mockClient.On("ScheduleRun", ctx, input).Return(output, nil)

        result, err := mockClient.ScheduleRun(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopJob", func(t *testing.T) {
        input := &devicefarm.StopJobInput{}
        output := &devicefarm.StopJobOutput{}

        mockClient.On("StopJob", ctx, input).Return(output, nil)

        result, err := mockClient.StopJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopRemoteAccessSession", func(t *testing.T) {
        input := &devicefarm.StopRemoteAccessSessionInput{}
        output := &devicefarm.StopRemoteAccessSessionOutput{}

        mockClient.On("StopRemoteAccessSession", ctx, input).Return(output, nil)

        result, err := mockClient.StopRemoteAccessSession(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopRun", func(t *testing.T) {
        input := &devicefarm.StopRunInput{}
        output := &devicefarm.StopRunOutput{}

        mockClient.On("StopRun", ctx, input).Return(output, nil)

        result, err := mockClient.StopRun(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &devicefarm.TagResourceInput{}
        output := &devicefarm.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &devicefarm.UntagResourceInput{}
        output := &devicefarm.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateDeviceInstance", func(t *testing.T) {
        input := &devicefarm.UpdateDeviceInstanceInput{}
        output := &devicefarm.UpdateDeviceInstanceOutput{}

        mockClient.On("UpdateDeviceInstance", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateDeviceInstance(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateDevicePool", func(t *testing.T) {
        input := &devicefarm.UpdateDevicePoolInput{}
        output := &devicefarm.UpdateDevicePoolOutput{}

        mockClient.On("UpdateDevicePool", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateDevicePool(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateInstanceProfile", func(t *testing.T) {
        input := &devicefarm.UpdateInstanceProfileInput{}
        output := &devicefarm.UpdateInstanceProfileOutput{}

        mockClient.On("UpdateInstanceProfile", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateInstanceProfile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateNetworkProfile", func(t *testing.T) {
        input := &devicefarm.UpdateNetworkProfileInput{}
        output := &devicefarm.UpdateNetworkProfileOutput{}

        mockClient.On("UpdateNetworkProfile", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateNetworkProfile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateProject", func(t *testing.T) {
        input := &devicefarm.UpdateProjectInput{}
        output := &devicefarm.UpdateProjectOutput{}

        mockClient.On("UpdateProject", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateProject(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateTestGridProject", func(t *testing.T) {
        input := &devicefarm.UpdateTestGridProjectInput{}
        output := &devicefarm.UpdateTestGridProjectOutput{}

        mockClient.On("UpdateTestGridProject", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateTestGridProject(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateUpload", func(t *testing.T) {
        input := &devicefarm.UpdateUploadInput{}
        output := &devicefarm.UpdateUploadOutput{}

        mockClient.On("UpdateUpload", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateUpload(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateVPCEConfiguration", func(t *testing.T) {
        input := &devicefarm.UpdateVPCEConfigurationInput{}
        output := &devicefarm.UpdateVPCEConfigurationOutput{}

        mockClient.On("UpdateVPCEConfiguration", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateVPCEConfiguration(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
