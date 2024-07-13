package nimble_test

// tests for the nimble service interface for this ../../../service/nimble/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/nimble"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/nimble/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/nimble/nimble_iface"
	"github.com/stretchr/testify/assert"
)

func TestNimbleServiceCanBeMocked(t *testing.T) {
	var iface nimble_iface.IClient
	iface = &nimble.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := nimble.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAcceptEulas", func(t *testing.T) {
        input := &nimble.AcceptEulasInput{}
        output := &nimble.AcceptEulasOutput{}

        mockClient.On("AcceptEulas", ctx, input).Return(output, nil)

        result, err := mockClient.AcceptEulas(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateLaunchProfile", func(t *testing.T) {
        input := &nimble.CreateLaunchProfileInput{}
        output := &nimble.CreateLaunchProfileOutput{}

        mockClient.On("CreateLaunchProfile", ctx, input).Return(output, nil)

        result, err := mockClient.CreateLaunchProfile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateStreamingImage", func(t *testing.T) {
        input := &nimble.CreateStreamingImageInput{}
        output := &nimble.CreateStreamingImageOutput{}

        mockClient.On("CreateStreamingImage", ctx, input).Return(output, nil)

        result, err := mockClient.CreateStreamingImage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateStreamingSession", func(t *testing.T) {
        input := &nimble.CreateStreamingSessionInput{}
        output := &nimble.CreateStreamingSessionOutput{}

        mockClient.On("CreateStreamingSession", ctx, input).Return(output, nil)

        result, err := mockClient.CreateStreamingSession(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateStreamingSessionStream", func(t *testing.T) {
        input := &nimble.CreateStreamingSessionStreamInput{}
        output := &nimble.CreateStreamingSessionStreamOutput{}

        mockClient.On("CreateStreamingSessionStream", ctx, input).Return(output, nil)

        result, err := mockClient.CreateStreamingSessionStream(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateStudio", func(t *testing.T) {
        input := &nimble.CreateStudioInput{}
        output := &nimble.CreateStudioOutput{}

        mockClient.On("CreateStudio", ctx, input).Return(output, nil)

        result, err := mockClient.CreateStudio(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateStudioComponent", func(t *testing.T) {
        input := &nimble.CreateStudioComponentInput{}
        output := &nimble.CreateStudioComponentOutput{}

        mockClient.On("CreateStudioComponent", ctx, input).Return(output, nil)

        result, err := mockClient.CreateStudioComponent(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteLaunchProfile", func(t *testing.T) {
        input := &nimble.DeleteLaunchProfileInput{}
        output := &nimble.DeleteLaunchProfileOutput{}

        mockClient.On("DeleteLaunchProfile", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteLaunchProfile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteLaunchProfileMember", func(t *testing.T) {
        input := &nimble.DeleteLaunchProfileMemberInput{}
        output := &nimble.DeleteLaunchProfileMemberOutput{}

        mockClient.On("DeleteLaunchProfileMember", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteLaunchProfileMember(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteStreamingImage", func(t *testing.T) {
        input := &nimble.DeleteStreamingImageInput{}
        output := &nimble.DeleteStreamingImageOutput{}

        mockClient.On("DeleteStreamingImage", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteStreamingImage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteStreamingSession", func(t *testing.T) {
        input := &nimble.DeleteStreamingSessionInput{}
        output := &nimble.DeleteStreamingSessionOutput{}

        mockClient.On("DeleteStreamingSession", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteStreamingSession(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteStudio", func(t *testing.T) {
        input := &nimble.DeleteStudioInput{}
        output := &nimble.DeleteStudioOutput{}

        mockClient.On("DeleteStudio", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteStudio(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteStudioComponent", func(t *testing.T) {
        input := &nimble.DeleteStudioComponentInput{}
        output := &nimble.DeleteStudioComponentOutput{}

        mockClient.On("DeleteStudioComponent", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteStudioComponent(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteStudioMember", func(t *testing.T) {
        input := &nimble.DeleteStudioMemberInput{}
        output := &nimble.DeleteStudioMemberOutput{}

        mockClient.On("DeleteStudioMember", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteStudioMember(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetEula", func(t *testing.T) {
        input := &nimble.GetEulaInput{}
        output := &nimble.GetEulaOutput{}

        mockClient.On("GetEula", ctx, input).Return(output, nil)

        result, err := mockClient.GetEula(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetLaunchProfile", func(t *testing.T) {
        input := &nimble.GetLaunchProfileInput{}
        output := &nimble.GetLaunchProfileOutput{}

        mockClient.On("GetLaunchProfile", ctx, input).Return(output, nil)

        result, err := mockClient.GetLaunchProfile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetLaunchProfileDetails", func(t *testing.T) {
        input := &nimble.GetLaunchProfileDetailsInput{}
        output := &nimble.GetLaunchProfileDetailsOutput{}

        mockClient.On("GetLaunchProfileDetails", ctx, input).Return(output, nil)

        result, err := mockClient.GetLaunchProfileDetails(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetLaunchProfileInitialization", func(t *testing.T) {
        input := &nimble.GetLaunchProfileInitializationInput{}
        output := &nimble.GetLaunchProfileInitializationOutput{}

        mockClient.On("GetLaunchProfileInitialization", ctx, input).Return(output, nil)

        result, err := mockClient.GetLaunchProfileInitialization(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetLaunchProfileMember", func(t *testing.T) {
        input := &nimble.GetLaunchProfileMemberInput{}
        output := &nimble.GetLaunchProfileMemberOutput{}

        mockClient.On("GetLaunchProfileMember", ctx, input).Return(output, nil)

        result, err := mockClient.GetLaunchProfileMember(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetStreamingImage", func(t *testing.T) {
        input := &nimble.GetStreamingImageInput{}
        output := &nimble.GetStreamingImageOutput{}

        mockClient.On("GetStreamingImage", ctx, input).Return(output, nil)

        result, err := mockClient.GetStreamingImage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetStreamingSession", func(t *testing.T) {
        input := &nimble.GetStreamingSessionInput{}
        output := &nimble.GetStreamingSessionOutput{}

        mockClient.On("GetStreamingSession", ctx, input).Return(output, nil)

        result, err := mockClient.GetStreamingSession(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetStreamingSessionBackup", func(t *testing.T) {
        input := &nimble.GetStreamingSessionBackupInput{}
        output := &nimble.GetStreamingSessionBackupOutput{}

        mockClient.On("GetStreamingSessionBackup", ctx, input).Return(output, nil)

        result, err := mockClient.GetStreamingSessionBackup(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetStreamingSessionStream", func(t *testing.T) {
        input := &nimble.GetStreamingSessionStreamInput{}
        output := &nimble.GetStreamingSessionStreamOutput{}

        mockClient.On("GetStreamingSessionStream", ctx, input).Return(output, nil)

        result, err := mockClient.GetStreamingSessionStream(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetStudio", func(t *testing.T) {
        input := &nimble.GetStudioInput{}
        output := &nimble.GetStudioOutput{}

        mockClient.On("GetStudio", ctx, input).Return(output, nil)

        result, err := mockClient.GetStudio(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetStudioComponent", func(t *testing.T) {
        input := &nimble.GetStudioComponentInput{}
        output := &nimble.GetStudioComponentOutput{}

        mockClient.On("GetStudioComponent", ctx, input).Return(output, nil)

        result, err := mockClient.GetStudioComponent(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetStudioMember", func(t *testing.T) {
        input := &nimble.GetStudioMemberInput{}
        output := &nimble.GetStudioMemberOutput{}

        mockClient.On("GetStudioMember", ctx, input).Return(output, nil)

        result, err := mockClient.GetStudioMember(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListEulaAcceptances", func(t *testing.T) {
        input := &nimble.ListEulaAcceptancesInput{}
        output := &nimble.ListEulaAcceptancesOutput{}

        mockClient.On("ListEulaAcceptances", ctx, input).Return(output, nil)

        result, err := mockClient.ListEulaAcceptances(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListEulas", func(t *testing.T) {
        input := &nimble.ListEulasInput{}
        output := &nimble.ListEulasOutput{}

        mockClient.On("ListEulas", ctx, input).Return(output, nil)

        result, err := mockClient.ListEulas(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListLaunchProfileMembers", func(t *testing.T) {
        input := &nimble.ListLaunchProfileMembersInput{}
        output := &nimble.ListLaunchProfileMembersOutput{}

        mockClient.On("ListLaunchProfileMembers", ctx, input).Return(output, nil)

        result, err := mockClient.ListLaunchProfileMembers(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListLaunchProfiles", func(t *testing.T) {
        input := &nimble.ListLaunchProfilesInput{}
        output := &nimble.ListLaunchProfilesOutput{}

        mockClient.On("ListLaunchProfiles", ctx, input).Return(output, nil)

        result, err := mockClient.ListLaunchProfiles(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListStreamingImages", func(t *testing.T) {
        input := &nimble.ListStreamingImagesInput{}
        output := &nimble.ListStreamingImagesOutput{}

        mockClient.On("ListStreamingImages", ctx, input).Return(output, nil)

        result, err := mockClient.ListStreamingImages(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListStreamingSessionBackups", func(t *testing.T) {
        input := &nimble.ListStreamingSessionBackupsInput{}
        output := &nimble.ListStreamingSessionBackupsOutput{}

        mockClient.On("ListStreamingSessionBackups", ctx, input).Return(output, nil)

        result, err := mockClient.ListStreamingSessionBackups(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListStreamingSessions", func(t *testing.T) {
        input := &nimble.ListStreamingSessionsInput{}
        output := &nimble.ListStreamingSessionsOutput{}

        mockClient.On("ListStreamingSessions", ctx, input).Return(output, nil)

        result, err := mockClient.ListStreamingSessions(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListStudioComponents", func(t *testing.T) {
        input := &nimble.ListStudioComponentsInput{}
        output := &nimble.ListStudioComponentsOutput{}

        mockClient.On("ListStudioComponents", ctx, input).Return(output, nil)

        result, err := mockClient.ListStudioComponents(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListStudioMembers", func(t *testing.T) {
        input := &nimble.ListStudioMembersInput{}
        output := &nimble.ListStudioMembersOutput{}

        mockClient.On("ListStudioMembers", ctx, input).Return(output, nil)

        result, err := mockClient.ListStudioMembers(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListStudios", func(t *testing.T) {
        input := &nimble.ListStudiosInput{}
        output := &nimble.ListStudiosOutput{}

        mockClient.On("ListStudios", ctx, input).Return(output, nil)

        result, err := mockClient.ListStudios(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &nimble.ListTagsForResourceInput{}
        output := &nimble.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutLaunchProfileMembers", func(t *testing.T) {
        input := &nimble.PutLaunchProfileMembersInput{}
        output := &nimble.PutLaunchProfileMembersOutput{}

        mockClient.On("PutLaunchProfileMembers", ctx, input).Return(output, nil)

        result, err := mockClient.PutLaunchProfileMembers(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutStudioMembers", func(t *testing.T) {
        input := &nimble.PutStudioMembersInput{}
        output := &nimble.PutStudioMembersOutput{}

        mockClient.On("PutStudioMembers", ctx, input).Return(output, nil)

        result, err := mockClient.PutStudioMembers(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartStreamingSession", func(t *testing.T) {
        input := &nimble.StartStreamingSessionInput{}
        output := &nimble.StartStreamingSessionOutput{}

        mockClient.On("StartStreamingSession", ctx, input).Return(output, nil)

        result, err := mockClient.StartStreamingSession(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartStudioSSOConfigurationRepair", func(t *testing.T) {
        input := &nimble.StartStudioSSOConfigurationRepairInput{}
        output := &nimble.StartStudioSSOConfigurationRepairOutput{}

        mockClient.On("StartStudioSSOConfigurationRepair", ctx, input).Return(output, nil)

        result, err := mockClient.StartStudioSSOConfigurationRepair(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStopStreamingSession", func(t *testing.T) {
        input := &nimble.StopStreamingSessionInput{}
        output := &nimble.StopStreamingSessionOutput{}

        mockClient.On("StopStreamingSession", ctx, input).Return(output, nil)

        result, err := mockClient.StopStreamingSession(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &nimble.TagResourceInput{}
        output := &nimble.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &nimble.UntagResourceInput{}
        output := &nimble.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateLaunchProfile", func(t *testing.T) {
        input := &nimble.UpdateLaunchProfileInput{}
        output := &nimble.UpdateLaunchProfileOutput{}

        mockClient.On("UpdateLaunchProfile", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateLaunchProfile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateLaunchProfileMember", func(t *testing.T) {
        input := &nimble.UpdateLaunchProfileMemberInput{}
        output := &nimble.UpdateLaunchProfileMemberOutput{}

        mockClient.On("UpdateLaunchProfileMember", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateLaunchProfileMember(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateStreamingImage", func(t *testing.T) {
        input := &nimble.UpdateStreamingImageInput{}
        output := &nimble.UpdateStreamingImageOutput{}

        mockClient.On("UpdateStreamingImage", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateStreamingImage(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateStudio", func(t *testing.T) {
        input := &nimble.UpdateStudioInput{}
        output := &nimble.UpdateStudioOutput{}

        mockClient.On("UpdateStudio", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateStudio(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateStudioComponent", func(t *testing.T) {
        input := &nimble.UpdateStudioComponentInput{}
        output := &nimble.UpdateStudioComponentOutput{}

        mockClient.On("UpdateStudioComponent", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateStudioComponent(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
