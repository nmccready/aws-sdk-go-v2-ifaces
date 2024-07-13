package rolesanywhere_test

// tests for the rolesanywhere service interface for this ../../../service/rolesanywhere/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/rolesanywhere"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/rolesanywhere/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/rolesanywhere/rolesanywhere_iface"
	"github.com/stretchr/testify/assert"
)

func TestRolesanywhereServiceCanBeMocked(t *testing.T) {
	var iface rolesanywhere_iface.IClient
	iface = &rolesanywhere.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := rolesanywhere.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateProfile", func(t *testing.T) {
        input := &rolesanywhere.CreateProfileInput{}
        output := &rolesanywhere.CreateProfileOutput{}

        mockClient.On("CreateProfile", ctx, input).Return(output, nil)

        result, err := mockClient.CreateProfile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateTrustAnchor", func(t *testing.T) {
        input := &rolesanywhere.CreateTrustAnchorInput{}
        output := &rolesanywhere.CreateTrustAnchorOutput{}

        mockClient.On("CreateTrustAnchor", ctx, input).Return(output, nil)

        result, err := mockClient.CreateTrustAnchor(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteAttributeMapping", func(t *testing.T) {
        input := &rolesanywhere.DeleteAttributeMappingInput{}
        output := &rolesanywhere.DeleteAttributeMappingOutput{}

        mockClient.On("DeleteAttributeMapping", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteAttributeMapping(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteCrl", func(t *testing.T) {
        input := &rolesanywhere.DeleteCrlInput{}
        output := &rolesanywhere.DeleteCrlOutput{}

        mockClient.On("DeleteCrl", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteCrl(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteProfile", func(t *testing.T) {
        input := &rolesanywhere.DeleteProfileInput{}
        output := &rolesanywhere.DeleteProfileOutput{}

        mockClient.On("DeleteProfile", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteProfile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteTrustAnchor", func(t *testing.T) {
        input := &rolesanywhere.DeleteTrustAnchorInput{}
        output := &rolesanywhere.DeleteTrustAnchorOutput{}

        mockClient.On("DeleteTrustAnchor", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteTrustAnchor(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisableCrl", func(t *testing.T) {
        input := &rolesanywhere.DisableCrlInput{}
        output := &rolesanywhere.DisableCrlOutput{}

        mockClient.On("DisableCrl", ctx, input).Return(output, nil)

        result, err := mockClient.DisableCrl(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisableProfile", func(t *testing.T) {
        input := &rolesanywhere.DisableProfileInput{}
        output := &rolesanywhere.DisableProfileOutput{}

        mockClient.On("DisableProfile", ctx, input).Return(output, nil)

        result, err := mockClient.DisableProfile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisableTrustAnchor", func(t *testing.T) {
        input := &rolesanywhere.DisableTrustAnchorInput{}
        output := &rolesanywhere.DisableTrustAnchorOutput{}

        mockClient.On("DisableTrustAnchor", ctx, input).Return(output, nil)

        result, err := mockClient.DisableTrustAnchor(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestEnableCrl", func(t *testing.T) {
        input := &rolesanywhere.EnableCrlInput{}
        output := &rolesanywhere.EnableCrlOutput{}

        mockClient.On("EnableCrl", ctx, input).Return(output, nil)

        result, err := mockClient.EnableCrl(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestEnableProfile", func(t *testing.T) {
        input := &rolesanywhere.EnableProfileInput{}
        output := &rolesanywhere.EnableProfileOutput{}

        mockClient.On("EnableProfile", ctx, input).Return(output, nil)

        result, err := mockClient.EnableProfile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestEnableTrustAnchor", func(t *testing.T) {
        input := &rolesanywhere.EnableTrustAnchorInput{}
        output := &rolesanywhere.EnableTrustAnchorOutput{}

        mockClient.On("EnableTrustAnchor", ctx, input).Return(output, nil)

        result, err := mockClient.EnableTrustAnchor(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetCrl", func(t *testing.T) {
        input := &rolesanywhere.GetCrlInput{}
        output := &rolesanywhere.GetCrlOutput{}

        mockClient.On("GetCrl", ctx, input).Return(output, nil)

        result, err := mockClient.GetCrl(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetProfile", func(t *testing.T) {
        input := &rolesanywhere.GetProfileInput{}
        output := &rolesanywhere.GetProfileOutput{}

        mockClient.On("GetProfile", ctx, input).Return(output, nil)

        result, err := mockClient.GetProfile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetSubject", func(t *testing.T) {
        input := &rolesanywhere.GetSubjectInput{}
        output := &rolesanywhere.GetSubjectOutput{}

        mockClient.On("GetSubject", ctx, input).Return(output, nil)

        result, err := mockClient.GetSubject(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetTrustAnchor", func(t *testing.T) {
        input := &rolesanywhere.GetTrustAnchorInput{}
        output := &rolesanywhere.GetTrustAnchorOutput{}

        mockClient.On("GetTrustAnchor", ctx, input).Return(output, nil)

        result, err := mockClient.GetTrustAnchor(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestImportCrl", func(t *testing.T) {
        input := &rolesanywhere.ImportCrlInput{}
        output := &rolesanywhere.ImportCrlOutput{}

        mockClient.On("ImportCrl", ctx, input).Return(output, nil)

        result, err := mockClient.ImportCrl(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListCrls", func(t *testing.T) {
        input := &rolesanywhere.ListCrlsInput{}
        output := &rolesanywhere.ListCrlsOutput{}

        mockClient.On("ListCrls", ctx, input).Return(output, nil)

        result, err := mockClient.ListCrls(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListProfiles", func(t *testing.T) {
        input := &rolesanywhere.ListProfilesInput{}
        output := &rolesanywhere.ListProfilesOutput{}

        mockClient.On("ListProfiles", ctx, input).Return(output, nil)

        result, err := mockClient.ListProfiles(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListSubjects", func(t *testing.T) {
        input := &rolesanywhere.ListSubjectsInput{}
        output := &rolesanywhere.ListSubjectsOutput{}

        mockClient.On("ListSubjects", ctx, input).Return(output, nil)

        result, err := mockClient.ListSubjects(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &rolesanywhere.ListTagsForResourceInput{}
        output := &rolesanywhere.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTrustAnchors", func(t *testing.T) {
        input := &rolesanywhere.ListTrustAnchorsInput{}
        output := &rolesanywhere.ListTrustAnchorsOutput{}

        mockClient.On("ListTrustAnchors", ctx, input).Return(output, nil)

        result, err := mockClient.ListTrustAnchors(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutAttributeMapping", func(t *testing.T) {
        input := &rolesanywhere.PutAttributeMappingInput{}
        output := &rolesanywhere.PutAttributeMappingOutput{}

        mockClient.On("PutAttributeMapping", ctx, input).Return(output, nil)

        result, err := mockClient.PutAttributeMapping(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestPutNotificationSettings", func(t *testing.T) {
        input := &rolesanywhere.PutNotificationSettingsInput{}
        output := &rolesanywhere.PutNotificationSettingsOutput{}

        mockClient.On("PutNotificationSettings", ctx, input).Return(output, nil)

        result, err := mockClient.PutNotificationSettings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestResetNotificationSettings", func(t *testing.T) {
        input := &rolesanywhere.ResetNotificationSettingsInput{}
        output := &rolesanywhere.ResetNotificationSettingsOutput{}

        mockClient.On("ResetNotificationSettings", ctx, input).Return(output, nil)

        result, err := mockClient.ResetNotificationSettings(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &rolesanywhere.TagResourceInput{}
        output := &rolesanywhere.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &rolesanywhere.UntagResourceInput{}
        output := &rolesanywhere.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateCrl", func(t *testing.T) {
        input := &rolesanywhere.UpdateCrlInput{}
        output := &rolesanywhere.UpdateCrlOutput{}

        mockClient.On("UpdateCrl", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateCrl(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateProfile", func(t *testing.T) {
        input := &rolesanywhere.UpdateProfileInput{}
        output := &rolesanywhere.UpdateProfileOutput{}

        mockClient.On("UpdateProfile", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateProfile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateTrustAnchor", func(t *testing.T) {
        input := &rolesanywhere.UpdateTrustAnchorInput{}
        output := &rolesanywhere.UpdateTrustAnchorOutput{}

        mockClient.On("UpdateTrustAnchor", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateTrustAnchor(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
