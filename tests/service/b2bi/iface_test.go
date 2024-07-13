package b2bi_test

// tests for the b2bi service interface for this ../../../service/b2bi/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/b2bi"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/b2bi/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/b2bi/b2bi_iface"
	"github.com/stretchr/testify/assert"
)

func TestB2biServiceCanBeMocked(t *testing.T) {
	var iface b2bi_iface.IClient
	iface = &b2bi.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := b2bi.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateCapability", func(t *testing.T) {
        input := &b2bi.CreateCapabilityInput{}
        output := &b2bi.CreateCapabilityOutput{}

        mockClient.On("CreateCapability", ctx, input).Return(output, nil)

        result, err := mockClient.CreateCapability(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreatePartnership", func(t *testing.T) {
        input := &b2bi.CreatePartnershipInput{}
        output := &b2bi.CreatePartnershipOutput{}

        mockClient.On("CreatePartnership", ctx, input).Return(output, nil)

        result, err := mockClient.CreatePartnership(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateProfile", func(t *testing.T) {
        input := &b2bi.CreateProfileInput{}
        output := &b2bi.CreateProfileOutput{}

        mockClient.On("CreateProfile", ctx, input).Return(output, nil)

        result, err := mockClient.CreateProfile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestCreateTransformer", func(t *testing.T) {
        input := &b2bi.CreateTransformerInput{}
        output := &b2bi.CreateTransformerOutput{}

        mockClient.On("CreateTransformer", ctx, input).Return(output, nil)

        result, err := mockClient.CreateTransformer(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteCapability", func(t *testing.T) {
        input := &b2bi.DeleteCapabilityInput{}
        output := &b2bi.DeleteCapabilityOutput{}

        mockClient.On("DeleteCapability", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteCapability(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeletePartnership", func(t *testing.T) {
        input := &b2bi.DeletePartnershipInput{}
        output := &b2bi.DeletePartnershipOutput{}

        mockClient.On("DeletePartnership", ctx, input).Return(output, nil)

        result, err := mockClient.DeletePartnership(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteProfile", func(t *testing.T) {
        input := &b2bi.DeleteProfileInput{}
        output := &b2bi.DeleteProfileOutput{}

        mockClient.On("DeleteProfile", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteProfile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestDeleteTransformer", func(t *testing.T) {
        input := &b2bi.DeleteTransformerInput{}
        output := &b2bi.DeleteTransformerOutput{}

        mockClient.On("DeleteTransformer", ctx, input).Return(output, nil)

        result, err := mockClient.DeleteTransformer(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetCapability", func(t *testing.T) {
        input := &b2bi.GetCapabilityInput{}
        output := &b2bi.GetCapabilityOutput{}

        mockClient.On("GetCapability", ctx, input).Return(output, nil)

        result, err := mockClient.GetCapability(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetPartnership", func(t *testing.T) {
        input := &b2bi.GetPartnershipInput{}
        output := &b2bi.GetPartnershipOutput{}

        mockClient.On("GetPartnership", ctx, input).Return(output, nil)

        result, err := mockClient.GetPartnership(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetProfile", func(t *testing.T) {
        input := &b2bi.GetProfileInput{}
        output := &b2bi.GetProfileOutput{}

        mockClient.On("GetProfile", ctx, input).Return(output, nil)

        result, err := mockClient.GetProfile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetTransformer", func(t *testing.T) {
        input := &b2bi.GetTransformerInput{}
        output := &b2bi.GetTransformerOutput{}

        mockClient.On("GetTransformer", ctx, input).Return(output, nil)

        result, err := mockClient.GetTransformer(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestGetTransformerJob", func(t *testing.T) {
        input := &b2bi.GetTransformerJobInput{}
        output := &b2bi.GetTransformerJobOutput{}

        mockClient.On("GetTransformerJob", ctx, input).Return(output, nil)

        result, err := mockClient.GetTransformerJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListCapabilities", func(t *testing.T) {
        input := &b2bi.ListCapabilitiesInput{}
        output := &b2bi.ListCapabilitiesOutput{}

        mockClient.On("ListCapabilities", ctx, input).Return(output, nil)

        result, err := mockClient.ListCapabilities(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListPartnerships", func(t *testing.T) {
        input := &b2bi.ListPartnershipsInput{}
        output := &b2bi.ListPartnershipsOutput{}

        mockClient.On("ListPartnerships", ctx, input).Return(output, nil)

        result, err := mockClient.ListPartnerships(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListProfiles", func(t *testing.T) {
        input := &b2bi.ListProfilesInput{}
        output := &b2bi.ListProfilesOutput{}

        mockClient.On("ListProfiles", ctx, input).Return(output, nil)

        result, err := mockClient.ListProfiles(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTagsForResource", func(t *testing.T) {
        input := &b2bi.ListTagsForResourceInput{}
        output := &b2bi.ListTagsForResourceOutput{}

        mockClient.On("ListTagsForResource", ctx, input).Return(output, nil)

        result, err := mockClient.ListTagsForResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestListTransformers", func(t *testing.T) {
        input := &b2bi.ListTransformersInput{}
        output := &b2bi.ListTransformersOutput{}

        mockClient.On("ListTransformers", ctx, input).Return(output, nil)

        result, err := mockClient.ListTransformers(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestStartTransformerJob", func(t *testing.T) {
        input := &b2bi.StartTransformerJobInput{}
        output := &b2bi.StartTransformerJobOutput{}

        mockClient.On("StartTransformerJob", ctx, input).Return(output, nil)

        result, err := mockClient.StartTransformerJob(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTagResource", func(t *testing.T) {
        input := &b2bi.TagResourceInput{}
        output := &b2bi.TagResourceOutput{}

        mockClient.On("TagResource", ctx, input).Return(output, nil)

        result, err := mockClient.TagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTestMapping", func(t *testing.T) {
        input := &b2bi.TestMappingInput{}
        output := &b2bi.TestMappingOutput{}

        mockClient.On("TestMapping", ctx, input).Return(output, nil)

        result, err := mockClient.TestMapping(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestTestParsing", func(t *testing.T) {
        input := &b2bi.TestParsingInput{}
        output := &b2bi.TestParsingOutput{}

        mockClient.On("TestParsing", ctx, input).Return(output, nil)

        result, err := mockClient.TestParsing(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUntagResource", func(t *testing.T) {
        input := &b2bi.UntagResourceInput{}
        output := &b2bi.UntagResourceOutput{}

        mockClient.On("UntagResource", ctx, input).Return(output, nil)

        result, err := mockClient.UntagResource(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateCapability", func(t *testing.T) {
        input := &b2bi.UpdateCapabilityInput{}
        output := &b2bi.UpdateCapabilityOutput{}

        mockClient.On("UpdateCapability", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateCapability(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdatePartnership", func(t *testing.T) {
        input := &b2bi.UpdatePartnershipInput{}
        output := &b2bi.UpdatePartnershipOutput{}

        mockClient.On("UpdatePartnership", ctx, input).Return(output, nil)

        result, err := mockClient.UpdatePartnership(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateProfile", func(t *testing.T) {
        input := &b2bi.UpdateProfileInput{}
        output := &b2bi.UpdateProfileOutput{}

        mockClient.On("UpdateProfile", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateProfile(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestUpdateTransformer", func(t *testing.T) {
        input := &b2bi.UpdateTransformerInput{}
        output := &b2bi.UpdateTransformerOutput{}

        mockClient.On("UpdateTransformer", ctx, input).Return(output, nil)

        result, err := mockClient.UpdateTransformer(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
