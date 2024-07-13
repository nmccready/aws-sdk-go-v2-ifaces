package route53profiles_test

// tests for the route53profiles service interface for this ../../../service/route53profiles/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/route53profiles"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/route53profiles/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/route53profiles/route53profiles_iface"
	"github.com/stretchr/testify/assert"
)

func TestRoute53profilesServiceCanBeMocked(t *testing.T) {
	var iface route53profiles_iface.IClient
	iface = &route53profiles.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := route53profiles.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateProfile", func(t *testing.T) {
        input := &route53profiles.AssociateProfileInput{}
        output := &route53profiles.AssociateProfileOutput{}

        mockClient.On("AssociateProfile", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateProfile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestAssociateResourceToProfile", func(t *testing.T) {
        input := &route53profiles.AssociateResourceToProfileInput{}
        output := &route53profiles.AssociateResourceToProfileOutput{}

        mockClient.On("AssociateResourceToProfile", ctx, input).Return(output, nil)

        result, err := mockClient.AssociateResourceToProfile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateProfile", func(t *testing.T) {
        input := &route53profiles.CreateProfileInput{}
        output := &route53profiles.CreateProfileOutput{}

        mockClient.On("CreateProfile", ctx, input).Return(output, nil)

        result, err := mockClient.CreateProfile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteProfile", func(t *testing.T) {
        input := &route53profiles.DeleteProfileInput{}
        output := &route53profiles.DeleteProfileOutput{}

        mockClient.On("DeleteProfile", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteProfile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateProfile", func(t *testing.T) {
        input := &route53profiles.DisassociateProfileInput{}
        output := &route53profiles.DisassociateProfileOutput{}

        mockClient.On("DisassociateProfile", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateProfile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDisassociateResourceFromProfile", func(t *testing.T) {
        input := &route53profiles.DisassociateResourceFromProfileInput{}
        output := &route53profiles.DisassociateResourceFromProfileOutput{}

        mockClient.On("DisassociateResourceFromProfile", ctx, input).Return(output, nil)

        result, err := mockClient.DisassociateResourceFromProfile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetProfile", func(t *testing.T) {
        input := &route53profiles.GetProfileInput{}
        output := &route53profiles.GetProfileOutput{}

        mockClient.On("GetProfile", ctx, input).Return(output, nil)

        result, err := mockClient.GetProfile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetProfileAssociation", func(t *testing.T) {
        input := &route53profiles.GetProfileAssociationInput{}
        output := &route53profiles.GetProfileAssociationOutput{}

        mockClient.On("GetProfileAssociation", ctx, input).Return(output, nil)

        result, err := mockClient.GetProfileAssociation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetProfileResourceAssociation", func(t *testing.T) {
        input := &route53profiles.GetProfileResourceAssociationInput{}
        output := &route53profiles.GetProfileResourceAssociationOutput{}

        mockClient.On("GetProfileResourceAssociation", ctx, input).Return(output, nil)

        result, err := mockClient.GetProfileResourceAssociation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListProfileAssociations", func(t *testing.T) {
        input := &route53profiles.ListProfileAssociationsInput{}
        output := &route53profiles.ListProfileAssociationsOutput{}

        mockClient.On("ListProfileAssociations", ctx, input).Return(output, nil)

        result, err := mockClient.ListProfileAssociations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListProfileResourceAssociations", func(t *testing.T) {
        input := &route53profiles.ListProfileResourceAssociationsInput{}
        output := &route53profiles.ListProfileResourceAssociationsOutput{}

        mockClient.On("ListProfileResourceAssociations", ctx, input).Return(output, nil)

        result, err := mockClient.ListProfileResourceAssociations(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListProfiles", func(t *testing.T) {
        input := &route53profiles.ListProfilesInput{}
        output := &route53profiles.ListProfilesOutput{}

        mockClient.On("ListProfiles", ctx, input).Return(output, nil)

        result, err := mockClient.ListProfiles(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &route53profiles.ListTagsForResourceInput{}
        output := &route53profiles.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &route53profiles.TagResourceInput{}
        output := &route53profiles.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &route53profiles.UntagResourceInput{}
        output := &route53profiles.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateProfileResourceAssociation", func(t *testing.T) {
        input := &route53profiles.UpdateProfileResourceAssociationInput{}
        output := &route53profiles.UpdateProfileResourceAssociationOutput{}

        mockClient.On("UpdateProfileResourceAssociation", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateProfileResourceAssociation(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
